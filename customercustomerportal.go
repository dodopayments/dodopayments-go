// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
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
}

// NewCustomerCustomerPortalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerCustomerPortalService(opts ...option.RequestOption) (r *CustomerCustomerPortalService) {
	r = &CustomerCustomerPortalService{}
	r.Options = opts
	return
}

func (r *CustomerCustomerPortalService) New(ctx context.Context, customerID string, body CustomerCustomerPortalNewParams, opts ...option.RequestOption) (res *CustomerPortalSession, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/customer-portal/session", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CustomerCustomerPortalNewParams struct {
	// If true, will send link to user.
	SendEmail param.Field[bool] `query:"send_email"`
}

// URLQuery serializes [CustomerCustomerPortalNewParams]'s query parameters as
// `url.Values`.
func (r CustomerCustomerPortalNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
