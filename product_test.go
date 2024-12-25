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
		Price: dodopayments.F[dodopayments.ProductNewParamsPriceUnion](dodopayments.ProductNewParamsPriceOneTimePrice{
			Currency:              dodopayments.F(dodopayments.ProductNewParamsPriceOneTimePriceCurrencyAed),
			Discount:              dodopayments.F(0.000000),
			Price:                 dodopayments.F(int64(0)),
			PurchasingPowerParity: dodopayments.F(true),
			Type:                  dodopayments.F(dodopayments.ProductNewParamsPriceOneTimePriceTypeOneTimePrice),
		}),
		TaxCategory:                 dodopayments.F(dodopayments.ProductNewParamsTaxCategoryDigitalProducts),
		Description:                 dodopayments.F("description"),
		LicenseKeyActivationMessage: dodopayments.F("license_key_activation_message"),
		LicenseKeyActivationsLimit:  dodopayments.F(int64(0)),
		LicenseKeyDuration: dodopayments.F(dodopayments.ProductNewParamsLicenseKeyDuration{
			Count:    dodopayments.F(int64(0)),
			Interval: dodopayments.F(dodopayments.ProductNewParamsLicenseKeyDurationIntervalDay),
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
			Description:                 dodopayments.F("description"),
			LicenseKeyActivationMessage: dodopayments.F("license_key_activation_message"),
			LicenseKeyActivationsLimit:  dodopayments.F(int64(0)),
			LicenseKeyDuration: dodopayments.F(dodopayments.ProductUpdateParamsLicenseKeyDuration{
				Count:    dodopayments.F(int64(0)),
				Interval: dodopayments.F(dodopayments.ProductUpdateParamsLicenseKeyDurationIntervalDay),
			}),
			LicenseKeyEnabled: dodopayments.F(true),
			Name:              dodopayments.F("name"),
			Price: dodopayments.F[dodopayments.ProductUpdateParamsPriceUnion](dodopayments.ProductUpdateParamsPriceOneTimePrice{
				Currency:              dodopayments.F(dodopayments.ProductUpdateParamsPriceOneTimePriceCurrencyAed),
				Discount:              dodopayments.F(0.000000),
				Price:                 dodopayments.F(int64(0)),
				PurchasingPowerParity: dodopayments.F(true),
				Type:                  dodopayments.F(dodopayments.ProductUpdateParamsPriceOneTimePriceTypeOneTimePrice),
			}),
			TaxCategory: dodopayments.F(dodopayments.ProductUpdateParamsTaxCategoryDigitalProducts),
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
