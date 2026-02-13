// Package main demonstrates advanced caching patterns with Dodo Payments SDK.
// Run this example separately: go run advanced_example.go
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
	// Example 1: Redis cache with distributed caching
	redisExample()

	// Example 2: Cache invalidation strategies
	invalidationExample()

	// Example 3: Custom key generator
	customKeyGeneratorExample()

	// Example 4: Hybrid invalidation strategy
	hybridStrategyExample()
}

// redisExample demonstrates using Redis as a cache backend.
func redisExample() {
	fmt.Println("=== Redis Cache Example ===\n")

	// Note: This is a simplified example. In production, you would use an actual Redis client.
	// Example with go-redis/redis/v9:
	/*
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		redisCache := cache.NewRedisCache(rdb)
	*/

	// For this example, we'll use memory cache
	memCache := cache.NewMemoryCache()

	// Configure cache with Redis backend
	cacheConfig := cache.NewCacheConfig(memCache)
	cacheConfig.WithDefaultTTL(10 * time.Minute)
	cacheConfig.WithInvalidationStrategy(cache.Hybrid) // Use hybrid invalidation with Redis

	_ = cache.NewCachedClient(dodopayments.NewClient(), cacheConfig)

	fmt.Printf("Redis cache configured for distributed caching\n")
	fmt.Printf("Invalidation strategy: Hybrid\n")
	fmt.Printf("This ensures cache consistency across multiple instances\n")
	fmt.Println()
}

// invalidationExample demonstrates cache invalidation strategies.
func invalidationExample() {
	fmt.Println("=== Cache Invalidation Strategies ===\n")

	memCache := cache.NewMemoryCache()

	// Time-based invalidation (TTL only)
	timeBasedConfig := cache.NewCacheConfig(memCache)
	timeBasedConfig.WithDefaultTTL(5 * time.Minute)
	timeBasedConfig.WithInvalidationStrategy(cache.TimeBased)

	_ = timeBasedConfig

	fmt.Println("Time-based invalidation:")
	fmt.Println("- Cache entries expire after TTL")
	fmt.Println("- No automatic invalidation on write operations")
	fmt.Println("- Simple and predictable")

	// Operation-based invalidation
	operationBasedConfig := cache.NewCacheConfig(memCache)
	operationBasedConfig.WithDefaultTTL(1 * time.Hour)
	operationBasedConfig.WithInvalidationStrategy(cache.OperationBased)

	_ = operationBasedConfig

	fmt.Println("\nOperation-based invalidation:")
	fmt.Println("- Cache entries invalidated on related write operations")
	fmt.Println("- Example: Creating a payment invalidates payment list cache")
	fmt.Println("- Ensures data consistency without waiting for TTL")

	// Hybrid strategy (recommended)
	hybridConfig := cache.NewCacheConfig(memCache)
	hybridConfig.WithDefaultTTL(5 * time.Minute)
	hybridConfig.WithInvalidationStrategy(cache.Hybrid)

	_ = hybridConfig

	fmt.Println("\nHybrid invalidation (recommended):")
	fmt.Println("- Combines TTL and operation-based invalidation")
	fmt.Println("- Best of both worlds")
	fmt.Println("- Automatic fallback to TTL if invalidation fails")
	fmt.Println()
}

// customKeyGeneratorExample demonstrates using a custom cache key generator.
func customKeyGeneratorExample() {
	fmt.Println("=== Custom Key Generator Example ===\n")

	client := dodopayments.NewClient()
	memCache := cache.NewMemoryCache()

	// Custom key generator that includes user ID for multi-tenant caching
	customGenerator := func(service, method string, params interface{}) string {
		// In a real application, you would extract user ID from context
		// userID := getUserIDFromContext(ctx)
		userID := "user_123"
		return fmt.Sprintf("tenant:%s:dodo:%s:%s:%v", userID, service, method, params)
	}

	cacheConfig := cache.NewCacheConfig(memCache)
	cacheConfig.WithDefaultTTL(5 * time.Minute)
	cacheConfig.WithKeyGenerator(customGenerator)

	_ = cache.NewCachedClient(client, cacheConfig)

	fmt.Println("Custom key generator configured")
	fmt.Println("Keys include tenant/user ID for multi-tenant isolation")
	fmt.Println("Example key: tenant:user_123:dodo:Payment:Get:pay_123")
	fmt.Println()
}

// hybridStrategyExample demonstrates the hybrid invalidation strategy in action.
func hybridStrategyExample() {
	fmt.Println("=== Hybrid Strategy Example ===\n")

	client := dodopayments.NewClient()
	memCache := cache.NewMemoryCache()

	// Configure with hybrid strategy
	cacheConfig := cache.NewCacheConfig(memCache)
	cacheConfig.WithDefaultTTL(5 * time.Minute)
	cacheConfig.WithInvalidationStrategy(cache.Hybrid)

	cachedClient := cache.NewCachedClient(client, cacheConfig)

	ctx := context.Background()
	productService := cache.NewCachedService(cachedClient, "Product")

	// Scenario 1: Read product (cached)
	fmt.Println("1. Fetching product list (will cache)...")
	var products dodopayments.ProductListResponse
	err := productService.Execute(
		ctx,
		"List",
		dodopayments.ProductListParams{},
		&products,
		func() error {
			// Actual API call
			_, err := client.Products.List(ctx, dodopayments.ProductListParams{})
			return err
		},
	)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println("   Product list cached")

	// Scenario 2: Create new product (invalidates product cache)
	fmt.Println("\n2. Creating new product...")
	var newProduct dodopayments.Product
	err = productService.ExecuteWithInvalidate(
		ctx,
		"New",
		dodopayments.ProductNewParams{
			Name: dodopayments.F("New Product"),
		},
		&newProduct,
		func() error {
			// Actual API call to create product
			_, err := client.Products.New(ctx, dodopayments.ProductNewParams{})
			return err
		},
	)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println("   Product created")
	fmt.Println("   Product list cache invalidated")

	// Scenario 3: Read product list again (fresh from API)
	fmt.Println("\n3. Fetching product list again (cache miss, fresh data)...")
	_ = productService.Execute(
		ctx,
		"List",
		dodopayments.ProductListParams{},
		&products,
		func() error {
			_, err := client.Products.List(ctx, dodopayments.ProductListParams{})
			return err
		},
	)
	fmt.Println("   Fresh product list fetched from API")
	fmt.Println("   New list cached")

	fmt.Println("\nHybrid strategy ensures:")
	fmt.Println("- Fast reads from cache")
	fmt.Println("- Automatic invalidation on writes")
	fmt.Println("- Fallback to TTL if needed")
	fmt.Println()
}
