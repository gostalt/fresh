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

// Run uses the values from the .env file in the root directory
// to start the server.
func (a *App) Run() error {
	r := a.Container.Get("router").(*mux.Router)
	env := a.Container.Get("env").(map[string]string)

	srv := &http.Server{
		Handler:      r,
		Addr:         env["ADDRESS"],
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting web server on %s\n", env["ADDRESS"])
	return srv.ListenAndServe()
}