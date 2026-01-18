package algo

// Why interviewers ask this:
// Binary search is one of the most fundamental algorithms. It demonstrates understanding
// of divide-and-conquer, time complexity, and edge case handling. Almost every technical
// interview includes at least one binary search problem or variant.

// Common pitfalls:
// - Off-by-one errors in loop conditions
// - Integer overflow when calculating mid: (left + right) / 2
// - Not handling empty arrays or single elements
// - Forgetting that array must be sorted
// - Infinite loops from incorrect boundary updates

// Key takeaway:
// Binary search reduces search space by half each iteration: O(log n) time.
// Always use left + (right - left) / 2 to avoid overflow.
// Works only on sorted arrays.

// BinarySearch finds target in sorted array, returns index or -1
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		// Avoid overflow: use left + (right - left) / 2
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Not found
}

// BinarySearchRecursive implements binary search recursively
// Time Complexity: O(log n)
// Space Complexity: O(log n) due to call stack
func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchHelper(arr, target, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchHelper(arr, target, mid+1, right)
	} else {
		return binarySearchHelper(arr, target, left, mid-1)
	}
}

// FindFirstOccurrence finds the first occurrence of target
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindFirstOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching left
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// FindLastOccurrence finds the last occurrence of target
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindLastOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			left = mid + 1 // Continue searching right
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// SearchInsertPosition finds position where target should be inserted
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchInsertPosition(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left // Insert position
}

// SearchRotatedArray searches in rotated sorted array
// Example: [4,5,6,7,0,1,2], target = 0 -> index 4
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SearchRotatedArray(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		// Determine which half is sorted
		if arr[left] <= arr[mid] {
			// Left half is sorted
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// FindPeakElement finds any peak element (element greater than neighbors)
// Time Complexity: O(log n)
// Space Complexity: O(1)
func FindPeakElement(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] > arr[mid+1] {
			// Peak is on the left (including mid)
			right = mid
		} else {
			// Peak is on the right
			left = mid + 1
		}
	}

	return left
}

// SquareRoot finds integer square root using binary search
// Time Complexity: O(log n)
// Space Complexity: O(1)
func SquareRoot(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	result := 0

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
