package services

import (
	"context"
	"fmt"
	"gostalt/config"

	"github.com/kabukky/httpscerts"
	"github.com/sarulabs/di"
	"golang.org/x/crypto/acme/autocert"
)

type TLSServiceProvider struct{}

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

	if !p.certificatesExist() {
		if err := p.generateCertificates(); err != nil {
			panic("could not generate certificates")
		}
	}
}

// certificatesExist determines if a certificate and private key
// pair already exist in the certificate directory. If they do,
// they don't need generating again.
func (p TLSServiceProvider) certificatesExist() bool {
	if err := httpscerts.Check(
		p.certificateDirectory("cert.pem"),
		p.certificateDirectory("key.pem"),
	); err != nil {
		return false
	}

	return true
}

// generateCertificates uses the httpscerts library to generate
// a certificate and private key pair for local development.
func (p TLSServiceProvider) generateCertificates() error {
	return httpscerts.Generate(
		p.certificateDirectory("cert.pem"),
		p.certificateDirectory("key.pem"),
		config.Get("app", "address"),
	)
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
