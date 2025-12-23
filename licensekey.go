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
func NewLicenseKeyService(opts ...option.RequestOption) (r LicenseKeyService) {
	r = LicenseKeyService{}
	r.Options = opts
	return
}

func (r *LicenseKeyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LicenseKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("license_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *LicenseKeyService) Update(ctx context.Context, id string, body LicenseKeyUpdateParams, opts ...option.RequestOption) (res *LicenseKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("license_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *LicenseKeyService) List(ctx context.Context, query LicenseKeyListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[LicenseKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "license_keys"
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

func (r *LicenseKeyService) ListAutoPaging(ctx context.Context, query LicenseKeyListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[LicenseKey] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
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
	ProductID string `json:"product_id,required"`
	// The current status of the license key (e.g., active, inactive, expired).
	//
	// Any of "active", "expired", "disabled".
	Status LicenseKeyStatus `json:"status,required"`
	// The maximum number of activations allowed for this license key.
	ActivationsLimit int64 `json:"activations_limit,nullable"`
	// The timestamp indicating when the license key expires, in UTC.
	ExpiresAt time.Time `json:"expires_at,nullable" format:"date-time"`
	// The unique identifier of the subscription associated with the license key, if
	// any.
	SubscriptionID string `json:"subscription_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		BusinessID       respjson.Field
		CreatedAt        respjson.Field
		CustomerID       respjson.Field
		InstancesCount   respjson.Field
		Key              respjson.Field
		PaymentID        respjson.Field
		ProductID        respjson.Field
		Status           respjson.Field
		ActivationsLimit respjson.Field
		ExpiresAt        respjson.Field
		SubscriptionID   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LicenseKey) RawJSON() string { return r.JSON.raw }
func (r *LicenseKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LicenseKeyStatus string

const (
	LicenseKeyStatusActive   LicenseKeyStatus = "active"
	LicenseKeyStatusExpired  LicenseKeyStatus = "expired"
	LicenseKeyStatusDisabled LicenseKeyStatus = "disabled"
)

type LicenseKeyUpdateParams struct {
	// The updated activation limit for the license key. Use `null` to remove the
	// limit, or omit this field to leave it unchanged.
	ActivationsLimit param.Opt[int64] `json:"activations_limit,omitzero"`
	// Indicates whether the license key should be disabled. A value of `true` disables
	// the key, while `false` enables it. Omit this field to leave it unchanged.
	Disabled param.Opt[bool] `json:"disabled,omitzero"`
	// The updated expiration timestamp for the license key in UTC. Use `null` to
	// remove the expiration date, or omit this field to leave it unchanged.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	paramObj
}

func (r LicenseKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LicenseKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LicenseKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LicenseKeyListParams struct {
	// Filter by customer ID
	CustomerID param.Opt[string] `query:"customer_id,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter by product ID
	ProductID param.Opt[string] `query:"product_id,omitzero" json:"-"`
	// Filter by license key status
	//
	// Any of "active", "expired", "disabled".
	Status LicenseKeyListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LicenseKeyListParams]'s query parameters as `url.Values`.
func (r LicenseKeyListParams) URLQuery() (v url.Values, err error) {
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
