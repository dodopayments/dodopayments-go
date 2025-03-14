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
	DisputeID     string               `json:"dispute_id,required"`
	DisputeStage  DisputeDisputeStage  `json:"dispute_stage,required"`
	DisputeStatus DisputeDisputeStatus `json:"dispute_status,required"`
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

type DisputeDisputeStage string

const (
	DisputeDisputeStagePreDispute     DisputeDisputeStage = "pre_dispute"
	DisputeDisputeStageDispute        DisputeDisputeStage = "dispute"
	DisputeDisputeStagePreArbitration DisputeDisputeStage = "pre_arbitration"
)

func (r DisputeDisputeStage) IsKnown() bool {
	switch r {
	case DisputeDisputeStagePreDispute, DisputeDisputeStageDispute, DisputeDisputeStagePreArbitration:
		return true
	}
	return false
}

type DisputeDisputeStatus string

const (
	DisputeDisputeStatusDisputeOpened     DisputeDisputeStatus = "dispute_opened"
	DisputeDisputeStatusDisputeExpired    DisputeDisputeStatus = "dispute_expired"
	DisputeDisputeStatusDisputeAccepted   DisputeDisputeStatus = "dispute_accepted"
	DisputeDisputeStatusDisputeCancelled  DisputeDisputeStatus = "dispute_cancelled"
	DisputeDisputeStatusDisputeChallenged DisputeDisputeStatus = "dispute_challenged"
	DisputeDisputeStatusDisputeWon        DisputeDisputeStatus = "dispute_won"
	DisputeDisputeStatusDisputeLost       DisputeDisputeStatus = "dispute_lost"
)

func (r DisputeDisputeStatus) IsKnown() bool {
	switch r {
	case DisputeDisputeStatusDisputeOpened, DisputeDisputeStatusDisputeExpired, DisputeDisputeStatusDisputeAccepted, DisputeDisputeStatusDisputeCancelled, DisputeDisputeStatusDisputeChallenged, DisputeDisputeStatusDisputeWon, DisputeDisputeStatusDisputeLost:
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
	DisputeStage param.Field[DisputeListParamsDisputeStage] `query:"dispute_stage"`
	// Filter by dispute status
	DisputeStatus param.Field[DisputeListParamsDisputeStatus] `query:"dispute_status"`
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

// Filter by dispute stage
type DisputeListParamsDisputeStage string

const (
	DisputeListParamsDisputeStagePreDispute     DisputeListParamsDisputeStage = "pre_dispute"
	DisputeListParamsDisputeStageDispute        DisputeListParamsDisputeStage = "dispute"
	DisputeListParamsDisputeStagePreArbitration DisputeListParamsDisputeStage = "pre_arbitration"
)

func (r DisputeListParamsDisputeStage) IsKnown() bool {
	switch r {
	case DisputeListParamsDisputeStagePreDispute, DisputeListParamsDisputeStageDispute, DisputeListParamsDisputeStagePreArbitration:
		return true
	}
	return false
}

// Filter by dispute status
type DisputeListParamsDisputeStatus string

const (
	DisputeListParamsDisputeStatusDisputeOpened     DisputeListParamsDisputeStatus = "dispute_opened"
	DisputeListParamsDisputeStatusDisputeExpired    DisputeListParamsDisputeStatus = "dispute_expired"
	DisputeListParamsDisputeStatusDisputeAccepted   DisputeListParamsDisputeStatus = "dispute_accepted"
	DisputeListParamsDisputeStatusDisputeCancelled  DisputeListParamsDisputeStatus = "dispute_cancelled"
	DisputeListParamsDisputeStatusDisputeChallenged DisputeListParamsDisputeStatus = "dispute_challenged"
	DisputeListParamsDisputeStatusDisputeWon        DisputeListParamsDisputeStatus = "dispute_won"
	DisputeListParamsDisputeStatusDisputeLost       DisputeListParamsDisputeStatus = "dispute_lost"
)

func (r DisputeListParamsDisputeStatus) IsKnown() bool {
	switch r {
	case DisputeListParamsDisputeStatusDisputeOpened, DisputeListParamsDisputeStatusDisputeExpired, DisputeListParamsDisputeStatusDisputeAccepted, DisputeListParamsDisputeStatusDisputeCancelled, DisputeListParamsDisputeStatusDisputeChallenged, DisputeListParamsDisputeStatusDisputeWon, DisputeListParamsDisputeStatusDisputeLost:
		return true
	}
	return false
}
