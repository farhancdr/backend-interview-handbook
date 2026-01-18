# Implementation Guide - Backend Interview Handbook

## Purpose
This guide provides step-by-step instructions for AI agents to add new topics to the Backend Interview Handbook while maintaining consistency and quality.

## Standard Implementation Pattern

### Step 1: Create Implementation File

**File**: `internal/[package]/[topic_name].go`

**Template**:
```go
package [packagename]

// Why interviewers ask this:
// [2-3 sentences explaining why this topic is important for interviews]
// [What interviewers are looking for when they ask about this]

// Common pitfalls:
// - [Pitfall 1]
// - [Pitfall 2]
// - [Pitfall 3]
// - [Pitfall 4]

// Key takeaway:
// [1-2 sentences summarizing the most important concept]
// [What the candidate must remember]

// [Function implementations demonstrating the concept]
```

**Example** (`internal/concurrency/select_stmt.go`):
```go
package concurrency

// Why interviewers ask this:
// Select statements are crucial for coordinating multiple channels and implementing
// timeouts, cancellation, and non-blocking operations. Interviewers want to ensure
// you understand how select chooses between multiple channel operations.

// Common pitfalls:
// - Not understanding that select picks randomly when multiple cases are ready
// - Forgetting that select blocks if no case is ready (unless there's a default)
// - Not knowing that nil channels in select are ignored
// - Confusion about default case behavior

// Key takeaway:
// Select lets you wait on multiple channel operations. If multiple cases are ready,
// one is chosen at random. Default case makes select non-blocking.

// SelectBasic demonstrates basic select usage
func SelectBasic() string {
    ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)
    
    ch1 <- "from ch1"
    ch2 <- "from ch2"
    
    select {
    case msg := <-ch1:
        return msg
    case msg := <-ch2:
        return msg
    }
}
```

### Step 2: Create Test File

**File**: `internal/[package]/[topic_name]_test.go`

**Template**:
```go
package [packagename]

import "testing"

func Test[Topic]_[Behavior](t *testing.T) {
    // Arrange
    // Act
    // Assert
}
```

**Test Naming Convention**:
- `Test[Topic]_[SpecificBehavior]`
- Be descriptive about what's being tested
- Examples:
  - `TestSelect_RandomSelection`
  - `TestSelect_DefaultCase`
  - `TestSelect_NilChannel`

**Example**:
```go
package concurrency

import "testing"

func TestSelect_Basic(t *testing.T) {
    result := SelectBasic()
    
    // Should receive from one of the channels
    if result != "from ch1" && result != "from ch2" {
        t.Errorf("unexpected result: %s", result)
    }
}

func TestSelect_DefaultCase(t *testing.T) {
    ch := make(chan string)
    
    select {
    case <-ch:
        t.Error("should not receive from empty channel")
    default:
        // Expected: default case executes
    }
}
```

### Step 3: Implement Multiple Examples

Each topic should have 5-10 functions demonstrating different aspects:

1. **Basic usage** - Simplest possible example
2. **Common patterns** - Real-world usage
3. **Edge cases** - Nil values, empty inputs, etc.
4. **Gotchas** - Common mistakes (commented out if they would panic/deadlock)
5. **Best practices** - Idiomatic Go

**Example Structure**:
```go
// Basic usage
func BasicExample() { }

// Common pattern
func CommonPattern() { }

// Edge case
func EdgeCase() { }

// Gotcha (commented to prevent panic)
func GotchaExample() {
    // This would panic:
    // var m map[string]int
    // m["key"] = 42
    
    // Correct way:
    m := make(map[string]int)
    m["key"] = 42
}

// Best practice
func BestPractice() { }
```

## Package-Specific Guidelines

### Basics Package
- Focus on fundamentals
- Keep examples simple
- Emphasize differences (e.g., value vs reference)
- Cover zero values

### Intermediate Package
- Build on basics
- Show common interview traps
- Demonstrate proper error handling
- Include JSON/marshaling edge cases

### Advanced Package
- Include "Interview Notes" comment block
- Show performance implications
- Demonstrate escape analysis
- Cover context patterns

### Concurrency Package
- **CRITICAL**: Tests must be deterministic
- Use `sync.WaitGroup` for coordination
- Avoid `time.Sleep` in tests (use channels instead)
- Document potential deadlocks (don't trigger them)
- Show both correct and incorrect patterns

**Concurrency Test Pattern**:
```go
func TestConcurrency_Pattern(t *testing.T) {
    var wg sync.WaitGroup
    done := make(chan bool)
    
    wg.Add(1)
    go func() {
        defer wg.Done()
        // Work
        done <- true
    }()
    
    wg.Wait()
    <-done
}
```

### Memory Package
- Include benchmarks (`BenchmarkXxx`)
- Show capacity growth
- Demonstrate allocation behavior
- Use `testing.B` for benchmarks

**Benchmark Pattern**:
```go
func BenchmarkSliceAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := make([]int, 0)
        for j := 0; j < 1000; j++ {
            s = append(s, j)
        }
    }
}
```

### Data Structures Package
- Include time/space complexity comments
- Implement from scratch (no imports)
- Test edge cases (empty, single element, etc.)
- Show both iterative and recursive approaches where applicable

**Complexity Comment Pattern**:
```go
// BinarySearch performs binary search on a sorted slice
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearch(arr []int, target int) int {
    // implementation
}
```

### Algorithms Package
- Include problem statement as comment
- Show brute force first, then optimized
- Explain complexity trade-offs
- Provide multiple test cases

**Algorithm Pattern**:
```go
// Problem: Find the maximum subarray sum
// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Output: 6 (subarray [4, -1, 2, 1])
//
// Approach: Kadane's Algorithm
// Time Complexity: O(n)
// Space Complexity: O(1)
func MaxSubarraySum(nums []int) int {
    // implementation
}
```

## Testing Best Practices

### 1. Deterministic Tests
```go
// ❌ BAD: Flaky, timing-dependent
func TestBad(t *testing.T) {
    go doWork()
    time.Sleep(100 * time.Millisecond)
    // Hope work is done
}

// ✅ GOOD: Deterministic
func TestGood(t *testing.T) {
    done := make(chan bool)
    go func() {
        doWork()
        done <- true
    }()
    <-done // Wait for completion
}
```

### 2. Table-Driven Tests
```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"zero", 0, 0},
        {"positive", 5, 25},
        {"negative", -3, 9},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Square(tt.input)
            if result != tt.expected {
                t.Errorf("expected %d, got %d", tt.expected, result)
            }
        })
    }
}
```

### 3. Test Helpers
```go
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

## Common Patterns

### Pattern 1: Demonstrating Gotchas
```go
// DemonstrateNilMapGotcha shows the nil map panic
func DemonstrateNilMapGotcha() {
    var m map[string]int // nil map
    
    // Reading is safe
    _ = m["key"] // Returns zero value
    
    // Writing panics (don't actually do this in code)
    // m["key"] = 42 // PANIC!
    
    // Correct way:
    m = make(map[string]int)
    m["key"] = 42 // Safe
}
```

### Pattern 2: Before/After Comparison
```go
// WrongWay demonstrates incorrect approach
func WrongWay() {
    // Incorrect implementation
}

// RightWay demonstrates correct approach  
func RightWay() {
    // Correct implementation
}
```

### Pattern 3: Progressive Complexity
```go
// Level1_Basic - Simplest example
func Level1_Basic() { }

// Level2_Intermediate - More realistic
func Level2_Intermediate() { }

// Level3_Advanced - Production-ready
func Level3_Advanced() { }
```

## Checklist for New Topics

Before committing a new topic, verify:

- [ ] Implementation file has required comment blocks
- [ ] At least 5 test functions
- [ ] Tests cover edge cases
- [ ] Tests are deterministic (no `time.Sleep` in assertions)
- [ ] All tests pass: `go test ./internal/[package]/`
- [ ] No race conditions: `go test -race ./internal/[package]/`
- [ ] Vet passes: `go vet ./internal/[package]/`
- [ ] Code is formatted: `go fmt ./internal/[package]/`
- [ ] Coverage > 85%: `go test -cover ./internal/[package]/`
- [ ] Test names are descriptive
- [ ] Comments explain "why", not "what"
- [ ] No external dependencies added

## File Size Guidelines

- Implementation file: 100-300 lines
- Test file: 150-400 lines
- If larger, consider splitting into multiple topics

## Example: Complete Topic Implementation

See these reference implementations:
- `internal/basics/interfaces.go` - Good example of interface patterns
- `internal/concurrency/channels.go` - Good example of channel patterns
- `internal/concurrency/worker_pool.go` - Good example of practical patterns

## Workflow Summary

1. Choose next topic from task.md
2. Create `[topic].go` with required comments
3. Implement 5-10 functions demonstrating concept
4. Create `[topic]_test.go` with comprehensive tests
5. Run tests: `go test ./internal/[package]/ -v`
6. Run race detector: `go test -race ./internal/[package]/`
7. Check coverage: `go test -cover ./internal/[package]/`
8. Update task.md to mark topic complete
9. Commit changes

## Tips for AI Agents

1. **Start simple**: Begin with basic example, then add complexity
2. **Test as you go**: Write test after each function
3. **Follow patterns**: Look at existing files for inspiration
4. **Be consistent**: Match naming and structure of existing code
5. **Document gotchas**: Interview questions often focus on edge cases
6. **Think interview**: What would an interviewer ask about this topic?

## Common Mistakes to Avoid

1. ❌ Forgetting required comment blocks
2. ❌ Using `time.Sleep` in tests
3. ❌ Not handling edge cases
4. ❌ Tests that can fail randomly
5. ❌ Adding external dependencies
6. ❌ Overly complex examples
7. ❌ Not explaining "why"
8. ❌ Copying code without understanding

## Quality Standards

Every topic should be:
- **Runnable**: Tests demonstrate behavior
- **Clear**: Easy to understand for juniors
- **Complete**: Covers common interview questions
- **Correct**: No bugs or race conditions
- **Concise**: No unnecessary complexity

---

**Remember**: The goal is executable documentation that helps backend engineers prepare for interviews through hands-on Go practice. Keep it simple, clear, and focused on interview-relevant content.
