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

func TestCustomerWalletLedgerEntryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Wallets.LedgerEntries.New(
		context.TODO(),
		"customer_id",
		dodopayments.CustomerWalletLedgerEntryNewParams{
			Amount:         dodopayments.F(int64(0)),
			Currency:       dodopayments.F(dodopayments.CurrencyAed),
			EntryType:      dodopayments.F(dodopayments.CustomerWalletLedgerEntryNewParamsEntryTypeCredit),
			IdempotencyKey: dodopayments.F("idempotency_key"),
			Reason:         dodopayments.F("reason"),
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

func TestCustomerWalletLedgerEntryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Customers.Wallets.LedgerEntries.List(
		context.TODO(),
		"customer_id",
		dodopayments.CustomerWalletLedgerEntryListParams{
			Currency:   dodopayments.F(dodopayments.CurrencyAed),
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
