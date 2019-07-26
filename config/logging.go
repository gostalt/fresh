package config

func logging(env map[string]string) map[string]string {
	return map[string]string{
		// Driver is the type of log to use. You're free to make
		// your own and add the appropriate logic in the logging
		// service provider's getLogger method.
		//
		// The default values available are: `stdout` and `file`.
		"driver": "file",

		// Dir determines the directory to save log files to, if
		// a file driver is chosen.
		"dir": "./storage/logs/",
	}
}
