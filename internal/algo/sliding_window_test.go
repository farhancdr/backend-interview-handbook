package algo

import (
	"reflect"
	"testing"
)

func TestMaxSumSubarray(t *testing.T) {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3

	result := MaxSumSubarray(arr, k)
	expected := 9 // [5, 1, 3]

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestMaxSumSubarray_KLargerThanArray(t *testing.T) {
	arr := []int{1, 2, 3}
	k := 5

	result := MaxSumSubarray(arr, k)
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"eceba", 2, 3}, // "ece"
		{"aa", 1, 2},    // "aa"
		{"abcba", 2, 3}, // "bcb" or "cbc"
		{"", 2, 0},      // empty
		{"abc", 0, 0},   // k=0
	}

	for _, tt := range tests {
		result := LongestSubstringKDistinct(tt.s, tt.k)
		if result != tt.expected {
			t.Errorf("LongestSubstringKDistinct(%s, %d): expected %d, got %d",
				tt.s, tt.k, tt.expected, result)
		}
	}
}

func TestMinSubarraySum(t *testing.T) {
	arr := []int{2, 3, 1, 2, 4, 3}
	target := 7

	result := MinSubarraySum(arr, target)
	expected := 2 // [4, 3]

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestMinSubarraySum_NotFound(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	target := 10

	result := MinSubarraySum(arr, target)
	if result != 0 {
		t.Errorf("expected 0 (not found), got %d", result)
	}
}

func TestLongestSubstringWithoutRepeating(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"abcabcbb", 3}, // "abc"
		{"bbbbb", 1},    // "b"
		{"pwwkew", 3},   // "wke"
		{"", 0},         // empty
		{"au", 2},       // "au"
	}

	for _, tt := range tests {
		result := LongestSubstringWithoutRepeating(tt.s)
		if result != tt.expected {
			t.Errorf("LongestSubstringWithoutRepeating(%s): expected %d, got %d",
				tt.s, tt.expected, result)
		}
	}
}

func TestMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	for _, tt := range tests {
		result := MaxConsecutiveOnes(tt.arr, tt.k)
		if result != tt.expected {
			t.Errorf("MaxConsecutiveOnes(k=%d): expected %d, got %d",
				tt.k, tt.expected, result)
		}
	}
}

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"

	result := FindAnagrams(s, p)
	expected := []int{0, 6} // "cba" at 0, "bac" at 6

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFindAnagrams_NoMatch(t *testing.T) {
	s := "abab"
	p := "xyz"

	result := FindAnagrams(s, p)
	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"ABAB", 2, 4},    // Replace both B's
		{"AABABBA", 1, 4}, // "AABA" or "ABBB"
		{"AAAA", 0, 4},    // Already all same
	}

	for _, tt := range tests {
		result := CharacterReplacement(tt.s, tt.k)
		if result != tt.expected {
			t.Errorf("CharacterReplacement(%s, %d): expected %d, got %d",
				tt.s, tt.k, tt.expected, result)
		}
	}
}

func TestSlidingWindow_EdgeCases(t *testing.T) {
	// Empty array
	if MaxSumSubarray([]int{}, 1) != 0 {
		t.Error("MaxSumSubarray should return 0 for empty array")
	}

	// Single element
	arr := []int{5}
	if MaxSumSubarray(arr, 1) != 5 {
		t.Error("MaxSumSubarray should handle single element")
	}

	// All negative numbers
	negArr := []int{-1, -2, -3, -4}
	result := MaxSumSubarray(negArr, 2)
	if result != -3 { // [-1, -2]
		t.Errorf("expected -3, got %d", result)
	}
}
