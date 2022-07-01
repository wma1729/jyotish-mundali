package views

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
 * HTTP Error
 */
type HTTPError struct {
	StatusCode int    `json:"status"`
	Error      string `json:"error"`
	Detail     string `json:"details"`
}

func GetHTTPError(code int, err error, detail string) *HTTPError {
	return &HTTPError{
		StatusCode: code,
		Error:      err.Error(),
		Detail:     detail,
	}
}

func (httpError *HTTPError) Send(w http.ResponseWriter) {
	data, err := json.MarshalIndent(httpError, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("failed to send HTTP error: %s", err)
	}

	http.Error(w, httpError.Detail, httpError.StatusCode)
}
