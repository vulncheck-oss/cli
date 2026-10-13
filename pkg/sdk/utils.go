package sdk

import (
	"encoding/json"
	"net/http"
)

// UnmarshalJSON decodes all three error shapes: v3's {"error": true, "errors": [...]},
// v4's {"error": "<message>"} and the NVD-compatible endpoints' {"message": "<message>",
// "status": 400}. The string forms are promoted into Errors because that is the only
// field defaultMessageFor reads.
func (m *MetaError) UnmarshalJSON(data []byte) error {
	var raw struct {
		Error   json.RawMessage `json:"error"`
		Errors  []string        `json:"errors"`
		Message string          `json:"message"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.Errors = raw.Errors

	if len(m.Errors) == 0 && raw.Message != "" {
		m.Error = true
		m.Errors = []string{raw.Message}
	}

	if len(raw.Error) == 0 {
		return nil
	}

	var flag bool
	if err := json.Unmarshal(raw.Error, &flag); err == nil {
		m.Error = flag
		return nil
	}

	var message string
	if err := json.Unmarshal(raw.Error, &message); err == nil && message != "" {
		m.Error = true
		if len(m.Errors) == 0 {
			m.Errors = []string{message}
		}
	}
	return nil
}

// MetaError is a struct that represents the error response from the API
func handleErrorResponse(resp *http.Response) error {
	var metaError MetaError
	_ = json.NewDecoder(LimitedBody(resp.Body)).Decode(&metaError)

	return ReqError{
		StatusCode: resp.StatusCode,
		Reason:     metaError,
	}
}
