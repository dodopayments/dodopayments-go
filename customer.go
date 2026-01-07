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

// CustomerService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options        []option.RequestOption
	CustomerPortal *CustomerCustomerPortalService
	Wallets        *CustomerWalletService
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
	r = &CustomerService{}
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
	JSON        customerJSON      `json:"-"`
}

// customerJSON contains the JSON metadata for the struct [Customer]
type customerJSON struct {
	BusinessID  apijson.Field
	CreatedAt   apijson.Field
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Metadata    apijson.Field
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Customer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerJSON) RawJSON() string {
	return r.raw
}

type CustomerPortalSession struct {
	Link string                    `json:"link,required"`
	JSON customerPortalSessionJSON `json:"-"`
}

// customerPortalSessionJSON contains the JSON metadata for the struct
// [CustomerPortalSession]
type customerPortalSessionJSON struct {
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerPortalSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerPortalSessionJSON) RawJSON() string {
	return r.raw
}

type CustomerGetPaymentMethodsResponse struct {
	Items []CustomerGetPaymentMethodsResponseItem `json:"items,required"`
	JSON  customerGetPaymentMethodsResponseJSON   `json:"-"`
}

// customerGetPaymentMethodsResponseJSON contains the JSON metadata for the struct
// [CustomerGetPaymentMethodsResponse]
type customerGetPaymentMethodsResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerGetPaymentMethodsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerGetPaymentMethodsResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerGetPaymentMethodsResponseItem struct {
	PaymentMethod     CustomerGetPaymentMethodsResponseItemsPaymentMethod `json:"payment_method,required"`
	PaymentMethodID   string                                              `json:"payment_method_id,required"`
	Card              CustomerGetPaymentMethodsResponseItemsCard          `json:"card,nullable"`
	LastUsedAt        time.Time                                           `json:"last_used_at,nullable" format:"date-time"`
	PaymentMethodType PaymentMethodTypes                                  `json:"payment_method_type,nullable"`
	RecurringEnabled  bool                                                `json:"recurring_enabled,nullable"`
	JSON              customerGetPaymentMethodsResponseItemJSON           `json:"-"`
}

// customerGetPaymentMethodsResponseItemJSON contains the JSON metadata for the
// struct [CustomerGetPaymentMethodsResponseItem]
type customerGetPaymentMethodsResponseItemJSON struct {
	PaymentMethod     apijson.Field
	PaymentMethodID   apijson.Field
	Card              apijson.Field
	LastUsedAt        apijson.Field
	PaymentMethodType apijson.Field
	RecurringEnabled  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CustomerGetPaymentMethodsResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerGetPaymentMethodsResponseItemJSON) RawJSON() string {
	return r.raw
}

type CustomerGetPaymentMethodsResponseItemsPaymentMethod string

const (
	CustomerGetPaymentMethodsResponseItemsPaymentMethodCard            CustomerGetPaymentMethodsResponseItemsPaymentMethod = "card"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodCardRedirect    CustomerGetPaymentMethodsResponseItemsPaymentMethod = "card_redirect"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodPayLater        CustomerGetPaymentMethodsResponseItemsPaymentMethod = "pay_later"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodWallet          CustomerGetPaymentMethodsResponseItemsPaymentMethod = "wallet"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodBankRedirect    CustomerGetPaymentMethodsResponseItemsPaymentMethod = "bank_redirect"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodBankTransfer    CustomerGetPaymentMethodsResponseItemsPaymentMethod = "bank_transfer"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodCrypto          CustomerGetPaymentMethodsResponseItemsPaymentMethod = "crypto"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodBankDebit       CustomerGetPaymentMethodsResponseItemsPaymentMethod = "bank_debit"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodReward          CustomerGetPaymentMethodsResponseItemsPaymentMethod = "reward"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodRealTimePayment CustomerGetPaymentMethodsResponseItemsPaymentMethod = "real_time_payment"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodUpi             CustomerGetPaymentMethodsResponseItemsPaymentMethod = "upi"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodVoucher         CustomerGetPaymentMethodsResponseItemsPaymentMethod = "voucher"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodGiftCard        CustomerGetPaymentMethodsResponseItemsPaymentMethod = "gift_card"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodOpenBanking     CustomerGetPaymentMethodsResponseItemsPaymentMethod = "open_banking"
	CustomerGetPaymentMethodsResponseItemsPaymentMethodMobilePayment   CustomerGetPaymentMethodsResponseItemsPaymentMethod = "mobile_payment"
)

func (r CustomerGetPaymentMethodsResponseItemsPaymentMethod) IsKnown() bool {
	switch r {
	case CustomerGetPaymentMethodsResponseItemsPaymentMethodCard, CustomerGetPaymentMethodsResponseItemsPaymentMethodCardRedirect, CustomerGetPaymentMethodsResponseItemsPaymentMethodPayLater, CustomerGetPaymentMethodsResponseItemsPaymentMethodWallet, CustomerGetPaymentMethodsResponseItemsPaymentMethodBankRedirect, CustomerGetPaymentMethodsResponseItemsPaymentMethodBankTransfer, CustomerGetPaymentMethodsResponseItemsPaymentMethodCrypto, CustomerGetPaymentMethodsResponseItemsPaymentMethodBankDebit, CustomerGetPaymentMethodsResponseItemsPaymentMethodReward, CustomerGetPaymentMethodsResponseItemsPaymentMethodRealTimePayment, CustomerGetPaymentMethodsResponseItemsPaymentMethodUpi, CustomerGetPaymentMethodsResponseItemsPaymentMethodVoucher, CustomerGetPaymentMethodsResponseItemsPaymentMethodGiftCard, CustomerGetPaymentMethodsResponseItemsPaymentMethodOpenBanking, CustomerGetPaymentMethodsResponseItemsPaymentMethodMobilePayment:
		return true
	}
	return false
}

type CustomerGetPaymentMethodsResponseItemsCard struct {
	CardHolderName string `json:"card_holder_name,nullable"`
	// ISO country code alpha2 variant
	CardIssuingCountry CountryCode                                    `json:"card_issuing_country,nullable"`
	CardNetwork        string                                         `json:"card_network,nullable"`
	CardType           string                                         `json:"card_type,nullable"`
	ExpiryMonth        string                                         `json:"expiry_month,nullable"`
	ExpiryYear         string                                         `json:"expiry_year,nullable"`
	Last4Digits        string                                         `json:"last4_digits,nullable"`
	JSON               customerGetPaymentMethodsResponseItemsCardJSON `json:"-"`
}

// customerGetPaymentMethodsResponseItemsCardJSON contains the JSON metadata for
// the struct [CustomerGetPaymentMethodsResponseItemsCard]
type customerGetPaymentMethodsResponseItemsCardJSON struct {
	CardHolderName     apijson.Field
	CardIssuingCountry apijson.Field
	CardNetwork        apijson.Field
	CardType           apijson.Field
	ExpiryMonth        apijson.Field
	ExpiryYear         apijson.Field
	Last4Digits        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerGetPaymentMethodsResponseItemsCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerGetPaymentMethodsResponseItemsCardJSON) RawJSON() string {
	return r.raw
}

type CustomerNewParams struct {
	Email param.Field[string] `json:"email,required"`
	Name  param.Field[string] `json:"name,required"`
	// Additional metadata for the customer
	Metadata    param.Field[map[string]string] `json:"metadata"`
	PhoneNumber param.Field[string]            `json:"phone_number"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParams struct {
	// Additional metadata for the customer
	Metadata    param.Field[map[string]string] `json:"metadata"`
	Name        param.Field[string]            `json:"name"`
	PhoneNumber param.Field[string]            `json:"phone_number"`
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerListParams struct {
	// Filter by customer email
	Email param.Field[string] `query:"email"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
