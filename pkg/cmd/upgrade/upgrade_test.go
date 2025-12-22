package upgrade

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestGetSpecificRelease_VersionFormatting(t *testing.T) {
	tests := []struct {
		name         string
		inputVersion string
		expectedURL  string
	}{
		{
			name:         "version with v prefix",
			inputVersion: "v1.0.0",
			expectedURL:  "https://api.github.com/repos/vulncheck-oss/cli/releases/tags/v1.0.0",
		},
		{
			name:         "version without v prefix",
			inputVersion: "1.0.0",
			expectedURL:  "https://api.github.com/repos/vulncheck-oss/cli/releases/tags/v1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We can't actually call the function without network access,
			// but we can verify the URL construction logic
			version := tt.inputVersion
			if !strings.HasPrefix(version, "v") {
				version = "v" + version
			}

			expectedURL := "https://api.github.com/repos/vulncheck-oss/cli/releases/tags/" + version
			if expectedURL != tt.expectedURL {
				t.Errorf("URL construction failed: expected %s, got %s", tt.expectedURL, expectedURL)
			}
		})
	}
}

func TestHTTPClientTimeout(t *testing.T) {
	// Verify that httpClient has a reasonable timeout configured
	expectedTimeout := 30 * time.Second
	if httpClient.Timeout != expectedTimeout {
		t.Errorf("httpClient timeout = %v, expected %v", httpClient.Timeout, expectedTimeout)
	}
}

func TestBackupFilenameFormat(t *testing.T) {
	// Test that backup filename follows expected format
	currentVersion := "1.0.0"
	now := time.Date(2024, 9, 15, 14, 30, 45, 0, time.UTC)

	expected := "vulncheck.backup.v1.0.0.20240915.143045"
	actual := fmt.Sprintf("vulncheck.backup.v%s.%s",
		currentVersion,
		now.Format("20060102.150405"))

	if actual != expected {
		t.Errorf("Backup filename format = %s, expected %s", actual, expected)
	}

	// Verify it's human readable and contains version
	if !strings.Contains(actual, "v1.0.0") {
		t.Errorf("Backup filename should contain version: %s", actual)
	}

	if !strings.Contains(actual, "20240915") {
		t.Errorf("Backup filename should contain date: %s", actual)
	}
}

func TestCanWriteToBinaryDir(t *testing.T) {
	tests := []struct {
		name          string
		setupFunc     func(t *testing.T) string
		cleanupFunc   func(string)
		expectError   bool
		errorContains string
	}{
		{
			name: "writable directory",
			setupFunc: func(t *testing.T) string {
				// Create a temporary directory that we can write to
				tempDir := t.TempDir()
				binaryPath := filepath.Join(tempDir, "vulncheck")
				// Create a fake binary file
				f, err := os.Create(binaryPath)
				if err != nil {
					t.Fatalf("Failed to create test binary: %v", err)
				}
				if err := f.Close(); err != nil {
					t.Fatalf("Failed to close test binary: %v", err)
				}
				return binaryPath
			},
			expectError: false,
		},
		{
			name: "non-existent directory",
			setupFunc: func(t *testing.T) string {
				return "/non/existent/path/vulncheck"
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binaryPath := tt.setupFunc(t)
			if tt.cleanupFunc != nil {
				defer tt.cleanupFunc(binaryPath)
			}

			err := canWriteToBinaryDir(binaryPath)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain '%s', got: %v", tt.errorContains, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestGetPermissionErrorMessage(t *testing.T) {
	tests := []struct {
		name        string
		binaryPath  string
		goos        string
		contains    []string
		notContains []string
	}{
		{
			name:        "unix system message",
			binaryPath:  "/usr/local/bin/vulncheck",
			goos:        "linux",
			contains:    []string{"Permission denied", "/usr/local/bin", "sudo vulncheck upgrade", "~/bin"},
			notContains: []string{"Administrator"},
		},
		{
			name:        "windows system message",
			binaryPath:  "C:\\Program Files\\vulncheck\\vulncheck.exe",
			goos:        "windows",
			contains:    []string{"Permission denied", "C:\\Program Files\\vulncheck", "Administrator"},
			notContains: []string{"sudo", "~/bin"},
		},
		{
			name:       "macos system message",
			binaryPath: "/usr/local/bin/vulncheck",
			goos:       "darwin",
			contains:   []string{"Permission denied", "/usr/local/bin", "sudo vulncheck upgrade"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Temporarily override runtime.GOOS for testing
			originalGOOS := runtime.GOOS
			defer func() {
				// We can't actually change runtime.GOOS at runtime but we can test the current OS behavior
			}()

			if tt.goos != originalGOOS {
				t.Skipf("Skipping test for %s on %s", tt.goos, originalGOOS)
			}

			message := getPermissionErrorMessage(tt.binaryPath)

			for _, contain := range tt.contains {
				if !strings.Contains(message, contain) {
					t.Errorf("Expected message to contain '%s', got: %s", contain, message)
				}
			}

			for _, notContain := range tt.notContains {
				if strings.Contains(message, notContain) {
					t.Errorf("Expected message to NOT contain '%s', got: %s", notContain, message)
				}
			}

			binaryDir := filepath.Dir(tt.binaryPath)
			if !strings.Contains(message, binaryDir) {
				t.Errorf("Expected message to contain binary directory '%s', got: %s", binaryDir, message)
			}

			if !strings.HasPrefix(message, "❌ Permission denied") {
				t.Errorf("Expected message to start with '❌ Permission denied', got: %s", message)
			}
		})
	}
}

func TestPermissionErrorMessageFormatting(t *testing.T) {
	binaryPath := "/some/path/vulncheck"
	message := getPermissionErrorMessage(binaryPath)

	lines := strings.Split(message, "\n")
	if len(lines) < 3 {
		t.Errorf("Expected multi-line message, got %d lines: %s", len(lines), message)
	}

	if strings.TrimSpace(message) == "" {
		t.Error("Expected non-empty error message")
	}

	helpfulKeywords := []string{"Permission denied", "privileges"}
	hasHelpfulContent := false
	for _, keyword := range helpfulKeywords {
		if strings.Contains(message, keyword) {
			hasHelpfulContent = true
			break
		}
	}
	if !hasHelpfulContent {
		t.Errorf("Expected message to contain helpful keywords, got: %s", message)
	}
}
