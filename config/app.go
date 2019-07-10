package config

func app(env map[string]string) map[string]string {
	return map[string]string{
		"address":  env["ADDRESS"],
		"app_name": env["APP_NAME"],
	}
}
