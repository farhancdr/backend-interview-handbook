package concurrency

import (
	"sync"
	"testing"
)

func TestMutex_Counter(t *testing.T) {
	c := &Counter{}

	// Increment concurrently
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()

	if c.Value() != 100 {
		t.Errorf("expected counter to be 100, got %d", c.Value())
	}
}

func TestRWMutex_Counter(t *testing.T) {
	c := &RWCounter{}

	// Increment concurrently
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()

	if c.Value() != 100 {
		t.Errorf("expected counter to be 100, got %d", c.Value())
	}
}

func TestMutex_SafeMap(t *testing.T) {
	sm := NewSafeMap()

	// Set values concurrently
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Set("key", id)
		}(i)
	}

	wg.Wait()

	// Get value (should be one of the set values)
	val, ok := sm.Get("key")
	if !ok {
		t.Error("expected key to exist")
	}

	if val < 0 || val >= 10 {
		t.Errorf("expected value between 0-9, got %d", val)
	}
}

func TestMutex_SafeConcurrentIncrement(t *testing.T) {
	n := 1000
	result := SafeConcurrentIncrement(n)

	if result != n {
		t.Errorf("expected %d, got %d", n, result)
	}
}

func TestMutex_WithDefer(t *testing.T) {
	result := MutexWithDefer()

	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestRWMutex_Performance(t *testing.T) {
	readers := 100
	results := RWMutexPerformance(readers)

	if len(results) != readers {
		t.Errorf("expected %d results, got %d", readers, len(results))
	}

	// All readers should see the same value
	for i, val := range results {
		if val != 100 {
			t.Errorf("expected results[%d] to be 100, got %d", i, val)
		}
	}
}

func TestMutex_ZeroValue(t *testing.T) {
	result := MutexZeroValue()

	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestMutex_MultipleGoroutines(t *testing.T) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	// Launch many goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if counter != 1000 {
		t.Errorf("expected counter to be 1000, got %d", counter)
	}
}

func TestRWMutex_MultipleReaders(t *testing.T) {
	var mu sync.RWMutex
	var wg sync.WaitGroup
	value := 42

	// Many concurrent readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.RLock()
			_ = value // Read
			mu.RUnlock()
		}()
	}

	wg.Wait()
}

func TestMutex_SafeMapConcurrent(t *testing.T) {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Set("key"+string(rune('0'+id%10)), id)
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Get("key" + string(rune('0'+id%10)))
		}(i)
	}

	wg.Wait()
}
