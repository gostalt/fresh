package services

import (
	"gostalt/config"

	// Import the postgres driver for the database.
	"github.com/gostalt/framework/service"
	"github.com/gostalt/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sarulabs/di"
)

type DatabaseServiceProvider struct {
	service.BaseProvider
}

func (p DatabaseServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "database",
		Build: func(c di.Container) (interface{}, error) {
			db, err := sqlx.Connect(
				config.Get("database", "driver"),
				config.Get("database", "string"),
			)

			if err != nil {
				return nil, err
			}

			return db, nil
		},
	})
}

func (p DatabaseServiceProvider) Boot(c di.Container) {
	logger := c.Get("logger").(logger.Logger)
	logger.Debug([]byte("Database booted"))
}
