// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// PayoutBreakupService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayoutBreakupService] method instead.
type PayoutBreakupService struct {
	Options []option.RequestOption
	Details *PayoutBreakupDetailService
}

// NewPayoutBreakupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPayoutBreakupService(opts ...option.RequestOption) (r *PayoutBreakupService) {
	r = &PayoutBreakupService{}
	r.Options = opts
	r.Details = NewPayoutBreakupDetailService(opts...)
	return
}

// Returns the breakdown of a payout by event type (payments, refunds, disputes,
// fees, etc.) in the payout's currency. Each amount is proportionally allocated
// based on USD equivalent values, ensuring the total sums exactly to the payout
// amount.
func (r *PayoutBreakupService) Get(ctx context.Context, payoutID string, opts ...option.RequestOption) (res *[]PayoutBreakupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if payoutID == "" {
		err = errors.New("missing required payout_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("payouts/%s/breakup", payoutID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Payout breakup aggregated by event type, with amounts in the payout's currency.
type PayoutBreakupGetResponse struct {
	// The type of balance ledger event (e.g., "payment", "refund", "dispute",
	// "payment_fees").
	EventType string `json:"event_type" api:"required"`
	// Total amount for this event type in the payout's currency (in smallest currency
	// unit, e.g., cents).
	Total int64                        `json:"total" api:"required"`
	JSON  payoutBreakupGetResponseJSON `json:"-"`
}

// payoutBreakupGetResponseJSON contains the JSON metadata for the struct
// [PayoutBreakupGetResponse]
type payoutBreakupGetResponseJSON struct {
	EventType   apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayoutBreakupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutBreakupGetResponseJSON) RawJSON() string {
	return r.raw
}
