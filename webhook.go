// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
