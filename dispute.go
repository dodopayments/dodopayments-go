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

// DisputeService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDisputeService] method instead.
type DisputeService struct {
	Options []option.RequestOption
}

// NewDisputeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDisputeService(opts ...option.RequestOption) (r *DisputeService) {
	r = &DisputeService{}
	r.Options = opts
	return
}

func (r *DisputeService) Get(ctx context.Context, disputeID string, opts ...option.RequestOption) (res *Dispute, err error) {
	opts = append(r.Options[:], opts...)
	if disputeID == "" {
		err = errors.New("missing required dispute_id parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s", disputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *DisputeService) List(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Dispute], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "disputes"
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

func (r *DisputeService) ListAutoPaging(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Dispute] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type Dispute struct {
	// The amount involved in the dispute, represented as a string to accommodate
	// precision.
	Amount string `json:"amount,required"`
	// The unique identifier of the business involved in the dispute.
	BusinessID string `json:"business_id,required"`
	// The timestamp of when the dispute was created, in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the disputed amount, represented as an ISO 4217 currency code.
	Currency string `json:"currency,required"`
	// The unique identifier of the dispute.
	DisputeID     string        `json:"dispute_id,required"`
	DisputeStage  DisputeStage  `json:"dispute_stage,required"`
	DisputeStatus DisputeStatus `json:"dispute_status,required"`
	// The unique identifier of the payment associated with the dispute.
	PaymentID string      `json:"payment_id,required"`
	JSON      disputeJSON `json:"-"`
}

// disputeJSON contains the JSON metadata for the struct [Dispute]
type disputeJSON struct {
	Amount        apijson.Field
	BusinessID    apijson.Field
	CreatedAt     apijson.Field
	Currency      apijson.Field
	DisputeID     apijson.Field
	DisputeStage  apijson.Field
	DisputeStatus apijson.Field
	PaymentID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Dispute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeJSON) RawJSON() string {
	return r.raw
}

type DisputeStage string

const (
	DisputeStagePreDispute     DisputeStage = "pre_dispute"
	DisputeStageDispute        DisputeStage = "dispute"
	DisputeStagePreArbitration DisputeStage = "pre_arbitration"
)

func (r DisputeStage) IsKnown() bool {
	switch r {
	case DisputeStagePreDispute, DisputeStageDispute, DisputeStagePreArbitration:
		return true
	}
	return false
}

type DisputeStatus string

const (
	DisputeStatusDisputeOpened     DisputeStatus = "dispute_opened"
	DisputeStatusDisputeExpired    DisputeStatus = "dispute_expired"
	DisputeStatusDisputeAccepted   DisputeStatus = "dispute_accepted"
	DisputeStatusDisputeCancelled  DisputeStatus = "dispute_cancelled"
	DisputeStatusDisputeChallenged DisputeStatus = "dispute_challenged"
	DisputeStatusDisputeWon        DisputeStatus = "dispute_won"
	DisputeStatusDisputeLost       DisputeStatus = "dispute_lost"
)

func (r DisputeStatus) IsKnown() bool {
	switch r {
	case DisputeStatusDisputeOpened, DisputeStatusDisputeExpired, DisputeStatusDisputeAccepted, DisputeStatusDisputeCancelled, DisputeStatusDisputeChallenged, DisputeStatusDisputeWon, DisputeStatusDisputeLost:
		return true
	}
	return false
}

type DisputeListParams struct {
	// Get events after this created time
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get events created before this time
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Filter by customer_id
	CustomerID param.Field[string] `query:"customer_id"`
	// Filter by dispute stage
	DisputeStage param.Field[DisputeStage] `query:"dispute_stage"`
	// Filter by dispute status
	DisputeStatus param.Field[DisputeStatus] `query:"dispute_status"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [DisputeListParams]'s query parameters as `url.Values`.
func (r DisputeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
