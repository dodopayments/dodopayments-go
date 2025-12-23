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
func NewDisputeService(opts ...option.RequestOption) (r DisputeService) {
	r = DisputeService{}
	r.Options = opts
	return
}

func (r *DisputeService) Get(ctx context.Context, disputeID string, opts ...option.RequestOption) (res *GetDispute, err error) {
	opts = slices.Concat(r.Options, opts)
	if disputeID == "" {
		err = errors.New("missing required dispute_id parameter")
		return
	}
	path := fmt.Sprintf("disputes/%s", disputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *DisputeService) List(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[DisputeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

func (r *DisputeService) ListAutoPaging(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[DisputeListResponse] {
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
	DisputeID string `json:"dispute_id,required"`
	// The current stage of the dispute process.
	//
	// Any of "pre_dispute", "dispute", "pre_arbitration".
	DisputeStage DisputeStage `json:"dispute_stage,required"`
	// The current status of the dispute.
	//
	// Any of "dispute_opened", "dispute_expired", "dispute_accepted",
	// "dispute_cancelled", "dispute_challenged", "dispute_won", "dispute_lost".
	DisputeStatus DisputeStatus `json:"dispute_status,required"`
	// The unique identifier of the payment associated with the dispute.
	PaymentID string `json:"payment_id,required"`
	// Remarks
	Remarks string `json:"remarks,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		BusinessID    respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		DisputeID     respjson.Field
		DisputeStage  respjson.Field
		DisputeStatus respjson.Field
		PaymentID     respjson.Field
		Remarks       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Dispute) RawJSON() string { return r.JSON.raw }
func (r *Dispute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeStage string

const (
	DisputeStagePreDispute     DisputeStage = "pre_dispute"
	DisputeStageDispute        DisputeStage = "dispute"
	DisputeStagePreArbitration DisputeStage = "pre_arbitration"
)

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

type GetDispute struct {
	// The amount involved in the dispute, represented as a string to accommodate
	// precision.
	Amount string `json:"amount,required"`
	// The unique identifier of the business involved in the dispute.
	BusinessID string `json:"business_id,required"`
	// The timestamp of when the dispute was created, in UTC.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the disputed amount, represented as an ISO 4217 currency code.
	Currency string `json:"currency,required"`
	// The customer who filed the dispute
	Customer CustomerLimitedDetails `json:"customer,required"`
	// The unique identifier of the dispute.
	DisputeID string `json:"dispute_id,required"`
	// The current stage of the dispute process.
	//
	// Any of "pre_dispute", "dispute", "pre_arbitration".
	DisputeStage DisputeStage `json:"dispute_stage,required"`
	// The current status of the dispute.
	//
	// Any of "dispute_opened", "dispute_expired", "dispute_accepted",
	// "dispute_cancelled", "dispute_challenged", "dispute_won", "dispute_lost".
	DisputeStatus DisputeStatus `json:"dispute_status,required"`
	// The unique identifier of the payment associated with the dispute.
	PaymentID string `json:"payment_id,required"`
	// Reason for the dispute
	Reason string `json:"reason,nullable"`
	// Remarks
	Remarks string `json:"remarks,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		BusinessID    respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Customer      respjson.Field
		DisputeID     respjson.Field
		DisputeStage  respjson.Field
		DisputeStatus respjson.Field
		PaymentID     respjson.Field
		Reason        respjson.Field
		Remarks       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetDispute) RawJSON() string { return r.JSON.raw }
func (r *GetDispute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeListResponse struct {
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
	DisputeID string `json:"dispute_id,required"`
	// The current stage of the dispute process.
	//
	// Any of "pre_dispute", "dispute", "pre_arbitration".
	DisputeStage DisputeStage `json:"dispute_stage,required"`
	// The current status of the dispute.
	//
	// Any of "dispute_opened", "dispute_expired", "dispute_accepted",
	// "dispute_cancelled", "dispute_challenged", "dispute_won", "dispute_lost".
	DisputeStatus DisputeStatus `json:"dispute_status,required"`
	// The unique identifier of the payment associated with the dispute.
	PaymentID string `json:"payment_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		BusinessID    respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		DisputeID     respjson.Field
		DisputeStage  respjson.Field
		DisputeStatus respjson.Field
		PaymentID     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeListResponse) RawJSON() string { return r.JSON.raw }
func (r *DisputeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeListParams struct {
	// Get events after this created time
	CreatedAtGte param.Opt[time.Time] `query:"created_at_gte,omitzero" format:"date-time" json:"-"`
	// Get events created before this time
	CreatedAtLte param.Opt[time.Time] `query:"created_at_lte,omitzero" format:"date-time" json:"-"`
	// Filter by customer_id
	CustomerID param.Opt[string] `query:"customer_id,omitzero" json:"-"`
	// Page number default is 0
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Page size default is 10 max is 100
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter by dispute stage
	//
	// Any of "pre_dispute", "dispute", "pre_arbitration".
	DisputeStage DisputeListParamsDisputeStage `query:"dispute_stage,omitzero" json:"-"`
	// Filter by dispute status
	//
	// Any of "dispute_opened", "dispute_expired", "dispute_accepted",
	// "dispute_cancelled", "dispute_challenged", "dispute_won", "dispute_lost".
	DisputeStatus DisputeListParamsDisputeStatus `query:"dispute_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DisputeListParams]'s query parameters as `url.Values`.
func (r DisputeListParams) URLQuery() (v url.Values, err error) {
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
