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

func TestGetBackups(t *testing.T) {
	req := httptest.NewRequest("GET", "/backup", nil)
	w := httptest.NewRecorder()

	mockJson := `{"_benchmark":0.02508,"data":[{"name":"a10","description":"A10 Networks Security Advisories","href":"https://api.vulncheck.com/v3/index/a10"}]}`

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockJson)
	})

	handler.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)

	var backupResp BackupsResponse
	err := json.Unmarshal(body, &backupResp)
	if err != nil {
		t.Error("error unmarshalling response")
	}

	t.Run("indexes response is parsed", func(t *testing.T) {
		assert.Equal(t, "a10", backupResp.Data[0].Name)
	})
}

// LIVE TESTS - these tests require a valid token to run
// func TestGetBackupsLive(t *testing.T) {
// 	t.Run("GetBackups", func(t *testing.T) {
// 		client := Connect("", "")

// 		data, err := client.GetBackups()
// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		if data == nil {
// 			t.Errorf("Expected a string, got %v", data)
// 		}
// 	})
// }
