package routes

import (
	"gostalt/app/http/middleware"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

func WebRoutes(r *mux.Router, container di.Container) {
	s := r.PathPrefix("/").Subrouter()

	// Middleware can be defined on the subrouter, and this
	// affects all routes then registered. Middleware runs
	// in the order that it is defined below.
	s.Use(
		middleware.AddURIParametersToRequest,
	)

	views := container.Get("views").(*template.Template)

	s.
		Methods(http.MethodGet).
		Path("/hello/{name}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			views.ExecuteTemplate(w, "welcome.html", nil)
		})
}
