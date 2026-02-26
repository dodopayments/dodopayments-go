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

// Deprecated: deprecated
func (r *PaymentService) New(ctx context.Context, body PaymentNewParams, opts ...option.RequestOption) (res *PaymentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *PaymentService) Get(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *Payment, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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

func (r *PaymentService) GetLineItems(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *PaymentGetLineItemsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentID == "" {
		err = errors.New("missing required payment_id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s/line-items", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AttachExistingCustomerParam struct {
	CustomerID param.Field[string] `json:"customer_id" api:"required"`
}

func (r AttachExistingCustomerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AttachExistingCustomerParam) implementsCustomerRequestUnionParam() {}

type BillingAddress struct {
	// Two-letter ISO country code (ISO 3166-1 alpha-2)
	Country CountryCode `json:"country" api:"required"`
	// City name
	City string `json:"city" api:"nullable"`
	// State or province name
	State string `json:"state" api:"nullable"`
	// Street address including house number and unit/apartment if applicable
	Street string `json:"street" api:"nullable"`
	// Postal code or ZIP code
	Zipcode string             `json:"zipcode" api:"nullable"`
	JSON    billingAddressJSON `json:"-"`
}

// billingAddressJSON contains the JSON metadata for the struct [BillingAddress]
type billingAddressJSON struct {
	Country     apijson.Field
	City        apijson.Field
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

func (r BillingAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerLimitedDetails struct {
	// Unique identifier for the customer
	CustomerID string `json:"customer_id" api:"required"`
	// Email address of the customer
	Email string `json:"email" api:"required"`
	// Full name of the customer
	Name string `json:"name" api:"required"`
	// Additional metadata associated with the customer
	Metadata map[string]string `json:"metadata"`
	// Phone number of the customer
	PhoneNumber string                     `json:"phone_number" api:"nullable"`
	JSON        customerLimitedDetailsJSON `json:"-"`
}

// customerLimitedDetailsJSON contains the JSON metadata for the struct
// [CustomerLimitedDetails]
type customerLimitedDetailsJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Metadata    apijson.Field
	PhoneNumber apijson.Field
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
	CustomerID param.Field[string] `json:"customer_id"`
	// Email is required for creating a new customer
	Email param.Field[string] `json:"email"`
	// Optional full name of the customer. If provided during session creation, it is
	// persisted and becomes immutable for the session. If omitted here, it can be
	// provided later via the confirm API.
	Name        param.Field[string] `json:"name"`
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r CustomerRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CustomerRequestParam) implementsCustomerRequestUnionParam() {}

// Satisfied by [AttachExistingCustomerParam], [NewCustomerParam],
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

type NewCustomerParam struct {
	// Email is required for creating a new customer
	Email param.Field[string] `json:"email" api:"required"`
	// Optional full name of the customer. If provided during session creation, it is
	// persisted and becomes immutable for the session. If omitted here, it can be
	// provided later via the confirm API.
	Name        param.Field[string] `json:"name"`
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r NewCustomerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewCustomerParam) implementsCustomerRequestUnionParam() {}

type OneTimeProductCartItem struct {
	ProductID string `json:"product_id" api:"required"`
	Quantity  int64  `json:"quantity" api:"required"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Amount int64                      `json:"amount" api:"nullable"`
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
	ProductID param.Field[string] `json:"product_id" api:"required"`
	Quantity  param.Field[int64]  `json:"quantity" api:"required"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Amount param.Field[int64] `json:"amount"`
}

func (r OneTimeProductCartItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Payment struct {
	// Billing address details for payments
	Billing BillingAddress `json:"billing" api:"required"`
	// brand id this payment belongs to
	BrandID string `json:"brand_id" api:"required"`
	// Identifier of the business associated with the payment
	BusinessID string `json:"business_id" api:"required"`
	// Timestamp when the payment was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Currency used for the payment
	Currency Currency `json:"currency" api:"required"`
	// Details about the customer who made the payment
	Customer CustomerLimitedDetails `json:"customer" api:"required"`
	// brand id this payment belongs to
	DigitalProductsDelivered bool `json:"digital_products_delivered" api:"required"`
	// List of disputes associated with this payment
	Disputes []Dispute `json:"disputes" api:"required"`
	// Additional custom data associated with the payment
	Metadata map[string]string `json:"metadata" api:"required"`
	// Unique identifier for the payment
	PaymentID string `json:"payment_id" api:"required"`
	// List of refunds issued for this payment
	Refunds []PaymentRefund `json:"refunds" api:"required"`
	// The amount that will be credited to your Dodo balance after currency conversion
	// and processing. Especially relevant for adaptive pricing where the customer's
	// payment currency differs from your settlement currency.
	SettlementAmount int64 `json:"settlement_amount" api:"required"`
	// The currency in which the settlement_amount will be credited to your Dodo
	// balance. This may differ from the customer's payment currency in adaptive
	// pricing scenarios.
	SettlementCurrency Currency `json:"settlement_currency" api:"required"`
	// Total amount charged to the customer including tax, in smallest currency unit
	// (e.g. cents)
	TotalAmount int64 `json:"total_amount" api:"required"`
	// Cardholder name
	CardHolderName string `json:"card_holder_name" api:"nullable"`
	// ISO2 country code of the card
	CardIssuingCountry CountryCode `json:"card_issuing_country" api:"nullable"`
	// The last four digits of the card
	CardLastFour string `json:"card_last_four" api:"nullable"`
	// Card network like VISA, MASTERCARD etc.
	CardNetwork string `json:"card_network" api:"nullable"`
	// The type of card DEBIT or CREDIT
	CardType string `json:"card_type" api:"nullable"`
	// If payment is made using a checkout session, this field is set to the id of the
	// session.
	CheckoutSessionID string `json:"checkout_session_id" api:"nullable"`
	// Customer's responses to custom fields collected during checkout
	CustomFieldResponses []PaymentCustomFieldResponse `json:"custom_field_responses" api:"nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id" api:"nullable"`
	// An error code if the payment failed
	ErrorCode string `json:"error_code" api:"nullable"`
	// An error message if the payment failed
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Invoice ID for this payment. Uses India-specific invoice ID if available.
	InvoiceID string `json:"invoice_id" api:"nullable"`
	// URL to download the invoice PDF for this payment.
	InvoiceURL string `json:"invoice_url" api:"nullable"`
	// Checkout URL
	PaymentLink string `json:"payment_link" api:"nullable"`
	// Payment method used by customer (e.g. "card", "bank_transfer")
	PaymentMethod string `json:"payment_method" api:"nullable"`
	// Specific type of payment method (e.g. "visa", "mastercard")
	PaymentMethodType string `json:"payment_method_type" api:"nullable"`
	// List of products purchased in a one-time payment
	ProductCart []PaymentProductCart `json:"product_cart" api:"nullable"`
	// Summary of the refund status for this payment. None if no succeeded refunds
	// exist.
	RefundStatus PaymentRefundStatus `json:"refund_status" api:"nullable"`
	// This represents the portion of settlement_amount that corresponds to taxes
	// collected. Especially relevant for adaptive pricing where the tax component must
	// be tracked separately in your Dodo balance.
	SettlementTax int64 `json:"settlement_tax" api:"nullable"`
	// Current status of the payment intent
	Status IntentStatus `json:"status" api:"nullable"`
	// Identifier of the subscription if payment is part of a subscription
	SubscriptionID string `json:"subscription_id" api:"nullable"`
	// Amount of tax collected in smallest currency unit (e.g. cents)
	Tax int64 `json:"tax" api:"nullable"`
	// Timestamp when the payment was last updated
	UpdatedAt time.Time   `json:"updated_at" api:"nullable" format:"date-time"`
	JSON      paymentJSON `json:"-"`
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	Billing                  apijson.Field
	BrandID                  apijson.Field
	BusinessID               apijson.Field
	CreatedAt                apijson.Field
	Currency                 apijson.Field
	Customer                 apijson.Field
	DigitalProductsDelivered apijson.Field
	Disputes                 apijson.Field
	Metadata                 apijson.Field
	PaymentID                apijson.Field
	Refunds                  apijson.Field
	SettlementAmount         apijson.Field
	SettlementCurrency       apijson.Field
	TotalAmount              apijson.Field
	CardHolderName           apijson.Field
	CardIssuingCountry       apijson.Field
	CardLastFour             apijson.Field
	CardNetwork              apijson.Field
	CardType                 apijson.Field
	CheckoutSessionID        apijson.Field
	CustomFieldResponses     apijson.Field
	DiscountID               apijson.Field
	ErrorCode                apijson.Field
	ErrorMessage             apijson.Field
	InvoiceID                apijson.Field
	InvoiceURL               apijson.Field
	PaymentLink              apijson.Field
	PaymentMethod            apijson.Field
	PaymentMethodType        apijson.Field
	ProductCart              apijson.Field
	RefundStatus             apijson.Field
	SettlementTax            apijson.Field
	Status                   apijson.Field
	SubscriptionID           apijson.Field
	Tax                      apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Payment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentJSON) RawJSON() string {
	return r.raw
}

type PaymentRefund struct {
	// The unique identifier of the business issuing the refund.
	BusinessID string `json:"business_id" api:"required"`
	// The timestamp of when the refund was created in UTC.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// If true the refund is a partial refund
	IsPartial bool `json:"is_partial" api:"required"`
	// The unique identifier of the payment associated with the refund.
	PaymentID string `json:"payment_id" api:"required"`
	// The unique identifier of the refund.
	RefundID string `json:"refund_id" api:"required"`
	// The current status of the refund.
	Status RefundStatus `json:"status" api:"required"`
	// The refunded amount.
	Amount int64 `json:"amount" api:"nullable"`
	// The currency of the refund, represented as an ISO 4217 currency code.
	Currency Currency `json:"currency" api:"nullable"`
	// The reason provided for the refund, if any. Optional.
	Reason string            `json:"reason" api:"nullable"`
	JSON   paymentRefundJSON `json:"-"`
}

// paymentRefundJSON contains the JSON metadata for the struct [PaymentRefund]
type paymentRefundJSON struct {
	BusinessID  apijson.Field
	CreatedAt   apijson.Field
	IsPartial   apijson.Field
	PaymentID   apijson.Field
	RefundID    apijson.Field
	Status      apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentRefund) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentRefundJSON) RawJSON() string {
	return r.raw
}

// Customer's response to a custom field
type PaymentCustomFieldResponse struct {
	// Key matching the custom field definition
	Key string `json:"key" api:"required"`
	// Value provided by customer
	Value string                         `json:"value" api:"required"`
	JSON  paymentCustomFieldResponseJSON `json:"-"`
}

// paymentCustomFieldResponseJSON contains the JSON metadata for the struct
// [PaymentCustomFieldResponse]
type paymentCustomFieldResponseJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentCustomFieldResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentCustomFieldResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentProductCart struct {
	ProductID string                 `json:"product_id" api:"required"`
	Quantity  int64                  `json:"quantity" api:"required"`
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

// Summary of the refund status for this payment. None if no succeeded refunds
// exist.
type PaymentRefundStatus string

const (
	PaymentRefundStatusPartial PaymentRefundStatus = "partial"
	PaymentRefundStatusFull    PaymentRefundStatus = "full"
)

func (r PaymentRefundStatus) IsKnown() bool {
	switch r {
	case PaymentRefundStatusPartial, PaymentRefundStatusFull:
		return true
	}
	return false
}

type PaymentMethodTypes string

const (
	PaymentMethodTypesACH                        PaymentMethodTypes = "ach"
	PaymentMethodTypesAffirm                     PaymentMethodTypes = "affirm"
	PaymentMethodTypesAfterpayClearpay           PaymentMethodTypes = "afterpay_clearpay"
	PaymentMethodTypesAlfamart                   PaymentMethodTypes = "alfamart"
	PaymentMethodTypesAliPay                     PaymentMethodTypes = "ali_pay"
	PaymentMethodTypesAliPayHk                   PaymentMethodTypes = "ali_pay_hk"
	PaymentMethodTypesAlma                       PaymentMethodTypes = "alma"
	PaymentMethodTypesAmazonPay                  PaymentMethodTypes = "amazon_pay"
	PaymentMethodTypesApplePay                   PaymentMethodTypes = "apple_pay"
	PaymentMethodTypesAtome                      PaymentMethodTypes = "atome"
	PaymentMethodTypesBacs                       PaymentMethodTypes = "bacs"
	PaymentMethodTypesBancontactCard             PaymentMethodTypes = "bancontact_card"
	PaymentMethodTypesBecs                       PaymentMethodTypes = "becs"
	PaymentMethodTypesBenefit                    PaymentMethodTypes = "benefit"
	PaymentMethodTypesBizum                      PaymentMethodTypes = "bizum"
	PaymentMethodTypesBlik                       PaymentMethodTypes = "blik"
	PaymentMethodTypesBoleto                     PaymentMethodTypes = "boleto"
	PaymentMethodTypesBcaBankTransfer            PaymentMethodTypes = "bca_bank_transfer"
	PaymentMethodTypesBniVa                      PaymentMethodTypes = "bni_va"
	PaymentMethodTypesBriVa                      PaymentMethodTypes = "bri_va"
	PaymentMethodTypesCardRedirect               PaymentMethodTypes = "card_redirect"
	PaymentMethodTypesCimbVa                     PaymentMethodTypes = "cimb_va"
	PaymentMethodTypesClassic                    PaymentMethodTypes = "classic"
	PaymentMethodTypesCredit                     PaymentMethodTypes = "credit"
	PaymentMethodTypesCryptoCurrency             PaymentMethodTypes = "crypto_currency"
	PaymentMethodTypesCashapp                    PaymentMethodTypes = "cashapp"
	PaymentMethodTypesDana                       PaymentMethodTypes = "dana"
	PaymentMethodTypesDanamonVa                  PaymentMethodTypes = "danamon_va"
	PaymentMethodTypesDebit                      PaymentMethodTypes = "debit"
	PaymentMethodTypesDuitNow                    PaymentMethodTypes = "duit_now"
	PaymentMethodTypesEfecty                     PaymentMethodTypes = "efecty"
	PaymentMethodTypesEft                        PaymentMethodTypes = "eft"
	PaymentMethodTypesEps                        PaymentMethodTypes = "eps"
	PaymentMethodTypesFps                        PaymentMethodTypes = "fps"
	PaymentMethodTypesEvoucher                   PaymentMethodTypes = "evoucher"
	PaymentMethodTypesGiropay                    PaymentMethodTypes = "giropay"
	PaymentMethodTypesGivex                      PaymentMethodTypes = "givex"
	PaymentMethodTypesGooglePay                  PaymentMethodTypes = "google_pay"
	PaymentMethodTypesGoPay                      PaymentMethodTypes = "go_pay"
	PaymentMethodTypesGcash                      PaymentMethodTypes = "gcash"
	PaymentMethodTypesIdeal                      PaymentMethodTypes = "ideal"
	PaymentMethodTypesInterac                    PaymentMethodTypes = "interac"
	PaymentMethodTypesIndomaret                  PaymentMethodTypes = "indomaret"
	PaymentMethodTypesKlarna                     PaymentMethodTypes = "klarna"
	PaymentMethodTypesKakaoPay                   PaymentMethodTypes = "kakao_pay"
	PaymentMethodTypesLocalBankRedirect          PaymentMethodTypes = "local_bank_redirect"
	PaymentMethodTypesMandiriVa                  PaymentMethodTypes = "mandiri_va"
	PaymentMethodTypesKnet                       PaymentMethodTypes = "knet"
	PaymentMethodTypesMBWay                      PaymentMethodTypes = "mb_way"
	PaymentMethodTypesMobilePay                  PaymentMethodTypes = "mobile_pay"
	PaymentMethodTypesMomo                       PaymentMethodTypes = "momo"
	PaymentMethodTypesMomoAtm                    PaymentMethodTypes = "momo_atm"
	PaymentMethodTypesMultibanco                 PaymentMethodTypes = "multibanco"
	PaymentMethodTypesOnlineBankingThailand      PaymentMethodTypes = "online_banking_thailand"
	PaymentMethodTypesOnlineBankingCzechRepublic PaymentMethodTypes = "online_banking_czech_republic"
	PaymentMethodTypesOnlineBankingFinland       PaymentMethodTypes = "online_banking_finland"
	PaymentMethodTypesOnlineBankingFpx           PaymentMethodTypes = "online_banking_fpx"
	PaymentMethodTypesOnlineBankingPoland        PaymentMethodTypes = "online_banking_poland"
	PaymentMethodTypesOnlineBankingSlovakia      PaymentMethodTypes = "online_banking_slovakia"
	PaymentMethodTypesOxxo                       PaymentMethodTypes = "oxxo"
	PaymentMethodTypesPagoEfectivo               PaymentMethodTypes = "pago_efectivo"
	PaymentMethodTypesPermataBankTransfer        PaymentMethodTypes = "permata_bank_transfer"
	PaymentMethodTypesOpenBankingUk              PaymentMethodTypes = "open_banking_uk"
	PaymentMethodTypesPayBright                  PaymentMethodTypes = "pay_bright"
	PaymentMethodTypesPaypal                     PaymentMethodTypes = "paypal"
	PaymentMethodTypesPaze                       PaymentMethodTypes = "paze"
	PaymentMethodTypesPix                        PaymentMethodTypes = "pix"
	PaymentMethodTypesPaySafeCard                PaymentMethodTypes = "pay_safe_card"
	PaymentMethodTypesPrzelewy24                 PaymentMethodTypes = "przelewy24"
	PaymentMethodTypesPromptPay                  PaymentMethodTypes = "prompt_pay"
	PaymentMethodTypesPse                        PaymentMethodTypes = "pse"
	PaymentMethodTypesRedCompra                  PaymentMethodTypes = "red_compra"
	PaymentMethodTypesRedPagos                   PaymentMethodTypes = "red_pagos"
	PaymentMethodTypesSamsungPay                 PaymentMethodTypes = "samsung_pay"
	PaymentMethodTypesSepa                       PaymentMethodTypes = "sepa"
	PaymentMethodTypesSepaBankTransfer           PaymentMethodTypes = "sepa_bank_transfer"
	PaymentMethodTypesSofort                     PaymentMethodTypes = "sofort"
	PaymentMethodTypesSwish                      PaymentMethodTypes = "swish"
	PaymentMethodTypesTouchNGo                   PaymentMethodTypes = "touch_n_go"
	PaymentMethodTypesTrustly                    PaymentMethodTypes = "trustly"
	PaymentMethodTypesTwint                      PaymentMethodTypes = "twint"
	PaymentMethodTypesUpiCollect                 PaymentMethodTypes = "upi_collect"
	PaymentMethodTypesUpiIntent                  PaymentMethodTypes = "upi_intent"
	PaymentMethodTypesVipps                      PaymentMethodTypes = "vipps"
	PaymentMethodTypesVietQr                     PaymentMethodTypes = "viet_qr"
	PaymentMethodTypesVenmo                      PaymentMethodTypes = "venmo"
	PaymentMethodTypesWalley                     PaymentMethodTypes = "walley"
	PaymentMethodTypesWeChatPay                  PaymentMethodTypes = "we_chat_pay"
	PaymentMethodTypesSevenEleven                PaymentMethodTypes = "seven_eleven"
	PaymentMethodTypesLawson                     PaymentMethodTypes = "lawson"
	PaymentMethodTypesMiniStop                   PaymentMethodTypes = "mini_stop"
	PaymentMethodTypesFamilyMart                 PaymentMethodTypes = "family_mart"
	PaymentMethodTypesSeicomart                  PaymentMethodTypes = "seicomart"
	PaymentMethodTypesPayEasy                    PaymentMethodTypes = "pay_easy"
	PaymentMethodTypesLocalBankTransfer          PaymentMethodTypes = "local_bank_transfer"
	PaymentMethodTypesMifinity                   PaymentMethodTypes = "mifinity"
	PaymentMethodTypesOpenBankingPis             PaymentMethodTypes = "open_banking_pis"
	PaymentMethodTypesDirectCarrierBilling       PaymentMethodTypes = "direct_carrier_billing"
	PaymentMethodTypesInstantBankTransfer        PaymentMethodTypes = "instant_bank_transfer"
	PaymentMethodTypesBillie                     PaymentMethodTypes = "billie"
	PaymentMethodTypesZip                        PaymentMethodTypes = "zip"
	PaymentMethodTypesRevolutPay                 PaymentMethodTypes = "revolut_pay"
	PaymentMethodTypesNaverPay                   PaymentMethodTypes = "naver_pay"
	PaymentMethodTypesPayco                      PaymentMethodTypes = "payco"
)

func (r PaymentMethodTypes) IsKnown() bool {
	switch r {
	case PaymentMethodTypesACH, PaymentMethodTypesAffirm, PaymentMethodTypesAfterpayClearpay, PaymentMethodTypesAlfamart, PaymentMethodTypesAliPay, PaymentMethodTypesAliPayHk, PaymentMethodTypesAlma, PaymentMethodTypesAmazonPay, PaymentMethodTypesApplePay, PaymentMethodTypesAtome, PaymentMethodTypesBacs, PaymentMethodTypesBancontactCard, PaymentMethodTypesBecs, PaymentMethodTypesBenefit, PaymentMethodTypesBizum, PaymentMethodTypesBlik, PaymentMethodTypesBoleto, PaymentMethodTypesBcaBankTransfer, PaymentMethodTypesBniVa, PaymentMethodTypesBriVa, PaymentMethodTypesCardRedirect, PaymentMethodTypesCimbVa, PaymentMethodTypesClassic, PaymentMethodTypesCredit, PaymentMethodTypesCryptoCurrency, PaymentMethodTypesCashapp, PaymentMethodTypesDana, PaymentMethodTypesDanamonVa, PaymentMethodTypesDebit, PaymentMethodTypesDuitNow, PaymentMethodTypesEfecty, PaymentMethodTypesEft, PaymentMethodTypesEps, PaymentMethodTypesFps, PaymentMethodTypesEvoucher, PaymentMethodTypesGiropay, PaymentMethodTypesGivex, PaymentMethodTypesGooglePay, PaymentMethodTypesGoPay, PaymentMethodTypesGcash, PaymentMethodTypesIdeal, PaymentMethodTypesInterac, PaymentMethodTypesIndomaret, PaymentMethodTypesKlarna, PaymentMethodTypesKakaoPay, PaymentMethodTypesLocalBankRedirect, PaymentMethodTypesMandiriVa, PaymentMethodTypesKnet, PaymentMethodTypesMBWay, PaymentMethodTypesMobilePay, PaymentMethodTypesMomo, PaymentMethodTypesMomoAtm, PaymentMethodTypesMultibanco, PaymentMethodTypesOnlineBankingThailand, PaymentMethodTypesOnlineBankingCzechRepublic, PaymentMethodTypesOnlineBankingFinland, PaymentMethodTypesOnlineBankingFpx, PaymentMethodTypesOnlineBankingPoland, PaymentMethodTypesOnlineBankingSlovakia, PaymentMethodTypesOxxo, PaymentMethodTypesPagoEfectivo, PaymentMethodTypesPermataBankTransfer, PaymentMethodTypesOpenBankingUk, PaymentMethodTypesPayBright, PaymentMethodTypesPaypal, PaymentMethodTypesPaze, PaymentMethodTypesPix, PaymentMethodTypesPaySafeCard, PaymentMethodTypesPrzelewy24, PaymentMethodTypesPromptPay, PaymentMethodTypesPse, PaymentMethodTypesRedCompra, PaymentMethodTypesRedPagos, PaymentMethodTypesSamsungPay, PaymentMethodTypesSepa, PaymentMethodTypesSepaBankTransfer, PaymentMethodTypesSofort, PaymentMethodTypesSwish, PaymentMethodTypesTouchNGo, PaymentMethodTypesTrustly, PaymentMethodTypesTwint, PaymentMethodTypesUpiCollect, PaymentMethodTypesUpiIntent, PaymentMethodTypesVipps, PaymentMethodTypesVietQr, PaymentMethodTypesVenmo, PaymentMethodTypesWalley, PaymentMethodTypesWeChatPay, PaymentMethodTypesSevenEleven, PaymentMethodTypesLawson, PaymentMethodTypesMiniStop, PaymentMethodTypesFamilyMart, PaymentMethodTypesSeicomart, PaymentMethodTypesPayEasy, PaymentMethodTypesLocalBankTransfer, PaymentMethodTypesMifinity, PaymentMethodTypesOpenBankingPis, PaymentMethodTypesDirectCarrierBilling, PaymentMethodTypesInstantBankTransfer, PaymentMethodTypesBillie, PaymentMethodTypesZip, PaymentMethodTypesRevolutPay, PaymentMethodTypesNaverPay, PaymentMethodTypesPayco:
		return true
	}
	return false
}

type PaymentNewResponse struct {
	// Client secret used to load Dodo checkout SDK NOTE : Dodo checkout SDK will be
	// coming soon
	ClientSecret string `json:"client_secret" api:"required"`
	// Limited details about the customer making the payment
	Customer CustomerLimitedDetails `json:"customer" api:"required"`
	// Additional metadata associated with the payment
	Metadata map[string]string `json:"metadata" api:"required"`
	// Unique identifier for the payment
	PaymentID string `json:"payment_id" api:"required"`
	// Total amount of the payment in smallest currency unit (e.g. cents)
	TotalAmount int64 `json:"total_amount" api:"required"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id" api:"nullable"`
	// Expiry timestamp of the payment link
	ExpiresOn time.Time `json:"expires_on" api:"nullable" format:"date-time"`
	// Optional URL to a hosted payment page
	PaymentLink string `json:"payment_link" api:"nullable"`
	// Optional list of products included in the payment
	ProductCart []OneTimeProductCartItem `json:"product_cart" api:"nullable"`
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
	ExpiresOn    apijson.Field
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
	BrandID                  string                 `json:"brand_id" api:"required"`
	CreatedAt                time.Time              `json:"created_at" api:"required" format:"date-time"`
	Currency                 Currency               `json:"currency" api:"required"`
	Customer                 CustomerLimitedDetails `json:"customer" api:"required"`
	DigitalProductsDelivered bool                   `json:"digital_products_delivered" api:"required"`
	HasLicenseKey            bool                   `json:"has_license_key" api:"required"`
	Metadata                 map[string]string      `json:"metadata" api:"required"`
	PaymentID                string                 `json:"payment_id" api:"required"`
	TotalAmount              int64                  `json:"total_amount" api:"required"`
	// Invoice ID for this payment. Uses India-specific invoice ID if available.
	InvoiceID string `json:"invoice_id" api:"nullable"`
	// URL to download the invoice PDF for this payment.
	InvoiceURL        string                  `json:"invoice_url" api:"nullable"`
	PaymentMethod     string                  `json:"payment_method" api:"nullable"`
	PaymentMethodType string                  `json:"payment_method_type" api:"nullable"`
	Status            IntentStatus            `json:"status" api:"nullable"`
	SubscriptionID    string                  `json:"subscription_id" api:"nullable"`
	JSON              paymentListResponseJSON `json:"-"`
}

// paymentListResponseJSON contains the JSON metadata for the struct
// [PaymentListResponse]
type paymentListResponseJSON struct {
	BrandID                  apijson.Field
	CreatedAt                apijson.Field
	Currency                 apijson.Field
	Customer                 apijson.Field
	DigitalProductsDelivered apijson.Field
	HasLicenseKey            apijson.Field
	Metadata                 apijson.Field
	PaymentID                apijson.Field
	TotalAmount              apijson.Field
	InvoiceID                apijson.Field
	InvoiceURL               apijson.Field
	PaymentMethod            apijson.Field
	PaymentMethodType        apijson.Field
	Status                   apijson.Field
	SubscriptionID           apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *PaymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentGetLineItemsResponse struct {
	Currency Currency                          `json:"currency" api:"required"`
	Items    []PaymentGetLineItemsResponseItem `json:"items" api:"required"`
	JSON     paymentGetLineItemsResponseJSON   `json:"-"`
}

// paymentGetLineItemsResponseJSON contains the JSON metadata for the struct
// [PaymentGetLineItemsResponse]
type paymentGetLineItemsResponseJSON struct {
	Currency    apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentGetLineItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGetLineItemsResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentGetLineItemsResponseItem struct {
	Amount           int64                               `json:"amount" api:"required"`
	ItemsID          string                              `json:"items_id" api:"required"`
	RefundableAmount int64                               `json:"refundable_amount" api:"required"`
	Tax              int64                               `json:"tax" api:"required"`
	Description      string                              `json:"description" api:"nullable"`
	Name             string                              `json:"name" api:"nullable"`
	JSON             paymentGetLineItemsResponseItemJSON `json:"-"`
}

// paymentGetLineItemsResponseItemJSON contains the JSON metadata for the struct
// [PaymentGetLineItemsResponseItem]
type paymentGetLineItemsResponseItemJSON struct {
	Amount           apijson.Field
	ItemsID          apijson.Field
	RefundableAmount apijson.Field
	Tax              apijson.Field
	Description      apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PaymentGetLineItemsResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentGetLineItemsResponseItemJSON) RawJSON() string {
	return r.raw
}

type PaymentNewParams struct {
	// Billing address details for the payment
	Billing param.Field[BillingAddressParam] `json:"billing" api:"required"`
	// Customer information for the payment
	Customer param.Field[CustomerRequestUnionParam] `json:"customer" api:"required"`
	// List of products in the cart. Must contain at least 1 and at most 100 items.
	ProductCart param.Field[[]OneTimeProductCartItemParam] `json:"product_cart" api:"required"`
	// List of payment methods allowed during checkout.
	//
	// Customers will **never** see payment methods that are **not** in this list.
	// However, adding a method here **does not guarantee** customers will see it.
	// Availability still depends on other factors (e.g., customer location, merchant
	// settings).
	AllowedPaymentMethodTypes param.Field[[]PaymentMethodTypes] `json:"allowed_payment_method_types"`
	// Fix the currency in which the end customer is billed. If Dodo Payments cannot
	// support that currency for this transaction, it will not proceed
	BillingCurrency param.Field[Currency] `json:"billing_currency"`
	// Discount Code to apply to the transaction
	DiscountCode param.Field[string] `json:"discount_code"`
	// Override merchant default 3DS behaviour for this payment
	Force3DS param.Field[bool] `json:"force_3ds"`
	// Additional metadata associated with the payment. Defaults to empty if not
	// provided.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Whether to generate a payment link. Defaults to false if not specified.
	PaymentLink param.Field[bool] `json:"payment_link"`
	// Optional payment method ID to use for this payment. If provided, customer_id
	// must also be provided. The payment method will be validated for eligibility with
	// the payment's currency.
	PaymentMethodID param.Field[string] `json:"payment_method_id"`
	// If true, redirects the customer immediately after payment completion False by
	// default
	RedirectImmediately param.Field[bool] `json:"redirect_immediately"`
	// Optional URL to redirect the customer after payment. Must be a valid URL if
	// provided.
	ReturnURL param.Field[string] `json:"return_url"`
	// If true, returns a shortened payment link. Defaults to false if not specified.
	ShortLink param.Field[bool] `json:"short_link"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Field[bool] `json:"show_saved_payment_methods"`
	// Tax ID in case the payment is B2B. If tax id validation fails the payment
	// creation will fail
	TaxID param.Field[string] `json:"tax_id"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentListParams struct {
	// filter by Brand id
	BrandID param.Field[string] `query:"brand_id"`
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
	// Filter by product id
	ProductID param.Field[string] `query:"product_id"`
	// Filter by status
	Status param.Field[PaymentListParamsStatus] `query:"status"`
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

// Filter by status
type PaymentListParamsStatus string

const (
	PaymentListParamsStatusSucceeded                      PaymentListParamsStatus = "succeeded"
	PaymentListParamsStatusFailed                         PaymentListParamsStatus = "failed"
	PaymentListParamsStatusCancelled                      PaymentListParamsStatus = "cancelled"
	PaymentListParamsStatusProcessing                     PaymentListParamsStatus = "processing"
	PaymentListParamsStatusRequiresCustomerAction         PaymentListParamsStatus = "requires_customer_action"
	PaymentListParamsStatusRequiresMerchantAction         PaymentListParamsStatus = "requires_merchant_action"
	PaymentListParamsStatusRequiresPaymentMethod          PaymentListParamsStatus = "requires_payment_method"
	PaymentListParamsStatusRequiresConfirmation           PaymentListParamsStatus = "requires_confirmation"
	PaymentListParamsStatusRequiresCapture                PaymentListParamsStatus = "requires_capture"
	PaymentListParamsStatusPartiallyCaptured              PaymentListParamsStatus = "partially_captured"
	PaymentListParamsStatusPartiallyCapturedAndCapturable PaymentListParamsStatus = "partially_captured_and_capturable"
)

func (r PaymentListParamsStatus) IsKnown() bool {
	switch r {
	case PaymentListParamsStatusSucceeded, PaymentListParamsStatusFailed, PaymentListParamsStatusCancelled, PaymentListParamsStatusProcessing, PaymentListParamsStatusRequiresCustomerAction, PaymentListParamsStatusRequiresMerchantAction, PaymentListParamsStatusRequiresPaymentMethod, PaymentListParamsStatusRequiresConfirmation, PaymentListParamsStatusRequiresCapture, PaymentListParamsStatusPartiallyCaptured, PaymentListParamsStatusPartiallyCapturedAndCapturable:
		return true
	}
	return false
}
