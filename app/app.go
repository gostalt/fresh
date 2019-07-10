package app

import (
	"fmt"
	"gostalt/app/services"
	"gostalt/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// App is where the magic happens.
type App struct {
	Container di.Container
}

// Make creates and instance of our app and fills the Container.
func Make() *App {
	builder, _ := di.NewBuilder()

	for _, provider := range services.Providers {
		provider.Register(builder)
	}

	builder.Add(services.Services...)

	a := &App{
		Container: builder.Build(),
	}

	config.Load(a.Container)
	return a
}

// Run uses the values from the .env file in the root directory
// to start the server.
func (a *App) Run() error {
	r := a.Container.Get("router").(*mux.Router)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Get("main", "address"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting web server on %s\n", config.Get("main", "address"))
	return srv.ListenAndServe()
}
