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
			ProductCart: dodopayments.F([]dodopayments.CheckoutSessionRequestProductCartParam{{
				ProductID: dodopayments.F("product_id"),
				Quantity:  dodopayments.F(int64(0)),
				Addons: dodopayments.F([]dodopayments.AttachAddonParam{{
					AddonID:  dodopayments.F("addon_id"),
					Quantity: dodopayments.F(int64(0)),
				}}),
				Amount: dodopayments.F(int64(0)),
			}}),
			AllowedPaymentMethodTypes: dodopayments.F([]dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesCredit}),
			BillingAddress: dodopayments.F(dodopayments.CheckoutSessionRequestBillingAddressParam{
				Country: dodopayments.F(dodopayments.CountryCodeAf),
				City:    dodopayments.F("city"),
				State:   dodopayments.F("state"),
				Street:  dodopayments.F("street"),
				Zipcode: dodopayments.F("zipcode"),
			}),
			BillingCurrency: dodopayments.F(dodopayments.CurrencyAed),
			Confirm:         dodopayments.F(true),
			Customer: dodopayments.F[dodopayments.CustomerRequestUnionParam](dodopayments.AttachExistingCustomerParam{
				CustomerID: dodopayments.F("customer_id"),
			}),
			Customization: dodopayments.F(dodopayments.CheckoutSessionRequestCustomizationParam{
				ForceLanguage:    dodopayments.F("force_language"),
				ShowOnDemandTag:  dodopayments.F(true),
				ShowOrderDetails: dodopayments.F(true),
				Theme:            dodopayments.F(dodopayments.CheckoutSessionRequestCustomizationThemeDark),
			}),
			DiscountCode: dodopayments.F("discount_code"),
			FeatureFlags: dodopayments.F(dodopayments.CheckoutSessionRequestFeatureFlagsParam{
				AllowCurrencySelection:      dodopayments.F(true),
				AllowCustomerEditingCity:    dodopayments.F(true),
				AllowCustomerEditingCountry: dodopayments.F(true),
				AllowCustomerEditingEmail:   dodopayments.F(true),
				AllowCustomerEditingName:    dodopayments.F(true),
				AllowCustomerEditingState:   dodopayments.F(true),
				AllowCustomerEditingStreet:  dodopayments.F(true),
				AllowCustomerEditingZipcode: dodopayments.F(true),
				AllowDiscountCode:           dodopayments.F(true),
				AllowPhoneNumberCollection:  dodopayments.F(true),
				AllowTaxID:                  dodopayments.F(true),
				AlwaysCreateNewCustomer:     dodopayments.F(true),
				RedirectImmediately:         dodopayments.F(true),
			}),
			Force3DS: dodopayments.F(true),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
			MinimalAddress:          dodopayments.F(true),
			ReturnURL:               dodopayments.F("return_url"),
			ShortLink:               dodopayments.F(true),
			ShowSavedPaymentMethods: dodopayments.F(true),
			SubscriptionData: dodopayments.F(dodopayments.CheckoutSessionRequestSubscriptionDataParam{
				OnDemand: dodopayments.F(dodopayments.OnDemandSubscriptionParam{
					MandateOnly:                   dodopayments.F(true),
					AdaptiveCurrencyFeesInclusive: dodopayments.F(true),
					ProductCurrency:               dodopayments.F(dodopayments.CurrencyAed),
					ProductDescription:            dodopayments.F("product_description"),
					ProductPrice:                  dodopayments.F(int64(0)),
				}),
				TrialPeriodDays: dodopayments.F(int64(0)),
			}),
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
