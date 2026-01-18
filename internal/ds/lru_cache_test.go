package ds

import "testing"

func TestLRUCache_PutAndGet(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	if cache.Size() != 3 {
		t.Errorf("expected size 3, got %d", cache.Size())
	}

	val, ok := cache.Get("a")
	if !ok || val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	val, ok = cache.Get("b")
	if !ok || val != 2 {
		t.Errorf("expected 2, got %v", val)
	}
}

func TestLRUCache_GetNonExistent(t *testing.T) {
	cache := NewLRUCache(3)

	val, ok := cache.Get("nonexistent")
	if ok {
		t.Error("get should fail for non-existent key")
	}
	if val != nil {
		t.Errorf("expected nil, got %v", val)
	}
}

func TestLRUCache_UpdateValue(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("key", "value1")
	cache.Put("key", "value2")

	if cache.Size() != 1 {
		t.Errorf("expected size 1, got %d", cache.Size())
	}

	val, ok := cache.Get("key")
	if !ok || val != "value2" {
		t.Errorf("expected 'value2', got %v", val)
	}
}

func TestLRUCache_Eviction(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", 1)
	cache.Put("b", 2)

	// Access "a" to make it recently used
	cache.Get("a")

	// Add "c", should evict "b" (least recently used)
	cache.Put("c", 3)

	if cache.Size() != 2 {
		t.Errorf("expected size 2, got %d", cache.Size())
	}

	// "b" should be evicted
	_, ok := cache.Get("b")
	if ok {
		t.Error("'b' should have been evicted")
	}

	// "a" and "c" should exist
	if _, ok := cache.Get("a"); !ok {
		t.Error("'a' should still exist")
	}
	if _, ok := cache.Get("c"); !ok {
		t.Error("'c' should still exist")
	}
}

func TestLRUCache_EvictionOrder(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Add "d", should evict "a" (oldest)
	cache.Put("d", 4)

	_, ok := cache.Get("a")
	if ok {
		t.Error("'a' should have been evicted")
	}

	// b, c, d should exist
	if _, ok := cache.Get("b"); !ok {
		t.Error("'b' should exist")
	}
	if _, ok := cache.Get("c"); !ok {
		t.Error("'c' should exist")
	}
	if _, ok := cache.Get("d"); !ok {
		t.Error("'d' should exist")
	}
}

func TestLRUCache_GetUpdatesOrder(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", 1)
	cache.Put("b", 2)

	// Access "a" to make it recently used
	cache.Get("a")

	// Add "c", should evict "b"
	cache.Put("c", 3)

	_, ok := cache.Get("b")
	if ok {
		t.Error("'b' should have been evicted")
	}

	if _, ok := cache.Get("a"); !ok {
		t.Error("'a' should still exist")
	}
	if _, ok := cache.Get("c"); !ok {
		t.Error("'c' should still exist")
	}
}

func TestLRUCache_Delete(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	if !cache.Delete("b") {
		t.Error("delete should succeed")
	}

	if cache.Size() != 2 {
		t.Errorf("expected size 2, got %d", cache.Size())
	}

	_, ok := cache.Get("b")
	if ok {
		t.Error("'b' should be deleted")
	}
}

func TestLRUCache_DeleteNonExistent(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)

	if cache.Delete("nonexistent") {
		t.Error("delete of non-existent key should fail")
	}

	if cache.Size() != 1 {
		t.Errorf("size should remain 1, got %d", cache.Size())
	}
}

func TestLRUCache_Clear(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	cache.Clear()

	if cache.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", cache.Size())
	}

	_, ok := cache.Get("a")
	if ok {
		t.Error("cache should be empty after clear")
	}
}

func TestLRUCache_Capacity(t *testing.T) {
	cache := NewLRUCache(5)

	if cache.Capacity() != 5 {
		t.Errorf("expected capacity 5, got %d", cache.Capacity())
	}
}

func TestLRUCache_Keys(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	keys := cache.Keys()

	if len(keys) != 3 {
		t.Errorf("expected 3 keys, got %d", len(keys))
	}

	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	if !keyMap["a"] || !keyMap["b"] || !keyMap["c"] {
		t.Error("not all keys returned")
	}
}

func TestLRUCache_GetOldest(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	oldest, ok := cache.GetOldest()
	if !ok || oldest != "a" {
		t.Errorf("expected oldest 'a', got %s", oldest)
	}
}

func TestLRUCache_GetNewest(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	newest, ok := cache.GetNewest()
	if !ok || newest != "c" {
		t.Errorf("expected newest 'c', got %s", newest)
	}
}

func TestLRUCache_GetOldestEmpty(t *testing.T) {
	cache := NewLRUCache(3)

	_, ok := cache.GetOldest()
	if ok {
		t.Error("should fail on empty cache")
	}
}

func TestLRUCache_GetNewestEmpty(t *testing.T) {
	cache := NewLRUCache(3)

	_, ok := cache.GetNewest()
	if ok {
		t.Error("should fail on empty cache")
	}
}

func TestLRUCache_CapacityOne(t *testing.T) {
	cache := NewLRUCache(1)

	cache.Put("a", 1)
	cache.Put("b", 2)

	if cache.Size() != 1 {
		t.Errorf("expected size 1, got %d", cache.Size())
	}

	// Only "b" should exist
	_, ok := cache.Get("a")
	if ok {
		t.Error("'a' should have been evicted")
	}

	val, ok := cache.Get("b")
	if !ok || val != 2 {
		t.Errorf("expected 2, got %v", val)
	}
}

func TestLRUCache_ComplexScenario(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Access "a" to make it recently used
	cache.Get("a")

	// Update "b"
	cache.Put("b", 20)

	// Add "d", should evict "c" (oldest)
	cache.Put("d", 4)

	// Verify state
	_, ok := cache.Get("c")
	if ok {
		t.Error("'c' should have been evicted")
	}

	val, _ := cache.Get("a")
	if val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	val, _ = cache.Get("b")
	if val != 20 {
		t.Errorf("expected 20, got %v", val)
	}

	val, _ = cache.Get("d")
	if val != 4 {
		t.Errorf("expected 4, got %v", val)
	}
}

func TestLRUCache_PutUpdatesOrder(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put("a", 1)
	cache.Put("b", 2)

	// Update "a" with Put
	cache.Put("a", 10)

	// Add "c", should evict "b"
	cache.Put("c", 3)

	_, ok := cache.Get("b")
	if ok {
		t.Error("'b' should have been evicted")
	}

	val, _ := cache.Get("a")
	if val != 10 {
		t.Errorf("expected 10, got %v", val)
	}
}
