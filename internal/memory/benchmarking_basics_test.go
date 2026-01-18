package memory

import (
	"fmt"
	"testing"
)

// Test functions to ensure correctness before benchmarking

func TestEfficientVsInefficient(t *testing.T) {
	n := 10
	inefficient := InefficientFunction(n)
	efficient := EfficientFunction(n)

	if inefficient != efficient {
		t.Errorf("implementations differ: %d vs %d", inefficient, efficient)
	}
}

func TestAllocations(t *testing.T) {
	n := 10
	s1 := FunctionWithAllocation(n)

	s2 := make([]int, n)
	FunctionWithoutAllocation(s2)

	for i := 0; i < n; i++ {
		if s1[i] != i || s2[i] != i {
			t.Error("incorrect values")
		}
	}
}

// Benchmarks

// BenchmarkInefficient demonstrates measuring a slow function
func BenchmarkInefficient(b *testing.B) {
	// Setup (excluded from time if ResetTimer is used, but here it's fast)
	n := 100

	b.ResetTimer() // Start timing here
	for i := 0; i < b.N; i++ {
		// Use result to prevent compiler optimization
		result = InefficientFunction(n)
	}
}

// BenchmarkEfficient demonstrates measuring a fast function
func BenchmarkEfficient(b *testing.B) {
	n := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = EfficientFunction(n)
	}
}

// BenchmarkAllocations shows how to measure allocations
// Run with: go test -bench=Alloc -benchmem
func BenchmarkWithAllocation(b *testing.B) {
	b.ReportAllocs() // Report memory allocations

	for i := 0; i < b.N; i++ {
		_ = FunctionWithAllocation(1000)
	}
}

// BenchmarkNoAllocation shows zero allocations
func BenchmarkWithoutAllocation(b *testing.B) {
	b.ReportAllocs()

	// Pre-allocate buffer outside loop
	buf := make([]int, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FunctionWithoutAllocation(buf)
	}
}

// BenchmarkSorting compares sorting algorithms
func BenchmarkBubbleSort(b *testing.B) {
	size := 1000
	// We need a fresh slice for each iteration to properly benchmark sorting
	// But generating it takes time.
	// For reliable sorting benchmarks, we often pause/resume timer.

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer() // Pause timer
		data := GenerateRandomSlice(size)
		b.StartTimer() // Resume timer

		BubbleSort(data)
	}
}

func BenchmarkStandardSort(b *testing.B) {
	size := 1000

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := GenerateRandomSlice(size)
		b.StartTimer()

		StandardSort(data)
	}
}

// BenchmarkSubtest demonstrates using sub-benchmarks
func BenchmarkComparison(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Efficient-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = EfficientFunction(size)
			}
		})

		b.Run(fmt.Sprintf("Inefficient-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = InefficientFunction(size)
			}
		})
	}
}
