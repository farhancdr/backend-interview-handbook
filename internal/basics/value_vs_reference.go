package basics

// Why interviewers ask this:
// Understanding value vs reference types is fundamental to Go. Interviewers want
// to ensure you know when data is copied vs when it's shared, which affects
// function behavior, memory usage, and potential bugs.

// Common pitfalls:
// - Assuming all types are references (like in Java/Python)
// - Not understanding that slices/maps/channels are reference types but structs are values
// - Forgetting that passing a struct to a function creates a copy
// - Confusion about pointer receivers vs value receivers

// Key takeaway:
// Value types (int, float, bool, string, array, struct) are copied on assignment.
// Reference types (slice, map, channel, pointer) share the underlying data.

// ValueType demonstrates that basic types and structs are copied
type ValueType struct {
	Value int
}

// ModifyValue attempts to modify a value type (won't affect original)
func ModifyValue(v ValueType) {
	v.Value = 100
}

// ModifyValuePointer modifies via pointer (affects original)
func ModifyValuePointer(v *ValueType) {
	v.Value = 100
}

// ReferenceType demonstrates reference behavior with slices
type ReferenceType struct {
	Data []int
}

// ModifySlice modifies the underlying array (affects original)
func ModifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
}

// ModifyMap modifies the map (affects original)
func ModifyMap(m map[string]int) {
	m["key"] = 999
}

// DemonstrateValueCopy shows that structs are copied
func DemonstrateValueCopy() (original, modified ValueType) {
	original = ValueType{Value: 42}
	modified = original // This creates a copy
	modified.Value = 100
	return original, modified // original.Value is still 42
}

// DemonstrateSliceReference shows that slices share underlying data
func DemonstrateSliceReference() ([]int, []int) {
	original := []int{1, 2, 3}
	reference := original // Both point to same underlying array
	reference[0] = 999
	return original, reference // Both show [999, 2, 3]
}

// DemonstrateMapReference shows that maps are references
func DemonstrateMapReference() (map[string]int, map[string]int) {
	original := map[string]int{"a": 1}
	reference := original // Both point to same map
	reference["a"] = 999
	return original, reference // Both show {"a": 999}
}
