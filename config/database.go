package config

func database(env map[string]string) domain {
	return domain{
		"driver":              env["DB_DRIVER"],
		"string":              getDbConnectionString(env),
		"migration_directory": "./database/migrations",
	}
}

func getDbConnectionString(env map[string]string) string {
	// If the database driver is sqlite3, then the connection string
	// should just be a memory connection - sqlite is just for testing.
	if env["DB_DRIVER"] == "sqlite3" {
		return "file:ent?mode=memory&cache=shared&_fk=1"
	}

	return env["DB_CONNECTION_STRING"]
}
