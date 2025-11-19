// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/tidwall/gjson"
)

// WebhookService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
	Headers *WebhookHeaderService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	r.Headers = NewWebhookHeaderService(opts...)
	return
}

// Create a new webhook
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *WebhookDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a webhook by id
func (r *WebhookService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch a webhook by id
func (r *WebhookService) Update(ctx context.Context, webhookID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookDetails, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all webhooks
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *pagination.CursorPagePagination[WebhookDetails], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "webhooks"
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

// List all webhooks
func (r *WebhookService) ListAutoPaging(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) *pagination.CursorPagePaginationAutoPager[WebhookDetails] {
	return pagination.NewCursorPagePaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a webhook by id
func (r *WebhookService) Delete(ctx context.Context, webhookID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get webhook secret by id
func (r *WebhookService) GetSecret(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookGetSecretResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/secret", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *WebhookService) UnsafeUnwrap(payload []byte, opts ...option.RequestOption) (*UnsafeUnwrapWebhookEvent, error) {
	res := &UnsafeUnwrapWebhookEvent{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEvent, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookKey
	if key == "" {
		return nil, errors.New("The WebhookKey option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &UnwrapWebhookEvent{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type WebhookDetails struct {
	// The webhook's ID.
	ID string `json:"id,required"`
	// Created at timestamp
	CreatedAt string `json:"created_at,required"`
	// An example webhook name.
	Description string `json:"description,required"`
	// Metadata of the webhook
	Metadata map[string]string `json:"metadata,required"`
	// Updated at timestamp
	UpdatedAt string `json:"updated_at,required"`
	// Url endpoint of the webhook
	URL string `json:"url,required"`
	// Status of the webhook.
	//
	// If true, events are not sent
	Disabled bool `json:"disabled,nullable"`
	// Filter events to the webhook.
	//
	// Webhook event will only be sent for events in the list.
	FilterTypes []string `json:"filter_types,nullable"`
	// Configured rate limit
	RateLimit int64              `json:"rate_limit,nullable"`
	JSON      webhookDetailsJSON `json:"-"`
}

// webhookDetailsJSON contains the JSON metadata for the struct [WebhookDetails]
type webhookDetailsJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Metadata    apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	Disabled    apijson.Field
	FilterTypes apijson.Field
	RateLimit   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDetailsJSON) RawJSON() string {
	return r.raw
}

type WebhookGetSecretResponse struct {
	Secret string                       `json:"secret,required"`
	JSON   webhookGetSecretResponseJSON `json:"-"`
}

// webhookGetSecretResponseJSON contains the JSON metadata for the struct
// [WebhookGetSecretResponse]
type webhookGetSecretResponseJSON struct {
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetSecretResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetSecretResponseJSON) RawJSON() string {
	return r.raw
}

type DisputeAcceptedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeAcceptedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type DisputeAcceptedWebhookEventType `json:"type,required"`
	JSON disputeAcceptedWebhookEventJSON `json:"-"`
}

// disputeAcceptedWebhookEventJSON contains the JSON metadata for the struct
// [DisputeAcceptedWebhookEvent]
type disputeAcceptedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeAcceptedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeAcceptedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeAcceptedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r DisputeAcceptedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type DisputeAcceptedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType DisputeAcceptedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        disputeAcceptedWebhookEventDataJSON        `json:"-"`
	Dispute
}

// disputeAcceptedWebhookEventDataJSON contains the JSON metadata for the struct
// [DisputeAcceptedWebhookEventData]
type disputeAcceptedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeAcceptedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeAcceptedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type DisputeAcceptedWebhookEventDataPayloadType string

const (
	DisputeAcceptedWebhookEventDataPayloadTypeDispute DisputeAcceptedWebhookEventDataPayloadType = "Dispute"
)

func (r DisputeAcceptedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case DisputeAcceptedWebhookEventDataPayloadTypeDispute:
		return true
	}
	return false
}

// The event type
type DisputeAcceptedWebhookEventType string

const (
	DisputeAcceptedWebhookEventTypeDisputeAccepted DisputeAcceptedWebhookEventType = "dispute.accepted"
)

func (r DisputeAcceptedWebhookEventType) IsKnown() bool {
	switch r {
	case DisputeAcceptedWebhookEventTypeDisputeAccepted:
		return true
	}
	return false
}

type DisputeCancelledWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeCancelledWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type DisputeCancelledWebhookEventType `json:"type,required"`
	JSON disputeCancelledWebhookEventJSON `json:"-"`
}

// disputeCancelledWebhookEventJSON contains the JSON metadata for the struct
// [DisputeCancelledWebhookEvent]
type disputeCancelledWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeCancelledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeCancelledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeCancelledWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r DisputeCancelledWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type DisputeCancelledWebhookEventData struct {
	// The type of payload in the data field
	PayloadType DisputeCancelledWebhookEventDataPayloadType `json:"payload_type"`
	JSON        disputeCancelledWebhookEventDataJSON        `json:"-"`
	Dispute
}

// disputeCancelledWebhookEventDataJSON contains the JSON metadata for the struct
// [DisputeCancelledWebhookEventData]
type disputeCancelledWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeCancelledWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeCancelledWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type DisputeCancelledWebhookEventDataPayloadType string

const (
	DisputeCancelledWebhookEventDataPayloadTypeDispute DisputeCancelledWebhookEventDataPayloadType = "Dispute"
)

func (r DisputeCancelledWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case DisputeCancelledWebhookEventDataPayloadTypeDispute:
		return true
	}
	return false
}

// The event type
type DisputeCancelledWebhookEventType string

const (
	DisputeCancelledWebhookEventTypeDisputeCancelled DisputeCancelledWebhookEventType = "dispute.cancelled"
)

func (r DisputeCancelledWebhookEventType) IsKnown() bool {
	switch r {
	case DisputeCancelledWebhookEventTypeDisputeCancelled:
		return true
	}
	return false
}

type DisputeChallengedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeChallengedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type DisputeChallengedWebhookEventType `json:"type,required"`
	JSON disputeChallengedWebhookEventJSON `json:"-"`
}

// disputeChallengedWebhookEventJSON contains the JSON metadata for the struct
// [DisputeChallengedWebhookEvent]
type disputeChallengedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeChallengedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeChallengedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeChallengedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r DisputeChallengedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type DisputeChallengedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType DisputeChallengedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        disputeChallengedWebhookEventDataJSON        `json:"-"`
	Dispute
}

// disputeChallengedWebhookEventDataJSON contains the JSON metadata for the struct
// [DisputeChallengedWebhookEventData]
type disputeChallengedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeChallengedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeChallengedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type DisputeChallengedWebhookEventDataPayloadType string

const (
	DisputeChallengedWebhookEventDataPayloadTypeDispute DisputeChallengedWebhookEventDataPayloadType = "Dispute"
)

func (r DisputeChallengedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case DisputeChallengedWebhookEventDataPayloadTypeDispute:
		return true
	}
	return false
}

// The event type
type DisputeChallengedWebhookEventType string

const (
	DisputeChallengedWebhookEventTypeDisputeChallenged DisputeChallengedWebhookEventType = "dispute.challenged"
)

func (r DisputeChallengedWebhookEventType) IsKnown() bool {
	switch r {
	case DisputeChallengedWebhookEventTypeDisputeChallenged:
		return true
	}
	return false
}

type DisputeExpiredWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeExpiredWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type DisputeExpiredWebhookEventType `json:"type,required"`
	JSON disputeExpiredWebhookEventJSON `json:"-"`
}

// disputeExpiredWebhookEventJSON contains the JSON metadata for the struct
// [DisputeExpiredWebhookEvent]
type disputeExpiredWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeExpiredWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeExpiredWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeExpiredWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r DisputeExpiredWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type DisputeExpiredWebhookEventData struct {
	// The type of payload in the data field
	PayloadType DisputeExpiredWebhookEventDataPayloadType `json:"payload_type"`
	JSON        disputeExpiredWebhookEventDataJSON        `json:"-"`
	Dispute
}

// disputeExpiredWebhookEventDataJSON contains the JSON metadata for the struct
// [DisputeExpiredWebhookEventData]
type disputeExpiredWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeExpiredWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeExpiredWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type DisputeExpiredWebhookEventDataPayloadType string

const (
	DisputeExpiredWebhookEventDataPayloadTypeDispute DisputeExpiredWebhookEventDataPayloadType = "Dispute"
)

func (r DisputeExpiredWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case DisputeExpiredWebhookEventDataPayloadTypeDispute:
		return true
	}
	return false
}

// The event type
type DisputeExpiredWebhookEventType string

const (
	DisputeExpiredWebhookEventTypeDisputeExpired DisputeExpiredWebhookEventType = "dispute.expired"
)

func (r DisputeExpiredWebhookEventType) IsKnown() bool {
	switch r {
	case DisputeExpiredWebhookEventTypeDisputeExpired:
		return true
	}
	return false
}

type DisputeLostWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeLostWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type DisputeLostWebhookEventType `json:"type,required"`
	JSON disputeLostWebhookEventJSON `json:"-"`
}

// disputeLostWebhookEventJSON contains the JSON metadata for the struct
// [DisputeLostWebhookEvent]
type disputeLostWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeLostWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeLostWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeLostWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r DisputeLostWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type DisputeLostWebhookEventData struct {
	// The type of payload in the data field
	PayloadType DisputeLostWebhookEventDataPayloadType `json:"payload_type"`
	JSON        disputeLostWebhookEventDataJSON        `json:"-"`
	Dispute
}

// disputeLostWebhookEventDataJSON contains the JSON metadata for the struct
// [DisputeLostWebhookEventData]
type disputeLostWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeLostWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeLostWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type DisputeLostWebhookEventDataPayloadType string

const (
	DisputeLostWebhookEventDataPayloadTypeDispute DisputeLostWebhookEventDataPayloadType = "Dispute"
)

func (r DisputeLostWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case DisputeLostWebhookEventDataPayloadTypeDispute:
		return true
	}
	return false
}

// The event type
type DisputeLostWebhookEventType string

const (
	DisputeLostWebhookEventTypeDisputeLost DisputeLostWebhookEventType = "dispute.lost"
)

func (r DisputeLostWebhookEventType) IsKnown() bool {
	switch r {
	case DisputeLostWebhookEventTypeDisputeLost:
		return true
	}
	return false
}

type DisputeOpenedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeOpenedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type DisputeOpenedWebhookEventType `json:"type,required"`
	JSON disputeOpenedWebhookEventJSON `json:"-"`
}

// disputeOpenedWebhookEventJSON contains the JSON metadata for the struct
// [DisputeOpenedWebhookEvent]
type disputeOpenedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeOpenedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeOpenedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeOpenedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r DisputeOpenedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type DisputeOpenedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType DisputeOpenedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        disputeOpenedWebhookEventDataJSON        `json:"-"`
	Dispute
}

// disputeOpenedWebhookEventDataJSON contains the JSON metadata for the struct
// [DisputeOpenedWebhookEventData]
type disputeOpenedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeOpenedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeOpenedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type DisputeOpenedWebhookEventDataPayloadType string

const (
	DisputeOpenedWebhookEventDataPayloadTypeDispute DisputeOpenedWebhookEventDataPayloadType = "Dispute"
)

func (r DisputeOpenedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case DisputeOpenedWebhookEventDataPayloadTypeDispute:
		return true
	}
	return false
}

// The event type
type DisputeOpenedWebhookEventType string

const (
	DisputeOpenedWebhookEventTypeDisputeOpened DisputeOpenedWebhookEventType = "dispute.opened"
)

func (r DisputeOpenedWebhookEventType) IsKnown() bool {
	switch r {
	case DisputeOpenedWebhookEventTypeDisputeOpened:
		return true
	}
	return false
}

type DisputeWonWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeWonWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type DisputeWonWebhookEventType `json:"type,required"`
	JSON disputeWonWebhookEventJSON `json:"-"`
}

// disputeWonWebhookEventJSON contains the JSON metadata for the struct
// [DisputeWonWebhookEvent]
type disputeWonWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeWonWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeWonWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DisputeWonWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r DisputeWonWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type DisputeWonWebhookEventData struct {
	// The type of payload in the data field
	PayloadType DisputeWonWebhookEventDataPayloadType `json:"payload_type"`
	JSON        disputeWonWebhookEventDataJSON        `json:"-"`
	Dispute
}

// disputeWonWebhookEventDataJSON contains the JSON metadata for the struct
// [DisputeWonWebhookEventData]
type disputeWonWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisputeWonWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r disputeWonWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type DisputeWonWebhookEventDataPayloadType string

const (
	DisputeWonWebhookEventDataPayloadTypeDispute DisputeWonWebhookEventDataPayloadType = "Dispute"
)

func (r DisputeWonWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case DisputeWonWebhookEventDataPayloadTypeDispute:
		return true
	}
	return false
}

// The event type
type DisputeWonWebhookEventType string

const (
	DisputeWonWebhookEventTypeDisputeWon DisputeWonWebhookEventType = "dispute.won"
)

func (r DisputeWonWebhookEventType) IsKnown() bool {
	switch r {
	case DisputeWonWebhookEventTypeDisputeWon:
		return true
	}
	return false
}

type LicenseKeyCreatedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data LicenseKeyCreatedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type LicenseKeyCreatedWebhookEventType `json:"type,required"`
	JSON licenseKeyCreatedWebhookEventJSON `json:"-"`
}

// licenseKeyCreatedWebhookEventJSON contains the JSON metadata for the struct
// [LicenseKeyCreatedWebhookEvent]
type licenseKeyCreatedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseKeyCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseKeyCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r LicenseKeyCreatedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r LicenseKeyCreatedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type LicenseKeyCreatedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType LicenseKeyCreatedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        licenseKeyCreatedWebhookEventDataJSON        `json:"-"`
	LicenseKey
}

// licenseKeyCreatedWebhookEventDataJSON contains the JSON metadata for the struct
// [LicenseKeyCreatedWebhookEventData]
type licenseKeyCreatedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LicenseKeyCreatedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r licenseKeyCreatedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type LicenseKeyCreatedWebhookEventDataPayloadType string

const (
	LicenseKeyCreatedWebhookEventDataPayloadTypeLicenseKey LicenseKeyCreatedWebhookEventDataPayloadType = "LicenseKey"
)

func (r LicenseKeyCreatedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case LicenseKeyCreatedWebhookEventDataPayloadTypeLicenseKey:
		return true
	}
	return false
}

// The event type
type LicenseKeyCreatedWebhookEventType string

const (
	LicenseKeyCreatedWebhookEventTypeLicenseKeyCreated LicenseKeyCreatedWebhookEventType = "license_key.created"
)

func (r LicenseKeyCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case LicenseKeyCreatedWebhookEventTypeLicenseKeyCreated:
		return true
	}
	return false
}

type PaymentCancelledWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentCancelledWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type PaymentCancelledWebhookEventType `json:"type,required"`
	JSON paymentCancelledWebhookEventJSON `json:"-"`
}

// paymentCancelledWebhookEventJSON contains the JSON metadata for the struct
// [PaymentCancelledWebhookEvent]
type paymentCancelledWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentCancelledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentCancelledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentCancelledWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r PaymentCancelledWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type PaymentCancelledWebhookEventData struct {
	// The type of payload in the data field
	PayloadType PaymentCancelledWebhookEventDataPayloadType `json:"payload_type"`
	JSON        paymentCancelledWebhookEventDataJSON        `json:"-"`
	Payment
}

// paymentCancelledWebhookEventDataJSON contains the JSON metadata for the struct
// [PaymentCancelledWebhookEventData]
type paymentCancelledWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentCancelledWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentCancelledWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type PaymentCancelledWebhookEventDataPayloadType string

const (
	PaymentCancelledWebhookEventDataPayloadTypePayment PaymentCancelledWebhookEventDataPayloadType = "Payment"
)

func (r PaymentCancelledWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case PaymentCancelledWebhookEventDataPayloadTypePayment:
		return true
	}
	return false
}

// The event type
type PaymentCancelledWebhookEventType string

const (
	PaymentCancelledWebhookEventTypePaymentCancelled PaymentCancelledWebhookEventType = "payment.cancelled"
)

func (r PaymentCancelledWebhookEventType) IsKnown() bool {
	switch r {
	case PaymentCancelledWebhookEventTypePaymentCancelled:
		return true
	}
	return false
}

type PaymentFailedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentFailedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type PaymentFailedWebhookEventType `json:"type,required"`
	JSON paymentFailedWebhookEventJSON `json:"-"`
}

// paymentFailedWebhookEventJSON contains the JSON metadata for the struct
// [PaymentFailedWebhookEvent]
type paymentFailedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentFailedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r PaymentFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type PaymentFailedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType PaymentFailedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        paymentFailedWebhookEventDataJSON        `json:"-"`
	Payment
}

// paymentFailedWebhookEventDataJSON contains the JSON metadata for the struct
// [PaymentFailedWebhookEventData]
type paymentFailedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentFailedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFailedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type PaymentFailedWebhookEventDataPayloadType string

const (
	PaymentFailedWebhookEventDataPayloadTypePayment PaymentFailedWebhookEventDataPayloadType = "Payment"
)

func (r PaymentFailedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case PaymentFailedWebhookEventDataPayloadTypePayment:
		return true
	}
	return false
}

// The event type
type PaymentFailedWebhookEventType string

const (
	PaymentFailedWebhookEventTypePaymentFailed PaymentFailedWebhookEventType = "payment.failed"
)

func (r PaymentFailedWebhookEventType) IsKnown() bool {
	switch r {
	case PaymentFailedWebhookEventTypePaymentFailed:
		return true
	}
	return false
}

type PaymentProcessingWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentProcessingWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type PaymentProcessingWebhookEventType `json:"type,required"`
	JSON paymentProcessingWebhookEventJSON `json:"-"`
}

// paymentProcessingWebhookEventJSON contains the JSON metadata for the struct
// [PaymentProcessingWebhookEvent]
type paymentProcessingWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentProcessingWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentProcessingWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentProcessingWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r PaymentProcessingWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type PaymentProcessingWebhookEventData struct {
	// The type of payload in the data field
	PayloadType PaymentProcessingWebhookEventDataPayloadType `json:"payload_type"`
	JSON        paymentProcessingWebhookEventDataJSON        `json:"-"`
	Payment
}

// paymentProcessingWebhookEventDataJSON contains the JSON metadata for the struct
// [PaymentProcessingWebhookEventData]
type paymentProcessingWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentProcessingWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentProcessingWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type PaymentProcessingWebhookEventDataPayloadType string

const (
	PaymentProcessingWebhookEventDataPayloadTypePayment PaymentProcessingWebhookEventDataPayloadType = "Payment"
)

func (r PaymentProcessingWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case PaymentProcessingWebhookEventDataPayloadTypePayment:
		return true
	}
	return false
}

// The event type
type PaymentProcessingWebhookEventType string

const (
	PaymentProcessingWebhookEventTypePaymentProcessing PaymentProcessingWebhookEventType = "payment.processing"
)

func (r PaymentProcessingWebhookEventType) IsKnown() bool {
	switch r {
	case PaymentProcessingWebhookEventTypePaymentProcessing:
		return true
	}
	return false
}

type PaymentSucceededWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentSucceededWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type PaymentSucceededWebhookEventType `json:"type,required"`
	JSON paymentSucceededWebhookEventJSON `json:"-"`
}

// paymentSucceededWebhookEventJSON contains the JSON metadata for the struct
// [PaymentSucceededWebhookEvent]
type paymentSucceededWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentSucceededWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r PaymentSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type PaymentSucceededWebhookEventData struct {
	// The type of payload in the data field
	PayloadType PaymentSucceededWebhookEventDataPayloadType `json:"payload_type"`
	JSON        paymentSucceededWebhookEventDataJSON        `json:"-"`
	Payment
}

// paymentSucceededWebhookEventDataJSON contains the JSON metadata for the struct
// [PaymentSucceededWebhookEventData]
type paymentSucceededWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentSucceededWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentSucceededWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type PaymentSucceededWebhookEventDataPayloadType string

const (
	PaymentSucceededWebhookEventDataPayloadTypePayment PaymentSucceededWebhookEventDataPayloadType = "Payment"
)

func (r PaymentSucceededWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case PaymentSucceededWebhookEventDataPayloadTypePayment:
		return true
	}
	return false
}

// The event type
type PaymentSucceededWebhookEventType string

const (
	PaymentSucceededWebhookEventTypePaymentSucceeded PaymentSucceededWebhookEventType = "payment.succeeded"
)

func (r PaymentSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case PaymentSucceededWebhookEventTypePaymentSucceeded:
		return true
	}
	return false
}

type RefundFailedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data RefundFailedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type RefundFailedWebhookEventType `json:"type,required"`
	JSON refundFailedWebhookEventJSON `json:"-"`
}

// refundFailedWebhookEventJSON contains the JSON metadata for the struct
// [RefundFailedWebhookEvent]
type refundFailedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefundFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refundFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RefundFailedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r RefundFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type RefundFailedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType RefundFailedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        refundFailedWebhookEventDataJSON        `json:"-"`
	Refund
}

// refundFailedWebhookEventDataJSON contains the JSON metadata for the struct
// [RefundFailedWebhookEventData]
type refundFailedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefundFailedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refundFailedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type RefundFailedWebhookEventDataPayloadType string

const (
	RefundFailedWebhookEventDataPayloadTypeRefund RefundFailedWebhookEventDataPayloadType = "Refund"
)

func (r RefundFailedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case RefundFailedWebhookEventDataPayloadTypeRefund:
		return true
	}
	return false
}

// The event type
type RefundFailedWebhookEventType string

const (
	RefundFailedWebhookEventTypeRefundFailed RefundFailedWebhookEventType = "refund.failed"
)

func (r RefundFailedWebhookEventType) IsKnown() bool {
	switch r {
	case RefundFailedWebhookEventTypeRefundFailed:
		return true
	}
	return false
}

type RefundSucceededWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data RefundSucceededWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type RefundSucceededWebhookEventType `json:"type,required"`
	JSON refundSucceededWebhookEventJSON `json:"-"`
}

// refundSucceededWebhookEventJSON contains the JSON metadata for the struct
// [RefundSucceededWebhookEvent]
type refundSucceededWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefundSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refundSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r RefundSucceededWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r RefundSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type RefundSucceededWebhookEventData struct {
	// The type of payload in the data field
	PayloadType RefundSucceededWebhookEventDataPayloadType `json:"payload_type"`
	JSON        refundSucceededWebhookEventDataJSON        `json:"-"`
	Refund
}

// refundSucceededWebhookEventDataJSON contains the JSON metadata for the struct
// [RefundSucceededWebhookEventData]
type refundSucceededWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RefundSucceededWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r refundSucceededWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type RefundSucceededWebhookEventDataPayloadType string

const (
	RefundSucceededWebhookEventDataPayloadTypeRefund RefundSucceededWebhookEventDataPayloadType = "Refund"
)

func (r RefundSucceededWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case RefundSucceededWebhookEventDataPayloadTypeRefund:
		return true
	}
	return false
}

// The event type
type RefundSucceededWebhookEventType string

const (
	RefundSucceededWebhookEventTypeRefundSucceeded RefundSucceededWebhookEventType = "refund.succeeded"
)

func (r RefundSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case RefundSucceededWebhookEventTypeRefundSucceeded:
		return true
	}
	return false
}

type SubscriptionActiveWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionActiveWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type SubscriptionActiveWebhookEventType `json:"type,required"`
	JSON subscriptionActiveWebhookEventJSON `json:"-"`
}

// subscriptionActiveWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionActiveWebhookEvent]
type subscriptionActiveWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionActiveWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionActiveWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionActiveWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r SubscriptionActiveWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type SubscriptionActiveWebhookEventData struct {
	// The type of payload in the data field
	PayloadType SubscriptionActiveWebhookEventDataPayloadType `json:"payload_type"`
	JSON        subscriptionActiveWebhookEventDataJSON        `json:"-"`
	Subscription
}

// subscriptionActiveWebhookEventDataJSON contains the JSON metadata for the struct
// [SubscriptionActiveWebhookEventData]
type subscriptionActiveWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionActiveWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionActiveWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type SubscriptionActiveWebhookEventDataPayloadType string

const (
	SubscriptionActiveWebhookEventDataPayloadTypeSubscription SubscriptionActiveWebhookEventDataPayloadType = "Subscription"
)

func (r SubscriptionActiveWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case SubscriptionActiveWebhookEventDataPayloadTypeSubscription:
		return true
	}
	return false
}

// The event type
type SubscriptionActiveWebhookEventType string

const (
	SubscriptionActiveWebhookEventTypeSubscriptionActive SubscriptionActiveWebhookEventType = "subscription.active"
)

func (r SubscriptionActiveWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionActiveWebhookEventTypeSubscriptionActive:
		return true
	}
	return false
}

type SubscriptionCancelledWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionCancelledWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type SubscriptionCancelledWebhookEventType `json:"type,required"`
	JSON subscriptionCancelledWebhookEventJSON `json:"-"`
}

// subscriptionCancelledWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionCancelledWebhookEvent]
type subscriptionCancelledWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCancelledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancelledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionCancelledWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r SubscriptionCancelledWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type SubscriptionCancelledWebhookEventData struct {
	// The type of payload in the data field
	PayloadType SubscriptionCancelledWebhookEventDataPayloadType `json:"payload_type"`
	JSON        subscriptionCancelledWebhookEventDataJSON        `json:"-"`
	Subscription
}

// subscriptionCancelledWebhookEventDataJSON contains the JSON metadata for the
// struct [SubscriptionCancelledWebhookEventData]
type subscriptionCancelledWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCancelledWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancelledWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type SubscriptionCancelledWebhookEventDataPayloadType string

const (
	SubscriptionCancelledWebhookEventDataPayloadTypeSubscription SubscriptionCancelledWebhookEventDataPayloadType = "Subscription"
)

func (r SubscriptionCancelledWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case SubscriptionCancelledWebhookEventDataPayloadTypeSubscription:
		return true
	}
	return false
}

// The event type
type SubscriptionCancelledWebhookEventType string

const (
	SubscriptionCancelledWebhookEventTypeSubscriptionCancelled SubscriptionCancelledWebhookEventType = "subscription.cancelled"
)

func (r SubscriptionCancelledWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionCancelledWebhookEventTypeSubscriptionCancelled:
		return true
	}
	return false
}

type SubscriptionExpiredWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionExpiredWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type SubscriptionExpiredWebhookEventType `json:"type,required"`
	JSON subscriptionExpiredWebhookEventJSON `json:"-"`
}

// subscriptionExpiredWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionExpiredWebhookEvent]
type subscriptionExpiredWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionExpiredWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionExpiredWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionExpiredWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r SubscriptionExpiredWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type SubscriptionExpiredWebhookEventData struct {
	// The type of payload in the data field
	PayloadType SubscriptionExpiredWebhookEventDataPayloadType `json:"payload_type"`
	JSON        subscriptionExpiredWebhookEventDataJSON        `json:"-"`
	Subscription
}

// subscriptionExpiredWebhookEventDataJSON contains the JSON metadata for the
// struct [SubscriptionExpiredWebhookEventData]
type subscriptionExpiredWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionExpiredWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionExpiredWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type SubscriptionExpiredWebhookEventDataPayloadType string

const (
	SubscriptionExpiredWebhookEventDataPayloadTypeSubscription SubscriptionExpiredWebhookEventDataPayloadType = "Subscription"
)

func (r SubscriptionExpiredWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case SubscriptionExpiredWebhookEventDataPayloadTypeSubscription:
		return true
	}
	return false
}

// The event type
type SubscriptionExpiredWebhookEventType string

const (
	SubscriptionExpiredWebhookEventTypeSubscriptionExpired SubscriptionExpiredWebhookEventType = "subscription.expired"
)

func (r SubscriptionExpiredWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionExpiredWebhookEventTypeSubscriptionExpired:
		return true
	}
	return false
}

type SubscriptionFailedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionFailedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type SubscriptionFailedWebhookEventType `json:"type,required"`
	JSON subscriptionFailedWebhookEventJSON `json:"-"`
}

// subscriptionFailedWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionFailedWebhookEvent]
type subscriptionFailedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionFailedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r SubscriptionFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type SubscriptionFailedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType SubscriptionFailedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        subscriptionFailedWebhookEventDataJSON        `json:"-"`
	Subscription
}

// subscriptionFailedWebhookEventDataJSON contains the JSON metadata for the struct
// [SubscriptionFailedWebhookEventData]
type subscriptionFailedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionFailedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFailedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type SubscriptionFailedWebhookEventDataPayloadType string

const (
	SubscriptionFailedWebhookEventDataPayloadTypeSubscription SubscriptionFailedWebhookEventDataPayloadType = "Subscription"
)

func (r SubscriptionFailedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case SubscriptionFailedWebhookEventDataPayloadTypeSubscription:
		return true
	}
	return false
}

// The event type
type SubscriptionFailedWebhookEventType string

const (
	SubscriptionFailedWebhookEventTypeSubscriptionFailed SubscriptionFailedWebhookEventType = "subscription.failed"
)

func (r SubscriptionFailedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionFailedWebhookEventTypeSubscriptionFailed:
		return true
	}
	return false
}

type SubscriptionOnHoldWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionOnHoldWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type SubscriptionOnHoldWebhookEventType `json:"type,required"`
	JSON subscriptionOnHoldWebhookEventJSON `json:"-"`
}

// subscriptionOnHoldWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionOnHoldWebhookEvent]
type subscriptionOnHoldWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionOnHoldWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionOnHoldWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionOnHoldWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r SubscriptionOnHoldWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type SubscriptionOnHoldWebhookEventData struct {
	// The type of payload in the data field
	PayloadType SubscriptionOnHoldWebhookEventDataPayloadType `json:"payload_type"`
	JSON        subscriptionOnHoldWebhookEventDataJSON        `json:"-"`
	Subscription
}

// subscriptionOnHoldWebhookEventDataJSON contains the JSON metadata for the struct
// [SubscriptionOnHoldWebhookEventData]
type subscriptionOnHoldWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionOnHoldWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionOnHoldWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type SubscriptionOnHoldWebhookEventDataPayloadType string

const (
	SubscriptionOnHoldWebhookEventDataPayloadTypeSubscription SubscriptionOnHoldWebhookEventDataPayloadType = "Subscription"
)

func (r SubscriptionOnHoldWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case SubscriptionOnHoldWebhookEventDataPayloadTypeSubscription:
		return true
	}
	return false
}

// The event type
type SubscriptionOnHoldWebhookEventType string

const (
	SubscriptionOnHoldWebhookEventTypeSubscriptionOnHold SubscriptionOnHoldWebhookEventType = "subscription.on_hold"
)

func (r SubscriptionOnHoldWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionOnHoldWebhookEventTypeSubscriptionOnHold:
		return true
	}
	return false
}

type SubscriptionPlanChangedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionPlanChangedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type SubscriptionPlanChangedWebhookEventType `json:"type,required"`
	JSON subscriptionPlanChangedWebhookEventJSON `json:"-"`
}

// subscriptionPlanChangedWebhookEventJSON contains the JSON metadata for the
// struct [SubscriptionPlanChangedWebhookEvent]
type subscriptionPlanChangedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionPlanChangedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanChangedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPlanChangedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r SubscriptionPlanChangedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type SubscriptionPlanChangedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType SubscriptionPlanChangedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        subscriptionPlanChangedWebhookEventDataJSON        `json:"-"`
	Subscription
}

// subscriptionPlanChangedWebhookEventDataJSON contains the JSON metadata for the
// struct [SubscriptionPlanChangedWebhookEventData]
type subscriptionPlanChangedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionPlanChangedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanChangedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type SubscriptionPlanChangedWebhookEventDataPayloadType string

const (
	SubscriptionPlanChangedWebhookEventDataPayloadTypeSubscription SubscriptionPlanChangedWebhookEventDataPayloadType = "Subscription"
)

func (r SubscriptionPlanChangedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case SubscriptionPlanChangedWebhookEventDataPayloadTypeSubscription:
		return true
	}
	return false
}

// The event type
type SubscriptionPlanChangedWebhookEventType string

const (
	SubscriptionPlanChangedWebhookEventTypeSubscriptionPlanChanged SubscriptionPlanChangedWebhookEventType = "subscription.plan_changed"
)

func (r SubscriptionPlanChangedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionPlanChangedWebhookEventTypeSubscriptionPlanChanged:
		return true
	}
	return false
}

type SubscriptionRenewedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionRenewedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type SubscriptionRenewedWebhookEventType `json:"type,required"`
	JSON subscriptionRenewedWebhookEventJSON `json:"-"`
}

// subscriptionRenewedWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionRenewedWebhookEvent]
type subscriptionRenewedWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionRenewedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionRenewedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionRenewedWebhookEvent) implementsUnsafeUnwrapWebhookEvent() {}

func (r SubscriptionRenewedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Event-specific data
type SubscriptionRenewedWebhookEventData struct {
	// The type of payload in the data field
	PayloadType SubscriptionRenewedWebhookEventDataPayloadType `json:"payload_type"`
	JSON        subscriptionRenewedWebhookEventDataJSON        `json:"-"`
	Subscription
}

// subscriptionRenewedWebhookEventDataJSON contains the JSON metadata for the
// struct [SubscriptionRenewedWebhookEventData]
type subscriptionRenewedWebhookEventDataJSON struct {
	PayloadType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionRenewedWebhookEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionRenewedWebhookEventDataJSON) RawJSON() string {
	return r.raw
}

// The type of payload in the data field
type SubscriptionRenewedWebhookEventDataPayloadType string

const (
	SubscriptionRenewedWebhookEventDataPayloadTypeSubscription SubscriptionRenewedWebhookEventDataPayloadType = "Subscription"
)

func (r SubscriptionRenewedWebhookEventDataPayloadType) IsKnown() bool {
	switch r {
	case SubscriptionRenewedWebhookEventDataPayloadTypeSubscription:
		return true
	}
	return false
}

// The event type
type SubscriptionRenewedWebhookEventType string

const (
	SubscriptionRenewedWebhookEventTypeSubscriptionRenewed SubscriptionRenewedWebhookEventType = "subscription.renewed"
)

func (r SubscriptionRenewedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionRenewedWebhookEventTypeSubscriptionRenewed:
		return true
	}
	return false
}

type UnsafeUnwrapWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// This field can have the runtime type of [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData],
	// [LicenseKeyCreatedWebhookEventData], [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData], [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData], [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData].
	Data interface{} `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type  UnsafeUnwrapWebhookEventType `json:"type,required"`
	JSON  unsafeUnwrapWebhookEventJSON `json:"-"`
	union UnsafeUnwrapWebhookEventUnion
}

// unsafeUnwrapWebhookEventJSON contains the JSON metadata for the struct
// [UnsafeUnwrapWebhookEvent]
type unsafeUnwrapWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r unsafeUnwrapWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *UnsafeUnwrapWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = UnsafeUnwrapWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UnsafeUnwrapWebhookEventUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [DisputeAcceptedWebhookEvent],
// [DisputeCancelledWebhookEvent], [DisputeChallengedWebhookEvent],
// [DisputeExpiredWebhookEvent], [DisputeLostWebhookEvent],
// [DisputeOpenedWebhookEvent], [DisputeWonWebhookEvent],
// [LicenseKeyCreatedWebhookEvent], [PaymentCancelledWebhookEvent],
// [PaymentFailedWebhookEvent], [PaymentProcessingWebhookEvent],
// [PaymentSucceededWebhookEvent], [RefundFailedWebhookEvent],
// [RefundSucceededWebhookEvent], [SubscriptionActiveWebhookEvent],
// [SubscriptionCancelledWebhookEvent], [SubscriptionExpiredWebhookEvent],
// [SubscriptionFailedWebhookEvent], [SubscriptionOnHoldWebhookEvent],
// [SubscriptionPlanChangedWebhookEvent], [SubscriptionRenewedWebhookEvent].
func (r UnsafeUnwrapWebhookEvent) AsUnion() UnsafeUnwrapWebhookEventUnion {
	return r.union
}

// Union satisfied by [DisputeAcceptedWebhookEvent],
// [DisputeCancelledWebhookEvent], [DisputeChallengedWebhookEvent],
// [DisputeExpiredWebhookEvent], [DisputeLostWebhookEvent],
// [DisputeOpenedWebhookEvent], [DisputeWonWebhookEvent],
// [LicenseKeyCreatedWebhookEvent], [PaymentCancelledWebhookEvent],
// [PaymentFailedWebhookEvent], [PaymentProcessingWebhookEvent],
// [PaymentSucceededWebhookEvent], [RefundFailedWebhookEvent],
// [RefundSucceededWebhookEvent], [SubscriptionActiveWebhookEvent],
// [SubscriptionCancelledWebhookEvent], [SubscriptionExpiredWebhookEvent],
// [SubscriptionFailedWebhookEvent], [SubscriptionOnHoldWebhookEvent],
// [SubscriptionPlanChangedWebhookEvent] or [SubscriptionRenewedWebhookEvent].
type UnsafeUnwrapWebhookEventUnion interface {
	implementsUnsafeUnwrapWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnsafeUnwrapWebhookEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeAcceptedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeCancelledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeChallengedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeExpiredWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeLostWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeOpenedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeWonWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LicenseKeyCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentCancelledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentProcessingWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RefundFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RefundSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionActiveWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionCancelledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionExpiredWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionOnHoldWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPlanChangedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionRenewedWebhookEvent{}),
		},
	)
}

// The event type
type UnsafeUnwrapWebhookEventType string

const (
	UnsafeUnwrapWebhookEventTypeDisputeAccepted         UnsafeUnwrapWebhookEventType = "dispute.accepted"
	UnsafeUnwrapWebhookEventTypeDisputeCancelled        UnsafeUnwrapWebhookEventType = "dispute.cancelled"
	UnsafeUnwrapWebhookEventTypeDisputeChallenged       UnsafeUnwrapWebhookEventType = "dispute.challenged"
	UnsafeUnwrapWebhookEventTypeDisputeExpired          UnsafeUnwrapWebhookEventType = "dispute.expired"
	UnsafeUnwrapWebhookEventTypeDisputeLost             UnsafeUnwrapWebhookEventType = "dispute.lost"
	UnsafeUnwrapWebhookEventTypeDisputeOpened           UnsafeUnwrapWebhookEventType = "dispute.opened"
	UnsafeUnwrapWebhookEventTypeDisputeWon              UnsafeUnwrapWebhookEventType = "dispute.won"
	UnsafeUnwrapWebhookEventTypeLicenseKeyCreated       UnsafeUnwrapWebhookEventType = "license_key.created"
	UnsafeUnwrapWebhookEventTypePaymentCancelled        UnsafeUnwrapWebhookEventType = "payment.cancelled"
	UnsafeUnwrapWebhookEventTypePaymentFailed           UnsafeUnwrapWebhookEventType = "payment.failed"
	UnsafeUnwrapWebhookEventTypePaymentProcessing       UnsafeUnwrapWebhookEventType = "payment.processing"
	UnsafeUnwrapWebhookEventTypePaymentSucceeded        UnsafeUnwrapWebhookEventType = "payment.succeeded"
	UnsafeUnwrapWebhookEventTypeRefundFailed            UnsafeUnwrapWebhookEventType = "refund.failed"
	UnsafeUnwrapWebhookEventTypeRefundSucceeded         UnsafeUnwrapWebhookEventType = "refund.succeeded"
	UnsafeUnwrapWebhookEventTypeSubscriptionActive      UnsafeUnwrapWebhookEventType = "subscription.active"
	UnsafeUnwrapWebhookEventTypeSubscriptionCancelled   UnsafeUnwrapWebhookEventType = "subscription.cancelled"
	UnsafeUnwrapWebhookEventTypeSubscriptionExpired     UnsafeUnwrapWebhookEventType = "subscription.expired"
	UnsafeUnwrapWebhookEventTypeSubscriptionFailed      UnsafeUnwrapWebhookEventType = "subscription.failed"
	UnsafeUnwrapWebhookEventTypeSubscriptionOnHold      UnsafeUnwrapWebhookEventType = "subscription.on_hold"
	UnsafeUnwrapWebhookEventTypeSubscriptionPlanChanged UnsafeUnwrapWebhookEventType = "subscription.plan_changed"
	UnsafeUnwrapWebhookEventTypeSubscriptionRenewed     UnsafeUnwrapWebhookEventType = "subscription.renewed"
)

func (r UnsafeUnwrapWebhookEventType) IsKnown() bool {
	switch r {
	case UnsafeUnwrapWebhookEventTypeDisputeAccepted, UnsafeUnwrapWebhookEventTypeDisputeCancelled, UnsafeUnwrapWebhookEventTypeDisputeChallenged, UnsafeUnwrapWebhookEventTypeDisputeExpired, UnsafeUnwrapWebhookEventTypeDisputeLost, UnsafeUnwrapWebhookEventTypeDisputeOpened, UnsafeUnwrapWebhookEventTypeDisputeWon, UnsafeUnwrapWebhookEventTypeLicenseKeyCreated, UnsafeUnwrapWebhookEventTypePaymentCancelled, UnsafeUnwrapWebhookEventTypePaymentFailed, UnsafeUnwrapWebhookEventTypePaymentProcessing, UnsafeUnwrapWebhookEventTypePaymentSucceeded, UnsafeUnwrapWebhookEventTypeRefundFailed, UnsafeUnwrapWebhookEventTypeRefundSucceeded, UnsafeUnwrapWebhookEventTypeSubscriptionActive, UnsafeUnwrapWebhookEventTypeSubscriptionCancelled, UnsafeUnwrapWebhookEventTypeSubscriptionExpired, UnsafeUnwrapWebhookEventTypeSubscriptionFailed, UnsafeUnwrapWebhookEventTypeSubscriptionOnHold, UnsafeUnwrapWebhookEventTypeSubscriptionPlanChanged, UnsafeUnwrapWebhookEventTypeSubscriptionRenewed:
		return true
	}
	return false
}

type UnwrapWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// This field can have the runtime type of [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData],
	// [LicenseKeyCreatedWebhookEventData], [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData], [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData], [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData].
	Data interface{} `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	Type  UnwrapWebhookEventType `json:"type,required"`
	JSON  unwrapWebhookEventJSON `json:"-"`
	union UnwrapWebhookEventUnion
}

// unwrapWebhookEventJSON contains the JSON metadata for the struct
// [UnwrapWebhookEvent]
type unwrapWebhookEventJSON struct {
	BusinessID  apijson.Field
	Data        apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r unwrapWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *UnwrapWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = UnwrapWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UnwrapWebhookEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DisputeAcceptedWebhookEvent],
// [DisputeCancelledWebhookEvent], [DisputeChallengedWebhookEvent],
// [DisputeExpiredWebhookEvent], [DisputeLostWebhookEvent],
// [DisputeOpenedWebhookEvent], [DisputeWonWebhookEvent],
// [LicenseKeyCreatedWebhookEvent], [PaymentCancelledWebhookEvent],
// [PaymentFailedWebhookEvent], [PaymentProcessingWebhookEvent],
// [PaymentSucceededWebhookEvent], [RefundFailedWebhookEvent],
// [RefundSucceededWebhookEvent], [SubscriptionActiveWebhookEvent],
// [SubscriptionCancelledWebhookEvent], [SubscriptionExpiredWebhookEvent],
// [SubscriptionFailedWebhookEvent], [SubscriptionOnHoldWebhookEvent],
// [SubscriptionPlanChangedWebhookEvent], [SubscriptionRenewedWebhookEvent].
func (r UnwrapWebhookEvent) AsUnion() UnwrapWebhookEventUnion {
	return r.union
}

// Union satisfied by [DisputeAcceptedWebhookEvent],
// [DisputeCancelledWebhookEvent], [DisputeChallengedWebhookEvent],
// [DisputeExpiredWebhookEvent], [DisputeLostWebhookEvent],
// [DisputeOpenedWebhookEvent], [DisputeWonWebhookEvent],
// [LicenseKeyCreatedWebhookEvent], [PaymentCancelledWebhookEvent],
// [PaymentFailedWebhookEvent], [PaymentProcessingWebhookEvent],
// [PaymentSucceededWebhookEvent], [RefundFailedWebhookEvent],
// [RefundSucceededWebhookEvent], [SubscriptionActiveWebhookEvent],
// [SubscriptionCancelledWebhookEvent], [SubscriptionExpiredWebhookEvent],
// [SubscriptionFailedWebhookEvent], [SubscriptionOnHoldWebhookEvent],
// [SubscriptionPlanChangedWebhookEvent] or [SubscriptionRenewedWebhookEvent].
type UnwrapWebhookEventUnion interface {
	implementsUnwrapWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnwrapWebhookEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeAcceptedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeCancelledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeChallengedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeExpiredWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeLostWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeOpenedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DisputeWonWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LicenseKeyCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentCancelledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentProcessingWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RefundFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RefundSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionActiveWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionCancelledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionExpiredWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionOnHoldWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPlanChangedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionRenewedWebhookEvent{}),
		},
	)
}

// The event type
type UnwrapWebhookEventType string

const (
	UnwrapWebhookEventTypeDisputeAccepted         UnwrapWebhookEventType = "dispute.accepted"
	UnwrapWebhookEventTypeDisputeCancelled        UnwrapWebhookEventType = "dispute.cancelled"
	UnwrapWebhookEventTypeDisputeChallenged       UnwrapWebhookEventType = "dispute.challenged"
	UnwrapWebhookEventTypeDisputeExpired          UnwrapWebhookEventType = "dispute.expired"
	UnwrapWebhookEventTypeDisputeLost             UnwrapWebhookEventType = "dispute.lost"
	UnwrapWebhookEventTypeDisputeOpened           UnwrapWebhookEventType = "dispute.opened"
	UnwrapWebhookEventTypeDisputeWon              UnwrapWebhookEventType = "dispute.won"
	UnwrapWebhookEventTypeLicenseKeyCreated       UnwrapWebhookEventType = "license_key.created"
	UnwrapWebhookEventTypePaymentCancelled        UnwrapWebhookEventType = "payment.cancelled"
	UnwrapWebhookEventTypePaymentFailed           UnwrapWebhookEventType = "payment.failed"
	UnwrapWebhookEventTypePaymentProcessing       UnwrapWebhookEventType = "payment.processing"
	UnwrapWebhookEventTypePaymentSucceeded        UnwrapWebhookEventType = "payment.succeeded"
	UnwrapWebhookEventTypeRefundFailed            UnwrapWebhookEventType = "refund.failed"
	UnwrapWebhookEventTypeRefundSucceeded         UnwrapWebhookEventType = "refund.succeeded"
	UnwrapWebhookEventTypeSubscriptionActive      UnwrapWebhookEventType = "subscription.active"
	UnwrapWebhookEventTypeSubscriptionCancelled   UnwrapWebhookEventType = "subscription.cancelled"
	UnwrapWebhookEventTypeSubscriptionExpired     UnwrapWebhookEventType = "subscription.expired"
	UnwrapWebhookEventTypeSubscriptionFailed      UnwrapWebhookEventType = "subscription.failed"
	UnwrapWebhookEventTypeSubscriptionOnHold      UnwrapWebhookEventType = "subscription.on_hold"
	UnwrapWebhookEventTypeSubscriptionPlanChanged UnwrapWebhookEventType = "subscription.plan_changed"
	UnwrapWebhookEventTypeSubscriptionRenewed     UnwrapWebhookEventType = "subscription.renewed"
)

func (r UnwrapWebhookEventType) IsKnown() bool {
	switch r {
	case UnwrapWebhookEventTypeDisputeAccepted, UnwrapWebhookEventTypeDisputeCancelled, UnwrapWebhookEventTypeDisputeChallenged, UnwrapWebhookEventTypeDisputeExpired, UnwrapWebhookEventTypeDisputeLost, UnwrapWebhookEventTypeDisputeOpened, UnwrapWebhookEventTypeDisputeWon, UnwrapWebhookEventTypeLicenseKeyCreated, UnwrapWebhookEventTypePaymentCancelled, UnwrapWebhookEventTypePaymentFailed, UnwrapWebhookEventTypePaymentProcessing, UnwrapWebhookEventTypePaymentSucceeded, UnwrapWebhookEventTypeRefundFailed, UnwrapWebhookEventTypeRefundSucceeded, UnwrapWebhookEventTypeSubscriptionActive, UnwrapWebhookEventTypeSubscriptionCancelled, UnwrapWebhookEventTypeSubscriptionExpired, UnwrapWebhookEventTypeSubscriptionFailed, UnwrapWebhookEventTypeSubscriptionOnHold, UnwrapWebhookEventTypeSubscriptionPlanChanged, UnwrapWebhookEventTypeSubscriptionRenewed:
		return true
	}
	return false
}

type WebhookNewParams struct {
	// Url of the webhook
	URL         param.Field[string] `json:"url,required"`
	Description param.Field[string] `json:"description"`
	// Create the webhook in a disabled state.
	//
	// Default is false
	Disabled param.Field[bool] `json:"disabled"`
	// Filter events to the webhook.
	//
	// Webhook event will only be sent for events in the list.
	FilterTypes param.Field[[]WebhookEventType] `json:"filter_types"`
	// Custom headers to be passed
	Headers param.Field[map[string]string] `json:"headers"`
	// The request's idempotency key
	IdempotencyKey param.Field[string] `json:"idempotency_key"`
	// Metadata to be passed to the webhook Defaut is {}
	Metadata  param.Field[map[string]string] `json:"metadata"`
	RateLimit param.Field[int64]             `json:"rate_limit"`
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookUpdateParams struct {
	// Description of the webhook
	Description param.Field[string] `json:"description"`
	// To Disable the endpoint, set it to true.
	Disabled param.Field[bool] `json:"disabled"`
	// Filter events to the endpoint.
	//
	// Webhook event will only be sent for events in the list.
	FilterTypes param.Field[[]WebhookEventType] `json:"filter_types"`
	// Metadata
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Rate limit
	RateLimit param.Field[int64] `json:"rate_limit"`
	// Url endpoint
	URL param.Field[string] `json:"url"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookListParams struct {
	// The iterator returned from a prior invocation
	Iterator param.Field[string] `query:"iterator"`
	// Limit the number of returned items
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
