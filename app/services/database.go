package services

import (
	"database/sql"
	"fmt"
	"gostalt/app/entity"
	"gostalt/config"
	"time"

	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"

	// Import the postgres driver for the database.
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	entsql "github.com/facebookincubator/ent/dialect/sql"
)

type DatabaseServiceProvider struct {
	service.BaseProvider
}

func (p DatabaseServiceProvider) Register(b *di.Builder) error {
	err := b.Add(
		di.Def{
			Name: "entity-client",
			Build: func(c di.Container) (interface{}, error) {
				db := c.Get("database").(*sql.DB)
				db.SetMaxIdleConns(10)
				db.SetMaxOpenConns(100)
				db.SetConnMaxLifetime(time.Hour)

				drv := entsql.OpenDB(config.Get("database", "driver"), db)
				return entity.NewClient(entity.Driver(drv)), nil
			},
		},
		di.Def{
			Name: "database",
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
		},
	)

	if err != nil {
		return fmt.Errorf("unable to register database service: %w", err)
	}

	return nil
}
