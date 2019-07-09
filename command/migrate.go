package command

import (
	"log"

	"bitbucket.org/liamstask/goose/lib/goose"
	"time"

	"github.com/spf13/cobra"
)

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
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		f, err := goose.CreateMigration(name, "sql", "./db/migrations", time.Now())
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
		conf, err := goose.NewDBConf("db", "development", "")
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
		conf, err := goose.NewDBConf("db", "development", "")
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