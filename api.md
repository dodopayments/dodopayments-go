# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewResponse">PaymentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="post /payments">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewParams">PaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewResponse">PaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{payment_id}">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListResponse">PaymentListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>

Methods:

- <code title="post /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionUpdateParams">SubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionListParams">SubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Licenses

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateResponse">LicenseValidateResponse</a>

Methods:

- <code title="post /licenses/activate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseActivateParams">LicenseActivateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /licenses/deactivate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Deactivate">Deactivate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseDeactivateParams">LicenseDeactivateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /licenses/validate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateParams">LicenseValidateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateResponse">LicenseValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# LicenseKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKey">LicenseKey</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyListResponse">LicenseKeyListResponse</a>

Methods:

- <code title="get /license_keys/{id}">client.LicenseKeys.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKey">LicenseKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /license_keys/{id}">client.LicenseKeys.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyUpdateParams">LicenseKeyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKey">LicenseKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /license_keys">client.LicenseKeys.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyListParams">LicenseKeyListParams</a>) ([]<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyListResponse">LicenseKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# LicenseKeyInstances

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceListResponse">LicenseKeyInstanceListResponse</a>

Methods:

- <code title="get /license_key_instances/{id}">client.LicenseKeyInstances.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /license_key_instances/{id}">client.LicenseKeyInstances.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceUpdateParams">LicenseKeyInstanceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /license_key_instances">client.LicenseKeyInstances.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceListParams">LicenseKeyInstanceListParams</a>) ([]<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstanceListResponse">LicenseKeyInstanceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{customer_id}">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /customers/{customer_id}">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>

Methods:

- <code title="post /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundNewParams">RefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds/{refund_id}">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Disputes

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Dispute">Dispute</a>

Methods:

- <code title="get /disputes/{dispute_id}">client.Disputes.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, disputeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Dispute">Dispute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /disputes">client.Disputes.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeListParams">DisputeListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Dispute">Dispute</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PayoutListResponse">PayoutListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# WebhookEvents

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookEvent">WebhookEvent</a>

Methods:

- <code title="get /webhook_events/{webhook_event_id}">client.WebhookEvents.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEventID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookEvent">WebhookEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhook_events">client.WebhookEvents.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookEventListParams">WebhookEventListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#WebhookEvent">WebhookEvent</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Products

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="post /products">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductNewParams">ProductNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductUpdateParams">ProductUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /products">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListParams">ProductListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListResponse">ProductListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>

Methods:

- <code title="put /products/{id}/images">client.Products.Images.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Misc

## SupportedCountries

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>

Methods:

- <code title="get /checkout/supported_countries">client.Misc.SupportedCountries.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#MiscSupportedCountryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
