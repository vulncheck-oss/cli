package selfupgrade

import (
	"testing"
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
			expected: "vulncheck_1.0.0_macOS_arm64.zip",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getPlatformAssetName(tt.version)
			if err != nil {
				t.Fatalf("getPlatformAssetName() error = %v", err)
			}

			if len(result) == 0 {
				t.Errorf("getPlatformAssetName() returned empty string")
			}
			// Should contain version
			if !contains(result, tt.version) {
				t.Errorf("getPlatformAssetName() = %v, should contain version %v", result, tt.version)
			}
		})
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
