package auth

import (
	"fmt"
	"net/http"
	"simple_auth/internal/controllers"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
)

func Authhandler(r chi.Router) {
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(("The auth middlware is in action."))
			next.ServeHTTP(w, r)
		})
	})
	r.Post("/signup", controllers.HandleSignup)
	r.Route("/signin", func (router chi.Router){
		router.Use(httprate.Limit(4, time.Minute, httprate.WithKeyFuncs(
			httprate.KeyByIP,
			httprate.KeyByEndpoint,
		)))
		router.Post("/", controllers.HandleSignin)
	})
}
