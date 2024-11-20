package utils

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

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

func TestExtractFile(t *testing.T) {
	tests := []struct {
		name    string
		urlStr  string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid URL",
			urlStr:  "https://example.com/path/to/file.zip",
			want:    "path/to/file.zip",
			wantErr: false,
		},
		{
			name:    "Invalid URL",
			urlStr:  "://invalid-url",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Non-zip file",
			urlStr:  "https://example.com/path/to/file.txt",
			want:    "",
			wantErr: true,
		},
		{
			name:    "URL with query parameters",
			urlStr:  "https://example.com/path/to/file.zip?param=value",
			want:    "path/to/file.zip",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractFile(tt.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		name string
		date string
		want string
	}{
		{
			name: "Valid RFC3339 date (PM)",
			date: "2023-05-15T14:30:00Z",
			want: "May 15, 2023, 2:30:00 pm, UTC",
		},
		{
			name: "Valid RFC3339 date (AM)",
			date: "2023-05-15T02:30:00Z",
			want: "May 15, 2023, 2:30:00 am, UTC",
		},
		{
			name: "Invalid date format",
			date: "2023-05-15",
			want: "",
		},
		{
			name: "Empty string",
			date: "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDate(tt.date); got != tt.want {
				t.Errorf("ParseDate():  %v != %v", got, tt.want)
			}
		})
	}
}

func TestUnzip(t *testing.T) {
	// Create a temporary directory for the test
	tempDir := os.TempDir()

	// Create a test zip file
	zipPath := filepath.Join(tempDir, "test.zip")
	if err := createTestZip(zipPath); err != nil {
		t.Fatalf("Failed to create test zip: %v", err)
	}

	// Create a destination directory with appropriate permissions
	destDir := filepath.Join(tempDir, "dest")
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create destination directory: %v", err)
	}

	// Test the Unzip function
	err = Unzip(zipPath, destDir)
	if err != nil {
		t.Fatalf("Unzip failed: %v", err)
	}

	// Check if files were extracted correctly
	expectedFiles := []string{"file1.txt", "dir/file2.txt"}
	for _, file := range expectedFiles {
		path := filepath.Join(destDir, file)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("Expected file %s does not exist", file)
		}
	}
}

func createTestZip(zipPath string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add a file to the root of the zip
	file1, err := zipWriter.Create("file1.txt")
	if err != nil {
		return err
	}
	_, err = file1.Write([]byte("Content of file 1"))
	if err != nil {
		return err
	}

	// Add a file in a subdirectory
	file2, err := zipWriter.Create("dir/file2.txt")
	if err != nil {
		return err
	}
	_, err = file2.Write([]byte("Content of file 2"))
	if err != nil {
		return err
	}

	return nil
}
