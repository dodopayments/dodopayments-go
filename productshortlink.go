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
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// ProductShortLinkService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductShortLinkService] method instead.
type ProductShortLinkService struct {
	Options []option.RequestOption
}

// NewProductShortLinkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductShortLinkService(opts ...option.RequestOption) (r ProductShortLinkService) {
	r = ProductShortLinkService{}
	r.Options = opts
	return
}

// Gives a Short Checkout URL with custom slug for a product. Uses a Static
// Checkout URL under the hood.
func (r *ProductShortLinkService) New(ctx context.Context, id string, body ProductShortLinkNewParams, opts ...option.RequestOption) (res *ProductShortLinkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/short_links", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all short links created by the business.
func (r *ProductShortLinkService) List(ctx context.Context, query ProductShortLinkListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[ProductShortLinkListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "products/short_links"
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

// Lists all short links created by the business.
func (r *ProductShortLinkService) ListAutoPaging(ctx context.Context, query ProductShortLinkListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[ProductShortLinkListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type ProductShortLinkNewResponse struct {
	// Full URL.
	FullURL string `json:"full_url,required"`
	// Short URL.
	ShortURL string `json:"short_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullURL     respjson.Field
		ShortURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductShortLinkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductShortLinkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductShortLinkListResponse struct {
	// When the short url was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Full URL the short url redirects to
	FullURL string `json:"full_url,required"`
	// Product ID associated with the short link
	ProductID string `json:"product_id,required"`
	// Short URL
	ShortURL string `json:"short_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		FullURL     respjson.Field
		ProductID   respjson.Field
		ShortURL    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductShortLinkListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProductShortLinkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductShortLinkNewParams struct {
	// Slug for the short link.
	Slug string `json:"slug,required"`
	// Static Checkout URL parameters to apply to the resulting short URL.
	StaticCheckoutParams map[string]string `json:"static_checkout_params,omitzero"`
	paramObj
}

func (r ProductShortLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProductShortLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProductShortLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductShortLinkListParams struct {
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter by product ID
	ProductID param.Opt[string] `query:"product_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProductShortLinkListParams]'s query parameters as
// `url.Values`.
func (r ProductShortLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
