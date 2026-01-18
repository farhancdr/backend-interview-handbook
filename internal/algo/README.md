# 🧮 Algorithms

> **Essential algorithms for coding interviews**

This package implements common algorithmic patterns that appear frequently in technical interviews. Master these patterns to solve most coding problems.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **Binary Search** | [binary_search.go](binary_search.go) | Divide and conquer, search space reduction, O(log n) |
| **Sliding Window** | [sliding_window.go](sliding_window.go) | Fixed/variable window, two pointers, substring problems |
| **Two Pointers** | [two_pointers.go](two_pointers.go) | Left-right pointers, fast-slow pointers, in-place operations |
| **Sorting** | [sorting.go](sorting.go) | Quick sort, merge sort, heap sort, stability |
| **Dynamic Programming** | [dynamic_programming.go](dynamic_programming.go) | Memoization, tabulation, optimal substructure |

---

## 🚀 Quick Start

```bash
# Run all tests
go test -v ./algo/

# Run specific algorithm
go test -v ./algo/ -run TestBinarySearch

# Run benchmarks
go test -bench=. ./algo/
```

---

## 🎓 Learning Guide

### Recommended Order
1. **Two Pointers** - Simplest pattern
2. **Sliding Window** - Extension of two pointers
3. **Binary Search** - Fundamental divide and conquer
4. **Sorting** - Understanding O(n log n) algorithms
5. **Dynamic Programming** - Most challenging

### Common Interview Questions

**Binary Search:**
- "Find element in sorted array"
- "Find first/last occurrence"
- "Search in rotated sorted array"

**Sliding Window:**
- "Longest substring without repeating characters"
- "Minimum window substring"
- "Maximum sum subarray of size K"

**Two Pointers:**
- "Two sum in sorted array"
- "Remove duplicates from sorted array"
- "Container with most water"

**Dynamic Programming:**
- "Fibonacci sequence"
- "Longest common subsequence"
- "0/1 Knapsack"

---

## 💡 Key Takeaways

### Binary Search
```go
// Template:
left, right := 0, len(arr)-1
for left <= right {
    mid := left + (right-left)/2  // Prevent overflow
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        left = mid + 1
    } else {
        right = mid - 1
    }
}

// Time: O(log n)
// Space: O(1)

// Variations:
// - Find first occurrence
// - Find last occurrence
// - Find insertion position
// - Search in rotated array
```

### Sliding Window
```go
// Fixed window:
for i := 0; i < len(arr)-k+1; i++ {
    window := arr[i:i+k]
    // Process window
}

// Variable window:
left := 0
for right := 0; right < len(arr); right++ {
    // Expand window
    add(arr[right])
    
    // Shrink window if needed
    for !valid() {
        remove(arr[left])
        left++
    }
    
    // Update result
}

// Time: O(n)
// Space: O(1) or O(k)
```

### Two Pointers
```go
// Pattern 1: Opposite ends
left, right := 0, len(arr)-1
for left < right {
    if condition {
        left++
    } else {
        right--
    }
}

// Pattern 2: Same direction (fast-slow)
slow, fast := 0, 0
for fast < len(arr) {
    if condition {
        arr[slow] = arr[fast]
        slow++
    }
    fast++
}

// Time: O(n)
// Space: O(1)
```

### Sorting Algorithms
```go
// Quick Sort:
// - Average: O(n log n)
// - Worst: O(n²)
// - Space: O(log n)
// - NOT stable

// Merge Sort:
// - Always: O(n log n)
// - Space: O(n)
// - Stable

// Heap Sort:
// - Always: O(n log n)
// - Space: O(1)
// - NOT stable

// Go's sort.Sort uses:
// - Quicksort for large arrays
// - Insertion sort for small arrays
// - Heapsort as fallback
```

### Dynamic Programming
```go
// Two approaches:

// 1. Memoization (Top-down)
memo := make(map[int]int)
func fib(n int) int {
    if n <= 1 {
        return n
    }
    if val, ok := memo[n]; ok {
        return val
    }
    memo[n] = fib(n-1) + fib(n-2)
    return memo[n]
}

// 2. Tabulation (Bottom-up)
func fib(n int) int {
    if n <= 1 {
        return n
    }
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}

// Steps:
// 1. Define subproblem
// 2. Find recurrence relation
// 3. Identify base cases
// 4. Determine computation order
```

---

## 🧪 Pattern Recognition

### When to Use Each Pattern

**Binary Search:**
- Sorted array
- Search space can be divided
- "Find minimum/maximum satisfying condition"

**Sliding Window:**
- Contiguous subarray/substring
- "Longest/shortest/maximum/minimum"
- Fixed or variable window size

**Two Pointers:**
- Sorted array
- In-place modification
- Pair finding
- Cycle detection

**Dynamic Programming:**
- Optimal substructure
- Overlapping subproblems
- "Count ways", "Maximize/Minimize"
- Can be solved recursively

---

## ⚠️ Common Pitfalls

1. **Binary Search**
   - Integer overflow: Use `mid = left + (right-left)/2`
   - Off-by-one errors: Check `left <= right` vs `left < right`
   - Infinite loops: Ensure left/right always move

2. **Sliding Window**
   - Not shrinking window properly
   - Forgetting to update result
   - Wrong window size calculation

3. **Two Pointers**
   - Not handling duplicates
   - Wrong pointer movement
   - Missing edge cases (empty array, single element)

4. **Dynamic Programming**
   - Wrong base cases
   - Incorrect recurrence relation
   - Not optimizing space (can often reduce from O(n) to O(1))

---

## 🔗 Related Topics

- **[Data Structures](../ds/)** - Data structures used by algorithms
- **[LeetCode](../leetcode/)** - Practice problems
- **[Basics](../basics/)** - Slice and array fundamentals

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
