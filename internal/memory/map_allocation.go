package memory

import (
	"strconv"
)

// Why interviewers ask this:
// Maps are heavily used in Go applications, and understanding their allocation and growth
// behavior is important for performance. Interviewers want to see if you know when to
// pre-allocate maps, how they grow, and the memory implications of map operations.

// Common pitfalls:
// - Not pre-allocating maps when size is known
// - Assuming maps shrink when elements are deleted (they don't)
// - Not understanding that nil maps can't be written to
// - Forgetting that map iteration order is random
// - Not knowing that maps are reference types

// Key takeaway:
// Maps are reference types. Nil maps can be read but not written. Use make(map[K]V, size)
// to pre-allocate when size is known. Maps grow but never shrink - deleted keys free values
// but not buckets. For large maps that shrink significantly, create a new map and copy.

// NilMapBehavior demonstrates nil map characteristics
func NilMapBehavior() (canRead, canWrite bool, readValue int) {
	var m map[string]int

	// Reading from nil map is safe (returns zero value)
	canRead = true
	readValue = m["key"]

	// Writing to nil map causes panic
	canWrite = false
	// m["key"] = 42 // This would panic!

	return canRead, canWrite, readValue
}

// MapPreallocation shows map initialization patterns
func MapPreallocation(size int) (withSize, withoutSize map[string]int) {
	// Without size hint
	withoutSize = make(map[string]int)

	// With size hint (pre-allocates buckets)
	withSize = make(map[string]int, size)

	return withSize, withoutSize
}

// MapGrowthPattern demonstrates how maps grow
func MapGrowthPattern(n int) map[string]int {
	m := make(map[string]int)

	for i := 0; i < n; i++ {
		m[string(rune('a'+i%26))+strconv.Itoa(i)] = i
	}

	return m
}

// MapDoesNotShrink shows that maps don't release memory when elements are deleted
func MapDoesNotShrink() (before, after int) {
	m := make(map[int]int)

	// Add many elements
	for i := 0; i < 10000; i++ {
		m[i] = i
	}

	before = len(m)

	// Delete most elements
	for i := 0; i < 9000; i++ {
		delete(m, i)
	}

	after = len(m)
	// Note: The underlying buckets are not freed, only the values

	return before, after
}

// MapIsReferenceType shows maps are passed by reference
func MapIsReferenceType() (original, modified map[string]int, same bool) {
	original = map[string]int{"a": 1, "b": 2}

	// Pass to function that modifies it
	modifyMap(original)

	// Check if original was modified
	modified = original
	same = (original["a"] == 99)

	return original, modified, same
}

func modifyMap(m map[string]int) {
	m["a"] = 99
}

// MapLiteralVsMake compares initialization methods
func MapLiteralVsMake() (literal, withMake map[string]int) {
	// Map literal
	literal = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// Using make
	withMake = make(map[string]int)
	withMake["one"] = 1
	withMake["two"] = 2
	withMake["three"] = 3

	return literal, withMake
}

// CheckKeyExists demonstrates the comma-ok idiom
func CheckKeyExists(m map[string]int, key string) (value int, exists bool) {
	value, exists = m[key]
	return value, exists
}

// SafeMapAccess shows safe ways to access maps
func SafeMapAccess(m map[string]int, key string) int {
	if m == nil {
		return 0
	}

	if val, ok := m[key]; ok {
		return val
	}

	return -1 // Default value when key doesn't exist
}

// MapKeysAreUnordered demonstrates random iteration order
func MapKeysAreUnordered() []string {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// CopyMap creates a new map with same contents
// This is useful when you want to shrink a large map
func CopyMap(original map[string]int) map[string]int {
	copied := make(map[string]int, len(original))

	for k, v := range original {
		copied[k] = v
	}

	return copied
}

// DeleteAllKeys removes all keys from a map
// Note: This doesn't free the underlying buckets
func DeleteAllKeys(m map[string]int) {
	for k := range m {
		delete(m, k)
	}
}

// ReplaceMapToShrink shows how to reclaim memory from a large map
func ReplaceMapToShrink(m map[string]int, keysToKeep []string) map[string]int {
	// Create new map with only needed elements
	newMap := make(map[string]int, len(keysToKeep))

	for _, key := range keysToKeep {
		if val, ok := m[key]; ok {
			newMap[key] = val
		}
	}

	return newMap
}
