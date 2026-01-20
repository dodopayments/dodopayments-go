// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/internal/testutil"
	"github.com/dodopayments/dodopayments-go/option"
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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Payments.New(context.TODO(), dodopayments.PaymentNewParams{
		Billing: dodopayments.F(dodopayments.BillingAddressParam{
			Country: dodopayments.F(dodopayments.CountryCodeAf),
			City:    dodopayments.F("city"),
			State:   dodopayments.F("state"),
			Street:  dodopayments.F("street"),
			Zipcode: dodopayments.F("zipcode"),
		}),
		Customer: dodopayments.F[dodopayments.CustomerRequestUnionParam](dodopayments.AttachExistingCustomerParam{
			CustomerID: dodopayments.F("customer_id"),
		}),
		ProductCart: dodopayments.F([]dodopayments.OneTimeProductCartItemParam{{
			ProductID: dodopayments.F("product_id"),
			Quantity:  dodopayments.F(int64(0)),
			Amount:    dodopayments.F(int64(0)),
		}}),
		AllowedPaymentMethodTypes: dodopayments.F([]dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesCredit}),
		BillingCurrency:           dodopayments.F(dodopayments.CurrencyAed),
		DiscountCode:              dodopayments.F("discount_code"),
		Force3DS:                  dodopayments.F(true),
		Metadata: dodopayments.F(map[string]string{
			"foo": "string",
		}),
		PaymentLink:             dodopayments.F(true),
		PaymentMethodID:         dodopayments.F("payment_method_id"),
		RedirectImmediately:     dodopayments.F(true),
		ReturnURL:               dodopayments.F("return_url"),
		ShortLink:               dodopayments.F(true),
		ShowSavedPaymentMethods: dodopayments.F(true),
		TaxID:                   dodopayments.F("tax_id"),
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
		option.WithBearerToken("My Bearer Token"),
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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Payments.List(context.TODO(), dodopayments.PaymentListParams{
		BrandID:        dodopayments.F("brand_id"),
		CreatedAtGte:   dodopayments.F(time.Now()),
		CreatedAtLte:   dodopayments.F(time.Now()),
		CustomerID:     dodopayments.F("customer_id"),
		PageNumber:     dodopayments.F(int64(0)),
		PageSize:       dodopayments.F(int64(0)),
		Status:         dodopayments.F(dodopayments.PaymentListParamsStatusSucceeded),
		SubscriptionID: dodopayments.F("subscription_id"),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentGetLineItems(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dodopayments.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Payments.GetLineItems(context.TODO(), "payment_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
