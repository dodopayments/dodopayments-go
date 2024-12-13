// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dodopayments_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/dodo-payments-go"
	"github.com/stainless-sdks/dodo-payments-go/internal/testutil"
	"github.com/stainless-sdks/dodo-payments-go/option"
)

func TestOutgoingWebhookNew(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dodopayments.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.OutgoingWebhooks.New(context.TODO(), dodopayments.OutgoingWebhookNewParams{
		BusinessID: dodopayments.F("business_id"),
		Data: dodopayments.F[dodopayments.OutgoingWebhookNewParamsDataUnion](dodopayments.OutgoingWebhookNewParamsDataPayment(dodopayments.OutgoingWebhookNewParamsDataPayment{
			BusinessID: "business_id",
			CreatedAt:  time.Now(),
			Currency:   dodopayments.PaymentCurrencyAed,
			Customer: dodopayments.PaymentCustomerParam{
				CustomerID: dodopayments.F("customer_id"),
				Email:      dodopayments.F("email"),
				Name:       dodopayments.F("name"),
			},
			Disputes: []dodopayments.DisputeParam{{
				Amount:        dodopayments.F("amount"),
				BusinessID:    dodopayments.F("business_id"),
				CreatedAt:     dodopayments.F(time.Now()),
				Currency:      dodopayments.F("currency"),
				DisputeID:     dodopayments.F("dispute_id"),
				DisputeStage:  dodopayments.F(dodopayments.DisputeDisputeStagePreDispute),
				DisputeStatus: dodopayments.F(dodopayments.DisputeDisputeStatusDisputeOpened),
				PaymentID:     dodopayments.F("payment_id"),
			}},
			PaymentID: "payment_id",
			Refunds: []dodopayments.RefundParam{{
				BusinessID: dodopayments.F("business_id"),
				CreatedAt:  dodopayments.F(time.Now()),
				PaymentID:  dodopayments.F("payment_id"),
				RefundID:   dodopayments.F("refund_id"),
				Status:     dodopayments.F(dodopayments.RefundStatusSucceeded),
				Amount:     dodopayments.F(int64(0)),
				Currency:   dodopayments.F(dodopayments.RefundCurrencyAed),
				Reason:     dodopayments.F("reason"),
			}},
			TotalAmount:       int64(0),
			PaymentLink:       "payment_link",
			PaymentMethod:     "payment_method",
			PaymentMethodType: "payment_method_type",
			ProductCart: []dodopayments.PaymentProductCartParam{{
				ProductID: dodopayments.F("product_id"),
				Quantity:  dodopayments.F(int64(0)),
			}},
			Status:         dodopayments.PaymentStatusSucceeded,
			SubscriptionID: "subscription_id",
			Tax:            int64(0),
			UpdatedAt:      time.Now(),
			PayloadType:    dodopayments.OutgoingWebhookNewParamsDataPaymentPayloadTypePayment,
		})),
		Timestamp:        dodopayments.F(time.Now()),
		Type:             dodopayments.F(dodopayments.OutgoingWebhookNewParamsTypePaymentSucceeded),
		WebhookID:        dodopayments.F("webhook-id"),
		WebhookSignature: dodopayments.F("webhook-signature"),
		WebhookTimestamp: dodopayments.F("webhook-timestamp"),
	})
	if err != nil {
		var apierr *dodopayments.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
