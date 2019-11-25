package command

import (
	"gostalt/app"

	"github.com/gostalt/framework/schedule"
	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Runs jobs on a schedule",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()

		scheduler := a.Container.Get("scheduler").(*schedule.Runner)

		scheduler.Run()
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
