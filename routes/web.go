package routes

import (
	"github.com/gostalt/router"

	"gostalt/app/http/handler/web"
	"gostalt/app/http/middleware"
)

func Web(r *router.Router) {
	r.Group(
		r.Get("/", web.Welcome),
		r.Get("/home", web.Home).Middleware(middleware.Authenticate),
	)
}
