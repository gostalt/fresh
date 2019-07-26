package config

var cfg map[string]map[string]string

func Load(env map[string]string) map[string]map[string]string {
	// If the number of items in cfg is greater than zero, then
	// the config has already been loaded and we don't need to
	// reload it.
	if len(cfg) > 0 {
		return cfg
	}

	// TODO: Would be ideal to map these dynamically for each
	// file that does exist in the config directory.
	cfg = map[string]map[string]string{
		"app":      app(env),
		"database": database(env),
		"logging":  logging(env),
	}

	return cfg
}

// Get returns the value of a key from a specific domain. A non-
// existant domain or key will return an empty string.
func Get(domain string, key string) string {
	return cfg[domain][key]
}
