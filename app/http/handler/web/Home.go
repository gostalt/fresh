package web

import (
	"gostalt/app/entity"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sarulabs/di/v2"
)

// Home is an example route that can be safely deleted. After loggin
// in or registering, users are redirected to /home, so ensure that
// a new value is set in app/http/handler/auth/Provider when deleting
// this route.
func Home(w http.ResponseWriter, r *http.Request) {
	store := di.Get(r, "session").(*sessions.CookieStore)
	session, _ := store.Get(r, "gostalt")
	user := session.Values["user"].(*entity.User)

	views := di.Get(r, "views").(*template.Template)
	views.ExecuteTemplate(w, "home", user)
}
