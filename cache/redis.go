package cache

import (
	"context"
	"fmt"
	"time"
)

// RedisClient defines the interface for Redis client operations.
// This allows the cache to work with different Redis client libraries.
type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, keys ...string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	FlushDB(ctx context.Context) error
}

// RedisCache is a Redis implementation of the Cache interface.
// It's suitable for distributed applications and provides persistence.
type RedisCache struct {
	client RedisClient
}

// NewRedisCache creates a new Redis cache.
func NewRedisCache(client RedisClient) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

// Get retrieves a value from Redis.
func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := r.client.Get(ctx, key)
	if err != nil {
		// Return nil for key not found (common Redis behavior)
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, fmt.Errorf("redis get error: %w", err)
	}
	
	if val == "" {
		return nil, nil
	}
	
	return []byte(val), nil
}

// Set stores a value in Redis with an expiration time.
func (r *RedisCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	err := r.client.Set(ctx, key, value, ttl)
	if err != nil {
		return fmt.Errorf("redis set error: %w", err)
	}
	return nil
}

// Delete removes a value from Redis.
func (r *RedisCache) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key)
	if err != nil {
		return fmt.Errorf("redis del error: %w", err)
	}
	return nil
}

// DeletePrefix removes all keys with a given prefix.
func (r *RedisCache) DeletePrefix(ctx context.Context, prefix string) error {
	// Use SCAN to find all keys with the prefix
	keys, err := r.client.Keys(ctx, prefix+"*")
	if err != nil {
		return fmt.Errorf("redis keys error: %w", err)
	}
	
	// Delete all matching keys
	if len(keys) > 0 {
		err = r.client.Del(ctx, keys...)
		if err != nil {
			return fmt.Errorf("redis del error: %w", err)
		}
	}
	
	return nil
}

// Clear removes all items from the current Redis database.
func (r *RedisCache) Clear(ctx context.Context) error {
	err := r.client.FlushDB(ctx)
	if err != nil {
		return fmt.Errorf("redis flushdb error: %w", err)
	}
	return nil
}
