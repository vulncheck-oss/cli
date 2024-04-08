package utils

import "testing"

func TestNormalizeString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "snake_case",
			in:   "snake_case",
			out:  "snake-case",
		},
		{
			name: "kebab-case",
			in:   "kebab-case",
			out:  "kebab-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "UPPERCASE",
			in:   "UPPERCASE",
			out:  "uppercase",
		},
		{
			name: "lowercase",
			in:   "lowercase",
			out:  "lowercase",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
		{
			name: "Title Case",
			in:   "Title Case",
			out:  "title-case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeString(tt.in); got != tt.out {
				t.Errorf("NormalizeString() = %v, want %v", got, tt.out)
			}
		})
	}
}
