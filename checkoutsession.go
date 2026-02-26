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

func (r *CheckoutSessionService) Preview(ctx context.Context, body CheckoutSessionPreviewParams, opts ...option.RequestOption) (res *CheckoutSessionPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "checkouts/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CheckoutSessionBillingAddressParam struct {
	// Two-letter ISO country code (ISO 3166-1 alpha-2)
	Country param.Field[CountryCode] `json:"country" api:"required"`
	// City name
	City param.Field[string] `json:"city"`
	// State or province name
	State param.Field[string] `json:"state"`
	// Street address including house number and unit/apartment if applicable
	Street param.Field[string] `json:"street"`
	// Postal code or ZIP code
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r CheckoutSessionBillingAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckoutSessionCustomizationParam struct {
	// Force the checkout interface to render in a specific language (e.g. `en`, `es`)
	ForceLanguage param.Field[string] `json:"force_language"`
	// Show on demand tag
	//
	// Default is true
	ShowOnDemandTag param.Field[bool] `json:"show_on_demand_tag"`
	// Show order details by default
	//
	// Default is true
	ShowOrderDetails param.Field[bool] `json:"show_order_details"`
	// Theme of the page (determines which mode - light/dark/system - to use)
	//
	// Default is `System`.
	Theme param.Field[CheckoutSessionCustomizationTheme] `json:"theme"`
	// Optional custom theme configuration with colors for light and dark modes
	ThemeConfig param.Field[ThemeConfigParam] `json:"theme_config"`
}

func (r CheckoutSessionCustomizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Theme of the page (determines which mode - light/dark/system - to use)
//
// Default is `System`.
type CheckoutSessionCustomizationTheme string

const (
	CheckoutSessionCustomizationThemeDark   CheckoutSessionCustomizationTheme = "dark"
	CheckoutSessionCustomizationThemeLight  CheckoutSessionCustomizationTheme = "light"
	CheckoutSessionCustomizationThemeSystem CheckoutSessionCustomizationTheme = "system"
)

func (r CheckoutSessionCustomizationTheme) IsKnown() bool {
	switch r {
	case CheckoutSessionCustomizationThemeDark, CheckoutSessionCustomizationThemeLight, CheckoutSessionCustomizationThemeSystem:
		return true
	}
	return false
}

type CheckoutSessionFlagsParam struct {
	// if customer is allowed to change currency, set it to true
	//
	// Default is true
	AllowCurrencySelection      param.Field[bool] `json:"allow_currency_selection"`
	AllowCustomerEditingCity    param.Field[bool] `json:"allow_customer_editing_city"`
	AllowCustomerEditingCountry param.Field[bool] `json:"allow_customer_editing_country"`
	AllowCustomerEditingEmail   param.Field[bool] `json:"allow_customer_editing_email"`
	AllowCustomerEditingName    param.Field[bool] `json:"allow_customer_editing_name"`
	AllowCustomerEditingState   param.Field[bool] `json:"allow_customer_editing_state"`
	AllowCustomerEditingStreet  param.Field[bool] `json:"allow_customer_editing_street"`
	AllowCustomerEditingTaxID   param.Field[bool] `json:"allow_customer_editing_tax_id"`
	AllowCustomerEditingZipcode param.Field[bool] `json:"allow_customer_editing_zipcode"`
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
	// If true, redirects the customer immediately after payment completion
	//
	// Default is false
	RedirectImmediately param.Field[bool] `json:"redirect_immediately"`
}

func (r CheckoutSessionFlagsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckoutSessionRequestParam struct {
	ProductCart param.Field[[]ProductItemReqParam] `json:"product_cart" api:"required"`
	// Customers will never see payment methods that are not in this list. However,
	// adding a method here does not guarantee customers will see it. Availability
	// still depends on other factors (e.g., customer location, merchant settings).
	//
	// Disclaimar: Always provide 'credit' and 'debit' as a fallback. If all payment
	// methods are unavailable, checkout session will fail.
	AllowedPaymentMethodTypes param.Field[[]PaymentMethodTypes] `json:"allowed_payment_method_types"`
	// Billing address information for the session
	BillingAddress param.Field[CheckoutSessionBillingAddressParam] `json:"billing_address"`
	// This field is ingored if adaptive pricing is disabled
	BillingCurrency param.Field[Currency] `json:"billing_currency"`
	// If confirm is true, all the details will be finalized. If required data is
	// missing, an API error is thrown.
	Confirm param.Field[bool] `json:"confirm"`
	// Custom fields to collect from customer during checkout (max 5 fields)
	CustomFields param.Field[[]CustomFieldParam] `json:"custom_fields"`
	// Customer details for the session
	Customer param.Field[CustomerRequestUnionParam] `json:"customer"`
	// Customization for the checkout session page
	Customization param.Field[CheckoutSessionCustomizationParam] `json:"customization"`
	DiscountCode  param.Field[string]                            `json:"discount_code"`
	FeatureFlags  param.Field[CheckoutSessionFlagsParam]         `json:"feature_flags"`
	// Override merchant default 3DS behaviour for this session
	Force3DS param.Field[bool] `json:"force_3ds"`
	// Additional metadata associated with the payment. Defaults to empty if not
	// provided.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// If true, only zipcode is required when confirm is true; other address fields
	// remain optional
	MinimalAddress param.Field[bool] `json:"minimal_address"`
	// Optional payment method ID to use for this checkout session. Only allowed when
	// `confirm` is true. If provided, existing customer id must also be provided.
	PaymentMethodID param.Field[string] `json:"payment_method_id"`
	// Product collection ID for collection-based checkout flow
	ProductCollectionID param.Field[string] `json:"product_collection_id"`
	// The url to redirect after payment failure or success.
	ReturnURL param.Field[string] `json:"return_url"`
	// If true, returns a shortened checkout URL. Defaults to false if not specified.
	ShortLink param.Field[bool] `json:"short_link"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Field[bool]                  `json:"show_saved_payment_methods"`
	SubscriptionData        param.Field[SubscriptionDataParam] `json:"subscription_data"`
	// Tax ID for the customer (e.g. VAT number). Requires billing_address with
	// country.
	TaxID param.Field[string] `json:"tax_id"`
}

func (r CheckoutSessionRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CheckoutSessionResponse struct {
	// The ID of the created checkout session
	SessionID string `json:"session_id" api:"required"`
	// Checkout url (None when payment_method_id is provided)
	CheckoutURL string                      `json:"checkout_url" api:"nullable"`
	JSON        checkoutSessionResponseJSON `json:"-"`
}

// checkoutSessionResponseJSON contains the JSON metadata for the struct
// [CheckoutSessionResponse]
type checkoutSessionResponseJSON struct {
	SessionID   apijson.Field
	CheckoutURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckoutSessionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionResponseJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionStatus struct {
	// Id of the checkout session
	ID string `json:"id" api:"required"`
	// Created at timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Customer email: prefers payment's customer, falls back to session
	CustomerEmail string `json:"customer_email" api:"nullable"`
	// Customer name: prefers payment's customer, falls back to session
	CustomerName string `json:"customer_name" api:"nullable"`
	// Id of the payment created by the checkout sessions.
	//
	// Null if checkout sessions is still at the details collection stage.
	PaymentID string `json:"payment_id" api:"nullable"`
	// status of the payment.
	//
	// Null if checkout sessions is still at the details collection stage.
	PaymentStatus IntentStatus              `json:"payment_status" api:"nullable"`
	JSON          checkoutSessionStatusJSON `json:"-"`
}

// checkoutSessionStatusJSON contains the JSON metadata for the struct
// [CheckoutSessionStatus]
type checkoutSessionStatusJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	CustomerEmail apijson.Field
	CustomerName  apijson.Field
	PaymentID     apijson.Field
	PaymentStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CheckoutSessionStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionStatusJSON) RawJSON() string {
	return r.raw
}

// Definition of a custom field for checkout
type CustomFieldParam struct {
	// Type of field determining validation rules
	FieldType param.Field[CustomFieldFieldType] `json:"field_type" api:"required"`
	// Unique identifier for this field (used as key in responses)
	Key param.Field[string] `json:"key" api:"required"`
	// Display label shown to customer
	Label param.Field[string] `json:"label" api:"required"`
	// Options for dropdown type (required for dropdown, ignored for others)
	Options param.Field[[]string] `json:"options"`
	// Placeholder text for the input
	Placeholder param.Field[string] `json:"placeholder"`
	// Whether this field is required
	Required param.Field[bool] `json:"required"`
}

func (r CustomFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of field determining validation rules
type CustomFieldFieldType string

const (
	CustomFieldFieldTypeText     CustomFieldFieldType = "text"
	CustomFieldFieldTypeNumber   CustomFieldFieldType = "number"
	CustomFieldFieldTypeEmail    CustomFieldFieldType = "email"
	CustomFieldFieldTypeURL      CustomFieldFieldType = "url"
	CustomFieldFieldTypeDate     CustomFieldFieldType = "date"
	CustomFieldFieldTypeDropdown CustomFieldFieldType = "dropdown"
	CustomFieldFieldTypeBoolean  CustomFieldFieldType = "boolean"
)

func (r CustomFieldFieldType) IsKnown() bool {
	switch r {
	case CustomFieldFieldTypeText, CustomFieldFieldTypeNumber, CustomFieldFieldTypeEmail, CustomFieldFieldTypeURL, CustomFieldFieldTypeDate, CustomFieldFieldTypeDropdown, CustomFieldFieldTypeBoolean:
		return true
	}
	return false
}

type ProductItemReqParam struct {
	// unique id of the product
	ProductID param.Field[string] `json:"product_id" api:"required"`
	Quantity  param.Field[int64]  `json:"quantity" api:"required"`
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

func (r ProductItemReqParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionDataParam struct {
	OnDemand param.Field[OnDemandSubscriptionParam] `json:"on_demand"`
	// Optional trial period in days If specified, this value overrides the trial
	// period set in the product's price Must be between 0 and 10000 days
	TrialPeriodDays param.Field[int64] `json:"trial_period_days"`
}

func (r SubscriptionDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom theme configuration with colors for light and dark modes.
type ThemeConfigParam struct {
	// Dark mode color configuration
	Dark param.Field[ThemeModeConfigParam] `json:"dark"`
	// URL for the primary font. Must be a valid https:// URL.
	FontPrimaryURL param.Field[string] `json:"font_primary_url"`
	// URL for the secondary font. Must be a valid https:// URL.
	FontSecondaryURL param.Field[string] `json:"font_secondary_url"`
	// Font size for the checkout UI
	FontSize param.Field[ThemeConfigFontSize] `json:"font_size"`
	// Font weight for the checkout UI
	FontWeight param.Field[ThemeConfigFontWeight] `json:"font_weight"`
	// Light mode color configuration
	Light param.Field[ThemeModeConfigParam] `json:"light"`
	// Custom text for the pay button (e.g., "Complete Purchase", "Subscribe Now"). Max
	// 100 characters.
	PayButtonText param.Field[string] `json:"pay_button_text"`
	// Border radius for UI elements. Must be a number followed by px, rem, or em
	// (e.g., "4px", "0.5rem", "1em")
	Radius param.Field[string] `json:"radius"`
}

func (r ThemeConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Font size for the checkout UI
type ThemeConfigFontSize string

const (
	ThemeConfigFontSizeXs  ThemeConfigFontSize = "xs"
	ThemeConfigFontSizeSm  ThemeConfigFontSize = "sm"
	ThemeConfigFontSizeMd  ThemeConfigFontSize = "md"
	ThemeConfigFontSizeLg  ThemeConfigFontSize = "lg"
	ThemeConfigFontSizeXl  ThemeConfigFontSize = "xl"
	ThemeConfigFontSize2xl ThemeConfigFontSize = "2xl"
)

func (r ThemeConfigFontSize) IsKnown() bool {
	switch r {
	case ThemeConfigFontSizeXs, ThemeConfigFontSizeSm, ThemeConfigFontSizeMd, ThemeConfigFontSizeLg, ThemeConfigFontSizeXl, ThemeConfigFontSize2xl:
		return true
	}
	return false
}

// Font weight for the checkout UI
type ThemeConfigFontWeight string

const (
	ThemeConfigFontWeightNormal    ThemeConfigFontWeight = "normal"
	ThemeConfigFontWeightMedium    ThemeConfigFontWeight = "medium"
	ThemeConfigFontWeightBold      ThemeConfigFontWeight = "bold"
	ThemeConfigFontWeightExtraBold ThemeConfigFontWeight = "extraBold"
)

func (r ThemeConfigFontWeight) IsKnown() bool {
	switch r {
	case ThemeConfigFontWeightNormal, ThemeConfigFontWeightMedium, ThemeConfigFontWeightBold, ThemeConfigFontWeightExtraBold:
		return true
	}
	return false
}

// Color configuration for a single theme mode (light or dark).
//
// All color fields accept standard CSS color formats:
//
// - Hex: `#fff`, `#ffffff`, `#ffffffff` (with or without # prefix)
// - RGB/RGBA: `rgb(255, 255, 255)`, `rgba(255, 255, 255, 0.5)`
// - HSL/HSLA: `hsl(120, 100%, 50%)`, `hsla(120, 100%, 50%, 0.5)`
// - Named colors: `red`, `blue`, `transparent`, etc.
// - Advanced: `hwb()`, `lab()`, `lch()`, `oklab()`, `oklch()`, `color()`
type ThemeModeConfigParam struct {
	// Background primary color
	//
	// Examples: `"#ffffff"`, `"rgb(255, 255, 255)"`, `"white"`
	BgPrimary param.Field[string] `json:"bg_primary"`
	// Background secondary color
	BgSecondary param.Field[string] `json:"bg_secondary"`
	// Border primary color
	BorderPrimary param.Field[string] `json:"border_primary"`
	// Border secondary color
	BorderSecondary param.Field[string] `json:"border_secondary"`
	// Primary button background color
	ButtonPrimary param.Field[string] `json:"button_primary"`
	// Primary button hover color
	ButtonPrimaryHover param.Field[string] `json:"button_primary_hover"`
	// Secondary button background color
	ButtonSecondary param.Field[string] `json:"button_secondary"`
	// Secondary button hover color
	ButtonSecondaryHover param.Field[string] `json:"button_secondary_hover"`
	// Primary button text color
	ButtonTextPrimary param.Field[string] `json:"button_text_primary"`
	// Secondary button text color
	ButtonTextSecondary param.Field[string] `json:"button_text_secondary"`
	// Input focus border color
	InputFocusBorder param.Field[string] `json:"input_focus_border"`
	// Text error color
	TextError param.Field[string] `json:"text_error"`
	// Text placeholder color
	TextPlaceholder param.Field[string] `json:"text_placeholder"`
	// Text primary color
	TextPrimary param.Field[string] `json:"text_primary"`
	// Text secondary color
	TextSecondary param.Field[string] `json:"text_secondary"`
	// Text success color
	TextSuccess param.Field[string] `json:"text_success"`
}

func (r ThemeModeConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data returned by the calculate checkout session API
type CheckoutSessionPreviewResponse struct {
	// Billing country
	BillingCountry CountryCode `json:"billing_country" api:"required"`
	// Currency in which the calculations were made
	Currency Currency `json:"currency" api:"required"`
	// Breakup of the current payment
	CurrentBreakup CheckoutSessionPreviewResponseCurrentBreakup `json:"current_breakup" api:"required"`
	// The total product cart
	ProductCart []CheckoutSessionPreviewResponseProductCart `json:"product_cart" api:"required"`
	// Total calculate price of the product cart
	TotalPrice int64 `json:"total_price" api:"required"`
	// Breakup of recurring payments (None for one-time only)
	RecurringBreakup CheckoutSessionPreviewResponseRecurringBreakup `json:"recurring_breakup" api:"nullable"`
	// Error message if tax ID validation failed
	TaxIDErrMsg string `json:"tax_id_err_msg" api:"nullable"`
	// Total tax
	TotalTax int64                              `json:"total_tax" api:"nullable"`
	JSON     checkoutSessionPreviewResponseJSON `json:"-"`
}

// checkoutSessionPreviewResponseJSON contains the JSON metadata for the struct
// [CheckoutSessionPreviewResponse]
type checkoutSessionPreviewResponseJSON struct {
	BillingCountry   apijson.Field
	Currency         apijson.Field
	CurrentBreakup   apijson.Field
	ProductCart      apijson.Field
	TotalPrice       apijson.Field
	RecurringBreakup apijson.Field
	TaxIDErrMsg      apijson.Field
	TotalTax         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseJSON) RawJSON() string {
	return r.raw
}

// Breakup of the current payment
type CheckoutSessionPreviewResponseCurrentBreakup struct {
	// Total discount amount
	Discount int64 `json:"discount" api:"required"`
	// Subtotal before discount (pre-tax original prices)
	Subtotal int64 `json:"subtotal" api:"required"`
	// Total amount to be charged (final amount after all calculations)
	TotalAmount int64 `json:"total_amount" api:"required"`
	// Total tax amount
	Tax  int64                                            `json:"tax" api:"nullable"`
	JSON checkoutSessionPreviewResponseCurrentBreakupJSON `json:"-"`
}

// checkoutSessionPreviewResponseCurrentBreakupJSON contains the JSON metadata for
// the struct [CheckoutSessionPreviewResponseCurrentBreakup]
type checkoutSessionPreviewResponseCurrentBreakupJSON struct {
	Discount    apijson.Field
	Subtotal    apijson.Field
	TotalAmount apijson.Field
	Tax         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponseCurrentBreakup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseCurrentBreakupJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionPreviewResponseProductCart struct {
	// Credit entitlements that will be granted upon purchase
	CreditEntitlements []CheckoutSessionPreviewResponseProductCartCreditEntitlement `json:"credit_entitlements" api:"required"`
	// the currency in which the calculatiosn were made
	Currency Currency `json:"currency" api:"required"`
	// discounted price
	DiscountedPrice int64 `json:"discounted_price" api:"required"`
	// Whether this is a subscription product (affects tax calculation in breakup)
	IsSubscription bool                                             `json:"is_subscription" api:"required"`
	IsUsageBased   bool                                             `json:"is_usage_based" api:"required"`
	Meters         []CheckoutSessionPreviewResponseProductCartMeter `json:"meters" api:"required"`
	// the product currency
	OgCurrency Currency `json:"og_currency" api:"required"`
	// original price of the product
	OgPrice int64 `json:"og_price" api:"required"`
	// unique id of the product
	ProductID string `json:"product_id" api:"required"`
	// Quanitity
	Quantity int64 `json:"quantity" api:"required"`
	// tax category
	TaxCategory TaxCategory `json:"tax_category" api:"required"`
	// Whether tax is included in the price
	TaxInclusive bool `json:"tax_inclusive" api:"required"`
	// tax rate
	TaxRate     int64                                            `json:"tax_rate" api:"required"`
	Addons      []CheckoutSessionPreviewResponseProductCartAddon `json:"addons" api:"nullable"`
	Description string                                           `json:"description" api:"nullable"`
	// discount percentage
	DiscountAmount int64 `json:"discount_amount" api:"nullable"`
	// number of cycles the discount will apply
	DiscountCycle int64 `json:"discount_cycle" api:"nullable"`
	// name of the product
	Name string `json:"name" api:"nullable"`
	// total tax
	Tax  int64                                         `json:"tax" api:"nullable"`
	JSON checkoutSessionPreviewResponseProductCartJSON `json:"-"`
}

// checkoutSessionPreviewResponseProductCartJSON contains the JSON metadata for the
// struct [CheckoutSessionPreviewResponseProductCart]
type checkoutSessionPreviewResponseProductCartJSON struct {
	CreditEntitlements apijson.Field
	Currency           apijson.Field
	DiscountedPrice    apijson.Field
	IsSubscription     apijson.Field
	IsUsageBased       apijson.Field
	Meters             apijson.Field
	OgCurrency         apijson.Field
	OgPrice            apijson.Field
	ProductID          apijson.Field
	Quantity           apijson.Field
	TaxCategory        apijson.Field
	TaxInclusive       apijson.Field
	TaxRate            apijson.Field
	Addons             apijson.Field
	Description        apijson.Field
	DiscountAmount     apijson.Field
	DiscountCycle      apijson.Field
	Name               apijson.Field
	Tax                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponseProductCart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseProductCartJSON) RawJSON() string {
	return r.raw
}

// Minimal credit entitlement info shown at checkout — what credits the customer
// will receive
type CheckoutSessionPreviewResponseProductCartCreditEntitlement struct {
	// ID of the credit entitlement
	CreditEntitlementID string `json:"credit_entitlement_id" api:"required"`
	// Name of the credit entitlement
	CreditEntitlementName string `json:"credit_entitlement_name" api:"required"`
	// Unit label (e.g. "API Calls", "Tokens")
	CreditEntitlementUnit string `json:"credit_entitlement_unit" api:"required"`
	// Number of credits granted
	CreditsAmount string                                                         `json:"credits_amount" api:"required"`
	JSON          checkoutSessionPreviewResponseProductCartCreditEntitlementJSON `json:"-"`
}

// checkoutSessionPreviewResponseProductCartCreditEntitlementJSON contains the JSON
// metadata for the struct
// [CheckoutSessionPreviewResponseProductCartCreditEntitlement]
type checkoutSessionPreviewResponseProductCartCreditEntitlementJSON struct {
	CreditEntitlementID   apijson.Field
	CreditEntitlementName apijson.Field
	CreditEntitlementUnit apijson.Field
	CreditsAmount         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponseProductCartCreditEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseProductCartCreditEntitlementJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionPreviewResponseProductCartMeter struct {
	MeasurementUnit string                                             `json:"measurement_unit" api:"required"`
	Name            string                                             `json:"name" api:"required"`
	PricePerUnit    string                                             `json:"price_per_unit" api:"required"`
	Description     string                                             `json:"description" api:"nullable"`
	FreeThreshold   int64                                              `json:"free_threshold" api:"nullable"`
	JSON            checkoutSessionPreviewResponseProductCartMeterJSON `json:"-"`
}

// checkoutSessionPreviewResponseProductCartMeterJSON contains the JSON metadata
// for the struct [CheckoutSessionPreviewResponseProductCartMeter]
type checkoutSessionPreviewResponseProductCartMeterJSON struct {
	MeasurementUnit apijson.Field
	Name            apijson.Field
	PricePerUnit    apijson.Field
	Description     apijson.Field
	FreeThreshold   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponseProductCartMeter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseProductCartMeterJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionPreviewResponseProductCartAddon struct {
	AddonID         string   `json:"addon_id" api:"required"`
	Currency        Currency `json:"currency" api:"required"`
	DiscountedPrice int64    `json:"discounted_price" api:"required"`
	Name            string   `json:"name" api:"required"`
	OgCurrency      Currency `json:"og_currency" api:"required"`
	OgPrice         int64    `json:"og_price" api:"required"`
	Quantity        int64    `json:"quantity" api:"required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory    TaxCategory                                        `json:"tax_category" api:"required"`
	TaxInclusive   bool                                               `json:"tax_inclusive" api:"required"`
	TaxRate        int64                                              `json:"tax_rate" api:"required"`
	Description    string                                             `json:"description" api:"nullable"`
	DiscountAmount int64                                              `json:"discount_amount" api:"nullable"`
	Tax            int64                                              `json:"tax" api:"nullable"`
	JSON           checkoutSessionPreviewResponseProductCartAddonJSON `json:"-"`
}

// checkoutSessionPreviewResponseProductCartAddonJSON contains the JSON metadata
// for the struct [CheckoutSessionPreviewResponseProductCartAddon]
type checkoutSessionPreviewResponseProductCartAddonJSON struct {
	AddonID         apijson.Field
	Currency        apijson.Field
	DiscountedPrice apijson.Field
	Name            apijson.Field
	OgCurrency      apijson.Field
	OgPrice         apijson.Field
	Quantity        apijson.Field
	TaxCategory     apijson.Field
	TaxInclusive    apijson.Field
	TaxRate         apijson.Field
	Description     apijson.Field
	DiscountAmount  apijson.Field
	Tax             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponseProductCartAddon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseProductCartAddonJSON) RawJSON() string {
	return r.raw
}

// Breakup of recurring payments (None for one-time only)
type CheckoutSessionPreviewResponseRecurringBreakup struct {
	// Total discount amount
	Discount int64 `json:"discount" api:"required"`
	// Subtotal before discount (pre-tax original prices)
	Subtotal int64 `json:"subtotal" api:"required"`
	// Total recurring amount including tax
	TotalAmount int64 `json:"total_amount" api:"required"`
	// Total tax on recurring payments
	Tax  int64                                              `json:"tax" api:"nullable"`
	JSON checkoutSessionPreviewResponseRecurringBreakupJSON `json:"-"`
}

// checkoutSessionPreviewResponseRecurringBreakupJSON contains the JSON metadata
// for the struct [CheckoutSessionPreviewResponseRecurringBreakup]
type checkoutSessionPreviewResponseRecurringBreakupJSON struct {
	Discount    apijson.Field
	Subtotal    apijson.Field
	TotalAmount apijson.Field
	Tax         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponseRecurringBreakup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseRecurringBreakupJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionNewParams struct {
	CheckoutSessionRequest CheckoutSessionRequestParam `json:"checkout_session_request" api:"required"`
}

func (r CheckoutSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CheckoutSessionRequest)
}

type CheckoutSessionPreviewParams struct {
	CheckoutSessionRequest CheckoutSessionRequestParam `json:"checkout_session_request" api:"required"`
}

func (r CheckoutSessionPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CheckoutSessionRequest)
}
