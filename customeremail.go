// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/apiquery"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/dodopayments/dodopayments-go/packages/pagination"
)

// CustomerEmailService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerEmailService] method instead.
type CustomerEmailService struct {
	Options []option.RequestOption
}

// NewCustomerEmailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerEmailService(opts ...option.RequestOption) (r *CustomerEmailService) {
	r = &CustomerEmailService{}
	r.Options = opts
	return
}

// Returns every transactional email sent to this customer in the last 180 days,
// newest first, with its delivery outcome. Delivery status comes from the email
// provider and is as fresh as replication, typically seconds.
func (r *CustomerEmailService) List(ctx context.Context, customerID string, query CustomerEmailListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[EmailLogItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/%s/emails", customerID)
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

// Returns every transactional email sent to this customer in the last 180 days,
// newest first, with its delivery outcome. Delivery status comes from the email
// provider and is as fresh as replication, typically seconds.
func (r *CustomerEmailService) ListAutoPaging(ctx context.Context, customerID string, query CustomerEmailListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[EmailLogItem] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, customerID, query, opts...))
}

// Returns the email exactly as it was sent, plus the reason it failed when it did.
// Some emails have no body to show: an authentication email carries a live login
// token, a blocked email never reached the provider, and the provider clears
// bodies at 180 days.
func (r *CustomerEmailService) GetBody(ctx context.Context, customerID string, emailLogID string, opts ...option.RequestOption) (res *EmailBody, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return nil, err
	}
	if emailLogID == "" {
		err = errors.New("missing required email_log_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/%s/emails/%s/body", customerID, emailLogID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type EmailBody struct {
	// Whether the merchant wrote this content. It is true for the recovery and dunning
	// emails, which the merchant writes.
	//
	// The content is email HTML. Render it in a sandbox, whatever this value is.
	MerchantAuthored bool `json:"merchant_authored" api:"required"`
	// Why the email did not arrive. It is null unless the email failed.
	FailureCode EmailFailureCode `json:"failure_code" api:"nullable"`
	// A sentence that explains `failure_code`. It is null unless the email failed.
	FailureReason string `json:"failure_reason" api:"nullable"`
	// The stored HTML. It is null on a text-only email.
	HTML string `json:"html" api:"nullable"`
	// The stored plain text.
	Text string        `json:"text" api:"nullable"`
	JSON emailBodyJSON `json:"-"`
}

// emailBodyJSON contains the JSON metadata for the struct [EmailBody]
type emailBodyJSON struct {
	MerchantAuthored apijson.Field
	FailureCode      apijson.Field
	FailureReason    apijson.Field
	HTML             apijson.Field
	Text             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EmailBody) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailBodyJSON) RawJSON() string {
	return r.raw
}

// Why an email did not reach the recipient.
//
// The code is stable. `send_failed` is the catch-all: it covers every failure that
// the other codes do not name.
type EmailFailureCode string

const (
	EmailFailureCodeMailboxNotFound   EmailFailureCode = "mailbox_not_found"
	EmailFailureCodeAddressRejected   EmailFailureCode = "address_rejected"
	EmailFailureCodeAddressSuppressed EmailFailureCode = "address_suppressed"
	EmailFailureCodeMailboxFull       EmailFailureCode = "mailbox_full"
	EmailFailureCodeTemporaryFailure  EmailFailureCode = "temporary_failure"
	EmailFailureCodeMessageTooLarge   EmailFailureCode = "message_too_large"
	EmailFailureCodeMarkedAsSpam      EmailFailureCode = "marked_as_spam"
	EmailFailureCodeSendFailed        EmailFailureCode = "send_failed"
)

func (r EmailFailureCode) IsKnown() bool {
	switch r {
	case EmailFailureCodeMailboxNotFound, EmailFailureCodeAddressRejected, EmailFailureCodeAddressSuppressed, EmailFailureCodeMailboxFull, EmailFailureCodeTemporaryFailure, EmailFailureCodeMessageTooLarge, EmailFailureCodeMarkedAsSpam, EmailFailureCodeSendFailed:
		return true
	}
	return false
}

type EmailLogItem struct {
	// The group this email belongs to: payments, refunds, subscriptions,
	// dunning_recovery, entitlements or auth.
	Category string `json:"category" api:"required"`
	// When this email was sent.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Identifies this email. Use it to read the body or to send it again.
	EmailLogID string `json:"email_log_id" api:"required"`
	// What kind of email this is, for example `payment_successful`.
	EmailType string `json:"email_type" api:"required"`
	// Whether this email has content to show. The content endpoint can still refuse,
	// because the content is removed after 180 days.
	HasPreview bool `json:"has_preview" api:"required"`
	// What you may do with this email.
	Policies EmailPolicies `json:"policies" api:"required"`
	// Where the email got to: sent, delivered, failed, complained or blocked.
	Status EmailLogStatus `json:"status" api:"required"`
	// Why the email did not arrive. It is null unless the email failed.
	FailureCode EmailFailureCode `json:"failure_code" api:"nullable"`
	// A sentence that explains `failure_code`. It is null unless the email failed.
	FailureReason string `json:"failure_reason" api:"nullable"`
	// The address the email was sent from.
	From string `json:"from" api:"nullable"`
	// What the merchant typed, when test mode redirected the send to the business
	// owner.
	IntendedRecipient string `json:"intended_recipient" api:"nullable"`
	// The address the email reached.
	Recipient string `json:"recipient" api:"nullable"`
	// The subject line as it was sent. Empty until the provider replicates.
	Subject string           `json:"subject" api:"nullable"`
	JSON    emailLogItemJSON `json:"-"`
}

// emailLogItemJSON contains the JSON metadata for the struct [EmailLogItem]
type emailLogItemJSON struct {
	Category          apijson.Field
	CreatedAt         apijson.Field
	EmailLogID        apijson.Field
	EmailType         apijson.Field
	HasPreview        apijson.Field
	Policies          apijson.Field
	Status            apijson.Field
	FailureCode       apijson.Field
	FailureReason     apijson.Field
	From              apijson.Field
	IntendedRecipient apijson.Field
	Recipient         apijson.Field
	Subject           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EmailLogItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailLogItemJSON) RawJSON() string {
	return r.raw
}

// The delivery status of one email.
//
// `sent` also covers an email that is still on its way. A status only becomes
// `delivered`, `failed` or `complained` when the mail server answers.
type EmailLogStatus string

const (
	EmailLogStatusSent       EmailLogStatus = "sent"
	EmailLogStatusDelivered  EmailLogStatus = "delivered"
	EmailLogStatusFailed     EmailLogStatus = "failed"
	EmailLogStatusComplained EmailLogStatus = "complained"
	EmailLogStatusBlocked    EmailLogStatus = "blocked"
)

func (r EmailLogStatus) IsKnown() bool {
	switch r {
	case EmailLogStatusSent, EmailLogStatusDelivered, EmailLogStatusFailed, EmailLogStatusComplained, EmailLogStatusBlocked:
		return true
	}
	return false
}

// What the merchant may do with one row. The server decides; the client never
// derives eligibility itself.
type EmailPolicies struct {
	// A permanent failure was recorded, so the same address would be a no-op.
	RequiresDifferentAddress bool `json:"requires_different_address" api:"required"`
	// The row was delivered and may be sent again.
	ResendAllowed bool `json:"resend_allowed" api:"required"`
	// How many sends are left in this email's chain.
	ResendsRemaining int64 `json:"resends_remaining" api:"required"`
	// The row failed and may be sent again.
	RetryAllowed bool              `json:"retry_allowed" api:"required"`
	JSON         emailPoliciesJSON `json:"-"`
}

// emailPoliciesJSON contains the JSON metadata for the struct [EmailPolicies]
type emailPoliciesJSON struct {
	RequiresDifferentAddress apijson.Field
	ResendAllowed            apijson.Field
	ResendsRemaining         apijson.Field
	RetryAllowed             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *EmailPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailPoliciesJSON) RawJSON() string {
	return r.raw
}

type CustomerEmailListParams struct {
	// Which page to return. The default is 0.
	PageNumber param.Field[int64] `query:"page_number"`
	// How many emails to return. The default is 10 and the maximum is 100.
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [CustomerEmailListParams]'s query parameters as
// `url.Values`.
func (r CustomerEmailListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
