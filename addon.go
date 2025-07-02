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

// AddonService contains methods and other services that help with interacting with
// the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddonService] method instead.
type AddonService struct {
	Options []option.RequestOption
}

// NewAddonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddonService(opts ...option.RequestOption) (r *AddonService) {
	r = &AddonService{}
	r.Options = opts
	return
}

func (r *AddonService) New(ctx context.Context, body AddonNewParams, opts ...option.RequestOption) (res *AddonResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "addons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *AddonService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AddonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("addons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *AddonService) Update(ctx context.Context, id string, body AddonUpdateParams, opts ...option.RequestOption) (res *AddonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("addons/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *AddonService) List(ctx context.Context, query AddonListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[AddonResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "addons"
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

func (r *AddonService) ListAutoPaging(ctx context.Context, query AddonListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[AddonResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *AddonService) UpdateImages(ctx context.Context, id string, opts ...option.RequestOption) (res *AddonUpdateImagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("addons/%s/images", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type AddonResponse struct {
	// id of the Addon
	ID string `json:"id,required"`
	// Unique identifier for the business to which the addon belongs.
	BusinessID string `json:"business_id,required"`
	// Created time
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Currency of the Addon
	Currency Currency `json:"currency,required"`
	// Name of the Addon
	Name string `json:"name,required"`
	// Amount of the addon
	Price int64 `json:"price,required"`
	// Tax category applied to this Addon
	TaxCategory TaxCategory `json:"tax_category,required"`
	// Updated time
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Optional description of the Addon
	Description string `json:"description,nullable"`
	// Image of the Addon
	Image string            `json:"image,nullable"`
	JSON  addonResponseJSON `json:"-"`
}

// addonResponseJSON contains the JSON metadata for the struct [AddonResponse]
type addonResponseJSON struct {
	ID          apijson.Field
	BusinessID  apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	TaxCategory apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addonResponseJSON) RawJSON() string {
	return r.raw
}

type AddonUpdateImagesResponse struct {
	ImageID string                        `json:"image_id,required" format:"uuid"`
	URL     string                        `json:"url,required"`
	JSON    addonUpdateImagesResponseJSON `json:"-"`
}

// addonUpdateImagesResponseJSON contains the JSON metadata for the struct
// [AddonUpdateImagesResponse]
type addonUpdateImagesResponseJSON struct {
	ImageID     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddonUpdateImagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addonUpdateImagesResponseJSON) RawJSON() string {
	return r.raw
}

type AddonNewParams struct {
	// The currency of the Addon
	Currency param.Field[Currency] `json:"currency,required"`
	// Name of the Addon
	Name param.Field[string] `json:"name,required"`
	// Amount of the addon
	Price param.Field[int64] `json:"price,required"`
	// Tax category applied to this Addon
	TaxCategory param.Field[TaxCategory] `json:"tax_category,required"`
	// Optional description of the Addon
	Description param.Field[string] `json:"description"`
}

func (r AddonNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddonUpdateParams struct {
	// The currency of the Addon
	Currency param.Field[Currency] `json:"currency"`
	// Description of the Addon, optional and must be at most 1000 characters.
	Description param.Field[string] `json:"description"`
	// Addon image id after its uploaded to S3
	ImageID param.Field[string] `json:"image_id" format:"uuid"`
	// Name of the Addon, optional and must be at most 100 characters.
	Name param.Field[string] `json:"name"`
	// Amount of the addon
	Price param.Field[int64] `json:"price"`
	// Tax category of the Addon.
	TaxCategory param.Field[TaxCategory] `json:"tax_category"`
}

func (r AddonUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddonListParams struct {
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [AddonListParams]'s query parameters as `url.Values`.
func (r AddonListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
