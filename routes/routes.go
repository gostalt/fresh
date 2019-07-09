package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// Load uses the router and container to register routes for
// the application.
func Load(
	router *mux.Router,
	container di.Container,
) {
	router.PathPrefix("/assets/").Handler(
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("./assets")),
		),
	)

	apiRoutes(router, container)
	webRoutes(router, container)
}

func makeSubRouter(prefix string, r *mux.Router) *mux.Router {
	s := r.PathPrefix(prefix).Subrouter()

	return s
}
