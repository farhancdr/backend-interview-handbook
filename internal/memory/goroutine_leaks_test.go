package memory

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func TestLeakyGoroutine(t *testing.T) {
	leaked := LeakyGoroutine()

	if leaked < 1 {
		t.Error("expected at least 1 goroutine to leak")
	}
}

func TestFixedWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	leaked := FixedWithContext(ctx)

	if leaked < 1 {
		t.Error("goroutine should have started")
	}

	// Cancel context
	cancel()
	time.Sleep(20 * time.Millisecond)

	// Goroutine should exit (we can't easily verify this in test)
}

func TestLeakyChannelWait(t *testing.T) {
	leaked := LeakyChannelWait()

	if leaked < 1 {
		t.Error("expected at least 1 goroutine to leak")
	}
}

func TestFixedChannelClose(t *testing.T) {
	leaked := FixedChannelClose()

	// Should be 0 or very small (goroutine properly cleaned up)
	if leaked > 1 {
		t.Errorf("expected no leaks, got %d goroutines", leaked)
	}
}

func TestLeakyWorkerPool(t *testing.T) {
	numWorkers := 5
	leaked := LeakyWorkerPool(numWorkers)

	if leaked < numWorkers {
		t.Errorf("expected at least %d goroutines to leak, got %d", numWorkers, leaked)
	}
}

func TestFixedWorkerPool(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numWorkers := 5
	leaked := FixedWorkerPool(ctx, numWorkers)

	// Should be 0 or very small (workers properly cleaned up)
	if leaked > 1 {
		t.Errorf("expected no leaks, got %d goroutines", leaked)
	}
}

func TestDetectGoroutineLeaks(t *testing.T) {
	before := 10
	after := 15

	if !DetectGoroutineLeaks(before, after) {
		t.Error("should detect leak when after > before")
	}

	if DetectGoroutineLeaks(10, 10) {
		t.Error("should not detect leak when counts are equal")
	}
}

func TestGetGoroutineCount(t *testing.T) {
	count := GetGoroutineCount()

	if count < 1 {
		t.Error("should have at least 1 goroutine (the test)")
	}
}

func TestTimeoutGoroutine(t *testing.T) {
	// Test with long timeout (should complete)
	completed := TimeoutGoroutine(100 * time.Millisecond)
	if !completed {
		t.Error("goroutine should complete within timeout")
	}

	// Test with short timeout (should timeout)
	completed = TimeoutGoroutine(10 * time.Millisecond)
	if completed {
		t.Error("goroutine should timeout")
	}
}

func TestBufferedChannelLeak(t *testing.T) {
	leaked := BufferedChannelLeak()

	if leaked < 1 {
		t.Error("expected at least 1 goroutine to leak")
	}
}

func TestFixedBufferedChannel(t *testing.T) {
	leaked := FixedBufferedChannel()

	// Should be 0 or very small
	if leaked > 1 {
		t.Errorf("expected no leaks, got %d goroutines", leaked)
	}
}

func TestGoroutineMemoryUsage(t *testing.T) {
	numGoroutines, estimatedMemoryKB := GoroutineMemoryUsage()

	if numGoroutines < 1 {
		t.Error("should have at least 1 goroutine")
	}

	if estimatedMemoryKB < 2 {
		t.Error("estimated memory should be at least 2KB")
	}

	// Verify calculation
	if estimatedMemoryKB != numGoroutines*2 {
		t.Error("memory calculation incorrect")
	}
}

func TestGoroutineCleanup(t *testing.T) {
	before := runtime.NumGoroutine()

	ctx, cancel := context.WithCancel(context.Background())

	// Start some goroutines
	for i := 0; i < 10; i++ {
		go func() {
			<-ctx.Done()
		}()
	}

	time.Sleep(20 * time.Millisecond)
	during := runtime.NumGoroutine()

	if during <= before {
		t.Error("goroutines should have started")
	}

	// Cancel and wait for cleanup
	cancel()
	time.Sleep(50 * time.Millisecond)

	after := runtime.NumGoroutine()

	// Should be close to original count
	if after > before+2 {
		t.Errorf("goroutines may not have cleaned up properly: before=%d after=%d", before, after)
	}
}

// Benchmark goroutine creation overhead

func BenchmarkGoroutineCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := make(chan bool)
		go func() {
			done <- true
		}()
		<-done
	}
}

func BenchmarkGoroutineWithContext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		done := make(chan bool)
		go func() {
			<-ctx.Done()
			done <- true
		}()
		cancel()
		<-done
	}
}
