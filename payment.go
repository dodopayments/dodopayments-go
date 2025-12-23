// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
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
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
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

// The property CustomerID is required.
type AttachExistingCustomerParam struct {
	CustomerID string `json:"customer_id,required"`
	paramObj
}

func (r AttachExistingCustomerParam) MarshalJSON() (data []byte, err error) {
	type shadow AttachExistingCustomerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttachExistingCustomerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingAddress struct {
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
	Country CountryCode `json:"country,required"`
	// City name
	City string `json:"city,nullable"`
	// State or province name
	State string `json:"state,nullable"`
	// Street address including house number and unit/apartment if applicable
	Street string `json:"street,nullable"`
	// Postal code or ZIP code
	Zipcode string `json:"zipcode,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		City        respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Zipcode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingAddress) RawJSON() string { return r.JSON.raw }
func (r *BillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BillingAddress to a BillingAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BillingAddressParam.Overrides()
func (r BillingAddress) ToParam() BillingAddressParam {
	return param.Override[BillingAddressParam](json.RawMessage(r.RawJSON()))
}

// The property Country is required.
type BillingAddressParam struct {
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

func (r BillingAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow BillingAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillingAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerLimitedDetails struct {
	// Unique identifier for the customer
	CustomerID string `json:"customer_id,required"`
	// Email address of the customer
	Email string `json:"email,required"`
	// Full name of the customer
	Name string `json:"name,required"`
	// Additional metadata associated with the customer
	Metadata map[string]string `json:"metadata"`
	// Phone number of the customer
	PhoneNumber string `json:"phone_number,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomerID  respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		Metadata    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerLimitedDetails) RawJSON() string { return r.JSON.raw }
func (r *CustomerLimitedDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func CustomerRequestParamOfAttachExistingCustomer(customerID string) CustomerRequestUnionParam {
	var variant AttachExistingCustomerParam
	variant.CustomerID = customerID
	return CustomerRequestUnionParam{OfAttachExistingCustomer: &variant}
}

func CustomerRequestParamOfNewCustomer(email string) CustomerRequestUnionParam {
	var variant NewCustomerParam
	variant.Email = email
	return CustomerRequestUnionParam{OfNewCustomer: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomerRequestUnionParam struct {
	OfAttachExistingCustomer *AttachExistingCustomerParam `json:",omitzero,inline"`
	OfNewCustomer            *NewCustomerParam            `json:",omitzero,inline"`
	paramUnion
}

func (u CustomerRequestUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttachExistingCustomer, u.OfNewCustomer)
}
func (u *CustomerRequestUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomerRequestUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAttachExistingCustomer) {
		return u.OfAttachExistingCustomer
	} else if !param.IsOmitted(u.OfNewCustomer) {
		return u.OfNewCustomer
	}
	return nil
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

// The property Email is required.
type NewCustomerParam struct {
	// Email is required for creating a new customer
	Email string `json:"email,required"`
	// Optional full name of the customer. If provided during session creation, it is
	// persisted and becomes immutable for the session. If omitted here, it can be
	// provided later via the confirm API.
	Name        param.Opt[string] `json:"name,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	paramObj
}

func (r NewCustomerParam) MarshalJSON() (data []byte, err error) {
	type shadow NewCustomerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewCustomerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OneTimeProductCartItem struct {
	ProductID string `json:"product_id,required"`
	Quantity  int64  `json:"quantity,required"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Amount int64 `json:"amount,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductID   respjson.Field
		Quantity    respjson.Field
		Amount      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OneTimeProductCartItem) RawJSON() string { return r.JSON.raw }
func (r *OneTimeProductCartItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OneTimeProductCartItem to a OneTimeProductCartItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OneTimeProductCartItemParam.Overrides()
func (r OneTimeProductCartItem) ToParam() OneTimeProductCartItemParam {
	return param.Override[OneTimeProductCartItemParam](json.RawMessage(r.RawJSON()))
}

// The properties ProductID, Quantity are required.
type OneTimeProductCartItemParam struct {
	ProductID string `json:"product_id,required"`
	Quantity  int64  `json:"quantity,required"`
	// Amount the customer pays if pay_what_you_want is enabled. If disabled then
	// amount will be ignored Represented in the lowest denomination of the currency
	// (e.g., cents for USD). For example, to charge $1.00, pass `100`.
	Amount param.Opt[int64] `json:"amount,omitzero"`
	paramObj
}

func (r OneTimeProductCartItemParam) MarshalJSON() (data []byte, err error) {
	type shadow OneTimeProductCartItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OneTimeProductCartItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Payment struct {
	// Billing address details for payments
	Billing BillingAddress `json:"billing,required"`
	// brand id this payment belongs to
	BrandID string `json:"brand_id,required"`
	// Identifier of the business associated with the payment
	BusinessID string `json:"business_id,required"`
	// Timestamp when the payment was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Currency used for the payment
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
	Currency Currency `json:"currency,required"`
	// Details about the customer who made the payment
	Customer CustomerLimitedDetails `json:"customer,required"`
	// brand id this payment belongs to
	DigitalProductsDelivered bool `json:"digital_products_delivered,required"`
	// List of disputes associated with this payment
	Disputes []Dispute `json:"disputes,required"`
	// Additional custom data associated with the payment
	Metadata map[string]string `json:"metadata,required"`
	// Unique identifier for the payment
	PaymentID string `json:"payment_id,required"`
	// List of refunds issued for this payment
	Refunds []PaymentRefund `json:"refunds,required"`
	// The amount that will be credited to your Dodo balance after currency conversion
	// and processing. Especially relevant for adaptive pricing where the customer's
	// payment currency differs from your settlement currency.
	SettlementAmount int64 `json:"settlement_amount,required"`
	// The currency in which the settlement_amount will be credited to your Dodo
	// balance. This may differ from the customer's payment currency in adaptive
	// pricing scenarios.
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
	SettlementCurrency Currency `json:"settlement_currency,required"`
	// Total amount charged to the customer including tax, in smallest currency unit
	// (e.g. cents)
	TotalAmount int64 `json:"total_amount,required"`
	// ISO2 country code of the card
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
	CardIssuingCountry CountryCode `json:"card_issuing_country,nullable"`
	// The last four digits of the card
	CardLastFour string `json:"card_last_four,nullable"`
	// Card network like VISA, MASTERCARD etc.
	CardNetwork string `json:"card_network,nullable"`
	// The type of card DEBIT or CREDIT
	CardType string `json:"card_type,nullable"`
	// If payment is made using a checkout session, this field is set to the id of the
	// session.
	CheckoutSessionID string `json:"checkout_session_id,nullable"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id,nullable"`
	// An error code if the payment failed
	ErrorCode string `json:"error_code,nullable"`
	// An error message if the payment failed
	ErrorMessage string `json:"error_message,nullable"`
	// Invoice ID for this payment. Uses India-specific invoice ID if available.
	InvoiceID string `json:"invoice_id,nullable"`
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
	SettlementTax int64 `json:"settlement_tax,nullable"`
	// Current status of the payment intent
	//
	// Any of "succeeded", "failed", "cancelled", "processing",
	// "requires_customer_action", "requires_merchant_action",
	// "requires_payment_method", "requires_confirmation", "requires_capture",
	// "partially_captured", "partially_captured_and_capturable".
	Status IntentStatus `json:"status,nullable"`
	// Identifier of the subscription if payment is part of a subscription
	SubscriptionID string `json:"subscription_id,nullable"`
	// Amount of tax collected in smallest currency unit (e.g. cents)
	Tax int64 `json:"tax,nullable"`
	// Timestamp when the payment was last updated
	UpdatedAt time.Time `json:"updated_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing                  respjson.Field
		BrandID                  respjson.Field
		BusinessID               respjson.Field
		CreatedAt                respjson.Field
		Currency                 respjson.Field
		Customer                 respjson.Field
		DigitalProductsDelivered respjson.Field
		Disputes                 respjson.Field
		Metadata                 respjson.Field
		PaymentID                respjson.Field
		Refunds                  respjson.Field
		SettlementAmount         respjson.Field
		SettlementCurrency       respjson.Field
		TotalAmount              respjson.Field
		CardIssuingCountry       respjson.Field
		CardLastFour             respjson.Field
		CardNetwork              respjson.Field
		CardType                 respjson.Field
		CheckoutSessionID        respjson.Field
		DiscountID               respjson.Field
		ErrorCode                respjson.Field
		ErrorMessage             respjson.Field
		InvoiceID                respjson.Field
		PaymentLink              respjson.Field
		PaymentMethod            respjson.Field
		PaymentMethodType        respjson.Field
		ProductCart              respjson.Field
		SettlementTax            respjson.Field
		Status                   respjson.Field
		SubscriptionID           respjson.Field
		Tax                      respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Payment) RawJSON() string { return r.JSON.raw }
func (r *Payment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentRefund struct {
	// The unique identifier of the business issuing the refund.
	BusinessID string `json:"business_id,required"`
	// The timestamp of when the refund was created in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If true the refund is a partial refund
	IsPartial bool `json:"is_partial,required"`
	// The unique identifier of the payment associated with the refund.
	PaymentID string `json:"payment_id,required"`
	// The unique identifier of the refund.
	RefundID string `json:"refund_id,required"`
	// The current status of the refund.
	//
	// Any of "succeeded", "failed", "pending", "review".
	Status RefundStatus `json:"status,required"`
	// The refunded amount.
	Amount int64 `json:"amount,nullable"`
	// The currency of the refund, represented as an ISO 4217 currency code.
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
	Currency Currency `json:"currency,nullable"`
	// The reason provided for the refund, if any. Optional.
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		CreatedAt   respjson.Field
		IsPartial   respjson.Field
		PaymentID   respjson.Field
		RefundID    respjson.Field
		Status      respjson.Field
		Amount      respjson.Field
		Currency    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentRefund) RawJSON() string { return r.JSON.raw }
func (r *PaymentRefund) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentProductCart struct {
	ProductID string `json:"product_id,required"`
	Quantity  int64  `json:"quantity,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductID   respjson.Field
		Quantity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentProductCart) RawJSON() string { return r.JSON.raw }
func (r *PaymentProductCart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentMethodTypes string

const (
	PaymentMethodTypesCredit           PaymentMethodTypes = "credit"
	PaymentMethodTypesDebit            PaymentMethodTypes = "debit"
	PaymentMethodTypesUpiCollect       PaymentMethodTypes = "upi_collect"
	PaymentMethodTypesUpiIntent        PaymentMethodTypes = "upi_intent"
	PaymentMethodTypesApplePay         PaymentMethodTypes = "apple_pay"
	PaymentMethodTypesCashapp          PaymentMethodTypes = "cashapp"
	PaymentMethodTypesGooglePay        PaymentMethodTypes = "google_pay"
	PaymentMethodTypesMultibanco       PaymentMethodTypes = "multibanco"
	PaymentMethodTypesBancontactCard   PaymentMethodTypes = "bancontact_card"
	PaymentMethodTypesEps              PaymentMethodTypes = "eps"
	PaymentMethodTypesIdeal            PaymentMethodTypes = "ideal"
	PaymentMethodTypesPrzelewy24       PaymentMethodTypes = "przelewy24"
	PaymentMethodTypesPaypal           PaymentMethodTypes = "paypal"
	PaymentMethodTypesAffirm           PaymentMethodTypes = "affirm"
	PaymentMethodTypesKlarna           PaymentMethodTypes = "klarna"
	PaymentMethodTypesSepa             PaymentMethodTypes = "sepa"
	PaymentMethodTypesACH              PaymentMethodTypes = "ach"
	PaymentMethodTypesAmazonPay        PaymentMethodTypes = "amazon_pay"
	PaymentMethodTypesAfterpayClearpay PaymentMethodTypes = "afterpay_clearpay"
)

type PaymentNewResponse struct {
	// Client secret used to load Dodo checkout SDK NOTE : Dodo checkout SDK will be
	// coming soon
	ClientSecret string `json:"client_secret,required"`
	// Limited details about the customer making the payment
	Customer CustomerLimitedDetails `json:"customer,required"`
	// Additional metadata associated with the payment
	Metadata map[string]string `json:"metadata,required"`
	// Unique identifier for the payment
	PaymentID string `json:"payment_id,required"`
	// Total amount of the payment in smallest currency unit (e.g. cents)
	TotalAmount int64 `json:"total_amount,required"`
	// The discount id if discount is applied
	DiscountID string `json:"discount_id,nullable"`
	// Expiry timestamp of the payment link
	ExpiresOn time.Time `json:"expires_on,nullable" format:"date-time"`
	// Optional URL to a hosted payment page
	PaymentLink string `json:"payment_link,nullable"`
	// Optional list of products included in the payment
	ProductCart []OneTimeProductCartItem `json:"product_cart,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientSecret respjson.Field
		Customer     respjson.Field
		Metadata     respjson.Field
		PaymentID    respjson.Field
		TotalAmount  respjson.Field
		DiscountID   respjson.Field
		ExpiresOn    respjson.Field
		PaymentLink  respjson.Field
		ProductCart  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentListResponse struct {
	BrandID   string    `json:"brand_id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
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
	Currency                 Currency               `json:"currency,required"`
	Customer                 CustomerLimitedDetails `json:"customer,required"`
	DigitalProductsDelivered bool                   `json:"digital_products_delivered,required"`
	Metadata                 map[string]string      `json:"metadata,required"`
	PaymentID                string                 `json:"payment_id,required"`
	TotalAmount              int64                  `json:"total_amount,required"`
	PaymentMethod            string                 `json:"payment_method,nullable"`
	PaymentMethodType        string                 `json:"payment_method_type,nullable"`
	// Any of "succeeded", "failed", "cancelled", "processing",
	// "requires_customer_action", "requires_merchant_action",
	// "requires_payment_method", "requires_confirmation", "requires_capture",
	// "partially_captured", "partially_captured_and_capturable".
	Status         IntentStatus `json:"status,nullable"`
	SubscriptionID string       `json:"subscription_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID                  respjson.Field
		CreatedAt                respjson.Field
		Currency                 respjson.Field
		Customer                 respjson.Field
		DigitalProductsDelivered respjson.Field
		Metadata                 respjson.Field
		PaymentID                respjson.Field
		TotalAmount              respjson.Field
		PaymentMethod            respjson.Field
		PaymentMethodType        respjson.Field
		Status                   respjson.Field
		SubscriptionID           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentGetLineItemsResponse struct {
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
	Currency Currency                          `json:"currency,required"`
	Items    []PaymentGetLineItemsResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGetLineItemsResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentGetLineItemsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentGetLineItemsResponseItem struct {
	Amount           int64  `json:"amount,required"`
	ItemsID          string `json:"items_id,required"`
	RefundableAmount int64  `json:"refundable_amount,required"`
	Tax              int64  `json:"tax,required"`
	Description      string `json:"description,nullable"`
	Name             string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount           respjson.Field
		ItemsID          respjson.Field
		RefundableAmount respjson.Field
		Tax              respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGetLineItemsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *PaymentGetLineItemsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewParams struct {
	// Billing address details for the payment
	Billing BillingAddressParam `json:"billing,omitzero,required"`
	// Customer information for the payment
	Customer CustomerRequestUnionParam `json:"customer,omitzero,required"`
	// List of products in the cart. Must contain at least 1 and at most 100 items.
	ProductCart []OneTimeProductCartItemParam `json:"product_cart,omitzero,required"`
	// Discount Code to apply to the transaction
	DiscountCode param.Opt[string] `json:"discount_code,omitzero"`
	// Override merchant default 3DS behaviour for this payment
	Force3DS param.Opt[bool] `json:"force_3ds,omitzero"`
	// Whether to generate a payment link. Defaults to false if not specified.
	PaymentLink param.Opt[bool] `json:"payment_link,omitzero"`
	// Optional URL to redirect the customer after payment. Must be a valid URL if
	// provided.
	ReturnURL param.Opt[string] `json:"return_url,omitzero"`
	// If true, returns a shortened payment link. Defaults to false if not specified.
	ShortLink param.Opt[bool] `json:"short_link,omitzero"`
	// Tax ID in case the payment is B2B. If tax id validation fails the payment
	// creation will fail
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// If true, redirects the customer immediately after payment completion False by
	// default
	RedirectImmediately param.Opt[bool] `json:"redirect_immediately,omitzero"`
	// Display saved payment methods of a returning customer False by default
	ShowSavedPaymentMethods param.Opt[bool] `json:"show_saved_payment_methods,omitzero"`
	// List of payment methods allowed during checkout.
	//
	// Customers will **never** see payment methods that are **not** in this list.
	// However, adding a method here **does not guarantee** customers will see it.
	// Availability still depends on other factors (e.g., customer location, merchant
	// settings).
	AllowedPaymentMethodTypes []PaymentMethodTypes `json:"allowed_payment_method_types,omitzero"`
	// Fix the currency in which the end customer is billed. If Dodo Payments cannot
	// support that currency for this transaction, it will not proceed
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
	// Additional metadata associated with the payment. Defaults to empty if not
	// provided.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentListParams struct {
	// filter by Brand id
	BrandID param.Opt[string] `query:"brand_id,omitzero" json:"-"`
	// Get events after this created time
	CreatedAtGte param.Opt[time.Time] `query:"created_at_gte,omitzero" format:"date-time" json:"-"`
	// Get events created before this time
	CreatedAtLte param.Opt[time.Time] `query:"created_at_lte,omitzero" format:"date-time" json:"-"`
	// Filter by customer id
	CustomerID param.Opt[string] `query:"customer_id,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter by subscription id
	SubscriptionID param.Opt[string] `query:"subscription_id,omitzero" json:"-"`
	// Filter by status
	//
	// Any of "succeeded", "failed", "cancelled", "processing",
	// "requires_customer_action", "requires_merchant_action",
	// "requires_payment_method", "requires_confirmation", "requires_capture",
	// "partially_captured", "partially_captured_and_capturable".
	Status PaymentListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values, err error) {
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
