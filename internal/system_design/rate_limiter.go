package systemdesign

import (
	"sync"
	"time"
)

// Why interviewers ask this:
// Variable rate limiting is crucial for protecting APIs from abuse and "thundering herd"
// scenarios. Implementing the Token Bucket algorithm demonstrates knowledge of concurrency,
// time-based logic, and resource management.

// Common pitfalls:
// - Busy waiting or spinning loops
// - Integer overflow on tokens
// - Non-atomic updates to token counts
// - Inaccurate refill logic (refilling one by one vs calculating elapsed time)

// Key takeaway:
// Calculate tokens based on elapsed time: current_tokens = min(capacity, old_tokens + (elapsed * rate)).
// Use a mutex to protect state.

type RateLimiter struct {
	mu         sync.Mutex
	capacity   float64 // Maximum number of tokens
	tokens     float64 // Current number of tokens
	refillRate float64 // Tokens per second
	lastRefill time.Time
}

func NewRateLimiter(capacity, refillRate float64) *RateLimiter {
	return &RateLimiter{
		capacity:   capacity,
		tokens:     capacity, // Start full
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Allow checks if a request can proceed. If yes, it consumes 1 token.
func (rl *RateLimiter) Allow() bool {
	return rl.AllowN(1)
}

// AllowN checks if a request for n tokens can proceed.
func (rl *RateLimiter) AllowN(n float64) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.refill()

	if rl.tokens >= n {
		rl.tokens -= n
		return true
	}

	return false
}

// refill adds tokens based on elapsed time without exceeding capacity
func (rl *RateLimiter) refill() {
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()

	// Calculate tokens to add
	tokensToAdd := elapsed * rl.refillRate

	if tokensToAdd > 0 {
		rl.tokens = min(rl.capacity, rl.tokens+tokensToAdd)
		rl.lastRefill = now
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
