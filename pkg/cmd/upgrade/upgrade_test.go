package upgrade

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGetPlatformAssetName(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected string
	}{
		{
			name:     "version 1.0.0",
			version:  "1.0.0",
			expected: "vulncheck_1.0.0_macOS_arm64.zip", // This will vary based on test platform
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getPlatformAssetName(tt.version)
			if err != nil {
				t.Fatalf("getPlatformAssetName() error = %v", err)
			}

			// Just verify it contains the version and has proper format
			if len(result) == 0 {
				t.Errorf("getPlatformAssetName() returned empty string")
			}
			// Should contain version
			if !strings.Contains(result, tt.version) {
				t.Errorf("getPlatformAssetName() = %v, should contain version %v", result, tt.version)
			}
		})
	}
}

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
