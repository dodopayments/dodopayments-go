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

// EntitlementGrantService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitlementGrantService] method instead.
type EntitlementGrantService struct {
	Options []option.RequestOption
}

// NewEntitlementGrantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntitlementGrantService(opts ...option.RequestOption) (r *EntitlementGrantService) {
	r = &EntitlementGrantService{}
	r.Options = opts
	return
}

// GET /entitlements/{id}/grants (public API)
func (r *EntitlementGrantService) List(ctx context.Context, id string, query EntitlementGrantListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[EntitlementGrant], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entitlements/%s/grants", id)
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

// GET /entitlements/{id}/grants (public API)
func (r *EntitlementGrantService) ListAutoPaging(ctx context.Context, id string, query EntitlementGrantListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[EntitlementGrant] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, id, query, opts...))
}

// Revoke a single grant. Idempotent: re-revoking an already-revoked grant returns
// the grant in its current state.
func (r *EntitlementGrantService) Revoke(ctx context.Context, id string, grantID string, opts ...option.RequestOption) (res *EntitlementGrant, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if grantID == "" {
		err = errors.New("missing required grant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entitlements/%s/grants/%s", id, grantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Detailed view of a single entitlement grant: who it's for, its lifecycle state,
// and any integration-specific delivery payload.
type EntitlementGrant struct {
	// Unique identifier of the grant.
	ID string `json:"id" api:"required"`
	// Identifier of the business that owns the grant.
	BusinessID string `json:"business_id" api:"required"`
	// Timestamp when the grant was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Identifier of the customer the grant was issued to.
	CustomerID string `json:"customer_id" api:"required"`
	// Identifier of the entitlement this grant was issued from.
	EntitlementID string `json:"entitlement_id" api:"required"`
	// Arbitrary key-value metadata recorded on the grant.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Lifecycle status of the grant.
	Status EntitlementGrantStatus `json:"status" api:"required"`
	// Timestamp when the grant was last modified.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Timestamp when the grant transitioned to `delivered`, when applicable.
	DeliveredAt time.Time `json:"delivered_at" api:"nullable" format:"date-time"`
	// Digital-product-delivery payload, present when the entitlement integration is
	// `digital_files`.
	DigitalProductDelivery DigitalProductDelivery `json:"digital_product_delivery" api:"nullable"`
	// Machine-readable code reported when delivery failed, when applicable.
	ErrorCode string `json:"error_code" api:"nullable"`
	// Human-readable message reported when delivery failed, when applicable.
	ErrorMessage string `json:"error_message" api:"nullable"`
	// License-key delivery payload, present when the entitlement integration is
	// `license_key`.
	LicenseKey LicenseKeyGrant `json:"license_key" api:"nullable"`
	// Timestamp when `oauth_url` stops being valid, when applicable.
	OAuthExpiresAt time.Time `json:"oauth_expires_at" api:"nullable" format:"date-time"`
	// Customer-facing OAuth URL for OAuth-style integrations. Populated during the
	// customer-portal accept flow; `null` until the customer completes that step, and
	// on grants for non-OAuth integrations.
	OAuthURL string `json:"oauth_url" api:"nullable"`
	// Identifier of the payment that triggered this grant, when applicable.
	PaymentID string `json:"payment_id" api:"nullable"`
	// Reason recorded when the grant was revoked, when applicable.
	RevocationReason string `json:"revocation_reason" api:"nullable"`
	// Timestamp when the grant transitioned to `revoked`, when applicable.
	RevokedAt time.Time `json:"revoked_at" api:"nullable" format:"date-time"`
	// Identifier of the subscription that triggered this grant, when applicable.
	SubscriptionID string               `json:"subscription_id" api:"nullable"`
	JSON           entitlementGrantJSON `json:"-"`
}

// entitlementGrantJSON contains the JSON metadata for the struct
// [EntitlementGrant]
type entitlementGrantJSON struct {
	ID                     apijson.Field
	BusinessID             apijson.Field
	CreatedAt              apijson.Field
	CustomerID             apijson.Field
	EntitlementID          apijson.Field
	Metadata               apijson.Field
	Status                 apijson.Field
	UpdatedAt              apijson.Field
	DeliveredAt            apijson.Field
	DigitalProductDelivery apijson.Field
	ErrorCode              apijson.Field
	ErrorMessage           apijson.Field
	LicenseKey             apijson.Field
	OAuthExpiresAt         apijson.Field
	OAuthURL               apijson.Field
	PaymentID              apijson.Field
	RevocationReason       apijson.Field
	RevokedAt              apijson.Field
	SubscriptionID         apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *EntitlementGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGrantJSON) RawJSON() string {
	return r.raw
}

// Lifecycle status of the grant.
type EntitlementGrantStatus string

const (
	EntitlementGrantStatusPending   EntitlementGrantStatus = "Pending"
	EntitlementGrantStatusDelivered EntitlementGrantStatus = "Delivered"
	EntitlementGrantStatusFailed    EntitlementGrantStatus = "Failed"
	EntitlementGrantStatusRevoked   EntitlementGrantStatus = "Revoked"
)

func (r EntitlementGrantStatus) IsKnown() bool {
	switch r {
	case EntitlementGrantStatusPending, EntitlementGrantStatusDelivered, EntitlementGrantStatusFailed, EntitlementGrantStatusRevoked:
		return true
	}
	return false
}

// License-key delivery payload, present on grants for `license_key` entitlements.
// The grant's top-level `status` is the source of truth for the grant's lifecycle.
type LicenseKeyGrant struct {
	// Number of activations consumed so far.
	ActivationsUsed int64 `json:"activations_used" api:"required"`
	// Issued license key.
	Key string `json:"key" api:"required"`
	// Maximum activations allowed by the entitlement, when set.
	ActivationsLimit int64 `json:"activations_limit" api:"nullable"`
	// When the license key expires, when applicable.
	ExpiresAt time.Time           `json:"expires_at" api:"nullable" format:"date-time"`
	JSON      licenseKeyGrantJSON `json:"-"`
}

// licenseKeyGrantJSON contains the JSON metadata for the struct [LicenseKeyGrant]
type licenseKeyGrantJSON struct {
	ActivationsUsed  apijson.Field
	Key              apijson.Field
	ActivationsLimit apijson.Field
	ExpiresAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LicenseKeyGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseKeyGrantJSON) RawJSON() string {
	return r.raw
}

type EntitlementGrantListParams struct {
	// Filter by customer ID
	CustomerID param.Field[string] `query:"customer_id"`
	// Page number (default 0)
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size (default 10, max 100)
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by grant status
	Status param.Field[EntitlementGrantListParamsStatus] `query:"status"`
}

// URLQuery serializes [EntitlementGrantListParams]'s query parameters as
// `url.Values`.
func (r EntitlementGrantListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by grant status
type EntitlementGrantListParamsStatus string

const (
	EntitlementGrantListParamsStatusPending   EntitlementGrantListParamsStatus = "Pending"
	EntitlementGrantListParamsStatusDelivered EntitlementGrantListParamsStatus = "Delivered"
	EntitlementGrantListParamsStatusFailed    EntitlementGrantListParamsStatus = "Failed"
	EntitlementGrantListParamsStatusRevoked   EntitlementGrantListParamsStatus = "Revoked"
)

func (r EntitlementGrantListParamsStatus) IsKnown() bool {
	switch r {
	case EntitlementGrantListParamsStatusPending, EntitlementGrantListParamsStatusDelivered, EntitlementGrantListParamsStatusFailed, EntitlementGrantListParamsStatusRevoked:
		return true
	}
	return false
}
