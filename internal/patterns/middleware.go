package patterns

import (
	"context"
	"errors"
	"log"
)

// Why interviewers ask this:
// Middleware is essential for cross-cutting concerns (logging, auth, metrics).
// Interviewers want to see if you understand how to wrap functions to execute code check
// before and after the main logic. This is commonly seen in `net/http` handlers.

// Common pitfalls:
// - Not calling the next handler (breaking the chain)
// - Modifying state that isn't thread-safe
// - Losing context values

// Key takeaway:
// Middleware is a function that takes a handler and returns a handler:
// `func(Handler) Handler`. It allows you to compose behavior like onion layers.

// Handler defines the clear business logic signature
type Handler func(ctx context.Context, input string) error

// Middleware defines the wrapper signature
type Middleware func(Handler) Handler

// ChainMiddleware composes multiple middlewares
func ChainMiddleware(h Handler, mws ...Middleware) Handler {
	// Apply in reverse order so the first in list is the outer-most wrapper
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

// LoggingMiddleware logs the start and end of a request
func LoggingMiddleware(logger *log.Logger) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, input string) error {
			logger.Printf("START: %s", input)
			err := next(ctx, input)
			logger.Printf("END: %s (err: %v)", input, err)
			return err
		}
	}
}

// AuthMiddleware simulates checking a context key for authorization
func AuthMiddleware(requiredRole string) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, input string) error {
			role, ok := ctx.Value("role").(string)
			if !ok || role != requiredRole {
				return errors.New("unauthorized")
			}
			return next(ctx, input)
		}
	}
}
