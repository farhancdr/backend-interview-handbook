package basics

// Why interviewers ask this:
// Arrays and slices are often confused. Interviewers want to ensure you understand
// the fundamental differences: arrays are fixed-size value types, slices are
// dynamic-size reference types. This affects performance, memory, and API design.

// Common pitfalls:
// - Thinking arrays and slices are interchangeable
// - Not understanding that slices are references to underlying arrays
// - Passing large arrays by value (expensive copy)
// - Confusion about slice capacity vs length
// - Modifying a slice affects other slices sharing the same backing array

// Key takeaway:
// Arrays: fixed size, value type, rarely used directly.
// Slices: dynamic size, reference to array, preferred in Go.

// ArrayExample demonstrates array behavior (fixed size, value type)
func ArrayExample() [3]int {
	var arr [3]int // Array of 3 integers, initialized to [0, 0, 0]
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	return arr
}

// SliceExample demonstrates slice behavior (dynamic size, reference type)
func SliceExample() []int {
	var s []int // Slice, initially nil
	s = append(s, 1, 2, 3)
	return s
}

// ArrayPassedByValue demonstrates that arrays are copied when passed to functions
func ArrayPassedByValue(arr [3]int) {
	arr[0] = 999 // Modifies the copy, not the original
}

// SlicePassedByReference demonstrates that slices share underlying data
func SlicePassedByReference(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Modifies the original underlying array
	}
}

// SliceCapacityExample demonstrates slice length vs capacity
func SliceCapacityExample() (length, capacity int) {
	s := make([]int, 3, 5) // length=3, capacity=5
	return len(s), cap(s)
}

// SliceAppendExample demonstrates append behavior and capacity growth
func SliceAppendExample() (original, appended []int) {
	original = []int{1, 2, 3}
	appended = append(original, 4) // May or may not share backing array
	return original, appended
}

// SliceResliceExample demonstrates slicing creates a view of the same array
func SliceResliceExample() (original, subslice []int) {
	original = []int{1, 2, 3, 4, 5}
	subslice = original[1:4] // [2, 3, 4] - shares backing array
	subslice[0] = 999        // Modifies original[1]
	return original, subslice
}

// ArrayToSlice converts an array to a slice
func ArrayToSlice(arr [3]int) []int {
	return arr[:] // Creates a slice from the array
}

// CompareArrays compares two arrays for equality
func CompareArrays(a, b [3]int) bool {
	return a == b // Arrays can be compared with ==
}

// CompareSlices demonstrates that slices cannot be compared with ==
func CompareSlices(a, b []int) bool {
	// This won't compile: return a == b
	// Must compare element by element
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// NilSliceVsEmptySlice demonstrates the difference
func NilSliceVsEmptySlice() (nilSlice, emptySlice []int, nilIsNil, emptyIsNil bool) {
	var nilSlice2 []int    // nil slice
	emptySlice2 := []int{} // empty slice (not nil)
	return nilSlice2, emptySlice2, nilSlice2 == nil, emptySlice2 == nil
}
