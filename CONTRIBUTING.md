# Contributing to Golang Interview Handbook

Thank you for your interest in contributing! This guide will help you add new topics while maintaining consistency.

## 🎯 Core Principles

1. **Executable Learning**: Every concept must have runnable tests
2. **Code > Theory**: Minimal prose, maximum clarity via code comments
3. **Isolation**: Each topic in its own file, minimal coupling
4. **Interview-Oriented**: Focus on "why interviewers ask this"
5. **Standard Library First**: Avoid external dependencies

## 📁 File Structure

Each topic requires two files:

```
internal/[package]/
├── [topic_name].go       # Implementation
└── [topic_name]_test.go  # Tests demonstrating behavior
```

### Example: Adding a New Topic

```
internal/concurrency/
├── channels.go
└── channels_test.go
```

## ✍️ Writing Implementation Files

### Required Structure

Every `.go` file must include:

```go
package [packagename]

// Why interviewers ask this:
// [Explain the interview relevance]

// Common pitfalls:
// [List common mistakes]

// Key takeaway:
// [Main concept to remember]

// [Your implementation here]
```

### Example Implementation

```go
package concurrency

// Why interviewers ask this:
// Channels are fundamental to Go's concurrency model. Interviewers want to
// ensure you understand blocking behavior, buffering, and proper channel usage.

// Common pitfalls:
// - Sending to unbuffered channel without receiver causes deadlock
// - Forgetting to close channels can leak goroutines
// - Closing a channel multiple times causes panic

// Key takeaway:
// Unbuffered channels are synchronous (sender blocks until receiver ready).
// Buffered channels are asynchronous up to their capacity.

// SendToUnbuffered demonstrates blocking behavior of unbuffered channels
func SendToUnbuffered() {
    ch := make(chan int) // unbuffered
    
    // This would deadlock if uncommented:
    // ch <- 1 // blocks forever, no receiver
    
    // Correct usage with goroutine:
    go func() {
        ch <- 1
    }()
    
    val := <-ch // receives the value
    _ = val
}
```

## 🧪 Writing Test Files

### Required Structure

```go
package [packagename]

import "testing"

func Test[TopicName]_[Behavior](t *testing.T) {
    // Arrange
    // Act
    // Assert
}
```

### Test Naming Convention

Use descriptive names that explain the behavior:

✅ Good:
- `TestChannel_BufferedVsUnbuffered`
- `TestDefer_ExecutionOrder`
- `TestInterface_NilInterfaceVsNilValue`

❌ Bad:
- `TestChannel`
- `TestDefer`
- `TestInterface`

### Example Test

```go
package concurrency

import (
    "testing"
    "time"
)

func TestChannel_UnbufferedBlocking(t *testing.T) {
    ch := make(chan int)
    
    done := make(chan bool)
    
    // Start receiver in goroutine
    go func() {
        val := <-ch
        if val != 42 {
            t.Errorf("expected 42, got %d", val)
        }
        done <- true
    }()
    
    // Send value (blocks until receiver ready)
    ch <- 42
    
    // Wait for completion
    <-done
}

func TestChannel_BufferedNonBlocking(t *testing.T) {
    ch := make(chan int, 1) // buffered with capacity 1
    
    // This doesn't block because buffer has space
    ch <- 42
    
    // Receive the value
    val := <-ch
    if val != 42 {
        t.Errorf("expected 42, got %d", val)
    }
}
```

## 📊 Complexity Analysis (for Algorithms & Data Structures)

Include time and space complexity in comments:

```go
// BinarySearch performs binary search on a sorted slice
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearch(arr []int, target int) int {
    // implementation
}
```

## 🎨 Code Style

### General Guidelines

1. **Use standard Go formatting**: Run `go fmt` before committing
2. **Clear variable names**: Prefer clarity over brevity
3. **Comment non-obvious code**: Explain "why", not "what"
4. **Keep functions small**: One concept per function
5. **Avoid global state**: Each test should be independent

### Comments Style

```go
// Good: Explains why
// We use a buffered channel here to prevent goroutine blocking
// when the receiver might not be ready immediately

// Bad: States the obvious
// Create a buffered channel with capacity 10
```

## ✅ Testing Requirements

### All Tests Must:

1. **Pass independently**: `go test ./internal/[package]/ -run TestSpecific`
2. **Be deterministic**: No flaky tests based on timing
3. **Use proper cleanup**: Close channels, cancel contexts
4. **Avoid sleeps**: Use synchronization primitives instead
5. **Handle edge cases**: Test nil values, empty inputs, etc.

### Concurrency Tests

Use `sync.WaitGroup` for deterministic tests:

```go
func TestConcurrency_WorkerPool(t *testing.T) {
    var wg sync.WaitGroup
    
    wg.Add(3)
    for i := 0; i < 3; i++ {
        go func(id int) {
            defer wg.Done()
            // worker logic
        }(i)
    }
    
    wg.Wait() // Wait for all workers
}
```

### Avoid Timing-Based Tests

❌ Bad (flaky):
```go
time.Sleep(100 * time.Millisecond)
// hope the goroutine finished
```

✅ Good (deterministic):
```go
done := make(chan bool)
go func() {
    // work
    done <- true
}()
<-done // wait for completion
```

## 📦 Package Organization

### Package Guidelines

- **One concept per file**: Don't mix unrelated topics
- **Minimal coupling**: Avoid importing other internal packages
- **Standard library only**: No external dependencies unless approved
- **Clear package names**: Use singular form (e.g., `concurrency`, not `concurrencies`)

### Package Documentation

Add package-level documentation:

```go
// Package concurrency demonstrates Go's concurrency primitives including
// goroutines, channels, select statements, and synchronization mechanisms.
// These are critical topics for backend engineering interviews.
package concurrency
```

## 🔍 Code Review Checklist

Before submitting:

- [ ] Code passes `go test ./...`
- [ ] Code passes `go vet ./...`
- [ ] Code formatted with `go fmt`
- [ ] Includes "Why interviewers ask this" comment
- [ ] Includes "Common pitfalls" comment
- [ ] Includes "Key takeaway" comment
- [ ] Tests are deterministic (no sleeps or race conditions)
- [ ] Test names clearly describe behavior
- [ ] Complexity analysis included (for algorithms/data structures)
- [ ] No external dependencies added
- [ ] Documentation is clear and concise

## 🚀 Adding a New Package

If adding an entirely new package:

1. Create directory: `internal/[newpackage]/`
2. Add package documentation
3. Create at least one topic with implementation + test
4. Update main README.md with package description
5. Add to recommended study order if relevant

## 📝 Documentation Standards

### Inline Comments

- Use complete sentences
- Start with capital letter
- End with period
- Be concise but clear

### Function Comments

```go
// CalculateFibonacci returns the nth Fibonacci number using dynamic programming.
// Time Complexity: O(n), Space Complexity: O(n)
func CalculateFibonacci(n int) int {
    // implementation
}
```

## 🐛 Reporting Issues

When reporting issues:

1. Specify which package/file
2. Include Go version: `go version`
3. Provide minimal reproduction steps
4. Include expected vs actual behavior

## 💡 Suggesting New Topics

Great topics for this handbook:

- Frequently asked in interviews
- Demonstrate Go-specific concepts
- Have clear, testable behavior
- Relevant to backend engineering

## 🎓 Learning Resources

When adding topics, consider referencing:

- Go official documentation
- Effective Go
- Go blog posts
- Common interview questions

## ⚡ Quick Commands

```bash
# Format code
go fmt ./...

# Run tests
go test ./...

# Run with race detector
go test -race ./...

# Run vet
go vet ./...

# Run specific test
go test ./internal/[package]/ -run TestName

# Run benchmarks
go test -bench=. ./internal/[package]/
```

## 🤝 Questions?

If you have questions about contributing:

1. Check existing implementations for examples
2. Review this guide
3. Open an issue for clarification

---

Thank you for helping make this handbook better! 🙏
