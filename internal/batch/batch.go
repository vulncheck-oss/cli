// Package batch collects multi-input arguments (positional args,
// stdin, or --from-file) and renders per-input results as a stable
// JSON envelope.
//
// The envelope shape is the contract: agents that piped 50 PURLs into
// the CLI can rely on receiving a 50-element array with matching .input
// keys and per-row .data or .error. Never break this shape.
package batch

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Result is one row of the batch output envelope.
type Result struct {
	Input string `json:"input"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// CollectInputs resolves the set of inputs to operate on.
//
//	args      - positional CLI arguments (caller passes cobra's args slice)
//	fromFile  - value of --from-file (empty = unset)
//	stdin     - reader to use for stdin fallback; nil defaults to os.Stdin
//
// Precedence: --from-file wins, then positional args, then stdin if it is
// not a terminal. Empty/whitespace lines and leading '#' comment lines
// are skipped. Returns an empty slice if no source provided any input.
func CollectInputs(args []string, fromFile string, stdin io.Reader) ([]string, error) {
	if fromFile != "" {
		f, err := os.Open(fromFile)
		if err != nil {
			return nil, fmt.Errorf("--from-file: %w", err)
		}
		defer f.Close()
		return readLines(f), nil
	}
	if len(args) > 0 {
		return args, nil
	}
	if stdin == nil {
		stdin = os.Stdin
	}
	// Only consume stdin when it is being piped to us — interactive shells
	// would block forever. Caller is expected to have checked TTY via
	// output.IsTTY(os.Stdin) when wiring this.
	return readLines(stdin), nil
}

func readLines(r io.Reader) []string {
	var lines []string
	sc := bufio.NewScanner(r)
	// 4 MiB lines; large enough for verbose CPE strings.
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

// ResultFromError builds a failure Result. If the error is nil this is
// equivalent to ResultFromData with data=nil — usually you'd call
// ResultFromData directly in that case.
func ResultFromError(input string, err error) Result {
	if err == nil {
		return Result{Input: input}
	}
	return Result{Input: input, Error: err.Error()}
}

// ResultFromData builds a success Result.
func ResultFromData(input string, data any) Result {
	return Result{Input: input, Data: data}
}

// ErrNoInputs is returned by CollectInputs callers when they need to
// surface "you provided nothing" as a validation error.
var ErrNoInputs = errors.New("no inputs provided")
