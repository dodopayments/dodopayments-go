package cache

import (
	"context"
	"testing"
	"time"
)

// TestMemoryCache_BasicOperations tests basic memory cache operations.
func TestMemoryCache_BasicOperations(t *testing.T) {
	cache := NewMemoryCache()
	ctx := context.Background()

	// Test Set and Get
	key := "test:key"
	value := []byte("test value")
	ttl := time.Minute

	err := cache.Set(ctx, key, value, ttl)
	if err != nil {
		t.Fatalf("Failed to set value: %v", err)
	}

	retrieved, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if retrieved == nil {
		t.Fatal("Expected non-nil value")
	}

	if string(retrieved) != string(value) {
		t.Errorf("Expected %s, got %s", string(value), string(retrieved))
	}
}

// TestMemoryCache_Expiration tests cache entry expiration.
func TestMemoryCache_Expiration(t *testing.T) {
	cache := NewMemoryCache()
	ctx := context.Background()

	key := "test:expired"
	value := []byte("expired value")
	ttl := 10 * time.Millisecond

	err := cache.Set(ctx, key, value, ttl)
	if err != nil {
		t.Fatalf("Failed to set value: %v", err)
	}

	// Wait for expiration
	time.Sleep(20 * time.Millisecond)

	retrieved, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if retrieved != nil {
		t.Error("Expected nil value after expiration, got non-nil")
	}
}

// TestMemoryCache_Delete tests deletion from cache.
func TestMemoryCache_Delete(t *testing.T) {
	cache := NewMemoryCache()
	ctx := context.Background()

	key := "test:delete"
	value := []byte("delete me")

	err := cache.Set(ctx, key, value, time.Minute)
	if err != nil {
		t.Fatalf("Failed to set value: %v", err)
	}

	err = cache.Delete(ctx, key)
	if err != nil {
		t.Fatalf("Failed to delete value: %v", err)
	}

	retrieved, err := cache.Get(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get value: %v", err)
	}

	if retrieved != nil {
		t.Error("Expected nil value after deletion")
	}
}

// TestMemoryCache_DeletePrefix tests prefix-based deletion.
func TestMemoryCache_DeletePrefix(t *testing.T) {
	cache := NewMemoryCache()
	ctx := context.Background()

	// Set multiple keys with the same prefix
	prefix := "test:prefix"
	keys := []string{
		prefix + ":key1",
		prefix + ":key2",
		prefix + ":key3",
		"other:key",
	}

	for _, key := range keys {
		value := []byte(key)
		err := cache.Set(ctx, key, value, time.Minute)
		if err != nil {
			t.Fatalf("Failed to set value for key %s: %v", key, err)
		}
	}

	// Delete all keys with prefix
	err := cache.DeletePrefix(ctx, prefix+":")
	if err != nil {
		t.Fatalf("Failed to delete prefix: %v", err)
	}

	// Verify deletion
	for _, key := range keys[:3] {
		retrieved, err := cache.Get(ctx, key)
		if err != nil {
			t.Fatalf("Failed to get value for key %s: %v", key, err)
		}
		if retrieved != nil {
			t.Errorf("Expected nil value for key %s after prefix deletion", key)
		}
	}

	// Verify other key still exists
	retrieved, err := cache.Get(ctx, "other:key")
	if err != nil {
		t.Fatalf("Failed to get value for other key: %v", err)
	}
	if retrieved == nil {
		t.Error("Expected non-nil value for other key")
	}
}

// TestMemoryCache_Clear tests clearing the entire cache.
func TestMemoryCache_Clear(t *testing.T) {
	cache := NewMemoryCache()
	ctx := context.Background()

	// Set some values
	for i := 0; i < 10; i++ {
		key := "test:clear:" + string(rune('0'+i))
		value := []byte("value")
		err := cache.Set(ctx, key, value, time.Minute)
		if err != nil {
			t.Fatalf("Failed to set value: %v", err)
		}
	}

	// Clear cache
	err := cache.Clear(ctx)
	if err != nil {
		t.Fatalf("Failed to clear cache: %v", err)
	}

	// Verify all values are gone
	if cache.Size() != 0 {
		t.Errorf("Expected cache size 0 after clear, got %d", cache.Size())
	}
}

// TestCacheConfig_Defaults tests default cache configuration.
func TestCacheConfig_Defaults(t *testing.T) {
	memCache := NewMemoryCache()
	config := NewCacheConfig(memCache)

	if config.DefaultTTL != 5*time.Minute {
		t.Errorf("Expected default TTL 5m, got %v", config.DefaultTTL)
	}

	if config.InvalidationStrategy != Hybrid {
		t.Errorf("Expected default strategy Hybrid, got %v", config.InvalidationStrategy)
	}

	if config.KeyGenerator == nil {
		t.Error("Expected default key generator")
	}
}

// TestCacheConfig_IsOperationEnabled tests operation enablement logic.
func TestCacheConfig_IsOperationEnabled(t *testing.T) {
	memCache := NewMemoryCache()
	config := NewCacheConfig(memCache)

	tests := []struct {
		name     string
		service  string
		method   string
		expected bool
	}{
		{"Get should be enabled by default", "Payment", "Get", true},
		{"List should be enabled by default", "Product", "List", true},
		{"New should not be enabled by default", "Payment", "New", false},
		{"Update should not be enabled by default", "Customer", "Update", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := config.IsOperationEnabled(tt.service, tt.method)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestCacheConfig_WithEnabledOperation tests explicit operation enablement.
func TestCacheConfig_WithEnabledOperation(t *testing.T) {
	memCache := NewMemoryCache()
	config := NewCacheConfig(memCache)

	// Enable New operation (which is disabled by default)
	config.WithEnabledOperation("Payment", "New")

	result := config.IsOperationEnabled("Payment", "New")
	if !result {
		t.Error("Expected New operation to be enabled")
	}
}

// TestCacheConfig_WithDisabledOperation tests explicit operation disablement.
func TestCacheConfig_WithDisabledOperation(t *testing.T) {
	memCache := NewMemoryCache()
	config := NewCacheConfig(memCache)

	// Disable Get operation (which is enabled by default)
	config.WithDisabledOperation("Payment", "Get")

	result := config.IsOperationEnabled("Payment", "Get")
	if result {
		t.Error("Expected Get operation to be disabled")
	}
}

// TestCacheConfig_GetTTL tests TTL retrieval.
func TestCacheConfig_GetTTL(t *testing.T) {
	memCache := NewMemoryCache()
	config := NewCacheConfig(memCache)

	// Test default TTL
	ttl := config.GetTTL("Payment", "Get")
	if ttl != config.DefaultTTL {
		t.Errorf("Expected default TTL %v, got %v", config.DefaultTTL, ttl)
	}

	// Test custom TTL
	customTTL := 10 * time.Minute
	config.WithOperationTTL("Payment", "Get", customTTL)

	ttl = config.GetTTL("Payment", "Get")
	if ttl != customTTL {
		t.Errorf("Expected custom TTL %v, got %v", customTTL, ttl)
	}
}

// TestDefaultKeyGenerator tests the default cache key generator.
func TestDefaultKeyGenerator(t *testing.T) {
	tests := []struct {
		name     string
		service  string
		method   string
		params   interface{}
		expected string
	}{
		{
			name:    "No params",
			service: "Payment",
			method:  "Get",
			params:  nil,
			expected: "dodo:Payment:Get:nil",
		},
		{
			name:    "With params",
			service: "Product",
			method:  "List",
			params:  map[string]string{"limit": "10"},
			expected: "dodo:Product:list:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DefaultKeyGenerator(tt.service, tt.method, tt.params)
			if len(result) < len(tt.expected) {
				t.Errorf("Expected key to start with %s, got %s", tt.expected, result)
			}
		})
	}
}

// TestInvalidationRegistry tests the invalidation registry.
func TestInvalidationRegistry(t *testing.T) {
	registry := NewInvalidationRegistry()

	// Test default rule for Payment.New
	invalidations := registry.GetInvalidations("Payment.New")
	if len(invalidations) == 0 {
		t.Error("Expected invalidations for Payment.New")
	}

	// Test custom rule
	registry.AddRule(InvalidationRule{
		Operation:   "Test.Create",
		Invalidates: []string{"Test.Get", "Test.List"},
	})

	invalidations = registry.GetInvalidations("Test.Create")
	if len(invalidations) != 2 {
		t.Errorf("Expected 2 invalidations, got %d", len(invalidations))
	}
}

// BenchmarkMemoryCache_Set benchmarks cache Set operation.
func BenchmarkMemoryCache_Set(b *testing.B) {
	cache := NewMemoryCache()
	ctx := context.Background()
	value := []byte("benchmark value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "bench:set:" + string(rune('0'+i%10))
		cache.Set(ctx, key, value, time.Minute)
	}
}

// BenchmarkMemoryCache_Get benchmarks cache Get operation.
func BenchmarkMemoryCache_Get(b *testing.B) {
	cache := NewMemoryCache()
	ctx := context.Background()
	key := "bench:get"
	value := []byte("benchmark value")
	cache.Set(ctx, key, value, time.Minute)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(ctx, key)
	}
}
