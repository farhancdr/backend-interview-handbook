package ds

import "testing"

func TestStack_PushAndPop(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}

	val := s.Pop()
	if val != 3 {
		t.Errorf("expected 3, got %v", val)
	}

	val = s.Pop()
	if val != 2 {
		t.Errorf("expected 2, got %v", val)
	}

	if s.Size() != 1 {
		t.Errorf("expected size 1, got %d", s.Size())
	}
}

func TestStack_PopEmpty(t *testing.T) {
	s := NewStack()

	val := s.Pop()
	if val != nil {
		t.Errorf("expected nil from empty stack, got %v", val)
	}
}

func TestStack_Peek(t *testing.T) {
	s := NewStack()

	s.Push("first")
	s.Push("second")

	val := s.Peek()
	if val != "second" {
		t.Errorf("expected 'second', got %v", val)
	}

	// Peek should not remove element
	if s.Size() != 2 {
		t.Errorf("expected size 2 after peek, got %d", s.Size())
	}
}

func TestStack_PeekEmpty(t *testing.T) {
	s := NewStack()

	val := s.Peek()
	if val != nil {
		t.Errorf("expected nil from empty stack peek, got %v", val)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := NewStack()

	if !s.IsEmpty() {
		t.Error("new stack should be empty")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Error("stack with element should not be empty")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Error("stack should be empty after popping last element")
	}
}

func TestStack_Clear(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Clear()

	if !s.IsEmpty() {
		t.Error("stack should be empty after clear")
	}

	if s.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", s.Size())
	}
}

func TestStack_MixedTypes(t *testing.T) {
	s := NewStack()

	s.Push(42)
	s.Push("string")
	s.Push(3.14)
	s.Push(true)

	if val := s.Pop(); val != true {
		t.Errorf("expected true, got %v", val)
	}

	if val := s.Pop(); val != 3.14 {
		t.Errorf("expected 3.14, got %v", val)
	}

	if val := s.Pop(); val != "string" {
		t.Errorf("expected 'string', got %v", val)
	}

	if val := s.Pop(); val != 42 {
		t.Errorf("expected 42, got %v", val)
	}
}

func TestStack_LIFO(t *testing.T) {
	s := NewStack()

	// Push in order
	for i := 1; i <= 5; i++ {
		s.Push(i)
	}

	// Pop should return in reverse order (LIFO)
	for i := 5; i >= 1; i-- {
		val := s.Pop()
		if val != i {
			t.Errorf("expected %d, got %v", i, val)
		}
	}
}

func TestIntStack_TypeSafety(t *testing.T) {
	s := NewIntStack()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	val, ok := s.Pop()
	if !ok {
		t.Error("expected successful pop")
	}
	if val != 30 {
		t.Errorf("expected 30, got %d", val)
	}
}

func TestIntStack_PopEmpty(t *testing.T) {
	s := NewIntStack()

	val, ok := s.Pop()
	if ok {
		t.Error("expected pop to fail on empty stack")
	}
	if val != 0 {
		t.Errorf("expected 0 for failed pop, got %d", val)
	}
}

func TestIntStack_Peek(t *testing.T) {
	s := NewIntStack()

	s.Push(100)

	val, ok := s.Peek()
	if !ok {
		t.Error("expected successful peek")
	}
	if val != 100 {
		t.Errorf("expected 100, got %d", val)
	}

	// Size should remain unchanged
	if s.Size() != 1 {
		t.Errorf("expected size 1 after peek, got %d", s.Size())
	}
}

func TestIntStack_PeekEmpty(t *testing.T) {
	s := NewIntStack()

	val, ok := s.Peek()
	if ok {
		t.Error("expected peek to fail on empty stack")
	}
	if val != 0 {
		t.Errorf("expected 0 for failed peek, got %d", val)
	}
}

func TestIntStack_MultipleOperations(t *testing.T) {
	s := NewIntStack()

	// Test sequence of operations
	s.Push(1)
	s.Push(2)

	val, _ := s.Peek()
	if val != 2 {
		t.Errorf("expected peek 2, got %d", val)
	}

	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}

	s.Pop()
	s.Pop()

	val, _ = s.Peek()
	if val != 1 {
		t.Errorf("expected peek 1, got %d", val)
	}
}

func TestStack_SingleElement(t *testing.T) {
	s := NewStack()

	s.Push("only")

	if s.IsEmpty() {
		t.Error("stack should not be empty")
	}

	val := s.Peek()
	if val != "only" {
		t.Errorf("expected 'only', got %v", val)
	}

	val = s.Pop()
	if val != "only" {
		t.Errorf("expected 'only', got %v", val)
	}

	if !s.IsEmpty() {
		t.Error("stack should be empty after popping single element")
	}
}
