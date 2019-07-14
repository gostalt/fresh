package command

import (
	"gostalt/app"
	"os"

	"github.com/spf13/cobra"
)

// serveCmd builds our app and runs it.
var maintenance = &cobra.Command{
	Use:   "maintenance",
	Short: "Controls maintenance of the application",
}

var up = &cobra.Command{
	Use:   "up",
	Short: "Puts the site up for maintenance",
	Run: func(cmd *cobra.Command, args []string) {
		app.Make()

		f, err := os.OpenFile(
			"./storage/maintenance/maintenance",
			os.O_RDWR|os.O_APPEND|os.O_CREATE,
			0666,
		)
		if err != nil {
			panic(err)
		}

		f.Close()
	},
}

var down = &cobra.Command{
	Use:   "down",
	Short: "Stop maintenance on the site",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove("./storage/maintenance/maintenance")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	maintenance.AddCommand(up, down)
	rootCmd.AddCommand(maintenance)
}
