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

// LicenseKeyInstanceService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLicenseKeyInstanceService] method instead.
type LicenseKeyInstanceService struct {
	Options []option.RequestOption
}

// NewLicenseKeyInstanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLicenseKeyInstanceService(opts ...option.RequestOption) (r *LicenseKeyInstanceService) {
	r = &LicenseKeyInstanceService{}
	r.Options = opts
	return
}

func (r *LicenseKeyInstanceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LicenseKeyInstance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("license_key_instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *LicenseKeyInstanceService) Update(ctx context.Context, id string, body LicenseKeyInstanceUpdateParams, opts ...option.RequestOption) (res *LicenseKeyInstance, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("license_key_instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *LicenseKeyInstanceService) List(ctx context.Context, query LicenseKeyInstanceListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[LicenseKeyInstance], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "license_key_instances"
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

func (r *LicenseKeyInstanceService) ListAutoPaging(ctx context.Context, query LicenseKeyInstanceListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[LicenseKeyInstance] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type LicenseKeyInstance struct {
	ID           string                 `json:"id,required"`
	BusinessID   string                 `json:"business_id,required"`
	CreatedAt    time.Time              `json:"created_at,required" format:"date-time"`
	LicenseKeyID string                 `json:"license_key_id,required"`
	Name         string                 `json:"name,required"`
	JSON         licenseKeyInstanceJSON `json:"-"`
}

// licenseKeyInstanceJSON contains the JSON metadata for the struct
// [LicenseKeyInstance]
type licenseKeyInstanceJSON struct {
	ID           apijson.Field
	BusinessID   apijson.Field
	CreatedAt    apijson.Field
	LicenseKeyID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LicenseKeyInstance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseKeyInstanceJSON) RawJSON() string {
	return r.raw
}

type LicenseKeyInstanceUpdateParams struct {
	Name param.Field[string] `json:"name,required"`
}

func (r LicenseKeyInstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LicenseKeyInstanceListParams struct {
	// Filter by license key ID
	LicenseKeyID param.Field[string] `query:"license_key_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [LicenseKeyInstanceListParams]'s query parameters as
// `url.Values`.
func (r LicenseKeyInstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
