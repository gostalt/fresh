package routes

import (
	"github.com/gostalt/framework/route/middleware"
	. "github.com/gostalt/router"

	"gostalt/app/http/handler/api"
)

func API(r *Router) {
	r.Group(
		Get("/hello", api.Welcome),
	).Prefix("api").Middleware(middleware.AddJSONContentTypeHeader)
}
