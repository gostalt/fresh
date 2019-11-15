package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var cfg map[string]map[string]string

func Load() map[string]map[string]string {
	// If the number of items in cfg is greater than zero, then
	// the config has already been loaded and we don't need to
	// reload it.
	if len(cfg) > 0 {
		return cfg
	}

	godotenv.Overload(".env")

	environ := loadEnviron()

	// TODO: Would be ideal to map these dynamically for each
	// file that does exist in the config directory.
	cfg = map[string]map[string]string{
		"app":      app(environ),
		"database": database(environ),
		"logging":  logging(environ),
		"maker":    maker(environ),
	}

	return cfg
}

// Get returns the value of a key from a specific domain. A non-
// existant domain or key will return an empty string.
func Get(domain string, key string) string {
	return cfg[domain][key]
}

// loadEnviron turns the current environment's variables into a
// map that can be used inside each config file.
func loadEnviron() map[string]string {
	environ := make(map[string]string)

	for _, value := range os.Environ() {
		vals := strings.Split(value, "=")

		for _, exclusion := range environExclusions() {
			if vals[0] == exclusion {
				continue
			}

			environ[vals[0]] = vals[1]
		}
	}

	return environ
}

// environExclusions returns the environment variables that should
// not be available within the Gostalt app.
func environExclusions() []string {
	return []string{
		"SHELL",
		"USER",
		"PATH",
		"LANG",
	}
}
