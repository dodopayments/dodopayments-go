// Package main demonstrates basic caching patterns with Dodo Payments SDK.
// Run this example separately: go run basic_example.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/cache"
)

func main() {
	// Example 1: Basic usage with in-memory cache
	basicExample()

	// Example 2: Custom TTL configuration
	customTTLExample()

	// Example 3: Operation-specific caching
	operationSpecificExample()
}

// basicExample demonstrates the simplest usage of caching with Dodo Payments SDK.
func basicExample() {
	fmt.Println("=== Basic Caching Example ===\n")

	// Create a new Dodo Payments client
	client := dodopayments.NewClient()

	// Create an in-memory cache
	memCache := cache.NewMemoryCache()

	// Create cache configuration with defaults
	cacheConfig := cache.NewCacheConfig(memCache)
	cacheConfig.DefaultTTL = 5 * time.Minute

	// Wrap the client with caching
	_ = cache.NewCachedClient(client, cacheConfig)

	// Create a cached service wrapper
	cachedClient := cache.NewCachedClient(client, cacheConfig)
	paymentService := cache.NewCachedService(cachedClient, "Payment")

	ctx := context.Background()

	// First call - will hit the API and cache the result
	var payment dodopayments.Payment
	err := paymentService.Execute(
		ctx,
		"Get",
		"pay_123",
		&payment,
		func() error {
			// This is the actual API call
			_, err := client.Payments.Get(ctx, "pay_123")
			return err
		},
	)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Payment fetched (API call): %+v\n", payment.PaymentID)

	// Second call - will return from cache
	var cachedPayment dodopayments.Payment
	err = paymentService.Execute(
		ctx,
		"Get",
		"pay_123",
		&cachedPayment,
		func() error {
			// This won't be called on cache hit
			_, err := client.Payments.Get(ctx, "pay_123")
			return err
		},
	)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Payment fetched (from cache): %+v\n", cachedPayment.PaymentID)
	fmt.Println()
}

// customTTLExample demonstrates how to set custom TTLs for different operations.
func customTTLExample() {
	fmt.Println("=== Custom TTL Example ===\n")

	client := dodopayments.NewClient()
	memCache := cache.NewMemoryCache()

	// Create cache configuration with custom TTLs
	cacheConfig := cache.NewCacheConfig(memCache)
	cacheConfig.WithDefaultTTL(10 * time.Minute)
	cacheConfig.WithOperationTTL("Payment", "Get", 1*time.Minute)      // Short TTL for payments
	cacheConfig.WithOperationTTL("Product", "Get", 30*time.Minute)     // Longer TTL for products
	cacheConfig.WithOperationTTL("Product", "List", 15*time.Minute)    // Medium TTL for product lists
	cacheConfig.WithOperationTTL("Discount", "Get", 5*time.Minute)     // Short TTL for discounts

	_ = cache.NewCachedClient(client, cacheConfig)

	fmt.Printf("Cache configuration created with custom TTLs\n")
	fmt.Printf("- Payment.Get: 1 minute\n")
	fmt.Printf("- Product.Get: 30 minutes\n")
	fmt.Printf("- Product.List: 15 minutes\n")
	fmt.Printf("- Discount.Get: 5 minutes\n")
	fmt.Printf("- Default: 10 minutes\n")
	fmt.Println()
}

// operationSpecificExample demonstrates how to enable/disable caching for specific operations.
func operationSpecificExample() {
	fmt.Println("=== Operation-Specific Caching Example ===\n")

	client := dodopayments.NewClient()
	memCache := cache.NewMemoryCache()

	// Create cache configuration with operation-specific settings
	cacheConfig := cache.NewCacheConfig(memCache)
	// Enable caching for expensive operations
	cacheConfig.WithEnabledOperation("Subscription", "List")
	cacheConfig.WithEnabledOperation("License", "Validate")
	// Disable caching for operations that need fresh data
	cacheConfig.WithDisabledOperation("Payment", "Get")
	cacheConfig.WithDisabledOperation("WebhookEvent", "List")

	_ = cache.NewCachedClient(client, cacheConfig)

	fmt.Printf("Cache configuration created with operation-specific settings\n")
	fmt.Printf("Enabled: Subscription.List, License.Validate\n")
	fmt.Printf("Disabled: Payment.Get, WebhookEvent.List\n")
	fmt.Printf("Default rules apply to other operations\n")
	fmt.Println()
}
