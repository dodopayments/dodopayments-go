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

// ProductCollectionService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductCollectionService] method instead.
type ProductCollectionService struct {
	Options []option.RequestOption
	Groups  *ProductCollectionGroupService
}

// NewProductCollectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductCollectionService(opts ...option.RequestOption) (r *ProductCollectionService) {
	r = &ProductCollectionService{}
	r.Options = opts
	r.Groups = NewProductCollectionGroupService(opts...)
	return
}

func (r *ProductCollectionService) New(ctx context.Context, body ProductCollectionNewParams, opts ...option.RequestOption) (res *ProductCollectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "product-collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ProductCollectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductCollectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("product-collections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ProductCollectionService) Update(ctx context.Context, id string, body ProductCollectionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("product-collections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

func (r *ProductCollectionService) List(ctx context.Context, query ProductCollectionListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[ProductCollectionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "product-collections"
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

func (r *ProductCollectionService) ListAutoPaging(ctx context.Context, query ProductCollectionListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[ProductCollectionListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *ProductCollectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("product-collections/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func (r *ProductCollectionService) Unarchive(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductCollectionUnarchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("product-collections/%s/unarchive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

func (r *ProductCollectionService) UpdateImages(ctx context.Context, id string, body ProductCollectionUpdateImagesParams, opts ...option.RequestOption) (res *ProductCollectionUpdateImagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("product-collections/%s/images", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type ProductCollectionNewResponse struct {
	// Unique identifier for the product collection
	ID string `json:"id" api:"required"`
	// Brand ID for the collection
	BrandID string `json:"brand_id" api:"required"`
	// Timestamp when the collection was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Groups in this collection
	Groups []ProductCollectionNewResponseGroup `json:"groups" api:"required"`
	// Name of the collection
	Name string `json:"name" api:"required"`
	// Timestamp when the collection was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Description of the collection
	Description string `json:"description" api:"nullable"`
	// URL of the collection image
	Image string                           `json:"image" api:"nullable"`
	JSON  productCollectionNewResponseJSON `json:"-"`
}

// productCollectionNewResponseJSON contains the JSON metadata for the struct
// [ProductCollectionNewResponse]
type productCollectionNewResponseJSON struct {
	ID          apijson.Field
	BrandID     apijson.Field
	CreatedAt   apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductCollectionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionNewResponseGroup struct {
	GroupID   string                                      `json:"group_id" api:"required"`
	Products  []ProductCollectionNewResponseGroupsProduct `json:"products" api:"required"`
	Status    bool                                        `json:"status" api:"required"`
	GroupName string                                      `json:"group_name" api:"nullable"`
	JSON      productCollectionNewResponseGroupJSON       `json:"-"`
}

// productCollectionNewResponseGroupJSON contains the JSON metadata for the struct
// [ProductCollectionNewResponseGroup]
type productCollectionNewResponseGroupJSON struct {
	GroupID     apijson.Field
	Products    apijson.Field
	Status      apijson.Field
	GroupName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductCollectionNewResponseGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionNewResponseGroupJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionNewResponseGroupsProduct struct {
	ID          string `json:"id" api:"required"`
	AddonsCount int64  `json:"addons_count" api:"required"`
	FilesCount  int64  `json:"files_count" api:"required"`
	// Whether this product has any credit entitlements attached
	HasCreditEntitlements bool     `json:"has_credit_entitlements" api:"required"`
	IsRecurring           bool     `json:"is_recurring" api:"required"`
	LicenseKeyEnabled     bool     `json:"license_key_enabled" api:"required"`
	MetersCount           int64    `json:"meters_count" api:"required"`
	ProductID             string   `json:"product_id" api:"required"`
	Status                bool     `json:"status" api:"required"`
	Currency              Currency `json:"currency" api:"nullable"`
	Description           string   `json:"description" api:"nullable"`
	Name                  string   `json:"name" api:"nullable"`
	Price                 int64    `json:"price" api:"nullable"`
	// One-time price details.
	PriceDetail Price `json:"price_detail" api:"nullable"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory  TaxCategory                                   `json:"tax_category" api:"nullable"`
	TaxInclusive bool                                          `json:"tax_inclusive" api:"nullable"`
	JSON         productCollectionNewResponseGroupsProductJSON `json:"-"`
}

// productCollectionNewResponseGroupsProductJSON contains the JSON metadata for the
// struct [ProductCollectionNewResponseGroupsProduct]
type productCollectionNewResponseGroupsProductJSON struct {
	ID                    apijson.Field
	AddonsCount           apijson.Field
	FilesCount            apijson.Field
	HasCreditEntitlements apijson.Field
	IsRecurring           apijson.Field
	LicenseKeyEnabled     apijson.Field
	MetersCount           apijson.Field
	ProductID             apijson.Field
	Status                apijson.Field
	Currency              apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	Price                 apijson.Field
	PriceDetail           apijson.Field
	TaxCategory           apijson.Field
	TaxInclusive          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProductCollectionNewResponseGroupsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionNewResponseGroupsProductJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionGetResponse struct {
	// Unique identifier for the product collection
	ID string `json:"id" api:"required"`
	// Brand ID for the collection
	BrandID string `json:"brand_id" api:"required"`
	// Timestamp when the collection was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Groups in this collection
	Groups []ProductCollectionGetResponseGroup `json:"groups" api:"required"`
	// Name of the collection
	Name string `json:"name" api:"required"`
	// Timestamp when the collection was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Description of the collection
	Description string `json:"description" api:"nullable"`
	// URL of the collection image
	Image string                           `json:"image" api:"nullable"`
	JSON  productCollectionGetResponseJSON `json:"-"`
}

// productCollectionGetResponseJSON contains the JSON metadata for the struct
// [ProductCollectionGetResponse]
type productCollectionGetResponseJSON struct {
	ID          apijson.Field
	BrandID     apijson.Field
	CreatedAt   apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductCollectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionGetResponseGroup struct {
	GroupID   string                                      `json:"group_id" api:"required"`
	Products  []ProductCollectionGetResponseGroupsProduct `json:"products" api:"required"`
	Status    bool                                        `json:"status" api:"required"`
	GroupName string                                      `json:"group_name" api:"nullable"`
	JSON      productCollectionGetResponseGroupJSON       `json:"-"`
}

// productCollectionGetResponseGroupJSON contains the JSON metadata for the struct
// [ProductCollectionGetResponseGroup]
type productCollectionGetResponseGroupJSON struct {
	GroupID     apijson.Field
	Products    apijson.Field
	Status      apijson.Field
	GroupName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductCollectionGetResponseGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionGetResponseGroupJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionGetResponseGroupsProduct struct {
	ID          string `json:"id" api:"required"`
	AddonsCount int64  `json:"addons_count" api:"required"`
	FilesCount  int64  `json:"files_count" api:"required"`
	// Whether this product has any credit entitlements attached
	HasCreditEntitlements bool     `json:"has_credit_entitlements" api:"required"`
	IsRecurring           bool     `json:"is_recurring" api:"required"`
	LicenseKeyEnabled     bool     `json:"license_key_enabled" api:"required"`
	MetersCount           int64    `json:"meters_count" api:"required"`
	ProductID             string   `json:"product_id" api:"required"`
	Status                bool     `json:"status" api:"required"`
	Currency              Currency `json:"currency" api:"nullable"`
	Description           string   `json:"description" api:"nullable"`
	Name                  string   `json:"name" api:"nullable"`
	Price                 int64    `json:"price" api:"nullable"`
	// One-time price details.
	PriceDetail Price `json:"price_detail" api:"nullable"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory  TaxCategory                                   `json:"tax_category" api:"nullable"`
	TaxInclusive bool                                          `json:"tax_inclusive" api:"nullable"`
	JSON         productCollectionGetResponseGroupsProductJSON `json:"-"`
}

// productCollectionGetResponseGroupsProductJSON contains the JSON metadata for the
// struct [ProductCollectionGetResponseGroupsProduct]
type productCollectionGetResponseGroupsProductJSON struct {
	ID                    apijson.Field
	AddonsCount           apijson.Field
	FilesCount            apijson.Field
	HasCreditEntitlements apijson.Field
	IsRecurring           apijson.Field
	LicenseKeyEnabled     apijson.Field
	MetersCount           apijson.Field
	ProductID             apijson.Field
	Status                apijson.Field
	Currency              apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	Price                 apijson.Field
	PriceDetail           apijson.Field
	TaxCategory           apijson.Field
	TaxInclusive          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ProductCollectionGetResponseGroupsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionGetResponseGroupsProductJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionListResponse struct {
	// Collection ID
	ID string `json:"id" api:"required"`
	// Timestamp when created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Collection name
	Name string `json:"name" api:"required"`
	// Number of products in the collection
	ProductsCount int64 `json:"products_count" api:"required"`
	// Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Collection description
	Description string `json:"description" api:"nullable"`
	// Collection image URL
	Image string                            `json:"image" api:"nullable"`
	JSON  productCollectionListResponseJSON `json:"-"`
}

// productCollectionListResponseJSON contains the JSON metadata for the struct
// [ProductCollectionListResponse]
type productCollectionListResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Name          apijson.Field
	ProductsCount apijson.Field
	UpdatedAt     apijson.Field
	Description   apijson.Field
	Image         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProductCollectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionListResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionUnarchiveResponse struct {
	// Collection ID that was unarchived
	CollectionID string `json:"collection_id" api:"required"`
	// Product IDs that were excluded because they are archived
	ExcludedProductIDs []string `json:"excluded_product_ids" api:"required"`
	// Success message
	Message string                                 `json:"message" api:"required"`
	JSON    productCollectionUnarchiveResponseJSON `json:"-"`
}

// productCollectionUnarchiveResponseJSON contains the JSON metadata for the struct
// [ProductCollectionUnarchiveResponse]
type productCollectionUnarchiveResponseJSON struct {
	CollectionID       apijson.Field
	ExcludedProductIDs apijson.Field
	Message            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProductCollectionUnarchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionUnarchiveResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionUpdateImagesResponse struct {
	// Presigned S3 URL for uploading the image
	URL string `json:"url" api:"required"`
	// Optional image ID (present when force_update is true)
	ImageID string                                    `json:"image_id" api:"nullable" format:"uuid"`
	JSON    productCollectionUpdateImagesResponseJSON `json:"-"`
}

// productCollectionUpdateImagesResponseJSON contains the JSON metadata for the
// struct [ProductCollectionUpdateImagesResponse]
type productCollectionUpdateImagesResponseJSON struct {
	URL         apijson.Field
	ImageID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductCollectionUpdateImagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionUpdateImagesResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionNewParams struct {
	// Groups of products in this collection
	Groups param.Field[[]ProductCollectionNewParamsGroup] `json:"groups" api:"required"`
	// Name of the product collection
	Name param.Field[string] `json:"name" api:"required"`
	// Brand id for the collection, if not provided will default to primary brand
	BrandID param.Field[string] `json:"brand_id"`
	// Optional description of the product collection
	Description param.Field[string] `json:"description"`
}

func (r ProductCollectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionNewParamsGroup struct {
	// Products in this group
	Products param.Field[[]ProductCollectionNewParamsGroupsProduct] `json:"products" api:"required"`
	// Optional group name. Multiple groups can have null names, but named groups must
	// be unique per collection
	GroupName param.Field[string] `json:"group_name"`
	// Status of the group (defaults to true if not provided)
	Status param.Field[bool] `json:"status"`
}

func (r ProductCollectionNewParamsGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionNewParamsGroupsProduct struct {
	// Product ID to include in the group
	ProductID param.Field[string] `json:"product_id" api:"required"`
	// Status of the product in this group (defaults to true if not provided)
	Status param.Field[bool] `json:"status"`
}

func (r ProductCollectionNewParamsGroupsProduct) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionUpdateParams struct {
	// Optional brand_id update
	BrandID param.Field[string] `json:"brand_id"`
	// Optional description update - pass null to remove, omit to keep unchanged
	Description param.Field[string] `json:"description"`
	// Optional new order for groups (array of group UUIDs in desired order)
	GroupOrder param.Field[[]string] `json:"group_order" format:"uuid"`
	// Optional image update - pass null to remove, omit to keep unchanged
	ImageID param.Field[string] `json:"image_id" format:"uuid"`
	// Optional new name for the collection
	Name param.Field[string] `json:"name"`
}

func (r ProductCollectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionListParams struct {
	// List archived collections
	Archived param.Field[bool] `query:"archived"`
	// Filter by Brand id
	BrandID param.Field[string] `query:"brand_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [ProductCollectionListParams]'s query parameters as
// `url.Values`.
func (r ProductCollectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProductCollectionUpdateImagesParams struct {
	// If true, generates a new image ID to force cache invalidation
	ForceUpdate param.Field[bool] `query:"force_update"`
}

// URLQuery serializes [ProductCollectionUpdateImagesParams]'s query parameters as
// `url.Values`.
func (r ProductCollectionUpdateImagesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
