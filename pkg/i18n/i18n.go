package i18n

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
	ScanPurlsFound             string

	ErrorNoToken      string
	ErrorUnauthorized string
}

var C Copy

func Init() {
	C = En
	// TODO: after a 2nd language is added, detect the system language and set the lang variable accordingly
	// look at the LANG environment variable
	// bonus: missing keys of the 2nd language fallback to En
}
