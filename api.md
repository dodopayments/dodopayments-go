# Checkout

## SupportedCountries

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CountryCodeAlpha2">CountryCodeAlpha2</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CountryCodeAlpha2">CountryCodeAlpha2</a>

Methods:

- <code title="get /checkout/supported_countries">client.Checkout.SupportedCountries.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CheckoutSupportedCountryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CountryCodeAlpha2">CountryCodeAlpha2</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="get /customers/{customer_id}">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Disputes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Dispute">Dispute</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#DisputeListResponse">DisputeListResponse</a>

Methods:

- <code title="get /disputes/{dispute_id}">client.Disputes.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#DisputeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, disputeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Dispute">Dispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /disputes">client.Disputes.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#DisputeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#DisputeListParams">DisputeListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#DisputeListResponse">DisputeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentNewResponse">PaymentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="post /payments">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentNewParams">PaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentNewResponse">PaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{payment_id}">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#PayoutListResponse">PayoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Products

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductNewResponse">ProductNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="post /products">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductNewParams">ProductNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductNewResponse">ProductNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductUpdateParams">ProductUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /products">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductListParams">ProductListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductListResponse">ProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>

Methods:

- <code title="put /products/{id}/images">client.Products.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductImageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#RefundListResponse">RefundListResponse</a>

Methods:

- <code title="post /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#RefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#RefundNewParams">RefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds/{refund_id}">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#RefundListResponse">RefundListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Subscription">Subscription</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionListResponse">SubscriptionListResponse</a>

Methods:

- <code title="post /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionUpdateParams">SubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionListParams">SubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#SubscriptionListResponse">SubscriptionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookEvents

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#WebhookEventLog">WebhookEventLog</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#WebhookEventListResponse">WebhookEventListResponse</a>

Methods:

- <code title="get /webhook_events/{webhook_event_id}">client.WebhookEvents.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#WebhookEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEventID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#WebhookEventLog">WebhookEventLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhook_events">client.WebhookEvents.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#WebhookEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#WebhookEventListParams">WebhookEventListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/dodo-payments-go#WebhookEventListResponse">WebhookEventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
