package basics

import (
	"errors"
	"testing"
)

func TestError_BasicDivision(t *testing.T) {
	// Successful division
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %f", result)
	}

	// Division by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("expected error for division by zero")
	}
}

func TestError_WithContext(t *testing.T) {
	_, err := DivideWithContext(10, 0)
	if err == nil {
		t.Error("expected error for division by zero")
	}

	expected := "cannot divide 10.000000 by zero"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}

func TestError_ProcessValue(t *testing.T) {
	// Valid value
	err := ProcessValue(50)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Negative value
	err = ProcessValue(-1)
	if err == nil {
		t.Error("expected error for negative value")
	}

	// Value too large
	err = ProcessValue(101)
	if err == nil {
		t.Error("expected error for value > 100")
	}
}

func TestError_MultipleReturns(t *testing.T) {
	// Success
	str, length, err := MultipleReturns("hello")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if str != "hello" {
		t.Errorf("expected hello, got %s", str)
	}
	if length != 5 {
		t.Errorf("expected length 5, got %d", length)
	}

	// Error
	_, _, err = MultipleReturns("")
	if err == nil {
		t.Error("expected error for empty input")
	}
}

func TestError_ChainedOperations(t *testing.T) {
	// Success
	result, err := ChainedOperations(20, 2, 5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("expected 2, got %f", result)
	}

	// First operation fails
	_, err = ChainedOperations(10, 0, 5)
	if err == nil {
		t.Error("expected error from first division")
	}

	// Second operation fails
	_, err = ChainedOperations(10, 2, 0)
	if err == nil {
		t.Error("expected error from second division")
	}
}

func TestError_NilMeansSuccess(t *testing.T) {
	err := NilError()
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestError_CreateError(t *testing.T) {
	// Simple error
	err := CreateError("simple")
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != "simple error" {
		t.Errorf("expected 'simple error', got %s", err.Error())
	}

	// Formatted error
	err = CreateError("formatted")
	if err == nil {
		t.Error("expected error")
	}

	// Nil error
	err = CreateError("nil")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
}

func TestError_ErrorInDefer(t *testing.T) {
	err := ErrorInDefer()
	if err == nil {
		t.Error("expected error")
	}

	// Error should be wrapped by defer
	expected := "deferred: original error"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}

func TestError_SentinelError(t *testing.T) {
	// Not found
	_, err := FindValue(0)
	if err == nil {
		t.Error("expected error")
	}

	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}

	// Found
	value, err := FindValue(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if value != "value-1" {
		t.Errorf("expected value-1, got %s", value)
	}
}

func TestError_CheckErrorType(t *testing.T) {
	// Check ErrNotFound
	if !CheckErrorType(ErrNotFound) {
		t.Error("expected ErrNotFound to match")
	}

	// Check different error
	otherErr := errors.New("other error")
	if CheckErrorType(otherErr) {
		t.Error("expected other error to not match")
	}

	// Check nil
	if CheckErrorType(nil) {
		t.Error("expected nil to not match")
	}
}

func TestError_ErrorString(t *testing.T) {
	err := errors.New("test error")

	// Error implements error interface with Error() method
	if err.Error() != "test error" {
		t.Errorf("expected 'test error', got %s", err.Error())
	}
}

func TestError_NilCheck(t *testing.T) {
	var err error

	// Nil error
	if err != nil {
		t.Error("expected nil error")
	}

	// Non-nil error
	err = errors.New("error")
	if err == nil {
		t.Error("expected non-nil error")
	}
}

func TestError_IdomaticPattern(t *testing.T) {
	// Idiomatic error handling pattern
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Use result
	if result != 5 {
		t.Errorf("expected 5, got %f", result)
	}
}
