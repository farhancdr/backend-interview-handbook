package patterns

import (
	"context"
	"errors"
	"strings"
)

// Why interviewers ask this:
// The Service Layer pattern puts business logic in one place, separate from
// HTTP handlers or database code. Interviewers look for "thin handlers, fat services".
// It ensures that business rules (like "email must be valid" or "user must be unique")
// are enforced regardless of how the application is accessed (API, CLI, etc.).

// Common pitfalls:
// - Mixing business logic with HTTP transport details (e.g., reading JSON in service)
// - Depending on concrete repository types instead of interfaces
// - Not propagating context for cancellation
// - Returning HTTP-specific errors instead of domain errors

// Key takeaway:
// Services should depend on Repository Interfaces. They implement the core business
// rules and workflows. They should accept and return domain objects/errors, not HTTP types.

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// RegisterUser implements the registration workflow
func (s *UserService) RegisterUser(ctx context.Context, id, name, email string) error {
	// 1. Validation Logic
	if id == "" || name == "" || email == "" {
		return errors.New("missing fields")
	}

	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}

	// 2. Business Rule: Check for existence
	existing, err := s.repo.Get(ctx, id)
	if err == nil && existing != nil {
		return errors.New("user already exists")
	}
	if err != nil && err != ErrUserNotFound {
		return err // Return internal repository errors as-is
	}

	// 3. Create User
	user := &User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	return s.repo.Get(ctx, id)
}
