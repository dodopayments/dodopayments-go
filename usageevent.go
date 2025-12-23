// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"encoding/json"
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

// UsageEventService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageEventService] method instead.
type UsageEventService struct {
	Options []option.RequestOption
}

// NewUsageEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUsageEventService(opts ...option.RequestOption) (r UsageEventService) {
	r = UsageEventService{}
	r.Options = opts
	return
}

// Fetch detailed information about a single event using its unique event ID. This
// endpoint is useful for:
//
// - Debugging specific event ingestion issues
// - Retrieving event details for customer support
// - Validating that events were processed correctly
// - Getting the complete metadata for an event
//
// ## Event ID Format:
//
// The event ID should be the same value that was provided during event ingestion
// via the `/events/ingest` endpoint. Event IDs are case-sensitive and must match
// exactly.
//
// ## Response Details:
//
// The response includes all event data including:
//
// - Complete metadata key-value pairs
// - Original timestamp (preserved from ingestion)
// - Customer and business association
// - Event name and processing information
//
// ## Example Usage:
//
// ```text
// GET /events/api_call_12345
// ```
func (r *UsageEventService) Get(ctx context.Context, eventID string, opts ...option.RequestOption) (res *Event, err error) {
	opts = slices.Concat(r.Options, opts)
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("events/%s", eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch events from your account with powerful filtering capabilities. This
// endpoint is ideal for:
//
// - Debugging event ingestion issues
// - Analyzing customer usage patterns
// - Building custom analytics dashboards
// - Auditing billing-related events
//
// ## Filtering Options:
//
//   - **Customer filtering**: Filter by specific customer ID
//   - **Event name filtering**: Filter by event type/name
//   - **Meter-based filtering**: Use a meter ID to apply the meter's event name and
//     filter criteria automatically
//   - **Time range filtering**: Filter events within a specific date range
//   - **Pagination**: Navigate through large result sets
//
// ## Meter Integration:
//
// When using `meter_id`, the endpoint automatically applies:
//
// - The meter's configured `event_name` filter
// - The meter's custom filter criteria (if any)
// - If you also provide `event_name`, it must match the meter's event name
//
// ## Example Queries:
//
//   - Get all events for a customer: `?customer_id=cus_abc123`
//   - Get API request events: `?event_name=api_request`
//   - Get events from last 24 hours:
//     `?start=2024-01-14T10:30:00Z&end=2024-01-15T10:30:00Z`
//   - Get events with meter filtering: `?meter_id=mtr_xyz789`
//   - Paginate results: `?page_size=50&page_number=2`
func (r *UsageEventService) List(ctx context.Context, query UsageEventListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Event], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "events"
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

// Fetch events from your account with powerful filtering capabilities. This
// endpoint is ideal for:
//
// - Debugging event ingestion issues
// - Analyzing customer usage patterns
// - Building custom analytics dashboards
// - Auditing billing-related events
//
// ## Filtering Options:
//
//   - **Customer filtering**: Filter by specific customer ID
//   - **Event name filtering**: Filter by event type/name
//   - **Meter-based filtering**: Use a meter ID to apply the meter's event name and
//     filter criteria automatically
//   - **Time range filtering**: Filter events within a specific date range
//   - **Pagination**: Navigate through large result sets
//
// ## Meter Integration:
//
// When using `meter_id`, the endpoint automatically applies:
//
// - The meter's configured `event_name` filter
// - The meter's custom filter criteria (if any)
// - If you also provide `event_name`, it must match the meter's event name
//
// ## Example Queries:
//
//   - Get all events for a customer: `?customer_id=cus_abc123`
//   - Get API request events: `?event_name=api_request`
//   - Get events from last 24 hours:
//     `?start=2024-01-14T10:30:00Z&end=2024-01-15T10:30:00Z`
//   - Get events with meter filtering: `?meter_id=mtr_xyz789`
//   - Paginate results: `?page_size=50&page_number=2`
func (r *UsageEventService) ListAutoPaging(ctx context.Context, query UsageEventListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Event] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// This endpoint allows you to ingest custom events that can be used for:
//
// - Usage-based billing and metering
// - Analytics and reporting
// - Customer behavior tracking
//
// ## Important Notes:
//
// - **Duplicate Prevention**:
//   - Duplicate `event_id` values within the same request are rejected (entire
//     request fails)
//   - Subsequent requests with existing `event_id` values are ignored (idempotent
//     behavior)
//   - **Rate Limiting**: Maximum 1000 events per request
//   - **Time Validation**: Events with timestamps older than 1 hour or more than 5
//     minutes in the future will be rejected
//   - **Metadata Limits**: Maximum 50 key-value pairs per event, keys max 100 chars,
//     values max 500 chars
//
// ## Example Usage:
//
// ```json
//
//	{
//	  "events": [
//	    {
//	      "event_id": "api_call_12345",
//	      "customer_id": "cus_abc123",
//	      "event_name": "api_request",
//	      "timestamp": "2024-01-15T10:30:00Z",
//	      "metadata": {
//	        "endpoint": "/api/v1/users",
//	        "method": "GET",
//	        "tokens_used": "150"
//	      }
//	    }
//	  ]
//	}
//
// ```
func (r *UsageEventService) Ingest(ctx context.Context, body UsageEventIngestParams, opts ...option.RequestOption) (res *UsageEventIngestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "events/ingest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Event struct {
	BusinessID string    `json:"business_id,required"`
	CustomerID string    `json:"customer_id,required"`
	EventID    string    `json:"event_id,required"`
	EventName  string    `json:"event_name,required"`
	Timestamp  time.Time `json:"timestamp,required" format:"date-time"`
	// Arbitrary key-value metadata. Values can be string, integer, number, or boolean.
	Metadata map[string]EventMetadataUnion `json:"metadata,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		CustomerID  respjson.Field
		EventID     respjson.Field
		EventName   respjson.Field
		Timestamp   respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventMetadataUnion contains all possible properties and values from [string],
// [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type EventMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u EventMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *EventMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CustomerID, EventID, EventName are required.
type EventInputParam struct {
	// customer_id of the customer whose usage needs to be tracked
	CustomerID string `json:"customer_id,required"`
	// Event Id acts as an idempotency key. Any subsequent requests with the same
	// event_id will be ignored
	EventID string `json:"event_id,required"`
	// Name of the event
	EventName string `json:"event_name,required"`
	// Custom Timestamp. Defaults to current timestamp in UTC. Timestamps that are
	// older that 1 hour or after 5 mins, from current timestamp, will be rejected.
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	// Custom metadata. Only key value pairs are accepted, objects or arrays submitted
	// will be rejected.
	Metadata map[string]EventInputMetadataUnionParam `json:"metadata,omitzero"`
	paramObj
}

func (r EventInputParam) MarshalJSON() (data []byte, err error) {
	type shadow EventInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventInputMetadataUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u EventInputMetadataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *EventInputMetadataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EventInputMetadataUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type UsageEventIngestResponse struct {
	IngestedCount int64 `json:"ingested_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IngestedCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageEventIngestResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageEventIngestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageEventListParams struct {
	// Filter events by customer ID
	CustomerID param.Opt[string] `query:"customer_id,omitzero" json:"-"`
	// Filter events created before this timestamp
	End param.Opt[time.Time] `query:"end,omitzero" format:"date-time" json:"-"`
	// Filter events by event name. If both event_name and meter_id are provided, they
	// must match the meter's configured event_name
	EventName param.Opt[string] `query:"event_name,omitzero" json:"-"`
	// Filter events by meter ID. When provided, only events that match the meter's
	// event_name and filter criteria will be returned
	MeterID param.Opt[string] `query:"meter_id,omitzero" json:"-"`
	// Page number (0-based, default: 0)
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// Number of events to return per page (default: 10)
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter events created after this timestamp
	Start param.Opt[time.Time] `query:"start,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [UsageEventListParams]'s query parameters as `url.Values`.
func (r UsageEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UsageEventIngestParams struct {
	// List of events to be pushed
	Events []EventInputParam `json:"events,omitzero,required"`
	paramObj
}

func (r UsageEventIngestParams) MarshalJSON() (data []byte, err error) {
	type shadow UsageEventIngestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsageEventIngestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
