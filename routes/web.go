package routes

import (
	"net/http"

	"github.com/gostalt/framework/route"
	mw "github.com/gostalt/framework/route/middleware"

	"gostalt/app/http/handler/web"
)

var Web = route.Collection(
	route.Get("/", http.Handler(http.HandlerFunc(web.Welcome))),
).Middleware(mw.AddURIParametersToRequest)
