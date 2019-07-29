package routes

import (
	"net/http"

	"github.com/gostalt/framework/route"

	"gostalt/app/http/handler/api"
)

var API = route.Collection{
	route.Get("welcome", http.Handler(http.HandlerFunc(api.Hello))),
	route.Post("post-test", api.PostTest{}),
}
