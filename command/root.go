package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gostalt",
	Short: "Gostalt is a Go framework",
	Long:  "Gostalt is a Go framework",
}

// Execute is the way in to the interactive CLI portion of the app.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
