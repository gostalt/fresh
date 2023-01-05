package routes

import (
	"github.com/gostalt/framework/route"
	. "github.com/gostalt/router"
)

// Web routes are endpoints that are typically accessed using a web browser. The
// routes declared here are made available to your application by being loaded
// inside the services.RouteServiceProvider's Register method.
func Web(r *Router) {
	r.Group(
		// `route.View` is a convenience function that loads the named template
		// from the `resources/views` directory - use dot notation to access
		// any views in subdirectories.
		Get("/", route.View("welcome")),
	)
}
