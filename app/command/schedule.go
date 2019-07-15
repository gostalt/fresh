package command

import (
	"gostalt/app"
	"gostalt/app/cron"

	"github.com/spf13/cobra"
)

// serveCmd builds our app and runs it.
var schedule = &cobra.Command{
	Use:   "schedule",
	Short: "Runs jobs on a schedule",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()

		scheduler := a.Container.Get("scheduler").(*cron.Scheduler)

		scheduler.Run()
	},
}

func init() {
	rootCmd.AddCommand(schedule)
}
