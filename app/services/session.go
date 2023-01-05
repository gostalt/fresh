package services

import (
	"fmt"
	"gostalt/config"

	"github.com/gorilla/sessions"
	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"
)

type SessionServiceProvider struct {
	service.BaseProvider
}

func (p SessionServiceProvider) Register(b *di.Builder) error {
	err := b.Add(di.Def{
		Name: "session",
		Build: func(c di.Container) (interface{}, error) {
			store := sessions.NewCookieStore(
				[]byte(config.Get("app", "key")),
			)

			store.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   86400 * 7,
				HttpOnly: true,
			}
			return store, nil
		},
	})

	if err != nil {
		return fmt.Errorf("unable to register session service: %w", err)
	}

	return nil
}
