// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
)

// PaymentService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r *PaymentService) {
	r = &PaymentService{}
	r.Options = opts
	return
}

func (r *PaymentService) New(ctx context.Context, body PaymentNewParams, opts ...option.RequestOption) (res *PaymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *PaymentService) Get(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	if paymentID == "" {
		err = errors.New("missing required payment_id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[PaymentListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "payments"
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

func (r *PaymentService) ListAutoPaging(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[PaymentListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type AttachExistingCustomerParam struct {
	CustomerID param.Field[string] `json:"customer_id,required"`
}

func (r AttachExistingCustomerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AttachExistingCustomerParam) implementsCustomerRequestUnionParam() {}

type BillingAddress struct {
	// City name
	City string `json:"city,required"`
	// ISO country code alpha2 variant
	Country CountryCode `json:"country,required"`
	// State or province name
	State string `json:"state,required"`
	// Street address including house number and unit/apartment if applicable
	Street string `json:"street,required"`
	// Postal code or ZIP code
	Zipcode string             `json:"zipcode,required"`
	JSON    billingAddressJSON `json:"-"`
}

// billingAddressJSON contains the JSON metadata for the struct [BillingAddress]
type billingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zipcode     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingAddressJSON) RawJSON() string {
	return r.raw
}

type BillingAddressParam struct {
	// City name
	City param.Field[string] `json:"city,required"`
	// ISO country code alpha2 variant
	Country param.Field[CountryCode] `json:"country,required"`
	// State or province name
	State param.Field[string] `json:"state,required"`
	// Street address including house number and unit/apartment if applicable
	Street param.Field[string] `json:"street,required"`
	// Postal code or ZIP code
	Zipcode param.Field[string] `json:"zipcode,required"`
}

func (r BillingAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreateNewCustomerParam struct {
	Email param.Field[string] `json:"email,required"`
	Name  param.Field[string] `json:"name,required"`
	// When false, the most recently created customer object with the given email is
	// used if exists. When true, a new customer object is always created False by
	// default
	CreateNewCustomer param.Field[bool]   `json:"create_new_customer"`
	PhoneNumber       param.Field[string] `json:"phone_number"`
}

func (r CreateNewCustomerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateNewCustomerParam) implementsCustomerRequestUnionParam() {}

type CustomerLimitedDetails struct {
	// Unique identifier for the customer
	CustomerID string `json:"customer_id,required"`
	// Email address of the customer
	Email string `json:"email,required"`
	// Full name of the customer
	Name string                     `json:"name,required"`
	JSON customerLimitedDetailsJSON `json:"-"`
}

// customerLimitedDetailsJSON contains the JSON metadata for the struct
// [CustomerLimitedDetails]
type customerLimitedDetailsJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerLimitedDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerLimitedDetailsJSON) RawJSON() string {
	return r.raw
}

type CustomerRequestParam struct {
	// When false, the most recently created customer object with the given email is
	// used if exists. When true, a new customer object is always created False by
	// default
	CreateNewCustomer param.Field[bool]   `json:"create_new_customer"`
	CustomerID        param.Field[string] `json:"customer_id"`
	Email             param.Field[string] `json:"email"`
	Name              param.Field[string] `json:"name"`
	PhoneNumber       param.Field[string] `json:"phone_number"`
}

func (r CustomerRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerRequestParam) implementsCustomerRequestUnionParam() {}

// Satisfied by [AttachExistingCustomerParam], [CreateNewCustomerParam],
// [CustomerRequestParam].
type CustomerRequestUnionParam interface {
	implementsCustomerRequestUnionParam()
}

type IntentStatus string

const (
	IntentStatusSucceeded                      IntentStatus = "succeeded"
	IntentStatusFailed                         IntentStatus = "failed"
	IntentStatusCancelled                      IntentStatus = "cancelled"
	IntentStatusProcessing                     IntentStatus = "processing"
	IntentStatusRequiresCustomerAction         IntentStatus = "requires_customer_action"
	IntentStatusRequiresMerchantAction         IntentStatus = "requires_merchant_action"
	IntentStatusRequiresPaymentMethod          IntentStatus = "requires_payment_method"
	IntentStatusRequiresConfirmation           IntentStatus = "requires_confirmation"
	IntentStatusRequiresCapture                IntentStatus = "requires_capture"
	IntentStatusPartiallyCaptured              IntentStatus = "partially_captured"
	IntentStatusPartiallyCapturedAndCapturable IntentStatus = "partially_captured_and_capturable"
)

func (r IntentStatus) IsKnown() bool {
	switch r {
	case IntentStatusSucceeded, IntentStatusFailed, IntentStatusCancelled, IntentStatusProcessing, IntentStatusRequiresCustomerAction, IntentStatusRequiresMerchantAction, IntentStatusRequiresPaymentMethod, IntentStatusRequiresConfirmation, IntentStatusRequiresCapture, IntentStatusPartiallyCaptured, IntentStatusPartiallyCapturedAndCapturable:
		return true
	}
	return false
}

type OneTimeProductCartItem struct {
	ProductID string `json:"product_id,required"`
	Quantity  int64  `json:"quantity,required"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Amount int64                      `json:"amount,nullable"`
	JSON   oneTimeProductCartItemJSON `json:"-"`
}

// oneTimeProductCartItemJSON contains the JSON metadata for the struct
// [OneTimeProductCartItem]
type oneTimeProductCartItemJSON struct {
	ProductID   apijson.Field
	Quantity    apijson.Field
	Amount      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OneTimeProductCartItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r oneTimeProductCartItemJSON) RawJSON() string {
	return r.raw
}

type OneTimeProductCartItemParam struct {
	ProductID param.Field[string] `json:"product_id,required"`
	Quantity  param.Field[int64]  `json:"quantity,required"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Amount param.Field[int64] `json:"amount"`
}

func (r OneTimeProductCartItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Payment struct {
	Billing BillingAddress `json:"billing,required"`
	// Identifier of the business associated with the payment
	BusinessID string `json:"business_id,required"`
	// Timestamp when the payment was created
	CreatedAt time.Time              `json:"created_at,required" format:"date-time"`
	Currency  Currency               `json:"currency,required"`
	Customer  CustomerLimitedDetails `json:"customer,required"`
	// List of disputes associated with this payment
	Disputes []Dispute         `json:"disputes,required"`
	Metadata map[string]string `json:"metadata,required"`
	// Unique identifier for the payment
	PaymentID string `json:"payment_id,required"`
	// List of refunds issued for this payment
	Refunds []Refund `json:"refunds,required"`
	// The amount that will be credited to your Dodo balance after currency conversion
	// and processing. Especially relevant for adaptive pricing where the customer's
	// payment currency differs from your settlement currency.
	SettlementAmount   int64    `json:"settlement_amount,required"`
	SettlementCurrency Currency `json:"settlement_currency,required"`
	// Total amount charged to the customer including tax, in smallest currency unit
	// (e.g. cents)
	TotalAmount int64 `json:"total_amount,required"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id,nullable"`
	// An error message if the payment failed
	ErrorMessage string `json:"error_message,nullable"`
	// Checkout URL
	PaymentLink string `json:"payment_link,nullable"`
	// Payment method used by customer (e.g. "card", "bank_transfer")
	PaymentMethod string `json:"payment_method,nullable"`
	// Specific type of payment method (e.g. "visa", "mastercard")
	PaymentMethodType string `json:"payment_method_type,nullable"`
	// List of products purchased in a one-time payment
	ProductCart []PaymentProductCart `json:"product_cart,nullable"`
	// This represents the portion of settlement_amount that corresponds to taxes
	// collected. Especially relevant for adaptive pricing where the tax component must
	// be tracked separately in your Dodo balance.
	SettlementTax int64        `json:"settlement_tax,nullable"`
	Status        IntentStatus `json:"status,nullable"`
	// Identifier of the subscription if payment is part of a subscription
	SubscriptionID string `json:"subscription_id,nullable"`
	// Amount of tax collected in smallest currency unit (e.g. cents)
	Tax int64 `json:"tax,nullable"`
	// Timestamp when the payment was last updated
	UpdatedAt time.Time   `json:"updated_at,nullable" format:"date-time"`
	JSON      paymentJSON `json:"-"`
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	Billing            apijson.Field
	BusinessID         apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Disputes           apijson.Field
	Metadata           apijson.Field
	PaymentID          apijson.Field
	Refunds            apijson.Field
	SettlementAmount   apijson.Field
	SettlementCurrency apijson.Field
	TotalAmount        apijson.Field
	DiscountID         apijson.Field
	ErrorMessage       apijson.Field
	PaymentLink        apijson.Field
	PaymentMethod      apijson.Field
	PaymentMethodType  apijson.Field
	ProductCart        apijson.Field
	SettlementTax      apijson.Field
	Status             apijson.Field
	SubscriptionID     apijson.Field
	Tax                apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Payment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentJSON) RawJSON() string {
	return r.raw
}

type PaymentProductCart struct {
	ProductID string                 `json:"product_id,required"`
	Quantity  int64                  `json:"quantity,required"`
	JSON      paymentProductCartJSON `json:"-"`
}

// paymentProductCartJSON contains the JSON metadata for the struct
// [PaymentProductCart]
type paymentProductCartJSON struct {
	ProductID   apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentProductCart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentProductCartJSON) RawJSON() string {
	return r.raw
}

type PaymentNewResponse struct {
	// Client secret used to load Dodo checkout SDK NOTE : Dodo checkout SDK will be
	// coming soon
	ClientSecret string                 `json:"client_secret,required"`
	Customer     CustomerLimitedDetails `json:"customer,required"`
	Metadata     map[string]string      `json:"metadata,required"`
	// Unique identifier for the payment
	PaymentID string `json:"payment_id,required"`
	// Total amount of the payment in smallest currency unit (e.g. cents)
	TotalAmount int64 `json:"total_amount,required"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id,nullable"`
	// Optional URL to a hosted payment page
	PaymentLink string `json:"payment_link,nullable"`
	// Optional list of products included in the payment
	ProductCart []OneTimeProductCartItem `json:"product_cart,nullable"`
	JSON        paymentNewResponseJSON   `json:"-"`
}

// paymentNewResponseJSON contains the JSON metadata for the struct
// [PaymentNewResponse]
type paymentNewResponseJSON struct {
	ClientSecret apijson.Field
	Customer     apijson.Field
	Metadata     apijson.Field
	PaymentID    apijson.Field
	TotalAmount  apijson.Field
	DiscountID   apijson.Field
	PaymentLink  apijson.Field
	ProductCart  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PaymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentNewResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponse struct {
	CreatedAt         time.Time               `json:"created_at,required" format:"date-time"`
	Currency          Currency                `json:"currency,required"`
	Customer          CustomerLimitedDetails  `json:"customer,required"`
	Metadata          map[string]string       `json:"metadata,required"`
	PaymentID         string                  `json:"payment_id,required"`
	TotalAmount       int64                   `json:"total_amount,required"`
	PaymentMethod     string                  `json:"payment_method,nullable"`
	PaymentMethodType string                  `json:"payment_method_type,nullable"`
	Status            IntentStatus            `json:"status,nullable"`
	SubscriptionID    string                  `json:"subscription_id,nullable"`
	JSON              paymentListResponseJSON `json:"-"`
}

// paymentListResponseJSON contains the JSON metadata for the struct
// [PaymentListResponse]
type paymentListResponseJSON struct {
	CreatedAt         apijson.Field
	Currency          apijson.Field
	Customer          apijson.Field
	Metadata          apijson.Field
	PaymentID         apijson.Field
	TotalAmount       apijson.Field
	PaymentMethod     apijson.Field
	PaymentMethodType apijson.Field
	Status            apijson.Field
	SubscriptionID    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PaymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentNewParams struct {
	Billing  param.Field[BillingAddressParam]       `json:"billing,required"`
	Customer param.Field[CustomerRequestUnionParam] `json:"customer,required"`
	// List of products in the cart. Must contain at least 1 and at most 100 items.
	ProductCart param.Field[[]OneTimeProductCartItemParam] `json:"product_cart,required"`
	// List of payment methods allowed during checkout.
	//
	// Customers will **never** see payment methods that are **not** in this list.
	// However, adding a method here **does not guarantee** customers will see it.
	// Availability still depends on other factors (e.g., customer location, merchant
	// settings).
	AllowedPaymentMethodTypes param.Field[[]PaymentNewParamsAllowedPaymentMethodType] `json:"allowed_payment_method_types"`
	BillingCurrency           param.Field[Currency]                                   `json:"billing_currency"`
	// Discount Code to apply to the transaction
	DiscountCode param.Field[string]            `json:"discount_code"`
	Metadata     param.Field[map[string]string] `json:"metadata"`
	// Whether to generate a payment link. Defaults to false if not specified.
	PaymentLink param.Field[bool] `json:"payment_link"`
	// Optional URL to redirect the customer after payment. Must be a valid URL if
	// provided.
	ReturnURL param.Field[string] `json:"return_url"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Field[bool] `json:"show_saved_payment_methods"`
	// Tax ID in case the payment is B2B. If tax id validation fails the payment
	// creation will fail
	TaxID param.Field[string] `json:"tax_id"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsAllowedPaymentMethodType string

const (
	PaymentNewParamsAllowedPaymentMethodTypeCredit           PaymentNewParamsAllowedPaymentMethodType = "credit"
	PaymentNewParamsAllowedPaymentMethodTypeDebit            PaymentNewParamsAllowedPaymentMethodType = "debit"
	PaymentNewParamsAllowedPaymentMethodTypeUpiCollect       PaymentNewParamsAllowedPaymentMethodType = "upi_collect"
	PaymentNewParamsAllowedPaymentMethodTypeUpiIntent        PaymentNewParamsAllowedPaymentMethodType = "upi_intent"
	PaymentNewParamsAllowedPaymentMethodTypeApplePay         PaymentNewParamsAllowedPaymentMethodType = "apple_pay"
	PaymentNewParamsAllowedPaymentMethodTypeCashapp          PaymentNewParamsAllowedPaymentMethodType = "cashapp"
	PaymentNewParamsAllowedPaymentMethodTypeGooglePay        PaymentNewParamsAllowedPaymentMethodType = "google_pay"
	PaymentNewParamsAllowedPaymentMethodTypeMultibanco       PaymentNewParamsAllowedPaymentMethodType = "multibanco"
	PaymentNewParamsAllowedPaymentMethodTypeBancontactCard   PaymentNewParamsAllowedPaymentMethodType = "bancontact_card"
	PaymentNewParamsAllowedPaymentMethodTypeEps              PaymentNewParamsAllowedPaymentMethodType = "eps"
	PaymentNewParamsAllowedPaymentMethodTypeIdeal            PaymentNewParamsAllowedPaymentMethodType = "ideal"
	PaymentNewParamsAllowedPaymentMethodTypePrzelewy24       PaymentNewParamsAllowedPaymentMethodType = "przelewy24"
	PaymentNewParamsAllowedPaymentMethodTypeAffirm           PaymentNewParamsAllowedPaymentMethodType = "affirm"
	PaymentNewParamsAllowedPaymentMethodTypeKlarna           PaymentNewParamsAllowedPaymentMethodType = "klarna"
	PaymentNewParamsAllowedPaymentMethodTypeSepa             PaymentNewParamsAllowedPaymentMethodType = "sepa"
	PaymentNewParamsAllowedPaymentMethodTypeACH              PaymentNewParamsAllowedPaymentMethodType = "ach"
	PaymentNewParamsAllowedPaymentMethodTypeAmazonPay        PaymentNewParamsAllowedPaymentMethodType = "amazon_pay"
	PaymentNewParamsAllowedPaymentMethodTypeAfterpayClearpay PaymentNewParamsAllowedPaymentMethodType = "afterpay_clearpay"
)

func (r PaymentNewParamsAllowedPaymentMethodType) IsKnown() bool {
	switch r {
	case PaymentNewParamsAllowedPaymentMethodTypeCredit, PaymentNewParamsAllowedPaymentMethodTypeDebit, PaymentNewParamsAllowedPaymentMethodTypeUpiCollect, PaymentNewParamsAllowedPaymentMethodTypeUpiIntent, PaymentNewParamsAllowedPaymentMethodTypeApplePay, PaymentNewParamsAllowedPaymentMethodTypeCashapp, PaymentNewParamsAllowedPaymentMethodTypeGooglePay, PaymentNewParamsAllowedPaymentMethodTypeMultibanco, PaymentNewParamsAllowedPaymentMethodTypeBancontactCard, PaymentNewParamsAllowedPaymentMethodTypeEps, PaymentNewParamsAllowedPaymentMethodTypeIdeal, PaymentNewParamsAllowedPaymentMethodTypePrzelewy24, PaymentNewParamsAllowedPaymentMethodTypeAffirm, PaymentNewParamsAllowedPaymentMethodTypeKlarna, PaymentNewParamsAllowedPaymentMethodTypeSepa, PaymentNewParamsAllowedPaymentMethodTypeACH, PaymentNewParamsAllowedPaymentMethodTypeAmazonPay, PaymentNewParamsAllowedPaymentMethodTypeAfterpayClearpay:
		return true
	}
	return false
}

type PaymentListParams struct {
	// Get events after this created time
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get events created before this time
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Filter by customer id
	CustomerID param.Field[string] `query:"customer_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by status
	Status param.Field[IntentStatus] `query:"status"`
	// Filter by subscription id
	SubscriptionID param.Field[string] `query:"subscription_id"`
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
