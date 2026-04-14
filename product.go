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

// ProductService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options    []option.RequestOption
	Images     *ProductImageService
	ShortLinks *ProductShortLinkService
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r *ProductService) {
	r = &ProductService{}
	r.Options = opts
	r.Images = NewProductImageService(opts...)
	r.ShortLinks = NewProductShortLinkService(opts...)
	return
}

func (r *ProductService) New(ctx context.Context, body ProductNewParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ProductService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ProductService) Update(ctx context.Context, id string, body ProductUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

func (r *ProductService) List(ctx context.Context, query ProductListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[ProductListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "products"
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

func (r *ProductService) ListAutoPaging(ctx context.Context, query ProductListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[ProductListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *ProductService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ProductService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("products/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

func (r *ProductService) UpdateFiles(ctx context.Context, id string, body ProductUpdateFilesParams, opts ...option.RequestOption) (res *ProductUpdateFilesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s/files", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type AddMeterToPrice struct {
	MeterID string `json:"meter_id" api:"required"`
	// Optional credit entitlement ID to link this meter to for credit-based billing
	CreditEntitlementID string `json:"credit_entitlement_id" api:"nullable"`
	// Meter description. Will ignored on Request, but will be shown in response
	Description   string `json:"description" api:"nullable"`
	FreeThreshold int64  `json:"free_threshold" api:"nullable"`
	// Meter measurement unit. Will ignored on Request, but will be shown in response
	MeasurementUnit string `json:"measurement_unit" api:"nullable"`
	// Number of meter units that equal one credit. Required when credit_entitlement_id
	// is set.
	MeterUnitsPerCredit string `json:"meter_units_per_credit" api:"nullable"`
	// Meter name. Will ignored on Request, but will be shown in response
	Name string `json:"name" api:"nullable"`
	// The price per unit in lowest denomination. Must be greater than zero. Supports
	// up to 5 digits before decimal point and 12 decimal places.
	PricePerUnit string              `json:"price_per_unit" api:"nullable"`
	JSON         addMeterToPriceJSON `json:"-"`
}

// addMeterToPriceJSON contains the JSON metadata for the struct [AddMeterToPrice]
type addMeterToPriceJSON struct {
	MeterID             apijson.Field
	CreditEntitlementID apijson.Field
	Description         apijson.Field
	FreeThreshold       apijson.Field
	MeasurementUnit     apijson.Field
	MeterUnitsPerCredit apijson.Field
	Name                apijson.Field
	PricePerUnit        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AddMeterToPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addMeterToPriceJSON) RawJSON() string {
	return r.raw
}

type AddMeterToPriceParam struct {
	MeterID param.Field[string] `json:"meter_id" api:"required"`
	// Optional credit entitlement ID to link this meter to for credit-based billing
	CreditEntitlementID param.Field[string] `json:"credit_entitlement_id"`
	// Meter description. Will ignored on Request, but will be shown in response
	Description   param.Field[string] `json:"description"`
	FreeThreshold param.Field[int64]  `json:"free_threshold"`
	// Meter measurement unit. Will ignored on Request, but will be shown in response
	MeasurementUnit param.Field[string] `json:"measurement_unit"`
	// Number of meter units that equal one credit. Required when credit_entitlement_id
	// is set.
	MeterUnitsPerCredit param.Field[string] `json:"meter_units_per_credit"`
	// Meter name. Will ignored on Request, but will be shown in response
	Name param.Field[string] `json:"name"`
	// The price per unit in lowest denomination. Must be greater than zero. Supports
	// up to 5 digits before decimal point and 12 decimal places.
	PricePerUnit param.Field[string] `json:"price_per_unit"`
}

func (r AddMeterToPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Request struct for attaching a credit entitlement to a product
type AttachCreditEntitlementParam struct {
	// ID of the credit entitlement to attach
	CreditEntitlementID param.Field[string] `json:"credit_entitlement_id" api:"required"`
	// Number of credits to grant when this product is purchased
	CreditsAmount param.Field[string] `json:"credits_amount" api:"required"`
	// Currency for credit-related pricing
	Currency param.Field[Currency] `json:"currency"`
	// Number of days after which credits expire
	ExpiresAfterDays param.Field[int64] `json:"expires_after_days"`
	// Balance threshold percentage for low balance notifications (0-100)
	LowBalanceThresholdPercent param.Field[int64] `json:"low_balance_threshold_percent"`
	// Maximum number of rollover cycles allowed
	MaxRolloverCount param.Field[int64] `json:"max_rollover_count"`
	// Controls how overage is handled at billing cycle end.
	OverageBehavior param.Field[CbbOverageBehavior] `json:"overage_behavior"`
	// Whether overage usage is allowed beyond credit balance
	OverageEnabled param.Field[bool] `json:"overage_enabled"`
	// Maximum amount of overage allowed
	OverageLimit param.Field[string] `json:"overage_limit"`
	// Price per credit unit for purchasing additional credits
	PricePerUnit param.Field[string] `json:"price_per_unit"`
	// Proration behavior for credit grants during plan changes
	ProrationBehavior param.Field[CbbProrationBehavior] `json:"proration_behavior"`
	// Whether unused credits can roll over to the next billing period
	RolloverEnabled param.Field[bool] `json:"rollover_enabled"`
	// Percentage of unused credits that can roll over (0-100)
	RolloverPercentage param.Field[int64] `json:"rollover_percentage"`
	// Number of timeframe units for rollover window
	RolloverTimeframeCount param.Field[int64] `json:"rollover_timeframe_count"`
	// Time interval for rollover window (day, week, month, year)
	RolloverTimeframeInterval param.Field[TimeInterval] `json:"rollover_timeframe_interval"`
	// Credits granted during trial period
	TrialCredits param.Field[string] `json:"trial_credits"`
	// Whether trial credits expire when trial ends
	TrialCreditsExpireAfterTrial param.Field[bool] `json:"trial_credits_expire_after_trial"`
}

func (r AttachCreditEntitlementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CbbProrationBehavior string

const (
	CbbProrationBehaviorProrate   CbbProrationBehavior = "prorate"
	CbbProrationBehaviorNoProrate CbbProrationBehavior = "no_prorate"
)

func (r CbbProrationBehavior) IsKnown() bool {
	switch r {
	case CbbProrationBehaviorProrate, CbbProrationBehaviorNoProrate:
		return true
	}
	return false
}

// Response struct for credit entitlement mapping
type CreditEntitlementMappingResponse struct {
	// Unique ID of this mapping
	ID string `json:"id" api:"required" format:"uuid"`
	// ID of the credit entitlement
	CreditEntitlementID string `json:"credit_entitlement_id" api:"required"`
	// Name of the credit entitlement
	CreditEntitlementName string `json:"credit_entitlement_name" api:"required"`
	// Unit label for the credit entitlement
	CreditEntitlementUnit string `json:"credit_entitlement_unit" api:"required"`
	// Number of credits granted
	CreditsAmount string `json:"credits_amount" api:"required"`
	// Controls how overage is handled at billing cycle end.
	OverageBehavior CbbOverageBehavior `json:"overage_behavior" api:"required"`
	// Whether overage is enabled
	OverageEnabled bool `json:"overage_enabled" api:"required"`
	// Proration behavior for credit grants during plan changes
	ProrationBehavior CbbProrationBehavior `json:"proration_behavior" api:"required"`
	// Whether rollover is enabled
	RolloverEnabled bool `json:"rollover_enabled" api:"required"`
	// Whether trial credits expire after trial
	TrialCreditsExpireAfterTrial bool `json:"trial_credits_expire_after_trial" api:"required"`
	// Currency
	Currency Currency `json:"currency" api:"nullable"`
	// Days until credits expire
	ExpiresAfterDays int64 `json:"expires_after_days" api:"nullable"`
	// Low balance threshold percentage
	LowBalanceThresholdPercent int64 `json:"low_balance_threshold_percent" api:"nullable"`
	// Maximum rollover cycles
	MaxRolloverCount int64 `json:"max_rollover_count" api:"nullable"`
	// Overage limit
	OverageLimit string `json:"overage_limit" api:"nullable"`
	// Price per unit
	PricePerUnit string `json:"price_per_unit" api:"nullable"`
	// Rollover percentage
	RolloverPercentage int64 `json:"rollover_percentage" api:"nullable"`
	// Rollover timeframe count
	RolloverTimeframeCount int64 `json:"rollover_timeframe_count" api:"nullable"`
	// Rollover timeframe interval
	RolloverTimeframeInterval TimeInterval `json:"rollover_timeframe_interval" api:"nullable"`
	// Trial credits
	TrialCredits string                               `json:"trial_credits" api:"nullable"`
	JSON         creditEntitlementMappingResponseJSON `json:"-"`
}

// creditEntitlementMappingResponseJSON contains the JSON metadata for the struct
// [CreditEntitlementMappingResponse]
type creditEntitlementMappingResponseJSON struct {
	ID                           apijson.Field
	CreditEntitlementID          apijson.Field
	CreditEntitlementName        apijson.Field
	CreditEntitlementUnit        apijson.Field
	CreditsAmount                apijson.Field
	OverageBehavior              apijson.Field
	OverageEnabled               apijson.Field
	ProrationBehavior            apijson.Field
	RolloverEnabled              apijson.Field
	TrialCreditsExpireAfterTrial apijson.Field
	Currency                     apijson.Field
	ExpiresAfterDays             apijson.Field
	LowBalanceThresholdPercent   apijson.Field
	MaxRolloverCount             apijson.Field
	OverageLimit                 apijson.Field
	PricePerUnit                 apijson.Field
	RolloverPercentage           apijson.Field
	RolloverTimeframeCount       apijson.Field
	RolloverTimeframeInterval    apijson.Field
	TrialCredits                 apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CreditEntitlementMappingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditEntitlementMappingResponseJSON) RawJSON() string {
	return r.raw
}

type DigitalProductDelivery struct {
	// External URL to digital product
	ExternalURL string `json:"external_url" api:"nullable"`
	// Uploaded files ids of digital product
	Files []DigitalProductDeliveryFile `json:"files" api:"nullable"`
	// Instructions to download and use the digital product
	Instructions string                     `json:"instructions" api:"nullable"`
	JSON         digitalProductDeliveryJSON `json:"-"`
}

// digitalProductDeliveryJSON contains the JSON metadata for the struct
// [DigitalProductDelivery]
type digitalProductDeliveryJSON struct {
	ExternalURL  apijson.Field
	Files        apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DigitalProductDelivery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalProductDeliveryJSON) RawJSON() string {
	return r.raw
}

type DigitalProductDeliveryFile struct {
	FileID   string                         `json:"file_id" api:"required" format:"uuid"`
	FileName string                         `json:"file_name" api:"required"`
	URL      string                         `json:"url" api:"required"`
	JSON     digitalProductDeliveryFileJSON `json:"-"`
}

// digitalProductDeliveryFileJSON contains the JSON metadata for the struct
// [DigitalProductDeliveryFile]
type digitalProductDeliveryFileJSON struct {
	FileID      apijson.Field
	FileName    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalProductDeliveryFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalProductDeliveryFileJSON) RawJSON() string {
	return r.raw
}

type LicenseKeyDuration struct {
	Count    int64                  `json:"count" api:"required"`
	Interval TimeInterval           `json:"interval" api:"required"`
	JSON     licenseKeyDurationJSON `json:"-"`
}

// licenseKeyDurationJSON contains the JSON metadata for the struct
// [LicenseKeyDuration]
type licenseKeyDurationJSON struct {
	Count       apijson.Field
	Interval    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseKeyDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseKeyDurationJSON) RawJSON() string {
	return r.raw
}

type LicenseKeyDurationParam struct {
	Count    param.Field[int64]        `json:"count" api:"required"`
	Interval param.Field[TimeInterval] `json:"interval" api:"required"`
}

func (r LicenseKeyDurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One-time price details.
type Price struct {
	// The currency in which the payment is made.
	Currency Currency `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity bool      `json:"purchasing_power_parity" api:"required"`
	Type                  PriceType `json:"type" api:"required"`
	// The fixed payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	FixedPrice int64 `json:"fixed_price"`
	// This field can have the runtime type of [[]AddMeterToPrice].
	Meters interface{} `json:"meters"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant bool `json:"pay_what_you_want"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount int64 `json:"payment_frequency_count"`
	// The time interval for the payment frequency (e.g., day, month, year).
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price int64 `json:"price"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount int64 `json:"subscription_period_count"`
	// The time interval for the subscription period (e.g., day, month, year).
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice int64 `json:"suggested_price" api:"nullable"`
	// Indicates if the price is tax inclusive.
	TaxInclusive bool `json:"tax_inclusive" api:"nullable"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays int64     `json:"trial_period_days"`
	JSON            priceJSON `json:"-"`
	union           PriceUnion
}

// priceJSON contains the JSON metadata for the struct [Price]
type priceJSON struct {
	Currency                   apijson.Field
	Discount                   apijson.Field
	PurchasingPowerParity      apijson.Field
	Type                       apijson.Field
	FixedPrice                 apijson.Field
	Meters                     apijson.Field
	PayWhatYouWant             apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	Price                      apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	SuggestedPrice             apijson.Field
	TaxInclusive               apijson.Field
	TrialPeriodDays            apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r priceJSON) RawJSON() string {
	return r.raw
}

func (r *Price) UnmarshalJSON(data []byte) (err error) {
	*r = Price{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PriceUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [PriceOneTimePrice],
// [PriceRecurringPrice], [PriceUsageBasedPrice].
func (r Price) AsUnion() PriceUnion {
	return r.union
}

// One-time price details.
//
// Union satisfied by [PriceOneTimePrice], [PriceRecurringPrice] or
// [PriceUsageBasedPrice].
type PriceUnion interface {
	implementsPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PriceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PriceOneTimePrice{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PriceRecurringPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PriceUsageBasedPrice{}),
		},
	)
}

// One-time price details.
type PriceOneTimePrice struct {
	// The currency in which the payment is made.
	Currency Currency `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount" api:"required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price int64 `json:"price" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity bool                  `json:"purchasing_power_parity" api:"required"`
	Type                  PriceOneTimePriceType `json:"type" api:"required"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant bool `json:"pay_what_you_want"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice int64 `json:"suggested_price" api:"nullable"`
	// Indicates if the price is tax inclusive.
	TaxInclusive bool                  `json:"tax_inclusive" api:"nullable"`
	JSON         priceOneTimePriceJSON `json:"-"`
}

// priceOneTimePriceJSON contains the JSON metadata for the struct
// [PriceOneTimePrice]
type priceOneTimePriceJSON struct {
	Currency              apijson.Field
	Discount              apijson.Field
	Price                 apijson.Field
	PurchasingPowerParity apijson.Field
	Type                  apijson.Field
	PayWhatYouWant        apijson.Field
	SuggestedPrice        apijson.Field
	TaxInclusive          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PriceOneTimePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceOneTimePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceOneTimePrice) implementsPrice() {}

type PriceOneTimePriceType string

const (
	PriceOneTimePriceTypeOneTimePrice PriceOneTimePriceType = "one_time_price"
)

func (r PriceOneTimePriceType) IsKnown() bool {
	switch r {
	case PriceOneTimePriceTypeOneTimePrice:
		return true
	}
	return false
}

// Recurring price details.
type PriceRecurringPrice struct {
	// The currency in which the payment is made.
	Currency Currency `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount" api:"required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount int64 `json:"payment_frequency_count" api:"required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval" api:"required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price int64 `json:"price" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity" api:"required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount int64 `json:"subscription_period_count" api:"required"`
	// The time interval for the subscription period (e.g., day, month, year).
	SubscriptionPeriodInterval TimeInterval            `json:"subscription_period_interval" api:"required"`
	Type                       PriceRecurringPriceType `json:"type" api:"required"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool `json:"tax_inclusive" api:"nullable"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays int64                   `json:"trial_period_days"`
	JSON            priceRecurringPriceJSON `json:"-"`
}

// priceRecurringPriceJSON contains the JSON metadata for the struct
// [PriceRecurringPrice]
type priceRecurringPriceJSON struct {
	Currency                   apijson.Field
	Discount                   apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	Price                      apijson.Field
	PurchasingPowerParity      apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	Type                       apijson.Field
	TaxInclusive               apijson.Field
	TrialPeriodDays            apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PriceRecurringPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceRecurringPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceRecurringPrice) implementsPrice() {}

type PriceRecurringPriceType string

const (
	PriceRecurringPriceTypeRecurringPrice PriceRecurringPriceType = "recurring_price"
)

func (r PriceRecurringPriceType) IsKnown() bool {
	switch r {
	case PriceRecurringPriceTypeRecurringPrice:
		return true
	}
	return false
}

// Usage Based price details.
type PriceUsageBasedPrice struct {
	// The currency in which the payment is made.
	Currency Currency `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount" api:"required"`
	// The fixed payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	FixedPrice int64 `json:"fixed_price" api:"required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount int64 `json:"payment_frequency_count" api:"required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity" api:"required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount int64 `json:"subscription_period_count" api:"required"`
	// The time interval for the subscription period (e.g., day, month, year).
	SubscriptionPeriodInterval TimeInterval             `json:"subscription_period_interval" api:"required"`
	Type                       PriceUsageBasedPriceType `json:"type" api:"required"`
	Meters                     []AddMeterToPrice        `json:"meters" api:"nullable"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool                     `json:"tax_inclusive" api:"nullable"`
	JSON         priceUsageBasedPriceJSON `json:"-"`
}

// priceUsageBasedPriceJSON contains the JSON metadata for the struct
// [PriceUsageBasedPrice]
type priceUsageBasedPriceJSON struct {
	Currency                   apijson.Field
	Discount                   apijson.Field
	FixedPrice                 apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	PurchasingPowerParity      apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	Type                       apijson.Field
	Meters                     apijson.Field
	TaxInclusive               apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PriceUsageBasedPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUsageBasedPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceUsageBasedPrice) implementsPrice() {}

type PriceUsageBasedPriceType string

const (
	PriceUsageBasedPriceTypeUsageBasedPrice PriceUsageBasedPriceType = "usage_based_price"
)

func (r PriceUsageBasedPriceType) IsKnown() bool {
	switch r {
	case PriceUsageBasedPriceTypeUsageBasedPrice:
		return true
	}
	return false
}

type PriceType string

const (
	PriceTypeOneTimePrice    PriceType = "one_time_price"
	PriceTypeRecurringPrice  PriceType = "recurring_price"
	PriceTypeUsageBasedPrice PriceType = "usage_based_price"
)

func (r PriceType) IsKnown() bool {
	switch r {
	case PriceTypeOneTimePrice, PriceTypeRecurringPrice, PriceTypeUsageBasedPrice:
		return true
	}
	return false
}

// One-time price details.
type PriceParam struct {
	// The currency in which the payment is made.
	Currency param.Field[Currency] `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[int64] `json:"discount" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity param.Field[bool]      `json:"purchasing_power_parity" api:"required"`
	Type                  param.Field[PriceType] `json:"type" api:"required"`
	// The fixed payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	FixedPrice param.Field[int64]       `json:"fixed_price"`
	Meters     param.Field[interface{}] `json:"meters"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant param.Field[bool] `json:"pay_what_you_want"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount param.Field[int64] `json:"payment_frequency_count"`
	// The time interval for the payment frequency (e.g., day, month, year).
	PaymentFrequencyInterval param.Field[TimeInterval] `json:"payment_frequency_interval"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price param.Field[int64] `json:"price"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount param.Field[int64] `json:"subscription_period_count"`
	// The time interval for the subscription period (e.g., day, month, year).
	SubscriptionPeriodInterval param.Field[TimeInterval] `json:"subscription_period_interval"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice param.Field[int64] `json:"suggested_price"`
	// Indicates if the price is tax inclusive.
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r PriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceParam) implementsPriceUnionParam() {}

// One-time price details.
//
// Satisfied by [PriceOneTimePriceParam], [PriceRecurringPriceParam],
// [PriceUsageBasedPriceParam], [PriceParam].
type PriceUnionParam interface {
	implementsPriceUnionParam()
}

// One-time price details.
type PriceOneTimePriceParam struct {
	// The currency in which the payment is made.
	Currency param.Field[Currency] `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[int64] `json:"discount" api:"required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price param.Field[int64] `json:"price" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity param.Field[bool]                  `json:"purchasing_power_parity" api:"required"`
	Type                  param.Field[PriceOneTimePriceType] `json:"type" api:"required"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant param.Field[bool] `json:"pay_what_you_want"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice param.Field[int64] `json:"suggested_price"`
	// Indicates if the price is tax inclusive.
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
}

func (r PriceOneTimePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceOneTimePriceParam) implementsPriceUnionParam() {}

// Recurring price details.
type PriceRecurringPriceParam struct {
	// The currency in which the payment is made.
	Currency param.Field[Currency] `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[int64] `json:"discount" api:"required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount param.Field[int64] `json:"payment_frequency_count" api:"required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	PaymentFrequencyInterval param.Field[TimeInterval] `json:"payment_frequency_interval" api:"required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool] `json:"purchasing_power_parity" api:"required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount param.Field[int64] `json:"subscription_period_count" api:"required"`
	// The time interval for the subscription period (e.g., day, month, year).
	SubscriptionPeriodInterval param.Field[TimeInterval]            `json:"subscription_period_interval" api:"required"`
	Type                       param.Field[PriceRecurringPriceType] `json:"type" api:"required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r PriceRecurringPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceRecurringPriceParam) implementsPriceUnionParam() {}

// Usage Based price details.
type PriceUsageBasedPriceParam struct {
	// The currency in which the payment is made.
	Currency param.Field[Currency] `json:"currency" api:"required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[int64] `json:"discount" api:"required"`
	// The fixed payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	FixedPrice param.Field[int64] `json:"fixed_price" api:"required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount param.Field[int64] `json:"payment_frequency_count" api:"required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	PaymentFrequencyInterval param.Field[TimeInterval] `json:"payment_frequency_interval" api:"required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool] `json:"purchasing_power_parity" api:"required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount param.Field[int64] `json:"subscription_period_count" api:"required"`
	// The time interval for the subscription period (e.g., day, month, year).
	SubscriptionPeriodInterval param.Field[TimeInterval]             `json:"subscription_period_interval" api:"required"`
	Type                       param.Field[PriceUsageBasedPriceType] `json:"type" api:"required"`
	Meters                     param.Field[[]AddMeterToPriceParam]   `json:"meters"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
}

func (r PriceUsageBasedPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceUsageBasedPriceParam) implementsPriceUnionParam() {}

type Product struct {
	BrandID string `json:"brand_id" api:"required"`
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id" api:"required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Attached credit entitlements with settings
	CreditEntitlements []CreditEntitlementMappingResponse `json:"credit_entitlements" api:"required"`
	// Attached entitlements (integration-based access grants)
	Entitlements []ProductEntitlement `json:"entitlements" api:"required"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring" api:"required"`
	// Indicates whether the product requires a license key.
	//
	// Deprecated: deprecated
	LicenseKeyEnabled bool `json:"license_key_enabled" api:"required"`
	// Additional custom data associated with the product
	Metadata map[string]string `json:"metadata" api:"required"`
	// Pricing information for the product.
	Price Price `json:"price" api:"required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id" api:"required"`
	// Tax category associated with the product.
	TaxCategory TaxCategory `json:"tax_category" api:"required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Available Addons for subscription products
	Addons []string `json:"addons" api:"nullable"`
	// Description of the product, optional.
	Description            string                 `json:"description" api:"nullable"`
	DigitalProductDelivery DigitalProductDelivery `json:"digital_product_delivery" api:"nullable"`
	// URL of the product image, optional.
	Image string `json:"image" api:"nullable"`
	// Message sent upon license key activation, if applicable.
	//
	// Deprecated: deprecated
	LicenseKeyActivationMessage string `json:"license_key_activation_message" api:"nullable"`
	// Limit on the number of activations for the license key, if enabled.
	//
	// Deprecated: deprecated
	LicenseKeyActivationsLimit int64 `json:"license_key_activations_limit" api:"nullable"`
	// Duration of the license key validity, if enabled.
	LicenseKeyDuration LicenseKeyDuration `json:"license_key_duration" api:"nullable"`
	// Name of the product, optional.
	Name string `json:"name" api:"nullable"`
	// The product collection ID this product belongs to, if any
	ProductCollectionID string      `json:"product_collection_id" api:"nullable"`
	JSON                productJSON `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	BrandID                     apijson.Field
	BusinessID                  apijson.Field
	CreatedAt                   apijson.Field
	CreditEntitlements          apijson.Field
	Entitlements                apijson.Field
	IsRecurring                 apijson.Field
	LicenseKeyEnabled           apijson.Field
	Metadata                    apijson.Field
	Price                       apijson.Field
	ProductID                   apijson.Field
	TaxCategory                 apijson.Field
	UpdatedAt                   apijson.Field
	Addons                      apijson.Field
	Description                 apijson.Field
	DigitalProductDelivery      apijson.Field
	Image                       apijson.Field
	LicenseKeyActivationMessage apijson.Field
	LicenseKeyActivationsLimit  apijson.Field
	LicenseKeyDuration          apijson.Field
	Name                        apijson.Field
	ProductCollectionID         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Product) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productJSON) RawJSON() string {
	return r.raw
}

// Summary of an entitlement attached to a product
type ProductEntitlement struct {
	ID string `json:"id" api:"required"`
	// Platform-specific configuration for an entitlement. Each variant uses unique
	// field names so `#[serde(untagged)]` can disambiguate correctly.
	IntegrationConfig ProductEntitlementsIntegrationConfig `json:"integration_config" api:"required"`
	IntegrationType   ProductEntitlementsIntegrationType   `json:"integration_type" api:"required"`
	Name              string                               `json:"name" api:"required"`
	Description       string                               `json:"description" api:"nullable"`
	JSON              productEntitlementJSON               `json:"-"`
}

// productEntitlementJSON contains the JSON metadata for the struct
// [ProductEntitlement]
type productEntitlementJSON struct {
	ID                apijson.Field
	IntegrationConfig apijson.Field
	IntegrationType   apijson.Field
	Name              apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProductEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementJSON) RawJSON() string {
	return r.raw
}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
type ProductEntitlementsIntegrationConfig struct {
	ActivationMessage string `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64  `json:"activations_limit" api:"nullable"`
	ChatID            string `json:"chat_id"`
	// This field can have the runtime type of [[]string].
	DigitalFileIDs   interface{} `json:"digital_file_ids"`
	DurationCount    int64       `json:"duration_count" api:"nullable"`
	DurationInterval string      `json:"duration_interval" api:"nullable"`
	ExternalURL      string      `json:"external_url" api:"nullable"`
	FigmaFileID      string      `json:"figma_file_id"`
	FramerTemplateID string      `json:"framer_template_id"`
	GuildID          string      `json:"guild_id"`
	Instructions     string      `json:"instructions" api:"nullable"`
	NotionTemplateID string      `json:"notion_template_id"`
	// One of: pull, push, admin, maintain, triage
	Permission string                                   `json:"permission"`
	RoleID     string                                   `json:"role_id" api:"nullable"`
	TargetID   string                                   `json:"target_id"`
	JSON       productEntitlementsIntegrationConfigJSON `json:"-"`
	union      ProductEntitlementsIntegrationConfigUnion
}

// productEntitlementsIntegrationConfigJSON contains the JSON metadata for the
// struct [ProductEntitlementsIntegrationConfig]
type productEntitlementsIntegrationConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	ChatID            apijson.Field
	DigitalFileIDs    apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	ExternalURL       apijson.Field
	FigmaFileID       apijson.Field
	FramerTemplateID  apijson.Field
	GuildID           apijson.Field
	Instructions      apijson.Field
	NotionTemplateID  apijson.Field
	Permission        apijson.Field
	RoleID            apijson.Field
	TargetID          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r productEntitlementsIntegrationConfigJSON) RawJSON() string {
	return r.raw
}

func (r *ProductEntitlementsIntegrationConfig) UnmarshalJSON(data []byte) (err error) {
	*r = ProductEntitlementsIntegrationConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProductEntitlementsIntegrationConfigUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProductEntitlementsIntegrationConfigGitHubConfig],
// [ProductEntitlementsIntegrationConfigDiscordConfig],
// [ProductEntitlementsIntegrationConfigTelegramConfig],
// [ProductEntitlementsIntegrationConfigFigmaConfig],
// [ProductEntitlementsIntegrationConfigFramerConfig],
// [ProductEntitlementsIntegrationConfigNotionConfig],
// [ProductEntitlementsIntegrationConfigDigitalFilesConfig],
// [ProductEntitlementsIntegrationConfigLicenseKeyConfig].
func (r ProductEntitlementsIntegrationConfig) AsUnion() ProductEntitlementsIntegrationConfigUnion {
	return r.union
}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
//
// Union satisfied by [ProductEntitlementsIntegrationConfigGitHubConfig],
// [ProductEntitlementsIntegrationConfigDiscordConfig],
// [ProductEntitlementsIntegrationConfigTelegramConfig],
// [ProductEntitlementsIntegrationConfigFigmaConfig],
// [ProductEntitlementsIntegrationConfigFramerConfig],
// [ProductEntitlementsIntegrationConfigNotionConfig],
// [ProductEntitlementsIntegrationConfigDigitalFilesConfig] or
// [ProductEntitlementsIntegrationConfigLicenseKeyConfig].
type ProductEntitlementsIntegrationConfigUnion interface {
	implementsProductEntitlementsIntegrationConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProductEntitlementsIntegrationConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigGitHubConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigDiscordConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigTelegramConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigFigmaConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigFramerConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigNotionConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigDigitalFilesConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductEntitlementsIntegrationConfigLicenseKeyConfig{}),
		},
	)
}

type ProductEntitlementsIntegrationConfigGitHubConfig struct {
	// One of: pull, push, admin, maintain, triage
	Permission string                                               `json:"permission" api:"required"`
	TargetID   string                                               `json:"target_id" api:"required"`
	JSON       productEntitlementsIntegrationConfigGitHubConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigGitHubConfigJSON contains the JSON metadata
// for the struct [ProductEntitlementsIntegrationConfigGitHubConfig]
type productEntitlementsIntegrationConfigGitHubConfigJSON struct {
	Permission  apijson.Field
	TargetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigGitHubConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigGitHubConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationConfigDiscordConfig struct {
	GuildID string                                                `json:"guild_id" api:"required"`
	RoleID  string                                                `json:"role_id" api:"nullable"`
	JSON    productEntitlementsIntegrationConfigDiscordConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigDiscordConfigJSON contains the JSON metadata
// for the struct [ProductEntitlementsIntegrationConfigDiscordConfig]
type productEntitlementsIntegrationConfigDiscordConfigJSON struct {
	GuildID     apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigDiscordConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigDiscordConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigDiscordConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationConfigTelegramConfig struct {
	ChatID string                                                 `json:"chat_id" api:"required"`
	JSON   productEntitlementsIntegrationConfigTelegramConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigTelegramConfigJSON contains the JSON
// metadata for the struct [ProductEntitlementsIntegrationConfigTelegramConfig]
type productEntitlementsIntegrationConfigTelegramConfigJSON struct {
	ChatID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigTelegramConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigTelegramConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigTelegramConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationConfigFigmaConfig struct {
	FigmaFileID string                                              `json:"figma_file_id" api:"required"`
	JSON        productEntitlementsIntegrationConfigFigmaConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigFigmaConfigJSON contains the JSON metadata
// for the struct [ProductEntitlementsIntegrationConfigFigmaConfig]
type productEntitlementsIntegrationConfigFigmaConfigJSON struct {
	FigmaFileID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigFigmaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigFigmaConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigFigmaConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationConfigFramerConfig struct {
	FramerTemplateID string                                               `json:"framer_template_id" api:"required"`
	JSON             productEntitlementsIntegrationConfigFramerConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigFramerConfigJSON contains the JSON metadata
// for the struct [ProductEntitlementsIntegrationConfigFramerConfig]
type productEntitlementsIntegrationConfigFramerConfigJSON struct {
	FramerTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigFramerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigFramerConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigFramerConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationConfigNotionConfig struct {
	NotionTemplateID string                                               `json:"notion_template_id" api:"required"`
	JSON             productEntitlementsIntegrationConfigNotionConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigNotionConfigJSON contains the JSON metadata
// for the struct [ProductEntitlementsIntegrationConfigNotionConfig]
type productEntitlementsIntegrationConfigNotionConfigJSON struct {
	NotionTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigNotionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigNotionConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigNotionConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationConfigDigitalFilesConfig struct {
	DigitalFileIDs []string                                                   `json:"digital_file_ids" api:"required"`
	ExternalURL    string                                                     `json:"external_url" api:"nullable"`
	Instructions   string                                                     `json:"instructions" api:"nullable"`
	JSON           productEntitlementsIntegrationConfigDigitalFilesConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigDigitalFilesConfigJSON contains the JSON
// metadata for the struct [ProductEntitlementsIntegrationConfigDigitalFilesConfig]
type productEntitlementsIntegrationConfigDigitalFilesConfigJSON struct {
	DigitalFileIDs apijson.Field
	ExternalURL    apijson.Field
	Instructions   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigDigitalFilesConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigDigitalFilesConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigDigitalFilesConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage string                                                   `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64                                                    `json:"activations_limit" api:"nullable"`
	DurationCount     int64                                                    `json:"duration_count" api:"nullable"`
	DurationInterval  string                                                   `json:"duration_interval" api:"nullable"`
	JSON              productEntitlementsIntegrationConfigLicenseKeyConfigJSON `json:"-"`
}

// productEntitlementsIntegrationConfigLicenseKeyConfigJSON contains the JSON
// metadata for the struct [ProductEntitlementsIntegrationConfigLicenseKeyConfig]
type productEntitlementsIntegrationConfigLicenseKeyConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProductEntitlementsIntegrationConfigLicenseKeyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productEntitlementsIntegrationConfigLicenseKeyConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductEntitlementsIntegrationConfigLicenseKeyConfig) implementsProductEntitlementsIntegrationConfig() {
}

type ProductEntitlementsIntegrationType string

const (
	ProductEntitlementsIntegrationTypeDiscord      ProductEntitlementsIntegrationType = "discord"
	ProductEntitlementsIntegrationTypeTelegram     ProductEntitlementsIntegrationType = "telegram"
	ProductEntitlementsIntegrationTypeGitHub       ProductEntitlementsIntegrationType = "github"
	ProductEntitlementsIntegrationTypeFigma        ProductEntitlementsIntegrationType = "figma"
	ProductEntitlementsIntegrationTypeFramer       ProductEntitlementsIntegrationType = "framer"
	ProductEntitlementsIntegrationTypeNotion       ProductEntitlementsIntegrationType = "notion"
	ProductEntitlementsIntegrationTypeDigitalFiles ProductEntitlementsIntegrationType = "digital_files"
	ProductEntitlementsIntegrationTypeLicenseKey   ProductEntitlementsIntegrationType = "license_key"
)

func (r ProductEntitlementsIntegrationType) IsKnown() bool {
	switch r {
	case ProductEntitlementsIntegrationTypeDiscord, ProductEntitlementsIntegrationTypeTelegram, ProductEntitlementsIntegrationTypeGitHub, ProductEntitlementsIntegrationTypeFigma, ProductEntitlementsIntegrationTypeFramer, ProductEntitlementsIntegrationTypeNotion, ProductEntitlementsIntegrationTypeDigitalFiles, ProductEntitlementsIntegrationTypeLicenseKey:
		return true
	}
	return false
}

type ProductListResponse struct {
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id" api:"required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Entitlements linked to this product
	Entitlements []ProductListResponseEntitlement `json:"entitlements" api:"required"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring" api:"required"`
	// Additional custom data associated with the product
	Metadata map[string]string `json:"metadata" api:"required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id" api:"required"`
	// Tax category associated with the product.
	TaxCategory TaxCategory `json:"tax_category" api:"required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Currency of the price
	Currency Currency `json:"currency" api:"nullable"`
	// Description of the product, optional.
	Description string `json:"description" api:"nullable"`
	// URL of the product image, optional.
	Image string `json:"image" api:"nullable"`
	// Name of the product, optional.
	Name string `json:"name" api:"nullable"`
	// Price of the product, optional.
	//
	// The price is represented in the lowest denomination of the currency. For
	// example:
	//
	// - In USD, a price of `$12.34` would be represented as `1234` (cents).
	// - In JPY, a price of `¥1500` would be represented as `1500` (yen).
	// - In INR, a price of `₹1234.56` would be represented as `123456` (paise).
	//
	// This ensures precision and avoids floating-point rounding errors.
	Price int64 `json:"price" api:"nullable"`
	// Details of the price
	PriceDetail Price `json:"price_detail" api:"nullable"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool                    `json:"tax_inclusive" api:"nullable"`
	JSON         productListResponseJSON `json:"-"`
}

// productListResponseJSON contains the JSON metadata for the struct
// [ProductListResponse]
type productListResponseJSON struct {
	BusinessID   apijson.Field
	CreatedAt    apijson.Field
	Entitlements apijson.Field
	IsRecurring  apijson.Field
	Metadata     apijson.Field
	ProductID    apijson.Field
	TaxCategory  apijson.Field
	UpdatedAt    apijson.Field
	Currency     apijson.Field
	Description  apijson.Field
	Image        apijson.Field
	Name         apijson.Field
	Price        apijson.Field
	PriceDetail  apijson.Field
	TaxInclusive apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseJSON) RawJSON() string {
	return r.raw
}

// Summary of an entitlement attached to a product
type ProductListResponseEntitlement struct {
	ID string `json:"id" api:"required"`
	// Platform-specific configuration for an entitlement. Each variant uses unique
	// field names so `#[serde(untagged)]` can disambiguate correctly.
	IntegrationConfig ProductListResponseEntitlementsIntegrationConfig `json:"integration_config" api:"required"`
	IntegrationType   ProductListResponseEntitlementsIntegrationType   `json:"integration_type" api:"required"`
	Name              string                                           `json:"name" api:"required"`
	Description       string                                           `json:"description" api:"nullable"`
	JSON              productListResponseEntitlementJSON               `json:"-"`
}

// productListResponseEntitlementJSON contains the JSON metadata for the struct
// [ProductListResponseEntitlement]
type productListResponseEntitlementJSON struct {
	ID                apijson.Field
	IntegrationConfig apijson.Field
	IntegrationType   apijson.Field
	Name              apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProductListResponseEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementJSON) RawJSON() string {
	return r.raw
}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
type ProductListResponseEntitlementsIntegrationConfig struct {
	ActivationMessage string `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64  `json:"activations_limit" api:"nullable"`
	ChatID            string `json:"chat_id"`
	// This field can have the runtime type of [[]string].
	DigitalFileIDs   interface{} `json:"digital_file_ids"`
	DurationCount    int64       `json:"duration_count" api:"nullable"`
	DurationInterval string      `json:"duration_interval" api:"nullable"`
	ExternalURL      string      `json:"external_url" api:"nullable"`
	FigmaFileID      string      `json:"figma_file_id"`
	FramerTemplateID string      `json:"framer_template_id"`
	GuildID          string      `json:"guild_id"`
	Instructions     string      `json:"instructions" api:"nullable"`
	NotionTemplateID string      `json:"notion_template_id"`
	// One of: pull, push, admin, maintain, triage
	Permission string                                               `json:"permission"`
	RoleID     string                                               `json:"role_id" api:"nullable"`
	TargetID   string                                               `json:"target_id"`
	JSON       productListResponseEntitlementsIntegrationConfigJSON `json:"-"`
	union      ProductListResponseEntitlementsIntegrationConfigUnion
}

// productListResponseEntitlementsIntegrationConfigJSON contains the JSON metadata
// for the struct [ProductListResponseEntitlementsIntegrationConfig]
type productListResponseEntitlementsIntegrationConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	ChatID            apijson.Field
	DigitalFileIDs    apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	ExternalURL       apijson.Field
	FigmaFileID       apijson.Field
	FramerTemplateID  apijson.Field
	GuildID           apijson.Field
	Instructions      apijson.Field
	NotionTemplateID  apijson.Field
	Permission        apijson.Field
	RoleID            apijson.Field
	TargetID          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r productListResponseEntitlementsIntegrationConfigJSON) RawJSON() string {
	return r.raw
}

func (r *ProductListResponseEntitlementsIntegrationConfig) UnmarshalJSON(data []byte) (err error) {
	*r = ProductListResponseEntitlementsIntegrationConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProductListResponseEntitlementsIntegrationConfigUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProductListResponseEntitlementsIntegrationConfigGitHubConfig],
// [ProductListResponseEntitlementsIntegrationConfigDiscordConfig],
// [ProductListResponseEntitlementsIntegrationConfigTelegramConfig],
// [ProductListResponseEntitlementsIntegrationConfigFigmaConfig],
// [ProductListResponseEntitlementsIntegrationConfigFramerConfig],
// [ProductListResponseEntitlementsIntegrationConfigNotionConfig],
// [ProductListResponseEntitlementsIntegrationConfigDigitalFilesConfig],
// [ProductListResponseEntitlementsIntegrationConfigLicenseKeyConfig].
func (r ProductListResponseEntitlementsIntegrationConfig) AsUnion() ProductListResponseEntitlementsIntegrationConfigUnion {
	return r.union
}

// Platform-specific configuration for an entitlement. Each variant uses unique
// field names so `#[serde(untagged)]` can disambiguate correctly.
//
// Union satisfied by
// [ProductListResponseEntitlementsIntegrationConfigGitHubConfig],
// [ProductListResponseEntitlementsIntegrationConfigDiscordConfig],
// [ProductListResponseEntitlementsIntegrationConfigTelegramConfig],
// [ProductListResponseEntitlementsIntegrationConfigFigmaConfig],
// [ProductListResponseEntitlementsIntegrationConfigFramerConfig],
// [ProductListResponseEntitlementsIntegrationConfigNotionConfig],
// [ProductListResponseEntitlementsIntegrationConfigDigitalFilesConfig] or
// [ProductListResponseEntitlementsIntegrationConfigLicenseKeyConfig].
type ProductListResponseEntitlementsIntegrationConfigUnion interface {
	implementsProductListResponseEntitlementsIntegrationConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProductListResponseEntitlementsIntegrationConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigGitHubConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigDiscordConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigTelegramConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigFigmaConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigFramerConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigNotionConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigDigitalFilesConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProductListResponseEntitlementsIntegrationConfigLicenseKeyConfig{}),
		},
	)
}

type ProductListResponseEntitlementsIntegrationConfigGitHubConfig struct {
	// One of: pull, push, admin, maintain, triage
	Permission string                                                           `json:"permission" api:"required"`
	TargetID   string                                                           `json:"target_id" api:"required"`
	JSON       productListResponseEntitlementsIntegrationConfigGitHubConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigGitHubConfigJSON contains the
// JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigGitHubConfig]
type productListResponseEntitlementsIntegrationConfigGitHubConfigJSON struct {
	Permission  apijson.Field
	TargetID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigGitHubConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigGitHubConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationConfigDiscordConfig struct {
	GuildID string                                                            `json:"guild_id" api:"required"`
	RoleID  string                                                            `json:"role_id" api:"nullable"`
	JSON    productListResponseEntitlementsIntegrationConfigDiscordConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigDiscordConfigJSON contains the
// JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigDiscordConfig]
type productListResponseEntitlementsIntegrationConfigDiscordConfigJSON struct {
	GuildID     apijson.Field
	RoleID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigDiscordConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigDiscordConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigDiscordConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationConfigTelegramConfig struct {
	ChatID string                                                             `json:"chat_id" api:"required"`
	JSON   productListResponseEntitlementsIntegrationConfigTelegramConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigTelegramConfigJSON contains the
// JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigTelegramConfig]
type productListResponseEntitlementsIntegrationConfigTelegramConfigJSON struct {
	ChatID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigTelegramConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigTelegramConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigTelegramConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationConfigFigmaConfig struct {
	FigmaFileID string                                                          `json:"figma_file_id" api:"required"`
	JSON        productListResponseEntitlementsIntegrationConfigFigmaConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigFigmaConfigJSON contains the
// JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigFigmaConfig]
type productListResponseEntitlementsIntegrationConfigFigmaConfigJSON struct {
	FigmaFileID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigFigmaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigFigmaConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigFigmaConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationConfigFramerConfig struct {
	FramerTemplateID string                                                           `json:"framer_template_id" api:"required"`
	JSON             productListResponseEntitlementsIntegrationConfigFramerConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigFramerConfigJSON contains the
// JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigFramerConfig]
type productListResponseEntitlementsIntegrationConfigFramerConfigJSON struct {
	FramerTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigFramerConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigFramerConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigFramerConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationConfigNotionConfig struct {
	NotionTemplateID string                                                           `json:"notion_template_id" api:"required"`
	JSON             productListResponseEntitlementsIntegrationConfigNotionConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigNotionConfigJSON contains the
// JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigNotionConfig]
type productListResponseEntitlementsIntegrationConfigNotionConfigJSON struct {
	NotionTemplateID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigNotionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigNotionConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigNotionConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationConfigDigitalFilesConfig struct {
	DigitalFileIDs []string                                                               `json:"digital_file_ids" api:"required"`
	ExternalURL    string                                                                 `json:"external_url" api:"nullable"`
	Instructions   string                                                                 `json:"instructions" api:"nullable"`
	JSON           productListResponseEntitlementsIntegrationConfigDigitalFilesConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigDigitalFilesConfigJSON contains
// the JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigDigitalFilesConfig]
type productListResponseEntitlementsIntegrationConfigDigitalFilesConfigJSON struct {
	DigitalFileIDs apijson.Field
	ExternalURL    apijson.Field
	Instructions   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigDigitalFilesConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigDigitalFilesConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigDigitalFilesConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationConfigLicenseKeyConfig struct {
	ActivationMessage string                                                               `json:"activation_message" api:"nullable"`
	ActivationsLimit  int64                                                                `json:"activations_limit" api:"nullable"`
	DurationCount     int64                                                                `json:"duration_count" api:"nullable"`
	DurationInterval  string                                                               `json:"duration_interval" api:"nullable"`
	JSON              productListResponseEntitlementsIntegrationConfigLicenseKeyConfigJSON `json:"-"`
}

// productListResponseEntitlementsIntegrationConfigLicenseKeyConfigJSON contains
// the JSON metadata for the struct
// [ProductListResponseEntitlementsIntegrationConfigLicenseKeyConfig]
type productListResponseEntitlementsIntegrationConfigLicenseKeyConfigJSON struct {
	ActivationMessage apijson.Field
	ActivationsLimit  apijson.Field
	DurationCount     apijson.Field
	DurationInterval  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProductListResponseEntitlementsIntegrationConfigLicenseKeyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseEntitlementsIntegrationConfigLicenseKeyConfigJSON) RawJSON() string {
	return r.raw
}

func (r ProductListResponseEntitlementsIntegrationConfigLicenseKeyConfig) implementsProductListResponseEntitlementsIntegrationConfig() {
}

type ProductListResponseEntitlementsIntegrationType string

const (
	ProductListResponseEntitlementsIntegrationTypeDiscord      ProductListResponseEntitlementsIntegrationType = "discord"
	ProductListResponseEntitlementsIntegrationTypeTelegram     ProductListResponseEntitlementsIntegrationType = "telegram"
	ProductListResponseEntitlementsIntegrationTypeGitHub       ProductListResponseEntitlementsIntegrationType = "github"
	ProductListResponseEntitlementsIntegrationTypeFigma        ProductListResponseEntitlementsIntegrationType = "figma"
	ProductListResponseEntitlementsIntegrationTypeFramer       ProductListResponseEntitlementsIntegrationType = "framer"
	ProductListResponseEntitlementsIntegrationTypeNotion       ProductListResponseEntitlementsIntegrationType = "notion"
	ProductListResponseEntitlementsIntegrationTypeDigitalFiles ProductListResponseEntitlementsIntegrationType = "digital_files"
	ProductListResponseEntitlementsIntegrationTypeLicenseKey   ProductListResponseEntitlementsIntegrationType = "license_key"
)

func (r ProductListResponseEntitlementsIntegrationType) IsKnown() bool {
	switch r {
	case ProductListResponseEntitlementsIntegrationTypeDiscord, ProductListResponseEntitlementsIntegrationTypeTelegram, ProductListResponseEntitlementsIntegrationTypeGitHub, ProductListResponseEntitlementsIntegrationTypeFigma, ProductListResponseEntitlementsIntegrationTypeFramer, ProductListResponseEntitlementsIntegrationTypeNotion, ProductListResponseEntitlementsIntegrationTypeDigitalFiles, ProductListResponseEntitlementsIntegrationTypeLicenseKey:
		return true
	}
	return false
}

type ProductUpdateFilesResponse struct {
	FileID string                         `json:"file_id" api:"required" format:"uuid"`
	URL    string                         `json:"url" api:"required"`
	JSON   productUpdateFilesResponseJSON `json:"-"`
}

// productUpdateFilesResponseJSON contains the JSON metadata for the struct
// [ProductUpdateFilesResponse]
type productUpdateFilesResponseJSON struct {
	FileID      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductUpdateFilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productUpdateFilesResponseJSON) RawJSON() string {
	return r.raw
}

type ProductNewParams struct {
	// Name of the product
	Name param.Field[string] `json:"name" api:"required"`
	// Price configuration for the product
	Price param.Field[PriceUnionParam] `json:"price" api:"required"`
	// Tax category applied to this product
	TaxCategory param.Field[TaxCategory] `json:"tax_category" api:"required"`
	// Addons available for subscription product
	Addons param.Field[[]string] `json:"addons"`
	// Brand id for the product, if not provided will default to primary brand
	BrandID param.Field[string] `json:"brand_id"`
	// Optional credit entitlements to attach (max 3)
	CreditEntitlements param.Field[[]AttachCreditEntitlementParam] `json:"credit_entitlements"`
	// Optional description of the product
	Description param.Field[string] `json:"description"`
	// Choose how you would like you digital product delivered
	DigitalProductDelivery param.Field[ProductNewParamsDigitalProductDelivery] `json:"digital_product_delivery"`
	// Optional entitlement IDs to attach to this product (max 20)
	EntitlementIDs param.Field[[]string] `json:"entitlement_ids"`
	// Optional message displayed during license key activation
	LicenseKeyActivationMessage param.Field[string] `json:"license_key_activation_message"`
	// The number of times the license key can be activated. Must be 0 or greater
	LicenseKeyActivationsLimit param.Field[int64] `json:"license_key_activations_limit"`
	// Duration configuration for the license key. Set to null if you don't want the
	// license key to expire. For subscriptions, the lifetime of the license key is
	// tied to the subscription period
	LicenseKeyDuration param.Field[LicenseKeyDurationParam] `json:"license_key_duration"`
	// When true, generates and sends a license key to your customer. Defaults to false
	LicenseKeyEnabled param.Field[bool] `json:"license_key_enabled"`
	// Additional metadata for the product
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how you would like you digital product delivered
type ProductNewParamsDigitalProductDelivery struct {
	// External URL to digital product
	ExternalURL param.Field[string] `json:"external_url"`
	// Instructions to download and use the digital product
	Instructions param.Field[string] `json:"instructions"`
}

func (r ProductNewParamsDigitalProductDelivery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductUpdateParams struct {
	// Available Addons for subscription products
	Addons  param.Field[[]string] `json:"addons"`
	BrandID param.Field[string]   `json:"brand_id"`
	// Credit entitlements to update (replaces all existing when present) Send empty
	// array to remove all, omit field to leave unchanged
	CreditEntitlements param.Field[[]AttachCreditEntitlementParam] `json:"credit_entitlements"`
	// Description of the product, optional and must be at most 1000 characters.
	Description param.Field[string] `json:"description"`
	// Choose how you would like you digital product delivered
	DigitalProductDelivery param.Field[ProductUpdateParamsDigitalProductDelivery] `json:"digital_product_delivery"`
	// Entitlement IDs to attach (replaces all existing when present) Send empty array
	// to remove all, omit field to leave unchanged
	EntitlementIDs param.Field[[]string] `json:"entitlement_ids"`
	// Product image id after its uploaded to S3
	ImageID param.Field[string] `json:"image_id" format:"uuid"`
	// Message sent to the customer upon license key activation.
	//
	// Only applicable if `license_key_enabled` is `true`. This message contains
	// instructions for activating the license key.
	LicenseKeyActivationMessage param.Field[string] `json:"license_key_activation_message"`
	// Limit for the number of activations for the license key.
	//
	// Only applicable if `license_key_enabled` is `true`. Represents the maximum
	// number of times the license key can be activated.
	LicenseKeyActivationsLimit param.Field[int64] `json:"license_key_activations_limit"`
	// Duration of the license key if enabled.
	//
	// Only applicable if `license_key_enabled` is `true`. Represents the duration in
	// days for which the license key is valid.
	LicenseKeyDuration param.Field[LicenseKeyDurationParam] `json:"license_key_duration"`
	// Whether the product requires a license key.
	//
	// If `true`, additional fields related to license key (duration, activations
	// limit, activation message) become applicable.
	LicenseKeyEnabled param.Field[bool] `json:"license_key_enabled"`
	// Additional metadata for the product
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Name of the product, optional and must be at most 100 characters.
	Name param.Field[string] `json:"name"`
	// Price details of the product.
	Price param.Field[PriceUnionParam] `json:"price"`
	// Tax category of the product.
	TaxCategory param.Field[TaxCategory] `json:"tax_category"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose how you would like you digital product delivered
type ProductUpdateParamsDigitalProductDelivery struct {
	// External URL to digital product
	ExternalURL param.Field[string] `json:"external_url"`
	// Uploaded files ids of digital product
	Files param.Field[[]string] `json:"files" format:"uuid"`
	// Instructions to download and use the digital product
	Instructions param.Field[string] `json:"instructions"`
}

func (r ProductUpdateParamsDigitalProductDelivery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductListParams struct {
	// List archived products
	Archived param.Field[bool] `query:"archived"`
	// filter by Brand id
	BrandID param.Field[string] `query:"brand_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter products by pricing type:
	//
	// - `true`: Show only recurring pricing products (e.g. subscriptions)
	// - `false`: Show only one-time price products
	// - `null` or absent: Show both types of products
	Recurring param.Field[bool] `query:"recurring"`
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProductUpdateFilesParams struct {
	FileName param.Field[string] `json:"file_name" api:"required"`
}

func (r ProductUpdateFilesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
