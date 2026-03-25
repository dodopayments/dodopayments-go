// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// ProductCollectionGroupService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductCollectionGroupService] method instead.
type ProductCollectionGroupService struct {
	Options []option.RequestOption
	Items   *ProductCollectionGroupItemService
}

// NewProductCollectionGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductCollectionGroupService(opts ...option.RequestOption) (r *ProductCollectionGroupService) {
	r = &ProductCollectionGroupService{}
	r.Options = opts
	r.Items = NewProductCollectionGroupItemService(opts...)
	return
}

func (r *ProductCollectionGroupService) New(ctx context.Context, id string, body ProductCollectionGroupNewParams, opts ...option.RequestOption) (res *ProductCollectionGroupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("product-collections/%s/groups", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ProductCollectionGroupService) Update(ctx context.Context, id string, groupID string, body ProductCollectionGroupUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return err
	}
	path := fmt.Sprintf("product-collections/%s/groups/%s", id, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

func (r *ProductCollectionGroupService) Delete(ctx context.Context, id string, groupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return err
	}
	path := fmt.Sprintf("product-collections/%s/groups/%s", id, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ProductCollectionGroupNewResponse struct {
	GroupID   string                                     `json:"group_id" api:"required"`
	Products  []ProductCollectionGroupNewResponseProduct `json:"products" api:"required"`
	Status    bool                                       `json:"status" api:"required"`
	GroupName string                                     `json:"group_name" api:"nullable"`
	JSON      productCollectionGroupNewResponseJSON      `json:"-"`
}

// productCollectionGroupNewResponseJSON contains the JSON metadata for the struct
// [ProductCollectionGroupNewResponse]
type productCollectionGroupNewResponseJSON struct {
	GroupID     apijson.Field
	Products    apijson.Field
	Status      apijson.Field
	GroupName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductCollectionGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionGroupNewResponseProduct struct {
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
	TaxCategory  TaxCategory                                  `json:"tax_category" api:"nullable"`
	TaxInclusive bool                                         `json:"tax_inclusive" api:"nullable"`
	JSON         productCollectionGroupNewResponseProductJSON `json:"-"`
}

// productCollectionGroupNewResponseProductJSON contains the JSON metadata for the
// struct [ProductCollectionGroupNewResponseProduct]
type productCollectionGroupNewResponseProductJSON struct {
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

func (r *ProductCollectionGroupNewResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionGroupNewResponseProductJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionGroupNewParams struct {
	// Products in this group
	Products param.Field[[]ProductCollectionGroupNewParamsProduct] `json:"products" api:"required"`
	// Optional group name. Multiple groups can have null names, but named groups must
	// be unique per collection
	GroupName param.Field[string] `json:"group_name"`
	// Status of the group (defaults to true if not provided)
	Status param.Field[bool] `json:"status"`
}

func (r ProductCollectionGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionGroupNewParamsProduct struct {
	// Product ID to include in the group
	ProductID param.Field[string] `json:"product_id" api:"required"`
	// Status of the product in this group (defaults to true if not provided)
	Status param.Field[bool] `json:"status"`
}

func (r ProductCollectionGroupNewParamsProduct) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionGroupUpdateParams struct {
	// Optional group name update: Some(Some(name)) = set name, Some(None) = clear
	// name, None = no change
	GroupName param.Field[string] `json:"group_name"`
	// Optional new order for products in this group (array of
	// product_collection_group_pdts UUIDs)
	ProductOrder param.Field[[]string] `json:"product_order" format:"uuid"`
	// Optional status update
	Status param.Field[bool] `json:"status"`
}

func (r ProductCollectionGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
