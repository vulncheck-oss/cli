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

	AuthShort:       "Authenticate vc with the VulnCheck portal",
	AuthStatusShort: "Check authentication status",
	AuthStatusLong:  "Check if you're currently authenticated and if so, display the account information",
	AuthLoginShort:  "Authenticate with the VulnCheck portal",
	AuthLoginLong: heredoc.Docf(`
			Authenticate with a VulnCheck account.

			The default authentication mode is a web-based browser flow.

			Alternatively, use %[1]stoken%[1]s to specify an issued token directly.

			Alternatively, vc will use the authentication token found in the %[1]sVC_TOKEN%[1]s environment variable.
			This method is most suitable for "headless" use of vc such as in automation.
		`, "`"),
	AuthLoginExample: heredoc.Doc(`
			# Start interactive authentication
			$ vc auth login

			# Authenticate with vulncheck.com by passing in a token
			$ vc auth login token vulncheck_******************
	`),
	AuthLoginErrorCI: "This command is interactive and cannot be run in a CI environment, use the VC_TOKEN environment variable instead",

	AuthLoginToken: "Connect a VulnCheck account using an authentication token",
	AuthLoginWeb:   "Log in with a VulnCheck account using a web browser",

	AuthLogoutShort:             "Invalidate and remove your current authentication token",
	AuthLogoutTokenRemoved:      "Token invalidated and removed",
	AuthLogoutErrorFailed:       "Failed to remove token",
	AuthLogoutErrorInvalidToken: "Token was invalid, removing from config",

	IndicesShort: "View indices",

	ListIndicesShort:  "List indices",
	ListIndicesSearch: "Listing %d indices searching for \"%s\"",
	ListIndicesFull:   "Listing %d indices",

	BrowseIndicesShort:  "Browse indices",
	BrowseIndicesSearch: "Listing %d indices searching for \"%s\"",
	BrowseIndicesFull:   "Listing %d indices",

	IndexShort:         "Browse or list an index",
	IndexListShort:     "List documents of a specified index",
	IndexBrowseShort:   "Browse documents of an index interactively",
	IndexErrorRequired: "index name is required",

	BackupShort:         "Download a backup of a specified index",
	BackupUrlShort:      "Get the temporary signed URL of the backup of an index",
	BackupDownloadShort: "Download the backup of an index",

	BackupDownloadInfo:     "Downloading backup of %s, created on %s",
	BackupDownloadProgress: "Downloading backup as %s",
	BackupDownloadComplete: "Backup downloaded successfully",

	CpeShort:               "Look up a specified cpe for any related CVEs",
	CpeExample:             "vc cpe \"%s\"",
	CpeNoCves:              "No CVEs were found for cpe %s",
	CpeCvesFound:           "%d CVEs were found for cpe %s",
	CpeErrorSchemeRequired: "cpe scheme is required",

	PurlShort:               "Look up a specified PURL for any CVES or vulnerabilities",
	PurlExample:             "vc purl \"%s\"",
	PurlErrorSchemeRequired: "purl scheme is required",

	PurlNoVulns:    "No Vulnerabilities were found for purl %s",
	PurlVulnFound:  "1 Vulnerability were found for purl %s",
	PurlVulnsFound: "%d Vulnerabilities were found for purl %s",

	ScanShort:                  "Scan a directory for vulnerabilities",
	ScanExample:                "vc scan /path/to/directory",
	ScanPackagesFound:          "SBOM generated, scanning %d found packages",
	ScanCvesFound:              "Collecting details of %d vulnerabilities found in the %d packages",
	ScanErrorDirectoryRequired: "Error: Directory is required",

	ErrorUnauthorized: "Error: Unauthorized, Try authenticating with: vc auth login",
	ErrorNoToken:      "No token found. Please run `vc auth login` to authenticate or populate the environment variable `VC_TOKEN`.",
}
