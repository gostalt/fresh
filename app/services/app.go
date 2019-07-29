package services

import (
	"database/sql"
	"gostalt/app/http/middleware"

	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di"
)

// AppServiceProvider is a more generic ServiceProvider that you
// can use for any misc initialisation that doesn't warrant a
// dedicated ServiceProvider.
type AppServiceProvider struct {
	service.BaseProvider
}

var services = []di.Def{
	{
		Name: "TokenAuthentication",
		Build: func(c di.Container) (interface{}, error) {
			mw := middleware.TokenAuthentication{
				DB: c.Get("database").(*sql.DB),
			}

			return mw, nil
		},
	},
}

// Register registers the list of services in the Container's
// build definition.
func (p AppServiceProvider) Register(b *di.Builder) {
	b.Add(services...)
}
