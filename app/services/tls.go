package services

import (
	"gostalt/config"

	"github.com/kabukky/httpscerts"
	"github.com/sarulabs/di"
)

type TLSServiceProvider struct{}

// certificateDirectory returns the path of a file in the config
// `certificate_directory` property.
func (p TLSServiceProvider) certificateDirectory(file string) string {
	return config.Get("app", "certificate_directory") + "/" + file
}

// Register checks the application's environment and generates
// some TLS certificates if it is not running in production.
func (p TLSServiceProvider) Register(_ *di.Builder) {
	if config.Get("app", "env") == "production" {
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
