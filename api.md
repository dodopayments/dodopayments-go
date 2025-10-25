# CheckoutSessions

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CheckoutSessionRequestParam">CheckoutSessionRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CheckoutSessionResponse">CheckoutSessionResponse</a>

Methods:

- <code title="post /checkouts">client.CheckoutSessions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CheckoutSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CheckoutSessionNewParams">CheckoutSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CheckoutSessionResponse">CheckoutSessionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AttachExistingCustomerParam">AttachExistingCustomerParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BillingAddressParam">BillingAddressParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerRequestUnionParam">CustomerRequestUnionParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#NewCustomerParam">NewCustomerParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#OneTimeProductCartItemParam">OneTimeProductCartItemParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentMethodTypes">PaymentMethodTypes</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BillingAddress">BillingAddress</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerLimitedDetails">CustomerLimitedDetails</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#IntentStatus">IntentStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#OneTimeProductCartItem">OneTimeProductCartItem</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewResponse">PaymentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListResponse">PaymentListResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentGetLineItemsResponse">PaymentGetLineItemsResponse</a>

Methods:

- <code title="post /payments">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewParams">PaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewResponse">PaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{payment_id}">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListResponse">PaymentListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{payment_id}/line-items">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.GetLineItems">GetLineItems</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentGetLineItemsResponse">PaymentGetLineItemsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AttachAddonParam">AttachAddonParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#OnDemandSubscriptionParam">OnDemandSubscriptionParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionStatus">SubscriptionStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#TimeInterval">TimeInterval</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonCartResponseItem">AddonCartResponseItem</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionStatus">SubscriptionStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#TimeInterval">TimeInterval</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionListResponse">SubscriptionListResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionChargeResponse">SubscriptionChargeResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionGetUsageHistoryResponse">SubscriptionGetUsageHistoryResponse</a>

Methods:

- <code title="post /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionUpdateParams">SubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionListParams">SubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionListResponse">SubscriptionListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /subscriptions/{subscription_id}/change-plan">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.ChangePlan">ChangePlan</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionChangePlanParams">SubscriptionChangePlanParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /subscriptions/{subscription_id}/charge">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.Charge">Charge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionChargeParams">SubscriptionChargeParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionChargeResponse">SubscriptionChargeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions/{subscription_id}/usage-history">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.GetUsageHistory">GetUsageHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionGetUsageHistoryParams">SubscriptionGetUsageHistoryParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionGetUsageHistoryResponse">SubscriptionGetUsageHistoryResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoices

## Payments

Methods:

- <code title="get /invoices/payments/{payment_id}">client.Invoices.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#InvoicePaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /invoices/refunds/{refund_id}">client.Invoices.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#InvoicePaymentService.GetRefund">GetRefund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Licenses

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseActivateResponse">LicenseActivateResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateResponse">LicenseValidateResponse</a>

Methods:

- <code title="post /licenses/activate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseActivateParams">LicenseActivateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseActivateResponse">LicenseActivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /licenses/deactivate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Deactivate">Deactivate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseDeactivateParams">LicenseDeactivateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /licenses/validate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateParams">LicenseValidateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateResponse">LicenseValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# LicenseKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKey">LicenseKey</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyStatus">LicenseKeyStatus</a>

Methods:

- <code title="get /license_keys/{id}">client.LicenseKeys.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKey">LicenseKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /license_keys/{id}">client.LicenseKeys.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyUpdateParams">LicenseKeyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKey">LicenseKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /license_keys">client.LicenseKeys.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyListParams">LicenseKeyListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKey">LicenseKey</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# LicenseKeyInstances

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>

Methods:

- <code title="get /license_key_instances/{id}">client.LicenseKeyInstances.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /license_key_instances/{id}">client.LicenseKeyInstances.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceUpdateParams">LicenseKeyInstanceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /license_key_instances">client.LicenseKeyInstances.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceListParams">LicenseKeyInstanceListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerPortalSession">CustomerPortalSession</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{customer_id}">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /customers/{customer_id}">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CustomerPortal

Methods:

- <code title="post /customers/{customer_id}/customer-portal/session">client.Customers.CustomerPortal.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerCustomerPortalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerCustomerPortalNewParams">CustomerCustomerPortalNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerPortalSession">CustomerPortalSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Wallets

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWallet">CustomerWallet</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletListResponse">CustomerWalletListResponse</a>

Methods:

- <code title="get /customers/{customer_id}/wallets">client.Customers.Wallets.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletListResponse">CustomerWalletListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### LedgerEntries

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletTransaction">CustomerWalletTransaction</a>

Methods:

- <code title="post /customers/{customer_id}/wallets/ledger-entries">client.Customers.Wallets.LedgerEntries.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletLedgerEntryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletLedgerEntryNewParams">CustomerWalletLedgerEntryNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWallet">CustomerWallet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{customer_id}/wallets/ledger-entries">client.Customers.Wallets.LedgerEntries.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletLedgerEntryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletLedgerEntryListParams">CustomerWalletLedgerEntryListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerWalletTransaction">CustomerWalletTransaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundStatus">RefundStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundListResponse">RefundListResponse</a>

Methods:

- <code title="post /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundNewParams">RefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds/{refund_id}">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundListResponse">RefundListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Disputes

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Dispute">Dispute</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeStage">DisputeStage</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeStatus">DisputeStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#GetDispute">GetDispute</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeListResponse">DisputeListResponse</a>

Methods:

- <code title="get /disputes/{dispute_id}">client.Disputes.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, disputeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#GetDispute">GetDispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /disputes">client.Disputes.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeListParams">DisputeListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeListResponse">DisputeListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutListResponse">PayoutListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Products

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddMeterToPriceParam">AddMeterToPriceParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyDurationParam">LicenseKeyDurationParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PriceUnionParam">PriceUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddMeterToPrice">AddMeterToPrice</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyDuration">LicenseKeyDuration</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Price">Price</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListResponse">ProductListResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductUpdateFilesResponse">ProductUpdateFilesResponse</a>

Methods:

- <code title="post /products">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductNewParams">ProductNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductUpdateParams">ProductUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /products">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListParams">ProductListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListResponse">ProductListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /products/{id}/unarchive">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Unarchive">Unarchive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /products/{id}/files">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.UpdateFiles">UpdateFiles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductUpdateFilesParams">ProductUpdateFilesParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductUpdateFilesResponse">ProductUpdateFilesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>

Methods:

- <code title="put /products/{id}/images">client.Products.Images.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateParams">ProductImageUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Misc

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Currency">Currency</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#TaxCategory">TaxCategory</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Currency">Currency</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#TaxCategory">TaxCategory</a>

Methods:

- <code title="get /checkout/supported_countries">client.Misc.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MiscService.ListSupportedCountries">ListSupportedCountries</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Discounts

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountType">DiscountType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Discount">Discount</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountType">DiscountType</a>

Methods:

- <code title="post /discounts">client.Discounts.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountNewParams">DiscountNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Discount">Discount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /discounts/{discount_id}">client.Discounts.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, discountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Discount">Discount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /discounts/{discount_id}">client.Discounts.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, discountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountUpdateParams">DiscountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Discount">Discount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /discounts">client.Discounts.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountListParams">DiscountListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Discount">Discount</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /discounts/{discount_id}">client.Discounts.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DiscountService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, discountID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Addons

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonResponse">AddonResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonUpdateImagesResponse">AddonUpdateImagesResponse</a>

Methods:

- <code title="post /addons">client.Addons.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonNewParams">AddonNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonResponse">AddonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /addons/{id}">client.Addons.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonResponse">AddonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /addons/{id}">client.Addons.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonUpdateParams">AddonUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonResponse">AddonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /addons">client.Addons.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonListParams">AddonListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonResponse">AddonResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /addons/{id}/images">client.Addons.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonService.UpdateImages">UpdateImages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AddonUpdateImagesResponse">AddonUpdateImagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Brands

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Brand">Brand</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandListResponse">BrandListResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandUpdateImagesResponse">BrandUpdateImagesResponse</a>

Methods:

- <code title="post /brands">client.Brands.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandNewParams">BrandNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /brands/{id}">client.Brands.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /brands/{id}">client.Brands.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandUpdateParams">BrandUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /brands">client.Brands.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandListResponse">BrandListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /brands/{id}/images">client.Brands.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandService.UpdateImages">UpdateImages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BrandUpdateImagesResponse">BrandUpdateImagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookDetails">WebhookDetails</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookGetSecretResponse">WebhookGetSecretResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeAcceptedWebhookEvent">DisputeAcceptedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeCancelledWebhookEvent">DisputeCancelledWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeChallengedWebhookEvent">DisputeChallengedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeExpiredWebhookEvent">DisputeExpiredWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeLostWebhookEvent">DisputeLostWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeOpenedWebhookEvent">DisputeOpenedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeWonWebhookEvent">DisputeWonWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyCreatedWebhookEvent">LicenseKeyCreatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentCancelledWebhookEvent">PaymentCancelledWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentFailedWebhookEvent">PaymentFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentProcessingWebhookEvent">PaymentProcessingWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentSucceededWebhookEvent">PaymentSucceededWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundFailedWebhookEvent">RefundFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundSucceededWebhookEvent">RefundSucceededWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionActiveWebhookEvent">SubscriptionActiveWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionCancelledWebhookEvent">SubscriptionCancelledWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionExpiredWebhookEvent">SubscriptionExpiredWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionFailedWebhookEvent">SubscriptionFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionOnHoldWebhookEvent">SubscriptionOnHoldWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionPlanChangedWebhookEvent">SubscriptionPlanChangedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionRenewedWebhookEvent">SubscriptionRenewedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeAcceptedWebhookEvent">DisputeAcceptedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeCancelledWebhookEvent">DisputeCancelledWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeChallengedWebhookEvent">DisputeChallengedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeExpiredWebhookEvent">DisputeExpiredWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeLostWebhookEvent">DisputeLostWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeOpenedWebhookEvent">DisputeOpenedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeWonWebhookEvent">DisputeWonWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyCreatedWebhookEvent">LicenseKeyCreatedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentCancelledWebhookEvent">PaymentCancelledWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentFailedWebhookEvent">PaymentFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentProcessingWebhookEvent">PaymentProcessingWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentSucceededWebhookEvent">PaymentSucceededWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundFailedWebhookEvent">RefundFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundSucceededWebhookEvent">RefundSucceededWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionActiveWebhookEvent">SubscriptionActiveWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionCancelledWebhookEvent">SubscriptionCancelledWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionExpiredWebhookEvent">SubscriptionExpiredWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionFailedWebhookEvent">SubscriptionFailedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionOnHoldWebhookEvent">SubscriptionOnHoldWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionPlanChangedWebhookEvent">SubscriptionPlanChangedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionRenewedWebhookEvent">SubscriptionRenewedWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UnsafeUnwrapWebhookEvent">UnsafeUnwrapWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UnwrapWebhookEvent">UnwrapWebhookEvent</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookNewParams">WebhookNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookDetails">WebhookDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/{webhook_id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookDetails">WebhookDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /webhooks/{webhook_id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookUpdateParams">WebhookUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookDetails">WebhookDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookListParams">WebhookListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#CursorPagePagination">CursorPagePagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookDetails">WebhookDetails</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhook_id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /webhooks/{webhook_id}/secret">client.Webhooks.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookService.GetSecret">GetSecret</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookGetSecretResponse">WebhookGetSecretResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Headers

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookHeaderGetResponse">WebhookHeaderGetResponse</a>

Methods:

- <code title="get /webhooks/{webhook_id}/headers">client.Webhooks.Headers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookHeaderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookHeaderGetResponse">WebhookHeaderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /webhooks/{webhook_id}/headers">client.Webhooks.Headers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookHeaderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookHeaderUpdateParams">WebhookHeaderUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# WebhookEvents

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookEventType">WebhookEventType</a>

# UsageEvents

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#EventInputParam">EventInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Event">Event</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UsageEventIngestResponse">UsageEventIngestResponse</a>

Methods:

- <code title="get /events/{event_id}">client.UsageEvents.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UsageEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Event">Event</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /events">client.UsageEvents.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UsageEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UsageEventListParams">UsageEventListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Event">Event</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /events/ingest">client.UsageEvents.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UsageEventService.Ingest">Ingest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UsageEventIngestParams">UsageEventIngestParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#UsageEventIngestResponse">UsageEventIngestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Meters

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterAggregationParam">MeterAggregationParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterFilterParam">MeterFilterParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Meter">Meter</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterAggregation">MeterAggregation</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterFilter">MeterFilter</a>

Methods:

- <code title="post /meters">client.Meters.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterNewParams">MeterNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Meter">Meter</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /meters/{id}">client.Meters.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Meter">Meter</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /meters">client.Meters.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterListParams">MeterListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Meter">Meter</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /meters/{id}">client.Meters.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /meters/{id}/unarchive">client.Meters.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MeterService.Unarchive">Unarchive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
