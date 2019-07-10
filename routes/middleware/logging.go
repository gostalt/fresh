package middleware

import (
	"log"
	"net/http"
	"time"
)

// Log prints the time a particular request has taken.
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Println(time.Since(t))
	})
}
