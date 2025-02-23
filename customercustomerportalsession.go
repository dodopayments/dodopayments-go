// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// CustomerCustomerPortalSessionService contains methods and other services that
// help with interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCustomerPortalSessionService] method instead.
type CustomerCustomerPortalSessionService struct {
	Options []option.RequestOption
}

// NewCustomerCustomerPortalSessionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomerCustomerPortalSessionService(opts ...option.RequestOption) (r *CustomerCustomerPortalSessionService) {
	r = &CustomerCustomerPortalSessionService{}
	r.Options = opts
	return
}

func (r *CustomerCustomerPortalSessionService) New(ctx context.Context, customerID string, body CustomerCustomerPortalSessionNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/customer-portal/session", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type CustomerCustomerPortalSessionNewParams struct {
	// If true, will send link to user.
	SendEmail param.Field[bool] `query:"send_email"`
}

// URLQuery serializes [CustomerCustomerPortalSessionNewParams]'s query parameters
// as `url.Values`.
func (r CustomerCustomerPortalSessionNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
