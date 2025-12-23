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
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// CustomerWalletService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerWalletService] method instead.
type CustomerWalletService struct {
	Options       []option.RequestOption
	LedgerEntries CustomerWalletLedgerEntryService
}

// NewCustomerWalletService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerWalletService(opts ...option.RequestOption) (r CustomerWalletService) {
	r = CustomerWalletService{}
	r.Options = opts
	r.LedgerEntries = NewCustomerWalletLedgerEntryService(opts...)
	return
}

func (r *CustomerWalletService) List(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerWalletListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/wallets", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomerWallet struct {
	Balance   int64     `json:"balance,required"`
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
	Currency   Currency  `json:"currency,required"`
	CustomerID string    `json:"customer_id,required"`
	UpdatedAt  time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balance     respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		CustomerID  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerWallet) RawJSON() string { return r.JSON.raw }
func (r *CustomerWallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerWalletListResponse struct {
	Items []CustomerWallet `json:"items,required"`
	// Sum of all wallet balances converted to USD (in smallest unit)
	TotalBalanceUsd int64 `json:"total_balance_usd,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items           respjson.Field
		TotalBalanceUsd respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerWalletListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerWalletListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
