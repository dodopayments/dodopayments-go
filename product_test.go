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
		Name: "name",
		Price: dodopayments.PriceUnionParam{
			OfOneTimePrice: &dodopayments.PriceOneTimePriceParam{
				Currency:              dodopayments.CurrencyAed,
				Discount:              0,
				Price:                 0,
				PurchasingPowerParity: true,
				Type:                  "one_time_price",
				PayWhatYouWant:        dodopayments.Bool(true),
				SuggestedPrice:        dodopayments.Int(0),
				TaxInclusive:          dodopayments.Bool(true),
			},
		},
		TaxCategory: dodopayments.TaxCategoryDigitalProducts,
		Addons:      []string{"string"},
		BrandID:     dodopayments.String("brand_id"),
		Description: dodopayments.String("description"),
		DigitalProductDelivery: dodopayments.ProductNewParamsDigitalProductDelivery{
			ExternalURL:  dodopayments.String("external_url"),
			Instructions: dodopayments.String("instructions"),
		},
		LicenseKeyActivationMessage: dodopayments.String("license_key_activation_message"),
		LicenseKeyActivationsLimit:  dodopayments.Int(0),
		LicenseKeyDuration: dodopayments.LicenseKeyDurationParam{
			Count:    0,
			Interval: dodopayments.TimeIntervalDay,
		},
		LicenseKeyEnabled: dodopayments.Bool(true),
		Metadata: map[string]string{
			"foo": "string",
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
			Addons:      []string{"string"},
			BrandID:     dodopayments.String("brand_id"),
			Description: dodopayments.String("description"),
			DigitalProductDelivery: dodopayments.ProductUpdateParamsDigitalProductDelivery{
				ExternalURL:  dodopayments.String("external_url"),
				Files:        []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
				Instructions: dodopayments.String("instructions"),
			},
			ImageID:                     dodopayments.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LicenseKeyActivationMessage: dodopayments.String("license_key_activation_message"),
			LicenseKeyActivationsLimit:  dodopayments.Int(0),
			LicenseKeyDuration: dodopayments.LicenseKeyDurationParam{
				Count:    0,
				Interval: dodopayments.TimeIntervalDay,
			},
			LicenseKeyEnabled: dodopayments.Bool(true),
			Metadata: map[string]string{
				"foo": "string",
			},
			Name: dodopayments.String("name"),
			Price: dodopayments.PriceUnionParam{
				OfOneTimePrice: &dodopayments.PriceOneTimePriceParam{
					Currency:              dodopayments.CurrencyAed,
					Discount:              0,
					Price:                 0,
					PurchasingPowerParity: true,
					Type:                  "one_time_price",
					PayWhatYouWant:        dodopayments.Bool(true),
					SuggestedPrice:        dodopayments.Int(0),
					TaxInclusive:          dodopayments.Bool(true),
				},
			},
			TaxCategory: dodopayments.TaxCategoryDigitalProducts,
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
		Archived:   dodopayments.Bool(true),
		BrandID:    dodopayments.String("brand_id"),
		PageNumber: dodopayments.Int(0),
		PageSize:   dodopayments.Int(0),
		Recurring:  dodopayments.Bool(true),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductArchive(t *testing.T) {
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
	err := client.Products.Archive(context.TODO(), "id")
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

func TestProductUpdateFiles(t *testing.T) {
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
	_, err := client.Products.UpdateFiles(
		context.TODO(),
		"id",
		dodopayments.ProductUpdateFilesParams{
			FileName: "file_name",
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
