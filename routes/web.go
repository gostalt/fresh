package routes

import (
	. "github.com/gostalt/router"
)

type View string

// Web routes are endpoints that are typically accessed using a web browser. The
// routes declared here are made available to your application by being loaded
// inside the services.RouteServiceProvider's Register method.
func Web(r *Router) {
	r.Group(
		// `View` is a convenience function that loads the named template from
		// the `resources/views` directory. Use dot notation to access any
		// views in subdirectories!
		Get("/", View("welcome")),
	)
}
