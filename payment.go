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
	// Identifier of the business associated with the payment
	BusinessID string `json:"business_id,required"`
	// Timestamp when the payment was created
	CreatedAt time.Time              `json:"created_at,required" format:"date-time"`
	Currency  PaymentCurrency        `json:"currency,required"`
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
	SettlementAmount   int64                     `json:"settlement_amount,required"`
	SettlementCurrency PaymentSettlementCurrency `json:"settlement_currency,required"`
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
	Status      IntentStatus         `json:"status,nullable"`
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

type PaymentCurrency string

const (
	PaymentCurrencyAed PaymentCurrency = "AED"
	PaymentCurrencyAll PaymentCurrency = "ALL"
	PaymentCurrencyAmd PaymentCurrency = "AMD"
	PaymentCurrencyAng PaymentCurrency = "ANG"
	PaymentCurrencyAoa PaymentCurrency = "AOA"
	PaymentCurrencyArs PaymentCurrency = "ARS"
	PaymentCurrencyAud PaymentCurrency = "AUD"
	PaymentCurrencyAwg PaymentCurrency = "AWG"
	PaymentCurrencyAzn PaymentCurrency = "AZN"
	PaymentCurrencyBam PaymentCurrency = "BAM"
	PaymentCurrencyBbd PaymentCurrency = "BBD"
	PaymentCurrencyBdt PaymentCurrency = "BDT"
	PaymentCurrencyBgn PaymentCurrency = "BGN"
	PaymentCurrencyBhd PaymentCurrency = "BHD"
	PaymentCurrencyBif PaymentCurrency = "BIF"
	PaymentCurrencyBmd PaymentCurrency = "BMD"
	PaymentCurrencyBnd PaymentCurrency = "BND"
	PaymentCurrencyBob PaymentCurrency = "BOB"
	PaymentCurrencyBrl PaymentCurrency = "BRL"
	PaymentCurrencyBsd PaymentCurrency = "BSD"
	PaymentCurrencyBwp PaymentCurrency = "BWP"
	PaymentCurrencyByn PaymentCurrency = "BYN"
	PaymentCurrencyBzd PaymentCurrency = "BZD"
	PaymentCurrencyCad PaymentCurrency = "CAD"
	PaymentCurrencyChf PaymentCurrency = "CHF"
	PaymentCurrencyClp PaymentCurrency = "CLP"
	PaymentCurrencyCny PaymentCurrency = "CNY"
	PaymentCurrencyCop PaymentCurrency = "COP"
	PaymentCurrencyCrc PaymentCurrency = "CRC"
	PaymentCurrencyCup PaymentCurrency = "CUP"
	PaymentCurrencyCve PaymentCurrency = "CVE"
	PaymentCurrencyCzk PaymentCurrency = "CZK"
	PaymentCurrencyDjf PaymentCurrency = "DJF"
	PaymentCurrencyDkk PaymentCurrency = "DKK"
	PaymentCurrencyDop PaymentCurrency = "DOP"
	PaymentCurrencyDzd PaymentCurrency = "DZD"
	PaymentCurrencyEgp PaymentCurrency = "EGP"
	PaymentCurrencyEtb PaymentCurrency = "ETB"
	PaymentCurrencyEur PaymentCurrency = "EUR"
	PaymentCurrencyFjd PaymentCurrency = "FJD"
	PaymentCurrencyFkp PaymentCurrency = "FKP"
	PaymentCurrencyGbp PaymentCurrency = "GBP"
	PaymentCurrencyGel PaymentCurrency = "GEL"
	PaymentCurrencyGhs PaymentCurrency = "GHS"
	PaymentCurrencyGip PaymentCurrency = "GIP"
	PaymentCurrencyGmd PaymentCurrency = "GMD"
	PaymentCurrencyGnf PaymentCurrency = "GNF"
	PaymentCurrencyGtq PaymentCurrency = "GTQ"
	PaymentCurrencyGyd PaymentCurrency = "GYD"
	PaymentCurrencyHkd PaymentCurrency = "HKD"
	PaymentCurrencyHnl PaymentCurrency = "HNL"
	PaymentCurrencyHrk PaymentCurrency = "HRK"
	PaymentCurrencyHtg PaymentCurrency = "HTG"
	PaymentCurrencyHuf PaymentCurrency = "HUF"
	PaymentCurrencyIdr PaymentCurrency = "IDR"
	PaymentCurrencyIls PaymentCurrency = "ILS"
	PaymentCurrencyInr PaymentCurrency = "INR"
	PaymentCurrencyIqd PaymentCurrency = "IQD"
	PaymentCurrencyJmd PaymentCurrency = "JMD"
	PaymentCurrencyJod PaymentCurrency = "JOD"
	PaymentCurrencyJpy PaymentCurrency = "JPY"
	PaymentCurrencyKes PaymentCurrency = "KES"
	PaymentCurrencyKgs PaymentCurrency = "KGS"
	PaymentCurrencyKhr PaymentCurrency = "KHR"
	PaymentCurrencyKmf PaymentCurrency = "KMF"
	PaymentCurrencyKrw PaymentCurrency = "KRW"
	PaymentCurrencyKwd PaymentCurrency = "KWD"
	PaymentCurrencyKyd PaymentCurrency = "KYD"
	PaymentCurrencyKzt PaymentCurrency = "KZT"
	PaymentCurrencyLak PaymentCurrency = "LAK"
	PaymentCurrencyLbp PaymentCurrency = "LBP"
	PaymentCurrencyLkr PaymentCurrency = "LKR"
	PaymentCurrencyLrd PaymentCurrency = "LRD"
	PaymentCurrencyLsl PaymentCurrency = "LSL"
	PaymentCurrencyLyd PaymentCurrency = "LYD"
	PaymentCurrencyMad PaymentCurrency = "MAD"
	PaymentCurrencyMdl PaymentCurrency = "MDL"
	PaymentCurrencyMga PaymentCurrency = "MGA"
	PaymentCurrencyMkd PaymentCurrency = "MKD"
	PaymentCurrencyMmk PaymentCurrency = "MMK"
	PaymentCurrencyMnt PaymentCurrency = "MNT"
	PaymentCurrencyMop PaymentCurrency = "MOP"
	PaymentCurrencyMru PaymentCurrency = "MRU"
	PaymentCurrencyMur PaymentCurrency = "MUR"
	PaymentCurrencyMvr PaymentCurrency = "MVR"
	PaymentCurrencyMwk PaymentCurrency = "MWK"
	PaymentCurrencyMxn PaymentCurrency = "MXN"
	PaymentCurrencyMyr PaymentCurrency = "MYR"
	PaymentCurrencyMzn PaymentCurrency = "MZN"
	PaymentCurrencyNad PaymentCurrency = "NAD"
	PaymentCurrencyNgn PaymentCurrency = "NGN"
	PaymentCurrencyNio PaymentCurrency = "NIO"
	PaymentCurrencyNok PaymentCurrency = "NOK"
	PaymentCurrencyNpr PaymentCurrency = "NPR"
	PaymentCurrencyNzd PaymentCurrency = "NZD"
	PaymentCurrencyOmr PaymentCurrency = "OMR"
	PaymentCurrencyPab PaymentCurrency = "PAB"
	PaymentCurrencyPen PaymentCurrency = "PEN"
	PaymentCurrencyPgk PaymentCurrency = "PGK"
	PaymentCurrencyPhp PaymentCurrency = "PHP"
	PaymentCurrencyPkr PaymentCurrency = "PKR"
	PaymentCurrencyPln PaymentCurrency = "PLN"
	PaymentCurrencyPyg PaymentCurrency = "PYG"
	PaymentCurrencyQar PaymentCurrency = "QAR"
	PaymentCurrencyRon PaymentCurrency = "RON"
	PaymentCurrencyRsd PaymentCurrency = "RSD"
	PaymentCurrencyRub PaymentCurrency = "RUB"
	PaymentCurrencyRwf PaymentCurrency = "RWF"
	PaymentCurrencySar PaymentCurrency = "SAR"
	PaymentCurrencySbd PaymentCurrency = "SBD"
	PaymentCurrencyScr PaymentCurrency = "SCR"
	PaymentCurrencySek PaymentCurrency = "SEK"
	PaymentCurrencySgd PaymentCurrency = "SGD"
	PaymentCurrencyShp PaymentCurrency = "SHP"
	PaymentCurrencySle PaymentCurrency = "SLE"
	PaymentCurrencySll PaymentCurrency = "SLL"
	PaymentCurrencySos PaymentCurrency = "SOS"
	PaymentCurrencySrd PaymentCurrency = "SRD"
	PaymentCurrencySsp PaymentCurrency = "SSP"
	PaymentCurrencyStn PaymentCurrency = "STN"
	PaymentCurrencySvc PaymentCurrency = "SVC"
	PaymentCurrencySzl PaymentCurrency = "SZL"
	PaymentCurrencyThb PaymentCurrency = "THB"
	PaymentCurrencyTnd PaymentCurrency = "TND"
	PaymentCurrencyTop PaymentCurrency = "TOP"
	PaymentCurrencyTry PaymentCurrency = "TRY"
	PaymentCurrencyTtd PaymentCurrency = "TTD"
	PaymentCurrencyTwd PaymentCurrency = "TWD"
	PaymentCurrencyTzs PaymentCurrency = "TZS"
	PaymentCurrencyUah PaymentCurrency = "UAH"
	PaymentCurrencyUgx PaymentCurrency = "UGX"
	PaymentCurrencyUsd PaymentCurrency = "USD"
	PaymentCurrencyUyu PaymentCurrency = "UYU"
	PaymentCurrencyUzs PaymentCurrency = "UZS"
	PaymentCurrencyVes PaymentCurrency = "VES"
	PaymentCurrencyVnd PaymentCurrency = "VND"
	PaymentCurrencyVuv PaymentCurrency = "VUV"
	PaymentCurrencyWst PaymentCurrency = "WST"
	PaymentCurrencyXaf PaymentCurrency = "XAF"
	PaymentCurrencyXcd PaymentCurrency = "XCD"
	PaymentCurrencyXof PaymentCurrency = "XOF"
	PaymentCurrencyXpf PaymentCurrency = "XPF"
	PaymentCurrencyYer PaymentCurrency = "YER"
	PaymentCurrencyZar PaymentCurrency = "ZAR"
	PaymentCurrencyZmw PaymentCurrency = "ZMW"
)

func (r PaymentCurrency) IsKnown() bool {
	switch r {
	case PaymentCurrencyAed, PaymentCurrencyAll, PaymentCurrencyAmd, PaymentCurrencyAng, PaymentCurrencyAoa, PaymentCurrencyArs, PaymentCurrencyAud, PaymentCurrencyAwg, PaymentCurrencyAzn, PaymentCurrencyBam, PaymentCurrencyBbd, PaymentCurrencyBdt, PaymentCurrencyBgn, PaymentCurrencyBhd, PaymentCurrencyBif, PaymentCurrencyBmd, PaymentCurrencyBnd, PaymentCurrencyBob, PaymentCurrencyBrl, PaymentCurrencyBsd, PaymentCurrencyBwp, PaymentCurrencyByn, PaymentCurrencyBzd, PaymentCurrencyCad, PaymentCurrencyChf, PaymentCurrencyClp, PaymentCurrencyCny, PaymentCurrencyCop, PaymentCurrencyCrc, PaymentCurrencyCup, PaymentCurrencyCve, PaymentCurrencyCzk, PaymentCurrencyDjf, PaymentCurrencyDkk, PaymentCurrencyDop, PaymentCurrencyDzd, PaymentCurrencyEgp, PaymentCurrencyEtb, PaymentCurrencyEur, PaymentCurrencyFjd, PaymentCurrencyFkp, PaymentCurrencyGbp, PaymentCurrencyGel, PaymentCurrencyGhs, PaymentCurrencyGip, PaymentCurrencyGmd, PaymentCurrencyGnf, PaymentCurrencyGtq, PaymentCurrencyGyd, PaymentCurrencyHkd, PaymentCurrencyHnl, PaymentCurrencyHrk, PaymentCurrencyHtg, PaymentCurrencyHuf, PaymentCurrencyIdr, PaymentCurrencyIls, PaymentCurrencyInr, PaymentCurrencyIqd, PaymentCurrencyJmd, PaymentCurrencyJod, PaymentCurrencyJpy, PaymentCurrencyKes, PaymentCurrencyKgs, PaymentCurrencyKhr, PaymentCurrencyKmf, PaymentCurrencyKrw, PaymentCurrencyKwd, PaymentCurrencyKyd, PaymentCurrencyKzt, PaymentCurrencyLak, PaymentCurrencyLbp, PaymentCurrencyLkr, PaymentCurrencyLrd, PaymentCurrencyLsl, PaymentCurrencyLyd, PaymentCurrencyMad, PaymentCurrencyMdl, PaymentCurrencyMga, PaymentCurrencyMkd, PaymentCurrencyMmk, PaymentCurrencyMnt, PaymentCurrencyMop, PaymentCurrencyMru, PaymentCurrencyMur, PaymentCurrencyMvr, PaymentCurrencyMwk, PaymentCurrencyMxn, PaymentCurrencyMyr, PaymentCurrencyMzn, PaymentCurrencyNad, PaymentCurrencyNgn, PaymentCurrencyNio, PaymentCurrencyNok, PaymentCurrencyNpr, PaymentCurrencyNzd, PaymentCurrencyOmr, PaymentCurrencyPab, PaymentCurrencyPen, PaymentCurrencyPgk, PaymentCurrencyPhp, PaymentCurrencyPkr, PaymentCurrencyPln, PaymentCurrencyPyg, PaymentCurrencyQar, PaymentCurrencyRon, PaymentCurrencyRsd, PaymentCurrencyRub, PaymentCurrencyRwf, PaymentCurrencySar, PaymentCurrencySbd, PaymentCurrencyScr, PaymentCurrencySek, PaymentCurrencySgd, PaymentCurrencyShp, PaymentCurrencySle, PaymentCurrencySll, PaymentCurrencySos, PaymentCurrencySrd, PaymentCurrencySsp, PaymentCurrencyStn, PaymentCurrencySvc, PaymentCurrencySzl, PaymentCurrencyThb, PaymentCurrencyTnd, PaymentCurrencyTop, PaymentCurrencyTry, PaymentCurrencyTtd, PaymentCurrencyTwd, PaymentCurrencyTzs, PaymentCurrencyUah, PaymentCurrencyUgx, PaymentCurrencyUsd, PaymentCurrencyUyu, PaymentCurrencyUzs, PaymentCurrencyVes, PaymentCurrencyVnd, PaymentCurrencyVuv, PaymentCurrencyWst, PaymentCurrencyXaf, PaymentCurrencyXcd, PaymentCurrencyXof, PaymentCurrencyXpf, PaymentCurrencyYer, PaymentCurrencyZar, PaymentCurrencyZmw:
		return true
	}
	return false
}

type PaymentSettlementCurrency string

const (
	PaymentSettlementCurrencyAed PaymentSettlementCurrency = "AED"
	PaymentSettlementCurrencyAll PaymentSettlementCurrency = "ALL"
	PaymentSettlementCurrencyAmd PaymentSettlementCurrency = "AMD"
	PaymentSettlementCurrencyAng PaymentSettlementCurrency = "ANG"
	PaymentSettlementCurrencyAoa PaymentSettlementCurrency = "AOA"
	PaymentSettlementCurrencyArs PaymentSettlementCurrency = "ARS"
	PaymentSettlementCurrencyAud PaymentSettlementCurrency = "AUD"
	PaymentSettlementCurrencyAwg PaymentSettlementCurrency = "AWG"
	PaymentSettlementCurrencyAzn PaymentSettlementCurrency = "AZN"
	PaymentSettlementCurrencyBam PaymentSettlementCurrency = "BAM"
	PaymentSettlementCurrencyBbd PaymentSettlementCurrency = "BBD"
	PaymentSettlementCurrencyBdt PaymentSettlementCurrency = "BDT"
	PaymentSettlementCurrencyBgn PaymentSettlementCurrency = "BGN"
	PaymentSettlementCurrencyBhd PaymentSettlementCurrency = "BHD"
	PaymentSettlementCurrencyBif PaymentSettlementCurrency = "BIF"
	PaymentSettlementCurrencyBmd PaymentSettlementCurrency = "BMD"
	PaymentSettlementCurrencyBnd PaymentSettlementCurrency = "BND"
	PaymentSettlementCurrencyBob PaymentSettlementCurrency = "BOB"
	PaymentSettlementCurrencyBrl PaymentSettlementCurrency = "BRL"
	PaymentSettlementCurrencyBsd PaymentSettlementCurrency = "BSD"
	PaymentSettlementCurrencyBwp PaymentSettlementCurrency = "BWP"
	PaymentSettlementCurrencyByn PaymentSettlementCurrency = "BYN"
	PaymentSettlementCurrencyBzd PaymentSettlementCurrency = "BZD"
	PaymentSettlementCurrencyCad PaymentSettlementCurrency = "CAD"
	PaymentSettlementCurrencyChf PaymentSettlementCurrency = "CHF"
	PaymentSettlementCurrencyClp PaymentSettlementCurrency = "CLP"
	PaymentSettlementCurrencyCny PaymentSettlementCurrency = "CNY"
	PaymentSettlementCurrencyCop PaymentSettlementCurrency = "COP"
	PaymentSettlementCurrencyCrc PaymentSettlementCurrency = "CRC"
	PaymentSettlementCurrencyCup PaymentSettlementCurrency = "CUP"
	PaymentSettlementCurrencyCve PaymentSettlementCurrency = "CVE"
	PaymentSettlementCurrencyCzk PaymentSettlementCurrency = "CZK"
	PaymentSettlementCurrencyDjf PaymentSettlementCurrency = "DJF"
	PaymentSettlementCurrencyDkk PaymentSettlementCurrency = "DKK"
	PaymentSettlementCurrencyDop PaymentSettlementCurrency = "DOP"
	PaymentSettlementCurrencyDzd PaymentSettlementCurrency = "DZD"
	PaymentSettlementCurrencyEgp PaymentSettlementCurrency = "EGP"
	PaymentSettlementCurrencyEtb PaymentSettlementCurrency = "ETB"
	PaymentSettlementCurrencyEur PaymentSettlementCurrency = "EUR"
	PaymentSettlementCurrencyFjd PaymentSettlementCurrency = "FJD"
	PaymentSettlementCurrencyFkp PaymentSettlementCurrency = "FKP"
	PaymentSettlementCurrencyGbp PaymentSettlementCurrency = "GBP"
	PaymentSettlementCurrencyGel PaymentSettlementCurrency = "GEL"
	PaymentSettlementCurrencyGhs PaymentSettlementCurrency = "GHS"
	PaymentSettlementCurrencyGip PaymentSettlementCurrency = "GIP"
	PaymentSettlementCurrencyGmd PaymentSettlementCurrency = "GMD"
	PaymentSettlementCurrencyGnf PaymentSettlementCurrency = "GNF"
	PaymentSettlementCurrencyGtq PaymentSettlementCurrency = "GTQ"
	PaymentSettlementCurrencyGyd PaymentSettlementCurrency = "GYD"
	PaymentSettlementCurrencyHkd PaymentSettlementCurrency = "HKD"
	PaymentSettlementCurrencyHnl PaymentSettlementCurrency = "HNL"
	PaymentSettlementCurrencyHrk PaymentSettlementCurrency = "HRK"
	PaymentSettlementCurrencyHtg PaymentSettlementCurrency = "HTG"
	PaymentSettlementCurrencyHuf PaymentSettlementCurrency = "HUF"
	PaymentSettlementCurrencyIdr PaymentSettlementCurrency = "IDR"
	PaymentSettlementCurrencyIls PaymentSettlementCurrency = "ILS"
	PaymentSettlementCurrencyInr PaymentSettlementCurrency = "INR"
	PaymentSettlementCurrencyIqd PaymentSettlementCurrency = "IQD"
	PaymentSettlementCurrencyJmd PaymentSettlementCurrency = "JMD"
	PaymentSettlementCurrencyJod PaymentSettlementCurrency = "JOD"
	PaymentSettlementCurrencyJpy PaymentSettlementCurrency = "JPY"
	PaymentSettlementCurrencyKes PaymentSettlementCurrency = "KES"
	PaymentSettlementCurrencyKgs PaymentSettlementCurrency = "KGS"
	PaymentSettlementCurrencyKhr PaymentSettlementCurrency = "KHR"
	PaymentSettlementCurrencyKmf PaymentSettlementCurrency = "KMF"
	PaymentSettlementCurrencyKrw PaymentSettlementCurrency = "KRW"
	PaymentSettlementCurrencyKwd PaymentSettlementCurrency = "KWD"
	PaymentSettlementCurrencyKyd PaymentSettlementCurrency = "KYD"
	PaymentSettlementCurrencyKzt PaymentSettlementCurrency = "KZT"
	PaymentSettlementCurrencyLak PaymentSettlementCurrency = "LAK"
	PaymentSettlementCurrencyLbp PaymentSettlementCurrency = "LBP"
	PaymentSettlementCurrencyLkr PaymentSettlementCurrency = "LKR"
	PaymentSettlementCurrencyLrd PaymentSettlementCurrency = "LRD"
	PaymentSettlementCurrencyLsl PaymentSettlementCurrency = "LSL"
	PaymentSettlementCurrencyLyd PaymentSettlementCurrency = "LYD"
	PaymentSettlementCurrencyMad PaymentSettlementCurrency = "MAD"
	PaymentSettlementCurrencyMdl PaymentSettlementCurrency = "MDL"
	PaymentSettlementCurrencyMga PaymentSettlementCurrency = "MGA"
	PaymentSettlementCurrencyMkd PaymentSettlementCurrency = "MKD"
	PaymentSettlementCurrencyMmk PaymentSettlementCurrency = "MMK"
	PaymentSettlementCurrencyMnt PaymentSettlementCurrency = "MNT"
	PaymentSettlementCurrencyMop PaymentSettlementCurrency = "MOP"
	PaymentSettlementCurrencyMru PaymentSettlementCurrency = "MRU"
	PaymentSettlementCurrencyMur PaymentSettlementCurrency = "MUR"
	PaymentSettlementCurrencyMvr PaymentSettlementCurrency = "MVR"
	PaymentSettlementCurrencyMwk PaymentSettlementCurrency = "MWK"
	PaymentSettlementCurrencyMxn PaymentSettlementCurrency = "MXN"
	PaymentSettlementCurrencyMyr PaymentSettlementCurrency = "MYR"
	PaymentSettlementCurrencyMzn PaymentSettlementCurrency = "MZN"
	PaymentSettlementCurrencyNad PaymentSettlementCurrency = "NAD"
	PaymentSettlementCurrencyNgn PaymentSettlementCurrency = "NGN"
	PaymentSettlementCurrencyNio PaymentSettlementCurrency = "NIO"
	PaymentSettlementCurrencyNok PaymentSettlementCurrency = "NOK"
	PaymentSettlementCurrencyNpr PaymentSettlementCurrency = "NPR"
	PaymentSettlementCurrencyNzd PaymentSettlementCurrency = "NZD"
	PaymentSettlementCurrencyOmr PaymentSettlementCurrency = "OMR"
	PaymentSettlementCurrencyPab PaymentSettlementCurrency = "PAB"
	PaymentSettlementCurrencyPen PaymentSettlementCurrency = "PEN"
	PaymentSettlementCurrencyPgk PaymentSettlementCurrency = "PGK"
	PaymentSettlementCurrencyPhp PaymentSettlementCurrency = "PHP"
	PaymentSettlementCurrencyPkr PaymentSettlementCurrency = "PKR"
	PaymentSettlementCurrencyPln PaymentSettlementCurrency = "PLN"
	PaymentSettlementCurrencyPyg PaymentSettlementCurrency = "PYG"
	PaymentSettlementCurrencyQar PaymentSettlementCurrency = "QAR"
	PaymentSettlementCurrencyRon PaymentSettlementCurrency = "RON"
	PaymentSettlementCurrencyRsd PaymentSettlementCurrency = "RSD"
	PaymentSettlementCurrencyRub PaymentSettlementCurrency = "RUB"
	PaymentSettlementCurrencyRwf PaymentSettlementCurrency = "RWF"
	PaymentSettlementCurrencySar PaymentSettlementCurrency = "SAR"
	PaymentSettlementCurrencySbd PaymentSettlementCurrency = "SBD"
	PaymentSettlementCurrencyScr PaymentSettlementCurrency = "SCR"
	PaymentSettlementCurrencySek PaymentSettlementCurrency = "SEK"
	PaymentSettlementCurrencySgd PaymentSettlementCurrency = "SGD"
	PaymentSettlementCurrencyShp PaymentSettlementCurrency = "SHP"
	PaymentSettlementCurrencySle PaymentSettlementCurrency = "SLE"
	PaymentSettlementCurrencySll PaymentSettlementCurrency = "SLL"
	PaymentSettlementCurrencySos PaymentSettlementCurrency = "SOS"
	PaymentSettlementCurrencySrd PaymentSettlementCurrency = "SRD"
	PaymentSettlementCurrencySsp PaymentSettlementCurrency = "SSP"
	PaymentSettlementCurrencyStn PaymentSettlementCurrency = "STN"
	PaymentSettlementCurrencySvc PaymentSettlementCurrency = "SVC"
	PaymentSettlementCurrencySzl PaymentSettlementCurrency = "SZL"
	PaymentSettlementCurrencyThb PaymentSettlementCurrency = "THB"
	PaymentSettlementCurrencyTnd PaymentSettlementCurrency = "TND"
	PaymentSettlementCurrencyTop PaymentSettlementCurrency = "TOP"
	PaymentSettlementCurrencyTry PaymentSettlementCurrency = "TRY"
	PaymentSettlementCurrencyTtd PaymentSettlementCurrency = "TTD"
	PaymentSettlementCurrencyTwd PaymentSettlementCurrency = "TWD"
	PaymentSettlementCurrencyTzs PaymentSettlementCurrency = "TZS"
	PaymentSettlementCurrencyUah PaymentSettlementCurrency = "UAH"
	PaymentSettlementCurrencyUgx PaymentSettlementCurrency = "UGX"
	PaymentSettlementCurrencyUsd PaymentSettlementCurrency = "USD"
	PaymentSettlementCurrencyUyu PaymentSettlementCurrency = "UYU"
	PaymentSettlementCurrencyUzs PaymentSettlementCurrency = "UZS"
	PaymentSettlementCurrencyVes PaymentSettlementCurrency = "VES"
	PaymentSettlementCurrencyVnd PaymentSettlementCurrency = "VND"
	PaymentSettlementCurrencyVuv PaymentSettlementCurrency = "VUV"
	PaymentSettlementCurrencyWst PaymentSettlementCurrency = "WST"
	PaymentSettlementCurrencyXaf PaymentSettlementCurrency = "XAF"
	PaymentSettlementCurrencyXcd PaymentSettlementCurrency = "XCD"
	PaymentSettlementCurrencyXof PaymentSettlementCurrency = "XOF"
	PaymentSettlementCurrencyXpf PaymentSettlementCurrency = "XPF"
	PaymentSettlementCurrencyYer PaymentSettlementCurrency = "YER"
	PaymentSettlementCurrencyZar PaymentSettlementCurrency = "ZAR"
	PaymentSettlementCurrencyZmw PaymentSettlementCurrency = "ZMW"
)

func (r PaymentSettlementCurrency) IsKnown() bool {
	switch r {
	case PaymentSettlementCurrencyAed, PaymentSettlementCurrencyAll, PaymentSettlementCurrencyAmd, PaymentSettlementCurrencyAng, PaymentSettlementCurrencyAoa, PaymentSettlementCurrencyArs, PaymentSettlementCurrencyAud, PaymentSettlementCurrencyAwg, PaymentSettlementCurrencyAzn, PaymentSettlementCurrencyBam, PaymentSettlementCurrencyBbd, PaymentSettlementCurrencyBdt, PaymentSettlementCurrencyBgn, PaymentSettlementCurrencyBhd, PaymentSettlementCurrencyBif, PaymentSettlementCurrencyBmd, PaymentSettlementCurrencyBnd, PaymentSettlementCurrencyBob, PaymentSettlementCurrencyBrl, PaymentSettlementCurrencyBsd, PaymentSettlementCurrencyBwp, PaymentSettlementCurrencyByn, PaymentSettlementCurrencyBzd, PaymentSettlementCurrencyCad, PaymentSettlementCurrencyChf, PaymentSettlementCurrencyClp, PaymentSettlementCurrencyCny, PaymentSettlementCurrencyCop, PaymentSettlementCurrencyCrc, PaymentSettlementCurrencyCup, PaymentSettlementCurrencyCve, PaymentSettlementCurrencyCzk, PaymentSettlementCurrencyDjf, PaymentSettlementCurrencyDkk, PaymentSettlementCurrencyDop, PaymentSettlementCurrencyDzd, PaymentSettlementCurrencyEgp, PaymentSettlementCurrencyEtb, PaymentSettlementCurrencyEur, PaymentSettlementCurrencyFjd, PaymentSettlementCurrencyFkp, PaymentSettlementCurrencyGbp, PaymentSettlementCurrencyGel, PaymentSettlementCurrencyGhs, PaymentSettlementCurrencyGip, PaymentSettlementCurrencyGmd, PaymentSettlementCurrencyGnf, PaymentSettlementCurrencyGtq, PaymentSettlementCurrencyGyd, PaymentSettlementCurrencyHkd, PaymentSettlementCurrencyHnl, PaymentSettlementCurrencyHrk, PaymentSettlementCurrencyHtg, PaymentSettlementCurrencyHuf, PaymentSettlementCurrencyIdr, PaymentSettlementCurrencyIls, PaymentSettlementCurrencyInr, PaymentSettlementCurrencyIqd, PaymentSettlementCurrencyJmd, PaymentSettlementCurrencyJod, PaymentSettlementCurrencyJpy, PaymentSettlementCurrencyKes, PaymentSettlementCurrencyKgs, PaymentSettlementCurrencyKhr, PaymentSettlementCurrencyKmf, PaymentSettlementCurrencyKrw, PaymentSettlementCurrencyKwd, PaymentSettlementCurrencyKyd, PaymentSettlementCurrencyKzt, PaymentSettlementCurrencyLak, PaymentSettlementCurrencyLbp, PaymentSettlementCurrencyLkr, PaymentSettlementCurrencyLrd, PaymentSettlementCurrencyLsl, PaymentSettlementCurrencyLyd, PaymentSettlementCurrencyMad, PaymentSettlementCurrencyMdl, PaymentSettlementCurrencyMga, PaymentSettlementCurrencyMkd, PaymentSettlementCurrencyMmk, PaymentSettlementCurrencyMnt, PaymentSettlementCurrencyMop, PaymentSettlementCurrencyMru, PaymentSettlementCurrencyMur, PaymentSettlementCurrencyMvr, PaymentSettlementCurrencyMwk, PaymentSettlementCurrencyMxn, PaymentSettlementCurrencyMyr, PaymentSettlementCurrencyMzn, PaymentSettlementCurrencyNad, PaymentSettlementCurrencyNgn, PaymentSettlementCurrencyNio, PaymentSettlementCurrencyNok, PaymentSettlementCurrencyNpr, PaymentSettlementCurrencyNzd, PaymentSettlementCurrencyOmr, PaymentSettlementCurrencyPab, PaymentSettlementCurrencyPen, PaymentSettlementCurrencyPgk, PaymentSettlementCurrencyPhp, PaymentSettlementCurrencyPkr, PaymentSettlementCurrencyPln, PaymentSettlementCurrencyPyg, PaymentSettlementCurrencyQar, PaymentSettlementCurrencyRon, PaymentSettlementCurrencyRsd, PaymentSettlementCurrencyRub, PaymentSettlementCurrencyRwf, PaymentSettlementCurrencySar, PaymentSettlementCurrencySbd, PaymentSettlementCurrencyScr, PaymentSettlementCurrencySek, PaymentSettlementCurrencySgd, PaymentSettlementCurrencyShp, PaymentSettlementCurrencySle, PaymentSettlementCurrencySll, PaymentSettlementCurrencySos, PaymentSettlementCurrencySrd, PaymentSettlementCurrencySsp, PaymentSettlementCurrencyStn, PaymentSettlementCurrencySvc, PaymentSettlementCurrencySzl, PaymentSettlementCurrencyThb, PaymentSettlementCurrencyTnd, PaymentSettlementCurrencyTop, PaymentSettlementCurrencyTry, PaymentSettlementCurrencyTtd, PaymentSettlementCurrencyTwd, PaymentSettlementCurrencyTzs, PaymentSettlementCurrencyUah, PaymentSettlementCurrencyUgx, PaymentSettlementCurrencyUsd, PaymentSettlementCurrencyUyu, PaymentSettlementCurrencyUzs, PaymentSettlementCurrencyVes, PaymentSettlementCurrencyVnd, PaymentSettlementCurrencyVuv, PaymentSettlementCurrencyWst, PaymentSettlementCurrencyXaf, PaymentSettlementCurrencyXcd, PaymentSettlementCurrencyXof, PaymentSettlementCurrencyXpf, PaymentSettlementCurrencyYer, PaymentSettlementCurrencyZar, PaymentSettlementCurrencyZmw:
		return true
	}
	return false
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
	CreatedAt         time.Time                   `json:"created_at,required" format:"date-time"`
	Currency          PaymentListResponseCurrency `json:"currency,required"`
	Customer          CustomerLimitedDetails      `json:"customer,required"`
	Metadata          map[string]string           `json:"metadata,required"`
	PaymentID         string                      `json:"payment_id,required"`
	TotalAmount       int64                       `json:"total_amount,required"`
	PaymentMethod     string                      `json:"payment_method,nullable"`
	PaymentMethodType string                      `json:"payment_method_type,nullable"`
	Status            IntentStatus                `json:"status,nullable"`
	SubscriptionID    string                      `json:"subscription_id,nullable"`
	JSON              paymentListResponseJSON     `json:"-"`
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

type PaymentListResponseCurrency string

const (
	PaymentListResponseCurrencyAed PaymentListResponseCurrency = "AED"
	PaymentListResponseCurrencyAll PaymentListResponseCurrency = "ALL"
	PaymentListResponseCurrencyAmd PaymentListResponseCurrency = "AMD"
	PaymentListResponseCurrencyAng PaymentListResponseCurrency = "ANG"
	PaymentListResponseCurrencyAoa PaymentListResponseCurrency = "AOA"
	PaymentListResponseCurrencyArs PaymentListResponseCurrency = "ARS"
	PaymentListResponseCurrencyAud PaymentListResponseCurrency = "AUD"
	PaymentListResponseCurrencyAwg PaymentListResponseCurrency = "AWG"
	PaymentListResponseCurrencyAzn PaymentListResponseCurrency = "AZN"
	PaymentListResponseCurrencyBam PaymentListResponseCurrency = "BAM"
	PaymentListResponseCurrencyBbd PaymentListResponseCurrency = "BBD"
	PaymentListResponseCurrencyBdt PaymentListResponseCurrency = "BDT"
	PaymentListResponseCurrencyBgn PaymentListResponseCurrency = "BGN"
	PaymentListResponseCurrencyBhd PaymentListResponseCurrency = "BHD"
	PaymentListResponseCurrencyBif PaymentListResponseCurrency = "BIF"
	PaymentListResponseCurrencyBmd PaymentListResponseCurrency = "BMD"
	PaymentListResponseCurrencyBnd PaymentListResponseCurrency = "BND"
	PaymentListResponseCurrencyBob PaymentListResponseCurrency = "BOB"
	PaymentListResponseCurrencyBrl PaymentListResponseCurrency = "BRL"
	PaymentListResponseCurrencyBsd PaymentListResponseCurrency = "BSD"
	PaymentListResponseCurrencyBwp PaymentListResponseCurrency = "BWP"
	PaymentListResponseCurrencyByn PaymentListResponseCurrency = "BYN"
	PaymentListResponseCurrencyBzd PaymentListResponseCurrency = "BZD"
	PaymentListResponseCurrencyCad PaymentListResponseCurrency = "CAD"
	PaymentListResponseCurrencyChf PaymentListResponseCurrency = "CHF"
	PaymentListResponseCurrencyClp PaymentListResponseCurrency = "CLP"
	PaymentListResponseCurrencyCny PaymentListResponseCurrency = "CNY"
	PaymentListResponseCurrencyCop PaymentListResponseCurrency = "COP"
	PaymentListResponseCurrencyCrc PaymentListResponseCurrency = "CRC"
	PaymentListResponseCurrencyCup PaymentListResponseCurrency = "CUP"
	PaymentListResponseCurrencyCve PaymentListResponseCurrency = "CVE"
	PaymentListResponseCurrencyCzk PaymentListResponseCurrency = "CZK"
	PaymentListResponseCurrencyDjf PaymentListResponseCurrency = "DJF"
	PaymentListResponseCurrencyDkk PaymentListResponseCurrency = "DKK"
	PaymentListResponseCurrencyDop PaymentListResponseCurrency = "DOP"
	PaymentListResponseCurrencyDzd PaymentListResponseCurrency = "DZD"
	PaymentListResponseCurrencyEgp PaymentListResponseCurrency = "EGP"
	PaymentListResponseCurrencyEtb PaymentListResponseCurrency = "ETB"
	PaymentListResponseCurrencyEur PaymentListResponseCurrency = "EUR"
	PaymentListResponseCurrencyFjd PaymentListResponseCurrency = "FJD"
	PaymentListResponseCurrencyFkp PaymentListResponseCurrency = "FKP"
	PaymentListResponseCurrencyGbp PaymentListResponseCurrency = "GBP"
	PaymentListResponseCurrencyGel PaymentListResponseCurrency = "GEL"
	PaymentListResponseCurrencyGhs PaymentListResponseCurrency = "GHS"
	PaymentListResponseCurrencyGip PaymentListResponseCurrency = "GIP"
	PaymentListResponseCurrencyGmd PaymentListResponseCurrency = "GMD"
	PaymentListResponseCurrencyGnf PaymentListResponseCurrency = "GNF"
	PaymentListResponseCurrencyGtq PaymentListResponseCurrency = "GTQ"
	PaymentListResponseCurrencyGyd PaymentListResponseCurrency = "GYD"
	PaymentListResponseCurrencyHkd PaymentListResponseCurrency = "HKD"
	PaymentListResponseCurrencyHnl PaymentListResponseCurrency = "HNL"
	PaymentListResponseCurrencyHrk PaymentListResponseCurrency = "HRK"
	PaymentListResponseCurrencyHtg PaymentListResponseCurrency = "HTG"
	PaymentListResponseCurrencyHuf PaymentListResponseCurrency = "HUF"
	PaymentListResponseCurrencyIdr PaymentListResponseCurrency = "IDR"
	PaymentListResponseCurrencyIls PaymentListResponseCurrency = "ILS"
	PaymentListResponseCurrencyInr PaymentListResponseCurrency = "INR"
	PaymentListResponseCurrencyIqd PaymentListResponseCurrency = "IQD"
	PaymentListResponseCurrencyJmd PaymentListResponseCurrency = "JMD"
	PaymentListResponseCurrencyJod PaymentListResponseCurrency = "JOD"
	PaymentListResponseCurrencyJpy PaymentListResponseCurrency = "JPY"
	PaymentListResponseCurrencyKes PaymentListResponseCurrency = "KES"
	PaymentListResponseCurrencyKgs PaymentListResponseCurrency = "KGS"
	PaymentListResponseCurrencyKhr PaymentListResponseCurrency = "KHR"
	PaymentListResponseCurrencyKmf PaymentListResponseCurrency = "KMF"
	PaymentListResponseCurrencyKrw PaymentListResponseCurrency = "KRW"
	PaymentListResponseCurrencyKwd PaymentListResponseCurrency = "KWD"
	PaymentListResponseCurrencyKyd PaymentListResponseCurrency = "KYD"
	PaymentListResponseCurrencyKzt PaymentListResponseCurrency = "KZT"
	PaymentListResponseCurrencyLak PaymentListResponseCurrency = "LAK"
	PaymentListResponseCurrencyLbp PaymentListResponseCurrency = "LBP"
	PaymentListResponseCurrencyLkr PaymentListResponseCurrency = "LKR"
	PaymentListResponseCurrencyLrd PaymentListResponseCurrency = "LRD"
	PaymentListResponseCurrencyLsl PaymentListResponseCurrency = "LSL"
	PaymentListResponseCurrencyLyd PaymentListResponseCurrency = "LYD"
	PaymentListResponseCurrencyMad PaymentListResponseCurrency = "MAD"
	PaymentListResponseCurrencyMdl PaymentListResponseCurrency = "MDL"
	PaymentListResponseCurrencyMga PaymentListResponseCurrency = "MGA"
	PaymentListResponseCurrencyMkd PaymentListResponseCurrency = "MKD"
	PaymentListResponseCurrencyMmk PaymentListResponseCurrency = "MMK"
	PaymentListResponseCurrencyMnt PaymentListResponseCurrency = "MNT"
	PaymentListResponseCurrencyMop PaymentListResponseCurrency = "MOP"
	PaymentListResponseCurrencyMru PaymentListResponseCurrency = "MRU"
	PaymentListResponseCurrencyMur PaymentListResponseCurrency = "MUR"
	PaymentListResponseCurrencyMvr PaymentListResponseCurrency = "MVR"
	PaymentListResponseCurrencyMwk PaymentListResponseCurrency = "MWK"
	PaymentListResponseCurrencyMxn PaymentListResponseCurrency = "MXN"
	PaymentListResponseCurrencyMyr PaymentListResponseCurrency = "MYR"
	PaymentListResponseCurrencyMzn PaymentListResponseCurrency = "MZN"
	PaymentListResponseCurrencyNad PaymentListResponseCurrency = "NAD"
	PaymentListResponseCurrencyNgn PaymentListResponseCurrency = "NGN"
	PaymentListResponseCurrencyNio PaymentListResponseCurrency = "NIO"
	PaymentListResponseCurrencyNok PaymentListResponseCurrency = "NOK"
	PaymentListResponseCurrencyNpr PaymentListResponseCurrency = "NPR"
	PaymentListResponseCurrencyNzd PaymentListResponseCurrency = "NZD"
	PaymentListResponseCurrencyOmr PaymentListResponseCurrency = "OMR"
	PaymentListResponseCurrencyPab PaymentListResponseCurrency = "PAB"
	PaymentListResponseCurrencyPen PaymentListResponseCurrency = "PEN"
	PaymentListResponseCurrencyPgk PaymentListResponseCurrency = "PGK"
	PaymentListResponseCurrencyPhp PaymentListResponseCurrency = "PHP"
	PaymentListResponseCurrencyPkr PaymentListResponseCurrency = "PKR"
	PaymentListResponseCurrencyPln PaymentListResponseCurrency = "PLN"
	PaymentListResponseCurrencyPyg PaymentListResponseCurrency = "PYG"
	PaymentListResponseCurrencyQar PaymentListResponseCurrency = "QAR"
	PaymentListResponseCurrencyRon PaymentListResponseCurrency = "RON"
	PaymentListResponseCurrencyRsd PaymentListResponseCurrency = "RSD"
	PaymentListResponseCurrencyRub PaymentListResponseCurrency = "RUB"
	PaymentListResponseCurrencyRwf PaymentListResponseCurrency = "RWF"
	PaymentListResponseCurrencySar PaymentListResponseCurrency = "SAR"
	PaymentListResponseCurrencySbd PaymentListResponseCurrency = "SBD"
	PaymentListResponseCurrencyScr PaymentListResponseCurrency = "SCR"
	PaymentListResponseCurrencySek PaymentListResponseCurrency = "SEK"
	PaymentListResponseCurrencySgd PaymentListResponseCurrency = "SGD"
	PaymentListResponseCurrencyShp PaymentListResponseCurrency = "SHP"
	PaymentListResponseCurrencySle PaymentListResponseCurrency = "SLE"
	PaymentListResponseCurrencySll PaymentListResponseCurrency = "SLL"
	PaymentListResponseCurrencySos PaymentListResponseCurrency = "SOS"
	PaymentListResponseCurrencySrd PaymentListResponseCurrency = "SRD"
	PaymentListResponseCurrencySsp PaymentListResponseCurrency = "SSP"
	PaymentListResponseCurrencyStn PaymentListResponseCurrency = "STN"
	PaymentListResponseCurrencySvc PaymentListResponseCurrency = "SVC"
	PaymentListResponseCurrencySzl PaymentListResponseCurrency = "SZL"
	PaymentListResponseCurrencyThb PaymentListResponseCurrency = "THB"
	PaymentListResponseCurrencyTnd PaymentListResponseCurrency = "TND"
	PaymentListResponseCurrencyTop PaymentListResponseCurrency = "TOP"
	PaymentListResponseCurrencyTry PaymentListResponseCurrency = "TRY"
	PaymentListResponseCurrencyTtd PaymentListResponseCurrency = "TTD"
	PaymentListResponseCurrencyTwd PaymentListResponseCurrency = "TWD"
	PaymentListResponseCurrencyTzs PaymentListResponseCurrency = "TZS"
	PaymentListResponseCurrencyUah PaymentListResponseCurrency = "UAH"
	PaymentListResponseCurrencyUgx PaymentListResponseCurrency = "UGX"
	PaymentListResponseCurrencyUsd PaymentListResponseCurrency = "USD"
	PaymentListResponseCurrencyUyu PaymentListResponseCurrency = "UYU"
	PaymentListResponseCurrencyUzs PaymentListResponseCurrency = "UZS"
	PaymentListResponseCurrencyVes PaymentListResponseCurrency = "VES"
	PaymentListResponseCurrencyVnd PaymentListResponseCurrency = "VND"
	PaymentListResponseCurrencyVuv PaymentListResponseCurrency = "VUV"
	PaymentListResponseCurrencyWst PaymentListResponseCurrency = "WST"
	PaymentListResponseCurrencyXaf PaymentListResponseCurrency = "XAF"
	PaymentListResponseCurrencyXcd PaymentListResponseCurrency = "XCD"
	PaymentListResponseCurrencyXof PaymentListResponseCurrency = "XOF"
	PaymentListResponseCurrencyXpf PaymentListResponseCurrency = "XPF"
	PaymentListResponseCurrencyYer PaymentListResponseCurrency = "YER"
	PaymentListResponseCurrencyZar PaymentListResponseCurrency = "ZAR"
	PaymentListResponseCurrencyZmw PaymentListResponseCurrency = "ZMW"
)

func (r PaymentListResponseCurrency) IsKnown() bool {
	switch r {
	case PaymentListResponseCurrencyAed, PaymentListResponseCurrencyAll, PaymentListResponseCurrencyAmd, PaymentListResponseCurrencyAng, PaymentListResponseCurrencyAoa, PaymentListResponseCurrencyArs, PaymentListResponseCurrencyAud, PaymentListResponseCurrencyAwg, PaymentListResponseCurrencyAzn, PaymentListResponseCurrencyBam, PaymentListResponseCurrencyBbd, PaymentListResponseCurrencyBdt, PaymentListResponseCurrencyBgn, PaymentListResponseCurrencyBhd, PaymentListResponseCurrencyBif, PaymentListResponseCurrencyBmd, PaymentListResponseCurrencyBnd, PaymentListResponseCurrencyBob, PaymentListResponseCurrencyBrl, PaymentListResponseCurrencyBsd, PaymentListResponseCurrencyBwp, PaymentListResponseCurrencyByn, PaymentListResponseCurrencyBzd, PaymentListResponseCurrencyCad, PaymentListResponseCurrencyChf, PaymentListResponseCurrencyClp, PaymentListResponseCurrencyCny, PaymentListResponseCurrencyCop, PaymentListResponseCurrencyCrc, PaymentListResponseCurrencyCup, PaymentListResponseCurrencyCve, PaymentListResponseCurrencyCzk, PaymentListResponseCurrencyDjf, PaymentListResponseCurrencyDkk, PaymentListResponseCurrencyDop, PaymentListResponseCurrencyDzd, PaymentListResponseCurrencyEgp, PaymentListResponseCurrencyEtb, PaymentListResponseCurrencyEur, PaymentListResponseCurrencyFjd, PaymentListResponseCurrencyFkp, PaymentListResponseCurrencyGbp, PaymentListResponseCurrencyGel, PaymentListResponseCurrencyGhs, PaymentListResponseCurrencyGip, PaymentListResponseCurrencyGmd, PaymentListResponseCurrencyGnf, PaymentListResponseCurrencyGtq, PaymentListResponseCurrencyGyd, PaymentListResponseCurrencyHkd, PaymentListResponseCurrencyHnl, PaymentListResponseCurrencyHrk, PaymentListResponseCurrencyHtg, PaymentListResponseCurrencyHuf, PaymentListResponseCurrencyIdr, PaymentListResponseCurrencyIls, PaymentListResponseCurrencyInr, PaymentListResponseCurrencyIqd, PaymentListResponseCurrencyJmd, PaymentListResponseCurrencyJod, PaymentListResponseCurrencyJpy, PaymentListResponseCurrencyKes, PaymentListResponseCurrencyKgs, PaymentListResponseCurrencyKhr, PaymentListResponseCurrencyKmf, PaymentListResponseCurrencyKrw, PaymentListResponseCurrencyKwd, PaymentListResponseCurrencyKyd, PaymentListResponseCurrencyKzt, PaymentListResponseCurrencyLak, PaymentListResponseCurrencyLbp, PaymentListResponseCurrencyLkr, PaymentListResponseCurrencyLrd, PaymentListResponseCurrencyLsl, PaymentListResponseCurrencyLyd, PaymentListResponseCurrencyMad, PaymentListResponseCurrencyMdl, PaymentListResponseCurrencyMga, PaymentListResponseCurrencyMkd, PaymentListResponseCurrencyMmk, PaymentListResponseCurrencyMnt, PaymentListResponseCurrencyMop, PaymentListResponseCurrencyMru, PaymentListResponseCurrencyMur, PaymentListResponseCurrencyMvr, PaymentListResponseCurrencyMwk, PaymentListResponseCurrencyMxn, PaymentListResponseCurrencyMyr, PaymentListResponseCurrencyMzn, PaymentListResponseCurrencyNad, PaymentListResponseCurrencyNgn, PaymentListResponseCurrencyNio, PaymentListResponseCurrencyNok, PaymentListResponseCurrencyNpr, PaymentListResponseCurrencyNzd, PaymentListResponseCurrencyOmr, PaymentListResponseCurrencyPab, PaymentListResponseCurrencyPen, PaymentListResponseCurrencyPgk, PaymentListResponseCurrencyPhp, PaymentListResponseCurrencyPkr, PaymentListResponseCurrencyPln, PaymentListResponseCurrencyPyg, PaymentListResponseCurrencyQar, PaymentListResponseCurrencyRon, PaymentListResponseCurrencyRsd, PaymentListResponseCurrencyRub, PaymentListResponseCurrencyRwf, PaymentListResponseCurrencySar, PaymentListResponseCurrencySbd, PaymentListResponseCurrencyScr, PaymentListResponseCurrencySek, PaymentListResponseCurrencySgd, PaymentListResponseCurrencyShp, PaymentListResponseCurrencySle, PaymentListResponseCurrencySll, PaymentListResponseCurrencySos, PaymentListResponseCurrencySrd, PaymentListResponseCurrencySsp, PaymentListResponseCurrencyStn, PaymentListResponseCurrencySvc, PaymentListResponseCurrencySzl, PaymentListResponseCurrencyThb, PaymentListResponseCurrencyTnd, PaymentListResponseCurrencyTop, PaymentListResponseCurrencyTry, PaymentListResponseCurrencyTtd, PaymentListResponseCurrencyTwd, PaymentListResponseCurrencyTzs, PaymentListResponseCurrencyUah, PaymentListResponseCurrencyUgx, PaymentListResponseCurrencyUsd, PaymentListResponseCurrencyUyu, PaymentListResponseCurrencyUzs, PaymentListResponseCurrencyVes, PaymentListResponseCurrencyVnd, PaymentListResponseCurrencyVuv, PaymentListResponseCurrencyWst, PaymentListResponseCurrencyXaf, PaymentListResponseCurrencyXcd, PaymentListResponseCurrencyXof, PaymentListResponseCurrencyXpf, PaymentListResponseCurrencyYer, PaymentListResponseCurrencyZar, PaymentListResponseCurrencyZmw:
		return true
	}
	return false
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
	BillingCurrency           param.Field[PaymentNewParamsBillingCurrency]            `json:"billing_currency"`
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

type PaymentNewParamsBillingCurrency string

const (
	PaymentNewParamsBillingCurrencyAed PaymentNewParamsBillingCurrency = "AED"
	PaymentNewParamsBillingCurrencyAll PaymentNewParamsBillingCurrency = "ALL"
	PaymentNewParamsBillingCurrencyAmd PaymentNewParamsBillingCurrency = "AMD"
	PaymentNewParamsBillingCurrencyAng PaymentNewParamsBillingCurrency = "ANG"
	PaymentNewParamsBillingCurrencyAoa PaymentNewParamsBillingCurrency = "AOA"
	PaymentNewParamsBillingCurrencyArs PaymentNewParamsBillingCurrency = "ARS"
	PaymentNewParamsBillingCurrencyAud PaymentNewParamsBillingCurrency = "AUD"
	PaymentNewParamsBillingCurrencyAwg PaymentNewParamsBillingCurrency = "AWG"
	PaymentNewParamsBillingCurrencyAzn PaymentNewParamsBillingCurrency = "AZN"
	PaymentNewParamsBillingCurrencyBam PaymentNewParamsBillingCurrency = "BAM"
	PaymentNewParamsBillingCurrencyBbd PaymentNewParamsBillingCurrency = "BBD"
	PaymentNewParamsBillingCurrencyBdt PaymentNewParamsBillingCurrency = "BDT"
	PaymentNewParamsBillingCurrencyBgn PaymentNewParamsBillingCurrency = "BGN"
	PaymentNewParamsBillingCurrencyBhd PaymentNewParamsBillingCurrency = "BHD"
	PaymentNewParamsBillingCurrencyBif PaymentNewParamsBillingCurrency = "BIF"
	PaymentNewParamsBillingCurrencyBmd PaymentNewParamsBillingCurrency = "BMD"
	PaymentNewParamsBillingCurrencyBnd PaymentNewParamsBillingCurrency = "BND"
	PaymentNewParamsBillingCurrencyBob PaymentNewParamsBillingCurrency = "BOB"
	PaymentNewParamsBillingCurrencyBrl PaymentNewParamsBillingCurrency = "BRL"
	PaymentNewParamsBillingCurrencyBsd PaymentNewParamsBillingCurrency = "BSD"
	PaymentNewParamsBillingCurrencyBwp PaymentNewParamsBillingCurrency = "BWP"
	PaymentNewParamsBillingCurrencyByn PaymentNewParamsBillingCurrency = "BYN"
	PaymentNewParamsBillingCurrencyBzd PaymentNewParamsBillingCurrency = "BZD"
	PaymentNewParamsBillingCurrencyCad PaymentNewParamsBillingCurrency = "CAD"
	PaymentNewParamsBillingCurrencyChf PaymentNewParamsBillingCurrency = "CHF"
	PaymentNewParamsBillingCurrencyClp PaymentNewParamsBillingCurrency = "CLP"
	PaymentNewParamsBillingCurrencyCny PaymentNewParamsBillingCurrency = "CNY"
	PaymentNewParamsBillingCurrencyCop PaymentNewParamsBillingCurrency = "COP"
	PaymentNewParamsBillingCurrencyCrc PaymentNewParamsBillingCurrency = "CRC"
	PaymentNewParamsBillingCurrencyCup PaymentNewParamsBillingCurrency = "CUP"
	PaymentNewParamsBillingCurrencyCve PaymentNewParamsBillingCurrency = "CVE"
	PaymentNewParamsBillingCurrencyCzk PaymentNewParamsBillingCurrency = "CZK"
	PaymentNewParamsBillingCurrencyDjf PaymentNewParamsBillingCurrency = "DJF"
	PaymentNewParamsBillingCurrencyDkk PaymentNewParamsBillingCurrency = "DKK"
	PaymentNewParamsBillingCurrencyDop PaymentNewParamsBillingCurrency = "DOP"
	PaymentNewParamsBillingCurrencyDzd PaymentNewParamsBillingCurrency = "DZD"
	PaymentNewParamsBillingCurrencyEgp PaymentNewParamsBillingCurrency = "EGP"
	PaymentNewParamsBillingCurrencyEtb PaymentNewParamsBillingCurrency = "ETB"
	PaymentNewParamsBillingCurrencyEur PaymentNewParamsBillingCurrency = "EUR"
	PaymentNewParamsBillingCurrencyFjd PaymentNewParamsBillingCurrency = "FJD"
	PaymentNewParamsBillingCurrencyFkp PaymentNewParamsBillingCurrency = "FKP"
	PaymentNewParamsBillingCurrencyGbp PaymentNewParamsBillingCurrency = "GBP"
	PaymentNewParamsBillingCurrencyGel PaymentNewParamsBillingCurrency = "GEL"
	PaymentNewParamsBillingCurrencyGhs PaymentNewParamsBillingCurrency = "GHS"
	PaymentNewParamsBillingCurrencyGip PaymentNewParamsBillingCurrency = "GIP"
	PaymentNewParamsBillingCurrencyGmd PaymentNewParamsBillingCurrency = "GMD"
	PaymentNewParamsBillingCurrencyGnf PaymentNewParamsBillingCurrency = "GNF"
	PaymentNewParamsBillingCurrencyGtq PaymentNewParamsBillingCurrency = "GTQ"
	PaymentNewParamsBillingCurrencyGyd PaymentNewParamsBillingCurrency = "GYD"
	PaymentNewParamsBillingCurrencyHkd PaymentNewParamsBillingCurrency = "HKD"
	PaymentNewParamsBillingCurrencyHnl PaymentNewParamsBillingCurrency = "HNL"
	PaymentNewParamsBillingCurrencyHrk PaymentNewParamsBillingCurrency = "HRK"
	PaymentNewParamsBillingCurrencyHtg PaymentNewParamsBillingCurrency = "HTG"
	PaymentNewParamsBillingCurrencyHuf PaymentNewParamsBillingCurrency = "HUF"
	PaymentNewParamsBillingCurrencyIdr PaymentNewParamsBillingCurrency = "IDR"
	PaymentNewParamsBillingCurrencyIls PaymentNewParamsBillingCurrency = "ILS"
	PaymentNewParamsBillingCurrencyInr PaymentNewParamsBillingCurrency = "INR"
	PaymentNewParamsBillingCurrencyIqd PaymentNewParamsBillingCurrency = "IQD"
	PaymentNewParamsBillingCurrencyJmd PaymentNewParamsBillingCurrency = "JMD"
	PaymentNewParamsBillingCurrencyJod PaymentNewParamsBillingCurrency = "JOD"
	PaymentNewParamsBillingCurrencyJpy PaymentNewParamsBillingCurrency = "JPY"
	PaymentNewParamsBillingCurrencyKes PaymentNewParamsBillingCurrency = "KES"
	PaymentNewParamsBillingCurrencyKgs PaymentNewParamsBillingCurrency = "KGS"
	PaymentNewParamsBillingCurrencyKhr PaymentNewParamsBillingCurrency = "KHR"
	PaymentNewParamsBillingCurrencyKmf PaymentNewParamsBillingCurrency = "KMF"
	PaymentNewParamsBillingCurrencyKrw PaymentNewParamsBillingCurrency = "KRW"
	PaymentNewParamsBillingCurrencyKwd PaymentNewParamsBillingCurrency = "KWD"
	PaymentNewParamsBillingCurrencyKyd PaymentNewParamsBillingCurrency = "KYD"
	PaymentNewParamsBillingCurrencyKzt PaymentNewParamsBillingCurrency = "KZT"
	PaymentNewParamsBillingCurrencyLak PaymentNewParamsBillingCurrency = "LAK"
	PaymentNewParamsBillingCurrencyLbp PaymentNewParamsBillingCurrency = "LBP"
	PaymentNewParamsBillingCurrencyLkr PaymentNewParamsBillingCurrency = "LKR"
	PaymentNewParamsBillingCurrencyLrd PaymentNewParamsBillingCurrency = "LRD"
	PaymentNewParamsBillingCurrencyLsl PaymentNewParamsBillingCurrency = "LSL"
	PaymentNewParamsBillingCurrencyLyd PaymentNewParamsBillingCurrency = "LYD"
	PaymentNewParamsBillingCurrencyMad PaymentNewParamsBillingCurrency = "MAD"
	PaymentNewParamsBillingCurrencyMdl PaymentNewParamsBillingCurrency = "MDL"
	PaymentNewParamsBillingCurrencyMga PaymentNewParamsBillingCurrency = "MGA"
	PaymentNewParamsBillingCurrencyMkd PaymentNewParamsBillingCurrency = "MKD"
	PaymentNewParamsBillingCurrencyMmk PaymentNewParamsBillingCurrency = "MMK"
	PaymentNewParamsBillingCurrencyMnt PaymentNewParamsBillingCurrency = "MNT"
	PaymentNewParamsBillingCurrencyMop PaymentNewParamsBillingCurrency = "MOP"
	PaymentNewParamsBillingCurrencyMru PaymentNewParamsBillingCurrency = "MRU"
	PaymentNewParamsBillingCurrencyMur PaymentNewParamsBillingCurrency = "MUR"
	PaymentNewParamsBillingCurrencyMvr PaymentNewParamsBillingCurrency = "MVR"
	PaymentNewParamsBillingCurrencyMwk PaymentNewParamsBillingCurrency = "MWK"
	PaymentNewParamsBillingCurrencyMxn PaymentNewParamsBillingCurrency = "MXN"
	PaymentNewParamsBillingCurrencyMyr PaymentNewParamsBillingCurrency = "MYR"
	PaymentNewParamsBillingCurrencyMzn PaymentNewParamsBillingCurrency = "MZN"
	PaymentNewParamsBillingCurrencyNad PaymentNewParamsBillingCurrency = "NAD"
	PaymentNewParamsBillingCurrencyNgn PaymentNewParamsBillingCurrency = "NGN"
	PaymentNewParamsBillingCurrencyNio PaymentNewParamsBillingCurrency = "NIO"
	PaymentNewParamsBillingCurrencyNok PaymentNewParamsBillingCurrency = "NOK"
	PaymentNewParamsBillingCurrencyNpr PaymentNewParamsBillingCurrency = "NPR"
	PaymentNewParamsBillingCurrencyNzd PaymentNewParamsBillingCurrency = "NZD"
	PaymentNewParamsBillingCurrencyOmr PaymentNewParamsBillingCurrency = "OMR"
	PaymentNewParamsBillingCurrencyPab PaymentNewParamsBillingCurrency = "PAB"
	PaymentNewParamsBillingCurrencyPen PaymentNewParamsBillingCurrency = "PEN"
	PaymentNewParamsBillingCurrencyPgk PaymentNewParamsBillingCurrency = "PGK"
	PaymentNewParamsBillingCurrencyPhp PaymentNewParamsBillingCurrency = "PHP"
	PaymentNewParamsBillingCurrencyPkr PaymentNewParamsBillingCurrency = "PKR"
	PaymentNewParamsBillingCurrencyPln PaymentNewParamsBillingCurrency = "PLN"
	PaymentNewParamsBillingCurrencyPyg PaymentNewParamsBillingCurrency = "PYG"
	PaymentNewParamsBillingCurrencyQar PaymentNewParamsBillingCurrency = "QAR"
	PaymentNewParamsBillingCurrencyRon PaymentNewParamsBillingCurrency = "RON"
	PaymentNewParamsBillingCurrencyRsd PaymentNewParamsBillingCurrency = "RSD"
	PaymentNewParamsBillingCurrencyRub PaymentNewParamsBillingCurrency = "RUB"
	PaymentNewParamsBillingCurrencyRwf PaymentNewParamsBillingCurrency = "RWF"
	PaymentNewParamsBillingCurrencySar PaymentNewParamsBillingCurrency = "SAR"
	PaymentNewParamsBillingCurrencySbd PaymentNewParamsBillingCurrency = "SBD"
	PaymentNewParamsBillingCurrencyScr PaymentNewParamsBillingCurrency = "SCR"
	PaymentNewParamsBillingCurrencySek PaymentNewParamsBillingCurrency = "SEK"
	PaymentNewParamsBillingCurrencySgd PaymentNewParamsBillingCurrency = "SGD"
	PaymentNewParamsBillingCurrencyShp PaymentNewParamsBillingCurrency = "SHP"
	PaymentNewParamsBillingCurrencySle PaymentNewParamsBillingCurrency = "SLE"
	PaymentNewParamsBillingCurrencySll PaymentNewParamsBillingCurrency = "SLL"
	PaymentNewParamsBillingCurrencySos PaymentNewParamsBillingCurrency = "SOS"
	PaymentNewParamsBillingCurrencySrd PaymentNewParamsBillingCurrency = "SRD"
	PaymentNewParamsBillingCurrencySsp PaymentNewParamsBillingCurrency = "SSP"
	PaymentNewParamsBillingCurrencyStn PaymentNewParamsBillingCurrency = "STN"
	PaymentNewParamsBillingCurrencySvc PaymentNewParamsBillingCurrency = "SVC"
	PaymentNewParamsBillingCurrencySzl PaymentNewParamsBillingCurrency = "SZL"
	PaymentNewParamsBillingCurrencyThb PaymentNewParamsBillingCurrency = "THB"
	PaymentNewParamsBillingCurrencyTnd PaymentNewParamsBillingCurrency = "TND"
	PaymentNewParamsBillingCurrencyTop PaymentNewParamsBillingCurrency = "TOP"
	PaymentNewParamsBillingCurrencyTry PaymentNewParamsBillingCurrency = "TRY"
	PaymentNewParamsBillingCurrencyTtd PaymentNewParamsBillingCurrency = "TTD"
	PaymentNewParamsBillingCurrencyTwd PaymentNewParamsBillingCurrency = "TWD"
	PaymentNewParamsBillingCurrencyTzs PaymentNewParamsBillingCurrency = "TZS"
	PaymentNewParamsBillingCurrencyUah PaymentNewParamsBillingCurrency = "UAH"
	PaymentNewParamsBillingCurrencyUgx PaymentNewParamsBillingCurrency = "UGX"
	PaymentNewParamsBillingCurrencyUsd PaymentNewParamsBillingCurrency = "USD"
	PaymentNewParamsBillingCurrencyUyu PaymentNewParamsBillingCurrency = "UYU"
	PaymentNewParamsBillingCurrencyUzs PaymentNewParamsBillingCurrency = "UZS"
	PaymentNewParamsBillingCurrencyVes PaymentNewParamsBillingCurrency = "VES"
	PaymentNewParamsBillingCurrencyVnd PaymentNewParamsBillingCurrency = "VND"
	PaymentNewParamsBillingCurrencyVuv PaymentNewParamsBillingCurrency = "VUV"
	PaymentNewParamsBillingCurrencyWst PaymentNewParamsBillingCurrency = "WST"
	PaymentNewParamsBillingCurrencyXaf PaymentNewParamsBillingCurrency = "XAF"
	PaymentNewParamsBillingCurrencyXcd PaymentNewParamsBillingCurrency = "XCD"
	PaymentNewParamsBillingCurrencyXof PaymentNewParamsBillingCurrency = "XOF"
	PaymentNewParamsBillingCurrencyXpf PaymentNewParamsBillingCurrency = "XPF"
	PaymentNewParamsBillingCurrencyYer PaymentNewParamsBillingCurrency = "YER"
	PaymentNewParamsBillingCurrencyZar PaymentNewParamsBillingCurrency = "ZAR"
	PaymentNewParamsBillingCurrencyZmw PaymentNewParamsBillingCurrency = "ZMW"
)

func (r PaymentNewParamsBillingCurrency) IsKnown() bool {
	switch r {
	case PaymentNewParamsBillingCurrencyAed, PaymentNewParamsBillingCurrencyAll, PaymentNewParamsBillingCurrencyAmd, PaymentNewParamsBillingCurrencyAng, PaymentNewParamsBillingCurrencyAoa, PaymentNewParamsBillingCurrencyArs, PaymentNewParamsBillingCurrencyAud, PaymentNewParamsBillingCurrencyAwg, PaymentNewParamsBillingCurrencyAzn, PaymentNewParamsBillingCurrencyBam, PaymentNewParamsBillingCurrencyBbd, PaymentNewParamsBillingCurrencyBdt, PaymentNewParamsBillingCurrencyBgn, PaymentNewParamsBillingCurrencyBhd, PaymentNewParamsBillingCurrencyBif, PaymentNewParamsBillingCurrencyBmd, PaymentNewParamsBillingCurrencyBnd, PaymentNewParamsBillingCurrencyBob, PaymentNewParamsBillingCurrencyBrl, PaymentNewParamsBillingCurrencyBsd, PaymentNewParamsBillingCurrencyBwp, PaymentNewParamsBillingCurrencyByn, PaymentNewParamsBillingCurrencyBzd, PaymentNewParamsBillingCurrencyCad, PaymentNewParamsBillingCurrencyChf, PaymentNewParamsBillingCurrencyClp, PaymentNewParamsBillingCurrencyCny, PaymentNewParamsBillingCurrencyCop, PaymentNewParamsBillingCurrencyCrc, PaymentNewParamsBillingCurrencyCup, PaymentNewParamsBillingCurrencyCve, PaymentNewParamsBillingCurrencyCzk, PaymentNewParamsBillingCurrencyDjf, PaymentNewParamsBillingCurrencyDkk, PaymentNewParamsBillingCurrencyDop, PaymentNewParamsBillingCurrencyDzd, PaymentNewParamsBillingCurrencyEgp, PaymentNewParamsBillingCurrencyEtb, PaymentNewParamsBillingCurrencyEur, PaymentNewParamsBillingCurrencyFjd, PaymentNewParamsBillingCurrencyFkp, PaymentNewParamsBillingCurrencyGbp, PaymentNewParamsBillingCurrencyGel, PaymentNewParamsBillingCurrencyGhs, PaymentNewParamsBillingCurrencyGip, PaymentNewParamsBillingCurrencyGmd, PaymentNewParamsBillingCurrencyGnf, PaymentNewParamsBillingCurrencyGtq, PaymentNewParamsBillingCurrencyGyd, PaymentNewParamsBillingCurrencyHkd, PaymentNewParamsBillingCurrencyHnl, PaymentNewParamsBillingCurrencyHrk, PaymentNewParamsBillingCurrencyHtg, PaymentNewParamsBillingCurrencyHuf, PaymentNewParamsBillingCurrencyIdr, PaymentNewParamsBillingCurrencyIls, PaymentNewParamsBillingCurrencyInr, PaymentNewParamsBillingCurrencyIqd, PaymentNewParamsBillingCurrencyJmd, PaymentNewParamsBillingCurrencyJod, PaymentNewParamsBillingCurrencyJpy, PaymentNewParamsBillingCurrencyKes, PaymentNewParamsBillingCurrencyKgs, PaymentNewParamsBillingCurrencyKhr, PaymentNewParamsBillingCurrencyKmf, PaymentNewParamsBillingCurrencyKrw, PaymentNewParamsBillingCurrencyKwd, PaymentNewParamsBillingCurrencyKyd, PaymentNewParamsBillingCurrencyKzt, PaymentNewParamsBillingCurrencyLak, PaymentNewParamsBillingCurrencyLbp, PaymentNewParamsBillingCurrencyLkr, PaymentNewParamsBillingCurrencyLrd, PaymentNewParamsBillingCurrencyLsl, PaymentNewParamsBillingCurrencyLyd, PaymentNewParamsBillingCurrencyMad, PaymentNewParamsBillingCurrencyMdl, PaymentNewParamsBillingCurrencyMga, PaymentNewParamsBillingCurrencyMkd, PaymentNewParamsBillingCurrencyMmk, PaymentNewParamsBillingCurrencyMnt, PaymentNewParamsBillingCurrencyMop, PaymentNewParamsBillingCurrencyMru, PaymentNewParamsBillingCurrencyMur, PaymentNewParamsBillingCurrencyMvr, PaymentNewParamsBillingCurrencyMwk, PaymentNewParamsBillingCurrencyMxn, PaymentNewParamsBillingCurrencyMyr, PaymentNewParamsBillingCurrencyMzn, PaymentNewParamsBillingCurrencyNad, PaymentNewParamsBillingCurrencyNgn, PaymentNewParamsBillingCurrencyNio, PaymentNewParamsBillingCurrencyNok, PaymentNewParamsBillingCurrencyNpr, PaymentNewParamsBillingCurrencyNzd, PaymentNewParamsBillingCurrencyOmr, PaymentNewParamsBillingCurrencyPab, PaymentNewParamsBillingCurrencyPen, PaymentNewParamsBillingCurrencyPgk, PaymentNewParamsBillingCurrencyPhp, PaymentNewParamsBillingCurrencyPkr, PaymentNewParamsBillingCurrencyPln, PaymentNewParamsBillingCurrencyPyg, PaymentNewParamsBillingCurrencyQar, PaymentNewParamsBillingCurrencyRon, PaymentNewParamsBillingCurrencyRsd, PaymentNewParamsBillingCurrencyRub, PaymentNewParamsBillingCurrencyRwf, PaymentNewParamsBillingCurrencySar, PaymentNewParamsBillingCurrencySbd, PaymentNewParamsBillingCurrencyScr, PaymentNewParamsBillingCurrencySek, PaymentNewParamsBillingCurrencySgd, PaymentNewParamsBillingCurrencyShp, PaymentNewParamsBillingCurrencySle, PaymentNewParamsBillingCurrencySll, PaymentNewParamsBillingCurrencySos, PaymentNewParamsBillingCurrencySrd, PaymentNewParamsBillingCurrencySsp, PaymentNewParamsBillingCurrencyStn, PaymentNewParamsBillingCurrencySvc, PaymentNewParamsBillingCurrencySzl, PaymentNewParamsBillingCurrencyThb, PaymentNewParamsBillingCurrencyTnd, PaymentNewParamsBillingCurrencyTop, PaymentNewParamsBillingCurrencyTry, PaymentNewParamsBillingCurrencyTtd, PaymentNewParamsBillingCurrencyTwd, PaymentNewParamsBillingCurrencyTzs, PaymentNewParamsBillingCurrencyUah, PaymentNewParamsBillingCurrencyUgx, PaymentNewParamsBillingCurrencyUsd, PaymentNewParamsBillingCurrencyUyu, PaymentNewParamsBillingCurrencyUzs, PaymentNewParamsBillingCurrencyVes, PaymentNewParamsBillingCurrencyVnd, PaymentNewParamsBillingCurrencyVuv, PaymentNewParamsBillingCurrencyWst, PaymentNewParamsBillingCurrencyXaf, PaymentNewParamsBillingCurrencyXcd, PaymentNewParamsBillingCurrencyXof, PaymentNewParamsBillingCurrencyXpf, PaymentNewParamsBillingCurrencyYer, PaymentNewParamsBillingCurrencyZar, PaymentNewParamsBillingCurrencyZmw:
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
