package systemdesign

import (
	"testing"
)

func TestIdempotencyManager(t *testing.T) {
	im := NewIdempotencyManager()
	key := "req-123"

	// 1. First Call - Should Execute
	executed := false
	res, err := im.ProcessWithIdempotency(key, func() (string, error) {
		executed = true
		return "success-result", nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !executed {
		t.Error("expected action to execute")
	}
	if res != "success-result" {
		t.Errorf("expected success-result, got %s", res)
	}

	// 2. Second Call - Should NOT Execute, but return cached result
	executed = false
	res, err = im.ProcessWithIdempotency(key, func() (string, error) {
		executed = true
		return "should-not-run", nil
	})

	if err != nil {
		t.Fatalf("unexpected duplicate error: %v", err)
	}
	if executed {
		t.Error("expected action to skip execution")
	}
	if res != "success-result" {
		t.Errorf("expected cached result 'success-result', got %s", res)
	}

	// 3. Concurrent Processing Simulation
	im2 := NewIdempotencyManager()
	im2.CheckAndSet("req-456") // Lock it

	_, err = im2.ProcessWithIdempotency("req-456", func() (string, error) { return "ok", nil })
	if err == nil || err.Error() != "request already in progress" {
		t.Errorf("expected in-progress error, got %v", err)
	}
}
