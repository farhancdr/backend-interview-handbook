package memory

import (
	"testing"
)

func TestNilMapBehavior(t *testing.T) {
	canRead, canWrite, readValue := NilMapBehavior()

	if !canRead {
		t.Error("should be able to read from nil map")
	}

	if canWrite {
		t.Error("should not be able to write to nil map")
	}

	if readValue != 0 {
		t.Errorf("expected zero value 0, got %d", readValue)
	}
}

func TestMapPreallocation(t *testing.T) {
	size := 100
	withSize, withoutSize := MapPreallocation(size)

	if withSize == nil {
		t.Error("withSize map should not be nil")
	}

	if withoutSize == nil {
		t.Error("withoutSize map should not be nil")
	}

	// Both should be writable
	withSize["test"] = 1
	withoutSize["test"] = 1

	if withSize["test"] != 1 || withoutSize["test"] != 1 {
		t.Error("both maps should be writable")
	}
}

func TestMapGrowthPattern(t *testing.T) {
	n := 100
	m := MapGrowthPattern(n)

	if len(m) != n {
		t.Errorf("expected map length %d, got %d", n, len(m))
	}
}

func TestMapDoesNotShrink(t *testing.T) {
	before, after := MapDoesNotShrink()

	if before != 10000 {
		t.Errorf("expected before = 10000, got %d", before)
	}

	if after != 1000 {
		t.Errorf("expected after = 1000, got %d", after)
	}

	// The key point: underlying memory is not reclaimed
	// (we can't directly test this without runtime internals)
}

func TestMapIsReferenceType(t *testing.T) {
	original, modified, same := MapIsReferenceType()

	if !same {
		t.Error("map should be modified by reference")
	}

	if original["a"] != 99 {
		t.Errorf("expected original['a'] = 99, got %d", original["a"])
	}

	if modified["a"] != 99 {
		t.Errorf("expected modified['a'] = 99, got %d", modified["a"])
	}
}

func TestMapLiteralVsMake(t *testing.T) {
	literal, withMake := MapLiteralVsMake()

	if len(literal) != 3 {
		t.Errorf("expected literal length 3, got %d", len(literal))
	}

	if len(withMake) != 3 {
		t.Errorf("expected withMake length 3, got %d", len(withMake))
	}

	// Check values
	if literal["one"] != 1 || withMake["one"] != 1 {
		t.Error("values should match")
	}
}

func TestCheckKeyExists(t *testing.T) {
	m := map[string]int{"exists": 42}

	value, exists := CheckKeyExists(m, "exists")
	if !exists || value != 42 {
		t.Errorf("expected exists=true, value=42, got exists=%v, value=%d", exists, value)
	}

	value, exists = CheckKeyExists(m, "notexists")
	if exists {
		t.Error("key should not exist")
	}
	if value != 0 {
		t.Errorf("expected zero value 0, got %d", value)
	}
}

func TestSafeMapAccess(t *testing.T) {
	m := map[string]int{"key": 42}

	value := SafeMapAccess(m, "key")
	if value != 42 {
		t.Errorf("expected 42, got %d", value)
	}

	value = SafeMapAccess(m, "notexists")
	if value != -1 {
		t.Errorf("expected -1 for missing key, got %d", value)
	}

	value = SafeMapAccess(nil, "key")
	if value != 0 {
		t.Errorf("expected 0 for nil map, got %d", value)
	}
}

func TestMapKeysAreUnordered(t *testing.T) {
	keys := MapKeysAreUnordered()

	if len(keys) != 5 {
		t.Errorf("expected 5 keys, got %d", len(keys))
	}

	// Verify all expected keys are present
	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	expected := []string{"a", "b", "c", "d", "e"}
	for _, k := range expected {
		if !keyMap[k] {
			t.Errorf("missing key %s", k)
		}
	}
}

func TestCopyMap(t *testing.T) {
	original := map[string]int{"a": 1, "b": 2, "c": 3}

	copied := CopyMap(original)

	if len(copied) != len(original) {
		t.Errorf("expected same length, got %d vs %d", len(copied), len(original))
	}

	// Modify copied
	copied["a"] = 99

	// Original should be unchanged
	if original["a"] != 1 {
		t.Error("original should not be modified")
	}

	if copied["a"] != 99 {
		t.Error("copied should be modified")
	}
}

func TestDeleteAllKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	DeleteAllKeys(m)

	if len(m) != 0 {
		t.Errorf("expected empty map, got length %d", len(m))
	}
}

func TestReplaceMapToShrink(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	keysToKeep := []string{"a", "c"}
	newMap := ReplaceMapToShrink(m, keysToKeep)

	if len(newMap) != 2 {
		t.Errorf("expected length 2, got %d", len(newMap))
	}

	if newMap["a"] != 1 || newMap["c"] != 3 {
		t.Error("expected keys a and c with correct values")
	}

	if _, ok := newMap["b"]; ok {
		t.Error("key b should not exist in new map")
	}
}

// Benchmarks

func BenchmarkMapWithoutPreallocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for j := 0; j < 1000; j++ {
			m[j] = j
		}
	}
}

func BenchmarkMapWithPreallocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, 1000)
		for j := 0; j < 1000; j++ {
			m[j] = j
		}
	}
}

func BenchmarkMapLiteral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
	}
}

func BenchmarkMapMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[string]int)
		m["a"] = 1
		m["b"] = 2
		m["c"] = 3
	}
}

func BenchmarkMapCopy(b *testing.B) {
	original := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		original[string(rune(i))] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyMap(original)
	}
}
