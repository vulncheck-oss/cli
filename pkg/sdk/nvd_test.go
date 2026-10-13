package sdk

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const nvdCvesBody = `{"resultsPerPage":1,"startIndex":0,"totalResults":837,"format":"NVD_CVE","version":"2.0",` +
	`"timestamp":"2026-10-13T00:00:00.000","vulnerabilities":[{"cve":{"id":"CVE-2021-44228","vulnStatus":"Analyzed"}}],` +
	`"nextCursor":"tok123"}`

func TestGetNvd(t *testing.T) {
	t.Run("sends the parameters as given and keeps the response as sent", func(t *testing.T) {
		var gotPath, gotQuery string
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotPath, gotQuery = r.URL.Path, r.URL.RawQuery
			_, _ = fmt.Fprint(w, nvdCvesBody)
		}))
		defer srv.Close()

		response, err := Connect(srv.URL, "tok").GetNvd(NvdCvesPath, []NvdParam{
			{Name: "hasKev", Value: "true"},
			{Name: "resultsPerPage", Value: "1"},
			{Name: "cursor", Value: ""},
		})
		require.NoError(t, err)

		assert.Equal(t, "/rest/json/cves/2.0", gotPath)
		assert.Equal(t, "cursor=&hasKev=true&resultsPerPage=1", gotQuery,
			"an empty cursor is sent: it is how a caller opts in to cursor paging")

		assert.JSONEq(t, nvdCvesBody, string(response.Raw()))
		assert.JSONEq(t, `[{"cve":{"id":"CVE-2021-44228","vulnStatus":"Analyzed"}}]`, string(response.Records("vulnerabilities")))
		assert.Equal(t, 1, response.ResultsPerPage)
		assert.Equal(t, 837, response.TotalResults)
		assert.Equal(t, "tok123", response.NextCursor)
	})

	t.Run("a missing record array is an empty one", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = fmt.Fprint(w, `{"resultsPerPage":0,"startIndex":0,"totalResults":0,"sources":null}`)
		}))
		defer srv.Close()

		response, err := Connect(srv.URL, "tok").GetNvd(NvdSourcePath, nil)
		require.NoError(t, err)
		assert.Equal(t, "[]", string(response.Records("sources")))
		assert.Equal(t, "[]", string(response.Records("vulnerabilities")))
	})

	t.Run("a refused request carries the API's reason", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, `{"message":"date range from lastModStartDate to lastModEndDate exceeds the 120-day maximum","status":400}`)
		}))
		defer srv.Close()

		_, err := Connect(srv.URL, "tok").GetNvd(NvdCvesPath, nil)

		var reqErr ReqError
		require.True(t, errors.As(err, &reqErr), "got %T", err)
		assert.Equal(t, http.StatusBadRequest, reqErr.StatusCode)
		assert.Equal(t, []string{"date range from lastModStartDate to lastModEndDate exceeds the 120-day maximum"}, reqErr.Reason.Errors)
	})
}

// TestGetNvdResetsQueryBetweenCalls guards the same regression TestGetCpeResetsQueryBetweenCalls
// does: c.Query appends, so a reused Client must not send one call's parameters with the next.
func TestGetNvdResetsQueryBetweenCalls(t *testing.T) {
	var received []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		received = append(received, r.URL.RawQuery)
		_, _ = fmt.Fprint(w, `{"resultsPerPage":0,"startIndex":0,"totalResults":0,"vulnerabilities":[]}`)
	}))
	defer srv.Close()

	client := Connect(srv.URL, "tok")
	_, err := client.GetNvd(NvdCvesPath, []NvdParam{{Name: "cveId", Value: "CVE-2021-44228"}})
	require.NoError(t, err)
	_, err = client.GetNvd(NvdCvesPath, []NvdParam{{Name: "cveId", Value: "CVE-2024-3400"}})
	require.NoError(t, err)

	assert.Equal(t, []string{"cveId=CVE-2021-44228", "cveId=CVE-2024-3400"}, received)
}
