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

func TestUsage(t *testing.T) {
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
	payment, err := client.Payments.New(context.TODO(), dodopayments.PaymentNewParams{
		Billing: dodopayments.F(dodopayments.PaymentNewParamsBilling{
			City:    dodopayments.F("city"),
			Country: dodopayments.F(dodopayments.CountryCodeAf),
			State:   dodopayments.F("state"),
			Street:  dodopayments.F("street"),
			Zipcode: dodopayments.F(int64(0)),
		}),
		Customer: dodopayments.F(dodopayments.PaymentNewParamsCustomer{
			Email: dodopayments.F("email"),
			Name:  dodopayments.F("name"),
		}),
		ProductCart: dodopayments.F([]dodopayments.PaymentNewParamsProductCart{{
			ProductID: dodopayments.F("product_id"),
			Quantity:  dodopayments.F(int64(0)),
		}}),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", payment.PaymentID)
}
