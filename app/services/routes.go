package services

import (
	"gostalt/app/http/middleware"
	"gostalt/routes"
	"net/http"

	"github.com/gostalt/framework/route"
	"github.com/gostalt/framework/service"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di/v2"
)

// RouteServiceProvider is responsible for registering the app's
// routes. That is, the URIs that call a handler.
type RouteServiceProvider struct {
	service.BaseProvider
}

var routeCollections = []*route.Group{
	routes.Web,
	routes.API,
}

// Register creates a new mux.Router instance in the container,
// and then registers the user defined routes inside of it.
func (p RouteServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "router",
		Build: func(c di.Container) (interface{}, error) {
			r := mux.NewRouter()

			// Apply the global middleware stack to the base router.
			for _, middleware := range p.globalMiddlewareStack(c) {
				r.Use(mux.MiddlewareFunc(middleware))
			}

			for _, rc := range routeCollections {
				route.TransformGorilla(r, rc)
			}

			p.registerFaviconRoute(r)
			p.registerAssetsRoute(r)

			return r, nil
		},
	})
}

func (p RouteServiceProvider) registerFaviconRoute(r *mux.Router) {
	r.Path("/favicon.ico").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./public/favicon.ico")
			return
		},
	)
}

func (p RouteServiceProvider) registerAssetsRoute(r *mux.Router) {
	r.PathPrefix("/assets/").Handler(
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("./public")),
		),
	)
}

// globalMiddlewareStack defines a middleware stack for the base
// router of the application. These middleware are ran in the
// order they are defined on every http request to the app.
func (p RouteServiceProvider) globalMiddlewareStack(c di.Container) []route.Middleware {
	containerResolver := middleware.ContainerResolver{c}.Handle

	return []route.Middleware{
		containerResolver,
		middleware.IsInMaintenanceMode,
	}

}
