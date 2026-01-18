package ds

// Why interviewers ask this:
// Queue demonstrates understanding of FIFO (First In First Out) principle and is fundamental
// to many algorithms (BFS, task scheduling, buffering). Interviewers want to see if you can
// implement efficient enqueue/dequeue operations and handle edge cases.

// Common pitfalls:
// - Inefficient implementation using slices without considering memory leaks
// - Not handling empty queue operations properly
// - Forgetting that removing from front of slice is O(n) operation
// - Not considering circular buffer optimization for fixed-size queues

// Key takeaway:
// Queue follows FIFO principle. Enqueue adds to rear, Dequeue removes from front.
// Slice-based implementation is simple but can be inefficient. Circular buffer or
// linked list provides better performance for frequent dequeue operations.

// Queue represents a FIFO (First In First Out) data structure
// Time Complexity: Enqueue O(1) amortized, Dequeue O(n) for slice-based, Peek O(1)
// Space Complexity: O(n) where n is the number of elements
type Queue struct {
	items []interface{}
}

// NewQueue creates and returns a new empty queue
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

// Enqueue adds an element to the rear of the queue
// Time Complexity: O(1) amortized
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element from the queue
// Returns nil if queue is empty
// Time Complexity: O(n) - due to slice re-slicing
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}

// Peek returns the front element without removing it
// Returns nil if queue is empty
// Time Complexity: O(1)
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.items[0]
}

// IsEmpty returns true if the queue has no elements
// Time Complexity: O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
// Time Complexity: O(1)
func (q *Queue) Size() int {
	return len(q.items)
}

// Clear removes all elements from the queue
// Time Complexity: O(1)
func (q *Queue) Clear() {
	q.items = make([]interface{}, 0)
}

// CircularQueue is a more efficient queue implementation using circular buffer
// This avoids the O(n) dequeue operation of slice-based queue
type CircularQueue struct {
	items    []interface{}
	front    int
	rear     int
	size     int
	capacity int
}

// NewCircularQueue creates a circular queue with given capacity
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items:    make([]interface{}, capacity),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

// Enqueue adds an element to the circular queue
// Returns false if queue is full
// Time Complexity: O(1)
func (q *CircularQueue) Enqueue(item interface{}) bool {
	if q.IsFull() {
		return false
	}

	q.rear = (q.rear + 1) % q.capacity
	q.items[q.rear] = item
	q.size++

	return true
}

// Dequeue removes and returns the front element
// Returns nil and false if queue is empty
// Time Complexity: O(1)
func (q *CircularQueue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	item := q.items[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--

	return item, true
}

// Peek returns the front element without removing it
// Returns nil and false if queue is empty
// Time Complexity: O(1)
func (q *CircularQueue) Peek() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	return q.items[q.front], true
}

// IsEmpty returns true if queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

// IsFull returns true if queue is at capacity
func (q *CircularQueue) IsFull() bool {
	return q.size == q.capacity
}

// Size returns the number of elements
func (q *CircularQueue) Size() int {
	return q.size
}
