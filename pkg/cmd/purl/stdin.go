package purl

import (
	"io"
	"os"

	"github.com/vulncheck-oss/cli/internal/output"
)

// stdinIfPiped returns os.Stdin when something is being piped into the
// process, and nil otherwise. Returning nil keeps batch.CollectInputs
// from blocking on a TTY waiting for the user to paste PURLs.
func stdinIfPiped() io.Reader {
	if output.IsTTY(os.Stdin) {
		return nil
	}
	return os.Stdin
}
