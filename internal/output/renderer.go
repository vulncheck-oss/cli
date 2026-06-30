// Package output centralises CLI output routing.
//
// Two design rules drive the package:
//
//  1. In JSON mode, stdout carries ONLY the JSON payload (or a structured
//     error envelope). Every status line, spinner, prompt, progress bar,
//     and human-readable note goes to stderr. Agents and `jq` pipelines
//     can rely on `vulncheck <cmd> --json` producing parseable output.
//
//  2. In text mode, everything goes to stdout, matching today's behaviour.
//     Styling is suppressed when stdout is not a TTY or when NO_COLOR is set.
//
// Commands obtain a Renderer once and call its methods instead of reaching
// for fmt.Print / pkg/ui directly. The Renderer is the single seam between
// "thing happened" and "bytes on a file descriptor".
package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Mode selects the output shape.
type Mode int

const (
	// ModeText emits human-readable, optionally styled output. Default.
	ModeText Mode = iota
	// ModeJSON emits a single JSON document (or structured error) on stdout.
	// All informational output is routed to stderr.
	ModeJSON
)

func (m Mode) String() string {
	switch m {
	case ModeJSON:
		return "json"
	default:
		return "text"
	}
}

// Renderer routes informational, payload, and error output to the right
// stream based on the selected Mode. Zero value is not useful; use New().
type Renderer struct {
	mode  Mode
	color bool

	stdout io.Writer
	stderr io.Writer

	// quiet suppresses informational lines (Info/Success/Stat) entirely.
	// Errors and payloads still render.
	quiet bool
}

// Options controls Renderer construction.
type Options struct {
	Mode   Mode
	Color  bool // when false, callers should render without ANSI sequences
	Quiet  bool
	Stdout io.Writer // defaults to os.Stdout
	Stderr io.Writer // defaults to os.Stderr
}

// New constructs a Renderer from explicit options. Callers that want
// environment defaults should use NewFromEnv.
func New(opts Options) *Renderer {
	r := &Renderer{
		mode:   opts.Mode,
		color:  opts.Color,
		quiet:  opts.Quiet,
		stdout: opts.Stdout,
		stderr: opts.Stderr,
	}
	if r.stdout == nil {
		r.stdout = os.Stdout
	}
	if r.stderr == nil {
		r.stderr = os.Stderr
	}
	return r
}

// NewFromEnv builds a Renderer using TTY / NO_COLOR / CI detection.
// Mode is taken from the caller (typically a --json flag); everything
// else is inferred from the environment.
func NewFromEnv(mode Mode) *Renderer {
	return New(Options{
		Mode:  mode,
		Color: ColorEnabled(),
	})
}

// Mode returns the renderer's mode. Useful for commands that need to
// choose between two payload shapes (table vs JSON) at a single call site.
func (r *Renderer) Mode() Mode { return r.mode }

// IsJSON is the common shorthand: `if r.IsJSON() { … }`.
func (r *Renderer) IsJSON() bool { return r.mode == ModeJSON }

// Color reports whether ANSI styling should be applied. Callers that
// build their own styled strings (lipgloss, etc.) should consult this.
func (r *Renderer) Color() bool { return r.color }

// Quiet reports whether informational output is suppressed.
func (r *Renderer) Quiet() bool { return r.quiet }

// Stdout returns the payload stream. Use for raw payload writes when
// the structured helpers (JSON, Println) are not a fit — e.g. a command
// that streams bytes (backup download to stdout).
func (r *Renderer) Stdout() io.Writer { return r.stdout }

// Stderr returns the diagnostic stream. Use for progress bars, spinners,
// and anything that must not contaminate stdout in JSON mode.
func (r *Renderer) Stderr() io.Writer { return r.stderr }

// InfoStream returns the stream that informational output should go to.
// stderr in JSON mode, stdout otherwise. Useful when wiring a third-party
// component (taskin, bubbletea) that takes an io.Writer.
func (r *Renderer) InfoStream() io.Writer {
	if r.mode == ModeJSON {
		return r.stderr
	}
	return r.stdout
}

// JSON marshals payload as indented JSON and writes it, followed by a
// newline, to stdout. Returns the marshal error so callers can decide
// whether to surface it (typically they should — a marshal failure is
// a programmer error worth seeing).
func (r *Renderer) JSON(payload any) error {
	buf, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return fmt.Errorf("output: marshal json: %w", err)
	}
	if _, err := fmt.Fprintln(r.stdout, string(buf)); err != nil {
		return fmt.Errorf("output: write stdout: %w", err)
	}
	return nil
}

// Println writes a plain text line to the payload stream.
// Has no effect in JSON mode (where stdout is reserved for the JSON document).
func (r *Renderer) Println(args ...any) {
	if r.mode == ModeJSON {
		return
	}
	fmt.Fprintln(r.stdout, args...)
}

// Printf writes a plain text fragment to the payload stream.
// Has no effect in JSON mode.
func (r *Renderer) Printf(format string, args ...any) {
	if r.mode == ModeJSON {
		return
	}
	fmt.Fprintf(r.stdout, format, args...)
}

// Info writes an informational status line. Routed to stderr in JSON mode
// so the stdout payload stays clean. Suppressed entirely when quiet.
func (r *Renderer) Info(format string, args ...any) {
	if r.quiet {
		return
	}
	fmt.Fprintln(r.InfoStream(), format2(format, args...))
}

// Success writes a positive status line. Same routing as Info.
func (r *Renderer) Success(format string, args ...any) {
	if r.quiet {
		return
	}
	fmt.Fprintln(r.InfoStream(), format2(format, args...))
}

// Stat writes a "label: value" pair. Same routing as Info.
func (r *Renderer) Stat(label, value string) {
	if r.quiet {
		return
	}
	fmt.Fprintf(r.InfoStream(), "%s: %s\n", label, value)
}

// Warn writes a warning line. Always emitted (even when quiet), since a
// warning that the user doesn't see is rarely the right call.
// Routed to stderr regardless of mode.
func (r *Renderer) Warn(format string, args ...any) {
	fmt.Fprintln(r.stderr, format2(format, args...))
}

// format2 lets callers pass either a preformatted string or a
// Printf-style format + args, without forcing two methods.
func format2(format string, args ...any) string {
	if len(args) == 0 {
		return format
	}
	return fmt.Sprintf(format, args...)
}
