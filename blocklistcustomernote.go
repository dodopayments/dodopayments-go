// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// BlocklistCustomerNoteService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlocklistCustomerNoteService] method instead.
type BlocklistCustomerNoteService struct {
	Options []option.RequestOption
}

// NewBlocklistCustomerNoteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlocklistCustomerNoteService(opts ...option.RequestOption) (r *BlocklistCustomerNoteService) {
	r = &BlocklistCustomerNoteService{}
	r.Options = opts
	return
}

func (r *BlocklistCustomerNoteService) New(ctx context.Context, entryID string, body BlocklistCustomerNoteNewParams, opts ...option.RequestOption) (res *BlockedCustomerNote, err error) {
	opts = slices.Concat(r.Options, opts)
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("blocklist/customers/%s/notes", entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

func (r *BlocklistCustomerNoteService) Update(ctx context.Context, entryID string, noteID string, body BlocklistCustomerNoteUpdateParams, opts ...option.RequestOption) (res *BlockedCustomerNote, err error) {
	opts = slices.Concat(r.Options, opts)
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	if noteID == "" {
		err = errors.New("missing required note_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("blocklist/customers/%s/notes/%s", entryID, noteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type BlockedCustomerNote struct {
	ID          string                  `json:"id" api:"required"`
	CreatedAt   time.Time               `json:"created_at" api:"required" format:"date-time"`
	Note        string                  `json:"note" api:"required"`
	AuthorEmail string                  `json:"author_email" api:"nullable"`
	UpdatedAt   time.Time               `json:"updated_at" api:"nullable" format:"date-time"`
	JSON        blockedCustomerNoteJSON `json:"-"`
}

// blockedCustomerNoteJSON contains the JSON metadata for the struct
// [BlockedCustomerNote]
type blockedCustomerNoteJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Note        apijson.Field
	AuthorEmail apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockedCustomerNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockedCustomerNoteJSON) RawJSON() string {
	return r.raw
}

type NoteRequestParam struct {
	Note param.Field[string] `json:"note" api:"required"`
}

func (r NoteRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BlocklistCustomerNoteNewParams struct {
	NoteRequest NoteRequestParam `json:"note_request" api:"required"`
}

func (r BlocklistCustomerNoteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NoteRequest)
}

type BlocklistCustomerNoteUpdateParams struct {
	NoteRequest NoteRequestParam `json:"note_request" api:"required"`
}

func (r BlocklistCustomerNoteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NoteRequest)
}
