package commands

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// synthetic builds a minimal cobra tree so we can exercise buildInfo
// without depending on the whole vulncheck command surface.
func synthetic() *cobra.Command {
	root := &cobra.Command{
		Use:   "vulncheck",
		Short: "root short",
	}
	root.PersistentFlags().Bool("json", false, "emit json")
	root.PersistentFlags().Bool("quiet", false, "suppress info")

	widget := &cobra.Command{
		Use:     "widget",
		Short:   "manage widgets",
		Aliases: []string{"w"},
	}
	widget.Annotations = map[string]string{"skipAuthCheck": "true"}

	widgetList := &cobra.Command{
		Use:   "list",
		Short: "list widgets",
	}
	widgetList.Flags().Int("limit", 0, "page size")
	widgetList.Flags().Bool("all", false, "paginate every page")

	widgetDelete := &cobra.Command{
		Use:        "delete <id>",
		Short:      "delete a widget",
		Hidden:     true,
		Deprecated: "use `widget remove` instead",
	}

	widget.AddCommand(widgetList)
	widget.AddCommand(widgetDelete)

	root.AddCommand(widget)
	// cobra adds these automatically at runtime — synthesise them so the
	// filter in buildInfo has something to skip.
	root.AddCommand(&cobra.Command{Use: "help"})
	root.AddCommand(&cobra.Command{Use: "completion"})

	return root
}

func TestBuildInfoRootExposesPersistentFlags(t *testing.T) {
	info := buildInfo(synthetic(), nil, true)

	if info.Name != "vulncheck" {
		t.Errorf("root name = %q, want vulncheck", info.Name)
	}
	if len(info.Path) != 0 {
		t.Errorf("root path should be empty, got %v", info.Path)
	}
	// Root's persistent flags are surfaced under InheritedFlags exactly
	// once so agents don't see them duplicated on every child.
	seen := map[string]bool{}
	for _, f := range info.InheritedFlags {
		seen[f.Name] = true
	}
	for _, want := range []string{"json", "quiet"} {
		if !seen[want] {
			t.Errorf("root.inherited_flags missing %q; got %v", want, seen)
		}
	}
}

func TestBuildInfoSkipsHelpAndCompletion(t *testing.T) {
	info := buildInfo(synthetic(), nil, true)
	for _, sub := range info.Subcommands {
		if sub.Name == "help" || sub.Name == "completion" {
			t.Fatalf("subcommand %q should have been filtered out", sub.Name)
		}
	}
}

func TestBuildInfoRecordsSubcommandPath(t *testing.T) {
	info := buildInfo(synthetic(), nil, true)
	var widget *commandInfo
	for i := range info.Subcommands {
		if info.Subcommands[i].Name == "widget" {
			widget = &info.Subcommands[i]
			break
		}
	}
	if widget == nil {
		t.Fatal("widget subcommand missing")
	}
	if strings.Join(widget.Path, ",") != "widget" {
		t.Errorf("widget path = %v, want [widget]", widget.Path)
	}
	var list *commandInfo
	for i := range widget.Subcommands {
		if widget.Subcommands[i].Name == "list" {
			list = &widget.Subcommands[i]
			break
		}
	}
	if list == nil {
		t.Fatal("widget list subcommand missing")
	}
	if strings.Join(list.Path, ",") != "widget,list" {
		t.Errorf("widget list path = %v, want [widget list]", list.Path)
	}
}

func TestBuildInfoPropagatesFlagMetadata(t *testing.T) {
	info := buildInfo(synthetic(), nil, true)
	list := findSub(t, info, "widget", "list")
	byName := map[string]flagInfo{}
	for _, f := range list.Flags {
		byName[f.Name] = f
	}
	limit, ok := byName["limit"]
	if !ok {
		t.Fatal("limit flag missing")
	}
	if limit.Type != "int" {
		t.Errorf("limit type = %q, want int", limit.Type)
	}
	if limit.Default != "0" {
		t.Errorf("limit default = %q, want 0", limit.Default)
	}
	if _, ok := byName["all"]; !ok {
		t.Fatal("all flag missing")
	}
}

func TestBuildInfoCarriesDeprecationHiddenAndAliases(t *testing.T) {
	info := buildInfo(synthetic(), nil, true)
	widget := findSub(t, info, "widget")
	if len(widget.Aliases) == 0 || widget.Aliases[0] != "w" {
		t.Errorf("widget aliases = %v, want [w]", widget.Aliases)
	}
	if !widget.SkipsAuth {
		t.Error("widget should be marked skips_auth (Annotations)")
	}

	del := findSub(t, info, "widget", "delete")
	if !del.Hidden {
		t.Error("delete should be marked hidden")
	}
	if del.Deprecated == "" {
		t.Error("delete should carry a deprecation message")
	}
}

// TestBuildInfoNoDuplicatedPersistentFlagsOnChildren locks in the
// design: root's persistent flags appear ONCE at the top, not repeated
// on every subcommand's Flags array. Agents walk up to the root to find
// them.
func TestBuildInfoNoDuplicatedPersistentFlagsOnChildren(t *testing.T) {
	info := buildInfo(synthetic(), nil, true)
	widget := findSub(t, info, "widget")
	for _, f := range widget.Flags {
		if f.Name == "json" || f.Name == "quiet" {
			t.Errorf("widget.flags leaked root persistent flag %q", f.Name)
		}
	}
}

func TestCommandCanBeConstructed(t *testing.T) {
	c := Command()
	if c.Name() != "commands" {
		t.Errorf("Name() = %q, want commands", c.Name())
	}
	if !c.Hidden {
		t.Error("commands should be hidden from top-level help listings")
	}
	if c.Annotations["skipAuthCheck"] != "true" {
		t.Error("commands should skip auth (it's a probe)")
	}
}

func findSub(t *testing.T, info commandInfo, names ...string) commandInfo {
	t.Helper()
	current := info
	for _, name := range names {
		var next *commandInfo
		for i := range current.Subcommands {
			if current.Subcommands[i].Name == name {
				next = &current.Subcommands[i]
				break
			}
		}
		if next == nil {
			t.Fatalf("subcommand %q not found under %q; siblings: %v",
				name, current.Name, subNames(current))
		}
		current = *next
	}
	return current
}

func subNames(info commandInfo) []string {
	out := make([]string, 0, len(info.Subcommands))
	for _, s := range info.Subcommands {
		out = append(out, s.Name)
	}
	return out
}
