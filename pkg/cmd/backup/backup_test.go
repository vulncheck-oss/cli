package backup

import (
	"github.com/vulncheck-oss/cli/pkg/utils"
	"testing"
)

func TestExtractFileBasename(t *testing.T) {
	tests := []struct {
		name    string
		urlStr  string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid URL with zip file",
			urlStr:  "http://example.com/path/to/file.zip",
			want:    "file.zip",
			wantErr: false,
		},
		{
			name:    "Invalid file format",
			urlStr:  "http://example.com/path/to/file.txt",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid URL",
			urlStr:  "http://a b.com/",
			want:    "",
			wantErr: true,
		},
		{
			name:    "URL with nested path",
			urlStr:  "http://example.com/deeply/nested/path/archive.zip",
			want:    "archive.zip",
			wantErr: false,
		},
		{
			name:    "URL with query parameters",
			urlStr:  "http://example.com/path/to/file.zip?token=abc123",
			want:    "file.zip",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.ExtractFileBasename(tt.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractFileBasename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractFileBasename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected string
	}{
		{
			name:     "Valid Date",
			date:     "2024-03-22T04:48:38.601Z",
			expected: "March 22, 2024, 4:48:38 am, UTC",
		},
		{
			name:     "Invalid Date",
			date:     "This is not a date",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.ParseDate(tt.date); got != tt.expected {
				t.Errorf("parseDate(%v) = %v, want %v", tt.date, got, tt.expected)
			}
		})
	}
}
