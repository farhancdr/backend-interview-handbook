package algo

import "testing"

func TestBinarySearch_Found(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 7

	result := BinarySearch(arr, target)
	if result != 3 {
		t.Errorf("expected index 3, got %d", result)
	}
}

func TestBinarySearch_NotFound(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 6

	result := BinarySearch(arr, target)
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
}

func TestBinarySearch_EmptyArray(t *testing.T) {
	arr := []int{}
	target := 5

	result := BinarySearch(arr, target)
	if result != -1 {
		t.Errorf("expected -1, got %d", result)
	}
}

func TestBinarySearch_SingleElement(t *testing.T) {
	arr := []int{5}

	// Found
	if BinarySearch(arr, 5) != 0 {
		t.Error("expected to find element at index 0")
	}

	// Not found
	if BinarySearch(arr, 3) != -1 {
		t.Error("expected -1 for not found")
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13}

	tests := []struct {
		target   int
		expected int
	}{
		{7, 3},
		{1, 0},
		{13, 6},
		{6, -1},
	}

	for _, tt := range tests {
		result := BinarySearchRecursive(arr, tt.target)
		if result != tt.expected {
			t.Errorf("target %d: expected %d, got %d", tt.target, tt.expected, result)
		}
	}
}

func TestFindFirstOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 4, 5}
	target := 2

	result := FindFirstOccurrence(arr, target)
	if result != 1 {
		t.Errorf("expected first occurrence at index 1, got %d", result)
	}
}

func TestFindLastOccurrence(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 4, 5}
	target := 2

	result := FindLastOccurrence(arr, target)
	if result != 3 {
		t.Errorf("expected last occurrence at index 3, got %d", result)
	}
}

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}

	for _, tt := range tests {
		result := SearchInsertPosition(tt.arr, tt.target)
		if result != tt.expected {
			t.Errorf("target %d: expected %d, got %d", tt.target, tt.expected, result)
		}
	}
}

func TestSearchRotatedArray(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}

	tests := []struct {
		target   int
		expected int
	}{
		{0, 4},
		{4, 0},
		{7, 3},
		{3, -1},
	}

	for _, tt := range tests {
		result := SearchRotatedArray(arr, tt.target)
		if result != tt.expected {
			t.Errorf("target %d: expected %d, got %d", tt.target, tt.expected, result)
		}
	}
}

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		arr []int
	}{
		{[]int{1, 2, 3, 1}},
		{[]int{1, 2, 1, 3, 5, 6, 4}},
	}

	for _, tt := range tests {
		peak := FindPeakElement(tt.arr)

		// Verify it's a peak
		if peak > 0 && tt.arr[peak] <= tt.arr[peak-1] {
			t.Errorf("not a peak: arr[%d]=%d <= arr[%d]=%d", peak, tt.arr[peak], peak-1, tt.arr[peak-1])
		}
		if peak < len(tt.arr)-1 && tt.arr[peak] <= tt.arr[peak+1] {
			t.Errorf("not a peak: arr[%d]=%d <= arr[%d]=%d", peak, tt.arr[peak], peak+1, tt.arr[peak+1])
		}
	}
}

func TestSquareRoot(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{4, 2},
		{8, 2},
		{16, 4},
		{1, 1},
		{0, 0},
		{10, 3},
	}

	for _, tt := range tests {
		result := SquareRoot(tt.x)
		if result != tt.expected {
			t.Errorf("sqrt(%d): expected %d, got %d", tt.x, tt.expected, result)
		}
	}
}

func TestBinarySearch_EdgeCases(t *testing.T) {
	// First element
	arr := []int{1, 3, 5, 7, 9}
	if BinarySearch(arr, 1) != 0 {
		t.Error("failed to find first element")
	}

	// Last element
	if BinarySearch(arr, 9) != 4 {
		t.Error("failed to find last element")
	}

	// Middle element
	if BinarySearch(arr, 5) != 2 {
		t.Error("failed to find middle element")
	}
}
