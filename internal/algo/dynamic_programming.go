package algo

// Why interviewers ask this:
// Dynamic programming is essential for optimization problems. It demonstrates
// understanding of overlapping subproblems, optimal substructure, and memoization.
// One of the most important algorithmic techniques for technical interviews.

// Common pitfalls:
// - Not identifying overlapping subproblems
// - Incorrect base cases
// - Wrong recurrence relation
// - Forgetting to memoize (leading to exponential time)
// - Confusion between top-down (memoization) vs bottom-up (tabulation)

// Key takeaway:
// DP = Recursion + Memoization. Break problem into subproblems, solve once, reuse.
// Two approaches: Top-down (recursive + memo) or Bottom-up (iterative + table).
// Time: O(n) or O(n²) typically. Space: O(n) for memo/table.

// Fibonacci calculates nth Fibonacci number using DP
// Time Complexity: O(n)
// Space Complexity: O(n)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	memo := make([]int, n+1)
	memo[0] = 0
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}

// FibonacciOptimized uses O(1) space
// Time Complexity: O(n)
// Space Complexity: O(1)
func FibonacciOptimized(n int) int {
	if n <= 1 {
		return n
	}

	prev2, prev1 := 0, 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// ClimbStairs calculates ways to climb n stairs (1 or 2 steps at a time)
// Time Complexity: O(n)
// Space Complexity: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// CoinChange finds minimum coins needed to make amount
// Time Complexity: O(amount * len(coins))
// Space Complexity: O(amount)
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	// Initialize with impossible value
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = minInt(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1 // Impossible
	}
	return dp[amount]
}

// LongestIncreasingSubsequence finds length of LIS
// Time Complexity: O(n²)
// Space Complexity: O(n)
func LongestIncreasingSubsequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // Each element is a subsequence of length 1
	}

	maxLen := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		maxLen = maxInt(maxLen, dp[i])
	}

	return maxLen
}

// MaxSubarraySum finds maximum sum of contiguous subarray (Kadane's algorithm)
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// Either extend current subarray or start new one
		currentSum = maxInt(nums[i], currentSum+nums[i])
		maxSum = maxInt(maxSum, currentSum)
	}

	return maxSum
}

// HouseRobber finds maximum money that can be robbed (can't rob adjacent houses)
// Time Complexity: O(n)
// Space Complexity: O(1)
func HouseRobber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2, prev1 := 0, 0

	for _, num := range nums {
		// Either rob current house + prev2, or skip and take prev1
		current := maxInt(prev1, prev2+num)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// UniquePaths finds number of unique paths in m x n grid
// Time Complexity: O(m * n)
// Space Complexity: O(m * n)
func UniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize first row and column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Fill the table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// Helper functions
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
