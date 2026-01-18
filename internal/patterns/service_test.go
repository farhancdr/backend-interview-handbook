package patterns

import (
	"context"
	"testing"
)

func TestUserService_RegisterUser(t *testing.T) {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)
	ctx := context.Background()

	// Scenario 1: Successful Registration
	err := service.RegisterUser(ctx, "1", "Bob", "bob@example.com")
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}

	// Scenario 2: Duplicate User
	err = service.RegisterUser(ctx, "1", "Bob", "bob@example.com")
	if err == nil || err.Error() != "user already exists" {
		t.Errorf("expected 'user already exists', got %v", err)
	}

	// Scenario 3: Validation Error
	err = service.RegisterUser(ctx, "2", "NoEmail", "invalid-email")
	if err == nil || err.Error() != "invalid email" {
		t.Errorf("expected 'invalid email', got %v", err)
	}
}
