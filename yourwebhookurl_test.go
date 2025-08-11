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

func TestYourWebhookURLNew(t *testing.T) {
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
	err := client.YourWebhookURL.New(context.TODO(), dodopayments.YourWebhookURLNewParams{
		WebhookPayload: dodopayments.WebhookPayloadParam{
			BusinessID: dodopayments.F("business_id"),
			Data: dodopayments.F[dodopayments.WebhookPayloadDataUnionParam](dodopayments.WebhookPayloadDataPaymentParam(dodopayments.WebhookPayloadDataPaymentParam{
				PaymentParam: dodopayments.PaymentParam{
					Billing: dodopayments.F(dodopayments.BillingAddressParam{
						City:    dodopayments.F("city"),
						Country: dodopayments.F(dodopayments.CountryCodeAf),
						State:   dodopayments.F("state"),
						Street:  dodopayments.F("street"),
						Zipcode: dodopayments.F("zipcode"),
					}),
					BrandID:    dodopayments.F("brand_id"),
					BusinessID: dodopayments.F("business_id"),
					CreatedAt:  dodopayments.F(time.Now()),
					Currency:   dodopayments.F(dodopayments.CurrencyAed),
					Customer: dodopayments.F(dodopayments.CustomerLimitedDetailsParam{
						CustomerID: dodopayments.F("customer_id"),
						Email:      dodopayments.F("email"),
						Name:       dodopayments.F("name"),
					}),
					DigitalProductsDelivered: dodopayments.F(true),
					Disputes: dodopayments.F([]dodopayments.DisputeParam{{
						Amount:        dodopayments.F("amount"),
						BusinessID:    dodopayments.F("business_id"),
						CreatedAt:     dodopayments.F(time.Now()),
						Currency:      dodopayments.F("currency"),
						DisputeID:     dodopayments.F("dispute_id"),
						DisputeStage:  dodopayments.F(dodopayments.DisputeStagePreDispute),
						DisputeStatus: dodopayments.F(dodopayments.DisputeStatusDisputeOpened),
						PaymentID:     dodopayments.F("payment_id"),
						Remarks:       dodopayments.F("remarks"),
					}}),
					Metadata: dodopayments.F(map[string]string{
						"foo": "string",
					}),
					PaymentID: dodopayments.F("payment_id"),
					Refunds: dodopayments.F([]dodopayments.RefundParam{{
						BusinessID: dodopayments.F("business_id"),
						CreatedAt:  dodopayments.F(time.Now()),
						IsPartial:  dodopayments.F(true),
						PaymentID:  dodopayments.F("payment_id"),
						RefundID:   dodopayments.F("refund_id"),
						Status:     dodopayments.F(dodopayments.RefundStatusSucceeded),
						Amount:     dodopayments.F(int64(0)),
						Currency:   dodopayments.F(dodopayments.CurrencyAed),
						Reason:     dodopayments.F("reason"),
					}}),
					SettlementAmount:   dodopayments.F(int64(0)),
					SettlementCurrency: dodopayments.F(dodopayments.CurrencyAed),
					TotalAmount:        dodopayments.F(int64(0)),
					CardIssuingCountry: dodopayments.F(dodopayments.CountryCodeAf),
					CardLastFour:       dodopayments.F("card_last_four"),
					CardNetwork:        dodopayments.F("card_network"),
					CardType:           dodopayments.F("card_type"),
					DiscountID:         dodopayments.F("discount_id"),
					ErrorCode:          dodopayments.F("error_code"),
					ErrorMessage:       dodopayments.F("error_message"),
					PaymentLink:        dodopayments.F("payment_link"),
					PaymentMethod:      dodopayments.F("payment_method"),
					PaymentMethodType:  dodopayments.F("payment_method_type"),
					ProductCart: dodopayments.F([]dodopayments.PaymentProductCartParam{{
						ProductID: dodopayments.F("product_id"),
						Quantity:  dodopayments.F(int64(0)),
					}}),
					SettlementTax:  dodopayments.F(int64(0)),
					Status:         dodopayments.F(dodopayments.IntentStatusSucceeded),
					SubscriptionID: dodopayments.F("subscription_id"),
					Tax:            dodopayments.F(int64(0)),
					UpdatedAt:      dodopayments.F(time.Now()),
				},
				PayloadType: dodopayments.F(dodopayments.WebhookPayloadDataPaymentPayloadTypePayment),
			})),
			Timestamp: dodopayments.F(time.Now()),
			Type:      dodopayments.F(dodopayments.WebhookEventTypePaymentSucceeded),
		},
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
