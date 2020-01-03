package routes

import (
	"gostalt/app/http/handler/auth"
	"gostalt/app/http/middleware"
	"net/http"

	"github.com/gostalt/framework/route"
)

var Auth = route.Collection(
	route.Get("login", http.HandlerFunc(auth.Login)),
	route.Post("login", http.HandlerFunc(auth.Login)),
	route.Get("register", http.HandlerFunc(auth.Register)),
	route.Post("register", http.HandlerFunc(auth.Register)),
).Middleware(middleware.RedirectIfAuthenticated)
