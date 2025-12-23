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
		ProductID: "product_id",
		Quantity:  0,
		Addons: []dodopayments.AttachAddonParam{{
			AddonID:  "addon_id",
			Quantity: 0,
		}},
		AllowedPaymentMethodTypes: []dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesCredit},
		BillingCurrency:           dodopayments.CurrencyAed,
		DiscountCode:              dodopayments.String("discount_code"),
		Force3DS:                  dodopayments.Bool(true),
		Metadata: map[string]string{
			"foo": "string",
		},
		OnDemand: dodopayments.OnDemandSubscriptionParam{
			MandateOnly:                   true,
			AdaptiveCurrencyFeesInclusive: dodopayments.Bool(true),
			ProductCurrency:               dodopayments.CurrencyAed,
			ProductDescription:            dodopayments.String("product_description"),
			ProductPrice:                  dodopayments.Int(0),
		},
		OneTimeProductCart: []dodopayments.OneTimeProductCartItemParam{{
			ProductID: "product_id",
			Quantity:  0,
			Amount:    dodopayments.Int(0),
		}},
		PaymentLink:             dodopayments.Bool(true),
		RedirectImmediately:     dodopayments.Bool(true),
		ReturnURL:               dodopayments.String("return_url"),
		ShortLink:               dodopayments.Bool(true),
		ShowSavedPaymentMethods: dodopayments.Bool(true),
		TaxID:                   dodopayments.String("tax_id"),
		TrialPeriodDays:         dodopayments.Int(0),
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
			Billing: dodopayments.BillingAddressParam{
				Country: dodopayments.CountryCodeAf,
				City:    dodopayments.String("city"),
				State:   dodopayments.String("state"),
				Street:  dodopayments.String("street"),
				Zipcode: dodopayments.String("zipcode"),
			},
			CancelAtNextBillingDate: dodopayments.Bool(true),
			CustomerName:            dodopayments.String("customer_name"),
			DisableOnDemand: dodopayments.SubscriptionUpdateParamsDisableOnDemand{
				NextBillingDate: time.Now(),
			},
			Metadata: map[string]string{
				"foo": "string",
			},
			NextBillingDate: dodopayments.Time(time.Now()),
			Status:          dodopayments.SubscriptionStatusPending,
			TaxID:           dodopayments.String("tax_id"),
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
		BrandID:      dodopayments.String("brand_id"),
		CreatedAtGte: dodopayments.Time(time.Now()),
		CreatedAtLte: dodopayments.Time(time.Now()),
		CustomerID:   dodopayments.String("customer_id"),
		PageNumber:   dodopayments.Int(0),
		PageSize:     dodopayments.Int(0),
		Status:       dodopayments.SubscriptionListParamsStatusPending,
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSubscriptionChangePlanWithOptionalParams(t *testing.T) {
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
	err := client.Subscriptions.ChangePlan(
		context.TODO(),
		"subscription_id",
		dodopayments.SubscriptionChangePlanParams{
			ProductID:            "product_id",
			ProrationBillingMode: dodopayments.SubscriptionChangePlanParamsProrationBillingModeProratedImmediately,
			Quantity:             0,
			Addons: []dodopayments.AttachAddonParam{{
				AddonID:  "addon_id",
				Quantity: 0,
			}},
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
			ProductPrice:                  0,
			AdaptiveCurrencyFeesInclusive: dodopayments.Bool(true),
			CustomerBalanceConfig: dodopayments.SubscriptionChargeParamsCustomerBalanceConfig{
				AllowCustomerCreditsPurchase: dodopayments.Bool(true),
				AllowCustomerCreditsUsage:    dodopayments.Bool(true),
			},
			Metadata: map[string]string{
				"foo": "string",
			},
			ProductCurrency:    dodopayments.CurrencyAed,
			ProductDescription: dodopayments.String("product_description"),
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

func TestSubscriptionPreviewChangePlanWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.PreviewChangePlan(
		context.TODO(),
		"subscription_id",
		dodopayments.SubscriptionPreviewChangePlanParams{
			ProductID:            "product_id",
			ProrationBillingMode: dodopayments.SubscriptionPreviewChangePlanParamsProrationBillingModeProratedImmediately,
			Quantity:             0,
			Addons: []dodopayments.AttachAddonParam{{
				AddonID:  "addon_id",
				Quantity: 0,
			}},
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

func TestSubscriptionGetUsageHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.GetUsageHistory(
		context.TODO(),
		"subscription_id",
		dodopayments.SubscriptionGetUsageHistoryParams{
			EndDate:    dodopayments.Time(time.Now()),
			MeterID:    dodopayments.String("meter_id"),
			PageNumber: dodopayments.Int(0),
			PageSize:   dodopayments.Int(0),
			StartDate:  dodopayments.Time(time.Now()),
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

func TestSubscriptionUpdatePaymentMethodWithOptionalParams(t *testing.T) {
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
	_, err := client.Subscriptions.UpdatePaymentMethod(
		context.TODO(),
		"subscription_id",
		dodopayments.SubscriptionUpdatePaymentMethodParams{
			OfNew: &dodopayments.SubscriptionUpdatePaymentMethodParamsBodyNew{
				Type:      "new",
				ReturnURL: dodopayments.String("return_url"),
			},
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
