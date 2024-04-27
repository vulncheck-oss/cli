package models

type ScanResult struct {
	Vulnerabilities []ScanResultVulnerabilities `json:"vulnerabilities"`
}

type ScanResultVulnerabilities struct {
	CVE               string `json:"cve"`
	CVSSBaseScore     string `json:"cvss_base_score"`
	CVSSTemporalScore string `json:"cvss_temporal_score"`
	FixedVersions     string `json:"fixed_versions"`
}
