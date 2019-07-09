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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


// probably roll own migration stuff:
// https://bitbucket.org/liamstask/goose/src/master/lib/goose/migration_go.go
// maybe use a wrapper around goose for now though :(