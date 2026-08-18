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
)

// BrandService contains methods and other services that help with interacting with
// the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	Options []option.RequestOption
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r *BrandService) {
	r = &BrandService{}
	r.Options = opts
	return
}

func (r *BrandService) New(ctx context.Context, body BrandNewParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Thin handler just calls `get_brand` and wraps in `Json(...)`
func (r *BrandService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *BrandService) Update(ctx context.Context, id string, body BrandUpdateParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *BrandService) List(ctx context.Context, query BrandListParams, opts ...option.RequestOption) (res *BrandListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Archive a brand. Its products, live subscriptions, and product collections move
// to the `move_products_to` brand. Archive is permanent.
func (r *BrandService) Archive(ctx context.Context, id string, body BrandArchiveParams, opts ...option.RequestOption) (res *BrandArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("brands/%s/archive", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BrandService) UpdateImages(ctx context.Context, id string, opts ...option.RequestOption) (res *BrandUpdateImagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("brands/%s/images", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type Brand struct {
	BrandID             string                  `json:"brand_id" api:"required"`
	BusinessID          string                  `json:"business_id" api:"required"`
	Enabled             bool                    `json:"enabled" api:"required"`
	StatementDescriptor string                  `json:"statement_descriptor" api:"required"`
	VerificationEnabled bool                    `json:"verification_enabled" api:"required"`
	VerificationStatus  BrandVerificationStatus `json:"verification_status" api:"required"`
	// Time the brand was archived. Null for an active brand.
	ArchivedAt  time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	Description string    `json:"description" api:"nullable"`
	Image       string    `json:"image" api:"nullable"`
	Name        string    `json:"name" api:"nullable"`
	// Incase the brand verification fails or is put on hold
	ReasonForHold string    `json:"reason_for_hold" api:"nullable"`
	SupportEmail  string    `json:"support_email" api:"nullable"`
	URL           string    `json:"url" api:"nullable"`
	JSON          brandJSON `json:"-"`
}

// brandJSON contains the JSON metadata for the struct [Brand]
type brandJSON struct {
	BrandID             apijson.Field
	BusinessID          apijson.Field
	Enabled             apijson.Field
	StatementDescriptor apijson.Field
	VerificationEnabled apijson.Field
	VerificationStatus  apijson.Field
	ArchivedAt          apijson.Field
	Description         apijson.Field
	Image               apijson.Field
	Name                apijson.Field
	ReasonForHold       apijson.Field
	SupportEmail        apijson.Field
	URL                 apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Brand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandJSON) RawJSON() string {
	return r.raw
}

type BrandVerificationStatus string

const (
	BrandVerificationStatusSuccess BrandVerificationStatus = "Success"
	BrandVerificationStatusFail    BrandVerificationStatus = "Fail"
	BrandVerificationStatusReview  BrandVerificationStatus = "Review"
	BrandVerificationStatusHold    BrandVerificationStatus = "Hold"
)

func (r BrandVerificationStatus) IsKnown() bool {
	switch r {
	case BrandVerificationStatusSuccess, BrandVerificationStatusFail, BrandVerificationStatusReview, BrandVerificationStatusHold:
		return true
	}
	return false
}

type BrandListResponse struct {
	// List of brands for this business
	Items []Brand               `json:"items" api:"required"`
	JSON  brandListResponseJSON `json:"-"`
}

// brandListResponseJSON contains the JSON metadata for the struct
// [BrandListResponse]
type brandListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandListResponseJSON) RawJSON() string {
	return r.raw
}

type BrandArchiveResponse struct {
	// Time the brand was archived.
	ArchivedAt time.Time `json:"archived_at" api:"required" format:"date-time"`
	// The archived brand.
	BrandID string `json:"brand_id" api:"required"`
	// Count of product collections moved to the target brand.
	CollectionsMoved int64 `json:"collections_moved" api:"required"`
	// Count of products moved to the target brand.
	ProductsMoved int64 `json:"products_moved" api:"required"`
	// Count of live subscriptions moved to the target brand.
	SubscriptionsMoved int64 `json:"subscriptions_moved" api:"required"`
	// Brand that received the moved records. Null when no target was given.
	MovedToBrandID string                   `json:"moved_to_brand_id" api:"nullable"`
	JSON           brandArchiveResponseJSON `json:"-"`
}

// brandArchiveResponseJSON contains the JSON metadata for the struct
// [BrandArchiveResponse]
type brandArchiveResponseJSON struct {
	ArchivedAt         apijson.Field
	BrandID            apijson.Field
	CollectionsMoved   apijson.Field
	ProductsMoved      apijson.Field
	SubscriptionsMoved apijson.Field
	MovedToBrandID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BrandArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type BrandUpdateImagesResponse struct {
	// UUID that will be used as the image identifier/key suffix
	ImageID string `json:"image_id" api:"required" format:"uuid"`
	// Presigned URL to upload the image
	URL  string                        `json:"url" api:"required"`
	JSON brandUpdateImagesResponseJSON `json:"-"`
}

// brandUpdateImagesResponseJSON contains the JSON metadata for the struct
// [BrandUpdateImagesResponse]
type brandUpdateImagesResponseJSON struct {
	ImageID     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandUpdateImagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandUpdateImagesResponseJSON) RawJSON() string {
	return r.raw
}

type BrandNewParams struct {
	Description         param.Field[string] `json:"description"`
	Name                param.Field[string] `json:"name"`
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	SupportEmail        param.Field[string] `json:"support_email"`
	URL                 param.Field[string] `json:"url"`
}

func (r BrandNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrandUpdateParams struct {
	Description param.Field[string] `json:"description"`
	// The UUID you got back from the presigned‐upload call
	ImageID             param.Field[string] `json:"image_id" format:"uuid"`
	Name                param.Field[string] `json:"name"`
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	SupportEmail        param.Field[string] `json:"support_email"`
	URL                 param.Field[string] `json:"url"`
}

func (r BrandUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrandListParams struct {
	// Set to true to also list archived brands. Default false.
	IncludeArchived param.Field[bool] `query:"include_archived"`
}

// URLQuery serializes [BrandListParams]'s query parameters as `url.Values`.
func (r BrandListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandArchiveParams struct {
	// Brand that takes over the products and the live subscriptions of the brand you
	// archive. It must be a brand of the same business, and it must not be archived.
	// The primary brand (its brand id is the business id) is a valid target. Omit this
	// field only when the brand holds no products and no live subscriptions.
	MoveProductsTo param.Field[string] `json:"move_products_to"`
}

func (r BrandArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
