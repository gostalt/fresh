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
	fileServer(router, "/assets/", "./public")

	apiRoutes(router, container)
	webRoutes(router, container)

	// Want other types of route definitions:
	// - redirect routes
	// - template routes

	// make some kind of load(router, []Route) method.
	// Route should be an interface
	// or maybe make a NewRoute() and tack bits on.

	// for _, route := range webRoutes {
	// route(router.NewRoute())
	// }

	// Route.Get("/user").Handler(xyz)
	// can already use regex on patterns: {name:[a-zA-Z]}

	// make some middleware to call mux.Vars on the request and add them back to the request?
}
