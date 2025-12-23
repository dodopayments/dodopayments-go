// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
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

// RefundService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefundService] method instead.
type RefundService struct {
	Options []option.RequestOption
}

// NewRefundService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRefundService(opts ...option.RequestOption) (r RefundService) {
	r = RefundService{}
	r.Options = opts
	return
}

func (r *RefundService) New(ctx context.Context, body RefundNewParams, opts ...option.RequestOption) (res *Refund, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "refunds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *RefundService) Get(ctx context.Context, refundID string, opts ...option.RequestOption) (res *Refund, err error) {
	opts = slices.Concat(r.Options, opts)
	if refundID == "" {
		err = errors.New("missing required refund_id parameter")
		return
	}
	path := fmt.Sprintf("refunds/%s", refundID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *RefundService) List(ctx context.Context, query RefundListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[RefundListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "refunds"
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

func (r *RefundService) ListAutoPaging(ctx context.Context, query RefundListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[RefundListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type Refund struct {
	// The unique identifier of the business issuing the refund.
	BusinessID string `json:"business_id,required"`
	// The timestamp of when the refund was created in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Details about the customer for this refund (from the associated payment)
	Customer CustomerLimitedDetails `json:"customer,required"`
	// If true the refund is a partial refund
	IsPartial bool `json:"is_partial,required"`
	// Additional metadata stored with the refund.
	Metadata map[string]string `json:"metadata,required"`
	// The unique identifier of the payment associated with the refund.
	PaymentID string `json:"payment_id,required"`
	// The unique identifier of the refund.
	RefundID string `json:"refund_id,required"`
	// The current status of the refund.
	//
	// Any of "succeeded", "failed", "pending", "review".
	Status RefundStatus `json:"status,required"`
	// The refunded amount.
	Amount int64 `json:"amount,nullable"`
	// The currency of the refund, represented as an ISO 4217 currency code.
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
	Currency Currency `json:"currency,nullable"`
	// The reason provided for the refund, if any. Optional.
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		CreatedAt   respjson.Field
		Customer    respjson.Field
		IsPartial   respjson.Field
		Metadata    respjson.Field
		PaymentID   respjson.Field
		RefundID    respjson.Field
		Status      respjson.Field
		Amount      respjson.Field
		Currency    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Refund) RawJSON() string { return r.JSON.raw }
func (r *Refund) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RefundStatus string

const (
	RefundStatusSucceeded RefundStatus = "succeeded"
	RefundStatusFailed    RefundStatus = "failed"
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusReview    RefundStatus = "review"
)

type RefundListResponse struct {
	// The unique identifier of the business issuing the refund.
	BusinessID string `json:"business_id,required"`
	// The timestamp of when the refund was created in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If true the refund is a partial refund
	IsPartial bool `json:"is_partial,required"`
	// The unique identifier of the payment associated with the refund.
	PaymentID string `json:"payment_id,required"`
	// The unique identifier of the refund.
	RefundID string `json:"refund_id,required"`
	// The current status of the refund.
	//
	// Any of "succeeded", "failed", "pending", "review".
	Status RefundStatus `json:"status,required"`
	// The refunded amount.
	Amount int64 `json:"amount,nullable"`
	// The currency of the refund, represented as an ISO 4217 currency code.
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
	Currency Currency `json:"currency,nullable"`
	// The reason provided for the refund, if any. Optional.
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		CreatedAt   respjson.Field
		IsPartial   respjson.Field
		PaymentID   respjson.Field
		RefundID    respjson.Field
		Status      respjson.Field
		Amount      respjson.Field
		Currency    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RefundListResponse) RawJSON() string { return r.JSON.raw }
func (r *RefundListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RefundNewParams struct {
	// The unique identifier of the payment to be refunded.
	PaymentID string `json:"payment_id,required"`
	// The reason for the refund, if any. Maximum length is 3000 characters. Optional.
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Partially Refund an Individual Item
	Items []RefundNewParamsItem `json:"items,omitzero"`
	// Additional metadata associated with the refund.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r RefundNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RefundNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RefundNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ItemID is required.
type RefundNewParamsItem struct {
	// The id of the item (i.e. `product_id` or `addon_id`)
	ItemID string `json:"item_id,required"`
	// The amount to refund. if None the whole item is refunded
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// Specify if tax is inclusive of the refund. Default true.
	TaxInclusive param.Opt[bool] `json:"tax_inclusive,omitzero"`
	paramObj
}

func (r RefundNewParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow RefundNewParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RefundNewParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RefundListParams struct {
	// Get events after this created time
	CreatedAtGte param.Opt[time.Time] `query:"created_at_gte,omitzero" format:"date-time" json:"-"`
	// Get events created before this time
	CreatedAtLte param.Opt[time.Time] `query:"created_at_lte,omitzero" format:"date-time" json:"-"`
	// Filter by customer_id
	CustomerID param.Opt[string] `query:"customer_id,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter by status
	//
	// Any of "succeeded", "failed", "pending", "review".
	Status RefundListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RefundListParams]'s query parameters as `url.Values`.
func (r RefundListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status
type RefundListParamsStatus string

const (
	RefundListParamsStatusSucceeded RefundListParamsStatus = "succeeded"
	RefundListParamsStatusFailed    RefundListParamsStatus = "failed"
	RefundListParamsStatusPending   RefundListParamsStatus = "pending"
	RefundListParamsStatusReview    RefundListParamsStatus = "review"
)
