package patterns

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetryWithBackoff(t *testing.T) {
	// 1. Success immediately
	ctx := context.Background()
	attempts := 0
	err := RetryWithBackoff(ctx, 3, 1*time.Millisecond, func(ctx context.Context) error {
		attempts++
		return nil
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if attempts != 1 {
		t.Errorf("expected 1 attempt, got %d", attempts)
	}

	// 2. Success after retry
	attempts = 0
	err = RetryWithBackoff(ctx, 3, 1*time.Millisecond, func(ctx context.Context) error {
		attempts++
		if attempts < 2 {
			return errors.New("fail")
		}
		return nil
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if attempts != 2 {
		t.Errorf("expected 2 attempts, got %d", attempts)
	}

	// 3. Max Retries Exceeded
	attempts = 0
	err = RetryWithBackoff(ctx, 3, 1*time.Millisecond, func(ctx context.Context) error {
		attempts++
		return errors.New("persistent fail")
	})

	if err == nil || err.Error() != "persistent fail" {
		t.Errorf("expected persistent fail, got %v", err)
	}
}
