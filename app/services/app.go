package services

import (
	"database/sql"
	"gostalt/app/http/middleware"
	"log"
	"os"

	"github.com/sarulabs/di"
	jww "github.com/spf13/jwalterweatherman"
)

// AppServiceProvider is a more generic ServiceProvider that you
// can use for any misc initialisation that doesn't warrant a
// dedicated ServiceProvider.
type AppServiceProvider struct {
	BaseServiceProvider
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
	{
		Name: "logger",
		Build: func(c di.Container) (interface{}, error) {
			f, err := os.OpenFile("./storage/logs/eg.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalln(err)
			}

			return jww.NewNotepad(
				jww.LevelInfo,
				jww.LevelTrace,
				os.Stdout,
				f,
				"",
				log.Ldate|log.Ltime,
			), nil
		},
	},
}

// Register registers the list of services in the Container's
// build definition.
func (p AppServiceProvider) Register(b *di.Builder) {
	b.Add(services...)
}
