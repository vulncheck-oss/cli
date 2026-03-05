package config

import (
	"os"
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
			want:  true,
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

func TestSaveAndLoadConfig(t *testing.T) {
	// Setup: Create a temporary directory for config
	tempDir := t.TempDir()
	homeDir := os.Getenv("HOME") // Save original HOME
	if err := os.Setenv("HOME", tempDir); err != nil {
		t.Fatalf("Failed to set HOME env var: %v", err)
	}
	defer func() {
		if err := os.Setenv("HOME", homeDir); err != nil {
			t.Errorf("Failed to restore HOME env var: %v", err)
		}
	}()

	expectedToken := "vulncheck_testtoken1234567890abcdefghijklmnopqrstuvw"
	config := &Config{Token: expectedToken}
	err := saveConfig(config)
	if err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	loadedConfig, err := loadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if loadedConfig.Token != expectedToken {
		t.Errorf("Expected token %s, got %s", expectedToken, loadedConfig.Token)
	}
}
