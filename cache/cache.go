// Package cache provides a caching layer for the Dodo Payments Go SDK.
// It supports multiple cache backends (memory, Redis) and cache invalidation strategies.
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"time"
)

// Cache defines the interface for cache storage backends.
type Cache interface {
	// Get retrieves a value from the cache.
	// Returns nil if the key doesn't exist.
	Get(ctx context.Context, key string) ([]byte, error)

	// Set stores a value in the cache with an expiration time.
	// If ttl is 0, the item doesn't expire.
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error

	// Delete removes a value from the cache.
	Delete(ctx context.Context, key string) error

	// DeletePrefix removes all keys with a given prefix.
	DeletePrefix(ctx context.Context, prefix string) error

	// Clear removes all items from the cache.
	Clear(ctx context.Context) error
}

// CacheConfig holds configuration for the caching layer.
type CacheConfig struct {
	// Cache is the backend cache storage.
	Cache Cache

	// DefaultTTL is the default time-to-live for cache entries.
	// If zero, entries don't expire by default.
	DefaultTTL time.Duration

	// TTLPerOperation specifies custom TTLs for specific operations.
	// Key format: "Service.Method" (e.g., "Payment.Get", "Product.List")
	TTLPerOperation map[string]time.Duration

	// EnabledOperations specifies which operations should be cached.
	// If nil, all safe operations (GET, LIST) are cached.
	// Key format: "Service.Method" (e.g., "Payment.Get", "Product.List")
	EnabledOperations map[string]bool

	// DisabledOperations specifies which operations should never be cached.
	// Takes precedence over EnabledOperations.
	// Key format: "Service.Method"
	DisabledOperations map[string]bool

	// InvalidationStrategy determines how cache invalidation works.
	InvalidationStrategy InvalidationStrategy

	// KeyGenerator is used to generate cache keys.
	// If nil, DefaultKeyGenerator is used.
	KeyGenerator KeyGenerator
}

// InvalidationStrategy determines how to invalidate cache entries.
type InvalidationStrategy int

const (
	// TimeBased invalidation uses TTL only.
	TimeBased InvalidationStrategy = iota

	// OperationBased invalidation invalidates related entries on write operations.
	// e.g., creating a payment invalidates payment list cache.
	OperationBased

	// Hybrid uses both time and operation-based invalidation.
	Hybrid
)

// KeyGenerator generates cache keys from operation metadata.
type KeyGenerator func(service, method string, params interface{}) string

// DefaultKeyGenerator is the default cache key generator.
func DefaultKeyGenerator(service, method string, params interface{}) string {
	// Generate a base key
	base := fmt.Sprintf("dodo:%s:%s:", service, method)

	// If no params, return base key
	if params == nil {
		return base + "nil"
	}

	// Hash params to create a consistent key
	hash := fnv.New64a()
	encoder := json.NewEncoder(hash)
	if err := encoder.Encode(params); err != nil {
		// Fallback to string representation if JSON encoding fails
		return fmt.Sprintf("%s%v", base, params)
	}

	return fmt.Sprintf("%s%x", base, hash.Sum64())
}

// CachedResponse wraps cached data with metadata.
type CachedResponse struct {
	Data      json.RawMessage `json:"data"`
	Timestamp time.Time       `json:"timestamp"`
}

// NewCacheConfig creates a new CacheConfig with sensible defaults.
func NewCacheConfig(cache Cache) *CacheConfig {
	return &CacheConfig{
		Cache:                cache,
		DefaultTTL:           5 * time.Minute,
		TTLPerOperation:      make(map[string]time.Duration),
		EnabledOperations:    nil,
		DisabledOperations:   make(map[string]bool),
		InvalidationStrategy: Hybrid,
		KeyGenerator:         DefaultKeyGenerator,
	}
}

// IsOperationEnabled checks if an operation should be cached.
func (c *CacheConfig) IsOperationEnabled(service, method string) bool {
	key := service + "." + method

	// Check disabled operations first (highest priority)
	if c.DisabledOperations != nil {
		if _, disabled := c.DisabledOperations[key]; disabled {
			return false
		}
	}

	// If no explicit enabled list, cache all safe operations
	if c.EnabledOperations == nil {
		return isSafeOperation(method)
	}

	// Check explicit enabled list
	enabled, ok := c.EnabledOperations[key]
	if !ok {
		return false
	}

	return enabled
}

// GetTTL returns the TTL for a specific operation.
func (c *CacheConfig) GetTTL(service, method string) time.Duration {
	key := service + "." + method
	if ttl, ok := c.TTLPerOperation[key]; ok {
		return ttl
	}
	return c.DefaultTTL
}

// isSafeOperation determines if an operation is safe to cache.
// Safe operations are typically read-only (GET, LIST).
func isSafeOperation(method string) bool {
	switch method {
	case "Get", "List", "ListAutoPaging", "Validate":
		return true
	default:
		return false
	}
}

// WithDefaultTTL sets the default TTL for cache entries.
func (c *CacheConfig) WithDefaultTTL(ttl time.Duration) *CacheConfig {
	c.DefaultTTL = ttl
	return c
}

// WithOperationTTL sets a custom TTL for a specific operation.
func (c *CacheConfig) WithOperationTTL(service, method string, ttl time.Duration) *CacheConfig {
	if c.TTLPerOperation == nil {
		c.TTLPerOperation = make(map[string]time.Duration)
	}
	c.TTLPerOperation[service+"."+method] = ttl
	return c
}

// WithEnabledOperation enables caching for a specific operation.
func (c *CacheConfig) WithEnabledOperation(service, method string) *CacheConfig {
	if c.EnabledOperations == nil {
		c.EnabledOperations = make(map[string]bool)
	}
	c.EnabledOperations[service+"."+method] = true
	return c
}

// WithDisabledOperation disables caching for a specific operation.
func (c *CacheConfig) WithDisabledOperation(service, method string) *CacheConfig {
	if c.DisabledOperations == nil {
		c.DisabledOperations = make(map[string]bool)
	}
	c.DisabledOperations[service+"."+method] = true
	return c
}

// WithInvalidationStrategy sets the cache invalidation strategy.
func (c *CacheConfig) WithInvalidationStrategy(strategy InvalidationStrategy) *CacheConfig {
	c.InvalidationStrategy = strategy
	return c
}

// WithKeyGenerator sets a custom cache key generator.
func (c *CacheConfig) WithKeyGenerator(generator KeyGenerator) *CacheConfig {
	c.KeyGenerator = generator
	return c
}
