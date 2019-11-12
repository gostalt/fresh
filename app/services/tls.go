package services

import (
	"context"
	"fmt"
	"gostalt/config"

	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di"
	"golang.org/x/crypto/acme/autocert"
)

type TLSServiceProvider struct {
	service.BaseProvider
}

// certificateDirectory returns the path of a file in the config
// `certificate_directory` property.
func (p TLSServiceProvider) certificateDirectory(file string) string {
	return config.Get("app", "certificate_directory") + "/" + file
}

// Register checks the application's environment and generates
// some TLS certificates if it is not running in production.
func (p TLSServiceProvider) Register(b *di.Builder) {
	// If the app environment is in production, then instead of
	// creating self-signed certificates, we instead use the
	// amazing acme/autocert package to interact with LE.
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

// buildAutocertManager
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
