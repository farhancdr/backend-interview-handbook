package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Why interviewers ask this:
// Goroutines are fundamental to Go's concurrency model. Understanding their
// lifecycle, scheduling, and proper usage is essential for any Go developer.
// This is one of the most frequently asked topics in Go interviews.

// Common pitfalls:
// - Forgetting to wait for goroutines to complete (goroutine leaks)
// - Not understanding that goroutines are not threads (they're lighter)
// - Assuming goroutines run in a specific order
// - Capturing loop variables incorrectly in closures
// - Not handling panics in goroutines

// Key takeaway:
// Goroutines are lightweight, managed by the Go runtime scheduler.
// Always ensure goroutines complete or are properly cancelled.
// Use sync.WaitGroup or channels to coordinate completion.

// SimpleGoroutine launches a single goroutine
func SimpleGoroutine() string {
	result := make(chan string, 1)

	go func() {
		result <- "Hello from goroutine"
	}()

	return <-result
}

// MultipleGoroutines launches multiple goroutines and waits for completion
func MultipleGoroutines(count int) []int {
	var wg sync.WaitGroup
	results := make(chan int, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			results <- id * 2
		}(i) // Pass i as parameter to avoid closure capture issue
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(results)

	// Collect results
	var output []int
	for result := range results {
		output = append(output, result)
	}

	return output
}

// ClosureCaptureWrong demonstrates the wrong way to capture loop variables
func ClosureCaptureWrong(count int) []int {
	var wg sync.WaitGroup
	results := make(chan int, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// BUG: All goroutines might see the same value of i
			results <- i
		}()
	}

	wg.Wait()
	close(results)

	var output []int
	for result := range results {
		output = append(output, result)
	}

	return output
}

// ClosureCaptureCorrect demonstrates the correct way to capture loop variables
func ClosureCaptureCorrect(count int) []int {
	var wg sync.WaitGroup
	results := make(chan int, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(id int) { // Pass as parameter
			defer wg.Done()
			results <- id
		}(i) // Capture current value of i
	}

	wg.Wait()
	close(results)

	var output []int
	for result := range results {
		output = append(output, result)
	}

	return output
}

// GoroutineWithPanic demonstrates panic handling in goroutines
func GoroutineWithPanic() (recovered bool) {
	done := make(chan bool)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				done <- true // Panic was recovered
			}
		}()

		panic("goroutine panic")
	}()

	return <-done
}

// GoroutineLeak demonstrates a potential goroutine leak
func GoroutineLeak() {
	// BAD: This goroutine will never complete
	// go func() {
	// 	ch := make(chan int) // Unbuffered channel
	// 	ch <- 1              // Blocks forever, no receiver
	// }()

	// The goroutine above would leak because it blocks forever
}

// GoroutineWithTimeout demonstrates timeout pattern
func GoroutineWithTimeout(duration time.Duration) string {
	result := make(chan string, 1)

	go func() {
		time.Sleep(duration)
		result <- "completed"
	}()

	select {
	case res := <-result:
		return res
	case <-time.After(100 * time.Millisecond):
		return "timeout"
	}
}

// AnonymousGoroutine demonstrates anonymous goroutine usage
func AnonymousGoroutine() string {
	done := make(chan string)

	go func() {
		done <- "anonymous goroutine"
	}()

	return <-done
}

// GoroutineCount demonstrates launching many goroutines
func GoroutineCount(n int) int {
	var wg sync.WaitGroup
	counter := 0
	var mu sync.Mutex

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	return counter
}

// GoroutineReturn demonstrates returning values from goroutines
func GoroutineReturn() (int, int) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 42
	}()

	go func() {
		ch2 <- 100
	}()

	return <-ch1, <-ch2
}

// PrintNumbers demonstrates concurrent printing (order not guaranteed)
func PrintNumbers(n int) []string {
	var wg sync.WaitGroup
	results := make(chan string, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			results <- fmt.Sprintf("Number: %d", num)
		}(i)
	}

	wg.Wait()
	close(results)

	var output []string
	for result := range results {
		output = append(output, result)
	}

	return output
}
