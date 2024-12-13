// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/dodo-payments-go"
	"github.com/stainless-sdks/dodo-payments-go/internal/testutil"
	"github.com/stainless-sdks/dodo-payments-go/option"
)

func TestPaymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.New(context.TODO(), dodopayments.PaymentNewParams{
		Billing: dodopayments.F(dodopayments.PaymentNewParamsBilling{
			City:    dodopayments.F("city"),
			Country: dodopayments.F(dodopayments.CountryCodeAlpha2Af),
			State:   dodopayments.F("state"),
			Street:  dodopayments.F("street"),
			Zipcode: dodopayments.F(int64(0)),
		}),
		Customer: dodopayments.F(dodopayments.PaymentNewParamsCustomer{
			Email:       dodopayments.F("email"),
			Name:        dodopayments.F("name"),
			PhoneNumber: dodopayments.F("phone_number"),
		}),
		ProductCart: dodopayments.F([]dodopayments.PaymentNewParamsProductCart{{
			ProductID: dodopayments.F("product_id"),
			Quantity:  dodopayments.F(int64(0)),
		}}),
		PaymentLink: dodopayments.F(true),
		ReturnURL:   dodopayments.F("return_url"),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentGet(t *testing.T) {
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
	_, err := client.Payments.Get(context.TODO(), "payment_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.List(context.TODO(), dodopayments.PaymentListParams{
		PageNumber: dodopayments.F(int64(0)),
		PageSize:   dodopayments.F(int64(0)),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
