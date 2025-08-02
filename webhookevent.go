// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"github.com/dodopayments/dodopayments-go/option"
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
