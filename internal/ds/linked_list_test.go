package ds

import (
	"reflect"
	"testing"
)

func TestLinkedList_InsertAtHead(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAtHead(3)
	ll.InsertAtHead(2)
	ll.InsertAtHead(1)

	if ll.Size() != 3 {
		t.Errorf("expected size 3, got %d", ll.Size())
	}

	expected := []interface{}{1, 2, 3}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_InsertAtTail(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	if ll.Size() != 3 {
		t.Errorf("expected size 3, got %d", ll.Size())
	}

	expected := []interface{}{1, 2, 3}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_InsertAtPosition(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAtTail(1)
	ll.InsertAtTail(3)

	// Insert at middle
	if !ll.InsertAtPosition(2, 1) {
		t.Error("insert at position 1 should succeed")
	}

	expected := []interface{}{1, 2, 3}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}

	// Insert at head (position 0)
	if !ll.InsertAtPosition(0, 0) {
		t.Error("insert at position 0 should succeed")
	}

	// Insert at tail (position size)
	if !ll.InsertAtPosition(4, ll.Size()) {
		t.Error("insert at tail position should succeed")
	}

	expected = []interface{}{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_InsertAtPositionInvalid(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)

	if ll.InsertAtPosition(99, -1) {
		t.Error("insert at negative position should fail")
	}

	if ll.InsertAtPosition(99, 10) {
		t.Error("insert at position > size should fail")
	}
}

func TestLinkedList_DeleteAtHead(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	val, ok := ll.DeleteAtHead()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	if ll.Size() != 2 {
		t.Errorf("expected size 2, got %d", ll.Size())
	}

	expected := []interface{}{2, 3}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_DeleteAtHeadEmpty(t *testing.T) {
	ll := NewLinkedList()

	val, ok := ll.DeleteAtHead()
	if ok {
		t.Error("delete from empty list should fail")
	}
	if val != nil {
		t.Errorf("expected nil, got %v", val)
	}
}

func TestLinkedList_DeleteAtTail(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	val, ok := ll.DeleteAtTail()
	if !ok || val != 3 {
		t.Errorf("expected 3, got %v", val)
	}

	if ll.Size() != 2 {
		t.Errorf("expected size 2, got %d", ll.Size())
	}

	expected := []interface{}{1, 2}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_DeleteAtTailSingleElement(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(42)

	val, ok := ll.DeleteAtTail()
	if !ok || val != 42 {
		t.Errorf("expected 42, got %v", val)
	}

	if !ll.IsEmpty() {
		t.Error("list should be empty")
	}
}

func TestLinkedList_DeleteValue(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)
	ll.InsertAtTail(4)

	// Delete middle value
	if !ll.DeleteValue(3) {
		t.Error("delete value 3 should succeed")
	}

	expected := []interface{}{1, 2, 4}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}

	// Delete head
	if !ll.DeleteValue(1) {
		t.Error("delete value 1 should succeed")
	}

	// Delete tail
	if !ll.DeleteValue(4) {
		t.Error("delete value 4 should succeed")
	}

	expected = []interface{}{2}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_DeleteValueNotFound(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)

	if ll.DeleteValue(99) {
		t.Error("delete non-existent value should fail")
	}

	if ll.Size() != 2 {
		t.Errorf("size should remain 2, got %d", ll.Size())
	}
}

func TestLinkedList_Search(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

	if !ll.Search(20) {
		t.Error("should find value 20")
	}

	if ll.Search(99) {
		t.Error("should not find value 99")
	}
}

func TestLinkedList_Get(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail("a")
	ll.InsertAtTail("b")
	ll.InsertAtTail("c")

	val, ok := ll.Get(1)
	if !ok || val != "b" {
		t.Errorf("expected 'b', got %v", val)
	}

	val, ok = ll.Get(0)
	if !ok || val != "a" {
		t.Errorf("expected 'a', got %v", val)
	}

	val, ok = ll.Get(2)
	if !ok || val != "c" {
		t.Errorf("expected 'c', got %v", val)
	}
}

func TestLinkedList_GetInvalid(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)

	val, ok := ll.Get(-1)
	if ok {
		t.Error("get at negative position should fail")
	}
	if val != nil {
		t.Errorf("expected nil, got %v", val)
	}

	val, ok = ll.Get(10)
	if ok {
		t.Error("get at position >= size should fail")
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)
	ll.InsertAtTail(4)

	ll.Reverse()

	expected := []interface{}{4, 3, 2, 1}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_ReverseEmpty(t *testing.T) {
	ll := NewLinkedList()
	ll.Reverse()

	if !ll.IsEmpty() {
		t.Error("reversed empty list should still be empty")
	}
}

func TestLinkedList_ReverseSingleElement(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(42)

	ll.Reverse()

	expected := []interface{}{42}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	ll := NewLinkedList()

	if !ll.IsEmpty() {
		t.Error("new list should be empty")
	}

	ll.InsertAtTail(1)
	if ll.IsEmpty() {
		t.Error("list with element should not be empty")
	}

	ll.DeleteAtHead()
	if !ll.IsEmpty() {
		t.Error("list should be empty after deleting last element")
	}
}

func TestLinkedList_Clear(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	ll.Clear()

	if !ll.IsEmpty() {
		t.Error("list should be empty after clear")
	}

	if ll.Size() != 0 {
		t.Errorf("expected size 0, got %d", ll.Size())
	}
}

func TestLinkedList_MixedOperations(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAtHead(2)
	ll.InsertAtTail(3)
	ll.InsertAtHead(1)
	ll.InsertAtTail(4)

	expected := []interface{}{1, 2, 3, 4}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}

	ll.DeleteValue(2)
	ll.InsertAtPosition(5, 2)

	expected = []interface{}{1, 3, 5, 4}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected %v, got %v", expected, ll.ToSlice())
	}
}
