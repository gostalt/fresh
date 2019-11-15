package services

import (
	"context"
	"fmt"
	"gostalt/config"

	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"
	"golang.org/x/crypto/acme/autocert"
)

type TLSServiceProvider struct {
	service.BaseProvider
}

func (p TLSServiceProvider) Register(b *di.Builder) {
	// If the app environment is in production, we the acme/autocert
	// package is added to the container to interact with LE.
	if config.Get("app", "environment") == "production" {
		b.Add(di.Def{
			Name: "autocert",
			Build: func(c di.Container) (interface{}, error) {
				return p.buildAutocertManager(), nil
			},
		})

		return
	}
}

func (p TLSServiceProvider) buildAutocertManager() *autocert.Manager {
	return &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		HostPolicy: func(ctx context.Context, host string) error {
			if host != config.Get("app", "host") {
				return fmt.Errorf("host name not allowed")
			}

			return nil
		},
		Cache: autocert.DirCache(
			config.Get("app", "certificate_directory"),
		),
	}
}
