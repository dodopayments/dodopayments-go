// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/dodo-payments-go/internal/apijson"
	"github.com/stainless-sdks/dodo-payments-go/internal/apiquery"
	"github.com/stainless-sdks/dodo-payments-go/internal/param"
	"github.com/stainless-sdks/dodo-payments-go/internal/requestconfig"
	"github.com/stainless-sdks/dodo-payments-go/option"
	"github.com/stainless-sdks/dodo-payments-go/packages/pagination"
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
	opts = append(r.Options[:], opts...)
	path := "subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *SubscriptionService) Get(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Subscription], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

func (r *SubscriptionService) ListAutoPaging(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Subscription] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type Subscription struct {
	CreatedAt                  time.Time                              `json:"created_at,required" format:"date-time"`
	Currency                   SubscriptionCurrency                   `json:"currency,required"`
	Customer                   SubscriptionCustomer                   `json:"customer,required"`
	NextBillingDate            time.Time                              `json:"next_billing_date,required" format:"date-time"`
	PaymentFrequencyCount      int64                                  `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval   SubscriptionPaymentFrequencyInterval   `json:"payment_frequency_interval,required"`
	ProductID                  string                                 `json:"product_id,required"`
	Quantity                   int64                                  `json:"quantity,required"`
	RecurringPreTaxAmount      int64                                  `json:"recurring_pre_tax_amount,required"`
	Status                     SubscriptionStatus                     `json:"status,required"`
	SubscriptionID             string                                 `json:"subscription_id,required"`
	SubscriptionPeriodCount    int64                                  `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval SubscriptionSubscriptionPeriodInterval `json:"subscription_period_interval,required"`
	TrialPeriodDays            int64                                  `json:"trial_period_days,required"`
	JSON                       subscriptionJSON                       `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	CreatedAt                  apijson.Field
	Currency                   apijson.Field
	Customer                   apijson.Field
	NextBillingDate            apijson.Field
	PaymentFrequencyCount      apijson.Field
	PaymentFrequencyInterval   apijson.Field
	ProductID                  apijson.Field
	Quantity                   apijson.Field
	RecurringPreTaxAmount      apijson.Field
	Status                     apijson.Field
	SubscriptionID             apijson.Field
	SubscriptionPeriodCount    apijson.Field
	SubscriptionPeriodInterval apijson.Field
	TrialPeriodDays            apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionCurrency string

const (
	SubscriptionCurrencyAed SubscriptionCurrency = "AED"
	SubscriptionCurrencyAll SubscriptionCurrency = "ALL"
	SubscriptionCurrencyAmd SubscriptionCurrency = "AMD"
	SubscriptionCurrencyAng SubscriptionCurrency = "ANG"
	SubscriptionCurrencyAoa SubscriptionCurrency = "AOA"
	SubscriptionCurrencyArs SubscriptionCurrency = "ARS"
	SubscriptionCurrencyAud SubscriptionCurrency = "AUD"
	SubscriptionCurrencyAwg SubscriptionCurrency = "AWG"
	SubscriptionCurrencyAzn SubscriptionCurrency = "AZN"
	SubscriptionCurrencyBam SubscriptionCurrency = "BAM"
	SubscriptionCurrencyBbd SubscriptionCurrency = "BBD"
	SubscriptionCurrencyBdt SubscriptionCurrency = "BDT"
	SubscriptionCurrencyBgn SubscriptionCurrency = "BGN"
	SubscriptionCurrencyBhd SubscriptionCurrency = "BHD"
	SubscriptionCurrencyBif SubscriptionCurrency = "BIF"
	SubscriptionCurrencyBmd SubscriptionCurrency = "BMD"
	SubscriptionCurrencyBnd SubscriptionCurrency = "BND"
	SubscriptionCurrencyBob SubscriptionCurrency = "BOB"
	SubscriptionCurrencyBrl SubscriptionCurrency = "BRL"
	SubscriptionCurrencyBsd SubscriptionCurrency = "BSD"
	SubscriptionCurrencyBwp SubscriptionCurrency = "BWP"
	SubscriptionCurrencyByn SubscriptionCurrency = "BYN"
	SubscriptionCurrencyBzd SubscriptionCurrency = "BZD"
	SubscriptionCurrencyCad SubscriptionCurrency = "CAD"
	SubscriptionCurrencyChf SubscriptionCurrency = "CHF"
	SubscriptionCurrencyClp SubscriptionCurrency = "CLP"
	SubscriptionCurrencyCny SubscriptionCurrency = "CNY"
	SubscriptionCurrencyCop SubscriptionCurrency = "COP"
	SubscriptionCurrencyCrc SubscriptionCurrency = "CRC"
	SubscriptionCurrencyCup SubscriptionCurrency = "CUP"
	SubscriptionCurrencyCve SubscriptionCurrency = "CVE"
	SubscriptionCurrencyCzk SubscriptionCurrency = "CZK"
	SubscriptionCurrencyDjf SubscriptionCurrency = "DJF"
	SubscriptionCurrencyDkk SubscriptionCurrency = "DKK"
	SubscriptionCurrencyDop SubscriptionCurrency = "DOP"
	SubscriptionCurrencyDzd SubscriptionCurrency = "DZD"
	SubscriptionCurrencyEgp SubscriptionCurrency = "EGP"
	SubscriptionCurrencyEtb SubscriptionCurrency = "ETB"
	SubscriptionCurrencyEur SubscriptionCurrency = "EUR"
	SubscriptionCurrencyFjd SubscriptionCurrency = "FJD"
	SubscriptionCurrencyFkp SubscriptionCurrency = "FKP"
	SubscriptionCurrencyGbp SubscriptionCurrency = "GBP"
	SubscriptionCurrencyGel SubscriptionCurrency = "GEL"
	SubscriptionCurrencyGhs SubscriptionCurrency = "GHS"
	SubscriptionCurrencyGip SubscriptionCurrency = "GIP"
	SubscriptionCurrencyGmd SubscriptionCurrency = "GMD"
	SubscriptionCurrencyGnf SubscriptionCurrency = "GNF"
	SubscriptionCurrencyGtq SubscriptionCurrency = "GTQ"
	SubscriptionCurrencyGyd SubscriptionCurrency = "GYD"
	SubscriptionCurrencyHkd SubscriptionCurrency = "HKD"
	SubscriptionCurrencyHnl SubscriptionCurrency = "HNL"
	SubscriptionCurrencyHrk SubscriptionCurrency = "HRK"
	SubscriptionCurrencyHtg SubscriptionCurrency = "HTG"
	SubscriptionCurrencyHuf SubscriptionCurrency = "HUF"
	SubscriptionCurrencyIdr SubscriptionCurrency = "IDR"
	SubscriptionCurrencyIls SubscriptionCurrency = "ILS"
	SubscriptionCurrencyInr SubscriptionCurrency = "INR"
	SubscriptionCurrencyIqd SubscriptionCurrency = "IQD"
	SubscriptionCurrencyJmd SubscriptionCurrency = "JMD"
	SubscriptionCurrencyJod SubscriptionCurrency = "JOD"
	SubscriptionCurrencyJpy SubscriptionCurrency = "JPY"
	SubscriptionCurrencyKes SubscriptionCurrency = "KES"
	SubscriptionCurrencyKgs SubscriptionCurrency = "KGS"
	SubscriptionCurrencyKhr SubscriptionCurrency = "KHR"
	SubscriptionCurrencyKmf SubscriptionCurrency = "KMF"
	SubscriptionCurrencyKrw SubscriptionCurrency = "KRW"
	SubscriptionCurrencyKwd SubscriptionCurrency = "KWD"
	SubscriptionCurrencyKyd SubscriptionCurrency = "KYD"
	SubscriptionCurrencyKzt SubscriptionCurrency = "KZT"
	SubscriptionCurrencyLak SubscriptionCurrency = "LAK"
	SubscriptionCurrencyLbp SubscriptionCurrency = "LBP"
	SubscriptionCurrencyLkr SubscriptionCurrency = "LKR"
	SubscriptionCurrencyLrd SubscriptionCurrency = "LRD"
	SubscriptionCurrencyLsl SubscriptionCurrency = "LSL"
	SubscriptionCurrencyLyd SubscriptionCurrency = "LYD"
	SubscriptionCurrencyMad SubscriptionCurrency = "MAD"
	SubscriptionCurrencyMdl SubscriptionCurrency = "MDL"
	SubscriptionCurrencyMga SubscriptionCurrency = "MGA"
	SubscriptionCurrencyMkd SubscriptionCurrency = "MKD"
	SubscriptionCurrencyMmk SubscriptionCurrency = "MMK"
	SubscriptionCurrencyMnt SubscriptionCurrency = "MNT"
	SubscriptionCurrencyMop SubscriptionCurrency = "MOP"
	SubscriptionCurrencyMru SubscriptionCurrency = "MRU"
	SubscriptionCurrencyMur SubscriptionCurrency = "MUR"
	SubscriptionCurrencyMvr SubscriptionCurrency = "MVR"
	SubscriptionCurrencyMwk SubscriptionCurrency = "MWK"
	SubscriptionCurrencyMxn SubscriptionCurrency = "MXN"
	SubscriptionCurrencyMyr SubscriptionCurrency = "MYR"
	SubscriptionCurrencyMzn SubscriptionCurrency = "MZN"
	SubscriptionCurrencyNad SubscriptionCurrency = "NAD"
	SubscriptionCurrencyNgn SubscriptionCurrency = "NGN"
	SubscriptionCurrencyNio SubscriptionCurrency = "NIO"
	SubscriptionCurrencyNok SubscriptionCurrency = "NOK"
	SubscriptionCurrencyNpr SubscriptionCurrency = "NPR"
	SubscriptionCurrencyNzd SubscriptionCurrency = "NZD"
	SubscriptionCurrencyOmr SubscriptionCurrency = "OMR"
	SubscriptionCurrencyPab SubscriptionCurrency = "PAB"
	SubscriptionCurrencyPen SubscriptionCurrency = "PEN"
	SubscriptionCurrencyPgk SubscriptionCurrency = "PGK"
	SubscriptionCurrencyPhp SubscriptionCurrency = "PHP"
	SubscriptionCurrencyPkr SubscriptionCurrency = "PKR"
	SubscriptionCurrencyPln SubscriptionCurrency = "PLN"
	SubscriptionCurrencyPyg SubscriptionCurrency = "PYG"
	SubscriptionCurrencyQar SubscriptionCurrency = "QAR"
	SubscriptionCurrencyRon SubscriptionCurrency = "RON"
	SubscriptionCurrencyRsd SubscriptionCurrency = "RSD"
	SubscriptionCurrencyRub SubscriptionCurrency = "RUB"
	SubscriptionCurrencyRwf SubscriptionCurrency = "RWF"
	SubscriptionCurrencySar SubscriptionCurrency = "SAR"
	SubscriptionCurrencySbd SubscriptionCurrency = "SBD"
	SubscriptionCurrencyScr SubscriptionCurrency = "SCR"
	SubscriptionCurrencySek SubscriptionCurrency = "SEK"
	SubscriptionCurrencySgd SubscriptionCurrency = "SGD"
	SubscriptionCurrencyShp SubscriptionCurrency = "SHP"
	SubscriptionCurrencySle SubscriptionCurrency = "SLE"
	SubscriptionCurrencySll SubscriptionCurrency = "SLL"
	SubscriptionCurrencySos SubscriptionCurrency = "SOS"
	SubscriptionCurrencySrd SubscriptionCurrency = "SRD"
	SubscriptionCurrencySsp SubscriptionCurrency = "SSP"
	SubscriptionCurrencyStn SubscriptionCurrency = "STN"
	SubscriptionCurrencySvc SubscriptionCurrency = "SVC"
	SubscriptionCurrencySzl SubscriptionCurrency = "SZL"
	SubscriptionCurrencyThb SubscriptionCurrency = "THB"
	SubscriptionCurrencyTnd SubscriptionCurrency = "TND"
	SubscriptionCurrencyTop SubscriptionCurrency = "TOP"
	SubscriptionCurrencyTry SubscriptionCurrency = "TRY"
	SubscriptionCurrencyTtd SubscriptionCurrency = "TTD"
	SubscriptionCurrencyTwd SubscriptionCurrency = "TWD"
	SubscriptionCurrencyTzs SubscriptionCurrency = "TZS"
	SubscriptionCurrencyUah SubscriptionCurrency = "UAH"
	SubscriptionCurrencyUgx SubscriptionCurrency = "UGX"
	SubscriptionCurrencyUsd SubscriptionCurrency = "USD"
	SubscriptionCurrencyUyu SubscriptionCurrency = "UYU"
	SubscriptionCurrencyUzs SubscriptionCurrency = "UZS"
	SubscriptionCurrencyVes SubscriptionCurrency = "VES"
	SubscriptionCurrencyVnd SubscriptionCurrency = "VND"
	SubscriptionCurrencyVuv SubscriptionCurrency = "VUV"
	SubscriptionCurrencyWst SubscriptionCurrency = "WST"
	SubscriptionCurrencyXaf SubscriptionCurrency = "XAF"
	SubscriptionCurrencyXcd SubscriptionCurrency = "XCD"
	SubscriptionCurrencyXof SubscriptionCurrency = "XOF"
	SubscriptionCurrencyXpf SubscriptionCurrency = "XPF"
	SubscriptionCurrencyYer SubscriptionCurrency = "YER"
	SubscriptionCurrencyZar SubscriptionCurrency = "ZAR"
	SubscriptionCurrencyZmw SubscriptionCurrency = "ZMW"
)

func (r SubscriptionCurrency) IsKnown() bool {
	switch r {
	case SubscriptionCurrencyAed, SubscriptionCurrencyAll, SubscriptionCurrencyAmd, SubscriptionCurrencyAng, SubscriptionCurrencyAoa, SubscriptionCurrencyArs, SubscriptionCurrencyAud, SubscriptionCurrencyAwg, SubscriptionCurrencyAzn, SubscriptionCurrencyBam, SubscriptionCurrencyBbd, SubscriptionCurrencyBdt, SubscriptionCurrencyBgn, SubscriptionCurrencyBhd, SubscriptionCurrencyBif, SubscriptionCurrencyBmd, SubscriptionCurrencyBnd, SubscriptionCurrencyBob, SubscriptionCurrencyBrl, SubscriptionCurrencyBsd, SubscriptionCurrencyBwp, SubscriptionCurrencyByn, SubscriptionCurrencyBzd, SubscriptionCurrencyCad, SubscriptionCurrencyChf, SubscriptionCurrencyClp, SubscriptionCurrencyCny, SubscriptionCurrencyCop, SubscriptionCurrencyCrc, SubscriptionCurrencyCup, SubscriptionCurrencyCve, SubscriptionCurrencyCzk, SubscriptionCurrencyDjf, SubscriptionCurrencyDkk, SubscriptionCurrencyDop, SubscriptionCurrencyDzd, SubscriptionCurrencyEgp, SubscriptionCurrencyEtb, SubscriptionCurrencyEur, SubscriptionCurrencyFjd, SubscriptionCurrencyFkp, SubscriptionCurrencyGbp, SubscriptionCurrencyGel, SubscriptionCurrencyGhs, SubscriptionCurrencyGip, SubscriptionCurrencyGmd, SubscriptionCurrencyGnf, SubscriptionCurrencyGtq, SubscriptionCurrencyGyd, SubscriptionCurrencyHkd, SubscriptionCurrencyHnl, SubscriptionCurrencyHrk, SubscriptionCurrencyHtg, SubscriptionCurrencyHuf, SubscriptionCurrencyIdr, SubscriptionCurrencyIls, SubscriptionCurrencyInr, SubscriptionCurrencyIqd, SubscriptionCurrencyJmd, SubscriptionCurrencyJod, SubscriptionCurrencyJpy, SubscriptionCurrencyKes, SubscriptionCurrencyKgs, SubscriptionCurrencyKhr, SubscriptionCurrencyKmf, SubscriptionCurrencyKrw, SubscriptionCurrencyKwd, SubscriptionCurrencyKyd, SubscriptionCurrencyKzt, SubscriptionCurrencyLak, SubscriptionCurrencyLbp, SubscriptionCurrencyLkr, SubscriptionCurrencyLrd, SubscriptionCurrencyLsl, SubscriptionCurrencyLyd, SubscriptionCurrencyMad, SubscriptionCurrencyMdl, SubscriptionCurrencyMga, SubscriptionCurrencyMkd, SubscriptionCurrencyMmk, SubscriptionCurrencyMnt, SubscriptionCurrencyMop, SubscriptionCurrencyMru, SubscriptionCurrencyMur, SubscriptionCurrencyMvr, SubscriptionCurrencyMwk, SubscriptionCurrencyMxn, SubscriptionCurrencyMyr, SubscriptionCurrencyMzn, SubscriptionCurrencyNad, SubscriptionCurrencyNgn, SubscriptionCurrencyNio, SubscriptionCurrencyNok, SubscriptionCurrencyNpr, SubscriptionCurrencyNzd, SubscriptionCurrencyOmr, SubscriptionCurrencyPab, SubscriptionCurrencyPen, SubscriptionCurrencyPgk, SubscriptionCurrencyPhp, SubscriptionCurrencyPkr, SubscriptionCurrencyPln, SubscriptionCurrencyPyg, SubscriptionCurrencyQar, SubscriptionCurrencyRon, SubscriptionCurrencyRsd, SubscriptionCurrencyRub, SubscriptionCurrencyRwf, SubscriptionCurrencySar, SubscriptionCurrencySbd, SubscriptionCurrencyScr, SubscriptionCurrencySek, SubscriptionCurrencySgd, SubscriptionCurrencyShp, SubscriptionCurrencySle, SubscriptionCurrencySll, SubscriptionCurrencySos, SubscriptionCurrencySrd, SubscriptionCurrencySsp, SubscriptionCurrencyStn, SubscriptionCurrencySvc, SubscriptionCurrencySzl, SubscriptionCurrencyThb, SubscriptionCurrencyTnd, SubscriptionCurrencyTop, SubscriptionCurrencyTry, SubscriptionCurrencyTtd, SubscriptionCurrencyTwd, SubscriptionCurrencyTzs, SubscriptionCurrencyUah, SubscriptionCurrencyUgx, SubscriptionCurrencyUsd, SubscriptionCurrencyUyu, SubscriptionCurrencyUzs, SubscriptionCurrencyVes, SubscriptionCurrencyVnd, SubscriptionCurrencyVuv, SubscriptionCurrencyWst, SubscriptionCurrencyXaf, SubscriptionCurrencyXcd, SubscriptionCurrencyXof, SubscriptionCurrencyXpf, SubscriptionCurrencyYer, SubscriptionCurrencyZar, SubscriptionCurrencyZmw:
		return true
	}
	return false
}

type SubscriptionCustomer struct {
	CustomerID string                   `json:"customer_id,required"`
	Email      string                   `json:"email,required"`
	Name       string                   `json:"name,required"`
	JSON       subscriptionCustomerJSON `json:"-"`
}

// subscriptionCustomerJSON contains the JSON metadata for the struct
// [SubscriptionCustomer]
type subscriptionCustomerJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCustomerJSON) RawJSON() string {
	return r.raw
}

type SubscriptionPaymentFrequencyInterval string

const (
	SubscriptionPaymentFrequencyIntervalDay   SubscriptionPaymentFrequencyInterval = "Day"
	SubscriptionPaymentFrequencyIntervalWeek  SubscriptionPaymentFrequencyInterval = "Week"
	SubscriptionPaymentFrequencyIntervalMonth SubscriptionPaymentFrequencyInterval = "Month"
	SubscriptionPaymentFrequencyIntervalYear  SubscriptionPaymentFrequencyInterval = "Year"
)

func (r SubscriptionPaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case SubscriptionPaymentFrequencyIntervalDay, SubscriptionPaymentFrequencyIntervalWeek, SubscriptionPaymentFrequencyIntervalMonth, SubscriptionPaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type SubscriptionStatus string

const (
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusOnHold    SubscriptionStatus = "on_hold"
	SubscriptionStatusPaused    SubscriptionStatus = "paused"
	SubscriptionStatusCancelled SubscriptionStatus = "cancelled"
	SubscriptionStatusFailed    SubscriptionStatus = "failed"
	SubscriptionStatusExpired   SubscriptionStatus = "expired"
)

func (r SubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionStatusPending, SubscriptionStatusActive, SubscriptionStatusOnHold, SubscriptionStatusPaused, SubscriptionStatusCancelled, SubscriptionStatusFailed, SubscriptionStatusExpired:
		return true
	}
	return false
}

type SubscriptionSubscriptionPeriodInterval string

const (
	SubscriptionSubscriptionPeriodIntervalDay   SubscriptionSubscriptionPeriodInterval = "Day"
	SubscriptionSubscriptionPeriodIntervalWeek  SubscriptionSubscriptionPeriodInterval = "Week"
	SubscriptionSubscriptionPeriodIntervalMonth SubscriptionSubscriptionPeriodInterval = "Month"
	SubscriptionSubscriptionPeriodIntervalYear  SubscriptionSubscriptionPeriodInterval = "Year"
)

func (r SubscriptionSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case SubscriptionSubscriptionPeriodIntervalDay, SubscriptionSubscriptionPeriodIntervalWeek, SubscriptionSubscriptionPeriodIntervalMonth, SubscriptionSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

type SubscriptionNewResponse struct {
	Customer              SubscriptionNewResponseCustomer `json:"customer,required"`
	RecurringPreTaxAmount int64                           `json:"recurring_pre_tax_amount,required"`
	SubscriptionID        string                          `json:"subscription_id,required"`
	ClientSecret          string                          `json:"client_secret,nullable"`
	PaymentLink           string                          `json:"payment_link,nullable"`
	JSON                  subscriptionNewResponseJSON     `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	Customer              apijson.Field
	RecurringPreTaxAmount apijson.Field
	SubscriptionID        apijson.Field
	ClientSecret          apijson.Field
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

type SubscriptionNewResponseCustomer struct {
	CustomerID string                              `json:"customer_id,required"`
	Email      string                              `json:"email,required"`
	Name       string                              `json:"name,required"`
	JSON       subscriptionNewResponseCustomerJSON `json:"-"`
}

// subscriptionNewResponseCustomerJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseCustomer]
type subscriptionNewResponseCustomerJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseCustomerJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewParams struct {
	Billing   param.Field[SubscriptionNewParamsBilling]  `json:"billing,required"`
	Customer  param.Field[SubscriptionNewParamsCustomer] `json:"customer,required"`
	ProductID param.Field[string]                        `json:"product_id,required"`
	Quantity  param.Field[int64]                         `json:"quantity,required"`
	// False by default
	PaymentLink param.Field[bool]   `json:"payment_link"`
	ReturnURL   param.Field[string] `json:"return_url"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsBilling struct {
	City param.Field[string] `json:"city,required"`
	// ISO country code alpha2 variant
	Country param.Field[CountryCode] `json:"country,required"`
	State   param.Field[string]      `json:"state,required"`
	Street  param.Field[string]      `json:"street,required"`
	Zipcode param.Field[int64]       `json:"zipcode,required"`
}

func (r SubscriptionNewParamsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsCustomer struct {
	Email       param.Field[string] `json:"email,required"`
	Name        param.Field[string] `json:"name,required"`
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r SubscriptionNewParamsCustomer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParams struct {
	Status param.Field[SubscriptionUpdateParamsStatus] `json:"status,required"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParamsStatus string

const (
	SubscriptionUpdateParamsStatusPending   SubscriptionUpdateParamsStatus = "pending"
	SubscriptionUpdateParamsStatusActive    SubscriptionUpdateParamsStatus = "active"
	SubscriptionUpdateParamsStatusOnHold    SubscriptionUpdateParamsStatus = "on_hold"
	SubscriptionUpdateParamsStatusPaused    SubscriptionUpdateParamsStatus = "paused"
	SubscriptionUpdateParamsStatusCancelled SubscriptionUpdateParamsStatus = "cancelled"
	SubscriptionUpdateParamsStatusFailed    SubscriptionUpdateParamsStatus = "failed"
	SubscriptionUpdateParamsStatusExpired   SubscriptionUpdateParamsStatus = "expired"
)

func (r SubscriptionUpdateParamsStatus) IsKnown() bool {
	switch r {
	case SubscriptionUpdateParamsStatusPending, SubscriptionUpdateParamsStatusActive, SubscriptionUpdateParamsStatusOnHold, SubscriptionUpdateParamsStatusPaused, SubscriptionUpdateParamsStatusCancelled, SubscriptionUpdateParamsStatusFailed, SubscriptionUpdateParamsStatusExpired:
		return true
	}
	return false
}

type SubscriptionListParams struct {
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
