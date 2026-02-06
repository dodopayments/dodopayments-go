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
			ProductCart: dodopayments.F([]dodopayments.ProductItemReqParam{{
				ProductID: dodopayments.F("product_id"),
				Quantity:  dodopayments.F(int64(0)),
				Addons: dodopayments.F([]dodopayments.AttachAddonParam{{
					AddonID:  dodopayments.F("addon_id"),
					Quantity: dodopayments.F(int64(0)),
				}}),
				Amount: dodopayments.F(int64(0)),
			}}),
			AllowedPaymentMethodTypes: dodopayments.F([]dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesACH}),
			BillingAddress: dodopayments.F(dodopayments.CheckoutSessionBillingAddressParam{
				Country: dodopayments.F(dodopayments.CountryCodeAf),
				City:    dodopayments.F("city"),
				State:   dodopayments.F("state"),
				Street:  dodopayments.F("street"),
				Zipcode: dodopayments.F("zipcode"),
			}),
			BillingCurrency: dodopayments.F(dodopayments.CurrencyAed),
			Confirm:         dodopayments.F(true),
			CustomFields: dodopayments.F([]dodopayments.CustomFieldParam{{
				FieldType:   dodopayments.F(dodopayments.CustomFieldFieldTypeText),
				Key:         dodopayments.F("key"),
				Label:       dodopayments.F("label"),
				Options:     dodopayments.F([]string{"string"}),
				Placeholder: dodopayments.F("placeholder"),
				Required:    dodopayments.F(true),
			}}),
			Customer: dodopayments.F[dodopayments.CustomerRequestUnionParam](dodopayments.AttachExistingCustomerParam{
				CustomerID: dodopayments.F("customer_id"),
			}),
			Customization: dodopayments.F(dodopayments.CheckoutSessionCustomizationParam{
				ForceLanguage:    dodopayments.F("force_language"),
				ShowOnDemandTag:  dodopayments.F(true),
				ShowOrderDetails: dodopayments.F(true),
				Theme:            dodopayments.F(dodopayments.CheckoutSessionCustomizationThemeDark),
				ThemeConfig: dodopayments.F(dodopayments.ThemeConfigParam{
					Dark: dodopayments.F(dodopayments.ThemeModeConfigParam{
						BgPrimary:            dodopayments.F("bg_primary"),
						BgSecondary:          dodopayments.F("bg_secondary"),
						BorderPrimary:        dodopayments.F("border_primary"),
						BorderSecondary:      dodopayments.F("border_secondary"),
						ButtonPrimary:        dodopayments.F("button_primary"),
						ButtonPrimaryHover:   dodopayments.F("button_primary_hover"),
						ButtonSecondary:      dodopayments.F("button_secondary"),
						ButtonSecondaryHover: dodopayments.F("button_secondary_hover"),
						ButtonTextPrimary:    dodopayments.F("button_text_primary"),
						ButtonTextSecondary:  dodopayments.F("button_text_secondary"),
						InputFocusBorder:     dodopayments.F("input_focus_border"),
						TextError:            dodopayments.F("text_error"),
						TextPlaceholder:      dodopayments.F("text_placeholder"),
						TextPrimary:          dodopayments.F("text_primary"),
						TextSecondary:        dodopayments.F("text_secondary"),
						TextSuccess:          dodopayments.F("text_success"),
					}),
					FontPrimaryURL:   dodopayments.F("font_primary_url"),
					FontSecondaryURL: dodopayments.F("font_secondary_url"),
					FontSize:         dodopayments.F(dodopayments.ThemeConfigFontSizeXs),
					FontWeight:       dodopayments.F(dodopayments.ThemeConfigFontWeightNormal),
					Light: dodopayments.F(dodopayments.ThemeModeConfigParam{
						BgPrimary:            dodopayments.F("bg_primary"),
						BgSecondary:          dodopayments.F("bg_secondary"),
						BorderPrimary:        dodopayments.F("border_primary"),
						BorderSecondary:      dodopayments.F("border_secondary"),
						ButtonPrimary:        dodopayments.F("button_primary"),
						ButtonPrimaryHover:   dodopayments.F("button_primary_hover"),
						ButtonSecondary:      dodopayments.F("button_secondary"),
						ButtonSecondaryHover: dodopayments.F("button_secondary_hover"),
						ButtonTextPrimary:    dodopayments.F("button_text_primary"),
						ButtonTextSecondary:  dodopayments.F("button_text_secondary"),
						InputFocusBorder:     dodopayments.F("input_focus_border"),
						TextError:            dodopayments.F("text_error"),
						TextPlaceholder:      dodopayments.F("text_placeholder"),
						TextPrimary:          dodopayments.F("text_primary"),
						TextSecondary:        dodopayments.F("text_secondary"),
						TextSuccess:          dodopayments.F("text_success"),
					}),
					PayButtonText: dodopayments.F("pay_button_text"),
					Radius:        dodopayments.F("radius"),
				}),
			}),
			DiscountCode: dodopayments.F("discount_code"),
			FeatureFlags: dodopayments.F(dodopayments.CheckoutSessionFlagsParam{
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
			PaymentMethodID:         dodopayments.F("payment_method_id"),
			ProductCollectionID:     dodopayments.F("product_collection_id"),
			ReturnURL:               dodopayments.F("return_url"),
			ShortLink:               dodopayments.F(true),
			ShowSavedPaymentMethods: dodopayments.F(true),
			SubscriptionData: dodopayments.F(dodopayments.SubscriptionDataParam{
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

func TestCheckoutSessionPreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.CheckoutSessions.Preview(context.TODO(), dodopayments.CheckoutSessionPreviewParams{
		CheckoutSessionRequest: dodopayments.CheckoutSessionRequestParam{
			ProductCart: dodopayments.F([]dodopayments.ProductItemReqParam{{
				ProductID: dodopayments.F("product_id"),
				Quantity:  dodopayments.F(int64(0)),
				Addons: dodopayments.F([]dodopayments.AttachAddonParam{{
					AddonID:  dodopayments.F("addon_id"),
					Quantity: dodopayments.F(int64(0)),
				}}),
				Amount: dodopayments.F(int64(0)),
			}}),
			AllowedPaymentMethodTypes: dodopayments.F([]dodopayments.PaymentMethodTypes{dodopayments.PaymentMethodTypesACH}),
			BillingAddress: dodopayments.F(dodopayments.CheckoutSessionBillingAddressParam{
				Country: dodopayments.F(dodopayments.CountryCodeAf),
				City:    dodopayments.F("city"),
				State:   dodopayments.F("state"),
				Street:  dodopayments.F("street"),
				Zipcode: dodopayments.F("zipcode"),
			}),
			BillingCurrency: dodopayments.F(dodopayments.CurrencyAed),
			Confirm:         dodopayments.F(true),
			CustomFields: dodopayments.F([]dodopayments.CustomFieldParam{{
				FieldType:   dodopayments.F(dodopayments.CustomFieldFieldTypeText),
				Key:         dodopayments.F("key"),
				Label:       dodopayments.F("label"),
				Options:     dodopayments.F([]string{"string"}),
				Placeholder: dodopayments.F("placeholder"),
				Required:    dodopayments.F(true),
			}}),
			Customer: dodopayments.F[dodopayments.CustomerRequestUnionParam](dodopayments.AttachExistingCustomerParam{
				CustomerID: dodopayments.F("customer_id"),
			}),
			Customization: dodopayments.F(dodopayments.CheckoutSessionCustomizationParam{
				ForceLanguage:    dodopayments.F("force_language"),
				ShowOnDemandTag:  dodopayments.F(true),
				ShowOrderDetails: dodopayments.F(true),
				Theme:            dodopayments.F(dodopayments.CheckoutSessionCustomizationThemeDark),
				ThemeConfig: dodopayments.F(dodopayments.ThemeConfigParam{
					Dark: dodopayments.F(dodopayments.ThemeModeConfigParam{
						BgPrimary:            dodopayments.F("bg_primary"),
						BgSecondary:          dodopayments.F("bg_secondary"),
						BorderPrimary:        dodopayments.F("border_primary"),
						BorderSecondary:      dodopayments.F("border_secondary"),
						ButtonPrimary:        dodopayments.F("button_primary"),
						ButtonPrimaryHover:   dodopayments.F("button_primary_hover"),
						ButtonSecondary:      dodopayments.F("button_secondary"),
						ButtonSecondaryHover: dodopayments.F("button_secondary_hover"),
						ButtonTextPrimary:    dodopayments.F("button_text_primary"),
						ButtonTextSecondary:  dodopayments.F("button_text_secondary"),
						InputFocusBorder:     dodopayments.F("input_focus_border"),
						TextError:            dodopayments.F("text_error"),
						TextPlaceholder:      dodopayments.F("text_placeholder"),
						TextPrimary:          dodopayments.F("text_primary"),
						TextSecondary:        dodopayments.F("text_secondary"),
						TextSuccess:          dodopayments.F("text_success"),
					}),
					FontPrimaryURL:   dodopayments.F("font_primary_url"),
					FontSecondaryURL: dodopayments.F("font_secondary_url"),
					FontSize:         dodopayments.F(dodopayments.ThemeConfigFontSizeXs),
					FontWeight:       dodopayments.F(dodopayments.ThemeConfigFontWeightNormal),
					Light: dodopayments.F(dodopayments.ThemeModeConfigParam{
						BgPrimary:            dodopayments.F("bg_primary"),
						BgSecondary:          dodopayments.F("bg_secondary"),
						BorderPrimary:        dodopayments.F("border_primary"),
						BorderSecondary:      dodopayments.F("border_secondary"),
						ButtonPrimary:        dodopayments.F("button_primary"),
						ButtonPrimaryHover:   dodopayments.F("button_primary_hover"),
						ButtonSecondary:      dodopayments.F("button_secondary"),
						ButtonSecondaryHover: dodopayments.F("button_secondary_hover"),
						ButtonTextPrimary:    dodopayments.F("button_text_primary"),
						ButtonTextSecondary:  dodopayments.F("button_text_secondary"),
						InputFocusBorder:     dodopayments.F("input_focus_border"),
						TextError:            dodopayments.F("text_error"),
						TextPlaceholder:      dodopayments.F("text_placeholder"),
						TextPrimary:          dodopayments.F("text_primary"),
						TextSecondary:        dodopayments.F("text_secondary"),
						TextSuccess:          dodopayments.F("text_success"),
					}),
					PayButtonText: dodopayments.F("pay_button_text"),
					Radius:        dodopayments.F("radius"),
				}),
			}),
			DiscountCode: dodopayments.F("discount_code"),
			FeatureFlags: dodopayments.F(dodopayments.CheckoutSessionFlagsParam{
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
			PaymentMethodID:         dodopayments.F("payment_method_id"),
			ProductCollectionID:     dodopayments.F("product_collection_id"),
			ReturnURL:               dodopayments.F("return_url"),
			ShortLink:               dodopayments.F(true),
			ShowSavedPaymentMethods: dodopayments.F(true),
			SubscriptionData: dodopayments.F(dodopayments.SubscriptionDataParam{
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
