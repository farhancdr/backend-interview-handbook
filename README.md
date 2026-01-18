# 🦄 The Ultimate Golang Interview Handbook

> **Executable. Comprehensive. Zero Dependencies.**
>
> Stop reading static tutorials. Start running code.

![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?style=flat-square&logo=go)
![Tests](https://img.shields.io/badge/tests-passing-success?style=flat-square)
![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)
![Dependencies](https://img.shields.io/badge/dependencies-zero-orange?style=flat-square)

A hands-on, battle-tested repository for backend engineers preparing for Go interviews. From **Goroutine leaks** to **System Design patterns**, everything is implemented using only the standard library.

---

## 🚀 Why This Handbook?

Most interview prep materials are static text. This repository is **executable code**.

*   **Don't just read about Race Conditions** -> Run `go test -race` and see them crash.
*   **Don't just memorize "Context"** -> Run a test that cancels a 50ms database call.
*   **Don't just theory-craft a Rate Limiter** -> Run the Token Bucket implementation.

Every file contains:
1.  🧠 **Why interviewers ask this**
2.  ⚠️ **Common Pitfalls**
3.  🔑 **Key Takeaways**
4.  🧪 **Runnable Tests**

---

## 🗺️ The Map

We've organized the chaos into 10 structured domains.

### 🏗️ Phase 1: The Foundations
| Package | Description | Key Concepts |
|:---|:---|:---|
| **[`internal/basics`](internal/basics)** | The "Must-Knows" | Value/Ref semantics, Slices vs Arrays, Maps |
| **[`internal/intermediate`](internal/intermediate)** | Idiomatic Go | Defer, Panic/Recover, Error Wrapping, Interfaces |
| **[`internal/advanced`](internal/advanced)** | Power User | Context Propagation, Generics, Reflection |

### ⚡ Phase 2: Concurrency & Runtime (The Hard Stuff)
| Package | Description | Key Concepts |
|:---|:---|:---|
| **[`internal/concurrency`](internal/concurrency)** | **🔥 Most Critical** | Channels, Worker Pools, Select, Mutex, Atomic |
| **[`internal/memory`](internal/memory)** | Performance | Escape Analysis, Slice Growth, GC Tuning Patterns |
| **[`internal/internals`](internal/internals)** | Under the Hood | GMP Scheduler, Garbage Collector usage |

### 🛠️ Phase 3: Real-World Engineering
| Package | Description | Key Concepts |
|:---|:---|:---|
| **[`internal/patterns`](internal/patterns)** | Architecture | Repository, Middleware, Functional Options, DI |
| **[`internal/system_design`](internal/system_design)** | Systems | Rate Limiter, Cache, Pub-Sub, Idempotency, Circuit Breaker |

### 🧠 Phase 4: CS Fundamentals
| Package | Description | Key Concepts |
|:---|:---|:---|
| **[`internal/ds`](internal/ds)** | Data Structures | LRU Cache, Heap, BST, Linked List, Stack/Queue |
| **[`internal/algo`](internal/algo)** | Algorithms | Sliding Window, Backtracking, DFS/BFS, Sorting |
| **[`internal/leetcode`](internal/leetcode)** | **Sandbox** | Two Sum, Valid Parentheses, Stock Best Time (100% Executable) |

---

## 🛠️ Quick Start

### 1. Clone & Run
```bash
git clone https://github.com/farhan/golang-interview-handbook.git
cd golang-interview-handbook

# Run EVERYTHING (Success means you're ready!)
go test ./...
```

### 2. Spot Check Concurrency
This is where 80% of candidates fail. Verify you understand it:
```bash
# Run with Race Detector
go test -race -v ./internal/concurrency/
```

### 3. Benchmark Memory
See the difference between pre-allocating slices vs appending blindly:
```bash
go test -bench=. ./internal/memory/
```

### 4. Practice LeetCode
Test your solutions against our pre-written test cases:
```bash
go test -v ./internal/leetcode/
```

---

## 🎓 How to Study

Each file is a self-contained lesson.

1.  **Open** `internal/concurrency/worker_pool.go`
2.  **Read** the "Why interviewers ask this" comment block.
3.  **Run** the test: `go test -v ./internal/concurrency/ -run TestWorkerPool`
4.  **Break it**: Remove the `wg.Wait()` or `close(results)`.
5.  **Observe**: Watch the test panic or hang. Now you know *why* that line exists.

### Example: Rate Limiter (System Design)
Go to `internal/system_design/rate_limiter.go`. You'll see a complete Token Bucket implementation using just `sync.Mutex` and `time.Now()`.

```go
// Real code snippet from the repo:
func (rl *RateLimiter) AllowN(n float64) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    rl.refill()
    if rl.tokens >= n {
        rl.tokens -= n
        return true
    }
    return false
}
```

---

## 🏆 Checklist for Mastery

- [ ] Can you implement a **Thread-Safe LRU Cache**? (See `internal/ds/lru_cache.go`)
- [ ] Do you know how to stop a **Goroutine leak**? (See `internal/memory/goroutine_leaks.go`)
- [ ] Can you implement **Graceful Shutdown**? (See `internal/patterns/server.go`)
- [ ] Do you know **Slice Capacity** growth rules? (See `internal/memory/slice_capacity.go`)
- [ ] Can you write a **Worker Pool** from scratch? (See `internal/concurrency/worker_pool.go`)
- [ ] Can you solve **Two Sum** in O(n) without Googling? (See `internal/leetcode/two_sum.go`)

---

## 🤝 Contributing

Found a bug? Want to add a graph algorithm?
PRs are welcome! Please ensure executing `go test ./...` passes.

---

**Made with ❤️ for the Go Community.**
*Run code, not mouth.*
