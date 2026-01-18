package basics

import "testing"

func TestMap_NilVsEmpty(t *testing.T) {
	nilMap := CreateNilMap()
	emptyMap := CreateEmptyMap()

	// Nil map is nil
	if nilMap != nil {
		t.Error("expected nil map to be nil")
	}

	// Empty map is not nil
	if emptyMap == nil {
		t.Error("expected empty map to not be nil")
	}

	// Both have length 0
	if len(nilMap) != 0 {
		t.Errorf("expected nil map length 0, got %d", len(nilMap))
	}
	if len(emptyMap) != 0 {
		t.Errorf("expected empty map length 0, got %d", len(emptyMap))
	}
}

func TestMap_ReadFromNil(t *testing.T) {
	nilMap := CreateNilMap()

	// Reading from nil map is safe, returns zero value
	value := nilMap["key"]
	if value != 0 {
		t.Errorf("expected zero value 0, got %d", value)
	}

	// Two-value read is also safe
	value, exists := SafeMapRead(nilMap, "key")
	if exists {
		t.Error("expected key to not exist in nil map")
	}
	if value != 0 {
		t.Errorf("expected zero value 0, got %d", value)
	}
}

func TestMap_WriteToEmpty(t *testing.T) {
	m := CreateEmptyMap()

	// Writing to empty map is safe
	m["key"] = 42

	if m["key"] != 42 {
		t.Errorf("expected m[key] to be 42, got %d", m["key"])
	}
}

func TestMap_Literal(t *testing.T) {
	m := CreateMapWithLiteral()

	if len(m) != 2 {
		t.Errorf("expected map length 2, got %d", len(m))
	}

	if m["a"] != 1 {
		t.Errorf("expected m[a] to be 1, got %d", m["a"])
	}

	if m["b"] != 2 {
		t.Errorf("expected m[b] to be 2, got %d", m["b"])
	}
}

func TestMap_SafeRead(t *testing.T) {
	m := map[string]int{"key": 42}

	// Key exists
	value, exists := SafeMapRead(m, "key")
	if !exists {
		t.Error("expected key to exist")
	}
	if value != 42 {
		t.Errorf("expected value 42, got %d", value)
	}

	// Key doesn't exist
	value, exists = SafeMapRead(m, "nonexistent")
	if exists {
		t.Error("expected key to not exist")
	}
	if value != 0 {
		t.Errorf("expected zero value 0, got %d", value)
	}
}

func TestMap_Delete(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	// Delete existing key
	DeleteFromMap(m, "a")
	if _, exists := m["a"]; exists {
		t.Error("expected key 'a' to be deleted")
	}

	// Delete non-existent key (safe, no panic)
	DeleteFromMap(m, "nonexistent")

	// Map should still have "b"
	if m["b"] != 2 {
		t.Errorf("expected m[b] to be 2, got %d", m["b"])
	}
}

func TestMap_Length(t *testing.T) {
	// Nil map
	var nilMap map[string]int
	if MapLength(nilMap) != 0 {
		t.Errorf("expected nil map length 0, got %d", MapLength(nilMap))
	}

	// Empty map
	emptyMap := make(map[string]int)
	if MapLength(emptyMap) != 0 {
		t.Errorf("expected empty map length 0, got %d", MapLength(emptyMap))
	}

	// Map with elements
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	if MapLength(m) != 3 {
		t.Errorf("expected map length 3, got %d", MapLength(m))
	}
}

func TestMap_Iteration(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	keys := IterateMap(m)

	// Should have all keys (order not guaranteed)
	if len(keys) != 3 {
		t.Errorf("expected 3 keys, got %d", len(keys))
	}

	// Verify all keys are present
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}

	if !keyMap["a"] || !keyMap["b"] || !keyMap["c"] {
		t.Error("expected all keys to be present")
	}
}

func TestMap_ZeroValue(t *testing.T) {
	m := map[string]int{"a": 1}

	// Existing key
	if MapZeroValue(m, "a") != 1 {
		t.Errorf("expected value 1, got %d", MapZeroValue(m, "a"))
	}

	// Non-existent key returns zero value
	if MapZeroValue(m, "nonexistent") != 0 {
		t.Errorf("expected zero value 0, got %d", MapZeroValue(m, "nonexistent"))
	}
}

func TestMap_CheckKeyExists(t *testing.T) {
	m := map[string]int{"a": 1, "b": 0} // Note: "b" has zero value

	// Key exists with non-zero value
	if !CheckKeyExists(m, "a") {
		t.Error("expected key 'a' to exist")
	}

	// Key exists with zero value (important distinction!)
	if !CheckKeyExists(m, "b") {
		t.Error("expected key 'b' to exist")
	}

	// Key doesn't exist
	if CheckKeyExists(m, "c") {
		t.Error("expected key 'c' to not exist")
	}
}

func TestMap_ReferenceSemantics(t *testing.T) {
	m1, m2 := MapAsReference()

	// Both should show the modification
	if m1["a"] != 999 {
		t.Errorf("expected m1[a] to be 999, got %d", m1["a"])
	}
	if m2["a"] != 999 {
		t.Errorf("expected m2[a] to be 999, got %d", m2["a"])
	}
}

func TestMap_Copy(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2}
	copy := CopyMap(original)

	// Modify copy
	copy["a"] = 999

	// Original should be unchanged
	if original["a"] != 1 {
		t.Errorf("expected original[a] to be 1, got %d", original["a"])
	}

	// Copy should be modified
	if copy["a"] != 999 {
		t.Errorf("expected copy[a] to be 999, got %d", copy["a"])
	}
}

func TestMap_DeleteFromNil(t *testing.T) {
	var m map[string]int // nil map

	// Deleting from nil map is safe (no panic)
	delete(m, "key")

	// Map is still nil
	if m != nil {
		t.Error("expected map to still be nil")
	}
}
