package root

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/cmd/upgrade"

	"github.com/vulncheck-oss/cli/pkg/cmd/offline"

	"github.com/vulncheck-oss/cli/pkg/cmd/token"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/about"
	"github.com/vulncheck-oss/cli/pkg/cmd/auth"
	"github.com/vulncheck-oss/cli/pkg/cmd/backup"
	"github.com/vulncheck-oss/cli/pkg/cmd/commands"
	"github.com/vulncheck-oss/cli/pkg/cmd/cpe"
	"github.com/vulncheck-oss/cli/pkg/cmd/index"
	"github.com/vulncheck-oss/cli/pkg/cmd/indices"
	"github.com/vulncheck-oss/cli/pkg/cmd/pdns"
	"github.com/vulncheck-oss/cli/pkg/cmd/purl"
	"github.com/vulncheck-oss/cli/pkg/cmd/rule"
	"github.com/vulncheck-oss/cli/pkg/cmd/scan"
	"github.com/vulncheck-oss/cli/pkg/cmd/tag"
	"github.com/vulncheck-oss/cli/pkg/cmd/version"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type AuthError struct {
	err error
}

func (ae *AuthError) Error() string {
	return ae.err.Error()
}

func NewCmdRoot() *cobra.Command {
	i18n.Init()
	cmd := &cobra.Command{
		Use:   "vulncheck <command> <subcommand> [flags]",
		Short: "VulnCheck CLI.",
		Long:  i18n.C.RootLong,
		Example: heredoc.Doc(`
			$ vulncheck indices list
			$ vulncheck index abb
			$ vulncheck backup abb
		`),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			environment.Init()
			config.Init()

			// Build the Renderer from global flags and stash it on the
			// command's context so every subcommand can pick it up via
			// output.FromCmd(cmd).
			r := rendererFromCmd(cmd)
			cmd.SetContext(output.WithContext(cmd.Context(), r))

			if session.IsAuthCheckEnabled(cmd) && cmd.Parent().Name() != "completion" && !session.CheckAuth() {
				// Auth help is human guidance, not payload — route to stderr
				// (or suppress when --json so callers see only the JSON error).
				if !r.IsJSON() {
					_, _ = fmt.Fprintln(r.Stderr(), authHelp())
				}
				return errs.AuthRequired(i18n.C.ErrorNoToken)
			}

			return nil
		},
	}

	cmd.SilenceUsage = true
	cmd.SilenceErrors = true

	cmd.PersistentFlags().BoolP("help", "h", false, "Show help for command")
	cmd.PersistentFlags().Bool("json", false, "Emit output as JSON on stdout; info/progress are routed to stderr")
	cmd.PersistentFlags().Bool("quiet", false, "Suppress informational output (errors and payloads still render)")
	cmd.PersistentFlags().Bool("no-color", false, "Disable ANSI colour output (also honours NO_COLOR env)")
	cmd.PersistentFlags().Bool("no-interactive", false, "Disable TUI prompts and confirmations; required for headless / CI use")

	cmd.AddGroup(&cobra.Group{
		ID:    "core",
		Title: "Core Commands",
	})

	cmd.AddCommand(version.Command())
	cmd.AddCommand(about.Command())
	cmd.AddCommand(commands.Command())
	cmd.AddCommand(auth.Command())
	cmd.AddCommand(token.Command())
	cmd.AddCommand(upgrade.Command())
	cmd.AddCommand(indices.Command())
	cmd.AddCommand(index.Command())
	cmd.AddCommand(backup.Command())
	cmd.AddCommand(cpe.Command())
	cmd.AddCommand(purl.Command())
	cmd.AddCommand(scan.Command())
	cmd.AddCommand(rule.Command())
	cmd.AddCommand(tag.Command())
	cmd.AddCommand(pdns.Command())
	cmd.AddCommand(offline.Command())

	return cmd
}

// rendererFromCmd resolves the Renderer to use for a given command
// invocation. It consults the persistent flags (--json / --quiet /
// --no-color / --no-interactive) and layers them on top of environment-
// derived defaults (NO_COLOR, TERM=dumb, TTY detection, CI=1).
func rendererFromCmd(cmd *cobra.Command) *output.Renderer {
	jsonOut, _ := cmd.Flags().GetBool("json")
	quiet, _ := cmd.Flags().GetBool("quiet")
	noColor, _ := cmd.Flags().GetBool("no-color")
	noInteractive, _ := cmd.Flags().GetBool("no-interactive")

	mode := output.ModeText
	if jsonOut {
		mode = output.ModeJSON
	}

	color := output.ColorEnabled()
	if noColor {
		color = false
	}

	// Interactivity baseline: TTY on stdin + stdout, not CI. Either an
	// explicit --no-interactive or --json flips it off.
	interactive := output.Interactive() && !noInteractive

	return output.New(output.Options{
		Mode:        mode,
		Color:       color,
		Quiet:       quiet,
		Interactive: interactive,
	})
}

// errorEnvelope is the structured error shape emitted in JSON mode.
// `code` and `http_status` are stable across releases — agents can key off
// `.error.code` (one of the errs.Kind values) for programmatic dispatch.
// `schema_version` lets agents detect a CLI whose envelope they don't
// speak — see internal/output.SchemaVersion.
type errorEnvelope struct {
	SchemaVersion int       `json:"schema_version"`
	Error         errorBody `json:"error"`
}

type errorBody struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"http_status,omitempty"`
}

func Execute() {
	// Install a single cancellable context that listens for SIGINT/SIGTERM
	// and propagates cancellation through every cmd.Context() consumer
	// (SDK HTTP requests, future scan/sync long-runners). Ctrl-C now cleanly
	// cancels in-flight work instead of leaving partial state.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	root := NewCmdRoot()
	root.SetContext(ctx)

	err := root.ExecuteContext(ctx)
	if err == nil {
		return
	}

	// Flags have been parsed by now even on error paths — rebuild a Renderer
	// from the root's persistent flags so we render the error in the right
	// shape. (PersistentPreRunE may not have run for every error path, e.g.
	// flag-parsing failures, so we don't rely on context.)
	r := rendererFromCmd(root)

	// If the context was cancelled by a signal, the underlying error may not
	// itself wrap context.Canceled — promote it explicitly so the user sees
	// "cancelled" and exit 130 rather than a network-shaped error.
	classified := errs.Classify(err)
	if ctx.Err() != nil && classified.Kind != errs.KindCancelled {
		classified = errs.Wrap(errs.KindCancelled, err, "cancelled")
	}

	if r.IsJSON() {
		_ = r.JSON(errorEnvelope{
			SchemaVersion: output.SchemaVersion,
			Error: errorBody{
				Code:       string(classified.Kind),
				Message:    classified.Message,
				HTTPStatus: classified.HTTPStatus,
			},
		})
	} else {
		_, _ = fmt.Fprintln(r.Stderr(), ui.Danger(classified.Message).Error())
	}

	os.Exit(classified.ExitCode())
}
