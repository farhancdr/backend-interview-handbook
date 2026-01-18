package advanced

import (
	"context"
	"fmt"
	"time"
)

// Why interviewers ask this:
// Context is fundamental to Go's concurrency model for cancellation, deadlines,
// and request-scoped values. Understanding context is essential for production
// Go code and demonstrates knowledge of proper resource management.

// Common pitfalls:
// - Not propagating context through call chain
// - Using context.Background() when context is available
// - Storing context in structs (anti-pattern)
// - Not checking context.Done() in long-running operations
// - Misusing context.Value (should be request-scoped only)

// Key takeaway:
// Context carries deadlines, cancellation signals, and request-scoped values.
// Always pass context as first parameter. Check ctx.Done() in loops.
// Use WithCancel, WithTimeout, WithDeadline for control flow.

// DoWorkWithContext demonstrates context-aware operation
func DoWorkWithContext(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// WithTimeout demonstrates timeout pattern
func WithTimeout(duration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel() // Always call cancel to release resources

	return DoWorkWithContext(ctx)
}

// WithCancel demonstrates cancellation pattern
func WithCancel() (string, error) {
	ctx, cancel := context.WithCancel(context.Background())

	// Simulate cancellation after 100ms
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	err := DoWorkWithContext(ctx)
	if err == context.Canceled {
		return "cancelled", err
	}
	return "completed", nil
}

// WithDeadline demonstrates deadline pattern
func WithDeadline(deadline time.Time) error {
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	return DoWorkWithContext(ctx)
}

// ContextValue demonstrates context value usage
type contextKey string

const userIDKey contextKey = "userID"

func WithContextValue(userID string) string {
	ctx := context.WithValue(context.Background(), userIDKey, userID)
	return GetUserID(ctx)
}

func GetUserID(ctx context.Context) string {
	if userID, ok := ctx.Value(userIDKey).(string); ok {
		return userID
	}
	return ""
}

// ChainedContext demonstrates context propagation
func ChainedContext(ctx context.Context) error {
	// Create child context with timeout
	childCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	return DoWorkWithContext(childCtx)
}

// MultipleGoroutinesWithContext demonstrates coordinated cancellation
func MultipleGoroutinesWithContext(ctx context.Context) []string {
	results := make(chan string, 3)

	for i := 0; i < 3; i++ {
		go func(id int) {
			select {
			case <-ctx.Done():
				results <- fmt.Sprintf("worker %d cancelled", id)
			case <-time.After(2 * time.Second):
				results <- fmt.Sprintf("worker %d completed", id)
			}
		}(i)
	}

	// Collect results
	var output []string
	for i := 0; i < 3; i++ {
		output = append(output, <-results)
	}

	return output
}

// LongRunningTask demonstrates proper context checking in loops
func LongRunningTask(ctx context.Context) (int, error) {
	count := 0

	for i := 0; i < 1000; i++ {
		// Check context periodically
		select {
		case <-ctx.Done():
			return count, ctx.Err()
		default:
			// Continue work
			count++
			time.Sleep(1 * time.Millisecond)
		}
	}

	return count, nil
}

// ContextPropagation demonstrates passing context through call chain
func ContextPropagation(ctx context.Context) error {
	return level1(ctx)
}

func level1(ctx context.Context) error {
	return level2(ctx)
}

func level2(ctx context.Context) error {
	return DoWorkWithContext(ctx)
}

// BackgroundVsTODO demonstrates context creation
func BackgroundVsTODO() (context.Context, context.Context) {
	// Use Background for main, init, tests
	bg := context.Background()

	// Use TODO when unsure which context to use
	todo := context.TODO()

	return bg, todo
}

// CancelCauseExample demonstrates context.WithCancelCause (Go 1.20+)
func CancelCauseExample() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Simulate cancellation
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	return DoWorkWithContext(ctx)
}

// TimeoutExample demonstrates real-world timeout usage
func TimeoutExample() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	resultCh := make(chan string, 1)
	errCh := make(chan error, 1)

	go func() {
		// Simulate work
		time.Sleep(200 * time.Millisecond)
		resultCh <- "result"
	}()

	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
