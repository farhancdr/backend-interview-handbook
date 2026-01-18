package advanced

// Why interviewers ask this:
// Generics (introduced in Go 1.18) are a major language feature. Understanding
// type parameters, constraints, and when to use generics demonstrates modern
// Go knowledge and ability to write reusable, type-safe code.

// Common pitfalls:
// - Overusing generics when interfaces would suffice
// - Not understanding constraint syntax
// - Forgetting that methods cannot have type parameters
// - Confusion about type inference
// - Not knowing standard constraints (comparable, any, etc.)

// Key takeaway:
// Generics enable type-safe reusable code without interface{} and type assertions.
// Use constraints to restrict type parameters. Type inference often works automatically.
// Use generics for data structures and algorithms, not for everything.

// Min returns the minimum of two values
func Min[T comparable](a, b T) T {
	// Note: This won't compile as-is because comparable doesn't include <
	// This is a simplified example
	return a
}

// MinOrdered returns minimum using constraints.Ordered
func MinOrdered[T interface{ ~int | ~float64 | ~string }](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MaxOrdered returns maximum using constraints
func MaxOrdered[T interface{ ~int | ~float64 | ~string }](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Stack is a generic stack data structure
type Stack[T any] struct {
	items []T
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Size returns the number of items
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// IsEmpty checks if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Map applies a function to each element
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter filters elements based on predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce reduces slice to single value
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

// Contains checks if slice contains element
func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// Keys returns all keys from a map
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values from a map
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Pair represents a key-value pair
type Pair[K, V any] struct {
	Key   K
	Value V
}

// NewPair creates a new pair
func NewPair[K, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{Key: key, Value: value}
}

// Swap swaps two values
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// Reverse reverses a slice in place
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Equal checks if two slices are equal
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Number constraint for numeric types
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum calculates sum of numbers
func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Average calculates average of numbers
func Average[T Number](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := Sum(numbers)
	return float64(sum) / float64(len(numbers))
}
