package systemdesign

import (
	"testing"
	"time"
)

func TestInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache()

	// 1. Set and Get valid
	cache.Set("key1", "value1", 1*time.Second)

	val, ok := cache.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("expected value1, got %v", val)
	}

	// 2. Expiration
	// Sleep to let it expire
	time.Sleep(1100 * time.Millisecond)

	val, ok = cache.Get("key1")
	if ok {
		t.Error("expected key to be expired")
	}

	// 3. Delete
	cache.Set("key2", "value2", 1*time.Minute)
	cache.Delete("key2")

	_, ok = cache.Get("key2")
	if ok {
		t.Error("expected key to be deleted")
	}
}
