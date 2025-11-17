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

// SubscriptionService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *SubscriptionService) Get(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[SubscriptionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "subscriptions"
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

func (r *SubscriptionService) ListAutoPaging(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[SubscriptionListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *SubscriptionService) ChangePlan(ctx context.Context, subscriptionID string, body SubscriptionChangePlanParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/change-plan", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

func (r *SubscriptionService) Charge(ctx context.Context, subscriptionID string, body SubscriptionChargeParams, opts ...option.RequestOption) (res *SubscriptionChargeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/charge", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get detailed usage history for a subscription that includes usage-based billing
// (metered components). This endpoint provides insights into customer usage
// patterns and billing calculations over time.
//
// ## What You'll Get:
//
//   - **Billing periods**: Each item represents a billing cycle with start and end
//     dates
//   - **Meter usage**: Detailed breakdown of usage for each meter configured on the
//     subscription
//   - **Usage calculations**: Total units consumed, free threshold units, and
//     chargeable units
//   - **Historical tracking**: Complete audit trail of usage-based charges
//
// ## Use Cases:
//
// - **Customer support**: Investigate billing questions and usage discrepancies
// - **Usage analytics**: Analyze customer consumption patterns over time
// - **Billing transparency**: Provide customers with detailed usage breakdowns
// - **Revenue optimization**: Identify usage trends to optimize pricing strategies
//
// ## Filtering Options:
//
// - **Date range filtering**: Get usage history for specific time periods
// - **Meter-specific filtering**: Focus on usage for a particular meter
// - **Pagination**: Navigate through large usage histories efficiently
//
// ## Important Notes:
//
//   - Only returns data for subscriptions with usage-based (metered) components
//   - Usage history is organized by billing periods (subscription cycles)
//   - Free threshold units are calculated and displayed separately from chargeable
//     units
//   - Historical data is preserved even if meter configurations change
//
// ## Example Query Patterns:
//
//   - Get last 3 months:
//     `?start_date=2024-01-01T00:00:00Z&end_date=2024-03-31T23:59:59Z`
//   - Filter by meter: `?meter_id=mtr_api_requests`
//   - Paginate results: `?page_size=20&page_number=1`
//   - Recent usage: `?start_date=2024-03-01T00:00:00Z` (from March 1st to now)
func (r *SubscriptionService) GetUsageHistory(ctx context.Context, subscriptionID string, query SubscriptionGetUsageHistoryParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[SubscriptionGetUsageHistoryResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/usage-history", subscriptionID)
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

// Get detailed usage history for a subscription that includes usage-based billing
// (metered components). This endpoint provides insights into customer usage
// patterns and billing calculations over time.
//
// ## What You'll Get:
//
//   - **Billing periods**: Each item represents a billing cycle with start and end
//     dates
//   - **Meter usage**: Detailed breakdown of usage for each meter configured on the
//     subscription
//   - **Usage calculations**: Total units consumed, free threshold units, and
//     chargeable units
//   - **Historical tracking**: Complete audit trail of usage-based charges
//
// ## Use Cases:
//
// - **Customer support**: Investigate billing questions and usage discrepancies
// - **Usage analytics**: Analyze customer consumption patterns over time
// - **Billing transparency**: Provide customers with detailed usage breakdowns
// - **Revenue optimization**: Identify usage trends to optimize pricing strategies
//
// ## Filtering Options:
//
// - **Date range filtering**: Get usage history for specific time periods
// - **Meter-specific filtering**: Focus on usage for a particular meter
// - **Pagination**: Navigate through large usage histories efficiently
//
// ## Important Notes:
//
//   - Only returns data for subscriptions with usage-based (metered) components
//   - Usage history is organized by billing periods (subscription cycles)
//   - Free threshold units are calculated and displayed separately from chargeable
//     units
//   - Historical data is preserved even if meter configurations change
//
// ## Example Query Patterns:
//
//   - Get last 3 months:
//     `?start_date=2024-01-01T00:00:00Z&end_date=2024-03-31T23:59:59Z`
//   - Filter by meter: `?meter_id=mtr_api_requests`
//   - Paginate results: `?page_size=20&page_number=1`
//   - Recent usage: `?start_date=2024-03-01T00:00:00Z` (from March 1st to now)
func (r *SubscriptionService) GetUsageHistoryAutoPaging(ctx context.Context, subscriptionID string, query SubscriptionGetUsageHistoryParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[SubscriptionGetUsageHistoryResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.GetUsageHistory(ctx, subscriptionID, query, opts...))
}

func (r *SubscriptionService) UpdatePaymentMethod(ctx context.Context, subscriptionID string, body SubscriptionUpdatePaymentMethodParams, opts ...option.RequestOption) (res *SubscriptionUpdatePaymentMethodResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/update-payment-method", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response struct representing subscription details
type AddonCartResponseItem struct {
	AddonID  string                    `json:"addon_id,required"`
	Quantity int64                     `json:"quantity,required"`
	JSON     addonCartResponseItemJSON `json:"-"`
}

// addonCartResponseItemJSON contains the JSON metadata for the struct
// [AddonCartResponseItem]
type addonCartResponseItemJSON struct {
	AddonID     apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddonCartResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addonCartResponseItemJSON) RawJSON() string {
	return r.raw
}

type AttachAddonParam struct {
	AddonID  param.Field[string] `json:"addon_id,required"`
	Quantity param.Field[int64]  `json:"quantity,required"`
}

func (r AttachAddonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OnDemandSubscriptionParam struct {
	// If set as True, does not perform any charge and only authorizes payment method
	// details for future use.
	MandateOnly param.Field[bool] `json:"mandate_only,required"`
	// Whether adaptive currency fees should be included in the product_price (true) or
	// added on top (false). This field is ignored if adaptive pricing is not enabled
	// for the business.
	AdaptiveCurrencyFeesInclusive param.Field[bool] `json:"adaptive_currency_fees_inclusive"`
	// Optional currency of the product price. If not specified, defaults to the
	// currency of the product.
	ProductCurrency param.Field[Currency] `json:"product_currency"`
	// Optional product description override for billing and line items. If not
	// specified, the stored description of the product will be used.
	ProductDescription param.Field[string] `json:"product_description"`
	// Product price for the initial charge to customer If not specified the stored
	// price of the product will be used Represented in the lowest denomination of the
	// currency (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	ProductPrice param.Field[int64] `json:"product_price"`
}

func (r OnDemandSubscriptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response struct representing subscription details
type Subscription struct {
	// Addons associated with this subscription
	Addons []AddonCartResponseItem `json:"addons,required"`
	// Billing address details for payments
	Billing BillingAddress `json:"billing,required"`
	// Indicates if the subscription will cancel at the next billing date
	CancelAtNextBillingDate bool `json:"cancel_at_next_billing_date,required"`
	// Timestamp when the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Currency used for the subscription payments
	Currency Currency `json:"currency,required"`
	// Customer details associated with the subscription
	Customer CustomerLimitedDetails `json:"customer,required"`
	// Additional custom data associated with the subscription
	Metadata map[string]string `json:"metadata,required"`
	// Meters associated with this subscription (for usage-based billing)
	Meters []SubscriptionMeter `json:"meters,required"`
	// Timestamp of the next scheduled billing. Indicates the end of current billing
	// period
	NextBillingDate time.Time `json:"next_billing_date,required" format:"date-time"`
	// Wether the subscription is on-demand or not
	OnDemand bool `json:"on_demand,required"`
	// Number of payment frequency intervals
	PaymentFrequencyCount int64 `json:"payment_frequency_count,required"`
	// Time interval for payment frequency (e.g. month, year)
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,required"`
	// Timestamp of the last payment. Indicates the start of current billing period
	PreviousBillingDate time.Time `json:"previous_billing_date,required" format:"date-time"`
	// Identifier of the product associated with this subscription
	ProductID string `json:"product_id,required"`
	// Number of units/items included in the subscription
	Quantity int64 `json:"quantity,required"`
	// Amount charged before tax for each recurring payment in smallest currency unit
	// (e.g. cents)
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount,required"`
	// Current status of the subscription
	Status SubscriptionStatus `json:"status,required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id,required"`
	// Number of subscription period intervals
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// Time interval for the subscription period (e.g. month, year)
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval,required"`
	// Indicates if the recurring_pre_tax_amount is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,required"`
	// Number of days in the trial period (0 if no trial)
	TrialPeriodDays int64 `json:"trial_period_days,required"`
	// Cancelled timestamp if the subscription is cancelled
	CancelledAt time.Time `json:"cancelled_at,nullable" format:"date-time"`
	// Number of remaining discount cycles if discount is applied
	DiscountCyclesRemaining int64 `json:"discount_cycles_remaining,nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id,nullable"`
	// Timestamp when the subscription will expire
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// Saved payment method id used for recurring charges
	PaymentMethodID string `json:"payment_method_id,nullable"`
	// Tax identifier provided for this subscription (if applicable)
	TaxID string           `json:"tax_id,nullable"`
	JSON  subscriptionJSON `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	Addons                     apijson.Field
	Billing                    apijson.Field
	CancelAtNextBillingDate    apijson.Field
	CreatedAt                  apijson.Field
	Currency                   apijson.Field
	Customer                   apijson.Field
	Metadata                   apijson.Field
	Meters                     apijson.Field
	NextBillingDate            apijson.Field
	OnDemand                   apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	PreviousBillingDate        apijson.Field
	ProductID                  apijson.Field
	Quantity                   apijson.Field
	RecurringPreTaxAmount      apijson.Field
	Status                     apijson.Field
	SubscriptionID             apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	TaxInclusive               apijson.Field
	TrialPeriodDays            apijson.Field
	CancelledAt                apijson.Field
	DiscountCyclesRemaining    apijson.Field
	DiscountID                 apijson.Field
	ExpiresAt                  apijson.Field
	PaymentMethodID            apijson.Field
	TaxID                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

// Response struct representing usage-based meter cart details for a subscription
type SubscriptionMeter struct {
	Currency        Currency              `json:"currency,required"`
	FreeThreshold   int64                 `json:"free_threshold,required"`
	MeasurementUnit string                `json:"measurement_unit,required"`
	MeterID         string                `json:"meter_id,required"`
	Name            string                `json:"name,required"`
	PricePerUnit    string                `json:"price_per_unit,required"`
	Description     string                `json:"description,nullable"`
	JSON            subscriptionMeterJSON `json:"-"`
}

// subscriptionMeterJSON contains the JSON metadata for the struct
// [SubscriptionMeter]
type subscriptionMeterJSON struct {
	Currency        apijson.Field
	FreeThreshold   apijson.Field
	MeasurementUnit apijson.Field
	MeterID         apijson.Field
	Name            apijson.Field
	PricePerUnit    apijson.Field
	Description     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionMeter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionMeterJSON) RawJSON() string {
	return r.raw
}

type SubscriptionStatus string

const (
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusOnHold    SubscriptionStatus = "on_hold"
	SubscriptionStatusCancelled SubscriptionStatus = "cancelled"
	SubscriptionStatusFailed    SubscriptionStatus = "failed"
	SubscriptionStatusExpired   SubscriptionStatus = "expired"
)

func (r SubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionStatusPending, SubscriptionStatusActive, SubscriptionStatusOnHold, SubscriptionStatusCancelled, SubscriptionStatusFailed, SubscriptionStatusExpired:
		return true
	}
	return false
}

type TimeInterval string

const (
	TimeIntervalDay   TimeInterval = "Day"
	TimeIntervalWeek  TimeInterval = "Week"
	TimeIntervalMonth TimeInterval = "Month"
	TimeIntervalYear  TimeInterval = "Year"
)

func (r TimeInterval) IsKnown() bool {
	switch r {
	case TimeIntervalDay, TimeIntervalWeek, TimeIntervalMonth, TimeIntervalYear:
		return true
	}
	return false
}

type SubscriptionNewResponse struct {
	// Addons associated with this subscription
	Addons []AddonCartResponseItem `json:"addons,required"`
	// Customer details associated with this subscription
	Customer CustomerLimitedDetails `json:"customer,required"`
	// Additional metadata associated with the subscription
	Metadata map[string]string `json:"metadata,required"`
	// First payment id for the subscription
	PaymentID string `json:"payment_id,required"`
	// Tax will be added to the amount and charged to the customer on each billing
	// cycle
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount,required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id,required"`
	// Client secret used to load Dodo checkout SDK NOTE : Dodo checkout SDK will be
	// coming soon
	ClientSecret string `json:"client_secret,nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id,nullable"`
	// Expiry timestamp of the payment link
	ExpiresOn time.Time `json:"expires_on,nullable" format:"date-time"`
	// URL to checkout page
	PaymentLink string                      `json:"payment_link,nullable"`
	JSON        subscriptionNewResponseJSON `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	Addons                apijson.Field
	Customer              apijson.Field
	Metadata              apijson.Field
	PaymentID             apijson.Field
	RecurringPreTaxAmount apijson.Field
	SubscriptionID        apijson.Field
	ClientSecret          apijson.Field
	DiscountID            apijson.Field
	ExpiresOn             apijson.Field
	PaymentLink           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseJSON) RawJSON() string {
	return r.raw
}

// Response struct representing subscription details
type SubscriptionListResponse struct {
	// Billing address details for payments
	Billing BillingAddress `json:"billing,required"`
	// Indicates if the subscription will cancel at the next billing date
	CancelAtNextBillingDate bool `json:"cancel_at_next_billing_date,required"`
	// Timestamp when the subscription was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Currency used for the subscription payments
	Currency Currency `json:"currency,required"`
	// Customer details associated with the subscription
	Customer CustomerLimitedDetails `json:"customer,required"`
	// Additional custom data associated with the subscription
	Metadata map[string]string `json:"metadata,required"`
	// Timestamp of the next scheduled billing. Indicates the end of current billing
	// period
	NextBillingDate time.Time `json:"next_billing_date,required" format:"date-time"`
	// Wether the subscription is on-demand or not
	OnDemand bool `json:"on_demand,required"`
	// Number of payment frequency intervals
	PaymentFrequencyCount int64 `json:"payment_frequency_count,required"`
	// Time interval for payment frequency (e.g. month, year)
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,required"`
	// Timestamp of the last payment. Indicates the start of current billing period
	PreviousBillingDate time.Time `json:"previous_billing_date,required" format:"date-time"`
	// Identifier of the product associated with this subscription
	ProductID string `json:"product_id,required"`
	// Number of units/items included in the subscription
	Quantity int64 `json:"quantity,required"`
	// Amount charged before tax for each recurring payment in smallest currency unit
	// (e.g. cents)
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount,required"`
	// Current status of the subscription
	Status SubscriptionStatus `json:"status,required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id,required"`
	// Number of subscription period intervals
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// Time interval for the subscription period (e.g. month, year)
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval,required"`
	// Indicates if the recurring_pre_tax_amount is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,required"`
	// Number of days in the trial period (0 if no trial)
	TrialPeriodDays int64 `json:"trial_period_days,required"`
	// Cancelled timestamp if the subscription is cancelled
	CancelledAt time.Time `json:"cancelled_at,nullable" format:"date-time"`
	// Number of remaining discount cycles if discount is applied
	DiscountCyclesRemaining int64 `json:"discount_cycles_remaining,nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id,nullable"`
	// Saved payment method id used for recurring charges
	PaymentMethodID string `json:"payment_method_id,nullable"`
	// Tax identifier provided for this subscription (if applicable)
	TaxID string                       `json:"tax_id,nullable"`
	JSON  subscriptionListResponseJSON `json:"-"`
}

// subscriptionListResponseJSON contains the JSON metadata for the struct
// [SubscriptionListResponse]
type subscriptionListResponseJSON struct {
	Billing                    apijson.Field
	CancelAtNextBillingDate    apijson.Field
	CreatedAt                  apijson.Field
	Currency                   apijson.Field
	Customer                   apijson.Field
	Metadata                   apijson.Field
	NextBillingDate            apijson.Field
	OnDemand                   apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	PreviousBillingDate        apijson.Field
	ProductID                  apijson.Field
	Quantity                   apijson.Field
	RecurringPreTaxAmount      apijson.Field
	Status                     apijson.Field
	SubscriptionID             apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	TaxInclusive               apijson.Field
	TrialPeriodDays            apijson.Field
	CancelledAt                apijson.Field
	DiscountCyclesRemaining    apijson.Field
	DiscountID                 apijson.Field
	PaymentMethodID            apijson.Field
	TaxID                      apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChargeResponse struct {
	PaymentID string                         `json:"payment_id,required"`
	JSON      subscriptionChargeResponseJSON `json:"-"`
}

// subscriptionChargeResponseJSON contains the JSON metadata for the struct
// [SubscriptionChargeResponse]
type subscriptionChargeResponseJSON struct {
	PaymentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChargeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChargeResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetUsageHistoryResponse struct {
	// End date of the billing period
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// List of meters and their usage for this billing period
	Meters []SubscriptionGetUsageHistoryResponseMeter `json:"meters,required"`
	// Start date of the billing period
	StartDate time.Time                               `json:"start_date,required" format:"date-time"`
	JSON      subscriptionGetUsageHistoryResponseJSON `json:"-"`
}

// subscriptionGetUsageHistoryResponseJSON contains the JSON metadata for the
// struct [SubscriptionGetUsageHistoryResponse]
type subscriptionGetUsageHistoryResponseJSON struct {
	EndDate     apijson.Field
	Meters      apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetUsageHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetUsageHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetUsageHistoryResponseMeter struct {
	// Meter identifier
	ID string `json:"id,required"`
	// Chargeable units (after free threshold) as string for precision
	ChargeableUnits string `json:"chargeable_units,required"`
	// Total units consumed as string for precision
	ConsumedUnits string `json:"consumed_units,required"`
	// Currency for the price per unit
	Currency Currency `json:"currency,required"`
	// Free threshold units for this meter
	FreeThreshold int64 `json:"free_threshold,required"`
	// Meter name
	Name string `json:"name,required"`
	// Price per unit in string format for precision
	PricePerUnit string `json:"price_per_unit,required"`
	// Total price charged for this meter in smallest currency unit (cents)
	TotalPrice int64                                        `json:"total_price,required"`
	JSON       subscriptionGetUsageHistoryResponseMeterJSON `json:"-"`
}

// subscriptionGetUsageHistoryResponseMeterJSON contains the JSON metadata for the
// struct [SubscriptionGetUsageHistoryResponseMeter]
type subscriptionGetUsageHistoryResponseMeterJSON struct {
	ID              apijson.Field
	ChargeableUnits apijson.Field
	ConsumedUnits   apijson.Field
	Currency        apijson.Field
	FreeThreshold   apijson.Field
	Name            apijson.Field
	PricePerUnit    apijson.Field
	TotalPrice      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionGetUsageHistoryResponseMeter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetUsageHistoryResponseMeterJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUpdatePaymentMethodResponse struct {
	ClientSecret string                                      `json:"client_secret,nullable"`
	ExpiresOn    time.Time                                   `json:"expires_on,nullable" format:"date-time"`
	PaymentID    string                                      `json:"payment_id,nullable"`
	PaymentLink  string                                      `json:"payment_link,nullable"`
	JSON         subscriptionUpdatePaymentMethodResponseJSON `json:"-"`
}

// subscriptionUpdatePaymentMethodResponseJSON contains the JSON metadata for the
// struct [SubscriptionUpdatePaymentMethodResponse]
type subscriptionUpdatePaymentMethodResponseJSON struct {
	ClientSecret apijson.Field
	ExpiresOn    apijson.Field
	PaymentID    apijson.Field
	PaymentLink  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionUpdatePaymentMethodResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdatePaymentMethodResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewParams struct {
	// Billing address information for the subscription
	Billing param.Field[BillingAddressParam] `json:"billing,required"`
	// Customer details for the subscription
	Customer param.Field[CustomerRequestUnionParam] `json:"customer,required"`
	// Unique identifier of the product to subscribe to
	ProductID param.Field[string] `json:"product_id,required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity param.Field[int64] `json:"quantity,required"`
	// Attach addons to this subscription
	Addons param.Field[[]AttachAddonParam] `json:"addons"`
	// List of payment methods allowed during checkout.
	//
	// Customers will **never** see payment methods that are **not** in this list.
	// However, adding a method here **does not guarantee** customers will see it.
	// Availability still depends on other factors (e.g., customer location, merchant
	// settings).
	AllowedPaymentMethodTypes param.Field[[]PaymentMethodTypes] `json:"allowed_payment_method_types"`
	// Fix the currency in which the end customer is billed. If Dodo Payments cannot
	// support that currency for this transaction, it will not proceed
	BillingCurrency param.Field[Currency] `json:"billing_currency"`
	// Discount Code to apply to the subscription
	DiscountCode param.Field[string] `json:"discount_code"`
	// Override merchant default 3DS behaviour for this subscription
	Force3DS param.Field[bool] `json:"force_3ds"`
	// Additional metadata for the subscription Defaults to empty if not specified
	Metadata param.Field[map[string]string]         `json:"metadata"`
	OnDemand param.Field[OnDemandSubscriptionParam] `json:"on_demand"`
	// If true, generates a payment link. Defaults to false if not specified.
	PaymentLink param.Field[bool] `json:"payment_link"`
	// Optional URL to redirect after successful subscription creation
	ReturnURL param.Field[string] `json:"return_url"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Field[bool] `json:"show_saved_payment_methods"`
	// Tax ID in case the payment is B2B. If tax id validation fails the payment
	// creation will fail
	TaxID param.Field[string] `json:"tax_id"`
	// Optional trial period in days If specified, this value overrides the trial
	// period set in the product's price Must be between 0 and 10000 days
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParams struct {
	Billing param.Field[BillingAddressParam] `json:"billing"`
	// When set, the subscription will remain active until the end of billing period
	CancelAtNextBillingDate param.Field[bool]                                    `json:"cancel_at_next_billing_date"`
	CustomerName            param.Field[string]                                  `json:"customer_name"`
	DisableOnDemand         param.Field[SubscriptionUpdateParamsDisableOnDemand] `json:"disable_on_demand"`
	Metadata                param.Field[map[string]string]                       `json:"metadata"`
	NextBillingDate         param.Field[time.Time]                               `json:"next_billing_date" format:"date-time"`
	Status                  param.Field[SubscriptionStatus]                      `json:"status"`
	TaxID                   param.Field[string]                                  `json:"tax_id"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParamsDisableOnDemand struct {
	NextBillingDate param.Field[time.Time] `json:"next_billing_date,required" format:"date-time"`
}

func (r SubscriptionUpdateParamsDisableOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionListParams struct {
	// filter by Brand id
	BrandID param.Field[string] `query:"brand_id"`
	// Get events after this created time
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get events created before this time
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Filter by customer id
	CustomerID param.Field[string] `query:"customer_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by status
	Status param.Field[SubscriptionListParamsStatus] `query:"status"`
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by status
type SubscriptionListParamsStatus string

const (
	SubscriptionListParamsStatusPending   SubscriptionListParamsStatus = "pending"
	SubscriptionListParamsStatusActive    SubscriptionListParamsStatus = "active"
	SubscriptionListParamsStatusOnHold    SubscriptionListParamsStatus = "on_hold"
	SubscriptionListParamsStatusCancelled SubscriptionListParamsStatus = "cancelled"
	SubscriptionListParamsStatusFailed    SubscriptionListParamsStatus = "failed"
	SubscriptionListParamsStatusExpired   SubscriptionListParamsStatus = "expired"
)

func (r SubscriptionListParamsStatus) IsKnown() bool {
	switch r {
	case SubscriptionListParamsStatusPending, SubscriptionListParamsStatusActive, SubscriptionListParamsStatusOnHold, SubscriptionListParamsStatusCancelled, SubscriptionListParamsStatusFailed, SubscriptionListParamsStatusExpired:
		return true
	}
	return false
}

type SubscriptionChangePlanParams struct {
	// Unique identifier of the product to subscribe to
	ProductID param.Field[string] `json:"product_id,required"`
	// Proration Billing Mode
	ProrationBillingMode param.Field[SubscriptionChangePlanParamsProrationBillingMode] `json:"proration_billing_mode,required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity param.Field[int64] `json:"quantity,required"`
	// Addons for the new plan. Note : Leaving this empty would remove any existing
	// addons
	Addons param.Field[[]AttachAddonParam] `json:"addons"`
}

func (r SubscriptionChangePlanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Proration Billing Mode
type SubscriptionChangePlanParamsProrationBillingMode string

const (
	SubscriptionChangePlanParamsProrationBillingModeProratedImmediately   SubscriptionChangePlanParamsProrationBillingMode = "prorated_immediately"
	SubscriptionChangePlanParamsProrationBillingModeFullImmediately       SubscriptionChangePlanParamsProrationBillingMode = "full_immediately"
	SubscriptionChangePlanParamsProrationBillingModeDifferenceImmediately SubscriptionChangePlanParamsProrationBillingMode = "difference_immediately"
)

func (r SubscriptionChangePlanParamsProrationBillingMode) IsKnown() bool {
	switch r {
	case SubscriptionChangePlanParamsProrationBillingModeProratedImmediately, SubscriptionChangePlanParamsProrationBillingModeFullImmediately, SubscriptionChangePlanParamsProrationBillingModeDifferenceImmediately:
		return true
	}
	return false
}

type SubscriptionChargeParams struct {
	// The product price. Represented in the lowest denomination of the currency (e.g.,
	// cents for USD). For example, to charge $1.00, pass `100`.
	ProductPrice param.Field[int64] `json:"product_price,required"`
	// Whether adaptive currency fees should be included in the product_price (true) or
	// added on top (false). This field is ignored if adaptive pricing is not enabled
	// for the business.
	AdaptiveCurrencyFeesInclusive param.Field[bool] `json:"adaptive_currency_fees_inclusive"`
	// Specify how customer balance is used for the payment
	CustomerBalanceConfig param.Field[SubscriptionChargeParamsCustomerBalanceConfig] `json:"customer_balance_config"`
	// Metadata for the payment. If not passed, the metadata of the subscription will
	// be taken
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Optional currency of the product price. If not specified, defaults to the
	// currency of the product.
	ProductCurrency param.Field[Currency] `json:"product_currency"`
	// Optional product description override for billing and line items. If not
	// specified, the stored description of the product will be used.
	ProductDescription param.Field[string] `json:"product_description"`
}

func (r SubscriptionChargeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how customer balance is used for the payment
type SubscriptionChargeParamsCustomerBalanceConfig struct {
	// Allows Customer Credit to be purchased to settle payments
	AllowCustomerCreditsPurchase param.Field[bool] `json:"allow_customer_credits_purchase"`
	// Allows Customer Credit Balance to be used to settle payments
	AllowCustomerCreditsUsage param.Field[bool] `json:"allow_customer_credits_usage"`
}

func (r SubscriptionChargeParamsCustomerBalanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionGetUsageHistoryParams struct {
	// Filter by end date (inclusive)
	EndDate param.Field[time.Time] `query:"end_date" format:"date-time"`
	// Filter by specific meter ID
	MeterID param.Field[string] `query:"meter_id"`
	// Page number (default: 0)
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size (default: 10, max: 100)
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by start date (inclusive)
	StartDate param.Field[time.Time] `query:"start_date" format:"date-time"`
}

// URLQuery serializes [SubscriptionGetUsageHistoryParams]'s query parameters as
// `url.Values`.
func (r SubscriptionGetUsageHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubscriptionUpdatePaymentMethodParams struct {
	Body SubscriptionUpdatePaymentMethodParamsBodyUnion `json:"body,required"`
}

func (r SubscriptionUpdatePaymentMethodParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SubscriptionUpdatePaymentMethodParamsBody struct {
	Type            param.Field[SubscriptionUpdatePaymentMethodParamsBodyType] `json:"type,required"`
	PaymentMethodID param.Field[string]                                        `json:"payment_method_id"`
	ReturnURL       param.Field[string]                                        `json:"return_url"`
}

func (r SubscriptionUpdatePaymentMethodParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionUpdatePaymentMethodParamsBody) implementsSubscriptionUpdatePaymentMethodParamsBodyUnion() {
}

// Satisfied by [SubscriptionUpdatePaymentMethodParamsBodyNew],
// [SubscriptionUpdatePaymentMethodParamsBodyExisting],
// [SubscriptionUpdatePaymentMethodParamsBody].
type SubscriptionUpdatePaymentMethodParamsBodyUnion interface {
	implementsSubscriptionUpdatePaymentMethodParamsBodyUnion()
}

type SubscriptionUpdatePaymentMethodParamsBodyNew struct {
	Type      param.Field[SubscriptionUpdatePaymentMethodParamsBodyNewType] `json:"type,required"`
	ReturnURL param.Field[string]                                           `json:"return_url"`
}

func (r SubscriptionUpdatePaymentMethodParamsBodyNew) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionUpdatePaymentMethodParamsBodyNew) implementsSubscriptionUpdatePaymentMethodParamsBodyUnion() {
}

type SubscriptionUpdatePaymentMethodParamsBodyNewType string

const (
	SubscriptionUpdatePaymentMethodParamsBodyNewTypeNew SubscriptionUpdatePaymentMethodParamsBodyNewType = "new"
)

func (r SubscriptionUpdatePaymentMethodParamsBodyNewType) IsKnown() bool {
	switch r {
	case SubscriptionUpdatePaymentMethodParamsBodyNewTypeNew:
		return true
	}
	return false
}

type SubscriptionUpdatePaymentMethodParamsBodyExisting struct {
	PaymentMethodID param.Field[string]                                                `json:"payment_method_id,required"`
	Type            param.Field[SubscriptionUpdatePaymentMethodParamsBodyExistingType] `json:"type,required"`
}

func (r SubscriptionUpdatePaymentMethodParamsBodyExisting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionUpdatePaymentMethodParamsBodyExisting) implementsSubscriptionUpdatePaymentMethodParamsBodyUnion() {
}

type SubscriptionUpdatePaymentMethodParamsBodyExistingType string

const (
	SubscriptionUpdatePaymentMethodParamsBodyExistingTypeExisting SubscriptionUpdatePaymentMethodParamsBodyExistingType = "existing"
)

func (r SubscriptionUpdatePaymentMethodParamsBodyExistingType) IsKnown() bool {
	switch r {
	case SubscriptionUpdatePaymentMethodParamsBodyExistingTypeExisting:
		return true
	}
	return false
}

type SubscriptionUpdatePaymentMethodParamsBodyType string

const (
	SubscriptionUpdatePaymentMethodParamsBodyTypeNew      SubscriptionUpdatePaymentMethodParamsBodyType = "new"
	SubscriptionUpdatePaymentMethodParamsBodyTypeExisting SubscriptionUpdatePaymentMethodParamsBodyType = "existing"
)

func (r SubscriptionUpdatePaymentMethodParamsBodyType) IsKnown() bool {
	switch r {
	case SubscriptionUpdatePaymentMethodParamsBodyTypeNew, SubscriptionUpdatePaymentMethodParamsBodyTypeExisting:
		return true
	}
	return false
}
