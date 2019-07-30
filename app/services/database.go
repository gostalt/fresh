package services

import (
	"database/sql"
	"gostalt/config"

	// Import the postgres driver for the database.
	"github.com/gostalt/framework/service"
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

	b.Add(di.Def{
		// TODO: Database Basic is a Go sql.DB item, rather
		// than the superior sqlx.DB. Goose migrations only
		// support sql.DB, so when the migrations are rewritten
		// make sure it supports sql.DB.
		Name: "database-basic",
		Build: func(c di.Container) (interface{}, error) {
			db, err := sql.Open(
				config.Get("database", "driver"),
				config.Get("database", "string"),
			)

			if err != nil {
				return nil, err
			}

			err = db.Ping()
			if err != nil {
				return nil, err
			}

			return db, nil
		},
	})
}
