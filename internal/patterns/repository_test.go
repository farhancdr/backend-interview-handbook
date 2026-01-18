package patterns

import (
	"context"
	"testing"
)

func TestInMemoryUserRepository(t *testing.T) {
	repo := NewInMemoryUserRepository()
	ctx := context.Background()

	// Test Create
	user := &User{ID: "1", Name: "Alice", Email: "alice@example.com"}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Test Get
	retrieved, err := repo.Get(ctx, "1")
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if retrieved.Name != user.Name {
		t.Errorf("expected name %s, got %s", user.Name, retrieved.Name)
	}

	// Test Get Not Found
	_, err = repo.Get(ctx, "999")
	if err != ErrUserNotFound {
		t.Errorf("expected ErrUserNotFound, got %v", err)
	}

	// Test List
	users, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("failed to list users: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}

	// Test Delete
	if err := repo.Delete(ctx, "1"); err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}

	// Verify Delete
	_, err = repo.Get(ctx, "1")
	if err != ErrUserNotFound {
		t.Error("expected user to be deleted")
	}
}

func TestRepositoryConcurrency(t *testing.T) {
	repo := NewInMemoryUserRepository()
	ctx := context.Background()

	// Run concurrent reads and writes
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func(id int) {
			user := &User{ID: "1", Name: "Alice", Email: "alice@example.com"}
			_ = repo.Create(ctx, user)
			_, _ = repo.Get(ctx, "1")
			done <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-done
	}
}
