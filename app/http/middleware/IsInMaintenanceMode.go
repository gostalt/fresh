package middleware

import (
	"net/http"
	"os"
)

func IsInMaintenanceMode(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, err := os.Stat("./storage/maintenance/maintenance")

			// If the file doesn't exist, the next http handler
			// can be executed.
			if os.IsNotExist(err) {
				next.ServeHTTP(w, r)
				return
			}

			w.WriteHeader(http.StatusServiceUnavailable)
			return
		},
	)
}
