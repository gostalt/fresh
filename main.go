package main

import (
	"gostalt/command"
)

func main() {
	// The entire app is handled by Cobra, so this is all we
	// need in the main function.
	//
	// For usage, build the app and run gostalt --help:
	//     `go build && ./gostalt --help`
	//
	// ... or, run `go run main.go --help`.
	command.Execute()
}
