package advanced

import "testing"

func TestNewServer_Defaults(t *testing.T) {
	server := NewServer()

	if server.Host() != "localhost" {
		t.Errorf("expected localhost, got %s", server.Host())
	}

	if server.Port() != 8080 {
		t.Errorf("expected 8080, got %d", server.Port())
	}

	if server.Timeout() != 30 {
		t.Errorf("expected 30, got %d", server.Timeout())
	}

	if server.MaxConn() != 100 {
		t.Errorf("expected 100, got %d", server.MaxConn())
	}

	if server.TLS() != false {
		t.Error("expected TLS to be false")
	}
}

func TestNewServer_WithOptions(t *testing.T) {
	server := NewServer(
		WithHost("example.com"),
		WithPort(443),
		WithTLS(true),
	)

	if server.Host() != "example.com" {
		t.Errorf("expected example.com, got %s", server.Host())
	}

	if server.Port() != 443 {
		t.Errorf("expected 443, got %d", server.Port())
	}

	if !server.TLS() {
		t.Error("expected TLS to be true")
	}

	// Unchanged options should have defaults
	if server.Timeout() != 30 {
		t.Errorf("expected default timeout 30, got %d", server.Timeout())
	}
}

func TestNewServer_SingleOption(t *testing.T) {
	server := NewServer(WithPort(9000))

	if server.Port() != 9000 {
		t.Errorf("expected 9000, got %d", server.Port())
	}

	// Other fields should have defaults
	if server.Host() != "localhost" {
		t.Errorf("expected localhost, got %s", server.Host())
	}
}

func TestNewDatabase_Defaults(t *testing.T) {
	db := NewDatabase()

	if db.Driver() != "postgres" {
		t.Errorf("expected postgres, got %s", db.Driver())
	}

	if db.Host() != "localhost" {
		t.Errorf("expected localhost, got %s", db.Host())
	}

	if db.Port() != 5432 {
		t.Errorf("expected 5432, got %d", db.Port())
	}

	if db.PoolSize() != 10 {
		t.Errorf("expected 10, got %d", db.PoolSize())
	}
}

func TestNewDatabase_WithOptions(t *testing.T) {
	db := NewDatabase(
		WithDriver("mysql"),
		WithDBHost("db.example.com"),
		WithDBPort(3306),
		WithCredentials("user", "pass"),
		WithDatabaseName("mydb"),
		WithPoolSize(20),
	)

	if db.Driver() != "mysql" {
		t.Errorf("expected mysql, got %s", db.Driver())
	}

	if db.Host() != "db.example.com" {
		t.Errorf("expected db.example.com, got %s", db.Host())
	}

	if db.Port() != 3306 {
		t.Errorf("expected 3306, got %d", db.Port())
	}

	if db.Username() != "user" {
		t.Errorf("expected user, got %s", db.Username())
	}

	if db.Password() != "pass" {
		t.Errorf("expected pass, got %s", db.Password())
	}

	if db.DBName() != "mydb" {
		t.Errorf("expected mydb, got %s", db.DBName())
	}

	if db.PoolSize() != 20 {
		t.Errorf("expected 20, got %d", db.PoolSize())
	}
}

func TestNewLogger_Defaults(t *testing.T) {
	logger := NewLogger()

	if logger.Level() != "info" {
		t.Errorf("expected info, got %s", logger.Level())
	}

	if logger.Output() != "stdout" {
		t.Errorf("expected stdout, got %s", logger.Output())
	}

	if logger.Format() != "json" {
		t.Errorf("expected json, got %s", logger.Format())
	}

	if !logger.Timestamps() {
		t.Error("expected timestamps to be true")
	}
}

func TestNewLogger_WithOptions(t *testing.T) {
	logger := NewLogger(
		WithLevel("debug"),
		WithOutput("file"),
		WithFormat("text"),
		WithTimestamps(false),
	)

	if logger.Level() != "debug" {
		t.Errorf("expected debug, got %s", logger.Level())
	}

	if logger.Output() != "file" {
		t.Errorf("expected file, got %s", logger.Output())
	}

	if logger.Format() != "text" {
		t.Errorf("expected text, got %s", logger.Format())
	}

	if logger.Timestamps() {
		t.Error("expected timestamps to be false")
	}
}

func TestWithValidatedPort(t *testing.T) {
	// Valid port
	server := NewServer(WithValidatedPort(3000))
	if server.Port() != 3000 {
		t.Errorf("expected 3000, got %d", server.Port())
	}

	// Invalid port (too high)
	server = NewServer(WithValidatedPort(70000))
	if server.Port() != 8080 { // Should keep default
		t.Errorf("expected default 8080, got %d", server.Port())
	}

	// Invalid port (negative)
	server = NewServer(WithValidatedPort(-1))
	if server.Port() != 8080 { // Should keep default
		t.Errorf("expected default 8080, got %d", server.Port())
	}
}

func TestProductionServer(t *testing.T) {
	server := NewServer(ProductionServer())

	if server.Host() != "0.0.0.0" {
		t.Errorf("expected 0.0.0.0, got %s", server.Host())
	}

	if server.Port() != 443 {
		t.Errorf("expected 443, got %d", server.Port())
	}

	if !server.TLS() {
		t.Error("expected TLS to be true")
	}

	if server.MaxConn() != 1000 {
		t.Errorf("expected 1000, got %d", server.MaxConn())
	}
}

func TestFunctionalOptions_Composability(t *testing.T) {
	// Options can be composed and reused
	commonOpts := []Option{
		WithHost("api.example.com"),
		WithTLS(true),
	}

	server1 := NewServer(append(commonOpts, WithPort(443))...)
	server2 := NewServer(append(commonOpts, WithPort(8443))...)

	if server1.Port() != 443 {
		t.Errorf("server1: expected 443, got %d", server1.Port())
	}

	if server2.Port() != 8443 {
		t.Errorf("server2: expected 8443, got %d", server2.Port())
	}

	// Both should have common options
	if server1.Host() != "api.example.com" || server2.Host() != "api.example.com" {
		t.Error("common options not applied correctly")
	}
}
