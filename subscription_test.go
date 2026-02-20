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
			Country: dodopayments.F(dodopayments.CountryCodeAf),
			City:    dodopayments.F("city"),
			State:   dodopayments.F("state"),
			Street:  dodopayments.F("street"),
			Zipcode: dodopayments.F("zipcode"),
		}),
		Customer: dodopayments.F[dodopayments.CustomerRequestUnionParam](dodopayments.AttachExistingCustomerParam{
			CustomerID: dodopayments.F("customer_id"),
		}),
		ProductID: dodopayments.F("product_id"),
		Quantity:  dodopayments.F(int64(0)),
		Addons: dodopayments.F([]dodopayments.AttachAddonParam{{
			AddonID:  dodopayments.F("addon_id"),
			Quantity: dodopayments.F(int64(0)),
		}}),
		AllowedPaymentMethodTypes: dodopayments.F([]dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesACH}),
		BillingCurrency:           dodopayments.F(dodopayments.CurrencyAed),
		DiscountCode:              dodopayments.F("discount_code"),
		Force3DS:                  dodopayments.F(true),
		Metadata: dodopayments.F(map[string]string{
			"foo": "string",
		}),
		OnDemand: dodopayments.F(dodopayments.OnDemandSubscriptionParam{
			MandateOnly:                   dodopayments.F(true),
			AdaptiveCurrencyFeesInclusive: dodopayments.F(true),
			ProductCurrency:               dodopayments.F(dodopayments.CurrencyAed),
			ProductDescription:            dodopayments.F("product_description"),
			ProductPrice:                  dodopayments.F(int64(0)),
		}),
		OneTimeProductCart: dodopayments.F([]dodopayments.OneTimeProductCartItemParam{{
			ProductID: dodopayments.F("product_id"),
			Quantity:  dodopayments.F(int64(0)),
			Amount:    dodopayments.F(int64(0)),
		}}),
		PaymentLink:             dodopayments.F(true),
		PaymentMethodID:         dodopayments.F("payment_method_id"),
		RedirectImmediately:     dodopayments.F(true),
		ReturnURL:               dodopayments.F("return_url"),
		ShortLink:               dodopayments.F(true),
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
				Country: dodopayments.F(dodopayments.CountryCodeAf),
				City:    dodopayments.F("city"),
				State:   dodopayments.F("state"),
				Street:  dodopayments.F("street"),
				Zipcode: dodopayments.F("zipcode"),
			}),
			CancelAtNextBillingDate: dodopayments.F(true),
			CreditEntitlementCart: dodopayments.F([]dodopayments.SubscriptionUpdateParamsCreditEntitlementCart{{
				CreditEntitlementID:        dodopayments.F("credit_entitlement_id"),
				CreditsAmount:              dodopayments.F("credits_amount"),
				ExpiresAfterDays:           dodopayments.F(int64(0)),
				LowBalanceThresholdPercent: dodopayments.F(int64(0)),
				MaxRolloverCount:           dodopayments.F(int64(0)),
				OverageChargeAtBilling:     dodopayments.F(true),
				OverageEnabled:             dodopayments.F(true),
				OverageLimit:               dodopayments.F("overage_limit"),
				RolloverEnabled:            dodopayments.F(true),
				RolloverPercentage:         dodopayments.F(int64(0)),
				RolloverTimeframeCount:     dodopayments.F(int64(0)),
				RolloverTimeframeInterval:  dodopayments.F(dodopayments.TimeIntervalDay),
			}}),
			CustomerName: dodopayments.F("customer_name"),
			DisableOnDemand: dodopayments.F(dodopayments.SubscriptionUpdateParamsDisableOnDemand{
				NextBillingDate: dodopayments.F(time.Now()),
			}),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
			NextBillingDate: dodopayments.F(time.Now()),
			Status:          dodopayments.F(dodopayments.SubscriptionStatusPending),
			TaxID:           dodopayments.F("tax_id"),
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
		ProductID:    dodopayments.F("product_id"),
		Status:       dodopayments.F(dodopayments.SubscriptionListParamsStatusPending),
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
			ProductID:            dodopayments.F("product_id"),
			ProrationBillingMode: dodopayments.F(dodopayments.SubscriptionChangePlanParamsProrationBillingModeProratedImmediately),
			Quantity:             dodopayments.F(int64(0)),
			Addons: dodopayments.F([]dodopayments.AttachAddonParam{{
				AddonID:  dodopayments.F("addon_id"),
				Quantity: dodopayments.F(int64(0)),
			}}),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
			OnPaymentFailure: dodopayments.F(dodopayments.SubscriptionChangePlanParamsOnPaymentFailurePreventChange),
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
			ProductPrice:                  dodopayments.F(int64(0)),
			AdaptiveCurrencyFeesInclusive: dodopayments.F(true),
			CustomerBalanceConfig: dodopayments.F(dodopayments.SubscriptionChargeParamsCustomerBalanceConfig{
				AllowCustomerCreditsPurchase: dodopayments.F(true),
				AllowCustomerCreditsUsage:    dodopayments.F(true),
			}),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
			ProductCurrency:    dodopayments.F(dodopayments.CurrencyAed),
			ProductDescription: dodopayments.F("product_description"),
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
			ProductID:            dodopayments.F("product_id"),
			ProrationBillingMode: dodopayments.F(dodopayments.SubscriptionPreviewChangePlanParamsProrationBillingModeProratedImmediately),
			Quantity:             dodopayments.F(int64(0)),
			Addons: dodopayments.F([]dodopayments.AttachAddonParam{{
				AddonID:  dodopayments.F("addon_id"),
				Quantity: dodopayments.F(int64(0)),
			}}),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
			OnPaymentFailure: dodopayments.F(dodopayments.SubscriptionPreviewChangePlanParamsOnPaymentFailurePreventChange),
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
			EndDate:    dodopayments.F(time.Now()),
			MeterID:    dodopayments.F("meter_id"),
			PageNumber: dodopayments.F(int64(0)),
			PageSize:   dodopayments.F(int64(0)),
			StartDate:  dodopayments.F(time.Now()),
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
			Body: dodopayments.SubscriptionUpdatePaymentMethodParamsBodyNew{
				Type:      dodopayments.F(dodopayments.SubscriptionUpdatePaymentMethodParamsBodyNewTypeNew),
				ReturnURL: dodopayments.F("return_url"),
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
