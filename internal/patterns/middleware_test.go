package patterns

import (
	"bytes"
	"context"
	"errors"
	"log"
	"strings"
	"testing"
)

func TestMiddleware(t *testing.T) {
	// Base Handler
	baseHandler := func(ctx context.Context, input string) error {
		if input == "error" {
			return errors.New("boom")
		}
		return nil
	}

	// Mock Logger
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)

	// Middleware Chain
	chain := ChainMiddleware(
		baseHandler,
		LoggingMiddleware(logger),
		AuthMiddleware("admin"),
	)

	// Scenario 1: Unauthorized
	ctx := context.Background()
	err := chain(ctx, "test")
	if err == nil || err.Error() != "unauthorized" {
		t.Errorf("expected unauthorized, got %v", err)
	}

	// Scenario 2: Authorized + Success
	ctx = context.WithValue(ctx, "role", "admin")
	err = chain(ctx, "success")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	logOutput := buf.String()
	if !strings.Contains(logOutput, "START: success") || !strings.Contains(logOutput, "END: success") {
		t.Errorf("logs missing start/end markers: %s", logOutput)
	}
}
