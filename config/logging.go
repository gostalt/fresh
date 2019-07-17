package config

func logging(env map[string]string) map[string]string {
	return map[string]string{
		// Dir determines the directory to save log files to.
		"dir": "./storage/logs/",
	}
}
