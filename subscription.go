// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	"github.com/tidwall/gjson"
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

// Deprecated: deprecated
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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

func (r *SubscriptionService) PreviewChangePlan(ctx context.Context, subscriptionID string, body SubscriptionPreviewChangePlanParams, opts ...option.RequestOption) (res *SubscriptionPreviewChangePlanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/change-plan/preview", subscriptionID)
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
	AddonID  string                    `json:"addon_id" api:"required"`
	Quantity int64                     `json:"quantity" api:"required"`
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
	AddonID  param.Field[string] `json:"addon_id" api:"required"`
	Quantity param.Field[int64]  `json:"quantity" api:"required"`
}

func (r AttachAddonParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OnDemandSubscriptionParam struct {
	// If set as True, does not perform any charge and only authorizes payment method
	// details for future use.
	MandateOnly param.Field[bool] `json:"mandate_only" api:"required"`
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
	Addons []AddonCartResponseItem `json:"addons" api:"required"`
	// Billing address details for payments
	Billing BillingAddress `json:"billing" api:"required"`
	// Indicates if the subscription will cancel at the next billing date
	CancelAtNextBillingDate bool `json:"cancel_at_next_billing_date" api:"required"`
	// Timestamp when the subscription was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Credit entitlement cart settings for this subscription
	CreditEntitlementCart []SubscriptionCreditEntitlementCart `json:"credit_entitlement_cart" api:"required"`
	// Currency used for the subscription payments
	Currency Currency `json:"currency" api:"required"`
	// Customer details associated with the subscription
	Customer CustomerLimitedDetails `json:"customer" api:"required"`
	// Additional custom data associated with the subscription
	Metadata map[string]string `json:"metadata" api:"required"`
	// Meter credit entitlement cart settings for this subscription
	MeterCreditEntitlementCart []SubscriptionMeterCreditEntitlementCart `json:"meter_credit_entitlement_cart" api:"required"`
	// Meters associated with this subscription (for usage-based billing)
	Meters []SubscriptionMeter `json:"meters" api:"required"`
	// Timestamp of the next scheduled billing. Indicates the end of current billing
	// period
	NextBillingDate time.Time `json:"next_billing_date" api:"required" format:"date-time"`
	// Wether the subscription is on-demand or not
	OnDemand bool `json:"on_demand" api:"required"`
	// Number of payment frequency intervals
	PaymentFrequencyCount int64 `json:"payment_frequency_count" api:"required"`
	// Time interval for payment frequency (e.g. month, year)
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval" api:"required"`
	// Timestamp of the last payment. Indicates the start of current billing period
	PreviousBillingDate time.Time `json:"previous_billing_date" api:"required" format:"date-time"`
	// Identifier of the product associated with this subscription
	ProductID string `json:"product_id" api:"required"`
	// Number of units/items included in the subscription
	Quantity int64 `json:"quantity" api:"required"`
	// Amount charged before tax for each recurring payment in smallest currency unit
	// (e.g. cents)
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount" api:"required"`
	// Current status of the subscription
	Status SubscriptionStatus `json:"status" api:"required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id" api:"required"`
	// Number of subscription period intervals
	SubscriptionPeriodCount int64 `json:"subscription_period_count" api:"required"`
	// Time interval for the subscription period (e.g. month, year)
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval" api:"required"`
	// Indicates if the recurring_pre_tax_amount is tax inclusive
	TaxInclusive bool `json:"tax_inclusive" api:"required"`
	// Number of days in the trial period (0 if no trial)
	TrialPeriodDays int64 `json:"trial_period_days" api:"required"`
	// Cancelled timestamp if the subscription is cancelled
	CancelledAt time.Time `json:"cancelled_at" api:"nullable" format:"date-time"`
	// Customer's responses to custom fields collected during checkout
	CustomFieldResponses []SubscriptionCustomFieldResponse `json:"custom_field_responses" api:"nullable"`
	// Number of remaining discount cycles if discount is applied
	DiscountCyclesRemaining int64 `json:"discount_cycles_remaining" api:"nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id" api:"nullable"`
	// Timestamp when the subscription will expire
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// Saved payment method id used for recurring charges
	PaymentMethodID string `json:"payment_method_id" api:"nullable"`
	// Tax identifier provided for this subscription (if applicable)
	TaxID string           `json:"tax_id" api:"nullable"`
	JSON  subscriptionJSON `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	Addons                     apijson.Field
	Billing                    apijson.Field
	CancelAtNextBillingDate    apijson.Field
	CreatedAt                  apijson.Field
	CreditEntitlementCart      apijson.Field
	Currency                   apijson.Field
	Customer                   apijson.Field
	Metadata                   apijson.Field
	MeterCreditEntitlementCart apijson.Field
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
	CustomFieldResponses       apijson.Field
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

// Response struct representing credit entitlement cart details for a subscription
type SubscriptionCreditEntitlementCart struct {
	CreditEntitlementID   string `json:"credit_entitlement_id" api:"required"`
	CreditEntitlementName string `json:"credit_entitlement_name" api:"required"`
	CreditsAmount         string `json:"credits_amount" api:"required"`
	// Customer's current overage balance for this entitlement
	OverageBalance         string `json:"overage_balance" api:"required"`
	OverageChargeAtBilling bool   `json:"overage_charge_at_billing" api:"required"`
	OverageEnabled         bool   `json:"overage_enabled" api:"required"`
	ProductID              string `json:"product_id" api:"required"`
	// Customer's current remaining credit balance for this entitlement
	RemainingBalance string `json:"remaining_balance" api:"required"`
	RolloverEnabled  bool   `json:"rollover_enabled" api:"required"`
	// Unit label for the credit entitlement (e.g., "API Calls", "Tokens")
	Unit                       string                                `json:"unit" api:"required"`
	ExpiresAfterDays           int64                                 `json:"expires_after_days" api:"nullable"`
	LowBalanceThresholdPercent int64                                 `json:"low_balance_threshold_percent" api:"nullable"`
	MaxRolloverCount           int64                                 `json:"max_rollover_count" api:"nullable"`
	OverageLimit               string                                `json:"overage_limit" api:"nullable"`
	RolloverPercentage         int64                                 `json:"rollover_percentage" api:"nullable"`
	RolloverTimeframeCount     int64                                 `json:"rollover_timeframe_count" api:"nullable"`
	RolloverTimeframeInterval  TimeInterval                          `json:"rollover_timeframe_interval" api:"nullable"`
	JSON                       subscriptionCreditEntitlementCartJSON `json:"-"`
}

// subscriptionCreditEntitlementCartJSON contains the JSON metadata for the struct
// [SubscriptionCreditEntitlementCart]
type subscriptionCreditEntitlementCartJSON struct {
	CreditEntitlementID        apijson.Field
	CreditEntitlementName      apijson.Field
	CreditsAmount              apijson.Field
	OverageBalance             apijson.Field
	OverageChargeAtBilling     apijson.Field
	OverageEnabled             apijson.Field
	ProductID                  apijson.Field
	RemainingBalance           apijson.Field
	RolloverEnabled            apijson.Field
	Unit                       apijson.Field
	ExpiresAfterDays           apijson.Field
	LowBalanceThresholdPercent apijson.Field
	MaxRolloverCount           apijson.Field
	OverageLimit               apijson.Field
	RolloverPercentage         apijson.Field
	RolloverTimeframeCount     apijson.Field
	RolloverTimeframeInterval  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SubscriptionCreditEntitlementCart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCreditEntitlementCartJSON) RawJSON() string {
	return r.raw
}

// Response struct representing meter-credit entitlement mapping cart details for a
// subscription
type SubscriptionMeterCreditEntitlementCart struct {
	CreditEntitlementID string                                     `json:"credit_entitlement_id" api:"required"`
	MeterID             string                                     `json:"meter_id" api:"required"`
	MeterName           string                                     `json:"meter_name" api:"required"`
	MeterUnitsPerCredit string                                     `json:"meter_units_per_credit" api:"required"`
	ProductID           string                                     `json:"product_id" api:"required"`
	JSON                subscriptionMeterCreditEntitlementCartJSON `json:"-"`
}

// subscriptionMeterCreditEntitlementCartJSON contains the JSON metadata for the
// struct [SubscriptionMeterCreditEntitlementCart]
type subscriptionMeterCreditEntitlementCartJSON struct {
	CreditEntitlementID apijson.Field
	MeterID             apijson.Field
	MeterName           apijson.Field
	MeterUnitsPerCredit apijson.Field
	ProductID           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SubscriptionMeterCreditEntitlementCart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionMeterCreditEntitlementCartJSON) RawJSON() string {
	return r.raw
}

// Response struct representing usage-based meter cart details for a subscription
type SubscriptionMeter struct {
	Currency        Currency              `json:"currency" api:"required"`
	FreeThreshold   int64                 `json:"free_threshold" api:"required"`
	MeasurementUnit string                `json:"measurement_unit" api:"required"`
	MeterID         string                `json:"meter_id" api:"required"`
	Name            string                `json:"name" api:"required"`
	PricePerUnit    string                `json:"price_per_unit" api:"required"`
	Description     string                `json:"description" api:"nullable"`
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

// Customer's response to a custom field
type SubscriptionCustomFieldResponse struct {
	// Key matching the custom field definition
	Key string `json:"key" api:"required"`
	// Value provided by customer
	Value string                              `json:"value" api:"required"`
	JSON  subscriptionCustomFieldResponseJSON `json:"-"`
}

// subscriptionCustomFieldResponseJSON contains the JSON metadata for the struct
// [SubscriptionCustomFieldResponse]
type subscriptionCustomFieldResponseJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCustomFieldResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCustomFieldResponseJSON) RawJSON() string {
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
	Addons []AddonCartResponseItem `json:"addons" api:"required"`
	// Customer details associated with this subscription
	Customer CustomerLimitedDetails `json:"customer" api:"required"`
	// Additional metadata associated with the subscription
	Metadata map[string]string `json:"metadata" api:"required"`
	// First payment id for the subscription
	PaymentID string `json:"payment_id" api:"required"`
	// Tax will be added to the amount and charged to the customer on each billing
	// cycle
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount" api:"required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id" api:"required"`
	// Client secret used to load Dodo checkout SDK NOTE : Dodo checkout SDK will be
	// coming soon
	ClientSecret string `json:"client_secret" api:"nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id" api:"nullable"`
	// Expiry timestamp of the payment link
	ExpiresOn time.Time `json:"expires_on" api:"nullable" format:"date-time"`
	// One time products associated with the purchase of subscription
	OneTimeProductCart []SubscriptionNewResponseOneTimeProductCart `json:"one_time_product_cart" api:"nullable"`
	// URL to checkout page
	PaymentLink string                      `json:"payment_link" api:"nullable"`
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
	OneTimeProductCart    apijson.Field
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

type SubscriptionNewResponseOneTimeProductCart struct {
	ProductID string                                        `json:"product_id" api:"required"`
	Quantity  int64                                         `json:"quantity" api:"required"`
	JSON      subscriptionNewResponseOneTimeProductCartJSON `json:"-"`
}

// subscriptionNewResponseOneTimeProductCartJSON contains the JSON metadata for the
// struct [SubscriptionNewResponseOneTimeProductCart]
type subscriptionNewResponseOneTimeProductCartJSON struct {
	ProductID   apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseOneTimeProductCart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseOneTimeProductCartJSON) RawJSON() string {
	return r.raw
}

// Response struct representing subscription details
type SubscriptionListResponse struct {
	// Billing address details for payments
	Billing BillingAddress `json:"billing" api:"required"`
	// Indicates if the subscription will cancel at the next billing date
	CancelAtNextBillingDate bool `json:"cancel_at_next_billing_date" api:"required"`
	// Timestamp when the subscription was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Currency used for the subscription payments
	Currency Currency `json:"currency" api:"required"`
	// Customer details associated with the subscription
	Customer CustomerLimitedDetails `json:"customer" api:"required"`
	// Additional custom data associated with the subscription
	Metadata map[string]string `json:"metadata" api:"required"`
	// Timestamp of the next scheduled billing. Indicates the end of current billing
	// period
	NextBillingDate time.Time `json:"next_billing_date" api:"required" format:"date-time"`
	// Wether the subscription is on-demand or not
	OnDemand bool `json:"on_demand" api:"required"`
	// Number of payment frequency intervals
	PaymentFrequencyCount int64 `json:"payment_frequency_count" api:"required"`
	// Time interval for payment frequency (e.g. month, year)
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval" api:"required"`
	// Timestamp of the last payment. Indicates the start of current billing period
	PreviousBillingDate time.Time `json:"previous_billing_date" api:"required" format:"date-time"`
	// Identifier of the product associated with this subscription
	ProductID string `json:"product_id" api:"required"`
	// Number of units/items included in the subscription
	Quantity int64 `json:"quantity" api:"required"`
	// Amount charged before tax for each recurring payment in smallest currency unit
	// (e.g. cents)
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount" api:"required"`
	// Current status of the subscription
	Status SubscriptionStatus `json:"status" api:"required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id" api:"required"`
	// Number of subscription period intervals
	SubscriptionPeriodCount int64 `json:"subscription_period_count" api:"required"`
	// Time interval for the subscription period (e.g. month, year)
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval" api:"required"`
	// Indicates if the recurring_pre_tax_amount is tax inclusive
	TaxInclusive bool `json:"tax_inclusive" api:"required"`
	// Number of days in the trial period (0 if no trial)
	TrialPeriodDays int64 `json:"trial_period_days" api:"required"`
	// Cancelled timestamp if the subscription is cancelled
	CancelledAt time.Time `json:"cancelled_at" api:"nullable" format:"date-time"`
	// Number of remaining discount cycles if discount is applied
	DiscountCyclesRemaining int64 `json:"discount_cycles_remaining" api:"nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id" api:"nullable"`
	// Saved payment method id used for recurring charges
	PaymentMethodID string `json:"payment_method_id" api:"nullable"`
	// Name of the product associated with this subscription
	ProductName string `json:"product_name" api:"nullable"`
	// Tax identifier provided for this subscription (if applicable)
	TaxID string                       `json:"tax_id" api:"nullable"`
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
	ProductName                apijson.Field
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
	PaymentID string                         `json:"payment_id" api:"required"`
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

type SubscriptionPreviewChangePlanResponse struct {
	ImmediateCharge SubscriptionPreviewChangePlanResponseImmediateCharge `json:"immediate_charge" api:"required"`
	// Response struct representing subscription details
	NewPlan Subscription                              `json:"new_plan" api:"required"`
	JSON    subscriptionPreviewChangePlanResponseJSON `json:"-"`
}

// subscriptionPreviewChangePlanResponseJSON contains the JSON metadata for the
// struct [SubscriptionPreviewChangePlanResponse]
type subscriptionPreviewChangePlanResponseJSON struct {
	ImmediateCharge apijson.Field
	NewPlan         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionPreviewChangePlanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPreviewChangePlanResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionPreviewChangePlanResponseImmediateCharge struct {
	LineItems []SubscriptionPreviewChangePlanResponseImmediateChargeLineItem `json:"line_items" api:"required"`
	Summary   SubscriptionPreviewChangePlanResponseImmediateChargeSummary    `json:"summary" api:"required"`
	JSON      subscriptionPreviewChangePlanResponseImmediateChargeJSON       `json:"-"`
}

// subscriptionPreviewChangePlanResponseImmediateChargeJSON contains the JSON
// metadata for the struct [SubscriptionPreviewChangePlanResponseImmediateCharge]
type subscriptionPreviewChangePlanResponseImmediateChargeJSON struct {
	LineItems   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionPreviewChangePlanResponseImmediateCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPreviewChangePlanResponseImmediateChargeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItem struct {
	ID              string                                                            `json:"id" api:"required"`
	Currency        Currency                                                          `json:"currency" api:"required"`
	TaxInclusive    bool                                                              `json:"tax_inclusive" api:"required"`
	Type            SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsType `json:"type" api:"required"`
	ChargeableUnits string                                                            `json:"chargeable_units"`
	Description     string                                                            `json:"description" api:"nullable"`
	FreeThreshold   int64                                                             `json:"free_threshold"`
	Name            string                                                            `json:"name" api:"nullable"`
	PricePerUnit    string                                                            `json:"price_per_unit"`
	ProductID       string                                                            `json:"product_id"`
	ProrationFactor float64                                                           `json:"proration_factor"`
	Quantity        int64                                                             `json:"quantity"`
	Subtotal        int64                                                             `json:"subtotal"`
	Tax             int64                                                             `json:"tax" api:"nullable"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory   TaxCategory                                                      `json:"tax_category"`
	TaxRate       float64                                                          `json:"tax_rate" api:"nullable"`
	UnitPrice     int64                                                            `json:"unit_price"`
	UnitsConsumed string                                                           `json:"units_consumed"`
	JSON          subscriptionPreviewChangePlanResponseImmediateChargeLineItemJSON `json:"-"`
	union         SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsUnion
}

// subscriptionPreviewChangePlanResponseImmediateChargeLineItemJSON contains the
// JSON metadata for the struct
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItem]
type subscriptionPreviewChangePlanResponseImmediateChargeLineItemJSON struct {
	ID              apijson.Field
	Currency        apijson.Field
	TaxInclusive    apijson.Field
	Type            apijson.Field
	ChargeableUnits apijson.Field
	Description     apijson.Field
	FreeThreshold   apijson.Field
	Name            apijson.Field
	PricePerUnit    apijson.Field
	ProductID       apijson.Field
	ProrationFactor apijson.Field
	Quantity        apijson.Field
	Subtotal        apijson.Field
	Tax             apijson.Field
	TaxCategory     apijson.Field
	TaxRate         apijson.Field
	UnitPrice       apijson.Field
	UnitsConsumed   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r subscriptionPreviewChangePlanResponseImmediateChargeLineItemJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItem) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionPreviewChangePlanResponseImmediateChargeLineItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscription],
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddon],
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeter].
func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItem) AsUnion() SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscription],
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddon] or
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeter].
type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsUnion interface {
	implementsSubscriptionPreviewChangePlanResponseImmediateChargeLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscription{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeter{}),
		},
	)
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscription struct {
	ID              string                                                                        `json:"id" api:"required"`
	Currency        Currency                                                                      `json:"currency" api:"required"`
	ProductID       string                                                                        `json:"product_id" api:"required"`
	ProrationFactor float64                                                                       `json:"proration_factor" api:"required"`
	Quantity        int64                                                                         `json:"quantity" api:"required"`
	TaxInclusive    bool                                                                          `json:"tax_inclusive" api:"required"`
	Type            SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionType `json:"type" api:"required"`
	UnitPrice       int64                                                                         `json:"unit_price" api:"required"`
	Description     string                                                                        `json:"description" api:"nullable"`
	Name            string                                                                        `json:"name" api:"nullable"`
	Tax             int64                                                                         `json:"tax" api:"nullable"`
	TaxRate         float64                                                                       `json:"tax_rate" api:"nullable"`
	JSON            subscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionJSON `json:"-"`
}

// subscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionJSON
// contains the JSON metadata for the struct
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscription]
type subscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionJSON struct {
	ID              apijson.Field
	Currency        apijson.Field
	ProductID       apijson.Field
	ProrationFactor apijson.Field
	Quantity        apijson.Field
	TaxInclusive    apijson.Field
	Type            apijson.Field
	UnitPrice       apijson.Field
	Description     apijson.Field
	Name            apijson.Field
	Tax             apijson.Field
	TaxRate         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscription) implementsSubscriptionPreviewChangePlanResponseImmediateChargeLineItem() {
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionType string

const (
	SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionTypeSubscription SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionType = "subscription"
)

func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionType) IsKnown() bool {
	switch r {
	case SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsSubscriptionTypeSubscription:
		return true
	}
	return false
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddon struct {
	ID              string   `json:"id" api:"required"`
	Currency        Currency `json:"currency" api:"required"`
	Name            string   `json:"name" api:"required"`
	ProrationFactor float64  `json:"proration_factor" api:"required"`
	Quantity        int64    `json:"quantity" api:"required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory  TaxCategory                                                            `json:"tax_category" api:"required"`
	TaxInclusive bool                                                                   `json:"tax_inclusive" api:"required"`
	TaxRate      float64                                                                `json:"tax_rate" api:"required"`
	Type         SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonType `json:"type" api:"required"`
	UnitPrice    int64                                                                  `json:"unit_price" api:"required"`
	Description  string                                                                 `json:"description" api:"nullable"`
	Tax          int64                                                                  `json:"tax" api:"nullable"`
	JSON         subscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonJSON `json:"-"`
}

// subscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonJSON contains
// the JSON metadata for the struct
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddon]
type subscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonJSON struct {
	ID              apijson.Field
	Currency        apijson.Field
	Name            apijson.Field
	ProrationFactor apijson.Field
	Quantity        apijson.Field
	TaxCategory     apijson.Field
	TaxInclusive    apijson.Field
	TaxRate         apijson.Field
	Type            apijson.Field
	UnitPrice       apijson.Field
	Description     apijson.Field
	Tax             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddon) implementsSubscriptionPreviewChangePlanResponseImmediateChargeLineItem() {
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonType string

const (
	SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonTypeAddon SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonType = "addon"
)

func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonType) IsKnown() bool {
	switch r {
	case SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsAddonTypeAddon:
		return true
	}
	return false
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeter struct {
	ID              string                                                                 `json:"id" api:"required"`
	ChargeableUnits string                                                                 `json:"chargeable_units" api:"required"`
	Currency        Currency                                                               `json:"currency" api:"required"`
	FreeThreshold   int64                                                                  `json:"free_threshold" api:"required"`
	Name            string                                                                 `json:"name" api:"required"`
	PricePerUnit    string                                                                 `json:"price_per_unit" api:"required"`
	Subtotal        int64                                                                  `json:"subtotal" api:"required"`
	TaxInclusive    bool                                                                   `json:"tax_inclusive" api:"required"`
	TaxRate         float64                                                                `json:"tax_rate" api:"required"`
	Type            SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterType `json:"type" api:"required"`
	UnitsConsumed   string                                                                 `json:"units_consumed" api:"required"`
	Description     string                                                                 `json:"description" api:"nullable"`
	Tax             int64                                                                  `json:"tax" api:"nullable"`
	JSON            subscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterJSON `json:"-"`
}

// subscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterJSON contains
// the JSON metadata for the struct
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeter]
type subscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterJSON struct {
	ID              apijson.Field
	ChargeableUnits apijson.Field
	Currency        apijson.Field
	FreeThreshold   apijson.Field
	Name            apijson.Field
	PricePerUnit    apijson.Field
	Subtotal        apijson.Field
	TaxInclusive    apijson.Field
	TaxRate         apijson.Field
	Type            apijson.Field
	UnitsConsumed   apijson.Field
	Description     apijson.Field
	Tax             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeter) implementsSubscriptionPreviewChangePlanResponseImmediateChargeLineItem() {
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterType string

const (
	SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterTypeMeter SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterType = "meter"
)

func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterType) IsKnown() bool {
	switch r {
	case SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsMeterTypeMeter:
		return true
	}
	return false
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsType string

const (
	SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsTypeSubscription SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsType = "subscription"
	SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsTypeAddon        SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsType = "addon"
	SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsTypeMeter        SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsType = "meter"
)

func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsType) IsKnown() bool {
	switch r {
	case SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsTypeSubscription, SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsTypeAddon, SubscriptionPreviewChangePlanResponseImmediateChargeLineItemsTypeMeter:
		return true
	}
	return false
}

type SubscriptionPreviewChangePlanResponseImmediateChargeSummary struct {
	Currency           Currency                                                        `json:"currency" api:"required"`
	CustomerCredits    int64                                                           `json:"customer_credits" api:"required"`
	SettlementAmount   int64                                                           `json:"settlement_amount" api:"required"`
	SettlementCurrency Currency                                                        `json:"settlement_currency" api:"required"`
	TotalAmount        int64                                                           `json:"total_amount" api:"required"`
	SettlementTax      int64                                                           `json:"settlement_tax" api:"nullable"`
	Tax                int64                                                           `json:"tax" api:"nullable"`
	JSON               subscriptionPreviewChangePlanResponseImmediateChargeSummaryJSON `json:"-"`
}

// subscriptionPreviewChangePlanResponseImmediateChargeSummaryJSON contains the
// JSON metadata for the struct
// [SubscriptionPreviewChangePlanResponseImmediateChargeSummary]
type subscriptionPreviewChangePlanResponseImmediateChargeSummaryJSON struct {
	Currency           apijson.Field
	CustomerCredits    apijson.Field
	SettlementAmount   apijson.Field
	SettlementCurrency apijson.Field
	TotalAmount        apijson.Field
	SettlementTax      apijson.Field
	Tax                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionPreviewChangePlanResponseImmediateChargeSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPreviewChangePlanResponseImmediateChargeSummaryJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetUsageHistoryResponse struct {
	// End date of the billing period
	EndDate time.Time `json:"end_date" api:"required" format:"date-time"`
	// List of meters and their usage for this billing period
	Meters []SubscriptionGetUsageHistoryResponseMeter `json:"meters" api:"required"`
	// Start date of the billing period
	StartDate time.Time                               `json:"start_date" api:"required" format:"date-time"`
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
	ID string `json:"id" api:"required"`
	// Chargeable units (after free threshold) as string for precision
	ChargeableUnits string `json:"chargeable_units" api:"required"`
	// Total units consumed as string for precision
	ConsumedUnits string `json:"consumed_units" api:"required"`
	// Currency for the price per unit
	Currency Currency `json:"currency" api:"required"`
	// Free threshold units for this meter
	FreeThreshold int64 `json:"free_threshold" api:"required"`
	// Meter name
	Name string `json:"name" api:"required"`
	// Price per unit in string format for precision
	PricePerUnit string `json:"price_per_unit" api:"required"`
	// Total price charged for this meter in smallest currency unit (cents)
	TotalPrice int64                                        `json:"total_price" api:"required"`
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
	ClientSecret string                                      `json:"client_secret" api:"nullable"`
	ExpiresOn    time.Time                                   `json:"expires_on" api:"nullable" format:"date-time"`
	PaymentID    string                                      `json:"payment_id" api:"nullable"`
	PaymentLink  string                                      `json:"payment_link" api:"nullable"`
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
	Billing param.Field[BillingAddressParam] `json:"billing" api:"required"`
	// Customer details for the subscription
	Customer param.Field[CustomerRequestUnionParam] `json:"customer" api:"required"`
	// Unique identifier of the product to subscribe to
	ProductID param.Field[string] `json:"product_id" api:"required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity param.Field[int64] `json:"quantity" api:"required"`
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
	// List of one time products that will be bundled with the first payment for this
	// subscription
	OneTimeProductCart param.Field[[]OneTimeProductCartItemParam] `json:"one_time_product_cart"`
	// If true, generates a payment link. Defaults to false if not specified.
	PaymentLink param.Field[bool] `json:"payment_link"`
	// Optional payment method ID to use for this subscription. If provided,
	// customer_id must also be provided (via AttachExistingCustomer). The payment
	// method will be validated for eligibility with the subscription's currency.
	PaymentMethodID param.Field[string] `json:"payment_method_id"`
	// If true, redirects the customer immediately after payment completion False by
	// default
	RedirectImmediately param.Field[bool] `json:"redirect_immediately"`
	// Optional URL to redirect after successful subscription creation
	ReturnURL param.Field[string] `json:"return_url"`
	// If true, returns a shortened payment link. Defaults to false if not specified.
	ShortLink param.Field[bool] `json:"short_link"`
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
	CancelAtNextBillingDate param.Field[bool] `json:"cancel_at_next_billing_date"`
	// Update credit entitlement cart settings
	CreditEntitlementCart param.Field[[]SubscriptionUpdateParamsCreditEntitlementCart] `json:"credit_entitlement_cart"`
	CustomerName          param.Field[string]                                          `json:"customer_name"`
	DisableOnDemand       param.Field[SubscriptionUpdateParamsDisableOnDemand]         `json:"disable_on_demand"`
	Metadata              param.Field[map[string]string]                               `json:"metadata"`
	NextBillingDate       param.Field[time.Time]                                       `json:"next_billing_date" format:"date-time"`
	Status                param.Field[SubscriptionStatus]                              `json:"status"`
	TaxID                 param.Field[string]                                          `json:"tax_id"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParamsCreditEntitlementCart struct {
	CreditEntitlementID        param.Field[string]       `json:"credit_entitlement_id" api:"required"`
	CreditsAmount              param.Field[string]       `json:"credits_amount"`
	ExpiresAfterDays           param.Field[int64]        `json:"expires_after_days"`
	LowBalanceThresholdPercent param.Field[int64]        `json:"low_balance_threshold_percent"`
	MaxRolloverCount           param.Field[int64]        `json:"max_rollover_count"`
	OverageChargeAtBilling     param.Field[bool]         `json:"overage_charge_at_billing"`
	OverageEnabled             param.Field[bool]         `json:"overage_enabled"`
	OverageLimit               param.Field[string]       `json:"overage_limit"`
	RolloverEnabled            param.Field[bool]         `json:"rollover_enabled"`
	RolloverPercentage         param.Field[int64]        `json:"rollover_percentage"`
	RolloverTimeframeCount     param.Field[int64]        `json:"rollover_timeframe_count"`
	RolloverTimeframeInterval  param.Field[TimeInterval] `json:"rollover_timeframe_interval"`
}

func (r SubscriptionUpdateParamsCreditEntitlementCart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParamsDisableOnDemand struct {
	NextBillingDate param.Field[time.Time] `json:"next_billing_date" api:"required" format:"date-time"`
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
	// Filter by product id
	ProductID param.Field[string] `query:"product_id"`
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
	ProductID param.Field[string] `json:"product_id" api:"required"`
	// Proration Billing Mode
	ProrationBillingMode param.Field[SubscriptionChangePlanParamsProrationBillingMode] `json:"proration_billing_mode" api:"required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity param.Field[int64] `json:"quantity" api:"required"`
	// Addons for the new plan. Note : Leaving this empty would remove any existing
	// addons
	Addons param.Field[[]AttachAddonParam] `json:"addons"`
	// Metadata for the payment. If not passed, the metadata of the subscription will
	// be taken
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Controls behavior when the plan change payment fails.
	//
	//   - `prevent_change`: Keep subscription on current plan until payment succeeds
	//   - `apply_change` (default): Apply plan change immediately regardless of payment
	//     outcome
	//
	// If not specified, uses the business-level default setting.
	OnPaymentFailure param.Field[SubscriptionChangePlanParamsOnPaymentFailure] `json:"on_payment_failure"`
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

// Controls behavior when the plan change payment fails.
//
//   - `prevent_change`: Keep subscription on current plan until payment succeeds
//   - `apply_change` (default): Apply plan change immediately regardless of payment
//     outcome
//
// If not specified, uses the business-level default setting.
type SubscriptionChangePlanParamsOnPaymentFailure string

const (
	SubscriptionChangePlanParamsOnPaymentFailurePreventChange SubscriptionChangePlanParamsOnPaymentFailure = "prevent_change"
	SubscriptionChangePlanParamsOnPaymentFailureApplyChange   SubscriptionChangePlanParamsOnPaymentFailure = "apply_change"
)

func (r SubscriptionChangePlanParamsOnPaymentFailure) IsKnown() bool {
	switch r {
	case SubscriptionChangePlanParamsOnPaymentFailurePreventChange, SubscriptionChangePlanParamsOnPaymentFailureApplyChange:
		return true
	}
	return false
}

type SubscriptionChargeParams struct {
	// The product price. Represented in the lowest denomination of the currency (e.g.,
	// cents for USD). For example, to charge $1.00, pass `100`.
	ProductPrice param.Field[int64] `json:"product_price" api:"required"`
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

type SubscriptionPreviewChangePlanParams struct {
	// Unique identifier of the product to subscribe to
	ProductID param.Field[string] `json:"product_id" api:"required"`
	// Proration Billing Mode
	ProrationBillingMode param.Field[SubscriptionPreviewChangePlanParamsProrationBillingMode] `json:"proration_billing_mode" api:"required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity param.Field[int64] `json:"quantity" api:"required"`
	// Addons for the new plan. Note : Leaving this empty would remove any existing
	// addons
	Addons param.Field[[]AttachAddonParam] `json:"addons"`
	// Metadata for the payment. If not passed, the metadata of the subscription will
	// be taken
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Controls behavior when the plan change payment fails.
	//
	//   - `prevent_change`: Keep subscription on current plan until payment succeeds
	//   - `apply_change` (default): Apply plan change immediately regardless of payment
	//     outcome
	//
	// If not specified, uses the business-level default setting.
	OnPaymentFailure param.Field[SubscriptionPreviewChangePlanParamsOnPaymentFailure] `json:"on_payment_failure"`
}

func (r SubscriptionPreviewChangePlanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Proration Billing Mode
type SubscriptionPreviewChangePlanParamsProrationBillingMode string

const (
	SubscriptionPreviewChangePlanParamsProrationBillingModeProratedImmediately   SubscriptionPreviewChangePlanParamsProrationBillingMode = "prorated_immediately"
	SubscriptionPreviewChangePlanParamsProrationBillingModeFullImmediately       SubscriptionPreviewChangePlanParamsProrationBillingMode = "full_immediately"
	SubscriptionPreviewChangePlanParamsProrationBillingModeDifferenceImmediately SubscriptionPreviewChangePlanParamsProrationBillingMode = "difference_immediately"
)

func (r SubscriptionPreviewChangePlanParamsProrationBillingMode) IsKnown() bool {
	switch r {
	case SubscriptionPreviewChangePlanParamsProrationBillingModeProratedImmediately, SubscriptionPreviewChangePlanParamsProrationBillingModeFullImmediately, SubscriptionPreviewChangePlanParamsProrationBillingModeDifferenceImmediately:
		return true
	}
	return false
}

// Controls behavior when the plan change payment fails.
//
//   - `prevent_change`: Keep subscription on current plan until payment succeeds
//   - `apply_change` (default): Apply plan change immediately regardless of payment
//     outcome
//
// If not specified, uses the business-level default setting.
type SubscriptionPreviewChangePlanParamsOnPaymentFailure string

const (
	SubscriptionPreviewChangePlanParamsOnPaymentFailurePreventChange SubscriptionPreviewChangePlanParamsOnPaymentFailure = "prevent_change"
	SubscriptionPreviewChangePlanParamsOnPaymentFailureApplyChange   SubscriptionPreviewChangePlanParamsOnPaymentFailure = "apply_change"
)

func (r SubscriptionPreviewChangePlanParamsOnPaymentFailure) IsKnown() bool {
	switch r {
	case SubscriptionPreviewChangePlanParamsOnPaymentFailurePreventChange, SubscriptionPreviewChangePlanParamsOnPaymentFailureApplyChange:
		return true
	}
	return false
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
	Body SubscriptionUpdatePaymentMethodParamsBodyUnion `json:"body" api:"required"`
}

func (r SubscriptionUpdatePaymentMethodParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SubscriptionUpdatePaymentMethodParamsBody struct {
	Type            param.Field[SubscriptionUpdatePaymentMethodParamsBodyType] `json:"type" api:"required"`
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
	Type      param.Field[SubscriptionUpdatePaymentMethodParamsBodyNewType] `json:"type" api:"required"`
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
	PaymentMethodID param.Field[string]                                                `json:"payment_method_id" api:"required"`
	Type            param.Field[SubscriptionUpdatePaymentMethodParamsBodyExistingType] `json:"type" api:"required"`
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
