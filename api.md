# Payments

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#AttachExistingCustomerParam">AttachExistingCustomerParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#BillingAddressParam">BillingAddressParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CreateNewCustomerParam">CreateNewCustomerParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerRequestUnionParam">CustomerRequestUnionParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#IntentStatus">IntentStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#OneTimeProductCartItemParam">OneTimeProductCartItemParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerLimitedDetails">CustomerLimitedDetails</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#IntentStatus">IntentStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#OneTimeProductCartItem">OneTimeProductCartItem</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewResponse">PaymentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="post /payments">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewParams">PaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentNewResponse">PaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{payment_id}">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PaymentListResponse">PaymentListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Subscriptions

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionStatus">SubscriptionStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#TimeInterval">TimeInterval</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionStatus">SubscriptionStatus</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#TimeInterval">TimeInterval</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>

Methods:

- <code title="post /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewParams">SubscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /subscriptions/{subscription_id}">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionUpdateParams">SubscriptionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /subscriptions">client.Subscriptions.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#SubscriptionListParams">SubscriptionListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Subscription">Subscription</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoices

## Payments

Methods:

- <code title="get /invoices/payments/{payment_id}">client.Invoices.Payments.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#InvoicePaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Licenses

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateResponse">LicenseValidateResponse</a>

Methods:

- <code title="post /licenses/activate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseActivateParams">LicenseActivateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyInstance">LicenseKeyInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /licenses/deactivate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Deactivate">Deactivate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseDeactivateParams">LicenseDeactivateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /licenses/validate">client.Licenses.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateParams">LicenseValidateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseValidateResponse">LicenseValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# LicenseKeys

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyStatus">LicenseKeyStatus</a>

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

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{customer_id}">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /customers/{customer_id}">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Customer">Customer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CustomerPortal

Methods:

- <code title="post /customers/{customer_id}/customer-portal/session">client.Customers.CustomerPortal.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerCustomerPortalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CustomerCustomerPortalNewParams">CustomerCustomerPortalNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Refunds

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundStatus">RefundStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundStatus">RefundStatus</a>

Methods:

- <code title="post /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundNewParams">RefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds/{refund_id}">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Refund">Refund</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Disputes

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeStage">DisputeStage</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeStatus">DisputeStatus</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Dispute">Dispute</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeStage">DisputeStage</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#DisputeStatus">DisputeStatus</a>

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

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyDurationParam">LicenseKeyDurationParam</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#PriceUnionParam">PriceUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#LicenseKeyDuration">LicenseKeyDuration</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Price">Price</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>
- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListResponse">ProductListResponse</a>

Methods:

- <code title="post /products">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductNewParams">ProductNewParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductUpdateParams">ProductUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /products">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListParams">ProductListParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go/packages/pagination#DefaultPageNumberPagination">DefaultPageNumberPagination</a>[<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductListResponse">ProductListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /products/{id}">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /products/{id}/unarchive">client.Products.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductService.Unarchive">Unarchive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>

Methods:

- <code title="put /products/{id}/images">client.Products.Images.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateParams">ProductImageUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#ProductImageUpdateResponse">ProductImageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Misc

Params Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go">dodopayments</a>.<a href="https://pkg.go.dev/github.com/dodopayments/dodopayments-go#CountryCode">CountryCode</a>

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
