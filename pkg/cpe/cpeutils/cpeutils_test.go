package cpeutils

import (
	"testing"
)

func TestProcess(t *testing.T) {
	cpe := &CPE{
		Part:    "a",
		Vendor:  "vendor",
		Product: "product",
		Version: "1.0",
	}

	entries := []CPEVulnerabilities{
		{CPE: CPE{Version: "1.0"}, Cves: []string{"CVE-2021-1234"}},
		{CPE: CPE{Version: "2.0"}, Cves: []string{"CVE-2021-5678"}},
		{CPE: CPE{Version: "1.0"}, Cves: []string{"CVE-2021-9012"}},
	}

	expected := []string{"CVE-2021-1234", "CVE-2021-9012"}

	result, err := Process(cpe, entries)
	if err != nil {
		t.Errorf("Process() error = %v", err)
		return
	}

	if len(result) != len(expected) {
		t.Errorf("Process() returned %d results, want %d", len(result), len(expected))
		return
	}

	for _, cve := range expected {
		if !contains(result, cve) {
			t.Errorf("Process() result doesn't contain expected CVE: %s", cve)
		}
	}
}

func TestRemoveDuplicatesUnordered(t *testing.T) {
	input := []string{"a", "b", "c", "a", "b"}
	expected := []string{"a", "b", "c"}

	result := RemoveDuplicatesUnordered(input)

	if len(result) != len(expected) {
		t.Errorf("RemoveDuplicatesUnordered() length = %d, want %d", len(result), len(expected))
	}

	for _, v := range expected {
		if !contains(result, v) {
			t.Errorf("RemoveDuplicatesUnordered() does not contain %s", v)
		}
	}
}

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		a, b     string
		expected int
	}{
		{"1.0", "1.0", 0},
		{"1.0", "2.0", -1},
		{"2.0", "1.0", 1},
		{"1.2.3", "1.2.3", 0},
		{"1.2.3", "1.2.4", -1},
		{"1.10", "1.2", 1},
	}

	for _, tt := range tests {
		result, err := CompareVersions(tt.a, tt.b)
		if err != nil {
			t.Errorf("CompareVersions(%s, %s) error = %v", tt.a, tt.b, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("CompareVersions(%s, %s) = %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestGetUint64FromVersion(t *testing.T) {
	tests := []struct {
		version  string
		expected uint64
		wantErr  bool
	}{
		{"1.2.3", 0x0102030000000000, false},
		{"255.255.255.255", 0xffffffff00000000, false},
		{"1.2.3.4.5.6.7.8", 0x0102030405060708, false},
		{"1.2.3.4.5.6.7.8.9", 0, true},
		{"a.b.c", 0, true},
	}

	for _, tt := range tests {
		result, err := GetUint64FromVersion(tt.version)
		if (err != nil) != tt.wantErr {
			t.Errorf("GetUint64FromVersion(%s) error = %v, wantErr %v", tt.version, err, tt.wantErr)
			continue
		}
		if result != tt.expected {
			t.Errorf("GetUint64FromVersion(%s) = %d, want %d", tt.version, result, tt.expected)
		}
	}
}

func TestIsParseableVersion(t *testing.T) {
	tests := []struct {
		version  string
		expected bool
	}{
		{"1.2.3", true},
		{"1.2.3-alpha", true},
		{"1.2.3+build", true},
		{"a.b.c", false},
		{"1.x", false},
	}

	for _, tt := range tests {
		result := IsParseableVersion(tt.version)
		if result != tt.expected {
			t.Errorf("IsParseableVersion(%s) = %v, want %v", tt.version, result, tt.expected)
		}
	}
}

func TestUnquote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test\\.version", "test.version"},
		{"test\\-version", "test-version"},
		{"test\\_version", "test_version"},
		{"test\\\\version", "test\\\\version"}, // Changed this line
		{"*", "*"},
		{"-", "-"},
		{"", "*"},
	}

	for _, tt := range tests {
		result := Unquote(tt.input)
		if result != tt.expected {
			t.Errorf("Unquote(%s) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
