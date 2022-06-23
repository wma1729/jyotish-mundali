package models

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

func (e HTTPError) Send(w http.ResponseWriter) {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}

	w.WriteHeader(e.StatusCode)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("failed to send HTTP error: %s", err)
	}
}
