package patterns

import "time"

// Why interviewers ask this:
// The Functional Options pattern is the idiomatic Go way to handle configuration with
// optional parameters. It avoids having multiple constructors (New, NewWithTimeout, etc.)
// or passing a massive Config struct with nil pointers.

// Common pitfalls:
// - Not making the Option function type exported (if needed)
// - modifying the options directly instead of the target config object
// - Not providing sensible defaults

// Key takeaway:
// Define a function type `type Option func(*Options)`. Create `New(...)` that applies
// defaults, then loops through provided options. It provides a clean, extensible API.

type Server struct {
	Host     string
	Port     int
	Timeout  time.Duration
	MaxConns int
}

type ServerOption func(*Server)

func NewServer(opts ...ServerOption) *Server {
	// 1. Set Defaults
	srv := &Server{
		Host:     "localhost",
		Port:     8080,
		Timeout:  30 * time.Second,
		MaxConns: 100,
	}

	// 2. Apply Options
	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

// WithPort sets the server port
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

// WithTimeout sets the request timeout
func WithTimeout(d time.Duration) ServerOption {
	return func(s *Server) {
		s.Timeout = d
	}
}

// WithHost sets the server host
func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.Host = host
	}
}
