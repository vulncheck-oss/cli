package config

import (
	"testing"
)

func TestValidToken(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  bool
	}{
		{
			name:  "valid token",
			token: "vulncheck_1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			want:  true,
		},
		{
			name:  "no token",
			token: "",
			want:  false,
		},
		{
			name:  "invalid token",
			token: "checkvuln_1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidToken(tt.token); got != tt.want {
				t.Errorf("ValidToken() = %v, want %v", got, tt.want)
			}
		})
	}

}
