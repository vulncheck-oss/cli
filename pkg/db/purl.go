package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/package-url/packageurl-go"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	_ "modernc.org/sqlite"
	dversion "pault.ag/go/debian/version"
)

type PurlEntry struct {
	Name            string                  `json:"name"`
	Version         string                  `json:"version"`
	Purl            []string                `json:"purl"`
	CVEs            []string                `json:"cves"`
	Vulnerabilities []sdk.PurlVulnerability `json:"vulnerabilities"`
}

func PURLSearch(indexName string, instance packageurl.PackageURL) ([]PurlEntry, *Stats, error) {
	// OS distribution indices (alpine-purls, ubuntu-purls, …) store one advisory
	// row per (package version, distro release). Matching them by installed
	// version is unreliable: the complete secdb for a package is spread across
	// rows keyed to *other* versions, and the rows keyed to the installed
	// version carry only a truncated/inconsistent subset. Instead we mirror the
	// API (see vulncheck/api pkg/purl): match every advisory row for the package
	// by name, then keep a CVE only when the installed version is below its
	// fixed version. Language ecosystem indices (npm, cargo, …) keep the simpler
	// per-version matching.

	if strings.HasSuffix(indexName, "-purls") {
		return purlSearchOS(indexName, instance)
	}
	return purlSearchExact(indexName, instance)
}

// purlSearchExact matches language ecosystem indices (npm, cargo, nuget, …)
// each row already describes the vulnerabilities affecting a specific package version.
func purlSearchExact(indexName string, instance packageurl.PackageURL) ([]PurlEntry, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	tableName := strings.ReplaceAll(indexName, "-", "_")
	query := fmt.Sprintf("SELECT name, version, purl, cves, vulnerabilities FROM `%s` WHERE purl LIKE ?", tableName)

	searchPURL := packageurl.PackageURL{
		Type:      instance.Type,
		Namespace: instance.Namespace,
		Name:      instance.Name,
		Version:   instance.Version,
		Subpath:   instance.Subpath,
	}.String()

	// Anchor on JSON string quotes so "1.0.1" doesn't bleed into rows storing
	// "1.0.10". When the caller supplies a versionless PURL, match any version
	// for that package by anchoring on "@" instead of a closing quote.
	pattern := `%"` + searchPURL + `"%`
	if instance.Version == "" {
		pattern = `%"` + searchPURL + `@%"%`
	}

	rows, err := db.Query(query, pattern)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			_ = err
		}
	}()

	var results []PurlEntry
	for rows.Next() {
		result, err := scanPurlRow(rows)
		if err != nil {
			return nil, nil, err
		}
		if len(result.Vulnerabilities) == 0 {
			continue
		}
		results = append(results, *result)
	}

	stats := &Stats{Duration: time.Since(startTime)}

	return results, stats, nil
}

// purlSearchOS matches OS distribution indices by package name (ignoring the
// installed version) and returns the CVEs whose fixed version is greater than
// the installed version - i.e. the package is still affected.
func purlSearchOS(indexName string, instance packageurl.PackageURL) ([]PurlEntry, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	tableName := strings.ReplaceAll(indexName, "-", "_")
	query := fmt.Sprintf("SELECT name, version, purl, cves, vulnerabilities FROM `%s` WHERE purl LIKE ?", tableName)

	// Match every advisory row for the package, regardless of version, e.g.
	// "pkg:apk/alpine/busybox@%".
	searchPrefix := packageurl.PackageURL{
		Type:      instance.Type,
		Namespace: instance.Namespace,
		Name:      instance.Name,
		Subpath:   instance.Subpath,
	}.String()

	rows, err := db.Query(query, "%"+searchPrefix+"@%")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			_ = err
		}
	}()

	// Deduplicate by CVE: the same advisory appears in many per-release rows.
	seen := make(map[string]struct{})
	var vulns []sdk.PurlVulnerability
	var cves []string

	for rows.Next() {
		result, err := scanPurlRow(rows)
		if err != nil {
			return nil, nil, err
		}

		for _, v := range result.Vulnerabilities {
			if v.FixedVersion == "" {
				continue
			}
			if _, exists := seen[v.Detection]; exists {
				continue
			}
			if !isAffectedVersion(instance.Version, v.FixedVersion) {
				continue
			}
			seen[v.Detection] = struct{}{}
			vulns = append(vulns, v)
			cves = append(cves, v.Detection)
		}
	}

	stats := &Stats{Duration: time.Since(startTime)}

	if len(vulns) == 0 {
		return nil, stats, nil
	}

	// Report against the installed package/version, matching the online API.
	return []PurlEntry{{
		Name:            instance.Name,
		Version:         instance.Version,
		CVEs:            cves,
		Vulnerabilities: vulns,
	}}, stats, nil
}

// scanPurlRow scans a single alpine/npm/… style row into a PurlEntry.
func scanPurlRow(rows *sql.Rows) (*PurlEntry, error) {
	var result PurlEntry
	var purlJSON, cvesJSON, vulnerabilitiesJSON sql.NullString

	if err := rows.Scan(
		&result.Name,
		&result.Version,
		&purlJSON,
		&cvesJSON,
		&vulnerabilitiesJSON,
	); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if purlJSON.Valid {
		if err := json.Unmarshal([]byte(purlJSON.String), &result.Purl); err != nil {
			return nil, fmt.Errorf("failed to unmarshal purl: %w", err)
		}
	}
	if cvesJSON.Valid {
		if err := json.Unmarshal([]byte(cvesJSON.String), &result.CVEs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cves: %w", err)
		}
	}
	if vulnerabilitiesJSON.Valid {
		if err := json.Unmarshal([]byte(vulnerabilitiesJSON.String), &result.Vulnerabilities); err != nil {
			return nil, fmt.Errorf("failed to unmarshal vulnerabilities: %w", err)
		}
	}

	return &result, nil
}

// versionReleaseSuffix strips distro release/build tags that the Debian version
// parser cannot handle (RPM-style "_4"/".elN" tags and Alpine "_gitNNNN" build
// suffixes), while preserving the upstream version and the "-rN" apk revision.
var versionReleaseSuffix = regexp.MustCompile(`(.el\d)*_\d*`)

// isAffectedVersion reports whether the installed version is below the fixed
// version Mirrors IsAffectedVersion in the vulncheck/api purl package
func isAffectedVersion(current, fixed string) bool {
	current = versionReleaseSuffix.ReplaceAllString(current, "$1")
	fixed = versionReleaseSuffix.ReplaceAllString(fixed, "$1")

	currentVersion, err1 := dversion.Parse(current)
	fixedVersion, err2 := dversion.Parse(fixed)
	if err1 != nil || err2 != nil {
		return false
	}
	return dversion.Compare(currentVersion, fixedVersion) < 0
}
