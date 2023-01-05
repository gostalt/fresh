package services

import (
	"encoding/gob"
	"fmt"
	"gostalt/app/entity"
	"gostalt/app/http/handler/auth"

	"github.com/gorilla/sessions"
	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"
)

type AuthServiceProvider struct {
	service.BaseProvider
}

func (p AuthServiceProvider) Register(b *di.Builder) error {
	err := b.Add(di.Def{
		Name: "auth",
		Build: func(c di.Container) (interface{}, error) {
			store := c.Get("session").(*sessions.CookieStore)
			return auth.NewProvider(store), nil
		},
	})

	if err != nil {
		return fmt.Errorf("unable to register auth service: %w", err)
	}

	return nil
}

func (p AuthServiceProvider) Boot(c di.Container) error {
	// Adding the User entity to gob allows it to be serialised
	// and deserialised as part of the session flow.
	gob.Register(&entity.User{})

	return nil
}
