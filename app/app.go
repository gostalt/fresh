package app

import (
	"fmt"
	"gostalt/app/services"
	"gostalt/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
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

	return &App{
		Container: builder.Build(),
	}
}

// Run uses the configured App to start a Web Server.
func (a *App) Run() error {
	r := a.Container.Get("router").(*mux.Router)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Get("app", "address"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting web server on %s\n", config.Get("app", "address"))
	return srv.ListenAndServeTLS(
		config.Get("app", "certificate_directory")+"/cert.pem",
		config.Get("app", "certificate_directory")+"/key.pem",
	)
}
