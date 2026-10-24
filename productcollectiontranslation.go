// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// ProductCollectionTranslationService contains methods and other services that
// help with interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductCollectionTranslationService] method instead.
type ProductCollectionTranslationService struct {
	Options []option.RequestOption
}

// NewProductCollectionTranslationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProductCollectionTranslationService(opts ...option.RequestOption) (r *ProductCollectionTranslationService) {
	r = &ProductCollectionTranslationService{}
	r.Options = opts
	return
}

// Creates or replaces the translation of one locale. The response is 201 for a new
// locale and 200 for a replaced locale.
func (r *ProductCollectionTranslationService) Update(ctx context.Context, id string, locale Locale, body ProductCollectionTranslationUpdateParams, opts ...option.RequestOption) (res *Translation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("product-collections/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Lists the translations, for all the locales or for the locales of the filter.
func (r *ProductCollectionTranslationService) List(ctx context.Context, id string, query ProductCollectionTranslationListParams, opts ...option.RequestOption) (res *TranslationsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("product-collections/%s/translations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the translation of one locale.
func (r *ProductCollectionTranslationService) Delete(ctx context.Context, id string, locale Locale, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("product-collections/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ProductCollectionTranslationUpdateParams struct {
	// The translated description.
	Description param.Field[string] `json:"description"`
	// The translated name.
	Name param.Field[string] `json:"name"`
}

func (r ProductCollectionTranslationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductCollectionTranslationListParams struct {
	// Returns only these locales, for example `?locale=fr&locale=de`. With no value,
	// the response has all the locales.
	Locale param.Field[[]string] `query:"locale"`
}

// URLQuery serializes [ProductCollectionTranslationListParams]'s query parameters
// as `url.Values`.
func (r ProductCollectionTranslationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
