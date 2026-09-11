package sync

import (
	"strings"
	"testing"
)

// A bare index name used to be dropped by cobra, leaving the sync to
// re-download the already-cached indices instead.
func TestArgsRejectsPositionalIndex(t *testing.T) {
	cmd := Command()

	err := cmd.Args(cmd, []string{"vulncheck-canaries-3d"})
	if err == nil {
		t.Fatal("expected an error for a positional index name, got nil")
	}
	if !strings.Contains(err.Error(), "--add vulncheck-canaries-3d") {
		t.Errorf("expected the error to suggest --add, got: %v", err)
	}
}

func TestArgsAllowsNoArgs(t *testing.T) {
	cmd := Command()

	if err := cmd.Args(cmd, []string{}); err != nil {
		t.Errorf("expected no error without args, got: %v", err)
	}
}
