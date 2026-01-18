package concurrency

import "sync"

// Why interviewers ask this:
// Mutexes are essential for protecting shared state in concurrent programs.
// Understanding when to use Mutex vs RWMutex, deadlock prevention, and proper
// locking patterns is critical. This is frequently asked in interviews.

// Common pitfalls:
// - Forgetting to unlock (use defer!)
// - Copying mutexes (they should never be copied)
// - Deadlocks from incorrect lock ordering
// - Not using RWMutex when you have many readers
// - Locking for too long (holding locks during I/O)

// Key takeaway:
// Mutex: exclusive access (one goroutine at a time)
// RWMutex: multiple readers OR one writer
// Always use defer mu.Unlock() immediately after locking

// Counter demonstrates mutex usage for protecting shared state
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value safely reads the counter value
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// RWCounter demonstrates RWMutex usage
type RWCounter struct {
	mu    sync.RWMutex
	value int
}

// Increment uses write lock
func (c *RWCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value uses read lock (multiple readers allowed)
func (c *RWCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

// SafeMap demonstrates mutex-protected map
type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

// NewSafeMap creates a new SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set safely sets a key-value pair
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get safely gets a value
func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.data[key]
	return val, ok
}

// ConcurrentIncrement demonstrates race condition without mutex
func ConcurrentIncrement(n int) int {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // RACE CONDITION!
		}()
	}

	wg.Wait()
	return counter // Will likely be less than n
}

// SafeConcurrentIncrement demonstrates proper mutex usage
func SafeConcurrentIncrement(n int) int {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	return counter // Will be exactly n
}

// MutexWithDefer demonstrates proper defer usage
func MutexWithDefer() int {
	var mu sync.Mutex
	value := 0

	mu.Lock()
	defer mu.Unlock() // Ensures unlock even if panic

	value = 42
	return value
}

// RWMutexPerformance demonstrates RWMutex advantage with many readers
func RWMutexPerformance(readers int) []int {
	rwc := &RWCounter{value: 100}
	var wg sync.WaitGroup
	results := make([]int, readers)

	// Many concurrent readers
	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			results[idx] = rwc.Value() // Read lock (concurrent)
		}(i)
	}

	wg.Wait()
	return results
}

// DeadlockExample demonstrates potential deadlock (commented out)
func DeadlockExample() {
	// var mu1, mu2 sync.Mutex

	// Goroutine 1: locks mu1 then mu2
	// go func() {
	// 	mu1.Lock()
	// 	mu2.Lock()
	// 	mu2.Unlock()
	// 	mu1.Unlock()
	// }()

	// Goroutine 2: locks mu2 then mu1 (DEADLOCK!)
	// go func() {
	// 	mu2.Lock()
	// 	mu1.Lock()
	// 	mu1.Unlock()
	// 	mu2.Unlock()
	// }()
}

// MutexZeroValue demonstrates that zero value mutex is ready to use
func MutexZeroValue() int {
	var mu sync.Mutex // Zero value is ready to use
	value := 0

	mu.Lock()
	value = 42
	mu.Unlock()

	return value
}
