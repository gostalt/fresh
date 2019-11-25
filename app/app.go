package app

import (
	"crypto/tls"
	"errors"
	"fmt"
	"gostalt/app/services"
	"gostalt/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gostalt/logger"
	"github.com/sarulabs/di/v2"
	"golang.org/x/crypto/acme/autocert"
)

type App struct {
	Container di.Container
}

// Make creates an instance of our app and fills the Container
// using the ServiceProviders defined in services.Providers.
func Make() *App {
	config.Load()

	// Create a new builder that will be used to populated and
	// used to create the app dependency injection container.
	builder, err := di.NewBuilder()
	if err != nil {
		panic("unable to create di builder")
	}

	for _, provider := range services.Providers {
		provider.Register(builder)
	}

	app := &App{
		Container: builder.Build(),
	}

	for _, provider := range services.Providers {
		provider.Boot(app.Container)
	}

	return app
}

// Run uses the configured App to start a Web Server.
func (a *App) Run() error {
	address := config.Get("app", "address")
	if config.Get("app", "environment") == "production" {
		address = ":443"
	}

	if address == "" {
		return errors.New("app cannot run. No address can be found in the environment")
	}

	srv := &http.Server{
		Handler:      a.Container.Get("router").(*mux.Router),
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	logger := a.Container.Get("logger").(logger.Logger)
	message := fmt.Sprintf("Server running at %s", address)
	logger.Info([]byte(message))

	if config.Get("app", "environment") == "production" {
		le := a.Container.Get("autocert").(*autocert.Manager)
		srv.TLSConfig = &tls.Config{GetCertificate: le.GetCertificate}

		// A non-TLS ListenAndServe is used here so Let's Encrypt
		// can use HTTP Challenge to make a new certificate.
		go http.ListenAndServe(":80", le.HTTPHandler(nil))
		return srv.ListenAndServeTLS("", "")
	} else {
		return srv.ListenAndServe()
	}
}
