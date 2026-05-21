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

func (r *ProductCollectionGroupService) New(ctx context.Context, id string, body ProductCollectionGroupNewParams, opts ...option.RequestOption) (res *ProductCollectionGroupResponse, err error) {
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

type GroupProductParam struct {
	// Product ID to include in the group
	ProductID param.Field[string] `json:"product_id" api:"required"`
	// Status of the product in this group (defaults to true if not provided)
	Status param.Field[bool] `json:"status"`
}

func (r GroupProductParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionGroupDetailsParam struct {
	// Products in this group
	Products param.Field[[]GroupProductParam] `json:"products" api:"required"`
	// Optional group name. Multiple groups can have null names, but named groups must
	// be unique per collection
	GroupName param.Field[string] `json:"group_name"`
	// Status of the group (defaults to true if not provided)
	Status param.Field[bool] `json:"status"`
}

func (r ProductCollectionGroupDetailsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionGroupResponse struct {
	GroupID   string                             `json:"group_id" api:"required"`
	Products  []ProductCollectionProduct         `json:"products" api:"required"`
	Status    bool                               `json:"status" api:"required"`
	GroupName string                             `json:"group_name" api:"nullable"`
	JSON      productCollectionGroupResponseJSON `json:"-"`
}

// productCollectionGroupResponseJSON contains the JSON metadata for the struct
// [ProductCollectionGroupResponse]
type productCollectionGroupResponseJSON struct {
	GroupID     apijson.Field
	Products    apijson.Field
	Status      apijson.Field
	GroupName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductCollectionGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionGroupResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionGroupNewParams struct {
	ProductCollectionGroupDetails ProductCollectionGroupDetailsParam `json:"product_collection_group_details" api:"required"`
}

func (r ProductCollectionGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ProductCollectionGroupDetails)
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
