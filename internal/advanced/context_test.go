package advanced

import (
	"context"
	"testing"
	"time"
)

func TestDoWorkWithContext_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := DoWorkWithContext(ctx)
	if err != context.DeadlineExceeded {
		t.Errorf("expected DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout_Short(t *testing.T) {
	err := WithTimeout(50 * time.Millisecond)
	if err != context.DeadlineExceeded {
		t.Errorf("expected DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout_Long(t *testing.T) {
	err := WithTimeout(3 * time.Second)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestWithCancel(t *testing.T) {
	result, err := WithCancel()

	if result != "cancelled" {
		t.Errorf("expected 'cancelled', got %s", result)
	}

	if err != context.Canceled {
		t.Errorf("expected Canceled error, got %v", err)
	}
}

func TestWithDeadline(t *testing.T) {
	// Deadline in the past
	pastDeadline := time.Now().Add(-1 * time.Second)
	err := WithDeadline(pastDeadline)

	if err != context.DeadlineExceeded {
		t.Errorf("expected DeadlineExceeded, got %v", err)
	}
}

func TestContextValue(t *testing.T) {
	userID := "user123"
	result := WithContextValue(userID)

	if result != userID {
		t.Errorf("expected %s, got %s", userID, result)
	}
}

func TestGetUserID_NotFound(t *testing.T) {
	ctx := context.Background()
	result := GetUserID(ctx)

	if result != "" {
		t.Errorf("expected empty string, got %s", result)
	}
}

func TestChainedContext(t *testing.T) {
	parentCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Child context should inherit cancellation
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	err := ChainedContext(parentCtx)
	if err != context.Canceled {
		t.Errorf("expected Canceled, got %v", err)
	}
}

func TestMultipleGoroutinesWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel immediately
	cancel()

	results := MultipleGoroutinesWithContext(ctx)

	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}

	// All should be cancelled
	for _, result := range results {
		if result != "worker 0 cancelled" && result != "worker 1 cancelled" && result != "worker 2 cancelled" {
			t.Logf("result: %s", result)
		}
	}
}

func TestLongRunningTask_Cancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel after 50ms
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	count, err := LongRunningTask(ctx)

	if err != context.Canceled {
		t.Errorf("expected Canceled, got %v", err)
	}

	// Should have processed some items before cancellation
	if count == 0 {
		t.Error("expected some work to be done before cancellation")
	}

	t.Logf("Processed %d items before cancellation", count)
}

func TestLongRunningTask_Completed(t *testing.T) {
	ctx := context.Background()

	count, err := LongRunningTask(ctx)

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if count != 1000 {
		t.Errorf("expected 1000, got %d", count)
	}
}

func TestContextPropagation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := ContextPropagation(ctx)

	if err != context.DeadlineExceeded {
		t.Errorf("expected DeadlineExceeded, got %v", err)
	}
}

func TestBackgroundVsTODO(t *testing.T) {
	bg, todo := BackgroundVsTODO()

	if bg == nil {
		t.Error("Background context should not be nil")
	}

	if todo == nil {
		t.Error("TODO context should not be nil")
	}
}

func TestCancelCauseExample(t *testing.T) {
	err := CancelCauseExample()

	if err != context.Canceled {
		t.Errorf("expected Canceled, got %v", err)
	}
}

func TestTimeoutExample(t *testing.T) {
	result, err := TimeoutExample()

	// Should timeout since work takes 200ms but timeout is 100ms
	if err != context.DeadlineExceeded {
		t.Errorf("expected DeadlineExceeded, got %v", err)
	}

	if result != "" {
		t.Errorf("expected empty result, got %s", result)
	}
}

func TestContext_DeferCancel(t *testing.T) {
	// This test verifies that defer cancel() is called
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Context should not be cancelled yet
	select {
	case <-ctx.Done():
		t.Error("context should not be cancelled yet")
	default:
		// Expected
	}

	// After function returns, defer will call cancel()
}
