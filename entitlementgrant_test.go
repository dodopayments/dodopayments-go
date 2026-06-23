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

func TestEntitlementGrantListWithOptionalParams(t *testing.T) {
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
	_, err := client.Entitlements.Grants.List(
		context.TODO(),
		"ent_jt7jcvI79Xh8eehqgWdcm",
		dodopayments.EntitlementGrantListParams{
			CustomerID: dodopayments.F("customer_id"),
			PageNumber: dodopayments.F(int64(0)),
			PageSize:   dodopayments.F(int64(0)),
			Status:     dodopayments.F(dodopayments.EntitlementGrantListParamsStatusPending),
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

func TestEntitlementGrantFulfillLicenseKeyWithOptionalParams(t *testing.T) {
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
	_, err := client.Entitlements.Grants.FulfillLicenseKey(
		context.TODO(),
		"entg_w0ZCJZgNXuNDdMVzvja6p",
		dodopayments.EntitlementGrantFulfillLicenseKeyParams{
			Key:              dodopayments.F("key"),
			ActivationsLimit: dodopayments.F(int64(0)),
			ExpiresAt:        dodopayments.F(time.Now()),
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

func TestEntitlementGrantRevoke(t *testing.T) {
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
	_, err := client.Entitlements.Grants.Revoke(
		context.TODO(),
		"ent_jt7jcvI79Xh8eehqgWdcm",
		"entg_w0ZCJZgNXuNDdMVzvja6p",
	)
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
