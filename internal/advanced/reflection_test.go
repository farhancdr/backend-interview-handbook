package advanced

import (
	"strings"
	"testing"
)

type Config struct {
	Host string `required:"true"`
	Port int    `required:"true"`
	Mode string // Optional
}

type NestedStruct struct {
	Name string
	Conf Config
}

func TestValidateStruct(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "Valid Struct",
			input:   Config{Host: "localhost", Port: 8080},
			wantErr: false,
		},
		{
			name:    "Missing Host",
			input:   Config{Port: 8080},
			wantErr: true,
		},
		{
			name:    "Missing Port",
			input:   Config{Host: "localhost"},
			wantErr: true,
		},
		{
			name:    "Pointer Valid",
			input:   &Config{Host: "localhost", Port: 8080},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStruct(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWalkStruct(t *testing.T) {
	input := NestedStruct{
		Name: "Main",
		Conf: Config{Host: "127.0.0.1", Port: 9000},
	}

	results := WalkStruct(input, 0)

	// We expect keys to be present
	expectedKeys := []string{"Name", "Conf", "Host", "Port"}

	for _, key := range expectedKeys {
		found := false
		for _, line := range results {
			if strings.Contains(line, key) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("WalkStruct() output missing key %s", key)
		}
	}
}
