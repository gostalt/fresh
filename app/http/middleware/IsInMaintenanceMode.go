package middleware

import (
	"net/http"
	"os"
)

// IsInMaintenanceMode restricts access to your application by
// returning a 503 Service Unavailable response on all requests.
// You can start Maintenance Mode by running
//
//     ./gostalt maintenance up
//
func IsInMaintenanceMode(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Attempt to get details about the maintenance file.
			_, err := os.Stat("./storage/maintenance/maintenance")

			// If the file doesn't exist, the next http handler
			// can be executed, because the app is not currently
			// running in Maintenance Mode.
			if os.IsNotExist(err) {
				next.ServeHTTP(w, r)
				return
			}

			// Otherwise, return a 503 response and stop the app
			// from executing any further.
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		},
	)
}
