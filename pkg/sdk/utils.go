package sdk

import (
	"encoding/json"
	"net/http"
)

// UnmarshalJSON decodes both error shapes: v3's {"error": true, "errors": [...]}
// and v4's {"error": "<message>"}. The string form is promoted into Errors
// because that is the only field defaultMessageFor reads.
func (m *MetaError) UnmarshalJSON(data []byte) error {
	var raw struct {
		Error  json.RawMessage `json:"error"`
		Errors []string        `json:"errors"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.Errors = raw.Errors

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
