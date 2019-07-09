package services

import (
	"gostalt/routes"
	"gostalt/views"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
)

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
}
