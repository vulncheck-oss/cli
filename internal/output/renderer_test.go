package output

import (
	"bytes"
	"strings"
	"testing"
)

func newTestRenderer(mode Mode) (*Renderer, *bytes.Buffer, *bytes.Buffer) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	r := New(Options{
		Mode:   mode,
		Stdout: stdout,
		Stderr: stderr,
	})
	return r, stdout, stderr
}

func TestJSONModeKeepsStdoutClean(t *testing.T) {
	r, stdout, stderr := newTestRenderer(ModeJSON)

	r.Info("Listing %d tokens", 3)
	r.Success("done")
	r.Stat("count", "3")

	if stdout.Len() != 0 {
		t.Fatalf("stdout should be empty before JSON payload, got: %q", stdout.String())
	}

	if err := r.JSON(map[string]any{"count": 3}); err != nil {
		t.Fatalf("JSON: %v", err)
	}

	out := stdout.String()
	if !strings.HasPrefix(out, "{") {
		t.Fatalf("stdout should start with JSON, got: %q", out)
	}
	if !strings.Contains(stderr.String(), "Listing 3 tokens") {
		t.Fatalf("info should be on stderr, got: %q", stderr.String())
	}
}

func TestTextModeRoutesInfoToStdout(t *testing.T) {
	r, stdout, stderr := newTestRenderer(ModeText)

	r.Info("hello")
	r.Println("world")

	if !strings.Contains(stdout.String(), "hello") {
		t.Fatalf("info should hit stdout in text mode, got: %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "world") {
		t.Fatalf("println should hit stdout in text mode, got: %q", stdout.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("stderr should be empty in text mode, got: %q", stderr.String())
	}
}

func TestPrintlnSuppressedInJSONMode(t *testing.T) {
	r, stdout, _ := newTestRenderer(ModeJSON)
	r.Println("would corrupt JSON")
	if stdout.Len() != 0 {
		t.Fatalf("Println must not write stdout in JSON mode, got: %q", stdout.String())
	}
}

func TestQuietSuppressesInfoButNotWarnOrJSON(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	r := New(Options{
		Mode:   ModeText,
		Quiet:  true,
		Stdout: stdout,
		Stderr: stderr,
	})

	r.Info("hidden")
	r.Success("also hidden")
	r.Stat("k", "v")
	r.Warn("not hidden")
	if err := r.JSON(map[string]int{"a": 1}); err != nil {
		t.Fatal(err)
	}

	if strings.Contains(stdout.String(), "hidden") {
		t.Fatalf("quiet should suppress info, got stdout: %q", stdout.String())
	}
	if !strings.Contains(stderr.String(), "not hidden") {
		t.Fatalf("quiet must not suppress warnings, got stderr: %q", stderr.String())
	}
	if !strings.Contains(stdout.String(), `"a": 1`) {
		t.Fatalf("quiet must not suppress JSON payload, got stdout: %q", stdout.String())
	}
}

func TestWarnAlwaysHitsStderr(t *testing.T) {
	for _, mode := range []Mode{ModeText, ModeJSON} {
		r, _, stderr := newTestRenderer(mode)
		r.Warn("careful")
		if !strings.Contains(stderr.String(), "careful") {
			t.Fatalf("mode=%s: warn must go to stderr, got: %q", mode, stderr.String())
		}
	}
}

func TestInfoStreamFollowsMode(t *testing.T) {
	rText, stdout, stderr := newTestRenderer(ModeText)
	if rText.InfoStream() != stdout {
		t.Fatal("text mode: InfoStream should be stdout")
	}

	rJSON, _, stderr2 := newTestRenderer(ModeJSON)
	if rJSON.InfoStream() != stderr2 {
		t.Fatal("json mode: InfoStream should be stderr")
	}
	_ = stderr
}

func TestModeString(t *testing.T) {
	if ModeText.String() != "text" {
		t.Errorf("ModeText.String() = %q, want %q", ModeText.String(), "text")
	}
	if ModeJSON.String() != "json" {
		t.Errorf("ModeJSON.String() = %q, want %q", ModeJSON.String(), "json")
	}
}

func TestJSONMarshalError(t *testing.T) {
	r, _, _ := newTestRenderer(ModeJSON)
	// channels can't be marshaled
	err := r.JSON(make(chan int))
	if err == nil {
		t.Fatal("expected marshal error, got nil")
	}
}
