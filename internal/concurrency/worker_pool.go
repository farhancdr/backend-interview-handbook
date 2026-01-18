package concurrency

import "sync"

// Why interviewers ask this:
// Worker pools are a common concurrency pattern for limiting parallelism and
// managing resources. Understanding how to implement and use worker pools is
// essential for production Go code and frequently asked in interviews.

// Common pitfalls:
// - Not closing job channel (workers never exit)
// - Not waiting for workers to finish
// - Creating too many or too few workers
// - Not handling errors from workers
// - Forgetting to close results channel

// Key takeaway:
// Worker pools limit concurrency by having a fixed number of goroutines
// processing jobs from a shared channel. This prevents resource exhaustion
// and provides backpressure.

// Job represents work to be done
type Job struct {
	ID    int
	Value int
}

// Result represents the result of a job
type Result struct {
	JobID int
	Value int
}

// WorkerPool demonstrates the worker pool pattern
func WorkerPool(numWorkers int, jobs []Job) []Result {
	jobChan := make(chan Job, len(jobs))
	resultChan := make(chan Result, len(jobs))
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobChan, resultChan, &wg)
	}

	// Send jobs
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan) // Signal no more jobs

	// Wait for workers to finish
	wg.Wait()
	close(resultChan)

	// Collect results
	var results []Result
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// worker processes jobs from the job channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Process job (square the value)
		result := Result{
			JobID: job.ID,
			Value: job.Value * job.Value,
		}
		results <- result
	}
}

// SimpleWorkerPool demonstrates a simpler worker pool
func SimpleWorkerPool(numWorkers int, numJobs int) []int {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job * 2
			}
		}()
	}

	// Send jobs
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait and close results
	wg.Wait()
	close(results)

	// Collect results
	var output []int
	for result := range results {
		output = append(output, result)
	}

	return output
}

// BoundedWorkerPool demonstrates limiting concurrent work
func BoundedWorkerPool(maxWorkers int, tasks []func() int) []int {
	sem := make(chan struct{}, maxWorkers) // Semaphore
	results := make(chan int, len(tasks))
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go func(t func() int) {
			defer wg.Done()

			sem <- struct{}{}        // Acquire semaphore
			defer func() { <-sem }() // Release semaphore

			results <- t()
		}(task)
	}

	wg.Wait()
	close(results)

	var output []int
	for result := range results {
		output = append(output, result)
	}

	return output
}

// WorkerPoolWithContext demonstrates cancellable worker pool
func WorkerPoolWithContext(numWorkers int, jobs []int, cancel <-chan struct{}) []int {
	jobChan := make(chan int, len(jobs))
	resultChan := make(chan int, len(jobs))
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case job, ok := <-jobChan:
					if !ok {
						return
					}
					resultChan <- job * 2
				case <-cancel:
					return
				}
			}
		}()
	}

	// Send jobs
	for _, job := range jobs {
		select {
		case jobChan <- job:
		case <-cancel:
			close(jobChan)
			wg.Wait()
			close(resultChan)

			var results []int
			for result := range resultChan {
				results = append(results, result)
			}
			return results
		}
	}
	close(jobChan)

	wg.Wait()
	close(resultChan)

	var results []int
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}
