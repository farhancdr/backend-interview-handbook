package algo

// Why interviewers ask this:
// Sliding window is a powerful technique for array/string problems involving
// subarrays or substrings. It optimizes brute force O(n²) or O(n³) solutions
// to O(n). Very common in substring and subarray problems.

// Common pitfalls:
// - Not understanding when to expand vs shrink window
// - Forgetting to update window state when shrinking
// - Off-by-one errors in window boundaries
// - Not handling edge cases (empty array, window larger than array)
// - Confusion about fixed vs variable window size

// Key takeaway:
// Sliding window maintains a window of elements and slides it across the array.
// Fixed window: move both pointers together. Variable window: expand/shrink based on condition.
// Reduces time complexity from O(n²) to O(n).

// MaxSumSubarray finds maximum sum of k consecutive elements
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSumSubarray(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}

	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	// Slide the window
	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// LongestSubstringKDistinct finds longest substring with at most k distinct characters
// Time Complexity: O(n)
// Space Complexity: O(k)
func LongestSubstringKDistinct(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	charCount := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		// Expand window
		charCount[s[right]]++

		// Shrink window if too many distinct characters
		for len(charCount) > k {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		// Update max length
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

// MinSubarraySum finds minimum length subarray with sum >= target
// Time Complexity: O(n)
// Space Complexity: O(1)
func MinSubarraySum(arr []int, target int) int {
	minLen := len(arr) + 1
	left := 0
	sum := 0

	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		// Shrink window while sum >= target
		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= arr[left]
			left++
		}
	}

	if minLen == len(arr)+1 {
		return 0 // Not found
	}
	return minLen
}

// LongestSubstringWithoutRepeating finds longest substring without repeating characters
// Time Complexity: O(n)
// Space Complexity: O(min(n, charset size))
func LongestSubstringWithoutRepeating(s string) int {
	charIndex := make(map[byte]int)
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		// If character seen before and in current window
		if idx, found := charIndex[s[right]]; found && idx >= left {
			left = idx + 1
		}

		charIndex[s[right]] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

// MaxConsecutiveOnes finds max consecutive 1s after flipping at most k 0s
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxConsecutiveOnes(arr []int, k int) int {
	left := 0
	zeros := 0
	maxLen := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] == 0 {
			zeros++
		}

		// Shrink window if too many zeros
		for zeros > k {
			if arr[left] == 0 {
				zeros--
			}
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

// FindAnagrams finds all start indices of anagrams of p in s
// Time Complexity: O(n)
// Space Complexity: O(1) - fixed size map (26 letters)
func FindAnagrams(s string, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}

	// Count characters in p
	pCount := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pCount[p[i]]++
	}

	windowCount := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		// Add character to window
		windowCount[s[right]]++

		// Shrink window if too large
		if right-left+1 > len(p) {
			windowCount[s[left]]--
			if windowCount[s[left]] == 0 {
				delete(windowCount, s[left])
			}
			left++
		}

		// Check if window is an anagram
		if right-left+1 == len(p) && mapsEqual(windowCount, pCount) {
			result = append(result, left)
		}
	}

	return result
}

// CharacterReplacement finds longest substring with same character after k replacements
// Time Complexity: O(n)
// Space Complexity: O(1) - fixed size map (26 letters)
func CharacterReplacement(s string, k int) int {
	charCount := make(map[byte]int)
	left := 0
	maxCount := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		// Track most frequent character count
		if charCount[s[right]] > maxCount {
			maxCount = charCount[s[right]]
		}

		// Shrink window if replacements needed > k
		for right-left+1-maxCount > k {
			charCount[s[left]]--
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

// Helper function
func mapsEqual(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
