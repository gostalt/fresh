package routes

import (
	"gostalt/routes/middleware"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

func webRoutes(r *mux.Router, container di.Container) {
	s := r.PathPrefix("/").Subrouter()

	// Middleware can be defined on the subrouter, and this
	// affects all routes then registered.
	s.Use(middleware.Log, middleware.Varsity)

	views := container.Get("views").(*template.Template)

	s.
		// but can middleware be defined on a single route?
		Methods(http.MethodGet).
		Path("/hello").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			views.ExecuteTemplate(w, "welcome.html", nil)
		})
}
