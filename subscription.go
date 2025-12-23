// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"encoding/json"
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
func NewSubscriptionService(opts ...option.RequestOption) (r SubscriptionService) {
	r = SubscriptionService{}
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
	AddonID  string `json:"addon_id,required"`
	Quantity int64  `json:"quantity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddonID     respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddonCartResponseItem) RawJSON() string { return r.JSON.raw }
func (r *AddonCartResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AddonID, Quantity are required.
type AttachAddonParam struct {
	AddonID  string `json:"addon_id,required"`
	Quantity int64  `json:"quantity,required"`
	paramObj
}

func (r AttachAddonParam) MarshalJSON() (data []byte, err error) {
	type shadow AttachAddonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachAddonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property MandateOnly is required.
type OnDemandSubscriptionParam struct {
	// If set as True, does not perform any charge and only authorizes payment method
	// details for future use.
	MandateOnly bool `json:"mandate_only,required"`
	// Whether adaptive currency fees should be included in the product_price (true) or
	// added on top (false). This field is ignored if adaptive pricing is not enabled
	// for the business.
	AdaptiveCurrencyFeesInclusive param.Opt[bool] `json:"adaptive_currency_fees_inclusive,omitzero"`
	// Optional product description override for billing and line items. If not
	// specified, the stored description of the product will be used.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// Product price for the initial charge to customer If not specified the stored
	// price of the product will be used Represented in the lowest denomination of the
	// currency (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	ProductPrice param.Opt[int64] `json:"product_price,omitzero"`
	// Optional currency of the product price. If not specified, defaults to the
	// currency of the product.
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
	ProductCurrency Currency `json:"product_currency,omitzero"`
	paramObj
}

func (r OnDemandSubscriptionParam) MarshalJSON() (data []byte, err error) {
	type shadow OnDemandSubscriptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OnDemandSubscriptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	//
	// Any of "Day", "Week", "Month", "Year".
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
	//
	// Any of "pending", "active", "on_hold", "cancelled", "failed", "expired".
	Status SubscriptionStatus `json:"status,required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id,required"`
	// Number of subscription period intervals
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// Time interval for the subscription period (e.g. month, year)
	//
	// Any of "Day", "Week", "Month", "Year".
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
	TaxID string `json:"tax_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addons                     respjson.Field
		Billing                    respjson.Field
		CancelAtNextBillingDate    respjson.Field
		CreatedAt                  respjson.Field
		Currency                   respjson.Field
		Customer                   respjson.Field
		Metadata                   respjson.Field
		Meters                     respjson.Field
		NextBillingDate            respjson.Field
		OnDemand                   respjson.Field
		PaymentFrequencyCount      respjson.Field
		PaymentFrequencyInterval   respjson.Field
		PreviousBillingDate        respjson.Field
		ProductID                  respjson.Field
		Quantity                   respjson.Field
		RecurringPreTaxAmount      respjson.Field
		Status                     respjson.Field
		SubscriptionID             respjson.Field
		SubscriptionPeriodCount    respjson.Field
		SubscriptionPeriodInterval respjson.Field
		TaxInclusive               respjson.Field
		TrialPeriodDays            respjson.Field
		CancelledAt                respjson.Field
		DiscountCyclesRemaining    respjson.Field
		DiscountID                 respjson.Field
		ExpiresAt                  respjson.Field
		PaymentMethodID            respjson.Field
		TaxID                      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subscription) RawJSON() string { return r.JSON.raw }
func (r *Subscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response struct representing usage-based meter cart details for a subscription
type SubscriptionMeter struct {
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
	Currency        Currency `json:"currency,required"`
	FreeThreshold   int64    `json:"free_threshold,required"`
	MeasurementUnit string   `json:"measurement_unit,required"`
	MeterID         string   `json:"meter_id,required"`
	Name            string   `json:"name,required"`
	PricePerUnit    string   `json:"price_per_unit,required"`
	Description     string   `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency        respjson.Field
		FreeThreshold   respjson.Field
		MeasurementUnit respjson.Field
		MeterID         respjson.Field
		Name            respjson.Field
		PricePerUnit    respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionMeter) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionMeter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type TimeInterval string

const (
	TimeIntervalDay   TimeInterval = "Day"
	TimeIntervalWeek  TimeInterval = "Week"
	TimeIntervalMonth TimeInterval = "Month"
	TimeIntervalYear  TimeInterval = "Year"
)

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
	// One time products associated with the purchase of subscription
	OneTimeProductCart []SubscriptionNewResponseOneTimeProductCart `json:"one_time_product_cart,nullable"`
	// URL to checkout page
	PaymentLink string `json:"payment_link,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addons                respjson.Field
		Customer              respjson.Field
		Metadata              respjson.Field
		PaymentID             respjson.Field
		RecurringPreTaxAmount respjson.Field
		SubscriptionID        respjson.Field
		ClientSecret          respjson.Field
		DiscountID            respjson.Field
		ExpiresOn             respjson.Field
		OneTimeProductCart    respjson.Field
		PaymentLink           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionNewResponseOneTimeProductCart struct {
	ProductID string `json:"product_id,required"`
	Quantity  int64  `json:"quantity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductID   respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionNewResponseOneTimeProductCart) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionNewResponseOneTimeProductCart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	//
	// Any of "Day", "Week", "Month", "Year".
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
	//
	// Any of "pending", "active", "on_hold", "cancelled", "failed", "expired".
	Status SubscriptionStatus `json:"status,required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id,required"`
	// Number of subscription period intervals
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// Time interval for the subscription period (e.g. month, year)
	//
	// Any of "Day", "Week", "Month", "Year".
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
	TaxID string `json:"tax_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing                    respjson.Field
		CancelAtNextBillingDate    respjson.Field
		CreatedAt                  respjson.Field
		Currency                   respjson.Field
		Customer                   respjson.Field
		Metadata                   respjson.Field
		NextBillingDate            respjson.Field
		OnDemand                   respjson.Field
		PaymentFrequencyCount      respjson.Field
		PaymentFrequencyInterval   respjson.Field
		PreviousBillingDate        respjson.Field
		ProductID                  respjson.Field
		Quantity                   respjson.Field
		RecurringPreTaxAmount      respjson.Field
		Status                     respjson.Field
		SubscriptionID             respjson.Field
		SubscriptionPeriodCount    respjson.Field
		SubscriptionPeriodInterval respjson.Field
		TaxInclusive               respjson.Field
		TrialPeriodDays            respjson.Field
		CancelledAt                respjson.Field
		DiscountCyclesRemaining    respjson.Field
		DiscountID                 respjson.Field
		PaymentMethodID            respjson.Field
		TaxID                      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionChargeResponse struct {
	PaymentID string `json:"payment_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionChargeResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionChargeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPreviewChangePlanResponse struct {
	ImmediateCharge SubscriptionPreviewChangePlanResponseImmediateCharge `json:"immediate_charge,required"`
	// Response struct representing subscription details
	NewPlan Subscription `json:"new_plan,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImmediateCharge respjson.Field
		NewPlan         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPreviewChangePlanResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionPreviewChangePlanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPreviewChangePlanResponseImmediateCharge struct {
	LineItems []SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion `json:"line_items,required"`
	Summary   SubscriptionPreviewChangePlanResponseImmediateChargeSummary         `json:"summary,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LineItems   respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPreviewChangePlanResponseImmediateCharge) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionPreviewChangePlanResponseImmediateCharge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion contains all
// possible properties and values from
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemSubscription],
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemAddon],
// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion struct {
	ID string `json:"id"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemSubscription].
	Currency Currency `json:"currency"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemSubscription].
	ProductID       string  `json:"product_id"`
	ProrationFactor float64 `json:"proration_factor"`
	Quantity        int64   `json:"quantity"`
	TaxInclusive    bool    `json:"tax_inclusive"`
	Type            string  `json:"type"`
	UnitPrice       int64   `json:"unit_price"`
	Description     string  `json:"description"`
	Name            string  `json:"name"`
	Tax             int64   `json:"tax"`
	TaxRate         float64 `json:"tax_rate"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemAddon].
	TaxCategory TaxCategory `json:"tax_category"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter].
	ChargeableUnits string `json:"chargeable_units"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter].
	FreeThreshold int64 `json:"free_threshold"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter].
	PricePerUnit string `json:"price_per_unit"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter].
	Subtotal int64 `json:"subtotal"`
	// This field is from variant
	// [SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter].
	UnitsConsumed string `json:"units_consumed"`
	JSON          struct {
		ID              respjson.Field
		Currency        respjson.Field
		ProductID       respjson.Field
		ProrationFactor respjson.Field
		Quantity        respjson.Field
		TaxInclusive    respjson.Field
		Type            respjson.Field
		UnitPrice       respjson.Field
		Description     respjson.Field
		Name            respjson.Field
		Tax             respjson.Field
		TaxRate         respjson.Field
		TaxCategory     respjson.Field
		ChargeableUnits respjson.Field
		FreeThreshold   respjson.Field
		PricePerUnit    respjson.Field
		Subtotal        respjson.Field
		UnitsConsumed   respjson.Field
		raw             string
	} `json:"-"`
}

func (u SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion) AsSubscription() (v SubscriptionPreviewChangePlanResponseImmediateChargeLineItemSubscription) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion) AsAddon() (v SubscriptionPreviewChangePlanResponseImmediateChargeLineItemAddon) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion) AsMeter() (v SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemSubscription struct {
	ID string `json:"id,required"`
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
	Currency        Currency `json:"currency,required"`
	ProductID       string   `json:"product_id,required"`
	ProrationFactor float64  `json:"proration_factor,required"`
	Quantity        int64    `json:"quantity,required"`
	TaxInclusive    bool     `json:"tax_inclusive,required"`
	// Any of "subscription".
	Type        string  `json:"type,required"`
	UnitPrice   int64   `json:"unit_price,required"`
	Description string  `json:"description,nullable"`
	Name        string  `json:"name,nullable"`
	Tax         int64   `json:"tax,nullable"`
	TaxRate     float64 `json:"tax_rate,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Currency        respjson.Field
		ProductID       respjson.Field
		ProrationFactor respjson.Field
		Quantity        respjson.Field
		TaxInclusive    respjson.Field
		Type            respjson.Field
		UnitPrice       respjson.Field
		Description     respjson.Field
		Name            respjson.Field
		Tax             respjson.Field
		TaxRate         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemSubscription) RawJSON() string {
	return r.JSON.raw
}
func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItemSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemAddon struct {
	ID string `json:"id,required"`
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
	Currency        Currency `json:"currency,required"`
	Name            string   `json:"name,required"`
	ProrationFactor float64  `json:"proration_factor,required"`
	Quantity        int64    `json:"quantity,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	//
	// Any of "digital_products", "saas", "e_book", "edtech".
	TaxCategory  TaxCategory `json:"tax_category,required"`
	TaxInclusive bool        `json:"tax_inclusive,required"`
	TaxRate      float64     `json:"tax_rate,required"`
	// Any of "addon".
	Type        string `json:"type,required"`
	UnitPrice   int64  `json:"unit_price,required"`
	Description string `json:"description,nullable"`
	Tax         int64  `json:"tax,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Currency        respjson.Field
		Name            respjson.Field
		ProrationFactor respjson.Field
		Quantity        respjson.Field
		TaxCategory     respjson.Field
		TaxInclusive    respjson.Field
		TaxRate         respjson.Field
		Type            respjson.Field
		UnitPrice       respjson.Field
		Description     respjson.Field
		Tax             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemAddon) RawJSON() string {
	return r.JSON.raw
}
func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItemAddon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter struct {
	ID              string `json:"id,required"`
	ChargeableUnits string `json:"chargeable_units,required"`
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
	Currency      Currency `json:"currency,required"`
	FreeThreshold int64    `json:"free_threshold,required"`
	Name          string   `json:"name,required"`
	PricePerUnit  string   `json:"price_per_unit,required"`
	Subtotal      int64    `json:"subtotal,required"`
	TaxInclusive  bool     `json:"tax_inclusive,required"`
	TaxRate       float64  `json:"tax_rate,required"`
	// Any of "meter".
	Type          string `json:"type,required"`
	UnitsConsumed string `json:"units_consumed,required"`
	Description   string `json:"description,nullable"`
	Tax           int64  `json:"tax,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChargeableUnits respjson.Field
		Currency        respjson.Field
		FreeThreshold   respjson.Field
		Name            respjson.Field
		PricePerUnit    respjson.Field
		Subtotal        respjson.Field
		TaxInclusive    respjson.Field
		TaxRate         respjson.Field
		Type            respjson.Field
		UnitsConsumed   respjson.Field
		Description     respjson.Field
		Tax             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter) RawJSON() string {
	return r.JSON.raw
}
func (r *SubscriptionPreviewChangePlanResponseImmediateChargeLineItemMeter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPreviewChangePlanResponseImmediateChargeSummary struct {
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
	Currency         Currency `json:"currency,required"`
	CustomerCredits  int64    `json:"customer_credits,required"`
	SettlementAmount int64    `json:"settlement_amount,required"`
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
	SettlementCurrency Currency `json:"settlement_currency,required"`
	TotalAmount        int64    `json:"total_amount,required"`
	SettlementTax      int64    `json:"settlement_tax,nullable"`
	Tax                int64    `json:"tax,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency           respjson.Field
		CustomerCredits    respjson.Field
		SettlementAmount   respjson.Field
		SettlementCurrency respjson.Field
		TotalAmount        respjson.Field
		SettlementTax      respjson.Field
		Tax                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPreviewChangePlanResponseImmediateChargeSummary) RawJSON() string {
	return r.JSON.raw
}
func (r *SubscriptionPreviewChangePlanResponseImmediateChargeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionGetUsageHistoryResponse struct {
	// End date of the billing period
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// List of meters and their usage for this billing period
	Meters []SubscriptionGetUsageHistoryResponseMeter `json:"meters,required"`
	// Start date of the billing period
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate     respjson.Field
		Meters      respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionGetUsageHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionGetUsageHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionGetUsageHistoryResponseMeter struct {
	// Meter identifier
	ID string `json:"id,required"`
	// Chargeable units (after free threshold) as string for precision
	ChargeableUnits string `json:"chargeable_units,required"`
	// Total units consumed as string for precision
	ConsumedUnits string `json:"consumed_units,required"`
	// Currency for the price per unit
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
	// Free threshold units for this meter
	FreeThreshold int64 `json:"free_threshold,required"`
	// Meter name
	Name string `json:"name,required"`
	// Price per unit in string format for precision
	PricePerUnit string `json:"price_per_unit,required"`
	// Total price charged for this meter in smallest currency unit (cents)
	TotalPrice int64 `json:"total_price,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChargeableUnits respjson.Field
		ConsumedUnits   respjson.Field
		Currency        respjson.Field
		FreeThreshold   respjson.Field
		Name            respjson.Field
		PricePerUnit    respjson.Field
		TotalPrice      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionGetUsageHistoryResponseMeter) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionGetUsageHistoryResponseMeter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionUpdatePaymentMethodResponse struct {
	ClientSecret string    `json:"client_secret,nullable"`
	ExpiresOn    time.Time `json:"expires_on,nullable" format:"date-time"`
	PaymentID    string    `json:"payment_id,nullable"`
	PaymentLink  string    `json:"payment_link,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientSecret respjson.Field
		ExpiresOn    respjson.Field
		PaymentID    respjson.Field
		PaymentLink  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionUpdatePaymentMethodResponse) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionUpdatePaymentMethodResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionNewParams struct {
	// Billing address information for the subscription
	Billing BillingAddressParam `json:"billing,omitzero,required"`
	// Customer details for the subscription
	Customer CustomerRequestUnionParam `json:"customer,omitzero,required"`
	// Unique identifier of the product to subscribe to
	ProductID string `json:"product_id,required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity int64 `json:"quantity,required"`
	// Discount Code to apply to the subscription
	DiscountCode param.Opt[string] `json:"discount_code,omitzero"`
	// Override merchant default 3DS behaviour for this subscription
	Force3DS param.Opt[bool] `json:"force_3ds,omitzero"`
	// If true, generates a payment link. Defaults to false if not specified.
	PaymentLink param.Opt[bool] `json:"payment_link,omitzero"`
	// Optional URL to redirect after successful subscription creation
	ReturnURL param.Opt[string] `json:"return_url,omitzero"`
	// If true, returns a shortened payment link. Defaults to false if not specified.
	ShortLink param.Opt[bool] `json:"short_link,omitzero"`
	// Tax ID in case the payment is B2B. If tax id validation fails the payment
	// creation will fail
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// Optional trial period in days If specified, this value overrides the trial
	// period set in the product's price Must be between 0 and 10000 days
	TrialPeriodDays param.Opt[int64] `json:"trial_period_days,omitzero"`
	// If true, redirects the customer immediately after payment completion False by
	// default
	RedirectImmediately param.Opt[bool] `json:"redirect_immediately,omitzero"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Opt[bool] `json:"show_saved_payment_methods,omitzero"`
	// Attach addons to this subscription
	Addons []AttachAddonParam `json:"addons,omitzero"`
	// List of payment methods allowed during checkout.
	//
	// Customers will **never** see payment methods that are **not** in this list.
	// However, adding a method here **does not guarantee** customers will see it.
	// Availability still depends on other factors (e.g., customer location, merchant
	// settings).
	AllowedPaymentMethodTypes []PaymentMethodTypes `json:"allowed_payment_method_types,omitzero"`
	// List of one time products that will be bundled with the first payment for this
	// subscription
	OneTimeProductCart []OneTimeProductCartItemParam `json:"one_time_product_cart,omitzero"`
	// Fix the currency in which the end customer is billed. If Dodo Payments cannot
	// support that currency for this transaction, it will not proceed
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
	BillingCurrency Currency `json:"billing_currency,omitzero"`
	// Additional metadata for the subscription Defaults to empty if not specified
	Metadata map[string]string         `json:"metadata,omitzero"`
	OnDemand OnDemandSubscriptionParam `json:"on_demand,omitzero"`
	paramObj
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionUpdateParams struct {
	// When set, the subscription will remain active until the end of billing period
	CancelAtNextBillingDate param.Opt[bool]                         `json:"cancel_at_next_billing_date,omitzero"`
	CustomerName            param.Opt[string]                       `json:"customer_name,omitzero"`
	NextBillingDate         param.Opt[time.Time]                    `json:"next_billing_date,omitzero" format:"date-time"`
	TaxID                   param.Opt[string]                       `json:"tax_id,omitzero"`
	DisableOnDemand         SubscriptionUpdateParamsDisableOnDemand `json:"disable_on_demand,omitzero"`
	Metadata                map[string]string                       `json:"metadata,omitzero"`
	Billing                 BillingAddressParam                     `json:"billing,omitzero"`
	// Any of "pending", "active", "on_hold", "cancelled", "failed", "expired".
	Status SubscriptionStatus `json:"status,omitzero"`
	paramObj
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property NextBillingDate is required.
type SubscriptionUpdateParamsDisableOnDemand struct {
	NextBillingDate time.Time `json:"next_billing_date,required" format:"date-time"`
	paramObj
}

func (r SubscriptionUpdateParamsDisableOnDemand) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionUpdateParamsDisableOnDemand
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionUpdateParamsDisableOnDemand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionListParams struct {
	// filter by Brand id
	BrandID param.Opt[string] `query:"brand_id,omitzero" json:"-"`
	// Get events after this created time
	CreatedAtGte param.Opt[time.Time] `query:"created_at_gte,omitzero" format:"date-time" json:"-"`
	// Get events created before this time
	CreatedAtLte param.Opt[time.Time] `query:"created_at_lte,omitzero" format:"date-time" json:"-"`
	// Filter by customer id
	CustomerID param.Opt[string] `query:"customer_id,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter by status
	//
	// Any of "pending", "active", "on_hold", "cancelled", "failed", "expired".
	Status SubscriptionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values, err error) {
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

type SubscriptionChangePlanParams struct {
	// Unique identifier of the product to subscribe to
	ProductID string `json:"product_id,required"`
	// Proration Billing Mode
	//
	// Any of "prorated_immediately", "full_immediately", "difference_immediately".
	ProrationBillingMode SubscriptionChangePlanParamsProrationBillingMode `json:"proration_billing_mode,omitzero,required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity int64 `json:"quantity,required"`
	// Addons for the new plan. Note : Leaving this empty would remove any existing
	// addons
	Addons []AttachAddonParam `json:"addons,omitzero"`
	paramObj
}

func (r SubscriptionChangePlanParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionChangePlanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionChangePlanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Proration Billing Mode
type SubscriptionChangePlanParamsProrationBillingMode string

const (
	SubscriptionChangePlanParamsProrationBillingModeProratedImmediately   SubscriptionChangePlanParamsProrationBillingMode = "prorated_immediately"
	SubscriptionChangePlanParamsProrationBillingModeFullImmediately       SubscriptionChangePlanParamsProrationBillingMode = "full_immediately"
	SubscriptionChangePlanParamsProrationBillingModeDifferenceImmediately SubscriptionChangePlanParamsProrationBillingMode = "difference_immediately"
)

type SubscriptionChargeParams struct {
	// The product price. Represented in the lowest denomination of the currency (e.g.,
	// cents for USD). For example, to charge $1.00, pass `100`.
	ProductPrice int64 `json:"product_price,required"`
	// Whether adaptive currency fees should be included in the product_price (true) or
	// added on top (false). This field is ignored if adaptive pricing is not enabled
	// for the business.
	AdaptiveCurrencyFeesInclusive param.Opt[bool] `json:"adaptive_currency_fees_inclusive,omitzero"`
	// Optional product description override for billing and line items. If not
	// specified, the stored description of the product will be used.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// Specify how customer balance is used for the payment
	CustomerBalanceConfig SubscriptionChargeParamsCustomerBalanceConfig `json:"customer_balance_config,omitzero"`
	// Metadata for the payment. If not passed, the metadata of the subscription will
	// be taken
	Metadata map[string]string `json:"metadata,omitzero"`
	// Optional currency of the product price. If not specified, defaults to the
	// currency of the product.
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
	ProductCurrency Currency `json:"product_currency,omitzero"`
	paramObj
}

func (r SubscriptionChargeParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionChargeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionChargeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify how customer balance is used for the payment
type SubscriptionChargeParamsCustomerBalanceConfig struct {
	// Allows Customer Credit to be purchased to settle payments
	AllowCustomerCreditsPurchase param.Opt[bool] `json:"allow_customer_credits_purchase,omitzero"`
	// Allows Customer Credit Balance to be used to settle payments
	AllowCustomerCreditsUsage param.Opt[bool] `json:"allow_customer_credits_usage,omitzero"`
	paramObj
}

func (r SubscriptionChargeParamsCustomerBalanceConfig) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionChargeParamsCustomerBalanceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionChargeParamsCustomerBalanceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionPreviewChangePlanParams struct {
	// Unique identifier of the product to subscribe to
	ProductID string `json:"product_id,required"`
	// Proration Billing Mode
	//
	// Any of "prorated_immediately", "full_immediately", "difference_immediately".
	ProrationBillingMode SubscriptionPreviewChangePlanParamsProrationBillingMode `json:"proration_billing_mode,omitzero,required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity int64 `json:"quantity,required"`
	// Addons for the new plan. Note : Leaving this empty would remove any existing
	// addons
	Addons []AttachAddonParam `json:"addons,omitzero"`
	paramObj
}

func (r SubscriptionPreviewChangePlanParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionPreviewChangePlanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionPreviewChangePlanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Proration Billing Mode
type SubscriptionPreviewChangePlanParamsProrationBillingMode string

const (
	SubscriptionPreviewChangePlanParamsProrationBillingModeProratedImmediately   SubscriptionPreviewChangePlanParamsProrationBillingMode = "prorated_immediately"
	SubscriptionPreviewChangePlanParamsProrationBillingModeFullImmediately       SubscriptionPreviewChangePlanParamsProrationBillingMode = "full_immediately"
	SubscriptionPreviewChangePlanParamsProrationBillingModeDifferenceImmediately SubscriptionPreviewChangePlanParamsProrationBillingMode = "difference_immediately"
)

type SubscriptionGetUsageHistoryParams struct {
	// Filter by end date (inclusive)
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Filter by specific meter ID
	MeterID param.Opt[string] `query:"meter_id,omitzero" json:"-"`
	// Page number (default: 0)
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size (default: 10, max: 100)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter by start date (inclusive)
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionGetUsageHistoryParams]'s query parameters as
// `url.Values`.
func (r SubscriptionGetUsageHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubscriptionUpdatePaymentMethodParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfNew *SubscriptionUpdatePaymentMethodParamsBodyNew `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfExisting *SubscriptionUpdatePaymentMethodParamsBodyExisting `json:",inline"`

	paramObj
}

func (u SubscriptionUpdatePaymentMethodParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNew, u.OfExisting)
}
func (r *SubscriptionUpdatePaymentMethodParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type SubscriptionUpdatePaymentMethodParamsBodyNew struct {
	// Any of "new".
	Type      string            `json:"type,omitzero,required"`
	ReturnURL param.Opt[string] `json:"return_url,omitzero"`
	paramObj
}

func (r SubscriptionUpdatePaymentMethodParamsBodyNew) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionUpdatePaymentMethodParamsBodyNew
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionUpdatePaymentMethodParamsBodyNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SubscriptionUpdatePaymentMethodParamsBodyNew](
		"type", "new",
	)
}

// The properties PaymentMethodID, Type are required.
type SubscriptionUpdatePaymentMethodParamsBodyExisting struct {
	PaymentMethodID string `json:"payment_method_id,required"`
	// Any of "existing".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r SubscriptionUpdatePaymentMethodParamsBodyExisting) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionUpdatePaymentMethodParamsBodyExisting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionUpdatePaymentMethodParamsBodyExisting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SubscriptionUpdatePaymentMethodParamsBodyExisting](
		"type", "existing",
	)
}
