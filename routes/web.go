package routes

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// WebRoutes is the place to define web routes for your app. The
// web middleware stack is applied to all routes automatically.
func WebRoutes(r *mux.Router, container di.Container) {
	views := container.Get("views").(*template.Template)

	r.
		Methods(http.MethodGet).
		Path("/hello/{name}").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			name := r.Form.Get(":name")
			views.ExecuteTemplate(w, "welcome", name)
		})
}
