package advanced

import (
	"reflect"
	"testing"
)

func TestMinOrdered(t *testing.T) {
	// Test with int
	if MinOrdered(5, 3) != 3 {
		t.Error("MinOrdered(5, 3) should be 3")
	}

	// Test with float64
	if MinOrdered(5.5, 3.3) != 3.3 {
		t.Error("MinOrdered(5.5, 3.3) should be 3.3")
	}

	// Test with string
	if MinOrdered("b", "a") != "a" {
		t.Error("MinOrdered(b, a) should be a")
	}
}

func TestMaxOrdered(t *testing.T) {
	if MaxOrdered(5, 3) != 5 {
		t.Error("MaxOrdered(5, 3) should be 5")
	}

	if MaxOrdered(5.5, 3.3) != 5.5 {
		t.Error("MaxOrdered(5.5, 3.3) should be 5.5")
	}
}

func TestStack_Int(t *testing.T) {
	stack := NewStack[int]()

	// Test empty
	if !stack.IsEmpty() {
		t.Error("new stack should be empty")
	}

	// Test push
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Errorf("expected size 3, got %d", stack.Size())
	}

	// Test peek
	val, ok := stack.Peek()
	if !ok || val != 3 {
		t.Errorf("expected peek to return 3, got %d", val)
	}

	// Test pop
	val, ok = stack.Pop()
	if !ok || val != 3 {
		t.Errorf("expected pop to return 3, got %d", val)
	}

	if stack.Size() != 2 {
		t.Errorf("expected size 2 after pop, got %d", stack.Size())
	}
}

func TestStack_String(t *testing.T) {
	stack := NewStack[string]()

	stack.Push("hello")
	stack.Push("world")

	val, ok := stack.Pop()
	if !ok || val != "world" {
		t.Errorf("expected 'world', got %s", val)
	}
}

func TestStack_Empty(t *testing.T) {
	stack := NewStack[int]()

	val, ok := stack.Pop()
	if ok {
		t.Error("pop on empty stack should return false")
	}
	if val != 0 {
		t.Errorf("pop on empty stack should return zero value, got %d", val)
	}
}

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	// Double each number
	doubled := Map(numbers, func(n int) int {
		return n * 2
	})

	expected := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(doubled, expected) {
		t.Errorf("expected %v, got %v", expected, doubled)
	}
}

func TestMap_StringToInt(t *testing.T) {
	strings := []string{"1", "2", "3"}

	// Convert to lengths
	lengths := Map(strings, func(s string) int {
		return len(s)
	})

	expected := []int{1, 1, 1}
	if !reflect.DeepEqual(lengths, expected) {
		t.Errorf("expected %v, got %v", expected, lengths)
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}

	// Filter even numbers
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(evens, expected) {
		t.Errorf("expected %v, got %v", expected, evens)
	}
}

func TestReduce(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	// Sum all numbers
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})

	if sum != 15 {
		t.Errorf("expected 15, got %d", sum)
	}
}

func TestContains(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	if !Contains(numbers, 3) {
		t.Error("should contain 3")
	}

	if Contains(numbers, 10) {
		t.Error("should not contain 10")
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	keys := Keys(m)

	if len(keys) != 3 {
		t.Errorf("expected 3 keys, got %d", len(keys))
	}

	// Check all keys are present (order not guaranteed)
	keyMap := make(map[string]bool)
	for _, k := range keys {
		keyMap[k] = true
	}

	if !keyMap["a"] || !keyMap["b"] || !keyMap["c"] {
		t.Error("missing expected keys")
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := Values(m)

	if len(values) != 3 {
		t.Errorf("expected 3 values, got %d", len(values))
	}
}

func TestPair(t *testing.T) {
	pair := NewPair("key", 42)

	if pair.Key != "key" {
		t.Errorf("expected key 'key', got %s", pair.Key)
	}

	if pair.Value != 42 {
		t.Errorf("expected value 42, got %d", pair.Value)
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)

	if a != 2 || b != 1 {
		t.Errorf("expected a=2, b=1, got a=%d, b=%d", a, b)
	}
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	Reverse(slice)

	expected := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("expected %v, got %v", expected, slice)
	}
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}

	if !Equal(a, b) {
		t.Error("a and b should be equal")
	}

	if Equal(a, c) {
		t.Error("a and c should not be equal")
	}
}

func TestSum(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	if Sum(ints) != 15 {
		t.Errorf("expected 15, got %d", Sum(ints))
	}

	floats := []float64{1.5, 2.5, 3.0}
	if Sum(floats) != 7.0 {
		t.Errorf("expected 7.0, got %f", Sum(floats))
	}
}

func TestAverage(t *testing.T) {
	numbers := []int{2, 4, 6, 8, 10}
	avg := Average(numbers)

	if avg != 6.0 {
		t.Errorf("expected 6.0, got %f", avg)
	}
}

func TestAverage_Empty(t *testing.T) {
	numbers := []int{}
	avg := Average(numbers)

	if avg != 0 {
		t.Errorf("expected 0, got %f", avg)
	}
}
