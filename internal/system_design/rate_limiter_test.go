package systemdesign

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	// Capacity 5, Refill 1 per second
	rl := NewRateLimiter(5, 1)

	// 1. Burst Allow
	for i := 0; i < 5; i++ {
		if !rl.Allow() {
			t.Errorf("expected to allow request %d", i)
		}
	}

	// 2. Limit Reached
	if rl.Allow() {
		t.Error("expected to deny request when empty")
	}

	// 3. Refill
	// Wait 1.1s to ensure at least 1 token is added
	time.Sleep(1100 * time.Millisecond)

	if !rl.Allow() {
		t.Error("expected to allow request after refill")
	}

	// Should be empty again
	if rl.Allow() {
		t.Error("expected to deny request after using refilled token")
	}
}

func TestRateLimiter_Concurrent(t *testing.T) {
	// High capacity to allow concurrency
	rl := NewRateLimiter(1000, 100)

	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func() {
			rl.Allow()
			done <- true
		}()
	}

	for i := 0; i < 100; i++ {
		<-done
	}

	// Just ensuring no panic/race occurred
}
