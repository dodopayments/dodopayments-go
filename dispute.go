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

// DisputeService contains methods and other services that help with interacting
// with the dodopayments API.
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

func (r *DisputeService) List(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) (res *DisputeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "disputes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Dispute struct {
	Amount        string               `json:"amount,required"`
	BusinessID    string               `json:"business_id,required"`
	CreatedAt     time.Time            `json:"created_at,required" format:"date-time"`
	Currency      string               `json:"currency,required"`
	DisputeID     string               `json:"dispute_id,required"`
	DisputeStage  DisputeDisputeStage  `json:"dispute_stage,required"`
	DisputeStatus DisputeDisputeStatus `json:"dispute_status,required"`
	PaymentID     string               `json:"payment_id,required"`
	JSON          disputeJSON          `json:"-"`
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

type DisputeParam struct {
	Amount        param.Field[string]               `json:"amount,required"`
	BusinessID    param.Field[string]               `json:"business_id,required"`
	CreatedAt     param.Field[time.Time]            `json:"created_at,required" format:"date-time"`
	Currency      param.Field[string]               `json:"currency,required"`
	DisputeID     param.Field[string]               `json:"dispute_id,required"`
	DisputeStage  param.Field[DisputeDisputeStage]  `json:"dispute_stage,required"`
	DisputeStatus param.Field[DisputeDisputeStatus] `json:"dispute_status,required"`
	PaymentID     param.Field[string]               `json:"payment_id,required"`
}

func (r DisputeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DisputeListResponse struct {
	Items []Dispute               `json:"items,required"`
	JSON  disputeListResponseJSON `json:"-"`
}

// disputeListResponseJSON contains the JSON metadata for the struct
// [DisputeListResponse]
type disputeListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeListResponseJSON) RawJSON() string {
	return r.raw
}

type DisputeListParams struct {
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
