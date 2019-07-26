package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AddURIParametersToRequest uses Gorilla's mux.Vars to load the
// parameters that occur in a URI and saves them to the request
// Form, making it easier to access params in route handlers.
func AddURIParametersToRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Load the parameters from the request. If there
			// are no URI params, we can return early here.
			params := mux.Vars(r)
			if len(params) == 0 {
				next.ServeHTTP(w, r)
				return
			}

			// If params do exist on the request, then parse the
			// request's Form field and add each URI's param to
			// it. To prevent collisions with existing r.Form
			// values, URI params are prepended with `:`.
			r.ParseForm()
			for param, value := range params {
				param = ":" + param
				r.Form.Add(param, value)
			}

			next.ServeHTTP(w, r)
		},
	)
}
