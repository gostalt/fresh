package config

func session(env map[string]string) map[string]string {
	return map[string]string{
		// Key is the encryption key used to secure your sessions.
		// Do not commit a hardcoded value to version control, or
		// disclose it in any way.
		"key": env["APP_KEY"],
	}
}
