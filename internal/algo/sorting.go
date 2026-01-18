package algo

// Why interviewers ask this:
// Sorting algorithms test understanding of time/space complexity, recursion,
// divide-and-conquer, and in-place operations. Knowing when to use which
// algorithm is crucial for optimization discussions.

// Common pitfalls:
// - Not understanding stability (maintaining relative order)
// - Incorrect pivot selection in QuickSort
// - Forgetting base cases in recursive sorts
// - Not handling empty or single-element arrays
// - Confusion about in-place vs out-of-place sorting

// Key takeaway:
// QuickSort: O(n log n) average, O(n²) worst, in-place, unstable
// MergeSort: O(n log n) always, O(n) space, stable
// Choose based on requirements: stability, space, worst-case guarantees

// QuickSort sorts array in-place using divide-and-conquer
// Time Complexity: O(n log n) average, O(n²) worst
// Space Complexity: O(log n) for recursion stack
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		// Partition and get pivot index
		pivotIndex := partition(arr, low, high)

		// Recursively sort left and right
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose last element as pivot
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place pivot in correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort sorts array using divide-and-conquer (stable sort)
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Merge while both have elements
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// BubbleSort sorts array using bubble sort (for educational purposes)
// Time Complexity: O(n²)
// Space Complexity: O(1)
func BubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// Optimization: if no swaps, array is sorted
		if !swapped {
			break
		}
	}
}

// InsertionSort sorts array using insertion sort
// Time Complexity: O(n²), but O(n) for nearly sorted arrays
// Space Complexity: O(1)
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements greater than key one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

// SelectionSort sorts array using selection sort
// Time Complexity: O(n²)
// Space Complexity: O(1)
func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// Find minimum in unsorted portion
		minIdx := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		// Swap with first unsorted element
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// HeapSort sorts array using heap sort
// Time Complexity: O(n log n)
// Space Complexity: O(1)
func HeapSort(arr []int) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// Heapify reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// IsSorted checks if array is sorted
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// KthLargest finds kth largest element using QuickSelect
// Time Complexity: O(n) average, O(n²) worst
// Space Complexity: O(1)
func KthLargest(arr []int, k int) int {
	// Convert to kth smallest from end
	return quickSelect(arr, 0, len(arr)-1, len(arr)-k)
}

func quickSelect(arr []int, low, high, k int) int {
	if low == high {
		return arr[low]
	}

	pivotIndex := partition(arr, low, high)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelect(arr, low, pivotIndex-1, k)
	} else {
		return quickSelect(arr, pivotIndex+1, high, k)
	}
}
