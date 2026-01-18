# 🔧 Internal Packages

> **Executable Go code for every backend interview topic**

This directory contains runnable Go implementations of everything from basic language features to advanced system design patterns. Every file is tested, documented, and ready to run.

---

## 🗺️ Learning Path

### 🏗️ Phase 1: The Foundations

Master Go fundamentals before diving into advanced topics.

| Package | Description | Topics Covered |
|:--------|:------------|:---------------|
| **[basics/](basics/)** | The "Must-Knows" | Arrays/Slices, Maps, Structs, Interfaces, Value vs Reference, Errors |
| **[intermediate/](intermediate/)** | Idiomatic Go | Defer, Panic/Recover, Error Wrapping, Interface Composition |
| **[advanced/](advanced/)** | Power User Features | Context, Generics, Reflection, Unsafe, Memory Alignment, Functional Options |

---

### ⚡ Phase 2: Concurrency & Runtime

**This is where 80% of candidates fail.** Master these topics.

| Package | Description | Topics Covered |
|:--------|:------------|:---------------|
| **[concurrency/](concurrency/)** | **🔥 Most Critical** | Channels, Goroutines, Mutex, Worker Pools, Select |
| **[memory/](memory/)** | Performance Optimization | Escape Analysis, Slice Growth, GC Patterns, Object Pooling, Goroutine Leaks |
| **[internals/](internals/)** | Under the Hood | GMP Scheduler, Garbage Collector Internals |

---

### 🛠️ Phase 3: Real-World Engineering

Apply your knowledge to production patterns.

| Package | Description | Topics Covered |
|:--------|:------------|:---------------|
| **[patterns/](patterns/)** | Design Patterns | Repository, Middleware, Functional Options, DI, Circuit Breaker, Retry, Service Layer |
| **[system_design/](system_design/)** | System Primitives | Rate Limiter, Cache, Pub-Sub, Idempotency, Pagination, Orchestrator |

---

### 🧠 Phase 4: CS Fundamentals

Data structures and algorithms implemented in Go.

| Package | Description | Topics Covered |
|:--------|:------------|:---------------|
| **[ds/](ds/)** | Data Structures | LRU Cache, Heap, BST, Binary Tree, Linked List, Stack, Queue, HashMap |
| **[algo/](algo/)** | Algorithms | Binary Search, Sliding Window, Two Pointers, Sorting, Dynamic Programming |
| **[leetcode/](leetcode/)** | **Practice Sandbox** | Two Sum, Valid Parentheses, Stock Best Time, Merge Lists |

---

## 🚀 Quick Start

### Run Everything
```bash
# Run all tests
go test ./...

# Run with race detector (CRITICAL for concurrency)
go test -race ./...

# Run benchmarks
go test -bench=. ./...
```

### Focus on Specific Topics

```bash
# Master concurrency (most important!)
go test -race -v ./concurrency/

# Understand memory optimization
go test -bench=. ./memory/

# Practice data structures
go test -v ./ds/

# System design patterns
go test -v ./system_design/
```

---

## 📚 How to Study

Each package follows the same structure:

1. **Read the code** - Every file has detailed comments explaining "Why interviewers ask this"
2. **Run the tests** - See the code in action
3. **Break it** - Remove a line, see what fails, understand why it was needed
4. **Modify it** - Add your own test cases

### Example Learning Flow

```
1. Open: internal/concurrency/worker_pool.go
2. Read: The "Why this matters" comment block
3. Run: go test -v ./concurrency/ -run TestWorkerPool
4. Break: Remove the wg.Wait() line
5. Observe: Test hangs - now you understand why WaitGroup is needed
```

---

## 🎯 Interview Preparation Checklist

### Must-Master Topics

- [ ] **Goroutines & Channels** - Can you implement a worker pool from scratch?
- [ ] **Race Conditions** - Can you identify and fix data races?
- [ ] **Context** - Do you know when to use context.WithTimeout vs context.WithCancel?
- [ ] **Interfaces** - Can you explain empty interface vs type assertions?
- [ ] **Slice Internals** - Do you understand capacity vs length?
- [ ] **Error Handling** - Can you implement custom errors with wrapping?
- [ ] **LRU Cache** - Can you build a thread-safe LRU cache?
- [ ] **Rate Limiter** - Can you implement token bucket algorithm?

### Common Interview Questions

**Concurrency:**
- "Implement a worker pool that processes jobs concurrently"
- "How do you prevent goroutine leaks?"
- "Explain the difference between buffered and unbuffered channels"

**Memory:**
- "When does a variable escape to the heap?"
- "How does Go's garbage collector work?"
- "What's the difference between make() and new()?"

**System Design:**
- "Implement a rate limiter"
- "Design a cache with TTL support"
- "Implement graceful shutdown for a web server"

---

## 🔗 Related Sections

- **[Fundamentals](../fundamentals/)** - Theory behind OS, Networking, and Databases
- **[Main README](../README.md)** - Project overview and getting started

---

## 💡 Pro Tips

1. **Start with basics** - Don't jump to advanced topics without mastering fundamentals
2. **Run with -race** - Always test concurrent code with the race detector
3. **Benchmark everything** - Use benchmarks to understand performance implications
4. **Read the tests** - Tests show you how to use the code
5. **Connect the dots** - Link Go concepts to OS fundamentals (goroutines → threads, GC → memory management)

---

**Made with ❤️ for the Go Community.**
*Run code, not mouth.*
