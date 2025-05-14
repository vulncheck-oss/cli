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

	OfflineStatusShort string
	OfflineStatusLong  string

	AuthLoginToken string
	AuthLoginWeb   string

	FlagSaveResults      string
	FlagOutputJson       string
	FlagSpecifyFile      string
	FlagSpecifySbomFile  string
	FlagSpecifySbomInput string
	FlagSpecifySbomOnly  string
	SavingResultsStart   string
	SavingResultsEnd     string

	TokenShort string

	ListTokensShort string
	ListTokensFull  string

	BrowseTokensShort string
	BrowseTokens      string

	CreateTokenShort         string
	CreateTokenLabelRequired string
	CreateTokenSuccess       string

	RemoveTokenShort      string
	RemoveTokenSuccess    string
	RemoveTokenIDRequired string

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

	RuleShort                 string
	RuleErrorRequired         string
	RuleExample               string
	RuleErrorRuleNameRequired string

	TagShort                string
	TagExample              string
	TagErrorTagNameRequired string

	PdnsShort                 string
	PdnsExample               string
	PdnsErrorListNameRequired string

	ScanShort                   string
	ScanExample                 string
	ScanErrorDirectoryRequired  string
	ScanSbomStart               string
	ScanSbomLoad                string
	ScanSbomEnd                 string
	ScanSbomLoaded              string
	ScanExtractPurlStart        string
	ScanExtractPurlEnd          string
	ScanScanPurlStart           string
	ScanScanPurlStartOffline    string
	ScanScanPurlProgress        string
	ScanScanPurlProgressOffline string
	ScanScanPurlEnd             string
	ScanScanPurlEndOffline      string
	ScanVulnMetaStart           string
	ScanVulnMetaEnd             string
	ScanNoCvesFound             string
	ScanBenchmark               string

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

	AuthShort:       "Authenticate vulncheck with the VulnCheck portal",
	AuthStatusShort: "Check authentication status",
	AuthStatusLong:  "Check if you're currently authenticated and if so, display the account information",
	AuthLoginShort:  "Authenticate with the VulnCheck portal",
	AuthLoginLong: heredoc.Docf(`
			Authenticate with a VulnCheck account.

			The default authentication mode is a web-based browser flow.

			Alternatively, use %[1]stoken%[1]s to specify an issued token directly.

			Alternatively, vulncheck will use the authentication token found in the %[1]sVC_TOKEN%[1]s environment variable.
			This method is most suitable for "headless" use of vulncheck such as in automation.
		`, "`"),
	AuthLoginExample: heredoc.Doc(`
			# Start interactive authentication
			$ vulncheck auth login

			# Authenticate with vulncheck.com by passing in a token
			$ vulncheck auth login token vulncheck_******************
	`),
	AuthLoginErrorCI: "This command is interactive and cannot be run in a CI environment, use the VC_TOKEN environment variable instead",

	AuthLoginToken: "Connect a VulnCheck account using an authentication token",
	AuthLoginWeb:   "Log in with a VulnCheck account using a web browser",

	AuthLogoutShort:             "Invalidate and remove your current authentication token",
	AuthLogoutTokenRemoved:      "Token invalidated and removed",
	AuthLogoutErrorFailed:       "Failed to remove token",
	AuthLogoutErrorInvalidToken: "Token was invalid, removing from config",

	FlagSaveResults:      "Save Results as a file",
	FlagOutputJson:       "Output JSON Results",
	FlagSpecifyFile:      "Specify the file to save the results to",
	FlagSpecifySbomFile:  "Specify the file to save your SBOM scan to",
	FlagSpecifySbomInput: "Specify an existing SBOM file to scan instead of creating one from a folder",
	FlagSpecifySbomOnly:  "Do not run a scan and only create a SBOM file",
	SavingResultsStart:   "Saving Results to %s",
	SavingResultsEnd:     "Results saved to %s",

	TokenShort: "Manage Tokens",

	ListTokensShort: "List tokens",
	ListTokensFull:  "Listing %d tokens",
	BrowseTokens:    "Browsing %d tokens, ESC or q to quit, c to create a token",

	BrowseTokensShort: "Browse tokens interactively",

	CreateTokenShort:         "Create a token",
	CreateTokenSuccess:       "Token %s created successfully: %s",
	CreateTokenLabelRequired: "Token Label missing or invalid",

	RemoveTokenShort:      "Remove a token",
	RemoveTokenSuccess:    "Token %s removed successfully",
	RemoveTokenIDRequired: "Token ID missing or invalid",

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
	CpeExample:             "vulncheck cpe \"%s\"",
	CpeNoCves:              "No CVEs were found for cpe %s",
	CpeCvesFound:           "%d CVEs were found for cpe %s",
	CpeErrorSchemeRequired: "cpe scheme is required",

	PurlShort:               "Look up a specified PURL for any CVEs or vulnerabilities",
	PurlExample:             "vulncheck purl \"%s\"",
	PurlErrorSchemeRequired: "purl scheme is required",

	PurlNoVulns:    "No Vulnerabilities were found for purl %s",
	PurlVulnFound:  "1 Vulnerability were found for purl %s",
	PurlVulnsFound: "%d Vulnerabilities were found for purl %s",

	RuleShort:                 "Look up a specified rule for Initial Access Intelligence",
	RuleErrorRequired:         "rule name is required",
	RuleExample:               "vulncheck rule \"%s\" \nvulncheck rule \"%s\"",
	RuleErrorRuleNameRequired: "rule name is required",

	TagShort:                "List IP Intelligence Tags",
	TagExample:              "vulncheck tag \"%s\"",
	TagErrorTagNameRequired: "tag name is required",

	PdnsShort:                 "List IP Intelligence Protective DNS records",
	PdnsExample:               "vulncheck pdns \"%s\"",
	PdnsErrorListNameRequired: "list name is required",

	ScanShort:                   "Scan a directory for vulnerabilities",
	ScanExample:                 "vulncheck scan /path/to/directory",
	ScanSbomStart:               "Generating SBOM",
	ScanSbomLoad:                "Loading SBOM file %s",
	ScanSbomEnd:                 "SBOM created",
	ScanSbomLoaded:              "SBOM file loadded",
	ScanExtractPurlStart:        "Extracting PURLs",
	ScanExtractPurlEnd:          "%d PURLs extracted",
	ScanScanPurlStart:           "Scanning PURLs",
	ScanScanPurlStartOffline:    "[OFFLINE] Scanning PURLs",
	ScanScanPurlProgress:        "Scanning PURLs [%d/%d]",
	ScanScanPurlProgressOffline: "[OFFLINE] Scanning PURLs [%d/%d]",
	ScanScanPurlEnd:             "Scanning PURLs: %d vulns found in %d packages",
	ScanScanPurlEndOffline:      "[OFFLINE] Scanning PURLs: %d vulns found in %d packages",

	ScanVulnMetaStart: "Fetching vulnerability metadata",
	ScanVulnMetaEnd:   "Vulnerability metadata fetched",

	ScanNoCvesFound:            "No vulnerabilities found in %d packages",
	ScanBenchmark:              "Scan completed in %s",
	ScanErrorDirectoryRequired: "Error: Directory is required",

	ErrorUnauthorized: "Error: Unauthorized, Try authenticating with: vulncheck auth login",
	ErrorNoToken:      "No token found. Please run `vulncheck auth login` to authenticate or populate the environment variable `VC_TOKEN`.",

	OfflineStatusShort: "Check the status of the offline database",
	OfflineStatusLong:  "Check the status of the offline database",
}

func Init() {
	C = En
	// TODO: after a 2nd language is added, detect the system language and set the lang variable accordingly
	// look at the LANG environment variable
	// bonus: missing keys of the 2nd language fallback to En
}
