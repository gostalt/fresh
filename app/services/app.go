package services

import (
	"database/sql"

	"github.com/sarulabs/di"
)

// AppServiceProvider is a more generic ServiceProvider that you
// can use for any misc initialisation that doesn't warrant a
// dedicated ServiceProvider.
type AppServiceProvider struct{}

var services = []di.Def{
	{
		Name: "database",
		Build: func(c di.Container) (interface{}, error) {
			env := c.Get("env").(map[string]string)

			return sql.Open(env["DB_DRIVER"], env["DB_CONNECTION_STRING"])
		},
	},
}

// Register registers the list of services in the Container's
// build definition.
func (p AppServiceProvider) Register(b *di.Builder) {
	b.Add(services...)
}
