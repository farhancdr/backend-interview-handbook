# 🧠 Memory Optimization

> **Write performant Go code by understanding memory**

This package covers memory management patterns that separate good Go developers from great ones. Learn to write code that's both correct and fast.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **Slice Capacity** | [slice_capacity.go](slice_capacity.go) | Growth algorithm, pre-allocation, capacity vs length |
| **Goroutine Leaks** | [goroutine_leaks.go](goroutine_leaks.go) | Detecting leaks, context cancellation, cleanup patterns |
| **Map Allocation** | [map_allocation.go](map_allocation.go) | Pre-sizing maps, load factor, memory overhead |
| **Object Pooling** | [object_pooling.go](object_pooling.go) | sync.Pool, reducing GC pressure, when to pool |
| **String Immutability** | [string_immutability.go](string_immutability.go) | String internals, []byte conversion, string builder |
| **Benchmarking Basics** | [benchmarking_basics.go](benchmarking_basics_test.go) | Writing benchmarks, interpreting results, profiling |

---

## 🚀 Quick Start

```bash
# Run benchmarks
go test -bench=. ./memory/

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem ./memory/

# Run specific benchmark
go test -bench=BenchmarkSlicePrealloc ./memory/

# Profile memory
go test -bench=. -memprofile=mem.out ./memory/
go tool pprof mem.out
```

---

## 🎓 Learning Guide

### Recommended Order
1. **Slice Capacity** - Most common performance issue
2. **Map Allocation** - Pre-sizing for performance
3. **String Immutability** - Efficient string operations
4. **Object Pooling** - Reduce GC pressure
5. **Goroutine Leaks** - Prevent memory leaks
6. **Benchmarking Basics** - Measure everything

### Common Interview Questions

**Slice Performance:**
- "How does slice capacity grow?"
- "When should you pre-allocate slices?"
- "What's the performance difference between append and pre-allocation?"

**Memory Management:**
- "How does Go's garbage collector work?"
- "What is escape analysis?"
- "When should you use sync.Pool?"

**Goroutine Leaks:**
- "How do you detect goroutine leaks?"
- "What causes goroutines to leak?"
- "How do you prevent leaks with context?"

---

## 💡 Key Takeaways

### Slice Capacity
```go
// ❌ Slow - multiple allocations
s := []int{}
for i := 0; i < 1000; i++ {
    s = append(s, i)
}

// ✅ Fast - single allocation
s := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    s = append(s, i)
}
```

**Growth Rule:** Capacity doubles until 1024, then grows by 25%

### Map Pre-sizing
```go
// ❌ Slow - multiple rehashes
m := make(map[string]int)

// ✅ Fast - pre-sized
m := make(map[string]int, 1000)
```

### String Operations
```go
// ❌ Slow - creates many intermediate strings
s := ""
for i := 0; i < 1000; i++ {
    s += "x"  // O(n²) complexity!
}

// ✅ Fast - uses strings.Builder
var sb strings.Builder
sb.Grow(1000)  // Pre-allocate
for i := 0; i < 1000; i++ {
    sb.WriteString("x")
}
s := sb.String()
```

### Object Pooling
```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

// Get from pool
buf := bufferPool.Get().(*bytes.Buffer)
defer bufferPool.Put(buf)
buf.Reset()  // CRITICAL: Reset before reuse
```

**When to Use:**
- High allocation rate (millions per second)
- Objects are expensive to create
- GC pressure is measurable problem

**When NOT to Use:**
- Premature optimization
- Objects hold resources (file handles, connections)
- Memory usage is more important than CPU

### Goroutine Leaks
```go
// ❌ Leaks - goroutine waits forever
func leak() {
    ch := make(chan int)
    go func() {
        <-ch  // Blocks forever if nothing sends
    }()
}

// ✅ Fixed - use context for cancellation
func noLeak(ctx context.Context) {
    ch := make(chan int)
    go func() {
        select {
        case <-ch:
        case <-ctx.Done():
            return
        }
    }()
}
```

---

## 🧪 Benchmarking

### Writing Benchmarks
```go
func BenchmarkExample(b *testing.B) {
    // Setup (not timed)
    data := make([]int, 1000)
    
    b.ResetTimer()  // Reset timer after setup
    
    for i := 0; i < b.N; i++ {
        // Code to benchmark
        process(data)
    }
}
```

### Interpreting Results
```
BenchmarkSliceAppend-8    1000000    1234 ns/op    2048 B/op    5 allocs/op
                          ^^^^^^^^   ^^^^^^^^^^    ^^^^^^^^^    ^^^^^^^^^^^^
                          iterations  time/op      bytes/op     allocations/op
```

---

## ⚠️ Common Pitfalls

1. **Premature Optimization** - Measure first, optimize second
2. **Ignoring Escape Analysis** - Variables escape to heap unexpectedly
3. **Over-using sync.Pool** - Not always faster, adds complexity
4. **Forgetting buf.Reset()** - Pooled objects must be reset
5. **Not Benchmarking** - Assumptions about performance are often wrong

---

## 🔗 Related Topics

- **[Concurrency](../concurrency/)** - Goroutine leaks prevention
- **[Internals](../internals/)** - GC and scheduler internals
- **[Basics](../basics/)** - Slice and map fundamentals

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
