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

func (p Provider) DefaultRedirect() string {
	return "/"
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
