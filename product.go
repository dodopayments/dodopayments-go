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
	return
}

func (r *ProductService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ProductService) Update(ctx context.Context, id string, body ProductUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *ProductService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

func (r *ProductService) UpdateFiles(ctx context.Context, id string, body ProductUpdateFilesParams, opts ...option.RequestOption) (res *ProductUpdateFilesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/files", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AddMeterToPrice struct {
	MeterID string `json:"meter_id" api:"required"`
	// The price per unit in lowest denomination. Must be greater than zero. Supports
	// up to 5 digits before decimal point and 12 decimal places.
	PricePerUnit string `json:"price_per_unit" api:"required"`
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
	Name string              `json:"name" api:"nullable"`
	JSON addMeterToPriceJSON `json:"-"`
}

// addMeterToPriceJSON contains the JSON metadata for the struct [AddMeterToPrice]
type addMeterToPriceJSON struct {
	MeterID             apijson.Field
	PricePerUnit        apijson.Field
	CreditEntitlementID apijson.Field
	Description         apijson.Field
	FreeThreshold       apijson.Field
	MeasurementUnit     apijson.Field
	MeterUnitsPerCredit apijson.Field
	Name                apijson.Field
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
	// The price per unit in lowest denomination. Must be greater than zero. Supports
	// up to 5 digits before decimal point and 12 decimal places.
	PricePerUnit param.Field[string] `json:"price_per_unit" api:"required"`
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
}

func (r AddMeterToPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	CreditEntitlements []ProductCreditEntitlement `json:"credit_entitlements" api:"required"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring" api:"required"`
	// Indicates whether the product requires a license key.
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
	Description            string                        `json:"description" api:"nullable"`
	DigitalProductDelivery ProductDigitalProductDelivery `json:"digital_product_delivery" api:"nullable"`
	// URL of the product image, optional.
	Image string `json:"image" api:"nullable"`
	// Message sent upon license key activation, if applicable.
	LicenseKeyActivationMessage string `json:"license_key_activation_message" api:"nullable"`
	// Limit on the number of activations for the license key, if enabled.
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

// Response struct for credit entitlement mapping
type ProductCreditEntitlement struct {
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
	// Whether new credit grants reduce existing overage
	CreditsReduceOverage bool `json:"credits_reduce_overage" api:"required"`
	// Whether overage is charged at billing
	OverageChargeAtBilling bool `json:"overage_charge_at_billing" api:"required"`
	// Whether overage is enabled
	OverageEnabled bool `json:"overage_enabled" api:"required"`
	// Whether to preserve overage balance when credits reset
	PreserveOverageAtReset bool `json:"preserve_overage_at_reset" api:"required"`
	// Proration behavior for credit grants during plan changes
	ProrationBehavior ProductCreditEntitlementsProrationBehavior `json:"proration_behavior" api:"required"`
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
	TrialCredits string                       `json:"trial_credits" api:"nullable"`
	JSON         productCreditEntitlementJSON `json:"-"`
}

// productCreditEntitlementJSON contains the JSON metadata for the struct
// [ProductCreditEntitlement]
type productCreditEntitlementJSON struct {
	ID                           apijson.Field
	CreditEntitlementID          apijson.Field
	CreditEntitlementName        apijson.Field
	CreditEntitlementUnit        apijson.Field
	CreditsAmount                apijson.Field
	CreditsReduceOverage         apijson.Field
	OverageChargeAtBilling       apijson.Field
	OverageEnabled               apijson.Field
	PreserveOverageAtReset       apijson.Field
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

func (r *ProductCreditEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCreditEntitlementJSON) RawJSON() string {
	return r.raw
}

// Proration behavior for credit grants during plan changes
type ProductCreditEntitlementsProrationBehavior string

const (
	ProductCreditEntitlementsProrationBehaviorProrate   ProductCreditEntitlementsProrationBehavior = "prorate"
	ProductCreditEntitlementsProrationBehaviorNoProrate ProductCreditEntitlementsProrationBehavior = "no_prorate"
)

func (r ProductCreditEntitlementsProrationBehavior) IsKnown() bool {
	switch r {
	case ProductCreditEntitlementsProrationBehaviorProrate, ProductCreditEntitlementsProrationBehaviorNoProrate:
		return true
	}
	return false
}

type ProductDigitalProductDelivery struct {
	// External URL to digital product
	ExternalURL string `json:"external_url" api:"nullable"`
	// Uploaded files ids of digital product
	Files []ProductDigitalProductDeliveryFile `json:"files" api:"nullable"`
	// Instructions to download and use the digital product
	Instructions string                            `json:"instructions" api:"nullable"`
	JSON         productDigitalProductDeliveryJSON `json:"-"`
}

// productDigitalProductDeliveryJSON contains the JSON metadata for the struct
// [ProductDigitalProductDelivery]
type productDigitalProductDeliveryJSON struct {
	ExternalURL  apijson.Field
	Files        apijson.Field
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProductDigitalProductDelivery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productDigitalProductDeliveryJSON) RawJSON() string {
	return r.raw
}

type ProductDigitalProductDeliveryFile struct {
	FileID   string                                `json:"file_id" api:"required" format:"uuid"`
	FileName string                                `json:"file_name" api:"required"`
	URL      string                                `json:"url" api:"required"`
	JSON     productDigitalProductDeliveryFileJSON `json:"-"`
}

// productDigitalProductDeliveryFileJSON contains the JSON metadata for the struct
// [ProductDigitalProductDeliveryFile]
type productDigitalProductDeliveryFileJSON struct {
	FileID      apijson.Field
	FileName    apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductDigitalProductDeliveryFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productDigitalProductDeliveryFileJSON) RawJSON() string {
	return r.raw
}

type ProductListResponse struct {
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id" api:"required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	CreditEntitlements param.Field[[]ProductNewParamsCreditEntitlement] `json:"credit_entitlements"`
	// Optional description of the product
	Description param.Field[string] `json:"description"`
	// Choose how you would like you digital product delivered
	DigitalProductDelivery param.Field[ProductNewParamsDigitalProductDelivery] `json:"digital_product_delivery"`
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

// Request struct for attaching a credit entitlement to a product
type ProductNewParamsCreditEntitlement struct {
	// ID of the credit entitlement to attach
	CreditEntitlementID param.Field[string] `json:"credit_entitlement_id" api:"required"`
	// Number of credits to grant when this product is purchased
	CreditsAmount param.Field[string] `json:"credits_amount" api:"required"`
	// Whether new credit grants reduce existing overage
	CreditsReduceOverage param.Field[bool] `json:"credits_reduce_overage"`
	// Currency for credit-related pricing
	Currency param.Field[Currency] `json:"currency"`
	// Number of days after which credits expire
	ExpiresAfterDays param.Field[int64] `json:"expires_after_days"`
	// Balance threshold percentage for low balance notifications (0-100)
	LowBalanceThresholdPercent param.Field[int64] `json:"low_balance_threshold_percent"`
	// Maximum number of rollover cycles allowed
	MaxRolloverCount param.Field[int64] `json:"max_rollover_count"`
	// Whether overage charges are applied at billing time
	OverageChargeAtBilling param.Field[bool] `json:"overage_charge_at_billing"`
	// Whether overage usage is allowed beyond credit balance
	OverageEnabled param.Field[bool] `json:"overage_enabled"`
	// Maximum amount of overage allowed
	OverageLimit param.Field[string] `json:"overage_limit"`
	// Whether to preserve overage balance when credits reset
	PreserveOverageAtReset param.Field[bool] `json:"preserve_overage_at_reset"`
	// Price per credit unit for purchasing additional credits
	PricePerUnit param.Field[string] `json:"price_per_unit"`
	// Proration behavior for credit grants during plan changes
	ProrationBehavior param.Field[ProductNewParamsCreditEntitlementsProrationBehavior] `json:"proration_behavior"`
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

func (r ProductNewParamsCreditEntitlement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Proration behavior for credit grants during plan changes
type ProductNewParamsCreditEntitlementsProrationBehavior string

const (
	ProductNewParamsCreditEntitlementsProrationBehaviorProrate   ProductNewParamsCreditEntitlementsProrationBehavior = "prorate"
	ProductNewParamsCreditEntitlementsProrationBehaviorNoProrate ProductNewParamsCreditEntitlementsProrationBehavior = "no_prorate"
)

func (r ProductNewParamsCreditEntitlementsProrationBehavior) IsKnown() bool {
	switch r {
	case ProductNewParamsCreditEntitlementsProrationBehaviorProrate, ProductNewParamsCreditEntitlementsProrationBehaviorNoProrate:
		return true
	}
	return false
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
	CreditEntitlements param.Field[[]ProductUpdateParamsCreditEntitlement] `json:"credit_entitlements"`
	// Description of the product, optional and must be at most 1000 characters.
	Description param.Field[string] `json:"description"`
	// Choose how you would like you digital product delivered
	DigitalProductDelivery param.Field[ProductUpdateParamsDigitalProductDelivery] `json:"digital_product_delivery"`
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

// Request struct for attaching a credit entitlement to a product
type ProductUpdateParamsCreditEntitlement struct {
	// ID of the credit entitlement to attach
	CreditEntitlementID param.Field[string] `json:"credit_entitlement_id" api:"required"`
	// Number of credits to grant when this product is purchased
	CreditsAmount param.Field[string] `json:"credits_amount" api:"required"`
	// Whether new credit grants reduce existing overage
	CreditsReduceOverage param.Field[bool] `json:"credits_reduce_overage"`
	// Currency for credit-related pricing
	Currency param.Field[Currency] `json:"currency"`
	// Number of days after which credits expire
	ExpiresAfterDays param.Field[int64] `json:"expires_after_days"`
	// Balance threshold percentage for low balance notifications (0-100)
	LowBalanceThresholdPercent param.Field[int64] `json:"low_balance_threshold_percent"`
	// Maximum number of rollover cycles allowed
	MaxRolloverCount param.Field[int64] `json:"max_rollover_count"`
	// Whether overage charges are applied at billing time
	OverageChargeAtBilling param.Field[bool] `json:"overage_charge_at_billing"`
	// Whether overage usage is allowed beyond credit balance
	OverageEnabled param.Field[bool] `json:"overage_enabled"`
	// Maximum amount of overage allowed
	OverageLimit param.Field[string] `json:"overage_limit"`
	// Whether to preserve overage balance when credits reset
	PreserveOverageAtReset param.Field[bool] `json:"preserve_overage_at_reset"`
	// Price per credit unit for purchasing additional credits
	PricePerUnit param.Field[string] `json:"price_per_unit"`
	// Proration behavior for credit grants during plan changes
	ProrationBehavior param.Field[ProductUpdateParamsCreditEntitlementsProrationBehavior] `json:"proration_behavior"`
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

func (r ProductUpdateParamsCreditEntitlement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Proration behavior for credit grants during plan changes
type ProductUpdateParamsCreditEntitlementsProrationBehavior string

const (
	ProductUpdateParamsCreditEntitlementsProrationBehaviorProrate   ProductUpdateParamsCreditEntitlementsProrationBehavior = "prorate"
	ProductUpdateParamsCreditEntitlementsProrationBehaviorNoProrate ProductUpdateParamsCreditEntitlementsProrationBehavior = "no_prorate"
)

func (r ProductUpdateParamsCreditEntitlementsProrationBehavior) IsKnown() bool {
	switch r {
	case ProductUpdateParamsCreditEntitlementsProrationBehaviorProrate, ProductUpdateParamsCreditEntitlementsProrationBehaviorNoProrate:
		return true
	}
	return false
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
