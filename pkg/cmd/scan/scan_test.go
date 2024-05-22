package scan

import (
	"testing"
)

func TestFormatSingleDecimal(t *testing.T) {
	tests := []struct {
		name     string
		value    *float32
		expected string
	}{
		{
			name:     "positive float",
			value:    float32Ptr(3.14159),
			expected: "3.1",
		},
		{
			name:     "negative float",
			value:    float32Ptr(-2.71828),
			expected: "-2.7",
		},
		{
			name:     "zero",
			value:    float32Ptr(0),
			expected: "0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatSingleDecimal(tt.value); got != tt.expected {
				t.Errorf("formatSingleDecimal() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func float32Ptr(v float32) *float32 {
	return &v
}
