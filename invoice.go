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
	"github.com/tidwall/gjson"
)

// InvoiceService contains methods and other services that help with interacting
// with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options  []option.RequestOption
	Payments *InvoicePaymentService
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	r.Payments = NewInvoicePaymentService(opts...)
	return
}

// Creates a draft. A repeat with the same `Idempotency-Key` returns the first
// draft.
func (r *InvoiceService) New(ctx context.Context, params InvoiceNewParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if params.IdempotencyKey.Present {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "billing/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get an invoice
func (r *InvoiceService) Get(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Changes only the sent fields. `null` clears a field that can be empty.
func (r *InvoiceService) Update(ctx context.Context, invoiceID string, params InvoiceUpdateParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists invoice-led invoices, newest first. The items have no `lines` and no
// `payments`.
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[Invoice], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "billing/invoices"
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

// Lists invoice-led invoices, newest first. The items have no `lines` and no
// `payments`.
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[Invoice] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a draft invoice
func (r *InvoiceService) Delete(ctx context.Context, invoiceID string, body InvoiceDeleteParams, opts ...option.RequestOption) (err error) {
	if body.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", body.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return err
	}
	path := fmt.Sprintf("billing/invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A note changes no field, so it needs no `If-Match`. The note shows in `/events`.
func (r *InvoiceService) AddNote(ctx context.Context, invoiceID string, body InvoiceAddNoteParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/notes", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the page where the customer pays. A copied link counts as a send, so the
// first call sets `sent_at`.
func (r *InvoiceService) NewHostedLink(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *InvoiceNewHostedLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/hosted-link", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Copies the customer, the product lines, the terms and the recipients into a new
// draft. Usage lines are not copied.
func (r *InvoiceService) Duplicate(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/duplicate", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Issues the draft. With `send: true`, the invoice email goes out after the issue.
// A failed send keeps the invoice issued. Call `/send` again with the new `ETag`.
func (r *InvoiceService) Finalize(ctx context.Context, invoiceID string, params InvoiceFinalizeParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/finalize", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the change history of the invoice, oldest first.
func (r *InvoiceService) ListEvents(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *InvoiceListEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/events", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Mark an invoice uncollectable
func (r *InvoiceService) MarkUncollectable(ctx context.Context, invoiceID string, body InvoiceMarkUncollectableParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if body.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", body.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/mark-uncollectable", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Prices the body as a draft and returns it. No row is saved.
func (r *InvoiceService) Preview(ctx context.Context, params InvoicePreviewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "billing/invoices/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Preview a draft invoice
func (r *InvoiceService) PreviewExisting(ctx context.Context, invoiceID string, query InvoicePreviewExistingParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return err
	}
	path := fmt.Sprintf("billing/invoices/%s/preview", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

// Charges the saved payment method of the invoice subscription now.
func (r *InvoiceService) Retry(ctx context.Context, invoiceID string, body InvoiceRetryParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if body.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", body.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/retry", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Emails the invoice to the customer email and the CC emails.
func (r *InvoiceService) Send(ctx context.Context, invoiceID string, body InvoiceSendParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if body.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", body.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/send", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Pauses, resumes or moves the reminders of an open or past-due invoice.
func (r *InvoiceService) UpdateDunning(ctx context.Context, invoiceID string, params InvoiceUpdateDunningParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if params.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/dunning", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Closes an open or past-due invoice. Applied credits go back to the customer
// wallet.
func (r *InvoiceService) Void(ctx context.Context, invoiceID string, body InvoiceVoidParams, opts ...option.RequestOption) (res *Invoice, err error) {
	if body.IfMatch.Present {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", body.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoice_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("billing/invoices/%s/void", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type CollectionMethod string

const (
	CollectionMethodAutoCharge CollectionMethod = "auto_charge"
	CollectionMethodManual     CollectionMethod = "manual"
)

func (r CollectionMethod) IsKnown() bool {
	switch r {
	case CollectionMethodAutoCharge, CollectionMethodManual:
		return true
	}
	return false
}

type DunningState string

const (
	DunningStateNotStarted DunningState = "not_started"
	DunningStateActive     DunningState = "active"
	DunningStatePaused     DunningState = "paused"
	DunningStateExhausted  DunningState = "exhausted"
)

func (r DunningState) IsKnown() bool {
	switch r {
	case DunningStateNotStarted, DunningStateActive, DunningStatePaused, DunningStateExhausted:
		return true
	}
	return false
}

// An invoice-led invoice. Every read and every `invoice.*` webhook uses this
// schema.
type Invoice struct {
	// The actions that the current status accepts.
	AllowedActions []InvoiceAllowedAction `json:"allowed_actions" api:"required"`
	// The amount that the customer must pay. Zero unless the invoice is open or past
	// due.
	AmountDue int64 `json:"amount_due" api:"required"`
	// The sum of the succeeded payments, in the smallest currency unit. A late payment
	// is not included after its late refund.
	AmountPaid int64 `json:"amount_paid" api:"required"`
	// The sum of the succeeded refunds, in the currency of each payment. The late
	// refund of a late payment is not included.
	AmountRefunded int64              `json:"amount_refunded" api:"required"`
	BillingMode    InvoiceBillingMode `json:"billing_mode" api:"required"`
	BusinessID     string             `json:"business_id" api:"required"`
	CreatedAt      time.Time          `json:"created_at" api:"required" format:"date-time"`
	// Customer credits applied at finalise, in the smallest currency unit. Zero or
	// less.
	Credits  int64           `json:"credits" api:"required"`
	Currency Currency        `json:"currency" api:"required"`
	Customer InvoiceCustomer `json:"customer" api:"required"`
	// The discount codes of the invoice.
	DiscountIDs []string `json:"discount_ids" api:"required"`
	// The fields that `PATCH` can change in the current status.
	EditableFields []InvoiceEditableField `json:"editable_fields" api:"required"`
	// How often the invoice lines bill. `mixed` covers several intervals, or one-time
	// lines next to a plan.
	Frequency InvoiceFrequency `json:"frequency" api:"required"`
	InvoiceID string           `json:"invoice_id" api:"required"`
	Metadata  interface{}      `json:"metadata" api:"required"`
	RefundIDs []string         `json:"refund_ids" api:"required"`
	// The number of failed automatic charges.
	RetryCount int64 `json:"retry_count" api:"required"`
	// Goes up by one on each change. Use it to ignore an older webhook delivery.
	RowVersion int64 `json:"row_version" api:"required"`
	// The `message` of each variant is the label of the invoice document.
	Status          InvoiceStatus `json:"status" api:"required"`
	SubscriptionIDs []string      `json:"subscription_ids" api:"required"`
	Tags            []InvoiceTag  `json:"tags" api:"required"`
	// Tax in the smallest currency unit.
	Tax int64 `json:"tax" api:"required"`
	// Total with tax, before credits, in the smallest currency unit.
	TotalAmount      int64            `json:"total_amount" api:"required"`
	CollectionMethod CollectionMethod `json:"collection_method" api:"nullable"`
	DueAt            time.Time        `json:"due_at" api:"nullable" format:"date-time"`
	// `null` for an invoice that has no collection: draft, paid, void or
	// uncollectable.
	DunningState DunningState `json:"dunning_state" api:"nullable"`
	FinalizedAt  time.Time    `json:"finalized_at" api:"nullable" format:"date-time"`
	// The page where the customer pays. `null` until the invoice is sent.
	HostedURL string `json:"hosted_url" api:"nullable"`
	// Only on a single invoice. A list leaves it out.
	Lines []InvoiceLine `json:"lines" api:"nullable"`
	Memo  string        `json:"memo" api:"nullable"`
	// The time of the next reminder email to the customer, before or after the due
	// date. `null` when no reminder email will go out.
	NextReminderAt time.Time `json:"next_reminder_at" api:"nullable" format:"date-time"`
	NextRetryAt    time.Time `json:"next_retry_at" api:"nullable" format:"date-time"`
	// Days from finalise to the due date. `null` when the invoice has a custom due
	// date.
	PaymentTermsDays int64 `json:"payment_terms_days" api:"nullable"`
	// Only on a single invoice. A list leaves it out.
	Payments []InvoicePayment `json:"payments" api:"nullable"`
	PdfURL   string           `json:"pdf_url" api:"nullable"`
	// `null` on a draft that keeps all reminder settings of the business.
	ReminderSettings InvoiceReminderSettings `json:"reminder_settings" api:"nullable"`
	// The time at which a draft that waits for review finalises.
	ReviewDeadline time.Time   `json:"review_deadline" api:"nullable" format:"date-time"`
	SentAt         time.Time   `json:"sent_at" api:"nullable" format:"date-time"`
	JSON           invoiceJSON `json:"-"`
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	AllowedActions   apijson.Field
	AmountDue        apijson.Field
	AmountPaid       apijson.Field
	AmountRefunded   apijson.Field
	BillingMode      apijson.Field
	BusinessID       apijson.Field
	CreatedAt        apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	Customer         apijson.Field
	DiscountIDs      apijson.Field
	EditableFields   apijson.Field
	Frequency        apijson.Field
	InvoiceID        apijson.Field
	Metadata         apijson.Field
	RefundIDs        apijson.Field
	RetryCount       apijson.Field
	RowVersion       apijson.Field
	Status           apijson.Field
	SubscriptionIDs  apijson.Field
	Tags             apijson.Field
	Tax              apijson.Field
	TotalAmount      apijson.Field
	CollectionMethod apijson.Field
	DueAt            apijson.Field
	DunningState     apijson.Field
	FinalizedAt      apijson.Field
	HostedURL        apijson.Field
	Lines            apijson.Field
	Memo             apijson.Field
	NextReminderAt   apijson.Field
	NextRetryAt      apijson.Field
	PaymentTermsDays apijson.Field
	Payments         apijson.Field
	PdfURL           apijson.Field
	ReminderSettings apijson.Field
	ReviewDeadline   apijson.Field
	SentAt           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceJSON) RawJSON() string {
	return r.raw
}

// An action that a merchant can request on the invoice in its current status.
type InvoiceAllowedAction string

const (
	InvoiceAllowedActionEdit              InvoiceAllowedAction = "edit"
	InvoiceAllowedActionFinalize          InvoiceAllowedAction = "finalize"
	InvoiceAllowedActionDelete            InvoiceAllowedAction = "delete"
	InvoiceAllowedActionSend              InvoiceAllowedAction = "send"
	InvoiceAllowedActionRetry             InvoiceAllowedAction = "retry"
	InvoiceAllowedActionVoid              InvoiceAllowedAction = "void"
	InvoiceAllowedActionMarkUncollectable InvoiceAllowedAction = "mark_uncollectable"
	InvoiceAllowedActionChangeDunning     InvoiceAllowedAction = "change_dunning"
	InvoiceAllowedActionAddNote           InvoiceAllowedAction = "add_note"
	InvoiceAllowedActionDuplicate         InvoiceAllowedAction = "duplicate"
)

func (r InvoiceAllowedAction) IsKnown() bool {
	switch r {
	case InvoiceAllowedActionEdit, InvoiceAllowedActionFinalize, InvoiceAllowedActionDelete, InvoiceAllowedActionSend, InvoiceAllowedActionRetry, InvoiceAllowedActionVoid, InvoiceAllowedActionMarkUncollectable, InvoiceAllowedActionChangeDunning, InvoiceAllowedActionAddNote, InvoiceAllowedActionDuplicate:
		return true
	}
	return false
}

type InvoiceBillingMode string

const (
	InvoiceBillingModeNormal     InvoiceBillingMode = "normal"
	InvoiceBillingModeInvoiceLed InvoiceBillingMode = "invoice_led"
)

func (r InvoiceBillingMode) IsKnown() bool {
	switch r {
	case InvoiceBillingModeNormal, InvoiceBillingModeInvoiceLed:
		return true
	}
	return false
}

type InvoiceCustomer struct {
	// A draft can have no country yet. Finalise needs it.
	BillingAddress InvoiceCustomerBillingAddress `json:"billing_address" api:"required"`
	CcEmails       []string                      `json:"cc_emails" api:"required"`
	CustomerID     string                        `json:"customer_id" api:"required"`
	Email          string                        `json:"email" api:"nullable"`
	TaxID          string                        `json:"tax_id" api:"nullable"`
	JSON           invoiceCustomerJSON           `json:"-"`
}

// invoiceCustomerJSON contains the JSON metadata for the struct [InvoiceCustomer]
type invoiceCustomerJSON struct {
	BillingAddress apijson.Field
	CcEmails       apijson.Field
	CustomerID     apijson.Field
	Email          apijson.Field
	TaxID          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerJSON) RawJSON() string {
	return r.raw
}

// A draft can have no country yet. Finalise needs it.
type InvoiceCustomerBillingAddress struct {
	City string `json:"city" api:"nullable"`
	// ISO country code alpha2 variant
	Country CountryCode                       `json:"country" api:"nullable"`
	State   string                            `json:"state" api:"nullable"`
	Street  string                            `json:"street" api:"nullable"`
	Zipcode string                            `json:"zipcode" api:"nullable"`
	JSON    invoiceCustomerBillingAddressJSON `json:"-"`
}

// invoiceCustomerBillingAddressJSON contains the JSON metadata for the struct
// [InvoiceCustomerBillingAddress]
type invoiceCustomerBillingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zipcode     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCustomerBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCustomerBillingAddressJSON) RawJSON() string {
	return r.raw
}

type InvoiceEditableField string

const (
	InvoiceEditableFieldCustomer         InvoiceEditableField = "customer"
	InvoiceEditableFieldItems            InvoiceEditableField = "items"
	InvoiceEditableFieldDiscounts        InvoiceEditableField = "discounts"
	InvoiceEditableFieldCurrency         InvoiceEditableField = "currency"
	InvoiceEditableFieldBillingAddress   InvoiceEditableField = "billing_address"
	InvoiceEditableFieldTaxID            InvoiceEditableField = "tax_id"
	InvoiceEditableFieldReviewRequired   InvoiceEditableField = "review_required"
	InvoiceEditableFieldDueAt            InvoiceEditableField = "due_at"
	InvoiceEditableFieldCollectionMethod InvoiceEditableField = "collection_method"
	InvoiceEditableFieldMemo             InvoiceEditableField = "memo"
	InvoiceEditableFieldCustomerEmail    InvoiceEditableField = "customer_email"
	InvoiceEditableFieldCcEmails         InvoiceEditableField = "cc_emails"
	InvoiceEditableFieldDunning          InvoiceEditableField = "dunning"
	InvoiceEditableFieldMetadata         InvoiceEditableField = "metadata"
)

func (r InvoiceEditableField) IsKnown() bool {
	switch r {
	case InvoiceEditableFieldCustomer, InvoiceEditableFieldItems, InvoiceEditableFieldDiscounts, InvoiceEditableFieldCurrency, InvoiceEditableFieldBillingAddress, InvoiceEditableFieldTaxID, InvoiceEditableFieldReviewRequired, InvoiceEditableFieldDueAt, InvoiceEditableFieldCollectionMethod, InvoiceEditableFieldMemo, InvoiceEditableFieldCustomerEmail, InvoiceEditableFieldCcEmails, InvoiceEditableFieldDunning, InvoiceEditableFieldMetadata:
		return true
	}
	return false
}

// How often the invoice lines bill. `mixed` covers several intervals, or one-time
// lines next to a plan.
type InvoiceFrequency struct {
	Type  InvoiceFrequencyType `json:"type" api:"required"`
	Count int64                `json:"count"`
	// Unit of a duration count (e.g. license-key validity period).
	Interval TimeInterval         `json:"interval"`
	JSON     invoiceFrequencyJSON `json:"-"`
	union    InvoiceFrequencyUnion
}

// invoiceFrequencyJSON contains the JSON metadata for the struct
// [InvoiceFrequency]
type invoiceFrequencyJSON struct {
	Type        apijson.Field
	Count       apijson.Field
	Interval    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r invoiceFrequencyJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceFrequency) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceFrequency{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceFrequencyUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [InvoiceFrequencyType],
// [InvoiceFrequencyObject], [InvoiceFrequencyType].
func (r InvoiceFrequency) AsUnion() InvoiceFrequencyUnion {
	return r.union
}

// How often the invoice lines bill. `mixed` covers several intervals, or one-time
// lines next to a plan.
//
// Union satisfied by [InvoiceFrequencyType], [InvoiceFrequencyObject] or
// [InvoiceFrequencyType].
type InvoiceFrequencyUnion interface {
	implementsInvoiceFrequency()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceFrequencyUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceFrequencyType{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceFrequencyObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceFrequencyType{}),
		},
	)
}

type InvoiceFrequencyType struct {
	Type InvoiceFrequencyTypeType `json:"type" api:"required"`
	JSON invoiceFrequencyTypeJSON `json:"-"`
}

// invoiceFrequencyTypeJSON contains the JSON metadata for the struct
// [InvoiceFrequencyType]
type invoiceFrequencyTypeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFrequencyType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFrequencyTypeJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFrequencyType) implementsInvoiceFrequency() {}

type InvoiceFrequencyTypeType string

const (
	InvoiceFrequencyTypeTypeOneTime InvoiceFrequencyTypeType = "one_time"
)

func (r InvoiceFrequencyTypeType) IsKnown() bool {
	switch r {
	case InvoiceFrequencyTypeTypeOneTime:
		return true
	}
	return false
}

type InvoiceFrequencyObject struct {
	Count int64 `json:"count" api:"required"`
	// Unit of a duration count (e.g. license-key validity period).
	Interval TimeInterval               `json:"interval" api:"required"`
	Type     InvoiceFrequencyObjectType `json:"type" api:"required"`
	JSON     invoiceFrequencyObjectJSON `json:"-"`
}

// invoiceFrequencyObjectJSON contains the JSON metadata for the struct
// [InvoiceFrequencyObject]
type invoiceFrequencyObjectJSON struct {
	Count       apijson.Field
	Interval    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFrequencyObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceFrequencyObjectJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceFrequencyObject) implementsInvoiceFrequency() {}

type InvoiceFrequencyObjectType string

const (
	InvoiceFrequencyObjectTypeRecurring InvoiceFrequencyObjectType = "recurring"
)

func (r InvoiceFrequencyObjectType) IsKnown() bool {
	switch r {
	case InvoiceFrequencyObjectTypeRecurring:
		return true
	}
	return false
}

type InvoiceTag string

const (
	InvoiceTagRefunded          InvoiceTag = "refunded"
	InvoiceTagPartiallyRefunded InvoiceTag = "partially_refunded"
	InvoiceTagDisputed          InvoiceTag = "disputed"
	InvoiceTagPaymentReversed   InvoiceTag = "payment_reversed"
)

func (r InvoiceTag) IsKnown() bool {
	switch r {
	case InvoiceTagRefunded, InvoiceTagPartiallyRefunded, InvoiceTagDisputed, InvoiceTagPaymentReversed:
		return true
	}
	return false
}

type InvoicePayment struct {
	Amount    int64     `json:"amount" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	PaymentID string    `json:"payment_id" api:"required"`
	// The automatic charge number. `null` for a payment that the customer started.
	RetryAttempt int64 `json:"retry_attempt" api:"nullable"`
	// `null` while the processor has not answered.
	Status IntentStatus       `json:"status" api:"nullable"`
	JSON   invoicePaymentJSON `json:"-"`
}

// invoicePaymentJSON contains the JSON metadata for the struct [InvoicePayment]
type invoicePaymentJSON struct {
	Amount       apijson.Field
	CreatedAt    apijson.Field
	PaymentID    apijson.Field
	RetryAttempt apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoicePayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentJSON) RawJSON() string {
	return r.raw
}

// `null` on a draft that keeps all reminder settings of the business.
type InvoiceReminderSettings struct {
	Kind InvoiceReminderSettingsKind `json:"kind" api:"required"`
	// This field can have the runtime type of [[]int64].
	OffsetsMinutes   interface{}                 `json:"offsets_minutes"`
	Paused           bool                        `json:"paused" api:"nullable"`
	RemindersEnabled bool                        `json:"reminders_enabled" api:"nullable"`
	JSON             invoiceReminderSettingsJSON `json:"-"`
	union            InvoiceReminderSettingsUnion
}

// invoiceReminderSettingsJSON contains the JSON metadata for the struct
// [InvoiceReminderSettings]
type invoiceReminderSettingsJSON struct {
	Kind             apijson.Field
	OffsetsMinutes   apijson.Field
	Paused           apijson.Field
	RemindersEnabled apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r invoiceReminderSettingsJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceReminderSettings) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceReminderSettings{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvoiceReminderSettingsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [InvoiceReminderSettingsObject],
// [InvoiceReminderSettingsObject].
func (r InvoiceReminderSettings) AsUnion() InvoiceReminderSettingsUnion {
	return r.union
}

// `null` on a draft that keeps all reminder settings of the business.
//
// Union satisfied by [InvoiceReminderSettingsObject] or
// [InvoiceReminderSettingsObject].
type InvoiceReminderSettingsUnion interface {
	implementsInvoiceReminderSettings()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceReminderSettingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceReminderSettingsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceReminderSettingsObject{}),
		},
	)
}

// The settings that finalise froze. The reminders run on them.
type InvoiceReminderSettingsObject struct {
	Kind             InvoiceReminderSettingsObjectKind `json:"kind" api:"required"`
	OffsetsMinutes   []int64                           `json:"offsets_minutes" api:"required"`
	Paused           bool                              `json:"paused" api:"required"`
	RemindersEnabled bool                              `json:"reminders_enabled" api:"required"`
	JSON             invoiceReminderSettingsObjectJSON `json:"-"`
}

// invoiceReminderSettingsObjectJSON contains the JSON metadata for the struct
// [InvoiceReminderSettingsObject]
type invoiceReminderSettingsObjectJSON struct {
	Kind             apijson.Field
	OffsetsMinutes   apijson.Field
	Paused           apijson.Field
	RemindersEnabled apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvoiceReminderSettingsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceReminderSettingsObjectJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceReminderSettingsObject) implementsInvoiceReminderSettings() {}

type InvoiceReminderSettingsObjectKind string

const (
	InvoiceReminderSettingsObjectKindSnapshot InvoiceReminderSettingsObjectKind = "snapshot"
)

func (r InvoiceReminderSettingsObjectKind) IsKnown() bool {
	switch r {
	case InvoiceReminderSettingsObjectKindSnapshot:
		return true
	}
	return false
}

type InvoiceReminderSettingsKind string

const (
	InvoiceReminderSettingsKindSnapshot InvoiceReminderSettingsKind = "snapshot"
	InvoiceReminderSettingsKindOverride InvoiceReminderSettingsKind = "override"
)

func (r InvoiceReminderSettingsKind) IsKnown() bool {
	switch r {
	case InvoiceReminderSettingsKindSnapshot, InvoiceReminderSettingsKindOverride:
		return true
	}
	return false
}

type InvoiceEvent struct {
	// A user ID, `api_key`, `customer` or `system`.
	Actor      string           `json:"actor" api:"required"`
	CreatedAt  time.Time        `json:"created_at" api:"required" format:"date-time"`
	EventID    string           `json:"event_id" api:"required"`
	EventType  string           `json:"event_type" api:"required"`
	Payload    interface{}      `json:"payload" api:"required"`
	RowVersion int64            `json:"row_version" api:"required"`
	ActorEmail string           `json:"actor_email" api:"nullable"`
	JSON       invoiceEventJSON `json:"-"`
}

// invoiceEventJSON contains the JSON metadata for the struct [InvoiceEvent]
type invoiceEventJSON struct {
	Actor       apijson.Field
	CreatedAt   apijson.Field
	EventID     apijson.Field
	EventType   apijson.Field
	Payload     apijson.Field
	RowVersion  apijson.Field
	ActorEmail  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEventJSON) RawJSON() string {
	return r.raw
}

// One frozen line. Amounts are in the smallest currency unit.
type InvoiceLine struct {
	Amount int64           `json:"amount" api:"required"`
	Kind   InvoiceLineKind `json:"kind" api:"required"`
	// A usage line has a quantity of 1. Its amount is the full usage charge.
	Quantity int64 `json:"quantity" api:"required"`
	Tax      int64 `json:"tax" api:"required"`
	// `true` when `amount` includes `tax`.
	TaxInclusive   bool      `json:"tax_inclusive" api:"required"`
	UnitAmount     int64     `json:"unit_amount" api:"required"`
	AddonID        string    `json:"addon_id" api:"nullable"`
	Description    string    `json:"description" api:"nullable"`
	Name           string    `json:"name" api:"nullable"`
	PeriodEnd      time.Time `json:"period_end" api:"nullable" format:"date-time"`
	PeriodStart    time.Time `json:"period_start" api:"nullable" format:"date-time"`
	ProductID      string    `json:"product_id" api:"nullable"`
	SubscriptionID string    `json:"subscription_id" api:"nullable"`
	// Tax percent times 100. For example, 1800 is 18%.
	TaxRate int64           `json:"tax_rate" api:"nullable"`
	JSON    invoiceLineJSON `json:"-"`
}

// invoiceLineJSON contains the JSON metadata for the struct [InvoiceLine]
type invoiceLineJSON struct {
	Amount         apijson.Field
	Kind           apijson.Field
	Quantity       apijson.Field
	Tax            apijson.Field
	TaxInclusive   apijson.Field
	UnitAmount     apijson.Field
	AddonID        apijson.Field
	Description    apijson.Field
	Name           apijson.Field
	PeriodEnd      apijson.Field
	PeriodStart    apijson.Field
	ProductID      apijson.Field
	SubscriptionID apijson.Field
	TaxRate        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceLineJSON) RawJSON() string {
	return r.raw
}

type InvoiceLineKind string

const (
	InvoiceLineKindSubscription  InvoiceLineKind = "subscription"
	InvoiceLineKindAddon         InvoiceLineKind = "addon"
	InvoiceLineKindMeter         InvoiceLineKind = "meter"
	InvoiceLineKindCreditOverage InvoiceLineKind = "credit_overage"
	InvoiceLineKindOneTime       InvoiceLineKind = "one_time"
)

func (r InvoiceLineKind) IsKnown() bool {
	switch r {
	case InvoiceLineKindSubscription, InvoiceLineKindAddon, InvoiceLineKindMeter, InvoiceLineKindCreditOverage, InvoiceLineKindOneTime:
		return true
	}
	return false
}

// The `message` of each variant is the label of the invoice document.
type InvoiceStatus string

const (
	InvoiceStatusOpen          InvoiceStatus = "open"
	InvoiceStatusPaid          InvoiceStatus = "paid"
	InvoiceStatusVoid          InvoiceStatus = "void"
	InvoiceStatusDraft         InvoiceStatus = "draft"
	InvoiceStatusPastDue       InvoiceStatus = "past_due"
	InvoiceStatusUncollectable InvoiceStatus = "uncollectable"
)

func (r InvoiceStatus) IsKnown() bool {
	switch r {
	case InvoiceStatusOpen, InvoiceStatusPaid, InvoiceStatusVoid, InvoiceStatusDraft, InvoiceStatusPastDue, InvoiceStatusUncollectable:
		return true
	}
	return false
}

type InvoiceNewHostedLinkResponse struct {
	ExpiresAt time.Time                        `json:"expires_at" api:"required" format:"date-time"`
	HostedURL string                           `json:"hosted_url" api:"required"`
	JSON      invoiceNewHostedLinkResponseJSON `json:"-"`
}

// invoiceNewHostedLinkResponseJSON contains the JSON metadata for the struct
// [InvoiceNewHostedLinkResponse]
type invoiceNewHostedLinkResponseJSON struct {
	ExpiresAt   apijson.Field
	HostedURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceNewHostedLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceNewHostedLinkResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceListEventsResponse struct {
	Items []InvoiceEvent                `json:"items" api:"required"`
	JSON  invoiceListEventsResponseJSON `json:"-"`
}

// invoiceListEventsResponseJSON contains the JSON metadata for the struct
// [InvoiceListEventsResponse]
type invoiceListEventsResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceListEventsResponseJSON) RawJSON() string {
	return r.raw
}

type InvoiceNewParams struct {
	CollectionMethod param.Field[CollectionMethod] `json:"collection_method" api:"required"`
	Currency         param.Field[Currency]         `json:"currency" api:"required"`
	CustomerID       param.Field[string]           `json:"customer_id" api:"required"`
	// 1 to 50 items.
	Items param.Field[[]InvoiceNewParamsItemUnion] `json:"items" api:"required"`
	// The billing address of an invoice request.
	BillingAddress param.Field[InvoiceNewParamsBillingAddress] `json:"billing_address"`
	// 10 addresses or fewer.
	CcEmails param.Field[[]string] `json:"cc_emails"`
	// The customer email is used when this is not set.
	CustomerEmail param.Field[string] `json:"customer_email"`
	// 20 codes or fewer.
	DiscountIDs param.Field[[]string] `json:"discount_ids"`
	// A custom due date. Do not send it with `payment_terms`.
	DueAt param.Field[time.Time] `json:"due_at" format:"date-time"`
	// The reminder override of one invoice. The business settings fill the other
	// dunning values.
	Dunning param.Field[InvoiceNewParamsDunning] `json:"dunning"`
	Memo    param.Field[string]                  `json:"memo"`
	// Arbitrary key-value metadata. Values can be string, integer, number, or boolean.
	Metadata param.Field[MetadataParam] `json:"metadata"`
	// Do not send it with `due_at`.
	PaymentTerms param.Field[InvoiceNewParamsPaymentTermsUnion] `json:"payment_terms"`
	// `true` keeps a renewal draft for review before it finalises.
	ReviewRequired param.Field[bool]   `json:"review_required"`
	TaxID          param.Field[string] `json:"tax_id"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A line that exists only on this invoice.
type InvoiceNewParamsItem struct {
	Addons      param.Field[interface{}]               `json:"addons"`
	Description param.Field[string]                    `json:"description"`
	Kind        param.Field[InvoiceNewParamsItemsKind] `json:"kind"`
	ProductID   param.Field[string]                    `json:"product_id"`
	// The default is 1.
	Quantity param.Field[int64] `json:"quantity"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category"`
	// The price in the smallest currency unit.
	UnitAmount param.Field[int64] `json:"unit_amount"`
}

func (r InvoiceNewParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceNewParamsItem) implementsInvoiceNewParamsItemUnion() {}

// A line that exists only on this invoice.
//
// Satisfied by [InvoiceNewParamsItemsCustomItem],
// [InvoiceNewParamsItemsProductItem], [InvoiceNewParamsItem].
type InvoiceNewParamsItemUnion interface {
	implementsInvoiceNewParamsItemUnion()
}

// A line that exists only on this invoice.
type InvoiceNewParamsItemsCustomItem struct {
	Description param.Field[string]                              `json:"description" api:"required"`
	Kind        param.Field[InvoiceNewParamsItemsCustomItemKind] `json:"kind" api:"required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category" api:"required"`
	// The price in the smallest currency unit.
	UnitAmount param.Field[int64] `json:"unit_amount" api:"required"`
	// The default is 1.
	Quantity param.Field[int64] `json:"quantity"`
}

func (r InvoiceNewParamsItemsCustomItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceNewParamsItemsCustomItem) implementsInvoiceNewParamsItemUnion() {}

type InvoiceNewParamsItemsCustomItemKind string

const (
	InvoiceNewParamsItemsCustomItemKindCustom InvoiceNewParamsItemsCustomItemKind = "custom"
)

func (r InvoiceNewParamsItemsCustomItemKind) IsKnown() bool {
	switch r {
	case InvoiceNewParamsItemsCustomItemKindCustom:
		return true
	}
	return false
}

type InvoiceNewParamsItemsProductItem struct {
	ProductID param.Field[string]             `json:"product_id" api:"required"`
	Quantity  param.Field[int64]              `json:"quantity" api:"required"`
	Addons    param.Field[[]AttachAddonParam] `json:"addons"`
}

func (r InvoiceNewParamsItemsProductItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceNewParamsItemsProductItem) implementsInvoiceNewParamsItemUnion() {}

type InvoiceNewParamsItemsKind string

const (
	InvoiceNewParamsItemsKindCustom InvoiceNewParamsItemsKind = "custom"
)

func (r InvoiceNewParamsItemsKind) IsKnown() bool {
	switch r {
	case InvoiceNewParamsItemsKindCustom:
		return true
	}
	return false
}

// The billing address of an invoice request.
type InvoiceNewParamsBillingAddress struct {
	// Two-letter ISO country code (ISO 3166-1 alpha-2)
	Country param.Field[CountryCode] `json:"country" api:"required"`
	City    param.Field[string]      `json:"city"`
	State   param.Field[string]      `json:"state"`
	Street  param.Field[string]      `json:"street"`
	Zipcode param.Field[string]      `json:"zipcode"`
}

func (r InvoiceNewParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reminder override of one invoice. The business settings fill the other
// dunning values.
type InvoiceNewParamsDunning struct {
	RemindersEnabled param.Field[bool] `json:"reminders_enabled" api:"required"`
	// Days from the due date. A negative value sends the reminder before the due date.
	OffsetsDays param.Field[[]int64] `json:"offsets_days"`
}

func (r InvoiceNewParamsDunning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Do not send it with `due_at`.
type InvoiceNewParamsPaymentTerms struct {
	Type param.Field[InvoiceNewParamsPaymentTermsType] `json:"type" api:"required"`
	Days param.Field[int64]                            `json:"days"`
}

func (r InvoiceNewParamsPaymentTerms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceNewParamsPaymentTerms) implementsInvoiceNewParamsPaymentTermsUnion() {}

// Do not send it with `due_at`.
//
// Satisfied by [InvoiceNewParamsPaymentTermsType],
// [InvoiceNewParamsPaymentTermsObject], [InvoiceNewParamsPaymentTerms].
type InvoiceNewParamsPaymentTermsUnion interface {
	implementsInvoiceNewParamsPaymentTermsUnion()
}

type InvoiceNewParamsPaymentTermsType struct {
	Type param.Field[InvoiceNewParamsPaymentTermsTypeType] `json:"type" api:"required"`
}

func (r InvoiceNewParamsPaymentTermsType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceNewParamsPaymentTermsType) implementsInvoiceNewParamsPaymentTermsUnion() {}

type InvoiceNewParamsPaymentTermsTypeType string

const (
	InvoiceNewParamsPaymentTermsTypeTypeDueOnReceipt InvoiceNewParamsPaymentTermsTypeType = "due_on_receipt"
)

func (r InvoiceNewParamsPaymentTermsTypeType) IsKnown() bool {
	switch r {
	case InvoiceNewParamsPaymentTermsTypeTypeDueOnReceipt:
		return true
	}
	return false
}

type InvoiceNewParamsPaymentTermsObject struct {
	Days param.Field[int64]                                  `json:"days" api:"required"`
	Type param.Field[InvoiceNewParamsPaymentTermsObjectType] `json:"type" api:"required"`
}

func (r InvoiceNewParamsPaymentTermsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceNewParamsPaymentTermsObject) implementsInvoiceNewParamsPaymentTermsUnion() {}

type InvoiceNewParamsPaymentTermsObjectType string

const (
	InvoiceNewParamsPaymentTermsObjectTypeNet InvoiceNewParamsPaymentTermsObjectType = "net"
)

func (r InvoiceNewParamsPaymentTermsObjectType) IsKnown() bool {
	switch r {
	case InvoiceNewParamsPaymentTermsObjectTypeNet:
		return true
	}
	return false
}

type InvoiceUpdateParams struct {
	IfMatch param.Field[string] `header:"If-Match" api:"required"`
	// The billing address of an invoice request.
	BillingAddress param.Field[InvoiceUpdateParamsBillingAddress] `json:"billing_address"`
	// 10 addresses or fewer.
	CcEmails         param.Field[[]string]         `json:"cc_emails"`
	CollectionMethod param.Field[CollectionMethod] `json:"collection_method"`
	Currency         param.Field[Currency]         `json:"currency"`
	CustomerEmail    param.Field[string]           `json:"customer_email"`
	CustomerID       param.Field[string]           `json:"customer_id"`
	// 20 codes or fewer.
	DiscountIDs param.Field[[]string]  `json:"discount_ids"`
	DueAt       param.Field[time.Time] `json:"due_at" format:"date-time"`
	// The reminder override of one invoice. The business settings fill the other
	// dunning values.
	Dunning param.Field[InvoiceUpdateParamsDunning] `json:"dunning"`
	// 1 to 50 items.
	Items param.Field[[]InvoiceUpdateParamsItemUnion] `json:"items"`
	Memo  param.Field[string]                         `json:"memo"`
	// Arbitrary key-value metadata. Values can be string, integer, number, or boolean.
	Metadata       param.Field[MetadataParam]                        `json:"metadata"`
	PaymentTerms   param.Field[InvoiceUpdateParamsPaymentTermsUnion] `json:"payment_terms"`
	ReviewRequired param.Field[bool]                                 `json:"review_required"`
	TaxID          param.Field[string]                               `json:"tax_id"`
}

func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The billing address of an invoice request.
type InvoiceUpdateParamsBillingAddress struct {
	// Two-letter ISO country code (ISO 3166-1 alpha-2)
	Country param.Field[CountryCode] `json:"country" api:"required"`
	City    param.Field[string]      `json:"city"`
	State   param.Field[string]      `json:"state"`
	Street  param.Field[string]      `json:"street"`
	Zipcode param.Field[string]      `json:"zipcode"`
}

func (r InvoiceUpdateParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reminder override of one invoice. The business settings fill the other
// dunning values.
type InvoiceUpdateParamsDunning struct {
	RemindersEnabled param.Field[bool] `json:"reminders_enabled" api:"required"`
	// Days from the due date. A negative value sends the reminder before the due date.
	OffsetsDays param.Field[[]int64] `json:"offsets_days"`
}

func (r InvoiceUpdateParamsDunning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A line that exists only on this invoice.
type InvoiceUpdateParamsItem struct {
	Addons      param.Field[interface{}]                  `json:"addons"`
	Description param.Field[string]                       `json:"description"`
	Kind        param.Field[InvoiceUpdateParamsItemsKind] `json:"kind"`
	ProductID   param.Field[string]                       `json:"product_id"`
	// The default is 1.
	Quantity param.Field[int64] `json:"quantity"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category"`
	// The price in the smallest currency unit.
	UnitAmount param.Field[int64] `json:"unit_amount"`
}

func (r InvoiceUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceUpdateParamsItem) implementsInvoiceUpdateParamsItemUnion() {}

// A line that exists only on this invoice.
//
// Satisfied by [InvoiceUpdateParamsItemsCustomItem],
// [InvoiceUpdateParamsItemsProductItem], [InvoiceUpdateParamsItem].
type InvoiceUpdateParamsItemUnion interface {
	implementsInvoiceUpdateParamsItemUnion()
}

// A line that exists only on this invoice.
type InvoiceUpdateParamsItemsCustomItem struct {
	Description param.Field[string]                                 `json:"description" api:"required"`
	Kind        param.Field[InvoiceUpdateParamsItemsCustomItemKind] `json:"kind" api:"required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category" api:"required"`
	// The price in the smallest currency unit.
	UnitAmount param.Field[int64] `json:"unit_amount" api:"required"`
	// The default is 1.
	Quantity param.Field[int64] `json:"quantity"`
}

func (r InvoiceUpdateParamsItemsCustomItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceUpdateParamsItemsCustomItem) implementsInvoiceUpdateParamsItemUnion() {}

type InvoiceUpdateParamsItemsCustomItemKind string

const (
	InvoiceUpdateParamsItemsCustomItemKindCustom InvoiceUpdateParamsItemsCustomItemKind = "custom"
)

func (r InvoiceUpdateParamsItemsCustomItemKind) IsKnown() bool {
	switch r {
	case InvoiceUpdateParamsItemsCustomItemKindCustom:
		return true
	}
	return false
}

type InvoiceUpdateParamsItemsProductItem struct {
	ProductID param.Field[string]             `json:"product_id" api:"required"`
	Quantity  param.Field[int64]              `json:"quantity" api:"required"`
	Addons    param.Field[[]AttachAddonParam] `json:"addons"`
}

func (r InvoiceUpdateParamsItemsProductItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceUpdateParamsItemsProductItem) implementsInvoiceUpdateParamsItemUnion() {}

type InvoiceUpdateParamsItemsKind string

const (
	InvoiceUpdateParamsItemsKindCustom InvoiceUpdateParamsItemsKind = "custom"
)

func (r InvoiceUpdateParamsItemsKind) IsKnown() bool {
	switch r {
	case InvoiceUpdateParamsItemsKindCustom:
		return true
	}
	return false
}

type InvoiceUpdateParamsPaymentTerms struct {
	Type param.Field[InvoiceUpdateParamsPaymentTermsType] `json:"type" api:"required"`
	Days param.Field[int64]                               `json:"days"`
}

func (r InvoiceUpdateParamsPaymentTerms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceUpdateParamsPaymentTerms) implementsInvoiceUpdateParamsPaymentTermsUnion() {}

// Satisfied by [InvoiceUpdateParamsPaymentTermsType],
// [InvoiceUpdateParamsPaymentTermsObject], [InvoiceUpdateParamsPaymentTerms].
type InvoiceUpdateParamsPaymentTermsUnion interface {
	implementsInvoiceUpdateParamsPaymentTermsUnion()
}

type InvoiceUpdateParamsPaymentTermsType struct {
	Type param.Field[InvoiceUpdateParamsPaymentTermsTypeType] `json:"type" api:"required"`
}

func (r InvoiceUpdateParamsPaymentTermsType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceUpdateParamsPaymentTermsType) implementsInvoiceUpdateParamsPaymentTermsUnion() {}

type InvoiceUpdateParamsPaymentTermsTypeType string

const (
	InvoiceUpdateParamsPaymentTermsTypeTypeDueOnReceipt InvoiceUpdateParamsPaymentTermsTypeType = "due_on_receipt"
)

func (r InvoiceUpdateParamsPaymentTermsTypeType) IsKnown() bool {
	switch r {
	case InvoiceUpdateParamsPaymentTermsTypeTypeDueOnReceipt:
		return true
	}
	return false
}

type InvoiceUpdateParamsPaymentTermsObject struct {
	Days param.Field[int64]                                     `json:"days" api:"required"`
	Type param.Field[InvoiceUpdateParamsPaymentTermsObjectType] `json:"type" api:"required"`
}

func (r InvoiceUpdateParamsPaymentTermsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoiceUpdateParamsPaymentTermsObject) implementsInvoiceUpdateParamsPaymentTermsUnion() {}

type InvoiceUpdateParamsPaymentTermsObjectType string

const (
	InvoiceUpdateParamsPaymentTermsObjectTypeNet InvoiceUpdateParamsPaymentTermsObjectType = "net"
)

func (r InvoiceUpdateParamsPaymentTermsObjectType) IsKnown() bool {
	switch r {
	case InvoiceUpdateParamsPaymentTermsObjectTypeNet:
		return true
	}
	return false
}

type InvoiceListParams struct {
	CollectionMethod param.Field[InvoiceListParamsCollectionMethod] `query:"collection_method"`
	CreatedAtGte     param.Field[time.Time]                         `query:"created_at_gte" format:"date-time"`
	CreatedAtLte     param.Field[time.Time]                         `query:"created_at_lte" format:"date-time"`
	CustomerID       param.Field[string]                            `query:"customer_id"`
	DueAtGte         param.Field[time.Time]                         `query:"due_at_gte" format:"date-time"`
	DueAtLte         param.Field[time.Time]                         `query:"due_at_lte" format:"date-time"`
	// Page number. The default is 0.
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size. The default is 10 and the maximum is 100.
	PageSize param.Field[int64] `query:"page_size"`
	// Repeat the key to match more than one status, for example
	// `status=open&status=past_due`.
	Status         param.Field[[]InvoiceListParamsStatus] `query:"status"`
	SubscriptionID param.Field[string]                    `query:"subscription_id"`
	// Repeat the key to match an invoice that has any one of the tags.
	Tags           param.Field[[]InvoiceListParamsTag] `query:"tags"`
	TotalAmountGte param.Field[int64]                  `query:"total_amount_gte"`
	TotalAmountLte param.Field[int64]                  `query:"total_amount_lte"`
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvoiceListParamsCollectionMethod string

const (
	InvoiceListParamsCollectionMethodAutoCharge InvoiceListParamsCollectionMethod = "auto_charge"
	InvoiceListParamsCollectionMethodManual     InvoiceListParamsCollectionMethod = "manual"
)

func (r InvoiceListParamsCollectionMethod) IsKnown() bool {
	switch r {
	case InvoiceListParamsCollectionMethodAutoCharge, InvoiceListParamsCollectionMethodManual:
		return true
	}
	return false
}

// The `message` of each variant is the label of the invoice document.
type InvoiceListParamsStatus string

const (
	InvoiceListParamsStatusOpen          InvoiceListParamsStatus = "open"
	InvoiceListParamsStatusPaid          InvoiceListParamsStatus = "paid"
	InvoiceListParamsStatusVoid          InvoiceListParamsStatus = "void"
	InvoiceListParamsStatusDraft         InvoiceListParamsStatus = "draft"
	InvoiceListParamsStatusPastDue       InvoiceListParamsStatus = "past_due"
	InvoiceListParamsStatusUncollectable InvoiceListParamsStatus = "uncollectable"
)

func (r InvoiceListParamsStatus) IsKnown() bool {
	switch r {
	case InvoiceListParamsStatusOpen, InvoiceListParamsStatusPaid, InvoiceListParamsStatusVoid, InvoiceListParamsStatusDraft, InvoiceListParamsStatusPastDue, InvoiceListParamsStatusUncollectable:
		return true
	}
	return false
}

type InvoiceListParamsTag string

const (
	InvoiceListParamsTagRefunded          InvoiceListParamsTag = "refunded"
	InvoiceListParamsTagPartiallyRefunded InvoiceListParamsTag = "partially_refunded"
	InvoiceListParamsTagDisputed          InvoiceListParamsTag = "disputed"
	InvoiceListParamsTagPaymentReversed   InvoiceListParamsTag = "payment_reversed"
)

func (r InvoiceListParamsTag) IsKnown() bool {
	switch r {
	case InvoiceListParamsTagRefunded, InvoiceListParamsTagPartiallyRefunded, InvoiceListParamsTagDisputed, InvoiceListParamsTagPaymentReversed:
		return true
	}
	return false
}

type InvoiceDeleteParams struct {
	IfMatch param.Field[string] `header:"If-Match" api:"required"`
}

type InvoiceAddNoteParams struct {
	NoteRequest NoteRequestParam `json:"note_request" api:"required"`
}

func (r InvoiceAddNoteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NoteRequest)
}

type InvoiceFinalizeParams struct {
	IfMatch param.Field[string] `header:"If-Match" api:"required"`
	// 10 addresses or fewer.
	CcEmails      param.Field[[]string] `json:"cc_emails"`
	CustomerEmail param.Field[string]   `json:"customer_email"`
	// `true` sends the invoice email after finalise.
	Send param.Field[bool] `json:"send"`
}

func (r InvoiceFinalizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceMarkUncollectableParams struct {
	IfMatch param.Field[string] `header:"If-Match" api:"required"`
}

type InvoicePreviewParams struct {
	CollectionMethod param.Field[CollectionMethod] `json:"collection_method" api:"required"`
	Currency         param.Field[Currency]         `json:"currency" api:"required"`
	CustomerID       param.Field[string]           `json:"customer_id" api:"required"`
	// 1 to 50 items.
	Items param.Field[[]InvoicePreviewParamsItemUnion] `json:"items" api:"required"`
	// The preview body. `json` returns the invoice JSON. `pdf` returns the PDF with a
	// draft mark. `email` returns the HTML of the invoice email. A preview saves no
	// row.
	Format param.Field[InvoicePreviewParamsFormat] `query:"format"`
	// The billing address of an invoice request.
	BillingAddress param.Field[InvoicePreviewParamsBillingAddress] `json:"billing_address"`
	// 10 addresses or fewer.
	CcEmails param.Field[[]string] `json:"cc_emails"`
	// The customer email is used when this is not set.
	CustomerEmail param.Field[string] `json:"customer_email"`
	// 20 codes or fewer.
	DiscountIDs param.Field[[]string] `json:"discount_ids"`
	// A custom due date. Do not send it with `payment_terms`.
	DueAt param.Field[time.Time] `json:"due_at" format:"date-time"`
	// The reminder override of one invoice. The business settings fill the other
	// dunning values.
	Dunning param.Field[InvoicePreviewParamsDunning] `json:"dunning"`
	Memo    param.Field[string]                      `json:"memo"`
	// Arbitrary key-value metadata. Values can be string, integer, number, or boolean.
	Metadata param.Field[MetadataParam] `json:"metadata"`
	// Do not send it with `due_at`.
	PaymentTerms param.Field[InvoicePreviewParamsPaymentTermsUnion] `json:"payment_terms"`
	// `true` keeps a renewal draft for review before it finalises.
	ReviewRequired param.Field[bool]   `json:"review_required"`
	TaxID          param.Field[string] `json:"tax_id"`
}

func (r InvoicePreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [InvoicePreviewParams]'s query parameters as `url.Values`.
func (r InvoicePreviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A line that exists only on this invoice.
type InvoicePreviewParamsItem struct {
	Addons      param.Field[interface{}]                   `json:"addons"`
	Description param.Field[string]                        `json:"description"`
	Kind        param.Field[InvoicePreviewParamsItemsKind] `json:"kind"`
	ProductID   param.Field[string]                        `json:"product_id"`
	// The default is 1.
	Quantity param.Field[int64] `json:"quantity"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category"`
	// The price in the smallest currency unit.
	UnitAmount param.Field[int64] `json:"unit_amount"`
}

func (r InvoicePreviewParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoicePreviewParamsItem) implementsInvoicePreviewParamsItemUnion() {}

// A line that exists only on this invoice.
//
// Satisfied by [InvoicePreviewParamsItemsCustomItem],
// [InvoicePreviewParamsItemsProductItem], [InvoicePreviewParamsItem].
type InvoicePreviewParamsItemUnion interface {
	implementsInvoicePreviewParamsItemUnion()
}

// A line that exists only on this invoice.
type InvoicePreviewParamsItemsCustomItem struct {
	Description param.Field[string]                                  `json:"description" api:"required"`
	Kind        param.Field[InvoicePreviewParamsItemsCustomItemKind] `json:"kind" api:"required"`
	// Represents the different categories of taxation applicable to various products
	// and services.
	TaxCategory param.Field[TaxCategory] `json:"tax_category" api:"required"`
	// The price in the smallest currency unit.
	UnitAmount param.Field[int64] `json:"unit_amount" api:"required"`
	// The default is 1.
	Quantity param.Field[int64] `json:"quantity"`
}

func (r InvoicePreviewParamsItemsCustomItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoicePreviewParamsItemsCustomItem) implementsInvoicePreviewParamsItemUnion() {}

type InvoicePreviewParamsItemsCustomItemKind string

const (
	InvoicePreviewParamsItemsCustomItemKindCustom InvoicePreviewParamsItemsCustomItemKind = "custom"
)

func (r InvoicePreviewParamsItemsCustomItemKind) IsKnown() bool {
	switch r {
	case InvoicePreviewParamsItemsCustomItemKindCustom:
		return true
	}
	return false
}

type InvoicePreviewParamsItemsProductItem struct {
	ProductID param.Field[string]             `json:"product_id" api:"required"`
	Quantity  param.Field[int64]              `json:"quantity" api:"required"`
	Addons    param.Field[[]AttachAddonParam] `json:"addons"`
}

func (r InvoicePreviewParamsItemsProductItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoicePreviewParamsItemsProductItem) implementsInvoicePreviewParamsItemUnion() {}

type InvoicePreviewParamsItemsKind string

const (
	InvoicePreviewParamsItemsKindCustom InvoicePreviewParamsItemsKind = "custom"
)

func (r InvoicePreviewParamsItemsKind) IsKnown() bool {
	switch r {
	case InvoicePreviewParamsItemsKindCustom:
		return true
	}
	return false
}

// The preview body. `json` returns the invoice JSON. `pdf` returns the PDF with a
// draft mark. `email` returns the HTML of the invoice email. A preview saves no
// row.
type InvoicePreviewParamsFormat string

const (
	InvoicePreviewParamsFormatJson  InvoicePreviewParamsFormat = "json"
	InvoicePreviewParamsFormatPdf   InvoicePreviewParamsFormat = "pdf"
	InvoicePreviewParamsFormatEmail InvoicePreviewParamsFormat = "email"
)

func (r InvoicePreviewParamsFormat) IsKnown() bool {
	switch r {
	case InvoicePreviewParamsFormatJson, InvoicePreviewParamsFormatPdf, InvoicePreviewParamsFormatEmail:
		return true
	}
	return false
}

// The billing address of an invoice request.
type InvoicePreviewParamsBillingAddress struct {
	// Two-letter ISO country code (ISO 3166-1 alpha-2)
	Country param.Field[CountryCode] `json:"country" api:"required"`
	City    param.Field[string]      `json:"city"`
	State   param.Field[string]      `json:"state"`
	Street  param.Field[string]      `json:"street"`
	Zipcode param.Field[string]      `json:"zipcode"`
}

func (r InvoicePreviewParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reminder override of one invoice. The business settings fill the other
// dunning values.
type InvoicePreviewParamsDunning struct {
	RemindersEnabled param.Field[bool] `json:"reminders_enabled" api:"required"`
	// Days from the due date. A negative value sends the reminder before the due date.
	OffsetsDays param.Field[[]int64] `json:"offsets_days"`
}

func (r InvoicePreviewParamsDunning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Do not send it with `due_at`.
type InvoicePreviewParamsPaymentTerms struct {
	Type param.Field[InvoicePreviewParamsPaymentTermsType] `json:"type" api:"required"`
	Days param.Field[int64]                                `json:"days"`
}

func (r InvoicePreviewParamsPaymentTerms) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoicePreviewParamsPaymentTerms) implementsInvoicePreviewParamsPaymentTermsUnion() {}

// Do not send it with `due_at`.
//
// Satisfied by [InvoicePreviewParamsPaymentTermsType],
// [InvoicePreviewParamsPaymentTermsObject], [InvoicePreviewParamsPaymentTerms].
type InvoicePreviewParamsPaymentTermsUnion interface {
	implementsInvoicePreviewParamsPaymentTermsUnion()
}

type InvoicePreviewParamsPaymentTermsType struct {
	Type param.Field[InvoicePreviewParamsPaymentTermsTypeType] `json:"type" api:"required"`
}

func (r InvoicePreviewParamsPaymentTermsType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoicePreviewParamsPaymentTermsType) implementsInvoicePreviewParamsPaymentTermsUnion() {}

type InvoicePreviewParamsPaymentTermsTypeType string

const (
	InvoicePreviewParamsPaymentTermsTypeTypeDueOnReceipt InvoicePreviewParamsPaymentTermsTypeType = "due_on_receipt"
)

func (r InvoicePreviewParamsPaymentTermsTypeType) IsKnown() bool {
	switch r {
	case InvoicePreviewParamsPaymentTermsTypeTypeDueOnReceipt:
		return true
	}
	return false
}

type InvoicePreviewParamsPaymentTermsObject struct {
	Days param.Field[int64]                                      `json:"days" api:"required"`
	Type param.Field[InvoicePreviewParamsPaymentTermsObjectType] `json:"type" api:"required"`
}

func (r InvoicePreviewParamsPaymentTermsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r InvoicePreviewParamsPaymentTermsObject) implementsInvoicePreviewParamsPaymentTermsUnion() {}

type InvoicePreviewParamsPaymentTermsObjectType string

const (
	InvoicePreviewParamsPaymentTermsObjectTypeNet InvoicePreviewParamsPaymentTermsObjectType = "net"
)

func (r InvoicePreviewParamsPaymentTermsObjectType) IsKnown() bool {
	switch r {
	case InvoicePreviewParamsPaymentTermsObjectTypeNet:
		return true
	}
	return false
}

type InvoicePreviewExistingParams struct {
	// The preview body. `json` returns the invoice JSON. `pdf` returns the PDF with a
	// draft mark. `email` returns the HTML of the invoice email. A preview saves no
	// row.
	Format param.Field[InvoicePreviewExistingParamsFormat] `query:"format"`
}

// URLQuery serializes [InvoicePreviewExistingParams]'s query parameters as
// `url.Values`.
func (r InvoicePreviewExistingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The preview body. `json` returns the invoice JSON. `pdf` returns the PDF with a
// draft mark. `email` returns the HTML of the invoice email. A preview saves no
// row.
type InvoicePreviewExistingParamsFormat string

const (
	InvoicePreviewExistingParamsFormatJson  InvoicePreviewExistingParamsFormat = "json"
	InvoicePreviewExistingParamsFormatPdf   InvoicePreviewExistingParamsFormat = "pdf"
	InvoicePreviewExistingParamsFormatEmail InvoicePreviewExistingParamsFormat = "email"
)

func (r InvoicePreviewExistingParamsFormat) IsKnown() bool {
	switch r {
	case InvoicePreviewExistingParamsFormatJson, InvoicePreviewExistingParamsFormatPdf, InvoicePreviewExistingParamsFormatEmail:
		return true
	}
	return false
}

type InvoiceRetryParams struct {
	IfMatch param.Field[string] `header:"If-Match" api:"required"`
}

type InvoiceSendParams struct {
	IfMatch param.Field[string] `header:"If-Match" api:"required"`
}

type InvoiceUpdateDunningParams struct {
	Action  param.Field[InvoiceUpdateDunningParamsAction] `json:"action" api:"required"`
	IfMatch param.Field[string]                           `header:"If-Match" api:"required"`
	// Needed for `reschedule`. Days from the due date.
	OffsetsDays param.Field[[]int64] `json:"offsets_days"`
}

func (r InvoiceUpdateDunningParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceUpdateDunningParamsAction string

const (
	InvoiceUpdateDunningParamsActionPause      InvoiceUpdateDunningParamsAction = "pause"
	InvoiceUpdateDunningParamsActionResume     InvoiceUpdateDunningParamsAction = "resume"
	InvoiceUpdateDunningParamsActionReschedule InvoiceUpdateDunningParamsAction = "reschedule"
)

func (r InvoiceUpdateDunningParamsAction) IsKnown() bool {
	switch r {
	case InvoiceUpdateDunningParamsActionPause, InvoiceUpdateDunningParamsActionResume, InvoiceUpdateDunningParamsActionReschedule:
		return true
	}
	return false
}

type InvoiceVoidParams struct {
	IfMatch param.Field[string] `header:"If-Match" api:"required"`
}
