package routes

import (
	"net/http"

	"github.com/gostalt/framework/route"
	mw "github.com/gostalt/framework/route/middleware"

	"gostalt/app/http/handler/web"
	"gostalt/app/http/middleware"
)

var Web = route.Collection(
	route.Get("/", http.HandlerFunc(web.Welcome)),
	route.Get("/home", middleware.Authenticate(http.HandlerFunc(web.Home))),
).Middleware(mw.AddURIParametersToRequest)
