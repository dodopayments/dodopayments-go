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
)

// LicenseKeyService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLicenseKeyService] method instead.
type LicenseKeyService struct {
	Options []option.RequestOption
}

// NewLicenseKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLicenseKeyService(opts ...option.RequestOption) (r *LicenseKeyService) {
	r = &LicenseKeyService{}
	r.Options = opts
	return
}

func (r *LicenseKeyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LicenseKey, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("license_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *LicenseKeyService) Update(ctx context.Context, id string, body LicenseKeyUpdateParams, opts ...option.RequestOption) (res *LicenseKey, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("license_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *LicenseKeyService) List(ctx context.Context, query LicenseKeyListParams, opts ...option.RequestOption) (res *[]LicenseKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "license_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LicenseKey struct {
	// The unique identifier of the license key.
	ID string `json:"id,required"`
	// The unique identifier of the business associated with the license key.
	BusinessID string `json:"business_id,required"`
	// The timestamp indicating when the license key was created, in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The unique identifier of the customer associated with the license key.
	CustomerID string `json:"customer_id,required"`
	// The current number of instances activated for this license key.
	InstancesCount int64 `json:"instances_count,required"`
	// The license key string.
	Key string `json:"key,required"`
	// The unique identifier of the payment associated with the license key.
	PaymentID string `json:"payment_id,required"`
	// The unique identifier of the product associated with the license key.
	ProductID string           `json:"product_id,required"`
	Status    LicenseKeyStatus `json:"status,required"`
	// The maximum number of activations allowed for this license key.
	ActivationsLimit int64 `json:"activations_limit,nullable"`
	// The timestamp indicating when the license key expires, in UTC.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// The unique identifier of the subscription associated with the license key, if
	// any.
	SubscriptionID string         `json:"subscription_id,nullable"`
	JSON           licenseKeyJSON `json:"-"`
}

// licenseKeyJSON contains the JSON metadata for the struct [LicenseKey]
type licenseKeyJSON struct {
	ID               apijson.Field
	BusinessID       apijson.Field
	CreatedAt        apijson.Field
	CustomerID       apijson.Field
	InstancesCount   apijson.Field
	Key              apijson.Field
	PaymentID        apijson.Field
	ProductID        apijson.Field
	Status           apijson.Field
	ActivationsLimit apijson.Field
	ExpiresAt        apijson.Field
	SubscriptionID   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LicenseKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseKeyJSON) RawJSON() string {
	return r.raw
}

type LicenseKeyStatus string

const (
	LicenseKeyStatusActive   LicenseKeyStatus = "active"
	LicenseKeyStatusExpired  LicenseKeyStatus = "expired"
	LicenseKeyStatusDisabled LicenseKeyStatus = "disabled"
)

func (r LicenseKeyStatus) IsKnown() bool {
	switch r {
	case LicenseKeyStatusActive, LicenseKeyStatusExpired, LicenseKeyStatusDisabled:
		return true
	}
	return false
}

type LicenseKeyListResponse struct {
	Items []LicenseKey               `json:"items,required"`
	JSON  licenseKeyListResponseJSON `json:"-"`
}

// licenseKeyListResponseJSON contains the JSON metadata for the struct
// [LicenseKeyListResponse]
type licenseKeyListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseKeyUpdateParams struct {
	// The updated activation limit for the license key. Use `null` to remove the
	// limit, or omit this field to leave it unchanged.
	ActivationsLimit param.Field[int64] `json:"activations_limit"`
	// Indicates whether the license key should be disabled. A value of `true` disables
	// the key, while `false` enables it. Omit this field to leave it unchanged.
	Disabled param.Field[bool] `json:"disabled"`
	// The updated expiration timestamp for the license key in UTC. Use `null` to
	// remove the expiration date, or omit this field to leave it unchanged.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
}

func (r LicenseKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LicenseKeyListParams struct {
	// Filter by customer ID
	CustomerID param.Field[string] `query:"customer_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by product ID
	ProductID param.Field[string] `query:"product_id"`
	// Filter by license key status
	Status param.Field[LicenseKeyListParamsStatus] `query:"status"`
}

// URLQuery serializes [LicenseKeyListParams]'s query parameters as `url.Values`.
func (r LicenseKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by license key status
type LicenseKeyListParamsStatus string

const (
	LicenseKeyListParamsStatusActive   LicenseKeyListParamsStatus = "active"
	LicenseKeyListParamsStatusExpired  LicenseKeyListParamsStatus = "expired"
	LicenseKeyListParamsStatusDisabled LicenseKeyListParamsStatus = "disabled"
)

func (r LicenseKeyListParamsStatus) IsKnown() bool {
	switch r {
	case LicenseKeyListParamsStatusActive, LicenseKeyListParamsStatusExpired, LicenseKeyListParamsStatusDisabled:
		return true
	}
	return false
}
