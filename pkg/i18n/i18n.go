package i18n

import "github.com/MakeNowJust/heredoc/v2"

type Copy struct {
	AboutInfo       string
	InteractiveOnly string
	RootLong        string

	AuthShort                   string
	AuthStatusShort             string
	AuthStatusLong              string
	AuthLoginShort              string
	AuthLoginLong               string
	AuthLoginExample            string
	AuthLoginErrorCI            string
	AuthLogoutShort             string
	AuthLogoutTokenRemoved      string
	AuthLogoutErrorFailed       string
	AuthLogoutErrorInvalidToken string

	AuthLoginToken string
	AuthLoginWeb   string

	FlagSaveResults    string
	FlagSpecifyFile    string
	SavingResultsStart string
	SavingResultsEnd   string

	IndicesShort string

	ListIndicesShort  string
	ListIndicesSearch string
	ListIndicesFull   string

	BrowseIndicesShort  string
	BrowseIndicesSearch string
	BrowseIndicesFull   string

	IndexShort         string
	IndexListShort     string
	IndexBrowseShort   string
	IndexErrorRequired string

	BackupShort            string
	BackupUrlShort         string
	BackupDownloadShort    string
	BackupDownloadInfo     string
	BackupDownloadProgress string
	BackupDownloadComplete string

	CpeShort               string
	CpeExample             string
	CpeNoCves              string
	CpeCvesFound           string
	CpeErrorSchemeRequired string

	PurlShort               string
	PurlExample             string
	PurlNoVulns             string
	PurlVulnFound           string
	PurlVulnsFound          string
	PurlErrorSchemeRequired string

	ScanShort                  string
	ScanExample                string
	ScanErrorDirectoryRequired string
	ScanSbomStart              string
	ScanSbomEnd                string
	ScanExtractPurlStart       string
	ScanExtractPurlEnd         string
	ScanScanPurlStart          string
	ScanScanPurlProgress       string
	ScanScanPurlEnd            string
	ScanVulnMetaStart          string
	ScanVulnMetaEnd            string
	ScanNoCvesFound            string
	ScanBenchmark              string

	ErrorNoToken      string
	ErrorUnauthorized string
}

var C Copy

var En = Copy{
	AboutInfo: heredoc.Doc(`
					The VulnCheck CLI is a command-line interface for the VulnCheck API
					For more information on our products, please visit https://vulncheck.com
					For API Documentation, please visit https://docs.vulncheck.com
	`),
	InteractiveOnly: "This command is interactive and cannot run in a CI environment, please try %s instead",
	RootLong:        "Work seamlessly with the VulnCheck API.",

	AuthShort:       "Authenticate vci with the VulnCheck portal",
	AuthStatusShort: "Check authentication status",
	AuthStatusLong:  "Check if you're currently authenticated and if so, display the account information",
	AuthLoginShort:  "Authenticate with the VulnCheck portal",
	AuthLoginLong: heredoc.Docf(`
			Authenticate with a VulnCheck account.

			The default authentication mode is a web-based browser flow.

			Alternatively, use %[1]stoken%[1]s to specify an issued token directly.

			Alternatively, vci will use the authentication token found in the %[1]sVC_TOKEN%[1]s environment variable.
			This method is most suitable for "headless" use of vci such as in automation.
		`, "`"),
	AuthLoginExample: heredoc.Doc(`
			# Start interactive authentication
			$ vci auth login

			# Authenticate with vulncheck.com by passing in a token
			$ vci auth login token vulncheck_******************
	`),
	AuthLoginErrorCI: "This command is interactive and cannot be run in a CI environment, use the VC_TOKEN environment variable instead",

	AuthLoginToken: "Connect a VulnCheck account using an authentication token",
	AuthLoginWeb:   "Log in with a VulnCheck account using a web browser",

	AuthLogoutShort:             "Invalidate and remove your current authentication token",
	AuthLogoutTokenRemoved:      "Token invalidated and removed",
	AuthLogoutErrorFailed:       "Failed to remove token",
	AuthLogoutErrorInvalidToken: "Token was invalid, removing from config",

	FlagSaveResults:    "Save Results as a file",
	FlagSpecifyFile:    "Specify the file to save the results to",
	SavingResultsStart: "Saving Results to %s",
	SavingResultsEnd:   "Results saved to %s",

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
	CpeExample:             "vci cpe \"%s\"",
	CpeNoCves:              "No CVEs were found for cpe %s",
	CpeCvesFound:           "%d CVEs were found for cpe %s",
	CpeErrorSchemeRequired: "cpe scheme is required",

	PurlShort:               "Look up a specified PURL for any CVEs or vulnerabilities",
	PurlExample:             "vci purl \"%s\"",
	PurlErrorSchemeRequired: "purl scheme is required",

	PurlNoVulns:    "No Vulnerabilities were found for purl %s",
	PurlVulnFound:  "1 Vulnerability were found for purl %s",
	PurlVulnsFound: "%d Vulnerabilities were found for purl %s",

	ScanShort:   "Scan a directory for vulnerabilities",
	ScanExample: "vci scan /path/to/directory",

	ScanSbomStart:        "Generating SBOM",
	ScanSbomEnd:          "SBOM created",
	ScanExtractPurlStart: "Extracting PURLs",
	ScanExtractPurlEnd:   "%d PURLs extracted",
	ScanScanPurlStart:    "Scanning PURLs",
	ScanScanPurlProgress: "Scanning PURLs [%d/%d]",
	ScanScanPurlEnd:      "Scanning PURLs: %d vulns found in %d packages",

	ScanVulnMetaStart: "Fetching vulnerability metadata",
	ScanVulnMetaEnd:   "Vulnerability metadata fetched",

	ScanNoCvesFound:            "No vulnerabilities found in %d packages",
	ScanBenchmark:              "Scan completed in %s",
	ScanErrorDirectoryRequired: "Error: Directory is required",

	ErrorUnauthorized: "Error: Unauthorized, Try authenticating with: vci auth login",
	ErrorNoToken:      "No token found. Please run `vci auth login` to authenticate or populate the environment variable `VC_TOKEN`.",
}

func Init() {
	C = En
	// TODO: after a 2nd language is added, detect the system language and set the lang variable accordingly
	// look at the LANG environment variable
	// bonus: missing keys of the 2nd language fallback to En
}
