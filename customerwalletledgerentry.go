// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
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

// CustomerWalletLedgerEntryService contains methods and other services that help
// with interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerWalletLedgerEntryService] method instead.
type CustomerWalletLedgerEntryService struct {
	Options []option.RequestOption
}

// NewCustomerWalletLedgerEntryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomerWalletLedgerEntryService(opts ...option.RequestOption) (r CustomerWalletLedgerEntryService) {
	r = CustomerWalletLedgerEntryService{}
	r.Options = opts
	return
}

func (r *CustomerWalletLedgerEntryService) New(ctx context.Context, customerID string, body CustomerWalletLedgerEntryNewParams, opts ...option.RequestOption) (res *CustomerWallet, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/wallets/ledger-entries", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *CustomerWalletLedgerEntryService) List(ctx context.Context, customerID string, query CustomerWalletLedgerEntryListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[CustomerWalletTransaction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/wallets/ledger-entries", customerID)
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

func (r *CustomerWalletLedgerEntryService) ListAutoPaging(ctx context.Context, customerID string, query CustomerWalletLedgerEntryListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[CustomerWalletTransaction] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, customerID, query, opts...))
}

type CustomerWalletTransaction struct {
	ID            string    `json:"id,required"`
	AfterBalance  int64     `json:"after_balance,required"`
	Amount        int64     `json:"amount,required"`
	BeforeBalance int64     `json:"before_balance,required"`
	BusinessID    string    `json:"business_id,required"`
	CreatedAt     time.Time `json:"created_at,required" format:"date-time"`
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
	Currency   Currency `json:"currency,required"`
	CustomerID string   `json:"customer_id,required"`
	// Any of "payment", "payment_reversal", "refund", "refund_reversal", "dispute",
	// "dispute_reversal", "merchant_adjustment".
	EventType         CustomerWalletTransactionEventType `json:"event_type,required"`
	IsCredit          bool                               `json:"is_credit,required"`
	Reason            string                             `json:"reason,nullable"`
	ReferenceObjectID string                             `json:"reference_object_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AfterBalance      respjson.Field
		Amount            respjson.Field
		BeforeBalance     respjson.Field
		BusinessID        respjson.Field
		CreatedAt         respjson.Field
		Currency          respjson.Field
		CustomerID        respjson.Field
		EventType         respjson.Field
		IsCredit          respjson.Field
		Reason            respjson.Field
		ReferenceObjectID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerWalletTransaction) RawJSON() string { return r.JSON.raw }
func (r *CustomerWalletTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerWalletTransactionEventType string

const (
	CustomerWalletTransactionEventTypePayment            CustomerWalletTransactionEventType = "payment"
	CustomerWalletTransactionEventTypePaymentReversal    CustomerWalletTransactionEventType = "payment_reversal"
	CustomerWalletTransactionEventTypeRefund             CustomerWalletTransactionEventType = "refund"
	CustomerWalletTransactionEventTypeRefundReversal     CustomerWalletTransactionEventType = "refund_reversal"
	CustomerWalletTransactionEventTypeDispute            CustomerWalletTransactionEventType = "dispute"
	CustomerWalletTransactionEventTypeDisputeReversal    CustomerWalletTransactionEventType = "dispute_reversal"
	CustomerWalletTransactionEventTypeMerchantAdjustment CustomerWalletTransactionEventType = "merchant_adjustment"
)

type CustomerWalletLedgerEntryNewParams struct {
	Amount int64 `json:"amount,required"`
	// Currency of the wallet to adjust
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
	// Type of ledger entry - credit or debit
	//
	// Any of "credit", "debit".
	EntryType CustomerWalletLedgerEntryNewParamsEntryType `json:"entry_type,omitzero,required"`
	// Optional idempotency key to prevent duplicate entries
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	Reason         param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r CustomerWalletLedgerEntryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerWalletLedgerEntryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerWalletLedgerEntryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of ledger entry - credit or debit
type CustomerWalletLedgerEntryNewParamsEntryType string

const (
	CustomerWalletLedgerEntryNewParamsEntryTypeCredit CustomerWalletLedgerEntryNewParamsEntryType = "credit"
	CustomerWalletLedgerEntryNewParamsEntryTypeDebit  CustomerWalletLedgerEntryNewParamsEntryType = "debit"
)

type CustomerWalletLedgerEntryListParams struct {
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Optional currency filter
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
	Currency Currency `query:"currency,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerWalletLedgerEntryListParams]'s query parameters as
// `url.Values`.
func (r CustomerWalletLedgerEntryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
