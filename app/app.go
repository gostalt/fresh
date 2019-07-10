package app

import (
	"fmt"
	"gostalt/app/services"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

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

func (a *App) Config(domain string, key string) string {
	env := a.Container.Get("env").(map[string]string)

	// TODO: Rather than hardcoding domains here, register them
	// a la services.Services and define keys in separate files.
	// (like Laravel Config).
	m := map[string]map[string]string{
		"main": {
			"address": env["ADDRESS"],
		},
	}

	return m[domain][key]
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
