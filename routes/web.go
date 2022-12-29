package routes

import (
	. "github.com/gostalt/router"

	"gostalt/app/http/handler/web"
	"gostalt/app/http/middleware"
)

func Web(r *Router) {
	r.Group(
		Get("/", web.Welcome),
		Get("/home", web.Home).Middleware(middleware.Authenticate),
	)
}
