package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPurl(t *testing.T) {
	req := httptest.NewRequest("GET", "/purl", nil)
	w := httptest.NewRecorder()

	mockJson := `{"_benchmark":0.046492,"_meta":{"purl_struct":{"type":"hex","namespace":"","name":"coherence","version":"0.1.2","qualifiers":null,"subpath":""},"timestamp":"2024-02-12T22:52:52.548053402Z","total_documents":1},"data":{"cves":["CVE-2018-20301"],"vulnerabilities":[{"detection":"CVE-2018-20301","fixed_version":"0.5.2"}]}}`

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, mockJson)
	})

	handler.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	var purlResp PurlResponse
	err := json.Unmarshal(body, &purlResp)
	if err != nil {
		t.Error("error unmarshalling response")
	}

	t.Run("purl response is parsed", func(t *testing.T) {
		assert.Equal(t, "CVE-2018-20301", purlResp.Data.Cves[0])
	})
}

// TestGetPurls covers the batch endpoint, which had no test. The _meta fields
// matter most: dropping them silently is how a partial scan would come back
// looking complete.
func TestGetPurls(t *testing.T) {
	t.Run("decodes findings alongside unprocessed purls", func(t *testing.T) {
		body := `{"_benchmark":0.01,"_meta":{"timestamp":"2026-09-11T00:00:00Z",` +
			`"total_documents":1,"total_submitted":3,"unprocessed":[` +
			`{"purl":"pkg:generic/openssl@1.1.1","reason":"unsupported_type"},` +
			`{"purl":"not-a-purl-at-all","reason":"unparseable"}]},` +
			`"data":[{"purl":"pkg:hex/coherence@0.1.2","purl_struct":{"name":"coherence","version":"0.1.2"},` +
			`"cves":["CVE-2018-20301"],"vulnerabilities":[{"detection":"CVE-2018-20301","fixed_version":"0.5.2"}]}]}`

		var gotBody []byte
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotBody, _ = io.ReadAll(r.Body)
			_, _ = fmt.Fprint(w, body)
		}))
		defer srv.Close()

		resp, err := Connect(srv.URL, "tok").GetPurls([]string{
			"pkg:hex/coherence@0.1.2",
			"pkg:generic/openssl@1.1.1",
			"not-a-purl-at-all",
		})
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// the endpoint takes a bare JSON array, not an object
		assert.JSONEq(t,
			`["pkg:hex/coherence@0.1.2","pkg:generic/openssl@1.1.1","not-a-purl-at-all"]`,
			string(gotBody))

		assert.Len(t, resp.PurlData, 1)
		assert.Equal(t, "pkg:hex/coherence@0.1.2", resp.PurlData[0].Purl)
		assert.Equal(t, float64(3), resp.Meta.TotalSubmitted)
		assert.Equal(t, []UnprocessedPurl{
			{Purl: "pkg:generic/openssl@1.1.1", Reason: "unsupported_type"},
			{Purl: "not-a-purl-at-all", Reason: "unparseable"},
		}, resp.Meta.Unprocessed)
	})

	t.Run("tolerates an API that does not report unprocessed purls", func(t *testing.T) {
		body := `{"_benchmark":0.01,"_meta":{"timestamp":"2026-09-11T00:00:00Z","total_documents":0},"data":[]}`

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = fmt.Fprint(w, body)
		}))
		defer srv.Close()

		resp, err := Connect(srv.URL, "tok").GetPurls([]string{"pkg:hex/coherence@0.1.2"})
		assert.NoError(t, err)
		assert.Empty(t, resp.Meta.Unprocessed)
		assert.Zero(t, resp.Meta.TotalSubmitted)
	})

	t.Run("surfaces a non-200 as an error", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(w, `{"error":true,"errors":["boom"]}`)
		}))
		defer srv.Close()

		resp, err := Connect(srv.URL, "tok").GetPurls([]string{"pkg:hex/coherence@0.1.2"})
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
