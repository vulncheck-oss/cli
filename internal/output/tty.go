package output

import (
	"os"

	"golang.org/x/term"
)

// IsTTY reports whether the given file is attached to a terminal.
func IsTTY(f *os.File) bool {
	if f == nil {
		return false
	}
	return term.IsTerminal(int(f.Fd()))
}

// ColorEnabled reports whether ANSI colour output should be emitted on stdout.
// Honours NO_COLOR (https://no-color.org) and TERM=dumb, and falls back to
// "stdout is a terminal" otherwise.
func ColorEnabled() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	return IsTTY(os.Stdout)
}

// IsCI reports whether the process appears to be running under a CI system.
// Mirrors the heuristic used elsewhere in the codebase; kept here so the
// output package has no dependency on pkg/config.
func IsCI() bool {
	return os.Getenv("CI") != "" ||
		os.Getenv("BUILD_NUMBER") != "" ||
		os.Getenv("RUN_ID") != ""
}

// Interactive reports whether the process can safely show interactive UI
// (prompts, spinners, full-screen TUIs). False whenever stdin or stdout is
// non-TTY, or when CI/NO_COLOR-style signals are set.
func Interactive() bool {
	if IsCI() {
		return false
	}
	if !IsTTY(os.Stdin) || !IsTTY(os.Stdout) {
		return false
	}
	return true
}
