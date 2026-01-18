package patterns

import (
	"testing"
	"time"
)

func TestFunctionalOptions(t *testing.T) {
	// Should use defaults
	defaultSrv := NewServer()
	if defaultSrv.Port != 8080 {
		t.Errorf("expected default port 8080, got %d", defaultSrv.Port)
	}

	// Should apply options
	customSrv := NewServer(
		WithPort(9090),
		WithTimeout(5*time.Second),
		WithHost("0.0.0.0"),
	)

	if customSrv.Port != 9090 {
		t.Errorf("expected port 9090, got %d", customSrv.Port)
	}
	if customSrv.Timeout != 5*time.Second {
		t.Errorf("expected timeout 5s, got %v", customSrv.Timeout)
	}
	if customSrv.Host != "0.0.0.0" {
		t.Errorf("expected host 0.0.0.0, got %s", customSrv.Host)
	}
}
