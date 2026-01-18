package patterns

import (
	"errors"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(2, 100*time.Millisecond)

	// 1. Start Closed
	if state := cb.State(); state != StateClosed {
		t.Errorf("expected closed, got %v", state)
	}

	// 2. Fail once
	cb.Execute(func() error { return errors.New("fail") })
	if state := cb.State(); state != StateClosed {
		t.Errorf("should stay closed after 1 failure")
	}

	// 3. Fail twice -> Open
	cb.Execute(func() error { return errors.New("fail") })
	if state := cb.State(); state != StateOpen {
		t.Errorf("should open after 2 failures")
	}

	// 4. Fail Fast
	err := cb.Execute(func() error { return nil }) // Action shouldn't run
	if err != ErrCircuitOpen {
		t.Errorf("expected ErrCircuitOpen, got %v", err)
	}

	// 5. Wait for timeout -> Half Open recovery
	time.Sleep(150 * time.Millisecond)

	// 6. Success -> Closed
	err = cb.Execute(func() error { return nil })
	if err != nil {
		t.Errorf("unexpected error during recovery: %v", err)
	}

	if state := cb.State(); state != StateClosed {
		t.Errorf("should return to closed after success")
	}
}
