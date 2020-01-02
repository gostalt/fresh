package http

import (
	"gostalt/app/entity"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sarulabs/di/v2"
)

type Auth struct {
	Store *sessions.CookieStore
}

func (a Auth) Login(w http.ResponseWriter, r *http.Request, user interface{}) error {
	session, err := a.Store.Get(r, "gostalt")
	if err != nil {
		return err
	}

	session.Values["user"] = user
	session.Save(r, w)

	return nil
}

func (a Auth) LoginRoute(w http.ResponseWriter, r *http.Request) {
	store := di.Get(r, "session").(*sessions.CookieStore)
	a.Store = store
	err := a.Login(w, r, entity.User{Name: "Tomy"})
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte("Logged in"))
}
