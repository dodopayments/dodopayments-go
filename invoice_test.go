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
	"github.com/dodopayments/dodopayments-go/shared"
)

func TestInvoiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.New(context.TODO(), dodopayments.InvoiceNewParams{
		CollectionMethod: dodopayments.F(dodopayments.CollectionMethodAutoCharge),
		Currency:         dodopayments.F(dodopayments.CurrencyAed),
		CustomerID:       dodopayments.F("customer_id"),
		Items: dodopayments.F([]dodopayments.InvoiceNewParamsItemUnion{dodopayments.InvoiceNewParamsItemsCustomItem{
			Description: dodopayments.F("description"),
			Kind:        dodopayments.F(dodopayments.InvoiceNewParamsItemsCustomItemKindCustom),
			TaxCategory: dodopayments.F(dodopayments.TaxCategoryDigitalProducts),
			UnitAmount:  dodopayments.F(int64(0)),
			Quantity:    dodopayments.F(int64(0)),
		}}),
		BillingAddress: dodopayments.F(dodopayments.InvoiceNewParamsBillingAddress{
			Country: dodopayments.F(dodopayments.CountryCodeAf),
			City:    dodopayments.F("city"),
			State:   dodopayments.F("state"),
			Street:  dodopayments.F("street"),
			Zipcode: dodopayments.F("zipcode"),
		}),
		CcEmails:      dodopayments.F([]string{"string"}),
		CustomerEmail: dodopayments.F("customer_email"),
		DiscountIDs:   dodopayments.F([]string{"string"}),
		DueAt:         dodopayments.F(time.Now()),
		Dunning: dodopayments.F(dodopayments.InvoiceNewParamsDunning{
			RemindersEnabled: dodopayments.F(true),
			OffsetsDays:      dodopayments.F([]int64{int64(0)}),
		}),
		Memo: dodopayments.F("memo"),
		Metadata: dodopayments.F(dodopayments.MetadataParam{
			"foo": shared.UnionString("string"),
		}),
		PaymentTerms: dodopayments.F[dodopayments.InvoiceNewParamsPaymentTermsUnion](dodopayments.InvoiceNewParamsPaymentTermsDueOnReceipt{
			Type: dodopayments.F(dodopayments.InvoiceNewParamsPaymentTermsDueOnReceiptTypeDueOnReceipt),
		}),
		ReviewRequired: dodopayments.F(true),
		TaxID:          dodopayments.F("tax_id"),
		IdempotencyKey: dodopayments.F("Idempotency-Key"),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceGet(t *testing.T) {
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
	_, err := client.Invoices.Get(context.TODO(), "invoice_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.Update(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceUpdateParams{
			IfMatch: dodopayments.F("If-Match"),
			BillingAddress: dodopayments.F(dodopayments.InvoiceUpdateParamsBillingAddress{
				Country: dodopayments.F(dodopayments.CountryCodeAf),
				City:    dodopayments.F("city"),
				State:   dodopayments.F("state"),
				Street:  dodopayments.F("street"),
				Zipcode: dodopayments.F("zipcode"),
			}),
			CcEmails:         dodopayments.F([]string{"string"}),
			CollectionMethod: dodopayments.F(dodopayments.CollectionMethodAutoCharge),
			Currency:         dodopayments.F(dodopayments.CurrencyAed),
			CustomerEmail:    dodopayments.F("customer_email"),
			CustomerID:       dodopayments.F("customer_id"),
			DiscountIDs:      dodopayments.F([]string{"string"}),
			DueAt:            dodopayments.F(time.Now()),
			Dunning: dodopayments.F(dodopayments.InvoiceUpdateParamsDunning{
				RemindersEnabled: dodopayments.F(true),
				OffsetsDays:      dodopayments.F([]int64{int64(0)}),
			}),
			Items: dodopayments.F([]dodopayments.InvoiceUpdateParamsItemUnion{dodopayments.InvoiceUpdateParamsItemsCustomItem{
				Description: dodopayments.F("description"),
				Kind:        dodopayments.F(dodopayments.InvoiceUpdateParamsItemsCustomItemKindCustom),
				TaxCategory: dodopayments.F(dodopayments.TaxCategoryDigitalProducts),
				UnitAmount:  dodopayments.F(int64(0)),
				Quantity:    dodopayments.F(int64(0)),
			}}),
			Memo: dodopayments.F("memo"),
			Metadata: dodopayments.F(dodopayments.MetadataParam{
				"foo": shared.UnionString("string"),
			}),
			PaymentTerms: dodopayments.F[dodopayments.InvoiceUpdateParamsPaymentTermsUnion](dodopayments.InvoiceUpdateParamsPaymentTermsDueOnReceipt{
				Type: dodopayments.F(dodopayments.InvoiceUpdateParamsPaymentTermsDueOnReceiptTypeDueOnReceipt),
			}),
			ReviewRequired: dodopayments.F(true),
			TaxID:          dodopayments.F("tax_id"),
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

func TestInvoiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.List(context.TODO(), dodopayments.InvoiceListParams{
		CollectionMethod: dodopayments.F(dodopayments.InvoiceListParamsCollectionMethodAutoCharge),
		CreatedAtGte:     dodopayments.F(time.Now()),
		CreatedAtLte:     dodopayments.F(time.Now()),
		CustomerID:       dodopayments.F("customer_id"),
		DueAtGte:         dodopayments.F(time.Now()),
		DueAtLte:         dodopayments.F(time.Now()),
		PageNumber:       dodopayments.F(int64(0)),
		PageSize:         dodopayments.F(int64(0)),
		Status:           dodopayments.F([]dodopayments.InvoiceListParamsStatus{dodopayments.InvoiceListParamsStatusOpen}),
		SubscriptionID:   dodopayments.F("subscription_id"),
		Tags:             dodopayments.F([]dodopayments.InvoiceListParamsTag{dodopayments.InvoiceListParamsTagRefunded}),
		TotalAmountGte:   dodopayments.F(int64(0)),
		TotalAmountLte:   dodopayments.F(int64(0)),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceDelete(t *testing.T) {
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
	err := client.Invoices.Delete(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceDeleteParams{
			IfMatch: dodopayments.F("If-Match"),
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

func TestInvoiceAddNote(t *testing.T) {
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
	_, err := client.Invoices.AddNote(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceAddNoteParams{
			NoteRequest: dodopayments.NoteRequestParam{
				Text: dodopayments.F("text"),
			},
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

func TestInvoiceNewHostedLink(t *testing.T) {
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
	_, err := client.Invoices.NewHostedLink(context.TODO(), "invoice_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceDuplicate(t *testing.T) {
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
	_, err := client.Invoices.Duplicate(context.TODO(), "invoice_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceFinalizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.Finalize(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceFinalizeParams{
			IfMatch:       dodopayments.F("If-Match"),
			CcEmails:      dodopayments.F([]string{"string"}),
			CustomerEmail: dodopayments.F("customer_email"),
			Send:          dodopayments.F(true),
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

func TestInvoiceListEvents(t *testing.T) {
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
	_, err := client.Invoices.ListEvents(context.TODO(), "invoice_id")
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoiceMarkUncollectable(t *testing.T) {
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
	_, err := client.Invoices.MarkUncollectable(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceMarkUncollectableParams{
			IfMatch: dodopayments.F("If-Match"),
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

func TestInvoicePreviewWithOptionalParams(t *testing.T) {
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
	err := client.Invoices.Preview(context.TODO(), dodopayments.InvoicePreviewParams{
		CollectionMethod: dodopayments.F(dodopayments.CollectionMethodAutoCharge),
		Currency:         dodopayments.F(dodopayments.CurrencyAed),
		CustomerID:       dodopayments.F("customer_id"),
		Items: dodopayments.F([]dodopayments.InvoicePreviewParamsItemUnion{dodopayments.InvoicePreviewParamsItemsCustomItem{
			Description: dodopayments.F("description"),
			Kind:        dodopayments.F(dodopayments.InvoicePreviewParamsItemsCustomItemKindCustom),
			TaxCategory: dodopayments.F(dodopayments.TaxCategoryDigitalProducts),
			UnitAmount:  dodopayments.F(int64(0)),
			Quantity:    dodopayments.F(int64(0)),
		}}),
		Format: dodopayments.F(dodopayments.InvoicePreviewParamsFormatJson),
		BillingAddress: dodopayments.F(dodopayments.InvoicePreviewParamsBillingAddress{
			Country: dodopayments.F(dodopayments.CountryCodeAf),
			City:    dodopayments.F("city"),
			State:   dodopayments.F("state"),
			Street:  dodopayments.F("street"),
			Zipcode: dodopayments.F("zipcode"),
		}),
		CcEmails:      dodopayments.F([]string{"string"}),
		CustomerEmail: dodopayments.F("customer_email"),
		DiscountIDs:   dodopayments.F([]string{"string"}),
		DueAt:         dodopayments.F(time.Now()),
		Dunning: dodopayments.F(dodopayments.InvoicePreviewParamsDunning{
			RemindersEnabled: dodopayments.F(true),
			OffsetsDays:      dodopayments.F([]int64{int64(0)}),
		}),
		Memo: dodopayments.F("memo"),
		Metadata: dodopayments.F(dodopayments.MetadataParam{
			"foo": shared.UnionString("string"),
		}),
		PaymentTerms: dodopayments.F[dodopayments.InvoicePreviewParamsPaymentTermsUnion](dodopayments.InvoicePreviewParamsPaymentTermsDueOnReceipt{
			Type: dodopayments.F(dodopayments.InvoicePreviewParamsPaymentTermsDueOnReceiptTypeDueOnReceipt),
		}),
		ReviewRequired: dodopayments.F(true),
		TaxID:          dodopayments.F("tax_id"),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvoicePreviewExistingWithOptionalParams(t *testing.T) {
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
	err := client.Invoices.PreviewExisting(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoicePreviewExistingParams{
			Format: dodopayments.F(dodopayments.InvoicePreviewExistingParamsFormatJson),
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

func TestInvoiceRetry(t *testing.T) {
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
	_, err := client.Invoices.Retry(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceRetryParams{
			IfMatch: dodopayments.F("If-Match"),
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

func TestInvoiceSend(t *testing.T) {
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
	_, err := client.Invoices.Send(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceSendParams{
			IfMatch: dodopayments.F("If-Match"),
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

func TestInvoiceUpdateDunningWithOptionalParams(t *testing.T) {
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
	_, err := client.Invoices.UpdateDunning(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceUpdateDunningParams{
			Action:      dodopayments.F(dodopayments.InvoiceUpdateDunningParamsActionPause),
			IfMatch:     dodopayments.F("If-Match"),
			OffsetsDays: dodopayments.F([]int64{int64(0)}),
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

func TestInvoiceVoid(t *testing.T) {
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
	_, err := client.Invoices.Void(
		context.TODO(),
		"invoice_id",
		dodopayments.InvoiceVoidParams{
			IfMatch: dodopayments.F("If-Match"),
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
