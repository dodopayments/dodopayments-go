// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// YourWebhookURLService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewYourWebhookURLService] method instead.
type YourWebhookURLService struct {
	Options []option.RequestOption
}

// NewYourWebhookURLService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewYourWebhookURLService(opts ...option.RequestOption) (r *YourWebhookURLService) {
	r = &YourWebhookURLService{}
	r.Options = opts
	return
}

func (r *YourWebhookURLService) New(ctx context.Context, params YourWebhookURLNewParams, opts ...option.RequestOption) (err error) {
	if params.WebhookID.Present {
		opts = append(opts, option.WithHeader("webhook-id", fmt.Sprintf("%v", params.WebhookID)))
	}
	if params.WebhookSignature.Present {
		opts = append(opts, option.WithHeader("webhook-signature", fmt.Sprintf("%v", params.WebhookSignature)))
	}
	if params.WebhookTimestamp.Present {
		opts = append(opts, option.WithHeader("webhook-timestamp", fmt.Sprintf("%v", params.WebhookTimestamp)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "your-webhook-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

type YourWebhookURLNewParams struct {
	WebhookPayload   WebhookPayloadParam `json:"webhook_payload" api:"required"`
	WebhookID        param.Field[string] `header:"webhook-id" api:"required"`
	WebhookSignature param.Field[string] `header:"webhook-signature" api:"required"`
	WebhookTimestamp param.Field[string] `header:"webhook-timestamp" api:"required"`
}

func (r YourWebhookURLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.WebhookPayload)
}
