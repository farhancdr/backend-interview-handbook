package ds

import (
	"reflect"
	"testing"
)

func TestMinHeap_Insert(t *testing.T) {
	h := NewMinHeap()

	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(1)

	if h.Size() != 4 {
		t.Errorf("expected size 4, got %d", h.Size())
	}

	// Min should be 1
	min, ok := h.Peek()
	if !ok || min != 1 {
		t.Errorf("expected min 1, got %d", min)
	}
}

func TestMinHeap_ExtractMin(t *testing.T) {
	h := NewMinHeap()

	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(1)
	h.Insert(9)

	// Extract in sorted order
	expected := []int{1, 3, 5, 7, 9}
	for _, exp := range expected {
		val, ok := h.ExtractMin()
		if !ok {
			t.Error("extract should succeed")
		}
		if val != exp {
			t.Errorf("expected %d, got %d", exp, val)
		}
	}

	if !h.IsEmpty() {
		t.Error("heap should be empty after extracting all elements")
	}
}

func TestMinHeap_ExtractMinEmpty(t *testing.T) {
	h := NewMinHeap()

	val, ok := h.ExtractMin()
	if ok {
		t.Error("extract from empty heap should fail")
	}
	if val != 0 {
		t.Errorf("expected 0, got %d", val)
	}
}

func TestMinHeap_Peek(t *testing.T) {
	h := NewMinHeap()

	h.Insert(10)
	h.Insert(5)
	h.Insert(20)

	val, ok := h.Peek()
	if !ok || val != 5 {
		t.Errorf("expected peek 5, got %d", val)
	}

	// Size should remain unchanged
	if h.Size() != 3 {
		t.Errorf("expected size 3 after peek, got %d", h.Size())
	}
}

func TestMinHeap_PeekEmpty(t *testing.T) {
	h := NewMinHeap()

	val, ok := h.Peek()
	if ok {
		t.Error("peek on empty heap should fail")
	}
	if val != 0 {
		t.Errorf("expected 0, got %d", val)
	}
}

func TestMinHeap_IsEmpty(t *testing.T) {
	h := NewMinHeap()

	if !h.IsEmpty() {
		t.Error("new heap should be empty")
	}

	h.Insert(1)
	if h.IsEmpty() {
		t.Error("heap with element should not be empty")
	}

	h.ExtractMin()
	if !h.IsEmpty() {
		t.Error("heap should be empty after extracting last element")
	}
}

func TestMinHeap_Clear(t *testing.T) {
	h := NewMinHeap()

	h.Insert(1)
	h.Insert(2)
	h.Insert(3)

	h.Clear()

	if !h.IsEmpty() {
		t.Error("heap should be empty after clear")
	}

	if h.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", h.Size())
	}
}

func TestMinHeap_BuildHeap(t *testing.T) {
	h := NewMinHeap()

	values := []int{9, 5, 6, 2, 3, 7, 1, 4, 8}
	h.BuildHeap(values)

	if h.Size() != len(values) {
		t.Errorf("expected size %d, got %d", len(values), h.Size())
	}

	// Extract all and verify sorted order
	var result []int
	for !h.IsEmpty() {
		val, _ := h.ExtractMin()
		result = append(result, val)
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMinHeap_HeapProperty(t *testing.T) {
	h := NewMinHeap()

	// Insert random values
	values := []int{15, 10, 20, 8, 25, 30, 5}
	for _, v := range values {
		h.Insert(v)
	}

	// Verify heap property: parent <= children
	items := h.ToSlice()
	for i := 0; i < len(items); i++ {
		leftChild := 2*i + 1
		rightChild := 2*i + 2

		if leftChild < len(items) && items[i] > items[leftChild] {
			t.Errorf("heap property violated at index %d", i)
		}

		if rightChild < len(items) && items[i] > items[rightChild] {
			t.Errorf("heap property violated at index %d", i)
		}
	}
}

func TestMaxHeap_Insert(t *testing.T) {
	h := NewMaxHeap()

	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(1)
	h.Insert(9)

	if h.Size() != 5 {
		t.Errorf("expected size 5, got %d", h.Size())
	}

	// Max should be 9
	max, ok := h.Peek()
	if !ok || max != 9 {
		t.Errorf("expected max 9, got %d", max)
	}
}

func TestMaxHeap_ExtractMax(t *testing.T) {
	h := NewMaxHeap()

	h.Insert(5)
	h.Insert(3)
	h.Insert(7)
	h.Insert(1)
	h.Insert(9)

	// Extract in descending order
	expected := []int{9, 7, 5, 3, 1}
	for _, exp := range expected {
		val, ok := h.ExtractMax()
		if !ok {
			t.Error("extract should succeed")
		}
		if val != exp {
			t.Errorf("expected %d, got %d", exp, val)
		}
	}

	if !h.IsEmpty() {
		t.Error("heap should be empty after extracting all elements")
	}
}

func TestMaxHeap_ExtractMaxEmpty(t *testing.T) {
	h := NewMaxHeap()

	val, ok := h.ExtractMax()
	if ok {
		t.Error("extract from empty heap should fail")
	}
	if val != 0 {
		t.Errorf("expected 0, got %d", val)
	}
}

func TestMaxHeap_Peek(t *testing.T) {
	h := NewMaxHeap()

	h.Insert(10)
	h.Insert(50)
	h.Insert(20)

	val, ok := h.Peek()
	if !ok || val != 50 {
		t.Errorf("expected peek 50, got %d", val)
	}

	if h.Size() != 3 {
		t.Errorf("expected size 3 after peek, got %d", h.Size())
	}
}

func TestMaxHeap_IsEmpty(t *testing.T) {
	h := NewMaxHeap()

	if !h.IsEmpty() {
		t.Error("new heap should be empty")
	}

	h.Insert(1)
	if h.IsEmpty() {
		t.Error("heap with element should not be empty")
	}

	h.ExtractMax()
	if !h.IsEmpty() {
		t.Error("heap should be empty after extracting last element")
	}
}

func TestMinHeap_DuplicateValues(t *testing.T) {
	h := NewMinHeap()

	h.Insert(5)
	h.Insert(5)
	h.Insert(5)
	h.Insert(3)
	h.Insert(3)

	if h.Size() != 5 {
		t.Errorf("expected size 5, got %d", h.Size())
	}

	// Extract all
	result := []int{}
	for !h.IsEmpty() {
		val, _ := h.ExtractMin()
		result = append(result, val)
	}

	expected := []int{3, 3, 5, 5, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMinHeap_SingleElement(t *testing.T) {
	h := NewMinHeap()

	h.Insert(42)

	val, ok := h.Peek()
	if !ok || val != 42 {
		t.Errorf("expected 42, got %d", val)
	}

	val, ok = h.ExtractMin()
	if !ok || val != 42 {
		t.Errorf("expected 42, got %d", val)
	}

	if !h.IsEmpty() {
		t.Error("heap should be empty")
	}
}
