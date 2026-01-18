package memory

// Why interviewers ask this:
// Goroutine leaks are a common source of memory leaks in Go applications. Interviewers want
// to see if you understand how goroutines can leak, how to detect them, and how to prevent
// them. This is crucial for writing production-grade concurrent Go code.

// Common pitfalls:
// - Starting goroutines that never exit (waiting on channels that never close)
// - Not using context for cancellation
// - Forgetting to close channels that goroutines are waiting on
// - Creating goroutines in loops without proper cleanup
// - Not handling goroutine lifecycle in long-running services

// Key takeaway:
// Goroutines are cheap but not free. Each goroutine uses memory (stack starts at 2KB).
// Leaked goroutines accumulate and cause memory leaks. Always ensure goroutines can exit.
// Use context for cancellation, close channels when done, and use sync.WaitGroup to track
// goroutine completion. Monitor goroutine count with runtime.NumGoroutine().

import (
	"context"
	"runtime"
	"sync"
	"time"
)

// LeakyGoroutine demonstrates a goroutine that never exits
// This causes a memory leak
func LeakyGoroutine() int {
	before := runtime.NumGoroutine()

	// This goroutine will never exit
	go func() {
		ch := make(chan int)
		<-ch // Blocks forever, channel never receives
	}()

	time.Sleep(10 * time.Millisecond) // Give goroutine time to start
	after := runtime.NumGoroutine()

	return after - before
}

// FixedWithContext shows proper goroutine cleanup using context
func FixedWithContext(ctx context.Context) int {
	before := runtime.NumGoroutine()

	go func() {
		<-ctx.Done() // Properly exits when context is cancelled
	}()

	time.Sleep(10 * time.Millisecond)
	after := runtime.NumGoroutine()

	return after - before
}

// LeakyChannelWait demonstrates goroutine waiting on channel that never closes
func LeakyChannelWait() int {
	before := runtime.NumGoroutine()

	ch := make(chan int)

	// Goroutine waits for channel that's never closed
	go func() {
		for range ch {
			// Process
		}
	}()

	time.Sleep(10 * time.Millisecond)
	after := runtime.NumGoroutine()

	// Note: ch is never closed, goroutine leaks
	return after - before
}

// FixedChannelClose shows proper channel closing
func FixedChannelClose() int {
	before := runtime.NumGoroutine()

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
			// Process
		}
	}()

	close(ch) // Close channel to allow goroutine to exit
	wg.Wait() // Wait for goroutine to finish

	after := runtime.NumGoroutine()
	return after - before
}

// LeakyWorkerPool demonstrates workers that never exit
func LeakyWorkerPool(numWorkers int) int {
	before := runtime.NumGoroutine()

	jobs := make(chan int)

	for i := 0; i < numWorkers; i++ {
		go func() {
			for job := range jobs {
				_ = job
			}
		}()
	}

	time.Sleep(10 * time.Millisecond)
	after := runtime.NumGoroutine()

	// jobs channel never closed, workers leak
	return after - before
}

// FixedWorkerPool shows proper worker pool cleanup
func FixedWorkerPool(ctx context.Context, numWorkers int) int {
	before := runtime.NumGoroutine()

	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case job, ok := <-jobs:
					if !ok {
						return // Channel closed
					}
					_ = job
				case <-ctx.Done():
					return // Context cancelled
				}
			}
		}()
	}

	close(jobs)
	wg.Wait()

	after := runtime.NumGoroutine()
	return after - before
}

// DetectGoroutineLeaks shows how to detect goroutine leaks
func DetectGoroutineLeaks(before, after int) bool {
	return after > before
}

// GetGoroutineCount returns current goroutine count
func GetGoroutineCount() int {
	return runtime.NumGoroutine()
}

// LeakyHTTPClient demonstrates goroutines leaked by not reading response body
// This is a common real-world leak
func LeakyHTTPClient() {
	// Simulated: In real code, not reading/closing response body leaks goroutines
	// http.Get() starts goroutines that wait for body to be read
	// Always: defer resp.Body.Close() and io.ReadAll(resp.Body)
}

// TimeoutGoroutine shows using timeout to prevent indefinite blocking
func TimeoutGoroutine(timeout time.Duration) (completed bool) {
	done := make(chan bool)

	go func() {
		// Simulate work
		time.Sleep(50 * time.Millisecond)
		select {
		case done <- true:
		default:
			// Timeout already occurred, don't block
		}
	}()

	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}

// BufferedChannelLeak shows that buffered channels can still leak
func BufferedChannelLeak() int {
	before := runtime.NumGoroutine()

	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i // Will block after buffer fills
		}
	}()

	time.Sleep(10 * time.Millisecond)
	after := runtime.NumGoroutine()

	return after - before
}

// FixedBufferedChannel shows proper handling
func FixedBufferedChannel() int {
	before := runtime.NumGoroutine()

	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			select {
			case ch <- i:
			default:
				// Channel full, handle appropriately
				return
			}
		}
	}()

	wg.Wait()
	after := runtime.NumGoroutine()

	return after - before
}

// GoroutineMemoryUsage estimates memory used by goroutines
func GoroutineMemoryUsage() (numGoroutines int, estimatedMemoryKB int) {
	numGoroutines = runtime.NumGoroutine()
	// Each goroutine starts with ~2KB stack (can grow)
	estimatedMemoryKB = numGoroutines * 2

	return numGoroutines, estimatedMemoryKB
}
