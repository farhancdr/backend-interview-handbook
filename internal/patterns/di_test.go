package patterns

import (
	"context"
	"testing"
)

func TestDependencyInjection(t *testing.T) {
	// This test verifies that the application is correctly wired
	app := NewApplication()

	if app.UserService == nil {
		t.Fatal("UserService should not be nil")
	}

	// Verify end-to-end functionality
	ctx := context.Background()
	err := app.RunStub(ctx)
	if err != nil {
		t.Fatalf("Application run failed: %v", err)
	}

	// Check side effect
	user, err := app.UserService.GetUser(ctx, "di-1")
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}
	if user.Name != "DI User" {
		t.Errorf("Expected 'DI User', got %s", user.Name)
	}
}
