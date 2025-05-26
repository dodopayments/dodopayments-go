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

func TestProductNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.New(context.TODO(), dodopayments.ProductNewParams{
		Price: dodopayments.F[dodopayments.PriceUnionParam](dodopayments.PriceOneTimePriceParam{
			Currency:              dodopayments.F(dodopayments.CurrencyAed),
			Discount:              dodopayments.F(0.000000),
			Price:                 dodopayments.F(int64(0)),
			PurchasingPowerParity: dodopayments.F(true),
			Type:                  dodopayments.F(dodopayments.PriceOneTimePriceTypeOneTimePrice),
			PayWhatYouWant:        dodopayments.F(true),
			SuggestedPrice:        dodopayments.F(int64(0)),
			TaxInclusive:          dodopayments.F(true),
		}),
		TaxCategory:                 dodopayments.F(dodopayments.TaxCategoryDigitalProducts),
		Addons:                      dodopayments.F([]string{"string"}),
		BrandID:                     dodopayments.F("brand_id"),
		Description:                 dodopayments.F("description"),
		LicenseKeyActivationMessage: dodopayments.F("license_key_activation_message"),
		LicenseKeyActivationsLimit:  dodopayments.F(int64(0)),
		LicenseKeyDuration: dodopayments.F(dodopayments.LicenseKeyDurationParam{
			Count:    dodopayments.F(int64(0)),
			Interval: dodopayments.F(dodopayments.TimeIntervalDay),
		}),
		LicenseKeyEnabled: dodopayments.F(true),
		Name:              dodopayments.F("name"),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductGet(t *testing.T) {
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
	_, err := client.Products.Get(context.TODO(), "id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Products.Update(
		context.TODO(),
		"id",
		dodopayments.ProductUpdateParams{
			Addons:                      dodopayments.F([]string{"string"}),
			BrandID:                     dodopayments.F("brand_id"),
			Description:                 dodopayments.F("description"),
			ImageID:                     dodopayments.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LicenseKeyActivationMessage: dodopayments.F("license_key_activation_message"),
			LicenseKeyActivationsLimit:  dodopayments.F(int64(0)),
			LicenseKeyDuration: dodopayments.F(dodopayments.LicenseKeyDurationParam{
				Count:    dodopayments.F(int64(0)),
				Interval: dodopayments.F(dodopayments.TimeIntervalDay),
			}),
			LicenseKeyEnabled: dodopayments.F(true),
			Name:              dodopayments.F("name"),
			Price: dodopayments.F[dodopayments.PriceUnionParam](dodopayments.PriceOneTimePriceParam{
				Currency:              dodopayments.F(dodopayments.CurrencyAed),
				Discount:              dodopayments.F(0.000000),
				Price:                 dodopayments.F(int64(0)),
				PurchasingPowerParity: dodopayments.F(true),
				Type:                  dodopayments.F(dodopayments.PriceOneTimePriceTypeOneTimePrice),
				PayWhatYouWant:        dodopayments.F(true),
				SuggestedPrice:        dodopayments.F(int64(0)),
				TaxInclusive:          dodopayments.F(true),
			}),
			TaxCategory: dodopayments.F(dodopayments.TaxCategoryDigitalProducts),
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

func TestProductListWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.List(context.TODO(), dodopayments.ProductListParams{
		Archived:   dodopayments.F(true),
		BrandID:    dodopayments.F("brand_id"),
		PageNumber: dodopayments.F(int64(0)),
		PageSize:   dodopayments.F(int64(0)),
		Recurring:  dodopayments.F(true),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductDelete(t *testing.T) {
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
	err := client.Products.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductUnarchive(t *testing.T) {
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
	err := client.Products.Unarchive(context.TODO(), "id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
