package app

import (
	"crypto/tls"
	"gostalt/app/services"
	"gostalt/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
	"golang.org/x/crypto/acme/autocert"
)

// App is where the magic happens.
type App struct {
	Container di.Container
}

// Make creates an instance of our app and fills the Container
// using the ServiceProviders defined in services.Providers.
func Make() *App {
	// Load the .env file into the app, and use it to populate
	// the different config domains defined in `./config`.
	if env, err := godotenv.Read(); err != nil {
		panic("unable to load environment")
	} else {
		config.Load(env)
	}

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
	srv := &http.Server{
		Handler:      a.Container.Get("router").(*mux.Router),
		Addr:         config.Get("app", "address"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if config.Get("app", "environment") == "production" {
		le := a.Container.Get("autocert").(*autocert.Manager)
		srv.TLSConfig = &tls.Config{GetCertificate: le.GetCertificate}

		// In production, the server address should just be 443.
		srv.Addr = ":443"

		// A non-TLS ListenAndServe is used here so Let's Encrypt
		// can use HTTP Challenge to make a new certificate.
		go http.ListenAndServe(":80", le.HTTPHandler(nil))

		return srv.ListenAndServeTLS("", "")
	}

	return srv.ListenAndServeTLS(
		config.Get("app", "certificate_directory")+"/cert.pem",
		config.Get("app", "certificate_directory")+"/key.pem",
	)
}
