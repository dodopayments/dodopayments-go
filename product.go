// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
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
	Options []option.RequestOption
	Images  *ProductImageService
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r *ProductService) {
	r = &ProductService{}
	r.Options = opts
	r.Images = NewProductImageService(opts...)
	return
}

func (r *ProductService) New(ctx context.Context, body ProductNewParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ProductService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ProductService) Update(ctx context.Context, id string, body ProductUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
	opts = append(r.Options[:], opts...)
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

func (r *ProductService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *ProductService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type LicenseKeyDuration struct {
	Count    int64                  `json:"count,required"`
	Interval TimeInterval           `json:"interval,required"`
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
	Count    param.Field[int64]        `json:"count,required"`
	Interval param.Field[TimeInterval] `json:"interval,required"`
}

func (r LicenseKeyDurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Price struct {
	Currency Currency `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount float64 `json:"discount,required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity bool      `json:"purchasing_power_parity,required"`
	Type                  PriceType `json:"type,required"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant bool `json:"pay_what_you_want"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    int64        `json:"payment_frequency_count"`
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    int64        `json:"subscription_period_count"`
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice int64 `json:"suggested_price,nullable"`
	// Indicates if the price is tax inclusive.
	TaxInclusive bool `json:"tax_inclusive,nullable"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays int64     `json:"trial_period_days"`
	JSON            priceJSON `json:"-"`
	union           PriceUnion
}

// priceJSON contains the JSON metadata for the struct [Price]
type priceJSON struct {
	Currency                   apijson.Field
	Discount                   apijson.Field
	Price                      apijson.Field
	PurchasingPowerParity      apijson.Field
	Type                       apijson.Field
	PayWhatYouWant             apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
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
// [PriceRecurringPrice].
func (r Price) AsUnion() PriceUnion {
	return r.union
}

// Union satisfied by [PriceOneTimePrice] or [PriceRecurringPrice].
type PriceUnion interface {
	implementsPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceOneTimePrice{}),
			DiscriminatorValue: "one_time_price",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceRecurringPrice{}),
			DiscriminatorValue: "recurring_price",
		},
	)
}

type PriceOneTimePrice struct {
	Currency Currency `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount float64 `json:"discount,required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity bool                  `json:"purchasing_power_parity,required"`
	Type                  PriceOneTimePriceType `json:"type,required"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant bool `json:"pay_what_you_want"`
	// A suggested price for the user to pay. This value is only considered if
	// [`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is
	// ignored.
	SuggestedPrice int64 `json:"suggested_price,nullable"`
	// Indicates if the price is tax inclusive.
	TaxInclusive bool                  `json:"tax_inclusive,nullable"`
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

type PriceRecurringPrice struct {
	Currency Currency `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount float64 `json:"discount,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    int64        `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    int64                   `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval TimeInterval            `json:"subscription_period_interval,required"`
	Type                       PriceRecurringPriceType `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,nullable"`
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

type PriceType string

const (
	PriceTypeOneTimePrice   PriceType = "one_time_price"
	PriceTypeRecurringPrice PriceType = "recurring_price"
)

func (r PriceType) IsKnown() bool {
	switch r {
	case PriceTypeOneTimePrice, PriceTypeRecurringPrice:
		return true
	}
	return false
}

type PriceParam struct {
	Currency param.Field[Currency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity param.Field[bool]      `json:"purchasing_power_parity,required"`
	Type                  param.Field[PriceType] `json:"type,required"`
	// Indicates whether the customer can pay any amount they choose. If set to `true`,
	// the [`price`](Self::price) field is the minimum amount.
	PayWhatYouWant param.Field[bool] `json:"pay_what_you_want"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    param.Field[int64]        `json:"payment_frequency_count"`
	PaymentFrequencyInterval param.Field[TimeInterval] `json:"payment_frequency_interval"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    param.Field[int64]        `json:"subscription_period_count"`
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

// Satisfied by [PriceOneTimePriceParam], [PriceRecurringPriceParam], [PriceParam].
type PriceUnionParam interface {
	implementsPriceUnionParam()
}

type PriceOneTimePriceParam struct {
	Currency param.Field[Currency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// The payment amount, in the smallest denomination of the currency (e.g., cents
	// for USD). For example, to charge $1.00, pass `100`.
	//
	// If [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field
	// represents the **minimum** amount the customer must pay.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now.
	PurchasingPowerParity param.Field[bool]                  `json:"purchasing_power_parity,required"`
	Type                  param.Field[PriceOneTimePriceType] `json:"type,required"`
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

type PriceRecurringPriceParam struct {
	Currency param.Field[Currency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    param.Field[int64]        `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval param.Field[TimeInterval] `json:"payment_frequency_interval,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool] `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    param.Field[int64]                   `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval param.Field[TimeInterval]            `json:"subscription_period_interval,required"`
	Type                       param.Field[PriceRecurringPriceType] `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r PriceRecurringPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceRecurringPriceParam) implementsPriceUnionParam() {}

type Product struct {
	BrandID string `json:"brand_id,required"`
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id,required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring,required"`
	// Indicates whether the product requires a license key.
	LicenseKeyEnabled bool  `json:"license_key_enabled,required"`
	Price             Price `json:"price,required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory TaxCategory `json:"tax_category,required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Available Addons for subscription products
	Addons []string `json:"addons,nullable"`
	// Description of the product, optional.
	Description string `json:"description,nullable"`
	// URL of the product image, optional.
	Image string `json:"image,nullable"`
	// Message sent upon license key activation, if applicable.
	LicenseKeyActivationMessage string `json:"license_key_activation_message,nullable"`
	// Limit on the number of activations for the license key, if enabled.
	LicenseKeyActivationsLimit int64              `json:"license_key_activations_limit,nullable"`
	LicenseKeyDuration         LicenseKeyDuration `json:"license_key_duration,nullable"`
	// Name of the product, optional.
	Name string      `json:"name,nullable"`
	JSON productJSON `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	BrandID                     apijson.Field
	BusinessID                  apijson.Field
	CreatedAt                   apijson.Field
	IsRecurring                 apijson.Field
	LicenseKeyEnabled           apijson.Field
	Price                       apijson.Field
	ProductID                   apijson.Field
	TaxCategory                 apijson.Field
	UpdatedAt                   apijson.Field
	Addons                      apijson.Field
	Description                 apijson.Field
	Image                       apijson.Field
	LicenseKeyActivationMessage apijson.Field
	LicenseKeyActivationsLimit  apijson.Field
	LicenseKeyDuration          apijson.Field
	Name                        apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Product) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productJSON) RawJSON() string {
	return r.raw
}

type ProductListResponse struct {
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id,required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring,required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory TaxCategory `json:"tax_category,required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	Currency  Currency  `json:"currency,nullable"`
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
	Price       int64 `json:"price,nullable"`
	PriceDetail Price `json:"price_detail,nullable"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool                    `json:"tax_inclusive,nullable"`
	JSON         productListResponseJSON `json:"-"`
}

// productListResponseJSON contains the JSON metadata for the struct
// [ProductListResponse]
type productListResponseJSON struct {
	BusinessID   apijson.Field
	CreatedAt    apijson.Field
	IsRecurring  apijson.Field
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

type ProductNewParams struct {
	Price param.Field[PriceUnionParam] `json:"price,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category,required"`
	// Addons available for subscription product
	Addons param.Field[[]string] `json:"addons"`
	// Brand id for the product, if not provided will default to primary brand
	BrandID param.Field[string] `json:"brand_id"`
	// Optional description of the product
	Description param.Field[string] `json:"description"`
	// Optional message displayed during license key activation
	LicenseKeyActivationMessage param.Field[string] `json:"license_key_activation_message"`
	// The number of times the license key can be activated. Must be 0 or greater
	LicenseKeyActivationsLimit param.Field[int64]                   `json:"license_key_activations_limit"`
	LicenseKeyDuration         param.Field[LicenseKeyDurationParam] `json:"license_key_duration"`
	// When true, generates and sends a license key to your customer. Defaults to false
	LicenseKeyEnabled param.Field[bool] `json:"license_key_enabled"`
	// Optional name of the product
	Name param.Field[string] `json:"name"`
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductUpdateParams struct {
	// Available Addons for subscription products
	Addons  param.Field[[]string] `json:"addons"`
	BrandID param.Field[string]   `json:"brand_id"`
	// Description of the product, optional and must be at most 1000 characters.
	Description param.Field[string] `json:"description"`
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
	LicenseKeyActivationsLimit param.Field[int64]                   `json:"license_key_activations_limit"`
	LicenseKeyDuration         param.Field[LicenseKeyDurationParam] `json:"license_key_duration"`
	// Whether the product requires a license key.
	//
	// If `true`, additional fields related to license key (duration, activations
	// limit, activation message) become applicable.
	LicenseKeyEnabled param.Field[bool] `json:"license_key_enabled"`
	// Name of the product, optional and must be at most 100 characters.
	Name  param.Field[string]          `json:"name"`
	Price param.Field[PriceUnionParam] `json:"price"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
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
