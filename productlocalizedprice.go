// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// ProductLocalizedPriceService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductLocalizedPriceService] method instead.
type ProductLocalizedPriceService struct {
	Options []option.RequestOption
}

// NewProductLocalizedPriceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductLocalizedPriceService(opts ...option.RequestOption) (r *ProductLocalizedPriceService) {
	r = &ProductLocalizedPriceService{}
	r.Options = opts
	return
}

func (r *ProductLocalizedPriceService) New(ctx context.Context, productID string, body ProductLocalizedPriceNewParams, opts ...option.RequestOption) (res *LocalizedPrice, err error) {
	opts = slices.Concat(r.Options, opts)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s/localized-prices", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *ProductLocalizedPriceService) Get(ctx context.Context, productID string, id string, opts ...option.RequestOption) (res *LocalizedPrice, err error) {
	opts = slices.Concat(r.Options, opts)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s/localized-prices/%s", productID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ProductLocalizedPriceService) Update(ctx context.Context, productID string, id string, body ProductLocalizedPriceUpdateParams, opts ...option.RequestOption) (res *LocalizedPrice, err error) {
	opts = slices.Concat(r.Options, opts)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s/localized-prices/%s", productID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

func (r *ProductLocalizedPriceService) List(ctx context.Context, productID string, opts ...option.RequestOption) (res *ListLocalizedPricesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s/localized-prices", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *ProductLocalizedPriceService) Archive(ctx context.Context, productID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("products/%s/localized-prices/%s", productID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ListLocalizedPricesResponse struct {
	Items []LocalizedPrice                `json:"items" api:"required"`
	JSON  listLocalizedPricesResponseJSON `json:"-"`
}

// listLocalizedPricesResponseJSON contains the JSON metadata for the struct
// [ListLocalizedPricesResponse]
type listLocalizedPricesResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListLocalizedPricesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listLocalizedPricesResponseJSON) RawJSON() string {
	return r.raw
}

type LocalizedPrice struct {
	// Unique identifier for the localized price.
	ID string `json:"id" api:"required"`
	// Amount in the smallest currency unit (e.g., cents).
	Amount int64 `json:"amount" api:"required"`
	// Timestamp when the localized price was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Currency to charge in.
	Currency Currency `json:"currency" api:"required"`
	// Pricing mode of the rule: by_currency or by_country.
	Mode PricingMode `json:"mode" api:"required"`
	// Product this localized price belongs to.
	ProductID string `json:"product_id" api:"required"`
	// Timestamp when the localized price was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Country the rule applies to. Only set when mode is by_country.
	CountryCode CountryCode        `json:"country_code" api:"nullable"`
	JSON        localizedPriceJSON `json:"-"`
}

// localizedPriceJSON contains the JSON metadata for the struct [LocalizedPrice]
type localizedPriceJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	Mode        apijson.Field
	ProductID   apijson.Field
	UpdatedAt   apijson.Field
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocalizedPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r localizedPriceJSON) RawJSON() string {
	return r.raw
}

type PricingMode string

const (
	PricingModeByCurrency PricingMode = "by_currency"
	PricingModeByCountry  PricingMode = "by_country"
)

func (r PricingMode) IsKnown() bool {
	switch r {
	case PricingModeByCurrency, PricingModeByCountry:
		return true
	}
	return false
}

type ProductLocalizedPriceNewParams struct {
	// Amount in the smallest currency unit (e.g., cents). Must be greater than zero.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// Currency to charge in. Must be a supported currency.
	Currency param.Field[Currency] `json:"currency" api:"required"`
	// Required when the product's pricing_mode is by_country; forbidden when
	// by_currency.
	CountryCode param.Field[CountryCode] `json:"country_code"`
}

func (r ProductLocalizedPriceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductLocalizedPriceUpdateParams struct {
	// New amount in the smallest currency unit (e.g., cents). Must be greater than
	// zero. The currency and country_code of an existing rule cannot be changed.
	Amount param.Field[int64] `json:"amount"`
}

func (r ProductLocalizedPriceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
