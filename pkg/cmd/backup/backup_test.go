package backup

import (
	"testing"
)

func TestExtractFile(t *testing.T) {
	tests := []struct {
		name    string
		urlStr  string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid URL with zip file",
			urlStr:  "http://example.com/path/to/file.zip",
			want:    "path/to/file.zip",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractFile(tt.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractFile() = %v, want %v", got, tt.want)
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
			if got := parseDate(tt.date); got != tt.expected {
				t.Errorf("parseDate(%v) = %v, want %v", tt.date, got, tt.expected)
			}
		})
	}
}
