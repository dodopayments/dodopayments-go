// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/param"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
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
func NewBrandService(opts ...option.RequestOption) (r *BrandService) {
	r = &BrandService{}
	r.Options = opts
	return
}

func (r *BrandService) New(ctx context.Context, body BrandNewParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = append(r.Options[:], opts...)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Thin handler just calls `get_brand` and wraps in `Json(...)`
func (r *BrandService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Brand, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *BrandService) Update(ctx context.Context, id string, body BrandUpdateParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *BrandService) List(ctx context.Context, opts ...option.RequestOption) (res *BrandListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *BrandService) UpdateImages(ctx context.Context, id string, opts ...option.RequestOption) (res *BrandUpdateImagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("brands/%s/images", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type Brand struct {
	BrandID             string                  `json:"brand_id,required"`
	BusinessID          string                  `json:"business_id,required"`
	Enabled             bool                    `json:"enabled,required"`
	StatementDescriptor string                  `json:"statement_descriptor,required"`
	VerificationEnabled bool                    `json:"verification_enabled,required"`
	VerificationStatus  BrandVerificationStatus `json:"verification_status,required"`
	Description         string                  `json:"description,nullable"`
	Image               string                  `json:"image,nullable"`
	Name                string                  `json:"name,nullable"`
	// Incase the brand verification fails or is put on hold
	ReasonForHold string    `json:"reason_for_hold,nullable"`
	SupportEmail  string    `json:"support_email,nullable"`
	URL           string    `json:"url,nullable"`
	JSON          brandJSON `json:"-"`
}

// brandJSON contains the JSON metadata for the struct [Brand]
type brandJSON struct {
	BrandID             apijson.Field
	BusinessID          apijson.Field
	Enabled             apijson.Field
	StatementDescriptor apijson.Field
	VerificationEnabled apijson.Field
	VerificationStatus  apijson.Field
	Description         apijson.Field
	Image               apijson.Field
	Name                apijson.Field
	ReasonForHold       apijson.Field
	SupportEmail        apijson.Field
	URL                 apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Brand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandJSON) RawJSON() string {
	return r.raw
}

type BrandVerificationStatus string

const (
	BrandVerificationStatusSuccess BrandVerificationStatus = "Success"
	BrandVerificationStatusFail    BrandVerificationStatus = "Fail"
	BrandVerificationStatusReview  BrandVerificationStatus = "Review"
	BrandVerificationStatusHold    BrandVerificationStatus = "Hold"
)

func (r BrandVerificationStatus) IsKnown() bool {
	switch r {
	case BrandVerificationStatusSuccess, BrandVerificationStatusFail, BrandVerificationStatusReview, BrandVerificationStatusHold:
		return true
	}
	return false
}

type BrandListResponse struct {
	// List of brands for this business
	Items []Brand               `json:"items,required"`
	JSON  brandListResponseJSON `json:"-"`
}

// brandListResponseJSON contains the JSON metadata for the struct
// [BrandListResponse]
type brandListResponseJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandListResponseJSON) RawJSON() string {
	return r.raw
}

type BrandUpdateImagesResponse struct {
	// UUID that will be used as the image identifier/key suffix
	ImageID string `json:"image_id,required" format:"uuid"`
	// Presigned URL to upload the image
	URL  string                        `json:"url,required"`
	JSON brandUpdateImagesResponseJSON `json:"-"`
}

// brandUpdateImagesResponseJSON contains the JSON metadata for the struct
// [BrandUpdateImagesResponse]
type brandUpdateImagesResponseJSON struct {
	ImageID     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandUpdateImagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandUpdateImagesResponseJSON) RawJSON() string {
	return r.raw
}

type BrandNewParams struct {
	Description         param.Field[string] `json:"description"`
	Name                param.Field[string] `json:"name"`
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	SupportEmail        param.Field[string] `json:"support_email"`
	URL                 param.Field[string] `json:"url"`
}

func (r BrandNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrandUpdateParams struct {
	// The UUID you got back from the presigned‐upload call
	ImageID             param.Field[string] `json:"image_id" format:"uuid"`
	Name                param.Field[string] `json:"name"`
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	SupportEmail        param.Field[string] `json:"support_email"`
}

func (r BrandUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
