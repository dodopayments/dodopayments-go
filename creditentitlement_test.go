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

func TestCreditEntitlementNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditEntitlements.New(context.TODO(), dodopayments.CreditEntitlementNewParams{
		Name:                      dodopayments.F("name"),
		OverageEnabled:            dodopayments.F(true),
		Precision:                 dodopayments.F(int64(0)),
		RolloverEnabled:           dodopayments.F(true),
		Unit:                      dodopayments.F("unit"),
		Currency:                  dodopayments.F(dodopayments.CurrencyAed),
		Description:               dodopayments.F("description"),
		ExpiresAfterDays:          dodopayments.F(int64(0)),
		MaxRolloverCount:          dodopayments.F(int64(0)),
		OverageBehavior:           dodopayments.F(dodopayments.CbbOverageBehaviorForgiveAtReset),
		OverageLimit:              dodopayments.F(int64(0)),
		PricePerUnit:              dodopayments.F("price_per_unit"),
		RolloverPercentage:        dodopayments.F(int64(0)),
		RolloverTimeframeCount:    dodopayments.F(int64(0)),
		RolloverTimeframeInterval: dodopayments.F(dodopayments.TimeIntervalDay),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditEntitlementGet(t *testing.T) {
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
	_, err := client.CreditEntitlements.Get(context.TODO(), "id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditEntitlementUpdateWithOptionalParams(t *testing.T) {
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
	err := client.CreditEntitlements.Update(
		context.TODO(),
		"id",
		dodopayments.CreditEntitlementUpdateParams{
			Currency:                  dodopayments.F(dodopayments.CurrencyAed),
			Description:               dodopayments.F("description"),
			ExpiresAfterDays:          dodopayments.F(int64(0)),
			MaxRolloverCount:          dodopayments.F(int64(0)),
			Name:                      dodopayments.F("name"),
			OverageBehavior:           dodopayments.F(dodopayments.CbbOverageBehaviorForgiveAtReset),
			OverageEnabled:            dodopayments.F(true),
			OverageLimit:              dodopayments.F(int64(0)),
			PricePerUnit:              dodopayments.F("price_per_unit"),
			RolloverEnabled:           dodopayments.F(true),
			RolloverPercentage:        dodopayments.F(int64(0)),
			RolloverTimeframeCount:    dodopayments.F(int64(0)),
			RolloverTimeframeInterval: dodopayments.F(dodopayments.TimeIntervalDay),
			Unit:                      dodopayments.F("unit"),
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

func TestCreditEntitlementListWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditEntitlements.List(context.TODO(), dodopayments.CreditEntitlementListParams{
		Deleted:    dodopayments.F(true),
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

func TestCreditEntitlementDelete(t *testing.T) {
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
	err := client.CreditEntitlements.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditEntitlementUndelete(t *testing.T) {
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
	err := client.CreditEntitlements.Undelete(context.TODO(), "id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
