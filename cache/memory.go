package cache

import (
	"context"
	"sync"
	"time"
)

// MemoryCache is an in-memory implementation of the Cache interface.
// It's thread-safe and suitable for single-instance applications.
type MemoryCache struct {
	mu    sync.RWMutex
	items map[string]*cacheItem
}

// cacheItem represents a single item in the cache.
type cacheItem struct {
	value      []byte
	expiration time.Time
}

// NewMemoryCache creates a new in-memory cache.
func NewMemoryCache() *MemoryCache {
	cache := &MemoryCache{
		items: make(map[string]*cacheItem),
	}
	
	// Start a goroutine to clean up expired items
	go cache.cleanupExpiredItems()
	
	return cache
}

// Get retrieves a value from the cache.
func (m *MemoryCache) Get(ctx context.Context, key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	item, ok := m.items[key]
	if !ok {
		return nil, nil
	}
	
	// Check if item has expired
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return nil, nil
	}
	
	return item.value, nil
}

// Set stores a value in the cache with an expiration time.
func (m *MemoryCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	var expiration time.Time
	if ttl > 0 {
		expiration = time.Now().Add(ttl)
	}
	
	m.items[key] = &cacheItem{
		value:      value,
		expiration: expiration,
	}
	
	return nil
}

// Delete removes a value from the cache.
func (m *MemoryCache) Delete(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	delete(m.items, key)
	return nil
}

// DeletePrefix removes all keys with a given prefix.
func (m *MemoryCache) DeletePrefix(ctx context.Context, prefix string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	for key := range m.items {
		if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			delete(m.items, key)
		}
	}
	
	return nil
}

// Clear removes all items from the cache.
func (m *MemoryCache) Clear(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.items = make(map[string]*cacheItem)
	return nil
}

// cleanupExpiredItems periodically removes expired items from the cache.
func (m *MemoryCache) cleanupExpiredItems() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		m.mu.Lock()
		now := time.Now()
		for key, item := range m.items {
			if !item.expiration.IsZero() && now.After(item.expiration) {
				delete(m.items, key)
			}
		}
		m.mu.Unlock()
	}
}

// Size returns the number of items in the cache.
func (m *MemoryCache) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.items)
}
