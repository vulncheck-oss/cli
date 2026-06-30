package output

import (
	"context"

	"github.com/spf13/cobra"
)

type ctxKey struct{}

// WithContext attaches a Renderer to ctx. The root command's
// PersistentPreRunE installs one; subcommands retrieve it via FromCmd.
func WithContext(ctx context.Context, r *Renderer) context.Context {
	return context.WithValue(ctx, ctxKey{}, r)
}

// FromContext returns the Renderer attached to ctx, or a sensible default
// (text mode, env-derived) when none is present. Always returns non-nil.
func FromContext(ctx context.Context) *Renderer {
	if ctx == nil {
		return NewFromEnv(ModeText)
	}
	if r, ok := ctx.Value(ctxKey{}).(*Renderer); ok && r != nil {
		return r
	}
	return NewFromEnv(ModeText)
}

// FromCmd is shorthand for FromContext(cmd.Context()). Use in command
// RunE bodies: `r := output.FromCmd(cmd)`.
func FromCmd(cmd *cobra.Command) *Renderer {
	if cmd == nil {
		return NewFromEnv(ModeText)
	}
	return FromContext(cmd.Context())
}
