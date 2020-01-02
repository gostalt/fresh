package command

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"gostalt/app"
	"gostalt/config"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gostalt/logger"
	"github.com/spf13/cobra"
)

var stdout bool

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Prints the application key to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()
		key := config.Get("session", "key")
		l := a.Container.Get("logger").(logger.Logger)

		l.Info(
			[]byte(fmt.Sprintf("Application key is `%s`.", key)),
		)
	},
}

var generateKeyCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new application key",
	Run: func(cmd *cobra.Command, args []string) {
		a := app.Make()
		l := a.Container.Get("logger").(logger.Logger)

		b := make([]byte, 64)
		rand.Read(b[:])
		key := base64.URLEncoding.EncodeToString(b)

		if stdout {
			l.Info(
				[]byte(fmt.Sprintf("Generated key is `%s`", key)),
			)

			return
		}

		writeKeyToEnv(key, l)
	},
}

func writeKeyToEnv(key string, l logger.Logger) {
	var output string

	file, err := ioutil.ReadFile(".env")
	if err != nil {
		l.Notice([]byte(".env file does not exist. Creating."))
		output = `APP_KEY="` + key + `"`
	} else {
		lines := strings.Split(string(file), "\n")

		for i, line := range lines {
			if strings.Contains(line, "APP_KEY") {
				lines[i] = `APP_KEY="` + key + `"`
			}
		}

		output = strings.Join(lines, "\n")
	}

	err = ioutil.WriteFile(".env", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	generateKeyCmd.Flags().BoolVar(&stdout, "stdout", false, "Prints the key to stdout")

	rootCmd.AddCommand(keyCmd)
	keyCmd.AddCommand(generateKeyCmd)
}
