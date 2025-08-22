// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
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
func NewCheckoutSessionService(opts ...option.RequestOption) (r *CheckoutSessionService) {
	r = &CheckoutSessionService{}
	r.Options = opts
	return
}

func (r *CheckoutSessionService) New(ctx context.Context, body CheckoutSessionNewParams, opts ...option.RequestOption) (res *CheckoutSessionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "checkouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CheckoutSessionRequestParam struct {
	ProductCart param.Field[[]CheckoutSessionRequestProductCartParam] `json:"product_cart,required"`
	// Customers will never see payment methods that are not in this list. However,
	// adding a method here does not guarantee customers will see it. Availability
	// still depends on other factors (e.g., customer location, merchant settings).
	//
	// Disclaimar: Always provide 'credit' and 'debit' as a fallback. If all payment
	// methods are unavailable, checkout session will fail.
	AllowedPaymentMethodTypes param.Field[[]PaymentMethodTypes] `json:"allowed_payment_method_types"`
	// Billing address information for the session
	BillingAddress param.Field[CheckoutSessionRequestBillingAddressParam] `json:"billing_address"`
	// This field is ingored if adaptive pricing is disabled
	BillingCurrency param.Field[Currency] `json:"billing_currency"`
	// If confirm is true, all the details will be finalized. If required data is
	// missing, an API error is thrown.
	Confirm param.Field[bool] `json:"confirm"`
	// Customer details for the session
	Customer param.Field[CustomerRequestUnionParam] `json:"customer"`
	// Customization for the checkout session page
	Customization param.Field[CheckoutSessionRequestCustomizationParam] `json:"customization"`
	DiscountCode  param.Field[string]                                   `json:"discount_code"`
	FeatureFlags  param.Field[CheckoutSessionRequestFeatureFlagsParam]  `json:"feature_flags"`
	// Additional metadata associated with the payment. Defaults to empty if not
	// provided.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The url to redirect after payment failure or success.
	ReturnURL param.Field[string] `json:"return_url"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Field[bool]                                        `json:"show_saved_payment_methods"`
	SubscriptionData        param.Field[CheckoutSessionRequestSubscriptionDataParam] `json:"subscription_data"`
}

func (r CheckoutSessionRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckoutSessionRequestProductCartParam struct {
	// unique id of the product
	ProductID param.Field[string] `json:"product_id,required"`
	Quantity  param.Field[int64]  `json:"quantity,required"`
	// only valid if product is a subscription
	Addons param.Field[[]AttachAddonParam] `json:"addons"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`. Only applicable
	// for one time payments
	//
	// If amount is not set for pay_what_you_want product, customer is allowed to
	// select the amount.
	Amount param.Field[int64] `json:"amount"`
}

func (r CheckoutSessionRequestProductCartParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Billing address information for the session
type CheckoutSessionRequestBillingAddressParam struct {
	// Two-letter ISO country code (ISO 3166-1 alpha-2)
	Country param.Field[CountryCode] `json:"country,required"`
	// City name
	City param.Field[string] `json:"city"`
	// State or province name
	State param.Field[string] `json:"state"`
	// Street address including house number and unit/apartment if applicable
	Street param.Field[string] `json:"street"`
	// Postal code or ZIP code
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r CheckoutSessionRequestBillingAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Customization for the checkout session page
type CheckoutSessionRequestCustomizationParam struct {
	// Show on demand tag
	//
	// Default is true
	ShowOnDemandTag param.Field[bool] `json:"show_on_demand_tag"`
	// Show order details by default
	//
	// Default is true
	ShowOrderDetails param.Field[bool] `json:"show_order_details"`
	// Theme of the page
	//
	// Default is `System`.
	Theme param.Field[CheckoutSessionRequestCustomizationTheme] `json:"theme"`
}

func (r CheckoutSessionRequestCustomizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Theme of the page
//
// Default is `System`.
type CheckoutSessionRequestCustomizationTheme string

const (
	CheckoutSessionRequestCustomizationThemeDark   CheckoutSessionRequestCustomizationTheme = "dark"
	CheckoutSessionRequestCustomizationThemeLight  CheckoutSessionRequestCustomizationTheme = "light"
	CheckoutSessionRequestCustomizationThemeSystem CheckoutSessionRequestCustomizationTheme = "system"
)

func (r CheckoutSessionRequestCustomizationTheme) IsKnown() bool {
	switch r {
	case CheckoutSessionRequestCustomizationThemeDark, CheckoutSessionRequestCustomizationThemeLight, CheckoutSessionRequestCustomizationThemeSystem:
		return true
	}
	return false
}

type CheckoutSessionRequestFeatureFlagsParam struct {
	// if customer is allowed to change currency, set it to true
	//
	// Default is true
	AllowCurrencySelection param.Field[bool] `json:"allow_currency_selection"`
	// If the customer is allowed to apply discount code, set it to true.
	//
	// Default is true
	AllowDiscountCode param.Field[bool] `json:"allow_discount_code"`
	// If phone number is collected from customer, set it to rue
	//
	// Default is true
	AllowPhoneNumberCollection param.Field[bool] `json:"allow_phone_number_collection"`
	// If the customer is allowed to add tax id, set it to true
	//
	// Default is true
	AllowTaxID param.Field[bool] `json:"allow_tax_id"`
	// Set to true if a new customer object should be created. By default email is used
	// to find an existing customer to attach the session to
	//
	// Default is false
	AlwaysCreateNewCustomer param.Field[bool] `json:"always_create_new_customer"`
}

func (r CheckoutSessionRequestFeatureFlagsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckoutSessionRequestSubscriptionDataParam struct {
	OnDemand param.Field[OnDemandSubscriptionParam] `json:"on_demand"`
	// Optional trial period in days If specified, this value overrides the trial
	// period set in the product's price Must be between 0 and 10000 days
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r CheckoutSessionRequestSubscriptionDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckoutSessionResponse struct {
	// Checkout url
	CheckoutURL string `json:"checkout_url,required"`
	// The ID of the created checkout session
	SessionID string                      `json:"session_id,required"`
	JSON      checkoutSessionResponseJSON `json:"-"`
}

// checkoutSessionResponseJSON contains the JSON metadata for the struct
// [CheckoutSessionResponse]
type checkoutSessionResponseJSON struct {
	CheckoutURL apijson.Field
	SessionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckoutSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionResponseJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionNewParams struct {
	CheckoutSessionRequest CheckoutSessionRequestParam `json:"checkout_session_request,required"`
}

func (r CheckoutSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CheckoutSessionRequest)
}
