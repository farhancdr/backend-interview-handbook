# 🛠️ Design Patterns

> **Production-ready patterns for backend systems**

This package implements common design patterns used in real-world Go applications. These patterns appear frequently in interviews and production code.

---

## 📖 Topics

| Topic | File | Key Concepts |
|:------|:-----|:-------------|
| **Repository Pattern** | [repository.go](repository.go) | Data access abstraction, dependency inversion, testing |
| **Middleware** | [middleware.go](middleware.go) | HTTP middleware, chain of responsibility, cross-cutting concerns |
| **Functional Options** | [functional_options.go](functional_options.go) | Builder pattern, optional parameters, API design |
| **Dependency Injection** | [di.go](di.go) | Constructor injection, interface-based design, testability |
| **Circuit Breaker** | [circuit_breaker.go](circuit_breaker.go) | Fault tolerance, failure detection, automatic recovery |
| **Retry Pattern** | [retry.go](retry.go) | Exponential backoff, jitter, idempotency |
| **Service Layer** | [service.go](service.go) | Business logic separation, transaction management |

---

## 🚀 Quick Start

```bash
# Run all tests
go test -v ./patterns/

# Run specific pattern
go test -v ./patterns/ -run TestCircuitBreaker

# Check test coverage
go test -cover ./patterns/
```

---

## 🎓 Learning Guide

### Recommended Order
1. **Dependency Injection** - Foundation for testable code
2. **Repository Pattern** - Data access abstraction
3. **Service Layer** - Business logic organization
4. **Middleware** - Cross-cutting concerns
5. **Functional Options** - Clean API design
6. **Circuit Breaker** - Fault tolerance
7. **Retry Pattern** - Resilience

### Common Interview Questions

**Repository Pattern:**
- "How do you abstract database access?"
- "Why use interfaces for repositories?"
- "How do you test code that uses databases?"

**Middleware:**
- "How do you implement HTTP middleware in Go?"
- "Explain the chain of responsibility pattern"
- "How do you handle errors in middleware?"

**Circuit Breaker:**
- "What problem does circuit breaker solve?"
- "Explain the three states: Closed, Open, Half-Open"
- "How do you implement circuit breaker in a distributed system?"

---

## 💡 Key Takeaways

### Repository Pattern
```go
// Interface for data access
type UserRepository interface {
    GetByID(id string) (*User, error)
    Save(user *User) error
}

// Benefits:
// - Testable (mock repository)
// - Database-agnostic
// - Single responsibility
```

### Middleware Pattern
```go
// Middleware signature
type Middleware func(http.Handler) http.Handler

// Chain middleware
handler := LoggingMiddleware(
    AuthMiddleware(
        RateLimitMiddleware(
            myHandler,
        ),
    ),
)

// Common use cases:
// - Logging
// - Authentication
// - Rate limiting
// - CORS
// - Request ID
```

### Functional Options
```go
// Clean API with optional parameters
type Server struct {
    host string
    port int
    timeout time.Duration
}

type Option func(*Server)

func WithPort(port int) Option {
    return func(s *Server) { s.port = port }
}

// Usage
server := NewServer(
    WithPort(8080),
    WithTimeout(30*time.Second),
)

// Benefits:
// - Backward compatible
// - Self-documenting
// - Optional parameters
```

### Circuit Breaker
```go
// Three states:
// 1. Closed - Normal operation
// 2. Open - Fast fail (circuit "tripped")
// 3. Half-Open - Testing if service recovered

// Prevents:
// - Cascading failures
// - Resource exhaustion
// - Slow responses

// Use when:
// - Calling external services
// - Database queries
// - Any operation that can fail
```

### Retry Pattern
```go
// Exponential backoff with jitter
func retry(fn func() error, maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        if err := fn(); err == nil {
            return nil
        }
        
        // Exponential backoff: 2^i * 100ms
        backoff := time.Duration(1<<i) * 100 * time.Millisecond
        
        // Add jitter to prevent thundering herd
        jitter := time.Duration(rand.Int63n(int64(backoff)))
        
        time.Sleep(backoff + jitter)
    }
    return errors.New("max retries exceeded")
}

// CRITICAL: Only retry idempotent operations!
```

---

## ⚠️ Common Pitfalls

1. **Over-engineering** - Don't add patterns you don't need
2. **Repository Leakage** - Don't expose database types in repository interface
3. **Middleware Order** - Order matters! Auth before rate limiting
4. **Circuit Breaker Tuning** - Threshold and timeout need careful tuning
5. **Retry Without Idempotency** - Can cause duplicate operations

---

## 🔗 Related Topics

- **[System Design](../system_design/)** - Apply patterns to system primitives
- **[Concurrency](../concurrency/)** - Thread-safe pattern implementations
- **[Advanced](../advanced/)** - Context usage in patterns

---

## 📚 Real-World Usage

These patterns are used in popular Go libraries:

- **Functional Options**: gRPC, Uber's Zap logger
- **Middleware**: Gin, Echo, Chi routers
- **Repository**: Most ORMs and data access layers
- **Circuit Breaker**: Netflix Hystrix, Sony Gobreaker
- **Retry**: Hashicorp's go-retryablehttp

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
