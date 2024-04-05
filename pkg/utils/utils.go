package utils

import "strings"

// NormalizeString normalizes a string by converting it to lowercase and replacing spaces and underscores with hyphens.
func NormalizeString(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")

	return s
}
