package patterns

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"time"
)

// Why interviewers ask this:
// Network calls are unreliable. "Retry with Backoff" handles transient failures gracefully
// without overwhelming the server (thundering herd problem).
// Interviewers look for jitter (randomness) and capped delays.

// Common pitfalls:
// - Retrying forever (no max attempts)
// - Retrying on non-transient errors (like 400 Bad Request)
// - Blocking without Context support

// Key takeaway:
// Loop with limited attempts. Use `time.Sleep` with exponential duration `base * 2^attempt`.
// Add jitter to avoid synchronized retries across clients.

var ErrMaxRetriesReached = errors.New("max retries reached")

type RetryableFunc func(ctx context.Context) error

func RetryWithBackoff(ctx context.Context, maxAttempts int, initialBackoff time.Duration, fn RetryableFunc) error {
	var err error

	for attempt := 0; attempt < maxAttempts; attempt++ {
		err = fn(ctx)
		if err == nil {
			return nil
		}

		// If context is cancelled, stop immediately
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Last attempt failed, return error
		if attempt == maxAttempts-1 {
			return err
		}

		// Calculate backoff: initial * 2^attempt
		delay := initialBackoff * time.Duration(math.Pow(2, float64(attempt)))

		// Add Jitter (±10%) to prevent thundering herd
		jitter := time.Duration(rand.Int63n(int64(delay)/10 + 1))
		delay = delay + jitter

		// Wait or Context Cancel
		select {
		case <-time.After(delay):
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return ErrMaxRetriesReached // Fallback
}
