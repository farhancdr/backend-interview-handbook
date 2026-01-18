package ds

// Why interviewers ask this:
// Stack is a fundamental data structure that demonstrates understanding of LIFO (Last In First Out)
// principle. It's used in many algorithms (DFS, expression evaluation, backtracking) and is often
// the first step in testing a candidate's ability to implement basic data structures.

// Common pitfalls:
// - Not handling empty stack operations (pop/peek on empty stack)
// - Forgetting to check capacity in fixed-size implementations
// - Not considering thread-safety in concurrent scenarios
// - Inefficient implementation using wrong underlying structure

// Key takeaway:
// Stack follows LIFO principle. Push adds to top, Pop removes from top, Peek views top without removal.
// In Go, slices make excellent stack backing structures with O(1) amortized push/pop operations.

// Stack represents a LIFO (Last In First Out) data structure
// Time Complexity: Push O(1) amortized, Pop O(1), Peek O(1), IsEmpty O(1)
// Space Complexity: O(n) where n is the number of elements
type Stack struct {
	items []interface{}
}

// NewStack creates and returns a new empty stack
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push adds an element to the top of the stack
// Time Complexity: O(1) amortized
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
// Returns nil if stack is empty
// Time Complexity: O(1)
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item
}

// Peek returns the top element without removing it
// Returns nil if stack is empty
// Time Complexity: O(1)
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.items[len(s.items)-1]
}

// IsEmpty returns true if the stack has no elements
// Time Complexity: O(1)
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
// Time Complexity: O(1)
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack
// Time Complexity: O(1)
func (s *Stack) Clear() {
	s.items = make([]interface{}, 0)
}

// IntStack is a type-safe stack for integers
// This demonstrates how to create specialized stacks without using interface{}
type IntStack struct {
	items []int
}

// NewIntStack creates a new integer stack
func NewIntStack() *IntStack {
	return &IntStack{
		items: make([]int, 0),
	}
}

// Push adds an integer to the stack
func (s *IntStack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top integer
// Returns 0 and false if stack is empty
func (s *IntStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item, true
}

// Peek returns the top integer without removing it
// Returns 0 and false if stack is empty
func (s *IntStack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty
func (s *IntStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements
func (s *IntStack) Size() int {
	return len(s.items)
}
