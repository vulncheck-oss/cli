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

func TestGetCpe(t *testing.T) {
	req := httptest.NewRequest("GET", "/cpe", nil)
	w := httptest.NewRecorder()

	mockJson := `{"_benchmark":0.131502,"_meta":{"cpe":"cpe:/a:microsoft:internet_explorer:8.0.6001:beta","cpe_struct":{"part":"a","vendor":"microsoft","product":"internet_explorer","version":"8\\.0\\.6001","update":"beta","edition":"*","language":"*","sw_edition":"*","target_sw":"*","target_hw":"*","other":"*"},"timestamp":"2024-02-12T23:11:51.481466662Z","total_documents":15},"data":["CVE-2010-0246","CVE-2010-0490","CVE-2010-1117","CVE-2009-2433","CVE-2008-4127","CVE-2002-2435","CVE-2010-0245","CVE-2010-0492","CVE-2010-0027","CVE-2010-0244","CVE-2010-0248","CVE-2010-0494","CVE-2012-1545","CVE-2010-5071","CVE-2011-2383"]}`

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockJson)
	})

	handler.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	var cpeResp CpeResponse
	err := json.Unmarshal(body, &cpeResp)
	if err != nil {
		t.Error("error unmarshalling response")
	}

	t.Run("cpe response is parsed", func(t *testing.T) {
		assert.Equal(t, "microsoft", cpeResp.Meta.CpeMeta.Vendor)
	})
}
