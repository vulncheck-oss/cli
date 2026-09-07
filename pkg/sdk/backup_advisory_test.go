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
	body := `{"feed":"ghsa","available":true,"url_mrap":"https://mrap/ghsa.zip","url_us-east-1":"https://use1/ghsa.zip","url_eu-west-2":"https://euw2/ghsa.zip","url_ap-southeast-2":"https://apse2/ghsa.zip","url_expires":"2026-09-02T16:42:21Z","url_ttl_minutes":15,"sha256":"01133e67"}`
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
	assert.Equal(t, "https://mrap/ghsa.zip", backup.URL)
}

func TestAdvisoryBackupURLFallsBackThroughRegions(t *testing.T) {
	assert.Equal(t, "https://use1/x.zip", AdvisoryBackup{URLUsEast1: "https://use1/x.zip"}.resolveURL())
	assert.Equal(t, "https://euw2/x.zip", AdvisoryBackup{URLEuWest2: "https://euw2/x.zip"}.resolveURL())
	assert.Empty(t, AdvisoryBackup{}.resolveURL())
}

// The API sends only the regional variants; url is resolved on decode so that
// `jq -r .url` works against a v3 and a v4 payload alike.
func TestAdvisoryBackupExposesAPlainURLField(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, `{"feed":"ghsa","available":true,"url_us-east-1":"https://use1/ghsa.zip"}`)
	}))
	defer srv.Close()

	backup, err := Connect(srv.URL, "tok").GetAdvisoryBackup("ghsa")
	assert.NoError(t, err)

	encoded, err := json.Marshal(backup)
	assert.NoError(t, err)
	assert.Contains(t, string(encoded), `"url":"https://use1/ghsa.zip"`)
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
