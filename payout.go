// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
)

// PayoutService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayoutService] method instead.
type PayoutService struct {
	Options []option.RequestOption
}

// NewPayoutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPayoutService(opts ...option.RequestOption) (r *PayoutService) {
	r = &PayoutService{}
	r.Options = opts
	return
}

func (r *PayoutService) List(ctx context.Context, query PayoutListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[PayoutListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "payouts"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PayoutService) ListAutoPaging(ctx context.Context, query PayoutListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[PayoutListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type PayoutListResponse struct {
	// The total amount of the payout.
	Amount int64 `json:"amount" api:"required"`
	// The unique identifier of the business associated with the payout.
	BusinessID string `json:"business_id" api:"required"`
	// The total value of chargebacks associated with the payout.
	//
	// Deprecated: deprecated
	Chargebacks int64 `json:"chargebacks" api:"required"`
	// The timestamp when the payout was created, in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The currency of the payout, represented as an ISO 4217 currency code.
	Currency Currency `json:"currency" api:"required"`
	// The fee charged for processing the payout.
	Fee int64 `json:"fee" api:"required"`
	// The payment method used for the payout (e.g., bank transfer, card, etc.).
	PaymentMethod string `json:"payment_method" api:"required"`
	// The unique identifier of the payout.
	PayoutID string `json:"payout_id" api:"required"`
	// The total value of refunds associated with the payout.
	//
	// Deprecated: deprecated
	Refunds int64 `json:"refunds" api:"required"`
	// The current status of the payout.
	Status PayoutListResponseStatus `json:"status" api:"required"`
	// The tax applied to the payout.
	//
	// Deprecated: deprecated
	Tax int64 `json:"tax" api:"required"`
	// The timestamp when the payout was last updated, in UTC.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The name of the payout recipient or purpose.
	Name string `json:"name" api:"nullable"`
	// The URL of the document associated with the payout.
	PayoutDocumentURL string `json:"payout_document_url" api:"nullable"`
	// Any additional remarks or notes associated with the payout.
	Remarks string                 `json:"remarks" api:"nullable"`
	JSON    payoutListResponseJSON `json:"-"`
}

// payoutListResponseJSON contains the JSON metadata for the struct
// [PayoutListResponse]
type payoutListResponseJSON struct {
	Amount            apijson.Field
	BusinessID        apijson.Field
	Chargebacks       apijson.Field
	CreatedAt         apijson.Field
	Currency          apijson.Field
	Fee               apijson.Field
	PaymentMethod     apijson.Field
	PayoutID          apijson.Field
	Refunds           apijson.Field
	Status            apijson.Field
	Tax               apijson.Field
	UpdatedAt         apijson.Field
	Name              apijson.Field
	PayoutDocumentURL apijson.Field
	Remarks           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PayoutListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutListResponseJSON) RawJSON() string {
	return r.raw
}

// The current status of the payout.
type PayoutListResponseStatus string

const (
	PayoutListResponseStatusNotInitiated PayoutListResponseStatus = "not_initiated"
	PayoutListResponseStatusInProgress   PayoutListResponseStatus = "in_progress"
	PayoutListResponseStatusOnHold       PayoutListResponseStatus = "on_hold"
	PayoutListResponseStatusFailed       PayoutListResponseStatus = "failed"
	PayoutListResponseStatusSuccess      PayoutListResponseStatus = "success"
)

func (r PayoutListResponseStatus) IsKnown() bool {
	switch r {
	case PayoutListResponseStatusNotInitiated, PayoutListResponseStatusInProgress, PayoutListResponseStatusOnHold, PayoutListResponseStatusFailed, PayoutListResponseStatusSuccess:
		return true
	}
	return false
}

type PayoutListParams struct {
	// Get payouts created after this time (inclusive)
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get payouts created before this time (inclusive)
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [PayoutListParams]'s query parameters as `url.Values`.
func (r PayoutListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
