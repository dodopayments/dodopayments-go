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
)

// EntitlementFileService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitlementFileService] method instead.
type EntitlementFileService struct {
	Options []option.RequestOption
}

// NewEntitlementFileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntitlementFileService(opts ...option.RequestOption) (r *EntitlementFileService) {
	r = &EntitlementFileService{}
	r.Options = opts
	return
}

// Companion to `post_entitlement_file`. Deletes the file from the Entitlements
// Engine (force=true) and atomically removes the `file_id` from the entitlement's
// `integration_config.digital_file_ids` JSONB array. EE delete happens first; if
// it fails we surface the error and leave local state untouched.
func (r *EntitlementFileService) Delete(ctx context.Context, id string, fileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return err
	}
	path := fmt.Sprintf("entitlements/%s/files/%s", id, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Streams a multipart/form-data body to the Entitlements Engine
// (`POST /api/digital-files/dodo/files/upload`) and appends the returned `file_id`
// to the entitlement's `integration_config.digital_file_ids` using a JSONB array
// append. Compensates EE-side on local DB write failure (best-effort delete of the
// just-uploaded file).
func (r *EntitlementFileService) Upload(ctx context.Context, id string, opts ...option.RequestOption) (res *EntitlementFileUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entitlements/%s/files", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type EntitlementFileUploadResponse struct {
	// EE-issued digital file id; appended to
	// `entitlements.integration_config.digital_file_ids`.
	FileID string                            `json:"file_id" api:"required"`
	JSON   entitlementFileUploadResponseJSON `json:"-"`
}

// entitlementFileUploadResponseJSON contains the JSON metadata for the struct
// [EntitlementFileUploadResponse]
type entitlementFileUploadResponseJSON struct {
	FileID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementFileUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementFileUploadResponseJSON) RawJSON() string {
	return r.raw
}
