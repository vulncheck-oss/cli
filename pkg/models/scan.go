package models

type ScanResult struct {
	Vulnerabilities []ScanResultVulnerabilities `json:"vulnerabilities"`
}

type ScanResultVulnerabilities struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	CVE               string `json:"cve"`
	InKEV             bool   `json:"in_kev"`
	CVSSBaseScore     string `json:"cvss_base_score"`
	CVSSTemporalScore string `json:"cvss_temporal_score"`
	FixedVersions     string `json:"fixed_versions"`
}
