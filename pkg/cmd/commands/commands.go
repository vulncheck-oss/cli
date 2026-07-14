// Package commands implements `vulncheck commands`, a machine-readable
// dump of the command tree used by agents and skills to discover the
// CLI's surface without parsing --help.
//
// The output shape is versioned via SchemaVersion. Additions are not
// breaking; renames / removals require a bump.
package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/session"
)

// dump is the top-level JSON envelope.
type dump struct {
	SchemaVersion int         `json:"schema_version"`
	Root          commandInfo `json:"root"`
}

// commandInfo describes one command in the tree.
type commandInfo struct {
	// Name is the short, non-namespaced identifier (e.g. "list").
	Name string `json:"name"`
	// Path is the full invocation path from the root, excluding the
	// binary name itself (e.g. ["token", "list"]). Empty on the root.
	Path []string `json:"path"`
	// Short is the one-line description shown in --help.
	Short string `json:"short,omitempty"`
	// Hidden mirrors cobra.Command.Hidden — agents can filter these out
	// unless they specifically want to probe internal commands.
	Hidden bool `json:"hidden,omitempty"`
	// Deprecated carries the human message when cobra marks the command
	// deprecated; empty means "supported".
	Deprecated string `json:"deprecated,omitempty"`
	// Aliases lists alternative names (e.g. cobra sub-command aliases).
	Aliases []string `json:"aliases,omitempty"`
	// Args is the positional-argument summary string (`<label>`, `<path>`).
	// Not machine-parseable, but useful for agents rendering help text.
	Args string `json:"args,omitempty"`
	// SkipsAuth flags commands that don't require an API token — agents
	// can probe them without authenticating first (version, auth, etc.).
	SkipsAuth bool `json:"skips_auth,omitempty"`
	// Flags declared on this command directly. Inherited persistent
	// flags from parents are NOT repeated here — walk up Path to find them.
	Flags []flagInfo `json:"flags,omitempty"`
	// InheritedFlags are the persistent flags added by parent commands
	// (only populated on the ROOT to declare them once).
	InheritedFlags []flagInfo `json:"inherited_flags,omitempty"`
	// Subcommands recurses; empty for leaves.
	Subcommands []commandInfo `json:"subcommands,omitempty"`
}

type flagInfo struct {
	Name       string `json:"name"`
	Shorthand  string `json:"shorthand,omitempty"`
	Type       string `json:"type"`             // "bool", "string", "int", "stringSlice", etc.
	Default    string `json:"default,omitempty"`
	Usage      string `json:"usage,omitempty"`
	Deprecated string `json:"deprecated,omitempty"`
	Hidden     bool   `json:"hidden,omitempty"`
}

// Command returns the `vulncheck commands` cobra command. It emits the
// tree as JSON regardless of the --json flag, because the sole reason to
// invoke it is machine consumption.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commands",
		Short: "Print the command tree as JSON for agent / skill capability discovery",
		// Deliberately visible in the top-level help listing — this is a
		// first-class discoverability surface, and hiding it defeats the
		// point of shipping a capability probe.
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			root := cmd.Root()
			payload := dump{
				SchemaVersion: output.SchemaVersion,
				Root:          buildInfo(root, nil, true),
			}
			return r.JSON(payload)
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}

// buildInfo walks the cobra tree recursively. When isRoot is true the
// command's persistent flags are copied to InheritedFlags on the root
// entry so agents pick them up in one place; nested commands only list
// their own local flags.
func buildInfo(cmd *cobra.Command, path []string, isRoot bool) commandInfo {
	fullPath := append([]string{}, path...)
	if !isRoot {
		fullPath = append(fullPath, cmd.Name())
	}

	info := commandInfo{
		Name:       cmd.Name(),
		Path:       fullPath,
		Short:      cmd.Short,
		Hidden:     cmd.Hidden,
		Deprecated: cmd.Deprecated,
		Aliases:    cmd.Aliases,
		Args:       cmd.Use, // e.g. "list <search>"
		SkipsAuth:  cmd.Annotations != nil && cmd.Annotations["skipAuthCheck"] == "true",
	}

	// Local (non-persistent, non-inherited) flags for this command.
	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		// The root's PersistentFlags show up on the root's LocalFlags
		// AND on every child's InheritedFlags. We surface them once at
		// the top under InheritedFlags to keep the payload compact.
		if isRoot && isPersistentFlag(cmd, f.Name) {
			info.InheritedFlags = append(info.InheritedFlags, flagFromPflag(f))
			return
		}
		info.Flags = append(info.Flags, flagFromPflag(f))
	})

	for _, sub := range cmd.Commands() {
		// Skip cobra's built-in help / completion machinery — agents get
		// help via `commands --json` itself, and completions via
		// `vulncheck completion <shell>` which is documented separately.
		if sub.Name() == "help" || sub.Name() == "completion" {
			continue
		}
		info.Subcommands = append(info.Subcommands, buildInfo(sub, fullPath, false))
	}
	return info
}

// isPersistentFlag reports whether the named flag is declared on the
// command's PersistentFlags set (as opposed to Flags). Cobra merges
// PersistentFlags into LocalFlags at runtime so we need this to
// distinguish them.
func isPersistentFlag(cmd *cobra.Command, name string) bool {
	return cmd.PersistentFlags().Lookup(name) != nil
}

func flagFromPflag(f *pflag.Flag) flagInfo {
	return flagInfo{
		Name:       f.Name,
		Shorthand:  f.Shorthand,
		Type:       f.Value.Type(),
		Default:    f.DefValue,
		Usage:      f.Usage,
		Deprecated: f.Deprecated,
		Hidden:     f.Hidden,
	}
}
