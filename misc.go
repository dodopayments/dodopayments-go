// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"github.com/dodopayments/dodopayments-go/option"
)

// MiscService contains methods and other services that help with interacting with
// the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMiscService] method instead.
type MiscService struct {
	Options            []option.RequestOption
	SupportedCountries *MiscSupportedCountryService
}

// NewMiscService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMiscService(opts ...option.RequestOption) (r *MiscService) {
	r = &MiscService{}
	r.Options = opts
	r.SupportedCountries = NewMiscSupportedCountryService(opts...)
	return
}
