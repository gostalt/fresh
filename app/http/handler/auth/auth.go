package auth

import (
	"log"

	"github.com/gostalt/validate"
)

// getErrorsFromMessage iterates through the validation errors
// provided by Validate and returns them as a simple slice
// of strings.
func getErrorsFromMessage(msgs validate.Message) []string {
	var errors []string

	for _, field := range msgs {
		for _, msg := range field {
			log.Println(msg)
			errors = append(errors, msg)
		}
	}

	return errors
}
