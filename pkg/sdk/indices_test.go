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

func TestGetIndices(t *testing.T) {
	req := httptest.NewRequest("GET", "/index", nil)
	w := httptest.NewRecorder()

	mockJson := `{"_benchmark":0.02508,"data":[{"name":"a10","description":"A10 Networks Security Advisories","href":"https://api.vulncheck.com/v3/index/a10"}]}`

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockJson)
	})

	handler.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	var indexResp IndicesResponse
	err := json.Unmarshal(body, &indexResp)
	if err != nil {
		t.Error("error unmarshalling response")
	}

	t.Run("indexes response is parsed", func(t *testing.T) {
		assert.Equal(t, "a10", indexResp.Data[0].Name)
	})
}
