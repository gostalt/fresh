package services

import (
	"database/sql"
	"gostalt/routes"
	"gostalt/views"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
)

// Services are items that will be added to the DI Container.
// The DI Container uses sarulabs/di.
var Services = []di.Def{
	{
		Name: "router",
		Build: func(c di.Container) (interface{}, error) {
			r := mux.NewRouter()
			routes.Load(r, c)
			return r, nil
		},
	},
	{
		Name: "views",
		Build: func(c di.Container) (interface{}, error) {
			return views.Load("views"), nil
		},
	},
	{
		Name: "env",
		Build: func(c di.Container) (interface{}, error) {
			return godotenv.Read()
		},
	},
	{
		Name: "database",
		Build: func(c di.Container) (interface{}, error) {
			env := c.Get("env").(map[string]string)

			return sql.Open(env["DB_DRIVER"], env["DB_CONNECTION_STRING"])
		},
	},
}
