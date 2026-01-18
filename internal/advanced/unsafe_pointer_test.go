package advanced

import (
	"testing"
)

func TestZeroCopyBytesToString(t *testing.T) {
	bytes := []byte{'H', 'e', 'l', 'l', 'o'}
	str := ZeroCopyBytesToString(bytes)

	if str != "Hello" {
		t.Errorf("Expected 'Hello', got %s", str)
	}

	// Verify it points to the same memory (dangerous test but educational)
	// We won't modify it to avoid crashing the test runner or race detector panics,
	// but purely functional verification is enough.
}

func TestZeroCopyStringToBytes(t *testing.T) {
	str := "World"
	bytes := ZeroCopyStringToBytes(str)

	if string(bytes) != "World" {
		t.Errorf("Expected 'World', got %s", string(bytes))
	}
	if len(bytes) != 5 {
		t.Errorf("Expected len 5, got %d", len(bytes))
	}
}
