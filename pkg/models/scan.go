package models

import (
	"github.com/vulncheck-oss/cli/pkg/client"
)

type ScanResult struct {
	Vulnerabilities []ScanResultVulnerabilities `json:"vulnerabilities"`
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
	Description       *[]client.ApiNVD20Description      `json:"description,omitempty"`
	CPE               string                             `json:"cpe"`
}
