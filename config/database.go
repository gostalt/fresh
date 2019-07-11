package config

func database(env map[string]string) map[string]string {
	return map[string]string{
		"driver":              env["DB_DRIVER"],
		"string":              env["DB_CONNECTION_STRING"],
		"migration_directory": "./database/migrations",
	}
}
