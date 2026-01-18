package concurrency

import "testing"

func TestWorkerPool_Basic(t *testing.T) {
	jobs := []Job{
		{ID: 1, Value: 2},
		{ID: 2, Value: 3},
		{ID: 3, Value: 4},
	}

	results := WorkerPool(2, jobs)

	if len(results) != len(jobs) {
		t.Errorf("expected %d results, got %d", len(jobs), len(results))
	}

	// Verify results (order not guaranteed)
	resultMap := make(map[int]int)
	for _, r := range results {
		resultMap[r.JobID] = r.Value
	}

	expected := map[int]int{
		1: 4,  // 2*2
		2: 9,  // 3*3
		3: 16, // 4*4
	}

	for id, expectedVal := range expected {
		if resultMap[id] != expectedVal {
			t.Errorf("expected result for job %d to be %d, got %d", id, expectedVal, resultMap[id])
		}
	}
}

func TestWorkerPool_Simple(t *testing.T) {
	numWorkers := 3
	numJobs := 10

	results := SimpleWorkerPool(numWorkers, numJobs)

	if len(results) != numJobs {
		t.Errorf("expected %d results, got %d", numJobs, len(results))
	}

	// Verify all results are present (order not guaranteed)
	resultMap := make(map[int]bool)
	for _, r := range results {
		resultMap[r] = true
	}

	for i := 0; i < numJobs; i++ {
		expected := i * 2
		if !resultMap[expected] {
			t.Errorf("expected result %d not found", expected)
		}
	}
}

func TestWorkerPool_Bounded(t *testing.T) {
	tasks := []func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
		func() int { return 4 },
		func() int { return 5 },
	}

	maxWorkers := 2
	results := BoundedWorkerPool(maxWorkers, tasks)

	if len(results) != len(tasks) {
		t.Errorf("expected %d results, got %d", len(tasks), len(results))
	}

	// Verify all results are present
	sum := 0
	for _, r := range results {
		sum += r
	}

	expectedSum := 1 + 2 + 3 + 4 + 5
	if sum != expectedSum {
		t.Errorf("expected sum %d, got %d", expectedSum, sum)
	}
}

func TestWorkerPool_WithContext(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}
	cancel := make(chan struct{})

	results := WorkerPoolWithContext(2, jobs, cancel)

	if len(results) != len(jobs) {
		t.Errorf("expected %d results, got %d", len(jobs), len(results))
	}
}

func TestWorkerPool_ManyJobs(t *testing.T) {
	numJobs := 100
	jobs := make([]Job, numJobs)
	for i := 0; i < numJobs; i++ {
		jobs[i] = Job{ID: i, Value: i}
	}

	results := WorkerPool(10, jobs)

	if len(results) != numJobs {
		t.Errorf("expected %d results, got %d", numJobs, len(results))
	}
}

func TestWorkerPool_SingleWorker(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}
	results := SimpleWorkerPool(1, len(jobs))

	if len(results) != len(jobs) {
		t.Errorf("expected %d results, got %d", len(jobs), len(results))
	}
}

func TestWorkerPool_ManyWorkers(t *testing.T) {
	numJobs := 10
	numWorkers := 20 // More workers than jobs

	results := SimpleWorkerPool(numWorkers, numJobs)

	if len(results) != numJobs {
		t.Errorf("expected %d results, got %d", numJobs, len(results))
	}
}
