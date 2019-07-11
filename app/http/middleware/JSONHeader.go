package middleware

import "net/http"

// JSONHeader middleware adds the `Content-Type` header to the
// response with a value of `application/json`, meaning that
// all routes that implement this middleware return JSON.
func JSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}
