package app

import (
	"fmt"
	"gostalt/app/services"
	"gostalt/config"
	"log"
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

// Make creates and instance of our app and fills the Container.
func Make() *App {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatalln(err)
	}

	config.Load(env)

	builder, _ := di.NewBuilder()

	for _, provider := range services.Providers {
		provider.Register(builder)
	}

	a := &App{
		Container: builder.Build(),
	}
	return a
}

// Run uses the values from the .env file in the root directory
// to start the server.
func (a *App) Run() error {
	r := a.Container.Get("router").(*mux.Router)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Get("app", "address"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting web server on %s\n", config.Get("app", "address"))
	return srv.ListenAndServe()
}
