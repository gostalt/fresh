package routes

import (
	"gostalt/app/http/handler/web"

	"net/http"

	"github.com/gostalt/framework/route"
)

var Web = route.Collection{
	route.Get("/", http.Handler(http.HandlerFunc(web.Welcome))),
	route.Redirect("old", "/"),
}
