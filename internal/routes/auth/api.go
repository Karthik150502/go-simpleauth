package auth

import (
	"fmt"
	"net/http"
	"simple_auth/internal/controllers"

	"github.com/go-chi/chi"
)

func Authhandler(r chi.Router) {
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(("The auth middlware is in action."))
			next.ServeHTTP(w, r)
		})
	})
	r.Post("/signup", controllers.HandleSignup)
	r.Post("/signin", controllers.HandleSignin)
}
