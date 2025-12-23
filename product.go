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

// ProductService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options []option.RequestOption
	Images  ProductImageService
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r ProductService) {
	r = ProductService{}
	r.Options = opts
	r.Images = NewProductImageService(opts...)
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
	MeterID string `json:"meter_id,required"`
	// The price per unit in lowest denomination. Must be greater than zero. Supports
	// up to 5 digits before decimal point and 12 decimal places.
	PricePerUnit string `json:"price_per_unit,required"`
	// Meter description. Will ignored on Request, but will be shown in response
	Description   string `json:"description,nullable"`
	FreeThreshold int64  `json:"free_threshold,nullable"`
	// Meter measurement unit. Will ignored on Request, but will be shown in response
	MeasurementUnit string `json:"measurement_unit,nullable"`
	// Meter name. Will ignored on Request, but will be shown in response
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MeterID         respjson.Field
		PricePerUnit    respjson.Field
		Description     respjson.Field
		FreeThreshold   respjson.Field
		MeasurementUnit respjson.Field
		Name            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddMeterToPrice) RawJSON() string { return r.JSON.raw }
func (r *AddMeterToPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AddMeterToPrice to a AddMeterToPriceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AddMeterToPriceParam.Overrides()
func (r AddMeterToPrice) ToParam() AddMeterToPriceParam {
	return param.Override[AddMeterToPriceParam](json.RawMessage(r.RawJSON()))
}

// The properties MeterID, PricePerUnit are required.
type AddMeterToPriceParam struct {
	MeterID string `json:"meter_id,required"`
	// The price per unit in lowest denomination. Must be greater than zero. Supports
	// up to 5 digits before decimal point and 12 decimal places.
	PricePerUnit string `json:"price_per_unit,required"`
	// Meter description. Will ignored on Request, but will be shown in response
	Description   param.Opt[string] `json:"description,omitzero"`
	FreeThreshold param.Opt[int64]  `json:"free_threshold,omitzero"`
	// Meter measurement unit. Will ignored on Request, but will be shown in response
	MeasurementUnit param.Opt[string] `json:"measurement_unit,omitzero"`
	// Meter name. Will ignored on Request, but will be shown in response
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AddMeterToPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow AddMeterToPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddMeterToPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LicenseKeyDuration struct {
	Count int64 `json:"count,required"`
	// Any of "Day", "Week", "Month", "Year".
	Interval TimeInterval `json:"interval,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Interval    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LicenseKeyDuration) RawJSON() string { return r.JSON.raw }
func (r *LicenseKeyDuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LicenseKeyDuration to a LicenseKeyDurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LicenseKeyDurationParam.Overrides()
func (r LicenseKeyDuration) ToParam() LicenseKeyDurationParam {
	return param.Override[LicenseKeyDurationParam](json.RawMessage(r.RawJSON()))
}

// The properties Count, Interval are required.
type LicenseKeyDurationParam struct {
	Count int64 `json:"count,required"`
	// Any of "Day", "Week", "Month", "Year".
	Interval TimeInterval `json:"interval,omitzero,required"`
	paramObj
}

func (r LicenseKeyDurationParam) MarshalJSON() (data []byte, err error) {
	type shadow LicenseKeyDurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LicenseKeyDurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PriceUnion contains all possible properties and values from [PriceOneTimePrice],
// [PriceRecurringPrice], [PriceUsageBasedPrice].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PriceUnion struct {
	// This field is from variant [PriceOneTimePrice].
	Currency              Currency `json:"currency"`
	Discount              int64    `json:"discount"`
	Price                 int64    `json:"price"`
	PurchasingPowerParity bool     `json:"purchasing_power_parity"`
	Type                  string   `json:"type"`
	// This field is from variant [PriceOneTimePrice].
	PayWhatYouWant bool `json:"pay_what_you_want"`
	// This field is from variant [PriceOneTimePrice].
	SuggestedPrice        int64 `json:"suggested_price"`
	TaxInclusive          bool  `json:"tax_inclusive"`
	PaymentFrequencyCount int64 `json:"payment_frequency_count"`
	// This field is from variant [PriceRecurringPrice].
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval"`
	SubscriptionPeriodCount  int64        `json:"subscription_period_count"`
	// This field is from variant [PriceRecurringPrice].
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval"`
	// This field is from variant [PriceRecurringPrice].
	TrialPeriodDays int64 `json:"trial_period_days"`
	// This field is from variant [PriceUsageBasedPrice].
	FixedPrice int64 `json:"fixed_price"`
	// This field is from variant [PriceUsageBasedPrice].
	Meters []AddMeterToPrice `json:"meters"`
	JSON   struct {
		Currency                   respjson.Field
		Discount                   respjson.Field
		Price                      respjson.Field
		PurchasingPowerParity      respjson.Field
		Type                       respjson.Field
		PayWhatYouWant             respjson.Field
		SuggestedPrice             respjson.Field
		TaxInclusive               respjson.Field
		PaymentFrequencyCount      respjson.Field
		PaymentFrequencyInterval   respjson.Field
		SubscriptionPeriodCount    respjson.Field
		SubscriptionPeriodInterval respjson.Field
		TrialPeriodDays            respjson.Field
		FixedPrice                 respjson.Field
		Meters                     respjson.Field
		raw                        string
	} `json:"-"`
}

func (u PriceUnion) AsOneTimePrice() (v PriceOneTimePrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PriceUnion) AsRecurringPrice() (v PriceRecurringPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PriceUnion) AsUsageBasedPrice() (v PriceUsageBasedPrice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PriceUnion) RawJSON() string { return u.JSON.raw }

func (r *PriceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PriceUnion to a PriceUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PriceUnionParam.Overrides()
func (r PriceUnion) ToParam() PriceUnionParam {
	return param.Override[PriceUnionParam](json.RawMessage(r.RawJSON()))
}

// One-time price details.
type PriceOneTimePrice struct {
	// The currency in which the payment is made.
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
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount,required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Any of "one_time_price".
	Type string `json:"type,required"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant bool `json:"pay_what_you_want"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice int64 `json:"suggested_price,nullable"`
	// Indicates if the price is tax inclusive.
	TaxInclusive bool `json:"tax_inclusive,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency              respjson.Field
		Discount              respjson.Field
		Price                 respjson.Field
		PurchasingPowerParity respjson.Field
		Type                  respjson.Field
		PayWhatYouWant        respjson.Field
		SuggestedPrice        respjson.Field
		TaxInclusive          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceOneTimePrice) RawJSON() string { return r.JSON.raw }
func (r *PriceOneTimePrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recurring price details.
type PriceRecurringPrice struct {
	// The currency in which the payment is made.
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
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount int64 `json:"payment_frequency_count,required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// The time interval for the subscription period (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval,required"`
	// Any of "recurring_price".
	Type string `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,nullable"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays int64 `json:"trial_period_days"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency                   respjson.Field
		Discount                   respjson.Field
		PaymentFrequencyCount      respjson.Field
		PaymentFrequencyInterval   respjson.Field
		Price                      respjson.Field
		PurchasingPowerParity      respjson.Field
		SubscriptionPeriodCount    respjson.Field
		SubscriptionPeriodInterval respjson.Field
		Type                       respjson.Field
		TaxInclusive               respjson.Field
		TrialPeriodDays            respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceRecurringPrice) RawJSON() string { return r.JSON.raw }
func (r *PriceRecurringPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage Based price details.
type PriceUsageBasedPrice struct {
	// The currency in which the payment is made.
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
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount,required"`
	// The fixed payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	FixedPrice int64 `json:"fixed_price,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount int64 `json:"payment_frequency_count,required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// The time interval for the subscription period (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval,required"`
	// Any of "usage_based_price".
	Type   string            `json:"type,required"`
	Meters []AddMeterToPrice `json:"meters,nullable"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency                   respjson.Field
		Discount                   respjson.Field
		FixedPrice                 respjson.Field
		PaymentFrequencyCount      respjson.Field
		PaymentFrequencyInterval   respjson.Field
		PurchasingPowerParity      respjson.Field
		SubscriptionPeriodCount    respjson.Field
		SubscriptionPeriodInterval respjson.Field
		Type                       respjson.Field
		Meters                     respjson.Field
		TaxInclusive               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceUsageBasedPrice) RawJSON() string { return r.JSON.raw }
func (r *PriceUsageBasedPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PriceUnionParam struct {
	OfOneTimePrice    *PriceOneTimePriceParam    `json:",omitzero,inline"`
	OfRecurringPrice  *PriceRecurringPriceParam  `json:",omitzero,inline"`
	OfUsageBasedPrice *PriceUsageBasedPriceParam `json:",omitzero,inline"`
	paramUnion
}

func (u PriceUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfOneTimePrice, u.OfRecurringPrice, u.OfUsageBasedPrice)
}
func (u *PriceUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PriceUnionParam) asAny() any {
	if !param.IsOmitted(u.OfOneTimePrice) {
		return u.OfOneTimePrice
	} else if !param.IsOmitted(u.OfRecurringPrice) {
		return u.OfRecurringPrice
	} else if !param.IsOmitted(u.OfUsageBasedPrice) {
		return u.OfUsageBasedPrice
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetPayWhatYouWant() *bool {
	if vt := u.OfOneTimePrice; vt != nil && vt.PayWhatYouWant.Valid() {
		return &vt.PayWhatYouWant.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetSuggestedPrice() *int64 {
	if vt := u.OfOneTimePrice; vt != nil && vt.SuggestedPrice.Valid() {
		return &vt.SuggestedPrice.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetTrialPeriodDays() *int64 {
	if vt := u.OfRecurringPrice; vt != nil && vt.TrialPeriodDays.Valid() {
		return &vt.TrialPeriodDays.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetFixedPrice() *int64 {
	if vt := u.OfUsageBasedPrice; vt != nil {
		return &vt.FixedPrice
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetMeters() []AddMeterToPriceParam {
	if vt := u.OfUsageBasedPrice; vt != nil {
		return vt.Meters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetCurrency() *string {
	if vt := u.OfOneTimePrice; vt != nil {
		return (*string)(&vt.Currency)
	} else if vt := u.OfRecurringPrice; vt != nil {
		return (*string)(&vt.Currency)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*string)(&vt.Currency)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetDiscount() *int64 {
	if vt := u.OfOneTimePrice; vt != nil {
		return (*int64)(&vt.Discount)
	} else if vt := u.OfRecurringPrice; vt != nil {
		return (*int64)(&vt.Discount)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*int64)(&vt.Discount)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetPrice() *int64 {
	if vt := u.OfOneTimePrice; vt != nil {
		return (*int64)(&vt.Price)
	} else if vt := u.OfRecurringPrice; vt != nil {
		return (*int64)(&vt.Price)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetPurchasingPowerParity() *bool {
	if vt := u.OfOneTimePrice; vt != nil {
		return (*bool)(&vt.PurchasingPowerParity)
	} else if vt := u.OfRecurringPrice; vt != nil {
		return (*bool)(&vt.PurchasingPowerParity)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*bool)(&vt.PurchasingPowerParity)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetType() *string {
	if vt := u.OfOneTimePrice; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRecurringPrice; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetTaxInclusive() *bool {
	if vt := u.OfOneTimePrice; vt != nil && vt.TaxInclusive.Valid() {
		return &vt.TaxInclusive.Value
	} else if vt := u.OfRecurringPrice; vt != nil && vt.TaxInclusive.Valid() {
		return &vt.TaxInclusive.Value
	} else if vt := u.OfUsageBasedPrice; vt != nil && vt.TaxInclusive.Valid() {
		return &vt.TaxInclusive.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetPaymentFrequencyCount() *int64 {
	if vt := u.OfRecurringPrice; vt != nil {
		return (*int64)(&vt.PaymentFrequencyCount)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*int64)(&vt.PaymentFrequencyCount)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetPaymentFrequencyInterval() *string {
	if vt := u.OfRecurringPrice; vt != nil {
		return (*string)(&vt.PaymentFrequencyInterval)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*string)(&vt.PaymentFrequencyInterval)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetSubscriptionPeriodCount() *int64 {
	if vt := u.OfRecurringPrice; vt != nil {
		return (*int64)(&vt.SubscriptionPeriodCount)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*int64)(&vt.SubscriptionPeriodCount)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PriceUnionParam) GetSubscriptionPeriodInterval() *string {
	if vt := u.OfRecurringPrice; vt != nil {
		return (*string)(&vt.SubscriptionPeriodInterval)
	} else if vt := u.OfUsageBasedPrice; vt != nil {
		return (*string)(&vt.SubscriptionPeriodInterval)
	}
	return nil
}

// One-time price details.
//
// The properties Currency, Discount, Price, PurchasingPowerParity, Type are
// required.
type PriceOneTimePriceParam struct {
	// The currency in which the payment is made.
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
	Currency Currency `json:"currency,omitzero,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount,required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Any of "one_time_price".
	Type string `json:"type,omitzero,required"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice param.Opt[int64] `json:"suggested_price,omitzero"`
	// Indicates if the price is tax inclusive.
	TaxInclusive param.Opt[bool] `json:"tax_inclusive,omitzero"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant param.Opt[bool] `json:"pay_what_you_want,omitzero"`
	paramObj
}

func (r PriceOneTimePriceParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceOneTimePriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceOneTimePriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PriceOneTimePriceParam](
		"type", "one_time_price",
	)
}

// Recurring price details.
//
// The properties Currency, Discount, PaymentFrequencyCount,
// PaymentFrequencyInterval, Price, PurchasingPowerParity, SubscriptionPeriodCount,
// SubscriptionPeriodInterval, Type are required.
type PriceRecurringPriceParam struct {
	// The currency in which the payment is made.
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
	Currency Currency `json:"currency,omitzero,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount int64 `json:"payment_frequency_count,required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,omitzero,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// The time interval for the subscription period (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval,omitzero,required"`
	// Any of "recurring_price".
	Type string `json:"type,omitzero,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Opt[bool] `json:"tax_inclusive,omitzero"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Opt[int64] `json:"trial_period_days,omitzero"`
	paramObj
}

func (r PriceRecurringPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceRecurringPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceRecurringPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PriceRecurringPriceParam](
		"type", "recurring_price",
	)
}

// Usage Based price details.
//
// The properties Currency, Discount, FixedPrice, PaymentFrequencyCount,
// PaymentFrequencyInterval, PurchasingPowerParity, SubscriptionPeriodCount,
// SubscriptionPeriodInterval, Type are required.
type PriceUsageBasedPriceParam struct {
	// The currency in which the payment is made.
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
	Currency Currency `json:"currency,omitzero,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount int64 `json:"discount,required"`
	// The fixed payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	FixedPrice int64 `json:"fixed_price,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount int64 `json:"payment_frequency_count,required"`
	// The time interval for the payment frequency (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,omitzero,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount int64 `json:"subscription_period_count,required"`
	// The time interval for the subscription period (e.g., day, month, year).
	//
	// Any of "Day", "Week", "Month", "Year".
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval,omitzero,required"`
	// Any of "usage_based_price".
	Type string `json:"type,omitzero,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Opt[bool]        `json:"tax_inclusive,omitzero"`
	Meters       []AddMeterToPriceParam `json:"meters,omitzero"`
	paramObj
}

func (r PriceUsageBasedPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow PriceUsageBasedPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PriceUsageBasedPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PriceUsageBasedPriceParam](
		"type", "usage_based_price",
	)
}

type Product struct {
	BrandID string `json:"brand_id,required"`
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id,required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring,required"`
	// Indicates whether the product requires a license key.
	LicenseKeyEnabled bool `json:"license_key_enabled,required"`
	// Additional custom data associated with the product
	Metadata map[string]string `json:"metadata,required"`
	// Pricing information for the product.
	Price PriceUnion `json:"price,required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id,required"`
	// Tax category associated with the product.
	//
	// Any of "digital_products", "saas", "e_book", "edtech".
	TaxCategory TaxCategory `json:"tax_category,required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Available Addons for subscription products
	Addons []string `json:"addons,nullable"`
	// Description of the product, optional.
	Description            string                        `json:"description,nullable"`
	DigitalProductDelivery ProductDigitalProductDelivery `json:"digital_product_delivery,nullable"`
	// URL of the product image, optional.
	Image string `json:"image,nullable"`
	// Message sent upon license key activation, if applicable.
	LicenseKeyActivationMessage string `json:"license_key_activation_message,nullable"`
	// Limit on the number of activations for the license key, if enabled.
	LicenseKeyActivationsLimit int64 `json:"license_key_activations_limit,nullable"`
	// Duration of the license key validity, if enabled.
	LicenseKeyDuration LicenseKeyDuration `json:"license_key_duration,nullable"`
	// Name of the product, optional.
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID                     respjson.Field
		BusinessID                  respjson.Field
		CreatedAt                   respjson.Field
		IsRecurring                 respjson.Field
		LicenseKeyEnabled           respjson.Field
		Metadata                    respjson.Field
		Price                       respjson.Field
		ProductID                   respjson.Field
		TaxCategory                 respjson.Field
		UpdatedAt                   respjson.Field
		Addons                      respjson.Field
		Description                 respjson.Field
		DigitalProductDelivery      respjson.Field
		Image                       respjson.Field
		LicenseKeyActivationMessage respjson.Field
		LicenseKeyActivationsLimit  respjson.Field
		LicenseKeyDuration          respjson.Field
		Name                        respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Product) RawJSON() string { return r.JSON.raw }
func (r *Product) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductDigitalProductDelivery struct {
	// External URL to digital product
	ExternalURL string `json:"external_url,nullable"`
	// Uploaded files ids of digital product
	Files []ProductDigitalProductDeliveryFile `json:"files,nullable"`
	// Instructions to download and use the digital product
	Instructions string `json:"instructions,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalURL  respjson.Field
		Files        respjson.Field
		Instructions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDigitalProductDelivery) RawJSON() string { return r.JSON.raw }
func (r *ProductDigitalProductDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductDigitalProductDeliveryFile struct {
	FileID   string `json:"file_id,required" format:"uuid"`
	FileName string `json:"file_name,required"`
	URL      string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		FileName    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDigitalProductDeliveryFile) RawJSON() string { return r.JSON.raw }
func (r *ProductDigitalProductDeliveryFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductListResponse struct {
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id,required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring,required"`
	// Additional custom data associated with the product
	Metadata map[string]string `json:"metadata,required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id,required"`
	// Tax category associated with the product.
	//
	// Any of "digital_products", "saas", "e_book", "edtech".
	TaxCategory TaxCategory `json:"tax_category,required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Currency of the price
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
	// Description of the product, optional.
	Description string `json:"description,nullable"`
	// URL of the product image, optional.
	Image string `json:"image,nullable"`
	// Name of the product, optional.
	Name string `json:"name,nullable"`
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
	Price int64 `json:"price,nullable"`
	// Details of the price
	PriceDetail PriceUnion `json:"price_detail,nullable"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID   respjson.Field
		CreatedAt    respjson.Field
		IsRecurring  respjson.Field
		Metadata     respjson.Field
		ProductID    respjson.Field
		TaxCategory  respjson.Field
		UpdatedAt    respjson.Field
		Currency     respjson.Field
		Description  respjson.Field
		Image        respjson.Field
		Name         respjson.Field
		Price        respjson.Field
		PriceDetail  respjson.Field
		TaxInclusive respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductUpdateFilesResponse struct {
	FileID string `json:"file_id,required" format:"uuid"`
	URL    string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductUpdateFilesResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductUpdateFilesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductNewParams struct {
	// Name of the product
	Name string `json:"name,required"`
	// Price configuration for the product
	Price PriceUnionParam `json:"price,omitzero,required"`
	// Tax category applied to this product
	//
	// Any of "digital_products", "saas", "e_book", "edtech".
	TaxCategory TaxCategory `json:"tax_category,omitzero,required"`
	// Brand id for the product, if not provided will default to primary brand
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// Optional description of the product
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional message displayed during license key activation
	LicenseKeyActivationMessage param.Opt[string] `json:"license_key_activation_message,omitzero"`
	// The number of times the license key can be activated. Must be 0 or greater
	LicenseKeyActivationsLimit param.Opt[int64] `json:"license_key_activations_limit,omitzero"`
	// When true, generates and sends a license key to your customer. Defaults to false
	LicenseKeyEnabled param.Opt[bool] `json:"license_key_enabled,omitzero"`
	// Addons available for subscription product
	Addons []string `json:"addons,omitzero"`
	// Choose how you would like you digital product delivered
	DigitalProductDelivery ProductNewParamsDigitalProductDelivery `json:"digital_product_delivery,omitzero"`
	// Duration configuration for the license key. Set to null if you don't want the
	// license key to expire. For subscriptions, the lifetime of the license key is
	// tied to the subscription period
	LicenseKeyDuration LicenseKeyDurationParam `json:"license_key_duration,omitzero"`
	// Additional metadata for the product
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Choose how you would like you digital product delivered
type ProductNewParamsDigitalProductDelivery struct {
	// External URL to digital product
	ExternalURL param.Opt[string] `json:"external_url,omitzero"`
	// Instructions to download and use the digital product
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r ProductNewParamsDigitalProductDelivery) MarshalJSON() (data []byte, err error) {
	type shadow ProductNewParamsDigitalProductDelivery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductNewParamsDigitalProductDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductUpdateParams struct {
	BrandID param.Opt[string] `json:"brand_id,omitzero"`
	// Description of the product, optional and must be at most 1000 characters.
	Description param.Opt[string] `json:"description,omitzero"`
	// Product image id after its uploaded to S3
	ImageID param.Opt[string] `json:"image_id,omitzero" format:"uuid"`
	// Message sent to the customer upon license key activation.
	//
	// Only applicable if `license_key_enabled` is `true`. This message contains
	// instructions for activating the license key.
	LicenseKeyActivationMessage param.Opt[string] `json:"license_key_activation_message,omitzero"`
	// Limit for the number of activations for the license key.
	//
	// Only applicable if `license_key_enabled` is `true`. Represents the maximum
	// number of times the license key can be activated.
	LicenseKeyActivationsLimit param.Opt[int64] `json:"license_key_activations_limit,omitzero"`
	// Whether the product requires a license key.
	//
	// If `true`, additional fields related to license key (duration, activations
	// limit, activation message) become applicable.
	LicenseKeyEnabled param.Opt[bool] `json:"license_key_enabled,omitzero"`
	// Name of the product, optional and must be at most 100 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// Available Addons for subscription products
	Addons []string `json:"addons,omitzero"`
	// Choose how you would like you digital product delivered
	DigitalProductDelivery ProductUpdateParamsDigitalProductDelivery `json:"digital_product_delivery,omitzero"`
	// Additional metadata for the product
	Metadata map[string]string `json:"metadata,omitzero"`
	// Duration of the license key if enabled.
	//
	// Only applicable if `license_key_enabled` is `true`. Represents the duration in
	// days for which the license key is valid.
	LicenseKeyDuration LicenseKeyDurationParam `json:"license_key_duration,omitzero"`
	// Price details of the product.
	Price PriceUnionParam `json:"price,omitzero"`
	// Tax category of the product.
	//
	// Any of "digital_products", "saas", "e_book", "edtech".
	TaxCategory TaxCategory `json:"tax_category,omitzero"`
	paramObj
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Choose how you would like you digital product delivered
type ProductUpdateParamsDigitalProductDelivery struct {
	// External URL to digital product
	ExternalURL param.Opt[string] `json:"external_url,omitzero"`
	// Instructions to download and use the digital product
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Uploaded files ids of digital product
	Files []string `json:"files,omitzero" format:"uuid"`
	paramObj
}

func (r ProductUpdateParamsDigitalProductDelivery) MarshalJSON() (data []byte, err error) {
	type shadow ProductUpdateParamsDigitalProductDelivery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductUpdateParamsDigitalProductDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductListParams struct {
	// List archived products
	Archived param.Opt[bool] `query:"archived,omitzero" json:"-"`
	// filter by Brand id
	BrandID param.Opt[string] `query:"brand_id,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter products by pricing type:
	//
	// - `true`: Show only recurring pricing products (e.g. subscriptions)
	// - `false`: Show only one-time price products
	// - `null` or absent: Show both types of products
	Recurring param.Opt[bool] `query:"recurring,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProductUpdateFilesParams struct {
	FileName string `json:"file_name,required"`
	paramObj
}

func (r ProductUpdateFilesParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductUpdateFilesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductUpdateFilesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
