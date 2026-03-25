// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
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
	WebhookEventTypeCreditAdded             WebhookEventType = "credit.added"
	WebhookEventTypeCreditDeducted          WebhookEventType = "credit.deducted"
	WebhookEventTypeCreditExpired           WebhookEventType = "credit.expired"
	WebhookEventTypeCreditRolledOver        WebhookEventType = "credit.rolled_over"
	WebhookEventTypeCreditRolloverForfeited WebhookEventType = "credit.rollover_forfeited"
	WebhookEventTypeCreditOverageCharged    WebhookEventType = "credit.overage_charged"
	WebhookEventTypeCreditManualAdjustment  WebhookEventType = "credit.manual_adjustment"
	WebhookEventTypeCreditBalanceLow        WebhookEventType = "credit.balance_low"
)

func (r WebhookEventType) IsKnown() bool {
	switch r {
	case WebhookEventTypePaymentSucceeded, WebhookEventTypePaymentFailed, WebhookEventTypePaymentProcessing, WebhookEventTypePaymentCancelled, WebhookEventTypeRefundSucceeded, WebhookEventTypeRefundFailed, WebhookEventTypeDisputeOpened, WebhookEventTypeDisputeExpired, WebhookEventTypeDisputeAccepted, WebhookEventTypeDisputeCancelled, WebhookEventTypeDisputeChallenged, WebhookEventTypeDisputeWon, WebhookEventTypeDisputeLost, WebhookEventTypeSubscriptionActive, WebhookEventTypeSubscriptionRenewed, WebhookEventTypeSubscriptionOnHold, WebhookEventTypeSubscriptionCancelled, WebhookEventTypeSubscriptionFailed, WebhookEventTypeSubscriptionExpired, WebhookEventTypeSubscriptionPlanChanged, WebhookEventTypeSubscriptionUpdated, WebhookEventTypeLicenseKeyCreated, WebhookEventTypePayoutNotInitiated, WebhookEventTypePayoutOnHold, WebhookEventTypePayoutInProgress, WebhookEventTypePayoutFailed, WebhookEventTypePayoutSuccess, WebhookEventTypeCreditAdded, WebhookEventTypeCreditDeducted, WebhookEventTypeCreditExpired, WebhookEventTypeCreditRolledOver, WebhookEventTypeCreditRolloverForfeited, WebhookEventTypeCreditOverageCharged, WebhookEventTypeCreditManualAdjustment, WebhookEventTypeCreditBalanceLow:
		return true
	}
	return false
}

type WebhookPayloadParam struct {
	BusinessID param.Field[string] `json:"business_id" api:"required"`
	// The latest data at the time of delivery attempt
	Data param.Field[WebhookPayloadDataUnionParam] `json:"data" api:"required"`
	// The timestamp of when the event occurred (not necessarily the same of when it
	// was delivered)
	Timestamp param.Field[time.Time] `json:"timestamp" api:"required" format:"date-time"`
	// Event types for Dodo events
	Type param.Field[WebhookEventType] `json:"type" api:"required"`
}

func (r WebhookPayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The latest data at the time of delivery attempt
//
// Satisfied by [WebhookPayloadDataPaymentParam],
// [WebhookPayloadDataSubscriptionParam], [WebhookPayloadDataRefundParam],
// [WebhookPayloadDataDisputeParam], [WebhookPayloadDataLicenseKeyParam],
// [WebhookPayloadDataCreditLedgerEntryParam],
// [WebhookPayloadDataCreditBalanceLowParam].
type WebhookPayloadDataUnionParam interface {
	implementsWebhookPayloadDataUnionParam()
}

type WebhookPayloadDataPaymentParam struct {
	PayloadType param.Field[WebhookPayloadDataPaymentPayloadType] `json:"payload_type" api:"required"`
	PaymentParam
}

func (r WebhookPayloadDataPaymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookPayloadDataPaymentParam) implementsWebhookPayloadDataUnionParam() {}

type WebhookPayloadDataPaymentPayloadType string

const (
	WebhookPayloadDataPaymentPayloadTypePayment WebhookPayloadDataPaymentPayloadType = "Payment"
)

func (r WebhookPayloadDataPaymentPayloadType) IsKnown() bool {
	switch r {
	case WebhookPayloadDataPaymentPayloadTypePayment:
		return true
	}
	return false
}

// Response struct representing subscription details
type WebhookPayloadDataSubscriptionParam struct {
	PayloadType param.Field[WebhookPayloadDataSubscriptionPayloadType] `json:"payload_type" api:"required"`
	SubscriptionParam
}

func (r WebhookPayloadDataSubscriptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookPayloadDataSubscriptionParam) implementsWebhookPayloadDataUnionParam() {}

type WebhookPayloadDataSubscriptionPayloadType string

const (
	WebhookPayloadDataSubscriptionPayloadTypeSubscription WebhookPayloadDataSubscriptionPayloadType = "Subscription"
)

func (r WebhookPayloadDataSubscriptionPayloadType) IsKnown() bool {
	switch r {
	case WebhookPayloadDataSubscriptionPayloadTypeSubscription:
		return true
	}
	return false
}

type WebhookPayloadDataRefundParam struct {
	PayloadType param.Field[WebhookPayloadDataRefundPayloadType] `json:"payload_type" api:"required"`
	RefundParam
}

func (r WebhookPayloadDataRefundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookPayloadDataRefundParam) implementsWebhookPayloadDataUnionParam() {}

type WebhookPayloadDataRefundPayloadType string

const (
	WebhookPayloadDataRefundPayloadTypeRefund WebhookPayloadDataRefundPayloadType = "Refund"
)

func (r WebhookPayloadDataRefundPayloadType) IsKnown() bool {
	switch r {
	case WebhookPayloadDataRefundPayloadTypeRefund:
		return true
	}
	return false
}

type WebhookPayloadDataDisputeParam struct {
	PayloadType param.Field[WebhookPayloadDataDisputePayloadType] `json:"payload_type" api:"required"`
	GetDisputeParam
}

func (r WebhookPayloadDataDisputeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookPayloadDataDisputeParam) implementsWebhookPayloadDataUnionParam() {}

type WebhookPayloadDataDisputePayloadType string

const (
	WebhookPayloadDataDisputePayloadTypeDispute WebhookPayloadDataDisputePayloadType = "Dispute"
)

func (r WebhookPayloadDataDisputePayloadType) IsKnown() bool {
	switch r {
	case WebhookPayloadDataDisputePayloadTypeDispute:
		return true
	}
	return false
}

type WebhookPayloadDataLicenseKeyParam struct {
	PayloadType param.Field[WebhookPayloadDataLicenseKeyPayloadType] `json:"payload_type" api:"required"`
	LicenseKeyParam
}

func (r WebhookPayloadDataLicenseKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookPayloadDataLicenseKeyParam) implementsWebhookPayloadDataUnionParam() {}

type WebhookPayloadDataLicenseKeyPayloadType string

const (
	WebhookPayloadDataLicenseKeyPayloadTypeLicenseKey WebhookPayloadDataLicenseKeyPayloadType = "LicenseKey"
)

func (r WebhookPayloadDataLicenseKeyPayloadType) IsKnown() bool {
	switch r {
	case WebhookPayloadDataLicenseKeyPayloadTypeLicenseKey:
		return true
	}
	return false
}

// Response for a ledger entry
type WebhookPayloadDataCreditLedgerEntryParam struct {
	PayloadType param.Field[WebhookPayloadDataCreditLedgerEntryPayloadType] `json:"payload_type" api:"required"`
	CreditLedgerEntryParam
}

func (r WebhookPayloadDataCreditLedgerEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookPayloadDataCreditLedgerEntryParam) implementsWebhookPayloadDataUnionParam() {}

type WebhookPayloadDataCreditLedgerEntryPayloadType string

const (
	WebhookPayloadDataCreditLedgerEntryPayloadTypeCreditLedgerEntry WebhookPayloadDataCreditLedgerEntryPayloadType = "CreditLedgerEntry"
)

func (r WebhookPayloadDataCreditLedgerEntryPayloadType) IsKnown() bool {
	switch r {
	case WebhookPayloadDataCreditLedgerEntryPayloadTypeCreditLedgerEntry:
		return true
	}
	return false
}

type WebhookPayloadDataCreditBalanceLowParam struct {
	AvailableBalance          param.Field[string]                                        `json:"available_balance" api:"required"`
	CreditEntitlementID       param.Field[string]                                        `json:"credit_entitlement_id" api:"required"`
	CreditEntitlementName     param.Field[string]                                        `json:"credit_entitlement_name" api:"required"`
	CustomerID                param.Field[string]                                        `json:"customer_id" api:"required"`
	PayloadType               param.Field[WebhookPayloadDataCreditBalanceLowPayloadType] `json:"payload_type" api:"required"`
	SubscriptionCreditsAmount param.Field[string]                                        `json:"subscription_credits_amount" api:"required"`
	SubscriptionID            param.Field[string]                                        `json:"subscription_id" api:"required"`
	ThresholdAmount           param.Field[string]                                        `json:"threshold_amount" api:"required"`
	ThresholdPercent          param.Field[int64]                                         `json:"threshold_percent" api:"required"`
}

func (r WebhookPayloadDataCreditBalanceLowParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebhookPayloadDataCreditBalanceLowParam) implementsWebhookPayloadDataUnionParam() {}

type WebhookPayloadDataCreditBalanceLowPayloadType string

const (
	WebhookPayloadDataCreditBalanceLowPayloadTypeCreditBalanceLow WebhookPayloadDataCreditBalanceLowPayloadType = "CreditBalanceLow"
)

func (r WebhookPayloadDataCreditBalanceLowPayloadType) IsKnown() bool {
	switch r {
	case WebhookPayloadDataCreditBalanceLowPayloadTypeCreditBalanceLow:
		return true
	}
	return false
}
