package sdk

import (
	"encoding/json"
	"net/http"
)

// MetaError is a struct that represents the error response from the API
func handleErrorResponse(resp *http.Response) error {
	var metaError MetaError
	_ = json.NewDecoder(resp.Body).Decode(&metaError)

	return ReqError{
		StatusCode: resp.StatusCode,
		Reason:     metaError,
	}
}
