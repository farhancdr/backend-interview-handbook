package memory

// Why interviewers ask this:
// Writing and interpreting benchmarks is essential for optimizing Go applications. Interviewers
// want to see if you know how to write benchmarks using the testing.B package, how to reset
// timers, how to avoid compiler optimizations that obscure results, and how to verify allocation
// counts.

// Common pitfalls:
// - Compiler optimizations eliminating code (e.g., dead code elimination)
// - Including setup time in benchmark measurement
// - Not using b.N loop correctly
// - Benchmarking with too few iterations or unstable environment
// - Ignoring memory allocation statistics

// Key takeaway:
// Use `go test -bench=.` to run benchmarks. Use `testing.B` with `for i := 0; i < b.N; i++`.
// Use `b.ResetTimer()` to exclude setup. Assign results to package-level variables to prevent
// compiler optimizations. Use `b.ReportAllocs()` to track memory allocations.

import (
	"math/rand"
	"sort"
	"time"
)

// Global variable to prevent compiler optimization
var result int

// InefficientFunction simulates a slow operation
func InefficientFunction(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		time.Sleep(1 * time.Microsecond) // Simulate work
		sum += i
	}
	return sum
}

// EfficientFunction simulates a fast operation
func EfficientFunction(n int) int {
	// Formula: n*(n-1)/2
	return n * (n - 1) / 2
}

// FunctionWithAllocation allocates memory
func FunctionWithAllocation(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

// FunctionWithoutAllocation avoids allocation (uses pre-allocated buffer)
func FunctionWithoutAllocation(buf []int) {
	for i := 0; i < len(buf); i++ {
		buf[i] = i
	}
}

// BubbleSort implementation (O(n^2))
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// QuickSort implementation (O(n log n))
// actually using stdlib sort which is usually efficient
func StandardSort(arr []int) {
	sort.Ints(arr)
}

// GenerateRandomSlice helper for benchmarks
func GenerateRandomSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}
	return s
}
