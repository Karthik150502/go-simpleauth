package user

import (
	"fmt"
	"net/http"
	"simple_auth/internal/controllers"
	"simple_auth/internal/middlewares/auth"

	"github.com/go-chi/chi"
)

func UserHandler(r chi.Router) {
	r.Use(auth.AuthMiddleware)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("The user middlware is in action.");
			next.ServeHTTP(w, r)
		})
	})
	r.Post("/balance", controllers.GetUserBalance)
}
