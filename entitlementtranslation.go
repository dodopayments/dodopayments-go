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

// EntitlementTranslationService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitlementTranslationService] method instead.
type EntitlementTranslationService struct {
	Options []option.RequestOption
}

// NewEntitlementTranslationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntitlementTranslationService(opts ...option.RequestOption) (r *EntitlementTranslationService) {
	r = &EntitlementTranslationService{}
	r.Options = opts
	return
}

// Creates or replaces the translation of one locale. The response is 201 for a new
// locale and 200 for a replaced locale.
func (r *EntitlementTranslationService) Update(ctx context.Context, id string, locale Locale, body EntitlementTranslationUpdateParams, opts ...option.RequestOption) (res *EntitlementTranslation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entitlements/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Lists the translations, for all the locales or for the locales of the filter.
func (r *EntitlementTranslationService) List(ctx context.Context, id string, query EntitlementTranslationListParams, opts ...option.RequestOption) (res *EntitlementTranslationsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entitlements/%s/translations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the translation of one locale.
func (r *EntitlementTranslationService) Delete(ctx context.Context, id string, locale Locale, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("entitlements/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// The translation of an entitlement in one locale.
type EntitlementTranslation struct {
	// The locale of the translation.
	Locale Locale `json:"locale" api:"required"`
	// The time of the last change.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The translated activation message of a `license_key` entitlement.
	ActivationMessage string `json:"activation_message" api:"nullable"`
	// The translated description. `null` shows that the description has no
	// translation.
	Description string `json:"description" api:"nullable"`
	// The translated delivery instructions of a `digital_files` entitlement.
	Instructions string `json:"instructions" api:"nullable"`
	// The translated name. `null` shows that the name has no translation.
	Name string                     `json:"name" api:"nullable"`
	JSON entitlementTranslationJSON `json:"-"`
}

// entitlementTranslationJSON contains the JSON metadata for the struct
// [EntitlementTranslation]
type entitlementTranslationJSON struct {
	Locale            apijson.Field
	UpdatedAt         apijson.Field
	ActivationMessage apijson.Field
	Description       apijson.Field
	Instructions      apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EntitlementTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementTranslationJSON) RawJSON() string {
	return r.raw
}

// The translations of an entitlement.
type EntitlementTranslationsListResponse struct {
	Items []EntitlementTranslation                `json:"items" api:"required"`
	JSON  entitlementTranslationsListResponseJSON `json:"-"`
}

// entitlementTranslationsListResponseJSON contains the JSON metadata for the
// struct [EntitlementTranslationsListResponse]
type entitlementTranslationsListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementTranslationsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementTranslationsListResponseJSON) RawJSON() string {
	return r.raw
}

type EntitlementTranslationUpdateParams struct {
	// The translated activation message. Only a `license_key` entitlement has it.
	ActivationMessage param.Field[string] `json:"activation_message"`
	// The translated description.
	Description param.Field[string] `json:"description"`
	// The translated delivery instructions. Only a `digital_files` entitlement has
	// them.
	Instructions param.Field[string] `json:"instructions"`
	// The translated name.
	Name param.Field[string] `json:"name"`
}

func (r EntitlementTranslationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntitlementTranslationListParams struct {
	// Returns only these locales, for example `?locale=fr&locale=de`. With no value,
	// the response has all the locales.
	Locale param.Field[[]string] `query:"locale"`
}

// URLQuery serializes [EntitlementTranslationListParams]'s query parameters as
// `url.Values`.
func (r EntitlementTranslationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
