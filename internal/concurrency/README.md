# ⚡ Concurrency

> **🔥 The most critical topic for Go interviews**

Concurrency is where Go shines and where most candidates fail. Master these patterns to stand out in interviews.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **Goroutines** | [goroutines.go](goroutines.go) | Lightweight threads, scheduling, goroutine lifecycle, leaks |
| **Channels** | [channels.go](channels.go) | Buffered vs unbuffered, send/receive, close semantics, select |
| **Mutex** | [mutex.go](mutex.go) | Mutual exclusion, RWMutex, critical sections, deadlocks |
| **Worker Pool** | [worker_pool.go](worker_pool.go) | Job distribution, bounded concurrency, graceful shutdown |

---

## 🚀 Quick Start

```bash
# ALWAYS run with race detector for concurrency code
go test -race -v ./concurrency/

# Run specific test
go test -race -v ./concurrency/ -run TestWorkerPool

# Benchmark concurrent performance
go test -bench=. ./concurrency/
```

---

## 🎓 Learning Guide

### Recommended Order
1. **Goroutines** - Understand the basics of concurrent execution
2. **Channels** - Learn Go's primary concurrency primitive
3. **Mutex** - Understand shared memory synchronization
4. **Worker Pool** - Apply everything to a real-world pattern

### Common Interview Questions

**Goroutines:**
- "What's the difference between goroutines and OS threads?"
- "How do you prevent goroutine leaks?"
- "What happens if a goroutine panics?"

**Channels:**
- "What's the difference between buffered and unbuffered channels?"
- "What happens if you send to a closed channel?"
- "How do you implement a timeout with channels?"

**Mutex:**
- "When should you use Mutex vs channels?"
- "What's the difference between Mutex and RWMutex?"
- "How do you prevent deadlocks?"

**Worker Pool:**
- "How do you implement bounded concurrency?"
- "How do you handle graceful shutdown?"
- "What's the difference between fan-out and fan-in?"

---

## 💡 Key Takeaways

### Goroutines
- **Extremely lightweight** - Can spawn millions
- Managed by Go runtime's **GMP scheduler**
- Always ensure goroutines can **exit** (prevent leaks)
- Use `runtime.NumGoroutine()` to detect leaks

### Channels
- **Unbuffered** - Synchronous (sender blocks until receiver ready)
- **Buffered** - Asynchronous (sender blocks only when buffer full)
- **Closing** - Signals "no more values" to receivers
- **Select** - Multiplexing multiple channel operations

### Channel Rules
```go
// ✅ Safe
close(ch)              // Close from sender
v, ok := <-ch          // Check if closed

// ❌ Panic
send to closed channel
close already closed channel
close nil channel
```

### Mutex vs Channels
- **Use Channels** when: Passing ownership, distributing work, signaling events
- **Use Mutex** when: Protecting shared state, simple critical sections, performance-critical

### Worker Pool Pattern
```go
// Key components:
1. Job channel (buffered)
2. Results channel (buffered)
3. WaitGroup for workers
4. Graceful shutdown (close job channel)
```

---

## ⚠️ Common Pitfalls

1. **Goroutine Leaks** - Goroutines waiting forever on channels
2. **Sending to Closed Channel** - Causes panic
3. **Forgetting WaitGroup.Done()** - Causes deadlock
4. **Mutex Deadlock** - Acquiring locks in inconsistent order
5. **Race Conditions** - Accessing shared memory without synchronization

---

## 🧪 Testing Concurrency

### Race Detector
```bash
# ALWAYS use -race flag
go test -race ./concurrency/

# Race detector finds:
# - Concurrent map access
# - Unsynchronized variable access
# - Channel misuse
```

### Detecting Goroutine Leaks
```go
before := runtime.NumGoroutine()
// ... your code ...
after := runtime.NumGoroutine()
if after > before {
    t.Error("goroutine leak detected")
}
```

---

## 🔗 Related Topics

- **[Memory](../memory/)** - Goroutine leaks, escape analysis
- **[Internals](../internals/)** - GMP scheduler deep dive
- **[Patterns](../patterns/)** - Circuit breaker, retry patterns
- **[System Design](../system_design/)** - Worker pools in production systems

---

## 📚 Further Reading

- [Go Concurrency Patterns (Rob Pike)](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
