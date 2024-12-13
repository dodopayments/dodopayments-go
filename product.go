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

	"github.com/stainless-sdks/dodo-payments-go/internal/apijson"
	"github.com/stainless-sdks/dodo-payments-go/internal/apiquery"
	"github.com/stainless-sdks/dodo-payments-go/internal/param"
	"github.com/stainless-sdks/dodo-payments-go/internal/requestconfig"
	"github.com/stainless-sdks/dodo-payments-go/option"
	"github.com/tidwall/gjson"
)

// ProductService contains methods and other services that help with interacting
// with the dodopayments API.
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

func (r *ProductService) New(ctx context.Context, body ProductNewParams, opts ...option.RequestOption) (res *ProductNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ProductService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductGetResponse, err error) {
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

func (r *ProductService) List(ctx context.Context, query ProductListParams, opts ...option.RequestOption) (res *ProductListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProductNewResponse struct {
	ProductID string                 `json:"product_id,required"`
	JSON      productNewResponseJSON `json:"-"`
}

// productNewResponseJSON contains the JSON metadata for the struct
// [ProductNewResponse]
type productNewResponseJSON struct {
	ProductID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProductGetResponse struct {
	BusinessID  string                  `json:"business_id,required"`
	CreatedAt   time.Time               `json:"created_at,required" format:"date-time"`
	IsRecurring bool                    `json:"is_recurring,required"`
	Price       ProductGetResponsePrice `json:"price,required"`
	ProductID   string                  `json:"product_id,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory ProductGetResponseTaxCategory `json:"tax_category,required"`
	UpdatedAt   time.Time                     `json:"updated_at,required" format:"date-time"`
	Description string                        `json:"description,nullable"`
	Image       string                        `json:"image,nullable"`
	Name        string                        `json:"name,nullable"`
	JSON        productGetResponseJSON        `json:"-"`
}

// productGetResponseJSON contains the JSON metadata for the struct
// [ProductGetResponse]
type productGetResponseJSON struct {
	BusinessID  apijson.Field
	CreatedAt   apijson.Field
	IsRecurring apijson.Field
	Price       apijson.Field
	ProductID   apijson.Field
	TaxCategory apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	Image       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProductGetResponsePrice struct {
	Currency ProductGetResponsePriceCurrency `json:"currency,required"`
	Discount float64                         `json:"discount,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                      int64                                             `json:"price,required"`
	PurchasingPowerParity      bool                                              `json:"purchasing_power_parity,required"`
	Type                       ProductGetResponsePriceType                       `json:"type,required"`
	PaymentFrequencyCount      int64                                             `json:"payment_frequency_count"`
	PaymentFrequencyInterval   ProductGetResponsePricePaymentFrequencyInterval   `json:"payment_frequency_interval"`
	SubscriptionPeriodCount    int64                                             `json:"subscription_period_count"`
	SubscriptionPeriodInterval ProductGetResponsePriceSubscriptionPeriodInterval `json:"subscription_period_interval"`
	TrialPeriodDays            int64                                             `json:"trial_period_days"`
	JSON                       productGetResponsePriceJSON                       `json:"-"`
	union                      ProductGetResponsePriceUnion
}

// productGetResponsePriceJSON contains the JSON metadata for the struct
// [ProductGetResponsePrice]
type productGetResponsePriceJSON struct {
	Currency                   apijson.Field
	Discount                   apijson.Field
	Price                      apijson.Field
	PurchasingPowerParity      apijson.Field
	Type                       apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	TrialPeriodDays            apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r productGetResponsePriceJSON) RawJSON() string {
	return r.raw
}

func (r *ProductGetResponsePrice) UnmarshalJSON(data []byte) (err error) {
	*r = ProductGetResponsePrice{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProductGetResponsePriceUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ProductGetResponsePriceOneTimePrice],
// [ProductGetResponsePriceRecurringPrice].
func (r ProductGetResponsePrice) AsUnion() ProductGetResponsePriceUnion {
	return r.union
}

// Union satisfied by [ProductGetResponsePriceOneTimePrice] or
// [ProductGetResponsePriceRecurringPrice].
type ProductGetResponsePriceUnion interface {
	implementsProductGetResponsePrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProductGetResponsePriceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductGetResponsePriceOneTimePrice{}),
			DiscriminatorValue: "one_time_price",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProductGetResponsePriceRecurringPrice{}),
			DiscriminatorValue: "recurring_price",
		},
	)
}

type ProductGetResponsePriceOneTimePrice struct {
	Currency ProductGetResponsePriceOneTimePriceCurrency `json:"currency,required"`
	Discount float64                                     `json:"discount,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                 int64                                   `json:"price,required"`
	PurchasingPowerParity bool                                    `json:"purchasing_power_parity,required"`
	Type                  ProductGetResponsePriceOneTimePriceType `json:"type,required"`
	JSON                  productGetResponsePriceOneTimePriceJSON `json:"-"`
}

// productGetResponsePriceOneTimePriceJSON contains the JSON metadata for the
// struct [ProductGetResponsePriceOneTimePrice]
type productGetResponsePriceOneTimePriceJSON struct {
	Currency              apijson.Field
	Discount              apijson.Field
	Price                 apijson.Field
	PurchasingPowerParity apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProductGetResponsePriceOneTimePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productGetResponsePriceOneTimePriceJSON) RawJSON() string {
	return r.raw
}

func (r ProductGetResponsePriceOneTimePrice) implementsProductGetResponsePrice() {}

type ProductGetResponsePriceOneTimePriceCurrency string

const (
	ProductGetResponsePriceOneTimePriceCurrencyAed ProductGetResponsePriceOneTimePriceCurrency = "AED"
	ProductGetResponsePriceOneTimePriceCurrencyAll ProductGetResponsePriceOneTimePriceCurrency = "ALL"
	ProductGetResponsePriceOneTimePriceCurrencyAmd ProductGetResponsePriceOneTimePriceCurrency = "AMD"
	ProductGetResponsePriceOneTimePriceCurrencyAng ProductGetResponsePriceOneTimePriceCurrency = "ANG"
	ProductGetResponsePriceOneTimePriceCurrencyAoa ProductGetResponsePriceOneTimePriceCurrency = "AOA"
	ProductGetResponsePriceOneTimePriceCurrencyArs ProductGetResponsePriceOneTimePriceCurrency = "ARS"
	ProductGetResponsePriceOneTimePriceCurrencyAud ProductGetResponsePriceOneTimePriceCurrency = "AUD"
	ProductGetResponsePriceOneTimePriceCurrencyAwg ProductGetResponsePriceOneTimePriceCurrency = "AWG"
	ProductGetResponsePriceOneTimePriceCurrencyAzn ProductGetResponsePriceOneTimePriceCurrency = "AZN"
	ProductGetResponsePriceOneTimePriceCurrencyBam ProductGetResponsePriceOneTimePriceCurrency = "BAM"
	ProductGetResponsePriceOneTimePriceCurrencyBbd ProductGetResponsePriceOneTimePriceCurrency = "BBD"
	ProductGetResponsePriceOneTimePriceCurrencyBdt ProductGetResponsePriceOneTimePriceCurrency = "BDT"
	ProductGetResponsePriceOneTimePriceCurrencyBgn ProductGetResponsePriceOneTimePriceCurrency = "BGN"
	ProductGetResponsePriceOneTimePriceCurrencyBhd ProductGetResponsePriceOneTimePriceCurrency = "BHD"
	ProductGetResponsePriceOneTimePriceCurrencyBif ProductGetResponsePriceOneTimePriceCurrency = "BIF"
	ProductGetResponsePriceOneTimePriceCurrencyBmd ProductGetResponsePriceOneTimePriceCurrency = "BMD"
	ProductGetResponsePriceOneTimePriceCurrencyBnd ProductGetResponsePriceOneTimePriceCurrency = "BND"
	ProductGetResponsePriceOneTimePriceCurrencyBob ProductGetResponsePriceOneTimePriceCurrency = "BOB"
	ProductGetResponsePriceOneTimePriceCurrencyBrl ProductGetResponsePriceOneTimePriceCurrency = "BRL"
	ProductGetResponsePriceOneTimePriceCurrencyBsd ProductGetResponsePriceOneTimePriceCurrency = "BSD"
	ProductGetResponsePriceOneTimePriceCurrencyBwp ProductGetResponsePriceOneTimePriceCurrency = "BWP"
	ProductGetResponsePriceOneTimePriceCurrencyByn ProductGetResponsePriceOneTimePriceCurrency = "BYN"
	ProductGetResponsePriceOneTimePriceCurrencyBzd ProductGetResponsePriceOneTimePriceCurrency = "BZD"
	ProductGetResponsePriceOneTimePriceCurrencyCad ProductGetResponsePriceOneTimePriceCurrency = "CAD"
	ProductGetResponsePriceOneTimePriceCurrencyChf ProductGetResponsePriceOneTimePriceCurrency = "CHF"
	ProductGetResponsePriceOneTimePriceCurrencyClp ProductGetResponsePriceOneTimePriceCurrency = "CLP"
	ProductGetResponsePriceOneTimePriceCurrencyCny ProductGetResponsePriceOneTimePriceCurrency = "CNY"
	ProductGetResponsePriceOneTimePriceCurrencyCop ProductGetResponsePriceOneTimePriceCurrency = "COP"
	ProductGetResponsePriceOneTimePriceCurrencyCrc ProductGetResponsePriceOneTimePriceCurrency = "CRC"
	ProductGetResponsePriceOneTimePriceCurrencyCup ProductGetResponsePriceOneTimePriceCurrency = "CUP"
	ProductGetResponsePriceOneTimePriceCurrencyCve ProductGetResponsePriceOneTimePriceCurrency = "CVE"
	ProductGetResponsePriceOneTimePriceCurrencyCzk ProductGetResponsePriceOneTimePriceCurrency = "CZK"
	ProductGetResponsePriceOneTimePriceCurrencyDjf ProductGetResponsePriceOneTimePriceCurrency = "DJF"
	ProductGetResponsePriceOneTimePriceCurrencyDkk ProductGetResponsePriceOneTimePriceCurrency = "DKK"
	ProductGetResponsePriceOneTimePriceCurrencyDop ProductGetResponsePriceOneTimePriceCurrency = "DOP"
	ProductGetResponsePriceOneTimePriceCurrencyDzd ProductGetResponsePriceOneTimePriceCurrency = "DZD"
	ProductGetResponsePriceOneTimePriceCurrencyEgp ProductGetResponsePriceOneTimePriceCurrency = "EGP"
	ProductGetResponsePriceOneTimePriceCurrencyEtb ProductGetResponsePriceOneTimePriceCurrency = "ETB"
	ProductGetResponsePriceOneTimePriceCurrencyEur ProductGetResponsePriceOneTimePriceCurrency = "EUR"
	ProductGetResponsePriceOneTimePriceCurrencyFjd ProductGetResponsePriceOneTimePriceCurrency = "FJD"
	ProductGetResponsePriceOneTimePriceCurrencyFkp ProductGetResponsePriceOneTimePriceCurrency = "FKP"
	ProductGetResponsePriceOneTimePriceCurrencyGbp ProductGetResponsePriceOneTimePriceCurrency = "GBP"
	ProductGetResponsePriceOneTimePriceCurrencyGel ProductGetResponsePriceOneTimePriceCurrency = "GEL"
	ProductGetResponsePriceOneTimePriceCurrencyGhs ProductGetResponsePriceOneTimePriceCurrency = "GHS"
	ProductGetResponsePriceOneTimePriceCurrencyGip ProductGetResponsePriceOneTimePriceCurrency = "GIP"
	ProductGetResponsePriceOneTimePriceCurrencyGmd ProductGetResponsePriceOneTimePriceCurrency = "GMD"
	ProductGetResponsePriceOneTimePriceCurrencyGnf ProductGetResponsePriceOneTimePriceCurrency = "GNF"
	ProductGetResponsePriceOneTimePriceCurrencyGtq ProductGetResponsePriceOneTimePriceCurrency = "GTQ"
	ProductGetResponsePriceOneTimePriceCurrencyGyd ProductGetResponsePriceOneTimePriceCurrency = "GYD"
	ProductGetResponsePriceOneTimePriceCurrencyHkd ProductGetResponsePriceOneTimePriceCurrency = "HKD"
	ProductGetResponsePriceOneTimePriceCurrencyHnl ProductGetResponsePriceOneTimePriceCurrency = "HNL"
	ProductGetResponsePriceOneTimePriceCurrencyHrk ProductGetResponsePriceOneTimePriceCurrency = "HRK"
	ProductGetResponsePriceOneTimePriceCurrencyHtg ProductGetResponsePriceOneTimePriceCurrency = "HTG"
	ProductGetResponsePriceOneTimePriceCurrencyHuf ProductGetResponsePriceOneTimePriceCurrency = "HUF"
	ProductGetResponsePriceOneTimePriceCurrencyIdr ProductGetResponsePriceOneTimePriceCurrency = "IDR"
	ProductGetResponsePriceOneTimePriceCurrencyIls ProductGetResponsePriceOneTimePriceCurrency = "ILS"
	ProductGetResponsePriceOneTimePriceCurrencyInr ProductGetResponsePriceOneTimePriceCurrency = "INR"
	ProductGetResponsePriceOneTimePriceCurrencyIqd ProductGetResponsePriceOneTimePriceCurrency = "IQD"
	ProductGetResponsePriceOneTimePriceCurrencyJmd ProductGetResponsePriceOneTimePriceCurrency = "JMD"
	ProductGetResponsePriceOneTimePriceCurrencyJod ProductGetResponsePriceOneTimePriceCurrency = "JOD"
	ProductGetResponsePriceOneTimePriceCurrencyJpy ProductGetResponsePriceOneTimePriceCurrency = "JPY"
	ProductGetResponsePriceOneTimePriceCurrencyKes ProductGetResponsePriceOneTimePriceCurrency = "KES"
	ProductGetResponsePriceOneTimePriceCurrencyKgs ProductGetResponsePriceOneTimePriceCurrency = "KGS"
	ProductGetResponsePriceOneTimePriceCurrencyKhr ProductGetResponsePriceOneTimePriceCurrency = "KHR"
	ProductGetResponsePriceOneTimePriceCurrencyKmf ProductGetResponsePriceOneTimePriceCurrency = "KMF"
	ProductGetResponsePriceOneTimePriceCurrencyKrw ProductGetResponsePriceOneTimePriceCurrency = "KRW"
	ProductGetResponsePriceOneTimePriceCurrencyKwd ProductGetResponsePriceOneTimePriceCurrency = "KWD"
	ProductGetResponsePriceOneTimePriceCurrencyKyd ProductGetResponsePriceOneTimePriceCurrency = "KYD"
	ProductGetResponsePriceOneTimePriceCurrencyKzt ProductGetResponsePriceOneTimePriceCurrency = "KZT"
	ProductGetResponsePriceOneTimePriceCurrencyLak ProductGetResponsePriceOneTimePriceCurrency = "LAK"
	ProductGetResponsePriceOneTimePriceCurrencyLbp ProductGetResponsePriceOneTimePriceCurrency = "LBP"
	ProductGetResponsePriceOneTimePriceCurrencyLkr ProductGetResponsePriceOneTimePriceCurrency = "LKR"
	ProductGetResponsePriceOneTimePriceCurrencyLrd ProductGetResponsePriceOneTimePriceCurrency = "LRD"
	ProductGetResponsePriceOneTimePriceCurrencyLsl ProductGetResponsePriceOneTimePriceCurrency = "LSL"
	ProductGetResponsePriceOneTimePriceCurrencyLyd ProductGetResponsePriceOneTimePriceCurrency = "LYD"
	ProductGetResponsePriceOneTimePriceCurrencyMad ProductGetResponsePriceOneTimePriceCurrency = "MAD"
	ProductGetResponsePriceOneTimePriceCurrencyMdl ProductGetResponsePriceOneTimePriceCurrency = "MDL"
	ProductGetResponsePriceOneTimePriceCurrencyMga ProductGetResponsePriceOneTimePriceCurrency = "MGA"
	ProductGetResponsePriceOneTimePriceCurrencyMkd ProductGetResponsePriceOneTimePriceCurrency = "MKD"
	ProductGetResponsePriceOneTimePriceCurrencyMmk ProductGetResponsePriceOneTimePriceCurrency = "MMK"
	ProductGetResponsePriceOneTimePriceCurrencyMnt ProductGetResponsePriceOneTimePriceCurrency = "MNT"
	ProductGetResponsePriceOneTimePriceCurrencyMop ProductGetResponsePriceOneTimePriceCurrency = "MOP"
	ProductGetResponsePriceOneTimePriceCurrencyMru ProductGetResponsePriceOneTimePriceCurrency = "MRU"
	ProductGetResponsePriceOneTimePriceCurrencyMur ProductGetResponsePriceOneTimePriceCurrency = "MUR"
	ProductGetResponsePriceOneTimePriceCurrencyMvr ProductGetResponsePriceOneTimePriceCurrency = "MVR"
	ProductGetResponsePriceOneTimePriceCurrencyMwk ProductGetResponsePriceOneTimePriceCurrency = "MWK"
	ProductGetResponsePriceOneTimePriceCurrencyMxn ProductGetResponsePriceOneTimePriceCurrency = "MXN"
	ProductGetResponsePriceOneTimePriceCurrencyMyr ProductGetResponsePriceOneTimePriceCurrency = "MYR"
	ProductGetResponsePriceOneTimePriceCurrencyMzn ProductGetResponsePriceOneTimePriceCurrency = "MZN"
	ProductGetResponsePriceOneTimePriceCurrencyNad ProductGetResponsePriceOneTimePriceCurrency = "NAD"
	ProductGetResponsePriceOneTimePriceCurrencyNgn ProductGetResponsePriceOneTimePriceCurrency = "NGN"
	ProductGetResponsePriceOneTimePriceCurrencyNio ProductGetResponsePriceOneTimePriceCurrency = "NIO"
	ProductGetResponsePriceOneTimePriceCurrencyNok ProductGetResponsePriceOneTimePriceCurrency = "NOK"
	ProductGetResponsePriceOneTimePriceCurrencyNpr ProductGetResponsePriceOneTimePriceCurrency = "NPR"
	ProductGetResponsePriceOneTimePriceCurrencyNzd ProductGetResponsePriceOneTimePriceCurrency = "NZD"
	ProductGetResponsePriceOneTimePriceCurrencyOmr ProductGetResponsePriceOneTimePriceCurrency = "OMR"
	ProductGetResponsePriceOneTimePriceCurrencyPab ProductGetResponsePriceOneTimePriceCurrency = "PAB"
	ProductGetResponsePriceOneTimePriceCurrencyPen ProductGetResponsePriceOneTimePriceCurrency = "PEN"
	ProductGetResponsePriceOneTimePriceCurrencyPgk ProductGetResponsePriceOneTimePriceCurrency = "PGK"
	ProductGetResponsePriceOneTimePriceCurrencyPhp ProductGetResponsePriceOneTimePriceCurrency = "PHP"
	ProductGetResponsePriceOneTimePriceCurrencyPkr ProductGetResponsePriceOneTimePriceCurrency = "PKR"
	ProductGetResponsePriceOneTimePriceCurrencyPln ProductGetResponsePriceOneTimePriceCurrency = "PLN"
	ProductGetResponsePriceOneTimePriceCurrencyPyg ProductGetResponsePriceOneTimePriceCurrency = "PYG"
	ProductGetResponsePriceOneTimePriceCurrencyQar ProductGetResponsePriceOneTimePriceCurrency = "QAR"
	ProductGetResponsePriceOneTimePriceCurrencyRon ProductGetResponsePriceOneTimePriceCurrency = "RON"
	ProductGetResponsePriceOneTimePriceCurrencyRsd ProductGetResponsePriceOneTimePriceCurrency = "RSD"
	ProductGetResponsePriceOneTimePriceCurrencyRub ProductGetResponsePriceOneTimePriceCurrency = "RUB"
	ProductGetResponsePriceOneTimePriceCurrencyRwf ProductGetResponsePriceOneTimePriceCurrency = "RWF"
	ProductGetResponsePriceOneTimePriceCurrencySar ProductGetResponsePriceOneTimePriceCurrency = "SAR"
	ProductGetResponsePriceOneTimePriceCurrencySbd ProductGetResponsePriceOneTimePriceCurrency = "SBD"
	ProductGetResponsePriceOneTimePriceCurrencyScr ProductGetResponsePriceOneTimePriceCurrency = "SCR"
	ProductGetResponsePriceOneTimePriceCurrencySek ProductGetResponsePriceOneTimePriceCurrency = "SEK"
	ProductGetResponsePriceOneTimePriceCurrencySgd ProductGetResponsePriceOneTimePriceCurrency = "SGD"
	ProductGetResponsePriceOneTimePriceCurrencyShp ProductGetResponsePriceOneTimePriceCurrency = "SHP"
	ProductGetResponsePriceOneTimePriceCurrencySle ProductGetResponsePriceOneTimePriceCurrency = "SLE"
	ProductGetResponsePriceOneTimePriceCurrencySll ProductGetResponsePriceOneTimePriceCurrency = "SLL"
	ProductGetResponsePriceOneTimePriceCurrencySos ProductGetResponsePriceOneTimePriceCurrency = "SOS"
	ProductGetResponsePriceOneTimePriceCurrencySrd ProductGetResponsePriceOneTimePriceCurrency = "SRD"
	ProductGetResponsePriceOneTimePriceCurrencySsp ProductGetResponsePriceOneTimePriceCurrency = "SSP"
	ProductGetResponsePriceOneTimePriceCurrencyStn ProductGetResponsePriceOneTimePriceCurrency = "STN"
	ProductGetResponsePriceOneTimePriceCurrencySvc ProductGetResponsePriceOneTimePriceCurrency = "SVC"
	ProductGetResponsePriceOneTimePriceCurrencySzl ProductGetResponsePriceOneTimePriceCurrency = "SZL"
	ProductGetResponsePriceOneTimePriceCurrencyThb ProductGetResponsePriceOneTimePriceCurrency = "THB"
	ProductGetResponsePriceOneTimePriceCurrencyTnd ProductGetResponsePriceOneTimePriceCurrency = "TND"
	ProductGetResponsePriceOneTimePriceCurrencyTop ProductGetResponsePriceOneTimePriceCurrency = "TOP"
	ProductGetResponsePriceOneTimePriceCurrencyTry ProductGetResponsePriceOneTimePriceCurrency = "TRY"
	ProductGetResponsePriceOneTimePriceCurrencyTtd ProductGetResponsePriceOneTimePriceCurrency = "TTD"
	ProductGetResponsePriceOneTimePriceCurrencyTwd ProductGetResponsePriceOneTimePriceCurrency = "TWD"
	ProductGetResponsePriceOneTimePriceCurrencyTzs ProductGetResponsePriceOneTimePriceCurrency = "TZS"
	ProductGetResponsePriceOneTimePriceCurrencyUah ProductGetResponsePriceOneTimePriceCurrency = "UAH"
	ProductGetResponsePriceOneTimePriceCurrencyUgx ProductGetResponsePriceOneTimePriceCurrency = "UGX"
	ProductGetResponsePriceOneTimePriceCurrencyUsd ProductGetResponsePriceOneTimePriceCurrency = "USD"
	ProductGetResponsePriceOneTimePriceCurrencyUyu ProductGetResponsePriceOneTimePriceCurrency = "UYU"
	ProductGetResponsePriceOneTimePriceCurrencyUzs ProductGetResponsePriceOneTimePriceCurrency = "UZS"
	ProductGetResponsePriceOneTimePriceCurrencyVes ProductGetResponsePriceOneTimePriceCurrency = "VES"
	ProductGetResponsePriceOneTimePriceCurrencyVnd ProductGetResponsePriceOneTimePriceCurrency = "VND"
	ProductGetResponsePriceOneTimePriceCurrencyVuv ProductGetResponsePriceOneTimePriceCurrency = "VUV"
	ProductGetResponsePriceOneTimePriceCurrencyWst ProductGetResponsePriceOneTimePriceCurrency = "WST"
	ProductGetResponsePriceOneTimePriceCurrencyXaf ProductGetResponsePriceOneTimePriceCurrency = "XAF"
	ProductGetResponsePriceOneTimePriceCurrencyXcd ProductGetResponsePriceOneTimePriceCurrency = "XCD"
	ProductGetResponsePriceOneTimePriceCurrencyXof ProductGetResponsePriceOneTimePriceCurrency = "XOF"
	ProductGetResponsePriceOneTimePriceCurrencyXpf ProductGetResponsePriceOneTimePriceCurrency = "XPF"
	ProductGetResponsePriceOneTimePriceCurrencyYer ProductGetResponsePriceOneTimePriceCurrency = "YER"
	ProductGetResponsePriceOneTimePriceCurrencyZar ProductGetResponsePriceOneTimePriceCurrency = "ZAR"
	ProductGetResponsePriceOneTimePriceCurrencyZmw ProductGetResponsePriceOneTimePriceCurrency = "ZMW"
)

func (r ProductGetResponsePriceOneTimePriceCurrency) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceOneTimePriceCurrencyAed, ProductGetResponsePriceOneTimePriceCurrencyAll, ProductGetResponsePriceOneTimePriceCurrencyAmd, ProductGetResponsePriceOneTimePriceCurrencyAng, ProductGetResponsePriceOneTimePriceCurrencyAoa, ProductGetResponsePriceOneTimePriceCurrencyArs, ProductGetResponsePriceOneTimePriceCurrencyAud, ProductGetResponsePriceOneTimePriceCurrencyAwg, ProductGetResponsePriceOneTimePriceCurrencyAzn, ProductGetResponsePriceOneTimePriceCurrencyBam, ProductGetResponsePriceOneTimePriceCurrencyBbd, ProductGetResponsePriceOneTimePriceCurrencyBdt, ProductGetResponsePriceOneTimePriceCurrencyBgn, ProductGetResponsePriceOneTimePriceCurrencyBhd, ProductGetResponsePriceOneTimePriceCurrencyBif, ProductGetResponsePriceOneTimePriceCurrencyBmd, ProductGetResponsePriceOneTimePriceCurrencyBnd, ProductGetResponsePriceOneTimePriceCurrencyBob, ProductGetResponsePriceOneTimePriceCurrencyBrl, ProductGetResponsePriceOneTimePriceCurrencyBsd, ProductGetResponsePriceOneTimePriceCurrencyBwp, ProductGetResponsePriceOneTimePriceCurrencyByn, ProductGetResponsePriceOneTimePriceCurrencyBzd, ProductGetResponsePriceOneTimePriceCurrencyCad, ProductGetResponsePriceOneTimePriceCurrencyChf, ProductGetResponsePriceOneTimePriceCurrencyClp, ProductGetResponsePriceOneTimePriceCurrencyCny, ProductGetResponsePriceOneTimePriceCurrencyCop, ProductGetResponsePriceOneTimePriceCurrencyCrc, ProductGetResponsePriceOneTimePriceCurrencyCup, ProductGetResponsePriceOneTimePriceCurrencyCve, ProductGetResponsePriceOneTimePriceCurrencyCzk, ProductGetResponsePriceOneTimePriceCurrencyDjf, ProductGetResponsePriceOneTimePriceCurrencyDkk, ProductGetResponsePriceOneTimePriceCurrencyDop, ProductGetResponsePriceOneTimePriceCurrencyDzd, ProductGetResponsePriceOneTimePriceCurrencyEgp, ProductGetResponsePriceOneTimePriceCurrencyEtb, ProductGetResponsePriceOneTimePriceCurrencyEur, ProductGetResponsePriceOneTimePriceCurrencyFjd, ProductGetResponsePriceOneTimePriceCurrencyFkp, ProductGetResponsePriceOneTimePriceCurrencyGbp, ProductGetResponsePriceOneTimePriceCurrencyGel, ProductGetResponsePriceOneTimePriceCurrencyGhs, ProductGetResponsePriceOneTimePriceCurrencyGip, ProductGetResponsePriceOneTimePriceCurrencyGmd, ProductGetResponsePriceOneTimePriceCurrencyGnf, ProductGetResponsePriceOneTimePriceCurrencyGtq, ProductGetResponsePriceOneTimePriceCurrencyGyd, ProductGetResponsePriceOneTimePriceCurrencyHkd, ProductGetResponsePriceOneTimePriceCurrencyHnl, ProductGetResponsePriceOneTimePriceCurrencyHrk, ProductGetResponsePriceOneTimePriceCurrencyHtg, ProductGetResponsePriceOneTimePriceCurrencyHuf, ProductGetResponsePriceOneTimePriceCurrencyIdr, ProductGetResponsePriceOneTimePriceCurrencyIls, ProductGetResponsePriceOneTimePriceCurrencyInr, ProductGetResponsePriceOneTimePriceCurrencyIqd, ProductGetResponsePriceOneTimePriceCurrencyJmd, ProductGetResponsePriceOneTimePriceCurrencyJod, ProductGetResponsePriceOneTimePriceCurrencyJpy, ProductGetResponsePriceOneTimePriceCurrencyKes, ProductGetResponsePriceOneTimePriceCurrencyKgs, ProductGetResponsePriceOneTimePriceCurrencyKhr, ProductGetResponsePriceOneTimePriceCurrencyKmf, ProductGetResponsePriceOneTimePriceCurrencyKrw, ProductGetResponsePriceOneTimePriceCurrencyKwd, ProductGetResponsePriceOneTimePriceCurrencyKyd, ProductGetResponsePriceOneTimePriceCurrencyKzt, ProductGetResponsePriceOneTimePriceCurrencyLak, ProductGetResponsePriceOneTimePriceCurrencyLbp, ProductGetResponsePriceOneTimePriceCurrencyLkr, ProductGetResponsePriceOneTimePriceCurrencyLrd, ProductGetResponsePriceOneTimePriceCurrencyLsl, ProductGetResponsePriceOneTimePriceCurrencyLyd, ProductGetResponsePriceOneTimePriceCurrencyMad, ProductGetResponsePriceOneTimePriceCurrencyMdl, ProductGetResponsePriceOneTimePriceCurrencyMga, ProductGetResponsePriceOneTimePriceCurrencyMkd, ProductGetResponsePriceOneTimePriceCurrencyMmk, ProductGetResponsePriceOneTimePriceCurrencyMnt, ProductGetResponsePriceOneTimePriceCurrencyMop, ProductGetResponsePriceOneTimePriceCurrencyMru, ProductGetResponsePriceOneTimePriceCurrencyMur, ProductGetResponsePriceOneTimePriceCurrencyMvr, ProductGetResponsePriceOneTimePriceCurrencyMwk, ProductGetResponsePriceOneTimePriceCurrencyMxn, ProductGetResponsePriceOneTimePriceCurrencyMyr, ProductGetResponsePriceOneTimePriceCurrencyMzn, ProductGetResponsePriceOneTimePriceCurrencyNad, ProductGetResponsePriceOneTimePriceCurrencyNgn, ProductGetResponsePriceOneTimePriceCurrencyNio, ProductGetResponsePriceOneTimePriceCurrencyNok, ProductGetResponsePriceOneTimePriceCurrencyNpr, ProductGetResponsePriceOneTimePriceCurrencyNzd, ProductGetResponsePriceOneTimePriceCurrencyOmr, ProductGetResponsePriceOneTimePriceCurrencyPab, ProductGetResponsePriceOneTimePriceCurrencyPen, ProductGetResponsePriceOneTimePriceCurrencyPgk, ProductGetResponsePriceOneTimePriceCurrencyPhp, ProductGetResponsePriceOneTimePriceCurrencyPkr, ProductGetResponsePriceOneTimePriceCurrencyPln, ProductGetResponsePriceOneTimePriceCurrencyPyg, ProductGetResponsePriceOneTimePriceCurrencyQar, ProductGetResponsePriceOneTimePriceCurrencyRon, ProductGetResponsePriceOneTimePriceCurrencyRsd, ProductGetResponsePriceOneTimePriceCurrencyRub, ProductGetResponsePriceOneTimePriceCurrencyRwf, ProductGetResponsePriceOneTimePriceCurrencySar, ProductGetResponsePriceOneTimePriceCurrencySbd, ProductGetResponsePriceOneTimePriceCurrencyScr, ProductGetResponsePriceOneTimePriceCurrencySek, ProductGetResponsePriceOneTimePriceCurrencySgd, ProductGetResponsePriceOneTimePriceCurrencyShp, ProductGetResponsePriceOneTimePriceCurrencySle, ProductGetResponsePriceOneTimePriceCurrencySll, ProductGetResponsePriceOneTimePriceCurrencySos, ProductGetResponsePriceOneTimePriceCurrencySrd, ProductGetResponsePriceOneTimePriceCurrencySsp, ProductGetResponsePriceOneTimePriceCurrencyStn, ProductGetResponsePriceOneTimePriceCurrencySvc, ProductGetResponsePriceOneTimePriceCurrencySzl, ProductGetResponsePriceOneTimePriceCurrencyThb, ProductGetResponsePriceOneTimePriceCurrencyTnd, ProductGetResponsePriceOneTimePriceCurrencyTop, ProductGetResponsePriceOneTimePriceCurrencyTry, ProductGetResponsePriceOneTimePriceCurrencyTtd, ProductGetResponsePriceOneTimePriceCurrencyTwd, ProductGetResponsePriceOneTimePriceCurrencyTzs, ProductGetResponsePriceOneTimePriceCurrencyUah, ProductGetResponsePriceOneTimePriceCurrencyUgx, ProductGetResponsePriceOneTimePriceCurrencyUsd, ProductGetResponsePriceOneTimePriceCurrencyUyu, ProductGetResponsePriceOneTimePriceCurrencyUzs, ProductGetResponsePriceOneTimePriceCurrencyVes, ProductGetResponsePriceOneTimePriceCurrencyVnd, ProductGetResponsePriceOneTimePriceCurrencyVuv, ProductGetResponsePriceOneTimePriceCurrencyWst, ProductGetResponsePriceOneTimePriceCurrencyXaf, ProductGetResponsePriceOneTimePriceCurrencyXcd, ProductGetResponsePriceOneTimePriceCurrencyXof, ProductGetResponsePriceOneTimePriceCurrencyXpf, ProductGetResponsePriceOneTimePriceCurrencyYer, ProductGetResponsePriceOneTimePriceCurrencyZar, ProductGetResponsePriceOneTimePriceCurrencyZmw:
		return true
	}
	return false
}

type ProductGetResponsePriceOneTimePriceType string

const (
	ProductGetResponsePriceOneTimePriceTypeOneTimePrice ProductGetResponsePriceOneTimePriceType = "one_time_price"
)

func (r ProductGetResponsePriceOneTimePriceType) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceOneTimePriceTypeOneTimePrice:
		return true
	}
	return false
}

type ProductGetResponsePriceRecurringPrice struct {
	Currency                 ProductGetResponsePriceRecurringPriceCurrency                 `json:"currency,required"`
	Discount                 float64                                                       `json:"discount,required"`
	PaymentFrequencyCount    int64                                                         `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval ProductGetResponsePriceRecurringPricePaymentFrequencyInterval `json:"payment_frequency_interval,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                      int64                                                           `json:"price,required"`
	PurchasingPowerParity      bool                                                            `json:"purchasing_power_parity,required"`
	SubscriptionPeriodCount    int64                                                           `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval ProductGetResponsePriceRecurringPriceSubscriptionPeriodInterval `json:"subscription_period_interval,required"`
	TrialPeriodDays            int64                                                           `json:"trial_period_days,required"`
	Type                       ProductGetResponsePriceRecurringPriceType                       `json:"type,required"`
	JSON                       productGetResponsePriceRecurringPriceJSON                       `json:"-"`
}

// productGetResponsePriceRecurringPriceJSON contains the JSON metadata for the
// struct [ProductGetResponsePriceRecurringPrice]
type productGetResponsePriceRecurringPriceJSON struct {
	Currency                   apijson.Field
	Discount                   apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	Price                      apijson.Field
	PurchasingPowerParity      apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	TrialPeriodDays            apijson.Field
	Type                       apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ProductGetResponsePriceRecurringPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productGetResponsePriceRecurringPriceJSON) RawJSON() string {
	return r.raw
}

func (r ProductGetResponsePriceRecurringPrice) implementsProductGetResponsePrice() {}

type ProductGetResponsePriceRecurringPriceCurrency string

const (
	ProductGetResponsePriceRecurringPriceCurrencyAed ProductGetResponsePriceRecurringPriceCurrency = "AED"
	ProductGetResponsePriceRecurringPriceCurrencyAll ProductGetResponsePriceRecurringPriceCurrency = "ALL"
	ProductGetResponsePriceRecurringPriceCurrencyAmd ProductGetResponsePriceRecurringPriceCurrency = "AMD"
	ProductGetResponsePriceRecurringPriceCurrencyAng ProductGetResponsePriceRecurringPriceCurrency = "ANG"
	ProductGetResponsePriceRecurringPriceCurrencyAoa ProductGetResponsePriceRecurringPriceCurrency = "AOA"
	ProductGetResponsePriceRecurringPriceCurrencyArs ProductGetResponsePriceRecurringPriceCurrency = "ARS"
	ProductGetResponsePriceRecurringPriceCurrencyAud ProductGetResponsePriceRecurringPriceCurrency = "AUD"
	ProductGetResponsePriceRecurringPriceCurrencyAwg ProductGetResponsePriceRecurringPriceCurrency = "AWG"
	ProductGetResponsePriceRecurringPriceCurrencyAzn ProductGetResponsePriceRecurringPriceCurrency = "AZN"
	ProductGetResponsePriceRecurringPriceCurrencyBam ProductGetResponsePriceRecurringPriceCurrency = "BAM"
	ProductGetResponsePriceRecurringPriceCurrencyBbd ProductGetResponsePriceRecurringPriceCurrency = "BBD"
	ProductGetResponsePriceRecurringPriceCurrencyBdt ProductGetResponsePriceRecurringPriceCurrency = "BDT"
	ProductGetResponsePriceRecurringPriceCurrencyBgn ProductGetResponsePriceRecurringPriceCurrency = "BGN"
	ProductGetResponsePriceRecurringPriceCurrencyBhd ProductGetResponsePriceRecurringPriceCurrency = "BHD"
	ProductGetResponsePriceRecurringPriceCurrencyBif ProductGetResponsePriceRecurringPriceCurrency = "BIF"
	ProductGetResponsePriceRecurringPriceCurrencyBmd ProductGetResponsePriceRecurringPriceCurrency = "BMD"
	ProductGetResponsePriceRecurringPriceCurrencyBnd ProductGetResponsePriceRecurringPriceCurrency = "BND"
	ProductGetResponsePriceRecurringPriceCurrencyBob ProductGetResponsePriceRecurringPriceCurrency = "BOB"
	ProductGetResponsePriceRecurringPriceCurrencyBrl ProductGetResponsePriceRecurringPriceCurrency = "BRL"
	ProductGetResponsePriceRecurringPriceCurrencyBsd ProductGetResponsePriceRecurringPriceCurrency = "BSD"
	ProductGetResponsePriceRecurringPriceCurrencyBwp ProductGetResponsePriceRecurringPriceCurrency = "BWP"
	ProductGetResponsePriceRecurringPriceCurrencyByn ProductGetResponsePriceRecurringPriceCurrency = "BYN"
	ProductGetResponsePriceRecurringPriceCurrencyBzd ProductGetResponsePriceRecurringPriceCurrency = "BZD"
	ProductGetResponsePriceRecurringPriceCurrencyCad ProductGetResponsePriceRecurringPriceCurrency = "CAD"
	ProductGetResponsePriceRecurringPriceCurrencyChf ProductGetResponsePriceRecurringPriceCurrency = "CHF"
	ProductGetResponsePriceRecurringPriceCurrencyClp ProductGetResponsePriceRecurringPriceCurrency = "CLP"
	ProductGetResponsePriceRecurringPriceCurrencyCny ProductGetResponsePriceRecurringPriceCurrency = "CNY"
	ProductGetResponsePriceRecurringPriceCurrencyCop ProductGetResponsePriceRecurringPriceCurrency = "COP"
	ProductGetResponsePriceRecurringPriceCurrencyCrc ProductGetResponsePriceRecurringPriceCurrency = "CRC"
	ProductGetResponsePriceRecurringPriceCurrencyCup ProductGetResponsePriceRecurringPriceCurrency = "CUP"
	ProductGetResponsePriceRecurringPriceCurrencyCve ProductGetResponsePriceRecurringPriceCurrency = "CVE"
	ProductGetResponsePriceRecurringPriceCurrencyCzk ProductGetResponsePriceRecurringPriceCurrency = "CZK"
	ProductGetResponsePriceRecurringPriceCurrencyDjf ProductGetResponsePriceRecurringPriceCurrency = "DJF"
	ProductGetResponsePriceRecurringPriceCurrencyDkk ProductGetResponsePriceRecurringPriceCurrency = "DKK"
	ProductGetResponsePriceRecurringPriceCurrencyDop ProductGetResponsePriceRecurringPriceCurrency = "DOP"
	ProductGetResponsePriceRecurringPriceCurrencyDzd ProductGetResponsePriceRecurringPriceCurrency = "DZD"
	ProductGetResponsePriceRecurringPriceCurrencyEgp ProductGetResponsePriceRecurringPriceCurrency = "EGP"
	ProductGetResponsePriceRecurringPriceCurrencyEtb ProductGetResponsePriceRecurringPriceCurrency = "ETB"
	ProductGetResponsePriceRecurringPriceCurrencyEur ProductGetResponsePriceRecurringPriceCurrency = "EUR"
	ProductGetResponsePriceRecurringPriceCurrencyFjd ProductGetResponsePriceRecurringPriceCurrency = "FJD"
	ProductGetResponsePriceRecurringPriceCurrencyFkp ProductGetResponsePriceRecurringPriceCurrency = "FKP"
	ProductGetResponsePriceRecurringPriceCurrencyGbp ProductGetResponsePriceRecurringPriceCurrency = "GBP"
	ProductGetResponsePriceRecurringPriceCurrencyGel ProductGetResponsePriceRecurringPriceCurrency = "GEL"
	ProductGetResponsePriceRecurringPriceCurrencyGhs ProductGetResponsePriceRecurringPriceCurrency = "GHS"
	ProductGetResponsePriceRecurringPriceCurrencyGip ProductGetResponsePriceRecurringPriceCurrency = "GIP"
	ProductGetResponsePriceRecurringPriceCurrencyGmd ProductGetResponsePriceRecurringPriceCurrency = "GMD"
	ProductGetResponsePriceRecurringPriceCurrencyGnf ProductGetResponsePriceRecurringPriceCurrency = "GNF"
	ProductGetResponsePriceRecurringPriceCurrencyGtq ProductGetResponsePriceRecurringPriceCurrency = "GTQ"
	ProductGetResponsePriceRecurringPriceCurrencyGyd ProductGetResponsePriceRecurringPriceCurrency = "GYD"
	ProductGetResponsePriceRecurringPriceCurrencyHkd ProductGetResponsePriceRecurringPriceCurrency = "HKD"
	ProductGetResponsePriceRecurringPriceCurrencyHnl ProductGetResponsePriceRecurringPriceCurrency = "HNL"
	ProductGetResponsePriceRecurringPriceCurrencyHrk ProductGetResponsePriceRecurringPriceCurrency = "HRK"
	ProductGetResponsePriceRecurringPriceCurrencyHtg ProductGetResponsePriceRecurringPriceCurrency = "HTG"
	ProductGetResponsePriceRecurringPriceCurrencyHuf ProductGetResponsePriceRecurringPriceCurrency = "HUF"
	ProductGetResponsePriceRecurringPriceCurrencyIdr ProductGetResponsePriceRecurringPriceCurrency = "IDR"
	ProductGetResponsePriceRecurringPriceCurrencyIls ProductGetResponsePriceRecurringPriceCurrency = "ILS"
	ProductGetResponsePriceRecurringPriceCurrencyInr ProductGetResponsePriceRecurringPriceCurrency = "INR"
	ProductGetResponsePriceRecurringPriceCurrencyIqd ProductGetResponsePriceRecurringPriceCurrency = "IQD"
	ProductGetResponsePriceRecurringPriceCurrencyJmd ProductGetResponsePriceRecurringPriceCurrency = "JMD"
	ProductGetResponsePriceRecurringPriceCurrencyJod ProductGetResponsePriceRecurringPriceCurrency = "JOD"
	ProductGetResponsePriceRecurringPriceCurrencyJpy ProductGetResponsePriceRecurringPriceCurrency = "JPY"
	ProductGetResponsePriceRecurringPriceCurrencyKes ProductGetResponsePriceRecurringPriceCurrency = "KES"
	ProductGetResponsePriceRecurringPriceCurrencyKgs ProductGetResponsePriceRecurringPriceCurrency = "KGS"
	ProductGetResponsePriceRecurringPriceCurrencyKhr ProductGetResponsePriceRecurringPriceCurrency = "KHR"
	ProductGetResponsePriceRecurringPriceCurrencyKmf ProductGetResponsePriceRecurringPriceCurrency = "KMF"
	ProductGetResponsePriceRecurringPriceCurrencyKrw ProductGetResponsePriceRecurringPriceCurrency = "KRW"
	ProductGetResponsePriceRecurringPriceCurrencyKwd ProductGetResponsePriceRecurringPriceCurrency = "KWD"
	ProductGetResponsePriceRecurringPriceCurrencyKyd ProductGetResponsePriceRecurringPriceCurrency = "KYD"
	ProductGetResponsePriceRecurringPriceCurrencyKzt ProductGetResponsePriceRecurringPriceCurrency = "KZT"
	ProductGetResponsePriceRecurringPriceCurrencyLak ProductGetResponsePriceRecurringPriceCurrency = "LAK"
	ProductGetResponsePriceRecurringPriceCurrencyLbp ProductGetResponsePriceRecurringPriceCurrency = "LBP"
	ProductGetResponsePriceRecurringPriceCurrencyLkr ProductGetResponsePriceRecurringPriceCurrency = "LKR"
	ProductGetResponsePriceRecurringPriceCurrencyLrd ProductGetResponsePriceRecurringPriceCurrency = "LRD"
	ProductGetResponsePriceRecurringPriceCurrencyLsl ProductGetResponsePriceRecurringPriceCurrency = "LSL"
	ProductGetResponsePriceRecurringPriceCurrencyLyd ProductGetResponsePriceRecurringPriceCurrency = "LYD"
	ProductGetResponsePriceRecurringPriceCurrencyMad ProductGetResponsePriceRecurringPriceCurrency = "MAD"
	ProductGetResponsePriceRecurringPriceCurrencyMdl ProductGetResponsePriceRecurringPriceCurrency = "MDL"
	ProductGetResponsePriceRecurringPriceCurrencyMga ProductGetResponsePriceRecurringPriceCurrency = "MGA"
	ProductGetResponsePriceRecurringPriceCurrencyMkd ProductGetResponsePriceRecurringPriceCurrency = "MKD"
	ProductGetResponsePriceRecurringPriceCurrencyMmk ProductGetResponsePriceRecurringPriceCurrency = "MMK"
	ProductGetResponsePriceRecurringPriceCurrencyMnt ProductGetResponsePriceRecurringPriceCurrency = "MNT"
	ProductGetResponsePriceRecurringPriceCurrencyMop ProductGetResponsePriceRecurringPriceCurrency = "MOP"
	ProductGetResponsePriceRecurringPriceCurrencyMru ProductGetResponsePriceRecurringPriceCurrency = "MRU"
	ProductGetResponsePriceRecurringPriceCurrencyMur ProductGetResponsePriceRecurringPriceCurrency = "MUR"
	ProductGetResponsePriceRecurringPriceCurrencyMvr ProductGetResponsePriceRecurringPriceCurrency = "MVR"
	ProductGetResponsePriceRecurringPriceCurrencyMwk ProductGetResponsePriceRecurringPriceCurrency = "MWK"
	ProductGetResponsePriceRecurringPriceCurrencyMxn ProductGetResponsePriceRecurringPriceCurrency = "MXN"
	ProductGetResponsePriceRecurringPriceCurrencyMyr ProductGetResponsePriceRecurringPriceCurrency = "MYR"
	ProductGetResponsePriceRecurringPriceCurrencyMzn ProductGetResponsePriceRecurringPriceCurrency = "MZN"
	ProductGetResponsePriceRecurringPriceCurrencyNad ProductGetResponsePriceRecurringPriceCurrency = "NAD"
	ProductGetResponsePriceRecurringPriceCurrencyNgn ProductGetResponsePriceRecurringPriceCurrency = "NGN"
	ProductGetResponsePriceRecurringPriceCurrencyNio ProductGetResponsePriceRecurringPriceCurrency = "NIO"
	ProductGetResponsePriceRecurringPriceCurrencyNok ProductGetResponsePriceRecurringPriceCurrency = "NOK"
	ProductGetResponsePriceRecurringPriceCurrencyNpr ProductGetResponsePriceRecurringPriceCurrency = "NPR"
	ProductGetResponsePriceRecurringPriceCurrencyNzd ProductGetResponsePriceRecurringPriceCurrency = "NZD"
	ProductGetResponsePriceRecurringPriceCurrencyOmr ProductGetResponsePriceRecurringPriceCurrency = "OMR"
	ProductGetResponsePriceRecurringPriceCurrencyPab ProductGetResponsePriceRecurringPriceCurrency = "PAB"
	ProductGetResponsePriceRecurringPriceCurrencyPen ProductGetResponsePriceRecurringPriceCurrency = "PEN"
	ProductGetResponsePriceRecurringPriceCurrencyPgk ProductGetResponsePriceRecurringPriceCurrency = "PGK"
	ProductGetResponsePriceRecurringPriceCurrencyPhp ProductGetResponsePriceRecurringPriceCurrency = "PHP"
	ProductGetResponsePriceRecurringPriceCurrencyPkr ProductGetResponsePriceRecurringPriceCurrency = "PKR"
	ProductGetResponsePriceRecurringPriceCurrencyPln ProductGetResponsePriceRecurringPriceCurrency = "PLN"
	ProductGetResponsePriceRecurringPriceCurrencyPyg ProductGetResponsePriceRecurringPriceCurrency = "PYG"
	ProductGetResponsePriceRecurringPriceCurrencyQar ProductGetResponsePriceRecurringPriceCurrency = "QAR"
	ProductGetResponsePriceRecurringPriceCurrencyRon ProductGetResponsePriceRecurringPriceCurrency = "RON"
	ProductGetResponsePriceRecurringPriceCurrencyRsd ProductGetResponsePriceRecurringPriceCurrency = "RSD"
	ProductGetResponsePriceRecurringPriceCurrencyRub ProductGetResponsePriceRecurringPriceCurrency = "RUB"
	ProductGetResponsePriceRecurringPriceCurrencyRwf ProductGetResponsePriceRecurringPriceCurrency = "RWF"
	ProductGetResponsePriceRecurringPriceCurrencySar ProductGetResponsePriceRecurringPriceCurrency = "SAR"
	ProductGetResponsePriceRecurringPriceCurrencySbd ProductGetResponsePriceRecurringPriceCurrency = "SBD"
	ProductGetResponsePriceRecurringPriceCurrencyScr ProductGetResponsePriceRecurringPriceCurrency = "SCR"
	ProductGetResponsePriceRecurringPriceCurrencySek ProductGetResponsePriceRecurringPriceCurrency = "SEK"
	ProductGetResponsePriceRecurringPriceCurrencySgd ProductGetResponsePriceRecurringPriceCurrency = "SGD"
	ProductGetResponsePriceRecurringPriceCurrencyShp ProductGetResponsePriceRecurringPriceCurrency = "SHP"
	ProductGetResponsePriceRecurringPriceCurrencySle ProductGetResponsePriceRecurringPriceCurrency = "SLE"
	ProductGetResponsePriceRecurringPriceCurrencySll ProductGetResponsePriceRecurringPriceCurrency = "SLL"
	ProductGetResponsePriceRecurringPriceCurrencySos ProductGetResponsePriceRecurringPriceCurrency = "SOS"
	ProductGetResponsePriceRecurringPriceCurrencySrd ProductGetResponsePriceRecurringPriceCurrency = "SRD"
	ProductGetResponsePriceRecurringPriceCurrencySsp ProductGetResponsePriceRecurringPriceCurrency = "SSP"
	ProductGetResponsePriceRecurringPriceCurrencyStn ProductGetResponsePriceRecurringPriceCurrency = "STN"
	ProductGetResponsePriceRecurringPriceCurrencySvc ProductGetResponsePriceRecurringPriceCurrency = "SVC"
	ProductGetResponsePriceRecurringPriceCurrencySzl ProductGetResponsePriceRecurringPriceCurrency = "SZL"
	ProductGetResponsePriceRecurringPriceCurrencyThb ProductGetResponsePriceRecurringPriceCurrency = "THB"
	ProductGetResponsePriceRecurringPriceCurrencyTnd ProductGetResponsePriceRecurringPriceCurrency = "TND"
	ProductGetResponsePriceRecurringPriceCurrencyTop ProductGetResponsePriceRecurringPriceCurrency = "TOP"
	ProductGetResponsePriceRecurringPriceCurrencyTry ProductGetResponsePriceRecurringPriceCurrency = "TRY"
	ProductGetResponsePriceRecurringPriceCurrencyTtd ProductGetResponsePriceRecurringPriceCurrency = "TTD"
	ProductGetResponsePriceRecurringPriceCurrencyTwd ProductGetResponsePriceRecurringPriceCurrency = "TWD"
	ProductGetResponsePriceRecurringPriceCurrencyTzs ProductGetResponsePriceRecurringPriceCurrency = "TZS"
	ProductGetResponsePriceRecurringPriceCurrencyUah ProductGetResponsePriceRecurringPriceCurrency = "UAH"
	ProductGetResponsePriceRecurringPriceCurrencyUgx ProductGetResponsePriceRecurringPriceCurrency = "UGX"
	ProductGetResponsePriceRecurringPriceCurrencyUsd ProductGetResponsePriceRecurringPriceCurrency = "USD"
	ProductGetResponsePriceRecurringPriceCurrencyUyu ProductGetResponsePriceRecurringPriceCurrency = "UYU"
	ProductGetResponsePriceRecurringPriceCurrencyUzs ProductGetResponsePriceRecurringPriceCurrency = "UZS"
	ProductGetResponsePriceRecurringPriceCurrencyVes ProductGetResponsePriceRecurringPriceCurrency = "VES"
	ProductGetResponsePriceRecurringPriceCurrencyVnd ProductGetResponsePriceRecurringPriceCurrency = "VND"
	ProductGetResponsePriceRecurringPriceCurrencyVuv ProductGetResponsePriceRecurringPriceCurrency = "VUV"
	ProductGetResponsePriceRecurringPriceCurrencyWst ProductGetResponsePriceRecurringPriceCurrency = "WST"
	ProductGetResponsePriceRecurringPriceCurrencyXaf ProductGetResponsePriceRecurringPriceCurrency = "XAF"
	ProductGetResponsePriceRecurringPriceCurrencyXcd ProductGetResponsePriceRecurringPriceCurrency = "XCD"
	ProductGetResponsePriceRecurringPriceCurrencyXof ProductGetResponsePriceRecurringPriceCurrency = "XOF"
	ProductGetResponsePriceRecurringPriceCurrencyXpf ProductGetResponsePriceRecurringPriceCurrency = "XPF"
	ProductGetResponsePriceRecurringPriceCurrencyYer ProductGetResponsePriceRecurringPriceCurrency = "YER"
	ProductGetResponsePriceRecurringPriceCurrencyZar ProductGetResponsePriceRecurringPriceCurrency = "ZAR"
	ProductGetResponsePriceRecurringPriceCurrencyZmw ProductGetResponsePriceRecurringPriceCurrency = "ZMW"
)

func (r ProductGetResponsePriceRecurringPriceCurrency) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceRecurringPriceCurrencyAed, ProductGetResponsePriceRecurringPriceCurrencyAll, ProductGetResponsePriceRecurringPriceCurrencyAmd, ProductGetResponsePriceRecurringPriceCurrencyAng, ProductGetResponsePriceRecurringPriceCurrencyAoa, ProductGetResponsePriceRecurringPriceCurrencyArs, ProductGetResponsePriceRecurringPriceCurrencyAud, ProductGetResponsePriceRecurringPriceCurrencyAwg, ProductGetResponsePriceRecurringPriceCurrencyAzn, ProductGetResponsePriceRecurringPriceCurrencyBam, ProductGetResponsePriceRecurringPriceCurrencyBbd, ProductGetResponsePriceRecurringPriceCurrencyBdt, ProductGetResponsePriceRecurringPriceCurrencyBgn, ProductGetResponsePriceRecurringPriceCurrencyBhd, ProductGetResponsePriceRecurringPriceCurrencyBif, ProductGetResponsePriceRecurringPriceCurrencyBmd, ProductGetResponsePriceRecurringPriceCurrencyBnd, ProductGetResponsePriceRecurringPriceCurrencyBob, ProductGetResponsePriceRecurringPriceCurrencyBrl, ProductGetResponsePriceRecurringPriceCurrencyBsd, ProductGetResponsePriceRecurringPriceCurrencyBwp, ProductGetResponsePriceRecurringPriceCurrencyByn, ProductGetResponsePriceRecurringPriceCurrencyBzd, ProductGetResponsePriceRecurringPriceCurrencyCad, ProductGetResponsePriceRecurringPriceCurrencyChf, ProductGetResponsePriceRecurringPriceCurrencyClp, ProductGetResponsePriceRecurringPriceCurrencyCny, ProductGetResponsePriceRecurringPriceCurrencyCop, ProductGetResponsePriceRecurringPriceCurrencyCrc, ProductGetResponsePriceRecurringPriceCurrencyCup, ProductGetResponsePriceRecurringPriceCurrencyCve, ProductGetResponsePriceRecurringPriceCurrencyCzk, ProductGetResponsePriceRecurringPriceCurrencyDjf, ProductGetResponsePriceRecurringPriceCurrencyDkk, ProductGetResponsePriceRecurringPriceCurrencyDop, ProductGetResponsePriceRecurringPriceCurrencyDzd, ProductGetResponsePriceRecurringPriceCurrencyEgp, ProductGetResponsePriceRecurringPriceCurrencyEtb, ProductGetResponsePriceRecurringPriceCurrencyEur, ProductGetResponsePriceRecurringPriceCurrencyFjd, ProductGetResponsePriceRecurringPriceCurrencyFkp, ProductGetResponsePriceRecurringPriceCurrencyGbp, ProductGetResponsePriceRecurringPriceCurrencyGel, ProductGetResponsePriceRecurringPriceCurrencyGhs, ProductGetResponsePriceRecurringPriceCurrencyGip, ProductGetResponsePriceRecurringPriceCurrencyGmd, ProductGetResponsePriceRecurringPriceCurrencyGnf, ProductGetResponsePriceRecurringPriceCurrencyGtq, ProductGetResponsePriceRecurringPriceCurrencyGyd, ProductGetResponsePriceRecurringPriceCurrencyHkd, ProductGetResponsePriceRecurringPriceCurrencyHnl, ProductGetResponsePriceRecurringPriceCurrencyHrk, ProductGetResponsePriceRecurringPriceCurrencyHtg, ProductGetResponsePriceRecurringPriceCurrencyHuf, ProductGetResponsePriceRecurringPriceCurrencyIdr, ProductGetResponsePriceRecurringPriceCurrencyIls, ProductGetResponsePriceRecurringPriceCurrencyInr, ProductGetResponsePriceRecurringPriceCurrencyIqd, ProductGetResponsePriceRecurringPriceCurrencyJmd, ProductGetResponsePriceRecurringPriceCurrencyJod, ProductGetResponsePriceRecurringPriceCurrencyJpy, ProductGetResponsePriceRecurringPriceCurrencyKes, ProductGetResponsePriceRecurringPriceCurrencyKgs, ProductGetResponsePriceRecurringPriceCurrencyKhr, ProductGetResponsePriceRecurringPriceCurrencyKmf, ProductGetResponsePriceRecurringPriceCurrencyKrw, ProductGetResponsePriceRecurringPriceCurrencyKwd, ProductGetResponsePriceRecurringPriceCurrencyKyd, ProductGetResponsePriceRecurringPriceCurrencyKzt, ProductGetResponsePriceRecurringPriceCurrencyLak, ProductGetResponsePriceRecurringPriceCurrencyLbp, ProductGetResponsePriceRecurringPriceCurrencyLkr, ProductGetResponsePriceRecurringPriceCurrencyLrd, ProductGetResponsePriceRecurringPriceCurrencyLsl, ProductGetResponsePriceRecurringPriceCurrencyLyd, ProductGetResponsePriceRecurringPriceCurrencyMad, ProductGetResponsePriceRecurringPriceCurrencyMdl, ProductGetResponsePriceRecurringPriceCurrencyMga, ProductGetResponsePriceRecurringPriceCurrencyMkd, ProductGetResponsePriceRecurringPriceCurrencyMmk, ProductGetResponsePriceRecurringPriceCurrencyMnt, ProductGetResponsePriceRecurringPriceCurrencyMop, ProductGetResponsePriceRecurringPriceCurrencyMru, ProductGetResponsePriceRecurringPriceCurrencyMur, ProductGetResponsePriceRecurringPriceCurrencyMvr, ProductGetResponsePriceRecurringPriceCurrencyMwk, ProductGetResponsePriceRecurringPriceCurrencyMxn, ProductGetResponsePriceRecurringPriceCurrencyMyr, ProductGetResponsePriceRecurringPriceCurrencyMzn, ProductGetResponsePriceRecurringPriceCurrencyNad, ProductGetResponsePriceRecurringPriceCurrencyNgn, ProductGetResponsePriceRecurringPriceCurrencyNio, ProductGetResponsePriceRecurringPriceCurrencyNok, ProductGetResponsePriceRecurringPriceCurrencyNpr, ProductGetResponsePriceRecurringPriceCurrencyNzd, ProductGetResponsePriceRecurringPriceCurrencyOmr, ProductGetResponsePriceRecurringPriceCurrencyPab, ProductGetResponsePriceRecurringPriceCurrencyPen, ProductGetResponsePriceRecurringPriceCurrencyPgk, ProductGetResponsePriceRecurringPriceCurrencyPhp, ProductGetResponsePriceRecurringPriceCurrencyPkr, ProductGetResponsePriceRecurringPriceCurrencyPln, ProductGetResponsePriceRecurringPriceCurrencyPyg, ProductGetResponsePriceRecurringPriceCurrencyQar, ProductGetResponsePriceRecurringPriceCurrencyRon, ProductGetResponsePriceRecurringPriceCurrencyRsd, ProductGetResponsePriceRecurringPriceCurrencyRub, ProductGetResponsePriceRecurringPriceCurrencyRwf, ProductGetResponsePriceRecurringPriceCurrencySar, ProductGetResponsePriceRecurringPriceCurrencySbd, ProductGetResponsePriceRecurringPriceCurrencyScr, ProductGetResponsePriceRecurringPriceCurrencySek, ProductGetResponsePriceRecurringPriceCurrencySgd, ProductGetResponsePriceRecurringPriceCurrencyShp, ProductGetResponsePriceRecurringPriceCurrencySle, ProductGetResponsePriceRecurringPriceCurrencySll, ProductGetResponsePriceRecurringPriceCurrencySos, ProductGetResponsePriceRecurringPriceCurrencySrd, ProductGetResponsePriceRecurringPriceCurrencySsp, ProductGetResponsePriceRecurringPriceCurrencyStn, ProductGetResponsePriceRecurringPriceCurrencySvc, ProductGetResponsePriceRecurringPriceCurrencySzl, ProductGetResponsePriceRecurringPriceCurrencyThb, ProductGetResponsePriceRecurringPriceCurrencyTnd, ProductGetResponsePriceRecurringPriceCurrencyTop, ProductGetResponsePriceRecurringPriceCurrencyTry, ProductGetResponsePriceRecurringPriceCurrencyTtd, ProductGetResponsePriceRecurringPriceCurrencyTwd, ProductGetResponsePriceRecurringPriceCurrencyTzs, ProductGetResponsePriceRecurringPriceCurrencyUah, ProductGetResponsePriceRecurringPriceCurrencyUgx, ProductGetResponsePriceRecurringPriceCurrencyUsd, ProductGetResponsePriceRecurringPriceCurrencyUyu, ProductGetResponsePriceRecurringPriceCurrencyUzs, ProductGetResponsePriceRecurringPriceCurrencyVes, ProductGetResponsePriceRecurringPriceCurrencyVnd, ProductGetResponsePriceRecurringPriceCurrencyVuv, ProductGetResponsePriceRecurringPriceCurrencyWst, ProductGetResponsePriceRecurringPriceCurrencyXaf, ProductGetResponsePriceRecurringPriceCurrencyXcd, ProductGetResponsePriceRecurringPriceCurrencyXof, ProductGetResponsePriceRecurringPriceCurrencyXpf, ProductGetResponsePriceRecurringPriceCurrencyYer, ProductGetResponsePriceRecurringPriceCurrencyZar, ProductGetResponsePriceRecurringPriceCurrencyZmw:
		return true
	}
	return false
}

type ProductGetResponsePriceRecurringPricePaymentFrequencyInterval string

const (
	ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalDay   ProductGetResponsePriceRecurringPricePaymentFrequencyInterval = "Day"
	ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalWeek  ProductGetResponsePriceRecurringPricePaymentFrequencyInterval = "Week"
	ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalMonth ProductGetResponsePriceRecurringPricePaymentFrequencyInterval = "Month"
	ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalYear  ProductGetResponsePriceRecurringPricePaymentFrequencyInterval = "Year"
)

func (r ProductGetResponsePriceRecurringPricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalDay, ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalWeek, ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalMonth, ProductGetResponsePriceRecurringPricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductGetResponsePriceRecurringPriceSubscriptionPeriodInterval string

const (
	ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalDay   ProductGetResponsePriceRecurringPriceSubscriptionPeriodInterval = "Day"
	ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalWeek  ProductGetResponsePriceRecurringPriceSubscriptionPeriodInterval = "Week"
	ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalMonth ProductGetResponsePriceRecurringPriceSubscriptionPeriodInterval = "Month"
	ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalYear  ProductGetResponsePriceRecurringPriceSubscriptionPeriodInterval = "Year"
)

func (r ProductGetResponsePriceRecurringPriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalDay, ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalWeek, ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalMonth, ProductGetResponsePriceRecurringPriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

type ProductGetResponsePriceRecurringPriceType string

const (
	ProductGetResponsePriceRecurringPriceTypeRecurringPrice ProductGetResponsePriceRecurringPriceType = "recurring_price"
)

func (r ProductGetResponsePriceRecurringPriceType) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceRecurringPriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductGetResponsePriceCurrency string

const (
	ProductGetResponsePriceCurrencyAed ProductGetResponsePriceCurrency = "AED"
	ProductGetResponsePriceCurrencyAll ProductGetResponsePriceCurrency = "ALL"
	ProductGetResponsePriceCurrencyAmd ProductGetResponsePriceCurrency = "AMD"
	ProductGetResponsePriceCurrencyAng ProductGetResponsePriceCurrency = "ANG"
	ProductGetResponsePriceCurrencyAoa ProductGetResponsePriceCurrency = "AOA"
	ProductGetResponsePriceCurrencyArs ProductGetResponsePriceCurrency = "ARS"
	ProductGetResponsePriceCurrencyAud ProductGetResponsePriceCurrency = "AUD"
	ProductGetResponsePriceCurrencyAwg ProductGetResponsePriceCurrency = "AWG"
	ProductGetResponsePriceCurrencyAzn ProductGetResponsePriceCurrency = "AZN"
	ProductGetResponsePriceCurrencyBam ProductGetResponsePriceCurrency = "BAM"
	ProductGetResponsePriceCurrencyBbd ProductGetResponsePriceCurrency = "BBD"
	ProductGetResponsePriceCurrencyBdt ProductGetResponsePriceCurrency = "BDT"
	ProductGetResponsePriceCurrencyBgn ProductGetResponsePriceCurrency = "BGN"
	ProductGetResponsePriceCurrencyBhd ProductGetResponsePriceCurrency = "BHD"
	ProductGetResponsePriceCurrencyBif ProductGetResponsePriceCurrency = "BIF"
	ProductGetResponsePriceCurrencyBmd ProductGetResponsePriceCurrency = "BMD"
	ProductGetResponsePriceCurrencyBnd ProductGetResponsePriceCurrency = "BND"
	ProductGetResponsePriceCurrencyBob ProductGetResponsePriceCurrency = "BOB"
	ProductGetResponsePriceCurrencyBrl ProductGetResponsePriceCurrency = "BRL"
	ProductGetResponsePriceCurrencyBsd ProductGetResponsePriceCurrency = "BSD"
	ProductGetResponsePriceCurrencyBwp ProductGetResponsePriceCurrency = "BWP"
	ProductGetResponsePriceCurrencyByn ProductGetResponsePriceCurrency = "BYN"
	ProductGetResponsePriceCurrencyBzd ProductGetResponsePriceCurrency = "BZD"
	ProductGetResponsePriceCurrencyCad ProductGetResponsePriceCurrency = "CAD"
	ProductGetResponsePriceCurrencyChf ProductGetResponsePriceCurrency = "CHF"
	ProductGetResponsePriceCurrencyClp ProductGetResponsePriceCurrency = "CLP"
	ProductGetResponsePriceCurrencyCny ProductGetResponsePriceCurrency = "CNY"
	ProductGetResponsePriceCurrencyCop ProductGetResponsePriceCurrency = "COP"
	ProductGetResponsePriceCurrencyCrc ProductGetResponsePriceCurrency = "CRC"
	ProductGetResponsePriceCurrencyCup ProductGetResponsePriceCurrency = "CUP"
	ProductGetResponsePriceCurrencyCve ProductGetResponsePriceCurrency = "CVE"
	ProductGetResponsePriceCurrencyCzk ProductGetResponsePriceCurrency = "CZK"
	ProductGetResponsePriceCurrencyDjf ProductGetResponsePriceCurrency = "DJF"
	ProductGetResponsePriceCurrencyDkk ProductGetResponsePriceCurrency = "DKK"
	ProductGetResponsePriceCurrencyDop ProductGetResponsePriceCurrency = "DOP"
	ProductGetResponsePriceCurrencyDzd ProductGetResponsePriceCurrency = "DZD"
	ProductGetResponsePriceCurrencyEgp ProductGetResponsePriceCurrency = "EGP"
	ProductGetResponsePriceCurrencyEtb ProductGetResponsePriceCurrency = "ETB"
	ProductGetResponsePriceCurrencyEur ProductGetResponsePriceCurrency = "EUR"
	ProductGetResponsePriceCurrencyFjd ProductGetResponsePriceCurrency = "FJD"
	ProductGetResponsePriceCurrencyFkp ProductGetResponsePriceCurrency = "FKP"
	ProductGetResponsePriceCurrencyGbp ProductGetResponsePriceCurrency = "GBP"
	ProductGetResponsePriceCurrencyGel ProductGetResponsePriceCurrency = "GEL"
	ProductGetResponsePriceCurrencyGhs ProductGetResponsePriceCurrency = "GHS"
	ProductGetResponsePriceCurrencyGip ProductGetResponsePriceCurrency = "GIP"
	ProductGetResponsePriceCurrencyGmd ProductGetResponsePriceCurrency = "GMD"
	ProductGetResponsePriceCurrencyGnf ProductGetResponsePriceCurrency = "GNF"
	ProductGetResponsePriceCurrencyGtq ProductGetResponsePriceCurrency = "GTQ"
	ProductGetResponsePriceCurrencyGyd ProductGetResponsePriceCurrency = "GYD"
	ProductGetResponsePriceCurrencyHkd ProductGetResponsePriceCurrency = "HKD"
	ProductGetResponsePriceCurrencyHnl ProductGetResponsePriceCurrency = "HNL"
	ProductGetResponsePriceCurrencyHrk ProductGetResponsePriceCurrency = "HRK"
	ProductGetResponsePriceCurrencyHtg ProductGetResponsePriceCurrency = "HTG"
	ProductGetResponsePriceCurrencyHuf ProductGetResponsePriceCurrency = "HUF"
	ProductGetResponsePriceCurrencyIdr ProductGetResponsePriceCurrency = "IDR"
	ProductGetResponsePriceCurrencyIls ProductGetResponsePriceCurrency = "ILS"
	ProductGetResponsePriceCurrencyInr ProductGetResponsePriceCurrency = "INR"
	ProductGetResponsePriceCurrencyIqd ProductGetResponsePriceCurrency = "IQD"
	ProductGetResponsePriceCurrencyJmd ProductGetResponsePriceCurrency = "JMD"
	ProductGetResponsePriceCurrencyJod ProductGetResponsePriceCurrency = "JOD"
	ProductGetResponsePriceCurrencyJpy ProductGetResponsePriceCurrency = "JPY"
	ProductGetResponsePriceCurrencyKes ProductGetResponsePriceCurrency = "KES"
	ProductGetResponsePriceCurrencyKgs ProductGetResponsePriceCurrency = "KGS"
	ProductGetResponsePriceCurrencyKhr ProductGetResponsePriceCurrency = "KHR"
	ProductGetResponsePriceCurrencyKmf ProductGetResponsePriceCurrency = "KMF"
	ProductGetResponsePriceCurrencyKrw ProductGetResponsePriceCurrency = "KRW"
	ProductGetResponsePriceCurrencyKwd ProductGetResponsePriceCurrency = "KWD"
	ProductGetResponsePriceCurrencyKyd ProductGetResponsePriceCurrency = "KYD"
	ProductGetResponsePriceCurrencyKzt ProductGetResponsePriceCurrency = "KZT"
	ProductGetResponsePriceCurrencyLak ProductGetResponsePriceCurrency = "LAK"
	ProductGetResponsePriceCurrencyLbp ProductGetResponsePriceCurrency = "LBP"
	ProductGetResponsePriceCurrencyLkr ProductGetResponsePriceCurrency = "LKR"
	ProductGetResponsePriceCurrencyLrd ProductGetResponsePriceCurrency = "LRD"
	ProductGetResponsePriceCurrencyLsl ProductGetResponsePriceCurrency = "LSL"
	ProductGetResponsePriceCurrencyLyd ProductGetResponsePriceCurrency = "LYD"
	ProductGetResponsePriceCurrencyMad ProductGetResponsePriceCurrency = "MAD"
	ProductGetResponsePriceCurrencyMdl ProductGetResponsePriceCurrency = "MDL"
	ProductGetResponsePriceCurrencyMga ProductGetResponsePriceCurrency = "MGA"
	ProductGetResponsePriceCurrencyMkd ProductGetResponsePriceCurrency = "MKD"
	ProductGetResponsePriceCurrencyMmk ProductGetResponsePriceCurrency = "MMK"
	ProductGetResponsePriceCurrencyMnt ProductGetResponsePriceCurrency = "MNT"
	ProductGetResponsePriceCurrencyMop ProductGetResponsePriceCurrency = "MOP"
	ProductGetResponsePriceCurrencyMru ProductGetResponsePriceCurrency = "MRU"
	ProductGetResponsePriceCurrencyMur ProductGetResponsePriceCurrency = "MUR"
	ProductGetResponsePriceCurrencyMvr ProductGetResponsePriceCurrency = "MVR"
	ProductGetResponsePriceCurrencyMwk ProductGetResponsePriceCurrency = "MWK"
	ProductGetResponsePriceCurrencyMxn ProductGetResponsePriceCurrency = "MXN"
	ProductGetResponsePriceCurrencyMyr ProductGetResponsePriceCurrency = "MYR"
	ProductGetResponsePriceCurrencyMzn ProductGetResponsePriceCurrency = "MZN"
	ProductGetResponsePriceCurrencyNad ProductGetResponsePriceCurrency = "NAD"
	ProductGetResponsePriceCurrencyNgn ProductGetResponsePriceCurrency = "NGN"
	ProductGetResponsePriceCurrencyNio ProductGetResponsePriceCurrency = "NIO"
	ProductGetResponsePriceCurrencyNok ProductGetResponsePriceCurrency = "NOK"
	ProductGetResponsePriceCurrencyNpr ProductGetResponsePriceCurrency = "NPR"
	ProductGetResponsePriceCurrencyNzd ProductGetResponsePriceCurrency = "NZD"
	ProductGetResponsePriceCurrencyOmr ProductGetResponsePriceCurrency = "OMR"
	ProductGetResponsePriceCurrencyPab ProductGetResponsePriceCurrency = "PAB"
	ProductGetResponsePriceCurrencyPen ProductGetResponsePriceCurrency = "PEN"
	ProductGetResponsePriceCurrencyPgk ProductGetResponsePriceCurrency = "PGK"
	ProductGetResponsePriceCurrencyPhp ProductGetResponsePriceCurrency = "PHP"
	ProductGetResponsePriceCurrencyPkr ProductGetResponsePriceCurrency = "PKR"
	ProductGetResponsePriceCurrencyPln ProductGetResponsePriceCurrency = "PLN"
	ProductGetResponsePriceCurrencyPyg ProductGetResponsePriceCurrency = "PYG"
	ProductGetResponsePriceCurrencyQar ProductGetResponsePriceCurrency = "QAR"
	ProductGetResponsePriceCurrencyRon ProductGetResponsePriceCurrency = "RON"
	ProductGetResponsePriceCurrencyRsd ProductGetResponsePriceCurrency = "RSD"
	ProductGetResponsePriceCurrencyRub ProductGetResponsePriceCurrency = "RUB"
	ProductGetResponsePriceCurrencyRwf ProductGetResponsePriceCurrency = "RWF"
	ProductGetResponsePriceCurrencySar ProductGetResponsePriceCurrency = "SAR"
	ProductGetResponsePriceCurrencySbd ProductGetResponsePriceCurrency = "SBD"
	ProductGetResponsePriceCurrencyScr ProductGetResponsePriceCurrency = "SCR"
	ProductGetResponsePriceCurrencySek ProductGetResponsePriceCurrency = "SEK"
	ProductGetResponsePriceCurrencySgd ProductGetResponsePriceCurrency = "SGD"
	ProductGetResponsePriceCurrencyShp ProductGetResponsePriceCurrency = "SHP"
	ProductGetResponsePriceCurrencySle ProductGetResponsePriceCurrency = "SLE"
	ProductGetResponsePriceCurrencySll ProductGetResponsePriceCurrency = "SLL"
	ProductGetResponsePriceCurrencySos ProductGetResponsePriceCurrency = "SOS"
	ProductGetResponsePriceCurrencySrd ProductGetResponsePriceCurrency = "SRD"
	ProductGetResponsePriceCurrencySsp ProductGetResponsePriceCurrency = "SSP"
	ProductGetResponsePriceCurrencyStn ProductGetResponsePriceCurrency = "STN"
	ProductGetResponsePriceCurrencySvc ProductGetResponsePriceCurrency = "SVC"
	ProductGetResponsePriceCurrencySzl ProductGetResponsePriceCurrency = "SZL"
	ProductGetResponsePriceCurrencyThb ProductGetResponsePriceCurrency = "THB"
	ProductGetResponsePriceCurrencyTnd ProductGetResponsePriceCurrency = "TND"
	ProductGetResponsePriceCurrencyTop ProductGetResponsePriceCurrency = "TOP"
	ProductGetResponsePriceCurrencyTry ProductGetResponsePriceCurrency = "TRY"
	ProductGetResponsePriceCurrencyTtd ProductGetResponsePriceCurrency = "TTD"
	ProductGetResponsePriceCurrencyTwd ProductGetResponsePriceCurrency = "TWD"
	ProductGetResponsePriceCurrencyTzs ProductGetResponsePriceCurrency = "TZS"
	ProductGetResponsePriceCurrencyUah ProductGetResponsePriceCurrency = "UAH"
	ProductGetResponsePriceCurrencyUgx ProductGetResponsePriceCurrency = "UGX"
	ProductGetResponsePriceCurrencyUsd ProductGetResponsePriceCurrency = "USD"
	ProductGetResponsePriceCurrencyUyu ProductGetResponsePriceCurrency = "UYU"
	ProductGetResponsePriceCurrencyUzs ProductGetResponsePriceCurrency = "UZS"
	ProductGetResponsePriceCurrencyVes ProductGetResponsePriceCurrency = "VES"
	ProductGetResponsePriceCurrencyVnd ProductGetResponsePriceCurrency = "VND"
	ProductGetResponsePriceCurrencyVuv ProductGetResponsePriceCurrency = "VUV"
	ProductGetResponsePriceCurrencyWst ProductGetResponsePriceCurrency = "WST"
	ProductGetResponsePriceCurrencyXaf ProductGetResponsePriceCurrency = "XAF"
	ProductGetResponsePriceCurrencyXcd ProductGetResponsePriceCurrency = "XCD"
	ProductGetResponsePriceCurrencyXof ProductGetResponsePriceCurrency = "XOF"
	ProductGetResponsePriceCurrencyXpf ProductGetResponsePriceCurrency = "XPF"
	ProductGetResponsePriceCurrencyYer ProductGetResponsePriceCurrency = "YER"
	ProductGetResponsePriceCurrencyZar ProductGetResponsePriceCurrency = "ZAR"
	ProductGetResponsePriceCurrencyZmw ProductGetResponsePriceCurrency = "ZMW"
)

func (r ProductGetResponsePriceCurrency) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceCurrencyAed, ProductGetResponsePriceCurrencyAll, ProductGetResponsePriceCurrencyAmd, ProductGetResponsePriceCurrencyAng, ProductGetResponsePriceCurrencyAoa, ProductGetResponsePriceCurrencyArs, ProductGetResponsePriceCurrencyAud, ProductGetResponsePriceCurrencyAwg, ProductGetResponsePriceCurrencyAzn, ProductGetResponsePriceCurrencyBam, ProductGetResponsePriceCurrencyBbd, ProductGetResponsePriceCurrencyBdt, ProductGetResponsePriceCurrencyBgn, ProductGetResponsePriceCurrencyBhd, ProductGetResponsePriceCurrencyBif, ProductGetResponsePriceCurrencyBmd, ProductGetResponsePriceCurrencyBnd, ProductGetResponsePriceCurrencyBob, ProductGetResponsePriceCurrencyBrl, ProductGetResponsePriceCurrencyBsd, ProductGetResponsePriceCurrencyBwp, ProductGetResponsePriceCurrencyByn, ProductGetResponsePriceCurrencyBzd, ProductGetResponsePriceCurrencyCad, ProductGetResponsePriceCurrencyChf, ProductGetResponsePriceCurrencyClp, ProductGetResponsePriceCurrencyCny, ProductGetResponsePriceCurrencyCop, ProductGetResponsePriceCurrencyCrc, ProductGetResponsePriceCurrencyCup, ProductGetResponsePriceCurrencyCve, ProductGetResponsePriceCurrencyCzk, ProductGetResponsePriceCurrencyDjf, ProductGetResponsePriceCurrencyDkk, ProductGetResponsePriceCurrencyDop, ProductGetResponsePriceCurrencyDzd, ProductGetResponsePriceCurrencyEgp, ProductGetResponsePriceCurrencyEtb, ProductGetResponsePriceCurrencyEur, ProductGetResponsePriceCurrencyFjd, ProductGetResponsePriceCurrencyFkp, ProductGetResponsePriceCurrencyGbp, ProductGetResponsePriceCurrencyGel, ProductGetResponsePriceCurrencyGhs, ProductGetResponsePriceCurrencyGip, ProductGetResponsePriceCurrencyGmd, ProductGetResponsePriceCurrencyGnf, ProductGetResponsePriceCurrencyGtq, ProductGetResponsePriceCurrencyGyd, ProductGetResponsePriceCurrencyHkd, ProductGetResponsePriceCurrencyHnl, ProductGetResponsePriceCurrencyHrk, ProductGetResponsePriceCurrencyHtg, ProductGetResponsePriceCurrencyHuf, ProductGetResponsePriceCurrencyIdr, ProductGetResponsePriceCurrencyIls, ProductGetResponsePriceCurrencyInr, ProductGetResponsePriceCurrencyIqd, ProductGetResponsePriceCurrencyJmd, ProductGetResponsePriceCurrencyJod, ProductGetResponsePriceCurrencyJpy, ProductGetResponsePriceCurrencyKes, ProductGetResponsePriceCurrencyKgs, ProductGetResponsePriceCurrencyKhr, ProductGetResponsePriceCurrencyKmf, ProductGetResponsePriceCurrencyKrw, ProductGetResponsePriceCurrencyKwd, ProductGetResponsePriceCurrencyKyd, ProductGetResponsePriceCurrencyKzt, ProductGetResponsePriceCurrencyLak, ProductGetResponsePriceCurrencyLbp, ProductGetResponsePriceCurrencyLkr, ProductGetResponsePriceCurrencyLrd, ProductGetResponsePriceCurrencyLsl, ProductGetResponsePriceCurrencyLyd, ProductGetResponsePriceCurrencyMad, ProductGetResponsePriceCurrencyMdl, ProductGetResponsePriceCurrencyMga, ProductGetResponsePriceCurrencyMkd, ProductGetResponsePriceCurrencyMmk, ProductGetResponsePriceCurrencyMnt, ProductGetResponsePriceCurrencyMop, ProductGetResponsePriceCurrencyMru, ProductGetResponsePriceCurrencyMur, ProductGetResponsePriceCurrencyMvr, ProductGetResponsePriceCurrencyMwk, ProductGetResponsePriceCurrencyMxn, ProductGetResponsePriceCurrencyMyr, ProductGetResponsePriceCurrencyMzn, ProductGetResponsePriceCurrencyNad, ProductGetResponsePriceCurrencyNgn, ProductGetResponsePriceCurrencyNio, ProductGetResponsePriceCurrencyNok, ProductGetResponsePriceCurrencyNpr, ProductGetResponsePriceCurrencyNzd, ProductGetResponsePriceCurrencyOmr, ProductGetResponsePriceCurrencyPab, ProductGetResponsePriceCurrencyPen, ProductGetResponsePriceCurrencyPgk, ProductGetResponsePriceCurrencyPhp, ProductGetResponsePriceCurrencyPkr, ProductGetResponsePriceCurrencyPln, ProductGetResponsePriceCurrencyPyg, ProductGetResponsePriceCurrencyQar, ProductGetResponsePriceCurrencyRon, ProductGetResponsePriceCurrencyRsd, ProductGetResponsePriceCurrencyRub, ProductGetResponsePriceCurrencyRwf, ProductGetResponsePriceCurrencySar, ProductGetResponsePriceCurrencySbd, ProductGetResponsePriceCurrencyScr, ProductGetResponsePriceCurrencySek, ProductGetResponsePriceCurrencySgd, ProductGetResponsePriceCurrencyShp, ProductGetResponsePriceCurrencySle, ProductGetResponsePriceCurrencySll, ProductGetResponsePriceCurrencySos, ProductGetResponsePriceCurrencySrd, ProductGetResponsePriceCurrencySsp, ProductGetResponsePriceCurrencyStn, ProductGetResponsePriceCurrencySvc, ProductGetResponsePriceCurrencySzl, ProductGetResponsePriceCurrencyThb, ProductGetResponsePriceCurrencyTnd, ProductGetResponsePriceCurrencyTop, ProductGetResponsePriceCurrencyTry, ProductGetResponsePriceCurrencyTtd, ProductGetResponsePriceCurrencyTwd, ProductGetResponsePriceCurrencyTzs, ProductGetResponsePriceCurrencyUah, ProductGetResponsePriceCurrencyUgx, ProductGetResponsePriceCurrencyUsd, ProductGetResponsePriceCurrencyUyu, ProductGetResponsePriceCurrencyUzs, ProductGetResponsePriceCurrencyVes, ProductGetResponsePriceCurrencyVnd, ProductGetResponsePriceCurrencyVuv, ProductGetResponsePriceCurrencyWst, ProductGetResponsePriceCurrencyXaf, ProductGetResponsePriceCurrencyXcd, ProductGetResponsePriceCurrencyXof, ProductGetResponsePriceCurrencyXpf, ProductGetResponsePriceCurrencyYer, ProductGetResponsePriceCurrencyZar, ProductGetResponsePriceCurrencyZmw:
		return true
	}
	return false
}

type ProductGetResponsePriceType string

const (
	ProductGetResponsePriceTypeOneTimePrice   ProductGetResponsePriceType = "one_time_price"
	ProductGetResponsePriceTypeRecurringPrice ProductGetResponsePriceType = "recurring_price"
)

func (r ProductGetResponsePriceType) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceTypeOneTimePrice, ProductGetResponsePriceTypeRecurringPrice:
		return true
	}
	return false
}

type ProductGetResponsePricePaymentFrequencyInterval string

const (
	ProductGetResponsePricePaymentFrequencyIntervalDay   ProductGetResponsePricePaymentFrequencyInterval = "Day"
	ProductGetResponsePricePaymentFrequencyIntervalWeek  ProductGetResponsePricePaymentFrequencyInterval = "Week"
	ProductGetResponsePricePaymentFrequencyIntervalMonth ProductGetResponsePricePaymentFrequencyInterval = "Month"
	ProductGetResponsePricePaymentFrequencyIntervalYear  ProductGetResponsePricePaymentFrequencyInterval = "Year"
)

func (r ProductGetResponsePricePaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case ProductGetResponsePricePaymentFrequencyIntervalDay, ProductGetResponsePricePaymentFrequencyIntervalWeek, ProductGetResponsePricePaymentFrequencyIntervalMonth, ProductGetResponsePricePaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type ProductGetResponsePriceSubscriptionPeriodInterval string

const (
	ProductGetResponsePriceSubscriptionPeriodIntervalDay   ProductGetResponsePriceSubscriptionPeriodInterval = "Day"
	ProductGetResponsePriceSubscriptionPeriodIntervalWeek  ProductGetResponsePriceSubscriptionPeriodInterval = "Week"
	ProductGetResponsePriceSubscriptionPeriodIntervalMonth ProductGetResponsePriceSubscriptionPeriodInterval = "Month"
	ProductGetResponsePriceSubscriptionPeriodIntervalYear  ProductGetResponsePriceSubscriptionPeriodInterval = "Year"
)

func (r ProductGetResponsePriceSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case ProductGetResponsePriceSubscriptionPeriodIntervalDay, ProductGetResponsePriceSubscriptionPeriodIntervalWeek, ProductGetResponsePriceSubscriptionPeriodIntervalMonth, ProductGetResponsePriceSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

// Represents the different categories of taxation applicable to various products
// and services.
type ProductGetResponseTaxCategory string

const (
	ProductGetResponseTaxCategoryDigitalProducts ProductGetResponseTaxCategory = "digital_products"
	ProductGetResponseTaxCategorySaas            ProductGetResponseTaxCategory = "saas"
	ProductGetResponseTaxCategoryEBook           ProductGetResponseTaxCategory = "e_book"
)

func (r ProductGetResponseTaxCategory) IsKnown() bool {
	switch r {
	case ProductGetResponseTaxCategoryDigitalProducts, ProductGetResponseTaxCategorySaas, ProductGetResponseTaxCategoryEBook:
		return true
	}
	return false
}

type ProductListResponse struct {
	Items []ProductListResponseItem `json:"items,required"`
	JSON  productListResponseJSON   `json:"-"`
}

// productListResponseJSON contains the JSON metadata for the struct
// [ProductListResponse]
type productListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseJSON) RawJSON() string {
	return r.raw
}

type ProductListResponseItem struct {
	BusinessID  string    `json:"business_id,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	IsRecurring bool      `json:"is_recurring,required"`
	ProductID   string    `json:"product_id,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory ProductListResponseItemsTaxCategory `json:"tax_category,required"`
	UpdatedAt   time.Time                           `json:"updated_at,required" format:"date-time"`
	Description string                              `json:"description,nullable"`
	Image       string                              `json:"image,nullable"`
	Name        string                              `json:"name,nullable"`
	Price       int64                               `json:"price,nullable"`
	JSON        productListResponseItemJSON         `json:"-"`
}

// productListResponseItemJSON contains the JSON metadata for the struct
// [ProductListResponseItem]
type productListResponseItemJSON struct {
	BusinessID  apijson.Field
	CreatedAt   apijson.Field
	IsRecurring apijson.Field
	ProductID   apijson.Field
	TaxCategory apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	Image       apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Represents the different categories of taxation applicable to various products
// and services.
type ProductListResponseItemsTaxCategory string

const (
	ProductListResponseItemsTaxCategoryDigitalProducts ProductListResponseItemsTaxCategory = "digital_products"
	ProductListResponseItemsTaxCategorySaas            ProductListResponseItemsTaxCategory = "saas"
	ProductListResponseItemsTaxCategoryEBook           ProductListResponseItemsTaxCategory = "e_book"
)

func (r ProductListResponseItemsTaxCategory) IsKnown() bool {
	switch r {
	case ProductListResponseItemsTaxCategoryDigitalProducts, ProductListResponseItemsTaxCategorySaas, ProductListResponseItemsTaxCategoryEBook:
		return true
	}
	return false
}

type ProductNewParams struct {
	Price param.Field[ProductNewParamsPriceUnion] `json:"price,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[ProductNewParamsTaxCategory] `json:"tax_category,required"`
	Description param.Field[string]                      `json:"description"`
	Name        param.Field[string]                      `json:"name"`
}

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductNewParamsPrice struct {
	Currency param.Field[ProductNewParamsPriceCurrency] `json:"currency,required"`
	Discount param.Field[float64]                       `json:"discount,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                      param.Field[int64]                                           `json:"price,required"`
	PurchasingPowerParity      param.Field[bool]                                            `json:"purchasing_power_parity,required"`
	Type                       param.Field[ProductNewParamsPriceType]                       `json:"type,required"`
	PaymentFrequencyCount      param.Field[int64]                                           `json:"payment_frequency_count"`
	PaymentFrequencyInterval   param.Field[ProductNewParamsPricePaymentFrequencyInterval]   `json:"payment_frequency_interval"`
	SubscriptionPeriodCount    param.Field[int64]                                           `json:"subscription_period_count"`
	SubscriptionPeriodInterval param.Field[ProductNewParamsPriceSubscriptionPeriodInterval] `json:"subscription_period_interval"`
	TrialPeriodDays            param.Field[int64]                                           `json:"trial_period_days"`
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
	Discount param.Field[float64]                                   `json:"discount,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                 param.Field[int64]                                 `json:"price,required"`
	PurchasingPowerParity param.Field[bool]                                  `json:"purchasing_power_parity,required"`
	Type                  param.Field[ProductNewParamsPriceOneTimePriceType] `json:"type,required"`
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
	Currency                 param.Field[ProductNewParamsPriceRecurringPriceCurrency]                 `json:"currency,required"`
	Discount                 param.Field[float64]                                                     `json:"discount,required"`
	PaymentFrequencyCount    param.Field[int64]                                                       `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval param.Field[ProductNewParamsPriceRecurringPricePaymentFrequencyInterval] `json:"payment_frequency_interval,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                      param.Field[int64]                                                         `json:"price,required"`
	PurchasingPowerParity      param.Field[bool]                                                          `json:"purchasing_power_parity,required"`
	SubscriptionPeriodCount    param.Field[int64]                                                         `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval param.Field[ProductNewParamsPriceRecurringPriceSubscriptionPeriodInterval] `json:"subscription_period_interval,required"`
	TrialPeriodDays            param.Field[int64]                                                         `json:"trial_period_days,required"`
	Type                       param.Field[ProductNewParamsPriceRecurringPriceType]                       `json:"type,required"`
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

type ProductUpdateParams struct {
	Description param.Field[string]                        `json:"description"`
	Name        param.Field[string]                        `json:"name"`
	Price       param.Field[ProductUpdateParamsPriceUnion] `json:"price"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[ProductUpdateParamsTaxCategory] `json:"tax_category"`
}

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductUpdateParamsPrice struct {
	Currency param.Field[ProductUpdateParamsPriceCurrency] `json:"currency,required"`
	Discount param.Field[float64]                          `json:"discount,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                      param.Field[int64]                                              `json:"price,required"`
	PurchasingPowerParity      param.Field[bool]                                               `json:"purchasing_power_parity,required"`
	Type                       param.Field[ProductUpdateParamsPriceType]                       `json:"type,required"`
	PaymentFrequencyCount      param.Field[int64]                                              `json:"payment_frequency_count"`
	PaymentFrequencyInterval   param.Field[ProductUpdateParamsPricePaymentFrequencyInterval]   `json:"payment_frequency_interval"`
	SubscriptionPeriodCount    param.Field[int64]                                              `json:"subscription_period_count"`
	SubscriptionPeriodInterval param.Field[ProductUpdateParamsPriceSubscriptionPeriodInterval] `json:"subscription_period_interval"`
	TrialPeriodDays            param.Field[int64]                                              `json:"trial_period_days"`
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
	Discount param.Field[float64]                                      `json:"discount,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                 param.Field[int64]                                    `json:"price,required"`
	PurchasingPowerParity param.Field[bool]                                     `json:"purchasing_power_parity,required"`
	Type                  param.Field[ProductUpdateParamsPriceOneTimePriceType] `json:"type,required"`
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
	Currency                 param.Field[ProductUpdateParamsPriceRecurringPriceCurrency]                 `json:"currency,required"`
	Discount                 param.Field[float64]                                                        `json:"discount,required"`
	PaymentFrequencyCount    param.Field[int64]                                                          `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval param.Field[ProductUpdateParamsPriceRecurringPricePaymentFrequencyInterval] `json:"payment_frequency_interval,required"`
	// The payment amount. Amount for the payment in the lowest denomination of the
	// currency, (i.e) in cents for USD denomination. E.g., Pass 100 to charge $1.00
	Price                      param.Field[int64]                                                            `json:"price,required"`
	PurchasingPowerParity      param.Field[bool]                                                             `json:"purchasing_power_parity,required"`
	SubscriptionPeriodCount    param.Field[int64]                                                            `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval param.Field[ProductUpdateParamsPriceRecurringPriceSubscriptionPeriodInterval] `json:"subscription_period_interval,required"`
	TrialPeriodDays            param.Field[int64]                                                            `json:"trial_period_days,required"`
	Type                       param.Field[ProductUpdateParamsPriceRecurringPriceType]                       `json:"type,required"`
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
