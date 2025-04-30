// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
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

func (r *SubscriptionService) ChangePlan(ctx context.Context, subscriptionID string, body SubscriptionChangePlanParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/change-plan", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

func (r *SubscriptionService) Charge(ctx context.Context, subscriptionID string, body SubscriptionChargeParams, opts ...option.RequestOption) (res *SubscriptionChargeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/charge", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response struct representing subscription details
type Subscription struct {
	Billing BillingAddress `json:"billing,required"`
	// Timestamp when the subscription was created
	CreatedAt time.Time              `json:"created_at,required" format:"date-time"`
	Currency  SubscriptionCurrency   `json:"currency,required"`
	Customer  CustomerLimitedDetails `json:"customer,required"`
	Metadata  map[string]string      `json:"metadata,required"`
	// Timestamp of the next scheduled billing. Indicates the end of current billing
	// period
	NextBillingDate time.Time `json:"next_billing_date,required" format:"date-time"`
	// Wether the subscription is on-demand or not
	OnDemand bool `json:"on_demand,required"`
	// Number of payment frequency intervals
	PaymentFrequencyCount    int64        `json:"payment_frequency_count,required"`
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval,required"`
	// Timestamp of the last payment. Indicates the start of current billing period
	PreviousBillingDate time.Time `json:"previous_billing_date,required" format:"date-time"`
	// Identifier of the product associated with this subscription
	ProductID string `json:"product_id,required"`
	// Number of units/items included in the subscription
	Quantity int64 `json:"quantity,required"`
	// Amount charged before tax for each recurring payment in smallest currency unit
	// (e.g. cents)
	RecurringPreTaxAmount int64              `json:"recurring_pre_tax_amount,required"`
	Status                SubscriptionStatus `json:"status,required"`
	// Unique identifier for the subscription
	SubscriptionID string `json:"subscription_id,required"`
	// Number of subscription period intervals
	SubscriptionPeriodCount    int64        `json:"subscription_period_count,required"`
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval,required"`
	// Indicates if the recurring_pre_tax_amount is tax inclusive
	TaxInclusive bool `json:"tax_inclusive,required"`
	// Number of days in the trial period (0 if no trial)
	TrialPeriodDays int64 `json:"trial_period_days,required"`
	// Cancelled timestamp if the subscription is cancelled
	CancelledAt time.Time `json:"cancelled_at,nullable" format:"date-time"`
	// The discount id if discount is applied
	DiscountID string           `json:"discount_id,nullable"`
	JSON       subscriptionJSON `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	Billing                    apijson.Field
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
	DiscountID                 apijson.Field
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
	Customer CustomerLimitedDetails `json:"customer,required"`
	Metadata map[string]string      `json:"metadata,required"`
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
	// URL to checkout page
	PaymentLink string                      `json:"payment_link,nullable"`
	JSON        subscriptionNewResponseJSON `json:"-"`
}

// subscriptionNewResponseJSON contains the JSON metadata for the struct
// [SubscriptionNewResponse]
type subscriptionNewResponseJSON struct {
	Customer              apijson.Field
	Metadata              apijson.Field
	RecurringPreTaxAmount apijson.Field
	SubscriptionID        apijson.Field
	ClientSecret          apijson.Field
	DiscountID            apijson.Field
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

type SubscriptionChargeResponse struct {
	PaymentID string                         `json:"payment_id,required"`
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

type SubscriptionNewParams struct {
	Billing  param.Field[BillingAddressParam]       `json:"billing,required"`
	Customer param.Field[CustomerRequestUnionParam] `json:"customer,required"`
	// Unique identifier of the product to subscribe to
	ProductID param.Field[string] `json:"product_id,required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity param.Field[int64] `json:"quantity,required"`
	// List of payment methods allowed during checkout.
	//
	// Customers will **never** see payment methods that are **not** in this list.
	// However, adding a method here **does not guarantee** customers will see it.
	// Availability still depends on other factors (e.g., customer location, merchant
	// settings).
	AllowedPaymentMethodTypes param.Field[[]SubscriptionNewParamsAllowedPaymentMethodType] `json:"allowed_payment_method_types"`
	BillingCurrency           param.Field[SubscriptionNewParamsBillingCurrency]            `json:"billing_currency"`
	// Discount Code to apply to the subscription
	DiscountCode param.Field[string]                        `json:"discount_code"`
	Metadata     param.Field[map[string]string]             `json:"metadata"`
	OnDemand     param.Field[SubscriptionNewParamsOnDemand] `json:"on_demand"`
	// If true, generates a payment link. Defaults to false if not specified.
	PaymentLink param.Field[bool] `json:"payment_link"`
	// Optional URL to redirect after successful subscription creation
	ReturnURL param.Field[string] `json:"return_url"`
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

type SubscriptionNewParamsAllowedPaymentMethodType string

const (
	SubscriptionNewParamsAllowedPaymentMethodTypeCredit           SubscriptionNewParamsAllowedPaymentMethodType = "credit"
	SubscriptionNewParamsAllowedPaymentMethodTypeDebit            SubscriptionNewParamsAllowedPaymentMethodType = "debit"
	SubscriptionNewParamsAllowedPaymentMethodTypeUpiCollect       SubscriptionNewParamsAllowedPaymentMethodType = "upi_collect"
	SubscriptionNewParamsAllowedPaymentMethodTypeUpiIntent        SubscriptionNewParamsAllowedPaymentMethodType = "upi_intent"
	SubscriptionNewParamsAllowedPaymentMethodTypeApplePay         SubscriptionNewParamsAllowedPaymentMethodType = "apple_pay"
	SubscriptionNewParamsAllowedPaymentMethodTypeCashapp          SubscriptionNewParamsAllowedPaymentMethodType = "cashapp"
	SubscriptionNewParamsAllowedPaymentMethodTypeGooglePay        SubscriptionNewParamsAllowedPaymentMethodType = "google_pay"
	SubscriptionNewParamsAllowedPaymentMethodTypeMultibanco       SubscriptionNewParamsAllowedPaymentMethodType = "multibanco"
	SubscriptionNewParamsAllowedPaymentMethodTypeBancontactCard   SubscriptionNewParamsAllowedPaymentMethodType = "bancontact_card"
	SubscriptionNewParamsAllowedPaymentMethodTypeEps              SubscriptionNewParamsAllowedPaymentMethodType = "eps"
	SubscriptionNewParamsAllowedPaymentMethodTypeIdeal            SubscriptionNewParamsAllowedPaymentMethodType = "ideal"
	SubscriptionNewParamsAllowedPaymentMethodTypePrzelewy24       SubscriptionNewParamsAllowedPaymentMethodType = "przelewy24"
	SubscriptionNewParamsAllowedPaymentMethodTypeAffirm           SubscriptionNewParamsAllowedPaymentMethodType = "affirm"
	SubscriptionNewParamsAllowedPaymentMethodTypeKlarna           SubscriptionNewParamsAllowedPaymentMethodType = "klarna"
	SubscriptionNewParamsAllowedPaymentMethodTypeSepa             SubscriptionNewParamsAllowedPaymentMethodType = "sepa"
	SubscriptionNewParamsAllowedPaymentMethodTypeACH              SubscriptionNewParamsAllowedPaymentMethodType = "ach"
	SubscriptionNewParamsAllowedPaymentMethodTypeAmazonPay        SubscriptionNewParamsAllowedPaymentMethodType = "amazon_pay"
	SubscriptionNewParamsAllowedPaymentMethodTypeAfterpayClearpay SubscriptionNewParamsAllowedPaymentMethodType = "afterpay_clearpay"
)

func (r SubscriptionNewParamsAllowedPaymentMethodType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAllowedPaymentMethodTypeCredit, SubscriptionNewParamsAllowedPaymentMethodTypeDebit, SubscriptionNewParamsAllowedPaymentMethodTypeUpiCollect, SubscriptionNewParamsAllowedPaymentMethodTypeUpiIntent, SubscriptionNewParamsAllowedPaymentMethodTypeApplePay, SubscriptionNewParamsAllowedPaymentMethodTypeCashapp, SubscriptionNewParamsAllowedPaymentMethodTypeGooglePay, SubscriptionNewParamsAllowedPaymentMethodTypeMultibanco, SubscriptionNewParamsAllowedPaymentMethodTypeBancontactCard, SubscriptionNewParamsAllowedPaymentMethodTypeEps, SubscriptionNewParamsAllowedPaymentMethodTypeIdeal, SubscriptionNewParamsAllowedPaymentMethodTypePrzelewy24, SubscriptionNewParamsAllowedPaymentMethodTypeAffirm, SubscriptionNewParamsAllowedPaymentMethodTypeKlarna, SubscriptionNewParamsAllowedPaymentMethodTypeSepa, SubscriptionNewParamsAllowedPaymentMethodTypeACH, SubscriptionNewParamsAllowedPaymentMethodTypeAmazonPay, SubscriptionNewParamsAllowedPaymentMethodTypeAfterpayClearpay:
		return true
	}
	return false
}

type SubscriptionNewParamsBillingCurrency string

const (
	SubscriptionNewParamsBillingCurrencyAed SubscriptionNewParamsBillingCurrency = "AED"
	SubscriptionNewParamsBillingCurrencyAll SubscriptionNewParamsBillingCurrency = "ALL"
	SubscriptionNewParamsBillingCurrencyAmd SubscriptionNewParamsBillingCurrency = "AMD"
	SubscriptionNewParamsBillingCurrencyAng SubscriptionNewParamsBillingCurrency = "ANG"
	SubscriptionNewParamsBillingCurrencyAoa SubscriptionNewParamsBillingCurrency = "AOA"
	SubscriptionNewParamsBillingCurrencyArs SubscriptionNewParamsBillingCurrency = "ARS"
	SubscriptionNewParamsBillingCurrencyAud SubscriptionNewParamsBillingCurrency = "AUD"
	SubscriptionNewParamsBillingCurrencyAwg SubscriptionNewParamsBillingCurrency = "AWG"
	SubscriptionNewParamsBillingCurrencyAzn SubscriptionNewParamsBillingCurrency = "AZN"
	SubscriptionNewParamsBillingCurrencyBam SubscriptionNewParamsBillingCurrency = "BAM"
	SubscriptionNewParamsBillingCurrencyBbd SubscriptionNewParamsBillingCurrency = "BBD"
	SubscriptionNewParamsBillingCurrencyBdt SubscriptionNewParamsBillingCurrency = "BDT"
	SubscriptionNewParamsBillingCurrencyBgn SubscriptionNewParamsBillingCurrency = "BGN"
	SubscriptionNewParamsBillingCurrencyBhd SubscriptionNewParamsBillingCurrency = "BHD"
	SubscriptionNewParamsBillingCurrencyBif SubscriptionNewParamsBillingCurrency = "BIF"
	SubscriptionNewParamsBillingCurrencyBmd SubscriptionNewParamsBillingCurrency = "BMD"
	SubscriptionNewParamsBillingCurrencyBnd SubscriptionNewParamsBillingCurrency = "BND"
	SubscriptionNewParamsBillingCurrencyBob SubscriptionNewParamsBillingCurrency = "BOB"
	SubscriptionNewParamsBillingCurrencyBrl SubscriptionNewParamsBillingCurrency = "BRL"
	SubscriptionNewParamsBillingCurrencyBsd SubscriptionNewParamsBillingCurrency = "BSD"
	SubscriptionNewParamsBillingCurrencyBwp SubscriptionNewParamsBillingCurrency = "BWP"
	SubscriptionNewParamsBillingCurrencyByn SubscriptionNewParamsBillingCurrency = "BYN"
	SubscriptionNewParamsBillingCurrencyBzd SubscriptionNewParamsBillingCurrency = "BZD"
	SubscriptionNewParamsBillingCurrencyCad SubscriptionNewParamsBillingCurrency = "CAD"
	SubscriptionNewParamsBillingCurrencyChf SubscriptionNewParamsBillingCurrency = "CHF"
	SubscriptionNewParamsBillingCurrencyClp SubscriptionNewParamsBillingCurrency = "CLP"
	SubscriptionNewParamsBillingCurrencyCny SubscriptionNewParamsBillingCurrency = "CNY"
	SubscriptionNewParamsBillingCurrencyCop SubscriptionNewParamsBillingCurrency = "COP"
	SubscriptionNewParamsBillingCurrencyCrc SubscriptionNewParamsBillingCurrency = "CRC"
	SubscriptionNewParamsBillingCurrencyCup SubscriptionNewParamsBillingCurrency = "CUP"
	SubscriptionNewParamsBillingCurrencyCve SubscriptionNewParamsBillingCurrency = "CVE"
	SubscriptionNewParamsBillingCurrencyCzk SubscriptionNewParamsBillingCurrency = "CZK"
	SubscriptionNewParamsBillingCurrencyDjf SubscriptionNewParamsBillingCurrency = "DJF"
	SubscriptionNewParamsBillingCurrencyDkk SubscriptionNewParamsBillingCurrency = "DKK"
	SubscriptionNewParamsBillingCurrencyDop SubscriptionNewParamsBillingCurrency = "DOP"
	SubscriptionNewParamsBillingCurrencyDzd SubscriptionNewParamsBillingCurrency = "DZD"
	SubscriptionNewParamsBillingCurrencyEgp SubscriptionNewParamsBillingCurrency = "EGP"
	SubscriptionNewParamsBillingCurrencyEtb SubscriptionNewParamsBillingCurrency = "ETB"
	SubscriptionNewParamsBillingCurrencyEur SubscriptionNewParamsBillingCurrency = "EUR"
	SubscriptionNewParamsBillingCurrencyFjd SubscriptionNewParamsBillingCurrency = "FJD"
	SubscriptionNewParamsBillingCurrencyFkp SubscriptionNewParamsBillingCurrency = "FKP"
	SubscriptionNewParamsBillingCurrencyGbp SubscriptionNewParamsBillingCurrency = "GBP"
	SubscriptionNewParamsBillingCurrencyGel SubscriptionNewParamsBillingCurrency = "GEL"
	SubscriptionNewParamsBillingCurrencyGhs SubscriptionNewParamsBillingCurrency = "GHS"
	SubscriptionNewParamsBillingCurrencyGip SubscriptionNewParamsBillingCurrency = "GIP"
	SubscriptionNewParamsBillingCurrencyGmd SubscriptionNewParamsBillingCurrency = "GMD"
	SubscriptionNewParamsBillingCurrencyGnf SubscriptionNewParamsBillingCurrency = "GNF"
	SubscriptionNewParamsBillingCurrencyGtq SubscriptionNewParamsBillingCurrency = "GTQ"
	SubscriptionNewParamsBillingCurrencyGyd SubscriptionNewParamsBillingCurrency = "GYD"
	SubscriptionNewParamsBillingCurrencyHkd SubscriptionNewParamsBillingCurrency = "HKD"
	SubscriptionNewParamsBillingCurrencyHnl SubscriptionNewParamsBillingCurrency = "HNL"
	SubscriptionNewParamsBillingCurrencyHrk SubscriptionNewParamsBillingCurrency = "HRK"
	SubscriptionNewParamsBillingCurrencyHtg SubscriptionNewParamsBillingCurrency = "HTG"
	SubscriptionNewParamsBillingCurrencyHuf SubscriptionNewParamsBillingCurrency = "HUF"
	SubscriptionNewParamsBillingCurrencyIdr SubscriptionNewParamsBillingCurrency = "IDR"
	SubscriptionNewParamsBillingCurrencyIls SubscriptionNewParamsBillingCurrency = "ILS"
	SubscriptionNewParamsBillingCurrencyInr SubscriptionNewParamsBillingCurrency = "INR"
	SubscriptionNewParamsBillingCurrencyIqd SubscriptionNewParamsBillingCurrency = "IQD"
	SubscriptionNewParamsBillingCurrencyJmd SubscriptionNewParamsBillingCurrency = "JMD"
	SubscriptionNewParamsBillingCurrencyJod SubscriptionNewParamsBillingCurrency = "JOD"
	SubscriptionNewParamsBillingCurrencyJpy SubscriptionNewParamsBillingCurrency = "JPY"
	SubscriptionNewParamsBillingCurrencyKes SubscriptionNewParamsBillingCurrency = "KES"
	SubscriptionNewParamsBillingCurrencyKgs SubscriptionNewParamsBillingCurrency = "KGS"
	SubscriptionNewParamsBillingCurrencyKhr SubscriptionNewParamsBillingCurrency = "KHR"
	SubscriptionNewParamsBillingCurrencyKmf SubscriptionNewParamsBillingCurrency = "KMF"
	SubscriptionNewParamsBillingCurrencyKrw SubscriptionNewParamsBillingCurrency = "KRW"
	SubscriptionNewParamsBillingCurrencyKwd SubscriptionNewParamsBillingCurrency = "KWD"
	SubscriptionNewParamsBillingCurrencyKyd SubscriptionNewParamsBillingCurrency = "KYD"
	SubscriptionNewParamsBillingCurrencyKzt SubscriptionNewParamsBillingCurrency = "KZT"
	SubscriptionNewParamsBillingCurrencyLak SubscriptionNewParamsBillingCurrency = "LAK"
	SubscriptionNewParamsBillingCurrencyLbp SubscriptionNewParamsBillingCurrency = "LBP"
	SubscriptionNewParamsBillingCurrencyLkr SubscriptionNewParamsBillingCurrency = "LKR"
	SubscriptionNewParamsBillingCurrencyLrd SubscriptionNewParamsBillingCurrency = "LRD"
	SubscriptionNewParamsBillingCurrencyLsl SubscriptionNewParamsBillingCurrency = "LSL"
	SubscriptionNewParamsBillingCurrencyLyd SubscriptionNewParamsBillingCurrency = "LYD"
	SubscriptionNewParamsBillingCurrencyMad SubscriptionNewParamsBillingCurrency = "MAD"
	SubscriptionNewParamsBillingCurrencyMdl SubscriptionNewParamsBillingCurrency = "MDL"
	SubscriptionNewParamsBillingCurrencyMga SubscriptionNewParamsBillingCurrency = "MGA"
	SubscriptionNewParamsBillingCurrencyMkd SubscriptionNewParamsBillingCurrency = "MKD"
	SubscriptionNewParamsBillingCurrencyMmk SubscriptionNewParamsBillingCurrency = "MMK"
	SubscriptionNewParamsBillingCurrencyMnt SubscriptionNewParamsBillingCurrency = "MNT"
	SubscriptionNewParamsBillingCurrencyMop SubscriptionNewParamsBillingCurrency = "MOP"
	SubscriptionNewParamsBillingCurrencyMru SubscriptionNewParamsBillingCurrency = "MRU"
	SubscriptionNewParamsBillingCurrencyMur SubscriptionNewParamsBillingCurrency = "MUR"
	SubscriptionNewParamsBillingCurrencyMvr SubscriptionNewParamsBillingCurrency = "MVR"
	SubscriptionNewParamsBillingCurrencyMwk SubscriptionNewParamsBillingCurrency = "MWK"
	SubscriptionNewParamsBillingCurrencyMxn SubscriptionNewParamsBillingCurrency = "MXN"
	SubscriptionNewParamsBillingCurrencyMyr SubscriptionNewParamsBillingCurrency = "MYR"
	SubscriptionNewParamsBillingCurrencyMzn SubscriptionNewParamsBillingCurrency = "MZN"
	SubscriptionNewParamsBillingCurrencyNad SubscriptionNewParamsBillingCurrency = "NAD"
	SubscriptionNewParamsBillingCurrencyNgn SubscriptionNewParamsBillingCurrency = "NGN"
	SubscriptionNewParamsBillingCurrencyNio SubscriptionNewParamsBillingCurrency = "NIO"
	SubscriptionNewParamsBillingCurrencyNok SubscriptionNewParamsBillingCurrency = "NOK"
	SubscriptionNewParamsBillingCurrencyNpr SubscriptionNewParamsBillingCurrency = "NPR"
	SubscriptionNewParamsBillingCurrencyNzd SubscriptionNewParamsBillingCurrency = "NZD"
	SubscriptionNewParamsBillingCurrencyOmr SubscriptionNewParamsBillingCurrency = "OMR"
	SubscriptionNewParamsBillingCurrencyPab SubscriptionNewParamsBillingCurrency = "PAB"
	SubscriptionNewParamsBillingCurrencyPen SubscriptionNewParamsBillingCurrency = "PEN"
	SubscriptionNewParamsBillingCurrencyPgk SubscriptionNewParamsBillingCurrency = "PGK"
	SubscriptionNewParamsBillingCurrencyPhp SubscriptionNewParamsBillingCurrency = "PHP"
	SubscriptionNewParamsBillingCurrencyPkr SubscriptionNewParamsBillingCurrency = "PKR"
	SubscriptionNewParamsBillingCurrencyPln SubscriptionNewParamsBillingCurrency = "PLN"
	SubscriptionNewParamsBillingCurrencyPyg SubscriptionNewParamsBillingCurrency = "PYG"
	SubscriptionNewParamsBillingCurrencyQar SubscriptionNewParamsBillingCurrency = "QAR"
	SubscriptionNewParamsBillingCurrencyRon SubscriptionNewParamsBillingCurrency = "RON"
	SubscriptionNewParamsBillingCurrencyRsd SubscriptionNewParamsBillingCurrency = "RSD"
	SubscriptionNewParamsBillingCurrencyRub SubscriptionNewParamsBillingCurrency = "RUB"
	SubscriptionNewParamsBillingCurrencyRwf SubscriptionNewParamsBillingCurrency = "RWF"
	SubscriptionNewParamsBillingCurrencySar SubscriptionNewParamsBillingCurrency = "SAR"
	SubscriptionNewParamsBillingCurrencySbd SubscriptionNewParamsBillingCurrency = "SBD"
	SubscriptionNewParamsBillingCurrencyScr SubscriptionNewParamsBillingCurrency = "SCR"
	SubscriptionNewParamsBillingCurrencySek SubscriptionNewParamsBillingCurrency = "SEK"
	SubscriptionNewParamsBillingCurrencySgd SubscriptionNewParamsBillingCurrency = "SGD"
	SubscriptionNewParamsBillingCurrencyShp SubscriptionNewParamsBillingCurrency = "SHP"
	SubscriptionNewParamsBillingCurrencySle SubscriptionNewParamsBillingCurrency = "SLE"
	SubscriptionNewParamsBillingCurrencySll SubscriptionNewParamsBillingCurrency = "SLL"
	SubscriptionNewParamsBillingCurrencySos SubscriptionNewParamsBillingCurrency = "SOS"
	SubscriptionNewParamsBillingCurrencySrd SubscriptionNewParamsBillingCurrency = "SRD"
	SubscriptionNewParamsBillingCurrencySsp SubscriptionNewParamsBillingCurrency = "SSP"
	SubscriptionNewParamsBillingCurrencyStn SubscriptionNewParamsBillingCurrency = "STN"
	SubscriptionNewParamsBillingCurrencySvc SubscriptionNewParamsBillingCurrency = "SVC"
	SubscriptionNewParamsBillingCurrencySzl SubscriptionNewParamsBillingCurrency = "SZL"
	SubscriptionNewParamsBillingCurrencyThb SubscriptionNewParamsBillingCurrency = "THB"
	SubscriptionNewParamsBillingCurrencyTnd SubscriptionNewParamsBillingCurrency = "TND"
	SubscriptionNewParamsBillingCurrencyTop SubscriptionNewParamsBillingCurrency = "TOP"
	SubscriptionNewParamsBillingCurrencyTry SubscriptionNewParamsBillingCurrency = "TRY"
	SubscriptionNewParamsBillingCurrencyTtd SubscriptionNewParamsBillingCurrency = "TTD"
	SubscriptionNewParamsBillingCurrencyTwd SubscriptionNewParamsBillingCurrency = "TWD"
	SubscriptionNewParamsBillingCurrencyTzs SubscriptionNewParamsBillingCurrency = "TZS"
	SubscriptionNewParamsBillingCurrencyUah SubscriptionNewParamsBillingCurrency = "UAH"
	SubscriptionNewParamsBillingCurrencyUgx SubscriptionNewParamsBillingCurrency = "UGX"
	SubscriptionNewParamsBillingCurrencyUsd SubscriptionNewParamsBillingCurrency = "USD"
	SubscriptionNewParamsBillingCurrencyUyu SubscriptionNewParamsBillingCurrency = "UYU"
	SubscriptionNewParamsBillingCurrencyUzs SubscriptionNewParamsBillingCurrency = "UZS"
	SubscriptionNewParamsBillingCurrencyVes SubscriptionNewParamsBillingCurrency = "VES"
	SubscriptionNewParamsBillingCurrencyVnd SubscriptionNewParamsBillingCurrency = "VND"
	SubscriptionNewParamsBillingCurrencyVuv SubscriptionNewParamsBillingCurrency = "VUV"
	SubscriptionNewParamsBillingCurrencyWst SubscriptionNewParamsBillingCurrency = "WST"
	SubscriptionNewParamsBillingCurrencyXaf SubscriptionNewParamsBillingCurrency = "XAF"
	SubscriptionNewParamsBillingCurrencyXcd SubscriptionNewParamsBillingCurrency = "XCD"
	SubscriptionNewParamsBillingCurrencyXof SubscriptionNewParamsBillingCurrency = "XOF"
	SubscriptionNewParamsBillingCurrencyXpf SubscriptionNewParamsBillingCurrency = "XPF"
	SubscriptionNewParamsBillingCurrencyYer SubscriptionNewParamsBillingCurrency = "YER"
	SubscriptionNewParamsBillingCurrencyZar SubscriptionNewParamsBillingCurrency = "ZAR"
	SubscriptionNewParamsBillingCurrencyZmw SubscriptionNewParamsBillingCurrency = "ZMW"
)

func (r SubscriptionNewParamsBillingCurrency) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsBillingCurrencyAed, SubscriptionNewParamsBillingCurrencyAll, SubscriptionNewParamsBillingCurrencyAmd, SubscriptionNewParamsBillingCurrencyAng, SubscriptionNewParamsBillingCurrencyAoa, SubscriptionNewParamsBillingCurrencyArs, SubscriptionNewParamsBillingCurrencyAud, SubscriptionNewParamsBillingCurrencyAwg, SubscriptionNewParamsBillingCurrencyAzn, SubscriptionNewParamsBillingCurrencyBam, SubscriptionNewParamsBillingCurrencyBbd, SubscriptionNewParamsBillingCurrencyBdt, SubscriptionNewParamsBillingCurrencyBgn, SubscriptionNewParamsBillingCurrencyBhd, SubscriptionNewParamsBillingCurrencyBif, SubscriptionNewParamsBillingCurrencyBmd, SubscriptionNewParamsBillingCurrencyBnd, SubscriptionNewParamsBillingCurrencyBob, SubscriptionNewParamsBillingCurrencyBrl, SubscriptionNewParamsBillingCurrencyBsd, SubscriptionNewParamsBillingCurrencyBwp, SubscriptionNewParamsBillingCurrencyByn, SubscriptionNewParamsBillingCurrencyBzd, SubscriptionNewParamsBillingCurrencyCad, SubscriptionNewParamsBillingCurrencyChf, SubscriptionNewParamsBillingCurrencyClp, SubscriptionNewParamsBillingCurrencyCny, SubscriptionNewParamsBillingCurrencyCop, SubscriptionNewParamsBillingCurrencyCrc, SubscriptionNewParamsBillingCurrencyCup, SubscriptionNewParamsBillingCurrencyCve, SubscriptionNewParamsBillingCurrencyCzk, SubscriptionNewParamsBillingCurrencyDjf, SubscriptionNewParamsBillingCurrencyDkk, SubscriptionNewParamsBillingCurrencyDop, SubscriptionNewParamsBillingCurrencyDzd, SubscriptionNewParamsBillingCurrencyEgp, SubscriptionNewParamsBillingCurrencyEtb, SubscriptionNewParamsBillingCurrencyEur, SubscriptionNewParamsBillingCurrencyFjd, SubscriptionNewParamsBillingCurrencyFkp, SubscriptionNewParamsBillingCurrencyGbp, SubscriptionNewParamsBillingCurrencyGel, SubscriptionNewParamsBillingCurrencyGhs, SubscriptionNewParamsBillingCurrencyGip, SubscriptionNewParamsBillingCurrencyGmd, SubscriptionNewParamsBillingCurrencyGnf, SubscriptionNewParamsBillingCurrencyGtq, SubscriptionNewParamsBillingCurrencyGyd, SubscriptionNewParamsBillingCurrencyHkd, SubscriptionNewParamsBillingCurrencyHnl, SubscriptionNewParamsBillingCurrencyHrk, SubscriptionNewParamsBillingCurrencyHtg, SubscriptionNewParamsBillingCurrencyHuf, SubscriptionNewParamsBillingCurrencyIdr, SubscriptionNewParamsBillingCurrencyIls, SubscriptionNewParamsBillingCurrencyInr, SubscriptionNewParamsBillingCurrencyIqd, SubscriptionNewParamsBillingCurrencyJmd, SubscriptionNewParamsBillingCurrencyJod, SubscriptionNewParamsBillingCurrencyJpy, SubscriptionNewParamsBillingCurrencyKes, SubscriptionNewParamsBillingCurrencyKgs, SubscriptionNewParamsBillingCurrencyKhr, SubscriptionNewParamsBillingCurrencyKmf, SubscriptionNewParamsBillingCurrencyKrw, SubscriptionNewParamsBillingCurrencyKwd, SubscriptionNewParamsBillingCurrencyKyd, SubscriptionNewParamsBillingCurrencyKzt, SubscriptionNewParamsBillingCurrencyLak, SubscriptionNewParamsBillingCurrencyLbp, SubscriptionNewParamsBillingCurrencyLkr, SubscriptionNewParamsBillingCurrencyLrd, SubscriptionNewParamsBillingCurrencyLsl, SubscriptionNewParamsBillingCurrencyLyd, SubscriptionNewParamsBillingCurrencyMad, SubscriptionNewParamsBillingCurrencyMdl, SubscriptionNewParamsBillingCurrencyMga, SubscriptionNewParamsBillingCurrencyMkd, SubscriptionNewParamsBillingCurrencyMmk, SubscriptionNewParamsBillingCurrencyMnt, SubscriptionNewParamsBillingCurrencyMop, SubscriptionNewParamsBillingCurrencyMru, SubscriptionNewParamsBillingCurrencyMur, SubscriptionNewParamsBillingCurrencyMvr, SubscriptionNewParamsBillingCurrencyMwk, SubscriptionNewParamsBillingCurrencyMxn, SubscriptionNewParamsBillingCurrencyMyr, SubscriptionNewParamsBillingCurrencyMzn, SubscriptionNewParamsBillingCurrencyNad, SubscriptionNewParamsBillingCurrencyNgn, SubscriptionNewParamsBillingCurrencyNio, SubscriptionNewParamsBillingCurrencyNok, SubscriptionNewParamsBillingCurrencyNpr, SubscriptionNewParamsBillingCurrencyNzd, SubscriptionNewParamsBillingCurrencyOmr, SubscriptionNewParamsBillingCurrencyPab, SubscriptionNewParamsBillingCurrencyPen, SubscriptionNewParamsBillingCurrencyPgk, SubscriptionNewParamsBillingCurrencyPhp, SubscriptionNewParamsBillingCurrencyPkr, SubscriptionNewParamsBillingCurrencyPln, SubscriptionNewParamsBillingCurrencyPyg, SubscriptionNewParamsBillingCurrencyQar, SubscriptionNewParamsBillingCurrencyRon, SubscriptionNewParamsBillingCurrencyRsd, SubscriptionNewParamsBillingCurrencyRub, SubscriptionNewParamsBillingCurrencyRwf, SubscriptionNewParamsBillingCurrencySar, SubscriptionNewParamsBillingCurrencySbd, SubscriptionNewParamsBillingCurrencyScr, SubscriptionNewParamsBillingCurrencySek, SubscriptionNewParamsBillingCurrencySgd, SubscriptionNewParamsBillingCurrencyShp, SubscriptionNewParamsBillingCurrencySle, SubscriptionNewParamsBillingCurrencySll, SubscriptionNewParamsBillingCurrencySos, SubscriptionNewParamsBillingCurrencySrd, SubscriptionNewParamsBillingCurrencySsp, SubscriptionNewParamsBillingCurrencyStn, SubscriptionNewParamsBillingCurrencySvc, SubscriptionNewParamsBillingCurrencySzl, SubscriptionNewParamsBillingCurrencyThb, SubscriptionNewParamsBillingCurrencyTnd, SubscriptionNewParamsBillingCurrencyTop, SubscriptionNewParamsBillingCurrencyTry, SubscriptionNewParamsBillingCurrencyTtd, SubscriptionNewParamsBillingCurrencyTwd, SubscriptionNewParamsBillingCurrencyTzs, SubscriptionNewParamsBillingCurrencyUah, SubscriptionNewParamsBillingCurrencyUgx, SubscriptionNewParamsBillingCurrencyUsd, SubscriptionNewParamsBillingCurrencyUyu, SubscriptionNewParamsBillingCurrencyUzs, SubscriptionNewParamsBillingCurrencyVes, SubscriptionNewParamsBillingCurrencyVnd, SubscriptionNewParamsBillingCurrencyVuv, SubscriptionNewParamsBillingCurrencyWst, SubscriptionNewParamsBillingCurrencyXaf, SubscriptionNewParamsBillingCurrencyXcd, SubscriptionNewParamsBillingCurrencyXof, SubscriptionNewParamsBillingCurrencyXpf, SubscriptionNewParamsBillingCurrencyYer, SubscriptionNewParamsBillingCurrencyZar, SubscriptionNewParamsBillingCurrencyZmw:
		return true
	}
	return false
}

type SubscriptionNewParamsOnDemand struct {
	// If set as True, does not perform any charge and only authorizes payment method
	// details for future use.
	MandateOnly param.Field[bool] `json:"mandate_only,required"`
	// Product price for the initial charge to customer If not specified the stored
	// price of the product will be used Represented in the lowest denomination of the
	// currency (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	ProductPrice param.Field[int64] `json:"product_price"`
}

func (r SubscriptionNewParamsOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParams struct {
	Billing         param.Field[BillingAddressParam]                     `json:"billing"`
	DisableOnDemand param.Field[SubscriptionUpdateParamsDisableOnDemand] `json:"disable_on_demand"`
	Metadata        param.Field[map[string]string]                       `json:"metadata"`
	Status          param.Field[SubscriptionStatus]                      `json:"status"`
	TaxID           param.Field[string]                                  `json:"tax_id"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateParamsDisableOnDemand struct {
	NextBillingDate param.Field[time.Time] `json:"next_billing_date,required" format:"date-time"`
}

func (r SubscriptionUpdateParamsDisableOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionListParams struct {
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
	// Filter by status
	Status param.Field[SubscriptionStatus] `query:"status"`
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubscriptionChangePlanParams struct {
	// Unique identifier of the product to subscribe to
	ProductID            param.Field[string]                                           `json:"product_id,required"`
	ProrationBillingMode param.Field[SubscriptionChangePlanParamsProrationBillingMode] `json:"proration_billing_mode,required"`
	// Number of units to subscribe for. Must be at least 1.
	Quantity param.Field[int64] `json:"quantity,required"`
}

func (r SubscriptionChangePlanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionChangePlanParamsProrationBillingMode string

const (
	SubscriptionChangePlanParamsProrationBillingModeProratedImmediately SubscriptionChangePlanParamsProrationBillingMode = "prorated_immediately"
)

func (r SubscriptionChangePlanParamsProrationBillingMode) IsKnown() bool {
	switch r {
	case SubscriptionChangePlanParamsProrationBillingModeProratedImmediately:
		return true
	}
	return false
}

type SubscriptionChargeParams struct {
	// The product price. Represented in the lowest denomination of the currency (e.g.,
	// cents for USD). For example, to charge $1.00, pass `100`.
	ProductPrice param.Field[int64] `json:"product_price,required"`
}

func (r SubscriptionChargeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
