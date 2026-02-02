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

// Event types for Dodo events
type WebhookEventType string

const (
	WebhookEventTypePaymentSucceeded        WebhookEventType = "payment.succeeded"
	WebhookEventTypePaymentFailed           WebhookEventType = "payment.failed"
	WebhookEventTypePaymentProcessing       WebhookEventType = "payment.processing"
	WebhookEventTypePaymentCancelled        WebhookEventType = "payment.cancelled"
	WebhookEventTypeRefundSucceeded         WebhookEventType = "refund.succeeded"
	WebhookEventTypeRefundFailed            WebhookEventType = "refund.failed"
	WebhookEventTypeDisputeOpened           WebhookEventType = "dispute.opened"
	WebhookEventTypeDisputeExpired          WebhookEventType = "dispute.expired"
	WebhookEventTypeDisputeAccepted         WebhookEventType = "dispute.accepted"
	WebhookEventTypeDisputeCancelled        WebhookEventType = "dispute.cancelled"
	WebhookEventTypeDisputeChallenged       WebhookEventType = "dispute.challenged"
	WebhookEventTypeDisputeWon              WebhookEventType = "dispute.won"
	WebhookEventTypeDisputeLost             WebhookEventType = "dispute.lost"
	WebhookEventTypeSubscriptionActive      WebhookEventType = "subscription.active"
	WebhookEventTypeSubscriptionRenewed     WebhookEventType = "subscription.renewed"
	WebhookEventTypeSubscriptionOnHold      WebhookEventType = "subscription.on_hold"
	WebhookEventTypeSubscriptionCancelled   WebhookEventType = "subscription.cancelled"
	WebhookEventTypeSubscriptionFailed      WebhookEventType = "subscription.failed"
	WebhookEventTypeSubscriptionExpired     WebhookEventType = "subscription.expired"
	WebhookEventTypeSubscriptionPlanChanged WebhookEventType = "subscription.plan_changed"
	WebhookEventTypeSubscriptionUpdated     WebhookEventType = "subscription.updated"
	WebhookEventTypeLicenseKeyCreated       WebhookEventType = "license_key.created"
	WebhookEventTypePayoutNotInitiated      WebhookEventType = "payout.not_initiated"
	WebhookEventTypePayoutOnHold            WebhookEventType = "payout.on_hold"
	WebhookEventTypePayoutInProgress        WebhookEventType = "payout.in_progress"
	WebhookEventTypePayoutFailed            WebhookEventType = "payout.failed"
	WebhookEventTypePayoutSuccess           WebhookEventType = "payout.success"
)

func (r WebhookEventType) IsKnown() bool {
	switch r {
	case WebhookEventTypePaymentSucceeded, WebhookEventTypePaymentFailed, WebhookEventTypePaymentProcessing, WebhookEventTypePaymentCancelled, WebhookEventTypeRefundSucceeded, WebhookEventTypeRefundFailed, WebhookEventTypeDisputeOpened, WebhookEventTypeDisputeExpired, WebhookEventTypeDisputeAccepted, WebhookEventTypeDisputeCancelled, WebhookEventTypeDisputeChallenged, WebhookEventTypeDisputeWon, WebhookEventTypeDisputeLost, WebhookEventTypeSubscriptionActive, WebhookEventTypeSubscriptionRenewed, WebhookEventTypeSubscriptionOnHold, WebhookEventTypeSubscriptionCancelled, WebhookEventTypeSubscriptionFailed, WebhookEventTypeSubscriptionExpired, WebhookEventTypeSubscriptionPlanChanged, WebhookEventTypeSubscriptionUpdated, WebhookEventTypeLicenseKeyCreated, WebhookEventTypePayoutNotInitiated, WebhookEventTypePayoutOnHold, WebhookEventTypePayoutInProgress, WebhookEventTypePayoutFailed, WebhookEventTypePayoutSuccess:
		return true
	}
	return false
}
