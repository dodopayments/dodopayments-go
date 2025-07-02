// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
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
func NewRefundService(opts ...option.RequestOption) (r *RefundService) {
	r = &RefundService{}
	r.Options = opts
	return
}

func (r *RefundService) New(ctx context.Context, body RefundNewParams, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	path := "refunds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *RefundService) Get(ctx context.Context, refundID string, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	if refundID == "" {
		err = errors.New("missing required refund_id parameter")
		return
	}
	path := fmt.Sprintf("refunds/%s", refundID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *RefundService) List(ctx context.Context, query RefundListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Refund], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

func (r *RefundService) ListAutoPaging(ctx context.Context, query RefundListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Refund] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type Refund struct {
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
	Status RefundStatus `json:"status,required"`
	// The refunded amount.
	Amount int64 `json:"amount,nullable"`
	// The currency of the refund, represented as an ISO 4217 currency code.
	Currency Currency `json:"currency,nullable"`
	// The reason provided for the refund, if any. Optional.
	Reason string     `json:"reason,nullable"`
	JSON   refundJSON `json:"-"`
}

// refundJSON contains the JSON metadata for the struct [Refund]
type refundJSON struct {
	BusinessID  apijson.Field
	CreatedAt   apijson.Field
	IsPartial   apijson.Field
	PaymentID   apijson.Field
	RefundID    apijson.Field
	Status      apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Refund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refundJSON) RawJSON() string {
	return r.raw
}

type RefundStatus string

const (
	RefundStatusSucceeded RefundStatus = "succeeded"
	RefundStatusFailed    RefundStatus = "failed"
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusReview    RefundStatus = "review"
)

func (r RefundStatus) IsKnown() bool {
	switch r {
	case RefundStatusSucceeded, RefundStatusFailed, RefundStatusPending, RefundStatusReview:
		return true
	}
	return false
}

type RefundNewParams struct {
	// The unique identifier of the payment to be refunded.
	PaymentID param.Field[string] `json:"payment_id,required"`
	// Partially Refund an Individual Item
	Items param.Field[[]RefundNewParamsItem] `json:"items"`
	// The reason for the refund, if any. Maximum length is 3000 characters. Optional.
	Reason param.Field[string] `json:"reason"`
}

func (r RefundNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RefundNewParamsItem struct {
	// The id of the item (i.e. `product_id` or `addon_id`)
	ItemID param.Field[string] `json:"item_id,required"`
	// The amount to refund. if None the whole item is refunded
	Amount param.Field[int64] `json:"amount"`
	// Specify if tax is inclusive of the refund. Default true.
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
}

func (r RefundNewParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RefundListParams struct {
	// Get events after this created time
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get events created before this time
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Filter by customer_id
	CustomerID param.Field[string] `query:"customer_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by status
	Status param.Field[RefundListParamsStatus] `query:"status"`
}

// URLQuery serializes [RefundListParams]'s query parameters as `url.Values`.
func (r RefundListParams) URLQuery() (v url.Values) {
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

func (r RefundListParamsStatus) IsKnown() bool {
	switch r {
	case RefundListParamsStatusSucceeded, RefundListParamsStatusFailed, RefundListParamsStatusPending, RefundListParamsStatusReview:
		return true
	}
	return false
}
