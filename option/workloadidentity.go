// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package option

import (
	"fmt"
	"os"

	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
)

// SubjectTokenProvider supplies the subject token (a JWT) that the client
// exchanges for an access token. The client calls it before each exchange. The
// audience is the origin of the client base URL, for example
// "https://live.dodopayments.com". Implementations must be safe for concurrent
// use.
type SubjectTokenProvider = requestconfig.SubjectTokenProvider

// SubjectTokenProviderFunc is a function that implements [SubjectTokenProvider].
type SubjectTokenProviderFunc = requestconfig.SubjectTokenProviderFunc

// WorkloadIdentity configures [WithWorkloadIdentity].
type WorkloadIdentity struct {
	// IdentityProviderID is the ID of the workload identity provider (wip_...).
	// If it is empty, the client reads DODO_PAYMENTS_IDENTITY_PROVIDER_ID.
	IdentityProviderID string
	// Provider supplies the subject token. If it is nil, the client detects a
	// source in the environment. See the workloadidentity package.
	Provider SubjectTokenProvider
}

// WithWorkloadIdentity returns a RequestOption that authenticates the client with
// workload identity federation. The client exchanges the subject token for a
// short-lived access token, caches the access token, and gets a new one when 5
// minutes or less remain. Do not use it together with [WithBearerToken].
func WithWorkloadIdentity(config WorkloadIdentity) RequestOption {
	id := config.IdentityProviderID
	if id == "" {
		id = os.Getenv("DODO_PAYMENTS_IDENTITY_PROVIDER_ID")
	}
	provider := config.Provider
	if provider == nil {
		provider = requestconfig.AutoSubjectTokenProvider()
	}
	state := requestconfig.NewWorkloadIdentityState(id, provider, true)
	return requestconfig.RequestOptionFunc(func(r *requestconfig.RequestConfig) error {
		if id == "" {
			return fmt.Errorf("requestoption: WithWorkloadIdentity needs an IdentityProviderID or the DODO_PAYMENTS_IDENTITY_PROVIDER_ID environment variable")
		}
		r.WorkloadIdentity = state
		return nil
	})
}
