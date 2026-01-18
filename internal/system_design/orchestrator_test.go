package systemdesign

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestOrchestrator(t *testing.T) {
	o := &Orchestrator{
		MaxRetries: 3,
		Backoff:    10 * time.Millisecond,
	}

	// 1. Success within attempt limit
	attempts := 0
	err := o.ExecuteReliably(context.Background(), func(ctx context.Context) error {
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

	// 2. Timeout Killing Logic
	// Global timeout 50ms. Backoff 20ms. 3 Retries would take 60ms+. Should fail.
	oSlow := &Orchestrator{MaxRetries: 5, Backoff: 20 * time.Millisecond}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	start := time.Now()
	err = oSlow.ExecuteReliably(ctx, func(ctx context.Context) error {
		return errors.New("forever fail")
	})

	if err != context.DeadlineExceeded {
		t.Errorf("expected DeadlineExceeded, got %v", err)
	}
	if time.Since(start) > 60*time.Millisecond {
		t.Error("operation took too long, didn't respect timeout")
	}
}
