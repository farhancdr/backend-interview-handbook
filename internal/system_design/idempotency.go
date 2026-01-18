package systemdesign

import (
	"errors"
	"sync"
)

// Why interviewers ask this:
// Idempotency is critical for payment processing and reliable distributed systems.
// Interviewers verify if you understand how to deduplicate requests when clients retry
// due to network timeouts (the "at-least-once" delivery problem).

// Common pitfalls:
// - Not handling the "in-progress" state (race condition where two requests come same time)
// - Returning different results for duplicate calls
// - Storing keys forever (need TTL)

// Key takeaway:
// Store the state of a request key: {Status: Processing | Completed, Result: ...}.
// If status is Processing -> Error (Conflict) or Wait.
// If status is Completed -> Return stored result immediately.

type RequestStatus int

const (
	StatusProcessing RequestStatus = iota
	StatusCompleted
	StatusFailed
)

type IdempotencyRecord struct {
	Status RequestStatus
	Result string // Simplified result storage
}

type IdempotencyManager struct {
	mu    sync.Mutex
	store map[string]IdempotencyRecord
}

func NewIdempotencyManager() *IdempotencyManager {
	return &IdempotencyManager{
		store: make(map[string]IdempotencyRecord),
	}
}

// CheckAndSet returns true if operation should proceed, false if it's a duplicate
func (im *IdempotencyManager) CheckAndSet(key string) (bool, *IdempotencyRecord) {
	im.mu.Lock()
	defer im.mu.Unlock()

	if record, exists := im.store[key]; exists {
		// Duplicate request
		return false, &record
	}

	// New request -> Lock it as Processing
	im.store[key] = IdempotencyRecord{Status: StatusProcessing}
	return true, nil
}

// UpdateResult saves the result after processing
func (im *IdempotencyManager) UpdateResult(key string, result string, success bool) {
	im.mu.Lock()
	defer im.mu.Unlock()

	status := StatusCompleted
	if !success {
		status = StatusFailed
	}

	im.store[key] = IdempotencyRecord{
		Status: status,
		Result: result,
	}
}

// ProcessWithIdempotency simulates a full flow
func (im *IdempotencyManager) ProcessWithIdempotency(key string, action func() (string, error)) (string, error) {
	proceed, record := im.CheckAndSet(key)
	if !proceed {
		if record.Status == StatusProcessing {
			return "", errors.New("request already in progress")
		}
		if record.Status == StatusFailed {
			return "", errors.New("previous attempt failed")
		}
		return record.Result, nil // Return cached result
	}

	// Execute Action
	result, err := action()

	// Save Result
	if err != nil {
		im.UpdateResult(key, "", false)
		return "", err
	}

	im.UpdateResult(key, result, true)
	return result, nil
}
