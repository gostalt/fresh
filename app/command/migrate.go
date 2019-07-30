package command

import (
	"database/sql"
	"gostalt/app"
	"gostalt/config"

	"github.com/pressly/goose"

	"github.com/spf13/cobra"
)

func init() {
	migrateCmd.AddCommand(migrateCreateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Handle database migrations",
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new database migration",
	// Args[0] is the name of the migration
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()
		db := a.Container.Get("database-basic").(*sql.DB)

		if err := goose.Create(
			db,
			config.Get("database", "migration_directory"),
			args[0],
			"sql",
		); err != nil {
			panic(err)
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run all pending migrations on the database",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()
		db := a.Container.Get("database-basic").(*sql.DB)

		if err := goose.Up(
			db,
			config.Get("database", "migration_directory"),
		); err != nil {
			panic(err)
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Drop the most recent migration",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()
		db := a.Container.Get("database-basic").(*sql.DB)

		if err := goose.Down(
			db,
			config.Get("database", "migration_directory"),
		); err != nil {
			panic(err)
		}
	},
}
