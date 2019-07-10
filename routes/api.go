package routes

import (
	"net/http"

	"gostalt/handler/api"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// api is the place to define API specific routes for the application.
// All API routes are prefixed with "/api/" automatically, but this
// behaviour can be changed by editing the `makeSubRouter` call.
//
// By default, the parent router and the container are injected
// into the api function. This allows the api routes to branch
// off of the base router, and resolve any dependencies out of
// the container so that we can use them. You can add additional
// dependencies in the call to api in routes.go
func apiRoutes(r *mux.Router, c di.Container) {
	s := r.PathPrefix("/api").Subrouter()

	// Here is the first route for the application. At the moment
	// this is just calling Gorilla's Mux behind the scenes, so
	// feel free to chain any methods onto the new route.
	s.
		Methods(http.MethodGet).
		Path("/welcome").
		Handler(api.Hello{Container: c})

	// ...
}
