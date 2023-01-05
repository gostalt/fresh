package services

import (
	"fmt"
	"gostalt/app/http/middleware"
	"gostalt/routes"
	"net/http"

	"github.com/gostalt/framework/service"
	"github.com/gostalt/router"

	"github.com/sarulabs/di/v2"
)

// RouteServiceProvider is responsible for registering the app's
// routes. That is, the URIs that call a handler.
type RouteServiceProvider struct {
	service.BaseProvider
}

type collectionParser func(*router.Router)

var routeCollections = []collectionParser{
	routes.API,
	routes.Web,
}

func (p RouteServiceProvider) Register(b *di.Builder) error {
	err := b.Add(di.Def{
		Name: "router",
		Build: func(c di.Container) (interface{}, error) {
			r := router.New()
			return r, nil
		},
	})

	if err != nil {
		return fmt.Errorf("unable to register routing service: %w", err)
	}

	return nil
}

func (p RouteServiceProvider) Boot(c di.Container) error {
	resp, err := c.SafeGet("router")
	if err != nil {
		return fmt.Errorf("unable to boot view service: cannot retrieve router from container: %w", err)
	}

	rtr, ok := resp.(*router.Router)
	if !ok {
		return fmt.Errorf("unable to boot view service: router is not of type *router.Router, got %T", rtr)
	}

	rtr.Middleware(p.globalMiddlewareStack(c)...)

	for _, rc := range routeCollections {
		rc(rtr)
	}

	p.registerFaviconRoute(rtr)
	p.registerAssetsRoute(rtr)

	return nil
}

func (p RouteServiceProvider) registerFaviconRoute(r *router.Router) {
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/favicon.ico")
		return
	})
}

func (p RouteServiceProvider) registerAssetsRoute(r *router.Router) {
	// TODO: Add the equivalent back with the Gostalt Router.
	// r.PathPrefix("/assets/").Handler(
	// 	http.StripPrefix(
	// 		"/assets/",
	// 		http.FileServer(http.Dir("./public")),
	// 	),
	// )
}

// globalMiddlewareStack defines a middleware stack for the base
// router of the application. These middleware are ran in the
// order they are defined on every http request to the app.
func (p RouteServiceProvider) globalMiddlewareStack(c di.Container) []router.Middleware {
	containerResolver := middleware.ContainerResolver{c}.Handle

	return []router.Middleware{
		containerResolver,
		middleware.IsInMaintenanceMode,
	}
}
