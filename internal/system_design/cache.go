package systemdesign

import (
	"sync"
	"time"
)

// Why interviewers ask this:
// Caching is a go-to solution for performance problems. Interviewers want to see if you
// can implement an expiration policy (TTL) safely in a concurrent environment.
// It also tests your understanding of memory management (how to cleanup expired items).

// Common pitfalls:
// - Deadlocks (holding mutex during long operations)
// - Memory leaks (never cleaning up expired items)
// - Race conditions on map access

// Key takeaway:
// Use `sync.RWMutex` for concurrent map access.
// Store {Value, ExpirationTime} in the map.
// Validate expiration on Get access (Lazy) OR run a cleanup goroutine (Active).

type CacheItem struct {
	Value      interface{}
	Expiration int64 // Unix nanoseconds
}

type InMemoryCache struct {
	mu    sync.RWMutex
	items map[string]CacheItem
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		items: make(map[string]CacheItem),
	}
}

// Set adds a key-value pair with a TTL
func (c *InMemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
}

// Get retrieves a value if it exists and hasn't expired
func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	// Check Expiration
	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

// Delete removes an item (manual invalidation)
func (c *InMemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}
