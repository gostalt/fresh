package command

import (
	"gostalt/app"
	"gostalt/config"
	"log"

	"time"

	"bitbucket.org/liamstask/goose/lib/goose"

	"github.com/spf13/cobra"
)

// probably roll own migration stuff:
// https://bitbucket.org/liamstask/goose/src/master/lib/goose/migration_go.go
// maybe use a wrapper around goose for now though :(

func init() {
	migrateCmd.AddCommand(migrateCreateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Handle database migrations",
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new database migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.Make()

		name := args[0]
		f, err := goose.CreateMigration(name, "sql", config.Get("database", "migration_directory"), time.Now())
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("%s created\n", f)
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run pending migrations",
	Run: func(cmd *cobra.Command, args []string) {
		// Ugh, ballache.
		// TODO: Make dbconf use env variables rather than dnconf.yml
		conf, err := goose.NewDBConf("database", "development", "")
		if err != nil {
			log.Fatalln(err)
		}

		target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
		if err != nil {
			log.Fatalln(err)
		}

		if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
			log.Fatalln(err)
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Drop the most recent migrations",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := goose.NewDBConf("database", "development", "")
		if err != nil {
			log.Fatalln(err)
		}

		current, err := goose.GetDBVersion(conf)
		if err != nil {
			log.Fatalln(err)
		}

		previous, err := goose.GetPreviousDBVersion(conf.MigrationsDir, current)
		if err != nil {
			log.Fatalln(err)
		}

		if err := goose.RunMigrations(conf, conf.MigrationsDir, previous); err != nil {
			log.Fatalln(err)
		}
	},
}
