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

// BlocklistCustomerService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlocklistCustomerService] method instead.
type BlocklistCustomerService struct {
	Options []option.RequestOption
	Notes   *BlocklistCustomerNoteService
}

// NewBlocklistCustomerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlocklistCustomerService(opts ...option.RequestOption) (r *BlocklistCustomerService) {
	r = &BlocklistCustomerService{}
	r.Options = opts
	r.Notes = NewBlocklistCustomerNoteService(opts...)
	return
}

func (r *BlocklistCustomerService) New(ctx context.Context, body BlocklistCustomerNewParams, opts ...option.RequestOption) (res *BlockedCustomer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "blocklist/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlocklistCustomerService) Get(ctx context.Context, entryID string, opts ...option.RequestOption) (res *BlockedCustomer, err error) {
	opts = slices.Concat(r.Options, opts)
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("blocklist/customers/%s", entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

func (r *BlocklistCustomerService) List(ctx context.Context, query BlocklistCustomerListParams, opts ...option.RequestOption) (res *pagination.DefaultPageNumberPagination[BlockedCustomer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "blocklist/customers"
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

func (r *BlocklistCustomerService) ListAutoPaging(ctx context.Context, query BlocklistCustomerListParams, opts ...option.RequestOption) *pagination.DefaultPageNumberPaginationAutoPager[BlockedCustomer] {
	return pagination.NewDefaultPageNumberPaginationAutoPager(r.List(ctx, query, opts...))
}

func (r *BlocklistCustomerService) Delete(ctx context.Context, entryID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return err
	}
	path := fmt.Sprintf("blocklist/customers/%s", entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type BlockByCustomerIDParam struct {
	// Customer to block. The block still applies to that customer's email.
	CustomerID param.Field[string] `json:"customer_id" api:"required"`
}

func (r BlockByCustomerIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockByEmailParam struct {
	// Email to block. It must belong to an existing customer of this business.
	Email param.Field[string] `json:"email" api:"required"`
}

func (r BlockByEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlockedCustomer struct {
	ID            string    `json:"id" api:"required"`
	CreatedAt     time.Time `json:"created_at" api:"required" format:"date-time"`
	CustomerEmail string    `json:"customer_email" api:"required"`
	CustomerID    string    `json:"customer_id" api:"required"`
	CustomerName  string    `json:"customer_name" api:"required"`
	// Customer id or email that the merchant supplied.
	Identifier string `json:"identifier" api:"required"`
	// Where a block came from. `Api` marks an API-key caller, which carries no
	// dashboard actor. The other values name the screen the merchant used.
	Source BlockedCustomerSource `json:"source" api:"required"`
	// Dashboard user who blocked the customer. `null` for an API-key caller.
	BlockedByEmail string `json:"blocked_by_email" api:"nullable"`
	// Subscriptions this block cancelled. Present on the create response only.
	CancelledSubscriptionIDs []string `json:"cancelled_subscription_ids" api:"nullable"`
	// Activity log. Present on the detail response only.
	Notes  []BlockedCustomerNote `json:"notes" api:"nullable"`
	Reason string                `json:"reason" api:"nullable"`
	// Subscriptions this block left live, because the cancel failed or the inline
	// batch filled up. Repeat the create call to continue; the block itself is already
	// in force.
	RemainingSubscriptionIDs []string `json:"remaining_subscription_ids" api:"nullable"`
	// False when the block left live subscriptions behind, including the case where
	// the sweep could not list them and `remaining_subscription_ids` is therefore
	// unknown. Repeat the create call until it reads true.
	SubscriptionsSwept bool                `json:"subscriptions_swept" api:"nullable"`
	UnblockedAt        time.Time           `json:"unblocked_at" api:"nullable" format:"date-time"`
	JSON               blockedCustomerJSON `json:"-"`
}

// blockedCustomerJSON contains the JSON metadata for the struct [BlockedCustomer]
type blockedCustomerJSON struct {
	ID                       apijson.Field
	CreatedAt                apijson.Field
	CustomerEmail            apijson.Field
	CustomerID               apijson.Field
	CustomerName             apijson.Field
	Identifier               apijson.Field
	Source                   apijson.Field
	BlockedByEmail           apijson.Field
	CancelledSubscriptionIDs apijson.Field
	Notes                    apijson.Field
	Reason                   apijson.Field
	RemainingSubscriptionIDs apijson.Field
	SubscriptionsSwept       apijson.Field
	UnblockedAt              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *BlockedCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockedCustomerJSON) RawJSON() string {
	return r.raw
}

// Where a block came from. `Api` marks an API-key caller, which carries no
// dashboard actor. The other values name the screen the merchant used.
type BlockedCustomerSource string

const (
	BlockedCustomerSourceBlocklistPage BlockedCustomerSource = "blocklist_page"
	BlockedCustomerSourceCustomerPage  BlockedCustomerSource = "customer_page"
	BlockedCustomerSourcePaymentPage   BlockedCustomerSource = "payment_page"
	BlockedCustomerSourceDisputePage   BlockedCustomerSource = "dispute_page"
	BlockedCustomerSourceAPI           BlockedCustomerSource = "api"
)

func (r BlockedCustomerSource) IsKnown() bool {
	switch r {
	case BlockedCustomerSourceBlocklistPage, BlockedCustomerSourceCustomerPage, BlockedCustomerSourcePaymentPage, BlockedCustomerSourceDisputePage, BlockedCustomerSourceAPI:
		return true
	}
	return false
}

// Satisfied by
// [CreateBlockedCustomerRequestBlocklistCustomersBlockByCustomerIDParam],
// [CreateBlockedCustomerRequestBlocklistCustomersBlockByEmailParam].
type CreateBlockedCustomerRequestUnionParam interface {
	implementsCreateBlockedCustomerRequestUnionParam()
}

type CreateBlockedCustomerRequestBlocklistCustomersBlockByCustomerIDParam struct {
	// Why the merchant blocked this customer. The entry page shows it.
	Reason param.Field[string] `json:"reason"`
	// Screen the merchant blocked from. Ignored for an API-key caller, whose entry
	// always records `api`. A dashboard caller that omits it records `blocklist_page`.
	Source param.Field[BlockedCustomerSource] `json:"source"`
	BlockByCustomerIDParam
}

func (r CreateBlockedCustomerRequestBlocklistCustomersBlockByCustomerIDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateBlockedCustomerRequestBlocklistCustomersBlockByCustomerIDParam) implementsCreateBlockedCustomerRequestUnionParam() {
}

type CreateBlockedCustomerRequestBlocklistCustomersBlockByEmailParam struct {
	// Why the merchant blocked this customer. The entry page shows it.
	Reason param.Field[string] `json:"reason"`
	// Screen the merchant blocked from. Ignored for an API-key caller, whose entry
	// always records `api`. A dashboard caller that omits it records `blocklist_page`.
	Source param.Field[BlockedCustomerSource] `json:"source"`
	BlockByEmailParam
}

func (r CreateBlockedCustomerRequestBlocklistCustomersBlockByEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CreateBlockedCustomerRequestBlocklistCustomersBlockByEmailParam) implementsCreateBlockedCustomerRequestUnionParam() {
}

type BlocklistCustomerNewParams struct {
	CreateBlockedCustomerRequest CreateBlockedCustomerRequestUnionParam `json:"create_blocked_customer_request" api:"required"`
}

func (r BlocklistCustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateBlockedCustomerRequest)
}

type BlocklistCustomerListParams struct {
	// Filter by the dashboard user who blocked the customer.
	BlockedByEmail param.Field[string] `query:"blocked_by_email"`
	// Blocked on or after this time.
	CreatedAtGte param.Field[time.Time] `query:"created_at_gte" format:"date-time"`
	// Blocked on or before this time.
	CreatedAtLte param.Field[time.Time] `query:"created_at_lte" format:"date-time"`
	// Partial, case-insensitive match on the email and on the customer id.
	Identifier param.Field[string] `query:"identifier"`
	// Page number. Default 0.
	PageNumber param.Field[int64] `query:"page_number"`
	// Page size. Default 10, maximum 100.
	PageSize param.Field[int64] `query:"page_size"`
}

// URLQuery serializes [BlocklistCustomerListParams]'s query parameters as
// `url.Values`.
func (r BlocklistCustomerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
