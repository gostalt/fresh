package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sarulabs/di/v2"
)

type Authenticate struct {
	di.Container
}

func (m Authenticate) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			store := di.Get(r, "session").(*sessions.CookieStore)
			sess, err := store.Get(r, "gostalt")

			user := sess.Values["user"]
			if user == "" || user == nil || err != nil {
				// TODO: Make forbidden route customisable.
				http.Redirect(w, r, "/forbidden", 302)
			}

			next.ServeHTTP(w, r)
		},
	)
}
