package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sarulabs/di/v2"
)

// RedirectIfAuthenticated provides the "inverse" of the Authenticate
// middleware: if a user is already logged in, they are prevented
// from accessing a page.
//
// This is used by Gostalt to prevent authenticated users from
// visiting the login and register pages.
func RedirectIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			store := di.Get(r, "session").(*sessions.CookieStore)
			sess, err := store.Get(r, "gostalt")

			user := sess.Values["user"]
			if user == "" || user == nil || err != nil {
				next.ServeHTTP(w, r)
				return
			}

			http.Redirect(w, r, "/home", 302)
		},
	)
}
