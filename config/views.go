package config

func views(env map[string]string) map[string]string {
	return map[string]string{
		"path": "resources/views",
		// cache determines whether to cache the view templates. If cache is set
		// to false, the View Service Provider will be built for each incoming
		// request, which can be slow.
		"cache": "false",
	}
}
