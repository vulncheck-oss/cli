package environment

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// Saving the current environment variable state to restore after tests
	originalEnv := os.Getenv("VC_ENV")
	defer os.Setenv("VC_ENV", originalEnv)

	// Test cases
	tests := []struct {
		name          string
		envVar        string
		expectedName  string
		expectToMatch bool
	}{
		{
			name:          "Setting to production",
			envVar:        "production",
			expectedName:  "production",
			expectToMatch: true,
		},
		{
			name:          "Setting to staging",
			envVar:        "staging",
			expectedName:  "staging",
			expectToMatch: true,
		},
		{
			name:          "Setting to development",
			envVar:        "dev",
			expectedName:  "dev",
			expectToMatch: true,
		},
		{
			name:          "Setting to an unknown environment",
			envVar:        "unknown",
			expectedName:  "development", // It should default to the current 'Env' setting which is [2] development at the start.
			expectToMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setting VC_ENV
			if err := os.Setenv("VC_ENV", tt.envVar); err != nil {
				t.Fatalf("Failed to set environment variable: %s", err)
			}
			Init() // Initialize based on the new environment variable

			// Checking if the environment was set correctly
			if (os.Getenv("VC_ENV") == tt.expectedName) != tt.expectToMatch {
				t.Errorf("Expected Env.Name to be '%s', got '%s'", tt.envVar, tt.expectedName)
			}
		})
	}
}
