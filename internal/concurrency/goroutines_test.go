package concurrency

import (
	"sync"
	"testing"
	"time"
)

func TestGoroutine_Simple(t *testing.T) {
	result := SimpleGoroutine()

	expected := "Hello from goroutine"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestGoroutine_Multiple(t *testing.T) {
	count := 5
	results := MultipleGoroutines(count)

	// Should have correct number of results
	if len(results) != count {
		t.Errorf("expected %d results, got %d", count, len(results))
	}

	// Each result should be id * 2 (order not guaranteed)
	resultMap := make(map[int]bool)
	for _, r := range results {
		resultMap[r] = true
	}

	// Check all expected values are present
	for i := 0; i < count; i++ {
		expected := i * 2
		if !resultMap[expected] {
			t.Errorf("expected result %d not found", expected)
		}
	}
}

func TestGoroutine_ClosureCaptureCorrect(t *testing.T) {
	count := 5
	results := ClosureCaptureCorrect(count)

	// Should have correct number of results
	if len(results) != count {
		t.Errorf("expected %d results, got %d", count, len(results))
	}

	// All values 0-4 should be present
	resultMap := make(map[int]bool)
	for _, r := range results {
		resultMap[r] = true
	}

	for i := 0; i < count; i++ {
		if !resultMap[i] {
			t.Errorf("expected value %d not found", i)
		}
	}
}

func TestGoroutine_WithPanic(t *testing.T) {
	recovered := GoroutineWithPanic()

	if !recovered {
		t.Error("expected panic to be recovered")
	}
}

func TestGoroutine_WithTimeout(t *testing.T) {
	// Should complete within timeout
	result := GoroutineWithTimeout(50 * time.Millisecond)
	if result != "completed" {
		t.Errorf("expected completed, got %s", result)
	}

	// Should timeout
	result = GoroutineWithTimeout(200 * time.Millisecond)
	if result != "timeout" {
		t.Errorf("expected timeout, got %s", result)
	}
}

func TestGoroutine_Anonymous(t *testing.T) {
	result := AnonymousGoroutine()

	expected := "anonymous goroutine"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestGoroutine_Count(t *testing.T) {
	n := 100
	count := GoroutineCount(n)

	if count != n {
		t.Errorf("expected count %d, got %d", n, count)
	}
}

func TestGoroutine_Return(t *testing.T) {
	val1, val2 := GoroutineReturn()

	if val1 != 42 {
		t.Errorf("expected val1 to be 42, got %d", val1)
	}

	if val2 != 100 {
		t.Errorf("expected val2 to be 100, got %d", val2)
	}
}

func TestGoroutine_PrintNumbers(t *testing.T) {
	n := 5
	results := PrintNumbers(n)

	if len(results) != n {
		t.Errorf("expected %d results, got %d", n, len(results))
	}

	// Verify all numbers are present (order not guaranteed)
	found := make(map[string]bool)
	for _, result := range results {
		found[result] = true
	}

	// Just check we have n unique results
	if len(found) != n {
		t.Errorf("expected %d unique results, got %d", n, len(found))
	}
}

func TestGoroutine_WaitGroup(t *testing.T) {
	// Demonstrate proper WaitGroup usage
	var wg sync.WaitGroup
	counter := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		counter++
	}()

	wg.Wait()

	if counter != 1 {
		t.Errorf("expected counter to be 1, got %d", counter)
	}
}

func TestGoroutine_NoWaitCausesRace(t *testing.T) {
	// Skip this test when running with race detector
	// This test intentionally demonstrates a race condition
	if testing.Short() {
		t.Skip("Skipping race condition demonstration in short mode")
	}

	// This demonstrates why we need synchronization
	counter := 0

	// Launch goroutine but don't wait
	go func() {
		counter++
	}()

	// Without synchronization, counter might still be 0
	// (This test might be flaky, which proves the point)
	time.Sleep(10 * time.Millisecond) // Give goroutine time to run

	// Counter should eventually be 1
	if counter != 1 {
		t.Logf("counter is %d (demonstrates race condition)", counter)
	}
}
