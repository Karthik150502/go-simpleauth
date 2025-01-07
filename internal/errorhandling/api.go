package errorhandling

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

func WriteError(w http.ResponseWriter, message string, code int) {
	var response = Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An unexpected server error occured, please try again later.", http.StatusInternalServerError)
	}
)
