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

// The v4 endpoints send {"error": "<message>"} where v3 sends
// {"error": true, "errors": [...]}. The string form must be promoted into
// Errors, because that is the only field the error classifier reads -- decoding
// it anywhere else leaves the user with no message at all.
func TestMetaErrorDecodesBothErrorShapes(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantFlag   bool
		wantErrors []string
	}{
		{
			name:       "v3 shape",
			body:       `{"error":true,"errors":["unauthorized"]}`,
			wantFlag:   true,
			wantErrors: []string{"unauthorized"},
		},
		{
			name:       "v3 shape, entitlement",
			body:       `{"error":true,"errors":["This endpoint requires a valid trial or paid subscription"]}`,
			wantFlag:   true,
			wantErrors: []string{"This endpoint requires a valid trial or paid subscription"},
		},
		{
			name:       "v4 string shape, backup feed not found",
			body:       `{"error":"feed not found: \"nosuchfeed\""}`,
			wantFlag:   true,
			wantErrors: []string{`feed not found: "nosuchfeed"`},
		},
		{
			name:       "v4 string shape, opensearch window",
			body:       `{"error":"failed to query data-source: opensearch response: status 400"}`,
			wantFlag:   true,
			wantErrors: []string{"failed to query data-source: opensearch response: status 400"},
		},
		{
			name:       "no error field",
			body:       `{}`,
			wantFlag:   false,
			wantErrors: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var metaError MetaError
			if err := json.Unmarshal([]byte(tt.body), &metaError); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if metaError.Error != tt.wantFlag {
				t.Errorf("Error = %v, want %v", metaError.Error, tt.wantFlag)
			}
			if len(metaError.Errors) != len(tt.wantErrors) {
				t.Fatalf("Errors = %#v, want %#v", metaError.Errors, tt.wantErrors)
			}
			for i, want := range tt.wantErrors {
				if metaError.Errors[i] != want {
					t.Errorf("Errors[%d] = %q, want %q", i, metaError.Errors[i], want)
				}
			}
		})
	}
}

// An explicit errors array wins over the string form rather than being replaced.
func TestMetaErrorKeepsAnExplicitErrorsArray(t *testing.T) {
	var metaError MetaError
	if err := json.Unmarshal([]byte(`{"error":"ignored","errors":["real message"]}`), &metaError); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(metaError.Errors) != 1 || metaError.Errors[0] != "real message" {
		t.Errorf("Errors = %#v, want [real message]", metaError.Errors)
	}
}
