package command

import (
	"context"
	"gostalt/app"
	"gostalt/app/entity"
	"gostalt/app/entity/migrate"
	"log"
	"os"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/spf13/cobra"
)

var dryrun bool
var destructive bool

func init() {
	migrateUpdateCmd.Flags().BoolVar(&dryrun, "dryrun", false, "Print SQL to logs")
	migrateUpdateCmd.Flags().BoolVar(&destructive, "destructive", false, "Allow columns and indexes to be dropped")

	migrateCmd.AddCommand(migrateUpdateCmd)

	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Handle database migrations",
}

var migrateUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update database tables using current entity data",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()
		client := a.Container.Get("entity-client").(*entity.Client)

		ctx := context.Background()
		var opts []schema.MigrateOption

		if destructive {
			opts = []schema.MigrateOption{
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
			}
		}

		if dryrun {
			// If the dryrun flag is checked, print the SQL to the
			// stdout and stop.
			if err := client.Schema.WriteTo(ctx, os.Stdout, opts...); err != nil {
				log.Fatalln("Unable to dryrun migration:", err)
			}

			return
		}

		if err := client.Schema.Create(context.Background(), opts...); err != nil {
			log.Fatalln(err)
		}
	},
}
