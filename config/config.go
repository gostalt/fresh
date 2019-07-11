package config

var cfg map[string]map[string]string

func Load(env map[string]string) map[string]map[string]string {
	// If the number of items in cfg is greater than zero, then
	// the config has already been loaded and we don't need to
	// reload it.
	if len(cfg) > 0 {
		return cfg
	}

	cfg = map[string]map[string]string{
		"app": app(env),
	}

	return cfg
}

// Get returns the value of a key from a specific domain.
func Get(domain string, key string) string {
	return cfg[domain][key]
}
