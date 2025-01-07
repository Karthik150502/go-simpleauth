package controllers

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"net/http"
	"simple_auth/internal/errorhandling"
	"simple_auth/internal/types"
)

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	// Get the user balance from the database.

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&types.MessageResponse{
		Message: "Fetched the user balance.",
		StatusCode:    http.StatusOK,
	})
	if err != nil {
		log.Error(err)
		errorhandling.InternalErrorHandler(w)
		return
	}
}
