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

func TestCreditEntitlementBalanceGet(t *testing.T) {
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
	_, err := client.CreditEntitlements.Balances.Get(
		context.TODO(),
		"credit_entitlement_id",
		"customer_id",
	)
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCreditEntitlementBalanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditEntitlements.Balances.List(
		context.TODO(),
		"credit_entitlement_id",
		dodopayments.CreditEntitlementBalanceListParams{
			CustomerID: dodopayments.F("customer_id"),
			PageNumber: dodopayments.F(int64(0)),
			PageSize:   dodopayments.F(int64(0)),
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

func TestCreditEntitlementBalanceNewLedgerEntryWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditEntitlements.Balances.NewLedgerEntry(
		context.TODO(),
		"credit_entitlement_id",
		"customer_id",
		dodopayments.CreditEntitlementBalanceNewLedgerEntryParams{
			Amount:         dodopayments.F("amount"),
			EntryType:      dodopayments.F(dodopayments.LedgerEntryTypeCredit),
			ExpiresAt:      dodopayments.F(time.Now()),
			IdempotencyKey: dodopayments.F("idempotency_key"),
			Metadata: dodopayments.F(map[string]string{
				"foo": "string",
			}),
			Reason: dodopayments.F("reason"),
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

func TestCreditEntitlementBalanceListGrantsWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditEntitlements.Balances.ListGrants(
		context.TODO(),
		"credit_entitlement_id",
		"customer_id",
		dodopayments.CreditEntitlementBalanceListGrantsParams{
			PageNumber: dodopayments.F(int64(0)),
			PageSize:   dodopayments.F(int64(0)),
			Status:     dodopayments.F(dodopayments.CreditEntitlementBalanceListGrantsParamsStatusActive),
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

func TestCreditEntitlementBalanceListLedgerWithOptionalParams(t *testing.T) {
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
	_, err := client.CreditEntitlements.Balances.ListLedger(
		context.TODO(),
		"credit_entitlement_id",
		"customer_id",
		dodopayments.CreditEntitlementBalanceListLedgerParams{
			EndDate:         dodopayments.F(time.Now()),
			PageNumber:      dodopayments.F(int64(0)),
			PageSize:        dodopayments.F(int64(0)),
			StartDate:       dodopayments.F(time.Now()),
			TransactionType: dodopayments.F("transaction_type"),
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
