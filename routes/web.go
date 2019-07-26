package routes

import (
	"gostalt/app/http/handler/web"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// WebRoutes is the place to define web routes for your app. The
// web middleware stack is applied to all routes automatically.
func WebRoutes(r *mux.Router, container di.Container) {
	r.
		Methods(http.MethodGet).
		Path("/hello/{name}").
		HandlerFunc(web.Welcome)
}
