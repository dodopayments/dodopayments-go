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
	Currency PriceCurrency `json:"currency,required"`
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
	Currency PriceOneTimePriceCurrency `json:"currency,required"`
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

type PriceOneTimePriceCurrency string

const (
	PriceOneTimePriceCurrencyAed PriceOneTimePriceCurrency = "AED"
	PriceOneTimePriceCurrencyAll PriceOneTimePriceCurrency = "ALL"
	PriceOneTimePriceCurrencyAmd PriceOneTimePriceCurrency = "AMD"
	PriceOneTimePriceCurrencyAng PriceOneTimePriceCurrency = "ANG"
	PriceOneTimePriceCurrencyAoa PriceOneTimePriceCurrency = "AOA"
	PriceOneTimePriceCurrencyArs PriceOneTimePriceCurrency = "ARS"
	PriceOneTimePriceCurrencyAud PriceOneTimePriceCurrency = "AUD"
	PriceOneTimePriceCurrencyAwg PriceOneTimePriceCurrency = "AWG"
	PriceOneTimePriceCurrencyAzn PriceOneTimePriceCurrency = "AZN"
	PriceOneTimePriceCurrencyBam PriceOneTimePriceCurrency = "BAM"
	PriceOneTimePriceCurrencyBbd PriceOneTimePriceCurrency = "BBD"
	PriceOneTimePriceCurrencyBdt PriceOneTimePriceCurrency = "BDT"
	PriceOneTimePriceCurrencyBgn PriceOneTimePriceCurrency = "BGN"
	PriceOneTimePriceCurrencyBhd PriceOneTimePriceCurrency = "BHD"
	PriceOneTimePriceCurrencyBif PriceOneTimePriceCurrency = "BIF"
	PriceOneTimePriceCurrencyBmd PriceOneTimePriceCurrency = "BMD"
	PriceOneTimePriceCurrencyBnd PriceOneTimePriceCurrency = "BND"
	PriceOneTimePriceCurrencyBob PriceOneTimePriceCurrency = "BOB"
	PriceOneTimePriceCurrencyBrl PriceOneTimePriceCurrency = "BRL"
	PriceOneTimePriceCurrencyBsd PriceOneTimePriceCurrency = "BSD"
	PriceOneTimePriceCurrencyBwp PriceOneTimePriceCurrency = "BWP"
	PriceOneTimePriceCurrencyByn PriceOneTimePriceCurrency = "BYN"
	PriceOneTimePriceCurrencyBzd PriceOneTimePriceCurrency = "BZD"
	PriceOneTimePriceCurrencyCad PriceOneTimePriceCurrency = "CAD"
	PriceOneTimePriceCurrencyChf PriceOneTimePriceCurrency = "CHF"
	PriceOneTimePriceCurrencyClp PriceOneTimePriceCurrency = "CLP"
	PriceOneTimePriceCurrencyCny PriceOneTimePriceCurrency = "CNY"
	PriceOneTimePriceCurrencyCop PriceOneTimePriceCurrency = "COP"
	PriceOneTimePriceCurrencyCrc PriceOneTimePriceCurrency = "CRC"
	PriceOneTimePriceCurrencyCup PriceOneTimePriceCurrency = "CUP"
	PriceOneTimePriceCurrencyCve PriceOneTimePriceCurrency = "CVE"
	PriceOneTimePriceCurrencyCzk PriceOneTimePriceCurrency = "CZK"
	PriceOneTimePriceCurrencyDjf PriceOneTimePriceCurrency = "DJF"
	PriceOneTimePriceCurrencyDkk PriceOneTimePriceCurrency = "DKK"
	PriceOneTimePriceCurrencyDop PriceOneTimePriceCurrency = "DOP"
	PriceOneTimePriceCurrencyDzd PriceOneTimePriceCurrency = "DZD"
	PriceOneTimePriceCurrencyEgp PriceOneTimePriceCurrency = "EGP"
	PriceOneTimePriceCurrencyEtb PriceOneTimePriceCurrency = "ETB"
	PriceOneTimePriceCurrencyEur PriceOneTimePriceCurrency = "EUR"
	PriceOneTimePriceCurrencyFjd PriceOneTimePriceCurrency = "FJD"
	PriceOneTimePriceCurrencyFkp PriceOneTimePriceCurrency = "FKP"
	PriceOneTimePriceCurrencyGbp PriceOneTimePriceCurrency = "GBP"
	PriceOneTimePriceCurrencyGel PriceOneTimePriceCurrency = "GEL"
	PriceOneTimePriceCurrencyGhs PriceOneTimePriceCurrency = "GHS"
	PriceOneTimePriceCurrencyGip PriceOneTimePriceCurrency = "GIP"
	PriceOneTimePriceCurrencyGmd PriceOneTimePriceCurrency = "GMD"
	PriceOneTimePriceCurrencyGnf PriceOneTimePriceCurrency = "GNF"
	PriceOneTimePriceCurrencyGtq PriceOneTimePriceCurrency = "GTQ"
	PriceOneTimePriceCurrencyGyd PriceOneTimePriceCurrency = "GYD"
	PriceOneTimePriceCurrencyHkd PriceOneTimePriceCurrency = "HKD"
	PriceOneTimePriceCurrencyHnl PriceOneTimePriceCurrency = "HNL"
	PriceOneTimePriceCurrencyHrk PriceOneTimePriceCurrency = "HRK"
	PriceOneTimePriceCurrencyHtg PriceOneTimePriceCurrency = "HTG"
	PriceOneTimePriceCurrencyHuf PriceOneTimePriceCurrency = "HUF"
	PriceOneTimePriceCurrencyIdr PriceOneTimePriceCurrency = "IDR"
	PriceOneTimePriceCurrencyIls PriceOneTimePriceCurrency = "ILS"
	PriceOneTimePriceCurrencyInr PriceOneTimePriceCurrency = "INR"
	PriceOneTimePriceCurrencyIqd PriceOneTimePriceCurrency = "IQD"
	PriceOneTimePriceCurrencyJmd PriceOneTimePriceCurrency = "JMD"
	PriceOneTimePriceCurrencyJod PriceOneTimePriceCurrency = "JOD"
	PriceOneTimePriceCurrencyJpy PriceOneTimePriceCurrency = "JPY"
	PriceOneTimePriceCurrencyKes PriceOneTimePriceCurrency = "KES"
	PriceOneTimePriceCurrencyKgs PriceOneTimePriceCurrency = "KGS"
	PriceOneTimePriceCurrencyKhr PriceOneTimePriceCurrency = "KHR"
	PriceOneTimePriceCurrencyKmf PriceOneTimePriceCurrency = "KMF"
	PriceOneTimePriceCurrencyKrw PriceOneTimePriceCurrency = "KRW"
	PriceOneTimePriceCurrencyKwd PriceOneTimePriceCurrency = "KWD"
	PriceOneTimePriceCurrencyKyd PriceOneTimePriceCurrency = "KYD"
	PriceOneTimePriceCurrencyKzt PriceOneTimePriceCurrency = "KZT"
	PriceOneTimePriceCurrencyLak PriceOneTimePriceCurrency = "LAK"
	PriceOneTimePriceCurrencyLbp PriceOneTimePriceCurrency = "LBP"
	PriceOneTimePriceCurrencyLkr PriceOneTimePriceCurrency = "LKR"
	PriceOneTimePriceCurrencyLrd PriceOneTimePriceCurrency = "LRD"
	PriceOneTimePriceCurrencyLsl PriceOneTimePriceCurrency = "LSL"
	PriceOneTimePriceCurrencyLyd PriceOneTimePriceCurrency = "LYD"
	PriceOneTimePriceCurrencyMad PriceOneTimePriceCurrency = "MAD"
	PriceOneTimePriceCurrencyMdl PriceOneTimePriceCurrency = "MDL"
	PriceOneTimePriceCurrencyMga PriceOneTimePriceCurrency = "MGA"
	PriceOneTimePriceCurrencyMkd PriceOneTimePriceCurrency = "MKD"
	PriceOneTimePriceCurrencyMmk PriceOneTimePriceCurrency = "MMK"
	PriceOneTimePriceCurrencyMnt PriceOneTimePriceCurrency = "MNT"
	PriceOneTimePriceCurrencyMop PriceOneTimePriceCurrency = "MOP"
	PriceOneTimePriceCurrencyMru PriceOneTimePriceCurrency = "MRU"
	PriceOneTimePriceCurrencyMur PriceOneTimePriceCurrency = "MUR"
	PriceOneTimePriceCurrencyMvr PriceOneTimePriceCurrency = "MVR"
	PriceOneTimePriceCurrencyMwk PriceOneTimePriceCurrency = "MWK"
	PriceOneTimePriceCurrencyMxn PriceOneTimePriceCurrency = "MXN"
	PriceOneTimePriceCurrencyMyr PriceOneTimePriceCurrency = "MYR"
	PriceOneTimePriceCurrencyMzn PriceOneTimePriceCurrency = "MZN"
	PriceOneTimePriceCurrencyNad PriceOneTimePriceCurrency = "NAD"
	PriceOneTimePriceCurrencyNgn PriceOneTimePriceCurrency = "NGN"
	PriceOneTimePriceCurrencyNio PriceOneTimePriceCurrency = "NIO"
	PriceOneTimePriceCurrencyNok PriceOneTimePriceCurrency = "NOK"
	PriceOneTimePriceCurrencyNpr PriceOneTimePriceCurrency = "NPR"
	PriceOneTimePriceCurrencyNzd PriceOneTimePriceCurrency = "NZD"
	PriceOneTimePriceCurrencyOmr PriceOneTimePriceCurrency = "OMR"
	PriceOneTimePriceCurrencyPab PriceOneTimePriceCurrency = "PAB"
	PriceOneTimePriceCurrencyPen PriceOneTimePriceCurrency = "PEN"
	PriceOneTimePriceCurrencyPgk PriceOneTimePriceCurrency = "PGK"
	PriceOneTimePriceCurrencyPhp PriceOneTimePriceCurrency = "PHP"
	PriceOneTimePriceCurrencyPkr PriceOneTimePriceCurrency = "PKR"
	PriceOneTimePriceCurrencyPln PriceOneTimePriceCurrency = "PLN"
	PriceOneTimePriceCurrencyPyg PriceOneTimePriceCurrency = "PYG"
	PriceOneTimePriceCurrencyQar PriceOneTimePriceCurrency = "QAR"
	PriceOneTimePriceCurrencyRon PriceOneTimePriceCurrency = "RON"
	PriceOneTimePriceCurrencyRsd PriceOneTimePriceCurrency = "RSD"
	PriceOneTimePriceCurrencyRub PriceOneTimePriceCurrency = "RUB"
	PriceOneTimePriceCurrencyRwf PriceOneTimePriceCurrency = "RWF"
	PriceOneTimePriceCurrencySar PriceOneTimePriceCurrency = "SAR"
	PriceOneTimePriceCurrencySbd PriceOneTimePriceCurrency = "SBD"
	PriceOneTimePriceCurrencyScr PriceOneTimePriceCurrency = "SCR"
	PriceOneTimePriceCurrencySek PriceOneTimePriceCurrency = "SEK"
	PriceOneTimePriceCurrencySgd PriceOneTimePriceCurrency = "SGD"
	PriceOneTimePriceCurrencyShp PriceOneTimePriceCurrency = "SHP"
	PriceOneTimePriceCurrencySle PriceOneTimePriceCurrency = "SLE"
	PriceOneTimePriceCurrencySll PriceOneTimePriceCurrency = "SLL"
	PriceOneTimePriceCurrencySos PriceOneTimePriceCurrency = "SOS"
	PriceOneTimePriceCurrencySrd PriceOneTimePriceCurrency = "SRD"
	PriceOneTimePriceCurrencySsp PriceOneTimePriceCurrency = "SSP"
	PriceOneTimePriceCurrencyStn PriceOneTimePriceCurrency = "STN"
	PriceOneTimePriceCurrencySvc PriceOneTimePriceCurrency = "SVC"
	PriceOneTimePriceCurrencySzl PriceOneTimePriceCurrency = "SZL"
	PriceOneTimePriceCurrencyThb PriceOneTimePriceCurrency = "THB"
	PriceOneTimePriceCurrencyTnd PriceOneTimePriceCurrency = "TND"
	PriceOneTimePriceCurrencyTop PriceOneTimePriceCurrency = "TOP"
	PriceOneTimePriceCurrencyTry PriceOneTimePriceCurrency = "TRY"
	PriceOneTimePriceCurrencyTtd PriceOneTimePriceCurrency = "TTD"
	PriceOneTimePriceCurrencyTwd PriceOneTimePriceCurrency = "TWD"
	PriceOneTimePriceCurrencyTzs PriceOneTimePriceCurrency = "TZS"
	PriceOneTimePriceCurrencyUah PriceOneTimePriceCurrency = "UAH"
	PriceOneTimePriceCurrencyUgx PriceOneTimePriceCurrency = "UGX"
	PriceOneTimePriceCurrencyUsd PriceOneTimePriceCurrency = "USD"
	PriceOneTimePriceCurrencyUyu PriceOneTimePriceCurrency = "UYU"
	PriceOneTimePriceCurrencyUzs PriceOneTimePriceCurrency = "UZS"
	PriceOneTimePriceCurrencyVes PriceOneTimePriceCurrency = "VES"
	PriceOneTimePriceCurrencyVnd PriceOneTimePriceCurrency = "VND"
	PriceOneTimePriceCurrencyVuv PriceOneTimePriceCurrency = "VUV"
	PriceOneTimePriceCurrencyWst PriceOneTimePriceCurrency = "WST"
	PriceOneTimePriceCurrencyXaf PriceOneTimePriceCurrency = "XAF"
	PriceOneTimePriceCurrencyXcd PriceOneTimePriceCurrency = "XCD"
	PriceOneTimePriceCurrencyXof PriceOneTimePriceCurrency = "XOF"
	PriceOneTimePriceCurrencyXpf PriceOneTimePriceCurrency = "XPF"
	PriceOneTimePriceCurrencyYer PriceOneTimePriceCurrency = "YER"
	PriceOneTimePriceCurrencyZar PriceOneTimePriceCurrency = "ZAR"
	PriceOneTimePriceCurrencyZmw PriceOneTimePriceCurrency = "ZMW"
)

func (r PriceOneTimePriceCurrency) IsKnown() bool {
	switch r {
	case PriceOneTimePriceCurrencyAed, PriceOneTimePriceCurrencyAll, PriceOneTimePriceCurrencyAmd, PriceOneTimePriceCurrencyAng, PriceOneTimePriceCurrencyAoa, PriceOneTimePriceCurrencyArs, PriceOneTimePriceCurrencyAud, PriceOneTimePriceCurrencyAwg, PriceOneTimePriceCurrencyAzn, PriceOneTimePriceCurrencyBam, PriceOneTimePriceCurrencyBbd, PriceOneTimePriceCurrencyBdt, PriceOneTimePriceCurrencyBgn, PriceOneTimePriceCurrencyBhd, PriceOneTimePriceCurrencyBif, PriceOneTimePriceCurrencyBmd, PriceOneTimePriceCurrencyBnd, PriceOneTimePriceCurrencyBob, PriceOneTimePriceCurrencyBrl, PriceOneTimePriceCurrencyBsd, PriceOneTimePriceCurrencyBwp, PriceOneTimePriceCurrencyByn, PriceOneTimePriceCurrencyBzd, PriceOneTimePriceCurrencyCad, PriceOneTimePriceCurrencyChf, PriceOneTimePriceCurrencyClp, PriceOneTimePriceCurrencyCny, PriceOneTimePriceCurrencyCop, PriceOneTimePriceCurrencyCrc, PriceOneTimePriceCurrencyCup, PriceOneTimePriceCurrencyCve, PriceOneTimePriceCurrencyCzk, PriceOneTimePriceCurrencyDjf, PriceOneTimePriceCurrencyDkk, PriceOneTimePriceCurrencyDop, PriceOneTimePriceCurrencyDzd, PriceOneTimePriceCurrencyEgp, PriceOneTimePriceCurrencyEtb, PriceOneTimePriceCurrencyEur, PriceOneTimePriceCurrencyFjd, PriceOneTimePriceCurrencyFkp, PriceOneTimePriceCurrencyGbp, PriceOneTimePriceCurrencyGel, PriceOneTimePriceCurrencyGhs, PriceOneTimePriceCurrencyGip, PriceOneTimePriceCurrencyGmd, PriceOneTimePriceCurrencyGnf, PriceOneTimePriceCurrencyGtq, PriceOneTimePriceCurrencyGyd, PriceOneTimePriceCurrencyHkd, PriceOneTimePriceCurrencyHnl, PriceOneTimePriceCurrencyHrk, PriceOneTimePriceCurrencyHtg, PriceOneTimePriceCurrencyHuf, PriceOneTimePriceCurrencyIdr, PriceOneTimePriceCurrencyIls, PriceOneTimePriceCurrencyInr, PriceOneTimePriceCurrencyIqd, PriceOneTimePriceCurrencyJmd, PriceOneTimePriceCurrencyJod, PriceOneTimePriceCurrencyJpy, PriceOneTimePriceCurrencyKes, PriceOneTimePriceCurrencyKgs, PriceOneTimePriceCurrencyKhr, PriceOneTimePriceCurrencyKmf, PriceOneTimePriceCurrencyKrw, PriceOneTimePriceCurrencyKwd, PriceOneTimePriceCurrencyKyd, PriceOneTimePriceCurrencyKzt, PriceOneTimePriceCurrencyLak, PriceOneTimePriceCurrencyLbp, PriceOneTimePriceCurrencyLkr, PriceOneTimePriceCurrencyLrd, PriceOneTimePriceCurrencyLsl, PriceOneTimePriceCurrencyLyd, PriceOneTimePriceCurrencyMad, PriceOneTimePriceCurrencyMdl, PriceOneTimePriceCurrencyMga, PriceOneTimePriceCurrencyMkd, PriceOneTimePriceCurrencyMmk, PriceOneTimePriceCurrencyMnt, PriceOneTimePriceCurrencyMop, PriceOneTimePriceCurrencyMru, PriceOneTimePriceCurrencyMur, PriceOneTimePriceCurrencyMvr, PriceOneTimePriceCurrencyMwk, PriceOneTimePriceCurrencyMxn, PriceOneTimePriceCurrencyMyr, PriceOneTimePriceCurrencyMzn, PriceOneTimePriceCurrencyNad, PriceOneTimePriceCurrencyNgn, PriceOneTimePriceCurrencyNio, PriceOneTimePriceCurrencyNok, PriceOneTimePriceCurrencyNpr, PriceOneTimePriceCurrencyNzd, PriceOneTimePriceCurrencyOmr, PriceOneTimePriceCurrencyPab, PriceOneTimePriceCurrencyPen, PriceOneTimePriceCurrencyPgk, PriceOneTimePriceCurrencyPhp, PriceOneTimePriceCurrencyPkr, PriceOneTimePriceCurrencyPln, PriceOneTimePriceCurrencyPyg, PriceOneTimePriceCurrencyQar, PriceOneTimePriceCurrencyRon, PriceOneTimePriceCurrencyRsd, PriceOneTimePriceCurrencyRub, PriceOneTimePriceCurrencyRwf, PriceOneTimePriceCurrencySar, PriceOneTimePriceCurrencySbd, PriceOneTimePriceCurrencyScr, PriceOneTimePriceCurrencySek, PriceOneTimePriceCurrencySgd, PriceOneTimePriceCurrencyShp, PriceOneTimePriceCurrencySle, PriceOneTimePriceCurrencySll, PriceOneTimePriceCurrencySos, PriceOneTimePriceCurrencySrd, PriceOneTimePriceCurrencySsp, PriceOneTimePriceCurrencyStn, PriceOneTimePriceCurrencySvc, PriceOneTimePriceCurrencySzl, PriceOneTimePriceCurrencyThb, PriceOneTimePriceCurrencyTnd, PriceOneTimePriceCurrencyTop, PriceOneTimePriceCurrencyTry, PriceOneTimePriceCurrencyTtd, PriceOneTimePriceCurrencyTwd, PriceOneTimePriceCurrencyTzs, PriceOneTimePriceCurrencyUah, PriceOneTimePriceCurrencyUgx, PriceOneTimePriceCurrencyUsd, PriceOneTimePriceCurrencyUyu, PriceOneTimePriceCurrencyUzs, PriceOneTimePriceCurrencyVes, PriceOneTimePriceCurrencyVnd, PriceOneTimePriceCurrencyVuv, PriceOneTimePriceCurrencyWst, PriceOneTimePriceCurrencyXaf, PriceOneTimePriceCurrencyXcd, PriceOneTimePriceCurrencyXof, PriceOneTimePriceCurrencyXpf, PriceOneTimePriceCurrencyYer, PriceOneTimePriceCurrencyZar, PriceOneTimePriceCurrencyZmw:
		return true
	}
	return false
}

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
	Currency PriceRecurringPriceCurrency `json:"currency,required"`
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

type PriceRecurringPriceCurrency string

const (
	PriceRecurringPriceCurrencyAed PriceRecurringPriceCurrency = "AED"
	PriceRecurringPriceCurrencyAll PriceRecurringPriceCurrency = "ALL"
	PriceRecurringPriceCurrencyAmd PriceRecurringPriceCurrency = "AMD"
	PriceRecurringPriceCurrencyAng PriceRecurringPriceCurrency = "ANG"
	PriceRecurringPriceCurrencyAoa PriceRecurringPriceCurrency = "AOA"
	PriceRecurringPriceCurrencyArs PriceRecurringPriceCurrency = "ARS"
	PriceRecurringPriceCurrencyAud PriceRecurringPriceCurrency = "AUD"
	PriceRecurringPriceCurrencyAwg PriceRecurringPriceCurrency = "AWG"
	PriceRecurringPriceCurrencyAzn PriceRecurringPriceCurrency = "AZN"
	PriceRecurringPriceCurrencyBam PriceRecurringPriceCurrency = "BAM"
	PriceRecurringPriceCurrencyBbd PriceRecurringPriceCurrency = "BBD"
	PriceRecurringPriceCurrencyBdt PriceRecurringPriceCurrency = "BDT"
	PriceRecurringPriceCurrencyBgn PriceRecurringPriceCurrency = "BGN"
	PriceRecurringPriceCurrencyBhd PriceRecurringPriceCurrency = "BHD"
	PriceRecurringPriceCurrencyBif PriceRecurringPriceCurrency = "BIF"
	PriceRecurringPriceCurrencyBmd PriceRecurringPriceCurrency = "BMD"
	PriceRecurringPriceCurrencyBnd PriceRecurringPriceCurrency = "BND"
	PriceRecurringPriceCurrencyBob PriceRecurringPriceCurrency = "BOB"
	PriceRecurringPriceCurrencyBrl PriceRecurringPriceCurrency = "BRL"
	PriceRecurringPriceCurrencyBsd PriceRecurringPriceCurrency = "BSD"
	PriceRecurringPriceCurrencyBwp PriceRecurringPriceCurrency = "BWP"
	PriceRecurringPriceCurrencyByn PriceRecurringPriceCurrency = "BYN"
	PriceRecurringPriceCurrencyBzd PriceRecurringPriceCurrency = "BZD"
	PriceRecurringPriceCurrencyCad PriceRecurringPriceCurrency = "CAD"
	PriceRecurringPriceCurrencyChf PriceRecurringPriceCurrency = "CHF"
	PriceRecurringPriceCurrencyClp PriceRecurringPriceCurrency = "CLP"
	PriceRecurringPriceCurrencyCny PriceRecurringPriceCurrency = "CNY"
	PriceRecurringPriceCurrencyCop PriceRecurringPriceCurrency = "COP"
	PriceRecurringPriceCurrencyCrc PriceRecurringPriceCurrency = "CRC"
	PriceRecurringPriceCurrencyCup PriceRecurringPriceCurrency = "CUP"
	PriceRecurringPriceCurrencyCve PriceRecurringPriceCurrency = "CVE"
	PriceRecurringPriceCurrencyCzk PriceRecurringPriceCurrency = "CZK"
	PriceRecurringPriceCurrencyDjf PriceRecurringPriceCurrency = "DJF"
	PriceRecurringPriceCurrencyDkk PriceRecurringPriceCurrency = "DKK"
	PriceRecurringPriceCurrencyDop PriceRecurringPriceCurrency = "DOP"
	PriceRecurringPriceCurrencyDzd PriceRecurringPriceCurrency = "DZD"
	PriceRecurringPriceCurrencyEgp PriceRecurringPriceCurrency = "EGP"
	PriceRecurringPriceCurrencyEtb PriceRecurringPriceCurrency = "ETB"
	PriceRecurringPriceCurrencyEur PriceRecurringPriceCurrency = "EUR"
	PriceRecurringPriceCurrencyFjd PriceRecurringPriceCurrency = "FJD"
	PriceRecurringPriceCurrencyFkp PriceRecurringPriceCurrency = "FKP"
	PriceRecurringPriceCurrencyGbp PriceRecurringPriceCurrency = "GBP"
	PriceRecurringPriceCurrencyGel PriceRecurringPriceCurrency = "GEL"
	PriceRecurringPriceCurrencyGhs PriceRecurringPriceCurrency = "GHS"
	PriceRecurringPriceCurrencyGip PriceRecurringPriceCurrency = "GIP"
	PriceRecurringPriceCurrencyGmd PriceRecurringPriceCurrency = "GMD"
	PriceRecurringPriceCurrencyGnf PriceRecurringPriceCurrency = "GNF"
	PriceRecurringPriceCurrencyGtq PriceRecurringPriceCurrency = "GTQ"
	PriceRecurringPriceCurrencyGyd PriceRecurringPriceCurrency = "GYD"
	PriceRecurringPriceCurrencyHkd PriceRecurringPriceCurrency = "HKD"
	PriceRecurringPriceCurrencyHnl PriceRecurringPriceCurrency = "HNL"
	PriceRecurringPriceCurrencyHrk PriceRecurringPriceCurrency = "HRK"
	PriceRecurringPriceCurrencyHtg PriceRecurringPriceCurrency = "HTG"
	PriceRecurringPriceCurrencyHuf PriceRecurringPriceCurrency = "HUF"
	PriceRecurringPriceCurrencyIdr PriceRecurringPriceCurrency = "IDR"
	PriceRecurringPriceCurrencyIls PriceRecurringPriceCurrency = "ILS"
	PriceRecurringPriceCurrencyInr PriceRecurringPriceCurrency = "INR"
	PriceRecurringPriceCurrencyIqd PriceRecurringPriceCurrency = "IQD"
	PriceRecurringPriceCurrencyJmd PriceRecurringPriceCurrency = "JMD"
	PriceRecurringPriceCurrencyJod PriceRecurringPriceCurrency = "JOD"
	PriceRecurringPriceCurrencyJpy PriceRecurringPriceCurrency = "JPY"
	PriceRecurringPriceCurrencyKes PriceRecurringPriceCurrency = "KES"
	PriceRecurringPriceCurrencyKgs PriceRecurringPriceCurrency = "KGS"
	PriceRecurringPriceCurrencyKhr PriceRecurringPriceCurrency = "KHR"
	PriceRecurringPriceCurrencyKmf PriceRecurringPriceCurrency = "KMF"
	PriceRecurringPriceCurrencyKrw PriceRecurringPriceCurrency = "KRW"
	PriceRecurringPriceCurrencyKwd PriceRecurringPriceCurrency = "KWD"
	PriceRecurringPriceCurrencyKyd PriceRecurringPriceCurrency = "KYD"
	PriceRecurringPriceCurrencyKzt PriceRecurringPriceCurrency = "KZT"
	PriceRecurringPriceCurrencyLak PriceRecurringPriceCurrency = "LAK"
	PriceRecurringPriceCurrencyLbp PriceRecurringPriceCurrency = "LBP"
	PriceRecurringPriceCurrencyLkr PriceRecurringPriceCurrency = "LKR"
	PriceRecurringPriceCurrencyLrd PriceRecurringPriceCurrency = "LRD"
	PriceRecurringPriceCurrencyLsl PriceRecurringPriceCurrency = "LSL"
	PriceRecurringPriceCurrencyLyd PriceRecurringPriceCurrency = "LYD"
	PriceRecurringPriceCurrencyMad PriceRecurringPriceCurrency = "MAD"
	PriceRecurringPriceCurrencyMdl PriceRecurringPriceCurrency = "MDL"
	PriceRecurringPriceCurrencyMga PriceRecurringPriceCurrency = "MGA"
	PriceRecurringPriceCurrencyMkd PriceRecurringPriceCurrency = "MKD"
	PriceRecurringPriceCurrencyMmk PriceRecurringPriceCurrency = "MMK"
	PriceRecurringPriceCurrencyMnt PriceRecurringPriceCurrency = "MNT"
	PriceRecurringPriceCurrencyMop PriceRecurringPriceCurrency = "MOP"
	PriceRecurringPriceCurrencyMru PriceRecurringPriceCurrency = "MRU"
	PriceRecurringPriceCurrencyMur PriceRecurringPriceCurrency = "MUR"
	PriceRecurringPriceCurrencyMvr PriceRecurringPriceCurrency = "MVR"
	PriceRecurringPriceCurrencyMwk PriceRecurringPriceCurrency = "MWK"
	PriceRecurringPriceCurrencyMxn PriceRecurringPriceCurrency = "MXN"
	PriceRecurringPriceCurrencyMyr PriceRecurringPriceCurrency = "MYR"
	PriceRecurringPriceCurrencyMzn PriceRecurringPriceCurrency = "MZN"
	PriceRecurringPriceCurrencyNad PriceRecurringPriceCurrency = "NAD"
	PriceRecurringPriceCurrencyNgn PriceRecurringPriceCurrency = "NGN"
	PriceRecurringPriceCurrencyNio PriceRecurringPriceCurrency = "NIO"
	PriceRecurringPriceCurrencyNok PriceRecurringPriceCurrency = "NOK"
	PriceRecurringPriceCurrencyNpr PriceRecurringPriceCurrency = "NPR"
	PriceRecurringPriceCurrencyNzd PriceRecurringPriceCurrency = "NZD"
	PriceRecurringPriceCurrencyOmr PriceRecurringPriceCurrency = "OMR"
	PriceRecurringPriceCurrencyPab PriceRecurringPriceCurrency = "PAB"
	PriceRecurringPriceCurrencyPen PriceRecurringPriceCurrency = "PEN"
	PriceRecurringPriceCurrencyPgk PriceRecurringPriceCurrency = "PGK"
	PriceRecurringPriceCurrencyPhp PriceRecurringPriceCurrency = "PHP"
	PriceRecurringPriceCurrencyPkr PriceRecurringPriceCurrency = "PKR"
	PriceRecurringPriceCurrencyPln PriceRecurringPriceCurrency = "PLN"
	PriceRecurringPriceCurrencyPyg PriceRecurringPriceCurrency = "PYG"
	PriceRecurringPriceCurrencyQar PriceRecurringPriceCurrency = "QAR"
	PriceRecurringPriceCurrencyRon PriceRecurringPriceCurrency = "RON"
	PriceRecurringPriceCurrencyRsd PriceRecurringPriceCurrency = "RSD"
	PriceRecurringPriceCurrencyRub PriceRecurringPriceCurrency = "RUB"
	PriceRecurringPriceCurrencyRwf PriceRecurringPriceCurrency = "RWF"
	PriceRecurringPriceCurrencySar PriceRecurringPriceCurrency = "SAR"
	PriceRecurringPriceCurrencySbd PriceRecurringPriceCurrency = "SBD"
	PriceRecurringPriceCurrencyScr PriceRecurringPriceCurrency = "SCR"
	PriceRecurringPriceCurrencySek PriceRecurringPriceCurrency = "SEK"
	PriceRecurringPriceCurrencySgd PriceRecurringPriceCurrency = "SGD"
	PriceRecurringPriceCurrencyShp PriceRecurringPriceCurrency = "SHP"
	PriceRecurringPriceCurrencySle PriceRecurringPriceCurrency = "SLE"
	PriceRecurringPriceCurrencySll PriceRecurringPriceCurrency = "SLL"
	PriceRecurringPriceCurrencySos PriceRecurringPriceCurrency = "SOS"
	PriceRecurringPriceCurrencySrd PriceRecurringPriceCurrency = "SRD"
	PriceRecurringPriceCurrencySsp PriceRecurringPriceCurrency = "SSP"
	PriceRecurringPriceCurrencyStn PriceRecurringPriceCurrency = "STN"
	PriceRecurringPriceCurrencySvc PriceRecurringPriceCurrency = "SVC"
	PriceRecurringPriceCurrencySzl PriceRecurringPriceCurrency = "SZL"
	PriceRecurringPriceCurrencyThb PriceRecurringPriceCurrency = "THB"
	PriceRecurringPriceCurrencyTnd PriceRecurringPriceCurrency = "TND"
	PriceRecurringPriceCurrencyTop PriceRecurringPriceCurrency = "TOP"
	PriceRecurringPriceCurrencyTry PriceRecurringPriceCurrency = "TRY"
	PriceRecurringPriceCurrencyTtd PriceRecurringPriceCurrency = "TTD"
	PriceRecurringPriceCurrencyTwd PriceRecurringPriceCurrency = "TWD"
	PriceRecurringPriceCurrencyTzs PriceRecurringPriceCurrency = "TZS"
	PriceRecurringPriceCurrencyUah PriceRecurringPriceCurrency = "UAH"
	PriceRecurringPriceCurrencyUgx PriceRecurringPriceCurrency = "UGX"
	PriceRecurringPriceCurrencyUsd PriceRecurringPriceCurrency = "USD"
	PriceRecurringPriceCurrencyUyu PriceRecurringPriceCurrency = "UYU"
	PriceRecurringPriceCurrencyUzs PriceRecurringPriceCurrency = "UZS"
	PriceRecurringPriceCurrencyVes PriceRecurringPriceCurrency = "VES"
	PriceRecurringPriceCurrencyVnd PriceRecurringPriceCurrency = "VND"
	PriceRecurringPriceCurrencyVuv PriceRecurringPriceCurrency = "VUV"
	PriceRecurringPriceCurrencyWst PriceRecurringPriceCurrency = "WST"
	PriceRecurringPriceCurrencyXaf PriceRecurringPriceCurrency = "XAF"
	PriceRecurringPriceCurrencyXcd PriceRecurringPriceCurrency = "XCD"
	PriceRecurringPriceCurrencyXof PriceRecurringPriceCurrency = "XOF"
	PriceRecurringPriceCurrencyXpf PriceRecurringPriceCurrency = "XPF"
	PriceRecurringPriceCurrencyYer PriceRecurringPriceCurrency = "YER"
	PriceRecurringPriceCurrencyZar PriceRecurringPriceCurrency = "ZAR"
	PriceRecurringPriceCurrencyZmw PriceRecurringPriceCurrency = "ZMW"
)

func (r PriceRecurringPriceCurrency) IsKnown() bool {
	switch r {
	case PriceRecurringPriceCurrencyAed, PriceRecurringPriceCurrencyAll, PriceRecurringPriceCurrencyAmd, PriceRecurringPriceCurrencyAng, PriceRecurringPriceCurrencyAoa, PriceRecurringPriceCurrencyArs, PriceRecurringPriceCurrencyAud, PriceRecurringPriceCurrencyAwg, PriceRecurringPriceCurrencyAzn, PriceRecurringPriceCurrencyBam, PriceRecurringPriceCurrencyBbd, PriceRecurringPriceCurrencyBdt, PriceRecurringPriceCurrencyBgn, PriceRecurringPriceCurrencyBhd, PriceRecurringPriceCurrencyBif, PriceRecurringPriceCurrencyBmd, PriceRecurringPriceCurrencyBnd, PriceRecurringPriceCurrencyBob, PriceRecurringPriceCurrencyBrl, PriceRecurringPriceCurrencyBsd, PriceRecurringPriceCurrencyBwp, PriceRecurringPriceCurrencyByn, PriceRecurringPriceCurrencyBzd, PriceRecurringPriceCurrencyCad, PriceRecurringPriceCurrencyChf, PriceRecurringPriceCurrencyClp, PriceRecurringPriceCurrencyCny, PriceRecurringPriceCurrencyCop, PriceRecurringPriceCurrencyCrc, PriceRecurringPriceCurrencyCup, PriceRecurringPriceCurrencyCve, PriceRecurringPriceCurrencyCzk, PriceRecurringPriceCurrencyDjf, PriceRecurringPriceCurrencyDkk, PriceRecurringPriceCurrencyDop, PriceRecurringPriceCurrencyDzd, PriceRecurringPriceCurrencyEgp, PriceRecurringPriceCurrencyEtb, PriceRecurringPriceCurrencyEur, PriceRecurringPriceCurrencyFjd, PriceRecurringPriceCurrencyFkp, PriceRecurringPriceCurrencyGbp, PriceRecurringPriceCurrencyGel, PriceRecurringPriceCurrencyGhs, PriceRecurringPriceCurrencyGip, PriceRecurringPriceCurrencyGmd, PriceRecurringPriceCurrencyGnf, PriceRecurringPriceCurrencyGtq, PriceRecurringPriceCurrencyGyd, PriceRecurringPriceCurrencyHkd, PriceRecurringPriceCurrencyHnl, PriceRecurringPriceCurrencyHrk, PriceRecurringPriceCurrencyHtg, PriceRecurringPriceCurrencyHuf, PriceRecurringPriceCurrencyIdr, PriceRecurringPriceCurrencyIls, PriceRecurringPriceCurrencyInr, PriceRecurringPriceCurrencyIqd, PriceRecurringPriceCurrencyJmd, PriceRecurringPriceCurrencyJod, PriceRecurringPriceCurrencyJpy, PriceRecurringPriceCurrencyKes, PriceRecurringPriceCurrencyKgs, PriceRecurringPriceCurrencyKhr, PriceRecurringPriceCurrencyKmf, PriceRecurringPriceCurrencyKrw, PriceRecurringPriceCurrencyKwd, PriceRecurringPriceCurrencyKyd, PriceRecurringPriceCurrencyKzt, PriceRecurringPriceCurrencyLak, PriceRecurringPriceCurrencyLbp, PriceRecurringPriceCurrencyLkr, PriceRecurringPriceCurrencyLrd, PriceRecurringPriceCurrencyLsl, PriceRecurringPriceCurrencyLyd, PriceRecurringPriceCurrencyMad, PriceRecurringPriceCurrencyMdl, PriceRecurringPriceCurrencyMga, PriceRecurringPriceCurrencyMkd, PriceRecurringPriceCurrencyMmk, PriceRecurringPriceCurrencyMnt, PriceRecurringPriceCurrencyMop, PriceRecurringPriceCurrencyMru, PriceRecurringPriceCurrencyMur, PriceRecurringPriceCurrencyMvr, PriceRecurringPriceCurrencyMwk, PriceRecurringPriceCurrencyMxn, PriceRecurringPriceCurrencyMyr, PriceRecurringPriceCurrencyMzn, PriceRecurringPriceCurrencyNad, PriceRecurringPriceCurrencyNgn, PriceRecurringPriceCurrencyNio, PriceRecurringPriceCurrencyNok, PriceRecurringPriceCurrencyNpr, PriceRecurringPriceCurrencyNzd, PriceRecurringPriceCurrencyOmr, PriceRecurringPriceCurrencyPab, PriceRecurringPriceCurrencyPen, PriceRecurringPriceCurrencyPgk, PriceRecurringPriceCurrencyPhp, PriceRecurringPriceCurrencyPkr, PriceRecurringPriceCurrencyPln, PriceRecurringPriceCurrencyPyg, PriceRecurringPriceCurrencyQar, PriceRecurringPriceCurrencyRon, PriceRecurringPriceCurrencyRsd, PriceRecurringPriceCurrencyRub, PriceRecurringPriceCurrencyRwf, PriceRecurringPriceCurrencySar, PriceRecurringPriceCurrencySbd, PriceRecurringPriceCurrencyScr, PriceRecurringPriceCurrencySek, PriceRecurringPriceCurrencySgd, PriceRecurringPriceCurrencyShp, PriceRecurringPriceCurrencySle, PriceRecurringPriceCurrencySll, PriceRecurringPriceCurrencySos, PriceRecurringPriceCurrencySrd, PriceRecurringPriceCurrencySsp, PriceRecurringPriceCurrencyStn, PriceRecurringPriceCurrencySvc, PriceRecurringPriceCurrencySzl, PriceRecurringPriceCurrencyThb, PriceRecurringPriceCurrencyTnd, PriceRecurringPriceCurrencyTop, PriceRecurringPriceCurrencyTry, PriceRecurringPriceCurrencyTtd, PriceRecurringPriceCurrencyTwd, PriceRecurringPriceCurrencyTzs, PriceRecurringPriceCurrencyUah, PriceRecurringPriceCurrencyUgx, PriceRecurringPriceCurrencyUsd, PriceRecurringPriceCurrencyUyu, PriceRecurringPriceCurrencyUzs, PriceRecurringPriceCurrencyVes, PriceRecurringPriceCurrencyVnd, PriceRecurringPriceCurrencyVuv, PriceRecurringPriceCurrencyWst, PriceRecurringPriceCurrencyXaf, PriceRecurringPriceCurrencyXcd, PriceRecurringPriceCurrencyXof, PriceRecurringPriceCurrencyXpf, PriceRecurringPriceCurrencyYer, PriceRecurringPriceCurrencyZar, PriceRecurringPriceCurrencyZmw:
		return true
	}
	return false
}

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

type PriceCurrency string

const (
	PriceCurrencyAed PriceCurrency = "AED"
	PriceCurrencyAll PriceCurrency = "ALL"
	PriceCurrencyAmd PriceCurrency = "AMD"
	PriceCurrencyAng PriceCurrency = "ANG"
	PriceCurrencyAoa PriceCurrency = "AOA"
	PriceCurrencyArs PriceCurrency = "ARS"
	PriceCurrencyAud PriceCurrency = "AUD"
	PriceCurrencyAwg PriceCurrency = "AWG"
	PriceCurrencyAzn PriceCurrency = "AZN"
	PriceCurrencyBam PriceCurrency = "BAM"
	PriceCurrencyBbd PriceCurrency = "BBD"
	PriceCurrencyBdt PriceCurrency = "BDT"
	PriceCurrencyBgn PriceCurrency = "BGN"
	PriceCurrencyBhd PriceCurrency = "BHD"
	PriceCurrencyBif PriceCurrency = "BIF"
	PriceCurrencyBmd PriceCurrency = "BMD"
	PriceCurrencyBnd PriceCurrency = "BND"
	PriceCurrencyBob PriceCurrency = "BOB"
	PriceCurrencyBrl PriceCurrency = "BRL"
	PriceCurrencyBsd PriceCurrency = "BSD"
	PriceCurrencyBwp PriceCurrency = "BWP"
	PriceCurrencyByn PriceCurrency = "BYN"
	PriceCurrencyBzd PriceCurrency = "BZD"
	PriceCurrencyCad PriceCurrency = "CAD"
	PriceCurrencyChf PriceCurrency = "CHF"
	PriceCurrencyClp PriceCurrency = "CLP"
	PriceCurrencyCny PriceCurrency = "CNY"
	PriceCurrencyCop PriceCurrency = "COP"
	PriceCurrencyCrc PriceCurrency = "CRC"
	PriceCurrencyCup PriceCurrency = "CUP"
	PriceCurrencyCve PriceCurrency = "CVE"
	PriceCurrencyCzk PriceCurrency = "CZK"
	PriceCurrencyDjf PriceCurrency = "DJF"
	PriceCurrencyDkk PriceCurrency = "DKK"
	PriceCurrencyDop PriceCurrency = "DOP"
	PriceCurrencyDzd PriceCurrency = "DZD"
	PriceCurrencyEgp PriceCurrency = "EGP"
	PriceCurrencyEtb PriceCurrency = "ETB"
	PriceCurrencyEur PriceCurrency = "EUR"
	PriceCurrencyFjd PriceCurrency = "FJD"
	PriceCurrencyFkp PriceCurrency = "FKP"
	PriceCurrencyGbp PriceCurrency = "GBP"
	PriceCurrencyGel PriceCurrency = "GEL"
	PriceCurrencyGhs PriceCurrency = "GHS"
	PriceCurrencyGip PriceCurrency = "GIP"
	PriceCurrencyGmd PriceCurrency = "GMD"
	PriceCurrencyGnf PriceCurrency = "GNF"
	PriceCurrencyGtq PriceCurrency = "GTQ"
	PriceCurrencyGyd PriceCurrency = "GYD"
	PriceCurrencyHkd PriceCurrency = "HKD"
	PriceCurrencyHnl PriceCurrency = "HNL"
	PriceCurrencyHrk PriceCurrency = "HRK"
	PriceCurrencyHtg PriceCurrency = "HTG"
	PriceCurrencyHuf PriceCurrency = "HUF"
	PriceCurrencyIdr PriceCurrency = "IDR"
	PriceCurrencyIls PriceCurrency = "ILS"
	PriceCurrencyInr PriceCurrency = "INR"
	PriceCurrencyIqd PriceCurrency = "IQD"
	PriceCurrencyJmd PriceCurrency = "JMD"
	PriceCurrencyJod PriceCurrency = "JOD"
	PriceCurrencyJpy PriceCurrency = "JPY"
	PriceCurrencyKes PriceCurrency = "KES"
	PriceCurrencyKgs PriceCurrency = "KGS"
	PriceCurrencyKhr PriceCurrency = "KHR"
	PriceCurrencyKmf PriceCurrency = "KMF"
	PriceCurrencyKrw PriceCurrency = "KRW"
	PriceCurrencyKwd PriceCurrency = "KWD"
	PriceCurrencyKyd PriceCurrency = "KYD"
	PriceCurrencyKzt PriceCurrency = "KZT"
	PriceCurrencyLak PriceCurrency = "LAK"
	PriceCurrencyLbp PriceCurrency = "LBP"
	PriceCurrencyLkr PriceCurrency = "LKR"
	PriceCurrencyLrd PriceCurrency = "LRD"
	PriceCurrencyLsl PriceCurrency = "LSL"
	PriceCurrencyLyd PriceCurrency = "LYD"
	PriceCurrencyMad PriceCurrency = "MAD"
	PriceCurrencyMdl PriceCurrency = "MDL"
	PriceCurrencyMga PriceCurrency = "MGA"
	PriceCurrencyMkd PriceCurrency = "MKD"
	PriceCurrencyMmk PriceCurrency = "MMK"
	PriceCurrencyMnt PriceCurrency = "MNT"
	PriceCurrencyMop PriceCurrency = "MOP"
	PriceCurrencyMru PriceCurrency = "MRU"
	PriceCurrencyMur PriceCurrency = "MUR"
	PriceCurrencyMvr PriceCurrency = "MVR"
	PriceCurrencyMwk PriceCurrency = "MWK"
	PriceCurrencyMxn PriceCurrency = "MXN"
	PriceCurrencyMyr PriceCurrency = "MYR"
	PriceCurrencyMzn PriceCurrency = "MZN"
	PriceCurrencyNad PriceCurrency = "NAD"
	PriceCurrencyNgn PriceCurrency = "NGN"
	PriceCurrencyNio PriceCurrency = "NIO"
	PriceCurrencyNok PriceCurrency = "NOK"
	PriceCurrencyNpr PriceCurrency = "NPR"
	PriceCurrencyNzd PriceCurrency = "NZD"
	PriceCurrencyOmr PriceCurrency = "OMR"
	PriceCurrencyPab PriceCurrency = "PAB"
	PriceCurrencyPen PriceCurrency = "PEN"
	PriceCurrencyPgk PriceCurrency = "PGK"
	PriceCurrencyPhp PriceCurrency = "PHP"
	PriceCurrencyPkr PriceCurrency = "PKR"
	PriceCurrencyPln PriceCurrency = "PLN"
	PriceCurrencyPyg PriceCurrency = "PYG"
	PriceCurrencyQar PriceCurrency = "QAR"
	PriceCurrencyRon PriceCurrency = "RON"
	PriceCurrencyRsd PriceCurrency = "RSD"
	PriceCurrencyRub PriceCurrency = "RUB"
	PriceCurrencyRwf PriceCurrency = "RWF"
	PriceCurrencySar PriceCurrency = "SAR"
	PriceCurrencySbd PriceCurrency = "SBD"
	PriceCurrencyScr PriceCurrency = "SCR"
	PriceCurrencySek PriceCurrency = "SEK"
	PriceCurrencySgd PriceCurrency = "SGD"
	PriceCurrencyShp PriceCurrency = "SHP"
	PriceCurrencySle PriceCurrency = "SLE"
	PriceCurrencySll PriceCurrency = "SLL"
	PriceCurrencySos PriceCurrency = "SOS"
	PriceCurrencySrd PriceCurrency = "SRD"
	PriceCurrencySsp PriceCurrency = "SSP"
	PriceCurrencyStn PriceCurrency = "STN"
	PriceCurrencySvc PriceCurrency = "SVC"
	PriceCurrencySzl PriceCurrency = "SZL"
	PriceCurrencyThb PriceCurrency = "THB"
	PriceCurrencyTnd PriceCurrency = "TND"
	PriceCurrencyTop PriceCurrency = "TOP"
	PriceCurrencyTry PriceCurrency = "TRY"
	PriceCurrencyTtd PriceCurrency = "TTD"
	PriceCurrencyTwd PriceCurrency = "TWD"
	PriceCurrencyTzs PriceCurrency = "TZS"
	PriceCurrencyUah PriceCurrency = "UAH"
	PriceCurrencyUgx PriceCurrency = "UGX"
	PriceCurrencyUsd PriceCurrency = "USD"
	PriceCurrencyUyu PriceCurrency = "UYU"
	PriceCurrencyUzs PriceCurrency = "UZS"
	PriceCurrencyVes PriceCurrency = "VES"
	PriceCurrencyVnd PriceCurrency = "VND"
	PriceCurrencyVuv PriceCurrency = "VUV"
	PriceCurrencyWst PriceCurrency = "WST"
	PriceCurrencyXaf PriceCurrency = "XAF"
	PriceCurrencyXcd PriceCurrency = "XCD"
	PriceCurrencyXof PriceCurrency = "XOF"
	PriceCurrencyXpf PriceCurrency = "XPF"
	PriceCurrencyYer PriceCurrency = "YER"
	PriceCurrencyZar PriceCurrency = "ZAR"
	PriceCurrencyZmw PriceCurrency = "ZMW"
)

func (r PriceCurrency) IsKnown() bool {
	switch r {
	case PriceCurrencyAed, PriceCurrencyAll, PriceCurrencyAmd, PriceCurrencyAng, PriceCurrencyAoa, PriceCurrencyArs, PriceCurrencyAud, PriceCurrencyAwg, PriceCurrencyAzn, PriceCurrencyBam, PriceCurrencyBbd, PriceCurrencyBdt, PriceCurrencyBgn, PriceCurrencyBhd, PriceCurrencyBif, PriceCurrencyBmd, PriceCurrencyBnd, PriceCurrencyBob, PriceCurrencyBrl, PriceCurrencyBsd, PriceCurrencyBwp, PriceCurrencyByn, PriceCurrencyBzd, PriceCurrencyCad, PriceCurrencyChf, PriceCurrencyClp, PriceCurrencyCny, PriceCurrencyCop, PriceCurrencyCrc, PriceCurrencyCup, PriceCurrencyCve, PriceCurrencyCzk, PriceCurrencyDjf, PriceCurrencyDkk, PriceCurrencyDop, PriceCurrencyDzd, PriceCurrencyEgp, PriceCurrencyEtb, PriceCurrencyEur, PriceCurrencyFjd, PriceCurrencyFkp, PriceCurrencyGbp, PriceCurrencyGel, PriceCurrencyGhs, PriceCurrencyGip, PriceCurrencyGmd, PriceCurrencyGnf, PriceCurrencyGtq, PriceCurrencyGyd, PriceCurrencyHkd, PriceCurrencyHnl, PriceCurrencyHrk, PriceCurrencyHtg, PriceCurrencyHuf, PriceCurrencyIdr, PriceCurrencyIls, PriceCurrencyInr, PriceCurrencyIqd, PriceCurrencyJmd, PriceCurrencyJod, PriceCurrencyJpy, PriceCurrencyKes, PriceCurrencyKgs, PriceCurrencyKhr, PriceCurrencyKmf, PriceCurrencyKrw, PriceCurrencyKwd, PriceCurrencyKyd, PriceCurrencyKzt, PriceCurrencyLak, PriceCurrencyLbp, PriceCurrencyLkr, PriceCurrencyLrd, PriceCurrencyLsl, PriceCurrencyLyd, PriceCurrencyMad, PriceCurrencyMdl, PriceCurrencyMga, PriceCurrencyMkd, PriceCurrencyMmk, PriceCurrencyMnt, PriceCurrencyMop, PriceCurrencyMru, PriceCurrencyMur, PriceCurrencyMvr, PriceCurrencyMwk, PriceCurrencyMxn, PriceCurrencyMyr, PriceCurrencyMzn, PriceCurrencyNad, PriceCurrencyNgn, PriceCurrencyNio, PriceCurrencyNok, PriceCurrencyNpr, PriceCurrencyNzd, PriceCurrencyOmr, PriceCurrencyPab, PriceCurrencyPen, PriceCurrencyPgk, PriceCurrencyPhp, PriceCurrencyPkr, PriceCurrencyPln, PriceCurrencyPyg, PriceCurrencyQar, PriceCurrencyRon, PriceCurrencyRsd, PriceCurrencyRub, PriceCurrencyRwf, PriceCurrencySar, PriceCurrencySbd, PriceCurrencyScr, PriceCurrencySek, PriceCurrencySgd, PriceCurrencyShp, PriceCurrencySle, PriceCurrencySll, PriceCurrencySos, PriceCurrencySrd, PriceCurrencySsp, PriceCurrencyStn, PriceCurrencySvc, PriceCurrencySzl, PriceCurrencyThb, PriceCurrencyTnd, PriceCurrencyTop, PriceCurrencyTry, PriceCurrencyTtd, PriceCurrencyTwd, PriceCurrencyTzs, PriceCurrencyUah, PriceCurrencyUgx, PriceCurrencyUsd, PriceCurrencyUyu, PriceCurrencyUzs, PriceCurrencyVes, PriceCurrencyVnd, PriceCurrencyVuv, PriceCurrencyWst, PriceCurrencyXaf, PriceCurrencyXcd, PriceCurrencyXof, PriceCurrencyXpf, PriceCurrencyYer, PriceCurrencyZar, PriceCurrencyZmw:
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
	Currency param.Field[PriceCurrency] `json:"currency,required"`
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
	Currency param.Field[PriceOneTimePriceCurrency] `json:"currency,required"`
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
	Currency param.Field[PriceRecurringPriceCurrency] `json:"currency,required"`
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
	TaxCategory ProductTaxCategory `json:"tax_category,required"`
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

// Represents the different categories of taxation applicable to various products
// and services.
type ProductTaxCategory string

const (
	ProductTaxCategoryDigitalProducts ProductTaxCategory = "digital_products"
	ProductTaxCategorySaas            ProductTaxCategory = "saas"
	ProductTaxCategoryEBook           ProductTaxCategory = "e_book"
	ProductTaxCategoryEdtech          ProductTaxCategory = "edtech"
)

func (r ProductTaxCategory) IsKnown() bool {
	switch r {
	case ProductTaxCategoryDigitalProducts, ProductTaxCategorySaas, ProductTaxCategoryEBook, ProductTaxCategoryEdtech:
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

// Represents the different categories of taxation applicable to various products
// and services.
type ProductListResponseTaxCategory string

const (
	ProductListResponseTaxCategoryDigitalProducts ProductListResponseTaxCategory = "digital_products"
	ProductListResponseTaxCategorySaas            ProductListResponseTaxCategory = "saas"
	ProductListResponseTaxCategoryEBook           ProductListResponseTaxCategory = "e_book"
	ProductListResponseTaxCategoryEdtech          ProductListResponseTaxCategory = "edtech"
)

func (r ProductListResponseTaxCategory) IsKnown() bool {
	switch r {
	case ProductListResponseTaxCategoryDigitalProducts, ProductListResponseTaxCategorySaas, ProductListResponseTaxCategoryEBook, ProductListResponseTaxCategoryEdtech:
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
	Price param.Field[PriceUnionParam] `json:"price,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[ProductNewParamsTaxCategory] `json:"tax_category,required"`
	// Addons available for subscription product
	Addons param.Field[[]string] `json:"addons"`
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

// Represents the different categories of taxation applicable to various products
// and services.
type ProductNewParamsTaxCategory string

const (
	ProductNewParamsTaxCategoryDigitalProducts ProductNewParamsTaxCategory = "digital_products"
	ProductNewParamsTaxCategorySaas            ProductNewParamsTaxCategory = "saas"
	ProductNewParamsTaxCategoryEBook           ProductNewParamsTaxCategory = "e_book"
	ProductNewParamsTaxCategoryEdtech          ProductNewParamsTaxCategory = "edtech"
)

func (r ProductNewParamsTaxCategory) IsKnown() bool {
	switch r {
	case ProductNewParamsTaxCategoryDigitalProducts, ProductNewParamsTaxCategorySaas, ProductNewParamsTaxCategoryEBook, ProductNewParamsTaxCategoryEdtech:
		return true
	}
	return false
}

type ProductUpdateParams struct {
	// Available Addons for subscription products
	Addons param.Field[[]string] `json:"addons"`
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
	TaxCategory param.Field[ProductUpdateParamsTaxCategory] `json:"tax_category"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Represents the different categories of taxation applicable to various products
// and services.
type ProductUpdateParamsTaxCategory string

const (
	ProductUpdateParamsTaxCategoryDigitalProducts ProductUpdateParamsTaxCategory = "digital_products"
	ProductUpdateParamsTaxCategorySaas            ProductUpdateParamsTaxCategory = "saas"
	ProductUpdateParamsTaxCategoryEBook           ProductUpdateParamsTaxCategory = "e_book"
	ProductUpdateParamsTaxCategoryEdtech          ProductUpdateParamsTaxCategory = "edtech"
)

func (r ProductUpdateParamsTaxCategory) IsKnown() bool {
	switch r {
	case ProductUpdateParamsTaxCategoryDigitalProducts, ProductUpdateParamsTaxCategorySaas, ProductUpdateParamsTaxCategoryEBook, ProductUpdateParamsTaxCategoryEdtech:
		return true
	}
	return false
}

type ProductListParams struct {
	// List archived products
	Archived param.Field[bool] `query:"archived"`
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
