package memory

// Why interviewers ask this:
// sync.Pool is a powerful tool for reducing GC pressure and improving performance in
// high-throughput applications. Interviewers want to see if you understand when to use
// object pooling, how sync.Pool works, and its limitations. This is important for
// writing high-performance Go services.

// Common pitfalls:
// - Using sync.Pool for objects that should have deterministic lifecycle
// - Storing pointers to pool objects beyond their use
// - Not resetting pooled objects before returning them
// - Using sync.Pool when allocation cost is negligible
// - Assuming pooled objects persist (they can be GC'd at any time)

// Key takeaway:
// sync.Pool is for temporary objects that are frequently allocated and deallocated.
// It reduces GC pressure by reusing objects. Objects in pool can be removed by GC at
// any time. Always reset objects before returning to pool. Use for high-frequency
// allocations where profiling shows allocation overhead. Not for connection pools or
// objects with important state.

import (
	"bytes"
	"sync"
)

// WithoutPool demonstrates allocation without pooling
func WithoutPool(n int) [][]byte {
	results := make([][]byte, n)

	for i := 0; i < n; i++ {
		// Allocate new buffer each time
		buf := make([]byte, 1024)
		buf[0] = byte(i)
		results[i] = buf
	}

	return results
}

// WithPool demonstrates using sync.Pool
func WithPool(n int) [][]byte {
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	results := make([][]byte, n)

	for i := 0; i < n; i++ {
		// Get from pool (or allocate if pool is empty)
		buf := pool.Get().([]byte)
		buf[0] = byte(i)
		results[i] = buf

		// Return to pool for reuse
		pool.Put(buf)
	}

	return results
}

// BufferPool demonstrates a common use case: pooling buffers
var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// GetBuffer gets a buffer from the pool
func GetBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

// PutBuffer returns a buffer to the pool after resetting it
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset() // Important: reset before returning to pool
	bufferPool.Put(buf)
}

// UseBufferPool demonstrates using the buffer pool
func UseBufferPool(data string) string {
	buf := GetBuffer()
	defer PutBuffer(buf)

	buf.WriteString(data)
	buf.WriteString(" processed")

	return buf.String()
}

// StructPool demonstrates pooling structs
type LargeStruct struct {
	Data  [1024]byte
	Count int
	Name  string
}

var structPool = sync.Pool{
	New: func() interface{} {
		return &LargeStruct{}
	},
}

// GetStruct gets a struct from the pool
func GetStruct() *LargeStruct {
	return structPool.Get().(*LargeStruct)
}

// PutStruct returns a struct to the pool after resetting
func PutStruct(s *LargeStruct) {
	// Reset fields
	s.Count = 0
	s.Name = ""
	// Note: Data array doesn't need explicit reset for this use case

	structPool.Put(s)
}

// ProcessWithStructPool demonstrates using struct pool
func ProcessWithStructPool(count int, name string) int {
	s := GetStruct()
	defer PutStruct(s)

	s.Count = count
	s.Name = name
	s.Data[0] = byte(count)

	return s.Count * 2
}

// PoolWithCustomReset shows proper reset pattern
type CustomObject struct {
	mu    sync.Mutex
	items []int
	name  string
}

func (c *CustomObject) Reset() {
	c.items = c.items[:0] // Keep capacity, reset length
	c.name = ""
}

var customPool = sync.Pool{
	New: func() interface{} {
		return &CustomObject{
			items: make([]int, 0, 100),
		}
	},
}

// GetCustomObject gets from pool
func GetCustomObject() *CustomObject {
	return customPool.Get().(*CustomObject)
}

// PutCustomObject returns to pool
func PutCustomObject(obj *CustomObject) {
	obj.Reset()
	customPool.Put(obj)
}

// WhenToUsePool demonstrates decision criteria
func WhenToUsePool(allocSize int, frequency int) bool {
	// Use pool when:
	// 1. Objects are allocated frequently
	// 2. Objects are short-lived
	// 3. Allocation cost is significant
	// 4. Profiling shows allocation overhead

	return allocSize > 1024 && frequency > 1000
}

// WhenNotToUsePool shows cases where pool is inappropriate
func WhenNotToUsePool() {
	// Don't use pool for:
	// - Connection pools (use custom pool with lifecycle management)
	// - Objects with important state that must persist
	// - Rarely allocated objects
	// - Small objects (allocation is cheap)
	// - Objects that need deterministic cleanup
}

// PoolLifecycle demonstrates that pool objects can be GC'd
func PoolLifecycle() {
	pool := &sync.Pool{
		New: func() interface{} {
			return &struct{ value int }{value: 42}
		},
	}

	// Put object in pool
	obj := pool.New()
	pool.Put(obj)

	// After GC, object might be gone
	// sync.Pool doesn't guarantee objects persist

	// Get might return nil or new object
	retrieved := pool.Get()
	_ = retrieved
}

// ConcurrentPoolUsage shows pool is safe for concurrent use
func ConcurrentPoolUsage(n int) {
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			buf := pool.Get().([]byte)
			// Use buf
			pool.Put(buf)
		}()
	}

	wg.Wait()
}
