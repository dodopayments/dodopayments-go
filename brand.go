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
	"github.com/dodopayments/dodopayments-go/packages/param"
	"github.com/dodopayments/dodopayments-go/packages/respjson"
)

// BrandService contains methods and other services that help with interacting with
// the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	Options []option.RequestOption
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r BrandService) {
	r = BrandService{}
	r.Options = opts
	return
}

func (r *BrandService) New(ctx context.Context, body BrandNewParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Thin handler just calls `get_brand` and wraps in `Json(...)`
func (r *BrandService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *BrandService) Update(ctx context.Context, id string, body BrandUpdateParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *BrandService) List(ctx context.Context, opts ...option.RequestOption) (res *BrandListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *BrandService) UpdateImages(ctx context.Context, id string, opts ...option.RequestOption) (res *BrandUpdateImagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s/images", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type Brand struct {
	BrandID             string `json:"brand_id,required"`
	BusinessID          string `json:"business_id,required"`
	Enabled             bool   `json:"enabled,required"`
	StatementDescriptor string `json:"statement_descriptor,required"`
	VerificationEnabled bool   `json:"verification_enabled,required"`
	// Any of "Success", "Fail", "Review", "Hold".
	VerificationStatus BrandVerificationStatus `json:"verification_status,required"`
	Description        string                  `json:"description,nullable"`
	Image              string                  `json:"image,nullable"`
	Name               string                  `json:"name,nullable"`
	// Incase the brand verification fails or is put on hold
	ReasonForHold string `json:"reason_for_hold,nullable"`
	SupportEmail  string `json:"support_email,nullable"`
	URL           string `json:"url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID             respjson.Field
		BusinessID          respjson.Field
		Enabled             respjson.Field
		StatementDescriptor respjson.Field
		VerificationEnabled respjson.Field
		VerificationStatus  respjson.Field
		Description         respjson.Field
		Image               respjson.Field
		Name                respjson.Field
		ReasonForHold       respjson.Field
		SupportEmail        respjson.Field
		URL                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Brand) RawJSON() string { return r.JSON.raw }
func (r *Brand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandVerificationStatus string

const (
	BrandVerificationStatusSuccess BrandVerificationStatus = "Success"
	BrandVerificationStatusFail    BrandVerificationStatus = "Fail"
	BrandVerificationStatusReview  BrandVerificationStatus = "Review"
	BrandVerificationStatusHold    BrandVerificationStatus = "Hold"
)

type BrandListResponse struct {
	// List of brands for this business
	Items []Brand `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandUpdateImagesResponse struct {
	// UUID that will be used as the image identifier/key suffix
	ImageID string `json:"image_id,required" format:"uuid"`
	// Presigned URL to upload the image
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageID     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandUpdateImagesResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandUpdateImagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandNewParams struct {
	Description         param.Opt[string] `json:"description,omitzero"`
	Name                param.Opt[string] `json:"name,omitzero"`
	StatementDescriptor param.Opt[string] `json:"statement_descriptor,omitzero"`
	SupportEmail        param.Opt[string] `json:"support_email,omitzero"`
	URL                 param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r BrandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandUpdateParams struct {
	// The UUID you got back from the presigned‐upload call
	ImageID             param.Opt[string] `json:"image_id,omitzero" format:"uuid"`
	Name                param.Opt[string] `json:"name,omitzero"`
	StatementDescriptor param.Opt[string] `json:"statement_descriptor,omitzero"`
	SupportEmail        param.Opt[string] `json:"support_email,omitzero"`
	paramObj
}

func (r BrandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BrandUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
