package config

func app(env map[string]string) map[string]string {
	return map[string]string{
		// Name is the name of your app.
		"name": env["APP_NAME"],

		// Environment dictates the "enviroment" that the app is
		// running in. This can be used to configure services
		// depending on the status of the application.
		"environment": env["APP_ENV"],

		// Address is the URL. Locally, this is likely to be a
		// localhost value, but this should be changed for an
		// app that is running in production.
		"address": env["ADDRESS"],

		// Certificate directory is the folder, relative to the
		// project root, that certificates should be stored in
		// when using TLS to serve the app locally.
		"certificate_directory": "./storage/certs",
	}
}
