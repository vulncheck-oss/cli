package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCpe(t *testing.T) {
	req := httptest.NewRequest("GET", "/cpe", nil)
	w := httptest.NewRecorder()

	mockJson := `{"_benchmark":0.131502,"_meta":{"cpe":"cpe:/a:microsoft:internet_explorer:8.0.6001:beta","cpe_struct":{"part":"a","vendor":"microsoft","product":"internet_explorer","version":"8\\.0\\.6001","update":"beta","edition":"*","language":"*","sw_edition":"*","target_sw":"*","target_hw":"*","other":"*"},"timestamp":"2024-02-12T23:11:51.481466662Z","total_documents":15},"data":["CVE-2010-0246","CVE-2010-0490","CVE-2010-1117","CVE-2009-2433","CVE-2008-4127","CVE-2002-2435","CVE-2010-0245","CVE-2010-0492","CVE-2010-0027","CVE-2010-0244","CVE-2010-0248","CVE-2010-0494","CVE-2012-1545","CVE-2010-5071","CVE-2011-2383"]}`

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, mockJson)
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

// TestGetCpeResetsQueryBetweenCalls guards against a copy-paste regression in
// GetCpe: c.Query appends via url.Values.Add, so a reused Client that skipped
// ResetQuery would send ?cpe=<prev>&cpe=<current> on the second call and the
// API keyed off the first param — every call after the first silently
// returned the previous CPE's CVE set. Surfaced by pkg/parity in the offline
// vs online CPE comparison.
func TestGetCpeResetsQueryBetweenCalls(t *testing.T) {
	var received []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = append(received, r.URL.RawQuery)
		_, _ = fmt.Fprintln(w, `{"_benchmark":0,"_meta":{},"data":[]}`)
	}))
	defer srv.Close()

	client := Connect(srv.URL, "tok")

	firstCpe := "cpe:2.3:a:vendor-a:product-a:1.0.0:*:*:*:*:*:*:*"
	secondCpe := "cpe:2.3:a:vendor-b:product-b:2.0.0:*:*:*:*:*:*:*"

	if _, err := client.GetCpe(firstCpe); err != nil {
		t.Fatalf("first GetCpe: %v", err)
	}
	if _, err := client.GetCpe(secondCpe); err != nil {
		t.Fatalf("second GetCpe: %v", err)
	}

	if len(received) != 2 {
		t.Fatalf("expected 2 upstream requests, got %d", len(received))
	}

	// Each request must carry exactly one cpe param, and it must be the
	// current call's value — not an accumulation from earlier calls.
	for i, want := range []string{firstCpe, secondCpe} {
		q, err := url.ParseQuery(received[i])
		if err != nil {
			t.Fatalf("parse request %d query: %v", i, err)
		}
		got := q["cpe"]
		if len(got) != 1 || got[0] != want {
			t.Errorf("request %d: got cpe=%v, want [%q]", i, got, want)
		}
	}
}
