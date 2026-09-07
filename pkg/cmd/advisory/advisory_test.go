package advisory

import (
	"sort"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
)

func init() { i18n.Init() }

func TestCommandStructure(t *testing.T) {
	cmd := Command()

	if cmd.Use != "advisory <command>" {
		t.Errorf("Use = %q, want %q", cmd.Use, "advisory <command>")
	}
	if cmd.Short == "" {
		t.Error("Short must be set from i18n")
	}

	var names []string
	for _, sub := range cmd.Commands() {
		names = append(names, sub.Name())
	}
	sort.Strings(names)

	want := []string{"browse", "feeds", "list"}
	if len(names) != len(want) {
		t.Fatalf("subcommands = %v, want %v", names, want)
	}
	for i, n := range want {
		if names[i] != n {
			t.Errorf("subcommands = %v, want %v", names, want)
			break
		}
	}
}

// The flag set is the CLI half of a closed contract: /v4/advisory answers an
// unrecognised parameter with an empty result set rather than an error, so a
// flag that no longer maps to a parameter is silent data loss. Pin it.
func TestQueryFlagsAreTheDocumentedSet(t *testing.T) {
	want := []string{
		"all", "cpe", "cursor", "cve", "description-lang", "feed", "full",
		"limit", "package-name", "page", "platform", "product", "purl",
		"reference-tag", "reference-url", "start-cursor", "updated-after",
		"updated-before", "vendor", "version",
	}

	for _, cmd := range []*cobra.Command{List(), Browse()} {
		var got []string
		cmd.Flags().VisitAll(func(f *pflag.Flag) { got = append(got, f.Name) })
		sort.Strings(got)

		if len(got) != len(want) {
			t.Fatalf("%s flags = %v, want %v", cmd.Name(), got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("%s flags = %v, want %v", cmd.Name(), got, want)
				break
			}
		}
	}
}

// --feed is deliberately renamed off its "name" wire parameter, because "name"
// collides with too much else in a CVE record to read well as a flag.
// The cursor flags exist so that --full's next_cursor can be handed back; the
// SDK plumbing they drive already runs on every --all.
func TestCursorFlagsReachTheQueryParameters(t *testing.T) {
	params := (&queryOptions{feed: "ghsa", cursor: "abc", startCursor: true}).params()
	if params.Cursor != "abc" {
		t.Errorf("Cursor = %q, want abc", params.Cursor)
	}
	if !params.StartCursor {
		t.Error("StartCursor should be set")
	}
}

func TestFeedFlagMapsToTheNameParameter(t *testing.T) {
	opts := &queryOptions{feed: "ghsa", cve: "CVE-2019-12255", packageName: "pkg"}
	params := opts.params()

	if params.Name != "ghsa" {
		t.Errorf("Name = %q, want %q", params.Name, "ghsa")
	}
	if params.CveID != "CVE-2019-12255" {
		t.Errorf("CveID = %q, want %q", params.CveID, "CVE-2019-12255")
	}
	if params.PackageName != "pkg" {
		t.Errorf("PackageName = %q, want %q", params.PackageName, "pkg")
	}
}

// --all overrides both page and limit, so accepting either silently would
// discard something the user asked for.
func TestValidateRejectsAllWithPagingFlags(t *testing.T) {
	for _, tt := range []struct {
		name string
		opts *queryOptions
		want string
	}{
		{"page", &queryOptions{feed: "ghsa", all: true, page: 3}, "--page"},
		{"limit below the maximum", &queryOptions{feed: "ghsa", all: true, limit: 5}, "--limit"},
		{"cursor", &queryOptions{feed: "ghsa", all: true, cursor: "abc"}, "--cursor"},
		{"start-cursor", &queryOptions{feed: "ghsa", all: true, startCursor: true}, "--cursor"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.opts.validate()
			if err == nil {
				t.Fatalf("expected --all with %s to be rejected", tt.want)
			}
			if !strings.Contains(err.Error(), tt.want) {
				t.Errorf("message %q should name %s", err.Error(), tt.want)
			}
		})
	}
}

// --limit 100 asks for exactly what --all already does, so rejecting it would
// be refusing a request that is being honoured.
func TestValidateAcceptsAllWithTheMaximumLimit(t *testing.T) {
	if err := (&queryOptions{feed: "ghsa", all: true, limit: sdk.MaxAdvisoryLimit}).validate(); err != nil {
		t.Errorf("--all --limit %d should be accepted: %v", sdk.MaxAdvisoryLimit, err)
	}
}

// --all pages by cursor, so the offset ceiling that would reject this query
// without --all must not apply to it.
func TestValidateIgnoresTheOffsetCeilingUnderAll(t *testing.T) {
	if err := (&queryOptions{feed: "ghsa", all: true}).validate(); err != nil {
		t.Errorf("--all should be accepted: %v", err)
	}
	if err := (&queryOptions{feed: "ghsa", page: 1001}).validate(); err == nil {
		t.Error("page 1001 at the default limit exceeds the result window and should be rejected")
	}
}

func TestValidateRequiresAFilter(t *testing.T) {
	if err := (&queryOptions{limit: 10}).validate(); err == nil {
		t.Error("an unfiltered query would walk the whole corpus and must be rejected")
	}
	if err := (&queryOptions{cve: "CVE-2019-12255"}).validate(); err != nil {
		t.Errorf("a filtered query should be accepted: %v", err)
	}
}

func TestHelpDoesNotError(t *testing.T) {
	for _, cmd := range []*cobra.Command{Command(), Feeds(), List(), Browse()} {
		cmd.SetArgs([]string{"--help"})
		if err := cmd.Help(); err != nil {
			t.Errorf("%s --help: %v", cmd.Name(), err)
		}
	}
}
