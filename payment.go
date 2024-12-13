// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/dodo-payments-go/internal/apijson"
	"github.com/stainless-sdks/dodo-payments-go/internal/apiquery"
	"github.com/stainless-sdks/dodo-payments-go/internal/param"
	"github.com/stainless-sdks/dodo-payments-go/internal/requestconfig"
	"github.com/stainless-sdks/dodo-payments-go/option"
	"github.com/stainless-sdks/dodo-payments-go/packages/pagination"
)

// PaymentService contains methods and other services that help with interacting
// with the dodopayments API.
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

type Payment struct {
	BusinessID string          `json:"business_id,required"`
	CreatedAt  time.Time       `json:"created_at,required" format:"date-time"`
	Currency   PaymentCurrency `json:"currency,required"`
	Customer   PaymentCustomer `json:"customer,required"`
	Disputes   []Dispute       `json:"disputes,required"`
	PaymentID  string          `json:"payment_id,required"`
	Refunds    []Refund        `json:"refunds,required"`
	// Total amount taken from the customer including tax
	TotalAmount       int64  `json:"total_amount,required"`
	PaymentLink       string `json:"payment_link,nullable"`
	PaymentMethod     string `json:"payment_method,nullable"`
	PaymentMethodType string `json:"payment_method_type,nullable"`
	// Product Cart of One time payment. In case of subscription/recurring payment
	// product id and quantity are available in Get Subscription Api
	ProductCart    []PaymentProductCart `json:"product_cart,nullable"`
	Status         PaymentStatus        `json:"status,nullable"`
	SubscriptionID string               `json:"subscription_id,nullable"`
	// Tax collected in this transaction
	Tax       int64       `json:"tax,nullable"`
	UpdatedAt time.Time   `json:"updated_at,nullable" format:"date-time"`
	JSON      paymentJSON `json:"-"`
}

// paymentJSON contains the JSON metadata for the struct [Payment]
type paymentJSON struct {
	BusinessID        apijson.Field
	CreatedAt         apijson.Field
	Currency          apijson.Field
	Customer          apijson.Field
	Disputes          apijson.Field
	PaymentID         apijson.Field
	Refunds           apijson.Field
	TotalAmount       apijson.Field
	PaymentLink       apijson.Field
	PaymentMethod     apijson.Field
	PaymentMethodType apijson.Field
	ProductCart       apijson.Field
	Status            apijson.Field
	SubscriptionID    apijson.Field
	Tax               apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
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

type PaymentCustomer struct {
	CustomerID string              `json:"customer_id,required"`
	Email      string              `json:"email,required"`
	Name       string              `json:"name,required"`
	JSON       paymentCustomerJSON `json:"-"`
}

// paymentCustomerJSON contains the JSON metadata for the struct [PaymentCustomer]
type paymentCustomerJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentCustomerJSON) RawJSON() string {
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

type PaymentStatus string

const (
	PaymentStatusSucceeded                      PaymentStatus = "succeeded"
	PaymentStatusFailed                         PaymentStatus = "failed"
	PaymentStatusCancelled                      PaymentStatus = "cancelled"
	PaymentStatusProcessing                     PaymentStatus = "processing"
	PaymentStatusRequiresCustomerAction         PaymentStatus = "requires_customer_action"
	PaymentStatusRequiresMerchantAction         PaymentStatus = "requires_merchant_action"
	PaymentStatusRequiresPaymentMethod          PaymentStatus = "requires_payment_method"
	PaymentStatusRequiresConfirmation           PaymentStatus = "requires_confirmation"
	PaymentStatusRequiresCapture                PaymentStatus = "requires_capture"
	PaymentStatusPartiallyCaptured              PaymentStatus = "partially_captured"
	PaymentStatusPartiallyCapturedAndCapturable PaymentStatus = "partially_captured_and_capturable"
)

func (r PaymentStatus) IsKnown() bool {
	switch r {
	case PaymentStatusSucceeded, PaymentStatusFailed, PaymentStatusCancelled, PaymentStatusProcessing, PaymentStatusRequiresCustomerAction, PaymentStatusRequiresMerchantAction, PaymentStatusRequiresPaymentMethod, PaymentStatusRequiresConfirmation, PaymentStatusRequiresCapture, PaymentStatusPartiallyCaptured, PaymentStatusPartiallyCapturedAndCapturable:
		return true
	}
	return false
}

type PaymentNewResponse struct {
	ClientSecret string                          `json:"client_secret,required"`
	Customer     PaymentNewResponseCustomer      `json:"customer,required"`
	PaymentID    string                          `json:"payment_id,required"`
	TotalAmount  int64                           `json:"total_amount,required"`
	PaymentLink  string                          `json:"payment_link,nullable"`
	ProductCart  []PaymentNewResponseProductCart `json:"product_cart,nullable"`
	JSON         paymentNewResponseJSON          `json:"-"`
}

// paymentNewResponseJSON contains the JSON metadata for the struct
// [PaymentNewResponse]
type paymentNewResponseJSON struct {
	ClientSecret apijson.Field
	Customer     apijson.Field
	PaymentID    apijson.Field
	TotalAmount  apijson.Field
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

type PaymentNewResponseCustomer struct {
	CustomerID string                         `json:"customer_id,required"`
	Email      string                         `json:"email,required"`
	Name       string                         `json:"name,required"`
	JSON       paymentNewResponseCustomerJSON `json:"-"`
}

// paymentNewResponseCustomerJSON contains the JSON metadata for the struct
// [PaymentNewResponseCustomer]
type paymentNewResponseCustomerJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponseCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentNewResponseCustomerJSON) RawJSON() string {
	return r.raw
}

type PaymentNewResponseProductCart struct {
	ProductID string                            `json:"product_id,required"`
	Quantity  int64                             `json:"quantity,required"`
	JSON      paymentNewResponseProductCartJSON `json:"-"`
}

// paymentNewResponseProductCartJSON contains the JSON metadata for the struct
// [PaymentNewResponseProductCart]
type paymentNewResponseProductCartJSON struct {
	ProductID   apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentNewResponseProductCart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentNewResponseProductCartJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponse struct {
	CreatedAt         time.Time                   `json:"created_at,required" format:"date-time"`
	Currency          PaymentListResponseCurrency `json:"currency,required"`
	Customer          PaymentListResponseCustomer `json:"customer,required"`
	PaymentID         string                      `json:"payment_id,required"`
	TotalAmount       int64                       `json:"total_amount,required"`
	PaymentMethod     string                      `json:"payment_method,nullable"`
	PaymentMethodType string                      `json:"payment_method_type,nullable"`
	Status            PaymentListResponseStatus   `json:"status,nullable"`
	SubscriptionID    string                      `json:"subscription_id,nullable"`
	JSON              paymentListResponseJSON     `json:"-"`
}

// paymentListResponseJSON contains the JSON metadata for the struct
// [PaymentListResponse]
type paymentListResponseJSON struct {
	CreatedAt         apijson.Field
	Currency          apijson.Field
	Customer          apijson.Field
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

type PaymentListResponseCustomer struct {
	CustomerID string                          `json:"customer_id,required"`
	Email      string                          `json:"email,required"`
	Name       string                          `json:"name,required"`
	JSON       paymentListResponseCustomerJSON `json:"-"`
}

// paymentListResponseCustomerJSON contains the JSON metadata for the struct
// [PaymentListResponseCustomer]
type paymentListResponseCustomerJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentListResponseCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseCustomerJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponseStatus string

const (
	PaymentListResponseStatusSucceeded                      PaymentListResponseStatus = "succeeded"
	PaymentListResponseStatusFailed                         PaymentListResponseStatus = "failed"
	PaymentListResponseStatusCancelled                      PaymentListResponseStatus = "cancelled"
	PaymentListResponseStatusProcessing                     PaymentListResponseStatus = "processing"
	PaymentListResponseStatusRequiresCustomerAction         PaymentListResponseStatus = "requires_customer_action"
	PaymentListResponseStatusRequiresMerchantAction         PaymentListResponseStatus = "requires_merchant_action"
	PaymentListResponseStatusRequiresPaymentMethod          PaymentListResponseStatus = "requires_payment_method"
	PaymentListResponseStatusRequiresConfirmation           PaymentListResponseStatus = "requires_confirmation"
	PaymentListResponseStatusRequiresCapture                PaymentListResponseStatus = "requires_capture"
	PaymentListResponseStatusPartiallyCaptured              PaymentListResponseStatus = "partially_captured"
	PaymentListResponseStatusPartiallyCapturedAndCapturable PaymentListResponseStatus = "partially_captured_and_capturable"
)

func (r PaymentListResponseStatus) IsKnown() bool {
	switch r {
	case PaymentListResponseStatusSucceeded, PaymentListResponseStatusFailed, PaymentListResponseStatusCancelled, PaymentListResponseStatusProcessing, PaymentListResponseStatusRequiresCustomerAction, PaymentListResponseStatusRequiresMerchantAction, PaymentListResponseStatusRequiresPaymentMethod, PaymentListResponseStatusRequiresConfirmation, PaymentListResponseStatusRequiresCapture, PaymentListResponseStatusPartiallyCaptured, PaymentListResponseStatusPartiallyCapturedAndCapturable:
		return true
	}
	return false
}

type PaymentNewParams struct {
	Billing     param.Field[PaymentNewParamsBilling]       `json:"billing,required"`
	Customer    param.Field[PaymentNewParamsCustomer]      `json:"customer,required"`
	ProductCart param.Field[[]PaymentNewParamsProductCart] `json:"product_cart,required"`
	PaymentLink param.Field[bool]                          `json:"payment_link"`
	ReturnURL   param.Field[string]                        `json:"return_url"`
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsBilling struct {
	City param.Field[string] `json:"city,required"`
	// ISO country code alpha2 variant
	Country param.Field[CountryCodeAlpha2] `json:"country,required"`
	State   param.Field[string]            `json:"state,required"`
	Street  param.Field[string]            `json:"street,required"`
	Zipcode param.Field[int64]             `json:"zipcode,required"`
}

func (r PaymentNewParamsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsCustomer struct {
	Email       param.Field[string] `json:"email,required"`
	Name        param.Field[string] `json:"name,required"`
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r PaymentNewParamsCustomer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentNewParamsProductCart struct {
	ProductID param.Field[string] `json:"product_id,required"`
	Quantity  param.Field[int64]  `json:"quantity,required"`
}

func (r PaymentNewParamsProductCart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentListParams struct {
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
