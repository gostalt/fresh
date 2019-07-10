package middleware

import (
	"log"
	"net/http"
	"time"
)

// RequestTimer prints the time a particular request has taken.
func RequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("Request took %s\n", time.Since(start))
		},
	)
}
