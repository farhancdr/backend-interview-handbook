package advanced

// Why interviewers ask this:
// Functional options pattern is idiomatic Go for clean, extensible APIs.
// It demonstrates understanding of closures, variadic functions, and API design.
// Widely used in production Go code (e.g., gRPC, many libraries).

// Common pitfalls:
// - Not providing sensible defaults
// - Making all options required (defeats the purpose)
// - Not validating options
// - Overcomplicating simple configurations
// - Not documenting what each option does

// Key takeaway:
// Functional options use variadic functions and closures for clean, backward-compatible APIs.
// Each option is a function that modifies the config. Provides flexibility without breaking changes.
// Pattern: type Option func(*Config), func New(opts ...Option) *Type

// Server represents a server with configuration
type Server struct {
	host    string
	port    int
	timeout int
	maxConn int
	tls     bool
}

// Option is a functional option for Server
type Option func(*Server)

// WithHost sets the host
func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

// WithPort sets the port
func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// WithServerTimeout sets the timeout in seconds
func WithServerTimeout(timeout int) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithMaxConnections sets max connections
func WithMaxConnections(maxConn int) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

// WithTLS enables TLS
func WithTLS(enabled bool) Option {
	return func(s *Server) {
		s.tls = enabled
	}
}

// NewServer creates a new server with options
func NewServer(opts ...Option) *Server {
	// Default configuration
	s := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30,
		maxConn: 100,
		tls:     false,
	}

	// Apply options
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Getters
func (s *Server) Host() string { return s.host }
func (s *Server) Port() int    { return s.port }
func (s *Server) Timeout() int { return s.timeout }
func (s *Server) MaxConn() int { return s.maxConn }
func (s *Server) TLS() bool    { return s.tls }

// Database represents a database connection
type Database struct {
	driver   string
	host     string
	port     int
	username string
	password string
	dbName   string
	poolSize int
}

// DBOption is a functional option for Database
type DBOption func(*Database)

// WithDriver sets the database driver
func WithDriver(driver string) DBOption {
	return func(db *Database) {
		db.driver = driver
	}
}

// WithDBHost sets the database host
func WithDBHost(host string) DBOption {
	return func(db *Database) {
		db.host = host
	}
}

// WithDBPort sets the database port
func WithDBPort(port int) DBOption {
	return func(db *Database) {
		db.port = port
	}
}

// WithCredentials sets username and password
func WithCredentials(username, password string) DBOption {
	return func(db *Database) {
		db.username = username
		db.password = password
	}
}

// WithDatabaseName sets the database name
func WithDatabaseName(name string) DBOption {
	return func(db *Database) {
		db.dbName = name
	}
}

// WithPoolSize sets connection pool size
func WithPoolSize(size int) DBOption {
	return func(db *Database) {
		db.poolSize = size
	}
}

// NewDatabase creates a new database connection
func NewDatabase(opts ...DBOption) *Database {
	// Defaults
	db := &Database{
		driver:   "postgres",
		host:     "localhost",
		port:     5432,
		poolSize: 10,
	}

	for _, opt := range opts {
		opt(db)
	}

	return db
}

// Getters
func (db *Database) Driver() string   { return db.driver }
func (db *Database) Host() string     { return db.host }
func (db *Database) Port() int        { return db.port }
func (db *Database) Username() string { return db.username }
func (db *Database) Password() string { return db.password }
func (db *Database) DBName() string   { return db.dbName }
func (db *Database) PoolSize() int    { return db.poolSize }

// Logger represents a logger with configuration
type Logger struct {
	level      string
	output     string
	format     string
	timestamps bool
}

// LoggerOption is a functional option for Logger
type LoggerOption func(*Logger)

// WithLevel sets log level
func WithLevel(level string) LoggerOption {
	return func(l *Logger) {
		l.level = level
	}
}

// WithOutput sets output destination
func WithOutput(output string) LoggerOption {
	return func(l *Logger) {
		l.output = output
	}
}

// WithFormat sets log format
func WithFormat(format string) LoggerOption {
	return func(l *Logger) {
		l.format = format
	}
}

// WithTimestamps enables/disables timestamps
func WithTimestamps(enabled bool) LoggerOption {
	return func(l *Logger) {
		l.timestamps = enabled
	}
}

// NewLogger creates a new logger
func NewLogger(opts ...LoggerOption) *Logger {
	l := &Logger{
		level:      "info",
		output:     "stdout",
		format:     "json",
		timestamps: true,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Getters
func (l *Logger) Level() string    { return l.level }
func (l *Logger) Output() string   { return l.output }
func (l *Logger) Format() string   { return l.format }
func (l *Logger) Timestamps() bool { return l.timestamps }

// ValidationOption demonstrates option validation
func WithValidatedPort(port int) Option {
	return func(s *Server) {
		if port > 0 && port < 65536 {
			s.port = port
		}
		// Invalid ports are ignored (keeps default)
	}
}

// ChainedOptions demonstrates option chaining
func ProductionServer() Option {
	return func(s *Server) {
		WithHost("0.0.0.0")(s)
		WithPort(443)(s)
		WithTLS(true)(s)
		WithMaxConnections(1000)(s)
	}
}
