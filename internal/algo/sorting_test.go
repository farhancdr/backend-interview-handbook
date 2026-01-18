package algo

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	QuickSort(arr)

	if !IsSorted(arr) {
		t.Errorf("array not sorted: %v", arr)
	}
}

func TestQuickSort_Empty(t *testing.T) {
	arr := []int{}
	QuickSort(arr)

	if len(arr) != 0 {
		t.Error("empty array should remain empty")
	}
}

func TestQuickSort_SingleElement(t *testing.T) {
	arr := []int{42}
	QuickSort(arr)

	if arr[0] != 42 {
		t.Errorf("expected [42], got %v", arr)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := MergeSort(arr)

	expected := []int{11, 12, 22, 25, 34, 64, 90}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("expected %v, got %v", expected, sorted)
	}
}

func TestMergeSort_AlreadySorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	sorted := MergeSort(arr)

	if !reflect.DeepEqual(sorted, arr) {
		t.Errorf("expected %v, got %v", arr, sorted)
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	BubbleSort(arr)

	if !IsSorted(arr) {
		t.Errorf("array not sorted: %v", arr)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	InsertionSort(arr)

	if !IsSorted(arr) {
		t.Errorf("array not sorted: %v", arr)
	}
}

func TestInsertionSort_NearlySorted(t *testing.T) {
	// Insertion sort is efficient for nearly sorted arrays
	arr := []int{1, 2, 3, 5, 4, 6, 7}
	InsertionSort(arr)

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	SelectionSort(arr)

	if !IsSorted(arr) {
		t.Errorf("array not sorted: %v", arr)
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	HeapSort(arr)

	if !IsSorted(arr) {
		t.Errorf("array not sorted: %v", arr)
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		arr      []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 3, 2, 4}, false},
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 1, 1}, true},
	}

	for _, tt := range tests {
		result := IsSorted(tt.arr)
		if result != tt.expected {
			t.Errorf("IsSorted(%v): expected %v, got %v", tt.arr, tt.expected, result)
		}
	}
}

func TestKthLargest(t *testing.T) {
	arr := []int{3, 2, 1, 5, 6, 4}

	tests := []struct {
		k        int
		expected int
	}{
		{1, 6}, // 1st largest
		{2, 5}, // 2nd largest
		{3, 4}, // 3rd largest
	}

	for _, tt := range tests {
		// Make a copy since quickSelect modifies array
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)

		result := KthLargest(arrCopy, tt.k)
		if result != tt.expected {
			t.Errorf("KthLargest(k=%d): expected %d, got %d", tt.k, tt.expected, result)
		}
	}
}

func TestSorting_Duplicates(t *testing.T) {
	arr := []int{5, 2, 3, 2, 1, 5, 3}

	// Test each sorting algorithm with duplicates
	testCases := []struct {
		name string
		fn   func([]int)
	}{
		{"QuickSort", QuickSort},
		{"BubbleSort", BubbleSort},
		{"InsertionSort", InsertionSort},
		{"SelectionSort", SelectionSort},
		{"HeapSort", HeapSort},
	}

	for _, tc := range testCases {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)

		tc.fn(arrCopy)

		if !IsSorted(arrCopy) {
			t.Errorf("%s failed with duplicates: %v", tc.name, arrCopy)
		}
	}
}

func TestSorting_NegativeNumbers(t *testing.T) {
	arr := []int{-5, 2, -3, 0, 1, -1}

	QuickSort(arr)

	expected := []int{-5, -3, -1, 0, 1, 2}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestMergeSort_Stability(t *testing.T) {
	// MergeSort should be stable (maintain relative order of equal elements)
	// This is harder to test with just integers, but we can verify it sorts correctly
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sorted := MergeSort(arr)

	if !IsSorted(sorted) {
		t.Errorf("MergeSort failed: %v", sorted)
	}
}
