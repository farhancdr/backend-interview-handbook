package patterns

import (
	"context"
	"errors"
	"sync"
)

// Why interviewers ask this:
// The Repository pattern is fundamental for abstracting data access logic from business logic.
// Interviewers want to see if you can create a clean separation of concerns, making the
// application easier to test and maintain. It allows swapping data sources (e.g., SQL to NoSQL,
// or Memory for testing) without changing business rules.

// Common pitfalls:
// - Leaking database-specific details (like sql.Rows) into the service layer
// - Not using Context for cancellation and timeouts
// - Creating a "God Repository" that does too much
// - Not handling "not found" vs "error" cases clearly

// Key takeaway:
// Define a repository interface that speaks the domain language (e.g., GetUser, SaveOrder).
// Implement it for your concrete data store. Use the interface in your service layer.

var (
	ErrUserNotFound = errors.New("user not found")
)

// User represents a domain entity
type User struct {
	ID    string
	Name  string
	Email string
}

// UserRepository defines the interface for data access
// This is what the service layer will depend on
type UserRepository interface {
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*User, error)
}

// InMemoryUserRepository is a concrete implementation useful for testing
// In a real app, you would have SQLUserRepository, MongoUserRepository, etc.
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*User),
	}
}

func (r *InMemoryUserRepository) Get(ctx context.Context, id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	// Simulate context check
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	
	if user, ok := r.users[id]; ok {
		// Return a copy to prevent external modification affecting the store
		return &User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}, nil
	}
	
	return nil, ErrUserNotFound
}

func (r *InMemoryUserRepository) Create(ctx context.Context, user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	
	if user.ID == "" {
		return errors.New("user ID required")
	}
	
	// Store a copy
	r.users[user.ID] = &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	
	if _, ok := r.users[id]; !ok {
		return ErrUserNotFound
	}
	
	delete(r.users, id)
	return nil
}

func (r *InMemoryUserRepository) List(ctx context.Context) ([]*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	
	users := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, &User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	
	return users, nil
}
