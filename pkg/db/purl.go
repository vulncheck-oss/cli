package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/package-url/packageurl-go"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	_ "modernc.org/sqlite"
)

type PurlEntry struct {
	Name            string                  `json:"name"`
	Version         string                  `json:"version"`
	Purl            []string                `json:"purl"`
	CVEs            []string                `json:"cves"`
	Vulnerabilities []sdk.PurlVulnerability `json:"vulnerabilities"`
}

func PURLSearch(indexName string, instance packageurl.PackageURL) ([]PurlEntry, *Stats, error) {
	startTime := time.Now()

	db, err := DB()
	if err != nil {
		return nil, nil, err
	}

	tableName := strings.ReplaceAll(indexName, "-", "_")
	query := fmt.Sprintf("SELECT name, version, purl, cves, vulnerabilities FROM `%s` WHERE purl LIKE ?", tableName)

	// Strip qualifiers and normalize the version for searching.
	//
	// The DB stores PURLs without distro qualifiers (e.g. pkg:apk/alpine/musl@1.1.22-3.10),
	// but SBOMs include qualifiers and Alpine APK revision suffixes
	// (e.g. pkg:apk/alpine/musl@1.1.22-r3?arch=x86_64&distro=alpine-3.10.0).
	//
	// For APK packages, the revision suffix "-rN" must be stripped so that the
	// upstream version ("1.1.22") is used as the search prefix, which then matches
	// all distro-specific rows in the DB (1.1.22-3.10, 1.1.22-3.11, …).
	searchVersion := instance.Version
	if instance.Type == "apk" {
		if idx := strings.LastIndex(searchVersion, "-r"); idx != -1 {
			if suffix := searchVersion[idx+2:]; len(suffix) > 0 && isDigits(suffix) {
				searchVersion = searchVersion[:idx]
			}
		}
	}

	searchPURL := packageurl.PackageURL{
		Type:      instance.Type,
		Namespace: instance.Namespace,
		Name:      instance.Name,
		Version:   searchVersion,
		Subpath:   instance.Subpath,
	}.String()

	rows, err := db.Query(query, "%"+searchPURL+"%")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			_ = err
		}
	}()

	// Deduplicate by (name, version, CVE) — APK searches may return one row per
	// Alpine distro release (3.10, 3.11, …) all with the same vulnerabilities.
	type vulnKey struct{ name, version, cve string }
	seen := make(map[vulnKey]struct{})

	var results []PurlEntry
	for rows.Next() {
		var result PurlEntry
		var purlJSON, cvesJSON, vulnerabilitiesJSON sql.NullString

		err := rows.Scan(
			&result.Name,
			&result.Version,
			&purlJSON,
			&cvesJSON,
			&vulnerabilitiesJSON,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if purlJSON.Valid {
			if err := json.Unmarshal([]byte(purlJSON.String), &result.Purl); err != nil {
				return nil, nil, fmt.Errorf("failed to unmarshal purl: %w", err)
			}
		}
		if cvesJSON.Valid {
			if err := json.Unmarshal([]byte(cvesJSON.String), &result.CVEs); err != nil {
				return nil, nil, fmt.Errorf("failed to unmarshal cves: %w", err)
			}
		}
		if vulnerabilitiesJSON.Valid {
			if err := json.Unmarshal([]byte(vulnerabilitiesJSON.String), &result.Vulnerabilities); err != nil {
				return nil, nil, fmt.Errorf("failed to unmarshal vulnerabilities: %w", err)
			}
		}

		if len(result.Vulnerabilities) == 0 {
			continue
		}

		// Filter out rows where every vulnerability has already been seen.
		var unique []sdk.PurlVulnerability

		for _, v := range result.Vulnerabilities {
			k := vulnKey{result.Name, result.Version, v.Detection}
			if _, exists := seen[k]; !exists {
				seen[k] = struct{}{}
				unique = append(unique, v)
			}
		}
		if len(unique) == 0 {
			continue
		}
		result.Vulnerabilities = unique
		results = append(results, result)
	}

	stats := &Stats{
		Duration: time.Since(startTime),
	}

	return results, stats, nil
}

func isDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}
