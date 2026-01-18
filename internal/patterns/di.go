package patterns

import "context"

// Why interviewers ask this:
// Dependency Injection (DI) is crucial for building loose-coupled, testable applications.
// In Go, this is typically done via "Constructor Injection". Interviewers want to see
// that you pass dependencies explicitly rather than using globals or init() magic.

// Common pitfalls:
// - Using global variables for dependencies (makes testing hard)
// - Using complex DI frameworks where simple manual wiring suffices
// - Not using interfaces for dependencies

// Key takeaway:
// Prefer explicit constructor injection. Accept interfaces, return structs.
// Wire your application in `main.go`.

// Application container showing how components are wired
type Application struct {
	UserService *UserService
}

// NewApplication "wires" the dependencies
func NewApplication() *Application {
	// 1. Create dependencies (Repositories)
	userRepo := NewInMemoryUserRepository()

	// 2. Inject dependencies into consumers (Services)
	userService := NewUserService(userRepo)

	// 3. Return the fully wired application
	return &Application{
		UserService: userService,
	}
}

// Example usage to prove it works
func (app *Application) RunStub(ctx context.Context) error {
	return app.UserService.RegisterUser(ctx, "di-1", "DI User", "di@example.com")
}
