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

	IndexShort       string
	IndexListShort   string
	IndexBrowseShort string

	BackupShort            string
	BackupUrlShort         string
	BackupDownloadShort    string
	BackupDownloadInfo     string
	BackupDownloadProgress string
	BackupDownloadComplete string

	CpeShort     string
	CpeExample   string
	CpeNoCves    string
	CpeCvesFound string

	PurlShort     string
	PurlExample   string
	PurlNoCves    string
	PurlCvesFound string

	SbomShort     string
	SbomListShort string

	SbomScanShort     string
	SbomScanListShort string

	ErrorNoToken      string
	ErrorUnauthorized string

	ErrorIndexRequired string

	ErrorCpeSchemeRequired  string
	ErrorPurlSchemeRequired string
}

var C Copy

func Init() {
	C = En
	// TODO: after a 2nd language is added, detect the system language and set the lang variable accordingly
	// look at the LANG environment variable
	// bonus: missing keys of the 2nd language fallback to En
}
