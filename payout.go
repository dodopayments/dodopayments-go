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
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
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
func NewPayoutService(opts ...option.RequestOption) (r PayoutService) {
	r = PayoutService{}
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
	Amount int64 `json:"amount,required"`
	// The unique identifier of the business associated with the payout.
	BusinessID string `json:"business_id,required"`
	// The total value of chargebacks associated with the payout.
	//
	// Deprecated: deprecated
	Chargebacks int64 `json:"chargebacks,required"`
	// The timestamp when the payout was created, in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the payout, represented as an ISO 4217 currency code.
	//
	// Any of "AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM",
	// "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP",
	// "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK",
	// "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL",
	// "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF",
	// "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF",
	// "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD",
	// "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN",
	// "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN",
	// "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR",
	// "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN",
	// "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX",
	// "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF",
	// "YER", "ZAR", "ZMW".
	Currency Currency `json:"currency,required"`
	// The fee charged for processing the payout.
	Fee int64 `json:"fee,required"`
	// The payment method used for the payout (e.g., bank transfer, card, etc.).
	PaymentMethod string `json:"payment_method,required"`
	// The unique identifier of the payout.
	PayoutID string `json:"payout_id,required"`
	// The total value of refunds associated with the payout.
	//
	// Deprecated: deprecated
	Refunds int64 `json:"refunds,required"`
	// The current status of the payout.
	//
	// Any of "not_initiated", "in_progress", "on_hold", "failed", "success".
	Status PayoutListResponseStatus `json:"status,required"`
	// The tax applied to the payout.
	//
	// Deprecated: deprecated
	Tax int64 `json:"tax,required"`
	// The timestamp when the payout was last updated, in UTC.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The name of the payout recipient or purpose.
	Name string `json:"name,nullable"`
	// The URL of the document associated with the payout.
	PayoutDocumentURL string `json:"payout_document_url,nullable"`
	// Any additional remarks or notes associated with the payout.
	Remarks string `json:"remarks,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount            respjson.Field
		BusinessID        respjson.Field
		Chargebacks       respjson.Field
		CreatedAt         respjson.Field
		Currency          respjson.Field
		Fee               respjson.Field
		PaymentMethod     respjson.Field
		PayoutID          respjson.Field
		Refunds           respjson.Field
		Status            respjson.Field
		Tax               respjson.Field
		UpdatedAt         respjson.Field
		Name              respjson.Field
		PayoutDocumentURL respjson.Field
		Remarks           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutListResponse) RawJSON() string { return r.JSON.raw }
func (r *PayoutListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type PayoutListParams struct {
	// Get payouts created after this time (inclusive)
	CreatedAtGte param.Opt[time.Time] `query:"created_at_gte,omitzero" format:"date-time" json:"-"`
	// Get payouts created before this time (inclusive)
	CreatedAtLte param.Opt[time.Time] `query:"created_at_lte,omitzero" format:"date-time" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PayoutListParams]'s query parameters as `url.Values`.
func (r PayoutListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
