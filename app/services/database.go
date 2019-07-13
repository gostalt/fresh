package services

import (
	"database/sql"
	"gostalt/config"

	// Import the postgres driver for the database.
	_ "github.com/lib/pq"

	"github.com/sarulabs/di"
	jww "github.com/spf13/jwalterweatherman"
)

type DatabaseServiceProvider struct {
	BaseServiceProvider
}

func (p DatabaseServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "database",
		Build: func(c di.Container) (interface{}, error) {
			db, err := sql.Open(config.Get("database", "driver"), config.Get("database", "string"))
			if err != nil {
				panic(err)
			}
			if err := db.Ping(); err != nil {
				panic(err)
			}

			return db, nil
		},
	})
}

func (p DatabaseServiceProvider) Boot(c di.Container) {
	logger := c.Get("logger").(*jww.Notepad)
	logger.INFO.Println("Database booted")
}
