package services

import (
	"encoding/gob"
	"gostalt/app/entity"
	"gostalt/app/http"

	"github.com/gorilla/sessions"
	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"
)

type AuthServiceProvider struct {
	service.BaseProvider
}

func (p AuthServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "auth",
		Build: func(c di.Container) (interface{}, error) {
			store := c.Get("session").(*sessions.CookieStore)
			return http.Auth{Store: store}, nil
		},
	})
}

func (p AuthServiceProvider) Boot(c di.Container) {
	// Adding the User entity to gob allows it to be serialised
	// and deserialised as part of the session flow.
	gob.Register(&entity.User{})
}
