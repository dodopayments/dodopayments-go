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
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
)

// CreditEntitlementBalanceService contains methods and other services that help
// with interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditEntitlementBalanceService] method instead.
type CreditEntitlementBalanceService struct {
	Options []option.RequestOption
}

// NewCreditEntitlementBalanceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCreditEntitlementBalanceService(opts ...option.RequestOption) (r *CreditEntitlementBalanceService) {
	r = &CreditEntitlementBalanceService{}
	r.Options = opts
	return
}

// Returns the credit balance details for a specific customer and credit
// entitlement.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
// - `customer_id` - The unique identifier of the customer
//
// # Responses
//
// - `200 OK` - Returns the customer's balance
// - `404 Not Found` - Credit entitlement or customer balance not found
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) Get(ctx context.Context, creditEntitlementID string, customerID string, opts ...option.RequestOption) (res *CustomerCreditBalance, err error) {
	opts = slices.Concat(r.Options, opts)
	if creditEntitlementID == "" {
		err = errors.New("missing required credit_entitlement_id parameter")
		return
	}
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s/balances/%s", creditEntitlementID, customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of customer credit balances for the given credit
// entitlement.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
//
// # Query Parameters
//
// - `page_size` - Number of items per page (default: 10, max: 100)
// - `page_number` - Zero-based page number (default: 0)
// - `customer_id` - Optional filter by specific customer
//
// # Responses
//
// - `200 OK` - Returns list of customer balances
// - `404 Not Found` - Credit entitlement not found
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) List(ctx context.Context, creditEntitlementID string, query CreditEntitlementBalanceListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[CustomerCreditBalance], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if creditEntitlementID == "" {
		err = errors.New("missing required credit_entitlement_id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s/balances", creditEntitlementID)
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

// Returns a paginated list of customer credit balances for the given credit
// entitlement.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
//
// # Query Parameters
//
// - `page_size` - Number of items per page (default: 10, max: 100)
// - `page_number` - Zero-based page number (default: 0)
// - `customer_id` - Optional filter by specific customer
//
// # Responses
//
// - `200 OK` - Returns list of customer balances
// - `404 Not Found` - Credit entitlement not found
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) ListAutoPaging(ctx context.Context, creditEntitlementID string, query CreditEntitlementBalanceListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[CustomerCreditBalance] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, creditEntitlementID, query, opts...))
}

// For credit entries, a new grant is created. For debit entries, credits are
// deducted from existing grants using FIFO (oldest first).
//
// # Authentication
//
// Requires an API key with `Editor` role.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
// - `customer_id` - The unique identifier of the customer
//
// # Request Body
//
// - `entry_type` - "credit" or "debit"
// - `amount` - Amount to credit or debit
// - `reason` - Optional human-readable reason
// - `expires_at` - Optional expiration for credited amount (only for credit type)
// - `idempotency_key` - Optional key to prevent duplicate entries
//
// # Responses
//
// - `201 Created` - Ledger entry created successfully
// - `400 Bad Request` - Invalid request (e.g., debit with insufficient balance)
// - `404 Not Found` - Credit entitlement or customer not found
// - `409 Conflict` - Idempotency key already exists
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) NewLedgerEntry(ctx context.Context, creditEntitlementID string, customerID string, body CreditEntitlementBalanceNewLedgerEntryParams, opts ...option.RequestOption) (res *CreditEntitlementBalanceNewLedgerEntryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if creditEntitlementID == "" {
		err = errors.New("missing required credit_entitlement_id parameter")
		return
	}
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s/balances/%s/ledger-entries", creditEntitlementID, customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a paginated list of credit grants with optional filtering by status.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
// - `customer_id` - The unique identifier of the customer
//
// # Query Parameters
//
// - `page_size` - Number of items per page (default: 10, max: 100)
// - `page_number` - Zero-based page number (default: 0)
// - `status` - Filter by status: active, expired, depleted
//
// # Responses
//
// - `200 OK` - Returns list of grants
// - `404 Not Found` - Credit entitlement not found
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) ListGrants(ctx context.Context, creditEntitlementID string, customerID string, query CreditEntitlementBalanceListGrantsParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[CreditEntitlementBalanceListGrantsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if creditEntitlementID == "" {
		err = errors.New("missing required credit_entitlement_id parameter")
		return
	}
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s/balances/%s/grants", creditEntitlementID, customerID)
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

// Returns a paginated list of credit grants with optional filtering by status.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
// - `customer_id` - The unique identifier of the customer
//
// # Query Parameters
//
// - `page_size` - Number of items per page (default: 10, max: 100)
// - `page_number` - Zero-based page number (default: 0)
// - `status` - Filter by status: active, expired, depleted
//
// # Responses
//
// - `200 OK` - Returns list of grants
// - `404 Not Found` - Credit entitlement not found
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) ListGrantsAutoPaging(ctx context.Context, creditEntitlementID string, customerID string, query CreditEntitlementBalanceListGrantsParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[CreditEntitlementBalanceListGrantsResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.ListGrants(ctx, creditEntitlementID, customerID, query, opts...))
}

// Returns a paginated list of credit transaction history with optional filtering.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
// - `customer_id` - The unique identifier of the customer
//
// # Query Parameters
//
// - `page_size` - Number of items per page (default: 10, max: 100)
// - `page_number` - Zero-based page number (default: 0)
// - `transaction_type` - Filter by transaction type
// - `start_date` - Filter entries from this date
// - `end_date` - Filter entries until this date
//
// # Responses
//
// - `200 OK` - Returns list of ledger entries
// - `404 Not Found` - Credit entitlement not found
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) ListLedger(ctx context.Context, creditEntitlementID string, customerID string, query CreditEntitlementBalanceListLedgerParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[CreditLedgerEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if creditEntitlementID == "" {
		err = errors.New("missing required credit_entitlement_id parameter")
		return
	}
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s/balances/%s/ledger", creditEntitlementID, customerID)
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

// Returns a paginated list of credit transaction history with optional filtering.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `credit_entitlement_id` - The unique identifier of the credit entitlement
// - `customer_id` - The unique identifier of the customer
//
// # Query Parameters
//
// - `page_size` - Number of items per page (default: 10, max: 100)
// - `page_number` - Zero-based page number (default: 0)
// - `transaction_type` - Filter by transaction type
// - `start_date` - Filter entries from this date
// - `end_date` - Filter entries until this date
//
// # Responses
//
// - `200 OK` - Returns list of ledger entries
// - `404 Not Found` - Credit entitlement not found
// - `500 Internal Server Error` - Database or server error
func (r *CreditEntitlementBalanceService) ListLedgerAutoPaging(ctx context.Context, creditEntitlementID string, customerID string, query CreditEntitlementBalanceListLedgerParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[CreditLedgerEntry] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.ListLedger(ctx, creditEntitlementID, customerID, query, opts...))
}

// Response for a ledger entry
type CreditLedgerEntry struct {
	ID                  string                           `json:"id" api:"required"`
	Amount              string                           `json:"amount" api:"required"`
	BalanceAfter        string                           `json:"balance_after" api:"required"`
	BalanceBefore       string                           `json:"balance_before" api:"required"`
	BusinessID          string                           `json:"business_id" api:"required"`
	CreatedAt           time.Time                        `json:"created_at" api:"required" format:"date-time"`
	CreditEntitlementID string                           `json:"credit_entitlement_id" api:"required"`
	CustomerID          string                           `json:"customer_id" api:"required"`
	IsCredit            bool                             `json:"is_credit" api:"required"`
	OverageAfter        string                           `json:"overage_after" api:"required"`
	OverageBefore       string                           `json:"overage_before" api:"required"`
	TransactionType     CreditLedgerEntryTransactionType `json:"transaction_type" api:"required"`
	Description         string                           `json:"description" api:"nullable"`
	GrantID             string                           `json:"grant_id" api:"nullable"`
	ReferenceID         string                           `json:"reference_id" api:"nullable"`
	ReferenceType       string                           `json:"reference_type" api:"nullable"`
	JSON                creditLedgerEntryJSON            `json:"-"`
}

// creditLedgerEntryJSON contains the JSON metadata for the struct
// [CreditLedgerEntry]
type creditLedgerEntryJSON struct {
	ID                  apijson.Field
	Amount              apijson.Field
	BalanceAfter        apijson.Field
	BalanceBefore       apijson.Field
	BusinessID          apijson.Field
	CreatedAt           apijson.Field
	CreditEntitlementID apijson.Field
	CustomerID          apijson.Field
	IsCredit            apijson.Field
	OverageAfter        apijson.Field
	OverageBefore       apijson.Field
	TransactionType     apijson.Field
	Description         apijson.Field
	GrantID             apijson.Field
	ReferenceID         apijson.Field
	ReferenceType       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CreditLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerEntryJSON) RawJSON() string {
	return r.raw
}

type CreditLedgerEntryTransactionType string

const (
	CreditLedgerEntryTransactionTypeCreditAdded       CreditLedgerEntryTransactionType = "credit_added"
	CreditLedgerEntryTransactionTypeCreditDeducted    CreditLedgerEntryTransactionType = "credit_deducted"
	CreditLedgerEntryTransactionTypeCreditExpired     CreditLedgerEntryTransactionType = "credit_expired"
	CreditLedgerEntryTransactionTypeCreditRolledOver  CreditLedgerEntryTransactionType = "credit_rolled_over"
	CreditLedgerEntryTransactionTypeRolloverForfeited CreditLedgerEntryTransactionType = "rollover_forfeited"
	CreditLedgerEntryTransactionTypeOverageCharged    CreditLedgerEntryTransactionType = "overage_charged"
	CreditLedgerEntryTransactionTypeAutoTopUp         CreditLedgerEntryTransactionType = "auto_top_up"
	CreditLedgerEntryTransactionTypeManualAdjustment  CreditLedgerEntryTransactionType = "manual_adjustment"
	CreditLedgerEntryTransactionTypeRefund            CreditLedgerEntryTransactionType = "refund"
)

func (r CreditLedgerEntryTransactionType) IsKnown() bool {
	switch r {
	case CreditLedgerEntryTransactionTypeCreditAdded, CreditLedgerEntryTransactionTypeCreditDeducted, CreditLedgerEntryTransactionTypeCreditExpired, CreditLedgerEntryTransactionTypeCreditRolledOver, CreditLedgerEntryTransactionTypeRolloverForfeited, CreditLedgerEntryTransactionTypeOverageCharged, CreditLedgerEntryTransactionTypeAutoTopUp, CreditLedgerEntryTransactionTypeManualAdjustment, CreditLedgerEntryTransactionTypeRefund:
		return true
	}
	return false
}

// Response for a customer's credit balance
type CustomerCreditBalance struct {
	ID                  string                    `json:"id" api:"required"`
	Balance             string                    `json:"balance" api:"required"`
	CreatedAt           time.Time                 `json:"created_at" api:"required" format:"date-time"`
	CreditEntitlementID string                    `json:"credit_entitlement_id" api:"required"`
	CustomerID          string                    `json:"customer_id" api:"required"`
	Overage             string                    `json:"overage" api:"required"`
	UpdatedAt           time.Time                 `json:"updated_at" api:"required" format:"date-time"`
	LastTransactionAt   time.Time                 `json:"last_transaction_at" api:"nullable" format:"date-time"`
	JSON                customerCreditBalanceJSON `json:"-"`
}

// customerCreditBalanceJSON contains the JSON metadata for the struct
// [CustomerCreditBalance]
type customerCreditBalanceJSON struct {
	ID                  apijson.Field
	Balance             apijson.Field
	CreatedAt           apijson.Field
	CreditEntitlementID apijson.Field
	CustomerID          apijson.Field
	Overage             apijson.Field
	UpdatedAt           apijson.Field
	LastTransactionAt   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerCreditBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceJSON) RawJSON() string {
	return r.raw
}

type LedgerEntryType string

const (
	LedgerEntryTypeCredit LedgerEntryType = "credit"
	LedgerEntryTypeDebit  LedgerEntryType = "debit"
)

func (r LedgerEntryType) IsKnown() bool {
	switch r {
	case LedgerEntryTypeCredit, LedgerEntryTypeDebit:
		return true
	}
	return false
}

// Response for creating a ledger entry
type CreditEntitlementBalanceNewLedgerEntryResponse struct {
	ID                  string                                             `json:"id" api:"required"`
	Amount              string                                             `json:"amount" api:"required"`
	BalanceAfter        string                                             `json:"balance_after" api:"required"`
	BalanceBefore       string                                             `json:"balance_before" api:"required"`
	CreatedAt           time.Time                                          `json:"created_at" api:"required" format:"date-time"`
	CreditEntitlementID string                                             `json:"credit_entitlement_id" api:"required"`
	CustomerID          string                                             `json:"customer_id" api:"required"`
	EntryType           LedgerEntryType                                    `json:"entry_type" api:"required"`
	IsCredit            bool                                               `json:"is_credit" api:"required"`
	OverageAfter        string                                             `json:"overage_after" api:"required"`
	OverageBefore       string                                             `json:"overage_before" api:"required"`
	GrantID             string                                             `json:"grant_id" api:"nullable"`
	Reason              string                                             `json:"reason" api:"nullable"`
	JSON                creditEntitlementBalanceNewLedgerEntryResponseJSON `json:"-"`
}

// creditEntitlementBalanceNewLedgerEntryResponseJSON contains the JSON metadata
// for the struct [CreditEntitlementBalanceNewLedgerEntryResponse]
type creditEntitlementBalanceNewLedgerEntryResponseJSON struct {
	ID                  apijson.Field
	Amount              apijson.Field
	BalanceAfter        apijson.Field
	BalanceBefore       apijson.Field
	CreatedAt           apijson.Field
	CreditEntitlementID apijson.Field
	CustomerID          apijson.Field
	EntryType           apijson.Field
	IsCredit            apijson.Field
	OverageAfter        apijson.Field
	OverageBefore       apijson.Field
	GrantID             apijson.Field
	Reason              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CreditEntitlementBalanceNewLedgerEntryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditEntitlementBalanceNewLedgerEntryResponseJSON) RawJSON() string {
	return r.raw
}

// Response for a credit grant
type CreditEntitlementBalanceListGrantsResponse struct {
	ID                  string                                               `json:"id" api:"required"`
	CreatedAt           time.Time                                            `json:"created_at" api:"required" format:"date-time"`
	CreditEntitlementID string                                               `json:"credit_entitlement_id" api:"required"`
	CustomerID          string                                               `json:"customer_id" api:"required"`
	InitialAmount       string                                               `json:"initial_amount" api:"required"`
	IsExpired           bool                                                 `json:"is_expired" api:"required"`
	IsRolledOver        bool                                                 `json:"is_rolled_over" api:"required"`
	RemainingAmount     string                                               `json:"remaining_amount" api:"required"`
	RolloverCount       int64                                                `json:"rollover_count" api:"required"`
	SourceType          CreditEntitlementBalanceListGrantsResponseSourceType `json:"source_type" api:"required"`
	UpdatedAt           time.Time                                            `json:"updated_at" api:"required" format:"date-time"`
	ExpiresAt           time.Time                                            `json:"expires_at" api:"nullable" format:"date-time"`
	Metadata            map[string]string                                    `json:"metadata" api:"nullable"`
	ParentGrantID       string                                               `json:"parent_grant_id" api:"nullable"`
	SourceID            string                                               `json:"source_id" api:"nullable"`
	JSON                creditEntitlementBalanceListGrantsResponseJSON       `json:"-"`
}

// creditEntitlementBalanceListGrantsResponseJSON contains the JSON metadata for
// the struct [CreditEntitlementBalanceListGrantsResponse]
type creditEntitlementBalanceListGrantsResponseJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	CreditEntitlementID apijson.Field
	CustomerID          apijson.Field
	InitialAmount       apijson.Field
	IsExpired           apijson.Field
	IsRolledOver        apijson.Field
	RemainingAmount     apijson.Field
	RolloverCount       apijson.Field
	SourceType          apijson.Field
	UpdatedAt           apijson.Field
	ExpiresAt           apijson.Field
	Metadata            apijson.Field
	ParentGrantID       apijson.Field
	SourceID            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CreditEntitlementBalanceListGrantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditEntitlementBalanceListGrantsResponseJSON) RawJSON() string {
	return r.raw
}

type CreditEntitlementBalanceListGrantsResponseSourceType string

const (
	CreditEntitlementBalanceListGrantsResponseSourceTypeSubscription CreditEntitlementBalanceListGrantsResponseSourceType = "subscription"
	CreditEntitlementBalanceListGrantsResponseSourceTypeOneTime      CreditEntitlementBalanceListGrantsResponseSourceType = "one_time"
	CreditEntitlementBalanceListGrantsResponseSourceTypeAddon        CreditEntitlementBalanceListGrantsResponseSourceType = "addon"
	CreditEntitlementBalanceListGrantsResponseSourceTypeAPI          CreditEntitlementBalanceListGrantsResponseSourceType = "api"
	CreditEntitlementBalanceListGrantsResponseSourceTypeRollover     CreditEntitlementBalanceListGrantsResponseSourceType = "rollover"
)

func (r CreditEntitlementBalanceListGrantsResponseSourceType) IsKnown() bool {
	switch r {
	case CreditEntitlementBalanceListGrantsResponseSourceTypeSubscription, CreditEntitlementBalanceListGrantsResponseSourceTypeOneTime, CreditEntitlementBalanceListGrantsResponseSourceTypeAddon, CreditEntitlementBalanceListGrantsResponseSourceTypeAPI, CreditEntitlementBalanceListGrantsResponseSourceTypeRollover:
		return true
	}
	return false
}

type CreditEntitlementBalanceListParams struct {
	// Filter by specific customer ID
	CustomerID param.Field[string] `query:"customer_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [CreditEntitlementBalanceListParams]'s query parameters as
// `url.Values`.
func (r CreditEntitlementBalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreditEntitlementBalanceNewLedgerEntryParams struct {
	// Amount to credit or debit
	Amount param.Field[string] `json:"amount" api:"required"`
	// Entry type: credit or debit
	EntryType param.Field[LedgerEntryType] `json:"entry_type" api:"required"`
	// Expiration for credited amount (only for credit type)
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// Idempotency key to prevent duplicate entries
	IdempotencyKey param.Field[string] `json:"idempotency_key"`
	// Optional metadata (max 50 key-value pairs, key max 40 chars, value max 500
	// chars)
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Human-readable reason for the entry
	Reason param.Field[string] `json:"reason"`
}

func (r CreditEntitlementBalanceNewLedgerEntryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditEntitlementBalanceListGrantsParams struct {
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by grant status: active, expired, depleted
	Status param.Field[CreditEntitlementBalanceListGrantsParamsStatus] `query:"status"`
}

// URLQuery serializes [CreditEntitlementBalanceListGrantsParams]'s query
// parameters as `url.Values`.
func (r CreditEntitlementBalanceListGrantsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by grant status: active, expired, depleted
type CreditEntitlementBalanceListGrantsParamsStatus string

const (
	CreditEntitlementBalanceListGrantsParamsStatusActive   CreditEntitlementBalanceListGrantsParamsStatus = "active"
	CreditEntitlementBalanceListGrantsParamsStatusExpired  CreditEntitlementBalanceListGrantsParamsStatus = "expired"
	CreditEntitlementBalanceListGrantsParamsStatusDepleted CreditEntitlementBalanceListGrantsParamsStatus = "depleted"
)

func (r CreditEntitlementBalanceListGrantsParamsStatus) IsKnown() bool {
	switch r {
	case CreditEntitlementBalanceListGrantsParamsStatusActive, CreditEntitlementBalanceListGrantsParamsStatusExpired, CreditEntitlementBalanceListGrantsParamsStatusDepleted:
		return true
	}
	return false
}

type CreditEntitlementBalanceListLedgerParams struct {
	// Filter by end date
	EndDate param.Field[time.Time] `query:"end_date" format:"date-time"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by start date
	StartDate param.Field[time.Time] `query:"start_date" format:"date-time"`
	// Filter by transaction type (snake_case: credit_added, credit_deducted,
	// credit_expired, etc.)
	TransactionType param.Field[string] `query:"transaction_type"`
}

// URLQuery serializes [CreditEntitlementBalanceListLedgerParams]'s query
// parameters as `url.Values`.
func (r CreditEntitlementBalanceListLedgerParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
