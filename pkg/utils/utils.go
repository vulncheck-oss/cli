package utils

import (
	"archive/zip"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// NormalizeString normalizes a string by converting it to lowercase and replacing spaces and underscores with hyphens.
func NormalizeString(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")

	return s
}

// ExtractFile extracts the file name from a URL string from an index backup.
func ExtractFile(urlStr string) (string, error) {

	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	path := strings.TrimPrefix(parsedUrl.Path, "/")
	if !strings.HasSuffix(path, ".zip") {
		return "", fmt.Errorf("invalid file format")
	}

	return path, nil
}

// ParseDate parses a date string in RFC3339 format and returns a formatted string.
func ParseDate(date string) string {
	dateAdded, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s, %s, %s",
		dateAdded.Format("January 2, 2006"), // Format date
		dateAdded.Format("3:04:05 pm"),      // Format time (lowercase pm)
		dateAdded.Format("MST"),             // Format timezone
	)
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	if err := os.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	for _, f := range r.File {
		err := extractZipFile(f, dest)
		if err != nil {
			return fmt.Errorf("failed to extract file: %w", err)
		}
	}

	return nil
}

func extractZipFile(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	path := filepath.Join(dest, f.Name)

	// Check for ZipSlip vulnerability
	if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
		return fmt.Errorf("illegal file path: %s", path)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	} else {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetDirectorySize calculates the total size of a directory and returns it as a human-readable string
func GetDirectorySize(path string) (uint64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		return 0, err
	}
	return uint64(size), nil
}

func GetSizeHuman(size uint64) string {
	return humanize.Bytes(size)
}

func GetDateHuman(date time.Time) string {
	return humanize.Time(date)
}
