// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"github.com/dodopayments/dodopayments-go/option"
)

// BlocklistService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlocklistService] method instead.
type BlocklistService struct {
	Options   []option.RequestOption
	Customers *BlocklistCustomerService
}

// NewBlocklistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBlocklistService(opts ...option.RequestOption) (r *BlocklistService) {
	r = &BlocklistService{}
	r.Options = opts
	r.Customers = NewBlocklistCustomerService(opts...)
	return
}
