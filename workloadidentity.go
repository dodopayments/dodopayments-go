// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// WorkloadIdentityService manages the access token that the client gets with
// [option.WithWorkloadIdentity]. You should not instantiate this service
// directly, and instead use the [NewWorkloadIdentityService] method instead.
type WorkloadIdentityService struct {
	Options []option.RequestOption
}

// NewWorkloadIdentityService generates a new service that applies the given
// options to each request.
func NewWorkloadIdentityService(opts ...option.RequestOption) (r *WorkloadIdentityService) {
	r = &WorkloadIdentityService{}
	r.Options = opts
	return
}

// Revoke revokes the cached access token and clears the cache. The next request
// exchanges a new subject token. Revoke does nothing when no access token is
// cached.
func (r *WorkloadIdentityService) Revoke(ctx context.Context, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, "workload-identity/token/revoke", nil, nil, opts...)
	if err != nil {
		return err
	}
	return cfg.RevokeWorkloadIdentityToken()
}
