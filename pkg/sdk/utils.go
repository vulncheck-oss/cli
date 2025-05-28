package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

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

// MetaError is a struct that represents the error response from the API
func handleErrorResponse(resp *http.Response) error {
	var metaError MetaError
	_ = json.NewDecoder(resp.Body).Decode(&metaError)

	return ReqError{
		StatusCode: resp.StatusCode,
		Reason:     metaError,
	}
}
