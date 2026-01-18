package basics

// Why interviewers ask this:
// Maps are heavily used in Go, and understanding nil maps vs empty maps is crucial
// to avoid runtime panics. Interviewers want to ensure you know safe map operations
// and the difference between nil and initialized maps.

// Common pitfalls:
// - Writing to a nil map causes panic
// - Not checking if a key exists before reading
// - Assuming maps are ordered (they're not)
// - Not understanding that maps are reference types
// - Forgetting that reading from a nil map returns zero value (doesn't panic)

// Key takeaway:
// Nil map: var m map[string]int (reading is safe, writing panics)
// Empty map: m := make(map[string]int) or m := map[string]int{} (safe for all operations)

// CreateNilMap returns a nil map
func CreateNilMap() map[string]int {
	var m map[string]int // nil map
	return m
}

// CreateEmptyMap returns an empty (but not nil) map
func CreateEmptyMap() map[string]int {
	m := make(map[string]int) // empty map, not nil
	return m
}

// CreateMapWithLiteral returns a map created with literal syntax
func CreateMapWithLiteral() map[string]int {
	m := map[string]int{"a": 1, "b": 2}
	return m
}

// SafeMapRead demonstrates safe reading from a map
func SafeMapRead(m map[string]int, key string) (value int, exists bool) {
	value, exists = m[key] // Two-value assignment
	return value, exists
}

// UnsafeMapWrite demonstrates that writing to nil map panics
func UnsafeMapWrite() {
	var m map[string]int // nil map
	// This would panic: m["key"] = 42
	_ = m
}

// SafeMapWrite demonstrates safe map writing
func SafeMapWrite() map[string]int {
	m := make(map[string]int) // Initialize first
	m["key"] = 42             // Safe to write
	return m
}

// DeleteFromMap demonstrates map deletion
func DeleteFromMap(m map[string]int, key string) {
	delete(m, key) // Safe even if key doesn't exist
}

// MapLength returns the number of key-value pairs
func MapLength(m map[string]int) int {
	return len(m) // Safe even for nil map (returns 0)
}

// IterateMap demonstrates map iteration
func IterateMap(m map[string]int) []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys // Order is not guaranteed!
}

// MapZeroValue demonstrates zero value behavior
func MapZeroValue(m map[string]int, key string) int {
	return m[key] // Returns 0 if key doesn't exist (for int)
}

// CheckKeyExists demonstrates the idiomatic way to check key existence
func CheckKeyExists(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

// MapAsReference demonstrates that maps are reference types
func MapAsReference() (map[string]int, map[string]int) {
	m1 := map[string]int{"a": 1}
	m2 := m1 // Both point to same map
	m2["a"] = 999
	return m1, m2 // Both show {"a": 999}
}

// CopyMap demonstrates how to create an independent copy
func CopyMap(original map[string]int) map[string]int {
	copy := make(map[string]int, len(original))
	for key, value := range original {
		copy[key] = value
	}
	return copy
}
