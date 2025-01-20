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

type Product struct {
	// Unique identifier for the business to which the product belongs.
	BusinessID string `json:"business_id,required"`
	// Timestamp when the product was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Indicates if the product is recurring (e.g., subscriptions).
	IsRecurring bool `json:"is_recurring,required"`
	// Indicates whether the product requires a license key.
	LicenseKeyEnabled bool         `json:"license_key_enabled,required"`
	Price             ProductPrice `json:"price,required"`
	// Unique identifier for the product.
	ProductID string `json:"product_id,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory ProductTaxCategory `json:"tax_category,required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Description of the product, optional.
	Description string `json:"description,nullable"`
	// URL of the product image, optional.
	Image string `json:"image,nullable"`
	// Message sent upon license key activation, if applicable.
	LicenseKeyActivationMessage string `json:"license_key_activation_message,nullable"`
	// Limit on the number of activations for the license key, if enabled.
	LicenseKeyActivationsLimit int64                     `json:"license_key_activations_limit,nullable"`
	LicenseKeyDuration         ProductLicenseKeyDuration `json:"license_key_duration,nullable"`
	// Name of the product, optional.
	Name string      `json:"name,nullable"`
	JSON productJSON `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	BusinessID                  apijson.Field
	CreatedAt                   apijson.Field
	IsRecurring                 apijson.Field
	LicenseKeyEnabled           apijson.Field
	Price                       apijson.Field
	ProductID                   apijson.Field
	TaxCategory                 apijson.Field
	UpdatedAt                   apijson.Field
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

type ProductPrice struct {
	Currency ProductPriceCurrency `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount float64 `json:"discount,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool             `json:"purchasing_power_parity,required"`
	Type                  ProductPriceType `json:"type,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    int64                                `json:"payment_frequency_count"`
	PaymentFrequencyInterval ProductPricePaymentFrequencyInterval `json:"payment_frequency_interval"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    int64                                  `json:"subscription_period_count"`
	SubscriptionPeriodInterval ProductPriceSubscriptionPeriodInterval `json:"subscription_period_interval"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,nullable"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays int64            `json:"trial_period_days"`
	JSON            productPriceJSON `json:"-"`
	union           ProductPriceUnion
}

// productPriceJSON contains the JSON metadata for the struct [ProductPrice]
type productPriceJSON struct {
	Currency                   apijson.Field
	Discount                   apijson.Field
	Price                      apijson.Field
	PurchasingPowerParity      apijson.Field
	Type                       apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	TaxInclusive               apijson.Field
	TrialPeriodDays            apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r productPriceJSON) RawJSON() string {
	return r.raw
}

func (r *ProductPrice) UnmarshalJSON(data []byte) (err error) {
	*r = ProductPrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProductPriceUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ProductPriceOneTimePrice],
// [ProductPriceRecurringPrice].
func (r ProductPrice) AsUnion() ProductPriceUnion {
	return r.union
}

// Union satisfied by [ProductPriceOneTimePrice] or [ProductPriceRecurringPrice].
type ProductPriceUnion interface {
	implementsProductPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProductPriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductPriceOneTimePrice{}),
			DiscriminatorValue: "one_time_price",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductPriceRecurringPrice{}),
			DiscriminatorValue: "recurring_price",
		},
	)
}

type ProductPriceOneTimePrice struct {
	Currency ProductPriceOneTimePriceCurrency `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount float64 `json:"discount,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool                         `json:"purchasing_power_parity,required"`
	Type                  ProductPriceOneTimePriceType `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool                         `json:"tax_inclusive,nullable"`
	JSON         productPriceOneTimePriceJSON `json:"-"`
}

// productPriceOneTimePriceJSON contains the JSON metadata for the struct
// [ProductPriceOneTimePrice]
type productPriceOneTimePriceJSON struct {
	Currency              apijson.Field
	Discount              apijson.Field
	Price                 apijson.Field
	PurchasingPowerParity apijson.Field
	Type                  apijson.Field
	TaxInclusive          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProductPriceOneTimePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productPriceOneTimePriceJSON) RawJSON() string {
	return r.raw
}

func (r ProductPriceOneTimePrice) implementsProductPrice() {}

type ProductPriceOneTimePriceCurrency string

const (
	ProductPriceOneTimePriceCurrencyAed ProductPriceOneTimePriceCurrency = "AED"
	ProductPriceOneTimePriceCurrencyAll ProductPriceOneTimePriceCurrency = "ALL"
	ProductPriceOneTimePriceCurrencyAmd ProductPriceOneTimePriceCurrency = "AMD"
	ProductPriceOneTimePriceCurrencyAng ProductPriceOneTimePriceCurrency = "ANG"
	ProductPriceOneTimePriceCurrencyAoa ProductPriceOneTimePriceCurrency = "AOA"
	ProductPriceOneTimePriceCurrencyArs ProductPriceOneTimePriceCurrency = "ARS"
	ProductPriceOneTimePriceCurrencyAud ProductPriceOneTimePriceCurrency = "AUD"
	ProductPriceOneTimePriceCurrencyAwg ProductPriceOneTimePriceCurrency = "AWG"
	ProductPriceOneTimePriceCurrencyAzn ProductPriceOneTimePriceCurrency = "AZN"
	ProductPriceOneTimePriceCurrencyBam ProductPriceOneTimePriceCurrency = "BAM"
	ProductPriceOneTimePriceCurrencyBbd ProductPriceOneTimePriceCurrency = "BBD"
	ProductPriceOneTimePriceCurrencyBdt ProductPriceOneTimePriceCurrency = "BDT"
	ProductPriceOneTimePriceCurrencyBgn ProductPriceOneTimePriceCurrency = "BGN"
	ProductPriceOneTimePriceCurrencyBhd ProductPriceOneTimePriceCurrency = "BHD"
	ProductPriceOneTimePriceCurrencyBif ProductPriceOneTimePriceCurrency = "BIF"
	ProductPriceOneTimePriceCurrencyBmd ProductPriceOneTimePriceCurrency = "BMD"
	ProductPriceOneTimePriceCurrencyBnd ProductPriceOneTimePriceCurrency = "BND"
	ProductPriceOneTimePriceCurrencyBob ProductPriceOneTimePriceCurrency = "BOB"
	ProductPriceOneTimePriceCurrencyBrl ProductPriceOneTimePriceCurrency = "BRL"
	ProductPriceOneTimePriceCurrencyBsd ProductPriceOneTimePriceCurrency = "BSD"
	ProductPriceOneTimePriceCurrencyBwp ProductPriceOneTimePriceCurrency = "BWP"
	ProductPriceOneTimePriceCurrencyByn ProductPriceOneTimePriceCurrency = "BYN"
	ProductPriceOneTimePriceCurrencyBzd ProductPriceOneTimePriceCurrency = "BZD"
	ProductPriceOneTimePriceCurrencyCad ProductPriceOneTimePriceCurrency = "CAD"
	ProductPriceOneTimePriceCurrencyChf ProductPriceOneTimePriceCurrency = "CHF"
	ProductPriceOneTimePriceCurrencyClp ProductPriceOneTimePriceCurrency = "CLP"
	ProductPriceOneTimePriceCurrencyCny ProductPriceOneTimePriceCurrency = "CNY"
	ProductPriceOneTimePriceCurrencyCop ProductPriceOneTimePriceCurrency = "COP"
	ProductPriceOneTimePriceCurrencyCrc ProductPriceOneTimePriceCurrency = "CRC"
	ProductPriceOneTimePriceCurrencyCup ProductPriceOneTimePriceCurrency = "CUP"
	ProductPriceOneTimePriceCurrencyCve ProductPriceOneTimePriceCurrency = "CVE"
	ProductPriceOneTimePriceCurrencyCzk ProductPriceOneTimePriceCurrency = "CZK"
	ProductPriceOneTimePriceCurrencyDjf ProductPriceOneTimePriceCurrency = "DJF"
	ProductPriceOneTimePriceCurrencyDkk ProductPriceOneTimePriceCurrency = "DKK"
	ProductPriceOneTimePriceCurrencyDop ProductPriceOneTimePriceCurrency = "DOP"
	ProductPriceOneTimePriceCurrencyDzd ProductPriceOneTimePriceCurrency = "DZD"
	ProductPriceOneTimePriceCurrencyEgp ProductPriceOneTimePriceCurrency = "EGP"
	ProductPriceOneTimePriceCurrencyEtb ProductPriceOneTimePriceCurrency = "ETB"
	ProductPriceOneTimePriceCurrencyEur ProductPriceOneTimePriceCurrency = "EUR"
	ProductPriceOneTimePriceCurrencyFjd ProductPriceOneTimePriceCurrency = "FJD"
	ProductPriceOneTimePriceCurrencyFkp ProductPriceOneTimePriceCurrency = "FKP"
	ProductPriceOneTimePriceCurrencyGbp ProductPriceOneTimePriceCurrency = "GBP"
	ProductPriceOneTimePriceCurrencyGel ProductPriceOneTimePriceCurrency = "GEL"
	ProductPriceOneTimePriceCurrencyGhs ProductPriceOneTimePriceCurrency = "GHS"
	ProductPriceOneTimePriceCurrencyGip ProductPriceOneTimePriceCurrency = "GIP"
	ProductPriceOneTimePriceCurrencyGmd ProductPriceOneTimePriceCurrency = "GMD"
	ProductPriceOneTimePriceCurrencyGnf ProductPriceOneTimePriceCurrency = "GNF"
	ProductPriceOneTimePriceCurrencyGtq ProductPriceOneTimePriceCurrency = "GTQ"
	ProductPriceOneTimePriceCurrencyGyd ProductPriceOneTimePriceCurrency = "GYD"
	ProductPriceOneTimePriceCurrencyHkd ProductPriceOneTimePriceCurrency = "HKD"
	ProductPriceOneTimePriceCurrencyHnl ProductPriceOneTimePriceCurrency = "HNL"
	ProductPriceOneTimePriceCurrencyHrk ProductPriceOneTimePriceCurrency = "HRK"
	ProductPriceOneTimePriceCurrencyHtg ProductPriceOneTimePriceCurrency = "HTG"
	ProductPriceOneTimePriceCurrencyHuf ProductPriceOneTimePriceCurrency = "HUF"
	ProductPriceOneTimePriceCurrencyIdr ProductPriceOneTimePriceCurrency = "IDR"
	ProductPriceOneTimePriceCurrencyIls ProductPriceOneTimePriceCurrency = "ILS"
	ProductPriceOneTimePriceCurrencyInr ProductPriceOneTimePriceCurrency = "INR"
	ProductPriceOneTimePriceCurrencyIqd ProductPriceOneTimePriceCurrency = "IQD"
	ProductPriceOneTimePriceCurrencyJmd ProductPriceOneTimePriceCurrency = "JMD"
	ProductPriceOneTimePriceCurrencyJod ProductPriceOneTimePriceCurrency = "JOD"
	ProductPriceOneTimePriceCurrencyJpy ProductPriceOneTimePriceCurrency = "JPY"
	ProductPriceOneTimePriceCurrencyKes ProductPriceOneTimePriceCurrency = "KES"
	ProductPriceOneTimePriceCurrencyKgs ProductPriceOneTimePriceCurrency = "KGS"
	ProductPriceOneTimePriceCurrencyKhr ProductPriceOneTimePriceCurrency = "KHR"
	ProductPriceOneTimePriceCurrencyKmf ProductPriceOneTimePriceCurrency = "KMF"
	ProductPriceOneTimePriceCurrencyKrw ProductPriceOneTimePriceCurrency = "KRW"
	ProductPriceOneTimePriceCurrencyKwd ProductPriceOneTimePriceCurrency = "KWD"
	ProductPriceOneTimePriceCurrencyKyd ProductPriceOneTimePriceCurrency = "KYD"
	ProductPriceOneTimePriceCurrencyKzt ProductPriceOneTimePriceCurrency = "KZT"
	ProductPriceOneTimePriceCurrencyLak ProductPriceOneTimePriceCurrency = "LAK"
	ProductPriceOneTimePriceCurrencyLbp ProductPriceOneTimePriceCurrency = "LBP"
	ProductPriceOneTimePriceCurrencyLkr ProductPriceOneTimePriceCurrency = "LKR"
	ProductPriceOneTimePriceCurrencyLrd ProductPriceOneTimePriceCurrency = "LRD"
	ProductPriceOneTimePriceCurrencyLsl ProductPriceOneTimePriceCurrency = "LSL"
	ProductPriceOneTimePriceCurrencyLyd ProductPriceOneTimePriceCurrency = "LYD"
	ProductPriceOneTimePriceCurrencyMad ProductPriceOneTimePriceCurrency = "MAD"
	ProductPriceOneTimePriceCurrencyMdl ProductPriceOneTimePriceCurrency = "MDL"
	ProductPriceOneTimePriceCurrencyMga ProductPriceOneTimePriceCurrency = "MGA"
	ProductPriceOneTimePriceCurrencyMkd ProductPriceOneTimePriceCurrency = "MKD"
	ProductPriceOneTimePriceCurrencyMmk ProductPriceOneTimePriceCurrency = "MMK"
	ProductPriceOneTimePriceCurrencyMnt ProductPriceOneTimePriceCurrency = "MNT"
	ProductPriceOneTimePriceCurrencyMop ProductPriceOneTimePriceCurrency = "MOP"
	ProductPriceOneTimePriceCurrencyMru ProductPriceOneTimePriceCurrency = "MRU"
	ProductPriceOneTimePriceCurrencyMur ProductPriceOneTimePriceCurrency = "MUR"
	ProductPriceOneTimePriceCurrencyMvr ProductPriceOneTimePriceCurrency = "MVR"
	ProductPriceOneTimePriceCurrencyMwk ProductPriceOneTimePriceCurrency = "MWK"
	ProductPriceOneTimePriceCurrencyMxn ProductPriceOneTimePriceCurrency = "MXN"
	ProductPriceOneTimePriceCurrencyMyr ProductPriceOneTimePriceCurrency = "MYR"
	ProductPriceOneTimePriceCurrencyMzn ProductPriceOneTimePriceCurrency = "MZN"
	ProductPriceOneTimePriceCurrencyNad ProductPriceOneTimePriceCurrency = "NAD"
	ProductPriceOneTimePriceCurrencyNgn ProductPriceOneTimePriceCurrency = "NGN"
	ProductPriceOneTimePriceCurrencyNio ProductPriceOneTimePriceCurrency = "NIO"
	ProductPriceOneTimePriceCurrencyNok ProductPriceOneTimePriceCurrency = "NOK"
	ProductPriceOneTimePriceCurrencyNpr ProductPriceOneTimePriceCurrency = "NPR"
	ProductPriceOneTimePriceCurrencyNzd ProductPriceOneTimePriceCurrency = "NZD"
	ProductPriceOneTimePriceCurrencyOmr ProductPriceOneTimePriceCurrency = "OMR"
	ProductPriceOneTimePriceCurrencyPab ProductPriceOneTimePriceCurrency = "PAB"
	ProductPriceOneTimePriceCurrencyPen ProductPriceOneTimePriceCurrency = "PEN"
	ProductPriceOneTimePriceCurrencyPgk ProductPriceOneTimePriceCurrency = "PGK"
	ProductPriceOneTimePriceCurrencyPhp ProductPriceOneTimePriceCurrency = "PHP"
	ProductPriceOneTimePriceCurrencyPkr ProductPriceOneTimePriceCurrency = "PKR"
	ProductPriceOneTimePriceCurrencyPln ProductPriceOneTimePriceCurrency = "PLN"
	ProductPriceOneTimePriceCurrencyPyg ProductPriceOneTimePriceCurrency = "PYG"
	ProductPriceOneTimePriceCurrencyQar ProductPriceOneTimePriceCurrency = "QAR"
	ProductPriceOneTimePriceCurrencyRon ProductPriceOneTimePriceCurrency = "RON"
	ProductPriceOneTimePriceCurrencyRsd ProductPriceOneTimePriceCurrency = "RSD"
	ProductPriceOneTimePriceCurrencyRub ProductPriceOneTimePriceCurrency = "RUB"
	ProductPriceOneTimePriceCurrencyRwf ProductPriceOneTimePriceCurrency = "RWF"
	ProductPriceOneTimePriceCurrencySar ProductPriceOneTimePriceCurrency = "SAR"
	ProductPriceOneTimePriceCurrencySbd ProductPriceOneTimePriceCurrency = "SBD"
	ProductPriceOneTimePriceCurrencyScr ProductPriceOneTimePriceCurrency = "SCR"
	ProductPriceOneTimePriceCurrencySek ProductPriceOneTimePriceCurrency = "SEK"
	ProductPriceOneTimePriceCurrencySgd ProductPriceOneTimePriceCurrency = "SGD"
	ProductPriceOneTimePriceCurrencyShp ProductPriceOneTimePriceCurrency = "SHP"
	ProductPriceOneTimePriceCurrencySle ProductPriceOneTimePriceCurrency = "SLE"
	ProductPriceOneTimePriceCurrencySll ProductPriceOneTimePriceCurrency = "SLL"
	ProductPriceOneTimePriceCurrencySos ProductPriceOneTimePriceCurrency = "SOS"
	ProductPriceOneTimePriceCurrencySrd ProductPriceOneTimePriceCurrency = "SRD"
	ProductPriceOneTimePriceCurrencySsp ProductPriceOneTimePriceCurrency = "SSP"
	ProductPriceOneTimePriceCurrencyStn ProductPriceOneTimePriceCurrency = "STN"
	ProductPriceOneTimePriceCurrencySvc ProductPriceOneTimePriceCurrency = "SVC"
	ProductPriceOneTimePriceCurrencySzl ProductPriceOneTimePriceCurrency = "SZL"
	ProductPriceOneTimePriceCurrencyThb ProductPriceOneTimePriceCurrency = "THB"
	ProductPriceOneTimePriceCurrencyTnd ProductPriceOneTimePriceCurrency = "TND"
	ProductPriceOneTimePriceCurrencyTop ProductPriceOneTimePriceCurrency = "TOP"
	ProductPriceOneTimePriceCurrencyTry ProductPriceOneTimePriceCurrency = "TRY"
	ProductPriceOneTimePriceCurrencyTtd ProductPriceOneTimePriceCurrency = "TTD"
	ProductPriceOneTimePriceCurrencyTwd ProductPriceOneTimePriceCurrency = "TWD"
	ProductPriceOneTimePriceCurrencyTzs ProductPriceOneTimePriceCurrency = "TZS"
	ProductPriceOneTimePriceCurrencyUah ProductPriceOneTimePriceCurrency = "UAH"
	ProductPriceOneTimePriceCurrencyUgx ProductPriceOneTimePriceCurrency = "UGX"
	ProductPriceOneTimePriceCurrencyUsd ProductPriceOneTimePriceCurrency = "USD"
	ProductPriceOneTimePriceCurrencyUyu ProductPriceOneTimePriceCurrency = "UYU"
	ProductPriceOneTimePriceCurrencyUzs ProductPriceOneTimePriceCurrency = "UZS"
	ProductPriceOneTimePriceCurrencyVes ProductPriceOneTimePriceCurrency = "VES"
	ProductPriceOneTimePriceCurrencyVnd ProductPriceOneTimePriceCurrency = "VND"
	ProductPriceOneTimePriceCurrencyVuv ProductPriceOneTimePriceCurrency = "VUV"
	ProductPriceOneTimePriceCurrencyWst ProductPriceOneTimePriceCurrency = "WST"
	ProductPriceOneTimePriceCurrencyXaf ProductPriceOneTimePriceCurrency = "XAF"
	ProductPriceOneTimePriceCurrencyXcd ProductPriceOneTimePriceCurrency = "XCD"
	ProductPriceOneTimePriceCurrencyXof ProductPriceOneTimePriceCurrency = "XOF"
	ProductPriceOneTimePriceCurrencyXpf ProductPriceOneTimePriceCurrency = "XPF"
	ProductPriceOneTimePriceCurrencyYer ProductPriceOneTimePriceCurrency = "YER"
	ProductPriceOneTimePriceCurrencyZar ProductPriceOneTimePriceCurrency = "ZAR"
	ProductPriceOneTimePriceCurrencyZmw ProductPriceOneTimePriceCurrency = "ZMW"
)

func (r ProductPriceOneTimePriceCurrency) IsKnown() bool {
	switch r {
	case ProductPriceOneTimePriceCurrencyAed, ProductPriceOneTimePriceCurrencyAll, ProductPriceOneTimePriceCurrencyAmd, ProductPriceOneTimePriceCurrencyAng, ProductPriceOneTimePriceCurrencyAoa, ProductPriceOneTimePriceCurrencyArs, ProductPriceOneTimePriceCurrencyAud, ProductPriceOneTimePriceCurrencyAwg, ProductPriceOneTimePriceCurrencyAzn, ProductPriceOneTimePriceCurrencyBam, ProductPriceOneTimePriceCurrencyBbd, ProductPriceOneTimePriceCurrencyBdt, ProductPriceOneTimePriceCurrencyBgn, ProductPriceOneTimePriceCurrencyBhd, ProductPriceOneTimePriceCurrencyBif, ProductPriceOneTimePriceCurrencyBmd, ProductPriceOneTimePriceCurrencyBnd, ProductPriceOneTimePriceCurrencyBob, ProductPriceOneTimePriceCurrencyBrl, ProductPriceOneTimePriceCurrencyBsd, ProductPriceOneTimePriceCurrencyBwp, ProductPriceOneTimePriceCurrencyByn, ProductPriceOneTimePriceCurrencyBzd, ProductPriceOneTimePriceCurrencyCad, ProductPriceOneTimePriceCurrencyChf, ProductPriceOneTimePriceCurrencyClp, ProductPriceOneTimePriceCurrencyCny, ProductPriceOneTimePriceCurrencyCop, ProductPriceOneTimePriceCurrencyCrc, ProductPriceOneTimePriceCurrencyCup, ProductPriceOneTimePriceCurrencyCve, ProductPriceOneTimePriceCurrencyCzk, ProductPriceOneTimePriceCurrencyDjf, ProductPriceOneTimePriceCurrencyDkk, ProductPriceOneTimePriceCurrencyDop, ProductPriceOneTimePriceCurrencyDzd, ProductPriceOneTimePriceCurrencyEgp, ProductPriceOneTimePriceCurrencyEtb, ProductPriceOneTimePriceCurrencyEur, ProductPriceOneTimePriceCurrencyFjd, ProductPriceOneTimePriceCurrencyFkp, ProductPriceOneTimePriceCurrencyGbp, ProductPriceOneTimePriceCurrencyGel, ProductPriceOneTimePriceCurrencyGhs, ProductPriceOneTimePriceCurrencyGip, ProductPriceOneTimePriceCurrencyGmd, ProductPriceOneTimePriceCurrencyGnf, ProductPriceOneTimePriceCurrencyGtq, ProductPriceOneTimePriceCurrencyGyd, ProductPriceOneTimePriceCurrencyHkd, ProductPriceOneTimePriceCurrencyHnl, ProductPriceOneTimePriceCurrencyHrk, ProductPriceOneTimePriceCurrencyHtg, ProductPriceOneTimePriceCurrencyHuf, ProductPriceOneTimePriceCurrencyIdr, ProductPriceOneTimePriceCurrencyIls, ProductPriceOneTimePriceCurrencyInr, ProductPriceOneTimePriceCurrencyIqd, ProductPriceOneTimePriceCurrencyJmd, ProductPriceOneTimePriceCurrencyJod, ProductPriceOneTimePriceCurrencyJpy, ProductPriceOneTimePriceCurrencyKes, ProductPriceOneTimePriceCurrencyKgs, ProductPriceOneTimePriceCurrencyKhr, ProductPriceOneTimePriceCurrencyKmf, ProductPriceOneTimePriceCurrencyKrw, ProductPriceOneTimePriceCurrencyKwd, ProductPriceOneTimePriceCurrencyKyd, ProductPriceOneTimePriceCurrencyKzt, ProductPriceOneTimePriceCurrencyLak, ProductPriceOneTimePriceCurrencyLbp, ProductPriceOneTimePriceCurrencyLkr, ProductPriceOneTimePriceCurrencyLrd, ProductPriceOneTimePriceCurrencyLsl, ProductPriceOneTimePriceCurrencyLyd, ProductPriceOneTimePriceCurrencyMad, ProductPriceOneTimePriceCurrencyMdl, ProductPriceOneTimePriceCurrencyMga, ProductPriceOneTimePriceCurrencyMkd, ProductPriceOneTimePriceCurrencyMmk, ProductPriceOneTimePriceCurrencyMnt, ProductPriceOneTimePriceCurrencyMop, ProductPriceOneTimePriceCurrencyMru, ProductPriceOneTimePriceCurrencyMur, ProductPriceOneTimePriceCurrencyMvr, ProductPriceOneTimePriceCurrencyMwk, ProductPriceOneTimePriceCurrencyMxn, ProductPriceOneTimePriceCurrencyMyr, ProductPriceOneTimePriceCurrencyMzn, ProductPriceOneTimePriceCurrencyNad, ProductPriceOneTimePriceCurrencyNgn, ProductPriceOneTimePriceCurrencyNio, ProductPriceOneTimePriceCurrencyNok, ProductPriceOneTimePriceCurrencyNpr, ProductPriceOneTimePriceCurrencyNzd, ProductPriceOneTimePriceCurrencyOmr, ProductPriceOneTimePriceCurrencyPab, ProductPriceOneTimePriceCurrencyPen, ProductPriceOneTimePriceCurrencyPgk, ProductPriceOneTimePriceCurrencyPhp, ProductPriceOneTimePriceCurrencyPkr, ProductPriceOneTimePriceCurrencyPln, ProductPriceOneTimePriceCurrencyPyg, ProductPriceOneTimePriceCurrencyQar, ProductPriceOneTimePriceCurrencyRon, ProductPriceOneTimePriceCurrencyRsd, ProductPriceOneTimePriceCurrencyRub, ProductPriceOneTimePriceCurrencyRwf, ProductPriceOneTimePriceCurrencySar, ProductPriceOneTimePriceCurrencySbd, ProductPriceOneTimePriceCurrencyScr, ProductPriceOneTimePriceCurrencySek, ProductPriceOneTimePriceCurrencySgd, ProductPriceOneTimePriceCurrencyShp, ProductPriceOneTimePriceCurrencySle, ProductPriceOneTimePriceCurrencySll, ProductPriceOneTimePriceCurrencySos, ProductPriceOneTimePriceCurrencySrd, ProductPriceOneTimePriceCurrencySsp, ProductPriceOneTimePriceCurrencyStn, ProductPriceOneTimePriceCurrencySvc, ProductPriceOneTimePriceCurrencySzl, ProductPriceOneTimePriceCurrencyThb, ProductPriceOneTimePriceCurrencyTnd, ProductPriceOneTimePriceCurrencyTop, ProductPriceOneTimePriceCurrencyTry, ProductPriceOneTimePriceCurrencyTtd, ProductPriceOneTimePriceCurrencyTwd, ProductPriceOneTimePriceCurrencyTzs, ProductPriceOneTimePriceCurrencyUah, ProductPriceOneTimePriceCurrencyUgx, ProductPriceOneTimePriceCurrencyUsd, ProductPriceOneTimePriceCurrencyUyu, ProductPriceOneTimePriceCurrencyUzs, ProductPriceOneTimePriceCurrencyVes, ProductPriceOneTimePriceCurrencyVnd, ProductPriceOneTimePriceCurrencyVuv, ProductPriceOneTimePriceCurrencyWst, ProductPriceOneTimePriceCurrencyXaf, ProductPriceOneTimePriceCurrencyXcd, ProductPriceOneTimePriceCurrencyXof, ProductPriceOneTimePriceCurrencyXpf, ProductPriceOneTimePriceCurrencyYer, ProductPriceOneTimePriceCurrencyZar, ProductPriceOneTimePriceCurrencyZmw:
		return true
	}
	return false
}

type ProductPriceOneTimePriceType string

const (
	ProductPriceOneTimePriceTypeOneTimePrice ProductPriceOneTimePriceType = "one_time_price"
)

func (r ProductPriceOneTimePriceType) IsKnown() bool {
	switch r {
	case ProductPriceOneTimePriceTypeOneTimePrice:
		return true
	}
	return false
}

type ProductPriceRecurringPrice struct {
	Currency ProductPriceRecurringPriceCurrency `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount float64 `json:"discount,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    int64                                              `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval ProductPriceRecurringPricePaymentFrequencyInterval `json:"payment_frequency_interval,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price int64 `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity bool `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    int64                                                `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval ProductPriceRecurringPriceSubscriptionPeriodInterval `json:"subscription_period_interval,required"`
	Type                       ProductPriceRecurringPriceType                       `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,nullable"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays int64                          `json:"trial_period_days"`
	JSON            productPriceRecurringPriceJSON `json:"-"`
}

// productPriceRecurringPriceJSON contains the JSON metadata for the struct
// [ProductPriceRecurringPrice]
type productPriceRecurringPriceJSON struct {
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

func (r *ProductPriceRecurringPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productPriceRecurringPriceJSON) RawJSON() string {
	return r.raw
}

func (r ProductPriceRecurringPrice) implementsProductPrice() {}

type ProductPriceRecurringPriceCurrency string

const (
	ProductPriceRecurringPriceCurrencyAed ProductPriceRecurringPriceCurrency = "AED"
	ProductPriceRecurringPriceCurrencyAll ProductPriceRecurringPriceCurrency = "ALL"
	ProductPriceRecurringPriceCurrencyAmd ProductPriceRecurringPriceCurrency = "AMD"
	ProductPriceRecurringPriceCurrencyAng ProductPriceRecurringPriceCurrency = "ANG"
	ProductPriceRecurringPriceCurrencyAoa ProductPriceRecurringPriceCurrency = "AOA"
	ProductPriceRecurringPriceCurrencyArs ProductPriceRecurringPriceCurrency = "ARS"
	ProductPriceRecurringPriceCurrencyAud ProductPriceRecurringPriceCurrency = "AUD"
	ProductPriceRecurringPriceCurrencyAwg ProductPriceRecurringPriceCurrency = "AWG"
	ProductPriceRecurringPriceCurrencyAzn ProductPriceRecurringPriceCurrency = "AZN"
	ProductPriceRecurringPriceCurrencyBam ProductPriceRecurringPriceCurrency = "BAM"
	ProductPriceRecurringPriceCurrencyBbd ProductPriceRecurringPriceCurrency = "BBD"
	ProductPriceRecurringPriceCurrencyBdt ProductPriceRecurringPriceCurrency = "BDT"
	ProductPriceRecurringPriceCurrencyBgn ProductPriceRecurringPriceCurrency = "BGN"
	ProductPriceRecurringPriceCurrencyBhd ProductPriceRecurringPriceCurrency = "BHD"
	ProductPriceRecurringPriceCurrencyBif ProductPriceRecurringPriceCurrency = "BIF"
	ProductPriceRecurringPriceCurrencyBmd ProductPriceRecurringPriceCurrency = "BMD"
	ProductPriceRecurringPriceCurrencyBnd ProductPriceRecurringPriceCurrency = "BND"
	ProductPriceRecurringPriceCurrencyBob ProductPriceRecurringPriceCurrency = "BOB"
	ProductPriceRecurringPriceCurrencyBrl ProductPriceRecurringPriceCurrency = "BRL"
	ProductPriceRecurringPriceCurrencyBsd ProductPriceRecurringPriceCurrency = "BSD"
	ProductPriceRecurringPriceCurrencyBwp ProductPriceRecurringPriceCurrency = "BWP"
	ProductPriceRecurringPriceCurrencyByn ProductPriceRecurringPriceCurrency = "BYN"
	ProductPriceRecurringPriceCurrencyBzd ProductPriceRecurringPriceCurrency = "BZD"
	ProductPriceRecurringPriceCurrencyCad ProductPriceRecurringPriceCurrency = "CAD"
	ProductPriceRecurringPriceCurrencyChf ProductPriceRecurringPriceCurrency = "CHF"
	ProductPriceRecurringPriceCurrencyClp ProductPriceRecurringPriceCurrency = "CLP"
	ProductPriceRecurringPriceCurrencyCny ProductPriceRecurringPriceCurrency = "CNY"
	ProductPriceRecurringPriceCurrencyCop ProductPriceRecurringPriceCurrency = "COP"
	ProductPriceRecurringPriceCurrencyCrc ProductPriceRecurringPriceCurrency = "CRC"
	ProductPriceRecurringPriceCurrencyCup ProductPriceRecurringPriceCurrency = "CUP"
	ProductPriceRecurringPriceCurrencyCve ProductPriceRecurringPriceCurrency = "CVE"
	ProductPriceRecurringPriceCurrencyCzk ProductPriceRecurringPriceCurrency = "CZK"
	ProductPriceRecurringPriceCurrencyDjf ProductPriceRecurringPriceCurrency = "DJF"
	ProductPriceRecurringPriceCurrencyDkk ProductPriceRecurringPriceCurrency = "DKK"
	ProductPriceRecurringPriceCurrencyDop ProductPriceRecurringPriceCurrency = "DOP"
	ProductPriceRecurringPriceCurrencyDzd ProductPriceRecurringPriceCurrency = "DZD"
	ProductPriceRecurringPriceCurrencyEgp ProductPriceRecurringPriceCurrency = "EGP"
	ProductPriceRecurringPriceCurrencyEtb ProductPriceRecurringPriceCurrency = "ETB"
	ProductPriceRecurringPriceCurrencyEur ProductPriceRecurringPriceCurrency = "EUR"
	ProductPriceRecurringPriceCurrencyFjd ProductPriceRecurringPriceCurrency = "FJD"
	ProductPriceRecurringPriceCurrencyFkp ProductPriceRecurringPriceCurrency = "FKP"
	ProductPriceRecurringPriceCurrencyGbp ProductPriceRecurringPriceCurrency = "GBP"
	ProductPriceRecurringPriceCurrencyGel ProductPriceRecurringPriceCurrency = "GEL"
	ProductPriceRecurringPriceCurrencyGhs ProductPriceRecurringPriceCurrency = "GHS"
	ProductPriceRecurringPriceCurrencyGip ProductPriceRecurringPriceCurrency = "GIP"
	ProductPriceRecurringPriceCurrencyGmd ProductPriceRecurringPriceCurrency = "GMD"
	ProductPriceRecurringPriceCurrencyGnf ProductPriceRecurringPriceCurrency = "GNF"
	ProductPriceRecurringPriceCurrencyGtq ProductPriceRecurringPriceCurrency = "GTQ"
	ProductPriceRecurringPriceCurrencyGyd ProductPriceRecurringPriceCurrency = "GYD"
	ProductPriceRecurringPriceCurrencyHkd ProductPriceRecurringPriceCurrency = "HKD"
	ProductPriceRecurringPriceCurrencyHnl ProductPriceRecurringPriceCurrency = "HNL"
	ProductPriceRecurringPriceCurrencyHrk ProductPriceRecurringPriceCurrency = "HRK"
	ProductPriceRecurringPriceCurrencyHtg ProductPriceRecurringPriceCurrency = "HTG"
	ProductPriceRecurringPriceCurrencyHuf ProductPriceRecurringPriceCurrency = "HUF"
	ProductPriceRecurringPriceCurrencyIdr ProductPriceRecurringPriceCurrency = "IDR"
	ProductPriceRecurringPriceCurrencyIls ProductPriceRecurringPriceCurrency = "ILS"
	ProductPriceRecurringPriceCurrencyInr ProductPriceRecurringPriceCurrency = "INR"
	ProductPriceRecurringPriceCurrencyIqd ProductPriceRecurringPriceCurrency = "IQD"
	ProductPriceRecurringPriceCurrencyJmd ProductPriceRecurringPriceCurrency = "JMD"
	ProductPriceRecurringPriceCurrencyJod ProductPriceRecurringPriceCurrency = "JOD"
	ProductPriceRecurringPriceCurrencyJpy ProductPriceRecurringPriceCurrency = "JPY"
	ProductPriceRecurringPriceCurrencyKes ProductPriceRecurringPriceCurrency = "KES"
	ProductPriceRecurringPriceCurrencyKgs ProductPriceRecurringPriceCurrency = "KGS"
	ProductPriceRecurringPriceCurrencyKhr ProductPriceRecurringPriceCurrency = "KHR"
	ProductPriceRecurringPriceCurrencyKmf ProductPriceRecurringPriceCurrency = "KMF"
	ProductPriceRecurringPriceCurrencyKrw ProductPriceRecurringPriceCurrency = "KRW"
	ProductPriceRecurringPriceCurrencyKwd ProductPriceRecurringPriceCurrency = "KWD"
	ProductPriceRecurringPriceCurrencyKyd ProductPriceRecurringPriceCurrency = "KYD"
	ProductPriceRecurringPriceCurrencyKzt ProductPriceRecurringPriceCurrency = "KZT"
	ProductPriceRecurringPriceCurrencyLak ProductPriceRecurringPriceCurrency = "LAK"
	ProductPriceRecurringPriceCurrencyLbp ProductPriceRecurringPriceCurrency = "LBP"
	ProductPriceRecurringPriceCurrencyLkr ProductPriceRecurringPriceCurrency = "LKR"
	ProductPriceRecurringPriceCurrencyLrd ProductPriceRecurringPriceCurrency = "LRD"
	ProductPriceRecurringPriceCurrencyLsl ProductPriceRecurringPriceCurrency = "LSL"
	ProductPriceRecurringPriceCurrencyLyd ProductPriceRecurringPriceCurrency = "LYD"
	ProductPriceRecurringPriceCurrencyMad ProductPriceRecurringPriceCurrency = "MAD"
	ProductPriceRecurringPriceCurrencyMdl ProductPriceRecurringPriceCurrency = "MDL"
	ProductPriceRecurringPriceCurrencyMga ProductPriceRecurringPriceCurrency = "MGA"
	ProductPriceRecurringPriceCurrencyMkd ProductPriceRecurringPriceCurrency = "MKD"
	ProductPriceRecurringPriceCurrencyMmk ProductPriceRecurringPriceCurrency = "MMK"
	ProductPriceRecurringPriceCurrencyMnt ProductPriceRecurringPriceCurrency = "MNT"
	ProductPriceRecurringPriceCurrencyMop ProductPriceRecurringPriceCurrency = "MOP"
	ProductPriceRecurringPriceCurrencyMru ProductPriceRecurringPriceCurrency = "MRU"
	ProductPriceRecurringPriceCurrencyMur ProductPriceRecurringPriceCurrency = "MUR"
	ProductPriceRecurringPriceCurrencyMvr ProductPriceRecurringPriceCurrency = "MVR"
	ProductPriceRecurringPriceCurrencyMwk ProductPriceRecurringPriceCurrency = "MWK"
	ProductPriceRecurringPriceCurrencyMxn ProductPriceRecurringPriceCurrency = "MXN"
	ProductPriceRecurringPriceCurrencyMyr ProductPriceRecurringPriceCurrency = "MYR"
	ProductPriceRecurringPriceCurrencyMzn ProductPriceRecurringPriceCurrency = "MZN"
	ProductPriceRecurringPriceCurrencyNad ProductPriceRecurringPriceCurrency = "NAD"
	ProductPriceRecurringPriceCurrencyNgn ProductPriceRecurringPriceCurrency = "NGN"
	ProductPriceRecurringPriceCurrencyNio ProductPriceRecurringPriceCurrency = "NIO"
	ProductPriceRecurringPriceCurrencyNok ProductPriceRecurringPriceCurrency = "NOK"
	ProductPriceRecurringPriceCurrencyNpr ProductPriceRecurringPriceCurrency = "NPR"
	ProductPriceRecurringPriceCurrencyNzd ProductPriceRecurringPriceCurrency = "NZD"
	ProductPriceRecurringPriceCurrencyOmr ProductPriceRecurringPriceCurrency = "OMR"
	ProductPriceRecurringPriceCurrencyPab ProductPriceRecurringPriceCurrency = "PAB"
	ProductPriceRecurringPriceCurrencyPen ProductPriceRecurringPriceCurrency = "PEN"
	ProductPriceRecurringPriceCurrencyPgk ProductPriceRecurringPriceCurrency = "PGK"
	ProductPriceRecurringPriceCurrencyPhp ProductPriceRecurringPriceCurrency = "PHP"
	ProductPriceRecurringPriceCurrencyPkr ProductPriceRecurringPriceCurrency = "PKR"
	ProductPriceRecurringPriceCurrencyPln ProductPriceRecurringPriceCurrency = "PLN"
	ProductPriceRecurringPriceCurrencyPyg ProductPriceRecurringPriceCurrency = "PYG"
	ProductPriceRecurringPriceCurrencyQar ProductPriceRecurringPriceCurrency = "QAR"
	ProductPriceRecurringPriceCurrencyRon ProductPriceRecurringPriceCurrency = "RON"
	ProductPriceRecurringPriceCurrencyRsd ProductPriceRecurringPriceCurrency = "RSD"
	ProductPriceRecurringPriceCurrencyRub ProductPriceRecurringPriceCurrency = "RUB"
	ProductPriceRecurringPriceCurrencyRwf ProductPriceRecurringPriceCurrency = "RWF"
	ProductPriceRecurringPriceCurrencySar ProductPriceRecurringPriceCurrency = "SAR"
	ProductPriceRecurringPriceCurrencySbd ProductPriceRecurringPriceCurrency = "SBD"
	ProductPriceRecurringPriceCurrencyScr ProductPriceRecurringPriceCurrency = "SCR"
	ProductPriceRecurringPriceCurrencySek ProductPriceRecurringPriceCurrency = "SEK"
	ProductPriceRecurringPriceCurrencySgd ProductPriceRecurringPriceCurrency = "SGD"
	ProductPriceRecurringPriceCurrencyShp ProductPriceRecurringPriceCurrency = "SHP"
	ProductPriceRecurringPriceCurrencySle ProductPriceRecurringPriceCurrency = "SLE"
	ProductPriceRecurringPriceCurrencySll ProductPriceRecurringPriceCurrency = "SLL"
	ProductPriceRecurringPriceCurrencySos ProductPriceRecurringPriceCurrency = "SOS"
	ProductPriceRecurringPriceCurrencySrd ProductPriceRecurringPriceCurrency = "SRD"
	ProductPriceRecurringPriceCurrencySsp ProductPriceRecurringPriceCurrency = "SSP"
	ProductPriceRecurringPriceCurrencyStn ProductPriceRecurringPriceCurrency = "STN"
	ProductPriceRecurringPriceCurrencySvc ProductPriceRecurringPriceCurrency = "SVC"
	ProductPriceRecurringPriceCurrencySzl ProductPriceRecurringPriceCurrency = "SZL"
	ProductPriceRecurringPriceCurrencyThb ProductPriceRecurringPriceCurrency = "THB"
	ProductPriceRecurringPriceCurrencyTnd ProductPriceRecurringPriceCurrency = "TND"
	ProductPriceRecurringPriceCurrencyTop ProductPriceRecurringPriceCurrency = "TOP"
	ProductPriceRecurringPriceCurrencyTry ProductPriceRecurringPriceCurrency = "TRY"
	ProductPriceRecurringPriceCurrencyTtd ProductPriceRecurringPriceCurrency = "TTD"
	ProductPriceRecurringPriceCurrencyTwd ProductPriceRecurringPriceCurrency = "TWD"
	ProductPriceRecurringPriceCurrencyTzs ProductPriceRecurringPriceCurrency = "TZS"
	ProductPriceRecurringPriceCurrencyUah ProductPriceRecurringPriceCurrency = "UAH"
	ProductPriceRecurringPriceCurrencyUgx ProductPriceRecurringPriceCurrency = "UGX"
	ProductPriceRecurringPriceCurrencyUsd ProductPriceRecurringPriceCurrency = "USD"
	ProductPriceRecurringPriceCurrencyUyu ProductPriceRecurringPriceCurrency = "UYU"
	ProductPriceRecurringPriceCurrencyUzs ProductPriceRecurringPriceCurrency = "UZS"
	ProductPriceRecurringPriceCurrencyVes ProductPriceRecurringPriceCurrency = "VES"
	ProductPriceRecurringPriceCurrencyVnd ProductPriceRecurringPriceCurrency = "VND"
	ProductPriceRecurringPriceCurrencyVuv ProductPriceRecurringPriceCurrency = "VUV"
	ProductPriceRecurringPriceCurrencyWst ProductPriceRecurringPriceCurrency = "WST"
	ProductPriceRecurringPriceCurrencyXaf ProductPriceRecurringPriceCurrency = "XAF"
	ProductPriceRecurringPriceCurrencyXcd ProductPriceRecurringPriceCurrency = "XCD"
	ProductPriceRecurringPriceCurrencyXof ProductPriceRecurringPriceCurrency = "XOF"
	ProductPriceRecurringPriceCurrencyXpf ProductPriceRecurringPriceCurrency = "XPF"
	ProductPriceRecurringPriceCurrencyYer ProductPriceRecurringPriceCurrency = "YER"
	ProductPriceRecurringPriceCurrencyZar ProductPriceRecurringPriceCurrency = "ZAR"
	ProductPriceRecurringPriceCurrencyZmw ProductPriceRecurringPriceCurrency = "ZMW"
)

func (r ProductPriceRecurringPriceCurrency) IsKnown() bool {
	switch r {
	case ProductPriceRecurringPriceCurrencyAed, ProductPriceRecurringPriceCurrencyAll, ProductPriceRecurringPriceCurrencyAmd, ProductPriceRecurringPriceCurrencyAng, ProductPriceRecurringPriceCurrencyAoa, ProductPriceRecurringPriceCurrencyArs, ProductPriceRecurringPriceCurrencyAud, ProductPriceRecurringPriceCurrencyAwg, ProductPriceRecurringPriceCurrencyAzn, ProductPriceRecurringPriceCurrencyBam, ProductPriceRecurringPriceCurrencyBbd, ProductPriceRecurringPriceCurrencyBdt, ProductPriceRecurringPriceCurrencyBgn, ProductPriceRecurringPriceCurrencyBhd, ProductPriceRecurringPriceCurrencyBif, ProductPriceRecurringPriceCurrencyBmd, ProductPriceRecurringPriceCurrencyBnd, ProductPriceRecurringPriceCurrencyBob, ProductPriceRecurringPriceCurrencyBrl, ProductPriceRecurringPriceCurrencyBsd, ProductPriceRecurringPriceCurrencyBwp, ProductPriceRecurringPriceCurrencyByn, ProductPriceRecurringPriceCurrencyBzd, ProductPriceRecurringPriceCurrencyCad, ProductPriceRecurringPriceCurrencyChf, ProductPriceRecurringPriceCurrencyClp, ProductPriceRecurringPriceCurrencyCny, ProductPriceRecurringPriceCurrencyCop, ProductPriceRecurringPriceCurrencyCrc, ProductPriceRecurringPriceCurrencyCup, ProductPriceRecurringPriceCurrencyCve, ProductPriceRecurringPriceCurrencyCzk, ProductPriceRecurringPriceCurrencyDjf, ProductPriceRecurringPriceCurrencyDkk, ProductPriceRecurringPriceCurrencyDop, ProductPriceRecurringPriceCurrencyDzd, ProductPriceRecurringPriceCurrencyEgp, ProductPriceRecurringPriceCurrencyEtb, ProductPriceRecurringPriceCurrencyEur, ProductPriceRecurringPriceCurrencyFjd, ProductPriceRecurringPriceCurrencyFkp, ProductPriceRecurringPriceCurrencyGbp, ProductPriceRecurringPriceCurrencyGel, ProductPriceRecurringPriceCurrencyGhs, ProductPriceRecurringPriceCurrencyGip, ProductPriceRecurringPriceCurrencyGmd, ProductPriceRecurringPriceCurrencyGnf, ProductPriceRecurringPriceCurrencyGtq, ProductPriceRecurringPriceCurrencyGyd, ProductPriceRecurringPriceCurrencyHkd, ProductPriceRecurringPriceCurrencyHnl, ProductPriceRecurringPriceCurrencyHrk, ProductPriceRecurringPriceCurrencyHtg, ProductPriceRecurringPriceCurrencyHuf, ProductPriceRecurringPriceCurrencyIdr, ProductPriceRecurringPriceCurrencyIls, ProductPriceRecurringPriceCurrencyInr, ProductPriceRecurringPriceCurrencyIqd, ProductPriceRecurringPriceCurrencyJmd, ProductPriceRecurringPriceCurrencyJod, ProductPriceRecurringPriceCurrencyJpy, ProductPriceRecurringPriceCurrencyKes, ProductPriceRecurringPriceCurrencyKgs, ProductPriceRecurringPriceCurrencyKhr, ProductPriceRecurringPriceCurrencyKmf, ProductPriceRecurringPriceCurrencyKrw, ProductPriceRecurringPriceCurrencyKwd, ProductPriceRecurringPriceCurrencyKyd, ProductPriceRecurringPriceCurrencyKzt, ProductPriceRecurringPriceCurrencyLak, ProductPriceRecurringPriceCurrencyLbp, ProductPriceRecurringPriceCurrencyLkr, ProductPriceRecurringPriceCurrencyLrd, ProductPriceRecurringPriceCurrencyLsl, ProductPriceRecurringPriceCurrencyLyd, ProductPriceRecurringPriceCurrencyMad, ProductPriceRecurringPriceCurrencyMdl, ProductPriceRecurringPriceCurrencyMga, ProductPriceRecurringPriceCurrencyMkd, ProductPriceRecurringPriceCurrencyMmk, ProductPriceRecurringPriceCurrencyMnt, ProductPriceRecurringPriceCurrencyMop, ProductPriceRecurringPriceCurrencyMru, ProductPriceRecurringPriceCurrencyMur, ProductPriceRecurringPriceCurrencyMvr, ProductPriceRecurringPriceCurrencyMwk, ProductPriceRecurringPriceCurrencyMxn, ProductPriceRecurringPriceCurrencyMyr, ProductPriceRecurringPriceCurrencyMzn, ProductPriceRecurringPriceCurrencyNad, ProductPriceRecurringPriceCurrencyNgn, ProductPriceRecurringPriceCurrencyNio, ProductPriceRecurringPriceCurrencyNok, ProductPriceRecurringPriceCurrencyNpr, ProductPriceRecurringPriceCurrencyNzd, ProductPriceRecurringPriceCurrencyOmr, ProductPriceRecurringPriceCurrencyPab, ProductPriceRecurringPriceCurrencyPen, ProductPriceRecurringPriceCurrencyPgk, ProductPriceRecurringPriceCurrencyPhp, ProductPriceRecurringPriceCurrencyPkr, ProductPriceRecurringPriceCurrencyPln, ProductPriceRecurringPriceCurrencyPyg, ProductPriceRecurringPriceCurrencyQar, ProductPriceRecurringPriceCurrencyRon, ProductPriceRecurringPriceCurrencyRsd, ProductPriceRecurringPriceCurrencyRub, ProductPriceRecurringPriceCurrencyRwf, ProductPriceRecurringPriceCurrencySar, ProductPriceRecurringPriceCurrencySbd, ProductPriceRecurringPriceCurrencyScr, ProductPriceRecurringPriceCurrencySek, ProductPriceRecurringPriceCurrencySgd, ProductPriceRecurringPriceCurrencyShp, ProductPriceRecurringPriceCurrencySle, ProductPriceRecurringPriceCurrencySll, ProductPriceRecurringPriceCurrencySos, ProductPriceRecurringPriceCurrencySrd, ProductPriceRecurringPriceCurrencySsp, ProductPriceRecurringPriceCurrencyStn, ProductPriceRecurringPriceCurrencySvc, ProductPriceRecurringPriceCurrencySzl, ProductPriceRecurringPriceCurrencyThb, ProductPriceRecurringPriceCurrencyTnd, ProductPriceRecurringPriceCurrencyTop, ProductPriceRecurringPriceCurrencyTry, ProductPriceRecurringPriceCurrencyTtd, ProductPriceRecurringPriceCurrencyTwd, ProductPriceRecurringPriceCurrencyTzs, ProductPriceRecurringPriceCurrencyUah, ProductPriceRecurringPriceCurrencyUgx, ProductPriceRecurringPriceCurrencyUsd, ProductPriceRecurringPriceCurrencyUyu, ProductPriceRecurringPriceCurrencyUzs, ProductPriceRecurringPriceCurrencyVes, ProductPriceRecurringPriceCurrencyVnd, ProductPriceRecurringPriceCurrencyVuv, ProductPriceRecurringPriceCurrencyWst, ProductPriceRecurringPriceCurrencyXaf, ProductPriceRecurringPriceCurrencyXcd, ProductPriceRecurringPriceCurrencyXof, ProductPriceRecurringPriceCurrencyXpf, ProductPriceRecurringPriceCurrencyYer, ProductPriceRecurringPriceCurrencyZar, ProductPriceRecurringPriceCurrencyZmw:
		return true
	}
	return false
}

type ProductPriceRecurringPricePaymentFrequencyInterval string

const (
	ProductPriceRecurringPricePaymentFrequencyIntervalDay   ProductPriceRecurringPricePaymentFrequencyInterval = "Day"
	ProductPriceRecurringPricePaymentFrequencyIntervalWeek  ProductPriceRecurringPricePaymentFrequencyInterval = "Week"
	ProductPriceRecurringPricePaymentFrequencyIntervalMonth ProductPriceRecurringPricePaymentFrequencyInterval = "Month"
	ProductPriceRecurringPricePaymentFrequencyIntervalYear  ProductPriceRecurringPricePaymentFrequencyInterval = "Year"
)

func (r ProductPriceRecurringPricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductPriceRecurringPricePaymentFrequencyIntervalDay, ProductPriceRecurringPricePaymentFrequencyIntervalWeek, ProductPriceRecurringPricePaymentFrequencyIntervalMonth, ProductPriceRecurringPricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductPriceRecurringPriceSubscriptionPeriodInterval string

const (
	ProductPriceRecurringPriceSubscriptionPeriodIntervalDay   ProductPriceRecurringPriceSubscriptionPeriodInterval = "Day"
	ProductPriceRecurringPriceSubscriptionPeriodIntervalWeek  ProductPriceRecurringPriceSubscriptionPeriodInterval = "Week"
	ProductPriceRecurringPriceSubscriptionPeriodIntervalMonth ProductPriceRecurringPriceSubscriptionPeriodInterval = "Month"
	ProductPriceRecurringPriceSubscriptionPeriodIntervalYear  ProductPriceRecurringPriceSubscriptionPeriodInterval = "Year"
)

func (r ProductPriceRecurringPriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductPriceRecurringPriceSubscriptionPeriodIntervalDay, ProductPriceRecurringPriceSubscriptionPeriodIntervalWeek, ProductPriceRecurringPriceSubscriptionPeriodIntervalMonth, ProductPriceRecurringPriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

type ProductPriceRecurringPriceType string

const (
	ProductPriceRecurringPriceTypeRecurringPrice ProductPriceRecurringPriceType = "recurring_price"
)

func (r ProductPriceRecurringPriceType) IsKnown() bool {
	switch r {
	case ProductPriceRecurringPriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductPriceCurrency string

const (
	ProductPriceCurrencyAed ProductPriceCurrency = "AED"
	ProductPriceCurrencyAll ProductPriceCurrency = "ALL"
	ProductPriceCurrencyAmd ProductPriceCurrency = "AMD"
	ProductPriceCurrencyAng ProductPriceCurrency = "ANG"
	ProductPriceCurrencyAoa ProductPriceCurrency = "AOA"
	ProductPriceCurrencyArs ProductPriceCurrency = "ARS"
	ProductPriceCurrencyAud ProductPriceCurrency = "AUD"
	ProductPriceCurrencyAwg ProductPriceCurrency = "AWG"
	ProductPriceCurrencyAzn ProductPriceCurrency = "AZN"
	ProductPriceCurrencyBam ProductPriceCurrency = "BAM"
	ProductPriceCurrencyBbd ProductPriceCurrency = "BBD"
	ProductPriceCurrencyBdt ProductPriceCurrency = "BDT"
	ProductPriceCurrencyBgn ProductPriceCurrency = "BGN"
	ProductPriceCurrencyBhd ProductPriceCurrency = "BHD"
	ProductPriceCurrencyBif ProductPriceCurrency = "BIF"
	ProductPriceCurrencyBmd ProductPriceCurrency = "BMD"
	ProductPriceCurrencyBnd ProductPriceCurrency = "BND"
	ProductPriceCurrencyBob ProductPriceCurrency = "BOB"
	ProductPriceCurrencyBrl ProductPriceCurrency = "BRL"
	ProductPriceCurrencyBsd ProductPriceCurrency = "BSD"
	ProductPriceCurrencyBwp ProductPriceCurrency = "BWP"
	ProductPriceCurrencyByn ProductPriceCurrency = "BYN"
	ProductPriceCurrencyBzd ProductPriceCurrency = "BZD"
	ProductPriceCurrencyCad ProductPriceCurrency = "CAD"
	ProductPriceCurrencyChf ProductPriceCurrency = "CHF"
	ProductPriceCurrencyClp ProductPriceCurrency = "CLP"
	ProductPriceCurrencyCny ProductPriceCurrency = "CNY"
	ProductPriceCurrencyCop ProductPriceCurrency = "COP"
	ProductPriceCurrencyCrc ProductPriceCurrency = "CRC"
	ProductPriceCurrencyCup ProductPriceCurrency = "CUP"
	ProductPriceCurrencyCve ProductPriceCurrency = "CVE"
	ProductPriceCurrencyCzk ProductPriceCurrency = "CZK"
	ProductPriceCurrencyDjf ProductPriceCurrency = "DJF"
	ProductPriceCurrencyDkk ProductPriceCurrency = "DKK"
	ProductPriceCurrencyDop ProductPriceCurrency = "DOP"
	ProductPriceCurrencyDzd ProductPriceCurrency = "DZD"
	ProductPriceCurrencyEgp ProductPriceCurrency = "EGP"
	ProductPriceCurrencyEtb ProductPriceCurrency = "ETB"
	ProductPriceCurrencyEur ProductPriceCurrency = "EUR"
	ProductPriceCurrencyFjd ProductPriceCurrency = "FJD"
	ProductPriceCurrencyFkp ProductPriceCurrency = "FKP"
	ProductPriceCurrencyGbp ProductPriceCurrency = "GBP"
	ProductPriceCurrencyGel ProductPriceCurrency = "GEL"
	ProductPriceCurrencyGhs ProductPriceCurrency = "GHS"
	ProductPriceCurrencyGip ProductPriceCurrency = "GIP"
	ProductPriceCurrencyGmd ProductPriceCurrency = "GMD"
	ProductPriceCurrencyGnf ProductPriceCurrency = "GNF"
	ProductPriceCurrencyGtq ProductPriceCurrency = "GTQ"
	ProductPriceCurrencyGyd ProductPriceCurrency = "GYD"
	ProductPriceCurrencyHkd ProductPriceCurrency = "HKD"
	ProductPriceCurrencyHnl ProductPriceCurrency = "HNL"
	ProductPriceCurrencyHrk ProductPriceCurrency = "HRK"
	ProductPriceCurrencyHtg ProductPriceCurrency = "HTG"
	ProductPriceCurrencyHuf ProductPriceCurrency = "HUF"
	ProductPriceCurrencyIdr ProductPriceCurrency = "IDR"
	ProductPriceCurrencyIls ProductPriceCurrency = "ILS"
	ProductPriceCurrencyInr ProductPriceCurrency = "INR"
	ProductPriceCurrencyIqd ProductPriceCurrency = "IQD"
	ProductPriceCurrencyJmd ProductPriceCurrency = "JMD"
	ProductPriceCurrencyJod ProductPriceCurrency = "JOD"
	ProductPriceCurrencyJpy ProductPriceCurrency = "JPY"
	ProductPriceCurrencyKes ProductPriceCurrency = "KES"
	ProductPriceCurrencyKgs ProductPriceCurrency = "KGS"
	ProductPriceCurrencyKhr ProductPriceCurrency = "KHR"
	ProductPriceCurrencyKmf ProductPriceCurrency = "KMF"
	ProductPriceCurrencyKrw ProductPriceCurrency = "KRW"
	ProductPriceCurrencyKwd ProductPriceCurrency = "KWD"
	ProductPriceCurrencyKyd ProductPriceCurrency = "KYD"
	ProductPriceCurrencyKzt ProductPriceCurrency = "KZT"
	ProductPriceCurrencyLak ProductPriceCurrency = "LAK"
	ProductPriceCurrencyLbp ProductPriceCurrency = "LBP"
	ProductPriceCurrencyLkr ProductPriceCurrency = "LKR"
	ProductPriceCurrencyLrd ProductPriceCurrency = "LRD"
	ProductPriceCurrencyLsl ProductPriceCurrency = "LSL"
	ProductPriceCurrencyLyd ProductPriceCurrency = "LYD"
	ProductPriceCurrencyMad ProductPriceCurrency = "MAD"
	ProductPriceCurrencyMdl ProductPriceCurrency = "MDL"
	ProductPriceCurrencyMga ProductPriceCurrency = "MGA"
	ProductPriceCurrencyMkd ProductPriceCurrency = "MKD"
	ProductPriceCurrencyMmk ProductPriceCurrency = "MMK"
	ProductPriceCurrencyMnt ProductPriceCurrency = "MNT"
	ProductPriceCurrencyMop ProductPriceCurrency = "MOP"
	ProductPriceCurrencyMru ProductPriceCurrency = "MRU"
	ProductPriceCurrencyMur ProductPriceCurrency = "MUR"
	ProductPriceCurrencyMvr ProductPriceCurrency = "MVR"
	ProductPriceCurrencyMwk ProductPriceCurrency = "MWK"
	ProductPriceCurrencyMxn ProductPriceCurrency = "MXN"
	ProductPriceCurrencyMyr ProductPriceCurrency = "MYR"
	ProductPriceCurrencyMzn ProductPriceCurrency = "MZN"
	ProductPriceCurrencyNad ProductPriceCurrency = "NAD"
	ProductPriceCurrencyNgn ProductPriceCurrency = "NGN"
	ProductPriceCurrencyNio ProductPriceCurrency = "NIO"
	ProductPriceCurrencyNok ProductPriceCurrency = "NOK"
	ProductPriceCurrencyNpr ProductPriceCurrency = "NPR"
	ProductPriceCurrencyNzd ProductPriceCurrency = "NZD"
	ProductPriceCurrencyOmr ProductPriceCurrency = "OMR"
	ProductPriceCurrencyPab ProductPriceCurrency = "PAB"
	ProductPriceCurrencyPen ProductPriceCurrency = "PEN"
	ProductPriceCurrencyPgk ProductPriceCurrency = "PGK"
	ProductPriceCurrencyPhp ProductPriceCurrency = "PHP"
	ProductPriceCurrencyPkr ProductPriceCurrency = "PKR"
	ProductPriceCurrencyPln ProductPriceCurrency = "PLN"
	ProductPriceCurrencyPyg ProductPriceCurrency = "PYG"
	ProductPriceCurrencyQar ProductPriceCurrency = "QAR"
	ProductPriceCurrencyRon ProductPriceCurrency = "RON"
	ProductPriceCurrencyRsd ProductPriceCurrency = "RSD"
	ProductPriceCurrencyRub ProductPriceCurrency = "RUB"
	ProductPriceCurrencyRwf ProductPriceCurrency = "RWF"
	ProductPriceCurrencySar ProductPriceCurrency = "SAR"
	ProductPriceCurrencySbd ProductPriceCurrency = "SBD"
	ProductPriceCurrencyScr ProductPriceCurrency = "SCR"
	ProductPriceCurrencySek ProductPriceCurrency = "SEK"
	ProductPriceCurrencySgd ProductPriceCurrency = "SGD"
	ProductPriceCurrencyShp ProductPriceCurrency = "SHP"
	ProductPriceCurrencySle ProductPriceCurrency = "SLE"
	ProductPriceCurrencySll ProductPriceCurrency = "SLL"
	ProductPriceCurrencySos ProductPriceCurrency = "SOS"
	ProductPriceCurrencySrd ProductPriceCurrency = "SRD"
	ProductPriceCurrencySsp ProductPriceCurrency = "SSP"
	ProductPriceCurrencyStn ProductPriceCurrency = "STN"
	ProductPriceCurrencySvc ProductPriceCurrency = "SVC"
	ProductPriceCurrencySzl ProductPriceCurrency = "SZL"
	ProductPriceCurrencyThb ProductPriceCurrency = "THB"
	ProductPriceCurrencyTnd ProductPriceCurrency = "TND"
	ProductPriceCurrencyTop ProductPriceCurrency = "TOP"
	ProductPriceCurrencyTry ProductPriceCurrency = "TRY"
	ProductPriceCurrencyTtd ProductPriceCurrency = "TTD"
	ProductPriceCurrencyTwd ProductPriceCurrency = "TWD"
	ProductPriceCurrencyTzs ProductPriceCurrency = "TZS"
	ProductPriceCurrencyUah ProductPriceCurrency = "UAH"
	ProductPriceCurrencyUgx ProductPriceCurrency = "UGX"
	ProductPriceCurrencyUsd ProductPriceCurrency = "USD"
	ProductPriceCurrencyUyu ProductPriceCurrency = "UYU"
	ProductPriceCurrencyUzs ProductPriceCurrency = "UZS"
	ProductPriceCurrencyVes ProductPriceCurrency = "VES"
	ProductPriceCurrencyVnd ProductPriceCurrency = "VND"
	ProductPriceCurrencyVuv ProductPriceCurrency = "VUV"
	ProductPriceCurrencyWst ProductPriceCurrency = "WST"
	ProductPriceCurrencyXaf ProductPriceCurrency = "XAF"
	ProductPriceCurrencyXcd ProductPriceCurrency = "XCD"
	ProductPriceCurrencyXof ProductPriceCurrency = "XOF"
	ProductPriceCurrencyXpf ProductPriceCurrency = "XPF"
	ProductPriceCurrencyYer ProductPriceCurrency = "YER"
	ProductPriceCurrencyZar ProductPriceCurrency = "ZAR"
	ProductPriceCurrencyZmw ProductPriceCurrency = "ZMW"
)

func (r ProductPriceCurrency) IsKnown() bool {
	switch r {
	case ProductPriceCurrencyAed, ProductPriceCurrencyAll, ProductPriceCurrencyAmd, ProductPriceCurrencyAng, ProductPriceCurrencyAoa, ProductPriceCurrencyArs, ProductPriceCurrencyAud, ProductPriceCurrencyAwg, ProductPriceCurrencyAzn, ProductPriceCurrencyBam, ProductPriceCurrencyBbd, ProductPriceCurrencyBdt, ProductPriceCurrencyBgn, ProductPriceCurrencyBhd, ProductPriceCurrencyBif, ProductPriceCurrencyBmd, ProductPriceCurrencyBnd, ProductPriceCurrencyBob, ProductPriceCurrencyBrl, ProductPriceCurrencyBsd, ProductPriceCurrencyBwp, ProductPriceCurrencyByn, ProductPriceCurrencyBzd, ProductPriceCurrencyCad, ProductPriceCurrencyChf, ProductPriceCurrencyClp, ProductPriceCurrencyCny, ProductPriceCurrencyCop, ProductPriceCurrencyCrc, ProductPriceCurrencyCup, ProductPriceCurrencyCve, ProductPriceCurrencyCzk, ProductPriceCurrencyDjf, ProductPriceCurrencyDkk, ProductPriceCurrencyDop, ProductPriceCurrencyDzd, ProductPriceCurrencyEgp, ProductPriceCurrencyEtb, ProductPriceCurrencyEur, ProductPriceCurrencyFjd, ProductPriceCurrencyFkp, ProductPriceCurrencyGbp, ProductPriceCurrencyGel, ProductPriceCurrencyGhs, ProductPriceCurrencyGip, ProductPriceCurrencyGmd, ProductPriceCurrencyGnf, ProductPriceCurrencyGtq, ProductPriceCurrencyGyd, ProductPriceCurrencyHkd, ProductPriceCurrencyHnl, ProductPriceCurrencyHrk, ProductPriceCurrencyHtg, ProductPriceCurrencyHuf, ProductPriceCurrencyIdr, ProductPriceCurrencyIls, ProductPriceCurrencyInr, ProductPriceCurrencyIqd, ProductPriceCurrencyJmd, ProductPriceCurrencyJod, ProductPriceCurrencyJpy, ProductPriceCurrencyKes, ProductPriceCurrencyKgs, ProductPriceCurrencyKhr, ProductPriceCurrencyKmf, ProductPriceCurrencyKrw, ProductPriceCurrencyKwd, ProductPriceCurrencyKyd, ProductPriceCurrencyKzt, ProductPriceCurrencyLak, ProductPriceCurrencyLbp, ProductPriceCurrencyLkr, ProductPriceCurrencyLrd, ProductPriceCurrencyLsl, ProductPriceCurrencyLyd, ProductPriceCurrencyMad, ProductPriceCurrencyMdl, ProductPriceCurrencyMga, ProductPriceCurrencyMkd, ProductPriceCurrencyMmk, ProductPriceCurrencyMnt, ProductPriceCurrencyMop, ProductPriceCurrencyMru, ProductPriceCurrencyMur, ProductPriceCurrencyMvr, ProductPriceCurrencyMwk, ProductPriceCurrencyMxn, ProductPriceCurrencyMyr, ProductPriceCurrencyMzn, ProductPriceCurrencyNad, ProductPriceCurrencyNgn, ProductPriceCurrencyNio, ProductPriceCurrencyNok, ProductPriceCurrencyNpr, ProductPriceCurrencyNzd, ProductPriceCurrencyOmr, ProductPriceCurrencyPab, ProductPriceCurrencyPen, ProductPriceCurrencyPgk, ProductPriceCurrencyPhp, ProductPriceCurrencyPkr, ProductPriceCurrencyPln, ProductPriceCurrencyPyg, ProductPriceCurrencyQar, ProductPriceCurrencyRon, ProductPriceCurrencyRsd, ProductPriceCurrencyRub, ProductPriceCurrencyRwf, ProductPriceCurrencySar, ProductPriceCurrencySbd, ProductPriceCurrencyScr, ProductPriceCurrencySek, ProductPriceCurrencySgd, ProductPriceCurrencyShp, ProductPriceCurrencySle, ProductPriceCurrencySll, ProductPriceCurrencySos, ProductPriceCurrencySrd, ProductPriceCurrencySsp, ProductPriceCurrencyStn, ProductPriceCurrencySvc, ProductPriceCurrencySzl, ProductPriceCurrencyThb, ProductPriceCurrencyTnd, ProductPriceCurrencyTop, ProductPriceCurrencyTry, ProductPriceCurrencyTtd, ProductPriceCurrencyTwd, ProductPriceCurrencyTzs, ProductPriceCurrencyUah, ProductPriceCurrencyUgx, ProductPriceCurrencyUsd, ProductPriceCurrencyUyu, ProductPriceCurrencyUzs, ProductPriceCurrencyVes, ProductPriceCurrencyVnd, ProductPriceCurrencyVuv, ProductPriceCurrencyWst, ProductPriceCurrencyXaf, ProductPriceCurrencyXcd, ProductPriceCurrencyXof, ProductPriceCurrencyXpf, ProductPriceCurrencyYer, ProductPriceCurrencyZar, ProductPriceCurrencyZmw:
		return true
	}
	return false
}

type ProductPriceType string

const (
	ProductPriceTypeOneTimePrice   ProductPriceType = "one_time_price"
	ProductPriceTypeRecurringPrice ProductPriceType = "recurring_price"
)

func (r ProductPriceType) IsKnown() bool {
	switch r {
	case ProductPriceTypeOneTimePrice, ProductPriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductPricePaymentFrequencyInterval string

const (
	ProductPricePaymentFrequencyIntervalDay   ProductPricePaymentFrequencyInterval = "Day"
	ProductPricePaymentFrequencyIntervalWeek  ProductPricePaymentFrequencyInterval = "Week"
	ProductPricePaymentFrequencyIntervalMonth ProductPricePaymentFrequencyInterval = "Month"
	ProductPricePaymentFrequencyIntervalYear  ProductPricePaymentFrequencyInterval = "Year"
)

func (r ProductPricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductPricePaymentFrequencyIntervalDay, ProductPricePaymentFrequencyIntervalWeek, ProductPricePaymentFrequencyIntervalMonth, ProductPricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductPriceSubscriptionPeriodInterval string

const (
	ProductPriceSubscriptionPeriodIntervalDay   ProductPriceSubscriptionPeriodInterval = "Day"
	ProductPriceSubscriptionPeriodIntervalWeek  ProductPriceSubscriptionPeriodInterval = "Week"
	ProductPriceSubscriptionPeriodIntervalMonth ProductPriceSubscriptionPeriodInterval = "Month"
	ProductPriceSubscriptionPeriodIntervalYear  ProductPriceSubscriptionPeriodInterval = "Year"
)

func (r ProductPriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductPriceSubscriptionPeriodIntervalDay, ProductPriceSubscriptionPeriodIntervalWeek, ProductPriceSubscriptionPeriodIntervalMonth, ProductPriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

// Represents the different categories of taxation applicable to various products
// and services.
type ProductTaxCategory string

const (
	ProductTaxCategoryDigitalProducts ProductTaxCategory = "digital_products"
	ProductTaxCategorySaas            ProductTaxCategory = "saas"
	ProductTaxCategoryEBook           ProductTaxCategory = "e_book"
)

func (r ProductTaxCategory) IsKnown() bool {
	switch r {
	case ProductTaxCategoryDigitalProducts, ProductTaxCategorySaas, ProductTaxCategoryEBook:
		return true
	}
	return false
}

type ProductLicenseKeyDuration struct {
	Count    int64                             `json:"count,required"`
	Interval ProductLicenseKeyDurationInterval `json:"interval,required"`
	JSON     productLicenseKeyDurationJSON     `json:"-"`
}

// productLicenseKeyDurationJSON contains the JSON metadata for the struct
// [ProductLicenseKeyDuration]
type productLicenseKeyDurationJSON struct {
	Count       apijson.Field
	Interval    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductLicenseKeyDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productLicenseKeyDurationJSON) RawJSON() string {
	return r.raw
}

type ProductLicenseKeyDurationInterval string

const (
	ProductLicenseKeyDurationIntervalDay   ProductLicenseKeyDurationInterval = "Day"
	ProductLicenseKeyDurationIntervalWeek  ProductLicenseKeyDurationInterval = "Week"
	ProductLicenseKeyDurationIntervalMonth ProductLicenseKeyDurationInterval = "Month"
	ProductLicenseKeyDurationIntervalYear  ProductLicenseKeyDurationInterval = "Year"
)

func (r ProductLicenseKeyDurationInterval) IsKnown() bool {
	switch r {
	case ProductLicenseKeyDurationIntervalDay, ProductLicenseKeyDurationIntervalWeek, ProductLicenseKeyDurationIntervalMonth, ProductLicenseKeyDurationIntervalYear:
		return true
	}
	return false
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
	TaxCategory ProductListResponseTaxCategory `json:"tax_category,required"`
	// Timestamp when the product was last updated.
	UpdatedAt time.Time                   `json:"updated_at,required" format:"date-time"`
	Currency  ProductListResponseCurrency `json:"currency,nullable"`
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
	Price        int64                   `json:"price,nullable"`
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

// Represents the different categories of taxation applicable to various products
// and services.
type ProductListResponseTaxCategory string

const (
	ProductListResponseTaxCategoryDigitalProducts ProductListResponseTaxCategory = "digital_products"
	ProductListResponseTaxCategorySaas            ProductListResponseTaxCategory = "saas"
	ProductListResponseTaxCategoryEBook           ProductListResponseTaxCategory = "e_book"
)

func (r ProductListResponseTaxCategory) IsKnown() bool {
	switch r {
	case ProductListResponseTaxCategoryDigitalProducts, ProductListResponseTaxCategorySaas, ProductListResponseTaxCategoryEBook:
		return true
	}
	return false
}

type ProductListResponseCurrency string

const (
	ProductListResponseCurrencyAed ProductListResponseCurrency = "AED"
	ProductListResponseCurrencyAll ProductListResponseCurrency = "ALL"
	ProductListResponseCurrencyAmd ProductListResponseCurrency = "AMD"
	ProductListResponseCurrencyAng ProductListResponseCurrency = "ANG"
	ProductListResponseCurrencyAoa ProductListResponseCurrency = "AOA"
	ProductListResponseCurrencyArs ProductListResponseCurrency = "ARS"
	ProductListResponseCurrencyAud ProductListResponseCurrency = "AUD"
	ProductListResponseCurrencyAwg ProductListResponseCurrency = "AWG"
	ProductListResponseCurrencyAzn ProductListResponseCurrency = "AZN"
	ProductListResponseCurrencyBam ProductListResponseCurrency = "BAM"
	ProductListResponseCurrencyBbd ProductListResponseCurrency = "BBD"
	ProductListResponseCurrencyBdt ProductListResponseCurrency = "BDT"
	ProductListResponseCurrencyBgn ProductListResponseCurrency = "BGN"
	ProductListResponseCurrencyBhd ProductListResponseCurrency = "BHD"
	ProductListResponseCurrencyBif ProductListResponseCurrency = "BIF"
	ProductListResponseCurrencyBmd ProductListResponseCurrency = "BMD"
	ProductListResponseCurrencyBnd ProductListResponseCurrency = "BND"
	ProductListResponseCurrencyBob ProductListResponseCurrency = "BOB"
	ProductListResponseCurrencyBrl ProductListResponseCurrency = "BRL"
	ProductListResponseCurrencyBsd ProductListResponseCurrency = "BSD"
	ProductListResponseCurrencyBwp ProductListResponseCurrency = "BWP"
	ProductListResponseCurrencyByn ProductListResponseCurrency = "BYN"
	ProductListResponseCurrencyBzd ProductListResponseCurrency = "BZD"
	ProductListResponseCurrencyCad ProductListResponseCurrency = "CAD"
	ProductListResponseCurrencyChf ProductListResponseCurrency = "CHF"
	ProductListResponseCurrencyClp ProductListResponseCurrency = "CLP"
	ProductListResponseCurrencyCny ProductListResponseCurrency = "CNY"
	ProductListResponseCurrencyCop ProductListResponseCurrency = "COP"
	ProductListResponseCurrencyCrc ProductListResponseCurrency = "CRC"
	ProductListResponseCurrencyCup ProductListResponseCurrency = "CUP"
	ProductListResponseCurrencyCve ProductListResponseCurrency = "CVE"
	ProductListResponseCurrencyCzk ProductListResponseCurrency = "CZK"
	ProductListResponseCurrencyDjf ProductListResponseCurrency = "DJF"
	ProductListResponseCurrencyDkk ProductListResponseCurrency = "DKK"
	ProductListResponseCurrencyDop ProductListResponseCurrency = "DOP"
	ProductListResponseCurrencyDzd ProductListResponseCurrency = "DZD"
	ProductListResponseCurrencyEgp ProductListResponseCurrency = "EGP"
	ProductListResponseCurrencyEtb ProductListResponseCurrency = "ETB"
	ProductListResponseCurrencyEur ProductListResponseCurrency = "EUR"
	ProductListResponseCurrencyFjd ProductListResponseCurrency = "FJD"
	ProductListResponseCurrencyFkp ProductListResponseCurrency = "FKP"
	ProductListResponseCurrencyGbp ProductListResponseCurrency = "GBP"
	ProductListResponseCurrencyGel ProductListResponseCurrency = "GEL"
	ProductListResponseCurrencyGhs ProductListResponseCurrency = "GHS"
	ProductListResponseCurrencyGip ProductListResponseCurrency = "GIP"
	ProductListResponseCurrencyGmd ProductListResponseCurrency = "GMD"
	ProductListResponseCurrencyGnf ProductListResponseCurrency = "GNF"
	ProductListResponseCurrencyGtq ProductListResponseCurrency = "GTQ"
	ProductListResponseCurrencyGyd ProductListResponseCurrency = "GYD"
	ProductListResponseCurrencyHkd ProductListResponseCurrency = "HKD"
	ProductListResponseCurrencyHnl ProductListResponseCurrency = "HNL"
	ProductListResponseCurrencyHrk ProductListResponseCurrency = "HRK"
	ProductListResponseCurrencyHtg ProductListResponseCurrency = "HTG"
	ProductListResponseCurrencyHuf ProductListResponseCurrency = "HUF"
	ProductListResponseCurrencyIdr ProductListResponseCurrency = "IDR"
	ProductListResponseCurrencyIls ProductListResponseCurrency = "ILS"
	ProductListResponseCurrencyInr ProductListResponseCurrency = "INR"
	ProductListResponseCurrencyIqd ProductListResponseCurrency = "IQD"
	ProductListResponseCurrencyJmd ProductListResponseCurrency = "JMD"
	ProductListResponseCurrencyJod ProductListResponseCurrency = "JOD"
	ProductListResponseCurrencyJpy ProductListResponseCurrency = "JPY"
	ProductListResponseCurrencyKes ProductListResponseCurrency = "KES"
	ProductListResponseCurrencyKgs ProductListResponseCurrency = "KGS"
	ProductListResponseCurrencyKhr ProductListResponseCurrency = "KHR"
	ProductListResponseCurrencyKmf ProductListResponseCurrency = "KMF"
	ProductListResponseCurrencyKrw ProductListResponseCurrency = "KRW"
	ProductListResponseCurrencyKwd ProductListResponseCurrency = "KWD"
	ProductListResponseCurrencyKyd ProductListResponseCurrency = "KYD"
	ProductListResponseCurrencyKzt ProductListResponseCurrency = "KZT"
	ProductListResponseCurrencyLak ProductListResponseCurrency = "LAK"
	ProductListResponseCurrencyLbp ProductListResponseCurrency = "LBP"
	ProductListResponseCurrencyLkr ProductListResponseCurrency = "LKR"
	ProductListResponseCurrencyLrd ProductListResponseCurrency = "LRD"
	ProductListResponseCurrencyLsl ProductListResponseCurrency = "LSL"
	ProductListResponseCurrencyLyd ProductListResponseCurrency = "LYD"
	ProductListResponseCurrencyMad ProductListResponseCurrency = "MAD"
	ProductListResponseCurrencyMdl ProductListResponseCurrency = "MDL"
	ProductListResponseCurrencyMga ProductListResponseCurrency = "MGA"
	ProductListResponseCurrencyMkd ProductListResponseCurrency = "MKD"
	ProductListResponseCurrencyMmk ProductListResponseCurrency = "MMK"
	ProductListResponseCurrencyMnt ProductListResponseCurrency = "MNT"
	ProductListResponseCurrencyMop ProductListResponseCurrency = "MOP"
	ProductListResponseCurrencyMru ProductListResponseCurrency = "MRU"
	ProductListResponseCurrencyMur ProductListResponseCurrency = "MUR"
	ProductListResponseCurrencyMvr ProductListResponseCurrency = "MVR"
	ProductListResponseCurrencyMwk ProductListResponseCurrency = "MWK"
	ProductListResponseCurrencyMxn ProductListResponseCurrency = "MXN"
	ProductListResponseCurrencyMyr ProductListResponseCurrency = "MYR"
	ProductListResponseCurrencyMzn ProductListResponseCurrency = "MZN"
	ProductListResponseCurrencyNad ProductListResponseCurrency = "NAD"
	ProductListResponseCurrencyNgn ProductListResponseCurrency = "NGN"
	ProductListResponseCurrencyNio ProductListResponseCurrency = "NIO"
	ProductListResponseCurrencyNok ProductListResponseCurrency = "NOK"
	ProductListResponseCurrencyNpr ProductListResponseCurrency = "NPR"
	ProductListResponseCurrencyNzd ProductListResponseCurrency = "NZD"
	ProductListResponseCurrencyOmr ProductListResponseCurrency = "OMR"
	ProductListResponseCurrencyPab ProductListResponseCurrency = "PAB"
	ProductListResponseCurrencyPen ProductListResponseCurrency = "PEN"
	ProductListResponseCurrencyPgk ProductListResponseCurrency = "PGK"
	ProductListResponseCurrencyPhp ProductListResponseCurrency = "PHP"
	ProductListResponseCurrencyPkr ProductListResponseCurrency = "PKR"
	ProductListResponseCurrencyPln ProductListResponseCurrency = "PLN"
	ProductListResponseCurrencyPyg ProductListResponseCurrency = "PYG"
	ProductListResponseCurrencyQar ProductListResponseCurrency = "QAR"
	ProductListResponseCurrencyRon ProductListResponseCurrency = "RON"
	ProductListResponseCurrencyRsd ProductListResponseCurrency = "RSD"
	ProductListResponseCurrencyRub ProductListResponseCurrency = "RUB"
	ProductListResponseCurrencyRwf ProductListResponseCurrency = "RWF"
	ProductListResponseCurrencySar ProductListResponseCurrency = "SAR"
	ProductListResponseCurrencySbd ProductListResponseCurrency = "SBD"
	ProductListResponseCurrencyScr ProductListResponseCurrency = "SCR"
	ProductListResponseCurrencySek ProductListResponseCurrency = "SEK"
	ProductListResponseCurrencySgd ProductListResponseCurrency = "SGD"
	ProductListResponseCurrencyShp ProductListResponseCurrency = "SHP"
	ProductListResponseCurrencySle ProductListResponseCurrency = "SLE"
	ProductListResponseCurrencySll ProductListResponseCurrency = "SLL"
	ProductListResponseCurrencySos ProductListResponseCurrency = "SOS"
	ProductListResponseCurrencySrd ProductListResponseCurrency = "SRD"
	ProductListResponseCurrencySsp ProductListResponseCurrency = "SSP"
	ProductListResponseCurrencyStn ProductListResponseCurrency = "STN"
	ProductListResponseCurrencySvc ProductListResponseCurrency = "SVC"
	ProductListResponseCurrencySzl ProductListResponseCurrency = "SZL"
	ProductListResponseCurrencyThb ProductListResponseCurrency = "THB"
	ProductListResponseCurrencyTnd ProductListResponseCurrency = "TND"
	ProductListResponseCurrencyTop ProductListResponseCurrency = "TOP"
	ProductListResponseCurrencyTry ProductListResponseCurrency = "TRY"
	ProductListResponseCurrencyTtd ProductListResponseCurrency = "TTD"
	ProductListResponseCurrencyTwd ProductListResponseCurrency = "TWD"
	ProductListResponseCurrencyTzs ProductListResponseCurrency = "TZS"
	ProductListResponseCurrencyUah ProductListResponseCurrency = "UAH"
	ProductListResponseCurrencyUgx ProductListResponseCurrency = "UGX"
	ProductListResponseCurrencyUsd ProductListResponseCurrency = "USD"
	ProductListResponseCurrencyUyu ProductListResponseCurrency = "UYU"
	ProductListResponseCurrencyUzs ProductListResponseCurrency = "UZS"
	ProductListResponseCurrencyVes ProductListResponseCurrency = "VES"
	ProductListResponseCurrencyVnd ProductListResponseCurrency = "VND"
	ProductListResponseCurrencyVuv ProductListResponseCurrency = "VUV"
	ProductListResponseCurrencyWst ProductListResponseCurrency = "WST"
	ProductListResponseCurrencyXaf ProductListResponseCurrency = "XAF"
	ProductListResponseCurrencyXcd ProductListResponseCurrency = "XCD"
	ProductListResponseCurrencyXof ProductListResponseCurrency = "XOF"
	ProductListResponseCurrencyXpf ProductListResponseCurrency = "XPF"
	ProductListResponseCurrencyYer ProductListResponseCurrency = "YER"
	ProductListResponseCurrencyZar ProductListResponseCurrency = "ZAR"
	ProductListResponseCurrencyZmw ProductListResponseCurrency = "ZMW"
)

func (r ProductListResponseCurrency) IsKnown() bool {
	switch r {
	case ProductListResponseCurrencyAed, ProductListResponseCurrencyAll, ProductListResponseCurrencyAmd, ProductListResponseCurrencyAng, ProductListResponseCurrencyAoa, ProductListResponseCurrencyArs, ProductListResponseCurrencyAud, ProductListResponseCurrencyAwg, ProductListResponseCurrencyAzn, ProductListResponseCurrencyBam, ProductListResponseCurrencyBbd, ProductListResponseCurrencyBdt, ProductListResponseCurrencyBgn, ProductListResponseCurrencyBhd, ProductListResponseCurrencyBif, ProductListResponseCurrencyBmd, ProductListResponseCurrencyBnd, ProductListResponseCurrencyBob, ProductListResponseCurrencyBrl, ProductListResponseCurrencyBsd, ProductListResponseCurrencyBwp, ProductListResponseCurrencyByn, ProductListResponseCurrencyBzd, ProductListResponseCurrencyCad, ProductListResponseCurrencyChf, ProductListResponseCurrencyClp, ProductListResponseCurrencyCny, ProductListResponseCurrencyCop, ProductListResponseCurrencyCrc, ProductListResponseCurrencyCup, ProductListResponseCurrencyCve, ProductListResponseCurrencyCzk, ProductListResponseCurrencyDjf, ProductListResponseCurrencyDkk, ProductListResponseCurrencyDop, ProductListResponseCurrencyDzd, ProductListResponseCurrencyEgp, ProductListResponseCurrencyEtb, ProductListResponseCurrencyEur, ProductListResponseCurrencyFjd, ProductListResponseCurrencyFkp, ProductListResponseCurrencyGbp, ProductListResponseCurrencyGel, ProductListResponseCurrencyGhs, ProductListResponseCurrencyGip, ProductListResponseCurrencyGmd, ProductListResponseCurrencyGnf, ProductListResponseCurrencyGtq, ProductListResponseCurrencyGyd, ProductListResponseCurrencyHkd, ProductListResponseCurrencyHnl, ProductListResponseCurrencyHrk, ProductListResponseCurrencyHtg, ProductListResponseCurrencyHuf, ProductListResponseCurrencyIdr, ProductListResponseCurrencyIls, ProductListResponseCurrencyInr, ProductListResponseCurrencyIqd, ProductListResponseCurrencyJmd, ProductListResponseCurrencyJod, ProductListResponseCurrencyJpy, ProductListResponseCurrencyKes, ProductListResponseCurrencyKgs, ProductListResponseCurrencyKhr, ProductListResponseCurrencyKmf, ProductListResponseCurrencyKrw, ProductListResponseCurrencyKwd, ProductListResponseCurrencyKyd, ProductListResponseCurrencyKzt, ProductListResponseCurrencyLak, ProductListResponseCurrencyLbp, ProductListResponseCurrencyLkr, ProductListResponseCurrencyLrd, ProductListResponseCurrencyLsl, ProductListResponseCurrencyLyd, ProductListResponseCurrencyMad, ProductListResponseCurrencyMdl, ProductListResponseCurrencyMga, ProductListResponseCurrencyMkd, ProductListResponseCurrencyMmk, ProductListResponseCurrencyMnt, ProductListResponseCurrencyMop, ProductListResponseCurrencyMru, ProductListResponseCurrencyMur, ProductListResponseCurrencyMvr, ProductListResponseCurrencyMwk, ProductListResponseCurrencyMxn, ProductListResponseCurrencyMyr, ProductListResponseCurrencyMzn, ProductListResponseCurrencyNad, ProductListResponseCurrencyNgn, ProductListResponseCurrencyNio, ProductListResponseCurrencyNok, ProductListResponseCurrencyNpr, ProductListResponseCurrencyNzd, ProductListResponseCurrencyOmr, ProductListResponseCurrencyPab, ProductListResponseCurrencyPen, ProductListResponseCurrencyPgk, ProductListResponseCurrencyPhp, ProductListResponseCurrencyPkr, ProductListResponseCurrencyPln, ProductListResponseCurrencyPyg, ProductListResponseCurrencyQar, ProductListResponseCurrencyRon, ProductListResponseCurrencyRsd, ProductListResponseCurrencyRub, ProductListResponseCurrencyRwf, ProductListResponseCurrencySar, ProductListResponseCurrencySbd, ProductListResponseCurrencyScr, ProductListResponseCurrencySek, ProductListResponseCurrencySgd, ProductListResponseCurrencyShp, ProductListResponseCurrencySle, ProductListResponseCurrencySll, ProductListResponseCurrencySos, ProductListResponseCurrencySrd, ProductListResponseCurrencySsp, ProductListResponseCurrencyStn, ProductListResponseCurrencySvc, ProductListResponseCurrencySzl, ProductListResponseCurrencyThb, ProductListResponseCurrencyTnd, ProductListResponseCurrencyTop, ProductListResponseCurrencyTry, ProductListResponseCurrencyTtd, ProductListResponseCurrencyTwd, ProductListResponseCurrencyTzs, ProductListResponseCurrencyUah, ProductListResponseCurrencyUgx, ProductListResponseCurrencyUsd, ProductListResponseCurrencyUyu, ProductListResponseCurrencyUzs, ProductListResponseCurrencyVes, ProductListResponseCurrencyVnd, ProductListResponseCurrencyVuv, ProductListResponseCurrencyWst, ProductListResponseCurrencyXaf, ProductListResponseCurrencyXcd, ProductListResponseCurrencyXof, ProductListResponseCurrencyXpf, ProductListResponseCurrencyYer, ProductListResponseCurrencyZar, ProductListResponseCurrencyZmw:
		return true
	}
	return false
}

type ProductNewParams struct {
	Price param.Field[ProductNewParamsPriceUnion] `json:"price,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[ProductNewParamsTaxCategory] `json:"tax_category,required"`
	// Optional description of the product
	Description param.Field[string] `json:"description"`
	// Optional message displayed during license key activation
	LicenseKeyActivationMessage param.Field[string] `json:"license_key_activation_message"`
	// The number of times the license key can be activated. Must be 0 or greater
	LicenseKeyActivationsLimit param.Field[int64]                              `json:"license_key_activations_limit"`
	LicenseKeyDuration         param.Field[ProductNewParamsLicenseKeyDuration] `json:"license_key_duration"`
	// When true, generates and sends a license key to your customer. Defaults to false
	LicenseKeyEnabled param.Field[bool] `json:"license_key_enabled"`
	// Optional name of the product
	Name param.Field[string] `json:"name"`
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductNewParamsPrice struct {
	Currency param.Field[ProductNewParamsPriceCurrency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool]                      `json:"purchasing_power_parity,required"`
	Type                  param.Field[ProductNewParamsPriceType] `json:"type,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    param.Field[int64]                                         `json:"payment_frequency_count"`
	PaymentFrequencyInterval param.Field[ProductNewParamsPricePaymentFrequencyInterval] `json:"payment_frequency_interval"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    param.Field[int64]                                           `json:"subscription_period_count"`
	SubscriptionPeriodInterval param.Field[ProductNewParamsPriceSubscriptionPeriodInterval] `json:"subscription_period_interval"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r ProductNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductNewParamsPrice) implementsProductNewParamsPriceUnion() {}

// Satisfied by [ProductNewParamsPriceOneTimePrice],
// [ProductNewParamsPriceRecurringPrice], [ProductNewParamsPrice].
type ProductNewParamsPriceUnion interface {
	implementsProductNewParamsPriceUnion()
}

type ProductNewParamsPriceOneTimePrice struct {
	Currency param.Field[ProductNewParamsPriceOneTimePriceCurrency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool]                                  `json:"purchasing_power_parity,required"`
	Type                  param.Field[ProductNewParamsPriceOneTimePriceType] `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
}

func (r ProductNewParamsPriceOneTimePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductNewParamsPriceOneTimePrice) implementsProductNewParamsPriceUnion() {}

type ProductNewParamsPriceOneTimePriceCurrency string

const (
	ProductNewParamsPriceOneTimePriceCurrencyAed ProductNewParamsPriceOneTimePriceCurrency = "AED"
	ProductNewParamsPriceOneTimePriceCurrencyAll ProductNewParamsPriceOneTimePriceCurrency = "ALL"
	ProductNewParamsPriceOneTimePriceCurrencyAmd ProductNewParamsPriceOneTimePriceCurrency = "AMD"
	ProductNewParamsPriceOneTimePriceCurrencyAng ProductNewParamsPriceOneTimePriceCurrency = "ANG"
	ProductNewParamsPriceOneTimePriceCurrencyAoa ProductNewParamsPriceOneTimePriceCurrency = "AOA"
	ProductNewParamsPriceOneTimePriceCurrencyArs ProductNewParamsPriceOneTimePriceCurrency = "ARS"
	ProductNewParamsPriceOneTimePriceCurrencyAud ProductNewParamsPriceOneTimePriceCurrency = "AUD"
	ProductNewParamsPriceOneTimePriceCurrencyAwg ProductNewParamsPriceOneTimePriceCurrency = "AWG"
	ProductNewParamsPriceOneTimePriceCurrencyAzn ProductNewParamsPriceOneTimePriceCurrency = "AZN"
	ProductNewParamsPriceOneTimePriceCurrencyBam ProductNewParamsPriceOneTimePriceCurrency = "BAM"
	ProductNewParamsPriceOneTimePriceCurrencyBbd ProductNewParamsPriceOneTimePriceCurrency = "BBD"
	ProductNewParamsPriceOneTimePriceCurrencyBdt ProductNewParamsPriceOneTimePriceCurrency = "BDT"
	ProductNewParamsPriceOneTimePriceCurrencyBgn ProductNewParamsPriceOneTimePriceCurrency = "BGN"
	ProductNewParamsPriceOneTimePriceCurrencyBhd ProductNewParamsPriceOneTimePriceCurrency = "BHD"
	ProductNewParamsPriceOneTimePriceCurrencyBif ProductNewParamsPriceOneTimePriceCurrency = "BIF"
	ProductNewParamsPriceOneTimePriceCurrencyBmd ProductNewParamsPriceOneTimePriceCurrency = "BMD"
	ProductNewParamsPriceOneTimePriceCurrencyBnd ProductNewParamsPriceOneTimePriceCurrency = "BND"
	ProductNewParamsPriceOneTimePriceCurrencyBob ProductNewParamsPriceOneTimePriceCurrency = "BOB"
	ProductNewParamsPriceOneTimePriceCurrencyBrl ProductNewParamsPriceOneTimePriceCurrency = "BRL"
	ProductNewParamsPriceOneTimePriceCurrencyBsd ProductNewParamsPriceOneTimePriceCurrency = "BSD"
	ProductNewParamsPriceOneTimePriceCurrencyBwp ProductNewParamsPriceOneTimePriceCurrency = "BWP"
	ProductNewParamsPriceOneTimePriceCurrencyByn ProductNewParamsPriceOneTimePriceCurrency = "BYN"
	ProductNewParamsPriceOneTimePriceCurrencyBzd ProductNewParamsPriceOneTimePriceCurrency = "BZD"
	ProductNewParamsPriceOneTimePriceCurrencyCad ProductNewParamsPriceOneTimePriceCurrency = "CAD"
	ProductNewParamsPriceOneTimePriceCurrencyChf ProductNewParamsPriceOneTimePriceCurrency = "CHF"
	ProductNewParamsPriceOneTimePriceCurrencyClp ProductNewParamsPriceOneTimePriceCurrency = "CLP"
	ProductNewParamsPriceOneTimePriceCurrencyCny ProductNewParamsPriceOneTimePriceCurrency = "CNY"
	ProductNewParamsPriceOneTimePriceCurrencyCop ProductNewParamsPriceOneTimePriceCurrency = "COP"
	ProductNewParamsPriceOneTimePriceCurrencyCrc ProductNewParamsPriceOneTimePriceCurrency = "CRC"
	ProductNewParamsPriceOneTimePriceCurrencyCup ProductNewParamsPriceOneTimePriceCurrency = "CUP"
	ProductNewParamsPriceOneTimePriceCurrencyCve ProductNewParamsPriceOneTimePriceCurrency = "CVE"
	ProductNewParamsPriceOneTimePriceCurrencyCzk ProductNewParamsPriceOneTimePriceCurrency = "CZK"
	ProductNewParamsPriceOneTimePriceCurrencyDjf ProductNewParamsPriceOneTimePriceCurrency = "DJF"
	ProductNewParamsPriceOneTimePriceCurrencyDkk ProductNewParamsPriceOneTimePriceCurrency = "DKK"
	ProductNewParamsPriceOneTimePriceCurrencyDop ProductNewParamsPriceOneTimePriceCurrency = "DOP"
	ProductNewParamsPriceOneTimePriceCurrencyDzd ProductNewParamsPriceOneTimePriceCurrency = "DZD"
	ProductNewParamsPriceOneTimePriceCurrencyEgp ProductNewParamsPriceOneTimePriceCurrency = "EGP"
	ProductNewParamsPriceOneTimePriceCurrencyEtb ProductNewParamsPriceOneTimePriceCurrency = "ETB"
	ProductNewParamsPriceOneTimePriceCurrencyEur ProductNewParamsPriceOneTimePriceCurrency = "EUR"
	ProductNewParamsPriceOneTimePriceCurrencyFjd ProductNewParamsPriceOneTimePriceCurrency = "FJD"
	ProductNewParamsPriceOneTimePriceCurrencyFkp ProductNewParamsPriceOneTimePriceCurrency = "FKP"
	ProductNewParamsPriceOneTimePriceCurrencyGbp ProductNewParamsPriceOneTimePriceCurrency = "GBP"
	ProductNewParamsPriceOneTimePriceCurrencyGel ProductNewParamsPriceOneTimePriceCurrency = "GEL"
	ProductNewParamsPriceOneTimePriceCurrencyGhs ProductNewParamsPriceOneTimePriceCurrency = "GHS"
	ProductNewParamsPriceOneTimePriceCurrencyGip ProductNewParamsPriceOneTimePriceCurrency = "GIP"
	ProductNewParamsPriceOneTimePriceCurrencyGmd ProductNewParamsPriceOneTimePriceCurrency = "GMD"
	ProductNewParamsPriceOneTimePriceCurrencyGnf ProductNewParamsPriceOneTimePriceCurrency = "GNF"
	ProductNewParamsPriceOneTimePriceCurrencyGtq ProductNewParamsPriceOneTimePriceCurrency = "GTQ"
	ProductNewParamsPriceOneTimePriceCurrencyGyd ProductNewParamsPriceOneTimePriceCurrency = "GYD"
	ProductNewParamsPriceOneTimePriceCurrencyHkd ProductNewParamsPriceOneTimePriceCurrency = "HKD"
	ProductNewParamsPriceOneTimePriceCurrencyHnl ProductNewParamsPriceOneTimePriceCurrency = "HNL"
	ProductNewParamsPriceOneTimePriceCurrencyHrk ProductNewParamsPriceOneTimePriceCurrency = "HRK"
	ProductNewParamsPriceOneTimePriceCurrencyHtg ProductNewParamsPriceOneTimePriceCurrency = "HTG"
	ProductNewParamsPriceOneTimePriceCurrencyHuf ProductNewParamsPriceOneTimePriceCurrency = "HUF"
	ProductNewParamsPriceOneTimePriceCurrencyIdr ProductNewParamsPriceOneTimePriceCurrency = "IDR"
	ProductNewParamsPriceOneTimePriceCurrencyIls ProductNewParamsPriceOneTimePriceCurrency = "ILS"
	ProductNewParamsPriceOneTimePriceCurrencyInr ProductNewParamsPriceOneTimePriceCurrency = "INR"
	ProductNewParamsPriceOneTimePriceCurrencyIqd ProductNewParamsPriceOneTimePriceCurrency = "IQD"
	ProductNewParamsPriceOneTimePriceCurrencyJmd ProductNewParamsPriceOneTimePriceCurrency = "JMD"
	ProductNewParamsPriceOneTimePriceCurrencyJod ProductNewParamsPriceOneTimePriceCurrency = "JOD"
	ProductNewParamsPriceOneTimePriceCurrencyJpy ProductNewParamsPriceOneTimePriceCurrency = "JPY"
	ProductNewParamsPriceOneTimePriceCurrencyKes ProductNewParamsPriceOneTimePriceCurrency = "KES"
	ProductNewParamsPriceOneTimePriceCurrencyKgs ProductNewParamsPriceOneTimePriceCurrency = "KGS"
	ProductNewParamsPriceOneTimePriceCurrencyKhr ProductNewParamsPriceOneTimePriceCurrency = "KHR"
	ProductNewParamsPriceOneTimePriceCurrencyKmf ProductNewParamsPriceOneTimePriceCurrency = "KMF"
	ProductNewParamsPriceOneTimePriceCurrencyKrw ProductNewParamsPriceOneTimePriceCurrency = "KRW"
	ProductNewParamsPriceOneTimePriceCurrencyKwd ProductNewParamsPriceOneTimePriceCurrency = "KWD"
	ProductNewParamsPriceOneTimePriceCurrencyKyd ProductNewParamsPriceOneTimePriceCurrency = "KYD"
	ProductNewParamsPriceOneTimePriceCurrencyKzt ProductNewParamsPriceOneTimePriceCurrency = "KZT"
	ProductNewParamsPriceOneTimePriceCurrencyLak ProductNewParamsPriceOneTimePriceCurrency = "LAK"
	ProductNewParamsPriceOneTimePriceCurrencyLbp ProductNewParamsPriceOneTimePriceCurrency = "LBP"
	ProductNewParamsPriceOneTimePriceCurrencyLkr ProductNewParamsPriceOneTimePriceCurrency = "LKR"
	ProductNewParamsPriceOneTimePriceCurrencyLrd ProductNewParamsPriceOneTimePriceCurrency = "LRD"
	ProductNewParamsPriceOneTimePriceCurrencyLsl ProductNewParamsPriceOneTimePriceCurrency = "LSL"
	ProductNewParamsPriceOneTimePriceCurrencyLyd ProductNewParamsPriceOneTimePriceCurrency = "LYD"
	ProductNewParamsPriceOneTimePriceCurrencyMad ProductNewParamsPriceOneTimePriceCurrency = "MAD"
	ProductNewParamsPriceOneTimePriceCurrencyMdl ProductNewParamsPriceOneTimePriceCurrency = "MDL"
	ProductNewParamsPriceOneTimePriceCurrencyMga ProductNewParamsPriceOneTimePriceCurrency = "MGA"
	ProductNewParamsPriceOneTimePriceCurrencyMkd ProductNewParamsPriceOneTimePriceCurrency = "MKD"
	ProductNewParamsPriceOneTimePriceCurrencyMmk ProductNewParamsPriceOneTimePriceCurrency = "MMK"
	ProductNewParamsPriceOneTimePriceCurrencyMnt ProductNewParamsPriceOneTimePriceCurrency = "MNT"
	ProductNewParamsPriceOneTimePriceCurrencyMop ProductNewParamsPriceOneTimePriceCurrency = "MOP"
	ProductNewParamsPriceOneTimePriceCurrencyMru ProductNewParamsPriceOneTimePriceCurrency = "MRU"
	ProductNewParamsPriceOneTimePriceCurrencyMur ProductNewParamsPriceOneTimePriceCurrency = "MUR"
	ProductNewParamsPriceOneTimePriceCurrencyMvr ProductNewParamsPriceOneTimePriceCurrency = "MVR"
	ProductNewParamsPriceOneTimePriceCurrencyMwk ProductNewParamsPriceOneTimePriceCurrency = "MWK"
	ProductNewParamsPriceOneTimePriceCurrencyMxn ProductNewParamsPriceOneTimePriceCurrency = "MXN"
	ProductNewParamsPriceOneTimePriceCurrencyMyr ProductNewParamsPriceOneTimePriceCurrency = "MYR"
	ProductNewParamsPriceOneTimePriceCurrencyMzn ProductNewParamsPriceOneTimePriceCurrency = "MZN"
	ProductNewParamsPriceOneTimePriceCurrencyNad ProductNewParamsPriceOneTimePriceCurrency = "NAD"
	ProductNewParamsPriceOneTimePriceCurrencyNgn ProductNewParamsPriceOneTimePriceCurrency = "NGN"
	ProductNewParamsPriceOneTimePriceCurrencyNio ProductNewParamsPriceOneTimePriceCurrency = "NIO"
	ProductNewParamsPriceOneTimePriceCurrencyNok ProductNewParamsPriceOneTimePriceCurrency = "NOK"
	ProductNewParamsPriceOneTimePriceCurrencyNpr ProductNewParamsPriceOneTimePriceCurrency = "NPR"
	ProductNewParamsPriceOneTimePriceCurrencyNzd ProductNewParamsPriceOneTimePriceCurrency = "NZD"
	ProductNewParamsPriceOneTimePriceCurrencyOmr ProductNewParamsPriceOneTimePriceCurrency = "OMR"
	ProductNewParamsPriceOneTimePriceCurrencyPab ProductNewParamsPriceOneTimePriceCurrency = "PAB"
	ProductNewParamsPriceOneTimePriceCurrencyPen ProductNewParamsPriceOneTimePriceCurrency = "PEN"
	ProductNewParamsPriceOneTimePriceCurrencyPgk ProductNewParamsPriceOneTimePriceCurrency = "PGK"
	ProductNewParamsPriceOneTimePriceCurrencyPhp ProductNewParamsPriceOneTimePriceCurrency = "PHP"
	ProductNewParamsPriceOneTimePriceCurrencyPkr ProductNewParamsPriceOneTimePriceCurrency = "PKR"
	ProductNewParamsPriceOneTimePriceCurrencyPln ProductNewParamsPriceOneTimePriceCurrency = "PLN"
	ProductNewParamsPriceOneTimePriceCurrencyPyg ProductNewParamsPriceOneTimePriceCurrency = "PYG"
	ProductNewParamsPriceOneTimePriceCurrencyQar ProductNewParamsPriceOneTimePriceCurrency = "QAR"
	ProductNewParamsPriceOneTimePriceCurrencyRon ProductNewParamsPriceOneTimePriceCurrency = "RON"
	ProductNewParamsPriceOneTimePriceCurrencyRsd ProductNewParamsPriceOneTimePriceCurrency = "RSD"
	ProductNewParamsPriceOneTimePriceCurrencyRub ProductNewParamsPriceOneTimePriceCurrency = "RUB"
	ProductNewParamsPriceOneTimePriceCurrencyRwf ProductNewParamsPriceOneTimePriceCurrency = "RWF"
	ProductNewParamsPriceOneTimePriceCurrencySar ProductNewParamsPriceOneTimePriceCurrency = "SAR"
	ProductNewParamsPriceOneTimePriceCurrencySbd ProductNewParamsPriceOneTimePriceCurrency = "SBD"
	ProductNewParamsPriceOneTimePriceCurrencyScr ProductNewParamsPriceOneTimePriceCurrency = "SCR"
	ProductNewParamsPriceOneTimePriceCurrencySek ProductNewParamsPriceOneTimePriceCurrency = "SEK"
	ProductNewParamsPriceOneTimePriceCurrencySgd ProductNewParamsPriceOneTimePriceCurrency = "SGD"
	ProductNewParamsPriceOneTimePriceCurrencyShp ProductNewParamsPriceOneTimePriceCurrency = "SHP"
	ProductNewParamsPriceOneTimePriceCurrencySle ProductNewParamsPriceOneTimePriceCurrency = "SLE"
	ProductNewParamsPriceOneTimePriceCurrencySll ProductNewParamsPriceOneTimePriceCurrency = "SLL"
	ProductNewParamsPriceOneTimePriceCurrencySos ProductNewParamsPriceOneTimePriceCurrency = "SOS"
	ProductNewParamsPriceOneTimePriceCurrencySrd ProductNewParamsPriceOneTimePriceCurrency = "SRD"
	ProductNewParamsPriceOneTimePriceCurrencySsp ProductNewParamsPriceOneTimePriceCurrency = "SSP"
	ProductNewParamsPriceOneTimePriceCurrencyStn ProductNewParamsPriceOneTimePriceCurrency = "STN"
	ProductNewParamsPriceOneTimePriceCurrencySvc ProductNewParamsPriceOneTimePriceCurrency = "SVC"
	ProductNewParamsPriceOneTimePriceCurrencySzl ProductNewParamsPriceOneTimePriceCurrency = "SZL"
	ProductNewParamsPriceOneTimePriceCurrencyThb ProductNewParamsPriceOneTimePriceCurrency = "THB"
	ProductNewParamsPriceOneTimePriceCurrencyTnd ProductNewParamsPriceOneTimePriceCurrency = "TND"
	ProductNewParamsPriceOneTimePriceCurrencyTop ProductNewParamsPriceOneTimePriceCurrency = "TOP"
	ProductNewParamsPriceOneTimePriceCurrencyTry ProductNewParamsPriceOneTimePriceCurrency = "TRY"
	ProductNewParamsPriceOneTimePriceCurrencyTtd ProductNewParamsPriceOneTimePriceCurrency = "TTD"
	ProductNewParamsPriceOneTimePriceCurrencyTwd ProductNewParamsPriceOneTimePriceCurrency = "TWD"
	ProductNewParamsPriceOneTimePriceCurrencyTzs ProductNewParamsPriceOneTimePriceCurrency = "TZS"
	ProductNewParamsPriceOneTimePriceCurrencyUah ProductNewParamsPriceOneTimePriceCurrency = "UAH"
	ProductNewParamsPriceOneTimePriceCurrencyUgx ProductNewParamsPriceOneTimePriceCurrency = "UGX"
	ProductNewParamsPriceOneTimePriceCurrencyUsd ProductNewParamsPriceOneTimePriceCurrency = "USD"
	ProductNewParamsPriceOneTimePriceCurrencyUyu ProductNewParamsPriceOneTimePriceCurrency = "UYU"
	ProductNewParamsPriceOneTimePriceCurrencyUzs ProductNewParamsPriceOneTimePriceCurrency = "UZS"
	ProductNewParamsPriceOneTimePriceCurrencyVes ProductNewParamsPriceOneTimePriceCurrency = "VES"
	ProductNewParamsPriceOneTimePriceCurrencyVnd ProductNewParamsPriceOneTimePriceCurrency = "VND"
	ProductNewParamsPriceOneTimePriceCurrencyVuv ProductNewParamsPriceOneTimePriceCurrency = "VUV"
	ProductNewParamsPriceOneTimePriceCurrencyWst ProductNewParamsPriceOneTimePriceCurrency = "WST"
	ProductNewParamsPriceOneTimePriceCurrencyXaf ProductNewParamsPriceOneTimePriceCurrency = "XAF"
	ProductNewParamsPriceOneTimePriceCurrencyXcd ProductNewParamsPriceOneTimePriceCurrency = "XCD"
	ProductNewParamsPriceOneTimePriceCurrencyXof ProductNewParamsPriceOneTimePriceCurrency = "XOF"
	ProductNewParamsPriceOneTimePriceCurrencyXpf ProductNewParamsPriceOneTimePriceCurrency = "XPF"
	ProductNewParamsPriceOneTimePriceCurrencyYer ProductNewParamsPriceOneTimePriceCurrency = "YER"
	ProductNewParamsPriceOneTimePriceCurrencyZar ProductNewParamsPriceOneTimePriceCurrency = "ZAR"
	ProductNewParamsPriceOneTimePriceCurrencyZmw ProductNewParamsPriceOneTimePriceCurrency = "ZMW"
)

func (r ProductNewParamsPriceOneTimePriceCurrency) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceOneTimePriceCurrencyAed, ProductNewParamsPriceOneTimePriceCurrencyAll, ProductNewParamsPriceOneTimePriceCurrencyAmd, ProductNewParamsPriceOneTimePriceCurrencyAng, ProductNewParamsPriceOneTimePriceCurrencyAoa, ProductNewParamsPriceOneTimePriceCurrencyArs, ProductNewParamsPriceOneTimePriceCurrencyAud, ProductNewParamsPriceOneTimePriceCurrencyAwg, ProductNewParamsPriceOneTimePriceCurrencyAzn, ProductNewParamsPriceOneTimePriceCurrencyBam, ProductNewParamsPriceOneTimePriceCurrencyBbd, ProductNewParamsPriceOneTimePriceCurrencyBdt, ProductNewParamsPriceOneTimePriceCurrencyBgn, ProductNewParamsPriceOneTimePriceCurrencyBhd, ProductNewParamsPriceOneTimePriceCurrencyBif, ProductNewParamsPriceOneTimePriceCurrencyBmd, ProductNewParamsPriceOneTimePriceCurrencyBnd, ProductNewParamsPriceOneTimePriceCurrencyBob, ProductNewParamsPriceOneTimePriceCurrencyBrl, ProductNewParamsPriceOneTimePriceCurrencyBsd, ProductNewParamsPriceOneTimePriceCurrencyBwp, ProductNewParamsPriceOneTimePriceCurrencyByn, ProductNewParamsPriceOneTimePriceCurrencyBzd, ProductNewParamsPriceOneTimePriceCurrencyCad, ProductNewParamsPriceOneTimePriceCurrencyChf, ProductNewParamsPriceOneTimePriceCurrencyClp, ProductNewParamsPriceOneTimePriceCurrencyCny, ProductNewParamsPriceOneTimePriceCurrencyCop, ProductNewParamsPriceOneTimePriceCurrencyCrc, ProductNewParamsPriceOneTimePriceCurrencyCup, ProductNewParamsPriceOneTimePriceCurrencyCve, ProductNewParamsPriceOneTimePriceCurrencyCzk, ProductNewParamsPriceOneTimePriceCurrencyDjf, ProductNewParamsPriceOneTimePriceCurrencyDkk, ProductNewParamsPriceOneTimePriceCurrencyDop, ProductNewParamsPriceOneTimePriceCurrencyDzd, ProductNewParamsPriceOneTimePriceCurrencyEgp, ProductNewParamsPriceOneTimePriceCurrencyEtb, ProductNewParamsPriceOneTimePriceCurrencyEur, ProductNewParamsPriceOneTimePriceCurrencyFjd, ProductNewParamsPriceOneTimePriceCurrencyFkp, ProductNewParamsPriceOneTimePriceCurrencyGbp, ProductNewParamsPriceOneTimePriceCurrencyGel, ProductNewParamsPriceOneTimePriceCurrencyGhs, ProductNewParamsPriceOneTimePriceCurrencyGip, ProductNewParamsPriceOneTimePriceCurrencyGmd, ProductNewParamsPriceOneTimePriceCurrencyGnf, ProductNewParamsPriceOneTimePriceCurrencyGtq, ProductNewParamsPriceOneTimePriceCurrencyGyd, ProductNewParamsPriceOneTimePriceCurrencyHkd, ProductNewParamsPriceOneTimePriceCurrencyHnl, ProductNewParamsPriceOneTimePriceCurrencyHrk, ProductNewParamsPriceOneTimePriceCurrencyHtg, ProductNewParamsPriceOneTimePriceCurrencyHuf, ProductNewParamsPriceOneTimePriceCurrencyIdr, ProductNewParamsPriceOneTimePriceCurrencyIls, ProductNewParamsPriceOneTimePriceCurrencyInr, ProductNewParamsPriceOneTimePriceCurrencyIqd, ProductNewParamsPriceOneTimePriceCurrencyJmd, ProductNewParamsPriceOneTimePriceCurrencyJod, ProductNewParamsPriceOneTimePriceCurrencyJpy, ProductNewParamsPriceOneTimePriceCurrencyKes, ProductNewParamsPriceOneTimePriceCurrencyKgs, ProductNewParamsPriceOneTimePriceCurrencyKhr, ProductNewParamsPriceOneTimePriceCurrencyKmf, ProductNewParamsPriceOneTimePriceCurrencyKrw, ProductNewParamsPriceOneTimePriceCurrencyKwd, ProductNewParamsPriceOneTimePriceCurrencyKyd, ProductNewParamsPriceOneTimePriceCurrencyKzt, ProductNewParamsPriceOneTimePriceCurrencyLak, ProductNewParamsPriceOneTimePriceCurrencyLbp, ProductNewParamsPriceOneTimePriceCurrencyLkr, ProductNewParamsPriceOneTimePriceCurrencyLrd, ProductNewParamsPriceOneTimePriceCurrencyLsl, ProductNewParamsPriceOneTimePriceCurrencyLyd, ProductNewParamsPriceOneTimePriceCurrencyMad, ProductNewParamsPriceOneTimePriceCurrencyMdl, ProductNewParamsPriceOneTimePriceCurrencyMga, ProductNewParamsPriceOneTimePriceCurrencyMkd, ProductNewParamsPriceOneTimePriceCurrencyMmk, ProductNewParamsPriceOneTimePriceCurrencyMnt, ProductNewParamsPriceOneTimePriceCurrencyMop, ProductNewParamsPriceOneTimePriceCurrencyMru, ProductNewParamsPriceOneTimePriceCurrencyMur, ProductNewParamsPriceOneTimePriceCurrencyMvr, ProductNewParamsPriceOneTimePriceCurrencyMwk, ProductNewParamsPriceOneTimePriceCurrencyMxn, ProductNewParamsPriceOneTimePriceCurrencyMyr, ProductNewParamsPriceOneTimePriceCurrencyMzn, ProductNewParamsPriceOneTimePriceCurrencyNad, ProductNewParamsPriceOneTimePriceCurrencyNgn, ProductNewParamsPriceOneTimePriceCurrencyNio, ProductNewParamsPriceOneTimePriceCurrencyNok, ProductNewParamsPriceOneTimePriceCurrencyNpr, ProductNewParamsPriceOneTimePriceCurrencyNzd, ProductNewParamsPriceOneTimePriceCurrencyOmr, ProductNewParamsPriceOneTimePriceCurrencyPab, ProductNewParamsPriceOneTimePriceCurrencyPen, ProductNewParamsPriceOneTimePriceCurrencyPgk, ProductNewParamsPriceOneTimePriceCurrencyPhp, ProductNewParamsPriceOneTimePriceCurrencyPkr, ProductNewParamsPriceOneTimePriceCurrencyPln, ProductNewParamsPriceOneTimePriceCurrencyPyg, ProductNewParamsPriceOneTimePriceCurrencyQar, ProductNewParamsPriceOneTimePriceCurrencyRon, ProductNewParamsPriceOneTimePriceCurrencyRsd, ProductNewParamsPriceOneTimePriceCurrencyRub, ProductNewParamsPriceOneTimePriceCurrencyRwf, ProductNewParamsPriceOneTimePriceCurrencySar, ProductNewParamsPriceOneTimePriceCurrencySbd, ProductNewParamsPriceOneTimePriceCurrencyScr, ProductNewParamsPriceOneTimePriceCurrencySek, ProductNewParamsPriceOneTimePriceCurrencySgd, ProductNewParamsPriceOneTimePriceCurrencyShp, ProductNewParamsPriceOneTimePriceCurrencySle, ProductNewParamsPriceOneTimePriceCurrencySll, ProductNewParamsPriceOneTimePriceCurrencySos, ProductNewParamsPriceOneTimePriceCurrencySrd, ProductNewParamsPriceOneTimePriceCurrencySsp, ProductNewParamsPriceOneTimePriceCurrencyStn, ProductNewParamsPriceOneTimePriceCurrencySvc, ProductNewParamsPriceOneTimePriceCurrencySzl, ProductNewParamsPriceOneTimePriceCurrencyThb, ProductNewParamsPriceOneTimePriceCurrencyTnd, ProductNewParamsPriceOneTimePriceCurrencyTop, ProductNewParamsPriceOneTimePriceCurrencyTry, ProductNewParamsPriceOneTimePriceCurrencyTtd, ProductNewParamsPriceOneTimePriceCurrencyTwd, ProductNewParamsPriceOneTimePriceCurrencyTzs, ProductNewParamsPriceOneTimePriceCurrencyUah, ProductNewParamsPriceOneTimePriceCurrencyUgx, ProductNewParamsPriceOneTimePriceCurrencyUsd, ProductNewParamsPriceOneTimePriceCurrencyUyu, ProductNewParamsPriceOneTimePriceCurrencyUzs, ProductNewParamsPriceOneTimePriceCurrencyVes, ProductNewParamsPriceOneTimePriceCurrencyVnd, ProductNewParamsPriceOneTimePriceCurrencyVuv, ProductNewParamsPriceOneTimePriceCurrencyWst, ProductNewParamsPriceOneTimePriceCurrencyXaf, ProductNewParamsPriceOneTimePriceCurrencyXcd, ProductNewParamsPriceOneTimePriceCurrencyXof, ProductNewParamsPriceOneTimePriceCurrencyXpf, ProductNewParamsPriceOneTimePriceCurrencyYer, ProductNewParamsPriceOneTimePriceCurrencyZar, ProductNewParamsPriceOneTimePriceCurrencyZmw:
		return true
	}
	return false
}

type ProductNewParamsPriceOneTimePriceType string

const (
	ProductNewParamsPriceOneTimePriceTypeOneTimePrice ProductNewParamsPriceOneTimePriceType = "one_time_price"
)

func (r ProductNewParamsPriceOneTimePriceType) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceOneTimePriceTypeOneTimePrice:
		return true
	}
	return false
}

type ProductNewParamsPriceRecurringPrice struct {
	Currency param.Field[ProductNewParamsPriceRecurringPriceCurrency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    param.Field[int64]                                                       `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval param.Field[ProductNewParamsPriceRecurringPricePaymentFrequencyInterval] `json:"payment_frequency_interval,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool] `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    param.Field[int64]                                                         `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval param.Field[ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval] `json:"subscription_period_interval,required"`
	Type                       param.Field[ProductNewParamsPriceRecurringPriceType]                       `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r ProductNewParamsPriceRecurringPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductNewParamsPriceRecurringPrice) implementsProductNewParamsPriceUnion() {}

type ProductNewParamsPriceRecurringPriceCurrency string

const (
	ProductNewParamsPriceRecurringPriceCurrencyAed ProductNewParamsPriceRecurringPriceCurrency = "AED"
	ProductNewParamsPriceRecurringPriceCurrencyAll ProductNewParamsPriceRecurringPriceCurrency = "ALL"
	ProductNewParamsPriceRecurringPriceCurrencyAmd ProductNewParamsPriceRecurringPriceCurrency = "AMD"
	ProductNewParamsPriceRecurringPriceCurrencyAng ProductNewParamsPriceRecurringPriceCurrency = "ANG"
	ProductNewParamsPriceRecurringPriceCurrencyAoa ProductNewParamsPriceRecurringPriceCurrency = "AOA"
	ProductNewParamsPriceRecurringPriceCurrencyArs ProductNewParamsPriceRecurringPriceCurrency = "ARS"
	ProductNewParamsPriceRecurringPriceCurrencyAud ProductNewParamsPriceRecurringPriceCurrency = "AUD"
	ProductNewParamsPriceRecurringPriceCurrencyAwg ProductNewParamsPriceRecurringPriceCurrency = "AWG"
	ProductNewParamsPriceRecurringPriceCurrencyAzn ProductNewParamsPriceRecurringPriceCurrency = "AZN"
	ProductNewParamsPriceRecurringPriceCurrencyBam ProductNewParamsPriceRecurringPriceCurrency = "BAM"
	ProductNewParamsPriceRecurringPriceCurrencyBbd ProductNewParamsPriceRecurringPriceCurrency = "BBD"
	ProductNewParamsPriceRecurringPriceCurrencyBdt ProductNewParamsPriceRecurringPriceCurrency = "BDT"
	ProductNewParamsPriceRecurringPriceCurrencyBgn ProductNewParamsPriceRecurringPriceCurrency = "BGN"
	ProductNewParamsPriceRecurringPriceCurrencyBhd ProductNewParamsPriceRecurringPriceCurrency = "BHD"
	ProductNewParamsPriceRecurringPriceCurrencyBif ProductNewParamsPriceRecurringPriceCurrency = "BIF"
	ProductNewParamsPriceRecurringPriceCurrencyBmd ProductNewParamsPriceRecurringPriceCurrency = "BMD"
	ProductNewParamsPriceRecurringPriceCurrencyBnd ProductNewParamsPriceRecurringPriceCurrency = "BND"
	ProductNewParamsPriceRecurringPriceCurrencyBob ProductNewParamsPriceRecurringPriceCurrency = "BOB"
	ProductNewParamsPriceRecurringPriceCurrencyBrl ProductNewParamsPriceRecurringPriceCurrency = "BRL"
	ProductNewParamsPriceRecurringPriceCurrencyBsd ProductNewParamsPriceRecurringPriceCurrency = "BSD"
	ProductNewParamsPriceRecurringPriceCurrencyBwp ProductNewParamsPriceRecurringPriceCurrency = "BWP"
	ProductNewParamsPriceRecurringPriceCurrencyByn ProductNewParamsPriceRecurringPriceCurrency = "BYN"
	ProductNewParamsPriceRecurringPriceCurrencyBzd ProductNewParamsPriceRecurringPriceCurrency = "BZD"
	ProductNewParamsPriceRecurringPriceCurrencyCad ProductNewParamsPriceRecurringPriceCurrency = "CAD"
	ProductNewParamsPriceRecurringPriceCurrencyChf ProductNewParamsPriceRecurringPriceCurrency = "CHF"
	ProductNewParamsPriceRecurringPriceCurrencyClp ProductNewParamsPriceRecurringPriceCurrency = "CLP"
	ProductNewParamsPriceRecurringPriceCurrencyCny ProductNewParamsPriceRecurringPriceCurrency = "CNY"
	ProductNewParamsPriceRecurringPriceCurrencyCop ProductNewParamsPriceRecurringPriceCurrency = "COP"
	ProductNewParamsPriceRecurringPriceCurrencyCrc ProductNewParamsPriceRecurringPriceCurrency = "CRC"
	ProductNewParamsPriceRecurringPriceCurrencyCup ProductNewParamsPriceRecurringPriceCurrency = "CUP"
	ProductNewParamsPriceRecurringPriceCurrencyCve ProductNewParamsPriceRecurringPriceCurrency = "CVE"
	ProductNewParamsPriceRecurringPriceCurrencyCzk ProductNewParamsPriceRecurringPriceCurrency = "CZK"
	ProductNewParamsPriceRecurringPriceCurrencyDjf ProductNewParamsPriceRecurringPriceCurrency = "DJF"
	ProductNewParamsPriceRecurringPriceCurrencyDkk ProductNewParamsPriceRecurringPriceCurrency = "DKK"
	ProductNewParamsPriceRecurringPriceCurrencyDop ProductNewParamsPriceRecurringPriceCurrency = "DOP"
	ProductNewParamsPriceRecurringPriceCurrencyDzd ProductNewParamsPriceRecurringPriceCurrency = "DZD"
	ProductNewParamsPriceRecurringPriceCurrencyEgp ProductNewParamsPriceRecurringPriceCurrency = "EGP"
	ProductNewParamsPriceRecurringPriceCurrencyEtb ProductNewParamsPriceRecurringPriceCurrency = "ETB"
	ProductNewParamsPriceRecurringPriceCurrencyEur ProductNewParamsPriceRecurringPriceCurrency = "EUR"
	ProductNewParamsPriceRecurringPriceCurrencyFjd ProductNewParamsPriceRecurringPriceCurrency = "FJD"
	ProductNewParamsPriceRecurringPriceCurrencyFkp ProductNewParamsPriceRecurringPriceCurrency = "FKP"
	ProductNewParamsPriceRecurringPriceCurrencyGbp ProductNewParamsPriceRecurringPriceCurrency = "GBP"
	ProductNewParamsPriceRecurringPriceCurrencyGel ProductNewParamsPriceRecurringPriceCurrency = "GEL"
	ProductNewParamsPriceRecurringPriceCurrencyGhs ProductNewParamsPriceRecurringPriceCurrency = "GHS"
	ProductNewParamsPriceRecurringPriceCurrencyGip ProductNewParamsPriceRecurringPriceCurrency = "GIP"
	ProductNewParamsPriceRecurringPriceCurrencyGmd ProductNewParamsPriceRecurringPriceCurrency = "GMD"
	ProductNewParamsPriceRecurringPriceCurrencyGnf ProductNewParamsPriceRecurringPriceCurrency = "GNF"
	ProductNewParamsPriceRecurringPriceCurrencyGtq ProductNewParamsPriceRecurringPriceCurrency = "GTQ"
	ProductNewParamsPriceRecurringPriceCurrencyGyd ProductNewParamsPriceRecurringPriceCurrency = "GYD"
	ProductNewParamsPriceRecurringPriceCurrencyHkd ProductNewParamsPriceRecurringPriceCurrency = "HKD"
	ProductNewParamsPriceRecurringPriceCurrencyHnl ProductNewParamsPriceRecurringPriceCurrency = "HNL"
	ProductNewParamsPriceRecurringPriceCurrencyHrk ProductNewParamsPriceRecurringPriceCurrency = "HRK"
	ProductNewParamsPriceRecurringPriceCurrencyHtg ProductNewParamsPriceRecurringPriceCurrency = "HTG"
	ProductNewParamsPriceRecurringPriceCurrencyHuf ProductNewParamsPriceRecurringPriceCurrency = "HUF"
	ProductNewParamsPriceRecurringPriceCurrencyIdr ProductNewParamsPriceRecurringPriceCurrency = "IDR"
	ProductNewParamsPriceRecurringPriceCurrencyIls ProductNewParamsPriceRecurringPriceCurrency = "ILS"
	ProductNewParamsPriceRecurringPriceCurrencyInr ProductNewParamsPriceRecurringPriceCurrency = "INR"
	ProductNewParamsPriceRecurringPriceCurrencyIqd ProductNewParamsPriceRecurringPriceCurrency = "IQD"
	ProductNewParamsPriceRecurringPriceCurrencyJmd ProductNewParamsPriceRecurringPriceCurrency = "JMD"
	ProductNewParamsPriceRecurringPriceCurrencyJod ProductNewParamsPriceRecurringPriceCurrency = "JOD"
	ProductNewParamsPriceRecurringPriceCurrencyJpy ProductNewParamsPriceRecurringPriceCurrency = "JPY"
	ProductNewParamsPriceRecurringPriceCurrencyKes ProductNewParamsPriceRecurringPriceCurrency = "KES"
	ProductNewParamsPriceRecurringPriceCurrencyKgs ProductNewParamsPriceRecurringPriceCurrency = "KGS"
	ProductNewParamsPriceRecurringPriceCurrencyKhr ProductNewParamsPriceRecurringPriceCurrency = "KHR"
	ProductNewParamsPriceRecurringPriceCurrencyKmf ProductNewParamsPriceRecurringPriceCurrency = "KMF"
	ProductNewParamsPriceRecurringPriceCurrencyKrw ProductNewParamsPriceRecurringPriceCurrency = "KRW"
	ProductNewParamsPriceRecurringPriceCurrencyKwd ProductNewParamsPriceRecurringPriceCurrency = "KWD"
	ProductNewParamsPriceRecurringPriceCurrencyKyd ProductNewParamsPriceRecurringPriceCurrency = "KYD"
	ProductNewParamsPriceRecurringPriceCurrencyKzt ProductNewParamsPriceRecurringPriceCurrency = "KZT"
	ProductNewParamsPriceRecurringPriceCurrencyLak ProductNewParamsPriceRecurringPriceCurrency = "LAK"
	ProductNewParamsPriceRecurringPriceCurrencyLbp ProductNewParamsPriceRecurringPriceCurrency = "LBP"
	ProductNewParamsPriceRecurringPriceCurrencyLkr ProductNewParamsPriceRecurringPriceCurrency = "LKR"
	ProductNewParamsPriceRecurringPriceCurrencyLrd ProductNewParamsPriceRecurringPriceCurrency = "LRD"
	ProductNewParamsPriceRecurringPriceCurrencyLsl ProductNewParamsPriceRecurringPriceCurrency = "LSL"
	ProductNewParamsPriceRecurringPriceCurrencyLyd ProductNewParamsPriceRecurringPriceCurrency = "LYD"
	ProductNewParamsPriceRecurringPriceCurrencyMad ProductNewParamsPriceRecurringPriceCurrency = "MAD"
	ProductNewParamsPriceRecurringPriceCurrencyMdl ProductNewParamsPriceRecurringPriceCurrency = "MDL"
	ProductNewParamsPriceRecurringPriceCurrencyMga ProductNewParamsPriceRecurringPriceCurrency = "MGA"
	ProductNewParamsPriceRecurringPriceCurrencyMkd ProductNewParamsPriceRecurringPriceCurrency = "MKD"
	ProductNewParamsPriceRecurringPriceCurrencyMmk ProductNewParamsPriceRecurringPriceCurrency = "MMK"
	ProductNewParamsPriceRecurringPriceCurrencyMnt ProductNewParamsPriceRecurringPriceCurrency = "MNT"
	ProductNewParamsPriceRecurringPriceCurrencyMop ProductNewParamsPriceRecurringPriceCurrency = "MOP"
	ProductNewParamsPriceRecurringPriceCurrencyMru ProductNewParamsPriceRecurringPriceCurrency = "MRU"
	ProductNewParamsPriceRecurringPriceCurrencyMur ProductNewParamsPriceRecurringPriceCurrency = "MUR"
	ProductNewParamsPriceRecurringPriceCurrencyMvr ProductNewParamsPriceRecurringPriceCurrency = "MVR"
	ProductNewParamsPriceRecurringPriceCurrencyMwk ProductNewParamsPriceRecurringPriceCurrency = "MWK"
	ProductNewParamsPriceRecurringPriceCurrencyMxn ProductNewParamsPriceRecurringPriceCurrency = "MXN"
	ProductNewParamsPriceRecurringPriceCurrencyMyr ProductNewParamsPriceRecurringPriceCurrency = "MYR"
	ProductNewParamsPriceRecurringPriceCurrencyMzn ProductNewParamsPriceRecurringPriceCurrency = "MZN"
	ProductNewParamsPriceRecurringPriceCurrencyNad ProductNewParamsPriceRecurringPriceCurrency = "NAD"
	ProductNewParamsPriceRecurringPriceCurrencyNgn ProductNewParamsPriceRecurringPriceCurrency = "NGN"
	ProductNewParamsPriceRecurringPriceCurrencyNio ProductNewParamsPriceRecurringPriceCurrency = "NIO"
	ProductNewParamsPriceRecurringPriceCurrencyNok ProductNewParamsPriceRecurringPriceCurrency = "NOK"
	ProductNewParamsPriceRecurringPriceCurrencyNpr ProductNewParamsPriceRecurringPriceCurrency = "NPR"
	ProductNewParamsPriceRecurringPriceCurrencyNzd ProductNewParamsPriceRecurringPriceCurrency = "NZD"
	ProductNewParamsPriceRecurringPriceCurrencyOmr ProductNewParamsPriceRecurringPriceCurrency = "OMR"
	ProductNewParamsPriceRecurringPriceCurrencyPab ProductNewParamsPriceRecurringPriceCurrency = "PAB"
	ProductNewParamsPriceRecurringPriceCurrencyPen ProductNewParamsPriceRecurringPriceCurrency = "PEN"
	ProductNewParamsPriceRecurringPriceCurrencyPgk ProductNewParamsPriceRecurringPriceCurrency = "PGK"
	ProductNewParamsPriceRecurringPriceCurrencyPhp ProductNewParamsPriceRecurringPriceCurrency = "PHP"
	ProductNewParamsPriceRecurringPriceCurrencyPkr ProductNewParamsPriceRecurringPriceCurrency = "PKR"
	ProductNewParamsPriceRecurringPriceCurrencyPln ProductNewParamsPriceRecurringPriceCurrency = "PLN"
	ProductNewParamsPriceRecurringPriceCurrencyPyg ProductNewParamsPriceRecurringPriceCurrency = "PYG"
	ProductNewParamsPriceRecurringPriceCurrencyQar ProductNewParamsPriceRecurringPriceCurrency = "QAR"
	ProductNewParamsPriceRecurringPriceCurrencyRon ProductNewParamsPriceRecurringPriceCurrency = "RON"
	ProductNewParamsPriceRecurringPriceCurrencyRsd ProductNewParamsPriceRecurringPriceCurrency = "RSD"
	ProductNewParamsPriceRecurringPriceCurrencyRub ProductNewParamsPriceRecurringPriceCurrency = "RUB"
	ProductNewParamsPriceRecurringPriceCurrencyRwf ProductNewParamsPriceRecurringPriceCurrency = "RWF"
	ProductNewParamsPriceRecurringPriceCurrencySar ProductNewParamsPriceRecurringPriceCurrency = "SAR"
	ProductNewParamsPriceRecurringPriceCurrencySbd ProductNewParamsPriceRecurringPriceCurrency = "SBD"
	ProductNewParamsPriceRecurringPriceCurrencyScr ProductNewParamsPriceRecurringPriceCurrency = "SCR"
	ProductNewParamsPriceRecurringPriceCurrencySek ProductNewParamsPriceRecurringPriceCurrency = "SEK"
	ProductNewParamsPriceRecurringPriceCurrencySgd ProductNewParamsPriceRecurringPriceCurrency = "SGD"
	ProductNewParamsPriceRecurringPriceCurrencyShp ProductNewParamsPriceRecurringPriceCurrency = "SHP"
	ProductNewParamsPriceRecurringPriceCurrencySle ProductNewParamsPriceRecurringPriceCurrency = "SLE"
	ProductNewParamsPriceRecurringPriceCurrencySll ProductNewParamsPriceRecurringPriceCurrency = "SLL"
	ProductNewParamsPriceRecurringPriceCurrencySos ProductNewParamsPriceRecurringPriceCurrency = "SOS"
	ProductNewParamsPriceRecurringPriceCurrencySrd ProductNewParamsPriceRecurringPriceCurrency = "SRD"
	ProductNewParamsPriceRecurringPriceCurrencySsp ProductNewParamsPriceRecurringPriceCurrency = "SSP"
	ProductNewParamsPriceRecurringPriceCurrencyStn ProductNewParamsPriceRecurringPriceCurrency = "STN"
	ProductNewParamsPriceRecurringPriceCurrencySvc ProductNewParamsPriceRecurringPriceCurrency = "SVC"
	ProductNewParamsPriceRecurringPriceCurrencySzl ProductNewParamsPriceRecurringPriceCurrency = "SZL"
	ProductNewParamsPriceRecurringPriceCurrencyThb ProductNewParamsPriceRecurringPriceCurrency = "THB"
	ProductNewParamsPriceRecurringPriceCurrencyTnd ProductNewParamsPriceRecurringPriceCurrency = "TND"
	ProductNewParamsPriceRecurringPriceCurrencyTop ProductNewParamsPriceRecurringPriceCurrency = "TOP"
	ProductNewParamsPriceRecurringPriceCurrencyTry ProductNewParamsPriceRecurringPriceCurrency = "TRY"
	ProductNewParamsPriceRecurringPriceCurrencyTtd ProductNewParamsPriceRecurringPriceCurrency = "TTD"
	ProductNewParamsPriceRecurringPriceCurrencyTwd ProductNewParamsPriceRecurringPriceCurrency = "TWD"
	ProductNewParamsPriceRecurringPriceCurrencyTzs ProductNewParamsPriceRecurringPriceCurrency = "TZS"
	ProductNewParamsPriceRecurringPriceCurrencyUah ProductNewParamsPriceRecurringPriceCurrency = "UAH"
	ProductNewParamsPriceRecurringPriceCurrencyUgx ProductNewParamsPriceRecurringPriceCurrency = "UGX"
	ProductNewParamsPriceRecurringPriceCurrencyUsd ProductNewParamsPriceRecurringPriceCurrency = "USD"
	ProductNewParamsPriceRecurringPriceCurrencyUyu ProductNewParamsPriceRecurringPriceCurrency = "UYU"
	ProductNewParamsPriceRecurringPriceCurrencyUzs ProductNewParamsPriceRecurringPriceCurrency = "UZS"
	ProductNewParamsPriceRecurringPriceCurrencyVes ProductNewParamsPriceRecurringPriceCurrency = "VES"
	ProductNewParamsPriceRecurringPriceCurrencyVnd ProductNewParamsPriceRecurringPriceCurrency = "VND"
	ProductNewParamsPriceRecurringPriceCurrencyVuv ProductNewParamsPriceRecurringPriceCurrency = "VUV"
	ProductNewParamsPriceRecurringPriceCurrencyWst ProductNewParamsPriceRecurringPriceCurrency = "WST"
	ProductNewParamsPriceRecurringPriceCurrencyXaf ProductNewParamsPriceRecurringPriceCurrency = "XAF"
	ProductNewParamsPriceRecurringPriceCurrencyXcd ProductNewParamsPriceRecurringPriceCurrency = "XCD"
	ProductNewParamsPriceRecurringPriceCurrencyXof ProductNewParamsPriceRecurringPriceCurrency = "XOF"
	ProductNewParamsPriceRecurringPriceCurrencyXpf ProductNewParamsPriceRecurringPriceCurrency = "XPF"
	ProductNewParamsPriceRecurringPriceCurrencyYer ProductNewParamsPriceRecurringPriceCurrency = "YER"
	ProductNewParamsPriceRecurringPriceCurrencyZar ProductNewParamsPriceRecurringPriceCurrency = "ZAR"
	ProductNewParamsPriceRecurringPriceCurrencyZmw ProductNewParamsPriceRecurringPriceCurrency = "ZMW"
)

func (r ProductNewParamsPriceRecurringPriceCurrency) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceRecurringPriceCurrencyAed, ProductNewParamsPriceRecurringPriceCurrencyAll, ProductNewParamsPriceRecurringPriceCurrencyAmd, ProductNewParamsPriceRecurringPriceCurrencyAng, ProductNewParamsPriceRecurringPriceCurrencyAoa, ProductNewParamsPriceRecurringPriceCurrencyArs, ProductNewParamsPriceRecurringPriceCurrencyAud, ProductNewParamsPriceRecurringPriceCurrencyAwg, ProductNewParamsPriceRecurringPriceCurrencyAzn, ProductNewParamsPriceRecurringPriceCurrencyBam, ProductNewParamsPriceRecurringPriceCurrencyBbd, ProductNewParamsPriceRecurringPriceCurrencyBdt, ProductNewParamsPriceRecurringPriceCurrencyBgn, ProductNewParamsPriceRecurringPriceCurrencyBhd, ProductNewParamsPriceRecurringPriceCurrencyBif, ProductNewParamsPriceRecurringPriceCurrencyBmd, ProductNewParamsPriceRecurringPriceCurrencyBnd, ProductNewParamsPriceRecurringPriceCurrencyBob, ProductNewParamsPriceRecurringPriceCurrencyBrl, ProductNewParamsPriceRecurringPriceCurrencyBsd, ProductNewParamsPriceRecurringPriceCurrencyBwp, ProductNewParamsPriceRecurringPriceCurrencyByn, ProductNewParamsPriceRecurringPriceCurrencyBzd, ProductNewParamsPriceRecurringPriceCurrencyCad, ProductNewParamsPriceRecurringPriceCurrencyChf, ProductNewParamsPriceRecurringPriceCurrencyClp, ProductNewParamsPriceRecurringPriceCurrencyCny, ProductNewParamsPriceRecurringPriceCurrencyCop, ProductNewParamsPriceRecurringPriceCurrencyCrc, ProductNewParamsPriceRecurringPriceCurrencyCup, ProductNewParamsPriceRecurringPriceCurrencyCve, ProductNewParamsPriceRecurringPriceCurrencyCzk, ProductNewParamsPriceRecurringPriceCurrencyDjf, ProductNewParamsPriceRecurringPriceCurrencyDkk, ProductNewParamsPriceRecurringPriceCurrencyDop, ProductNewParamsPriceRecurringPriceCurrencyDzd, ProductNewParamsPriceRecurringPriceCurrencyEgp, ProductNewParamsPriceRecurringPriceCurrencyEtb, ProductNewParamsPriceRecurringPriceCurrencyEur, ProductNewParamsPriceRecurringPriceCurrencyFjd, ProductNewParamsPriceRecurringPriceCurrencyFkp, ProductNewParamsPriceRecurringPriceCurrencyGbp, ProductNewParamsPriceRecurringPriceCurrencyGel, ProductNewParamsPriceRecurringPriceCurrencyGhs, ProductNewParamsPriceRecurringPriceCurrencyGip, ProductNewParamsPriceRecurringPriceCurrencyGmd, ProductNewParamsPriceRecurringPriceCurrencyGnf, ProductNewParamsPriceRecurringPriceCurrencyGtq, ProductNewParamsPriceRecurringPriceCurrencyGyd, ProductNewParamsPriceRecurringPriceCurrencyHkd, ProductNewParamsPriceRecurringPriceCurrencyHnl, ProductNewParamsPriceRecurringPriceCurrencyHrk, ProductNewParamsPriceRecurringPriceCurrencyHtg, ProductNewParamsPriceRecurringPriceCurrencyHuf, ProductNewParamsPriceRecurringPriceCurrencyIdr, ProductNewParamsPriceRecurringPriceCurrencyIls, ProductNewParamsPriceRecurringPriceCurrencyInr, ProductNewParamsPriceRecurringPriceCurrencyIqd, ProductNewParamsPriceRecurringPriceCurrencyJmd, ProductNewParamsPriceRecurringPriceCurrencyJod, ProductNewParamsPriceRecurringPriceCurrencyJpy, ProductNewParamsPriceRecurringPriceCurrencyKes, ProductNewParamsPriceRecurringPriceCurrencyKgs, ProductNewParamsPriceRecurringPriceCurrencyKhr, ProductNewParamsPriceRecurringPriceCurrencyKmf, ProductNewParamsPriceRecurringPriceCurrencyKrw, ProductNewParamsPriceRecurringPriceCurrencyKwd, ProductNewParamsPriceRecurringPriceCurrencyKyd, ProductNewParamsPriceRecurringPriceCurrencyKzt, ProductNewParamsPriceRecurringPriceCurrencyLak, ProductNewParamsPriceRecurringPriceCurrencyLbp, ProductNewParamsPriceRecurringPriceCurrencyLkr, ProductNewParamsPriceRecurringPriceCurrencyLrd, ProductNewParamsPriceRecurringPriceCurrencyLsl, ProductNewParamsPriceRecurringPriceCurrencyLyd, ProductNewParamsPriceRecurringPriceCurrencyMad, ProductNewParamsPriceRecurringPriceCurrencyMdl, ProductNewParamsPriceRecurringPriceCurrencyMga, ProductNewParamsPriceRecurringPriceCurrencyMkd, ProductNewParamsPriceRecurringPriceCurrencyMmk, ProductNewParamsPriceRecurringPriceCurrencyMnt, ProductNewParamsPriceRecurringPriceCurrencyMop, ProductNewParamsPriceRecurringPriceCurrencyMru, ProductNewParamsPriceRecurringPriceCurrencyMur, ProductNewParamsPriceRecurringPriceCurrencyMvr, ProductNewParamsPriceRecurringPriceCurrencyMwk, ProductNewParamsPriceRecurringPriceCurrencyMxn, ProductNewParamsPriceRecurringPriceCurrencyMyr, ProductNewParamsPriceRecurringPriceCurrencyMzn, ProductNewParamsPriceRecurringPriceCurrencyNad, ProductNewParamsPriceRecurringPriceCurrencyNgn, ProductNewParamsPriceRecurringPriceCurrencyNio, ProductNewParamsPriceRecurringPriceCurrencyNok, ProductNewParamsPriceRecurringPriceCurrencyNpr, ProductNewParamsPriceRecurringPriceCurrencyNzd, ProductNewParamsPriceRecurringPriceCurrencyOmr, ProductNewParamsPriceRecurringPriceCurrencyPab, ProductNewParamsPriceRecurringPriceCurrencyPen, ProductNewParamsPriceRecurringPriceCurrencyPgk, ProductNewParamsPriceRecurringPriceCurrencyPhp, ProductNewParamsPriceRecurringPriceCurrencyPkr, ProductNewParamsPriceRecurringPriceCurrencyPln, ProductNewParamsPriceRecurringPriceCurrencyPyg, ProductNewParamsPriceRecurringPriceCurrencyQar, ProductNewParamsPriceRecurringPriceCurrencyRon, ProductNewParamsPriceRecurringPriceCurrencyRsd, ProductNewParamsPriceRecurringPriceCurrencyRub, ProductNewParamsPriceRecurringPriceCurrencyRwf, ProductNewParamsPriceRecurringPriceCurrencySar, ProductNewParamsPriceRecurringPriceCurrencySbd, ProductNewParamsPriceRecurringPriceCurrencyScr, ProductNewParamsPriceRecurringPriceCurrencySek, ProductNewParamsPriceRecurringPriceCurrencySgd, ProductNewParamsPriceRecurringPriceCurrencyShp, ProductNewParamsPriceRecurringPriceCurrencySle, ProductNewParamsPriceRecurringPriceCurrencySll, ProductNewParamsPriceRecurringPriceCurrencySos, ProductNewParamsPriceRecurringPriceCurrencySrd, ProductNewParamsPriceRecurringPriceCurrencySsp, ProductNewParamsPriceRecurringPriceCurrencyStn, ProductNewParamsPriceRecurringPriceCurrencySvc, ProductNewParamsPriceRecurringPriceCurrencySzl, ProductNewParamsPriceRecurringPriceCurrencyThb, ProductNewParamsPriceRecurringPriceCurrencyTnd, ProductNewParamsPriceRecurringPriceCurrencyTop, ProductNewParamsPriceRecurringPriceCurrencyTry, ProductNewParamsPriceRecurringPriceCurrencyTtd, ProductNewParamsPriceRecurringPriceCurrencyTwd, ProductNewParamsPriceRecurringPriceCurrencyTzs, ProductNewParamsPriceRecurringPriceCurrencyUah, ProductNewParamsPriceRecurringPriceCurrencyUgx, ProductNewParamsPriceRecurringPriceCurrencyUsd, ProductNewParamsPriceRecurringPriceCurrencyUyu, ProductNewParamsPriceRecurringPriceCurrencyUzs, ProductNewParamsPriceRecurringPriceCurrencyVes, ProductNewParamsPriceRecurringPriceCurrencyVnd, ProductNewParamsPriceRecurringPriceCurrencyVuv, ProductNewParamsPriceRecurringPriceCurrencyWst, ProductNewParamsPriceRecurringPriceCurrencyXaf, ProductNewParamsPriceRecurringPriceCurrencyXcd, ProductNewParamsPriceRecurringPriceCurrencyXof, ProductNewParamsPriceRecurringPriceCurrencyXpf, ProductNewParamsPriceRecurringPriceCurrencyYer, ProductNewParamsPriceRecurringPriceCurrencyZar, ProductNewParamsPriceRecurringPriceCurrencyZmw:
		return true
	}
	return false
}

type ProductNewParamsPriceRecurringPricePaymentFrequencyInterval string

const (
	ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalDay   ProductNewParamsPriceRecurringPricePaymentFrequencyInterval = "Day"
	ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalWeek  ProductNewParamsPriceRecurringPricePaymentFrequencyInterval = "Week"
	ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalMonth ProductNewParamsPriceRecurringPricePaymentFrequencyInterval = "Month"
	ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalYear  ProductNewParamsPriceRecurringPricePaymentFrequencyInterval = "Year"
)

func (r ProductNewParamsPriceRecurringPricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalDay, ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalWeek, ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalMonth, ProductNewParamsPriceRecurringPricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval string

const (
	ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalDay   ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval = "Day"
	ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalWeek  ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval = "Week"
	ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalMonth ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval = "Month"
	ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalYear  ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval = "Year"
)

func (r ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalDay, ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalWeek, ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalMonth, ProductNewParamsPriceRecurringPriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

type ProductNewParamsPriceRecurringPriceType string

const (
	ProductNewParamsPriceRecurringPriceTypeRecurringPrice ProductNewParamsPriceRecurringPriceType = "recurring_price"
)

func (r ProductNewParamsPriceRecurringPriceType) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceRecurringPriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductNewParamsPriceCurrency string

const (
	ProductNewParamsPriceCurrencyAed ProductNewParamsPriceCurrency = "AED"
	ProductNewParamsPriceCurrencyAll ProductNewParamsPriceCurrency = "ALL"
	ProductNewParamsPriceCurrencyAmd ProductNewParamsPriceCurrency = "AMD"
	ProductNewParamsPriceCurrencyAng ProductNewParamsPriceCurrency = "ANG"
	ProductNewParamsPriceCurrencyAoa ProductNewParamsPriceCurrency = "AOA"
	ProductNewParamsPriceCurrencyArs ProductNewParamsPriceCurrency = "ARS"
	ProductNewParamsPriceCurrencyAud ProductNewParamsPriceCurrency = "AUD"
	ProductNewParamsPriceCurrencyAwg ProductNewParamsPriceCurrency = "AWG"
	ProductNewParamsPriceCurrencyAzn ProductNewParamsPriceCurrency = "AZN"
	ProductNewParamsPriceCurrencyBam ProductNewParamsPriceCurrency = "BAM"
	ProductNewParamsPriceCurrencyBbd ProductNewParamsPriceCurrency = "BBD"
	ProductNewParamsPriceCurrencyBdt ProductNewParamsPriceCurrency = "BDT"
	ProductNewParamsPriceCurrencyBgn ProductNewParamsPriceCurrency = "BGN"
	ProductNewParamsPriceCurrencyBhd ProductNewParamsPriceCurrency = "BHD"
	ProductNewParamsPriceCurrencyBif ProductNewParamsPriceCurrency = "BIF"
	ProductNewParamsPriceCurrencyBmd ProductNewParamsPriceCurrency = "BMD"
	ProductNewParamsPriceCurrencyBnd ProductNewParamsPriceCurrency = "BND"
	ProductNewParamsPriceCurrencyBob ProductNewParamsPriceCurrency = "BOB"
	ProductNewParamsPriceCurrencyBrl ProductNewParamsPriceCurrency = "BRL"
	ProductNewParamsPriceCurrencyBsd ProductNewParamsPriceCurrency = "BSD"
	ProductNewParamsPriceCurrencyBwp ProductNewParamsPriceCurrency = "BWP"
	ProductNewParamsPriceCurrencyByn ProductNewParamsPriceCurrency = "BYN"
	ProductNewParamsPriceCurrencyBzd ProductNewParamsPriceCurrency = "BZD"
	ProductNewParamsPriceCurrencyCad ProductNewParamsPriceCurrency = "CAD"
	ProductNewParamsPriceCurrencyChf ProductNewParamsPriceCurrency = "CHF"
	ProductNewParamsPriceCurrencyClp ProductNewParamsPriceCurrency = "CLP"
	ProductNewParamsPriceCurrencyCny ProductNewParamsPriceCurrency = "CNY"
	ProductNewParamsPriceCurrencyCop ProductNewParamsPriceCurrency = "COP"
	ProductNewParamsPriceCurrencyCrc ProductNewParamsPriceCurrency = "CRC"
	ProductNewParamsPriceCurrencyCup ProductNewParamsPriceCurrency = "CUP"
	ProductNewParamsPriceCurrencyCve ProductNewParamsPriceCurrency = "CVE"
	ProductNewParamsPriceCurrencyCzk ProductNewParamsPriceCurrency = "CZK"
	ProductNewParamsPriceCurrencyDjf ProductNewParamsPriceCurrency = "DJF"
	ProductNewParamsPriceCurrencyDkk ProductNewParamsPriceCurrency = "DKK"
	ProductNewParamsPriceCurrencyDop ProductNewParamsPriceCurrency = "DOP"
	ProductNewParamsPriceCurrencyDzd ProductNewParamsPriceCurrency = "DZD"
	ProductNewParamsPriceCurrencyEgp ProductNewParamsPriceCurrency = "EGP"
	ProductNewParamsPriceCurrencyEtb ProductNewParamsPriceCurrency = "ETB"
	ProductNewParamsPriceCurrencyEur ProductNewParamsPriceCurrency = "EUR"
	ProductNewParamsPriceCurrencyFjd ProductNewParamsPriceCurrency = "FJD"
	ProductNewParamsPriceCurrencyFkp ProductNewParamsPriceCurrency = "FKP"
	ProductNewParamsPriceCurrencyGbp ProductNewParamsPriceCurrency = "GBP"
	ProductNewParamsPriceCurrencyGel ProductNewParamsPriceCurrency = "GEL"
	ProductNewParamsPriceCurrencyGhs ProductNewParamsPriceCurrency = "GHS"
	ProductNewParamsPriceCurrencyGip ProductNewParamsPriceCurrency = "GIP"
	ProductNewParamsPriceCurrencyGmd ProductNewParamsPriceCurrency = "GMD"
	ProductNewParamsPriceCurrencyGnf ProductNewParamsPriceCurrency = "GNF"
	ProductNewParamsPriceCurrencyGtq ProductNewParamsPriceCurrency = "GTQ"
	ProductNewParamsPriceCurrencyGyd ProductNewParamsPriceCurrency = "GYD"
	ProductNewParamsPriceCurrencyHkd ProductNewParamsPriceCurrency = "HKD"
	ProductNewParamsPriceCurrencyHnl ProductNewParamsPriceCurrency = "HNL"
	ProductNewParamsPriceCurrencyHrk ProductNewParamsPriceCurrency = "HRK"
	ProductNewParamsPriceCurrencyHtg ProductNewParamsPriceCurrency = "HTG"
	ProductNewParamsPriceCurrencyHuf ProductNewParamsPriceCurrency = "HUF"
	ProductNewParamsPriceCurrencyIdr ProductNewParamsPriceCurrency = "IDR"
	ProductNewParamsPriceCurrencyIls ProductNewParamsPriceCurrency = "ILS"
	ProductNewParamsPriceCurrencyInr ProductNewParamsPriceCurrency = "INR"
	ProductNewParamsPriceCurrencyIqd ProductNewParamsPriceCurrency = "IQD"
	ProductNewParamsPriceCurrencyJmd ProductNewParamsPriceCurrency = "JMD"
	ProductNewParamsPriceCurrencyJod ProductNewParamsPriceCurrency = "JOD"
	ProductNewParamsPriceCurrencyJpy ProductNewParamsPriceCurrency = "JPY"
	ProductNewParamsPriceCurrencyKes ProductNewParamsPriceCurrency = "KES"
	ProductNewParamsPriceCurrencyKgs ProductNewParamsPriceCurrency = "KGS"
	ProductNewParamsPriceCurrencyKhr ProductNewParamsPriceCurrency = "KHR"
	ProductNewParamsPriceCurrencyKmf ProductNewParamsPriceCurrency = "KMF"
	ProductNewParamsPriceCurrencyKrw ProductNewParamsPriceCurrency = "KRW"
	ProductNewParamsPriceCurrencyKwd ProductNewParamsPriceCurrency = "KWD"
	ProductNewParamsPriceCurrencyKyd ProductNewParamsPriceCurrency = "KYD"
	ProductNewParamsPriceCurrencyKzt ProductNewParamsPriceCurrency = "KZT"
	ProductNewParamsPriceCurrencyLak ProductNewParamsPriceCurrency = "LAK"
	ProductNewParamsPriceCurrencyLbp ProductNewParamsPriceCurrency = "LBP"
	ProductNewParamsPriceCurrencyLkr ProductNewParamsPriceCurrency = "LKR"
	ProductNewParamsPriceCurrencyLrd ProductNewParamsPriceCurrency = "LRD"
	ProductNewParamsPriceCurrencyLsl ProductNewParamsPriceCurrency = "LSL"
	ProductNewParamsPriceCurrencyLyd ProductNewParamsPriceCurrency = "LYD"
	ProductNewParamsPriceCurrencyMad ProductNewParamsPriceCurrency = "MAD"
	ProductNewParamsPriceCurrencyMdl ProductNewParamsPriceCurrency = "MDL"
	ProductNewParamsPriceCurrencyMga ProductNewParamsPriceCurrency = "MGA"
	ProductNewParamsPriceCurrencyMkd ProductNewParamsPriceCurrency = "MKD"
	ProductNewParamsPriceCurrencyMmk ProductNewParamsPriceCurrency = "MMK"
	ProductNewParamsPriceCurrencyMnt ProductNewParamsPriceCurrency = "MNT"
	ProductNewParamsPriceCurrencyMop ProductNewParamsPriceCurrency = "MOP"
	ProductNewParamsPriceCurrencyMru ProductNewParamsPriceCurrency = "MRU"
	ProductNewParamsPriceCurrencyMur ProductNewParamsPriceCurrency = "MUR"
	ProductNewParamsPriceCurrencyMvr ProductNewParamsPriceCurrency = "MVR"
	ProductNewParamsPriceCurrencyMwk ProductNewParamsPriceCurrency = "MWK"
	ProductNewParamsPriceCurrencyMxn ProductNewParamsPriceCurrency = "MXN"
	ProductNewParamsPriceCurrencyMyr ProductNewParamsPriceCurrency = "MYR"
	ProductNewParamsPriceCurrencyMzn ProductNewParamsPriceCurrency = "MZN"
	ProductNewParamsPriceCurrencyNad ProductNewParamsPriceCurrency = "NAD"
	ProductNewParamsPriceCurrencyNgn ProductNewParamsPriceCurrency = "NGN"
	ProductNewParamsPriceCurrencyNio ProductNewParamsPriceCurrency = "NIO"
	ProductNewParamsPriceCurrencyNok ProductNewParamsPriceCurrency = "NOK"
	ProductNewParamsPriceCurrencyNpr ProductNewParamsPriceCurrency = "NPR"
	ProductNewParamsPriceCurrencyNzd ProductNewParamsPriceCurrency = "NZD"
	ProductNewParamsPriceCurrencyOmr ProductNewParamsPriceCurrency = "OMR"
	ProductNewParamsPriceCurrencyPab ProductNewParamsPriceCurrency = "PAB"
	ProductNewParamsPriceCurrencyPen ProductNewParamsPriceCurrency = "PEN"
	ProductNewParamsPriceCurrencyPgk ProductNewParamsPriceCurrency = "PGK"
	ProductNewParamsPriceCurrencyPhp ProductNewParamsPriceCurrency = "PHP"
	ProductNewParamsPriceCurrencyPkr ProductNewParamsPriceCurrency = "PKR"
	ProductNewParamsPriceCurrencyPln ProductNewParamsPriceCurrency = "PLN"
	ProductNewParamsPriceCurrencyPyg ProductNewParamsPriceCurrency = "PYG"
	ProductNewParamsPriceCurrencyQar ProductNewParamsPriceCurrency = "QAR"
	ProductNewParamsPriceCurrencyRon ProductNewParamsPriceCurrency = "RON"
	ProductNewParamsPriceCurrencyRsd ProductNewParamsPriceCurrency = "RSD"
	ProductNewParamsPriceCurrencyRub ProductNewParamsPriceCurrency = "RUB"
	ProductNewParamsPriceCurrencyRwf ProductNewParamsPriceCurrency = "RWF"
	ProductNewParamsPriceCurrencySar ProductNewParamsPriceCurrency = "SAR"
	ProductNewParamsPriceCurrencySbd ProductNewParamsPriceCurrency = "SBD"
	ProductNewParamsPriceCurrencyScr ProductNewParamsPriceCurrency = "SCR"
	ProductNewParamsPriceCurrencySek ProductNewParamsPriceCurrency = "SEK"
	ProductNewParamsPriceCurrencySgd ProductNewParamsPriceCurrency = "SGD"
	ProductNewParamsPriceCurrencyShp ProductNewParamsPriceCurrency = "SHP"
	ProductNewParamsPriceCurrencySle ProductNewParamsPriceCurrency = "SLE"
	ProductNewParamsPriceCurrencySll ProductNewParamsPriceCurrency = "SLL"
	ProductNewParamsPriceCurrencySos ProductNewParamsPriceCurrency = "SOS"
	ProductNewParamsPriceCurrencySrd ProductNewParamsPriceCurrency = "SRD"
	ProductNewParamsPriceCurrencySsp ProductNewParamsPriceCurrency = "SSP"
	ProductNewParamsPriceCurrencyStn ProductNewParamsPriceCurrency = "STN"
	ProductNewParamsPriceCurrencySvc ProductNewParamsPriceCurrency = "SVC"
	ProductNewParamsPriceCurrencySzl ProductNewParamsPriceCurrency = "SZL"
	ProductNewParamsPriceCurrencyThb ProductNewParamsPriceCurrency = "THB"
	ProductNewParamsPriceCurrencyTnd ProductNewParamsPriceCurrency = "TND"
	ProductNewParamsPriceCurrencyTop ProductNewParamsPriceCurrency = "TOP"
	ProductNewParamsPriceCurrencyTry ProductNewParamsPriceCurrency = "TRY"
	ProductNewParamsPriceCurrencyTtd ProductNewParamsPriceCurrency = "TTD"
	ProductNewParamsPriceCurrencyTwd ProductNewParamsPriceCurrency = "TWD"
	ProductNewParamsPriceCurrencyTzs ProductNewParamsPriceCurrency = "TZS"
	ProductNewParamsPriceCurrencyUah ProductNewParamsPriceCurrency = "UAH"
	ProductNewParamsPriceCurrencyUgx ProductNewParamsPriceCurrency = "UGX"
	ProductNewParamsPriceCurrencyUsd ProductNewParamsPriceCurrency = "USD"
	ProductNewParamsPriceCurrencyUyu ProductNewParamsPriceCurrency = "UYU"
	ProductNewParamsPriceCurrencyUzs ProductNewParamsPriceCurrency = "UZS"
	ProductNewParamsPriceCurrencyVes ProductNewParamsPriceCurrency = "VES"
	ProductNewParamsPriceCurrencyVnd ProductNewParamsPriceCurrency = "VND"
	ProductNewParamsPriceCurrencyVuv ProductNewParamsPriceCurrency = "VUV"
	ProductNewParamsPriceCurrencyWst ProductNewParamsPriceCurrency = "WST"
	ProductNewParamsPriceCurrencyXaf ProductNewParamsPriceCurrency = "XAF"
	ProductNewParamsPriceCurrencyXcd ProductNewParamsPriceCurrency = "XCD"
	ProductNewParamsPriceCurrencyXof ProductNewParamsPriceCurrency = "XOF"
	ProductNewParamsPriceCurrencyXpf ProductNewParamsPriceCurrency = "XPF"
	ProductNewParamsPriceCurrencyYer ProductNewParamsPriceCurrency = "YER"
	ProductNewParamsPriceCurrencyZar ProductNewParamsPriceCurrency = "ZAR"
	ProductNewParamsPriceCurrencyZmw ProductNewParamsPriceCurrency = "ZMW"
)

func (r ProductNewParamsPriceCurrency) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceCurrencyAed, ProductNewParamsPriceCurrencyAll, ProductNewParamsPriceCurrencyAmd, ProductNewParamsPriceCurrencyAng, ProductNewParamsPriceCurrencyAoa, ProductNewParamsPriceCurrencyArs, ProductNewParamsPriceCurrencyAud, ProductNewParamsPriceCurrencyAwg, ProductNewParamsPriceCurrencyAzn, ProductNewParamsPriceCurrencyBam, ProductNewParamsPriceCurrencyBbd, ProductNewParamsPriceCurrencyBdt, ProductNewParamsPriceCurrencyBgn, ProductNewParamsPriceCurrencyBhd, ProductNewParamsPriceCurrencyBif, ProductNewParamsPriceCurrencyBmd, ProductNewParamsPriceCurrencyBnd, ProductNewParamsPriceCurrencyBob, ProductNewParamsPriceCurrencyBrl, ProductNewParamsPriceCurrencyBsd, ProductNewParamsPriceCurrencyBwp, ProductNewParamsPriceCurrencyByn, ProductNewParamsPriceCurrencyBzd, ProductNewParamsPriceCurrencyCad, ProductNewParamsPriceCurrencyChf, ProductNewParamsPriceCurrencyClp, ProductNewParamsPriceCurrencyCny, ProductNewParamsPriceCurrencyCop, ProductNewParamsPriceCurrencyCrc, ProductNewParamsPriceCurrencyCup, ProductNewParamsPriceCurrencyCve, ProductNewParamsPriceCurrencyCzk, ProductNewParamsPriceCurrencyDjf, ProductNewParamsPriceCurrencyDkk, ProductNewParamsPriceCurrencyDop, ProductNewParamsPriceCurrencyDzd, ProductNewParamsPriceCurrencyEgp, ProductNewParamsPriceCurrencyEtb, ProductNewParamsPriceCurrencyEur, ProductNewParamsPriceCurrencyFjd, ProductNewParamsPriceCurrencyFkp, ProductNewParamsPriceCurrencyGbp, ProductNewParamsPriceCurrencyGel, ProductNewParamsPriceCurrencyGhs, ProductNewParamsPriceCurrencyGip, ProductNewParamsPriceCurrencyGmd, ProductNewParamsPriceCurrencyGnf, ProductNewParamsPriceCurrencyGtq, ProductNewParamsPriceCurrencyGyd, ProductNewParamsPriceCurrencyHkd, ProductNewParamsPriceCurrencyHnl, ProductNewParamsPriceCurrencyHrk, ProductNewParamsPriceCurrencyHtg, ProductNewParamsPriceCurrencyHuf, ProductNewParamsPriceCurrencyIdr, ProductNewParamsPriceCurrencyIls, ProductNewParamsPriceCurrencyInr, ProductNewParamsPriceCurrencyIqd, ProductNewParamsPriceCurrencyJmd, ProductNewParamsPriceCurrencyJod, ProductNewParamsPriceCurrencyJpy, ProductNewParamsPriceCurrencyKes, ProductNewParamsPriceCurrencyKgs, ProductNewParamsPriceCurrencyKhr, ProductNewParamsPriceCurrencyKmf, ProductNewParamsPriceCurrencyKrw, ProductNewParamsPriceCurrencyKwd, ProductNewParamsPriceCurrencyKyd, ProductNewParamsPriceCurrencyKzt, ProductNewParamsPriceCurrencyLak, ProductNewParamsPriceCurrencyLbp, ProductNewParamsPriceCurrencyLkr, ProductNewParamsPriceCurrencyLrd, ProductNewParamsPriceCurrencyLsl, ProductNewParamsPriceCurrencyLyd, ProductNewParamsPriceCurrencyMad, ProductNewParamsPriceCurrencyMdl, ProductNewParamsPriceCurrencyMga, ProductNewParamsPriceCurrencyMkd, ProductNewParamsPriceCurrencyMmk, ProductNewParamsPriceCurrencyMnt, ProductNewParamsPriceCurrencyMop, ProductNewParamsPriceCurrencyMru, ProductNewParamsPriceCurrencyMur, ProductNewParamsPriceCurrencyMvr, ProductNewParamsPriceCurrencyMwk, ProductNewParamsPriceCurrencyMxn, ProductNewParamsPriceCurrencyMyr, ProductNewParamsPriceCurrencyMzn, ProductNewParamsPriceCurrencyNad, ProductNewParamsPriceCurrencyNgn, ProductNewParamsPriceCurrencyNio, ProductNewParamsPriceCurrencyNok, ProductNewParamsPriceCurrencyNpr, ProductNewParamsPriceCurrencyNzd, ProductNewParamsPriceCurrencyOmr, ProductNewParamsPriceCurrencyPab, ProductNewParamsPriceCurrencyPen, ProductNewParamsPriceCurrencyPgk, ProductNewParamsPriceCurrencyPhp, ProductNewParamsPriceCurrencyPkr, ProductNewParamsPriceCurrencyPln, ProductNewParamsPriceCurrencyPyg, ProductNewParamsPriceCurrencyQar, ProductNewParamsPriceCurrencyRon, ProductNewParamsPriceCurrencyRsd, ProductNewParamsPriceCurrencyRub, ProductNewParamsPriceCurrencyRwf, ProductNewParamsPriceCurrencySar, ProductNewParamsPriceCurrencySbd, ProductNewParamsPriceCurrencyScr, ProductNewParamsPriceCurrencySek, ProductNewParamsPriceCurrencySgd, ProductNewParamsPriceCurrencyShp, ProductNewParamsPriceCurrencySle, ProductNewParamsPriceCurrencySll, ProductNewParamsPriceCurrencySos, ProductNewParamsPriceCurrencySrd, ProductNewParamsPriceCurrencySsp, ProductNewParamsPriceCurrencyStn, ProductNewParamsPriceCurrencySvc, ProductNewParamsPriceCurrencySzl, ProductNewParamsPriceCurrencyThb, ProductNewParamsPriceCurrencyTnd, ProductNewParamsPriceCurrencyTop, ProductNewParamsPriceCurrencyTry, ProductNewParamsPriceCurrencyTtd, ProductNewParamsPriceCurrencyTwd, ProductNewParamsPriceCurrencyTzs, ProductNewParamsPriceCurrencyUah, ProductNewParamsPriceCurrencyUgx, ProductNewParamsPriceCurrencyUsd, ProductNewParamsPriceCurrencyUyu, ProductNewParamsPriceCurrencyUzs, ProductNewParamsPriceCurrencyVes, ProductNewParamsPriceCurrencyVnd, ProductNewParamsPriceCurrencyVuv, ProductNewParamsPriceCurrencyWst, ProductNewParamsPriceCurrencyXaf, ProductNewParamsPriceCurrencyXcd, ProductNewParamsPriceCurrencyXof, ProductNewParamsPriceCurrencyXpf, ProductNewParamsPriceCurrencyYer, ProductNewParamsPriceCurrencyZar, ProductNewParamsPriceCurrencyZmw:
		return true
	}
	return false
}

type ProductNewParamsPriceType string

const (
	ProductNewParamsPriceTypeOneTimePrice   ProductNewParamsPriceType = "one_time_price"
	ProductNewParamsPriceTypeRecurringPrice ProductNewParamsPriceType = "recurring_price"
)

func (r ProductNewParamsPriceType) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceTypeOneTimePrice, ProductNewParamsPriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductNewParamsPricePaymentFrequencyInterval string

const (
	ProductNewParamsPricePaymentFrequencyIntervalDay   ProductNewParamsPricePaymentFrequencyInterval = "Day"
	ProductNewParamsPricePaymentFrequencyIntervalWeek  ProductNewParamsPricePaymentFrequencyInterval = "Week"
	ProductNewParamsPricePaymentFrequencyIntervalMonth ProductNewParamsPricePaymentFrequencyInterval = "Month"
	ProductNewParamsPricePaymentFrequencyIntervalYear  ProductNewParamsPricePaymentFrequencyInterval = "Year"
)

func (r ProductNewParamsPricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductNewParamsPricePaymentFrequencyIntervalDay, ProductNewParamsPricePaymentFrequencyIntervalWeek, ProductNewParamsPricePaymentFrequencyIntervalMonth, ProductNewParamsPricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductNewParamsPriceSubscriptionPeriodInterval string

const (
	ProductNewParamsPriceSubscriptionPeriodIntervalDay   ProductNewParamsPriceSubscriptionPeriodInterval = "Day"
	ProductNewParamsPriceSubscriptionPeriodIntervalWeek  ProductNewParamsPriceSubscriptionPeriodInterval = "Week"
	ProductNewParamsPriceSubscriptionPeriodIntervalMonth ProductNewParamsPriceSubscriptionPeriodInterval = "Month"
	ProductNewParamsPriceSubscriptionPeriodIntervalYear  ProductNewParamsPriceSubscriptionPeriodInterval = "Year"
)

func (r ProductNewParamsPriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductNewParamsPriceSubscriptionPeriodIntervalDay, ProductNewParamsPriceSubscriptionPeriodIntervalWeek, ProductNewParamsPriceSubscriptionPeriodIntervalMonth, ProductNewParamsPriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

// Represents the different categories of taxation applicable to various products
// and services.
type ProductNewParamsTaxCategory string

const (
	ProductNewParamsTaxCategoryDigitalProducts ProductNewParamsTaxCategory = "digital_products"
	ProductNewParamsTaxCategorySaas            ProductNewParamsTaxCategory = "saas"
	ProductNewParamsTaxCategoryEBook           ProductNewParamsTaxCategory = "e_book"
)

func (r ProductNewParamsTaxCategory) IsKnown() bool {
	switch r {
	case ProductNewParamsTaxCategoryDigitalProducts, ProductNewParamsTaxCategorySaas, ProductNewParamsTaxCategoryEBook:
		return true
	}
	return false
}

type ProductNewParamsLicenseKeyDuration struct {
	Count    param.Field[int64]                                      `json:"count,required"`
	Interval param.Field[ProductNewParamsLicenseKeyDurationInterval] `json:"interval,required"`
}

func (r ProductNewParamsLicenseKeyDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductNewParamsLicenseKeyDurationInterval string

const (
	ProductNewParamsLicenseKeyDurationIntervalDay   ProductNewParamsLicenseKeyDurationInterval = "Day"
	ProductNewParamsLicenseKeyDurationIntervalWeek  ProductNewParamsLicenseKeyDurationInterval = "Week"
	ProductNewParamsLicenseKeyDurationIntervalMonth ProductNewParamsLicenseKeyDurationInterval = "Month"
	ProductNewParamsLicenseKeyDurationIntervalYear  ProductNewParamsLicenseKeyDurationInterval = "Year"
)

func (r ProductNewParamsLicenseKeyDurationInterval) IsKnown() bool {
	switch r {
	case ProductNewParamsLicenseKeyDurationIntervalDay, ProductNewParamsLicenseKeyDurationIntervalWeek, ProductNewParamsLicenseKeyDurationIntervalMonth, ProductNewParamsLicenseKeyDurationIntervalYear:
		return true
	}
	return false
}

type ProductUpdateParams struct {
	// Description of the product, optional and must be at most 1000 characters.
	Description param.Field[string] `json:"description"`
	// Message sent to the customer upon license key activation.
	//
	// Only applicable if `license_key_enabled` is `true`. This message contains
	// instructions for activating the license key.
	LicenseKeyActivationMessage param.Field[string] `json:"license_key_activation_message"`
	// Limit for the number of activations for the license key.
	//
	// Only applicable if `license_key_enabled` is `true`. Represents the maximum
	// number of times the license key can be activated.
	LicenseKeyActivationsLimit param.Field[int64]                                 `json:"license_key_activations_limit"`
	LicenseKeyDuration         param.Field[ProductUpdateParamsLicenseKeyDuration] `json:"license_key_duration"`
	// Whether the product requires a license key.
	//
	// If `true`, additional fields related to license key (duration, activations
	// limit, activation message) become applicable.
	LicenseKeyEnabled param.Field[bool] `json:"license_key_enabled"`
	// Name of the product, optional and must be at most 100 characters.
	Name  param.Field[string]                        `json:"name"`
	Price param.Field[ProductUpdateParamsPriceUnion] `json:"price"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[ProductUpdateParamsTaxCategory] `json:"tax_category"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductUpdateParamsLicenseKeyDuration struct {
	Count    param.Field[int64]                                         `json:"count,required"`
	Interval param.Field[ProductUpdateParamsLicenseKeyDurationInterval] `json:"interval,required"`
}

func (r ProductUpdateParamsLicenseKeyDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductUpdateParamsLicenseKeyDurationInterval string

const (
	ProductUpdateParamsLicenseKeyDurationIntervalDay   ProductUpdateParamsLicenseKeyDurationInterval = "Day"
	ProductUpdateParamsLicenseKeyDurationIntervalWeek  ProductUpdateParamsLicenseKeyDurationInterval = "Week"
	ProductUpdateParamsLicenseKeyDurationIntervalMonth ProductUpdateParamsLicenseKeyDurationInterval = "Month"
	ProductUpdateParamsLicenseKeyDurationIntervalYear  ProductUpdateParamsLicenseKeyDurationInterval = "Year"
)

func (r ProductUpdateParamsLicenseKeyDurationInterval) IsKnown() bool {
	switch r {
	case ProductUpdateParamsLicenseKeyDurationIntervalDay, ProductUpdateParamsLicenseKeyDurationIntervalWeek, ProductUpdateParamsLicenseKeyDurationIntervalMonth, ProductUpdateParamsLicenseKeyDurationIntervalYear:
		return true
	}
	return false
}

type ProductUpdateParamsPrice struct {
	Currency param.Field[ProductUpdateParamsPriceCurrency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool]                         `json:"purchasing_power_parity,required"`
	Type                  param.Field[ProductUpdateParamsPriceType] `json:"type,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    param.Field[int64]                                            `json:"payment_frequency_count"`
	PaymentFrequencyInterval param.Field[ProductUpdateParamsPricePaymentFrequencyInterval] `json:"payment_frequency_interval"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    param.Field[int64]                                              `json:"subscription_period_count"`
	SubscriptionPeriodInterval param.Field[ProductUpdateParamsPriceSubscriptionPeriodInterval] `json:"subscription_period_interval"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r ProductUpdateParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductUpdateParamsPrice) implementsProductUpdateParamsPriceUnion() {}

// Satisfied by [ProductUpdateParamsPriceOneTimePrice],
// [ProductUpdateParamsPriceRecurringPrice], [ProductUpdateParamsPrice].
type ProductUpdateParamsPriceUnion interface {
	implementsProductUpdateParamsPriceUnion()
}

type ProductUpdateParamsPriceOneTimePrice struct {
	Currency param.Field[ProductUpdateParamsPriceOneTimePriceCurrency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool]                                     `json:"purchasing_power_parity,required"`
	Type                  param.Field[ProductUpdateParamsPriceOneTimePriceType] `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
}

func (r ProductUpdateParamsPriceOneTimePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductUpdateParamsPriceOneTimePrice) implementsProductUpdateParamsPriceUnion() {}

type ProductUpdateParamsPriceOneTimePriceCurrency string

const (
	ProductUpdateParamsPriceOneTimePriceCurrencyAed ProductUpdateParamsPriceOneTimePriceCurrency = "AED"
	ProductUpdateParamsPriceOneTimePriceCurrencyAll ProductUpdateParamsPriceOneTimePriceCurrency = "ALL"
	ProductUpdateParamsPriceOneTimePriceCurrencyAmd ProductUpdateParamsPriceOneTimePriceCurrency = "AMD"
	ProductUpdateParamsPriceOneTimePriceCurrencyAng ProductUpdateParamsPriceOneTimePriceCurrency = "ANG"
	ProductUpdateParamsPriceOneTimePriceCurrencyAoa ProductUpdateParamsPriceOneTimePriceCurrency = "AOA"
	ProductUpdateParamsPriceOneTimePriceCurrencyArs ProductUpdateParamsPriceOneTimePriceCurrency = "ARS"
	ProductUpdateParamsPriceOneTimePriceCurrencyAud ProductUpdateParamsPriceOneTimePriceCurrency = "AUD"
	ProductUpdateParamsPriceOneTimePriceCurrencyAwg ProductUpdateParamsPriceOneTimePriceCurrency = "AWG"
	ProductUpdateParamsPriceOneTimePriceCurrencyAzn ProductUpdateParamsPriceOneTimePriceCurrency = "AZN"
	ProductUpdateParamsPriceOneTimePriceCurrencyBam ProductUpdateParamsPriceOneTimePriceCurrency = "BAM"
	ProductUpdateParamsPriceOneTimePriceCurrencyBbd ProductUpdateParamsPriceOneTimePriceCurrency = "BBD"
	ProductUpdateParamsPriceOneTimePriceCurrencyBdt ProductUpdateParamsPriceOneTimePriceCurrency = "BDT"
	ProductUpdateParamsPriceOneTimePriceCurrencyBgn ProductUpdateParamsPriceOneTimePriceCurrency = "BGN"
	ProductUpdateParamsPriceOneTimePriceCurrencyBhd ProductUpdateParamsPriceOneTimePriceCurrency = "BHD"
	ProductUpdateParamsPriceOneTimePriceCurrencyBif ProductUpdateParamsPriceOneTimePriceCurrency = "BIF"
	ProductUpdateParamsPriceOneTimePriceCurrencyBmd ProductUpdateParamsPriceOneTimePriceCurrency = "BMD"
	ProductUpdateParamsPriceOneTimePriceCurrencyBnd ProductUpdateParamsPriceOneTimePriceCurrency = "BND"
	ProductUpdateParamsPriceOneTimePriceCurrencyBob ProductUpdateParamsPriceOneTimePriceCurrency = "BOB"
	ProductUpdateParamsPriceOneTimePriceCurrencyBrl ProductUpdateParamsPriceOneTimePriceCurrency = "BRL"
	ProductUpdateParamsPriceOneTimePriceCurrencyBsd ProductUpdateParamsPriceOneTimePriceCurrency = "BSD"
	ProductUpdateParamsPriceOneTimePriceCurrencyBwp ProductUpdateParamsPriceOneTimePriceCurrency = "BWP"
	ProductUpdateParamsPriceOneTimePriceCurrencyByn ProductUpdateParamsPriceOneTimePriceCurrency = "BYN"
	ProductUpdateParamsPriceOneTimePriceCurrencyBzd ProductUpdateParamsPriceOneTimePriceCurrency = "BZD"
	ProductUpdateParamsPriceOneTimePriceCurrencyCad ProductUpdateParamsPriceOneTimePriceCurrency = "CAD"
	ProductUpdateParamsPriceOneTimePriceCurrencyChf ProductUpdateParamsPriceOneTimePriceCurrency = "CHF"
	ProductUpdateParamsPriceOneTimePriceCurrencyClp ProductUpdateParamsPriceOneTimePriceCurrency = "CLP"
	ProductUpdateParamsPriceOneTimePriceCurrencyCny ProductUpdateParamsPriceOneTimePriceCurrency = "CNY"
	ProductUpdateParamsPriceOneTimePriceCurrencyCop ProductUpdateParamsPriceOneTimePriceCurrency = "COP"
	ProductUpdateParamsPriceOneTimePriceCurrencyCrc ProductUpdateParamsPriceOneTimePriceCurrency = "CRC"
	ProductUpdateParamsPriceOneTimePriceCurrencyCup ProductUpdateParamsPriceOneTimePriceCurrency = "CUP"
	ProductUpdateParamsPriceOneTimePriceCurrencyCve ProductUpdateParamsPriceOneTimePriceCurrency = "CVE"
	ProductUpdateParamsPriceOneTimePriceCurrencyCzk ProductUpdateParamsPriceOneTimePriceCurrency = "CZK"
	ProductUpdateParamsPriceOneTimePriceCurrencyDjf ProductUpdateParamsPriceOneTimePriceCurrency = "DJF"
	ProductUpdateParamsPriceOneTimePriceCurrencyDkk ProductUpdateParamsPriceOneTimePriceCurrency = "DKK"
	ProductUpdateParamsPriceOneTimePriceCurrencyDop ProductUpdateParamsPriceOneTimePriceCurrency = "DOP"
	ProductUpdateParamsPriceOneTimePriceCurrencyDzd ProductUpdateParamsPriceOneTimePriceCurrency = "DZD"
	ProductUpdateParamsPriceOneTimePriceCurrencyEgp ProductUpdateParamsPriceOneTimePriceCurrency = "EGP"
	ProductUpdateParamsPriceOneTimePriceCurrencyEtb ProductUpdateParamsPriceOneTimePriceCurrency = "ETB"
	ProductUpdateParamsPriceOneTimePriceCurrencyEur ProductUpdateParamsPriceOneTimePriceCurrency = "EUR"
	ProductUpdateParamsPriceOneTimePriceCurrencyFjd ProductUpdateParamsPriceOneTimePriceCurrency = "FJD"
	ProductUpdateParamsPriceOneTimePriceCurrencyFkp ProductUpdateParamsPriceOneTimePriceCurrency = "FKP"
	ProductUpdateParamsPriceOneTimePriceCurrencyGbp ProductUpdateParamsPriceOneTimePriceCurrency = "GBP"
	ProductUpdateParamsPriceOneTimePriceCurrencyGel ProductUpdateParamsPriceOneTimePriceCurrency = "GEL"
	ProductUpdateParamsPriceOneTimePriceCurrencyGhs ProductUpdateParamsPriceOneTimePriceCurrency = "GHS"
	ProductUpdateParamsPriceOneTimePriceCurrencyGip ProductUpdateParamsPriceOneTimePriceCurrency = "GIP"
	ProductUpdateParamsPriceOneTimePriceCurrencyGmd ProductUpdateParamsPriceOneTimePriceCurrency = "GMD"
	ProductUpdateParamsPriceOneTimePriceCurrencyGnf ProductUpdateParamsPriceOneTimePriceCurrency = "GNF"
	ProductUpdateParamsPriceOneTimePriceCurrencyGtq ProductUpdateParamsPriceOneTimePriceCurrency = "GTQ"
	ProductUpdateParamsPriceOneTimePriceCurrencyGyd ProductUpdateParamsPriceOneTimePriceCurrency = "GYD"
	ProductUpdateParamsPriceOneTimePriceCurrencyHkd ProductUpdateParamsPriceOneTimePriceCurrency = "HKD"
	ProductUpdateParamsPriceOneTimePriceCurrencyHnl ProductUpdateParamsPriceOneTimePriceCurrency = "HNL"
	ProductUpdateParamsPriceOneTimePriceCurrencyHrk ProductUpdateParamsPriceOneTimePriceCurrency = "HRK"
	ProductUpdateParamsPriceOneTimePriceCurrencyHtg ProductUpdateParamsPriceOneTimePriceCurrency = "HTG"
	ProductUpdateParamsPriceOneTimePriceCurrencyHuf ProductUpdateParamsPriceOneTimePriceCurrency = "HUF"
	ProductUpdateParamsPriceOneTimePriceCurrencyIdr ProductUpdateParamsPriceOneTimePriceCurrency = "IDR"
	ProductUpdateParamsPriceOneTimePriceCurrencyIls ProductUpdateParamsPriceOneTimePriceCurrency = "ILS"
	ProductUpdateParamsPriceOneTimePriceCurrencyInr ProductUpdateParamsPriceOneTimePriceCurrency = "INR"
	ProductUpdateParamsPriceOneTimePriceCurrencyIqd ProductUpdateParamsPriceOneTimePriceCurrency = "IQD"
	ProductUpdateParamsPriceOneTimePriceCurrencyJmd ProductUpdateParamsPriceOneTimePriceCurrency = "JMD"
	ProductUpdateParamsPriceOneTimePriceCurrencyJod ProductUpdateParamsPriceOneTimePriceCurrency = "JOD"
	ProductUpdateParamsPriceOneTimePriceCurrencyJpy ProductUpdateParamsPriceOneTimePriceCurrency = "JPY"
	ProductUpdateParamsPriceOneTimePriceCurrencyKes ProductUpdateParamsPriceOneTimePriceCurrency = "KES"
	ProductUpdateParamsPriceOneTimePriceCurrencyKgs ProductUpdateParamsPriceOneTimePriceCurrency = "KGS"
	ProductUpdateParamsPriceOneTimePriceCurrencyKhr ProductUpdateParamsPriceOneTimePriceCurrency = "KHR"
	ProductUpdateParamsPriceOneTimePriceCurrencyKmf ProductUpdateParamsPriceOneTimePriceCurrency = "KMF"
	ProductUpdateParamsPriceOneTimePriceCurrencyKrw ProductUpdateParamsPriceOneTimePriceCurrency = "KRW"
	ProductUpdateParamsPriceOneTimePriceCurrencyKwd ProductUpdateParamsPriceOneTimePriceCurrency = "KWD"
	ProductUpdateParamsPriceOneTimePriceCurrencyKyd ProductUpdateParamsPriceOneTimePriceCurrency = "KYD"
	ProductUpdateParamsPriceOneTimePriceCurrencyKzt ProductUpdateParamsPriceOneTimePriceCurrency = "KZT"
	ProductUpdateParamsPriceOneTimePriceCurrencyLak ProductUpdateParamsPriceOneTimePriceCurrency = "LAK"
	ProductUpdateParamsPriceOneTimePriceCurrencyLbp ProductUpdateParamsPriceOneTimePriceCurrency = "LBP"
	ProductUpdateParamsPriceOneTimePriceCurrencyLkr ProductUpdateParamsPriceOneTimePriceCurrency = "LKR"
	ProductUpdateParamsPriceOneTimePriceCurrencyLrd ProductUpdateParamsPriceOneTimePriceCurrency = "LRD"
	ProductUpdateParamsPriceOneTimePriceCurrencyLsl ProductUpdateParamsPriceOneTimePriceCurrency = "LSL"
	ProductUpdateParamsPriceOneTimePriceCurrencyLyd ProductUpdateParamsPriceOneTimePriceCurrency = "LYD"
	ProductUpdateParamsPriceOneTimePriceCurrencyMad ProductUpdateParamsPriceOneTimePriceCurrency = "MAD"
	ProductUpdateParamsPriceOneTimePriceCurrencyMdl ProductUpdateParamsPriceOneTimePriceCurrency = "MDL"
	ProductUpdateParamsPriceOneTimePriceCurrencyMga ProductUpdateParamsPriceOneTimePriceCurrency = "MGA"
	ProductUpdateParamsPriceOneTimePriceCurrencyMkd ProductUpdateParamsPriceOneTimePriceCurrency = "MKD"
	ProductUpdateParamsPriceOneTimePriceCurrencyMmk ProductUpdateParamsPriceOneTimePriceCurrency = "MMK"
	ProductUpdateParamsPriceOneTimePriceCurrencyMnt ProductUpdateParamsPriceOneTimePriceCurrency = "MNT"
	ProductUpdateParamsPriceOneTimePriceCurrencyMop ProductUpdateParamsPriceOneTimePriceCurrency = "MOP"
	ProductUpdateParamsPriceOneTimePriceCurrencyMru ProductUpdateParamsPriceOneTimePriceCurrency = "MRU"
	ProductUpdateParamsPriceOneTimePriceCurrencyMur ProductUpdateParamsPriceOneTimePriceCurrency = "MUR"
	ProductUpdateParamsPriceOneTimePriceCurrencyMvr ProductUpdateParamsPriceOneTimePriceCurrency = "MVR"
	ProductUpdateParamsPriceOneTimePriceCurrencyMwk ProductUpdateParamsPriceOneTimePriceCurrency = "MWK"
	ProductUpdateParamsPriceOneTimePriceCurrencyMxn ProductUpdateParamsPriceOneTimePriceCurrency = "MXN"
	ProductUpdateParamsPriceOneTimePriceCurrencyMyr ProductUpdateParamsPriceOneTimePriceCurrency = "MYR"
	ProductUpdateParamsPriceOneTimePriceCurrencyMzn ProductUpdateParamsPriceOneTimePriceCurrency = "MZN"
	ProductUpdateParamsPriceOneTimePriceCurrencyNad ProductUpdateParamsPriceOneTimePriceCurrency = "NAD"
	ProductUpdateParamsPriceOneTimePriceCurrencyNgn ProductUpdateParamsPriceOneTimePriceCurrency = "NGN"
	ProductUpdateParamsPriceOneTimePriceCurrencyNio ProductUpdateParamsPriceOneTimePriceCurrency = "NIO"
	ProductUpdateParamsPriceOneTimePriceCurrencyNok ProductUpdateParamsPriceOneTimePriceCurrency = "NOK"
	ProductUpdateParamsPriceOneTimePriceCurrencyNpr ProductUpdateParamsPriceOneTimePriceCurrency = "NPR"
	ProductUpdateParamsPriceOneTimePriceCurrencyNzd ProductUpdateParamsPriceOneTimePriceCurrency = "NZD"
	ProductUpdateParamsPriceOneTimePriceCurrencyOmr ProductUpdateParamsPriceOneTimePriceCurrency = "OMR"
	ProductUpdateParamsPriceOneTimePriceCurrencyPab ProductUpdateParamsPriceOneTimePriceCurrency = "PAB"
	ProductUpdateParamsPriceOneTimePriceCurrencyPen ProductUpdateParamsPriceOneTimePriceCurrency = "PEN"
	ProductUpdateParamsPriceOneTimePriceCurrencyPgk ProductUpdateParamsPriceOneTimePriceCurrency = "PGK"
	ProductUpdateParamsPriceOneTimePriceCurrencyPhp ProductUpdateParamsPriceOneTimePriceCurrency = "PHP"
	ProductUpdateParamsPriceOneTimePriceCurrencyPkr ProductUpdateParamsPriceOneTimePriceCurrency = "PKR"
	ProductUpdateParamsPriceOneTimePriceCurrencyPln ProductUpdateParamsPriceOneTimePriceCurrency = "PLN"
	ProductUpdateParamsPriceOneTimePriceCurrencyPyg ProductUpdateParamsPriceOneTimePriceCurrency = "PYG"
	ProductUpdateParamsPriceOneTimePriceCurrencyQar ProductUpdateParamsPriceOneTimePriceCurrency = "QAR"
	ProductUpdateParamsPriceOneTimePriceCurrencyRon ProductUpdateParamsPriceOneTimePriceCurrency = "RON"
	ProductUpdateParamsPriceOneTimePriceCurrencyRsd ProductUpdateParamsPriceOneTimePriceCurrency = "RSD"
	ProductUpdateParamsPriceOneTimePriceCurrencyRub ProductUpdateParamsPriceOneTimePriceCurrency = "RUB"
	ProductUpdateParamsPriceOneTimePriceCurrencyRwf ProductUpdateParamsPriceOneTimePriceCurrency = "RWF"
	ProductUpdateParamsPriceOneTimePriceCurrencySar ProductUpdateParamsPriceOneTimePriceCurrency = "SAR"
	ProductUpdateParamsPriceOneTimePriceCurrencySbd ProductUpdateParamsPriceOneTimePriceCurrency = "SBD"
	ProductUpdateParamsPriceOneTimePriceCurrencyScr ProductUpdateParamsPriceOneTimePriceCurrency = "SCR"
	ProductUpdateParamsPriceOneTimePriceCurrencySek ProductUpdateParamsPriceOneTimePriceCurrency = "SEK"
	ProductUpdateParamsPriceOneTimePriceCurrencySgd ProductUpdateParamsPriceOneTimePriceCurrency = "SGD"
	ProductUpdateParamsPriceOneTimePriceCurrencyShp ProductUpdateParamsPriceOneTimePriceCurrency = "SHP"
	ProductUpdateParamsPriceOneTimePriceCurrencySle ProductUpdateParamsPriceOneTimePriceCurrency = "SLE"
	ProductUpdateParamsPriceOneTimePriceCurrencySll ProductUpdateParamsPriceOneTimePriceCurrency = "SLL"
	ProductUpdateParamsPriceOneTimePriceCurrencySos ProductUpdateParamsPriceOneTimePriceCurrency = "SOS"
	ProductUpdateParamsPriceOneTimePriceCurrencySrd ProductUpdateParamsPriceOneTimePriceCurrency = "SRD"
	ProductUpdateParamsPriceOneTimePriceCurrencySsp ProductUpdateParamsPriceOneTimePriceCurrency = "SSP"
	ProductUpdateParamsPriceOneTimePriceCurrencyStn ProductUpdateParamsPriceOneTimePriceCurrency = "STN"
	ProductUpdateParamsPriceOneTimePriceCurrencySvc ProductUpdateParamsPriceOneTimePriceCurrency = "SVC"
	ProductUpdateParamsPriceOneTimePriceCurrencySzl ProductUpdateParamsPriceOneTimePriceCurrency = "SZL"
	ProductUpdateParamsPriceOneTimePriceCurrencyThb ProductUpdateParamsPriceOneTimePriceCurrency = "THB"
	ProductUpdateParamsPriceOneTimePriceCurrencyTnd ProductUpdateParamsPriceOneTimePriceCurrency = "TND"
	ProductUpdateParamsPriceOneTimePriceCurrencyTop ProductUpdateParamsPriceOneTimePriceCurrency = "TOP"
	ProductUpdateParamsPriceOneTimePriceCurrencyTry ProductUpdateParamsPriceOneTimePriceCurrency = "TRY"
	ProductUpdateParamsPriceOneTimePriceCurrencyTtd ProductUpdateParamsPriceOneTimePriceCurrency = "TTD"
	ProductUpdateParamsPriceOneTimePriceCurrencyTwd ProductUpdateParamsPriceOneTimePriceCurrency = "TWD"
	ProductUpdateParamsPriceOneTimePriceCurrencyTzs ProductUpdateParamsPriceOneTimePriceCurrency = "TZS"
	ProductUpdateParamsPriceOneTimePriceCurrencyUah ProductUpdateParamsPriceOneTimePriceCurrency = "UAH"
	ProductUpdateParamsPriceOneTimePriceCurrencyUgx ProductUpdateParamsPriceOneTimePriceCurrency = "UGX"
	ProductUpdateParamsPriceOneTimePriceCurrencyUsd ProductUpdateParamsPriceOneTimePriceCurrency = "USD"
	ProductUpdateParamsPriceOneTimePriceCurrencyUyu ProductUpdateParamsPriceOneTimePriceCurrency = "UYU"
	ProductUpdateParamsPriceOneTimePriceCurrencyUzs ProductUpdateParamsPriceOneTimePriceCurrency = "UZS"
	ProductUpdateParamsPriceOneTimePriceCurrencyVes ProductUpdateParamsPriceOneTimePriceCurrency = "VES"
	ProductUpdateParamsPriceOneTimePriceCurrencyVnd ProductUpdateParamsPriceOneTimePriceCurrency = "VND"
	ProductUpdateParamsPriceOneTimePriceCurrencyVuv ProductUpdateParamsPriceOneTimePriceCurrency = "VUV"
	ProductUpdateParamsPriceOneTimePriceCurrencyWst ProductUpdateParamsPriceOneTimePriceCurrency = "WST"
	ProductUpdateParamsPriceOneTimePriceCurrencyXaf ProductUpdateParamsPriceOneTimePriceCurrency = "XAF"
	ProductUpdateParamsPriceOneTimePriceCurrencyXcd ProductUpdateParamsPriceOneTimePriceCurrency = "XCD"
	ProductUpdateParamsPriceOneTimePriceCurrencyXof ProductUpdateParamsPriceOneTimePriceCurrency = "XOF"
	ProductUpdateParamsPriceOneTimePriceCurrencyXpf ProductUpdateParamsPriceOneTimePriceCurrency = "XPF"
	ProductUpdateParamsPriceOneTimePriceCurrencyYer ProductUpdateParamsPriceOneTimePriceCurrency = "YER"
	ProductUpdateParamsPriceOneTimePriceCurrencyZar ProductUpdateParamsPriceOneTimePriceCurrency = "ZAR"
	ProductUpdateParamsPriceOneTimePriceCurrencyZmw ProductUpdateParamsPriceOneTimePriceCurrency = "ZMW"
)

func (r ProductUpdateParamsPriceOneTimePriceCurrency) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceOneTimePriceCurrencyAed, ProductUpdateParamsPriceOneTimePriceCurrencyAll, ProductUpdateParamsPriceOneTimePriceCurrencyAmd, ProductUpdateParamsPriceOneTimePriceCurrencyAng, ProductUpdateParamsPriceOneTimePriceCurrencyAoa, ProductUpdateParamsPriceOneTimePriceCurrencyArs, ProductUpdateParamsPriceOneTimePriceCurrencyAud, ProductUpdateParamsPriceOneTimePriceCurrencyAwg, ProductUpdateParamsPriceOneTimePriceCurrencyAzn, ProductUpdateParamsPriceOneTimePriceCurrencyBam, ProductUpdateParamsPriceOneTimePriceCurrencyBbd, ProductUpdateParamsPriceOneTimePriceCurrencyBdt, ProductUpdateParamsPriceOneTimePriceCurrencyBgn, ProductUpdateParamsPriceOneTimePriceCurrencyBhd, ProductUpdateParamsPriceOneTimePriceCurrencyBif, ProductUpdateParamsPriceOneTimePriceCurrencyBmd, ProductUpdateParamsPriceOneTimePriceCurrencyBnd, ProductUpdateParamsPriceOneTimePriceCurrencyBob, ProductUpdateParamsPriceOneTimePriceCurrencyBrl, ProductUpdateParamsPriceOneTimePriceCurrencyBsd, ProductUpdateParamsPriceOneTimePriceCurrencyBwp, ProductUpdateParamsPriceOneTimePriceCurrencyByn, ProductUpdateParamsPriceOneTimePriceCurrencyBzd, ProductUpdateParamsPriceOneTimePriceCurrencyCad, ProductUpdateParamsPriceOneTimePriceCurrencyChf, ProductUpdateParamsPriceOneTimePriceCurrencyClp, ProductUpdateParamsPriceOneTimePriceCurrencyCny, ProductUpdateParamsPriceOneTimePriceCurrencyCop, ProductUpdateParamsPriceOneTimePriceCurrencyCrc, ProductUpdateParamsPriceOneTimePriceCurrencyCup, ProductUpdateParamsPriceOneTimePriceCurrencyCve, ProductUpdateParamsPriceOneTimePriceCurrencyCzk, ProductUpdateParamsPriceOneTimePriceCurrencyDjf, ProductUpdateParamsPriceOneTimePriceCurrencyDkk, ProductUpdateParamsPriceOneTimePriceCurrencyDop, ProductUpdateParamsPriceOneTimePriceCurrencyDzd, ProductUpdateParamsPriceOneTimePriceCurrencyEgp, ProductUpdateParamsPriceOneTimePriceCurrencyEtb, ProductUpdateParamsPriceOneTimePriceCurrencyEur, ProductUpdateParamsPriceOneTimePriceCurrencyFjd, ProductUpdateParamsPriceOneTimePriceCurrencyFkp, ProductUpdateParamsPriceOneTimePriceCurrencyGbp, ProductUpdateParamsPriceOneTimePriceCurrencyGel, ProductUpdateParamsPriceOneTimePriceCurrencyGhs, ProductUpdateParamsPriceOneTimePriceCurrencyGip, ProductUpdateParamsPriceOneTimePriceCurrencyGmd, ProductUpdateParamsPriceOneTimePriceCurrencyGnf, ProductUpdateParamsPriceOneTimePriceCurrencyGtq, ProductUpdateParamsPriceOneTimePriceCurrencyGyd, ProductUpdateParamsPriceOneTimePriceCurrencyHkd, ProductUpdateParamsPriceOneTimePriceCurrencyHnl, ProductUpdateParamsPriceOneTimePriceCurrencyHrk, ProductUpdateParamsPriceOneTimePriceCurrencyHtg, ProductUpdateParamsPriceOneTimePriceCurrencyHuf, ProductUpdateParamsPriceOneTimePriceCurrencyIdr, ProductUpdateParamsPriceOneTimePriceCurrencyIls, ProductUpdateParamsPriceOneTimePriceCurrencyInr, ProductUpdateParamsPriceOneTimePriceCurrencyIqd, ProductUpdateParamsPriceOneTimePriceCurrencyJmd, ProductUpdateParamsPriceOneTimePriceCurrencyJod, ProductUpdateParamsPriceOneTimePriceCurrencyJpy, ProductUpdateParamsPriceOneTimePriceCurrencyKes, ProductUpdateParamsPriceOneTimePriceCurrencyKgs, ProductUpdateParamsPriceOneTimePriceCurrencyKhr, ProductUpdateParamsPriceOneTimePriceCurrencyKmf, ProductUpdateParamsPriceOneTimePriceCurrencyKrw, ProductUpdateParamsPriceOneTimePriceCurrencyKwd, ProductUpdateParamsPriceOneTimePriceCurrencyKyd, ProductUpdateParamsPriceOneTimePriceCurrencyKzt, ProductUpdateParamsPriceOneTimePriceCurrencyLak, ProductUpdateParamsPriceOneTimePriceCurrencyLbp, ProductUpdateParamsPriceOneTimePriceCurrencyLkr, ProductUpdateParamsPriceOneTimePriceCurrencyLrd, ProductUpdateParamsPriceOneTimePriceCurrencyLsl, ProductUpdateParamsPriceOneTimePriceCurrencyLyd, ProductUpdateParamsPriceOneTimePriceCurrencyMad, ProductUpdateParamsPriceOneTimePriceCurrencyMdl, ProductUpdateParamsPriceOneTimePriceCurrencyMga, ProductUpdateParamsPriceOneTimePriceCurrencyMkd, ProductUpdateParamsPriceOneTimePriceCurrencyMmk, ProductUpdateParamsPriceOneTimePriceCurrencyMnt, ProductUpdateParamsPriceOneTimePriceCurrencyMop, ProductUpdateParamsPriceOneTimePriceCurrencyMru, ProductUpdateParamsPriceOneTimePriceCurrencyMur, ProductUpdateParamsPriceOneTimePriceCurrencyMvr, ProductUpdateParamsPriceOneTimePriceCurrencyMwk, ProductUpdateParamsPriceOneTimePriceCurrencyMxn, ProductUpdateParamsPriceOneTimePriceCurrencyMyr, ProductUpdateParamsPriceOneTimePriceCurrencyMzn, ProductUpdateParamsPriceOneTimePriceCurrencyNad, ProductUpdateParamsPriceOneTimePriceCurrencyNgn, ProductUpdateParamsPriceOneTimePriceCurrencyNio, ProductUpdateParamsPriceOneTimePriceCurrencyNok, ProductUpdateParamsPriceOneTimePriceCurrencyNpr, ProductUpdateParamsPriceOneTimePriceCurrencyNzd, ProductUpdateParamsPriceOneTimePriceCurrencyOmr, ProductUpdateParamsPriceOneTimePriceCurrencyPab, ProductUpdateParamsPriceOneTimePriceCurrencyPen, ProductUpdateParamsPriceOneTimePriceCurrencyPgk, ProductUpdateParamsPriceOneTimePriceCurrencyPhp, ProductUpdateParamsPriceOneTimePriceCurrencyPkr, ProductUpdateParamsPriceOneTimePriceCurrencyPln, ProductUpdateParamsPriceOneTimePriceCurrencyPyg, ProductUpdateParamsPriceOneTimePriceCurrencyQar, ProductUpdateParamsPriceOneTimePriceCurrencyRon, ProductUpdateParamsPriceOneTimePriceCurrencyRsd, ProductUpdateParamsPriceOneTimePriceCurrencyRub, ProductUpdateParamsPriceOneTimePriceCurrencyRwf, ProductUpdateParamsPriceOneTimePriceCurrencySar, ProductUpdateParamsPriceOneTimePriceCurrencySbd, ProductUpdateParamsPriceOneTimePriceCurrencyScr, ProductUpdateParamsPriceOneTimePriceCurrencySek, ProductUpdateParamsPriceOneTimePriceCurrencySgd, ProductUpdateParamsPriceOneTimePriceCurrencyShp, ProductUpdateParamsPriceOneTimePriceCurrencySle, ProductUpdateParamsPriceOneTimePriceCurrencySll, ProductUpdateParamsPriceOneTimePriceCurrencySos, ProductUpdateParamsPriceOneTimePriceCurrencySrd, ProductUpdateParamsPriceOneTimePriceCurrencySsp, ProductUpdateParamsPriceOneTimePriceCurrencyStn, ProductUpdateParamsPriceOneTimePriceCurrencySvc, ProductUpdateParamsPriceOneTimePriceCurrencySzl, ProductUpdateParamsPriceOneTimePriceCurrencyThb, ProductUpdateParamsPriceOneTimePriceCurrencyTnd, ProductUpdateParamsPriceOneTimePriceCurrencyTop, ProductUpdateParamsPriceOneTimePriceCurrencyTry, ProductUpdateParamsPriceOneTimePriceCurrencyTtd, ProductUpdateParamsPriceOneTimePriceCurrencyTwd, ProductUpdateParamsPriceOneTimePriceCurrencyTzs, ProductUpdateParamsPriceOneTimePriceCurrencyUah, ProductUpdateParamsPriceOneTimePriceCurrencyUgx, ProductUpdateParamsPriceOneTimePriceCurrencyUsd, ProductUpdateParamsPriceOneTimePriceCurrencyUyu, ProductUpdateParamsPriceOneTimePriceCurrencyUzs, ProductUpdateParamsPriceOneTimePriceCurrencyVes, ProductUpdateParamsPriceOneTimePriceCurrencyVnd, ProductUpdateParamsPriceOneTimePriceCurrencyVuv, ProductUpdateParamsPriceOneTimePriceCurrencyWst, ProductUpdateParamsPriceOneTimePriceCurrencyXaf, ProductUpdateParamsPriceOneTimePriceCurrencyXcd, ProductUpdateParamsPriceOneTimePriceCurrencyXof, ProductUpdateParamsPriceOneTimePriceCurrencyXpf, ProductUpdateParamsPriceOneTimePriceCurrencyYer, ProductUpdateParamsPriceOneTimePriceCurrencyZar, ProductUpdateParamsPriceOneTimePriceCurrencyZmw:
		return true
	}
	return false
}

type ProductUpdateParamsPriceOneTimePriceType string

const (
	ProductUpdateParamsPriceOneTimePriceTypeOneTimePrice ProductUpdateParamsPriceOneTimePriceType = "one_time_price"
)

func (r ProductUpdateParamsPriceOneTimePriceType) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceOneTimePriceTypeOneTimePrice:
		return true
	}
	return false
}

type ProductUpdateParamsPriceRecurringPrice struct {
	Currency param.Field[ProductUpdateParamsPriceRecurringPriceCurrency] `json:"currency,required"`
	// Discount applied to the price, represented as a percentage (0 to 100).
	Discount param.Field[float64] `json:"discount,required"`
	// Number of units for the payment frequency. For example, a value of `1` with a
	// `payment_frequency_interval` of `month` represents monthly payments.
	PaymentFrequencyCount    param.Field[int64]                                                          `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval param.Field[ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval] `json:"payment_frequency_interval,required"`
	// The payment amount. Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Price param.Field[int64] `json:"price,required"`
	// Indicates if purchasing power parity adjustments are applied to the price.
	// Purchasing power parity feature is not available as of now
	PurchasingPowerParity param.Field[bool] `json:"purchasing_power_parity,required"`
	// Number of units for the subscription period. For example, a value of `12` with a
	// `subscription_period_interval` of `month` represents a one-year subscription.
	SubscriptionPeriodCount    param.Field[int64]                                                            `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval param.Field[ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval] `json:"subscription_period_interval,required"`
	Type                       param.Field[ProductUpdateParamsPriceRecurringPriceType]                       `json:"type,required"`
	// Indicates if the price is tax inclusive
	TaxInclusive param.Field[bool] `json:"tax_inclusive"`
	// Number of days for the trial period. A value of `0` indicates no trial period.
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r ProductUpdateParamsPriceRecurringPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProductUpdateParamsPriceRecurringPrice) implementsProductUpdateParamsPriceUnion() {}

type ProductUpdateParamsPriceRecurringPriceCurrency string

const (
	ProductUpdateParamsPriceRecurringPriceCurrencyAed ProductUpdateParamsPriceRecurringPriceCurrency = "AED"
	ProductUpdateParamsPriceRecurringPriceCurrencyAll ProductUpdateParamsPriceRecurringPriceCurrency = "ALL"
	ProductUpdateParamsPriceRecurringPriceCurrencyAmd ProductUpdateParamsPriceRecurringPriceCurrency = "AMD"
	ProductUpdateParamsPriceRecurringPriceCurrencyAng ProductUpdateParamsPriceRecurringPriceCurrency = "ANG"
	ProductUpdateParamsPriceRecurringPriceCurrencyAoa ProductUpdateParamsPriceRecurringPriceCurrency = "AOA"
	ProductUpdateParamsPriceRecurringPriceCurrencyArs ProductUpdateParamsPriceRecurringPriceCurrency = "ARS"
	ProductUpdateParamsPriceRecurringPriceCurrencyAud ProductUpdateParamsPriceRecurringPriceCurrency = "AUD"
	ProductUpdateParamsPriceRecurringPriceCurrencyAwg ProductUpdateParamsPriceRecurringPriceCurrency = "AWG"
	ProductUpdateParamsPriceRecurringPriceCurrencyAzn ProductUpdateParamsPriceRecurringPriceCurrency = "AZN"
	ProductUpdateParamsPriceRecurringPriceCurrencyBam ProductUpdateParamsPriceRecurringPriceCurrency = "BAM"
	ProductUpdateParamsPriceRecurringPriceCurrencyBbd ProductUpdateParamsPriceRecurringPriceCurrency = "BBD"
	ProductUpdateParamsPriceRecurringPriceCurrencyBdt ProductUpdateParamsPriceRecurringPriceCurrency = "BDT"
	ProductUpdateParamsPriceRecurringPriceCurrencyBgn ProductUpdateParamsPriceRecurringPriceCurrency = "BGN"
	ProductUpdateParamsPriceRecurringPriceCurrencyBhd ProductUpdateParamsPriceRecurringPriceCurrency = "BHD"
	ProductUpdateParamsPriceRecurringPriceCurrencyBif ProductUpdateParamsPriceRecurringPriceCurrency = "BIF"
	ProductUpdateParamsPriceRecurringPriceCurrencyBmd ProductUpdateParamsPriceRecurringPriceCurrency = "BMD"
	ProductUpdateParamsPriceRecurringPriceCurrencyBnd ProductUpdateParamsPriceRecurringPriceCurrency = "BND"
	ProductUpdateParamsPriceRecurringPriceCurrencyBob ProductUpdateParamsPriceRecurringPriceCurrency = "BOB"
	ProductUpdateParamsPriceRecurringPriceCurrencyBrl ProductUpdateParamsPriceRecurringPriceCurrency = "BRL"
	ProductUpdateParamsPriceRecurringPriceCurrencyBsd ProductUpdateParamsPriceRecurringPriceCurrency = "BSD"
	ProductUpdateParamsPriceRecurringPriceCurrencyBwp ProductUpdateParamsPriceRecurringPriceCurrency = "BWP"
	ProductUpdateParamsPriceRecurringPriceCurrencyByn ProductUpdateParamsPriceRecurringPriceCurrency = "BYN"
	ProductUpdateParamsPriceRecurringPriceCurrencyBzd ProductUpdateParamsPriceRecurringPriceCurrency = "BZD"
	ProductUpdateParamsPriceRecurringPriceCurrencyCad ProductUpdateParamsPriceRecurringPriceCurrency = "CAD"
	ProductUpdateParamsPriceRecurringPriceCurrencyChf ProductUpdateParamsPriceRecurringPriceCurrency = "CHF"
	ProductUpdateParamsPriceRecurringPriceCurrencyClp ProductUpdateParamsPriceRecurringPriceCurrency = "CLP"
	ProductUpdateParamsPriceRecurringPriceCurrencyCny ProductUpdateParamsPriceRecurringPriceCurrency = "CNY"
	ProductUpdateParamsPriceRecurringPriceCurrencyCop ProductUpdateParamsPriceRecurringPriceCurrency = "COP"
	ProductUpdateParamsPriceRecurringPriceCurrencyCrc ProductUpdateParamsPriceRecurringPriceCurrency = "CRC"
	ProductUpdateParamsPriceRecurringPriceCurrencyCup ProductUpdateParamsPriceRecurringPriceCurrency = "CUP"
	ProductUpdateParamsPriceRecurringPriceCurrencyCve ProductUpdateParamsPriceRecurringPriceCurrency = "CVE"
	ProductUpdateParamsPriceRecurringPriceCurrencyCzk ProductUpdateParamsPriceRecurringPriceCurrency = "CZK"
	ProductUpdateParamsPriceRecurringPriceCurrencyDjf ProductUpdateParamsPriceRecurringPriceCurrency = "DJF"
	ProductUpdateParamsPriceRecurringPriceCurrencyDkk ProductUpdateParamsPriceRecurringPriceCurrency = "DKK"
	ProductUpdateParamsPriceRecurringPriceCurrencyDop ProductUpdateParamsPriceRecurringPriceCurrency = "DOP"
	ProductUpdateParamsPriceRecurringPriceCurrencyDzd ProductUpdateParamsPriceRecurringPriceCurrency = "DZD"
	ProductUpdateParamsPriceRecurringPriceCurrencyEgp ProductUpdateParamsPriceRecurringPriceCurrency = "EGP"
	ProductUpdateParamsPriceRecurringPriceCurrencyEtb ProductUpdateParamsPriceRecurringPriceCurrency = "ETB"
	ProductUpdateParamsPriceRecurringPriceCurrencyEur ProductUpdateParamsPriceRecurringPriceCurrency = "EUR"
	ProductUpdateParamsPriceRecurringPriceCurrencyFjd ProductUpdateParamsPriceRecurringPriceCurrency = "FJD"
	ProductUpdateParamsPriceRecurringPriceCurrencyFkp ProductUpdateParamsPriceRecurringPriceCurrency = "FKP"
	ProductUpdateParamsPriceRecurringPriceCurrencyGbp ProductUpdateParamsPriceRecurringPriceCurrency = "GBP"
	ProductUpdateParamsPriceRecurringPriceCurrencyGel ProductUpdateParamsPriceRecurringPriceCurrency = "GEL"
	ProductUpdateParamsPriceRecurringPriceCurrencyGhs ProductUpdateParamsPriceRecurringPriceCurrency = "GHS"
	ProductUpdateParamsPriceRecurringPriceCurrencyGip ProductUpdateParamsPriceRecurringPriceCurrency = "GIP"
	ProductUpdateParamsPriceRecurringPriceCurrencyGmd ProductUpdateParamsPriceRecurringPriceCurrency = "GMD"
	ProductUpdateParamsPriceRecurringPriceCurrencyGnf ProductUpdateParamsPriceRecurringPriceCurrency = "GNF"
	ProductUpdateParamsPriceRecurringPriceCurrencyGtq ProductUpdateParamsPriceRecurringPriceCurrency = "GTQ"
	ProductUpdateParamsPriceRecurringPriceCurrencyGyd ProductUpdateParamsPriceRecurringPriceCurrency = "GYD"
	ProductUpdateParamsPriceRecurringPriceCurrencyHkd ProductUpdateParamsPriceRecurringPriceCurrency = "HKD"
	ProductUpdateParamsPriceRecurringPriceCurrencyHnl ProductUpdateParamsPriceRecurringPriceCurrency = "HNL"
	ProductUpdateParamsPriceRecurringPriceCurrencyHrk ProductUpdateParamsPriceRecurringPriceCurrency = "HRK"
	ProductUpdateParamsPriceRecurringPriceCurrencyHtg ProductUpdateParamsPriceRecurringPriceCurrency = "HTG"
	ProductUpdateParamsPriceRecurringPriceCurrencyHuf ProductUpdateParamsPriceRecurringPriceCurrency = "HUF"
	ProductUpdateParamsPriceRecurringPriceCurrencyIdr ProductUpdateParamsPriceRecurringPriceCurrency = "IDR"
	ProductUpdateParamsPriceRecurringPriceCurrencyIls ProductUpdateParamsPriceRecurringPriceCurrency = "ILS"
	ProductUpdateParamsPriceRecurringPriceCurrencyInr ProductUpdateParamsPriceRecurringPriceCurrency = "INR"
	ProductUpdateParamsPriceRecurringPriceCurrencyIqd ProductUpdateParamsPriceRecurringPriceCurrency = "IQD"
	ProductUpdateParamsPriceRecurringPriceCurrencyJmd ProductUpdateParamsPriceRecurringPriceCurrency = "JMD"
	ProductUpdateParamsPriceRecurringPriceCurrencyJod ProductUpdateParamsPriceRecurringPriceCurrency = "JOD"
	ProductUpdateParamsPriceRecurringPriceCurrencyJpy ProductUpdateParamsPriceRecurringPriceCurrency = "JPY"
	ProductUpdateParamsPriceRecurringPriceCurrencyKes ProductUpdateParamsPriceRecurringPriceCurrency = "KES"
	ProductUpdateParamsPriceRecurringPriceCurrencyKgs ProductUpdateParamsPriceRecurringPriceCurrency = "KGS"
	ProductUpdateParamsPriceRecurringPriceCurrencyKhr ProductUpdateParamsPriceRecurringPriceCurrency = "KHR"
	ProductUpdateParamsPriceRecurringPriceCurrencyKmf ProductUpdateParamsPriceRecurringPriceCurrency = "KMF"
	ProductUpdateParamsPriceRecurringPriceCurrencyKrw ProductUpdateParamsPriceRecurringPriceCurrency = "KRW"
	ProductUpdateParamsPriceRecurringPriceCurrencyKwd ProductUpdateParamsPriceRecurringPriceCurrency = "KWD"
	ProductUpdateParamsPriceRecurringPriceCurrencyKyd ProductUpdateParamsPriceRecurringPriceCurrency = "KYD"
	ProductUpdateParamsPriceRecurringPriceCurrencyKzt ProductUpdateParamsPriceRecurringPriceCurrency = "KZT"
	ProductUpdateParamsPriceRecurringPriceCurrencyLak ProductUpdateParamsPriceRecurringPriceCurrency = "LAK"
	ProductUpdateParamsPriceRecurringPriceCurrencyLbp ProductUpdateParamsPriceRecurringPriceCurrency = "LBP"
	ProductUpdateParamsPriceRecurringPriceCurrencyLkr ProductUpdateParamsPriceRecurringPriceCurrency = "LKR"
	ProductUpdateParamsPriceRecurringPriceCurrencyLrd ProductUpdateParamsPriceRecurringPriceCurrency = "LRD"
	ProductUpdateParamsPriceRecurringPriceCurrencyLsl ProductUpdateParamsPriceRecurringPriceCurrency = "LSL"
	ProductUpdateParamsPriceRecurringPriceCurrencyLyd ProductUpdateParamsPriceRecurringPriceCurrency = "LYD"
	ProductUpdateParamsPriceRecurringPriceCurrencyMad ProductUpdateParamsPriceRecurringPriceCurrency = "MAD"
	ProductUpdateParamsPriceRecurringPriceCurrencyMdl ProductUpdateParamsPriceRecurringPriceCurrency = "MDL"
	ProductUpdateParamsPriceRecurringPriceCurrencyMga ProductUpdateParamsPriceRecurringPriceCurrency = "MGA"
	ProductUpdateParamsPriceRecurringPriceCurrencyMkd ProductUpdateParamsPriceRecurringPriceCurrency = "MKD"
	ProductUpdateParamsPriceRecurringPriceCurrencyMmk ProductUpdateParamsPriceRecurringPriceCurrency = "MMK"
	ProductUpdateParamsPriceRecurringPriceCurrencyMnt ProductUpdateParamsPriceRecurringPriceCurrency = "MNT"
	ProductUpdateParamsPriceRecurringPriceCurrencyMop ProductUpdateParamsPriceRecurringPriceCurrency = "MOP"
	ProductUpdateParamsPriceRecurringPriceCurrencyMru ProductUpdateParamsPriceRecurringPriceCurrency = "MRU"
	ProductUpdateParamsPriceRecurringPriceCurrencyMur ProductUpdateParamsPriceRecurringPriceCurrency = "MUR"
	ProductUpdateParamsPriceRecurringPriceCurrencyMvr ProductUpdateParamsPriceRecurringPriceCurrency = "MVR"
	ProductUpdateParamsPriceRecurringPriceCurrencyMwk ProductUpdateParamsPriceRecurringPriceCurrency = "MWK"
	ProductUpdateParamsPriceRecurringPriceCurrencyMxn ProductUpdateParamsPriceRecurringPriceCurrency = "MXN"
	ProductUpdateParamsPriceRecurringPriceCurrencyMyr ProductUpdateParamsPriceRecurringPriceCurrency = "MYR"
	ProductUpdateParamsPriceRecurringPriceCurrencyMzn ProductUpdateParamsPriceRecurringPriceCurrency = "MZN"
	ProductUpdateParamsPriceRecurringPriceCurrencyNad ProductUpdateParamsPriceRecurringPriceCurrency = "NAD"
	ProductUpdateParamsPriceRecurringPriceCurrencyNgn ProductUpdateParamsPriceRecurringPriceCurrency = "NGN"
	ProductUpdateParamsPriceRecurringPriceCurrencyNio ProductUpdateParamsPriceRecurringPriceCurrency = "NIO"
	ProductUpdateParamsPriceRecurringPriceCurrencyNok ProductUpdateParamsPriceRecurringPriceCurrency = "NOK"
	ProductUpdateParamsPriceRecurringPriceCurrencyNpr ProductUpdateParamsPriceRecurringPriceCurrency = "NPR"
	ProductUpdateParamsPriceRecurringPriceCurrencyNzd ProductUpdateParamsPriceRecurringPriceCurrency = "NZD"
	ProductUpdateParamsPriceRecurringPriceCurrencyOmr ProductUpdateParamsPriceRecurringPriceCurrency = "OMR"
	ProductUpdateParamsPriceRecurringPriceCurrencyPab ProductUpdateParamsPriceRecurringPriceCurrency = "PAB"
	ProductUpdateParamsPriceRecurringPriceCurrencyPen ProductUpdateParamsPriceRecurringPriceCurrency = "PEN"
	ProductUpdateParamsPriceRecurringPriceCurrencyPgk ProductUpdateParamsPriceRecurringPriceCurrency = "PGK"
	ProductUpdateParamsPriceRecurringPriceCurrencyPhp ProductUpdateParamsPriceRecurringPriceCurrency = "PHP"
	ProductUpdateParamsPriceRecurringPriceCurrencyPkr ProductUpdateParamsPriceRecurringPriceCurrency = "PKR"
	ProductUpdateParamsPriceRecurringPriceCurrencyPln ProductUpdateParamsPriceRecurringPriceCurrency = "PLN"
	ProductUpdateParamsPriceRecurringPriceCurrencyPyg ProductUpdateParamsPriceRecurringPriceCurrency = "PYG"
	ProductUpdateParamsPriceRecurringPriceCurrencyQar ProductUpdateParamsPriceRecurringPriceCurrency = "QAR"
	ProductUpdateParamsPriceRecurringPriceCurrencyRon ProductUpdateParamsPriceRecurringPriceCurrency = "RON"
	ProductUpdateParamsPriceRecurringPriceCurrencyRsd ProductUpdateParamsPriceRecurringPriceCurrency = "RSD"
	ProductUpdateParamsPriceRecurringPriceCurrencyRub ProductUpdateParamsPriceRecurringPriceCurrency = "RUB"
	ProductUpdateParamsPriceRecurringPriceCurrencyRwf ProductUpdateParamsPriceRecurringPriceCurrency = "RWF"
	ProductUpdateParamsPriceRecurringPriceCurrencySar ProductUpdateParamsPriceRecurringPriceCurrency = "SAR"
	ProductUpdateParamsPriceRecurringPriceCurrencySbd ProductUpdateParamsPriceRecurringPriceCurrency = "SBD"
	ProductUpdateParamsPriceRecurringPriceCurrencyScr ProductUpdateParamsPriceRecurringPriceCurrency = "SCR"
	ProductUpdateParamsPriceRecurringPriceCurrencySek ProductUpdateParamsPriceRecurringPriceCurrency = "SEK"
	ProductUpdateParamsPriceRecurringPriceCurrencySgd ProductUpdateParamsPriceRecurringPriceCurrency = "SGD"
	ProductUpdateParamsPriceRecurringPriceCurrencyShp ProductUpdateParamsPriceRecurringPriceCurrency = "SHP"
	ProductUpdateParamsPriceRecurringPriceCurrencySle ProductUpdateParamsPriceRecurringPriceCurrency = "SLE"
	ProductUpdateParamsPriceRecurringPriceCurrencySll ProductUpdateParamsPriceRecurringPriceCurrency = "SLL"
	ProductUpdateParamsPriceRecurringPriceCurrencySos ProductUpdateParamsPriceRecurringPriceCurrency = "SOS"
	ProductUpdateParamsPriceRecurringPriceCurrencySrd ProductUpdateParamsPriceRecurringPriceCurrency = "SRD"
	ProductUpdateParamsPriceRecurringPriceCurrencySsp ProductUpdateParamsPriceRecurringPriceCurrency = "SSP"
	ProductUpdateParamsPriceRecurringPriceCurrencyStn ProductUpdateParamsPriceRecurringPriceCurrency = "STN"
	ProductUpdateParamsPriceRecurringPriceCurrencySvc ProductUpdateParamsPriceRecurringPriceCurrency = "SVC"
	ProductUpdateParamsPriceRecurringPriceCurrencySzl ProductUpdateParamsPriceRecurringPriceCurrency = "SZL"
	ProductUpdateParamsPriceRecurringPriceCurrencyThb ProductUpdateParamsPriceRecurringPriceCurrency = "THB"
	ProductUpdateParamsPriceRecurringPriceCurrencyTnd ProductUpdateParamsPriceRecurringPriceCurrency = "TND"
	ProductUpdateParamsPriceRecurringPriceCurrencyTop ProductUpdateParamsPriceRecurringPriceCurrency = "TOP"
	ProductUpdateParamsPriceRecurringPriceCurrencyTry ProductUpdateParamsPriceRecurringPriceCurrency = "TRY"
	ProductUpdateParamsPriceRecurringPriceCurrencyTtd ProductUpdateParamsPriceRecurringPriceCurrency = "TTD"
	ProductUpdateParamsPriceRecurringPriceCurrencyTwd ProductUpdateParamsPriceRecurringPriceCurrency = "TWD"
	ProductUpdateParamsPriceRecurringPriceCurrencyTzs ProductUpdateParamsPriceRecurringPriceCurrency = "TZS"
	ProductUpdateParamsPriceRecurringPriceCurrencyUah ProductUpdateParamsPriceRecurringPriceCurrency = "UAH"
	ProductUpdateParamsPriceRecurringPriceCurrencyUgx ProductUpdateParamsPriceRecurringPriceCurrency = "UGX"
	ProductUpdateParamsPriceRecurringPriceCurrencyUsd ProductUpdateParamsPriceRecurringPriceCurrency = "USD"
	ProductUpdateParamsPriceRecurringPriceCurrencyUyu ProductUpdateParamsPriceRecurringPriceCurrency = "UYU"
	ProductUpdateParamsPriceRecurringPriceCurrencyUzs ProductUpdateParamsPriceRecurringPriceCurrency = "UZS"
	ProductUpdateParamsPriceRecurringPriceCurrencyVes ProductUpdateParamsPriceRecurringPriceCurrency = "VES"
	ProductUpdateParamsPriceRecurringPriceCurrencyVnd ProductUpdateParamsPriceRecurringPriceCurrency = "VND"
	ProductUpdateParamsPriceRecurringPriceCurrencyVuv ProductUpdateParamsPriceRecurringPriceCurrency = "VUV"
	ProductUpdateParamsPriceRecurringPriceCurrencyWst ProductUpdateParamsPriceRecurringPriceCurrency = "WST"
	ProductUpdateParamsPriceRecurringPriceCurrencyXaf ProductUpdateParamsPriceRecurringPriceCurrency = "XAF"
	ProductUpdateParamsPriceRecurringPriceCurrencyXcd ProductUpdateParamsPriceRecurringPriceCurrency = "XCD"
	ProductUpdateParamsPriceRecurringPriceCurrencyXof ProductUpdateParamsPriceRecurringPriceCurrency = "XOF"
	ProductUpdateParamsPriceRecurringPriceCurrencyXpf ProductUpdateParamsPriceRecurringPriceCurrency = "XPF"
	ProductUpdateParamsPriceRecurringPriceCurrencyYer ProductUpdateParamsPriceRecurringPriceCurrency = "YER"
	ProductUpdateParamsPriceRecurringPriceCurrencyZar ProductUpdateParamsPriceRecurringPriceCurrency = "ZAR"
	ProductUpdateParamsPriceRecurringPriceCurrencyZmw ProductUpdateParamsPriceRecurringPriceCurrency = "ZMW"
)

func (r ProductUpdateParamsPriceRecurringPriceCurrency) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceRecurringPriceCurrencyAed, ProductUpdateParamsPriceRecurringPriceCurrencyAll, ProductUpdateParamsPriceRecurringPriceCurrencyAmd, ProductUpdateParamsPriceRecurringPriceCurrencyAng, ProductUpdateParamsPriceRecurringPriceCurrencyAoa, ProductUpdateParamsPriceRecurringPriceCurrencyArs, ProductUpdateParamsPriceRecurringPriceCurrencyAud, ProductUpdateParamsPriceRecurringPriceCurrencyAwg, ProductUpdateParamsPriceRecurringPriceCurrencyAzn, ProductUpdateParamsPriceRecurringPriceCurrencyBam, ProductUpdateParamsPriceRecurringPriceCurrencyBbd, ProductUpdateParamsPriceRecurringPriceCurrencyBdt, ProductUpdateParamsPriceRecurringPriceCurrencyBgn, ProductUpdateParamsPriceRecurringPriceCurrencyBhd, ProductUpdateParamsPriceRecurringPriceCurrencyBif, ProductUpdateParamsPriceRecurringPriceCurrencyBmd, ProductUpdateParamsPriceRecurringPriceCurrencyBnd, ProductUpdateParamsPriceRecurringPriceCurrencyBob, ProductUpdateParamsPriceRecurringPriceCurrencyBrl, ProductUpdateParamsPriceRecurringPriceCurrencyBsd, ProductUpdateParamsPriceRecurringPriceCurrencyBwp, ProductUpdateParamsPriceRecurringPriceCurrencyByn, ProductUpdateParamsPriceRecurringPriceCurrencyBzd, ProductUpdateParamsPriceRecurringPriceCurrencyCad, ProductUpdateParamsPriceRecurringPriceCurrencyChf, ProductUpdateParamsPriceRecurringPriceCurrencyClp, ProductUpdateParamsPriceRecurringPriceCurrencyCny, ProductUpdateParamsPriceRecurringPriceCurrencyCop, ProductUpdateParamsPriceRecurringPriceCurrencyCrc, ProductUpdateParamsPriceRecurringPriceCurrencyCup, ProductUpdateParamsPriceRecurringPriceCurrencyCve, ProductUpdateParamsPriceRecurringPriceCurrencyCzk, ProductUpdateParamsPriceRecurringPriceCurrencyDjf, ProductUpdateParamsPriceRecurringPriceCurrencyDkk, ProductUpdateParamsPriceRecurringPriceCurrencyDop, ProductUpdateParamsPriceRecurringPriceCurrencyDzd, ProductUpdateParamsPriceRecurringPriceCurrencyEgp, ProductUpdateParamsPriceRecurringPriceCurrencyEtb, ProductUpdateParamsPriceRecurringPriceCurrencyEur, ProductUpdateParamsPriceRecurringPriceCurrencyFjd, ProductUpdateParamsPriceRecurringPriceCurrencyFkp, ProductUpdateParamsPriceRecurringPriceCurrencyGbp, ProductUpdateParamsPriceRecurringPriceCurrencyGel, ProductUpdateParamsPriceRecurringPriceCurrencyGhs, ProductUpdateParamsPriceRecurringPriceCurrencyGip, ProductUpdateParamsPriceRecurringPriceCurrencyGmd, ProductUpdateParamsPriceRecurringPriceCurrencyGnf, ProductUpdateParamsPriceRecurringPriceCurrencyGtq, ProductUpdateParamsPriceRecurringPriceCurrencyGyd, ProductUpdateParamsPriceRecurringPriceCurrencyHkd, ProductUpdateParamsPriceRecurringPriceCurrencyHnl, ProductUpdateParamsPriceRecurringPriceCurrencyHrk, ProductUpdateParamsPriceRecurringPriceCurrencyHtg, ProductUpdateParamsPriceRecurringPriceCurrencyHuf, ProductUpdateParamsPriceRecurringPriceCurrencyIdr, ProductUpdateParamsPriceRecurringPriceCurrencyIls, ProductUpdateParamsPriceRecurringPriceCurrencyInr, ProductUpdateParamsPriceRecurringPriceCurrencyIqd, ProductUpdateParamsPriceRecurringPriceCurrencyJmd, ProductUpdateParamsPriceRecurringPriceCurrencyJod, ProductUpdateParamsPriceRecurringPriceCurrencyJpy, ProductUpdateParamsPriceRecurringPriceCurrencyKes, ProductUpdateParamsPriceRecurringPriceCurrencyKgs, ProductUpdateParamsPriceRecurringPriceCurrencyKhr, ProductUpdateParamsPriceRecurringPriceCurrencyKmf, ProductUpdateParamsPriceRecurringPriceCurrencyKrw, ProductUpdateParamsPriceRecurringPriceCurrencyKwd, ProductUpdateParamsPriceRecurringPriceCurrencyKyd, ProductUpdateParamsPriceRecurringPriceCurrencyKzt, ProductUpdateParamsPriceRecurringPriceCurrencyLak, ProductUpdateParamsPriceRecurringPriceCurrencyLbp, ProductUpdateParamsPriceRecurringPriceCurrencyLkr, ProductUpdateParamsPriceRecurringPriceCurrencyLrd, ProductUpdateParamsPriceRecurringPriceCurrencyLsl, ProductUpdateParamsPriceRecurringPriceCurrencyLyd, ProductUpdateParamsPriceRecurringPriceCurrencyMad, ProductUpdateParamsPriceRecurringPriceCurrencyMdl, ProductUpdateParamsPriceRecurringPriceCurrencyMga, ProductUpdateParamsPriceRecurringPriceCurrencyMkd, ProductUpdateParamsPriceRecurringPriceCurrencyMmk, ProductUpdateParamsPriceRecurringPriceCurrencyMnt, ProductUpdateParamsPriceRecurringPriceCurrencyMop, ProductUpdateParamsPriceRecurringPriceCurrencyMru, ProductUpdateParamsPriceRecurringPriceCurrencyMur, ProductUpdateParamsPriceRecurringPriceCurrencyMvr, ProductUpdateParamsPriceRecurringPriceCurrencyMwk, ProductUpdateParamsPriceRecurringPriceCurrencyMxn, ProductUpdateParamsPriceRecurringPriceCurrencyMyr, ProductUpdateParamsPriceRecurringPriceCurrencyMzn, ProductUpdateParamsPriceRecurringPriceCurrencyNad, ProductUpdateParamsPriceRecurringPriceCurrencyNgn, ProductUpdateParamsPriceRecurringPriceCurrencyNio, ProductUpdateParamsPriceRecurringPriceCurrencyNok, ProductUpdateParamsPriceRecurringPriceCurrencyNpr, ProductUpdateParamsPriceRecurringPriceCurrencyNzd, ProductUpdateParamsPriceRecurringPriceCurrencyOmr, ProductUpdateParamsPriceRecurringPriceCurrencyPab, ProductUpdateParamsPriceRecurringPriceCurrencyPen, ProductUpdateParamsPriceRecurringPriceCurrencyPgk, ProductUpdateParamsPriceRecurringPriceCurrencyPhp, ProductUpdateParamsPriceRecurringPriceCurrencyPkr, ProductUpdateParamsPriceRecurringPriceCurrencyPln, ProductUpdateParamsPriceRecurringPriceCurrencyPyg, ProductUpdateParamsPriceRecurringPriceCurrencyQar, ProductUpdateParamsPriceRecurringPriceCurrencyRon, ProductUpdateParamsPriceRecurringPriceCurrencyRsd, ProductUpdateParamsPriceRecurringPriceCurrencyRub, ProductUpdateParamsPriceRecurringPriceCurrencyRwf, ProductUpdateParamsPriceRecurringPriceCurrencySar, ProductUpdateParamsPriceRecurringPriceCurrencySbd, ProductUpdateParamsPriceRecurringPriceCurrencyScr, ProductUpdateParamsPriceRecurringPriceCurrencySek, ProductUpdateParamsPriceRecurringPriceCurrencySgd, ProductUpdateParamsPriceRecurringPriceCurrencyShp, ProductUpdateParamsPriceRecurringPriceCurrencySle, ProductUpdateParamsPriceRecurringPriceCurrencySll, ProductUpdateParamsPriceRecurringPriceCurrencySos, ProductUpdateParamsPriceRecurringPriceCurrencySrd, ProductUpdateParamsPriceRecurringPriceCurrencySsp, ProductUpdateParamsPriceRecurringPriceCurrencyStn, ProductUpdateParamsPriceRecurringPriceCurrencySvc, ProductUpdateParamsPriceRecurringPriceCurrencySzl, ProductUpdateParamsPriceRecurringPriceCurrencyThb, ProductUpdateParamsPriceRecurringPriceCurrencyTnd, ProductUpdateParamsPriceRecurringPriceCurrencyTop, ProductUpdateParamsPriceRecurringPriceCurrencyTry, ProductUpdateParamsPriceRecurringPriceCurrencyTtd, ProductUpdateParamsPriceRecurringPriceCurrencyTwd, ProductUpdateParamsPriceRecurringPriceCurrencyTzs, ProductUpdateParamsPriceRecurringPriceCurrencyUah, ProductUpdateParamsPriceRecurringPriceCurrencyUgx, ProductUpdateParamsPriceRecurringPriceCurrencyUsd, ProductUpdateParamsPriceRecurringPriceCurrencyUyu, ProductUpdateParamsPriceRecurringPriceCurrencyUzs, ProductUpdateParamsPriceRecurringPriceCurrencyVes, ProductUpdateParamsPriceRecurringPriceCurrencyVnd, ProductUpdateParamsPriceRecurringPriceCurrencyVuv, ProductUpdateParamsPriceRecurringPriceCurrencyWst, ProductUpdateParamsPriceRecurringPriceCurrencyXaf, ProductUpdateParamsPriceRecurringPriceCurrencyXcd, ProductUpdateParamsPriceRecurringPriceCurrencyXof, ProductUpdateParamsPriceRecurringPriceCurrencyXpf, ProductUpdateParamsPriceRecurringPriceCurrencyYer, ProductUpdateParamsPriceRecurringPriceCurrencyZar, ProductUpdateParamsPriceRecurringPriceCurrencyZmw:
		return true
	}
	return false
}

type ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval string

const (
	ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalDay   ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval = "Day"
	ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalWeek  ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval = "Week"
	ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalMonth ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval = "Month"
	ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalYear  ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval = "Year"
)

func (r ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalDay, ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalWeek, ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalMonth, ProductUpdateParamsPriceRecurringPricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval string

const (
	ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalDay   ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval = "Day"
	ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalWeek  ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval = "Week"
	ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalMonth ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval = "Month"
	ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalYear  ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval = "Year"
)

func (r ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalDay, ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalWeek, ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalMonth, ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

type ProductUpdateParamsPriceRecurringPriceType string

const (
	ProductUpdateParamsPriceRecurringPriceTypeRecurringPrice ProductUpdateParamsPriceRecurringPriceType = "recurring_price"
)

func (r ProductUpdateParamsPriceRecurringPriceType) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceRecurringPriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductUpdateParamsPriceCurrency string

const (
	ProductUpdateParamsPriceCurrencyAed ProductUpdateParamsPriceCurrency = "AED"
	ProductUpdateParamsPriceCurrencyAll ProductUpdateParamsPriceCurrency = "ALL"
	ProductUpdateParamsPriceCurrencyAmd ProductUpdateParamsPriceCurrency = "AMD"
	ProductUpdateParamsPriceCurrencyAng ProductUpdateParamsPriceCurrency = "ANG"
	ProductUpdateParamsPriceCurrencyAoa ProductUpdateParamsPriceCurrency = "AOA"
	ProductUpdateParamsPriceCurrencyArs ProductUpdateParamsPriceCurrency = "ARS"
	ProductUpdateParamsPriceCurrencyAud ProductUpdateParamsPriceCurrency = "AUD"
	ProductUpdateParamsPriceCurrencyAwg ProductUpdateParamsPriceCurrency = "AWG"
	ProductUpdateParamsPriceCurrencyAzn ProductUpdateParamsPriceCurrency = "AZN"
	ProductUpdateParamsPriceCurrencyBam ProductUpdateParamsPriceCurrency = "BAM"
	ProductUpdateParamsPriceCurrencyBbd ProductUpdateParamsPriceCurrency = "BBD"
	ProductUpdateParamsPriceCurrencyBdt ProductUpdateParamsPriceCurrency = "BDT"
	ProductUpdateParamsPriceCurrencyBgn ProductUpdateParamsPriceCurrency = "BGN"
	ProductUpdateParamsPriceCurrencyBhd ProductUpdateParamsPriceCurrency = "BHD"
	ProductUpdateParamsPriceCurrencyBif ProductUpdateParamsPriceCurrency = "BIF"
	ProductUpdateParamsPriceCurrencyBmd ProductUpdateParamsPriceCurrency = "BMD"
	ProductUpdateParamsPriceCurrencyBnd ProductUpdateParamsPriceCurrency = "BND"
	ProductUpdateParamsPriceCurrencyBob ProductUpdateParamsPriceCurrency = "BOB"
	ProductUpdateParamsPriceCurrencyBrl ProductUpdateParamsPriceCurrency = "BRL"
	ProductUpdateParamsPriceCurrencyBsd ProductUpdateParamsPriceCurrency = "BSD"
	ProductUpdateParamsPriceCurrencyBwp ProductUpdateParamsPriceCurrency = "BWP"
	ProductUpdateParamsPriceCurrencyByn ProductUpdateParamsPriceCurrency = "BYN"
	ProductUpdateParamsPriceCurrencyBzd ProductUpdateParamsPriceCurrency = "BZD"
	ProductUpdateParamsPriceCurrencyCad ProductUpdateParamsPriceCurrency = "CAD"
	ProductUpdateParamsPriceCurrencyChf ProductUpdateParamsPriceCurrency = "CHF"
	ProductUpdateParamsPriceCurrencyClp ProductUpdateParamsPriceCurrency = "CLP"
	ProductUpdateParamsPriceCurrencyCny ProductUpdateParamsPriceCurrency = "CNY"
	ProductUpdateParamsPriceCurrencyCop ProductUpdateParamsPriceCurrency = "COP"
	ProductUpdateParamsPriceCurrencyCrc ProductUpdateParamsPriceCurrency = "CRC"
	ProductUpdateParamsPriceCurrencyCup ProductUpdateParamsPriceCurrency = "CUP"
	ProductUpdateParamsPriceCurrencyCve ProductUpdateParamsPriceCurrency = "CVE"
	ProductUpdateParamsPriceCurrencyCzk ProductUpdateParamsPriceCurrency = "CZK"
	ProductUpdateParamsPriceCurrencyDjf ProductUpdateParamsPriceCurrency = "DJF"
	ProductUpdateParamsPriceCurrencyDkk ProductUpdateParamsPriceCurrency = "DKK"
	ProductUpdateParamsPriceCurrencyDop ProductUpdateParamsPriceCurrency = "DOP"
	ProductUpdateParamsPriceCurrencyDzd ProductUpdateParamsPriceCurrency = "DZD"
	ProductUpdateParamsPriceCurrencyEgp ProductUpdateParamsPriceCurrency = "EGP"
	ProductUpdateParamsPriceCurrencyEtb ProductUpdateParamsPriceCurrency = "ETB"
	ProductUpdateParamsPriceCurrencyEur ProductUpdateParamsPriceCurrency = "EUR"
	ProductUpdateParamsPriceCurrencyFjd ProductUpdateParamsPriceCurrency = "FJD"
	ProductUpdateParamsPriceCurrencyFkp ProductUpdateParamsPriceCurrency = "FKP"
	ProductUpdateParamsPriceCurrencyGbp ProductUpdateParamsPriceCurrency = "GBP"
	ProductUpdateParamsPriceCurrencyGel ProductUpdateParamsPriceCurrency = "GEL"
	ProductUpdateParamsPriceCurrencyGhs ProductUpdateParamsPriceCurrency = "GHS"
	ProductUpdateParamsPriceCurrencyGip ProductUpdateParamsPriceCurrency = "GIP"
	ProductUpdateParamsPriceCurrencyGmd ProductUpdateParamsPriceCurrency = "GMD"
	ProductUpdateParamsPriceCurrencyGnf ProductUpdateParamsPriceCurrency = "GNF"
	ProductUpdateParamsPriceCurrencyGtq ProductUpdateParamsPriceCurrency = "GTQ"
	ProductUpdateParamsPriceCurrencyGyd ProductUpdateParamsPriceCurrency = "GYD"
	ProductUpdateParamsPriceCurrencyHkd ProductUpdateParamsPriceCurrency = "HKD"
	ProductUpdateParamsPriceCurrencyHnl ProductUpdateParamsPriceCurrency = "HNL"
	ProductUpdateParamsPriceCurrencyHrk ProductUpdateParamsPriceCurrency = "HRK"
	ProductUpdateParamsPriceCurrencyHtg ProductUpdateParamsPriceCurrency = "HTG"
	ProductUpdateParamsPriceCurrencyHuf ProductUpdateParamsPriceCurrency = "HUF"
	ProductUpdateParamsPriceCurrencyIdr ProductUpdateParamsPriceCurrency = "IDR"
	ProductUpdateParamsPriceCurrencyIls ProductUpdateParamsPriceCurrency = "ILS"
	ProductUpdateParamsPriceCurrencyInr ProductUpdateParamsPriceCurrency = "INR"
	ProductUpdateParamsPriceCurrencyIqd ProductUpdateParamsPriceCurrency = "IQD"
	ProductUpdateParamsPriceCurrencyJmd ProductUpdateParamsPriceCurrency = "JMD"
	ProductUpdateParamsPriceCurrencyJod ProductUpdateParamsPriceCurrency = "JOD"
	ProductUpdateParamsPriceCurrencyJpy ProductUpdateParamsPriceCurrency = "JPY"
	ProductUpdateParamsPriceCurrencyKes ProductUpdateParamsPriceCurrency = "KES"
	ProductUpdateParamsPriceCurrencyKgs ProductUpdateParamsPriceCurrency = "KGS"
	ProductUpdateParamsPriceCurrencyKhr ProductUpdateParamsPriceCurrency = "KHR"
	ProductUpdateParamsPriceCurrencyKmf ProductUpdateParamsPriceCurrency = "KMF"
	ProductUpdateParamsPriceCurrencyKrw ProductUpdateParamsPriceCurrency = "KRW"
	ProductUpdateParamsPriceCurrencyKwd ProductUpdateParamsPriceCurrency = "KWD"
	ProductUpdateParamsPriceCurrencyKyd ProductUpdateParamsPriceCurrency = "KYD"
	ProductUpdateParamsPriceCurrencyKzt ProductUpdateParamsPriceCurrency = "KZT"
	ProductUpdateParamsPriceCurrencyLak ProductUpdateParamsPriceCurrency = "LAK"
	ProductUpdateParamsPriceCurrencyLbp ProductUpdateParamsPriceCurrency = "LBP"
	ProductUpdateParamsPriceCurrencyLkr ProductUpdateParamsPriceCurrency = "LKR"
	ProductUpdateParamsPriceCurrencyLrd ProductUpdateParamsPriceCurrency = "LRD"
	ProductUpdateParamsPriceCurrencyLsl ProductUpdateParamsPriceCurrency = "LSL"
	ProductUpdateParamsPriceCurrencyLyd ProductUpdateParamsPriceCurrency = "LYD"
	ProductUpdateParamsPriceCurrencyMad ProductUpdateParamsPriceCurrency = "MAD"
	ProductUpdateParamsPriceCurrencyMdl ProductUpdateParamsPriceCurrency = "MDL"
	ProductUpdateParamsPriceCurrencyMga ProductUpdateParamsPriceCurrency = "MGA"
	ProductUpdateParamsPriceCurrencyMkd ProductUpdateParamsPriceCurrency = "MKD"
	ProductUpdateParamsPriceCurrencyMmk ProductUpdateParamsPriceCurrency = "MMK"
	ProductUpdateParamsPriceCurrencyMnt ProductUpdateParamsPriceCurrency = "MNT"
	ProductUpdateParamsPriceCurrencyMop ProductUpdateParamsPriceCurrency = "MOP"
	ProductUpdateParamsPriceCurrencyMru ProductUpdateParamsPriceCurrency = "MRU"
	ProductUpdateParamsPriceCurrencyMur ProductUpdateParamsPriceCurrency = "MUR"
	ProductUpdateParamsPriceCurrencyMvr ProductUpdateParamsPriceCurrency = "MVR"
	ProductUpdateParamsPriceCurrencyMwk ProductUpdateParamsPriceCurrency = "MWK"
	ProductUpdateParamsPriceCurrencyMxn ProductUpdateParamsPriceCurrency = "MXN"
	ProductUpdateParamsPriceCurrencyMyr ProductUpdateParamsPriceCurrency = "MYR"
	ProductUpdateParamsPriceCurrencyMzn ProductUpdateParamsPriceCurrency = "MZN"
	ProductUpdateParamsPriceCurrencyNad ProductUpdateParamsPriceCurrency = "NAD"
	ProductUpdateParamsPriceCurrencyNgn ProductUpdateParamsPriceCurrency = "NGN"
	ProductUpdateParamsPriceCurrencyNio ProductUpdateParamsPriceCurrency = "NIO"
	ProductUpdateParamsPriceCurrencyNok ProductUpdateParamsPriceCurrency = "NOK"
	ProductUpdateParamsPriceCurrencyNpr ProductUpdateParamsPriceCurrency = "NPR"
	ProductUpdateParamsPriceCurrencyNzd ProductUpdateParamsPriceCurrency = "NZD"
	ProductUpdateParamsPriceCurrencyOmr ProductUpdateParamsPriceCurrency = "OMR"
	ProductUpdateParamsPriceCurrencyPab ProductUpdateParamsPriceCurrency = "PAB"
	ProductUpdateParamsPriceCurrencyPen ProductUpdateParamsPriceCurrency = "PEN"
	ProductUpdateParamsPriceCurrencyPgk ProductUpdateParamsPriceCurrency = "PGK"
	ProductUpdateParamsPriceCurrencyPhp ProductUpdateParamsPriceCurrency = "PHP"
	ProductUpdateParamsPriceCurrencyPkr ProductUpdateParamsPriceCurrency = "PKR"
	ProductUpdateParamsPriceCurrencyPln ProductUpdateParamsPriceCurrency = "PLN"
	ProductUpdateParamsPriceCurrencyPyg ProductUpdateParamsPriceCurrency = "PYG"
	ProductUpdateParamsPriceCurrencyQar ProductUpdateParamsPriceCurrency = "QAR"
	ProductUpdateParamsPriceCurrencyRon ProductUpdateParamsPriceCurrency = "RON"
	ProductUpdateParamsPriceCurrencyRsd ProductUpdateParamsPriceCurrency = "RSD"
	ProductUpdateParamsPriceCurrencyRub ProductUpdateParamsPriceCurrency = "RUB"
	ProductUpdateParamsPriceCurrencyRwf ProductUpdateParamsPriceCurrency = "RWF"
	ProductUpdateParamsPriceCurrencySar ProductUpdateParamsPriceCurrency = "SAR"
	ProductUpdateParamsPriceCurrencySbd ProductUpdateParamsPriceCurrency = "SBD"
	ProductUpdateParamsPriceCurrencyScr ProductUpdateParamsPriceCurrency = "SCR"
	ProductUpdateParamsPriceCurrencySek ProductUpdateParamsPriceCurrency = "SEK"
	ProductUpdateParamsPriceCurrencySgd ProductUpdateParamsPriceCurrency = "SGD"
	ProductUpdateParamsPriceCurrencyShp ProductUpdateParamsPriceCurrency = "SHP"
	ProductUpdateParamsPriceCurrencySle ProductUpdateParamsPriceCurrency = "SLE"
	ProductUpdateParamsPriceCurrencySll ProductUpdateParamsPriceCurrency = "SLL"
	ProductUpdateParamsPriceCurrencySos ProductUpdateParamsPriceCurrency = "SOS"
	ProductUpdateParamsPriceCurrencySrd ProductUpdateParamsPriceCurrency = "SRD"
	ProductUpdateParamsPriceCurrencySsp ProductUpdateParamsPriceCurrency = "SSP"
	ProductUpdateParamsPriceCurrencyStn ProductUpdateParamsPriceCurrency = "STN"
	ProductUpdateParamsPriceCurrencySvc ProductUpdateParamsPriceCurrency = "SVC"
	ProductUpdateParamsPriceCurrencySzl ProductUpdateParamsPriceCurrency = "SZL"
	ProductUpdateParamsPriceCurrencyThb ProductUpdateParamsPriceCurrency = "THB"
	ProductUpdateParamsPriceCurrencyTnd ProductUpdateParamsPriceCurrency = "TND"
	ProductUpdateParamsPriceCurrencyTop ProductUpdateParamsPriceCurrency = "TOP"
	ProductUpdateParamsPriceCurrencyTry ProductUpdateParamsPriceCurrency = "TRY"
	ProductUpdateParamsPriceCurrencyTtd ProductUpdateParamsPriceCurrency = "TTD"
	ProductUpdateParamsPriceCurrencyTwd ProductUpdateParamsPriceCurrency = "TWD"
	ProductUpdateParamsPriceCurrencyTzs ProductUpdateParamsPriceCurrency = "TZS"
	ProductUpdateParamsPriceCurrencyUah ProductUpdateParamsPriceCurrency = "UAH"
	ProductUpdateParamsPriceCurrencyUgx ProductUpdateParamsPriceCurrency = "UGX"
	ProductUpdateParamsPriceCurrencyUsd ProductUpdateParamsPriceCurrency = "USD"
	ProductUpdateParamsPriceCurrencyUyu ProductUpdateParamsPriceCurrency = "UYU"
	ProductUpdateParamsPriceCurrencyUzs ProductUpdateParamsPriceCurrency = "UZS"
	ProductUpdateParamsPriceCurrencyVes ProductUpdateParamsPriceCurrency = "VES"
	ProductUpdateParamsPriceCurrencyVnd ProductUpdateParamsPriceCurrency = "VND"
	ProductUpdateParamsPriceCurrencyVuv ProductUpdateParamsPriceCurrency = "VUV"
	ProductUpdateParamsPriceCurrencyWst ProductUpdateParamsPriceCurrency = "WST"
	ProductUpdateParamsPriceCurrencyXaf ProductUpdateParamsPriceCurrency = "XAF"
	ProductUpdateParamsPriceCurrencyXcd ProductUpdateParamsPriceCurrency = "XCD"
	ProductUpdateParamsPriceCurrencyXof ProductUpdateParamsPriceCurrency = "XOF"
	ProductUpdateParamsPriceCurrencyXpf ProductUpdateParamsPriceCurrency = "XPF"
	ProductUpdateParamsPriceCurrencyYer ProductUpdateParamsPriceCurrency = "YER"
	ProductUpdateParamsPriceCurrencyZar ProductUpdateParamsPriceCurrency = "ZAR"
	ProductUpdateParamsPriceCurrencyZmw ProductUpdateParamsPriceCurrency = "ZMW"
)

func (r ProductUpdateParamsPriceCurrency) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceCurrencyAed, ProductUpdateParamsPriceCurrencyAll, ProductUpdateParamsPriceCurrencyAmd, ProductUpdateParamsPriceCurrencyAng, ProductUpdateParamsPriceCurrencyAoa, ProductUpdateParamsPriceCurrencyArs, ProductUpdateParamsPriceCurrencyAud, ProductUpdateParamsPriceCurrencyAwg, ProductUpdateParamsPriceCurrencyAzn, ProductUpdateParamsPriceCurrencyBam, ProductUpdateParamsPriceCurrencyBbd, ProductUpdateParamsPriceCurrencyBdt, ProductUpdateParamsPriceCurrencyBgn, ProductUpdateParamsPriceCurrencyBhd, ProductUpdateParamsPriceCurrencyBif, ProductUpdateParamsPriceCurrencyBmd, ProductUpdateParamsPriceCurrencyBnd, ProductUpdateParamsPriceCurrencyBob, ProductUpdateParamsPriceCurrencyBrl, ProductUpdateParamsPriceCurrencyBsd, ProductUpdateParamsPriceCurrencyBwp, ProductUpdateParamsPriceCurrencyByn, ProductUpdateParamsPriceCurrencyBzd, ProductUpdateParamsPriceCurrencyCad, ProductUpdateParamsPriceCurrencyChf, ProductUpdateParamsPriceCurrencyClp, ProductUpdateParamsPriceCurrencyCny, ProductUpdateParamsPriceCurrencyCop, ProductUpdateParamsPriceCurrencyCrc, ProductUpdateParamsPriceCurrencyCup, ProductUpdateParamsPriceCurrencyCve, ProductUpdateParamsPriceCurrencyCzk, ProductUpdateParamsPriceCurrencyDjf, ProductUpdateParamsPriceCurrencyDkk, ProductUpdateParamsPriceCurrencyDop, ProductUpdateParamsPriceCurrencyDzd, ProductUpdateParamsPriceCurrencyEgp, ProductUpdateParamsPriceCurrencyEtb, ProductUpdateParamsPriceCurrencyEur, ProductUpdateParamsPriceCurrencyFjd, ProductUpdateParamsPriceCurrencyFkp, ProductUpdateParamsPriceCurrencyGbp, ProductUpdateParamsPriceCurrencyGel, ProductUpdateParamsPriceCurrencyGhs, ProductUpdateParamsPriceCurrencyGip, ProductUpdateParamsPriceCurrencyGmd, ProductUpdateParamsPriceCurrencyGnf, ProductUpdateParamsPriceCurrencyGtq, ProductUpdateParamsPriceCurrencyGyd, ProductUpdateParamsPriceCurrencyHkd, ProductUpdateParamsPriceCurrencyHnl, ProductUpdateParamsPriceCurrencyHrk, ProductUpdateParamsPriceCurrencyHtg, ProductUpdateParamsPriceCurrencyHuf, ProductUpdateParamsPriceCurrencyIdr, ProductUpdateParamsPriceCurrencyIls, ProductUpdateParamsPriceCurrencyInr, ProductUpdateParamsPriceCurrencyIqd, ProductUpdateParamsPriceCurrencyJmd, ProductUpdateParamsPriceCurrencyJod, ProductUpdateParamsPriceCurrencyJpy, ProductUpdateParamsPriceCurrencyKes, ProductUpdateParamsPriceCurrencyKgs, ProductUpdateParamsPriceCurrencyKhr, ProductUpdateParamsPriceCurrencyKmf, ProductUpdateParamsPriceCurrencyKrw, ProductUpdateParamsPriceCurrencyKwd, ProductUpdateParamsPriceCurrencyKyd, ProductUpdateParamsPriceCurrencyKzt, ProductUpdateParamsPriceCurrencyLak, ProductUpdateParamsPriceCurrencyLbp, ProductUpdateParamsPriceCurrencyLkr, ProductUpdateParamsPriceCurrencyLrd, ProductUpdateParamsPriceCurrencyLsl, ProductUpdateParamsPriceCurrencyLyd, ProductUpdateParamsPriceCurrencyMad, ProductUpdateParamsPriceCurrencyMdl, ProductUpdateParamsPriceCurrencyMga, ProductUpdateParamsPriceCurrencyMkd, ProductUpdateParamsPriceCurrencyMmk, ProductUpdateParamsPriceCurrencyMnt, ProductUpdateParamsPriceCurrencyMop, ProductUpdateParamsPriceCurrencyMru, ProductUpdateParamsPriceCurrencyMur, ProductUpdateParamsPriceCurrencyMvr, ProductUpdateParamsPriceCurrencyMwk, ProductUpdateParamsPriceCurrencyMxn, ProductUpdateParamsPriceCurrencyMyr, ProductUpdateParamsPriceCurrencyMzn, ProductUpdateParamsPriceCurrencyNad, ProductUpdateParamsPriceCurrencyNgn, ProductUpdateParamsPriceCurrencyNio, ProductUpdateParamsPriceCurrencyNok, ProductUpdateParamsPriceCurrencyNpr, ProductUpdateParamsPriceCurrencyNzd, ProductUpdateParamsPriceCurrencyOmr, ProductUpdateParamsPriceCurrencyPab, ProductUpdateParamsPriceCurrencyPen, ProductUpdateParamsPriceCurrencyPgk, ProductUpdateParamsPriceCurrencyPhp, ProductUpdateParamsPriceCurrencyPkr, ProductUpdateParamsPriceCurrencyPln, ProductUpdateParamsPriceCurrencyPyg, ProductUpdateParamsPriceCurrencyQar, ProductUpdateParamsPriceCurrencyRon, ProductUpdateParamsPriceCurrencyRsd, ProductUpdateParamsPriceCurrencyRub, ProductUpdateParamsPriceCurrencyRwf, ProductUpdateParamsPriceCurrencySar, ProductUpdateParamsPriceCurrencySbd, ProductUpdateParamsPriceCurrencyScr, ProductUpdateParamsPriceCurrencySek, ProductUpdateParamsPriceCurrencySgd, ProductUpdateParamsPriceCurrencyShp, ProductUpdateParamsPriceCurrencySle, ProductUpdateParamsPriceCurrencySll, ProductUpdateParamsPriceCurrencySos, ProductUpdateParamsPriceCurrencySrd, ProductUpdateParamsPriceCurrencySsp, ProductUpdateParamsPriceCurrencyStn, ProductUpdateParamsPriceCurrencySvc, ProductUpdateParamsPriceCurrencySzl, ProductUpdateParamsPriceCurrencyThb, ProductUpdateParamsPriceCurrencyTnd, ProductUpdateParamsPriceCurrencyTop, ProductUpdateParamsPriceCurrencyTry, ProductUpdateParamsPriceCurrencyTtd, ProductUpdateParamsPriceCurrencyTwd, ProductUpdateParamsPriceCurrencyTzs, ProductUpdateParamsPriceCurrencyUah, ProductUpdateParamsPriceCurrencyUgx, ProductUpdateParamsPriceCurrencyUsd, ProductUpdateParamsPriceCurrencyUyu, ProductUpdateParamsPriceCurrencyUzs, ProductUpdateParamsPriceCurrencyVes, ProductUpdateParamsPriceCurrencyVnd, ProductUpdateParamsPriceCurrencyVuv, ProductUpdateParamsPriceCurrencyWst, ProductUpdateParamsPriceCurrencyXaf, ProductUpdateParamsPriceCurrencyXcd, ProductUpdateParamsPriceCurrencyXof, ProductUpdateParamsPriceCurrencyXpf, ProductUpdateParamsPriceCurrencyYer, ProductUpdateParamsPriceCurrencyZar, ProductUpdateParamsPriceCurrencyZmw:
		return true
	}
	return false
}

type ProductUpdateParamsPriceType string

const (
	ProductUpdateParamsPriceTypeOneTimePrice   ProductUpdateParamsPriceType = "one_time_price"
	ProductUpdateParamsPriceTypeRecurringPrice ProductUpdateParamsPriceType = "recurring_price"
)

func (r ProductUpdateParamsPriceType) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceTypeOneTimePrice, ProductUpdateParamsPriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductUpdateParamsPricePaymentFrequencyInterval string

const (
	ProductUpdateParamsPricePaymentFrequencyIntervalDay   ProductUpdateParamsPricePaymentFrequencyInterval = "Day"
	ProductUpdateParamsPricePaymentFrequencyIntervalWeek  ProductUpdateParamsPricePaymentFrequencyInterval = "Week"
	ProductUpdateParamsPricePaymentFrequencyIntervalMonth ProductUpdateParamsPricePaymentFrequencyInterval = "Month"
	ProductUpdateParamsPricePaymentFrequencyIntervalYear  ProductUpdateParamsPricePaymentFrequencyInterval = "Year"
)

func (r ProductUpdateParamsPricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPricePaymentFrequencyIntervalDay, ProductUpdateParamsPricePaymentFrequencyIntervalWeek, ProductUpdateParamsPricePaymentFrequencyIntervalMonth, ProductUpdateParamsPricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductUpdateParamsPriceSubscriptionPeriodInterval string

const (
	ProductUpdateParamsPriceSubscriptionPeriodIntervalDay   ProductUpdateParamsPriceSubscriptionPeriodInterval = "Day"
	ProductUpdateParamsPriceSubscriptionPeriodIntervalWeek  ProductUpdateParamsPriceSubscriptionPeriodInterval = "Week"
	ProductUpdateParamsPriceSubscriptionPeriodIntervalMonth ProductUpdateParamsPriceSubscriptionPeriodInterval = "Month"
	ProductUpdateParamsPriceSubscriptionPeriodIntervalYear  ProductUpdateParamsPriceSubscriptionPeriodInterval = "Year"
)

func (r ProductUpdateParamsPriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductUpdateParamsPriceSubscriptionPeriodIntervalDay, ProductUpdateParamsPriceSubscriptionPeriodIntervalWeek, ProductUpdateParamsPriceSubscriptionPeriodIntervalMonth, ProductUpdateParamsPriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

// Represents the different categories of taxation applicable to various products
// and services.
type ProductUpdateParamsTaxCategory string

const (
	ProductUpdateParamsTaxCategoryDigitalProducts ProductUpdateParamsTaxCategory = "digital_products"
	ProductUpdateParamsTaxCategorySaas            ProductUpdateParamsTaxCategory = "saas"
	ProductUpdateParamsTaxCategoryEBook           ProductUpdateParamsTaxCategory = "e_book"
)

func (r ProductUpdateParamsTaxCategory) IsKnown() bool {
	switch r {
	case ProductUpdateParamsTaxCategoryDigitalProducts, ProductUpdateParamsTaxCategorySaas, ProductUpdateParamsTaxCategoryEBook:
		return true
	}
	return false
}

type ProductListParams struct {
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
