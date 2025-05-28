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
		fmt.Fprintln(w, mockJson)
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
