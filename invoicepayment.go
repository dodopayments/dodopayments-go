// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// InvoicePaymentService contains methods and other services that help with
// interacting with the Dodo Payments API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoicePaymentService] method instead.
type InvoicePaymentService struct {
	Options []option.RequestOption
}

// NewInvoicePaymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvoicePaymentService(opts ...option.RequestOption) (r *InvoicePaymentService) {
	r = &InvoicePaymentService{}
	r.Options = opts
	return
}

func (r *InvoicePaymentService) Get(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if paymentID == "" {
		err = errors.New("missing required payment_id parameter")
		return
	}
	path := fmt.Sprintf("invoices/payments/%s", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
