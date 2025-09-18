package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
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
	defer func() {
		if err := r.Close(); err != nil {
			_ = err
		}
	}()

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
	defer func() {
		if err := rc.Close(); err != nil {
			_ = err
		}
	}()

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
		defer func() {
			if err := f.Close(); err != nil {
				_ = err
			}
		}()

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

func FormatCVE(bareCve string) string {
	// Convert to uppercase
	c := strings.ToUpper(bareCve)

	// Replace all special characters with a dash
	re := regexp.MustCompile(`[^A-Z0-9]`)
	c = re.ReplaceAllString(c, "-")

	// Check if it starts with "CVE-"
	if !strings.HasPrefix(c, "CVE-") {
		c = "CVE-" + c
	}

	// Ensure the format is "CVE-YYYY-NNNN"
	re = regexp.MustCompile(`CVE-(\d{4})-(\d{4,})`)
	matches := re.FindStringSubmatch(c)
	if len(matches) == 3 {
		year := matches[1]
		number := matches[2]
		c = fmt.Sprintf("CVE-%s-%04s", year, number)
	}

	return c
}
