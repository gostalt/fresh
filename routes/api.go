package routes

import (
	mw "github.com/gostalt/framework/route/middleware"
	"github.com/gostalt/router"

	"gostalt/app/http/handler/api"
)

func API(r *router.Router) {
	r.Group(
		router.Get("/hello", api.Welcome),
	).Prefix("api").Middleware(mw.AddJSONContentTypeHeader)
}
