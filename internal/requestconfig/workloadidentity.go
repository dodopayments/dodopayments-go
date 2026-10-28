// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package requestconfig

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apierror"
)

const (
	workloadIdentityTokenPath     = "/workload-identity/token"
	workloadIdentityProviderIDEnv = "DODO_PAYMENTS_IDENTITY_PROVIDER_ID"
	workloadIdentityAPIKeyEnv     = "DODO_PAYMENTS_API_KEY"
	workloadIdentityRefreshWindow = 5 * time.Minute
)

// SubjectTokenProvider supplies the subject token (a JWT) that the client
// exchanges for an access token. The client calls it before each exchange. The
// audience is the origin of the client base URL, for example
// "https://live.dodopayments.com".
type SubjectTokenProvider interface {
	SubjectToken(ctx context.Context, audience string) (string, error)
}

// SubjectTokenProviderFunc is a function that implements [SubjectTokenProvider].
type SubjectTokenProviderFunc func(ctx context.Context, audience string) (string, error)

func (f SubjectTokenProviderFunc) SubjectToken(ctx context.Context, audience string) (string, error) {
	return f(ctx, audience)
}

// WorkloadIdentityState holds the workload identity configuration and the
// cached access token. It is safe for concurrent use.
type WorkloadIdentityState struct {
	identityProviderID string
	provider           SubjectTokenProvider
	// explicit is false when the client reads the configuration from the environment.
	explicit bool

	mu          sync.Mutex
	accessToken string
	expiresAt   time.Time
}

// NewWorkloadIdentityState returns the state for one client. The identity
// provider ID must not be empty.
func NewWorkloadIdentityState(identityProviderID string, provider SubjectTokenProvider, explicit bool) *WorkloadIdentityState {
	return &WorkloadIdentityState{identityProviderID: identityProviderID, provider: provider, explicit: explicit}
}

// workloadIdentityError stops the retry loop in [RequestConfig.Execute].
type workloadIdentityError struct{ err error }

func (e *workloadIdentityError) Error() string { return e.err.Error() }
func (e *workloadIdentityError) Unwrap() error { return e.err }

// WithWorkloadIdentityFromEnvironment marks the bearer token that the
// environment supplies, and uses workload identity when the environment has an
// identity provider ID, no API key, and a subject token source.
func WithWorkloadIdentityFromEnvironment() RequestOption {
	var state *WorkloadIdentityState
	if _, ok := os.LookupEnv(workloadIdentityAPIKeyEnv); !ok {
		if id := os.Getenv(workloadIdentityProviderIDEnv); id != "" {
			if provider, found := DetectSubjectTokenProvider(); found {
				state = NewWorkloadIdentityState(id, provider, false)
			}
		}
	}
	return RequestOptionFunc(func(r *RequestConfig) error {
		r.ExplicitBearerToken = false
		if state != nil {
			r.WorkloadIdentity = state
		}
		return nil
	})
}

func (cfg *RequestConfig) withWorkloadIdentity(next middlewareNext) (middlewareNext, error) {
	state := cfg.WorkloadIdentity
	if state == nil {
		return next, nil
	}
	if cfg.ExplicitBearerToken {
		if state.explicit {
			return nil, fmt.Errorf("requestconfig: a bearer token and workload identity cannot be used together")
		}
		return next, nil
	}
	return func(req *http.Request) (*http.Response, error) {
		return state.roundTrip(cfg, next, req)
	}, nil
}

func (s *WorkloadIdentityState) roundTrip(cfg *RequestConfig, next middlewareNext, req *http.Request) (*http.Response, error) {
	token, err := s.token(cfg, next, req.Context())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := next(req)
	if err != nil || res.StatusCode != http.StatusUnauthorized {
		return res, err
	}

	// Exchange one time again and send the request one time again.
	retry := req.Clone(req.Context())
	if req.Body != nil && req.Body != http.NoBody {
		if req.GetBody == nil {
			return res, nil
		}
		if retry.Body, err = req.GetBody(); err != nil {
			return res, nil
		}
	}
	_ = res.Body.Close()
	s.invalidate(token)
	token, err = s.token(cfg, next, req.Context())
	if err != nil {
		return nil, err
	}
	retry.Header.Set("Authorization", "Bearer "+token)
	return next(retry)
}

func (s *WorkloadIdentityState) invalidate(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.accessToken == token {
		s.accessToken, s.expiresAt = "", time.Time{}
	}
}

// token returns the cached access token, or exchanges a new subject token. The
// mutex makes sure that concurrent requests do one exchange only.
func (s *WorkloadIdentityState) token(cfg *RequestConfig, next middlewareNext, ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.accessToken != "" && time.Until(s.expiresAt) > workloadIdentityRefreshWindow {
		return s.accessToken, nil
	}
	s.accessToken, s.expiresAt = "", time.Time{}

	subjectToken, err := s.provider.SubjectToken(ctx, cfg.BaseURL.Scheme+"://"+cfg.BaseURL.Host)
	if err != nil {
		return "", &workloadIdentityError{fmt.Errorf("requestconfig: cannot get the workload identity subject token: %w", err)}
	}
	if subjectToken == "" {
		return "", &workloadIdentityError{errors.New("requestconfig: the workload identity subject token is empty")}
	}

	body, err := json.Marshal(struct {
		GrantType          string `json:"grant_type"`
		SubjectTokenType   string `json:"subject_token_type"`
		SubjectToken       string `json:"subject_token"`
		IdentityProviderID string `json:"identity_provider_id"`
	}{
		GrantType:          "urn:ietf:params:oauth:grant-type:token-exchange",
		SubjectTokenType:   "urn:ietf:params:oauth:token-type:jwt",
		SubjectToken:       subjectToken,
		IdentityProviderID: s.identityProviderID,
	})
	if err != nil {
		return "", &workloadIdentityError{err}
	}
	u, err := cfg.BaseURL.Parse(strings.TrimLeft(workloadIdentityTokenPath, "/"))
	if err != nil {
		return "", &workloadIdentityError{err}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bytes.NewReader(body))
	if err != nil {
		return "", &workloadIdentityError{err}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	for k, v := range getDefaultHeaders() {
		req.Header.Set(k, v)
	}

	// A transport error or a 5xx status goes back to the retry loop of the request.
	res, err := next(req)
	if err != nil {
		return "", fmt.Errorf("requestconfig: workload identity token exchange failed: %w", err)
	}
	contents, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return "", fmt.Errorf("requestconfig: cannot read the workload identity token response: %w", err)
	}
	if res.StatusCode >= 400 {
		res.Body = io.NopCloser(bytes.NewReader(contents))
		apiErr := &apierror.Error{Request: req, Response: res, StatusCode: res.StatusCode}
		if err := apiErr.UnmarshalJSON(contents); err != nil {
			return "", &workloadIdentityError{err}
		}
		if res.StatusCode >= 500 {
			return "", apiErr
		}
		return "", &workloadIdentityError{apiErr}
	}

	var out struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		ExpiresAt   string `json:"expires_at"`
	}
	if err := json.Unmarshal(contents, &out); err != nil {
		return "", &workloadIdentityError{fmt.Errorf("requestconfig: cannot parse the workload identity token response: %w", err)}
	}
	if out.AccessToken == "" {
		return "", &workloadIdentityError{errors.New("requestconfig: the workload identity token response has no access_token")}
	}
	expiresAt := time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	if t, err := time.Parse(time.RFC3339, out.ExpiresAt); err == nil {
		expiresAt = t
	}
	s.accessToken, s.expiresAt = out.AccessToken, expiresAt
	return s.accessToken, nil
}

// RevokeWorkloadIdentityToken revokes the cached access token and clears the
// cache. It does nothing when no access token is cached.
func (cfg *RequestConfig) RevokeWorkloadIdentityToken() error {
	state := cfg.WorkloadIdentity
	if state == nil {
		return nil
	}
	state.mu.Lock()
	defer state.mu.Unlock()
	if state.accessToken == "" {
		return nil
	}
	cfg.WorkloadIdentity = nil
	cfg.Request.Header.Set("Authorization", "Bearer "+state.accessToken)
	err := cfg.Execute()
	var apiErr *apierror.Error
	if err == nil || errors.As(err, &apiErr) {
		state.accessToken, state.expiresAt = "", time.Time{}
	}
	return err
}

// DetectSubjectTokenProvider finds a subject token source in the environment.
// It examines GitHub Actions, Terraform Cloud, the
// DODO_PAYMENTS_IDENTITY_TOKEN_FILE variable, and the Kubernetes token path,
// in that order.
func DetectSubjectTokenProvider() (SubjectTokenProvider, bool) {
	if os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL") != "" && os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN") != "" {
		return GitHubActionsSubjectTokenProvider(""), true
	}
	if os.Getenv("TFC_WORKLOAD_IDENTITY_TOKEN_DODO") != "" || os.Getenv("TFC_WORKLOAD_IDENTITY_TOKEN") != "" {
		return TerraformCloudSubjectTokenProvider(), true
	}
	if os.Getenv("DODO_PAYMENTS_IDENTITY_TOKEN_FILE") != "" {
		return FileSubjectTokenProvider(""), true
	}
	if _, err := os.Stat(DefaultKubernetesTokenPath); err == nil {
		return FileSubjectTokenProvider(DefaultKubernetesTokenPath), true
	}
	return nil, false
}

// AutoSubjectTokenProvider detects a subject token source before each exchange.
func AutoSubjectTokenProvider() SubjectTokenProvider {
	return SubjectTokenProviderFunc(func(ctx context.Context, audience string) (string, error) {
		provider, ok := DetectSubjectTokenProvider()
		if !ok {
			return "", errors.New("no subject token source found in the environment")
		}
		return provider.SubjectToken(ctx, audience)
	})
}

// DefaultKubernetesTokenPath is the default path of a projected service account token.
const DefaultKubernetesTokenPath = "/var/run/secrets/tokens/dodo-token"

// GitHubActionsSubjectTokenProvider requests an OIDC token from GitHub Actions.
// An empty audience uses the origin of the client base URL.
func GitHubActionsSubjectTokenProvider(audience string) SubjectTokenProvider {
	return SubjectTokenProviderFunc(func(ctx context.Context, defaultAudience string) (string, error) {
		requestURL, requestToken := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL"), os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
		if requestURL == "" || requestToken == "" {
			return "", errors.New("ACTIONS_ID_TOKEN_REQUEST_URL and ACTIONS_ID_TOKEN_REQUEST_TOKEN must be set; give the job the id-token: write permission")
		}
		if audience != "" {
			defaultAudience = audience
		}
		u, err := url.Parse(requestURL)
		if err != nil {
			return "", err
		}
		query := u.Query()
		query.Set("audience", defaultAudience)
		u.RawQuery = query.Encode()
		body, err := getSubjectToken(ctx, u.String(), http.Header{"Authorization": {"bearer " + requestToken}})
		if err != nil {
			return "", err
		}
		var out struct {
			Value string `json:"value"`
		}
		if err := json.Unmarshal(body, &out); err != nil {
			return "", fmt.Errorf("cannot parse the GitHub Actions OIDC response: %w", err)
		}
		return out.Value, nil
	})
}

// GCPSubjectTokenProvider requests an identity token from the GCP metadata
// server. An empty audience uses the origin of the client base URL.
func GCPSubjectTokenProvider(audience string) SubjectTokenProvider {
	return SubjectTokenProviderFunc(func(ctx context.Context, defaultAudience string) (string, error) {
		if audience != "" {
			defaultAudience = audience
		}
		u := "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity?" +
			url.Values{"audience": {defaultAudience}, "format": {"full"}}.Encode()
		body, err := getSubjectToken(ctx, u, http.Header{"Metadata-Flavor": {"Google"}})
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(body)), nil
	})
}

func getSubjectToken(ctx context.Context, u string, header http.Header) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header = header
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the subject token request returned status %d", res.StatusCode)
	}
	return body, nil
}

// EnvSubjectTokenProvider reads the subject token from an environment variable.
func EnvSubjectTokenProvider(names ...string) SubjectTokenProvider {
	return SubjectTokenProviderFunc(func(context.Context, string) (string, error) {
		for _, name := range names {
			if v := strings.TrimSpace(os.Getenv(name)); v != "" {
				return v, nil
			}
		}
		return "", fmt.Errorf("the environment variable %s is not set", strings.Join(names, " or "))
	})
}

// TerraformCloudSubjectTokenProvider reads the Terraform Cloud workload
// identity token.
func TerraformCloudSubjectTokenProvider() SubjectTokenProvider {
	return EnvSubjectTokenProvider("TFC_WORKLOAD_IDENTITY_TOKEN_DODO", "TFC_WORKLOAD_IDENTITY_TOKEN")
}

// FileSubjectTokenProvider reads the subject token from a file before each
// exchange. An empty path uses the DODO_PAYMENTS_IDENTITY_TOKEN_FILE variable.
func FileSubjectTokenProvider(path string) SubjectTokenProvider {
	return SubjectTokenProviderFunc(func(context.Context, string) (string, error) {
		p := path
		if p == "" {
			p = os.Getenv("DODO_PAYMENTS_IDENTITY_TOKEN_FILE")
		}
		if p == "" {
			return "", errors.New("no token file path; set DODO_PAYMENTS_IDENTITY_TOKEN_FILE")
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(b)), nil
	})
}
