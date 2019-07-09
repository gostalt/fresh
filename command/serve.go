package command

import (
	"gostalt/app"

	"github.com/spf13/cobra"
)

// serveCmd builds our app and runs it.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Gostalt framework",
	Long:  "Runs the Gostalt framework on the address specified in the .env file",
	Run: func(cmd *cobra.Command, args []string) {
		app.Make().Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
