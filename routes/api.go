package routes

import (
	"net/http"

	"gostalt/app/http/handler/api"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// APIRoutes is where API specific routes for the app are defined.
// All API routes are prefixed with "/api/" automatically.
func APIRoutes(r *mux.Router, c di.Container) {
	r.Methods(http.MethodGet).Path("/welcome").HandlerFunc(api.Hello)
	r.Methods(http.MethodPost).Path("/post-test").Handler(api.PostTest{})
}
