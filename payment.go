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

func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *PaymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	Items []PaymentListResponseItem `json:"items,required"`
	JSON  paymentListResponseJSON   `json:"-"`
}

// paymentListResponseJSON contains the JSON metadata for the struct
// [PaymentListResponse]
type paymentListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponseItem struct {
	CreatedAt         time.Time                        `json:"created_at,required" format:"date-time"`
	Currency          PaymentListResponseItemsCurrency `json:"currency,required"`
	Customer          PaymentListResponseItemsCustomer `json:"customer,required"`
	PaymentID         string                           `json:"payment_id,required"`
	TotalAmount       int64                            `json:"total_amount,required"`
	PaymentMethod     string                           `json:"payment_method,nullable"`
	PaymentMethodType string                           `json:"payment_method_type,nullable"`
	Status            PaymentListResponseItemsStatus   `json:"status,nullable"`
	SubscriptionID    string                           `json:"subscription_id,nullable"`
	JSON              paymentListResponseItemJSON      `json:"-"`
}

// paymentListResponseItemJSON contains the JSON metadata for the struct
// [PaymentListResponseItem]
type paymentListResponseItemJSON struct {
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

func (r *PaymentListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseItemJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponseItemsCurrency string

const (
	PaymentListResponseItemsCurrencyAed PaymentListResponseItemsCurrency = "AED"
	PaymentListResponseItemsCurrencyAll PaymentListResponseItemsCurrency = "ALL"
	PaymentListResponseItemsCurrencyAmd PaymentListResponseItemsCurrency = "AMD"
	PaymentListResponseItemsCurrencyAng PaymentListResponseItemsCurrency = "ANG"
	PaymentListResponseItemsCurrencyAoa PaymentListResponseItemsCurrency = "AOA"
	PaymentListResponseItemsCurrencyArs PaymentListResponseItemsCurrency = "ARS"
	PaymentListResponseItemsCurrencyAud PaymentListResponseItemsCurrency = "AUD"
	PaymentListResponseItemsCurrencyAwg PaymentListResponseItemsCurrency = "AWG"
	PaymentListResponseItemsCurrencyAzn PaymentListResponseItemsCurrency = "AZN"
	PaymentListResponseItemsCurrencyBam PaymentListResponseItemsCurrency = "BAM"
	PaymentListResponseItemsCurrencyBbd PaymentListResponseItemsCurrency = "BBD"
	PaymentListResponseItemsCurrencyBdt PaymentListResponseItemsCurrency = "BDT"
	PaymentListResponseItemsCurrencyBgn PaymentListResponseItemsCurrency = "BGN"
	PaymentListResponseItemsCurrencyBhd PaymentListResponseItemsCurrency = "BHD"
	PaymentListResponseItemsCurrencyBif PaymentListResponseItemsCurrency = "BIF"
	PaymentListResponseItemsCurrencyBmd PaymentListResponseItemsCurrency = "BMD"
	PaymentListResponseItemsCurrencyBnd PaymentListResponseItemsCurrency = "BND"
	PaymentListResponseItemsCurrencyBob PaymentListResponseItemsCurrency = "BOB"
	PaymentListResponseItemsCurrencyBrl PaymentListResponseItemsCurrency = "BRL"
	PaymentListResponseItemsCurrencyBsd PaymentListResponseItemsCurrency = "BSD"
	PaymentListResponseItemsCurrencyBwp PaymentListResponseItemsCurrency = "BWP"
	PaymentListResponseItemsCurrencyByn PaymentListResponseItemsCurrency = "BYN"
	PaymentListResponseItemsCurrencyBzd PaymentListResponseItemsCurrency = "BZD"
	PaymentListResponseItemsCurrencyCad PaymentListResponseItemsCurrency = "CAD"
	PaymentListResponseItemsCurrencyChf PaymentListResponseItemsCurrency = "CHF"
	PaymentListResponseItemsCurrencyClp PaymentListResponseItemsCurrency = "CLP"
	PaymentListResponseItemsCurrencyCny PaymentListResponseItemsCurrency = "CNY"
	PaymentListResponseItemsCurrencyCop PaymentListResponseItemsCurrency = "COP"
	PaymentListResponseItemsCurrencyCrc PaymentListResponseItemsCurrency = "CRC"
	PaymentListResponseItemsCurrencyCup PaymentListResponseItemsCurrency = "CUP"
	PaymentListResponseItemsCurrencyCve PaymentListResponseItemsCurrency = "CVE"
	PaymentListResponseItemsCurrencyCzk PaymentListResponseItemsCurrency = "CZK"
	PaymentListResponseItemsCurrencyDjf PaymentListResponseItemsCurrency = "DJF"
	PaymentListResponseItemsCurrencyDkk PaymentListResponseItemsCurrency = "DKK"
	PaymentListResponseItemsCurrencyDop PaymentListResponseItemsCurrency = "DOP"
	PaymentListResponseItemsCurrencyDzd PaymentListResponseItemsCurrency = "DZD"
	PaymentListResponseItemsCurrencyEgp PaymentListResponseItemsCurrency = "EGP"
	PaymentListResponseItemsCurrencyEtb PaymentListResponseItemsCurrency = "ETB"
	PaymentListResponseItemsCurrencyEur PaymentListResponseItemsCurrency = "EUR"
	PaymentListResponseItemsCurrencyFjd PaymentListResponseItemsCurrency = "FJD"
	PaymentListResponseItemsCurrencyFkp PaymentListResponseItemsCurrency = "FKP"
	PaymentListResponseItemsCurrencyGbp PaymentListResponseItemsCurrency = "GBP"
	PaymentListResponseItemsCurrencyGel PaymentListResponseItemsCurrency = "GEL"
	PaymentListResponseItemsCurrencyGhs PaymentListResponseItemsCurrency = "GHS"
	PaymentListResponseItemsCurrencyGip PaymentListResponseItemsCurrency = "GIP"
	PaymentListResponseItemsCurrencyGmd PaymentListResponseItemsCurrency = "GMD"
	PaymentListResponseItemsCurrencyGnf PaymentListResponseItemsCurrency = "GNF"
	PaymentListResponseItemsCurrencyGtq PaymentListResponseItemsCurrency = "GTQ"
	PaymentListResponseItemsCurrencyGyd PaymentListResponseItemsCurrency = "GYD"
	PaymentListResponseItemsCurrencyHkd PaymentListResponseItemsCurrency = "HKD"
	PaymentListResponseItemsCurrencyHnl PaymentListResponseItemsCurrency = "HNL"
	PaymentListResponseItemsCurrencyHrk PaymentListResponseItemsCurrency = "HRK"
	PaymentListResponseItemsCurrencyHtg PaymentListResponseItemsCurrency = "HTG"
	PaymentListResponseItemsCurrencyHuf PaymentListResponseItemsCurrency = "HUF"
	PaymentListResponseItemsCurrencyIdr PaymentListResponseItemsCurrency = "IDR"
	PaymentListResponseItemsCurrencyIls PaymentListResponseItemsCurrency = "ILS"
	PaymentListResponseItemsCurrencyInr PaymentListResponseItemsCurrency = "INR"
	PaymentListResponseItemsCurrencyIqd PaymentListResponseItemsCurrency = "IQD"
	PaymentListResponseItemsCurrencyJmd PaymentListResponseItemsCurrency = "JMD"
	PaymentListResponseItemsCurrencyJod PaymentListResponseItemsCurrency = "JOD"
	PaymentListResponseItemsCurrencyJpy PaymentListResponseItemsCurrency = "JPY"
	PaymentListResponseItemsCurrencyKes PaymentListResponseItemsCurrency = "KES"
	PaymentListResponseItemsCurrencyKgs PaymentListResponseItemsCurrency = "KGS"
	PaymentListResponseItemsCurrencyKhr PaymentListResponseItemsCurrency = "KHR"
	PaymentListResponseItemsCurrencyKmf PaymentListResponseItemsCurrency = "KMF"
	PaymentListResponseItemsCurrencyKrw PaymentListResponseItemsCurrency = "KRW"
	PaymentListResponseItemsCurrencyKwd PaymentListResponseItemsCurrency = "KWD"
	PaymentListResponseItemsCurrencyKyd PaymentListResponseItemsCurrency = "KYD"
	PaymentListResponseItemsCurrencyKzt PaymentListResponseItemsCurrency = "KZT"
	PaymentListResponseItemsCurrencyLak PaymentListResponseItemsCurrency = "LAK"
	PaymentListResponseItemsCurrencyLbp PaymentListResponseItemsCurrency = "LBP"
	PaymentListResponseItemsCurrencyLkr PaymentListResponseItemsCurrency = "LKR"
	PaymentListResponseItemsCurrencyLrd PaymentListResponseItemsCurrency = "LRD"
	PaymentListResponseItemsCurrencyLsl PaymentListResponseItemsCurrency = "LSL"
	PaymentListResponseItemsCurrencyLyd PaymentListResponseItemsCurrency = "LYD"
	PaymentListResponseItemsCurrencyMad PaymentListResponseItemsCurrency = "MAD"
	PaymentListResponseItemsCurrencyMdl PaymentListResponseItemsCurrency = "MDL"
	PaymentListResponseItemsCurrencyMga PaymentListResponseItemsCurrency = "MGA"
	PaymentListResponseItemsCurrencyMkd PaymentListResponseItemsCurrency = "MKD"
	PaymentListResponseItemsCurrencyMmk PaymentListResponseItemsCurrency = "MMK"
	PaymentListResponseItemsCurrencyMnt PaymentListResponseItemsCurrency = "MNT"
	PaymentListResponseItemsCurrencyMop PaymentListResponseItemsCurrency = "MOP"
	PaymentListResponseItemsCurrencyMru PaymentListResponseItemsCurrency = "MRU"
	PaymentListResponseItemsCurrencyMur PaymentListResponseItemsCurrency = "MUR"
	PaymentListResponseItemsCurrencyMvr PaymentListResponseItemsCurrency = "MVR"
	PaymentListResponseItemsCurrencyMwk PaymentListResponseItemsCurrency = "MWK"
	PaymentListResponseItemsCurrencyMxn PaymentListResponseItemsCurrency = "MXN"
	PaymentListResponseItemsCurrencyMyr PaymentListResponseItemsCurrency = "MYR"
	PaymentListResponseItemsCurrencyMzn PaymentListResponseItemsCurrency = "MZN"
	PaymentListResponseItemsCurrencyNad PaymentListResponseItemsCurrency = "NAD"
	PaymentListResponseItemsCurrencyNgn PaymentListResponseItemsCurrency = "NGN"
	PaymentListResponseItemsCurrencyNio PaymentListResponseItemsCurrency = "NIO"
	PaymentListResponseItemsCurrencyNok PaymentListResponseItemsCurrency = "NOK"
	PaymentListResponseItemsCurrencyNpr PaymentListResponseItemsCurrency = "NPR"
	PaymentListResponseItemsCurrencyNzd PaymentListResponseItemsCurrency = "NZD"
	PaymentListResponseItemsCurrencyOmr PaymentListResponseItemsCurrency = "OMR"
	PaymentListResponseItemsCurrencyPab PaymentListResponseItemsCurrency = "PAB"
	PaymentListResponseItemsCurrencyPen PaymentListResponseItemsCurrency = "PEN"
	PaymentListResponseItemsCurrencyPgk PaymentListResponseItemsCurrency = "PGK"
	PaymentListResponseItemsCurrencyPhp PaymentListResponseItemsCurrency = "PHP"
	PaymentListResponseItemsCurrencyPkr PaymentListResponseItemsCurrency = "PKR"
	PaymentListResponseItemsCurrencyPln PaymentListResponseItemsCurrency = "PLN"
	PaymentListResponseItemsCurrencyPyg PaymentListResponseItemsCurrency = "PYG"
	PaymentListResponseItemsCurrencyQar PaymentListResponseItemsCurrency = "QAR"
	PaymentListResponseItemsCurrencyRon PaymentListResponseItemsCurrency = "RON"
	PaymentListResponseItemsCurrencyRsd PaymentListResponseItemsCurrency = "RSD"
	PaymentListResponseItemsCurrencyRub PaymentListResponseItemsCurrency = "RUB"
	PaymentListResponseItemsCurrencyRwf PaymentListResponseItemsCurrency = "RWF"
	PaymentListResponseItemsCurrencySar PaymentListResponseItemsCurrency = "SAR"
	PaymentListResponseItemsCurrencySbd PaymentListResponseItemsCurrency = "SBD"
	PaymentListResponseItemsCurrencyScr PaymentListResponseItemsCurrency = "SCR"
	PaymentListResponseItemsCurrencySek PaymentListResponseItemsCurrency = "SEK"
	PaymentListResponseItemsCurrencySgd PaymentListResponseItemsCurrency = "SGD"
	PaymentListResponseItemsCurrencyShp PaymentListResponseItemsCurrency = "SHP"
	PaymentListResponseItemsCurrencySle PaymentListResponseItemsCurrency = "SLE"
	PaymentListResponseItemsCurrencySll PaymentListResponseItemsCurrency = "SLL"
	PaymentListResponseItemsCurrencySos PaymentListResponseItemsCurrency = "SOS"
	PaymentListResponseItemsCurrencySrd PaymentListResponseItemsCurrency = "SRD"
	PaymentListResponseItemsCurrencySsp PaymentListResponseItemsCurrency = "SSP"
	PaymentListResponseItemsCurrencyStn PaymentListResponseItemsCurrency = "STN"
	PaymentListResponseItemsCurrencySvc PaymentListResponseItemsCurrency = "SVC"
	PaymentListResponseItemsCurrencySzl PaymentListResponseItemsCurrency = "SZL"
	PaymentListResponseItemsCurrencyThb PaymentListResponseItemsCurrency = "THB"
	PaymentListResponseItemsCurrencyTnd PaymentListResponseItemsCurrency = "TND"
	PaymentListResponseItemsCurrencyTop PaymentListResponseItemsCurrency = "TOP"
	PaymentListResponseItemsCurrencyTry PaymentListResponseItemsCurrency = "TRY"
	PaymentListResponseItemsCurrencyTtd PaymentListResponseItemsCurrency = "TTD"
	PaymentListResponseItemsCurrencyTwd PaymentListResponseItemsCurrency = "TWD"
	PaymentListResponseItemsCurrencyTzs PaymentListResponseItemsCurrency = "TZS"
	PaymentListResponseItemsCurrencyUah PaymentListResponseItemsCurrency = "UAH"
	PaymentListResponseItemsCurrencyUgx PaymentListResponseItemsCurrency = "UGX"
	PaymentListResponseItemsCurrencyUsd PaymentListResponseItemsCurrency = "USD"
	PaymentListResponseItemsCurrencyUyu PaymentListResponseItemsCurrency = "UYU"
	PaymentListResponseItemsCurrencyUzs PaymentListResponseItemsCurrency = "UZS"
	PaymentListResponseItemsCurrencyVes PaymentListResponseItemsCurrency = "VES"
	PaymentListResponseItemsCurrencyVnd PaymentListResponseItemsCurrency = "VND"
	PaymentListResponseItemsCurrencyVuv PaymentListResponseItemsCurrency = "VUV"
	PaymentListResponseItemsCurrencyWst PaymentListResponseItemsCurrency = "WST"
	PaymentListResponseItemsCurrencyXaf PaymentListResponseItemsCurrency = "XAF"
	PaymentListResponseItemsCurrencyXcd PaymentListResponseItemsCurrency = "XCD"
	PaymentListResponseItemsCurrencyXof PaymentListResponseItemsCurrency = "XOF"
	PaymentListResponseItemsCurrencyXpf PaymentListResponseItemsCurrency = "XPF"
	PaymentListResponseItemsCurrencyYer PaymentListResponseItemsCurrency = "YER"
	PaymentListResponseItemsCurrencyZar PaymentListResponseItemsCurrency = "ZAR"
	PaymentListResponseItemsCurrencyZmw PaymentListResponseItemsCurrency = "ZMW"
)

func (r PaymentListResponseItemsCurrency) IsKnown() bool {
	switch r {
	case PaymentListResponseItemsCurrencyAed, PaymentListResponseItemsCurrencyAll, PaymentListResponseItemsCurrencyAmd, PaymentListResponseItemsCurrencyAng, PaymentListResponseItemsCurrencyAoa, PaymentListResponseItemsCurrencyArs, PaymentListResponseItemsCurrencyAud, PaymentListResponseItemsCurrencyAwg, PaymentListResponseItemsCurrencyAzn, PaymentListResponseItemsCurrencyBam, PaymentListResponseItemsCurrencyBbd, PaymentListResponseItemsCurrencyBdt, PaymentListResponseItemsCurrencyBgn, PaymentListResponseItemsCurrencyBhd, PaymentListResponseItemsCurrencyBif, PaymentListResponseItemsCurrencyBmd, PaymentListResponseItemsCurrencyBnd, PaymentListResponseItemsCurrencyBob, PaymentListResponseItemsCurrencyBrl, PaymentListResponseItemsCurrencyBsd, PaymentListResponseItemsCurrencyBwp, PaymentListResponseItemsCurrencyByn, PaymentListResponseItemsCurrencyBzd, PaymentListResponseItemsCurrencyCad, PaymentListResponseItemsCurrencyChf, PaymentListResponseItemsCurrencyClp, PaymentListResponseItemsCurrencyCny, PaymentListResponseItemsCurrencyCop, PaymentListResponseItemsCurrencyCrc, PaymentListResponseItemsCurrencyCup, PaymentListResponseItemsCurrencyCve, PaymentListResponseItemsCurrencyCzk, PaymentListResponseItemsCurrencyDjf, PaymentListResponseItemsCurrencyDkk, PaymentListResponseItemsCurrencyDop, PaymentListResponseItemsCurrencyDzd, PaymentListResponseItemsCurrencyEgp, PaymentListResponseItemsCurrencyEtb, PaymentListResponseItemsCurrencyEur, PaymentListResponseItemsCurrencyFjd, PaymentListResponseItemsCurrencyFkp, PaymentListResponseItemsCurrencyGbp, PaymentListResponseItemsCurrencyGel, PaymentListResponseItemsCurrencyGhs, PaymentListResponseItemsCurrencyGip, PaymentListResponseItemsCurrencyGmd, PaymentListResponseItemsCurrencyGnf, PaymentListResponseItemsCurrencyGtq, PaymentListResponseItemsCurrencyGyd, PaymentListResponseItemsCurrencyHkd, PaymentListResponseItemsCurrencyHnl, PaymentListResponseItemsCurrencyHrk, PaymentListResponseItemsCurrencyHtg, PaymentListResponseItemsCurrencyHuf, PaymentListResponseItemsCurrencyIdr, PaymentListResponseItemsCurrencyIls, PaymentListResponseItemsCurrencyInr, PaymentListResponseItemsCurrencyIqd, PaymentListResponseItemsCurrencyJmd, PaymentListResponseItemsCurrencyJod, PaymentListResponseItemsCurrencyJpy, PaymentListResponseItemsCurrencyKes, PaymentListResponseItemsCurrencyKgs, PaymentListResponseItemsCurrencyKhr, PaymentListResponseItemsCurrencyKmf, PaymentListResponseItemsCurrencyKrw, PaymentListResponseItemsCurrencyKwd, PaymentListResponseItemsCurrencyKyd, PaymentListResponseItemsCurrencyKzt, PaymentListResponseItemsCurrencyLak, PaymentListResponseItemsCurrencyLbp, PaymentListResponseItemsCurrencyLkr, PaymentListResponseItemsCurrencyLrd, PaymentListResponseItemsCurrencyLsl, PaymentListResponseItemsCurrencyLyd, PaymentListResponseItemsCurrencyMad, PaymentListResponseItemsCurrencyMdl, PaymentListResponseItemsCurrencyMga, PaymentListResponseItemsCurrencyMkd, PaymentListResponseItemsCurrencyMmk, PaymentListResponseItemsCurrencyMnt, PaymentListResponseItemsCurrencyMop, PaymentListResponseItemsCurrencyMru, PaymentListResponseItemsCurrencyMur, PaymentListResponseItemsCurrencyMvr, PaymentListResponseItemsCurrencyMwk, PaymentListResponseItemsCurrencyMxn, PaymentListResponseItemsCurrencyMyr, PaymentListResponseItemsCurrencyMzn, PaymentListResponseItemsCurrencyNad, PaymentListResponseItemsCurrencyNgn, PaymentListResponseItemsCurrencyNio, PaymentListResponseItemsCurrencyNok, PaymentListResponseItemsCurrencyNpr, PaymentListResponseItemsCurrencyNzd, PaymentListResponseItemsCurrencyOmr, PaymentListResponseItemsCurrencyPab, PaymentListResponseItemsCurrencyPen, PaymentListResponseItemsCurrencyPgk, PaymentListResponseItemsCurrencyPhp, PaymentListResponseItemsCurrencyPkr, PaymentListResponseItemsCurrencyPln, PaymentListResponseItemsCurrencyPyg, PaymentListResponseItemsCurrencyQar, PaymentListResponseItemsCurrencyRon, PaymentListResponseItemsCurrencyRsd, PaymentListResponseItemsCurrencyRub, PaymentListResponseItemsCurrencyRwf, PaymentListResponseItemsCurrencySar, PaymentListResponseItemsCurrencySbd, PaymentListResponseItemsCurrencyScr, PaymentListResponseItemsCurrencySek, PaymentListResponseItemsCurrencySgd, PaymentListResponseItemsCurrencyShp, PaymentListResponseItemsCurrencySle, PaymentListResponseItemsCurrencySll, PaymentListResponseItemsCurrencySos, PaymentListResponseItemsCurrencySrd, PaymentListResponseItemsCurrencySsp, PaymentListResponseItemsCurrencyStn, PaymentListResponseItemsCurrencySvc, PaymentListResponseItemsCurrencySzl, PaymentListResponseItemsCurrencyThb, PaymentListResponseItemsCurrencyTnd, PaymentListResponseItemsCurrencyTop, PaymentListResponseItemsCurrencyTry, PaymentListResponseItemsCurrencyTtd, PaymentListResponseItemsCurrencyTwd, PaymentListResponseItemsCurrencyTzs, PaymentListResponseItemsCurrencyUah, PaymentListResponseItemsCurrencyUgx, PaymentListResponseItemsCurrencyUsd, PaymentListResponseItemsCurrencyUyu, PaymentListResponseItemsCurrencyUzs, PaymentListResponseItemsCurrencyVes, PaymentListResponseItemsCurrencyVnd, PaymentListResponseItemsCurrencyVuv, PaymentListResponseItemsCurrencyWst, PaymentListResponseItemsCurrencyXaf, PaymentListResponseItemsCurrencyXcd, PaymentListResponseItemsCurrencyXof, PaymentListResponseItemsCurrencyXpf, PaymentListResponseItemsCurrencyYer, PaymentListResponseItemsCurrencyZar, PaymentListResponseItemsCurrencyZmw:
		return true
	}
	return false
}

type PaymentListResponseItemsCustomer struct {
	CustomerID string                               `json:"customer_id,required"`
	Email      string                               `json:"email,required"`
	Name       string                               `json:"name,required"`
	JSON       paymentListResponseItemsCustomerJSON `json:"-"`
}

// paymentListResponseItemsCustomerJSON contains the JSON metadata for the struct
// [PaymentListResponseItemsCustomer]
type paymentListResponseItemsCustomerJSON struct {
	CustomerID  apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentListResponseItemsCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentListResponseItemsCustomerJSON) RawJSON() string {
	return r.raw
}

type PaymentListResponseItemsStatus string

const (
	PaymentListResponseItemsStatusSucceeded                      PaymentListResponseItemsStatus = "succeeded"
	PaymentListResponseItemsStatusFailed                         PaymentListResponseItemsStatus = "failed"
	PaymentListResponseItemsStatusCancelled                      PaymentListResponseItemsStatus = "cancelled"
	PaymentListResponseItemsStatusProcessing                     PaymentListResponseItemsStatus = "processing"
	PaymentListResponseItemsStatusRequiresCustomerAction         PaymentListResponseItemsStatus = "requires_customer_action"
	PaymentListResponseItemsStatusRequiresMerchantAction         PaymentListResponseItemsStatus = "requires_merchant_action"
	PaymentListResponseItemsStatusRequiresPaymentMethod          PaymentListResponseItemsStatus = "requires_payment_method"
	PaymentListResponseItemsStatusRequiresConfirmation           PaymentListResponseItemsStatus = "requires_confirmation"
	PaymentListResponseItemsStatusRequiresCapture                PaymentListResponseItemsStatus = "requires_capture"
	PaymentListResponseItemsStatusPartiallyCaptured              PaymentListResponseItemsStatus = "partially_captured"
	PaymentListResponseItemsStatusPartiallyCapturedAndCapturable PaymentListResponseItemsStatus = "partially_captured_and_capturable"
)

func (r PaymentListResponseItemsStatus) IsKnown() bool {
	switch r {
	case PaymentListResponseItemsStatusSucceeded, PaymentListResponseItemsStatusFailed, PaymentListResponseItemsStatusCancelled, PaymentListResponseItemsStatusProcessing, PaymentListResponseItemsStatusRequiresCustomerAction, PaymentListResponseItemsStatusRequiresMerchantAction, PaymentListResponseItemsStatusRequiresPaymentMethod, PaymentListResponseItemsStatusRequiresConfirmation, PaymentListResponseItemsStatusRequiresCapture, PaymentListResponseItemsStatusPartiallyCaptured, PaymentListResponseItemsStatusPartiallyCapturedAndCapturable:
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
