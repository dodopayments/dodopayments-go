// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/internal/testutil"
	"github.com/dodopayments/dodopayments-go/option"
)

func TestCheckoutSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CheckoutSessions.New(context.TODO(), dodopayments.CheckoutSessionNewParams{
		CheckoutSessionRequest: dodopayments.CheckoutSessionRequestParam{
			ProductCart: []dodopayments.CheckoutSessionRequestProductCartParam{{
				ProductID: "product_id",
				Quantity:  0,
				Addons: []dodopayments.AttachAddonParam{{
					AddonID:  "addon_id",
					Quantity: 0,
				}},
				Amount: dodopayments.Int(0),
			}},
			AllowedPaymentMethodTypes: []dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesCredit},
			BillingAddress: dodopayments.CheckoutSessionRequestBillingAddressParam{
				Country: dodopayments.CountryCodeAf,
				City:    dodopayments.String("city"),
				State:   dodopayments.String("state"),
				Street:  dodopayments.String("street"),
				Zipcode: dodopayments.String("zipcode"),
			},
			BillingCurrency: dodopayments.CurrencyAed,
			Confirm:         dodopayments.Bool(true),
			Customer: dodopayments.CustomerRequestUnionParam{
				OfAttachExistingCustomer: &dodopayments.AttachExistingCustomerParam{
					CustomerID: "customer_id",
				},
			},
			Customization: dodopayments.CheckoutSessionRequestCustomizationParam{
				ForceLanguage:    dodopayments.String("force_language"),
				ShowOnDemandTag:  dodopayments.Bool(true),
				ShowOrderDetails: dodopayments.Bool(true),
				Theme:            "dark",
			},
			DiscountCode: dodopayments.String("discount_code"),
			FeatureFlags: dodopayments.CheckoutSessionRequestFeatureFlagsParam{
				AllowCurrencySelection:      dodopayments.Bool(true),
				AllowCustomerEditingCity:    dodopayments.Bool(true),
				AllowCustomerEditingCountry: dodopayments.Bool(true),
				AllowCustomerEditingEmail:   dodopayments.Bool(true),
				AllowCustomerEditingName:    dodopayments.Bool(true),
				AllowCustomerEditingState:   dodopayments.Bool(true),
				AllowCustomerEditingStreet:  dodopayments.Bool(true),
				AllowCustomerEditingZipcode: dodopayments.Bool(true),
				AllowDiscountCode:           dodopayments.Bool(true),
				AllowPhoneNumberCollection:  dodopayments.Bool(true),
				AllowTaxID:                  dodopayments.Bool(true),
				AlwaysCreateNewCustomer:     dodopayments.Bool(true),
				RedirectImmediately:         dodopayments.Bool(true),
			},
			Force3DS: dodopayments.Bool(true),
			Metadata: map[string]string{
				"foo": "string",
			},
			MinimalAddress:          dodopayments.Bool(true),
			ReturnURL:               dodopayments.String("return_url"),
			ShortLink:               dodopayments.Bool(true),
			ShowSavedPaymentMethods: dodopayments.Bool(true),
			SubscriptionData: dodopayments.CheckoutSessionRequestSubscriptionDataParam{
				OnDemand: dodopayments.OnDemandSubscriptionParam{
					MandateOnly:                   true,
					AdaptiveCurrencyFeesInclusive: dodopayments.Bool(true),
					ProductCurrency:               dodopayments.CurrencyAed,
					ProductDescription:            dodopayments.String("product_description"),
					ProductPrice:                  dodopayments.Int(0),
				},
				TrialPeriodDays: dodopayments.Int(0),
			},
		},
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCheckoutSessionGet(t *testing.T) {
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
	_, err := client.CheckoutSessions.Get(context.TODO(), "id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
