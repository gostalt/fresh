package services

import (
	"gostalt/app/http/middleware"
	"gostalt/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// RouteServiceProvider is responsible for registering the app's
// routes. That is, the URIs that call a handler.
type RouteServiceProvider struct {
	BaseServiceProvider
}

// globalMiddlewareStack defines a middleware stack for the base
// router of the application. These middleware are ran in the
// order they are defined on every http request to the app.
var globalMiddlewareStack = []mux.MiddlewareFunc{
	// for example:
	middleware.IsInMaintenanceMode,
	middleware.RequestTimer,
}

// Register creates a new mux.Router instance in the container,
// and then registers the user defined routes inside of it.
func (p RouteServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "router",
		Build: func(c di.Container) (interface{}, error) {
			r := mux.NewRouter()

			// Apply the globalMiddlewareStack to the router.
			r.Use(middleware.ContainerResolver{c}.Handle)
			r.Use(globalMiddlewareStack...)

			r.Path("/favicon.ico").HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					http.ServeFile(w, r, "./public/favicon.ico")
					return
				},
			)

			p.registerWebRoutes(r, c)
			p.registerAPIRoutes(r, c)
			p.registerAssetsRoute(r)

			return r, nil
		},
	})
}

func (p RouteServiceProvider) registerWebRoutes(r *mux.Router, c di.Container) {
	web := r.PathPrefix("/").Subrouter()

	web.Use(
		middleware.AddURIParametersToRequest,
	)

	routes.WebRoutes(web, c)
}

func (p RouteServiceProvider) registerAPIRoutes(r *mux.Router, c di.Container) {
	api := r.PathPrefix("/api").Subrouter()

	api.Use(
		middleware.JSONHeader,
		c.Get("TokenAuthentication").(middleware.TokenAuthentication).Handle,
	)

	routes.APIRoutes(api, c)
}

func (p RouteServiceProvider) registerAssetsRoute(r *mux.Router) {
	r.PathPrefix("/assets/").Handler(
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("./public")),
		),
	)
}
