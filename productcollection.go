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

func (r *ProductCollectionService) New(ctx context.Context, body ProductCollectionNewParams, opts ...option.RequestOption) (res *ProductCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "product-collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ProductCollectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductCollection, err error) {
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

type ProductCollection struct {
	// Unique identifier for the product collection
	ID string `json:"id" api:"required"`
	// Brand ID for the collection
	BrandID string `json:"brand_id" api:"required"`
	// Timestamp when the collection was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Groups in this collection
	Groups []ProductCollectionGroupResponse `json:"groups" api:"required"`
	// Name of the collection
	Name string `json:"name" api:"required"`
	// Timestamp when the collection was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Description of the collection
	Description string `json:"description" api:"nullable"`
	// Default effective_at setting for subscription plan downgrades (null = inherit
	// from business)
	EffectiveAtOnDowngrade ProductCollectionEffectiveAtOnDowngrade `json:"effective_at_on_downgrade" api:"nullable"`
	// Default effective_at setting for subscription plan upgrades (null = inherit from
	// business)
	EffectiveAtOnUpgrade ProductCollectionEffectiveAtOnUpgrade `json:"effective_at_on_upgrade" api:"nullable"`
	// URL of the collection image
	Image string `json:"image" api:"nullable"`
	// Default behavior for subscription plan changes on payment failure (null =
	// inherit from business)
	OnPaymentFailure ProductCollectionOnPaymentFailure `json:"on_payment_failure" api:"nullable"`
	// Default proration billing mode for subscription plan downgrades (null = inherit
	// from business)
	ProrationBillingModeOnDowngrade ProductCollectionProrationBillingModeOnDowngrade `json:"proration_billing_mode_on_downgrade" api:"nullable"`
	// Default proration billing mode for subscription plan upgrades (null = inherit
	// from business)
	ProrationBillingModeOnUpgrade ProductCollectionProrationBillingModeOnUpgrade `json:"proration_billing_mode_on_upgrade" api:"nullable"`
	JSON                          productCollectionJSON                          `json:"-"`
}

// productCollectionJSON contains the JSON metadata for the struct
// [ProductCollection]
type productCollectionJSON struct {
	ID                              apijson.Field
	BrandID                         apijson.Field
	CreatedAt                       apijson.Field
	Groups                          apijson.Field
	Name                            apijson.Field
	UpdatedAt                       apijson.Field
	Description                     apijson.Field
	EffectiveAtOnDowngrade          apijson.Field
	EffectiveAtOnUpgrade            apijson.Field
	Image                           apijson.Field
	OnPaymentFailure                apijson.Field
	ProrationBillingModeOnDowngrade apijson.Field
	ProrationBillingModeOnUpgrade   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *ProductCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionJSON) RawJSON() string {
	return r.raw
}

// Default effective_at setting for subscription plan downgrades (null = inherit
// from business)
type ProductCollectionEffectiveAtOnDowngrade string

const (
	ProductCollectionEffectiveAtOnDowngradeImmediately     ProductCollectionEffectiveAtOnDowngrade = "immediately"
	ProductCollectionEffectiveAtOnDowngradeNextBillingDate ProductCollectionEffectiveAtOnDowngrade = "next_billing_date"
)

func (r ProductCollectionEffectiveAtOnDowngrade) IsKnown() bool {
	switch r {
	case ProductCollectionEffectiveAtOnDowngradeImmediately, ProductCollectionEffectiveAtOnDowngradeNextBillingDate:
		return true
	}
	return false
}

// Default effective_at setting for subscription plan upgrades (null = inherit from
// business)
type ProductCollectionEffectiveAtOnUpgrade string

const (
	ProductCollectionEffectiveAtOnUpgradeImmediately     ProductCollectionEffectiveAtOnUpgrade = "immediately"
	ProductCollectionEffectiveAtOnUpgradeNextBillingDate ProductCollectionEffectiveAtOnUpgrade = "next_billing_date"
)

func (r ProductCollectionEffectiveAtOnUpgrade) IsKnown() bool {
	switch r {
	case ProductCollectionEffectiveAtOnUpgradeImmediately, ProductCollectionEffectiveAtOnUpgradeNextBillingDate:
		return true
	}
	return false
}

// Default behavior for subscription plan changes on payment failure (null =
// inherit from business)
type ProductCollectionOnPaymentFailure string

const (
	ProductCollectionOnPaymentFailurePreventChange ProductCollectionOnPaymentFailure = "prevent_change"
	ProductCollectionOnPaymentFailureApplyChange   ProductCollectionOnPaymentFailure = "apply_change"
)

func (r ProductCollectionOnPaymentFailure) IsKnown() bool {
	switch r {
	case ProductCollectionOnPaymentFailurePreventChange, ProductCollectionOnPaymentFailureApplyChange:
		return true
	}
	return false
}

// Default proration billing mode for subscription plan downgrades (null = inherit
// from business)
type ProductCollectionProrationBillingModeOnDowngrade string

const (
	ProductCollectionProrationBillingModeOnDowngradeProratedImmediately   ProductCollectionProrationBillingModeOnDowngrade = "prorated_immediately"
	ProductCollectionProrationBillingModeOnDowngradeFullImmediately       ProductCollectionProrationBillingModeOnDowngrade = "full_immediately"
	ProductCollectionProrationBillingModeOnDowngradeDifferenceImmediately ProductCollectionProrationBillingModeOnDowngrade = "difference_immediately"
	ProductCollectionProrationBillingModeOnDowngradeDoNotBill             ProductCollectionProrationBillingModeOnDowngrade = "do_not_bill"
)

func (r ProductCollectionProrationBillingModeOnDowngrade) IsKnown() bool {
	switch r {
	case ProductCollectionProrationBillingModeOnDowngradeProratedImmediately, ProductCollectionProrationBillingModeOnDowngradeFullImmediately, ProductCollectionProrationBillingModeOnDowngradeDifferenceImmediately, ProductCollectionProrationBillingModeOnDowngradeDoNotBill:
		return true
	}
	return false
}

// Default proration billing mode for subscription plan upgrades (null = inherit
// from business)
type ProductCollectionProrationBillingModeOnUpgrade string

const (
	ProductCollectionProrationBillingModeOnUpgradeProratedImmediately   ProductCollectionProrationBillingModeOnUpgrade = "prorated_immediately"
	ProductCollectionProrationBillingModeOnUpgradeFullImmediately       ProductCollectionProrationBillingModeOnUpgrade = "full_immediately"
	ProductCollectionProrationBillingModeOnUpgradeDifferenceImmediately ProductCollectionProrationBillingModeOnUpgrade = "difference_immediately"
	ProductCollectionProrationBillingModeOnUpgradeDoNotBill             ProductCollectionProrationBillingModeOnUpgrade = "do_not_bill"
)

func (r ProductCollectionProrationBillingModeOnUpgrade) IsKnown() bool {
	switch r {
	case ProductCollectionProrationBillingModeOnUpgradeProratedImmediately, ProductCollectionProrationBillingModeOnUpgradeFullImmediately, ProductCollectionProrationBillingModeOnUpgradeDifferenceImmediately, ProductCollectionProrationBillingModeOnUpgradeDoNotBill:
		return true
	}
	return false
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
	Groups param.Field[[]ProductCollectionGroupDetailsParam] `json:"groups" api:"required"`
	// Name of the product collection
	Name param.Field[string] `json:"name" api:"required"`
	// Brand id for the collection, if not provided will default to primary brand
	BrandID param.Field[string] `json:"brand_id"`
	// Optional description of the product collection
	Description param.Field[string] `json:"description"`
	// Default effective_at setting for subscription plan downgrades (NULL = inherit
	// from business)
	EffectiveAtOnDowngrade param.Field[ProductCollectionNewParamsEffectiveAtOnDowngrade] `json:"effective_at_on_downgrade"`
	// Default effective_at setting for subscription plan upgrades (NULL = inherit from
	// business)
	EffectiveAtOnUpgrade param.Field[ProductCollectionNewParamsEffectiveAtOnUpgrade] `json:"effective_at_on_upgrade"`
	// Default behavior for subscription plan changes on payment failure (NULL =
	// inherit from business)
	OnPaymentFailure param.Field[ProductCollectionNewParamsOnPaymentFailure] `json:"on_payment_failure"`
	// Default proration billing mode for subscription plan downgrades (NULL = inherit
	// from business)
	ProrationBillingModeOnDowngrade param.Field[ProductCollectionNewParamsProrationBillingModeOnDowngrade] `json:"proration_billing_mode_on_downgrade"`
	// Default proration billing mode for subscription plan upgrades (NULL = inherit
	// from business)
	ProrationBillingModeOnUpgrade param.Field[ProductCollectionNewParamsProrationBillingModeOnUpgrade] `json:"proration_billing_mode_on_upgrade"`
}

func (r ProductCollectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Default effective_at setting for subscription plan downgrades (NULL = inherit
// from business)
type ProductCollectionNewParamsEffectiveAtOnDowngrade string

const (
	ProductCollectionNewParamsEffectiveAtOnDowngradeImmediately     ProductCollectionNewParamsEffectiveAtOnDowngrade = "immediately"
	ProductCollectionNewParamsEffectiveAtOnDowngradeNextBillingDate ProductCollectionNewParamsEffectiveAtOnDowngrade = "next_billing_date"
)

func (r ProductCollectionNewParamsEffectiveAtOnDowngrade) IsKnown() bool {
	switch r {
	case ProductCollectionNewParamsEffectiveAtOnDowngradeImmediately, ProductCollectionNewParamsEffectiveAtOnDowngradeNextBillingDate:
		return true
	}
	return false
}

// Default effective_at setting for subscription plan upgrades (NULL = inherit from
// business)
type ProductCollectionNewParamsEffectiveAtOnUpgrade string

const (
	ProductCollectionNewParamsEffectiveAtOnUpgradeImmediately     ProductCollectionNewParamsEffectiveAtOnUpgrade = "immediately"
	ProductCollectionNewParamsEffectiveAtOnUpgradeNextBillingDate ProductCollectionNewParamsEffectiveAtOnUpgrade = "next_billing_date"
)

func (r ProductCollectionNewParamsEffectiveAtOnUpgrade) IsKnown() bool {
	switch r {
	case ProductCollectionNewParamsEffectiveAtOnUpgradeImmediately, ProductCollectionNewParamsEffectiveAtOnUpgradeNextBillingDate:
		return true
	}
	return false
}

// Default behavior for subscription plan changes on payment failure (NULL =
// inherit from business)
type ProductCollectionNewParamsOnPaymentFailure string

const (
	ProductCollectionNewParamsOnPaymentFailurePreventChange ProductCollectionNewParamsOnPaymentFailure = "prevent_change"
	ProductCollectionNewParamsOnPaymentFailureApplyChange   ProductCollectionNewParamsOnPaymentFailure = "apply_change"
)

func (r ProductCollectionNewParamsOnPaymentFailure) IsKnown() bool {
	switch r {
	case ProductCollectionNewParamsOnPaymentFailurePreventChange, ProductCollectionNewParamsOnPaymentFailureApplyChange:
		return true
	}
	return false
}

// Default proration billing mode for subscription plan downgrades (NULL = inherit
// from business)
type ProductCollectionNewParamsProrationBillingModeOnDowngrade string

const (
	ProductCollectionNewParamsProrationBillingModeOnDowngradeProratedImmediately   ProductCollectionNewParamsProrationBillingModeOnDowngrade = "prorated_immediately"
	ProductCollectionNewParamsProrationBillingModeOnDowngradeFullImmediately       ProductCollectionNewParamsProrationBillingModeOnDowngrade = "full_immediately"
	ProductCollectionNewParamsProrationBillingModeOnDowngradeDifferenceImmediately ProductCollectionNewParamsProrationBillingModeOnDowngrade = "difference_immediately"
	ProductCollectionNewParamsProrationBillingModeOnDowngradeDoNotBill             ProductCollectionNewParamsProrationBillingModeOnDowngrade = "do_not_bill"
)

func (r ProductCollectionNewParamsProrationBillingModeOnDowngrade) IsKnown() bool {
	switch r {
	case ProductCollectionNewParamsProrationBillingModeOnDowngradeProratedImmediately, ProductCollectionNewParamsProrationBillingModeOnDowngradeFullImmediately, ProductCollectionNewParamsProrationBillingModeOnDowngradeDifferenceImmediately, ProductCollectionNewParamsProrationBillingModeOnDowngradeDoNotBill:
		return true
	}
	return false
}

// Default proration billing mode for subscription plan upgrades (NULL = inherit
// from business)
type ProductCollectionNewParamsProrationBillingModeOnUpgrade string

const (
	ProductCollectionNewParamsProrationBillingModeOnUpgradeProratedImmediately   ProductCollectionNewParamsProrationBillingModeOnUpgrade = "prorated_immediately"
	ProductCollectionNewParamsProrationBillingModeOnUpgradeFullImmediately       ProductCollectionNewParamsProrationBillingModeOnUpgrade = "full_immediately"
	ProductCollectionNewParamsProrationBillingModeOnUpgradeDifferenceImmediately ProductCollectionNewParamsProrationBillingModeOnUpgrade = "difference_immediately"
	ProductCollectionNewParamsProrationBillingModeOnUpgradeDoNotBill             ProductCollectionNewParamsProrationBillingModeOnUpgrade = "do_not_bill"
)

func (r ProductCollectionNewParamsProrationBillingModeOnUpgrade) IsKnown() bool {
	switch r {
	case ProductCollectionNewParamsProrationBillingModeOnUpgradeProratedImmediately, ProductCollectionNewParamsProrationBillingModeOnUpgradeFullImmediately, ProductCollectionNewParamsProrationBillingModeOnUpgradeDifferenceImmediately, ProductCollectionNewParamsProrationBillingModeOnUpgradeDoNotBill:
		return true
	}
	return false
}

type ProductCollectionUpdateParams struct {
	// Optional brand_id update
	BrandID param.Field[string] `json:"brand_id"`
	// Optional description update - pass null to remove, omit to keep unchanged
	Description param.Field[string] `json:"description"`
	// Effective_at setting for downgrades: Some(Some(val)) = set, Some(None) = clear
	// (inherit), None = no change
	EffectiveAtOnDowngrade param.Field[ProductCollectionUpdateParamsEffectiveAtOnDowngrade] `json:"effective_at_on_downgrade"`
	// Effective_at setting for upgrades: Some(Some(val)) = set, Some(None) = clear
	// (inherit), None = no change
	EffectiveAtOnUpgrade param.Field[ProductCollectionUpdateParamsEffectiveAtOnUpgrade] `json:"effective_at_on_upgrade"`
	// Optional new order for groups (array of group UUIDs in desired order)
	GroupOrder param.Field[[]string] `json:"group_order" format:"uuid"`
	// Optional image update - pass null to remove, omit to keep unchanged
	ImageID param.Field[string] `json:"image_id" format:"uuid"`
	// Optional new name for the collection
	Name param.Field[string] `json:"name"`
	// On payment failure behavior: Some(Some(val)) = set, Some(None) = clear
	// (inherit), None = no change
	OnPaymentFailure param.Field[ProductCollectionUpdateParamsOnPaymentFailure] `json:"on_payment_failure"`
	// Proration billing mode for downgrades: Some(Some(val)) = set, Some(None) = clear
	// (inherit), None = no change
	ProrationBillingModeOnDowngrade param.Field[ProductCollectionUpdateParamsProrationBillingModeOnDowngrade] `json:"proration_billing_mode_on_downgrade"`
	// Proration billing mode for upgrades: Some(Some(val)) = set, Some(None) = clear
	// (inherit), None = no change
	ProrationBillingModeOnUpgrade param.Field[ProductCollectionUpdateParamsProrationBillingModeOnUpgrade] `json:"proration_billing_mode_on_upgrade"`
}

func (r ProductCollectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Effective_at setting for downgrades: Some(Some(val)) = set, Some(None) = clear
// (inherit), None = no change
type ProductCollectionUpdateParamsEffectiveAtOnDowngrade string

const (
	ProductCollectionUpdateParamsEffectiveAtOnDowngradeImmediately     ProductCollectionUpdateParamsEffectiveAtOnDowngrade = "immediately"
	ProductCollectionUpdateParamsEffectiveAtOnDowngradeNextBillingDate ProductCollectionUpdateParamsEffectiveAtOnDowngrade = "next_billing_date"
)

func (r ProductCollectionUpdateParamsEffectiveAtOnDowngrade) IsKnown() bool {
	switch r {
	case ProductCollectionUpdateParamsEffectiveAtOnDowngradeImmediately, ProductCollectionUpdateParamsEffectiveAtOnDowngradeNextBillingDate:
		return true
	}
	return false
}

// Effective_at setting for upgrades: Some(Some(val)) = set, Some(None) = clear
// (inherit), None = no change
type ProductCollectionUpdateParamsEffectiveAtOnUpgrade string

const (
	ProductCollectionUpdateParamsEffectiveAtOnUpgradeImmediately     ProductCollectionUpdateParamsEffectiveAtOnUpgrade = "immediately"
	ProductCollectionUpdateParamsEffectiveAtOnUpgradeNextBillingDate ProductCollectionUpdateParamsEffectiveAtOnUpgrade = "next_billing_date"
)

func (r ProductCollectionUpdateParamsEffectiveAtOnUpgrade) IsKnown() bool {
	switch r {
	case ProductCollectionUpdateParamsEffectiveAtOnUpgradeImmediately, ProductCollectionUpdateParamsEffectiveAtOnUpgradeNextBillingDate:
		return true
	}
	return false
}

// On payment failure behavior: Some(Some(val)) = set, Some(None) = clear
// (inherit), None = no change
type ProductCollectionUpdateParamsOnPaymentFailure string

const (
	ProductCollectionUpdateParamsOnPaymentFailurePreventChange ProductCollectionUpdateParamsOnPaymentFailure = "prevent_change"
	ProductCollectionUpdateParamsOnPaymentFailureApplyChange   ProductCollectionUpdateParamsOnPaymentFailure = "apply_change"
)

func (r ProductCollectionUpdateParamsOnPaymentFailure) IsKnown() bool {
	switch r {
	case ProductCollectionUpdateParamsOnPaymentFailurePreventChange, ProductCollectionUpdateParamsOnPaymentFailureApplyChange:
		return true
	}
	return false
}

// Proration billing mode for downgrades: Some(Some(val)) = set, Some(None) = clear
// (inherit), None = no change
type ProductCollectionUpdateParamsProrationBillingModeOnDowngrade string

const (
	ProductCollectionUpdateParamsProrationBillingModeOnDowngradeProratedImmediately   ProductCollectionUpdateParamsProrationBillingModeOnDowngrade = "prorated_immediately"
	ProductCollectionUpdateParamsProrationBillingModeOnDowngradeFullImmediately       ProductCollectionUpdateParamsProrationBillingModeOnDowngrade = "full_immediately"
	ProductCollectionUpdateParamsProrationBillingModeOnDowngradeDifferenceImmediately ProductCollectionUpdateParamsProrationBillingModeOnDowngrade = "difference_immediately"
	ProductCollectionUpdateParamsProrationBillingModeOnDowngradeDoNotBill             ProductCollectionUpdateParamsProrationBillingModeOnDowngrade = "do_not_bill"
)

func (r ProductCollectionUpdateParamsProrationBillingModeOnDowngrade) IsKnown() bool {
	switch r {
	case ProductCollectionUpdateParamsProrationBillingModeOnDowngradeProratedImmediately, ProductCollectionUpdateParamsProrationBillingModeOnDowngradeFullImmediately, ProductCollectionUpdateParamsProrationBillingModeOnDowngradeDifferenceImmediately, ProductCollectionUpdateParamsProrationBillingModeOnDowngradeDoNotBill:
		return true
	}
	return false
}

// Proration billing mode for upgrades: Some(Some(val)) = set, Some(None) = clear
// (inherit), None = no change
type ProductCollectionUpdateParamsProrationBillingModeOnUpgrade string

const (
	ProductCollectionUpdateParamsProrationBillingModeOnUpgradeProratedImmediately   ProductCollectionUpdateParamsProrationBillingModeOnUpgrade = "prorated_immediately"
	ProductCollectionUpdateParamsProrationBillingModeOnUpgradeFullImmediately       ProductCollectionUpdateParamsProrationBillingModeOnUpgrade = "full_immediately"
	ProductCollectionUpdateParamsProrationBillingModeOnUpgradeDifferenceImmediately ProductCollectionUpdateParamsProrationBillingModeOnUpgrade = "difference_immediately"
	ProductCollectionUpdateParamsProrationBillingModeOnUpgradeDoNotBill             ProductCollectionUpdateParamsProrationBillingModeOnUpgrade = "do_not_bill"
)

func (r ProductCollectionUpdateParamsProrationBillingModeOnUpgrade) IsKnown() bool {
	switch r {
	case ProductCollectionUpdateParamsProrationBillingModeOnUpgradeProratedImmediately, ProductCollectionUpdateParamsProrationBillingModeOnUpgradeFullImmediately, ProductCollectionUpdateParamsProrationBillingModeOnUpgradeDifferenceImmediately, ProductCollectionUpdateParamsProrationBillingModeOnUpgradeDoNotBill:
		return true
	}
	return false
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
