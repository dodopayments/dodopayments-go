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

// CreditEntitlementTranslationService contains methods and other services that
// help with interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditEntitlementTranslationService] method instead.
type CreditEntitlementTranslationService struct {
	Options []option.RequestOption
}

// NewCreditEntitlementTranslationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCreditEntitlementTranslationService(opts ...option.RequestOption) (r *CreditEntitlementTranslationService) {
	r = &CreditEntitlementTranslationService{}
	r.Options = opts
	return
}

// Creates or replaces the translation of one locale. The response is 201 for a new
// locale and 200 for a replaced locale.
func (r *CreditEntitlementTranslationService) Update(ctx context.Context, id string, locale Locale, body CreditEntitlementTranslationUpdateParams, opts ...option.RequestOption) (res *CreditEntitlementTranslation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("credit-entitlements/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Lists the translations, for all the locales or for the locales of the filter.
func (r *CreditEntitlementTranslationService) List(ctx context.Context, id string, query CreditEntitlementTranslationListParams, opts ...option.RequestOption) (res *CreditEntitlementTranslationsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("credit-entitlements/%s/translations", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the translation of one locale.
func (r *CreditEntitlementTranslationService) Delete(ctx context.Context, id string, locale Locale, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("credit-entitlements/%s/translations/%v", id, locale)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// The translation of a credit entitlement in one locale.
type CreditEntitlementTranslation struct {
	// The locale of the translation.
	Locale Locale `json:"locale" api:"required"`
	// The time of the last change.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The translated name. `null` shows that the name has no translation.
	Name string `json:"name" api:"nullable"`
	// The translated unit. `null` shows that the unit has no translation.
	Unit string                           `json:"unit" api:"nullable"`
	JSON creditEntitlementTranslationJSON `json:"-"`
}

// creditEntitlementTranslationJSON contains the JSON metadata for the struct
// [CreditEntitlementTranslation]
type creditEntitlementTranslationJSON struct {
	Locale      apijson.Field
	UpdatedAt   apijson.Field
	Name        apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditEntitlementTranslation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditEntitlementTranslationJSON) RawJSON() string {
	return r.raw
}

// The translations of a credit entitlement.
type CreditEntitlementTranslationsListResponse struct {
	Items []CreditEntitlementTranslation                `json:"items" api:"required"`
	JSON  creditEntitlementTranslationsListResponseJSON `json:"-"`
}

// creditEntitlementTranslationsListResponseJSON contains the JSON metadata for the
// struct [CreditEntitlementTranslationsListResponse]
type creditEntitlementTranslationsListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditEntitlementTranslationsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditEntitlementTranslationsListResponseJSON) RawJSON() string {
	return r.raw
}

type CreditEntitlementTranslationUpdateParams struct {
	// The translated name.
	Name param.Field[string] `json:"name"`
	// The translated unit.
	Unit param.Field[string] `json:"unit"`
}

func (r CreditEntitlementTranslationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditEntitlementTranslationListParams struct {
	// Returns only these locales, for example `?locale=fr&locale=de`. With no value,
	// the response has all the locales.
	Locale param.Field[[]string] `query:"locale"`
}

// URLQuery serializes [CreditEntitlementTranslationListParams]'s query parameters
// as `url.Values`.
func (r CreditEntitlementTranslationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
