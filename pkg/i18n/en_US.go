package i18n

import "github.com/MakeNowJust/heredoc/v2"

var En = Copy{
	AboutInfo: heredoc.Doc(`
					The VulnCheck CLI is a command-line interface for the VulnCheck API
					For more information on our products, please visit https://vulncheck.com
					For API Documentation, please visit https://docs.vulncheck.com
	`),
	InteractiveOnly: "This command is interactive and cannot run in a CI environment, please try %s instead",
	RootLong:        "Work seamlessly with the VulnCheck API.",

	AuthShort:    "Authenticate vc with the VulnCheck portal",
	IndicesShort: "View indices",

	ListIndicesShort:  "List indices",
	ListIndicesSearch: "Listing %d indices searching for \"%s\"",
	ListIndicesFull:   "Listing %d indices",

	BrowseIndicesShort:  "Browse indices",
	BrowseIndicesSearch: "Listing %d indices searching for \"%s\"",
	BrowseIndicesFull:   "Listing %d indices",

	IndexShort:       "Browse or list an index",
	IndexListShort:   "List documents of a specified index",
	IndexBrowseShort: "Browse documents of an index interactively",

	BackupShort:         "Download a backup of a specified index",
	BackupUrlShort:      "Get the temporary signed URL of the backup of an index",
	BackupDownloadShort: "Download the backup of an index",

	BackupDownloadInfo:     "Downloading backup of %s, created on %s",
	BackupDownloadProgress: "Downloading backup as %s",
	BackupDownloadComplete: "Backup downloaded successfully",

	CpeShort:     "Look up a specified cpe for any related CVEs",
	CpeNoCves:    "No CVEs were found for cpe %s",
	CpeCvesFound: "%d CVEs were found for cpe %s",

	PurlShort: "Look up a specified PURL for any CVES or vulnerabilities",

	PurlNoCves:    "No CVEs were found for purl %s",
	PurlCvesFound: "%d CVEs were found for purl %s",

	ErrorNoToken:            "No valid token found",
	ErrorUnauthorized:       "Error: Unauthorized, Try authenticating with: vc auth login",
	ErrorIndexRequired:      "index name is required",
	ErrorCpeSchemeRequired:  "cpe scheme is required",
	ErrorPurlSchemeRequired: "purl scheme is required",
}
