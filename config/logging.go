package config

func logging(env map[string]string) map[string]string {
	return map[string]string{
		// Driver determines where messages should be printed to
		// when jww is used for logging. Choices are:
		//   - single: a single log file
		//   - daily: a log file for each day
		//   - stdout: logs are printed to the command line
		"driver": env["LOG_DRIVER"],

		// Dir determines the directory to save log files to.
		"dir": "./storage/logs/",
	}
}
