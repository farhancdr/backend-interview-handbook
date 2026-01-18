package algo

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(arr, target)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("expected [0, 1], got %v", result)
	}
}

func TestTwoSum_NotFound(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	target := 100

	result := TwoSum(arr, target)
	if !reflect.DeepEqual(result, []int{-1, -1}) {
		t.Errorf("expected [-1, -1], got %v", result)
	}
}

func TestThreeSum(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSum(arr)

	// Should find: [-1, -1, 2] and [-1, 0, 1]
	if len(result) != 2 {
		t.Errorf("expected 2 triplets, got %d", len(result))
	}
}

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{1, 1, 2, 2, 3, 4, 4, 5}
	newLen := RemoveDuplicates(arr)

	if newLen != 5 {
		t.Errorf("expected new length 5, got %d", newLen)
	}

	// Verify unique elements
	expected := []int{1, 2, 3, 4, 5}
	for i := 0; i < newLen; i++ {
		if arr[i] != expected[i] {
			t.Errorf("at index %d: expected %d, got %d", i, expected[i], arr[i])
		}
	}
}

func TestRemoveDuplicates_Empty(t *testing.T) {
	arr := []int{}
	newLen := RemoveDuplicates(arr)

	if newLen != 0 {
		t.Errorf("expected length 0, got %d", newLen)
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"a", true},
		{"", true},
		{"ab", false},
		{"aa", true},
	}

	for _, tt := range tests {
		result := IsPalindrome(tt.s)
		if result != tt.expected {
			t.Errorf("IsPalindrome(%s): expected %v, got %v", tt.s, tt.expected, result)
		}
	}
}

func TestReverseString(t *testing.T) {
	s := []byte("hello")
	ReverseString(s)

	expected := "olleh"
	if string(s) != expected {
		t.Errorf("expected %s, got %s", expected, string(s))
	}
}

func TestMoveZeroes(t *testing.T) {
	arr := []int{0, 1, 0, 3, 12}
	MoveZeroes(arr)

	expected := []int{1, 3, 12, 0, 0}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestMoveZeroes_AllZeros(t *testing.T) {
	arr := []int{0, 0, 0}
	MoveZeroes(arr)

	expected := []int{0, 0, 0}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestMoveZeroes_NoZeros(t *testing.T) {
	arr := []int{1, 2, 3}
	MoveZeroes(arr)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v, got %v", expected, arr)
	}
}

func TestContainerWithMostWater(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := ContainerWithMostWater(height)

	// Max area is 49 (between index 1 and 8)
	if result != 49 {
		t.Errorf("expected 49, got %d", result)
	}
}

func TestPartitionArray(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	pivot := 5

	partitionIndex := PartitionArray(arr, pivot)

	// Verify all elements before partition are < pivot
	for i := 0; i < partitionIndex; i++ {
		if arr[i] >= pivot {
			t.Errorf("element %d at index %d should be < %d", arr[i], i, pivot)
		}
	}

	// Verify all elements from partition are >= pivot
	for i := partitionIndex; i < len(arr); i++ {
		if arr[i] < pivot {
			t.Errorf("element %d at index %d should be >= %d", arr[i], i, pivot)
		}
	}
}

func TestTwoSum_MultipleValidPairs(t *testing.T) {
	// When there are multiple valid pairs, should return first one found
	arr := []int{1, 2, 3, 4, 5}
	target := 6

	result := TwoSum(arr, target)
	sum := arr[result[0]] + arr[result[1]]

	if sum != target {
		t.Errorf("expected sum %d, got %d", target, sum)
	}
}
