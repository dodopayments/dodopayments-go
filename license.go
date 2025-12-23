// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// LicenseService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLicenseService] method instead.
type LicenseService struct {
	Options []option.RequestOption
}

// NewLicenseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLicenseService(opts ...option.RequestOption) (r LicenseService) {
	r = LicenseService{}
	r.Options = opts
	return
}

func (r *LicenseService) Activate(ctx context.Context, body LicenseActivateParams, opts ...option.RequestOption) (res *LicenseActivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "licenses/activate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *LicenseService) Deactivate(ctx context.Context, body LicenseDeactivateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "licenses/deactivate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

func (r *LicenseService) Validate(ctx context.Context, body LicenseValidateParams, opts ...option.RequestOption) (res *LicenseValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "licenses/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LicenseActivateResponse struct {
	// License key instance ID
	ID string `json:"id,required"`
	// Business ID
	BusinessID string `json:"business_id,required"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Limited customer details associated with the license key.
	Customer CustomerLimitedDetails `json:"customer,required"`
	// Associated license key ID
	LicenseKeyID string `json:"license_key_id,required"`
	// Instance name
	Name string `json:"name,required"`
	// Related product info. Present if the license key is tied to a product.
	Product LicenseActivateResponseProduct `json:"product,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		BusinessID   respjson.Field
		CreatedAt    respjson.Field
		Customer     respjson.Field
		LicenseKeyID respjson.Field
		Name         respjson.Field
		Product      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LicenseActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *LicenseActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Related product info. Present if the license key is tied to a product.
type LicenseActivateResponseProduct struct {
	// Unique identifier for the product.
	ProductID string `json:"product_id,required"`
	// Name of the product, if set by the merchant.
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductID   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LicenseActivateResponseProduct) RawJSON() string { return r.JSON.raw }
func (r *LicenseActivateResponseProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LicenseValidateResponse struct {
	Valid bool `json:"valid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Valid       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LicenseValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *LicenseValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LicenseActivateParams struct {
	LicenseKey string `json:"license_key,required"`
	Name       string `json:"name,required"`
	paramObj
}

func (r LicenseActivateParams) MarshalJSON() (data []byte, err error) {
	type shadow LicenseActivateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LicenseActivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LicenseDeactivateParams struct {
	LicenseKey           string `json:"license_key,required"`
	LicenseKeyInstanceID string `json:"license_key_instance_id,required"`
	paramObj
}

func (r LicenseDeactivateParams) MarshalJSON() (data []byte, err error) {
	type shadow LicenseDeactivateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LicenseDeactivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LicenseValidateParams struct {
	LicenseKey           string            `json:"license_key,required"`
	LicenseKeyInstanceID param.Opt[string] `json:"license_key_instance_id,omitzero"`
	paramObj
}

func (r LicenseValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow LicenseValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LicenseValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
