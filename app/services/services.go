package services

import (
	"gostalt/config"

	"github.com/gostalt/framework/service"
)

// Build uses the provided Config to create Service Providers for the application.
func Build(conf config.Config) []service.Provider {
	return []service.Provider{
		service.NewViewServiceProvider(conf["views"]["path"], shouldCacheViews(conf)),
		&AppServiceProvider{},
		&AuthServiceProvider{},
		&DatabaseServiceProvider{},
		&LoggingServiceProvider{},
		&RouteServiceProvider{},
		&SchedulerServiceProvider{},
		&SessionServiceProvider{},
		&TLSServiceProvider{},
		// &ViewServiceProvider{},

		// Below are non-local services that are required by the
		// Gostalt framework. Remove at your own peril!
		// TODO: This needs thinking about...
		// &service.FileGeneratorServiceProvider{},
	}
}

func shouldCacheViews(conf config.Config) bool {
	if conf["views"]["cache"] == "true" {
		return true
	}

	return false
}
