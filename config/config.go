package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config map[string]domain

// domain is an area of config for the application that contains a key value list
// of config items.
type domain map[string]string

var cfg map[string]domain

func Load() Config {
	// If the number of items in cfg is greater than zero, then the config has
	// already been loaded and we don't need to reload it.
	if len(cfg) > 0 {
		return cfg
	}

	godotenv.Overload(".env")

	environ := loadEnviron()

	// TODO: Would be ideal to map these dynamically for each
	// file that does exist in the config directory.
	cfg = map[string]domain{
		"app":      app(environ),
		"database": database(environ),
		"logging":  logging(environ),
		"views":    views(environ),
	}

	return cfg
}

// Get returns the value of a key from a specific domain. Note that a non-existent
// domain or key will always return an empty string.
func Get(domain string, key string) string {
	return cfg[domain][key]
}

// loadEnviron turns the current OS environment variables into a map that can be
// used inside each config file.
func loadEnviron() map[string]string {
	environ := make(map[string]string)

	for _, value := range os.Environ() {
		vals := strings.Split(value, "=")

		for _, exclusion := range environExclusions() {
			if vals[0] == exclusion {
				continue
			}

			environ[vals[0]] = strings.Join(vals[1:], "=")
		}
	}

	return environ
}

// environExclusions returns the environment variables that should not be
// available within the Gostalt app.
func environExclusions() []string {
	return []string{
		"SHELL",
		"USER",
		"PATH",
		"LANG",
	}
}
