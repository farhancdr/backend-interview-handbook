package ds

// Why interviewers ask this:
// Hash maps demonstrate understanding of hashing, collision resolution, and amortized time complexity.
// They're fundamental to many algorithms and system design problems. Interviewers want to see if you
// understand how hash tables work under the hood and can implement collision resolution strategies.

// Common pitfalls:
// - Poor hash function leading to many collisions
// - Not handling collisions properly
// - Forgetting to resize when load factor is high
// - Not understanding amortized O(1) vs worst-case O(n)
// - Integer overflow in hash function

// Key takeaway:
// Hash map provides O(1) average case for insert/search/delete using hashing and collision resolution.
// Common collision strategies: chaining (linked lists) or open addressing (linear/quadratic probing).
// Load factor determines when to resize. Good hash function distributes keys uniformly.

// HashMapEntry represents a key-value pair in the hash map
type HashMapEntry struct {
	Key   string
	Value interface{}
	Next  *HashMapEntry // For chaining collision resolution
}

// HashMap represents a simplified hash map using chaining
// Time Complexity: Average O(1), Worst O(n) for insert/search/delete
// Space Complexity: O(n) where n is number of entries
type HashMap struct {
	buckets    []*HashMapEntry
	size       int
	capacity   int
	loadFactor float64
}

// NewHashMap creates a new hash map with initial capacity
func NewHashMap(capacity int) *HashMap {
	if capacity < 1 {
		capacity = 16
	}

	return &HashMap{
		buckets:    make([]*HashMapEntry, capacity),
		size:       0,
		capacity:   capacity,
		loadFactor: 0.75,
	}
}

// hash computes the hash value for a key
func (hm *HashMap) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash*31 + int(key[i])) % hm.capacity
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

// Put inserts or updates a key-value pair
// Time Complexity: O(1) average
func (hm *HashMap) Put(key string, value interface{}) {
	if float64(hm.size)/float64(hm.capacity) > hm.loadFactor {
		hm.resize()
	}

	index := hm.hash(key)

	// Check if key already exists
	current := hm.buckets[index]
	for current != nil {
		if current.Key == key {
			current.Value = value
			return
		}
		current = current.Next
	}

	// Insert new entry at head of chain
	newEntry := &HashMapEntry{
		Key:   key,
		Value: value,
		Next:  hm.buckets[index],
	}
	hm.buckets[index] = newEntry
	hm.size++
}

// Get retrieves the value for a key
// Returns nil and false if key doesn't exist
// Time Complexity: O(1) average
func (hm *HashMap) Get(key string) (interface{}, bool) {
	index := hm.hash(key)
	current := hm.buckets[index]

	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	return nil, false
}

// Delete removes a key-value pair
// Returns true if key was found and deleted
// Time Complexity: O(1) average
func (hm *HashMap) Delete(key string) bool {
	index := hm.hash(key)
	current := hm.buckets[index]
	var prev *HashMapEntry

	for current != nil {
		if current.Key == key {
			if prev == nil {
				// Deleting head of chain
				hm.buckets[index] = current.Next
			} else {
				prev.Next = current.Next
			}
			hm.size--
			return true
		}
		prev = current
		current = current.Next
	}

	return false
}

// Contains checks if a key exists
// Time Complexity: O(1) average
func (hm *HashMap) Contains(key string) bool {
	_, exists := hm.Get(key)
	return exists
}

// Size returns the number of key-value pairs
func (hm *HashMap) Size() int {
	return hm.size
}

// IsEmpty returns true if map has no entries
func (hm *HashMap) IsEmpty() bool {
	return hm.size == 0
}

// Clear removes all entries
func (hm *HashMap) Clear() {
	hm.buckets = make([]*HashMapEntry, hm.capacity)
	hm.size = 0
}

// Keys returns all keys in the map
func (hm *HashMap) Keys() []string {
	keys := make([]string, 0, hm.size)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			keys = append(keys, current.Key)
			current = current.Next
		}
	}

	return keys
}

// resize doubles the capacity and rehashes all entries
func (hm *HashMap) resize() {
	oldBuckets := hm.buckets
	hm.capacity *= 2
	hm.buckets = make([]*HashMapEntry, hm.capacity)
	hm.size = 0

	// Rehash all entries
	for _, bucket := range oldBuckets {
		current := bucket
		for current != nil {
			hm.Put(current.Key, current.Value)
			current = current.Next
		}
	}
}
