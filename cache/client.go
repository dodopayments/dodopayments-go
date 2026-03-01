package cache

import (
	"context"
	"encoding/json"
	"fmt"
)

// CachedClient wraps a Dodo Payments client with caching functionality.
type CachedClient struct {
	client      interface{} // The actual Dodo Payments client
	config      *CacheConfig
	invalidator *Invalidator
}

// NewCachedClient creates a new cached client wrapper.
// The client parameter should be a pointer to the actual Dodo Payments client.
func NewCachedClient(client interface{}, config *CacheConfig) *CachedClient {
	return &CachedClient{
		client:      client,
		config:      config,
		invalidator: NewInvalidator(config),
	}
}

// GetConfig returns the cache configuration.
func (c *CachedClient) GetConfig() *CacheConfig {
	return c.config
}

// GetCache returns the underlying cache backend.
func (c *CachedClient) GetCache() Cache {
	return c.config.Cache
}

// InvalidateCache invalidates cache entries for a given operation.
func (c *CachedClient) InvalidateCache(ctx context.Context, service, method string) error {
	return c.invalidator.Invalidate(ctx, service, method)
}

// CachedService provides cached access to service methods.
type CachedService struct {
	client     *CachedClient
	serviceName string
}

// NewCachedService creates a new cached service wrapper.
func NewCachedService(client *CachedClient, serviceName string) *CachedService {
	return &CachedService{
		client:      client,
		serviceName: serviceName,
	}
}

// Execute performs a cached operation.
// This is a generic method that can be used to wrap any service call.
func (s *CachedService) Execute(
	ctx context.Context,
	method string,
	params interface{},
	result interface{},
	operation func() error,
) error {
	// Check if caching is enabled for this operation
	if !s.client.config.IsOperationEnabled(s.serviceName, method) {
		// Caching disabled, execute directly
		return operation()
	}
	
	// Generate cache key
	key := s.client.config.KeyGenerator(s.serviceName, method, params)
	
	// Try to get from cache first
	if cachedData, err := s.client.config.Cache.Get(ctx, key); err == nil && cachedData != nil {
		// Cache hit - unmarshal and return
		var cachedResp CachedResponse
		if err := json.Unmarshal(cachedData, &cachedResp); err == nil {
			if err := json.Unmarshal(cachedResp.Data, result); err == nil {
				// Successfully unmarshaled from cache
				return nil
			}
		}
		// If unmarshal fails, fall through to execute operation
	}
	
	// Cache miss - execute the operation
	if err := operation(); err != nil {
		return err
	}
	
	// Cache the result
	resultData, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal result for caching: %w", err)
	}
	
	cachedResp := CachedResponse{
		Data: json.RawMessage(resultData),
	}
	
	cachedData, err := json.Marshal(cachedResp)
	if err != nil {
		return fmt.Errorf("failed to marshal cached response: %w", err)
	}
	
	ttl := s.client.config.GetTTL(s.serviceName, method)
	if err := s.client.config.Cache.Set(ctx, key, cachedData, ttl); err != nil {
		// Log error but don't fail the operation
		// In production, you might want to add proper logging here
	}
	
	return nil
}

// ExecuteWithInvalidate executes a write operation and invalidates related cache entries.
func (s *CachedService) ExecuteWithInvalidate(
	ctx context.Context,
	method string,
	params interface{},
	result interface{},
	operation func() error,
) error {
	// Execute the operation first
	if err := operation(); err != nil {
		return err
	}
	
	// Invalidate related cache entries
	if s.client.config.InvalidationStrategy != TimeBased {
		if err := s.client.InvalidateCache(ctx, s.serviceName, method); err != nil {
			// Log error but don't fail the operation
			// In production, you might want to add proper logging here
		}
	}
	
	return nil
}
