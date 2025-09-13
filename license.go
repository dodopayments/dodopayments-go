// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
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
func NewLicenseService(opts ...option.RequestOption) (r *LicenseService) {
	r = &LicenseService{}
	r.Options = opts
	return
}

func (r *LicenseService) Activate(ctx context.Context, body LicenseActivateParams, opts ...option.RequestOption) (res *LicenseActivateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "licenses/activate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *LicenseService) Deactivate(ctx context.Context, body LicenseDeactivateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "licenses/deactivate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

func (r *LicenseService) Validate(ctx context.Context, body LicenseValidateParams, opts ...option.RequestOption) (res *LicenseValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	JSON    licenseActivateResponseJSON    `json:"-"`
}

// licenseActivateResponseJSON contains the JSON metadata for the struct
// [LicenseActivateResponse]
type licenseActivateResponseJSON struct {
	ID           apijson.Field
	BusinessID   apijson.Field
	CreatedAt    apijson.Field
	Customer     apijson.Field
	LicenseKeyID apijson.Field
	Name         apijson.Field
	Product      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LicenseActivateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseActivateResponseJSON) RawJSON() string {
	return r.raw
}

// Related product info. Present if the license key is tied to a product.
type LicenseActivateResponseProduct struct {
	// Unique identifier for the product.
	ProductID string `json:"product_id,required"`
	// Name of the product, if set by the merchant.
	Name string                             `json:"name,nullable"`
	JSON licenseActivateResponseProductJSON `json:"-"`
}

// licenseActivateResponseProductJSON contains the JSON metadata for the struct
// [LicenseActivateResponseProduct]
type licenseActivateResponseProductJSON struct {
	ProductID   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseActivateResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseActivateResponseProductJSON) RawJSON() string {
	return r.raw
}

type LicenseValidateResponse struct {
	Valid bool                        `json:"valid,required"`
	JSON  licenseValidateResponseJSON `json:"-"`
}

// licenseValidateResponseJSON contains the JSON metadata for the struct
// [LicenseValidateResponse]
type licenseValidateResponseJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseValidateResponseJSON) RawJSON() string {
	return r.raw
}

type LicenseActivateParams struct {
	LicenseKey param.Field[string] `json:"license_key,required"`
	Name       param.Field[string] `json:"name,required"`
}

func (r LicenseActivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LicenseDeactivateParams struct {
	LicenseKey           param.Field[string] `json:"license_key,required"`
	LicenseKeyInstanceID param.Field[string] `json:"license_key_instance_id,required"`
}

func (r LicenseDeactivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LicenseValidateParams struct {
	LicenseKey           param.Field[string] `json:"license_key,required"`
	LicenseKeyInstanceID param.Field[string] `json:"license_key_instance_id"`
}

func (r LicenseValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
