package command

import (
	"gostalt/app"

	"github.com/gostalt/framework/maker"
	"github.com/spf13/cobra"
)

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Make a new file",
}

var makeHandlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "Make a handler",
	Long: `Use dot notation to create a nested handler. For
example, "api.Welcome" would become api/Welcome.go`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := app.Make()
		handler := args[0]

		m := app.Container.Get("HandlerMaker").(maker.HandlerMaker)
		m.Make(handler)
	},
}

func init() {
	makeCmd.AddCommand(makeHandlerCmd)
	rootCmd.AddCommand(makeCmd)
}
