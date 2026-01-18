package ds

import "testing"

func TestHashMap_PutAndGet(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("name", "John")
	hm.Put("age", 30)
	hm.Put("city", "NYC")

	if hm.Size() != 3 {
		t.Errorf("expected size 3, got %d", hm.Size())
	}

	val, ok := hm.Get("name")
	if !ok || val != "John" {
		t.Errorf("expected 'John', got %v", val)
	}

	val, ok = hm.Get("age")
	if !ok || val != 30 {
		t.Errorf("expected 30, got %v", val)
	}
}

func TestHashMap_GetNonExistent(t *testing.T) {
	hm := NewHashMap(4)

	val, ok := hm.Get("nonexistent")
	if ok {
		t.Error("get should fail for non-existent key")
	}
	if val != nil {
		t.Errorf("expected nil, got %v", val)
	}
}

func TestHashMap_UpdateValue(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("key", "value1")
	hm.Put("key", "value2")

	if hm.Size() != 1 {
		t.Errorf("expected size 1, got %d", hm.Size())
	}

	val, ok := hm.Get("key")
	if !ok || val != "value2" {
		t.Errorf("expected 'value2', got %v", val)
	}
}

func TestHashMap_Delete(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("key1", "value1")
	hm.Put("key2", "value2")
	hm.Put("key3", "value3")

	if !hm.Delete("key2") {
		t.Error("delete should succeed")
	}

	if hm.Size() != 2 {
		t.Errorf("expected size 2, got %d", hm.Size())
	}

	if hm.Contains("key2") {
		t.Error("key2 should be deleted")
	}

	if !hm.Contains("key1") || !hm.Contains("key3") {
		t.Error("key1 and key3 should still exist")
	}
}

func TestHashMap_DeleteNonExistent(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("key", "value")

	if hm.Delete("nonexistent") {
		t.Error("delete of non-existent key should fail")
	}

	if hm.Size() != 1 {
		t.Errorf("size should remain 1, got %d", hm.Size())
	}
}

func TestHashMap_Contains(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("exists", "value")

	if !hm.Contains("exists") {
		t.Error("should contain 'exists'")
	}

	if hm.Contains("notexists") {
		t.Error("should not contain 'notexists'")
	}
}

func TestHashMap_IsEmpty(t *testing.T) {
	hm := NewHashMap(4)

	if !hm.IsEmpty() {
		t.Error("new map should be empty")
	}

	hm.Put("key", "value")
	if hm.IsEmpty() {
		t.Error("map with entry should not be empty")
	}

	hm.Delete("key")
	if !hm.IsEmpty() {
		t.Error("map should be empty after deleting last entry")
	}
}

func TestHashMap_Clear(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("key1", "value1")
	hm.Put("key2", "value2")
	hm.Put("key3", "value3")

	hm.Clear()

	if !hm.IsEmpty() {
		t.Error("map should be empty after clear")
	}

	if hm.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", hm.Size())
	}
}

func TestHashMap_Keys(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("key1", "value1")
	hm.Put("key2", "value2")
	hm.Put("key3", "value3")

	keys := hm.Keys()

	if len(keys) != 3 {
		t.Errorf("expected 3 keys, got %d", len(keys))
	}

	// Check all keys are present
	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	if !keyMap["key1"] || !keyMap["key2"] || !keyMap["key3"] {
		t.Error("not all keys returned")
	}
}

func TestHashMap_Collisions(t *testing.T) {
	// Small capacity to force collisions
	hm := NewHashMap(2)

	hm.Put("key1", "value1")
	hm.Put("key2", "value2")
	hm.Put("key3", "value3")
	hm.Put("key4", "value4")

	// All values should be retrievable despite collisions
	val, ok := hm.Get("key1")
	if !ok || val != "value1" {
		t.Errorf("expected 'value1', got %v", val)
	}

	val, ok = hm.Get("key2")
	if !ok || val != "value2" {
		t.Errorf("expected 'value2', got %v", val)
	}

	val, ok = hm.Get("key3")
	if !ok || val != "value3" {
		t.Errorf("expected 'value3', got %v", val)
	}

	val, ok = hm.Get("key4")
	if !ok || val != "value4" {
		t.Errorf("expected 'value4', got %v", val)
	}
}

func TestHashMap_Resize(t *testing.T) {
	hm := NewHashMap(2)

	// Insert enough elements to trigger resize
	for i := 0; i < 10; i++ {
		key := string(rune('a' + i))
		hm.Put(key, i)
	}

	// All values should still be retrievable after resize
	for i := 0; i < 10; i++ {
		key := string(rune('a' + i))
		val, ok := hm.Get(key)
		if !ok || val != i {
			t.Errorf("expected %d for key %s, got %v", i, key, val)
		}
	}

	if hm.Size() != 10 {
		t.Errorf("expected size 10, got %d", hm.Size())
	}
}

func TestHashMap_MixedTypes(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("int", 42)
	hm.Put("string", "hello")
	hm.Put("bool", true)
	hm.Put("float", 3.14)

	val, _ := hm.Get("int")
	if val != 42 {
		t.Errorf("expected 42, got %v", val)
	}

	val, _ = hm.Get("string")
	if val != "hello" {
		t.Errorf("expected 'hello', got %v", val)
	}

	val, _ = hm.Get("bool")
	if val != true {
		t.Errorf("expected true, got %v", val)
	}

	val, _ = hm.Get("float")
	if val != 3.14 {
		t.Errorf("expected 3.14, got %v", val)
	}
}

func TestHashMap_EmptyKey(t *testing.T) {
	hm := NewHashMap(4)

	hm.Put("", "empty key value")

	val, ok := hm.Get("")
	if !ok || val != "empty key value" {
		t.Errorf("expected 'empty key value', got %v", val)
	}
}

func TestHashMap_DeleteFromChain(t *testing.T) {
	hm := NewHashMap(2)

	// Force collisions
	hm.Put("a", 1)
	hm.Put("b", 2)
	hm.Put("c", 3)

	// Delete middle of chain (if they collide)
	hm.Delete("b")

	// Verify a and c still accessible
	if !hm.Contains("a") {
		t.Error("key 'a' should still exist")
	}
	if !hm.Contains("c") {
		t.Error("key 'c' should still exist")
	}
	if hm.Contains("b") {
		t.Error("key 'b' should be deleted")
	}
}
