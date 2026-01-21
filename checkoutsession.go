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
	// Custom fields to collect from customer during checkout (max 5 fields)
	CustomFields param.Field[[]CheckoutSessionRequestCustomFieldParam] `json:"custom_fields"`
	// Customer details for the session
	Customer param.Field[CustomerRequestUnionParam] `json:"customer"`
	// Customization for the checkout session page
	Customization param.Field[CheckoutSessionRequestCustomizationParam] `json:"customization"`
	DiscountCode  param.Field[string]                                   `json:"discount_code"`
	FeatureFlags  param.Field[CheckoutSessionRequestFeatureFlagsParam]  `json:"feature_flags"`
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

// Definition of a custom field for checkout
type CheckoutSessionRequestCustomFieldParam struct {
	// Type of field determining validation rules
	FieldType param.Field[CheckoutSessionRequestCustomFieldsFieldType] `json:"field_type,required"`
	// Unique identifier for this field (used as key in responses)
	Key param.Field[string] `json:"key,required"`
	// Display label shown to customer
	Label param.Field[string] `json:"label,required"`
	// Options for dropdown type (required for dropdown, ignored for others)
	Options param.Field[[]string] `json:"options"`
	// Placeholder text for the input
	Placeholder param.Field[string] `json:"placeholder"`
	// Whether this field is required
	Required param.Field[bool] `json:"required"`
}

func (r CheckoutSessionRequestCustomFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of field determining validation rules
type CheckoutSessionRequestCustomFieldsFieldType string

const (
	CheckoutSessionRequestCustomFieldsFieldTypeText     CheckoutSessionRequestCustomFieldsFieldType = "text"
	CheckoutSessionRequestCustomFieldsFieldTypeNumber   CheckoutSessionRequestCustomFieldsFieldType = "number"
	CheckoutSessionRequestCustomFieldsFieldTypeEmail    CheckoutSessionRequestCustomFieldsFieldType = "email"
	CheckoutSessionRequestCustomFieldsFieldTypeURL      CheckoutSessionRequestCustomFieldsFieldType = "url"
	CheckoutSessionRequestCustomFieldsFieldTypePhone    CheckoutSessionRequestCustomFieldsFieldType = "phone"
	CheckoutSessionRequestCustomFieldsFieldTypeDate     CheckoutSessionRequestCustomFieldsFieldType = "date"
	CheckoutSessionRequestCustomFieldsFieldTypeDatetime CheckoutSessionRequestCustomFieldsFieldType = "datetime"
	CheckoutSessionRequestCustomFieldsFieldTypeDropdown CheckoutSessionRequestCustomFieldsFieldType = "dropdown"
	CheckoutSessionRequestCustomFieldsFieldTypeBoolean  CheckoutSessionRequestCustomFieldsFieldType = "boolean"
)

func (r CheckoutSessionRequestCustomFieldsFieldType) IsKnown() bool {
	switch r {
	case CheckoutSessionRequestCustomFieldsFieldTypeText, CheckoutSessionRequestCustomFieldsFieldTypeNumber, CheckoutSessionRequestCustomFieldsFieldTypeEmail, CheckoutSessionRequestCustomFieldsFieldTypeURL, CheckoutSessionRequestCustomFieldsFieldTypePhone, CheckoutSessionRequestCustomFieldsFieldTypeDate, CheckoutSessionRequestCustomFieldsFieldTypeDatetime, CheckoutSessionRequestCustomFieldsFieldTypeDropdown, CheckoutSessionRequestCustomFieldsFieldTypeBoolean:
		return true
	}
	return false
}

// Customization for the checkout session page
type CheckoutSessionRequestCustomizationParam struct {
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
	AllowCurrencySelection      param.Field[bool] `json:"allow_currency_selection"`
	AllowCustomerEditingCity    param.Field[bool] `json:"allow_customer_editing_city"`
	AllowCustomerEditingCountry param.Field[bool] `json:"allow_customer_editing_country"`
	AllowCustomerEditingEmail   param.Field[bool] `json:"allow_customer_editing_email"`
	AllowCustomerEditingName    param.Field[bool] `json:"allow_customer_editing_name"`
	AllowCustomerEditingState   param.Field[bool] `json:"allow_customer_editing_state"`
	AllowCustomerEditingStreet  param.Field[bool] `json:"allow_customer_editing_street"`
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
	// The ID of the created checkout session
	SessionID string `json:"session_id,required"`
	// Checkout url (None when payment_method_id is provided)
	CheckoutURL string                      `json:"checkout_url,nullable"`
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
	PaymentStatus IntentStatus              `json:"payment_status,nullable"`
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

// Data returned by the calculate checkout session API
type CheckoutSessionPreviewResponse struct {
	// Billing country
	BillingCountry CountryCode `json:"billing_country,required"`
	// Currency in which the calculations were made
	Currency Currency `json:"currency,required"`
	// Breakup of the current payment
	CurrentBreakup CheckoutSessionPreviewResponseCurrentBreakup `json:"current_breakup,required"`
	// The total product cart
	ProductCart []CheckoutSessionPreviewResponseProductCart `json:"product_cart,required"`
	// Total calculate price of the product cart
	TotalPrice int64 `json:"total_price,required"`
	// Breakup of recurring payments (None for one-time only)
	RecurringBreakup CheckoutSessionPreviewResponseRecurringBreakup `json:"recurring_breakup,nullable"`
	// Total tax
	TotalTax int64                              `json:"total_tax,nullable"`
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
	Discount int64 `json:"discount,required"`
	// Subtotal before discount (pre-tax original prices)
	Subtotal int64 `json:"subtotal,required"`
	// Total amount to be charged (final amount after all calculations)
	TotalAmount int64 `json:"total_amount,required"`
	// Total tax amount
	Tax  int64                                            `json:"tax,nullable"`
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
	// the currency in which the calculatiosn were made
	Currency Currency `json:"currency,required"`
	// discounted price
	DiscountedPrice int64 `json:"discounted_price,required"`
	// Whether this is a subscription product (affects tax calculation in breakup)
	IsSubscription bool                                             `json:"is_subscription,required"`
	IsUsageBased   bool                                             `json:"is_usage_based,required"`
	Meters         []CheckoutSessionPreviewResponseProductCartMeter `json:"meters,required"`
	// the product currency
	OgCurrency Currency `json:"og_currency,required"`
	// original price of the product
	OgPrice int64 `json:"og_price,required"`
	// unique id of the product
	ProductID string `json:"product_id,required"`
	// Quanitity
	Quantity int64 `json:"quantity,required"`
	// tax category
	TaxCategory TaxCategory `json:"tax_category,required"`
	// Whether tax is included in the price
	TaxInclusive bool `json:"tax_inclusive,required"`
	// tax rate
	TaxRate     int64                                            `json:"tax_rate,required"`
	Addons      []CheckoutSessionPreviewResponseProductCartAddon `json:"addons,nullable"`
	Description string                                           `json:"description,nullable"`
	// discount percentage
	DiscountAmount int64 `json:"discount_amount,nullable"`
	// number of cycles the discount will apply
	DiscountCycle int64 `json:"discount_cycle,nullable"`
	// name of the product
	Name string `json:"name,nullable"`
	// total tax
	Tax  int64                                         `json:"tax,nullable"`
	JSON checkoutSessionPreviewResponseProductCartJSON `json:"-"`
}

// checkoutSessionPreviewResponseProductCartJSON contains the JSON metadata for the
// struct [CheckoutSessionPreviewResponseProductCart]
type checkoutSessionPreviewResponseProductCartJSON struct {
	Currency        apijson.Field
	DiscountedPrice apijson.Field
	IsSubscription  apijson.Field
	IsUsageBased    apijson.Field
	Meters          apijson.Field
	OgCurrency      apijson.Field
	OgPrice         apijson.Field
	ProductID       apijson.Field
	Quantity        apijson.Field
	TaxCategory     apijson.Field
	TaxInclusive    apijson.Field
	TaxRate         apijson.Field
	Addons          apijson.Field
	Description     apijson.Field
	DiscountAmount  apijson.Field
	DiscountCycle   apijson.Field
	Name            apijson.Field
	Tax             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CheckoutSessionPreviewResponseProductCart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r checkoutSessionPreviewResponseProductCartJSON) RawJSON() string {
	return r.raw
}

type CheckoutSessionPreviewResponseProductCartMeter struct {
	MeasurementUnit string                                             `json:"measurement_unit,required"`
	Name            string                                             `json:"name,required"`
	PricePerUnit    string                                             `json:"price_per_unit,required"`
	Description     string                                             `json:"description,nullable"`
	FreeThreshold   int64                                              `json:"free_threshold,nullable"`
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
	AddonID         string   `json:"addon_id,required"`
	Currency        Currency `json:"currency,required"`
	DiscountedPrice int64    `json:"discounted_price,required"`
	Name            string   `json:"name,required"`
	OgCurrency      Currency `json:"og_currency,required"`
	OgPrice         int64    `json:"og_price,required"`
	Quantity        int64    `json:"quantity,required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory    TaxCategory                                        `json:"tax_category,required"`
	TaxInclusive   bool                                               `json:"tax_inclusive,required"`
	TaxRate        int64                                              `json:"tax_rate,required"`
	Description    string                                             `json:"description,nullable"`
	DiscountAmount int64                                              `json:"discount_amount,nullable"`
	Tax            int64                                              `json:"tax,nullable"`
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
	Discount int64 `json:"discount,required"`
	// Subtotal before discount (pre-tax original prices)
	Subtotal int64 `json:"subtotal,required"`
	// Total recurring amount including tax
	TotalAmount int64 `json:"total_amount,required"`
	// Total tax on recurring payments
	Tax  int64                                              `json:"tax,nullable"`
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
	CheckoutSessionRequest CheckoutSessionRequestParam `json:"checkout_session_request,required"`
}

func (r CheckoutSessionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CheckoutSessionRequest)
}

type CheckoutSessionPreviewParams struct {
	CheckoutSessionRequest CheckoutSessionRequestParam `json:"checkout_session_request,required"`
}

func (r CheckoutSessionPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CheckoutSessionRequest)
}
