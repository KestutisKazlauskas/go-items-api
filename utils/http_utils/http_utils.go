package http_utils

import (
	"net/http"
	"encoding/json"
	"github.com/KestutisKazlauskas/go-utils/rest_errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "applicatoin/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err *rest_errors.RestErr) {
	RespondJson(w, err.Status, err)
}