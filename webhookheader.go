// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// WebhookHeaderService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookHeaderService] method instead.
type WebhookHeaderService struct {
	Options []option.RequestOption
}

// NewWebhookHeaderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookHeaderService(opts ...option.RequestOption) (r WebhookHeaderService) {
	r = WebhookHeaderService{}
	r.Options = opts
	return
}

// Get a webhook by id
func (r *WebhookHeaderService) Get(ctx context.Context, webhookID string, opts ...option.RequestOption) (res *WebhookHeaderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/headers", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch a webhook by id
func (r *WebhookHeaderService) Update(ctx context.Context, webhookID string, body WebhookHeaderUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/headers", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// The value of the headers is returned in the `headers` field.
//
// Sensitive headers that have been redacted are returned in the sensitive field.
type WebhookHeaderGetResponse struct {
	// List of headers configured
	Headers map[string]string `json:"headers,required"`
	// Sensitive headers without the value
	Sensitive []string `json:"sensitive,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers     respjson.Field
		Sensitive   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookHeaderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookHeaderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookHeaderUpdateParams struct {
	// Object of header-value pair to update or add
	Headers map[string]string `json:"headers,omitzero,required"`
	paramObj
}

func (r WebhookHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookHeaderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookHeaderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
