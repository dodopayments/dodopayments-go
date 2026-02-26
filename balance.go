// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
)

// BalanceService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r *BalanceService) {
	r = &BalanceService{}
	r.Options = opts
	return
}

func (r *BalanceService) GetLedger(ctx context.Context, query BalanceGetLedgerParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[BalanceLedgerEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "balances/ledger"
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

func (r *BalanceService) GetLedgerAutoPaging(ctx context.Context, query BalanceGetLedgerParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[BalanceLedgerEntry] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.GetLedger(ctx, query, opts...))
}

type BalanceLedgerEntry struct {
	ID                  string                      `json:"id" api:"required"`
	Amount              int64                       `json:"amount" api:"required"`
	BusinessID          string                      `json:"business_id" api:"required"`
	CreatedAt           time.Time                   `json:"created_at" api:"required" format:"date-time"`
	Currency            Currency                    `json:"currency" api:"required"`
	EventType           BalanceLedgerEntryEventType `json:"event_type" api:"required"`
	IsCredit            bool                        `json:"is_credit" api:"required"`
	UsdEquivalentAmount int64                       `json:"usd_equivalent_amount" api:"required"`
	AfterBalance        int64                       `json:"after_balance" api:"nullable"`
	BeforeBalance       int64                       `json:"before_balance" api:"nullable"`
	Description         string                      `json:"description" api:"nullable"`
	ReferenceObjectID   string                      `json:"reference_object_id" api:"nullable"`
	JSON                balanceLedgerEntryJSON      `json:"-"`
}

// balanceLedgerEntryJSON contains the JSON metadata for the struct
// [BalanceLedgerEntry]
type balanceLedgerEntryJSON struct {
	ID                  apijson.Field
	Amount              apijson.Field
	BusinessID          apijson.Field
	CreatedAt           apijson.Field
	Currency            apijson.Field
	EventType           apijson.Field
	IsCredit            apijson.Field
	UsdEquivalentAmount apijson.Field
	AfterBalance        apijson.Field
	BeforeBalance       apijson.Field
	Description         apijson.Field
	ReferenceObjectID   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceLedgerEntryJSON) RawJSON() string {
	return r.raw
}

type BalanceLedgerEntryEventType string

const (
	BalanceLedgerEntryEventTypePayment            BalanceLedgerEntryEventType = "payment"
	BalanceLedgerEntryEventTypeRefund             BalanceLedgerEntryEventType = "refund"
	BalanceLedgerEntryEventTypeRefundReversal     BalanceLedgerEntryEventType = "refund_reversal"
	BalanceLedgerEntryEventTypeDispute            BalanceLedgerEntryEventType = "dispute"
	BalanceLedgerEntryEventTypeDisputeReversal    BalanceLedgerEntryEventType = "dispute_reversal"
	BalanceLedgerEntryEventTypeTax                BalanceLedgerEntryEventType = "tax"
	BalanceLedgerEntryEventTypeTaxReversal        BalanceLedgerEntryEventType = "tax_reversal"
	BalanceLedgerEntryEventTypePaymentFees        BalanceLedgerEntryEventType = "payment_fees"
	BalanceLedgerEntryEventTypeRefundFees         BalanceLedgerEntryEventType = "refund_fees"
	BalanceLedgerEntryEventTypeRefundFeesReversal BalanceLedgerEntryEventType = "refund_fees_reversal"
	BalanceLedgerEntryEventTypeDisputeFees        BalanceLedgerEntryEventType = "dispute_fees"
	BalanceLedgerEntryEventTypePayout             BalanceLedgerEntryEventType = "payout"
	BalanceLedgerEntryEventTypePayoutFees         BalanceLedgerEntryEventType = "payout_fees"
	BalanceLedgerEntryEventTypePayoutReversal     BalanceLedgerEntryEventType = "payout_reversal"
	BalanceLedgerEntryEventTypePayoutFeesReversal BalanceLedgerEntryEventType = "payout_fees_reversal"
	BalanceLedgerEntryEventTypeDodoCredits        BalanceLedgerEntryEventType = "dodo_credits"
	BalanceLedgerEntryEventTypeAdjustment         BalanceLedgerEntryEventType = "adjustment"
	BalanceLedgerEntryEventTypeCurrencyConversion BalanceLedgerEntryEventType = "currency_conversion"
)

func (r BalanceLedgerEntryEventType) IsKnown() bool {
	switch r {
	case BalanceLedgerEntryEventTypePayment, BalanceLedgerEntryEventTypeRefund, BalanceLedgerEntryEventTypeRefundReversal, BalanceLedgerEntryEventTypeDispute, BalanceLedgerEntryEventTypeDisputeReversal, BalanceLedgerEntryEventTypeTax, BalanceLedgerEntryEventTypeTaxReversal, BalanceLedgerEntryEventTypePaymentFees, BalanceLedgerEntryEventTypeRefundFees, BalanceLedgerEntryEventTypeRefundFeesReversal, BalanceLedgerEntryEventTypeDisputeFees, BalanceLedgerEntryEventTypePayout, BalanceLedgerEntryEventTypePayoutFees, BalanceLedgerEntryEventTypePayoutReversal, BalanceLedgerEntryEventTypePayoutFeesReversal, BalanceLedgerEntryEventTypeDodoCredits, BalanceLedgerEntryEventTypeAdjustment, BalanceLedgerEntryEventTypeCurrencyConversion:
		return true
	}
	return false
}

type BalanceGetLedgerParams struct {
	// Get events after this created time
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get events created before this time
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Filter by currency
	Currency param.Field[BalanceGetLedgerParamsCurrency] `query:"currency"`
	// Filter by Ledger Event Type
	EventType param.Field[BalanceGetLedgerParamsEventType] `query:"event_type"`
	// Min : 1, Max : 100, default 10
	Limit param.Field[int64] `query:"limit"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Get events history of a specific object like payment/subscription/refund/dispute
	ReferenceObjectID param.Field[string] `query:"reference_object_id"`
}

// URLQuery serializes [BalanceGetLedgerParams]'s query parameters as `url.Values`.
func (r BalanceGetLedgerParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by currency
type BalanceGetLedgerParamsCurrency string

const (
	BalanceGetLedgerParamsCurrencyAed BalanceGetLedgerParamsCurrency = "AED"
	BalanceGetLedgerParamsCurrencyAll BalanceGetLedgerParamsCurrency = "ALL"
	BalanceGetLedgerParamsCurrencyAmd BalanceGetLedgerParamsCurrency = "AMD"
	BalanceGetLedgerParamsCurrencyAng BalanceGetLedgerParamsCurrency = "ANG"
	BalanceGetLedgerParamsCurrencyAoa BalanceGetLedgerParamsCurrency = "AOA"
	BalanceGetLedgerParamsCurrencyArs BalanceGetLedgerParamsCurrency = "ARS"
	BalanceGetLedgerParamsCurrencyAud BalanceGetLedgerParamsCurrency = "AUD"
	BalanceGetLedgerParamsCurrencyAwg BalanceGetLedgerParamsCurrency = "AWG"
	BalanceGetLedgerParamsCurrencyAzn BalanceGetLedgerParamsCurrency = "AZN"
	BalanceGetLedgerParamsCurrencyBam BalanceGetLedgerParamsCurrency = "BAM"
	BalanceGetLedgerParamsCurrencyBbd BalanceGetLedgerParamsCurrency = "BBD"
	BalanceGetLedgerParamsCurrencyBdt BalanceGetLedgerParamsCurrency = "BDT"
	BalanceGetLedgerParamsCurrencyBgn BalanceGetLedgerParamsCurrency = "BGN"
	BalanceGetLedgerParamsCurrencyBhd BalanceGetLedgerParamsCurrency = "BHD"
	BalanceGetLedgerParamsCurrencyBif BalanceGetLedgerParamsCurrency = "BIF"
	BalanceGetLedgerParamsCurrencyBmd BalanceGetLedgerParamsCurrency = "BMD"
	BalanceGetLedgerParamsCurrencyBnd BalanceGetLedgerParamsCurrency = "BND"
	BalanceGetLedgerParamsCurrencyBob BalanceGetLedgerParamsCurrency = "BOB"
	BalanceGetLedgerParamsCurrencyBrl BalanceGetLedgerParamsCurrency = "BRL"
	BalanceGetLedgerParamsCurrencyBsd BalanceGetLedgerParamsCurrency = "BSD"
	BalanceGetLedgerParamsCurrencyBwp BalanceGetLedgerParamsCurrency = "BWP"
	BalanceGetLedgerParamsCurrencyByn BalanceGetLedgerParamsCurrency = "BYN"
	BalanceGetLedgerParamsCurrencyBzd BalanceGetLedgerParamsCurrency = "BZD"
	BalanceGetLedgerParamsCurrencyCad BalanceGetLedgerParamsCurrency = "CAD"
	BalanceGetLedgerParamsCurrencyChf BalanceGetLedgerParamsCurrency = "CHF"
	BalanceGetLedgerParamsCurrencyClp BalanceGetLedgerParamsCurrency = "CLP"
	BalanceGetLedgerParamsCurrencyCny BalanceGetLedgerParamsCurrency = "CNY"
	BalanceGetLedgerParamsCurrencyCop BalanceGetLedgerParamsCurrency = "COP"
	BalanceGetLedgerParamsCurrencyCrc BalanceGetLedgerParamsCurrency = "CRC"
	BalanceGetLedgerParamsCurrencyCup BalanceGetLedgerParamsCurrency = "CUP"
	BalanceGetLedgerParamsCurrencyCve BalanceGetLedgerParamsCurrency = "CVE"
	BalanceGetLedgerParamsCurrencyCzk BalanceGetLedgerParamsCurrency = "CZK"
	BalanceGetLedgerParamsCurrencyDjf BalanceGetLedgerParamsCurrency = "DJF"
	BalanceGetLedgerParamsCurrencyDkk BalanceGetLedgerParamsCurrency = "DKK"
	BalanceGetLedgerParamsCurrencyDop BalanceGetLedgerParamsCurrency = "DOP"
	BalanceGetLedgerParamsCurrencyDzd BalanceGetLedgerParamsCurrency = "DZD"
	BalanceGetLedgerParamsCurrencyEgp BalanceGetLedgerParamsCurrency = "EGP"
	BalanceGetLedgerParamsCurrencyEtb BalanceGetLedgerParamsCurrency = "ETB"
	BalanceGetLedgerParamsCurrencyEur BalanceGetLedgerParamsCurrency = "EUR"
	BalanceGetLedgerParamsCurrencyFjd BalanceGetLedgerParamsCurrency = "FJD"
	BalanceGetLedgerParamsCurrencyFkp BalanceGetLedgerParamsCurrency = "FKP"
	BalanceGetLedgerParamsCurrencyGbp BalanceGetLedgerParamsCurrency = "GBP"
	BalanceGetLedgerParamsCurrencyGel BalanceGetLedgerParamsCurrency = "GEL"
	BalanceGetLedgerParamsCurrencyGhs BalanceGetLedgerParamsCurrency = "GHS"
	BalanceGetLedgerParamsCurrencyGip BalanceGetLedgerParamsCurrency = "GIP"
	BalanceGetLedgerParamsCurrencyGmd BalanceGetLedgerParamsCurrency = "GMD"
	BalanceGetLedgerParamsCurrencyGnf BalanceGetLedgerParamsCurrency = "GNF"
	BalanceGetLedgerParamsCurrencyGtq BalanceGetLedgerParamsCurrency = "GTQ"
	BalanceGetLedgerParamsCurrencyGyd BalanceGetLedgerParamsCurrency = "GYD"
	BalanceGetLedgerParamsCurrencyHkd BalanceGetLedgerParamsCurrency = "HKD"
	BalanceGetLedgerParamsCurrencyHnl BalanceGetLedgerParamsCurrency = "HNL"
	BalanceGetLedgerParamsCurrencyHrk BalanceGetLedgerParamsCurrency = "HRK"
	BalanceGetLedgerParamsCurrencyHtg BalanceGetLedgerParamsCurrency = "HTG"
	BalanceGetLedgerParamsCurrencyHuf BalanceGetLedgerParamsCurrency = "HUF"
	BalanceGetLedgerParamsCurrencyIdr BalanceGetLedgerParamsCurrency = "IDR"
	BalanceGetLedgerParamsCurrencyIls BalanceGetLedgerParamsCurrency = "ILS"
	BalanceGetLedgerParamsCurrencyInr BalanceGetLedgerParamsCurrency = "INR"
	BalanceGetLedgerParamsCurrencyIqd BalanceGetLedgerParamsCurrency = "IQD"
	BalanceGetLedgerParamsCurrencyJmd BalanceGetLedgerParamsCurrency = "JMD"
	BalanceGetLedgerParamsCurrencyJod BalanceGetLedgerParamsCurrency = "JOD"
	BalanceGetLedgerParamsCurrencyJpy BalanceGetLedgerParamsCurrency = "JPY"
	BalanceGetLedgerParamsCurrencyKes BalanceGetLedgerParamsCurrency = "KES"
	BalanceGetLedgerParamsCurrencyKgs BalanceGetLedgerParamsCurrency = "KGS"
	BalanceGetLedgerParamsCurrencyKhr BalanceGetLedgerParamsCurrency = "KHR"
	BalanceGetLedgerParamsCurrencyKmf BalanceGetLedgerParamsCurrency = "KMF"
	BalanceGetLedgerParamsCurrencyKrw BalanceGetLedgerParamsCurrency = "KRW"
	BalanceGetLedgerParamsCurrencyKwd BalanceGetLedgerParamsCurrency = "KWD"
	BalanceGetLedgerParamsCurrencyKyd BalanceGetLedgerParamsCurrency = "KYD"
	BalanceGetLedgerParamsCurrencyKzt BalanceGetLedgerParamsCurrency = "KZT"
	BalanceGetLedgerParamsCurrencyLak BalanceGetLedgerParamsCurrency = "LAK"
	BalanceGetLedgerParamsCurrencyLbp BalanceGetLedgerParamsCurrency = "LBP"
	BalanceGetLedgerParamsCurrencyLkr BalanceGetLedgerParamsCurrency = "LKR"
	BalanceGetLedgerParamsCurrencyLrd BalanceGetLedgerParamsCurrency = "LRD"
	BalanceGetLedgerParamsCurrencyLsl BalanceGetLedgerParamsCurrency = "LSL"
	BalanceGetLedgerParamsCurrencyLyd BalanceGetLedgerParamsCurrency = "LYD"
	BalanceGetLedgerParamsCurrencyMad BalanceGetLedgerParamsCurrency = "MAD"
	BalanceGetLedgerParamsCurrencyMdl BalanceGetLedgerParamsCurrency = "MDL"
	BalanceGetLedgerParamsCurrencyMga BalanceGetLedgerParamsCurrency = "MGA"
	BalanceGetLedgerParamsCurrencyMkd BalanceGetLedgerParamsCurrency = "MKD"
	BalanceGetLedgerParamsCurrencyMmk BalanceGetLedgerParamsCurrency = "MMK"
	BalanceGetLedgerParamsCurrencyMnt BalanceGetLedgerParamsCurrency = "MNT"
	BalanceGetLedgerParamsCurrencyMop BalanceGetLedgerParamsCurrency = "MOP"
	BalanceGetLedgerParamsCurrencyMru BalanceGetLedgerParamsCurrency = "MRU"
	BalanceGetLedgerParamsCurrencyMur BalanceGetLedgerParamsCurrency = "MUR"
	BalanceGetLedgerParamsCurrencyMvr BalanceGetLedgerParamsCurrency = "MVR"
	BalanceGetLedgerParamsCurrencyMwk BalanceGetLedgerParamsCurrency = "MWK"
	BalanceGetLedgerParamsCurrencyMxn BalanceGetLedgerParamsCurrency = "MXN"
	BalanceGetLedgerParamsCurrencyMyr BalanceGetLedgerParamsCurrency = "MYR"
	BalanceGetLedgerParamsCurrencyMzn BalanceGetLedgerParamsCurrency = "MZN"
	BalanceGetLedgerParamsCurrencyNad BalanceGetLedgerParamsCurrency = "NAD"
	BalanceGetLedgerParamsCurrencyNgn BalanceGetLedgerParamsCurrency = "NGN"
	BalanceGetLedgerParamsCurrencyNio BalanceGetLedgerParamsCurrency = "NIO"
	BalanceGetLedgerParamsCurrencyNok BalanceGetLedgerParamsCurrency = "NOK"
	BalanceGetLedgerParamsCurrencyNpr BalanceGetLedgerParamsCurrency = "NPR"
	BalanceGetLedgerParamsCurrencyNzd BalanceGetLedgerParamsCurrency = "NZD"
	BalanceGetLedgerParamsCurrencyOmr BalanceGetLedgerParamsCurrency = "OMR"
	BalanceGetLedgerParamsCurrencyPab BalanceGetLedgerParamsCurrency = "PAB"
	BalanceGetLedgerParamsCurrencyPen BalanceGetLedgerParamsCurrency = "PEN"
	BalanceGetLedgerParamsCurrencyPgk BalanceGetLedgerParamsCurrency = "PGK"
	BalanceGetLedgerParamsCurrencyPhp BalanceGetLedgerParamsCurrency = "PHP"
	BalanceGetLedgerParamsCurrencyPkr BalanceGetLedgerParamsCurrency = "PKR"
	BalanceGetLedgerParamsCurrencyPln BalanceGetLedgerParamsCurrency = "PLN"
	BalanceGetLedgerParamsCurrencyPyg BalanceGetLedgerParamsCurrency = "PYG"
	BalanceGetLedgerParamsCurrencyQar BalanceGetLedgerParamsCurrency = "QAR"
	BalanceGetLedgerParamsCurrencyRon BalanceGetLedgerParamsCurrency = "RON"
	BalanceGetLedgerParamsCurrencyRsd BalanceGetLedgerParamsCurrency = "RSD"
	BalanceGetLedgerParamsCurrencyRub BalanceGetLedgerParamsCurrency = "RUB"
	BalanceGetLedgerParamsCurrencyRwf BalanceGetLedgerParamsCurrency = "RWF"
	BalanceGetLedgerParamsCurrencySar BalanceGetLedgerParamsCurrency = "SAR"
	BalanceGetLedgerParamsCurrencySbd BalanceGetLedgerParamsCurrency = "SBD"
	BalanceGetLedgerParamsCurrencyScr BalanceGetLedgerParamsCurrency = "SCR"
	BalanceGetLedgerParamsCurrencySek BalanceGetLedgerParamsCurrency = "SEK"
	BalanceGetLedgerParamsCurrencySgd BalanceGetLedgerParamsCurrency = "SGD"
	BalanceGetLedgerParamsCurrencyShp BalanceGetLedgerParamsCurrency = "SHP"
	BalanceGetLedgerParamsCurrencySle BalanceGetLedgerParamsCurrency = "SLE"
	BalanceGetLedgerParamsCurrencySll BalanceGetLedgerParamsCurrency = "SLL"
	BalanceGetLedgerParamsCurrencySos BalanceGetLedgerParamsCurrency = "SOS"
	BalanceGetLedgerParamsCurrencySrd BalanceGetLedgerParamsCurrency = "SRD"
	BalanceGetLedgerParamsCurrencySsp BalanceGetLedgerParamsCurrency = "SSP"
	BalanceGetLedgerParamsCurrencyStn BalanceGetLedgerParamsCurrency = "STN"
	BalanceGetLedgerParamsCurrencySvc BalanceGetLedgerParamsCurrency = "SVC"
	BalanceGetLedgerParamsCurrencySzl BalanceGetLedgerParamsCurrency = "SZL"
	BalanceGetLedgerParamsCurrencyThb BalanceGetLedgerParamsCurrency = "THB"
	BalanceGetLedgerParamsCurrencyTnd BalanceGetLedgerParamsCurrency = "TND"
	BalanceGetLedgerParamsCurrencyTop BalanceGetLedgerParamsCurrency = "TOP"
	BalanceGetLedgerParamsCurrencyTry BalanceGetLedgerParamsCurrency = "TRY"
	BalanceGetLedgerParamsCurrencyTtd BalanceGetLedgerParamsCurrency = "TTD"
	BalanceGetLedgerParamsCurrencyTwd BalanceGetLedgerParamsCurrency = "TWD"
	BalanceGetLedgerParamsCurrencyTzs BalanceGetLedgerParamsCurrency = "TZS"
	BalanceGetLedgerParamsCurrencyUah BalanceGetLedgerParamsCurrency = "UAH"
	BalanceGetLedgerParamsCurrencyUgx BalanceGetLedgerParamsCurrency = "UGX"
	BalanceGetLedgerParamsCurrencyUsd BalanceGetLedgerParamsCurrency = "USD"
	BalanceGetLedgerParamsCurrencyUyu BalanceGetLedgerParamsCurrency = "UYU"
	BalanceGetLedgerParamsCurrencyUzs BalanceGetLedgerParamsCurrency = "UZS"
	BalanceGetLedgerParamsCurrencyVes BalanceGetLedgerParamsCurrency = "VES"
	BalanceGetLedgerParamsCurrencyVnd BalanceGetLedgerParamsCurrency = "VND"
	BalanceGetLedgerParamsCurrencyVuv BalanceGetLedgerParamsCurrency = "VUV"
	BalanceGetLedgerParamsCurrencyWst BalanceGetLedgerParamsCurrency = "WST"
	BalanceGetLedgerParamsCurrencyXaf BalanceGetLedgerParamsCurrency = "XAF"
	BalanceGetLedgerParamsCurrencyXcd BalanceGetLedgerParamsCurrency = "XCD"
	BalanceGetLedgerParamsCurrencyXof BalanceGetLedgerParamsCurrency = "XOF"
	BalanceGetLedgerParamsCurrencyXpf BalanceGetLedgerParamsCurrency = "XPF"
	BalanceGetLedgerParamsCurrencyYer BalanceGetLedgerParamsCurrency = "YER"
	BalanceGetLedgerParamsCurrencyZar BalanceGetLedgerParamsCurrency = "ZAR"
	BalanceGetLedgerParamsCurrencyZmw BalanceGetLedgerParamsCurrency = "ZMW"
)

func (r BalanceGetLedgerParamsCurrency) IsKnown() bool {
	switch r {
	case BalanceGetLedgerParamsCurrencyAed, BalanceGetLedgerParamsCurrencyAll, BalanceGetLedgerParamsCurrencyAmd, BalanceGetLedgerParamsCurrencyAng, BalanceGetLedgerParamsCurrencyAoa, BalanceGetLedgerParamsCurrencyArs, BalanceGetLedgerParamsCurrencyAud, BalanceGetLedgerParamsCurrencyAwg, BalanceGetLedgerParamsCurrencyAzn, BalanceGetLedgerParamsCurrencyBam, BalanceGetLedgerParamsCurrencyBbd, BalanceGetLedgerParamsCurrencyBdt, BalanceGetLedgerParamsCurrencyBgn, BalanceGetLedgerParamsCurrencyBhd, BalanceGetLedgerParamsCurrencyBif, BalanceGetLedgerParamsCurrencyBmd, BalanceGetLedgerParamsCurrencyBnd, BalanceGetLedgerParamsCurrencyBob, BalanceGetLedgerParamsCurrencyBrl, BalanceGetLedgerParamsCurrencyBsd, BalanceGetLedgerParamsCurrencyBwp, BalanceGetLedgerParamsCurrencyByn, BalanceGetLedgerParamsCurrencyBzd, BalanceGetLedgerParamsCurrencyCad, BalanceGetLedgerParamsCurrencyChf, BalanceGetLedgerParamsCurrencyClp, BalanceGetLedgerParamsCurrencyCny, BalanceGetLedgerParamsCurrencyCop, BalanceGetLedgerParamsCurrencyCrc, BalanceGetLedgerParamsCurrencyCup, BalanceGetLedgerParamsCurrencyCve, BalanceGetLedgerParamsCurrencyCzk, BalanceGetLedgerParamsCurrencyDjf, BalanceGetLedgerParamsCurrencyDkk, BalanceGetLedgerParamsCurrencyDop, BalanceGetLedgerParamsCurrencyDzd, BalanceGetLedgerParamsCurrencyEgp, BalanceGetLedgerParamsCurrencyEtb, BalanceGetLedgerParamsCurrencyEur, BalanceGetLedgerParamsCurrencyFjd, BalanceGetLedgerParamsCurrencyFkp, BalanceGetLedgerParamsCurrencyGbp, BalanceGetLedgerParamsCurrencyGel, BalanceGetLedgerParamsCurrencyGhs, BalanceGetLedgerParamsCurrencyGip, BalanceGetLedgerParamsCurrencyGmd, BalanceGetLedgerParamsCurrencyGnf, BalanceGetLedgerParamsCurrencyGtq, BalanceGetLedgerParamsCurrencyGyd, BalanceGetLedgerParamsCurrencyHkd, BalanceGetLedgerParamsCurrencyHnl, BalanceGetLedgerParamsCurrencyHrk, BalanceGetLedgerParamsCurrencyHtg, BalanceGetLedgerParamsCurrencyHuf, BalanceGetLedgerParamsCurrencyIdr, BalanceGetLedgerParamsCurrencyIls, BalanceGetLedgerParamsCurrencyInr, BalanceGetLedgerParamsCurrencyIqd, BalanceGetLedgerParamsCurrencyJmd, BalanceGetLedgerParamsCurrencyJod, BalanceGetLedgerParamsCurrencyJpy, BalanceGetLedgerParamsCurrencyKes, BalanceGetLedgerParamsCurrencyKgs, BalanceGetLedgerParamsCurrencyKhr, BalanceGetLedgerParamsCurrencyKmf, BalanceGetLedgerParamsCurrencyKrw, BalanceGetLedgerParamsCurrencyKwd, BalanceGetLedgerParamsCurrencyKyd, BalanceGetLedgerParamsCurrencyKzt, BalanceGetLedgerParamsCurrencyLak, BalanceGetLedgerParamsCurrencyLbp, BalanceGetLedgerParamsCurrencyLkr, BalanceGetLedgerParamsCurrencyLrd, BalanceGetLedgerParamsCurrencyLsl, BalanceGetLedgerParamsCurrencyLyd, BalanceGetLedgerParamsCurrencyMad, BalanceGetLedgerParamsCurrencyMdl, BalanceGetLedgerParamsCurrencyMga, BalanceGetLedgerParamsCurrencyMkd, BalanceGetLedgerParamsCurrencyMmk, BalanceGetLedgerParamsCurrencyMnt, BalanceGetLedgerParamsCurrencyMop, BalanceGetLedgerParamsCurrencyMru, BalanceGetLedgerParamsCurrencyMur, BalanceGetLedgerParamsCurrencyMvr, BalanceGetLedgerParamsCurrencyMwk, BalanceGetLedgerParamsCurrencyMxn, BalanceGetLedgerParamsCurrencyMyr, BalanceGetLedgerParamsCurrencyMzn, BalanceGetLedgerParamsCurrencyNad, BalanceGetLedgerParamsCurrencyNgn, BalanceGetLedgerParamsCurrencyNio, BalanceGetLedgerParamsCurrencyNok, BalanceGetLedgerParamsCurrencyNpr, BalanceGetLedgerParamsCurrencyNzd, BalanceGetLedgerParamsCurrencyOmr, BalanceGetLedgerParamsCurrencyPab, BalanceGetLedgerParamsCurrencyPen, BalanceGetLedgerParamsCurrencyPgk, BalanceGetLedgerParamsCurrencyPhp, BalanceGetLedgerParamsCurrencyPkr, BalanceGetLedgerParamsCurrencyPln, BalanceGetLedgerParamsCurrencyPyg, BalanceGetLedgerParamsCurrencyQar, BalanceGetLedgerParamsCurrencyRon, BalanceGetLedgerParamsCurrencyRsd, BalanceGetLedgerParamsCurrencyRub, BalanceGetLedgerParamsCurrencyRwf, BalanceGetLedgerParamsCurrencySar, BalanceGetLedgerParamsCurrencySbd, BalanceGetLedgerParamsCurrencyScr, BalanceGetLedgerParamsCurrencySek, BalanceGetLedgerParamsCurrencySgd, BalanceGetLedgerParamsCurrencyShp, BalanceGetLedgerParamsCurrencySle, BalanceGetLedgerParamsCurrencySll, BalanceGetLedgerParamsCurrencySos, BalanceGetLedgerParamsCurrencySrd, BalanceGetLedgerParamsCurrencySsp, BalanceGetLedgerParamsCurrencyStn, BalanceGetLedgerParamsCurrencySvc, BalanceGetLedgerParamsCurrencySzl, BalanceGetLedgerParamsCurrencyThb, BalanceGetLedgerParamsCurrencyTnd, BalanceGetLedgerParamsCurrencyTop, BalanceGetLedgerParamsCurrencyTry, BalanceGetLedgerParamsCurrencyTtd, BalanceGetLedgerParamsCurrencyTwd, BalanceGetLedgerParamsCurrencyTzs, BalanceGetLedgerParamsCurrencyUah, BalanceGetLedgerParamsCurrencyUgx, BalanceGetLedgerParamsCurrencyUsd, BalanceGetLedgerParamsCurrencyUyu, BalanceGetLedgerParamsCurrencyUzs, BalanceGetLedgerParamsCurrencyVes, BalanceGetLedgerParamsCurrencyVnd, BalanceGetLedgerParamsCurrencyVuv, BalanceGetLedgerParamsCurrencyWst, BalanceGetLedgerParamsCurrencyXaf, BalanceGetLedgerParamsCurrencyXcd, BalanceGetLedgerParamsCurrencyXof, BalanceGetLedgerParamsCurrencyXpf, BalanceGetLedgerParamsCurrencyYer, BalanceGetLedgerParamsCurrencyZar, BalanceGetLedgerParamsCurrencyZmw:
		return true
	}
	return false
}

// Filter by Ledger Event Type
type BalanceGetLedgerParamsEventType string

const (
	BalanceGetLedgerParamsEventTypePayment            BalanceGetLedgerParamsEventType = "payment"
	BalanceGetLedgerParamsEventTypeRefund             BalanceGetLedgerParamsEventType = "refund"
	BalanceGetLedgerParamsEventTypeRefundReversal     BalanceGetLedgerParamsEventType = "refund_reversal"
	BalanceGetLedgerParamsEventTypeDispute            BalanceGetLedgerParamsEventType = "dispute"
	BalanceGetLedgerParamsEventTypeDisputeReversal    BalanceGetLedgerParamsEventType = "dispute_reversal"
	BalanceGetLedgerParamsEventTypeTax                BalanceGetLedgerParamsEventType = "tax"
	BalanceGetLedgerParamsEventTypeTaxReversal        BalanceGetLedgerParamsEventType = "tax_reversal"
	BalanceGetLedgerParamsEventTypePaymentFees        BalanceGetLedgerParamsEventType = "payment_fees"
	BalanceGetLedgerParamsEventTypeRefundFees         BalanceGetLedgerParamsEventType = "refund_fees"
	BalanceGetLedgerParamsEventTypeRefundFeesReversal BalanceGetLedgerParamsEventType = "refund_fees_reversal"
	BalanceGetLedgerParamsEventTypeDisputeFees        BalanceGetLedgerParamsEventType = "dispute_fees"
	BalanceGetLedgerParamsEventTypePayout             BalanceGetLedgerParamsEventType = "payout"
	BalanceGetLedgerParamsEventTypePayoutFees         BalanceGetLedgerParamsEventType = "payout_fees"
	BalanceGetLedgerParamsEventTypePayoutReversal     BalanceGetLedgerParamsEventType = "payout_reversal"
	BalanceGetLedgerParamsEventTypePayoutFeesReversal BalanceGetLedgerParamsEventType = "payout_fees_reversal"
	BalanceGetLedgerParamsEventTypeDodoCredits        BalanceGetLedgerParamsEventType = "dodo_credits"
	BalanceGetLedgerParamsEventTypeAdjustment         BalanceGetLedgerParamsEventType = "adjustment"
	BalanceGetLedgerParamsEventTypeCurrencyConversion BalanceGetLedgerParamsEventType = "currency_conversion"
)

func (r BalanceGetLedgerParamsEventType) IsKnown() bool {
	switch r {
	case BalanceGetLedgerParamsEventTypePayment, BalanceGetLedgerParamsEventTypeRefund, BalanceGetLedgerParamsEventTypeRefundReversal, BalanceGetLedgerParamsEventTypeDispute, BalanceGetLedgerParamsEventTypeDisputeReversal, BalanceGetLedgerParamsEventTypeTax, BalanceGetLedgerParamsEventTypeTaxReversal, BalanceGetLedgerParamsEventTypePaymentFees, BalanceGetLedgerParamsEventTypeRefundFees, BalanceGetLedgerParamsEventTypeRefundFeesReversal, BalanceGetLedgerParamsEventTypeDisputeFees, BalanceGetLedgerParamsEventTypePayout, BalanceGetLedgerParamsEventTypePayoutFees, BalanceGetLedgerParamsEventTypePayoutReversal, BalanceGetLedgerParamsEventTypePayoutFeesReversal, BalanceGetLedgerParamsEventTypeDodoCredits, BalanceGetLedgerParamsEventTypeAdjustment, BalanceGetLedgerParamsEventTypeCurrencyConversion:
		return true
	}
	return false
}
