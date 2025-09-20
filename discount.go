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
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
)

// DiscountService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiscountService] method instead.
type DiscountService struct {
	Options []option.RequestOption
}

// NewDiscountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDiscountService(opts ...option.RequestOption) (r *DiscountService) {
	r = &DiscountService{}
	r.Options = opts
	return
}

// POST /discounts If `code` is omitted or empty, a random 16-char uppercase code
// is generated.
func (r *DiscountService) New(ctx context.Context, body DiscountNewParams, opts ...option.RequestOption) (res *Discount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "discounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// GET /discounts/{discount_id}
func (r *DiscountService) Get(ctx context.Context, discountID string, opts ...option.RequestOption) (res *Discount, err error) {
	opts = slices.Concat(r.Options, opts)
	if discountID == "" {
		err = errors.New("missing required discount_id parameter")
		return
	}
	path := fmt.Sprintf("discounts/%s", discountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// PATCH /discounts/{discount_id}
func (r *DiscountService) Update(ctx context.Context, discountID string, body DiscountUpdateParams, opts ...option.RequestOption) (res *Discount, err error) {
	opts = slices.Concat(r.Options, opts)
	if discountID == "" {
		err = errors.New("missing required discount_id parameter")
		return
	}
	path := fmt.Sprintf("discounts/%s", discountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// GET /discounts
func (r *DiscountService) List(ctx context.Context, query DiscountListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Discount], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "discounts"
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

// GET /discounts
func (r *DiscountService) ListAutoPaging(ctx context.Context, query DiscountListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Discount] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// DELETE /discounts/{discount_id}
func (r *DiscountService) Delete(ctx context.Context, discountID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if discountID == "" {
		err = errors.New("missing required discount_id parameter")
		return
	}
	path := fmt.Sprintf("discounts/%s", discountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Discount struct {
	// The discount amount.
	//
	//   - If `discount_type` is `percentage`, this is in **basis points** (e.g., 540 =>
	//     5.4%).
	//   - Otherwise, this is **USD cents** (e.g., 100 => `$1.00`).
	Amount int64 `json:"amount,required"`
	// The business this discount belongs to.
	BusinessID string `json:"business_id,required"`
	// The discount code (up to 16 chars).
	Code string `json:"code,required"`
	// Timestamp when the discount is created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The unique discount ID
	DiscountID string `json:"discount_id,required"`
	// List of product IDs to which this discount is restricted.
	RestrictedTo []string `json:"restricted_to,required"`
	// How many times this discount has been used.
	TimesUsed int64 `json:"times_used,required"`
	// The type of discount, e.g. `percentage`, `flat`, or `flat_per_unit`.
	Type DiscountType `json:"type,required"`
	// Optional date/time after which discount is expired.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Name for the Discount
	Name string `json:"name,nullable"`
	// Number of subscription billing cycles this discount is valid for. If not
	// provided, the discount will be applied indefinitely to all recurring payments
	// related to the subscription.
	SubscriptionCycles int64 `json:"subscription_cycles,nullable"`
	// Usage limit for this discount, if any.
	UsageLimit int64        `json:"usage_limit,nullable"`
	JSON       discountJSON `json:"-"`
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	Amount             apijson.Field
	BusinessID         apijson.Field
	Code               apijson.Field
	CreatedAt          apijson.Field
	DiscountID         apijson.Field
	RestrictedTo       apijson.Field
	TimesUsed          apijson.Field
	Type               apijson.Field
	ExpiresAt          apijson.Field
	Name               apijson.Field
	SubscriptionCycles apijson.Field
	UsageLimit         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Discount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountJSON) RawJSON() string {
	return r.raw
}

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "percentage"
)

func (r DiscountType) IsKnown() bool {
	switch r {
	case DiscountTypePercentage:
		return true
	}
	return false
}

type DiscountNewParams struct {
	// The discount amount.
	//
	//   - If `discount_type` is **not** `percentage`, `amount` is in **USD cents**. For
	//     example, `100` means `$1.00`. Only USD is allowed.
	//   - If `discount_type` **is** `percentage`, `amount` is in **basis points**. For
	//     example, `540` means `5.4%`.
	//
	// Must be at least 1.
	Amount param.Field[int64] `json:"amount,required"`
	// The discount type (e.g. `percentage`, `flat`, or `flat_per_unit`).
	Type param.Field[DiscountType] `json:"type,required"`
	// Optionally supply a code (will be uppercased).
	//
	// - Must be at least 3 characters if provided.
	// - If omitted, a random 16-character code is generated.
	Code param.Field[string] `json:"code"`
	// When the discount expires, if ever.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	Name      param.Field[string]    `json:"name"`
	// List of product IDs to restrict usage (if any).
	RestrictedTo param.Field[[]string] `json:"restricted_to"`
	// Number of subscription billing cycles this discount is valid for. If not
	// provided, the discount will be applied indefinitely to all recurring payments
	// related to the subscription.
	SubscriptionCycles param.Field[int64] `json:"subscription_cycles"`
	// How many times this discount can be used (if any). Must be >= 1 if provided.
	UsageLimit param.Field[int64] `json:"usage_limit"`
}

func (r DiscountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DiscountUpdateParams struct {
	// If present, update the discount amount:
	//
	//   - If `discount_type` is `percentage`, this represents **basis points** (e.g.,
	//     `540` = `5.4%`).
	//   - Otherwise, this represents **USD cents** (e.g., `100` = `$1.00`).
	//
	// Must be at least 1 if provided.
	Amount param.Field[int64] `json:"amount"`
	// If present, update the discount code (uppercase).
	Code      param.Field[string]    `json:"code"`
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	Name      param.Field[string]    `json:"name"`
	// If present, replaces all restricted product IDs with this new set. To remove all
	// restrictions, send empty array
	RestrictedTo param.Field[[]string] `json:"restricted_to"`
	// Number of subscription billing cycles this discount is valid for. If not
	// provided, the discount will be applied indefinitely to all recurring payments
	// related to the subscription.
	SubscriptionCycles param.Field[int64] `json:"subscription_cycles"`
	// If present, update the discount type.
	Type       param.Field[DiscountType] `json:"type"`
	UsageLimit param.Field[int64]        `json:"usage_limit"`
}

func (r DiscountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DiscountListParams struct {
	// Page number (default = 0).
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size (default = 10, max = 100).
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [DiscountListParams]'s query parameters as `url.Values`.
func (r DiscountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
