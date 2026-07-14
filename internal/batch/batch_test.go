package batch

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCollectInputsFromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "inputs.txt")
	body := "pkg:foo/bar\n\n# this is a comment\npkg:baz/qux\n"
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := CollectInputs(nil, path, nil)
	if err != nil {
		t.Fatalf("CollectInputs: %v", err)
	}
	want := []string{"pkg:foo/bar", "pkg:baz/qux"}
	if len(got) != len(want) {
		t.Fatalf("len(got)=%d want %d (got=%v)", len(got), len(want), got)
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("got[%d]=%q want %q", i, got[i], w)
		}
	}
}

func TestCollectInputsArgsWinOverStdin(t *testing.T) {
	got, err := CollectInputs([]string{"a", "b"}, "", strings.NewReader("c\nd\n"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("expected [a b], got %v", got)
	}
}

func TestCollectInputsFromStdin(t *testing.T) {
	got, err := CollectInputs(nil, "", strings.NewReader("foo\nbar\n"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "foo" || got[1] != "bar" {
		t.Fatalf("expected [foo bar], got %v", got)
	}
}

func TestCollectInputsSkipsBlankAndComments(t *testing.T) {
	got, err := CollectInputs(nil, "", strings.NewReader("# hi\n\nfoo\n   \nbar\n"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0] != "foo" || got[1] != "bar" {
		t.Fatalf("expected [foo bar], got %v", got)
	}
}

func TestCollectInputsMissingFile(t *testing.T) {
	_, err := CollectInputs(nil, "/no/such/file", nil)
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
