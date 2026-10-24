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

// ProductTranslationService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductTranslationService] method instead.
type ProductTranslationService struct {
	Options []option.RequestOption
}

// NewProductTranslationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProductTranslationService(opts ...option.RequestOption) (r *ProductTranslationService) {
	r = &ProductTranslationService{}
	r.Options = opts
	return
}

// Creates or replaces the translation of one locale. The response is 201 for a new
// locale and 200 for a replaced locale.
func (r *ProductTranslationService) Update(ctx context.Context, id string, locale Locale, body ProductTranslationUpdateParams, opts ...option.RequestOption) (res *Translation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Lists the translations, for all the locales or for the locales of the filter.
func (r *ProductTranslationService) List(ctx context.Context, id string, query ProductTranslationListParams, opts ...option.RequestOption) (res *TranslationsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("products/%s/translations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the translation of one locale.
func (r *ProductTranslationService) Delete(ctx context.Context, id string, locale Locale, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("products/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// The translation of a product, an add-on, a product collection or a meter in one
// locale.
type Translation struct {
	// The locale of the translation.
	Locale Locale `json:"locale" api:"required"`
	// The time of the last change.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The translated description. `null` shows that the description has no
	// translation.
	Description string `json:"description" api:"nullable"`
	// The translated name. `null` shows that the name has no translation.
	Name string          `json:"name" api:"nullable"`
	JSON translationJSON `json:"-"`
}

// translationJSON contains the JSON metadata for the struct [Translation]
type translationJSON struct {
	Locale      apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Translation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translationJSON) RawJSON() string {
	return r.raw
}

// The translations of a product, an add-on, a product collection or a meter.
type TranslationsListResponse struct {
	Items []Translation                `json:"items" api:"required"`
	JSON  translationsListResponseJSON `json:"-"`
}

// translationsListResponseJSON contains the JSON metadata for the struct
// [TranslationsListResponse]
type translationsListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TranslationsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r translationsListResponseJSON) RawJSON() string {
	return r.raw
}

type ProductTranslationUpdateParams struct {
	// The translated description.
	Description param.Field[string] `json:"description"`
	// The translated name.
	Name param.Field[string] `json:"name"`
}

func (r ProductTranslationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProductTranslationListParams struct {
	// Returns only these locales, for example `?locale=fr&locale=de`. With no value,
	// the response has all the locales.
	Locale param.Field[[]string] `query:"locale"`
}

// URLQuery serializes [ProductTranslationListParams]'s query parameters as
// `url.Values`.
func (r ProductTranslationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
