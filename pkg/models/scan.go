package models

import (
	"github.com/vulncheck-oss/cli/pkg/client"
)

type ScanResult struct {
	Vulnerabilities []ScanResultVulnerabilities `json:"vulnerabilities"`

	// Omitted when empty: this field alone must not change the document
	// vulncheck-oss/action hashes to dedupe PR comments.
	Unprocessed []UnprocessedComponent `json:"unprocessed,omitempty"`
}

// UnprocessedComponent is a component the scan could not assess, and why. These
// never appear in Vulnerabilities, so without reporting them a partial scan
// reads as a clean one.
type UnprocessedComponent struct {
	Purl string `json:"purl"`

	// Passed through from the API. An open set; as of writing
	// "unsupported_type", "unparseable" or "unsupported_distro".
	Reason string `json:"reason"`
}

type PurlDetail struct {
	Purl        string   `json:"purl"`
	PackageType string   `json:"type"`
	Cataloger   string   `json:"cataloger"`
	Locations   []string `json:"locations"`
	SbomRef     string   `json:"sbom_ref"`
}

type ScanResultVulnerabilities struct {
	Name              string                             `json:"name"`
	Version           string                             `json:"version"`
	CVE               string                             `json:"cve"`
	InKEV             bool                               `json:"in_kev"`
	Published         string                             `json:"published"`
	CVSSBaseScore     string                             `json:"cvss_base_score"`
	CVSSTemporalScore string                             `json:"cvss_temporal_score"`
	Metrics           *client.ApiNVD20MetricExtended     `json:"metrics,omitempty"`
	FixedVersions     string                             `json:"fixed_versions"`
	PurlDetail        PurlDetail                         `json:"purl_detail,omitempty"`
	Weaknesses        *[]client.ApiNVD20WeaknessExtended `json:"weaknesses,omitempty"`
	Description       string                             `json:"description,omitempty"`
	CPE               string                             `json:"cpe"`
}
