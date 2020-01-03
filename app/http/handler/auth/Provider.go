package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Provider struct {
	store *sessions.CookieStore
}

func NewProvider(store *sessions.CookieStore) Provider {
	return Provider{
		store: store,
	}
}

// DefaultRedirect is the path to redirect to when a user
// successfully logs in or registers a new account.
func (p Provider) DefaultRedirect() string {
	return "/home"
}

func (a Provider) ProcessLogin(w http.ResponseWriter, r *http.Request, user interface{}) error {
	session, err := a.store.Get(r, "gostalt")
	if err != nil {
		return err
	}

	session.Values["user"] = user
	session.Save(r, w)

	return nil
}
