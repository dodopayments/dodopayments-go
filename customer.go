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
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// CustomerService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options        []option.RequestOption
	CustomerPortal CustomerCustomerPortalService
	Wallets        CustomerWalletService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r CustomerService) {
	r = CustomerService{}
	r.Options = opts
	r.CustomerPortal = NewCustomerCustomerPortalService(opts...)
	r.Wallets = NewCustomerWalletService(opts...)
	return
}

func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *CustomerService) Get(ctx context.Context, customerID string, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *CustomerService) Update(ctx context.Context, customerID string, body CustomerUpdateParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Customer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "customers"
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

func (r *CustomerService) ListAutoPaging(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Customer] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *CustomerService) GetPaymentMethods(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerGetPaymentMethodsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/payment-methods", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Customer struct {
	BusinessID string    `json:"business_id,required"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
	CustomerID string    `json:"customer_id,required"`
	Email      string    `json:"email,required"`
	Name       string    `json:"name,required"`
	// Additional metadata for the customer
	Metadata    map[string]string `json:"metadata"`
	PhoneNumber string            `json:"phone_number,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		CreatedAt   respjson.Field
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
func (r Customer) RawJSON() string { return r.JSON.raw }
func (r *Customer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerPortalSession struct {
	Link string `json:"link,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerPortalSession) RawJSON() string { return r.JSON.raw }
func (r *CustomerPortalSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetPaymentMethodsResponse struct {
	Items []CustomerGetPaymentMethodsResponseItem `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerGetPaymentMethodsResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerGetPaymentMethodsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetPaymentMethodsResponseItem struct {
	// PaymentMethod enum from hyperswitch
	//
	// https://github.com/juspay/hyperswitch/blob/ecd05d53c99ae701ac94893ec632a3988afe3238/crates/common_enums/src/enums.rs#L2097
	//
	// Any of "card", "card_redirect", "pay_later", "wallet", "bank_redirect",
	// "bank_transfer", "crypto", "bank_debit", "reward", "real_time_payment", "upi",
	// "voucher", "gift_card", "open_banking", "mobile_payment".
	PaymentMethod   string                                    `json:"payment_method,required"`
	PaymentMethodID string                                    `json:"payment_method_id,required"`
	Card            CustomerGetPaymentMethodsResponseItemCard `json:"card,nullable"`
	LastUsedAt      time.Time                                 `json:"last_used_at,nullable" format:"date-time"`
	// Any of "credit", "debit", "upi_collect", "upi_intent", "apple_pay", "cashapp",
	// "google_pay", "multibanco", "bancontact_card", "eps", "ideal", "przelewy24",
	// "paypal", "affirm", "klarna", "sepa", "ach", "amazon_pay", "afterpay_clearpay".
	PaymentMethodType PaymentMethodTypes `json:"payment_method_type,nullable"`
	RecurringEnabled  bool               `json:"recurring_enabled,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentMethod     respjson.Field
		PaymentMethodID   respjson.Field
		Card              respjson.Field
		LastUsedAt        respjson.Field
		PaymentMethodType respjson.Field
		RecurringEnabled  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerGetPaymentMethodsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *CustomerGetPaymentMethodsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerGetPaymentMethodsResponseItemCard struct {
	// ISO country code alpha2 variant
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
	CardNetwork        string      `json:"card_network,nullable"`
	CardType           string      `json:"card_type,nullable"`
	ExpiryMonth        string      `json:"expiry_month,nullable"`
	ExpiryYear         string      `json:"expiry_year,nullable"`
	Last4Digits        string      `json:"last4_digits,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardIssuingCountry respjson.Field
		CardNetwork        respjson.Field
		CardType           respjson.Field
		ExpiryMonth        respjson.Field
		ExpiryYear         respjson.Field
		Last4Digits        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerGetPaymentMethodsResponseItemCard) RawJSON() string { return r.JSON.raw }
func (r *CustomerGetPaymentMethodsResponseItemCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParams struct {
	Email       string            `json:"email,required"`
	Name        string            `json:"name,required"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Additional metadata for the customer
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUpdateParams struct {
	Name        param.Opt[string] `json:"name,omitzero"`
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Additional metadata for the customer
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListParams struct {
	// Filter by customer email
	Email param.Opt[string] `query:"email,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
