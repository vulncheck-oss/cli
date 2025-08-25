package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleErrorResponse(t *testing.T) {
	metaError := MetaError{
		Error:  true,
		Errors: []string{"error1", "error2"},
	}
	metaErrorJSON, _ := json.Marshal(metaError)

	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusBadRequest)
	resp.Body.Write(metaErrorJSON)

	err := handleErrorResponse(resp.Result())

	reqErr, ok := err.(ReqError)
	if !ok {
		t.Fatalf("expected ReqError, got %T", err)
	}

	if reqErr.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, reqErr.StatusCode)
	}

	expectedErrors := []string{"error1", "error2"}
	if len(reqErr.Reason.Errors) != len(expectedErrors) {
		t.Fatalf("expected %d errors, got %d", len(expectedErrors), len(reqErr.Reason.Errors))
	}
	for i, expectedError := range expectedErrors {
		if reqErr.Reason.Errors[i] != expectedError {
			t.Errorf("expected error %q, got %q", expectedError, reqErr.Reason.Errors[i])
		}
	}
}
