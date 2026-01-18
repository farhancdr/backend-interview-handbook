package ds

// Why interviewers ask this:
// LRU Cache is a classic system design problem that tests understanding of multiple data structures
// (hash map + doubly linked list), time complexity optimization, and cache eviction policies.
// It's commonly used in real systems (browser cache, database cache, CDN) and demonstrates
// ability to combine data structures for O(1) operations.

// Common pitfalls:
// - Using single data structure (either map or list alone) leading to O(n) operations
// - Not properly updating both map and list on every operation
// - Incorrect doubly linked list manipulation (losing references)
// - Forgetting to update head/tail pointers
// - Not handling edge cases (capacity 1, empty cache)

// Key takeaway:
// LRU Cache requires O(1) get and put operations. Achieve this by combining:
// 1. Hash map for O(1) key lookup
// 2. Doubly linked list for O(1) removal and insertion (maintains access order)
// Most recently used at head, least recently used at tail. Evict from tail when capacity reached.

// LRUNode represents a node in the doubly linked list
type LRUNode struct {
	Key   string
	Value interface{}
	Prev  *LRUNode
	Next  *LRUNode
}

// LRUCache implements a Least Recently Used cache
// Time Complexity: Get O(1), Put O(1)
// Space Complexity: O(capacity)
type LRUCache struct {
	capacity int
	cache    map[string]*LRUNode
	head     *LRUNode // Most recently used
	tail     *LRUNode // Least recently used
}

// NewLRUCache creates a new LRU cache with given capacity
func NewLRUCache(capacity int) *LRUCache {
	if capacity < 1 {
		capacity = 1
	}

	// Create dummy head and tail nodes
	head := &LRUNode{}
	tail := &LRUNode{}
	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*LRUNode),
		head:     head,
		tail:     tail,
	}
}

// Get retrieves a value from the cache
// Returns nil and false if key doesn't exist
// Moves accessed item to front (most recently used)
// Time Complexity: O(1)
func (lru *LRUCache) Get(key string) (interface{}, bool) {
	node, exists := lru.cache[key]
	if !exists {
		return nil, false
	}

	// Move to front (most recently used)
	lru.moveToFront(node)

	return node.Value, true
}

// Put adds or updates a key-value pair
// If key exists, updates value and moves to front
// If cache is at capacity, evicts least recently used item
// Time Complexity: O(1)
func (lru *LRUCache) Put(key string, value interface{}) {
	// Check if key already exists
	if node, exists := lru.cache[key]; exists {
		node.Value = value
		lru.moveToFront(node)
		return
	}

	// Create new node
	newNode := &LRUNode{
		Key:   key,
		Value: value,
	}

	// Add to cache and front of list
	lru.cache[key] = newNode
	lru.addToFront(newNode)

	// Check capacity and evict if necessary
	if len(lru.cache) > lru.capacity {
		lru.evictLRU()
	}
}

// Delete removes a key from the cache
// Returns true if key was found and deleted
// Time Complexity: O(1)
func (lru *LRUCache) Delete(key string) bool {
	node, exists := lru.cache[key]
	if !exists {
		return false
	}

	lru.removeNode(node)
	delete(lru.cache, key)

	return true
}

// Size returns the current number of items in cache
func (lru *LRUCache) Size() int {
	return len(lru.cache)
}

// Capacity returns the maximum capacity of the cache
func (lru *LRUCache) Capacity() int {
	return lru.capacity
}

// Clear removes all items from the cache
func (lru *LRUCache) Clear() {
	lru.cache = make(map[string]*LRUNode)
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
}

// Keys returns all keys in the cache (in no particular order)
func (lru *LRUCache) Keys() []string {
	keys := make([]string, 0, len(lru.cache))
	for k := range lru.cache {
		keys = append(keys, k)
	}
	return keys
}

// moveToFront moves a node to the front of the list (most recently used)
func (lru *LRUCache) moveToFront(node *LRUNode) {
	lru.removeNode(node)
	lru.addToFront(node)
}

// addToFront adds a node to the front of the list
func (lru *LRUCache) addToFront(node *LRUNode) {
	node.Next = lru.head.Next
	node.Prev = lru.head
	lru.head.Next.Prev = node
	lru.head.Next = node
}

// removeNode removes a node from the list
func (lru *LRUCache) removeNode(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// evictLRU removes the least recently used item (tail)
func (lru *LRUCache) evictLRU() {
	lruNode := lru.tail.Prev
	if lruNode == lru.head {
		return // Empty list
	}

	lru.removeNode(lruNode)
	delete(lru.cache, lruNode.Key)
}

// GetOldest returns the least recently used key without removing it
// Returns empty string and false if cache is empty
func (lru *LRUCache) GetOldest() (string, bool) {
	if lru.tail.Prev == lru.head {
		return "", false
	}

	return lru.tail.Prev.Key, true
}

// GetNewest returns the most recently used key without removing it
// Returns empty string and false if cache is empty
func (lru *LRUCache) GetNewest() (string, bool) {
	if lru.head.Next == lru.tail {
		return "", false
	}

	return lru.head.Next.Key, true
}
