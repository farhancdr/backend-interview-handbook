package ds

import "testing"

func TestQueue_EnqueueAndDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("expected size 3, got %d", q.Size())
	}

	val := q.Dequeue()
	if val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	val = q.Dequeue()
	if val != 2 {
		t.Errorf("expected 2, got %v", val)
	}

	if q.Size() != 1 {
		t.Errorf("expected size 1, got %d", q.Size())
	}
}

func TestQueue_DequeueEmpty(t *testing.T) {
	q := NewQueue()

	val := q.Dequeue()
	if val != nil {
		t.Errorf("expected nil from empty queue, got %v", val)
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue()

	q.Enqueue("first")
	q.Enqueue("second")

	val := q.Peek()
	if val != "first" {
		t.Errorf("expected 'first', got %v", val)
	}

	// Peek should not remove element
	if q.Size() != 2 {
		t.Errorf("expected size 2 after peek, got %d", q.Size())
	}
}

func TestQueue_PeekEmpty(t *testing.T) {
	q := NewQueue()

	val := q.Peek()
	if val != nil {
		t.Errorf("expected nil from empty queue peek, got %v", val)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := NewQueue()

	if !q.IsEmpty() {
		t.Error("new queue should be empty")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("queue with element should not be empty")
	}

	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("queue should be empty after dequeuing last element")
	}
}

func TestQueue_Clear(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if !q.IsEmpty() {
		t.Error("queue should be empty after clear")
	}

	if q.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", q.Size())
	}
}

func TestQueue_FIFO(t *testing.T) {
	q := NewQueue()

	// Enqueue in order
	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}

	// Dequeue should return in same order (FIFO)
	for i := 1; i <= 5; i++ {
		val := q.Dequeue()
		if val != i {
			t.Errorf("expected %d, got %v", i, val)
		}
	}
}

func TestQueue_MixedOperations(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)

	val := q.Dequeue()
	if val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	q.Enqueue(3)
	q.Enqueue(4)

	val = q.Peek()
	if val != 2 {
		t.Errorf("expected 2, got %v", val)
	}

	if q.Size() != 3 {
		t.Errorf("expected size 3, got %d", q.Size())
	}
}

func TestCircularQueue_BasicOperations(t *testing.T) {
	q := NewCircularQueue(3)

	if !q.Enqueue(1) {
		t.Error("enqueue should succeed")
	}
	if !q.Enqueue(2) {
		t.Error("enqueue should succeed")
	}
	if !q.Enqueue(3) {
		t.Error("enqueue should succeed")
	}

	if q.Enqueue(4) {
		t.Error("enqueue should fail when queue is full")
	}

	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %v", val)
	}
}

func TestCircularQueue_IsFull(t *testing.T) {
	q := NewCircularQueue(2)

	if q.IsFull() {
		t.Error("new queue should not be full")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if !q.IsFull() {
		t.Error("queue should be full")
	}

	q.Dequeue()

	if q.IsFull() {
		t.Error("queue should not be full after dequeue")
	}
}

func TestCircularQueue_CircularBehavior(t *testing.T) {
	q := NewCircularQueue(3)

	// Fill the queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Remove two elements
	q.Dequeue()
	q.Dequeue()

	// Add two more (should wrap around)
	if !q.Enqueue(4) {
		t.Error("enqueue should succeed")
	}
	if !q.Enqueue(5) {
		t.Error("enqueue should succeed")
	}

	// Verify FIFO order
	val, _ := q.Dequeue()
	if val != 3 {
		t.Errorf("expected 3, got %v", val)
	}

	val, _ = q.Dequeue()
	if val != 4 {
		t.Errorf("expected 4, got %v", val)
	}

	val, _ = q.Dequeue()
	if val != 5 {
		t.Errorf("expected 5, got %v", val)
	}
}

func TestCircularQueue_Peek(t *testing.T) {
	q := NewCircularQueue(3)

	q.Enqueue(100)

	val, ok := q.Peek()
	if !ok {
		t.Error("peek should succeed")
	}
	if val != 100 {
		t.Errorf("expected 100, got %v", val)
	}

	// Size should remain unchanged
	if q.Size() != 1 {
		t.Errorf("expected size 1 after peek, got %d", q.Size())
	}
}

func TestCircularQueue_PeekEmpty(t *testing.T) {
	q := NewCircularQueue(3)

	val, ok := q.Peek()
	if ok {
		t.Error("peek should fail on empty queue")
	}
	if val != nil {
		t.Errorf("expected nil for failed peek, got %v", val)
	}
}

func TestCircularQueue_DequeueEmpty(t *testing.T) {
	q := NewCircularQueue(3)

	val, ok := q.Dequeue()
	if ok {
		t.Error("dequeue should fail on empty queue")
	}
	if val != nil {
		t.Errorf("expected nil for failed dequeue, got %v", val)
	}
}

func TestQueue_SingleElement(t *testing.T) {
	q := NewQueue()

	q.Enqueue("only")

	if q.IsEmpty() {
		t.Error("queue should not be empty")
	}

	val := q.Peek()
	if val != "only" {
		t.Errorf("expected 'only', got %v", val)
	}

	val = q.Dequeue()
	if val != "only" {
		t.Errorf("expected 'only', got %v", val)
	}

	if !q.IsEmpty() {
		t.Error("queue should be empty after dequeuing single element")
	}
}
