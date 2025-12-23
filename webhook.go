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
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

// WebhookService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
	Headers WebhookHeaderService
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
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

func (r *WebhookService) UnsafeUnwrap(payload []byte, opts ...option.RequestOption) (*UnsafeUnwrapWebhookEventUnion, error) {
	res := &UnsafeUnwrapWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
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
	res := &UnwrapWebhookEventUnion{}
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
	RateLimit int64 `json:"rate_limit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Metadata    respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		Disabled    respjson.Field
		FilterTypes respjson.Field
		RateLimit   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDetails) RawJSON() string { return r.JSON.raw }
func (r *WebhookDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookGetSecretResponse struct {
	Secret string `json:"secret,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookGetSecretResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookGetSecretResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DisputeAcceptedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeAcceptedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "dispute.accepted".
	Type DisputeAcceptedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeAcceptedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DisputeAcceptedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type DisputeAcceptedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Dispute".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Dispute
}

// Returns the unmodified JSON received from the API
func (r DisputeAcceptedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DisputeAcceptedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type DisputeAcceptedWebhookEventType string

const (
	DisputeAcceptedWebhookEventTypeDisputeAccepted DisputeAcceptedWebhookEventType = "dispute.accepted"
)

type DisputeCancelledWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeCancelledWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "dispute.cancelled".
	Type DisputeCancelledWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeCancelledWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DisputeCancelledWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type DisputeCancelledWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Dispute".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Dispute
}

// Returns the unmodified JSON received from the API
func (r DisputeCancelledWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DisputeCancelledWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type DisputeCancelledWebhookEventType string

const (
	DisputeCancelledWebhookEventTypeDisputeCancelled DisputeCancelledWebhookEventType = "dispute.cancelled"
)

type DisputeChallengedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeChallengedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "dispute.challenged".
	Type DisputeChallengedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeChallengedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DisputeChallengedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type DisputeChallengedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Dispute".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Dispute
}

// Returns the unmodified JSON received from the API
func (r DisputeChallengedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DisputeChallengedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type DisputeChallengedWebhookEventType string

const (
	DisputeChallengedWebhookEventTypeDisputeChallenged DisputeChallengedWebhookEventType = "dispute.challenged"
)

type DisputeExpiredWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeExpiredWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "dispute.expired".
	Type DisputeExpiredWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeExpiredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DisputeExpiredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type DisputeExpiredWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Dispute".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Dispute
}

// Returns the unmodified JSON received from the API
func (r DisputeExpiredWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DisputeExpiredWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type DisputeExpiredWebhookEventType string

const (
	DisputeExpiredWebhookEventTypeDisputeExpired DisputeExpiredWebhookEventType = "dispute.expired"
)

type DisputeLostWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeLostWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "dispute.lost".
	Type DisputeLostWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeLostWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DisputeLostWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type DisputeLostWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Dispute".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Dispute
}

// Returns the unmodified JSON received from the API
func (r DisputeLostWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DisputeLostWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type DisputeLostWebhookEventType string

const (
	DisputeLostWebhookEventTypeDisputeLost DisputeLostWebhookEventType = "dispute.lost"
)

type DisputeOpenedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeOpenedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "dispute.opened".
	Type DisputeOpenedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeOpenedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DisputeOpenedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type DisputeOpenedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Dispute".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Dispute
}

// Returns the unmodified JSON received from the API
func (r DisputeOpenedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DisputeOpenedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type DisputeOpenedWebhookEventType string

const (
	DisputeOpenedWebhookEventTypeDisputeOpened DisputeOpenedWebhookEventType = "dispute.opened"
)

type DisputeWonWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data DisputeWonWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "dispute.won".
	Type DisputeWonWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DisputeWonWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DisputeWonWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type DisputeWonWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Dispute".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Dispute
}

// Returns the unmodified JSON received from the API
func (r DisputeWonWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DisputeWonWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type DisputeWonWebhookEventType string

const (
	DisputeWonWebhookEventTypeDisputeWon DisputeWonWebhookEventType = "dispute.won"
)

type LicenseKeyCreatedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data LicenseKeyCreatedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "license_key.created".
	Type LicenseKeyCreatedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LicenseKeyCreatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *LicenseKeyCreatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type LicenseKeyCreatedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "LicenseKey".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	LicenseKey
}

// Returns the unmodified JSON received from the API
func (r LicenseKeyCreatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *LicenseKeyCreatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type LicenseKeyCreatedWebhookEventType string

const (
	LicenseKeyCreatedWebhookEventTypeLicenseKeyCreated LicenseKeyCreatedWebhookEventType = "license_key.created"
)

type PaymentCancelledWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentCancelledWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "payment.cancelled".
	Type PaymentCancelledWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentCancelledWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentCancelledWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type PaymentCancelledWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Payment".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Payment
}

// Returns the unmodified JSON received from the API
func (r PaymentCancelledWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentCancelledWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type PaymentCancelledWebhookEventType string

const (
	PaymentCancelledWebhookEventTypePaymentCancelled PaymentCancelledWebhookEventType = "payment.cancelled"
)

type PaymentFailedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentFailedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "payment.failed".
	Type PaymentFailedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type PaymentFailedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Payment".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Payment
}

// Returns the unmodified JSON received from the API
func (r PaymentFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type PaymentFailedWebhookEventType string

const (
	PaymentFailedWebhookEventTypePaymentFailed PaymentFailedWebhookEventType = "payment.failed"
)

type PaymentProcessingWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentProcessingWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "payment.processing".
	Type PaymentProcessingWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentProcessingWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentProcessingWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type PaymentProcessingWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Payment".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Payment
}

// Returns the unmodified JSON received from the API
func (r PaymentProcessingWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentProcessingWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type PaymentProcessingWebhookEventType string

const (
	PaymentProcessingWebhookEventTypePaymentProcessing PaymentProcessingWebhookEventType = "payment.processing"
)

type PaymentSucceededWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data PaymentSucceededWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "payment.succeeded".
	Type PaymentSucceededWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentSucceededWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *PaymentSucceededWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type PaymentSucceededWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Payment".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Payment
}

// Returns the unmodified JSON received from the API
func (r PaymentSucceededWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *PaymentSucceededWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type PaymentSucceededWebhookEventType string

const (
	PaymentSucceededWebhookEventTypePaymentSucceeded PaymentSucceededWebhookEventType = "payment.succeeded"
)

type RefundFailedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data RefundFailedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "refund.failed".
	Type RefundFailedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RefundFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *RefundFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type RefundFailedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Refund".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Refund
}

// Returns the unmodified JSON received from the API
func (r RefundFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *RefundFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type RefundFailedWebhookEventType string

const (
	RefundFailedWebhookEventTypeRefundFailed RefundFailedWebhookEventType = "refund.failed"
)

type RefundSucceededWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data RefundSucceededWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "refund.succeeded".
	Type RefundSucceededWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RefundSucceededWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *RefundSucceededWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type RefundSucceededWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Refund".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Refund
}

// Returns the unmodified JSON received from the API
func (r RefundSucceededWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *RefundSucceededWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type RefundSucceededWebhookEventType string

const (
	RefundSucceededWebhookEventTypeRefundSucceeded RefundSucceededWebhookEventType = "refund.succeeded"
)

type SubscriptionActiveWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionActiveWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.active".
	Type SubscriptionActiveWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionActiveWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionActiveWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionActiveWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionActiveWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionActiveWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionActiveWebhookEventType string

const (
	SubscriptionActiveWebhookEventTypeSubscriptionActive SubscriptionActiveWebhookEventType = "subscription.active"
)

type SubscriptionCancelledWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionCancelledWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.cancelled".
	Type SubscriptionCancelledWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionCancelledWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionCancelledWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionCancelledWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionCancelledWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionCancelledWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionCancelledWebhookEventType string

const (
	SubscriptionCancelledWebhookEventTypeSubscriptionCancelled SubscriptionCancelledWebhookEventType = "subscription.cancelled"
)

type SubscriptionExpiredWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionExpiredWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.expired".
	Type SubscriptionExpiredWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionExpiredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionExpiredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionExpiredWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionExpiredWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionExpiredWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionExpiredWebhookEventType string

const (
	SubscriptionExpiredWebhookEventTypeSubscriptionExpired SubscriptionExpiredWebhookEventType = "subscription.expired"
)

type SubscriptionFailedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionFailedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.failed".
	Type SubscriptionFailedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionFailedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionFailedWebhookEventType string

const (
	SubscriptionFailedWebhookEventTypeSubscriptionFailed SubscriptionFailedWebhookEventType = "subscription.failed"
)

type SubscriptionOnHoldWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionOnHoldWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.on_hold".
	Type SubscriptionOnHoldWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionOnHoldWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionOnHoldWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionOnHoldWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionOnHoldWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionOnHoldWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionOnHoldWebhookEventType string

const (
	SubscriptionOnHoldWebhookEventTypeSubscriptionOnHold SubscriptionOnHoldWebhookEventType = "subscription.on_hold"
)

type SubscriptionPlanChangedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionPlanChangedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.plan_changed".
	Type SubscriptionPlanChangedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPlanChangedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionPlanChangedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionPlanChangedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionPlanChangedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionPlanChangedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionPlanChangedWebhookEventType string

const (
	SubscriptionPlanChangedWebhookEventTypeSubscriptionPlanChanged SubscriptionPlanChangedWebhookEventType = "subscription.plan_changed"
)

type SubscriptionRenewedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionRenewedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.renewed".
	Type SubscriptionRenewedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionRenewedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionRenewedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionRenewedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionRenewedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionRenewedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionRenewedWebhookEventType string

const (
	SubscriptionRenewedWebhookEventTypeSubscriptionRenewed SubscriptionRenewedWebhookEventType = "subscription.renewed"
)

type SubscriptionUpdatedWebhookEvent struct {
	// The business identifier
	BusinessID string `json:"business_id,required"`
	// Event-specific data
	Data SubscriptionUpdatedWebhookEventData `json:"data,required"`
	// The timestamp of when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The event type
	//
	// Any of "subscription.updated".
	Type SubscriptionUpdatedWebhookEventType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessID  respjson.Field
		Data        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-specific data
type SubscriptionUpdatedWebhookEventData struct {
	// The type of payload in the data field
	//
	// Any of "Subscription".
	PayloadType string `json:"payload_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Subscription
}

// Returns the unmodified JSON received from the API
func (r SubscriptionUpdatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionUpdatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event type
type SubscriptionUpdatedWebhookEventType string

const (
	SubscriptionUpdatedWebhookEventTypeSubscriptionUpdated SubscriptionUpdatedWebhookEventType = "subscription.updated"
)

// UnsafeUnwrapWebhookEventUnion contains all possible properties and values from
// [DisputeAcceptedWebhookEvent], [DisputeCancelledWebhookEvent],
// [DisputeChallengedWebhookEvent], [DisputeExpiredWebhookEvent],
// [DisputeLostWebhookEvent], [DisputeOpenedWebhookEvent],
// [DisputeWonWebhookEvent], [LicenseKeyCreatedWebhookEvent],
// [PaymentCancelledWebhookEvent], [PaymentFailedWebhookEvent],
// [PaymentProcessingWebhookEvent], [PaymentSucceededWebhookEvent],
// [RefundFailedWebhookEvent], [RefundSucceededWebhookEvent],
// [SubscriptionActiveWebhookEvent], [SubscriptionCancelledWebhookEvent],
// [SubscriptionExpiredWebhookEvent], [SubscriptionFailedWebhookEvent],
// [SubscriptionOnHoldWebhookEvent], [SubscriptionPlanChangedWebhookEvent],
// [SubscriptionRenewedWebhookEvent], [SubscriptionUpdatedWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnsafeUnwrapWebhookEventUnion struct {
	BusinessID string `json:"business_id"`
	// This field is a union of [DisputeAcceptedWebhookEventData],
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
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData]
	Data      UnsafeUnwrapWebhookEventUnionData `json:"data"`
	Timestamp time.Time                         `json:"timestamp"`
	Type      string                            `json:"type"`
	JSON      struct {
		BusinessID respjson.Field
		Data       respjson.Field
		Timestamp  respjson.Field
		Type       respjson.Field
		raw        string
	} `json:"-"`
}

func (u UnsafeUnwrapWebhookEventUnion) AsDisputeAcceptedWebhookEvent() (v DisputeAcceptedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsDisputeCancelledWebhookEvent() (v DisputeCancelledWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsDisputeChallengedWebhookEvent() (v DisputeChallengedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsDisputeExpiredWebhookEvent() (v DisputeExpiredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsDisputeLostWebhookEvent() (v DisputeLostWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsDisputeOpenedWebhookEvent() (v DisputeOpenedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsDisputeWonWebhookEvent() (v DisputeWonWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsLicenseKeyCreatedWebhookEvent() (v LicenseKeyCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsPaymentCancelledWebhookEvent() (v PaymentCancelledWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsPaymentFailedWebhookEvent() (v PaymentFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsPaymentProcessingWebhookEvent() (v PaymentProcessingWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsPaymentSucceededWebhookEvent() (v PaymentSucceededWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsRefundFailedWebhookEvent() (v RefundFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsRefundSucceededWebhookEvent() (v RefundSucceededWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionActiveWebhookEvent() (v SubscriptionActiveWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionCancelledWebhookEvent() (v SubscriptionCancelledWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionExpiredWebhookEvent() (v SubscriptionExpiredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionFailedWebhookEvent() (v SubscriptionFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionOnHoldWebhookEvent() (v SubscriptionOnHoldWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionPlanChangedWebhookEvent() (v SubscriptionPlanChangedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionRenewedWebhookEvent() (v SubscriptionRenewedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSubscriptionUpdatedWebhookEvent() (v SubscriptionUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnsafeUnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnsafeUnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionData is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionData struct {
	// This field is a union of [string], [int64]
	Amount     UnsafeUnwrapWebhookEventUnionDataAmount `json:"amount"`
	BusinessID string                                  `json:"business_id"`
	CreatedAt  time.Time                               `json:"created_at"`
	Currency   string                                  `json:"currency"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	DisputeID string `json:"dispute_id"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	DisputeStage DisputeStage `json:"dispute_stage"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	DisputeStatus DisputeStatus `json:"dispute_status"`
	PaymentID     string        `json:"payment_id"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	Remarks     string `json:"remarks"`
	PayloadType string `json:"payload_type"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	ID string `json:"id"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	CustomerID string `json:"customer_id"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	InstancesCount int64 `json:"instances_count"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	Key       string `json:"key"`
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	ActivationsLimit int64     `json:"activations_limit"`
	ExpiresAt        time.Time `json:"expires_at"`
	SubscriptionID   string    `json:"subscription_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Billing BillingAddress `json:"billing"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	BrandID string `json:"brand_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Customer CustomerLimitedDetails `json:"customer"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	DigitalProductsDelivered bool `json:"digital_products_delivered"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Disputes []Dispute `json:"disputes"`
	Metadata string    `json:"metadata"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Refunds []PaymentRefund `json:"refunds"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	SettlementAmount int64 `json:"settlement_amount"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	SettlementCurrency Currency `json:"settlement_currency"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	TotalAmount int64 `json:"total_amount"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardIssuingCountry CountryCode `json:"card_issuing_country"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardLastFour string `json:"card_last_four"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardNetwork string `json:"card_network"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardType string `json:"card_type"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CheckoutSessionID string `json:"checkout_session_id"`
	DiscountID        string `json:"discount_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	ErrorCode string `json:"error_code"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	ErrorMessage string `json:"error_message"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	InvoiceID string `json:"invoice_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	PaymentLink string `json:"payment_link"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	PaymentMethod string `json:"payment_method"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	PaymentMethodType string `json:"payment_method_type"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	ProductCart []PaymentProductCart `json:"product_cart"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	SettlementTax int64 `json:"settlement_tax"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Tax int64 `json:"tax"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData].
	IsPartial bool `json:"is_partial"`
	// This field is from variant [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData].
	RefundID string `json:"refund_id"`
	// This field is from variant [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData].
	Reason string `json:"reason"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	Addons []AddonCartResponseItem `json:"addons"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	CancelAtNextBillingDate bool `json:"cancel_at_next_billing_date"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	Meters []SubscriptionMeter `json:"meters"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	NextBillingDate time.Time `json:"next_billing_date"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	OnDemand bool `json:"on_demand"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PaymentFrequencyCount int64 `json:"payment_frequency_count"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PreviousBillingDate time.Time `json:"previous_billing_date"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	Quantity int64 `json:"quantity"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	SubscriptionPeriodCount int64 `json:"subscription_period_count"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	TaxInclusive bool `json:"tax_inclusive"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	TrialPeriodDays int64 `json:"trial_period_days"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	CancelledAt time.Time `json:"cancelled_at"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	DiscountCyclesRemaining int64 `json:"discount_cycles_remaining"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PaymentMethodID string `json:"payment_method_id"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	TaxID string `json:"tax_id"`
	JSON  struct {
		Amount                     respjson.Field
		BusinessID                 respjson.Field
		CreatedAt                  respjson.Field
		Currency                   respjson.Field
		DisputeID                  respjson.Field
		DisputeStage               respjson.Field
		DisputeStatus              respjson.Field
		PaymentID                  respjson.Field
		Remarks                    respjson.Field
		PayloadType                respjson.Field
		ID                         respjson.Field
		CustomerID                 respjson.Field
		InstancesCount             respjson.Field
		Key                        respjson.Field
		ProductID                  respjson.Field
		Status                     respjson.Field
		ActivationsLimit           respjson.Field
		ExpiresAt                  respjson.Field
		SubscriptionID             respjson.Field
		Billing                    respjson.Field
		BrandID                    respjson.Field
		Customer                   respjson.Field
		DigitalProductsDelivered   respjson.Field
		Disputes                   respjson.Field
		Metadata                   respjson.Field
		Refunds                    respjson.Field
		SettlementAmount           respjson.Field
		SettlementCurrency         respjson.Field
		TotalAmount                respjson.Field
		CardIssuingCountry         respjson.Field
		CardLastFour               respjson.Field
		CardNetwork                respjson.Field
		CardType                   respjson.Field
		CheckoutSessionID          respjson.Field
		DiscountID                 respjson.Field
		ErrorCode                  respjson.Field
		ErrorMessage               respjson.Field
		InvoiceID                  respjson.Field
		PaymentLink                respjson.Field
		PaymentMethod              respjson.Field
		PaymentMethodType          respjson.Field
		ProductCart                respjson.Field
		SettlementTax              respjson.Field
		Tax                        respjson.Field
		UpdatedAt                  respjson.Field
		IsPartial                  respjson.Field
		RefundID                   respjson.Field
		Reason                     respjson.Field
		Addons                     respjson.Field
		CancelAtNextBillingDate    respjson.Field
		Meters                     respjson.Field
		NextBillingDate            respjson.Field
		OnDemand                   respjson.Field
		PaymentFrequencyCount      respjson.Field
		PaymentFrequencyInterval   respjson.Field
		PreviousBillingDate        respjson.Field
		Quantity                   respjson.Field
		RecurringPreTaxAmount      respjson.Field
		SubscriptionPeriodCount    respjson.Field
		SubscriptionPeriodInterval respjson.Field
		TaxInclusive               respjson.Field
		TrialPeriodDays            respjson.Field
		CancelledAt                respjson.Field
		DiscountCyclesRemaining    respjson.Field
		PaymentMethodID            respjson.Field
		TaxID                      respjson.Field
		raw                        string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataAmount is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataAmount
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type UnsafeUnwrapWebhookEventUnionDataAmount struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnion contains all possible properties and values from
// [DisputeAcceptedWebhookEvent], [DisputeCancelledWebhookEvent],
// [DisputeChallengedWebhookEvent], [DisputeExpiredWebhookEvent],
// [DisputeLostWebhookEvent], [DisputeOpenedWebhookEvent],
// [DisputeWonWebhookEvent], [LicenseKeyCreatedWebhookEvent],
// [PaymentCancelledWebhookEvent], [PaymentFailedWebhookEvent],
// [PaymentProcessingWebhookEvent], [PaymentSucceededWebhookEvent],
// [RefundFailedWebhookEvent], [RefundSucceededWebhookEvent],
// [SubscriptionActiveWebhookEvent], [SubscriptionCancelledWebhookEvent],
// [SubscriptionExpiredWebhookEvent], [SubscriptionFailedWebhookEvent],
// [SubscriptionOnHoldWebhookEvent], [SubscriptionPlanChangedWebhookEvent],
// [SubscriptionRenewedWebhookEvent], [SubscriptionUpdatedWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	BusinessID string `json:"business_id"`
	// This field is a union of [DisputeAcceptedWebhookEventData],
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
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData]
	Data      UnwrapWebhookEventUnionData `json:"data"`
	Timestamp time.Time                   `json:"timestamp"`
	Type      string                      `json:"type"`
	JSON      struct {
		BusinessID respjson.Field
		Data       respjson.Field
		Timestamp  respjson.Field
		Type       respjson.Field
		raw        string
	} `json:"-"`
}

func (u UnwrapWebhookEventUnion) AsDisputeAcceptedWebhookEvent() (v DisputeAcceptedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsDisputeCancelledWebhookEvent() (v DisputeCancelledWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsDisputeChallengedWebhookEvent() (v DisputeChallengedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsDisputeExpiredWebhookEvent() (v DisputeExpiredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsDisputeLostWebhookEvent() (v DisputeLostWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsDisputeOpenedWebhookEvent() (v DisputeOpenedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsDisputeWonWebhookEvent() (v DisputeWonWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsLicenseKeyCreatedWebhookEvent() (v LicenseKeyCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentCancelledWebhookEvent() (v PaymentCancelledWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentFailedWebhookEvent() (v PaymentFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentProcessingWebhookEvent() (v PaymentProcessingWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsPaymentSucceededWebhookEvent() (v PaymentSucceededWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsRefundFailedWebhookEvent() (v RefundFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsRefundSucceededWebhookEvent() (v RefundSucceededWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionActiveWebhookEvent() (v SubscriptionActiveWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionCancelledWebhookEvent() (v SubscriptionCancelledWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionExpiredWebhookEvent() (v SubscriptionExpiredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionFailedWebhookEvent() (v SubscriptionFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionOnHoldWebhookEvent() (v SubscriptionOnHoldWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionPlanChangedWebhookEvent() (v SubscriptionPlanChangedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionRenewedWebhookEvent() (v SubscriptionRenewedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSubscriptionUpdatedWebhookEvent() (v SubscriptionUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionData is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionData struct {
	// This field is a union of [string], [int64]
	Amount     UnwrapWebhookEventUnionDataAmount `json:"amount"`
	BusinessID string                            `json:"business_id"`
	CreatedAt  time.Time                         `json:"created_at"`
	Currency   string                            `json:"currency"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	DisputeID string `json:"dispute_id"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	DisputeStage DisputeStage `json:"dispute_stage"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	DisputeStatus DisputeStatus `json:"dispute_status"`
	PaymentID     string        `json:"payment_id"`
	// This field is from variant [DisputeAcceptedWebhookEventData],
	// [DisputeCancelledWebhookEventData], [DisputeChallengedWebhookEventData],
	// [DisputeExpiredWebhookEventData], [DisputeLostWebhookEventData],
	// [DisputeOpenedWebhookEventData], [DisputeWonWebhookEventData].
	Remarks     string `json:"remarks"`
	PayloadType string `json:"payload_type"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	ID string `json:"id"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	CustomerID string `json:"customer_id"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	InstancesCount int64 `json:"instances_count"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	Key       string `json:"key"`
	ProductID string `json:"product_id"`
	Status    string `json:"status"`
	// This field is from variant [LicenseKeyCreatedWebhookEventData].
	ActivationsLimit int64     `json:"activations_limit"`
	ExpiresAt        time.Time `json:"expires_at"`
	SubscriptionID   string    `json:"subscription_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Billing BillingAddress `json:"billing"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	BrandID string `json:"brand_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Customer CustomerLimitedDetails `json:"customer"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	DigitalProductsDelivered bool `json:"digital_products_delivered"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Disputes []Dispute `json:"disputes"`
	Metadata string    `json:"metadata"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Refunds []PaymentRefund `json:"refunds"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	SettlementAmount int64 `json:"settlement_amount"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	SettlementCurrency Currency `json:"settlement_currency"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	TotalAmount int64 `json:"total_amount"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardIssuingCountry CountryCode `json:"card_issuing_country"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardLastFour string `json:"card_last_four"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardNetwork string `json:"card_network"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CardType string `json:"card_type"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	CheckoutSessionID string `json:"checkout_session_id"`
	DiscountID        string `json:"discount_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	ErrorCode string `json:"error_code"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	ErrorMessage string `json:"error_message"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	InvoiceID string `json:"invoice_id"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	PaymentLink string `json:"payment_link"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	PaymentMethod string `json:"payment_method"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	PaymentMethodType string `json:"payment_method_type"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	ProductCart []PaymentProductCart `json:"product_cart"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	SettlementTax int64 `json:"settlement_tax"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	Tax int64 `json:"tax"`
	// This field is from variant [PaymentCancelledWebhookEventData],
	// [PaymentFailedWebhookEventData], [PaymentProcessingWebhookEventData],
	// [PaymentSucceededWebhookEventData].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData].
	IsPartial bool `json:"is_partial"`
	// This field is from variant [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData].
	RefundID string `json:"refund_id"`
	// This field is from variant [RefundFailedWebhookEventData],
	// [RefundSucceededWebhookEventData].
	Reason string `json:"reason"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	Addons []AddonCartResponseItem `json:"addons"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	CancelAtNextBillingDate bool `json:"cancel_at_next_billing_date"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	Meters []SubscriptionMeter `json:"meters"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	NextBillingDate time.Time `json:"next_billing_date"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	OnDemand bool `json:"on_demand"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PaymentFrequencyCount int64 `json:"payment_frequency_count"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PaymentFrequencyInterval TimeInterval `json:"payment_frequency_interval"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PreviousBillingDate time.Time `json:"previous_billing_date"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	Quantity int64 `json:"quantity"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	RecurringPreTaxAmount int64 `json:"recurring_pre_tax_amount"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	SubscriptionPeriodCount int64 `json:"subscription_period_count"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	SubscriptionPeriodInterval TimeInterval `json:"subscription_period_interval"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	TaxInclusive bool `json:"tax_inclusive"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	TrialPeriodDays int64 `json:"trial_period_days"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	CancelledAt time.Time `json:"cancelled_at"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	DiscountCyclesRemaining int64 `json:"discount_cycles_remaining"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	PaymentMethodID string `json:"payment_method_id"`
	// This field is from variant [SubscriptionActiveWebhookEventData],
	// [SubscriptionCancelledWebhookEventData], [SubscriptionExpiredWebhookEventData],
	// [SubscriptionFailedWebhookEventData], [SubscriptionOnHoldWebhookEventData],
	// [SubscriptionPlanChangedWebhookEventData],
	// [SubscriptionRenewedWebhookEventData], [SubscriptionUpdatedWebhookEventData].
	TaxID string `json:"tax_id"`
	JSON  struct {
		Amount                     respjson.Field
		BusinessID                 respjson.Field
		CreatedAt                  respjson.Field
		Currency                   respjson.Field
		DisputeID                  respjson.Field
		DisputeStage               respjson.Field
		DisputeStatus              respjson.Field
		PaymentID                  respjson.Field
		Remarks                    respjson.Field
		PayloadType                respjson.Field
		ID                         respjson.Field
		CustomerID                 respjson.Field
		InstancesCount             respjson.Field
		Key                        respjson.Field
		ProductID                  respjson.Field
		Status                     respjson.Field
		ActivationsLimit           respjson.Field
		ExpiresAt                  respjson.Field
		SubscriptionID             respjson.Field
		Billing                    respjson.Field
		BrandID                    respjson.Field
		Customer                   respjson.Field
		DigitalProductsDelivered   respjson.Field
		Disputes                   respjson.Field
		Metadata                   respjson.Field
		Refunds                    respjson.Field
		SettlementAmount           respjson.Field
		SettlementCurrency         respjson.Field
		TotalAmount                respjson.Field
		CardIssuingCountry         respjson.Field
		CardLastFour               respjson.Field
		CardNetwork                respjson.Field
		CardType                   respjson.Field
		CheckoutSessionID          respjson.Field
		DiscountID                 respjson.Field
		ErrorCode                  respjson.Field
		ErrorMessage               respjson.Field
		InvoiceID                  respjson.Field
		PaymentLink                respjson.Field
		PaymentMethod              respjson.Field
		PaymentMethodType          respjson.Field
		ProductCart                respjson.Field
		SettlementTax              respjson.Field
		Tax                        respjson.Field
		UpdatedAt                  respjson.Field
		IsPartial                  respjson.Field
		RefundID                   respjson.Field
		Reason                     respjson.Field
		Addons                     respjson.Field
		CancelAtNextBillingDate    respjson.Field
		Meters                     respjson.Field
		NextBillingDate            respjson.Field
		OnDemand                   respjson.Field
		PaymentFrequencyCount      respjson.Field
		PaymentFrequencyInterval   respjson.Field
		PreviousBillingDate        respjson.Field
		Quantity                   respjson.Field
		RecurringPreTaxAmount      respjson.Field
		SubscriptionPeriodCount    respjson.Field
		SubscriptionPeriodInterval respjson.Field
		TaxInclusive               respjson.Field
		TrialPeriodDays            respjson.Field
		CancelledAt                respjson.Field
		DiscountCyclesRemaining    respjson.Field
		PaymentMethodID            respjson.Field
		TaxID                      respjson.Field
		raw                        string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataAmount is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataAmount provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type UnwrapWebhookEventUnionDataAmount struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// Url of the webhook
	URL         string            `json:"url,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Create the webhook in a disabled state.
	//
	// Default is false
	Disabled param.Opt[bool] `json:"disabled,omitzero"`
	// The request's idempotency key
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	RateLimit      param.Opt[int64]  `json:"rate_limit,omitzero"`
	// Custom headers to be passed
	Headers map[string]string `json:"headers,omitzero"`
	// Metadata to be passed to the webhook Defaut is {}
	Metadata map[string]string `json:"metadata,omitzero"`
	// Filter events to the webhook.
	//
	// Webhook event will only be sent for events in the list.
	FilterTypes []WebhookEventType `json:"filter_types,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	// Description of the webhook
	Description param.Opt[string] `json:"description,omitzero"`
	// To Disable the endpoint, set it to true.
	Disabled param.Opt[bool] `json:"disabled,omitzero"`
	// Rate limit
	RateLimit param.Opt[int64] `json:"rate_limit,omitzero"`
	// Url endpoint
	URL param.Opt[string] `json:"url,omitzero"`
	// Filter events to the endpoint.
	//
	// Webhook event will only be sent for events in the list.
	FilterTypes []WebhookEventType `json:"filter_types,omitzero"`
	// Metadata
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListParams struct {
	// The iterator returned from a prior invocation
	Iterator param.Opt[string] `query:"iterator,omitzero" json:"-"`
	// Limit the number of returned items
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
