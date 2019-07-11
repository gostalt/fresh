package services

import (
	"database/sql"

	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
)

// ServiceProvider defines an interface for providers that need
// more complex setup.
type ServiceProvider interface {
	Register(*di.Builder)
}

// Providers is a list of ServiceProviders that are registered
// and booted by the app when it is launched.
var Providers = []ServiceProvider{
	&RouteServiceProvider{},
	&ViewServiceProvider{},
}

// Services are items that will be added to the DI Container.
// The DI Container uses sarulabs/di.
var Services = []di.Def{
	// TODO: env needs to be booted earlier in the request cycle, and probably not in the container.
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
