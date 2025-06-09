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

// WebhookEventService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEventService] method instead.
type WebhookEventService struct {
	Options []option.RequestOption
}

// NewWebhookEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEventService(opts ...option.RequestOption) (r *WebhookEventService) {
	r = &WebhookEventService{}
	r.Options = opts
	return
}

func (r *WebhookEventService) Get(ctx context.Context, webhookEventID string, opts ...option.RequestOption) (res *WebhookEvent, err error) {
	opts = append(r.Options[:], opts...)
	if webhookEventID == "" {
		err = errors.New("missing required webhook_event_id parameter")
		return
	}
	path := fmt.Sprintf("webhook_events/%s", webhookEventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *WebhookEventService) List(ctx context.Context, query WebhookEventListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[WebhookEvent], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "webhook_events"
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

func (r *WebhookEventService) ListAutoPaging(ctx context.Context, query WebhookEventListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[WebhookEvent] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

type WebhookEvent struct {
	BusinessID        string           `json:"business_id,required"`
	CreatedAt         time.Time        `json:"created_at,required" format:"date-time"`
	EventID           string           `json:"event_id,required"`
	EventType         string           `json:"event_type,required"`
	ObjectID          string           `json:"object_id,required"`
	LatestAttemptedAt time.Time        `json:"latest_attempted_at,nullable" format:"date-time"`
	Request           string           `json:"request,nullable"`
	Response          string           `json:"response,nullable"`
	JSON              webhookEventJSON `json:"-"`
}

// webhookEventJSON contains the JSON metadata for the struct [WebhookEvent]
type webhookEventJSON struct {
	BusinessID        apijson.Field
	CreatedAt         apijson.Field
	EventID           apijson.Field
	EventType         apijson.Field
	ObjectID          apijson.Field
	LatestAttemptedAt apijson.Field
	Request           apijson.Field
	Response          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookEventJSON) RawJSON() string {
	return r.raw
}

type WebhookEventListParams struct {
	// Get events after this created time
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Get events created before this time
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Min : 1, Max : 100, default 10
	Limit param.Field[int64] `query:"limit"`
	// Get events history of a specific object like payment/subscription/refund/dispute
	ObjectID param.Field[string] `query:"object_id"`
	// Page number default is 0
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size default is 10 max is 100
	PageSize param.Field[int64] `query:"page_size"`
	// Filter by webhook destination
	WebhookID param.Field[string] `query:"webhook_id"`
}

// URLQuery serializes [WebhookEventListParams]'s query parameters as `url.Values`.
func (r WebhookEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
