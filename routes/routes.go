package routes

import (
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// Load uses the router and container to register routes for
// the application.
func Load(
	router *mux.Router,
	container di.Container,
) {
	fileServer(router, "/assets/", "./assets")
	apiRoutes(router, container)
	webRoutes(router, container)
}
