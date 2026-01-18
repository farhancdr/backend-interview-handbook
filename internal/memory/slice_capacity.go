package memory

// Why interviewers ask this:
// Understanding slice capacity growth is crucial for writing performant Go code. Interviewers
// want to see if you understand the difference between length and capacity, how slices grow
// when appending, and when to pre-allocate. This directly impacts application performance and
// memory usage patterns.

// Common pitfalls:
// - Not understanding the difference between len() and cap()
// - Repeatedly appending to slices without pre-allocation (causes multiple reallocations)
// - Over-allocating capacity (wasting memory)
// - Not knowing that capacity typically doubles when growing
// - Assuming append always modifies the original slice

// Key takeaway:
// Slices have both length (current elements) and capacity (allocated space). When appending
// beyond capacity, Go allocates a new array (typically 2x size), copies elements, and updates
// the slice. Pre-allocate with make([]T, 0, capacity) when final size is known to avoid
// multiple allocations and improve performance.

import "fmt"

// SliceGrowthPattern demonstrates how slice capacity grows
// Time Complexity: O(n) for n appends, but with reallocations
func SliceGrowthPattern(n int) []int {
	var s []int
	capacities := []int{}

	for i := 0; i < n; i++ {
		prevCap := cap(s)
		s = append(s, i)
		newCap := cap(s)

		if newCap != prevCap {
			capacities = append(capacities, newCap)
		}
	}

	return capacities
}

// AppendWithoutPreallocation shows inefficient pattern
// Multiple reallocations occur as slice grows
func AppendWithoutPreallocation(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

// AppendWithPreallocation shows efficient pattern
// Single allocation, no reallocations needed
func AppendWithPreallocation(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

// AppendWithLengthPreallocation uses make with length
// This pre-fills with zero values
func AppendWithLengthPreallocation(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

// SliceCapacityInfo returns length and capacity information
func SliceCapacityInfo(s []int) (length, capacity int) {
	return len(s), cap(s)
}

// DemonstrateSliceSharing shows how slices can share underlying arrays
func DemonstrateSliceSharing() (original, slice1, slice2 []int, modified bool) {
	original = []int{1, 2, 3, 4, 5}
	slice1 = original[:3] // [1, 2, 3] but shares array
	slice2 = original[2:] // [3, 4, 5] but shares array

	// Modify through slice1
	slice1[2] = 99

	// Check if original and slice2 are affected
	modified = (original[2] == 99 && slice2[0] == 99)

	return original, slice1, slice2, modified
}

// AppendCausingReallocation shows when append creates new array
func AppendCausingReallocation() (original, appended []int, sameArray bool) {
	original = make([]int, 3) // len=3, cap=3
	original[0], original[1], original[2] = 1, 2, 3

	// This will cause reallocation since cap is full
	appended = append(original, 4)

	// Modify appended slice
	appended[0] = 99

	// Check if they share the same underlying array
	sameArray = (original[0] == 99)

	return original, appended, sameArray
}

// CopyVsAppend demonstrates difference between copy and append
func CopyVsAppend(src []int) (copied, appended []int) {
	// Using copy - creates independent slice
	copied = make([]int, len(src))
	copy(copied, src)

	// Using append - may share array if capacity allows
	appended = append([]int{}, src...)

	return copied, appended
}

// SliceCapacityAfterSlicing shows capacity behavior after slicing
func SliceCapacityAfterSlicing() map[string]int {
	original := make([]int, 5, 10)

	results := make(map[string]int)
	results["original_len"] = len(original)
	results["original_cap"] = cap(original)

	sliced := original[1:3]
	results["sliced_len"] = len(sliced)
	results["sliced_cap"] = cap(sliced) // Capacity from sliced position to end

	return results
}

// FullSliceExpression demonstrates three-index slicing
// s[low:high:max] sets capacity to max-low
func FullSliceExpression() (normal, limited []int, normalCap, limitedCap int) {
	original := make([]int, 5, 10)
	for i := range original {
		original[i] = i
	}

	// Normal slicing: s[1:3]
	normal = original[1:3]
	normalCap = cap(normal) // Will be 9 (from index 1 to end of capacity)

	// Full slice expression: s[1:3:3]
	limited = original[1:3:3]
	limitedCap = cap(limited) // Will be 2 (3-1)

	return normal, limited, normalCap, limitedCap
}

// NilVsEmptySlice demonstrates the difference
func NilVsEmptySlice() (nilSlice, emptySlice []int, nilIsNil, emptyIsNil bool) {
	emptySlice = []int{}

	nilIsNil = (nilSlice == nil)
	emptyIsNil = (emptySlice == nil)

	return nilSlice, emptySlice, nilIsNil, emptyIsNil
}

// PrintSliceInfo prints detailed slice information (for debugging)
func PrintSliceInfo(name string, s []int) string {
	return fmt.Sprintf("%s: len=%d cap=%d values=%v", name, len(s), cap(s), s)
}
