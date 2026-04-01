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

// PayoutBreakupDetailService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayoutBreakupDetailService] method instead.
type PayoutBreakupDetailService struct {
	Options []option.RequestOption
}

// NewPayoutBreakupDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPayoutBreakupDetailService(opts ...option.RequestOption) (r *PayoutBreakupDetailService) {
	r = &PayoutBreakupDetailService{}
	r.Options = opts
	return
}

// Returns paginated individual balance ledger entries for a payout, with each
// entry's amount pro-rated into the payout's currency. Supports pagination via
// `page_size` (default 10, max 100) and `page_number` (default 0) query
// parameters.
func (r *PayoutBreakupDetailService) List(ctx context.Context, payoutID string, query PayoutBreakupDetailListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[PayoutBreakupDetailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if payoutID == "" {
		err = errors.New("missing required payout_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("payouts/%s/breakup/details", payoutID)
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

// Returns paginated individual balance ledger entries for a payout, with each
// entry's amount pro-rated into the payout's currency. Supports pagination via
// `page_size` (default 10, max 100) and `page_number` (default 0) query
// parameters.
func (r *PayoutBreakupDetailService) ListAutoPaging(ctx context.Context, payoutID string, query PayoutBreakupDetailListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[PayoutBreakupDetailListResponse] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, payoutID, query, opts...))
}

// Downloads the complete payout breakup as a CSV file. Each row represents a
// balance ledger entry with columns: Ledger ID, Event Type, Original Amount,
// Original Currency, Reference Object ID, Description, Created At, USD Equivalent
// Amount, and Payout Currency Amount.
func (r *PayoutBreakupDetailService) DownloadCsv(ctx context.Context, payoutID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if payoutID == "" {
		err = errors.New("missing required payout_id parameter")
		return err
	}
	path := fmt.Sprintf("payouts/%s/breakup/details/csv", payoutID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Individual balance ledger entry for a payout, with amounts pro-rated into the
// payout's currency.
type PayoutBreakupDetailListResponse struct {
	// Unique identifier of the balance ledger entry.
	ID string `json:"id" api:"required"`
	// Timestamp when this entry was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The type of balance ledger event (e.g., "payment", "refund", "dispute",
	// "payment_fees").
	EventType string `json:"event_type" api:"required"`
	// Original amount in the original currency (in smallest currency unit, e.g.,
	// cents).
	OriginalAmount int64 `json:"original_amount" api:"required"`
	// Original currency as ISO 4217 code (e.g., "USD", "EUR").
	OriginalCurrency string `json:"original_currency" api:"required"`
	// Amount in the payout's currency (in smallest currency unit). Uses cumulative
	// rounding to ensure sum matches payout total exactly.
	PayoutCurrencyAmount int64 `json:"payout_currency_amount" api:"required"`
	// USD equivalent of the original amount (in cents).
	UsdEquivalentAmount int64 `json:"usd_equivalent_amount" api:"required"`
	// Human-readable description of the transaction.
	Description string `json:"description" api:"nullable"`
	// ID of the related object (e.g., payment ID, refund ID) if applicable.
	ReferenceObjectID string                              `json:"reference_object_id" api:"nullable"`
	JSON              payoutBreakupDetailListResponseJSON `json:"-"`
}

// payoutBreakupDetailListResponseJSON contains the JSON metadata for the struct
// [PayoutBreakupDetailListResponse]
type payoutBreakupDetailListResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	EventType            apijson.Field
	OriginalAmount       apijson.Field
	OriginalCurrency     apijson.Field
	PayoutCurrencyAmount apijson.Field
	UsdEquivalentAmount  apijson.Field
	Description          apijson.Field
	ReferenceObjectID    apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PayoutBreakupDetailListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payoutBreakupDetailListResponseJSON) RawJSON() string {
	return r.raw
}

type PayoutBreakupDetailListParams struct {
	// Page number (0-indexed). Default: 0.
	PageNumber param.Field[int64] `query:"page_number"`
	// Number of items per page. Default: 10, Max: 100.
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [PayoutBreakupDetailListParams]'s query parameters as
// `url.Values`.
func (r PayoutBreakupDetailListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
