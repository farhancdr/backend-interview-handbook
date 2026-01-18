package algo

// Why interviewers ask this:
// Two pointers is a fundamental technique for array/string problems. It's efficient,
// elegant, and demonstrates understanding of space-time tradeoffs. Common in
// array manipulation, palindrome, and pair-finding problems.

// Common pitfalls:
// - Not handling empty arrays
// - Incorrect pointer movement logic
// - Off-by-one errors
// - Not considering duplicates
// - Forgetting to handle edge cases (single element, all same)

// Key takeaway:
// Two pointers reduces O(n²) brute force to O(n) by using two indices moving
// toward each other or in the same direction. Works great on sorted arrays.

// TwoSum finds two numbers that add up to target in sorted array
// Time Complexity: O(n)
// Space Complexity: O(1)
func TwoSum(arr []int, target int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1} // Not found
}

// ThreeSum finds all unique triplets that sum to zero
// Time Complexity: O(n²)
// Space Complexity: O(1) excluding output
func ThreeSum(arr []int) [][]int {
	// First, sort the array
	sortArray(arr)
	result := [][]int{}

	for i := 0; i < len(arr)-2; i++ {
		// Skip duplicates
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left, right := i+1, len(arr)-1
		target := -arr[i]

		for left < right {
			sum := arr[left] + arr[right]

			if sum == target {
				result = append(result, []int{arr[i], arr[left], arr[right]})

				// Skip duplicates
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// RemoveDuplicates removes duplicates from sorted array in-place
// Returns new length
// Time Complexity: O(n)
// Space Complexity: O(1)
func RemoveDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	slow := 0

	for fast := 1; fast < len(arr); fast++ {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
	}

	return slow + 1
}

// IsPalindrome checks if string is palindrome
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// ReverseString reverses string in-place
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// MoveZeroes moves all zeros to end while maintaining order
// Time Complexity: O(n)
// Space Complexity: O(1)
func MoveZeroes(arr []int) {
	slow := 0

	// Move all non-zeros to front
	for fast := 0; fast < len(arr); fast++ {
		if arr[fast] != 0 {
			arr[slow] = arr[fast]
			slow++
		}
	}

	// Fill rest with zeros
	for slow < len(arr) {
		arr[slow] = 0
		slow++
	}
}

// ContainerWithMostWater finds max area between two lines
// Time Complexity: O(n)
// Space Complexity: O(1)
func ContainerWithMostWater(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Area = width * min(height[left], height[right])
		width := right - left
		h := min(height[left], height[right])
		area := width * h

		if area > maxArea {
			maxArea = area
		}

		// Move pointer with smaller height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// PartitionArray partitions array around pivot
// All elements < pivot go to left, >= pivot go to right
// Time Complexity: O(n)
// Space Complexity: O(1)
func PartitionArray(arr []int, pivot int) int {
	left := 0

	for right := 0; right < len(arr); right++ {
		if arr[right] < pivot {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
	}

	return left
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Simple sort for ThreeSum (using insertion sort for simplicity)
func sortArray(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
