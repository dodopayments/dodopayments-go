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

// ProductCollectionGroupItemService contains methods and other services that help
// with interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductCollectionGroupItemService] method instead.
type ProductCollectionGroupItemService struct {
	Options []option.RequestOption
}

// NewProductCollectionGroupItemService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProductCollectionGroupItemService(opts ...option.RequestOption) (r *ProductCollectionGroupItemService) {
	r = &ProductCollectionGroupItemService{}
	r.Options = opts
	return
}

func (r *ProductCollectionGroupItemService) New(ctx context.Context, id string, groupID string, body ProductCollectionGroupItemNewParams, opts ...option.RequestOption) (res *[]ProductCollectionGroupItemNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("product-collections/%s/groups/%s/items", id, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ProductCollectionGroupItemService) Update(ctx context.Context, id string, groupID string, itemID string, body ProductCollectionGroupItemUpdateParams, opts ...option.RequestOption) (err error) {
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
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return err
	}
	path := fmt.Sprintf("product-collections/%s/groups/%s/items/%s", id, groupID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

func (r *ProductCollectionGroupItemService) Delete(ctx context.Context, id string, groupID string, itemID string, opts ...option.RequestOption) (err error) {
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
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return err
	}
	path := fmt.Sprintf("product-collections/%s/groups/%s/items/%s", id, groupID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ProductCollectionGroupItemNewResponse struct {
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
	TaxCategory  TaxCategory                               `json:"tax_category" api:"nullable"`
	TaxInclusive bool                                      `json:"tax_inclusive" api:"nullable"`
	JSON         productCollectionGroupItemNewResponseJSON `json:"-"`
}

// productCollectionGroupItemNewResponseJSON contains the JSON metadata for the
// struct [ProductCollectionGroupItemNewResponse]
type productCollectionGroupItemNewResponseJSON struct {
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

func (r *ProductCollectionGroupItemNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productCollectionGroupItemNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProductCollectionGroupItemNewParams struct {
	// Products to add to the group
	Products param.Field[[]ProductCollectionGroupItemNewParamsProduct] `json:"products" api:"required"`
}

func (r ProductCollectionGroupItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionGroupItemNewParamsProduct struct {
	// Product ID to include in the group
	ProductID param.Field[string] `json:"product_id" api:"required"`
	// Status of the product in this group (defaults to true if not provided)
	Status param.Field[bool] `json:"status"`
}

func (r ProductCollectionGroupItemNewParamsProduct) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionGroupItemUpdateParams struct {
	// Status of the product in the group
	Status param.Field[bool] `json:"status" api:"required"`
}

func (r ProductCollectionGroupItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
