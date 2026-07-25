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
	return res, err
}

// GET /discounts/{discount_id}
func (r *DiscountService) Get(ctx context.Context, discountID string, opts ...option.RequestOption) (res *Discount, err error) {
	opts = slices.Concat(r.Options, opts)
	if discountID == "" {
		err = errors.New("missing required discount_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("discounts/%s", discountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// PATCH /discounts/{discount_id}
func (r *DiscountService) Update(ctx context.Context, discountID string, body DiscountUpdateParams, opts ...option.RequestOption) (res *Discount, err error) {
	opts = slices.Concat(r.Options, opts)
	if discountID == "" {
		err = errors.New("missing required discount_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("discounts/%s", discountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if discountID == "" {
		err = errors.New("missing required discount_id parameter")
		return err
	}
	path := fmt.Sprintf("discounts/%s", discountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Validate and fetch a discount by its code name (e.g., "SAVE20"). This allows
// real-time validation directly against the API using the human-readable discount
// code instead of requiring the internal discount_id.
func (r *DiscountService) GetByCode(ctx context.Context, code string, opts ...option.RequestOption) (res *Discount, err error) {
	opts = slices.Concat(r.Options, opts)
	if code == "" {
		err = errors.New("missing required code parameter")
		return nil, err
	}
	path := fmt.Sprintf("discounts/code/%s", code)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Discount struct {
	// The discount amount in **basis points** (e.g., 540 => 5.4%).
	Amount int64 `json:"amount" api:"required"`
	// The business this discount belongs to.
	BusinessID string `json:"business_id" api:"required"`
	// The discount code (up to 16 chars).
	Code string `json:"code" api:"required"`
	// Timestamp when the discount is created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Who may redeem this discount code.
	CustomerEligibility DiscountCustomerEligibility `json:"customer_eligibility" api:"required"`
	// The unique discount ID
	DiscountID string `json:"discount_id" api:"required"`
	// Arbitrary key-value metadata. Values can be string, integer, number, or boolean.
	Metadata Metadata `json:"metadata" api:"required"`
	// Whether this discount should be preserved when a subscription changes plans.
	// Default: false (discount is removed on plan change)
	PreserveOnPlanChange bool `json:"preserve_on_plan_change" api:"required"`
	// List of product IDs to which this discount is restricted.
	RestrictedTo []string `json:"restricted_to" api:"required"`
	// How many times this discount has been used.
	TimesUsed int64 `json:"times_used" api:"required"`
	// The type of discount (`percentage` or `flat`).
	Type DiscountType `json:"type" api:"required"`
	// Per-currency options (flat deduction / percentage cap + minimum subtotal). Empty
	// for discounts without any configured currency options.
	CurrencyOptions []DiscountCurrencyOption `json:"currency_options"`
	// Optional date/time after which discount is expired.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Name for the Discount
	Name string `json:"name" api:"nullable"`
	// Maximum number of times a single customer may redeem this discount, if any.
	PerCustomerUsageLimit int64 `json:"per_customer_usage_limit" api:"nullable"`
	// Optional date/time before which the discount is not yet active. NULL = active
	// immediately.
	StartsAt time.Time `json:"starts_at" api:"nullable" format:"date-time"`
	// Number of subscription billing cycles this discount is valid for. If not
	// provided, the discount will be applied indefinitely to all recurring payments
	// related to the subscription.
	SubscriptionCycles int64 `json:"subscription_cycles" api:"nullable"`
	// Usage limit for this discount, if any.
	UsageLimit int64        `json:"usage_limit" api:"nullable"`
	JSON       discountJSON `json:"-"`
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	Amount                apijson.Field
	BusinessID            apijson.Field
	Code                  apijson.Field
	CreatedAt             apijson.Field
	CustomerEligibility   apijson.Field
	DiscountID            apijson.Field
	Metadata              apijson.Field
	PreserveOnPlanChange  apijson.Field
	RestrictedTo          apijson.Field
	TimesUsed             apijson.Field
	Type                  apijson.Field
	CurrencyOptions       apijson.Field
	ExpiresAt             apijson.Field
	Name                  apijson.Field
	PerCustomerUsageLimit apijson.Field
	StartsAt              apijson.Field
	SubscriptionCycles    apijson.Field
	UsageLimit            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Discount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountJSON) RawJSON() string {
	return r.raw
}

// Who may redeem this discount code.
type DiscountCustomerEligibility string

const (
	DiscountCustomerEligibilityAny       DiscountCustomerEligibility = "any"
	DiscountCustomerEligibilityFirstTime DiscountCustomerEligibility = "first_time"
	DiscountCustomerEligibilityExisting  DiscountCustomerEligibility = "existing"
	DiscountCustomerEligibilitySpecific  DiscountCustomerEligibility = "specific"
)

func (r DiscountCustomerEligibility) IsKnown() bool {
	switch r {
	case DiscountCustomerEligibilityAny, DiscountCustomerEligibilityFirstTime, DiscountCustomerEligibilityExisting, DiscountCustomerEligibilitySpecific:
		return true
	}
	return false
}

// A per-currency discount option (response shape). `max_amount_possible` mirrors
// the DB column of the same name.
type DiscountCurrencyOption struct {
	// The currency this option applies to.
	Currency Currency `json:"currency" api:"required"`
	// Whether this is the default row FX conversions pivot from.
	IsDefault bool `json:"is_default" api:"required"`
	// Eligible-cart threshold in this currency's subunits (0 = no minimum).
	MinimumSubtotal int64 `json:"minimum_subtotal" api:"required"`
	// The most this code discounts in this currency's subunits (flat deduction or
	// percentage cap).
	MaxAmountPossible int64                      `json:"max_amount_possible" api:"nullable"`
	JSON              discountCurrencyOptionJSON `json:"-"`
}

// discountCurrencyOptionJSON contains the JSON metadata for the struct
// [DiscountCurrencyOption]
type discountCurrencyOptionJSON struct {
	Currency          apijson.Field
	IsDefault         apijson.Field
	MinimumSubtotal   apijson.Field
	MaxAmountPossible apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DiscountCurrencyOption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountCurrencyOptionJSON) RawJSON() string {
	return r.raw
}

// Response struct for a discount with its position in a stack and optional
// cycle-tracking information (for subscriptions).
type DiscountDetail struct {
	// The discount amount in **basis points** (e.g., 540 => 5.4%).
	Amount int64 `json:"amount" api:"required"`
	// The business this discount belongs to
	BusinessID string `json:"business_id" api:"required"`
	// The discount code
	Code string `json:"code" api:"required"`
	// Timestamp when the discount was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The unique discount ID
	DiscountID string `json:"discount_id" api:"required"`
	// Additional metadata
	Metadata Metadata `json:"metadata" api:"required"`
	// Position of this discount in the stack (0-based)
	Position int64 `json:"position" api:"required"`
	// Whether this discount should be preserved when a subscription changes plans
	PreserveOnPlanChange bool `json:"preserve_on_plan_change" api:"required"`
	// List of product IDs to which this discount is restricted
	RestrictedTo []string `json:"restricted_to" api:"required"`
	// How many times this discount has been used
	TimesUsed int64 `json:"times_used" api:"required"`
	// The type of discount
	Type DiscountType `json:"type" api:"required"`
	// Remaining billing cycles for this discount on this subscription (None for
	// one-time payments)
	CyclesRemaining int64 `json:"cycles_remaining" api:"nullable"`
	// Optional date/time after which discount is expired
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Name for the Discount
	Name string `json:"name" api:"nullable"`
	// Number of subscription billing cycles this discount is valid for
	SubscriptionCycles int64 `json:"subscription_cycles" api:"nullable"`
	// Usage limit for this discount, if any
	UsageLimit int64              `json:"usage_limit" api:"nullable"`
	JSON       discountDetailJSON `json:"-"`
}

// discountDetailJSON contains the JSON metadata for the struct [DiscountDetail]
type discountDetailJSON struct {
	Amount               apijson.Field
	BusinessID           apijson.Field
	Code                 apijson.Field
	CreatedAt            apijson.Field
	DiscountID           apijson.Field
	Metadata             apijson.Field
	Position             apijson.Field
	PreserveOnPlanChange apijson.Field
	RestrictedTo         apijson.Field
	TimesUsed            apijson.Field
	Type                 apijson.Field
	CyclesRemaining      apijson.Field
	ExpiresAt            apijson.Field
	Name                 apijson.Field
	SubscriptionCycles   apijson.Field
	UsageLimit           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DiscountDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountDetailJSON) RawJSON() string {
	return r.raw
}

type DiscountType string

const (
	DiscountTypeFlat       DiscountType = "flat"
	DiscountTypePercentage DiscountType = "percentage"
)

func (r DiscountType) IsKnown() bool {
	switch r {
	case DiscountTypeFlat, DiscountTypePercentage:
		return true
	}
	return false
}

type DiscountNewParams struct {
	// The discount amount in **basis points** (e.g. `540` means `5.4%`, `10000` means
	// `100%`).
	//
	// Must be at least 1.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The discount type: `percentage` or `flat` (`flat_per_unit` stays blocked).
	Type param.Field[DiscountType] `json:"type" api:"required"`
	// Optionally supply a code (will be uppercased).
	//
	// - Must be at least 3 characters if provided.
	// - If omitted, a random 16-character code is generated.
	Code param.Field[string] `json:"code"`
	// Per-currency options (flat deduction / percentage cap + minimum subtotal).
	// Required for `flat` codes (must include a resolvable default); optional
	// per-currency caps for `percentage` codes. Per-row invariants are checked in
	// `normalize_currency_options`, not via `#[validate(nested)]`.
	CurrencyOptions param.Field[[]DiscountNewParamsCurrencyOption] `json:"currency_options"`
	// Who may redeem this discount code. Defaults to `any` (unrestricted). `specific`
	// starts with zero attached customers (fails closed) until customers are attached
	// via `POST /discounts/{id}/customers`.
	CustomerEligibility param.Field[DiscountNewParamsCustomerEligibility] `json:"customer_eligibility"`
	// When the discount expires, if ever.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Additional metadata for the discount
	Metadata param.Field[MetadataParam] `json:"metadata"`
	Name     param.Field[string]        `json:"name"`
	// Maximum number of times a single customer may redeem this discount. Must be
	// `<= usage_limit` when both are set.
	PerCustomerUsageLimit param.Field[int64] `json:"per_customer_usage_limit"`
	// Whether this discount should be preserved when a subscription changes plans.
	// Default: false (discount is removed on plan change)
	PreserveOnPlanChange param.Field[bool] `json:"preserve_on_plan_change"`
	// List of product IDs to restrict usage (if any).
	RestrictedTo param.Field[[]string] `json:"restricted_to"`
	// When the discount becomes active, if scheduled for the future. NULL = active
	// immediately. Must be strictly before `expires_at` when both are set.
	StartsAt param.Field[time.Time] `json:"starts_at" format:"date-time"`
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

// A per-currency discount option (request shape).
//
// `max_amount_possible` is the most this code discounts in this currency — the
// flat deduction for `flat` codes, or the max-discount cap for `percentage` codes.
// Maps to the DB column of the same name.
type DiscountNewParamsCurrencyOption struct {
	// The currency this option applies to.
	Currency param.Field[Currency] `json:"currency" api:"required"`
	// Whether this row is the default to convert from for unconfigured currencies. At
	// most one row per discount may be default.
	IsDefault param.Field[bool] `json:"is_default"`
	// The most this code discounts in this currency's subunits. For `flat` codes this
	// is the deduction; for `percentage` codes it is the max-discount cap. Must be > 0
	// if provided.
	MaxAmountPossible param.Field[int64] `json:"max_amount_possible"`
	// Eligible-cart threshold in this currency's subunits (0 = no minimum).
	MinimumSubtotal param.Field[int64] `json:"minimum_subtotal"`
}

func (r DiscountNewParamsCurrencyOption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Who may redeem this discount code. Defaults to `any` (unrestricted). `specific`
// starts with zero attached customers (fails closed) until customers are attached
// via `POST /discounts/{id}/customers`.
type DiscountNewParamsCustomerEligibility string

const (
	DiscountNewParamsCustomerEligibilityAny       DiscountNewParamsCustomerEligibility = "any"
	DiscountNewParamsCustomerEligibilityFirstTime DiscountNewParamsCustomerEligibility = "first_time"
	DiscountNewParamsCustomerEligibilityExisting  DiscountNewParamsCustomerEligibility = "existing"
	DiscountNewParamsCustomerEligibilitySpecific  DiscountNewParamsCustomerEligibility = "specific"
)

func (r DiscountNewParamsCustomerEligibility) IsKnown() bool {
	switch r {
	case DiscountNewParamsCustomerEligibilityAny, DiscountNewParamsCustomerEligibilityFirstTime, DiscountNewParamsCustomerEligibilityExisting, DiscountNewParamsCustomerEligibilitySpecific:
		return true
	}
	return false
}

type DiscountUpdateParams struct {
	// If present, update the discount amount in **basis points** (e.g., `540` =
	// `5.4%`, `10000` = `100%`).
	//
	// Must be at least 1 if provided.
	Amount param.Field[int64] `json:"amount"`
	// If present, update the discount code (uppercase).
	Code param.Field[string] `json:"code"`
	// If present, fully replaces the discount's currency options (replace-set
	// semantics, like `restricted_to`). Send an empty array to clear them.
	CurrencyOptions param.Field[[]DiscountUpdateParamsCurrencyOption] `json:"currency_options"`
	// If present, update who may redeem this discount. Plain field (not
	// double-option): the DB column is `NOT NULL`, so it can never be cleared back to
	// unset, only changed to another `CustomerEligibility` value.
	CustomerEligibility param.Field[DiscountUpdateParamsCustomerEligibility] `json:"customer_eligibility"`
	ExpiresAt           param.Field[time.Time]                               `json:"expires_at" format:"date-time"`
	// Additional metadata for the discount
	Metadata param.Field[MetadataParam] `json:"metadata"`
	Name     param.Field[string]        `json:"name"`
	// If present, update the per-customer usage limit (double-option: send `null` to
	// clear it back to unlimited). Must be `<= usage_limit` (the value in effect after
	// this patch) when both are set.
	PerCustomerUsageLimit param.Field[int64] `json:"per_customer_usage_limit"`
	// Whether this discount should be preserved when a subscription changes plans. If
	// not provided, the existing value is kept.
	PreserveOnPlanChange param.Field[bool] `json:"preserve_on_plan_change"`
	// If present, replaces all restricted product IDs with this new set. To remove all
	// restrictions, send empty array
	RestrictedTo param.Field[[]string] `json:"restricted_to"`
	// If present, update `starts_at` (double-option: send `null` to clear it).
	StartsAt param.Field[time.Time] `json:"starts_at" format:"date-time"`
	// Number of subscription billing cycles this discount is valid for. If not
	// provided, the discount will be applied indefinitely to all recurring payments
	// related to the subscription.
	SubscriptionCycles param.Field[int64] `json:"subscription_cycles"`
	// If present, update the discount type (`percentage` or `flat`).
	Type       param.Field[DiscountType] `json:"type"`
	UsageLimit param.Field[int64]        `json:"usage_limit"`
}

func (r DiscountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A per-currency discount option (request shape).
//
// `max_amount_possible` is the most this code discounts in this currency — the
// flat deduction for `flat` codes, or the max-discount cap for `percentage` codes.
// Maps to the DB column of the same name.
type DiscountUpdateParamsCurrencyOption struct {
	// The currency this option applies to.
	Currency param.Field[Currency] `json:"currency" api:"required"`
	// Whether this row is the default to convert from for unconfigured currencies. At
	// most one row per discount may be default.
	IsDefault param.Field[bool] `json:"is_default"`
	// The most this code discounts in this currency's subunits. For `flat` codes this
	// is the deduction; for `percentage` codes it is the max-discount cap. Must be > 0
	// if provided.
	MaxAmountPossible param.Field[int64] `json:"max_amount_possible"`
	// Eligible-cart threshold in this currency's subunits (0 = no minimum).
	MinimumSubtotal param.Field[int64] `json:"minimum_subtotal"`
}

func (r DiscountUpdateParamsCurrencyOption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If present, update who may redeem this discount. Plain field (not
// double-option): the DB column is `NOT NULL`, so it can never be cleared back to
// unset, only changed to another `CustomerEligibility` value.
type DiscountUpdateParamsCustomerEligibility string

const (
	DiscountUpdateParamsCustomerEligibilityAny       DiscountUpdateParamsCustomerEligibility = "any"
	DiscountUpdateParamsCustomerEligibilityFirstTime DiscountUpdateParamsCustomerEligibility = "first_time"
	DiscountUpdateParamsCustomerEligibilityExisting  DiscountUpdateParamsCustomerEligibility = "existing"
	DiscountUpdateParamsCustomerEligibilitySpecific  DiscountUpdateParamsCustomerEligibility = "specific"
)

func (r DiscountUpdateParamsCustomerEligibility) IsKnown() bool {
	switch r {
	case DiscountUpdateParamsCustomerEligibilityAny, DiscountUpdateParamsCustomerEligibilityFirstTime, DiscountUpdateParamsCustomerEligibilityExisting, DiscountUpdateParamsCustomerEligibilitySpecific:
		return true
	}
	return false
}

type DiscountListParams struct {
	// Filter by active status. `true` = currently redeemable (started, not expired,
	// not usage-exhausted). `false` = not currently redeemable (expired,
	// usage-exhausted, or pending a future `starts_at`).
	Active param.Field[bool] `query:"active"`
	// Filter by discount code (partial match, case-insensitive)
	Code param.Field[string] `query:"code"`
	// Filter by discount type
	DiscountType param.Field[DiscountType] `query:"discount_type"`
	// Page number (default = 0).
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size (default = 10, max = 100).
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by product restriction (only discounts that apply to this product)
	ProductID param.Field[string] `query:"product_id"`
}

// URLQuery serializes [DiscountListParams]'s query parameters as `url.Values`.
func (r DiscountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
