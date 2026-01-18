package ds

// Why interviewers ask this:
// Linked lists are fundamental for understanding pointer manipulation, dynamic memory allocation,
// and the trade-offs between array-based and pointer-based data structures. Many interview
// questions involve linked list manipulation (reversal, cycle detection, merging).

// Common pitfalls:
// - Losing reference to head when modifying the list
// - Not handling nil/empty list cases
// - Creating memory leaks by not properly updating pointers
// - Off-by-one errors in traversal
// - Not considering edge cases (single node, two nodes)

// Key takeaway:
// Linked lists provide O(1) insertion/deletion at known positions but O(n) search.
// Always handle nil cases and be careful with pointer manipulation. Drawing diagrams
// helps visualize pointer changes during operations.

// Node represents a single node in a singly linked list
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList represents a singly linked list
// Time Complexity: Insert O(1) at head, O(n) at tail/position
//
//	Delete O(1) at head, O(n) at tail/position
//	Search O(n)
//
// Space Complexity: O(n) where n is the number of nodes
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// NewLinkedList creates and returns a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// InsertAtHead adds a new node at the beginning of the list
// Time Complexity: O(1)
func (ll *LinkedList) InsertAtHead(value interface{}) {
	newNode := &Node{Value: value, Next: ll.head}
	ll.head = newNode

	if ll.tail == nil {
		ll.tail = newNode
	}

	ll.size++
}

// InsertAtTail adds a new node at the end of the list
// Time Complexity: O(1) with tail pointer, O(n) without
func (ll *LinkedList) InsertAtTail(value interface{}) {
	newNode := &Node{Value: value, Next: nil}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.Next = newNode
		ll.tail = newNode
	}

	ll.size++
}

// InsertAtPosition inserts a value at the specified position (0-indexed)
// Returns false if position is invalid
// Time Complexity: O(n)
func (ll *LinkedList) InsertAtPosition(value interface{}, position int) bool {
	if position < 0 || position > ll.size {
		return false
	}

	if position == 0 {
		ll.InsertAtHead(value)
		return true
	}

	if position == ll.size {
		ll.InsertAtTail(value)
		return true
	}

	newNode := &Node{Value: value}
	current := ll.head

	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	ll.size++

	return true
}

// DeleteAtHead removes the first node
// Returns the value and true if successful, nil and false if list is empty
// Time Complexity: O(1)
func (ll *LinkedList) DeleteAtHead() (interface{}, bool) {
	if ll.head == nil {
		return nil, false
	}

	value := ll.head.Value
	ll.head = ll.head.Next
	ll.size--

	if ll.head == nil {
		ll.tail = nil
	}

	return value, true
}

// DeleteAtTail removes the last node
// Returns the value and true if successful, nil and false if list is empty
// Time Complexity: O(n) - must traverse to second-to-last node
func (ll *LinkedList) DeleteAtTail() (interface{}, bool) {
	if ll.head == nil {
		return nil, false
	}

	if ll.head == ll.tail {
		value := ll.head.Value
		ll.head = nil
		ll.tail = nil
		ll.size--
		return value, true
	}

	current := ll.head
	for current.Next != ll.tail {
		current = current.Next
	}

	value := ll.tail.Value
	current.Next = nil
	ll.tail = current
	ll.size--

	return value, true
}

// DeleteValue removes the first occurrence of the value
// Returns true if value was found and deleted
// Time Complexity: O(n)
func (ll *LinkedList) DeleteValue(value interface{}) bool {
	if ll.head == nil {
		return false
	}

	if ll.head.Value == value {
		ll.DeleteAtHead()
		return true
	}

	current := ll.head
	for current.Next != nil {
		if current.Next.Value == value {
			// Found the value
			if current.Next == ll.tail {
				ll.tail = current
			}
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}

	return false
}

// Search finds the first occurrence of a value
// Returns true if found
// Time Complexity: O(n)
func (ll *LinkedList) Search(value interface{}) bool {
	current := ll.head

	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}

	return false
}

// Get returns the value at the specified position
// Returns nil and false if position is invalid
// Time Complexity: O(n)
func (ll *LinkedList) Get(position int) (interface{}, bool) {
	if position < 0 || position >= ll.size {
		return nil, false
	}

	current := ll.head
	for i := 0; i < position; i++ {
		current = current.Next
	}

	return current.Value, true
}

// Reverse reverses the linked list in place
// Time Complexity: O(n)
// Space Complexity: O(1)
func (ll *LinkedList) Reverse() {
	if ll.head == nil || ll.head.Next == nil {
		return
	}

	var prev *Node
	current := ll.head
	ll.tail = ll.head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.head = prev
}

// ToSlice converts the linked list to a slice
// Time Complexity: O(n)
func (ll *LinkedList) ToSlice() []interface{} {
	result := make([]interface{}, 0, ll.size)
	current := ll.head

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// IsEmpty returns true if the list has no nodes
func (ll *LinkedList) IsEmpty() bool {
	return ll.head == nil
}

// Size returns the number of nodes in the list
func (ll *LinkedList) Size() int {
	return ll.size
}

// Clear removes all nodes from the list
func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}
