// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"github.com/stainless-sdks/dodo-payments-go/option"
)

// CheckoutService contains methods and other services that help with interacting
// with the dodopayments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckoutService] method instead.
type CheckoutService struct {
	Options            []option.RequestOption
	SupportedCountries *CheckoutSupportedCountryService
}

// NewCheckoutService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckoutService(opts ...option.RequestOption) (r *CheckoutService) {
	r = &CheckoutService{}
	r.Options = opts
	r.SupportedCountries = NewCheckoutSupportedCountryService(opts...)
	return
}
