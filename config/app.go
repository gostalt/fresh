package config

import "encoding/base64"

func app(env map[string]string) map[string]string {
	return map[string]string{
		// Name is the name of your app.
		"name": env["APP_NAME"],

		// Key is the encryption key used to secure your application.
		// Do not commit a hardcoded value to version control, or
		// disclose it in any way.
		"key": getAppKey(env),

		// Environment dictates the "environment" that the app is
		// running in. This can be used to configure services
		// depending on the status of the application.
		"environment": env["APP_ENV"],

		// Address is the URL. Locally, this is likely to be a
		// localhost value, but this should be changed for an
		// app that is running in production.
		"address": getAddress(env),

		// Host is separate from the address in that it does not
		// and should not contain a port.
		"host": env["HOST"],

		// Certificate directory is the folder, relative to the
		// project root, that certificates should be stored in
		// when using TLS to serve the app locally.
		"certificate_directory": "./storage/certs",
	}
}

func getAddress(env map[string]string) string {
	address := env["HOST"]

	if env["PORT"] != "" {
		address = address + ":" + env["PORT"]
	}

	return address
}

func getAppKey(env map[string]string) string {
	key, err := base64.URLEncoding.DecodeString(env["APP_KEY"])
	if err != nil {
		panic("unable to decode app key")
	}

	return string(key)
}
