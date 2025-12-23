// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	shimjson "github.com/dodopayments/dodopayments-go/internal/encoding/json"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// CheckoutSessionService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCheckoutSessionService] method instead.
type CheckoutSessionService struct {
	Options []option.RequestOption
}

// NewCheckoutSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCheckoutSessionService(opts ...option.RequestOption) (r CheckoutSessionService) {
	r = CheckoutSessionService{}
	r.Options = opts
	return
}

func (r *CheckoutSessionService) New(ctx context.Context, body CheckoutSessionNewParams, opts ...option.RequestOption) (res *CheckoutSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "checkouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *CheckoutSessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CheckoutSessionStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("checkouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The property ProductCart is required.
type CheckoutSessionRequestParam struct {
	ProductCart  []CheckoutSessionRequestProductCartParam `json:"product_cart,omitzero,required"`
	DiscountCode param.Opt[string]                        `json:"discount_code,omitzero"`
	// Override merchant default 3DS behaviour for this session
	Force3DS param.Opt[bool] `json:"force_3ds,omitzero"`
	// The url to redirect after payment failure or success.
	ReturnURL param.Opt[string] `json:"return_url,omitzero"`
	// If confirm is true, all the details will be finalized. If required data is
	// missing, an API error is thrown.
	Confirm param.Opt[bool] `json:"confirm,omitzero"`
	// If true, only zipcode is required when confirm is true; other address fields
	// remain optional
	MinimalAddress param.Opt[bool] `json:"minimal_address,omitzero"`
	// If true, returns a shortened checkout URL. Defaults to false if not specified.
	ShortLink param.Opt[bool] `json:"short_link,omitzero"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Opt[bool] `json:"show_saved_payment_methods,omitzero"`
	// Customers will never see payment methods that are not in this list. However,
	// adding a method here does not guarantee customers will see it. Availability
	// still depends on other factors (e.g., customer location, merchant settings).
	//
	// Disclaimar: Always provide 'credit' and 'debit' as a fallback. If all payment
	// methods are unavailable, checkout session will fail.
	AllowedPaymentMethodTypes []PaymentMethodTypes `json:"allowed_payment_method_types,omitzero"`
	// Billing address information for the session
	BillingAddress CheckoutSessionRequestBillingAddressParam `json:"billing_address,omitzero"`
	// Additional metadata associated with the payment. Defaults to empty if not
	// provided.
	Metadata         map[string]string                           `json:"metadata,omitzero"`
	SubscriptionData CheckoutSessionRequestSubscriptionDataParam `json:"subscription_data,omitzero"`
	// This field is ingored if adaptive pricing is disabled
	//
	// Any of "AED", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM",
	// "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BWP",
	// "BYN", "BZD", "CAD", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK",
	// "DJF", "DKK", "DOP", "DZD", "EGP", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL",
	// "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF",
	// "IDR", "ILS", "INR", "IQD", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF",
	// "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD",
	// "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN",
	// "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN",
	// "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR",
	// "SBD", "SCR", "SEK", "SGD", "SHP", "SLE", "SLL", "SOS", "SRD", "SSP", "STN",
	// "SVC", "SZL", "THB", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX",
	// "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF",
	// "YER", "ZAR", "ZMW".
	BillingCurrency Currency `json:"billing_currency,omitzero"`
	// Customer details for the session
	Customer CustomerRequestUnionParam `json:"customer,omitzero"`
	// Customization for the checkout session page
	Customization CheckoutSessionRequestCustomizationParam `json:"customization,omitzero"`
	FeatureFlags  CheckoutSessionRequestFeatureFlagsParam  `json:"feature_flags,omitzero"`
	paramObj
}

func (r CheckoutSessionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSessionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSessionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ProductID, Quantity are required.
type CheckoutSessionRequestProductCartParam struct {
	// unique id of the product
	ProductID string `json:"product_id,required"`
	Quantity  int64  `json:"quantity,required"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`. Only applicable
	// for one time payments
	//
	// If amount is not set for pay_what_you_want product, customer is allowed to
	// select the amount.
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// only valid if product is a subscription
	Addons []AttachAddonParam `json:"addons,omitzero"`
	paramObj
}

func (r CheckoutSessionRequestProductCartParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSessionRequestProductCartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSessionRequestProductCartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing address information for the session
//
// The property Country is required.
type CheckoutSessionRequestBillingAddressParam struct {
	// Two-letter ISO country code (ISO 3166-1 alpha-2)
	//
	// Any of "AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM",
	// "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM",
	// "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "KH",
	// "CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG",
	// "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO",
	// "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK", "FO", "FJ", "FI", "FR", "GF",
	// "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU",
	// "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN",
	// "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE",
	// "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT",
	// "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU",
	// "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR",
	// "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MP", "NO", "OM", "PK",
	// "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE",
	// "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST",
	// "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS",
	// "SS", "ES", "LK", "SD", "SR", "SJ", "SZ", "SE", "CH", "SY", "TW", "TJ", "TZ",
	// "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA",
	// "AE", "GB", "UM", "US", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH",
	// "YE", "ZM", "ZW".
	Country CountryCode `json:"country,omitzero,required"`
	// City name
	City param.Opt[string] `json:"city,omitzero"`
	// State or province name
	State param.Opt[string] `json:"state,omitzero"`
	// Street address including house number and unit/apartment if applicable
	Street param.Opt[string] `json:"street,omitzero"`
	// Postal code or ZIP code
	Zipcode param.Opt[string] `json:"zipcode,omitzero"`
	paramObj
}

func (r CheckoutSessionRequestBillingAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSessionRequestBillingAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSessionRequestBillingAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customization for the checkout session page
type CheckoutSessionRequestCustomizationParam struct {
	// Force the checkout interface to render in a specific language (e.g. `en`, `es`)
	ForceLanguage param.Opt[string] `json:"force_language,omitzero"`
	// Show on demand tag
	//
	// Default is true
	ShowOnDemandTag param.Opt[bool] `json:"show_on_demand_tag,omitzero"`
	// Show order details by default
	//
	// Default is true
	ShowOrderDetails param.Opt[bool] `json:"show_order_details,omitzero"`
	// Theme of the page
	//
	// Default is `System`.
	//
	// Any of "dark", "light", "system".
	Theme string `json:"theme,omitzero"`
	paramObj
}

func (r CheckoutSessionRequestCustomizationParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSessionRequestCustomizationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSessionRequestCustomizationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CheckoutSessionRequestCustomizationParam](
		"theme", "dark", "light", "system",
	)
}

type CheckoutSessionRequestFeatureFlagsParam struct {
	// if customer is allowed to change currency, set it to true
	//
	// Default is true
	AllowCurrencySelection      param.Opt[bool] `json:"allow_currency_selection,omitzero"`
	AllowCustomerEditingCity    param.Opt[bool] `json:"allow_customer_editing_city,omitzero"`
	AllowCustomerEditingCountry param.Opt[bool] `json:"allow_customer_editing_country,omitzero"`
	AllowCustomerEditingEmail   param.Opt[bool] `json:"allow_customer_editing_email,omitzero"`
	AllowCustomerEditingName    param.Opt[bool] `json:"allow_customer_editing_name,omitzero"`
	AllowCustomerEditingState   param.Opt[bool] `json:"allow_customer_editing_state,omitzero"`
	AllowCustomerEditingStreet  param.Opt[bool] `json:"allow_customer_editing_street,omitzero"`
	AllowCustomerEditingZipcode param.Opt[bool] `json:"allow_customer_editing_zipcode,omitzero"`
	// If the customer is allowed to apply discount code, set it to true.
	//
	// Default is true
	AllowDiscountCode param.Opt[bool] `json:"allow_discount_code,omitzero"`
	// If phone number is collected from customer, set it to rue
	//
	// Default is true
	AllowPhoneNumberCollection param.Opt[bool] `json:"allow_phone_number_collection,omitzero"`
	// If the customer is allowed to add tax id, set it to true
	//
	// Default is true
	AllowTaxID param.Opt[bool] `json:"allow_tax_id,omitzero"`
	// Set to true if a new customer object should be created. By default email is used
	// to find an existing customer to attach the session to
	//
	// Default is false
	AlwaysCreateNewCustomer param.Opt[bool] `json:"always_create_new_customer,omitzero"`
	// If true, redirects the customer immediately after payment completion
	//
	// Default is false
	RedirectImmediately param.Opt[bool] `json:"redirect_immediately,omitzero"`
	paramObj
}

func (r CheckoutSessionRequestFeatureFlagsParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSessionRequestFeatureFlagsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSessionRequestFeatureFlagsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckoutSessionRequestSubscriptionDataParam struct {
	// Optional trial period in days If specified, this value overrides the trial
	// period set in the product's price Must be between 0 and 10000 days
	TrialPeriodDays param.Opt[int64]          `json:"trial_period_days,omitzero"`
	OnDemand        OnDemandSubscriptionParam `json:"on_demand,omitzero"`
	paramObj
}

func (r CheckoutSessionRequestSubscriptionDataParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSessionRequestSubscriptionDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSessionRequestSubscriptionDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckoutSessionResponse struct {
	// Checkout url
	CheckoutURL string `json:"checkout_url,required"`
	// The ID of the created checkout session
	SessionID string `json:"session_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckoutURL respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckoutSessionResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckoutSessionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckoutSessionStatus struct {
	// Id of the checkout session
	ID string `json:"id,required"`
	// Created at timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Customer email: prefers payment's customer, falls back to session
	CustomerEmail string `json:"customer_email,nullable"`
	// Customer name: prefers payment's customer, falls back to session
	CustomerName string `json:"customer_name,nullable"`
	// Id of the payment created by the checkout sessions.
	//
	// Null if checkout sessions is still at the details collection stage.
	PaymentID string `json:"payment_id,nullable"`
	// status of the payment.
	//
	// Null if checkout sessions is still at the details collection stage.
	//
	// Any of "succeeded", "failed", "cancelled", "processing",
	// "requires_customer_action", "requires_merchant_action",
	// "requires_payment_method", "requires_confirmation", "requires_capture",
	// "partially_captured", "partially_captured_and_capturable".
	PaymentStatus IntentStatus `json:"payment_status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		CustomerEmail respjson.Field
		CustomerName  respjson.Field
		PaymentID     respjson.Field
		PaymentStatus respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckoutSessionStatus) RawJSON() string { return r.JSON.raw }
func (r *CheckoutSessionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckoutSessionNewParams struct {
	CheckoutSessionRequest CheckoutSessionRequestParam
	paramObj
}

func (r CheckoutSessionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CheckoutSessionRequest)
}
func (r *CheckoutSessionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CheckoutSessionRequest)
}
