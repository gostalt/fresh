package services

import (
	"gostalt/app/http/middleware"
	"gostalt/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

type RouteServiceProvider struct {
	BaseServiceProvider
}

// globalMiddlewareStack defines a middleware stack for the base
// router of the application. These middleware are ran in the
// order they are defined on every http request to the app.
var globalMiddlewareStack = []mux.MiddlewareFunc{
	// for example:
	middleware.RequestTimer,
}

func (p RouteServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "router",
		Build: func(c di.Container) (interface{}, error) {
			// Create a new instance of a mux.Router and use the
			// globalMiddlewareStack for all incoming requests.
			r := mux.NewRouter()
			r.Use(globalMiddlewareStack...)

			// Register each set of routes with the Router. The
			// routes below are set in the ./routes directory.
			routes.APIRoutes(r, c)
			routes.WebRoutes(r, c)

			r.PathPrefix("/assets/").Handler(
				http.StripPrefix(
					"/assets/",
					http.FileServer(http.Dir("./public")),
				),
			)

			return r, nil
		},
	})
}
