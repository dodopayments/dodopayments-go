// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// CustomerWalletService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerWalletService] method instead.
type CustomerWalletService struct {
	Options       []option.RequestOption
	LedgerEntries *CustomerWalletLedgerEntryService
}

// NewCustomerWalletService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerWalletService(opts ...option.RequestOption) (r *CustomerWalletService) {
	r = &CustomerWalletService{}
	r.Options = opts
	r.LedgerEntries = NewCustomerWalletLedgerEntryService(opts...)
	return
}

func (r *CustomerWalletService) List(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerWalletListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/wallets", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomerWallet struct {
	Balance    int64              `json:"balance,required"`
	CreatedAt  time.Time          `json:"created_at,required" format:"date-time"`
	Currency   Currency           `json:"currency,required"`
	CustomerID string             `json:"customer_id,required"`
	UpdatedAt  time.Time          `json:"updated_at,required" format:"date-time"`
	JSON       customerWalletJSON `json:"-"`
}

// customerWalletJSON contains the JSON metadata for the struct [CustomerWallet]
type customerWalletJSON struct {
	Balance     apijson.Field
	CreatedAt   apijson.Field
	Currency    apijson.Field
	CustomerID  apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerWallet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerWalletJSON) RawJSON() string {
	return r.raw
}

type CustomerWalletListResponse struct {
	Items []CustomerWallet `json:"items,required"`
	// Sum of all wallet balances converted to USD (in smallest unit)
	TotalBalanceUsd int64                          `json:"total_balance_usd,required"`
	JSON            customerWalletListResponseJSON `json:"-"`
}

// customerWalletListResponseJSON contains the JSON metadata for the struct
// [CustomerWalletListResponse]
type customerWalletListResponseJSON struct {
	Items           apijson.Field
	TotalBalanceUsd apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerWalletListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerWalletListResponseJSON) RawJSON() string {
	return r.raw
}
