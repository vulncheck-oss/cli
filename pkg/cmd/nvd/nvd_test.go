package nvd

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/i18n"
)

func init() { i18n.Init() }

func TestCommandStructure(t *testing.T) {
	cmd := Command()
	assert.Equal(t, "nvd <command>", cmd.Use)
	assert.NotEmpty(t, cmd.Short)

	var names []string
	for _, sub := range cmd.Commands() {
		names = append(names, sub.Name())
		assert.NotEmpty(t, sub.Short, sub.Name())
	}
	sort.Strings(names)
	assert.Equal(t, []string{"cpematch", "cpes", "cvehistory", "cves", "source"}, names)
}

// TestParametersAreTheAPIs pins each command to its endpoint's query parameters, so the CLI
// neither drops one the API takes nor offers one it does not.
func TestParametersAreTheAPIs(t *testing.T) {
	paging := []string{"resultsPerPage", "startIndex", "cursor"}
	lastMod := []string{"lastModStartDate", "lastModEndDate"}

	want := map[string][]string{
		"cves": append(append([]string{
			"cveId", "cveIds", "cpeName", "isVulnerable", "virtualMatchString",
			"versionStart", "versionStartType", "versionEnd", "versionEndType",
			"vulnStatuses", "cveTag", "noRejected", "hasKev", "vcVulnerableCPEs",
			"pubStartDate", "pubEndDate", "sourceIdentifier",
		}, lastMod...), paging...),
		"cvehistory": append([]string{"cveId", "eventName", "changeStartDate", "changeEndDate"}, paging...),
		"cpes":       append(append([]string{"cpeNameId", "cpeMatchString", "matchCriteriaId"}, lastMod...), paging...),
		"cpematch":   append(append([]string{"cveId", "matchCriteriaId"}, lastMod...), paging...),
		// /source/2.0 has no cursor.
		"source": append([]string{"sourceIdentifier", "resultsPerPage", "startIndex"}, lastMod...),
	}

	for _, ep := range endpoints() {
		var got []string
		for _, p := range ep.params {
			got = append(got, p.name)
		}
		assert.ElementsMatch(t, want[ep.use], got, ep.use)
	}
}

// TestFlagsAreTheParametersInKebabCase pins the one naming rule: a flag is its API
// parameter in kebab-case, so --cve-id is cveId and nothing is renamed.
func TestFlagsAreTheParametersInKebabCase(t *testing.T) {
	kebab := func(name string) string {
		var b strings.Builder
		for i, r := range name {
			if unicode.IsUpper(r) && i > 0 && unicode.IsLower(rune(name[i-1])) {
				b.WriteByte('-')
			}
			b.WriteRune(unicode.ToLower(r))
		}
		return b.String()
	}

	for _, ep := range endpoints() {
		cmd := subcommand(ep)
		for _, p := range ep.params {
			assert.Equal(t, kebab(p.name), p.flag, "%s %s", ep.use, p.name)
			assert.NotNil(t, cmd.Flags().Lookup(p.flag), "%s --%s", ep.use, p.flag)
		}
		assert.NotNil(t, cmd.Flags().Lookup("full"), ep.use)
	}
}

// Every filter is a flag. cobra's default accepts a positional and ignores it, which would
// answer a broader question than the one asked, at exit 0.
func TestPositionalArgumentsAreRejected(t *testing.T) {
	for _, ep := range endpoints() {
		cmd := subcommand(ep)
		assert.Error(t, cmd.Args(cmd, []string{"CVE-2021-44228"}), ep.use)
		assert.NoError(t, cmd.Args(cmd, nil), ep.use)
	}
}

const cvesBody = `{"resultsPerPage":1,"startIndex":0,"totalResults":837,"format":"NVD_CVE","version":"2.0",` +
	`"timestamp":"2026-10-13T00:00:00.000","vulnerabilities":[{"cve":{"id":"CVE-2021-44228"}}],"nextCursor":"tok123"}`

// run executes `nvd <args>` against a fake API and returns what reached stdout and stderr.
func run(t *testing.T, mode output.Mode, handler http.HandlerFunc, args ...string) (string, string, error) {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	previous := environment.Env.API
	environment.Env.API = srv.URL
	t.Cleanup(func() { environment.Env.API = previous })
	t.Setenv(config.EnvToken, "vulncheck_test_token")

	var stdout, stderr bytes.Buffer
	renderer := output.New(output.Options{Mode: mode, Stdout: &stdout, Stderr: &stderr})

	cmd := Command()
	cmd.SetArgs(args)
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	err := cmd.ExecuteContext(output.WithContext(context.Background(), renderer))
	return stdout.String(), stderr.String(), err
}

func TestRun(t *testing.T) {
	t.Run("sends only the flags given, values unchanged", func(t *testing.T) {
		var gotPath, gotQuery string
		_, _, err := run(t, output.ModeJSON, func(w http.ResponseWriter, r *http.Request) {
			gotPath, gotQuery = r.URL.Path, r.URL.RawQuery
			_, _ = fmt.Fprint(w, cvesBody)
		}, "cves", "--cve-id", "CVE-2021-44228", "--has-kev", "--results-per-page", "1", "--cursor", "")
		require.NoError(t, err)

		assert.Equal(t, "/rest/json/cves/2.0", gotPath)
		assert.Equal(t, "cursor=&cveId=CVE-2021-44228&hasKev=true&resultsPerPage=1", gotQuery,
			"an unset flag is not sent, and an empty --cursor is")
	})

	t.Run("--json puts only the records on stdout", func(t *testing.T) {
		stdout, stderr, err := run(t, output.ModeJSON, func(w http.ResponseWriter, _ *http.Request) {
			_, _ = fmt.Fprint(w, cvesBody)
		}, "cves", "--cve-id", "CVE-2021-44228")
		require.NoError(t, err)

		assert.JSONEq(t, `[{"cve":{"id":"CVE-2021-44228"}}]`, stdout)
		assert.Contains(t, stderr, "1 of 837 results")
		assert.Contains(t, stderr, "Next page: --cursor tok123")
	})

	t.Run("--full puts the whole NVD response on stdout", func(t *testing.T) {
		stdout, _, err := run(t, output.ModeJSON, func(w http.ResponseWriter, _ *http.Request) {
			_, _ = fmt.Fprint(w, cvesBody)
		}, "cves", "--cve-id", "CVE-2021-44228", "--full")
		require.NoError(t, err)
		assert.JSONEq(t, cvesBody, stdout)
	})

	t.Run("--json output is the same on every run", func(t *testing.T) {
		handler := func(w http.ResponseWriter, _ *http.Request) { _, _ = fmt.Fprint(w, cvesBody) }
		first, _, err := run(t, output.ModeJSON, handler, "cves", "--full")
		require.NoError(t, err)
		second, _, err := run(t, output.ModeJSON, handler, "cves", "--full")
		require.NoError(t, err)
		assert.Equal(t, first, second)
	})

	t.Run("text mode keeps the status lines on stdout", func(t *testing.T) {
		stdout, stderr, err := run(t, output.ModeText, func(w http.ResponseWriter, _ *http.Request) {
			_, _ = fmt.Fprint(w, cvesBody)
		}, "cves", "--cve-id", "CVE-2021-44228")
		require.NoError(t, err)
		assert.Contains(t, stdout, "1 of 837 results")
		assert.Contains(t, stdout, `"id": "CVE-2021-44228"`)
		assert.Empty(t, stderr)
	})

	t.Run("source never sends a cursor", func(t *testing.T) {
		var gotQuery string
		_, _, err := run(t, output.ModeJSON, func(w http.ResponseWriter, r *http.Request) {
			gotQuery = r.URL.RawQuery
			_, _ = fmt.Fprint(w, `{"resultsPerPage":0,"startIndex":20,"totalResults":514,"sources":[]}`)
		}, "source", "--start-index", "20")
		require.NoError(t, err)
		assert.Equal(t, "startIndex=20", gotQuery)

		_, _, err = run(t, output.ModeJSON, nil, "source", "--cursor", "")
		assert.ErrorContains(t, err, "unknown flag: --cursor")
	})

	t.Run("a refused request becomes a bad_request error with the API's reason", func(t *testing.T) {
		_, _, err := run(t, output.ModeJSON, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, `{"message":"parameter \"keywordSearch\" is not supported in v1; deferred to v2","status":400}`)
		}, "cves")
		require.Error(t, err)

		classified := errs.Classify(err)
		assert.Equal(t, errs.KindBadRequest, classified.Kind)
		assert.Equal(t, `parameter "keywordSearch" is not supported in v1; deferred to v2`, classified.Message)
	})
}
