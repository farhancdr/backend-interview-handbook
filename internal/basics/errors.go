package basics

import (
	"errors"
	"fmt"
)

// Why interviewers ask this:
// Error handling is fundamental to Go. Unlike exceptions in other languages,
// Go uses explicit error returns. Interviewers want to ensure you understand
// idiomatic error handling, error creation, and nil error checks.

// Common pitfalls:
// - Not checking errors (ignoring return values)
// - Comparing errors with == instead of using errors.Is
// - Returning uninitialized errors
// - Not providing context in error messages
// - Confusion about nil errors

// Key takeaway:
// Errors are values. Always check errors. Return errors explicitly.
// nil error means success. Use errors.New or fmt.Errorf to create errors.

// Divide performs division and returns an error if divisor is zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// DivideWithContext returns an error with more context
func DivideWithContext(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %f by zero", a)
	}
	return a / b, nil
}

// ProcessValue demonstrates error checking pattern
func ProcessValue(value int) error {
	if value < 0 {
		return errors.New("value must be non-negative")
	}
	if value > 100 {
		return errors.New("value must not exceed 100")
	}
	// Success
	return nil
}

// MultipleReturns demonstrates multiple return values with error
func MultipleReturns(input string) (string, int, error) {
	if input == "" {
		return "", 0, errors.New("input cannot be empty")
	}
	return input, len(input), nil
}

// ChainedOperations demonstrates error handling in sequence
func ChainedOperations(a, b, c float64) (float64, error) {
	result1, err := Divide(a, b)
	if err != nil {
		return 0, err
	}

	result2, err := Divide(result1, c)
	if err != nil {
		return 0, err
	}

	return result2, nil
}

// IgnoreError demonstrates intentionally ignoring an error (use _ carefully)
func IgnoreError() {
	_, _ = Divide(10, 2) // Intentionally ignoring error (not recommended)
}

// NilError demonstrates that nil means success
func NilError() error {
	return nil // Success
}

// CreateError demonstrates different ways to create errors
func CreateError(errorType string) error {
	switch errorType {
	case "simple":
		return errors.New("simple error")
	case "formatted":
		return fmt.Errorf("formatted error: %s", "details")
	case "nil":
		return nil
	default:
		return fmt.Errorf("unknown error type: %s", errorType)
	}
}

// ErrorInDefer demonstrates error handling in defer
func ErrorInDefer() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("deferred: %w", err)
		}
	}()

	return errors.New("original error")
}

// SentinelError is a predefined error for comparison
var ErrNotFound = errors.New("not found")

// FindValue demonstrates sentinel errors
func FindValue(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	return fmt.Sprintf("value-%d", id), nil
}

// CheckErrorType demonstrates error type checking
func CheckErrorType(err error) bool {
	return err == ErrNotFound
}
