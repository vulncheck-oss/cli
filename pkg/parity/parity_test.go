//go:build parity

// Package parity contains an opt-in integration test suite that compares
// the CLI's offline emulation of server-side matching against the live
// VulnCheck API. Offline mode is meant to emulate the API for air-gapped
// users, but nothing in the tree otherwise asserts that the two paths
// actually agree on the same input — past drift has only surfaced via
// customer reports (alpine PURL SBOMs returning zero vulns offline while
// the same input returned many online; a linux_kernel CPE returning zero
// in an offline SBOM scan while returning thousands via the online CPE
// endpoint). Seeding those exact inputs here keeps them permanent
// regression cases.
//
// Run with:
//
//	VULNCHECK_API_TOKEN=... go test -tags=parity ./pkg/parity/...
//
// The build tag keeps this out of the default `go test ./...` run because it
// needs a live token and locally synced indices; the test skips gracefully if
// either precondition is missing.
package parity

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/package-url/packageurl-go"
	"github.com/vulncheck-oss/cli/pkg/bill"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/packages"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/models"
	"github.com/vulncheck-oss/cli/pkg/session"
)

// noopProgress is the iterator callback signature expected by the bill package.
// The parity test doesn't render progress; a real scan does.
func noopProgress(int, int) {}

// preflight loads the shared preconditions (auth + synced indices) or skips
// the test. Kept as a helper so both parity tests fail the same way.
func preflight(t *testing.T) cache.InfoFile {
	t.Helper()
	if config.Token() == "" {
		t.Skip("parity test requires VULNCHECK_API_TOKEN or VC_TOKEN in env (or a logged-in CLI config)")
	}
	indices, err := cache.Indices()
	if err != nil {
		t.Skipf("parity test requires synced offline indices: %v", err)
	}
	return indices
}

// loadFixture reads a fixture file, stripping blank lines and `#` comments so
// PURL / CPE lists can be annotated inline.
func loadFixture(t *testing.T, name string) []string {
	t.Helper()
	f, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("open fixture %s: %v", name, err)
	}
	defer func() { _ = f.Close() }()

	var out []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		out = append(out, line)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("read fixture %s: %v", name, err)
	}
	return out
}

// cveSet flattens a slice of vuln records into a sorted, deduplicated slice of
// CVE ids. Parity is about which vulnerabilities each path finds, not how they
// annotate them — CVSS / KEV / description differences are orthogonal.
func cveSet(vulns []models.ScanResultVulnerabilities) []string {
	seen := make(map[string]struct{}, len(vulns))
	for _, v := range vulns {
		if v.CVE == "" {
			continue
		}
		seen[v.CVE] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for cve := range seen {
		out = append(out, cve)
	}
	sort.Strings(out)
	return out
}

// diff returns the elements in a that are absent from b.
func diff(a, b []string) []string {
	bset := make(map[string]struct{}, len(b))
	for _, x := range b {
		bset[x] = struct{}{}
	}
	var out []string
	for _, x := range a {
		if _, ok := bset[x]; !ok {
			out = append(out, x)
		}
	}
	return out
}

// TestPURLParity walks every PURL in testdata/purls.txt through the online
// batch endpoint and the offline sqlite lookup, and fails if the CVE sets
// disagree. Regression seeds: alpine musl/busybox rows from the previously
// reported "alpine offline returns zero vulns" drift.
func TestPURLParity(t *testing.T) {
	indices := preflight(t)
	ctx := context.Background()
	purls := loadFixture(t, "purls.txt")

	for _, p := range purls {
		t.Run(p, func(t *testing.T) {
			instance, err := packageurl.FromString(p)
			if err != nil {
				t.Fatalf("parse purl: %v", err)
			}
			indexName := packages.IndexFromInstance(instance)
			if indices.GetIndex(indexName) == nil {
				t.Skipf("index %s not synced locally; run `vulncheck offline sync --add %s`", indexName, indexName)
			}

			detail := []models.PurlDetail{{Purl: p}}

			onlineVulns, unprocessed, err := bill.GetBatchVulns(ctx, detail, noopProgress)
			if err != nil {
				t.Fatalf("online lookup: %v", err)
			}
			// A purl the API cannot assess returns no findings, and offline
			// returns none either, so the CVE sets would agree trivially and the
			// row would pass for the wrong reason. Skip rather than bank a false
			// match.
			if len(unprocessed) > 0 {
				t.Skipf("API could not assess this purl: %s (%s)",
					unprocessed[0].Purl, unprocessed[0].Reason)
			}
			offlineVulns, err := bill.GetOfflineVulns(indices, detail, noopProgress, false)
			if err != nil {
				t.Fatalf("offline lookup: %v", err)
			}

			online, offline := cveSet(onlineVulns), cveSet(offlineVulns)
			missingOffline := diff(online, offline)
			extraOffline := diff(offline, online)
			if len(missingOffline) > 0 || len(extraOffline) > 0 {
				t.Errorf("PURL %s: online=%d offline=%d\n  only in online:  %v\n  only in offline: %v",
					p, len(online), len(offline), missingOffline, extraOffline)
			}
		})
	}
}

// TestCPEParity walks every CPE in testdata/cpes.txt through the online /v3/cpe
// endpoint and the offline cpecve lookup, and fails if the CVE sets disagree.
// Regression seed: linux_kernel CPE from the previously reported "SBOM CPE
// scan returns zero while the online CPE endpoint returns thousands" drift.
//
// Note: `bill.GetOfflineCpeVulns` returns entries keyed by (cpe, cve) with
// version/product metadata; we collapse to a CVE set to compare against the
// online endpoint which returns []string of CVE ids directly.
func TestCPEParity(t *testing.T) {
	indices := preflight(t)
	if indices.GetIndex("cpecve") == nil {
		t.Skip("index cpecve not synced locally; run `vulncheck offline sync --add cpecve`")
	}
	ctx := context.Background()
	client := session.ConnectWithContext(ctx, config.Token())
	cpes := loadFixture(t, "cpes.txt")

	for _, c := range cpes {
		t.Run(c, func(t *testing.T) {
			resp, err := client.GetCpe(c)
			if err != nil {
				t.Fatalf("online lookup: %v", err)
			}
			offlineVulns, err := bill.GetOfflineCpeVulns(indices, []string{c}, noopProgress, false)
			if err != nil {
				t.Fatalf("offline lookup: %v", err)
			}

			online := sortedUnique(resp.Data)
			offline := cveSet(offlineVulns)
			missingOffline := diff(online, offline)
			extraOffline := diff(offline, online)
			if len(missingOffline) > 0 || len(extraOffline) > 0 {
				t.Errorf("CPE %s: online=%d offline=%d\n  only in online:  %v\n  only in offline: %v",
					c, len(online), len(offline), missingOffline, extraOffline)
			}
		})
	}
}

func sortedUnique(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	for _, x := range in {
		if x == "" {
			continue
		}
		seen[x] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for x := range seen {
		out = append(out, x)
	}
	sort.Strings(out)
	return out
}
