package systemdesign

import (
	"context"
	"time"
)

// Why interviewers ask this:
// Real systems need to combine multiple patterns.
// "Orchestration" means managing timeouts across the whole request while retrying individual parts.
// Interviewers verify if you correctly select/cancel the context.

// Common pitfalls:
// - Retry logic resetting the context deadline (retries should fit WITHIN the parent timeout)
// - Not passing context to the dependency
// - Ignoring context cancellation errors

// Key takeaway:
// Create a parent context with a hard timeout.
// Pass that context to the retry loop.
// If the retry loop hits the timeout, it should abort immediately.

type Orchestrator struct {
	MaxRetries int
	Backoff    time.Duration
}

func (o *Orchestrator) ExecuteReliably(ctx context.Context, action func(context.Context) error) error {
	var err error

	for i := 0; i <= o.MaxRetries; i++ {
		// Check global timeout before attempt
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err = action(ctx)
		if err == nil {
			return nil
		}

		// If context cancelled during action, return immediately
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Wait before retry (if attempts remain)
		if i < o.MaxRetries {
			select {
			case <-time.After(o.Backoff):
				continue
			case <-ctx.Done():
				return ctx.Err() // Hard timeout hit during backoff
			}
		}
	}

	return err // Return last error if retries exhausted
}
