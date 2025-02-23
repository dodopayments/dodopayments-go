// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"github.com/dodopayments/dodopayments-go/option"
)

// CustomerCustomerPortalService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCustomerPortalService] method instead.
type CustomerCustomerPortalService struct {
	Options []option.RequestOption
	Session *CustomerCustomerPortalSessionService
}

// NewCustomerCustomerPortalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerCustomerPortalService(opts ...option.RequestOption) (r *CustomerCustomerPortalService) {
	r = &CustomerCustomerPortalService{}
	r.Options = opts
	r.Session = NewCustomerCustomerPortalSessionService(opts...)
	return
}
