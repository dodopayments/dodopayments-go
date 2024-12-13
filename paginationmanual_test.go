// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/dodo-payments-go"
	"github.com/stainless-sdks/dodo-payments-go/internal/testutil"
	"github.com/stainless-sdks/dodo-payments-go/option"
)

func TestManualPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dodopayments.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	page, err := client.Payments.List(context.TODO(), dodopayments.PaymentListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, payment := range page.Items {
		t.Logf("%+v\n", payment.PaymentID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, payment := range page.Items {
			t.Logf("%+v\n", payment.PaymentID)
		}
	}
}
