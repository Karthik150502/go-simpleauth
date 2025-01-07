package auth

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var token string = r.Header.Get("Authorization")
		// if token == "" {
		// 	// Return the error
		// 	return
		// }
		next.ServeHTTP(w, r)
	})
}
