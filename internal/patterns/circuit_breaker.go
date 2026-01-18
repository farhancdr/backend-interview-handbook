package patterns

import (
	"errors"
	"sync"
	"time"
)

// Why interviewers ask this:
// Distributed systems fail. The Circuit Breaker pattern prevents an application from
// repeatedly trying to execute an operation that's likely to fail. It prevents cascading
// failures and allows failing services time to recover.

// Common pitfalls:
// - Not handling concurrency (state changes must be atomic)
// - No half-open state (never checking if the service is back up)
// - Infinite timeouts

// Key takeaway:
// Three states: Closed (Normal), Open (Failing - fail fast), Half-Open (Testing recovery).

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

var ErrCircuitOpen = errors.New("circuit breaker is open")

type CircuitBreaker struct {
	mu               sync.Mutex
	state            State
	failures         int
	failureThreshold int
	resetTimeout     time.Duration
	lastFailure      time.Time
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		resetTimeout:     timeout,
	}
}

func (cb *CircuitBreaker) Execute(action func() error) error {
	cb.mu.Lock()

	// Transition logic: Open -> HalfOpen?
	if cb.state == StateOpen {
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			cb.state = StateHalfOpen
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	}
	cb.mu.Unlock()

	// Execute Action
	err := action()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		// Failure Logic
		cb.failures++
		cb.lastFailure = time.Now()

		if cb.failures >= cb.failureThreshold {
			cb.state = StateOpen
		}
		return err
	}

	// Success Logic
	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failures = 0
	} else {
		// Reset failures on success in Closed state (optional, or separate clean-up)
		cb.failures = 0
	}

	return nil
}

// State returns the current state (for testing)
func (cb *CircuitBreaker) State() State {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.state
}
