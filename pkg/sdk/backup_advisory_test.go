package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdvisoryBackupsDecodesTheFeedList(t *testing.T) {
	body := `{"data":[{"name":"epss","href":"https://api.vulncheck.com/v4/backup/epss","available":true,"backup_written_at":"2026-09-02T16:17:49Z"},{"name":"ghsa","href":"x","available":false,"backup_written_at":""}]}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, body)
	}))
	defer srv.Close()

	response, err := Connect(srv.URL, "tok").GetAdvisoryBackups()
	assert.NoError(t, err)
	assert.Len(t, response.GetData(), 2)
	assert.Equal(t, "epss", response.GetData()[0].Name)
	assert.True(t, response.GetData()[0].Available)
	assert.Equal(t, "2026-09-02T16:17:49Z", response.GetData()[0].BackupWrittenAt)
	assert.False(t, response.GetData()[1].Available)
}

// /v4/backup/{feed} is a flat object with no data array, no filename and no
// date_added, and url_ttl_minutes is a number where v3 sends a string.
func TestGetAdvisoryBackupDecodesTheFlatPayload(t *testing.T) {
	body := `{"feed":"ghsa","available":true,"url":"https://serve.vulncheck.com/v4DataBackups/ghsa.zip","url_mrap":"https://mrap/ghsa.zip","url_us-east-1":"https://use1/ghsa.zip","url_eu-west-2":"https://euw2/ghsa.zip","url_ap-southeast-2":"https://apse2/ghsa.zip","url_expires":"2026-09-02T16:42:21Z","url_ttl_minutes":15,"sha256":"01133e67"}`
	var path string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		_, _ = fmt.Fprint(w, body)
	}))
	defer srv.Close()

	backup, err := Connect(srv.URL, "tok").GetAdvisoryBackup("ghsa")
	assert.NoError(t, err)
	assert.Equal(t, "/v4/backup/ghsa", path)
	assert.Equal(t, "ghsa", backup.Feed)
	assert.True(t, backup.Available)
	assert.Equal(t, 15, backup.URLTTLMinutes)
	assert.Equal(t, "01133e67", backup.SHA256)
	assert.Equal(t, "https://serve.vulncheck.com/v4DataBackups/ghsa.zip", backup.URL)
	assert.Equal(t, "https://mrap/ghsa.zip", backup.URLMrap)
}

// The primary url is the CloudFront signed URL while CloudFront signing is on.
// Downloading from url_mrap instead would egress from S3 and drop the
// x-vc-req-id attribution the signed URL carries.
func TestGetAdvisoryBackupUsesTheCloudFrontURL(t *testing.T) {
	cf := "https://serve.vulncheck.com/v4DataBackups/ghsa.zip?x-vc-req-id=abc&Expires=1"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintf(w, `{"feed":"ghsa","available":true,"url":%q,"url_cloudfront":%q,"url_mrap":"https://mrap/ghsa.zip"}`, cf, cf)
	}))
	defer srv.Close()

	backup, err := Connect(srv.URL, "tok").GetAdvisoryBackup("ghsa")
	assert.NoError(t, err)
	assert.Equal(t, cf, backup.URL)
}

// BACKUP_CLOUDFRONT_SIGNING_ENABLED is a live kill-switch on the API. With it
// off the API puts the S3 pre-signed URL in url and omits url_cloudfront, so
// the CLI follows it back to S3 with no change of its own.
func TestGetAdvisoryBackupFollowsTheURLBackToS3(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, `{"feed":"ghsa","available":true,"url":"https://mrap/ghsa.zip","url_mrap":"https://mrap/ghsa.zip"}`)
	}))
	defer srv.Close()

	backup, err := Connect(srv.URL, "tok").GetAdvisoryBackup("ghsa")
	assert.NoError(t, err)
	assert.Equal(t, "https://mrap/ghsa.zip", backup.URL)
}

// The API always sends the regional alternates alongside the primary url;
// decoding must keep them so --json can expose every option.
func TestGetAdvisoryBackupKeepsTheRegionalAlternates(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, `{"feed":"ghsa","available":true,"url":"https://serve.vulncheck.com/ghsa.zip","url_mrap":"https://mrap/ghsa.zip","url_us-east-1":"https://use1/ghsa.zip","url_eu-west-2":"https://euw2/ghsa.zip","url_ap-southeast-2":"https://apse2/ghsa.zip"}`)
	}))
	defer srv.Close()

	backup, err := Connect(srv.URL, "tok").GetAdvisoryBackup("ghsa")
	assert.NoError(t, err)

	encoded, err := json.Marshal(backup)
	assert.NoError(t, err)
	for _, want := range []string{
		`"url":"https://serve.vulncheck.com/ghsa.zip"`,
		`"url_mrap":"https://mrap/ghsa.zip"`,
		`"url_us-east-1":"https://use1/ghsa.zip"`,
		`"url_eu-west-2":"https://euw2/ghsa.zip"`,
		`"url_ap-southeast-2":"https://apse2/ghsa.zip"`,
	} {
		assert.Contains(t, string(encoded), want)
	}
}

// PathEscape, not QueryEscape: a space in a path segment is %20, not "+".
func TestGetAdvisoryBackupEscapesTheFeedAsAPathSegment(t *testing.T) {
	var path string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.EscapedPath()
		_, _ = fmt.Fprint(w, `{"feed":"x"}`)
	}))
	defer srv.Close()

	_, err := Connect(srv.URL, "tok").GetAdvisoryBackup("a b")
	assert.NoError(t, err)
	assert.Equal(t, "/v4/backup/a%20b", path)
}
