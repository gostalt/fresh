package config

func maker(env map[string]string) map[string]string {
	return map[string]string{
		// Entity path is the directory where entities generated
		// with `make entity` will be stored.
		"entity_path": "app/entity/",

		// Repository path is the directory where repositories
		// generated with `make repository` will be stored.
		"repository_path": "app/repository/",
	}
}
