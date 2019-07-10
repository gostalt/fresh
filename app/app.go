package app

import (
	"fmt"
	"gostalt/app/config"
	"gostalt/app/services"
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
	builder.Add(services.Services...)

	return &App{
		Container: builder.Build(),
	}
}

// Config is a convenience method around the container's "env"
// key that makes it easy to retrieve a particular setting.
func (a *App) Config(domain string, key string) string {
	cfg := config.Load(a.Container)
	return cfg[domain][key]
}

// Run uses the values from the .env file in the root directory
// to start the server.
func (a *App) Run() error {
	r := a.Container.Get("router").(*mux.Router)

	srv := &http.Server{
		Handler:      r,
		Addr:         a.Config("main", "address"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting web server on %s\n", a.Config("main", "address"))
	return srv.ListenAndServe()
}
