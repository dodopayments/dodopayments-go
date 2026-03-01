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

// CreditEntitlementService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditEntitlementService] method instead.
type CreditEntitlementService struct {
	Options  []option.RequestOption
	Balances *CreditEntitlementBalanceService
}

// NewCreditEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCreditEntitlementService(opts ...option.RequestOption) (r *CreditEntitlementService) {
	r = &CreditEntitlementService{}
	r.Options = opts
	r.Balances = NewCreditEntitlementBalanceService(opts...)
	return
}

// Credit entitlements define reusable credit templates that can be attached to
// products. Each entitlement defines how credits behave in terms of expiration,
// rollover, and overage.
//
// # Authentication
//
// Requires an API key with `Editor` role.
//
// # Request Body
//
//   - `name` - Human-readable name of the credit entitlement (1-255 characters,
//     required)
//   - `description` - Optional description (max 1000 characters)
//   - `precision` - Decimal precision for credit amounts (0-10 decimal places)
//   - `unit` - Unit of measurement for the credit (e.g., "API Calls", "Tokens",
//     "Credits")
//   - `expires_after_days` - Number of days after which credits expire (optional)
//   - `rollover_enabled` - Whether unused credits can rollover to the next period
//   - `rollover_percentage` - Percentage of unused credits that rollover (0-100)
//   - `rollover_timeframe_count` - Count of timeframe periods for rollover limit
//   - `rollover_timeframe_interval` - Interval type (day, week, month, year)
//   - `max_rollover_count` - Maximum number of times credits can be rolled over
//   - `overage_enabled` - Whether overage charges apply when credits run out
//     (requires price_per_unit)
//   - `overage_limit` - Maximum overage units allowed (optional)
//   - `currency` - Currency for pricing (required if price_per_unit is set)
//   - `price_per_unit` - Price per credit unit (decimal)
//
// # Responses
//
//   - `201 Created` - Credit entitlement created successfully, returns the full
//     entitlement object
//   - `422 Unprocessable Entity` - Invalid request parameters or validation failure
//   - `500 Internal Server Error` - Database or server error
//
// # Business Logic
//
//   - A unique ID with prefix `cde_` is automatically generated for the entitlement
//   - Created and updated timestamps are automatically set
//   - Currency is required when price_per_unit is set
//   - price_per_unit is required when overage_enabled is true
//   - rollover_timeframe_count and rollover_timeframe_interval must both be set or
//     both be null
func (r *CreditEntitlementService) New(ctx context.Context, body CreditEntitlementNewParams, opts ...option.RequestOption) (res *CreditEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "credit-entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the full details of a single credit entitlement including all
// configuration settings for expiration, rollover, and overage policies.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Path Parameters
//
// - `id` - The unique identifier of the credit entitlement (format: `cde_...`)
//
// # Responses
//
//   - `200 OK` - Returns the full credit entitlement object
//   - `404 Not Found` - Credit entitlement does not exist or does not belong to the
//     authenticated business
//   - `500 Internal Server Error` - Database or server error
//
// # Business Logic
//
//   - Only non-deleted credit entitlements can be retrieved through this endpoint
//   - The entitlement must belong to the authenticated business (business_id check)
//   - Deleted entitlements return a 404 error and must be retrieved via the list
//     endpoint with `deleted=true`
func (r *CreditEntitlementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CreditEntitlement, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Allows partial updates to a credit entitlement's configuration. Only the fields
// provided in the request body will be updated; all other fields remain unchanged.
// This endpoint supports nullable fields using the double option pattern.
//
// # Authentication
//
// Requires an API key with `Editor` role.
//
// # Path Parameters
//
//   - `id` - The unique identifier of the credit entitlement to update (format:
//     `cde_...`)
//
// # Request Body (all fields optional)
//
// - `name` - Human-readable name of the credit entitlement (1-255 characters)
// - `description` - Optional description (max 1000 characters)
// - `unit` - Unit of measurement for the credit (1-50 characters)
//
// Note: `precision` cannot be modified after creation as it would invalidate
// existing grants.
//
//   - `expires_after_days` - Number of days after which credits expire (use `null`
//     to remove expiration)
//   - `rollover_enabled` - Whether unused credits can rollover to the next period
//   - `rollover_percentage` - Percentage of unused credits that rollover (0-100,
//     nullable)
//   - `rollover_timeframe_count` - Count of timeframe periods for rollover limit
//     (nullable)
//   - `rollover_timeframe_interval` - Interval type (day, week, month, year,
//     nullable)
//   - `max_rollover_count` - Maximum number of times credits can be rolled over
//     (nullable)
//   - `overage_enabled` - Whether overage charges apply when credits run out
//   - `overage_limit` - Maximum overage units allowed (nullable)
//   - `currency` - Currency for pricing (nullable)
//   - `price_per_unit` - Price per credit unit (decimal, nullable)
//
// # Responses
//
//   - `200 OK` - Credit entitlement updated successfully
//   - `404 Not Found` - Credit entitlement does not exist or does not belong to the
//     authenticated business
//   - `422 Unprocessable Entity` - Invalid request parameters or validation failure
//   - `500 Internal Server Error` - Database or server error
//
// # Business Logic
//
//   - Only non-deleted credit entitlements can be updated
//   - Fields set to `null` explicitly will clear the database value (using double
//     option pattern)
//   - The `updated_at` timestamp is automatically updated on successful modification
//   - Changes take effect immediately but do not retroactively affect existing
//     credit grants
//   - The merged state is validated: currency required with price, rollover
//     timeframe fields together, price required for overage
func (r *CreditEntitlementService) Update(ctx context.Context, id string, body CreditEntitlementUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Returns a paginated list of credit entitlements, allowing filtering of deleted
// entitlements. By default, only non-deleted entitlements are returned.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Query Parameters
//
//   - `page_size` - Number of items per page (default: 10, max: 100)
//   - `page_number` - Zero-based page number (default: 0)
//   - `deleted` - Boolean flag to list deleted entitlements instead of active ones
//     (default: false)
//
// # Responses
//
// - `200 OK` - Returns a list of credit entitlements wrapped in a response object
// - `422 Unprocessable Entity` - Invalid query parameters (e.g., page_size > 100)
// - `500 Internal Server Error` - Database or server error
//
// # Business Logic
//
// - Results are ordered by creation date in descending order (newest first)
// - Only entitlements belonging to the authenticated business are returned
// - The `deleted` parameter controls visibility of soft-deleted entitlements
// - Pagination uses offset-based pagination (offset = page_number \* page_size)
func (r *CreditEntitlementService) List(ctx context.Context, query CreditEntitlementListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[CreditEntitlement], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credit-entitlements"
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

// Returns a paginated list of credit entitlements, allowing filtering of deleted
// entitlements. By default, only non-deleted entitlements are returned.
//
// # Authentication
//
// Requires an API key with `Viewer` role or higher.
//
// # Query Parameters
//
//   - `page_size` - Number of items per page (default: 10, max: 100)
//   - `page_number` - Zero-based page number (default: 0)
//   - `deleted` - Boolean flag to list deleted entitlements instead of active ones
//     (default: false)
//
// # Responses
//
// - `200 OK` - Returns a list of credit entitlements wrapped in a response object
// - `422 Unprocessable Entity` - Invalid query parameters (e.g., page_size > 100)
// - `500 Internal Server Error` - Database or server error
//
// # Business Logic
//
// - Results are ordered by creation date in descending order (newest first)
// - Only entitlements belonging to the authenticated business are returned
// - The `deleted` parameter controls visibility of soft-deleted entitlements
// - Pagination uses offset-based pagination (offset = page_number \* page_size)
func (r *CreditEntitlementService) ListAutoPaging(ctx context.Context, query CreditEntitlementListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[CreditEntitlement] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *CreditEntitlementService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Undeletes a soft-deleted credit entitlement by clearing `deleted_at`, making it
// available again through standard list and get endpoints.
//
// # Authentication
//
// Requires an API key with `Editor` role.
//
// # Path Parameters
//
//   - `id` - The unique identifier of the credit entitlement to restore (format:
//     `cde_...`)
//
// # Responses
//
//   - `200 OK` - Credit entitlement restored successfully
//   - `500 Internal Server Error` - Database error, entitlement not found, or
//     entitlement is not deleted
//
// # Business Logic
//
//   - Only deleted credit entitlements can be restored
//   - The query filters for `deleted_at IS NOT NULL`, so non-deleted entitlements
//     will result in 0 rows affected
//   - If no rows are affected (entitlement doesn't exist, doesn't belong to
//     business, or is not deleted), returns 500
//   - The `updated_at` timestamp is automatically updated on successful restoration
//   - Once restored, the entitlement becomes immediately available in the standard
//     list and get endpoints
//   - All configuration settings are preserved during delete/restore operations
//
// # Error Handling
//
// This endpoint returns 500 Internal Server Error in several cases:
//
// - The credit entitlement does not exist
// - The credit entitlement belongs to a different business
// - The credit entitlement is not currently deleted (already active)
//
// Callers should verify the entitlement exists and is deleted before calling this
// endpoint.
func (r *CreditEntitlementService) Undelete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credit-entitlements/%s/undelete", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Controls how overage is handled at the end of a billing cycle.
//
// | Preset                     | Charge at billing | Credits reduce overage | Preserve overage at reset |
// | -------------------------- | :---------------: | :--------------------: | :-----------------------: |
// | `forgive_at_reset`         |        No         |           No           |            No             |
// | `invoice_at_billing`       |        Yes        |           No           |            No             |
// | `carry_deficit`            |        No         |           No           |            Yes            |
// | `carry_deficit_auto_repay` |        No         |          Yes           |            Yes            |
type CbbOverageBehavior string

const (
	CbbOverageBehaviorForgiveAtReset        CbbOverageBehavior = "forgive_at_reset"
	CbbOverageBehaviorInvoiceAtBilling      CbbOverageBehavior = "invoice_at_billing"
	CbbOverageBehaviorCarryDeficit          CbbOverageBehavior = "carry_deficit"
	CbbOverageBehaviorCarryDeficitAutoRepay CbbOverageBehavior = "carry_deficit_auto_repay"
)

func (r CbbOverageBehavior) IsKnown() bool {
	switch r {
	case CbbOverageBehaviorForgiveAtReset, CbbOverageBehaviorInvoiceAtBilling, CbbOverageBehaviorCarryDeficit, CbbOverageBehaviorCarryDeficitAutoRepay:
		return true
	}
	return false
}

type CreditEntitlement struct {
	ID         string    `json:"id" api:"required"`
	BusinessID string    `json:"business_id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	Name       string    `json:"name" api:"required"`
	// Controls how overage is handled at billing cycle end.
	OverageBehavior  CbbOverageBehavior `json:"overage_behavior" api:"required"`
	OverageEnabled   bool               `json:"overage_enabled" api:"required"`
	Precision        int64              `json:"precision" api:"required"`
	RolloverEnabled  bool               `json:"rollover_enabled" api:"required"`
	Unit             string             `json:"unit" api:"required"`
	UpdatedAt        time.Time          `json:"updated_at" api:"required" format:"date-time"`
	Currency         Currency           `json:"currency" api:"nullable"`
	Description      string             `json:"description" api:"nullable"`
	ExpiresAfterDays int64              `json:"expires_after_days" api:"nullable"`
	MaxRolloverCount int64              `json:"max_rollover_count" api:"nullable"`
	OverageLimit     int64              `json:"overage_limit" api:"nullable"`
	// Price per credit unit
	PricePerUnit              string                `json:"price_per_unit" api:"nullable"`
	RolloverPercentage        int64                 `json:"rollover_percentage" api:"nullable"`
	RolloverTimeframeCount    int64                 `json:"rollover_timeframe_count" api:"nullable"`
	RolloverTimeframeInterval TimeInterval          `json:"rollover_timeframe_interval" api:"nullable"`
	JSON                      creditEntitlementJSON `json:"-"`
}

// creditEntitlementJSON contains the JSON metadata for the struct
// [CreditEntitlement]
type creditEntitlementJSON struct {
	ID                        apijson.Field
	BusinessID                apijson.Field
	CreatedAt                 apijson.Field
	Name                      apijson.Field
	OverageBehavior           apijson.Field
	OverageEnabled            apijson.Field
	Precision                 apijson.Field
	RolloverEnabled           apijson.Field
	Unit                      apijson.Field
	UpdatedAt                 apijson.Field
	Currency                  apijson.Field
	Description               apijson.Field
	ExpiresAfterDays          apijson.Field
	MaxRolloverCount          apijson.Field
	OverageLimit              apijson.Field
	PricePerUnit              apijson.Field
	RolloverPercentage        apijson.Field
	RolloverTimeframeCount    apijson.Field
	RolloverTimeframeInterval apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CreditEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditEntitlementJSON) RawJSON() string {
	return r.raw
}

type CreditEntitlementNewParams struct {
	// Name of the credit entitlement
	Name param.Field[string] `json:"name" api:"required"`
	// Whether overage charges are enabled when credits run out
	OverageEnabled param.Field[bool] `json:"overage_enabled" api:"required"`
	// Precision for credit amounts (0-10 decimal places)
	Precision param.Field[int64] `json:"precision" api:"required"`
	// Whether rollover is enabled for unused credits
	RolloverEnabled param.Field[bool] `json:"rollover_enabled" api:"required"`
	// Unit of measurement for the credit (e.g., "API Calls", "Tokens", "Credits")
	Unit param.Field[string] `json:"unit" api:"required"`
	// Currency for pricing (required if price_per_unit is set)
	Currency param.Field[Currency] `json:"currency"`
	// Optional description of the credit entitlement
	Description param.Field[string] `json:"description"`
	// Number of days after which credits expire (optional)
	ExpiresAfterDays param.Field[int64] `json:"expires_after_days"`
	// Maximum number of times credits can be rolled over
	MaxRolloverCount param.Field[int64] `json:"max_rollover_count"`
	// Controls how overage is handled at billing cycle end. Defaults to
	// forgive_at_reset if not specified.
	OverageBehavior param.Field[CbbOverageBehavior] `json:"overage_behavior"`
	// Maximum overage units allowed (optional)
	OverageLimit param.Field[int64] `json:"overage_limit"`
	// Price per credit unit
	PricePerUnit param.Field[string] `json:"price_per_unit"`
	// Percentage of unused credits that can rollover (0-100)
	RolloverPercentage param.Field[int64] `json:"rollover_percentage"`
	// Count of timeframe periods for rollover limit
	RolloverTimeframeCount param.Field[int64] `json:"rollover_timeframe_count"`
	// Interval type for rollover timeframe
	RolloverTimeframeInterval param.Field[TimeInterval] `json:"rollover_timeframe_interval"`
}

func (r CreditEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditEntitlementUpdateParams struct {
	// Currency for pricing
	Currency param.Field[Currency] `json:"currency"`
	// Optional description of the credit entitlement
	Description param.Field[string] `json:"description"`
	// Number of days after which credits expire
	ExpiresAfterDays param.Field[int64] `json:"expires_after_days"`
	// Maximum number of times credits can be rolled over
	MaxRolloverCount param.Field[int64] `json:"max_rollover_count"`
	// Name of the credit entitlement
	Name param.Field[string] `json:"name"`
	// Controls how overage is handled at billing cycle end.
	OverageBehavior param.Field[CbbOverageBehavior] `json:"overage_behavior"`
	// Whether overage charges are enabled when credits run out
	OverageEnabled param.Field[bool] `json:"overage_enabled"`
	// Maximum overage units allowed
	OverageLimit param.Field[int64] `json:"overage_limit"`
	// Price per credit unit
	PricePerUnit param.Field[string] `json:"price_per_unit"`
	// Whether rollover is enabled for unused credits
	RolloverEnabled param.Field[bool] `json:"rollover_enabled"`
	// Percentage of unused credits that can rollover (0-100)
	RolloverPercentage param.Field[int64] `json:"rollover_percentage"`
	// Count of timeframe periods for rollover limit
	RolloverTimeframeCount param.Field[int64] `json:"rollover_timeframe_count"`
	// Interval type for rollover timeframe
	RolloverTimeframeInterval param.Field[TimeInterval] `json:"rollover_timeframe_interval"`
	// Unit of measurement for the credit (e.g., "API Calls", "Tokens", "Credits")
	Unit param.Field[string] `json:"unit"`
}

func (r CreditEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditEntitlementListParams struct {
	// List deleted credit entitlements
	Deleted param.Field[bool] `query:"deleted"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [CreditEntitlementListParams]'s query parameters as
// `url.Values`.
func (r CreditEntitlementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
