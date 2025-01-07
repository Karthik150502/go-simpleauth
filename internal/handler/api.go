package handler

import (
	"encoding/json"
	"net/http"
	"simple_auth/internal/errorhandling"
	"simple_auth/internal/routes/auth"
	"simple_auth/internal/routes/user"
	"simple_auth/internal/types"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
)

func Handler(r *chi.Mux) {
	// Parent middleware
	r.Use(chimiddle.StripSlashes) //Removes any trailing slashes in the endpoint.

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(&types.MessageResponse{
				Message: "Welcome to the Simple Auth API server",
				StatusCode:    http.StatusOK,
			})
			if err != nil {
				log.Error(err)
				errorhandling.InternalErrorHandler(w)
				return
			}
		})
		r.Route("/auth", auth.Authhandler)
		r.Route("/user", user.UserHandler)
	})
}
