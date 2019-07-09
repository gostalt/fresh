package app

import (

	// "github.com/gostalt/container"

	"gostalt/services"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sarulabs/di"
)

type App struct {
	Container   di.Container
	Environment map[string]string
}

// Make creates and instance of our app and fills the Container.
func Make() *App {
	builder, _ := di.NewBuilder()
	builder.Add(services.Services...)

	app := &App{
		Container: builder.Build(),
	}

	// TODO: The container should probably be populated by a
	// function that returns a map of strings and interfaces.

	// db, err := sql.Open(
	// 	"postgres",
	// 	app.Env("DB_CONNECTION", ""),
	// )

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// app.Add("DB", db)
	return app
}

// Run uses the values from the .env file in the root directory
// to start the server.
func (a *App) Run() error {
	r := a.Container.Get("router").(*mux.Router)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

// // Router is a convenient wrapper for returning the Router from
// // the Container. Is is automatically converted a mux.Router.
// func (c Container) Router() *mux.Router {
// 	return c["Router"].(*mux.Router)
// }

// // Views is a convenient wrapper for returning the Views from
// // the Container. It is automatically converted to a html
// // template.Template type, and is ready to be executed.
// func (c Container) Views() *template.Template {
// 	return c["Views"].(*template.Template)
// }

// // DB is a wrapper for returning the DB connection stored in
// // the Container. It is automatically converted to a *sql.DB
// // so you can start interacting with the database instantly.
// func (c Container) DB() *sql.DB {
// 	return c["DB"].(*sql.DB)
// }
