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
		Billing: dodopayments.BillingAddressParam{
			Country: dodopayments.CountryCodeAf,
			City:    dodopayments.String("city"),
			State:   dodopayments.String("state"),
			Street:  dodopayments.String("street"),
			Zipcode: dodopayments.String("zipcode"),
		},
		Customer: dodopayments.CustomerRequestUnionParam{
			OfAttachExistingCustomer: &dodopayments.AttachExistingCustomerParam{
				CustomerID: "customer_id",
			},
		},
		ProductCart: []dodopayments.OneTimeProductCartItemParam{{
			ProductID: "product_id",
			Quantity:  0,
			Amount:    dodopayments.Int(0),
		}},
		AllowedPaymentMethodTypes: []dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesCredit},
		BillingCurrency:           dodopayments.CurrencyAed,
		DiscountCode:              dodopayments.String("discount_code"),
		Force3DS:                  dodopayments.Bool(true),
		Metadata: map[string]string{
			"foo": "string",
		},
		PaymentLink:             dodopayments.Bool(true),
		RedirectImmediately:     dodopayments.Bool(true),
		ReturnURL:               dodopayments.String("return_url"),
		ShortLink:               dodopayments.Bool(true),
		ShowSavedPaymentMethods: dodopayments.Bool(true),
		TaxID:                   dodopayments.String("tax_id"),
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
		BrandID:        dodopayments.String("brand_id"),
		CreatedAtGte:   dodopayments.Time(time.Now()),
		CreatedAtLte:   dodopayments.Time(time.Now()),
		CustomerID:     dodopayments.String("customer_id"),
		PageNumber:     dodopayments.Int(0),
		PageSize:       dodopayments.Int(0),
		Status:         dodopayments.PaymentListParamsStatusSucceeded,
		SubscriptionID: dodopayments.String("subscription_id"),
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
