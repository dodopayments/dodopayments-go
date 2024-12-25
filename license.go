// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"

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

func (r *LicenseService) Activate(ctx context.Context, body LicenseActivateParams, opts ...option.RequestOption) (res *LicenseKeyInstance, err error) {
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
