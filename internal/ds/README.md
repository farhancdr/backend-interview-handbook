# 📊 Data Structures

> **Fundamental data structures implemented in Go**

This package contains Go implementations of essential data structures. Understanding these is crucial for both coding interviews and building efficient systems.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **LRU Cache** | [lru_cache.go](lru_cache.go) | Least Recently Used eviction, O(1) operations, doubly linked list + hash map |
| **Heap** | [heap.go](heap.go) | Min/max heap, priority queue, heapify, O(log n) operations |
| **Binary Search Tree** | [bst.go](bst.go) | BST properties, insert, delete, search, in-order traversal |
| **Binary Tree** | [binary_tree.go](binary_tree.go) | Tree traversals (pre/in/post-order), DFS, BFS, height, diameter |
| **Linked List** | [linked_list.go](linked_list.go) | Singly linked list, insert, delete, reverse, detect cycle |
| **Stack** | [stack.go](stack.go) | LIFO, push, pop, peek, applications |
| **Queue** | [queue.go](queue.go) | FIFO, enqueue, dequeue, circular queue |
| **HashMap** | [hashmap.go](hashmap.go) | Hash function, collision resolution, load factor |

---

## 🚀 Quick Start

```bash
# Run all tests
go test -v ./ds/

# Run specific data structure
go test -v ./ds/ -run TestLRUCache

# Check test coverage
go test -cover ./ds/
```

---

## 🎓 Learning Guide

### Recommended Order
1. **Stack & Queue** - Simplest structures
2. **Linked List** - Pointer manipulation
3. **Binary Tree** - Tree fundamentals
4. **Binary Search Tree** - Ordered tree
5. **Heap** - Complete binary tree
6. **HashMap** - Hash table implementation
7. **LRU Cache** - Combines linked list + hash map

### Common Interview Questions

**LRU Cache:**
- "Implement an LRU cache with O(1) get and put"
- "How do you make LRU cache thread-safe?"
- "Explain the data structures used in LRU"

**Heap:**
- "Find the Kth largest element"
- "Merge K sorted lists"
- "Implement a priority queue"

**Binary Tree:**
- "Find the lowest common ancestor"
- "Check if a tree is balanced"
- "Serialize and deserialize a binary tree"

**Linked List:**
- "Reverse a linked list"
- "Detect a cycle in a linked list"
- "Find the middle of a linked list"

---

## 💡 Key Takeaways

### LRU Cache
```go
// Data structures:
// 1. Doubly linked list (for LRU ordering)
// 2. Hash map (for O(1) lookup)

// Operations:
// - Get: O(1) - Move to front
// - Put: O(1) - Add to front, evict if full

// Real-world use:
// - Browser cache
// - Database query cache
// - CDN cache
```

### Heap
```go
// Properties:
// - Complete binary tree
// - Parent >= children (max heap)
// - Parent <= children (min heap)

// Operations:
// - Insert: O(log n)
// - Delete: O(log n)
// - Peek: O(1)

// Use cases:
// - Priority queue
// - Top K problems
// - Median finding
```

### Binary Search Tree
```go
// Properties:
// - Left subtree < root
// - Right subtree > root
// - In-order traversal gives sorted order

// Operations:
// - Search: O(log n) average, O(n) worst
// - Insert: O(log n) average, O(n) worst
// - Delete: O(log n) average, O(n) worst

// Balanced BST (AVL, Red-Black):
// - Guarantees O(log n) operations
```

### HashMap
```go
// Components:
// - Hash function
// - Buckets (array)
// - Collision resolution (chaining or open addressing)

// Load factor = n / capacity
// Rehash when load factor > 0.75

// Go's map:
// - Built-in hash map
// - NOT thread-safe
// - Random iteration order
```

---

## 🧪 Time Complexity Cheat Sheet

| Data Structure | Access | Search | Insert | Delete |
|:---------------|:-------|:-------|:-------|:-------|
| Array | O(1) | O(n) | O(n) | O(n) |
| Linked List | O(n) | O(n) | O(1)* | O(1)* |
| Stack | O(n) | O(n) | O(1) | O(1) |
| Queue | O(n) | O(n) | O(1) | O(1) |
| Hash Map | - | O(1)† | O(1)† | O(1)† |
| BST | O(log n)† | O(log n)† | O(log n)† | O(log n)† |
| Heap | - | O(n) | O(log n) | O(log n) |

\* If you have the pointer to the node  
† Average case, worst case can be O(n)

---

## ⚠️ Common Pitfalls

1. **Linked List**
   - Forgetting to update head/tail pointers
   - Not handling edge cases (empty list, single node)
   - Memory leaks (not breaking cycles)

2. **Binary Tree**
   - Confusing pre/in/post-order traversals
   - Not handling nil nodes
   - Stack overflow with deep recursion

3. **Heap**
   - Forgetting to heapify after insert/delete
   - Using wrong comparison for min/max heap
   - Not maintaining complete tree property

4. **LRU Cache**
   - Not updating order on get()
   - Memory leaks (not removing from map)
   - Race conditions in concurrent access

---

## 🔗 Related Topics

- **[Algorithms](../algo/)** - Algorithms that use these data structures
- **[LeetCode](../leetcode/)** - Practice problems
- **[System Design](../system_design/)** - LRU cache in production systems

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
