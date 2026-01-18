package algo

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
	}

	for _, tt := range tests {
		result := Fibonacci(tt.n)
		if result != tt.expected {
			t.Errorf("Fibonacci(%d): expected %d, got %d", tt.n, tt.expected, result)
		}
	}
}

func TestFibonacciOptimized(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{15, 610},
	}

	for _, tt := range tests {
		result := FibonacciOptimized(tt.n)
		if result != tt.expected {
			t.Errorf("FibonacciOptimized(%d): expected %d, got %d", tt.n, tt.expected, result)
		}
	}
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, tt := range tests {
		result := ClimbStairs(tt.n)
		if result != tt.expected {
			t.Errorf("ClimbStairs(%d): expected %d, got %d", tt.n, tt.expected, result)
		}
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},   // 5+5+1
		{[]int{2}, 3, -1},         // Impossible
		{[]int{1}, 0, 0},          // Amount 0
		{[]int{1, 2, 5}, 100, 20}, // 20 coins of 5
	}

	for _, tt := range tests {
		result := CoinChange(tt.coins, tt.amount)
		if result != tt.expected {
			t.Errorf("CoinChange(%v, %d): expected %d, got %d",
				tt.coins, tt.amount, tt.expected, result)
		}
	}
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, // [2,3,7,101]
		{[]int{0, 1, 0, 3, 2, 3}, 4},           // [0,1,2,3]
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},        // [7]
		{[]int{}, 0},                           // Empty
	}

	for _, tt := range tests {
		result := LongestIncreasingSubsequence(tt.nums)
		if result != tt.expected {
			t.Errorf("LongestIncreasingSubsequence(%v): expected %d, got %d",
				tt.nums, tt.expected, result)
		}
	}
}

func TestMaxSubarraySum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // [4,-1,2,1]
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1, -2, -3}, -1},
	}

	for _, tt := range tests {
		result := MaxSubarraySum(tt.nums)
		if result != tt.expected {
			t.Errorf("MaxSubarraySum(%v): expected %d, got %d",
				tt.nums, tt.expected, result)
		}
	}
}

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},     // Rob house 1 and 3
		{[]int{2, 7, 9, 3, 1}, 12}, // Rob house 1, 3, and 5
		{[]int{5}, 5},              // Single house
		{[]int{}, 0},               // No houses
		{[]int{2, 1, 1, 2}, 4},     // Rob house 1 and 4
	}

	for _, tt := range tests {
		result := HouseRobber(tt.nums)
		if result != tt.expected {
			t.Errorf("HouseRobber(%v): expected %d, got %d",
				tt.nums, tt.expected, result)
		}
	}
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m        int
		n        int
		expected int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{1, 1, 1},
		{2, 2, 2},
	}

	for _, tt := range tests {
		result := UniquePaths(tt.m, tt.n)
		if result != tt.expected {
			t.Errorf("UniquePaths(%d, %d): expected %d, got %d",
				tt.m, tt.n, tt.expected, result)
		}
	}
}

func TestDP_EdgeCases(t *testing.T) {
	// Fibonacci with 0
	if Fibonacci(0) != 0 {
		t.Error("Fibonacci(0) should be 0")
	}

	// Coin change with 0 amount
	if CoinChange([]int{1, 2, 5}, 0) != 0 {
		t.Error("CoinChange with amount 0 should return 0")
	}

	// Max subarray with single element
	if MaxSubarraySum([]int{-5}) != -5 {
		t.Error("MaxSubarraySum with single negative should return that number")
	}

	// Unique paths 1x1 grid
	if UniquePaths(1, 1) != 1 {
		t.Error("UniquePaths(1,1) should be 1")
	}
}
