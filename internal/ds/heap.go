package ds

// Why interviewers ask this:
// Heaps are essential for priority queues, heap sort, and finding k-th largest/smallest elements.
// They demonstrate understanding of complete binary trees, array representation of trees, and
// efficient priority-based operations. Common in system design (task scheduling, event processing).

// Common pitfalls:
// - Confusing min-heap and max-heap properties
// - Incorrect parent/child index calculations
// - Not maintaining heap property after insert/delete
// - Forgetting that heaps are complete binary trees (not BSTs)
// - Off-by-one errors in array indexing

// Key takeaway:
// Min-heap: parent <= children. Max-heap: parent >= children. Array-based implementation
// with parent at i, left child at 2i+1, right child at 2i+2. O(log n) insert/extract,
// O(1) peek. Used for priority queues and efficient k-th element problems.

// MinHeap represents a min-heap data structure
// Time Complexity: Insert O(log n), ExtractMin O(log n), Peek O(1)
// Space Complexity: O(n) where n is the number of elements
type MinHeap struct {
	items []int
}

// NewMinHeap creates a new empty min-heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		items: make([]int, 0),
	}
}

// Insert adds a value to the heap
// Time Complexity: O(log n)
func (h *MinHeap) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

// ExtractMin removes and returns the minimum value (root)
// Returns 0 and false if heap is empty
// Time Complexity: O(log n)
func (h *MinHeap) ExtractMin() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	min := h.items[0]
	lastIdx := len(h.items) - 1

	// Move last element to root
	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]

	// Restore heap property
	if len(h.items) > 0 {
		h.heapifyDown(0)
	}

	return min, true
}

// Peek returns the minimum value without removing it
// Returns 0 and false if heap is empty
// Time Complexity: O(1)
func (h *MinHeap) Peek() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	return h.items[0], true
}

// heapifyUp maintains heap property by moving element up
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) / 2

		if h.items[index] >= h.items[parentIdx] {
			break
		}

		// Swap with parent
		h.items[index], h.items[parentIdx] = h.items[parentIdx], h.items[index]
		index = parentIdx
	}
}

// heapifyDown maintains heap property by moving element down
func (h *MinHeap) heapifyDown(index int) {
	size := len(h.items)

	for {
		smallest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		// Check if left child is smaller
		if leftChild < size && h.items[leftChild] < h.items[smallest] {
			smallest = leftChild
		}

		// Check if right child is smaller
		if rightChild < size && h.items[rightChild] < h.items[smallest] {
			smallest = rightChild
		}

		// If current node is smallest, heap property is satisfied
		if smallest == index {
			break
		}

		// Swap with smallest child
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		index = smallest
	}
}

// IsEmpty returns true if heap has no elements
func (h *MinHeap) IsEmpty() bool {
	return len(h.items) == 0
}

// Size returns the number of elements in the heap
func (h *MinHeap) Size() int {
	return len(h.items)
}

// Clear removes all elements from the heap
func (h *MinHeap) Clear() {
	h.items = make([]int, 0)
}

// BuildHeap creates a heap from an array of values
// Time Complexity: O(n)
func (h *MinHeap) BuildHeap(values []int) {
	h.items = make([]int, len(values))
	copy(h.items, values)

	// Start from last non-leaf node and heapify down
	for i := len(h.items)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// ToSlice returns the heap as a slice (not sorted)
func (h *MinHeap) ToSlice() []int {
	result := make([]int, len(h.items))
	copy(result, h.items)
	return result
}

// MaxHeap represents a max-heap data structure
type MaxHeap struct {
	items []int
}

// NewMaxHeap creates a new empty max-heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		items: make([]int, 0),
	}
}

// Insert adds a value to the max-heap
// Time Complexity: O(log n)
func (h *MaxHeap) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

// ExtractMax removes and returns the maximum value (root)
// Returns 0 and false if heap is empty
// Time Complexity: O(log n)
func (h *MaxHeap) ExtractMax() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	max := h.items[0]
	lastIdx := len(h.items) - 1

	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]

	if len(h.items) > 0 {
		h.heapifyDown(0)
	}

	return max, true
}

// Peek returns the maximum value without removing it
// Time Complexity: O(1)
func (h *MaxHeap) Peek() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	return h.items[0], true
}

// heapifyUp maintains max-heap property by moving element up
func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) / 2

		if h.items[index] <= h.items[parentIdx] {
			break
		}

		h.items[index], h.items[parentIdx] = h.items[parentIdx], h.items[index]
		index = parentIdx
	}
}

// heapifyDown maintains max-heap property by moving element down
func (h *MaxHeap) heapifyDown(index int) {
	size := len(h.items)

	for {
		largest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		if leftChild < size && h.items[leftChild] > h.items[largest] {
			largest = leftChild
		}

		if rightChild < size && h.items[rightChild] > h.items[largest] {
			largest = rightChild
		}

		if largest == index {
			break
		}

		h.items[index], h.items[largest] = h.items[largest], h.items[index]
		index = largest
	}
}

// IsEmpty returns true if heap has no elements
func (h *MaxHeap) IsEmpty() bool {
	return len(h.items) == 0
}

// Size returns the number of elements in the heap
func (h *MaxHeap) Size() int {
	return len(h.items)
}
