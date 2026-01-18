package leetcode

import "testing"

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"valid simple", "()", true},
		{"valid mixed", "()[]{}", true},
		{"valid nested", "{[]}", true},
		{"invalid mismatch", "(]", false},
		{"invalid ordering", "([)]", false},
		{"invalid single", "[", false},
		{"invalid single close", "]", false},
		{"empty", "", true},
		{"complex valid", "(([]){})", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidParentheses(tt.s); got != tt.want {
				t.Errorf("IsValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
