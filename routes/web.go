package routes

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

func webRoutes(r *mux.Router, container di.Container) {
	s := r.PathPrefix("/").Subrouter()

	views := container.Get("views").(*template.Template)

	s.
		Methods(http.MethodGet).
		Path("/hello").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			views.ExecuteTemplate(w, "welcome.html", nil)
		})
}
