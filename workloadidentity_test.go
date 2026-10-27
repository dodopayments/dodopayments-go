// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/workloadidentity"
)

const (
	wifTokenPath  = "/workload-identity/token"
	wifRevokePath = "/workload-identity/token/revoke"
	wifSubject    = "subject-token-for-tests"
	wifProviderID = "wip_test"
)

// wifServer is a mock API server with the workload identity token endpoint.
type wifServer struct {
	*httptest.Server
	exchanges     atomic.Int32
	apiCalls      atomic.Int32
	expiresIn     int
	exchangeDelay time.Duration
	exchangeFail  bool
	apiStatus     func(call int32) int

	mu          sync.Mutex
	exchangeErr string
	apiAuth     []string
	revokeAuth  []string
}

func newWIFServer(t *testing.T) *wifServer {
	t.Helper()
	for _, name := range []string{
		"DODO_PAYMENTS_API_KEY", "DODO_PAYMENTS_IDENTITY_PROVIDER_ID", "DODO_PAYMENTS_BASE_URL", "DODO_PAYMENTS_IDENTITY_TOKEN_FILE",
		"ACTIONS_ID_TOKEN_REQUEST_URL", "ACTIONS_ID_TOKEN_REQUEST_TOKEN", "TFC_WORKLOAD_IDENTITY_TOKEN", "TFC_WORKLOAD_IDENTITY_TOKEN_DODO",
	} {
		t.Setenv(name, "")
		_ = os.Unsetenv(name)
	}
	s := &wifServer{expiresIn: 3600}
	s.Server = httptest.NewServer(http.HandlerFunc(s.handle))
	t.Cleanup(s.Close)
	return s
}

func (s *wifServer) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.URL.Path {
	case wifTokenPath:
		n := s.exchanges.Add(1)
		time.Sleep(s.exchangeDelay)
		if s.exchangeFail {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"error":"invalid_grant","error_description":"The subject token does not match the mapping"}`))
			return
		}
		var body map[string]any
		err := json.NewDecoder(r.Body).Decode(&body)
		want := map[string]any{
			"grant_type":           "urn:ietf:params:oauth:grant-type:token-exchange",
			"subject_token_type":   "urn:ietf:params:oauth:token-type:jwt",
			"subject_token":        wifSubject,
			"identity_provider_id": wifProviderID,
		}
		s.mu.Lock()
		switch {
		case err != nil || !reflect.DeepEqual(body, want):
			s.exchangeErr = "the exchange body is not the expected JSON"
		case r.Header.Get("Authorization") != "":
			s.exchangeErr = "the exchange request has an Authorization header"
		case r.Method != http.MethodPost:
			s.exchangeErr = "the exchange request is not a POST"
		}
		s.mu.Unlock()
		_, _ = fmt.Fprintf(w, `{"access_token":"access-%d","issued_token_type":"urn:ietf:params:oauth:token-type:access_token","token_type":"Bearer","expires_in":%d,"scope":"payments:read"}`, n, s.expiresIn)
	case wifRevokePath:
		s.mu.Lock()
		s.revokeAuth = append(s.revokeAuth, r.Method+" "+r.Header.Get("Authorization"))
		s.mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	default:
		n := s.apiCalls.Add(1)
		s.mu.Lock()
		s.apiAuth = append(s.apiAuth, r.Header.Get("Authorization"))
		s.mu.Unlock()
		status := http.StatusOK
		if s.apiStatus != nil {
			status = s.apiStatus(n)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(`{}`))
	}
}

func (s *wifServer) client(opts ...option.RequestOption) *dodopayments.Client {
	return dodopayments.NewClient(append([]option.RequestOption{
		option.WithBaseURL(s.URL),
		option.WithWorkloadIdentity(option.WorkloadIdentity{
			IdentityProviderID: wifProviderID,
			Provider:           workloadidentity.Static(wifSubject),
		}),
	}, opts...)...)
}

func (s *wifServer) headers(list *[]string) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), *list...)
}

func (s *wifServer) checkExchanges(t *testing.T) {
	t.Helper()
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.exchangeErr != "" {
		t.Error(s.exchangeErr)
	}
}

func wifCall(client *dodopayments.Client) error {
	var res map[string]any
	return client.Get(context.Background(), "ping", nil, &res)
}

func TestWorkloadIdentityExchangeThenBearer(t *testing.T) {
	server := newWIFServer(t)
	if err := wifCall(server.client()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	server.checkExchanges(t)
	if got := server.exchanges.Load(); got != 1 {
		t.Errorf("expected 1 exchange, got %d", got)
	}
	if !reflect.DeepEqual(server.headers(&server.apiAuth), []string{"Bearer access-1"}) {
		t.Error("the API request does not have the access token as a bearer token")
	}
}

func TestWorkloadIdentityConcurrentRequestsExchangeOnce(t *testing.T) {
	server := newWIFServer(t)
	server.exchangeDelay = 100 * time.Millisecond
	client := server.client()
	var wg sync.WaitGroup
	errs := make([]error, 2)
	for i := range errs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs[i] = wifCall(client)
		}()
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if got := server.exchanges.Load(); got != 1 {
		t.Errorf("expected 1 exchange, got %d", got)
	}
	if got := server.apiCalls.Load(); got != 2 {
		t.Errorf("expected 2 API calls, got %d", got)
	}
}

func TestWorkloadIdentityRefreshWindow(t *testing.T) {
	for _, tc := range []struct {
		expiresIn int
		exchanges int32
	}{{expiresIn: 200, exchanges: 2}, {expiresIn: 3600, exchanges: 1}} {
		t.Run(fmt.Sprintf("expires_in=%d", tc.expiresIn), func(t *testing.T) {
			server := newWIFServer(t)
			server.expiresIn = tc.expiresIn
			client := server.client()
			for range 2 {
				if err := wifCall(client); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			if got := server.exchanges.Load(); got != tc.exchanges {
				t.Errorf("expected %d exchanges, got %d", tc.exchanges, got)
			}
		})
	}
}

func TestWorkloadIdentityUnauthorizedReexchangesOnce(t *testing.T) {
	server := newWIFServer(t)
	server.apiStatus = func(int32) int { return http.StatusUnauthorized }
	err := wifCall(server.client())
	var apiErr *dodopayments.Error
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected a 401 API error, got %v", err)
	}
	if got := server.exchanges.Load(); got != 2 {
		t.Errorf("expected 2 exchanges, got %d", got)
	}
	if !reflect.DeepEqual(server.headers(&server.apiAuth), []string{"Bearer access-1", "Bearer access-2"}) {
		t.Error("the retry does not use the new access token")
	}
}

func TestWorkloadIdentityWithBearerTokenIsAnError(t *testing.T) {
	server := newWIFServer(t)
	err := wifCall(server.client(option.WithBearerToken("My Bearer Token")))
	if err == nil || !strings.Contains(err.Error(), "cannot be used together") {
		t.Fatalf("expected a configuration error, got %v", err)
	}
	if server.exchanges.Load() != 0 || server.apiCalls.Load() != 0 {
		t.Error("expected no requests")
	}
}

func TestWorkloadIdentityExchangeForbidden(t *testing.T) {
	server := newWIFServer(t)
	server.exchangeFail = true
	err := wifCall(server.client())
	var apiErr *dodopayments.Error
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusForbidden {
		t.Fatalf("expected a 403 API error, got %v", err)
	}
	if !strings.Contains(err.Error(), "invalid_grant") || !strings.Contains(err.Error(), "does not match the mapping") {
		t.Errorf("expected the error to contain the RFC 6749 error and description, got %v", err)
	}
	if got := server.exchanges.Load(); got != 1 {
		t.Errorf("expected 1 exchange, got %d", got)
	}
	if got := server.apiCalls.Load(); got != 0 {
		t.Errorf("expected no API calls, got %d", got)
	}
}

func TestWorkloadIdentityRevoke(t *testing.T) {
	server := newWIFServer(t)
	client := server.client()
	if err := client.WorkloadIdentity.Revoke(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(server.headers(&server.revokeAuth)) != 0 {
		t.Fatal("expected no revoke request without a cached token")
	}
	if err := wifCall(client); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := client.WorkloadIdentity.Revoke(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(server.headers(&server.revokeAuth), []string{"POST Bearer access-1"}) {
		t.Error("the revoke request does not have the cached access token as a bearer token")
	}
	if err := wifCall(client); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := server.exchanges.Load(); got != 2 {
		t.Errorf("expected a new exchange after revoke, got %d exchanges", got)
	}
}

func TestWorkloadIdentityFromEnvironment(t *testing.T) {
	server := newWIFServer(t)
	tokenFile := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFile, []byte(" "+wifSubject+"  "), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DODO_PAYMENTS_IDENTITY_PROVIDER_ID", wifProviderID)
	t.Setenv("DODO_PAYMENTS_IDENTITY_TOKEN_FILE", tokenFile)
	if err := wifCall(dodopayments.NewClient(option.WithBaseURL(server.URL))); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	server.checkExchanges(t)
	if !reflect.DeepEqual(server.headers(&server.apiAuth), []string{"Bearer access-1"}) {
		t.Error("the API request does not have the access token as a bearer token")
	}
}

func TestWorkloadIdentityOverridesAPIKeyFromEnvironment(t *testing.T) {
	server := newWIFServer(t)
	t.Setenv("DODO_PAYMENTS_API_KEY", "api-key-from-environment")
	if err := wifCall(server.client()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(server.headers(&server.apiAuth), []string{"Bearer access-1"}) {
		t.Error("the API request does not have the access token as a bearer token")
	}
}
