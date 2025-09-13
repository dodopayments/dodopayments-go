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
func NewCustomerWalletLedgerEntryService(opts ...option.RequestOption) (r *CustomerWalletLedgerEntryService) {
	r = &CustomerWalletLedgerEntryService{}
	r.Options = opts
	return
}

func (r *CustomerWalletLedgerEntryService) New(ctx context.Context, customerID string, body CustomerWalletLedgerEntryNewParams, opts ...option.RequestOption) (res *CustomerWallet, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	ID                string                             `json:"id,required"`
	AfterBalance      int64                              `json:"after_balance,required"`
	Amount            int64                              `json:"amount,required"`
	BeforeBalance     int64                              `json:"before_balance,required"`
	BusinessID        string                             `json:"business_id,required"`
	CreatedAt         time.Time                          `json:"created_at,required" format:"date-time"`
	Currency          Currency                           `json:"currency,required"`
	CustomerID        string                             `json:"customer_id,required"`
	EventType         CustomerWalletTransactionEventType `json:"event_type,required"`
	IsCredit          bool                               `json:"is_credit,required"`
	Reason            string                             `json:"reason,nullable"`
	ReferenceObjectID string                             `json:"reference_object_id,nullable"`
	JSON              customerWalletTransactionJSON      `json:"-"`
}

// customerWalletTransactionJSON contains the JSON metadata for the struct
// [CustomerWalletTransaction]
type customerWalletTransactionJSON struct {
	ID                apijson.Field
	AfterBalance      apijson.Field
	Amount            apijson.Field
	BeforeBalance     apijson.Field
	BusinessID        apijson.Field
	CreatedAt         apijson.Field
	Currency          apijson.Field
	CustomerID        apijson.Field
	EventType         apijson.Field
	IsCredit          apijson.Field
	Reason            apijson.Field
	ReferenceObjectID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerWalletTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerWalletTransactionJSON) RawJSON() string {
	return r.raw
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

func (r CustomerWalletTransactionEventType) IsKnown() bool {
	switch r {
	case CustomerWalletTransactionEventTypePayment, CustomerWalletTransactionEventTypePaymentReversal, CustomerWalletTransactionEventTypeRefund, CustomerWalletTransactionEventTypeRefundReversal, CustomerWalletTransactionEventTypeDispute, CustomerWalletTransactionEventTypeDisputeReversal, CustomerWalletTransactionEventTypeMerchantAdjustment:
		return true
	}
	return false
}

type CustomerWalletLedgerEntryNewParams struct {
	Amount param.Field[int64] `json:"amount,required"`
	// Currency of the wallet to adjust
	Currency param.Field[Currency] `json:"currency,required"`
	// Type of ledger entry - credit or debit
	EntryType param.Field[CustomerWalletLedgerEntryNewParamsEntryType] `json:"entry_type,required"`
	// Optional idempotency key to prevent duplicate entries
	IdempotencyKey param.Field[string] `json:"idempotency_key"`
	Reason         param.Field[string] `json:"reason"`
}

func (r CustomerWalletLedgerEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of ledger entry - credit or debit
type CustomerWalletLedgerEntryNewParamsEntryType string

const (
	CustomerWalletLedgerEntryNewParamsEntryTypeCredit CustomerWalletLedgerEntryNewParamsEntryType = "credit"
	CustomerWalletLedgerEntryNewParamsEntryTypeDebit  CustomerWalletLedgerEntryNewParamsEntryType = "debit"
)

func (r CustomerWalletLedgerEntryNewParamsEntryType) IsKnown() bool {
	switch r {
	case CustomerWalletLedgerEntryNewParamsEntryTypeCredit, CustomerWalletLedgerEntryNewParamsEntryTypeDebit:
		return true
	}
	return false
}

type CustomerWalletLedgerEntryListParams struct {
	// Optional currency filter
	Currency   param.Field[Currency] `query:"currency"`
	PageNumber param.Field[int64]    `query:"page_number"`
	PageSize   param.Field[int64]    `query:"page_size"`
}

// URLQuery serializes [CustomerWalletLedgerEntryListParams]'s query parameters as
// `url.Values`.
func (r CustomerWalletLedgerEntryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
