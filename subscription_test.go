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

func TestSubscriptionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.New(context.TODO(), dodopayments.SubscriptionNewParams{
		Billing: dodopayments.F(dodopayments.BillingAddressParam{
			City:    dodopayments.F("city"),
			Country: dodopayments.F(dodopayments.CountryCodeAf),
			State:   dodopayments.F("state"),
			Street:  dodopayments.F("street"),
			Zipcode: dodopayments.F("zipcode"),
		}),
		Customer: dodopayments.F[dodopayments.CustomerRequestUnionParam](dodopayments.AttachExistingCustomerParam{
			CustomerID: dodopayments.F("customer_id"),
		}),
		ProductID: dodopayments.F("product_id"),
		Quantity:  dodopayments.F(int64(0)),
		Addons: dodopayments.F([]dodopayments.SubscriptionNewParamsAddon{{
			AddonID:  dodopayments.F("addon_id"),
			Quantity: dodopayments.F(int64(0)),
		}}),
		AllowedPaymentMethodTypes: dodopayments.F([]dodopayments.SubscriptionNewParamsAllowedPaymentMethodType{dodopayments.SubscriptionNewParamsAllowedPaymentMethodTypeCredit}),
		BillingCurrency:           dodopayments.F(dodopayments.CurrencyAed),
		DiscountCode:              dodopayments.F("discount_code"),
		Metadata: dodopayments.F(map[string]string{
			"foo": "string",
		}),
		OnDemand: dodopayments.F(dodopayments.SubscriptionNewParamsOnDemand{
			MandateOnly:  dodopayments.F(true),
			ProductPrice: dodopayments.F(int64(0)),
		}),
		PaymentLink:             dodopayments.F(true),
		ReturnURL:               dodopayments.F("return_url"),
		ShowSavedPaymentMethods: dodopayments.F(true),
		TaxID:                   dodopayments.F("tax_id"),
		TrialPeriodDays:         dodopayments.F(int64(0)),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionGet(t *testing.T) {
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
	_, err := client.Subscriptions.Get(context.TODO(), "subscription_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.Update(
		context.TODO(),
		"subscription_id",
		dodopayments.SubscriptionUpdateParams{
			Billing: dodopayments.F(dodopayments.BillingAddressParam{
				City:    dodopayments.F("city"),
				Country: dodopayments.F(dodopayments.CountryCodeAf),
				State:   dodopayments.F("state"),
				Street:  dodopayments.F("street"),
				Zipcode: dodopayments.F("zipcode"),
			}),
			DisableOnDemand: dodopayments.F(dodopayments.SubscriptionUpdateParamsDisableOnDemand{
				NextBillingDate: dodopayments.F(time.Now()),
			}),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
			Status: dodopayments.F(dodopayments.SubscriptionStatusPending),
			TaxID:  dodopayments.F("tax_id"),
		},
	)
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.List(context.TODO(), dodopayments.SubscriptionListParams{
		BrandID:      dodopayments.F("brand_id"),
		CreatedAtGte: dodopayments.F(time.Now()),
		CreatedAtLte: dodopayments.F(time.Now()),
		CustomerID:   dodopayments.F("customer_id"),
		PageNumber:   dodopayments.F(int64(0)),
		PageSize:     dodopayments.F(int64(0)),
		Status:       dodopayments.F(dodopayments.SubscriptionStatusPending),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionChangePlan(t *testing.T) {
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
	err := client.Subscriptions.ChangePlan(context.TODO(), "subscription_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionChargeWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.Charge(
		context.TODO(),
		"subscription_id",
		dodopayments.SubscriptionChargeParams{
			ProductPrice: dodopayments.F(int64(0)),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
		},
	)
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
