package errorhandling

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	StatusCode int
	Message    string
}
type ValidationError struct {
	StatusCode      int               `json:"statusCode"`
	Message         string            `json:"message"`
	ValidationError map[string]string `json: "validationErrors"`
}

func WriteError(w http.ResponseWriter, message string, code int) {
	var response = Error{
		StatusCode: code,
		Message:    message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func WriteValidationError(w http.ResponseWriter, validationErrors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ValidationError{
		StatusCode:      http.StatusBadRequest,
		Message:         "Input validation has failed",
		ValidationError: validationErrors,
	})
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, message string, code int) {
		WriteError(w, message, code)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		WriteError(w, "An unexpected server error occured, please try again later.", http.StatusInternalServerError)
	}
	ValidationErrorHandler = func(w http.ResponseWriter, errors map[string]string) {
		WriteValidationError(w, errors)
	}
)
