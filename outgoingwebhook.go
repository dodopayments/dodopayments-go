// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/dodo-payments-go/internal/apijson"
	"github.com/stainless-sdks/dodo-payments-go/internal/param"
	"github.com/stainless-sdks/dodo-payments-go/internal/requestconfig"
	"github.com/stainless-sdks/dodo-payments-go/option"
)

// OutgoingWebhookService contains methods and other services that help with
// interacting with the dodopayments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOutgoingWebhookService] method instead.
type OutgoingWebhookService struct {
	Options []option.RequestOption
}

// NewOutgoingWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOutgoingWebhookService(opts ...option.RequestOption) (r *OutgoingWebhookService) {
	r = &OutgoingWebhookService{}
	r.Options = opts
	return
}

func (r *OutgoingWebhookService) New(ctx context.Context, params OutgoingWebhookNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "your-webhook-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type OutgoingWebhookNewParams struct {
	BusinessID param.Field[string]                            `json:"business_id,required"`
	Data       param.Field[OutgoingWebhookNewParamsDataUnion] `json:"data,required"`
	// The timestamp of when the event occurred (not necessarily the same of when it
	// was delivered)
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Event types for Dodo events
	Type             param.Field[OutgoingWebhookNewParamsType] `json:"type,required"`
	WebhookID        param.Field[string]                       `header:"webhook-id,required"`
	WebhookSignature param.Field[string]                       `header:"webhook-signature,required"`
	WebhookTimestamp param.Field[string]                       `header:"webhook-timestamp,required"`
}

func (r OutgoingWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [OutgoingWebhookNewParamsDataPayment],
// [OutgoingWebhookNewParamsDataSubscription],
// [OutgoingWebhookNewParamsDataRefund], [OutgoingWebhookNewParamsDataDispute].
type OutgoingWebhookNewParamsDataUnion interface {
	implementsOutgoingWebhookNewParamsDataUnion()
}

type OutgoingWebhookNewParamsDataPayment struct {
	PayloadType param.Field[OutgoingWebhookNewParamsDataPaymentPayloadType] `json:"payload_type,required"`
	PaymentParam
}

func (r OutgoingWebhookNewParamsDataPayment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OutgoingWebhookNewParamsDataPayment) implementsOutgoingWebhookNewParamsDataUnion() {}

type OutgoingWebhookNewParamsDataPaymentPayloadType string

const (
	OutgoingWebhookNewParamsDataPaymentPayloadTypePayment OutgoingWebhookNewParamsDataPaymentPayloadType = "Payment"
)

func (r OutgoingWebhookNewParamsDataPaymentPayloadType) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataPaymentPayloadTypePayment:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataSubscription struct {
	PayloadType param.Field[OutgoingWebhookNewParamsDataSubscriptionPayloadType] `json:"payload_type,required"`
	SubscriptionParam
}

func (r OutgoingWebhookNewParamsDataSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OutgoingWebhookNewParamsDataSubscription) implementsOutgoingWebhookNewParamsDataUnion() {}

type OutgoingWebhookNewParamsDataSubscriptionPayloadType string

const (
	OutgoingWebhookNewParamsDataSubscriptionPayloadTypeSubscription OutgoingWebhookNewParamsDataSubscriptionPayloadType = "Subscription"
)

func (r OutgoingWebhookNewParamsDataSubscriptionPayloadType) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataSubscriptionPayloadTypeSubscription:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataRefund struct {
	PayloadType param.Field[OutgoingWebhookNewParamsDataRefundPayloadType] `json:"payload_type,required"`
	RefundParam
}

func (r OutgoingWebhookNewParamsDataRefund) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OutgoingWebhookNewParamsDataRefund) implementsOutgoingWebhookNewParamsDataUnion() {}

type OutgoingWebhookNewParamsDataRefundPayloadType string

const (
	OutgoingWebhookNewParamsDataRefundPayloadTypeRefund OutgoingWebhookNewParamsDataRefundPayloadType = "Refund"
)

func (r OutgoingWebhookNewParamsDataRefundPayloadType) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataRefundPayloadTypeRefund:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataDispute struct {
	PayloadType param.Field[OutgoingWebhookNewParamsDataDisputePayloadType] `json:"payload_type,required"`
	DisputeParam
}

func (r OutgoingWebhookNewParamsDataDispute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OutgoingWebhookNewParamsDataDispute) implementsOutgoingWebhookNewParamsDataUnion() {}

type OutgoingWebhookNewParamsDataDisputePayloadType string

const (
	OutgoingWebhookNewParamsDataDisputePayloadTypeDispute OutgoingWebhookNewParamsDataDisputePayloadType = "Dispute"
)

func (r OutgoingWebhookNewParamsDataDisputePayloadType) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataDisputePayloadTypeDispute:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataDisputeStage string

const (
	OutgoingWebhookNewParamsDataDisputeStagePreDispute     OutgoingWebhookNewParamsDataDisputeStage = "pre_dispute"
	OutgoingWebhookNewParamsDataDisputeStageDispute        OutgoingWebhookNewParamsDataDisputeStage = "dispute"
	OutgoingWebhookNewParamsDataDisputeStagePreArbitration OutgoingWebhookNewParamsDataDisputeStage = "pre_arbitration"
)

func (r OutgoingWebhookNewParamsDataDisputeStage) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataDisputeStagePreDispute, OutgoingWebhookNewParamsDataDisputeStageDispute, OutgoingWebhookNewParamsDataDisputeStagePreArbitration:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataDisputeStatus string

const (
	OutgoingWebhookNewParamsDataDisputeStatusDisputeOpened     OutgoingWebhookNewParamsDataDisputeStatus = "dispute_opened"
	OutgoingWebhookNewParamsDataDisputeStatusDisputeExpired    OutgoingWebhookNewParamsDataDisputeStatus = "dispute_expired"
	OutgoingWebhookNewParamsDataDisputeStatusDisputeAccepted   OutgoingWebhookNewParamsDataDisputeStatus = "dispute_accepted"
	OutgoingWebhookNewParamsDataDisputeStatusDisputeCancelled  OutgoingWebhookNewParamsDataDisputeStatus = "dispute_cancelled"
	OutgoingWebhookNewParamsDataDisputeStatusDisputeChallenged OutgoingWebhookNewParamsDataDisputeStatus = "dispute_challenged"
	OutgoingWebhookNewParamsDataDisputeStatusDisputeWon        OutgoingWebhookNewParamsDataDisputeStatus = "dispute_won"
	OutgoingWebhookNewParamsDataDisputeStatusDisputeLost       OutgoingWebhookNewParamsDataDisputeStatus = "dispute_lost"
)

func (r OutgoingWebhookNewParamsDataDisputeStatus) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataDisputeStatusDisputeOpened, OutgoingWebhookNewParamsDataDisputeStatusDisputeExpired, OutgoingWebhookNewParamsDataDisputeStatusDisputeAccepted, OutgoingWebhookNewParamsDataDisputeStatusDisputeCancelled, OutgoingWebhookNewParamsDataDisputeStatusDisputeChallenged, OutgoingWebhookNewParamsDataDisputeStatusDisputeWon, OutgoingWebhookNewParamsDataDisputeStatusDisputeLost:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataPayloadType string

const (
	OutgoingWebhookNewParamsDataPayloadTypePayment      OutgoingWebhookNewParamsDataPayloadType = "Payment"
	OutgoingWebhookNewParamsDataPayloadTypeSubscription OutgoingWebhookNewParamsDataPayloadType = "Subscription"
	OutgoingWebhookNewParamsDataPayloadTypeRefund       OutgoingWebhookNewParamsDataPayloadType = "Refund"
	OutgoingWebhookNewParamsDataPayloadTypeDispute      OutgoingWebhookNewParamsDataPayloadType = "Dispute"
)

func (r OutgoingWebhookNewParamsDataPayloadType) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataPayloadTypePayment, OutgoingWebhookNewParamsDataPayloadTypeSubscription, OutgoingWebhookNewParamsDataPayloadTypeRefund, OutgoingWebhookNewParamsDataPayloadTypeDispute:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataPaymentFrequencyInterval string

const (
	OutgoingWebhookNewParamsDataPaymentFrequencyIntervalDay   OutgoingWebhookNewParamsDataPaymentFrequencyInterval = "Day"
	OutgoingWebhookNewParamsDataPaymentFrequencyIntervalWeek  OutgoingWebhookNewParamsDataPaymentFrequencyInterval = "Week"
	OutgoingWebhookNewParamsDataPaymentFrequencyIntervalMonth OutgoingWebhookNewParamsDataPaymentFrequencyInterval = "Month"
	OutgoingWebhookNewParamsDataPaymentFrequencyIntervalYear  OutgoingWebhookNewParamsDataPaymentFrequencyInterval = "Year"
)

func (r OutgoingWebhookNewParamsDataPaymentFrequencyInterval) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataPaymentFrequencyIntervalDay, OutgoingWebhookNewParamsDataPaymentFrequencyIntervalWeek, OutgoingWebhookNewParamsDataPaymentFrequencyIntervalMonth, OutgoingWebhookNewParamsDataPaymentFrequencyIntervalYear:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataStatus string

const (
	OutgoingWebhookNewParamsDataStatusSucceeded                      OutgoingWebhookNewParamsDataStatus = "succeeded"
	OutgoingWebhookNewParamsDataStatusFailed                         OutgoingWebhookNewParamsDataStatus = "failed"
	OutgoingWebhookNewParamsDataStatusCancelled                      OutgoingWebhookNewParamsDataStatus = "cancelled"
	OutgoingWebhookNewParamsDataStatusProcessing                     OutgoingWebhookNewParamsDataStatus = "processing"
	OutgoingWebhookNewParamsDataStatusRequiresCustomerAction         OutgoingWebhookNewParamsDataStatus = "requires_customer_action"
	OutgoingWebhookNewParamsDataStatusRequiresMerchantAction         OutgoingWebhookNewParamsDataStatus = "requires_merchant_action"
	OutgoingWebhookNewParamsDataStatusRequiresPaymentMethod          OutgoingWebhookNewParamsDataStatus = "requires_payment_method"
	OutgoingWebhookNewParamsDataStatusRequiresConfirmation           OutgoingWebhookNewParamsDataStatus = "requires_confirmation"
	OutgoingWebhookNewParamsDataStatusRequiresCapture                OutgoingWebhookNewParamsDataStatus = "requires_capture"
	OutgoingWebhookNewParamsDataStatusPartiallyCaptured              OutgoingWebhookNewParamsDataStatus = "partially_captured"
	OutgoingWebhookNewParamsDataStatusPartiallyCapturedAndCapturable OutgoingWebhookNewParamsDataStatus = "partially_captured_and_capturable"
	OutgoingWebhookNewParamsDataStatusPending                        OutgoingWebhookNewParamsDataStatus = "pending"
	OutgoingWebhookNewParamsDataStatusActive                         OutgoingWebhookNewParamsDataStatus = "active"
	OutgoingWebhookNewParamsDataStatusOnHold                         OutgoingWebhookNewParamsDataStatus = "on_hold"
	OutgoingWebhookNewParamsDataStatusPaused                         OutgoingWebhookNewParamsDataStatus = "paused"
	OutgoingWebhookNewParamsDataStatusExpired                        OutgoingWebhookNewParamsDataStatus = "expired"
	OutgoingWebhookNewParamsDataStatusReview                         OutgoingWebhookNewParamsDataStatus = "review"
)

func (r OutgoingWebhookNewParamsDataStatus) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataStatusSucceeded, OutgoingWebhookNewParamsDataStatusFailed, OutgoingWebhookNewParamsDataStatusCancelled, OutgoingWebhookNewParamsDataStatusProcessing, OutgoingWebhookNewParamsDataStatusRequiresCustomerAction, OutgoingWebhookNewParamsDataStatusRequiresMerchantAction, OutgoingWebhookNewParamsDataStatusRequiresPaymentMethod, OutgoingWebhookNewParamsDataStatusRequiresConfirmation, OutgoingWebhookNewParamsDataStatusRequiresCapture, OutgoingWebhookNewParamsDataStatusPartiallyCaptured, OutgoingWebhookNewParamsDataStatusPartiallyCapturedAndCapturable, OutgoingWebhookNewParamsDataStatusPending, OutgoingWebhookNewParamsDataStatusActive, OutgoingWebhookNewParamsDataStatusOnHold, OutgoingWebhookNewParamsDataStatusPaused, OutgoingWebhookNewParamsDataStatusExpired, OutgoingWebhookNewParamsDataStatusReview:
		return true
	}
	return false
}

type OutgoingWebhookNewParamsDataSubscriptionPeriodInterval string

const (
	OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalDay   OutgoingWebhookNewParamsDataSubscriptionPeriodInterval = "Day"
	OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalWeek  OutgoingWebhookNewParamsDataSubscriptionPeriodInterval = "Week"
	OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalMonth OutgoingWebhookNewParamsDataSubscriptionPeriodInterval = "Month"
	OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalYear  OutgoingWebhookNewParamsDataSubscriptionPeriodInterval = "Year"
)

func (r OutgoingWebhookNewParamsDataSubscriptionPeriodInterval) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalDay, OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalWeek, OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalMonth, OutgoingWebhookNewParamsDataSubscriptionPeriodIntervalYear:
		return true
	}
	return false
}

// Event types for Dodo events
type OutgoingWebhookNewParamsType string

const (
	OutgoingWebhookNewParamsTypePaymentSucceeded      OutgoingWebhookNewParamsType = "payment.succeeded"
	OutgoingWebhookNewParamsTypePaymentFailed         OutgoingWebhookNewParamsType = "payment.failed"
	OutgoingWebhookNewParamsTypePaymentProcessing     OutgoingWebhookNewParamsType = "payment.processing"
	OutgoingWebhookNewParamsTypePaymentCancelled      OutgoingWebhookNewParamsType = "payment.cancelled"
	OutgoingWebhookNewParamsTypeRefundSucceeded       OutgoingWebhookNewParamsType = "refund.succeeded"
	OutgoingWebhookNewParamsTypeRefundFailed          OutgoingWebhookNewParamsType = "refund.failed"
	OutgoingWebhookNewParamsTypeDisputeOpened         OutgoingWebhookNewParamsType = "dispute.opened"
	OutgoingWebhookNewParamsTypeDisputeExpired        OutgoingWebhookNewParamsType = "dispute.expired"
	OutgoingWebhookNewParamsTypeDisputeAccepted       OutgoingWebhookNewParamsType = "dispute.accepted"
	OutgoingWebhookNewParamsTypeDisputeCancelled      OutgoingWebhookNewParamsType = "dispute.cancelled"
	OutgoingWebhookNewParamsTypeDisputeChallenged     OutgoingWebhookNewParamsType = "dispute.challenged"
	OutgoingWebhookNewParamsTypeDisputeWon            OutgoingWebhookNewParamsType = "dispute.won"
	OutgoingWebhookNewParamsTypeDisputeLost           OutgoingWebhookNewParamsType = "dispute.lost"
	OutgoingWebhookNewParamsTypeSubscriptionActive    OutgoingWebhookNewParamsType = "subscription.active"
	OutgoingWebhookNewParamsTypeSubscriptionOnHold    OutgoingWebhookNewParamsType = "subscription.on_hold"
	OutgoingWebhookNewParamsTypeSubscriptionPaused    OutgoingWebhookNewParamsType = "subscription.paused"
	OutgoingWebhookNewParamsTypeSubscriptionCancelled OutgoingWebhookNewParamsType = "subscription.cancelled"
	OutgoingWebhookNewParamsTypeSubscriptionFailed    OutgoingWebhookNewParamsType = "subscription.failed"
	OutgoingWebhookNewParamsTypeSubscriptionExpired   OutgoingWebhookNewParamsType = "subscription.expired"
)

func (r OutgoingWebhookNewParamsType) IsKnown() bool {
	switch r {
	case OutgoingWebhookNewParamsTypePaymentSucceeded, OutgoingWebhookNewParamsTypePaymentFailed, OutgoingWebhookNewParamsTypePaymentProcessing, OutgoingWebhookNewParamsTypePaymentCancelled, OutgoingWebhookNewParamsTypeRefundSucceeded, OutgoingWebhookNewParamsTypeRefundFailed, OutgoingWebhookNewParamsTypeDisputeOpened, OutgoingWebhookNewParamsTypeDisputeExpired, OutgoingWebhookNewParamsTypeDisputeAccepted, OutgoingWebhookNewParamsTypeDisputeCancelled, OutgoingWebhookNewParamsTypeDisputeChallenged, OutgoingWebhookNewParamsTypeDisputeWon, OutgoingWebhookNewParamsTypeDisputeLost, OutgoingWebhookNewParamsTypeSubscriptionActive, OutgoingWebhookNewParamsTypeSubscriptionOnHold, OutgoingWebhookNewParamsTypeSubscriptionPaused, OutgoingWebhookNewParamsTypeSubscriptionCancelled, OutgoingWebhookNewParamsTypeSubscriptionFailed, OutgoingWebhookNewParamsTypeSubscriptionExpired:
		return true
	}
	return false
}
