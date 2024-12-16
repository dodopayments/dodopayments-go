// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dodopayments/dodopayments-go/internal/apijson"
	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// ProductImageService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductImageService] method instead.
type ProductImageService struct {
	Options []option.RequestOption
}

// NewProductImageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProductImageService(opts ...option.RequestOption) (r *ProductImageService) {
	r = &ProductImageService{}
	r.Options = opts
	return
}

func (r *ProductImageService) Update(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductImageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("products/%s/images", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type ProductImageUpdateResponse struct {
	URL  string                         `json:"url,required"`
	JSON productImageUpdateResponseJSON `json:"-"`
}

// productImageUpdateResponseJSON contains the JSON metadata for the struct
// [ProductImageUpdateResponse]
type productImageUpdateResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProductImageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productImageUpdateResponseJSON) RawJSON() string {
	return r.raw
}
