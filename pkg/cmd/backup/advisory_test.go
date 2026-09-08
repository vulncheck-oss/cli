package backup

import (
	"sort"
	"testing"

	"github.com/vulncheck-oss/cli/pkg/i18n"
)

func init() { i18n.Init() }

// v4 backups sit under `backup` so that everything bulk-download-shaped is
// discoverable in one place, and are named for the corpus rather than the API
// version: the v4 zip for a given name holds different data from the v3 one.
func TestBackupCommandStructure(t *testing.T) {
	var names []string
	for _, sub := range Command().Commands() {
		names = append(names, sub.Name())
	}
	sort.Strings(names)

	want := []string{"advisory", "download", "list", "url"}
	if len(names) != len(want) {
		t.Fatalf("subcommands = %v, want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Errorf("subcommands = %v, want %v", names, want)
			break
		}
	}
}

func TestBackupAdvisoryCommandStructure(t *testing.T) {
	cmd := advisoryCommand()

	if cmd.Use != "advisory <command>" {
		t.Errorf("Use = %q, want %q", cmd.Use, "advisory <command>")
	}

	var names []string
	for _, sub := range cmd.Commands() {
		names = append(names, sub.Name())
	}
	sort.Strings(names)

	want := []string{"download", "list", "url"}
	if len(names) != len(want) {
		t.Fatalf("subcommands = %v, want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Errorf("subcommands = %v, want %v", names, want)
			break
		}
	}
}

// The v3 subcommands take an <index>; the v4 ones take a <feed>. The wording is
// the only thing telling a user which corpus they are addressing, since most v4
// feed names are also v3 index names.
func TestBackupSubcommandsNameTheirArgument(t *testing.T) {
	uses := map[string]string{}
	for _, sub := range Command().Commands() {
		uses[sub.Name()] = sub.Use
	}
	for _, sub := range advisoryCommand().Commands() {
		uses["advisory "+sub.Name()] = sub.Use
	}

	cases := map[string]string{
		"url":               "url <index>",
		"download":          "download <index>",
		"advisory url":      "url <feed>",
		"advisory download": "download <feed>",
	}
	for name, want := range cases {
		if uses[name] != want {
			t.Errorf("%s: Use = %q, want %q", name, uses[name], want)
		}
	}
}

func TestBackupHelpDoesNotError(t *testing.T) {
	cmd := Command()
	if err := cmd.Help(); err != nil {
		t.Errorf("backup --help: %v", err)
	}
	for _, sub := range cmd.Commands() {
		if err := sub.Help(); err != nil {
			t.Errorf("backup %s --help: %v", sub.Name(), err)
		}
	}
}
