package services

import "github.com/gostalt/framework/service"

// Providers is a list of ServiceProviders that are registered
// and booted by the app when it is launched.
var Providers = []service.Provider{
	&AppServiceProvider{},
	&AuthServiceProvider{},
	&DatabaseServiceProvider{},
	&LoggingServiceProvider{},
	&RouteServiceProvider{},
	&SchedulerServiceProvider{},
	&SessionServiceProvider{},
	&TLSServiceProvider{},
	&ViewServiceProvider{},

	// Below are non-local services that are required by the
	// Gostalt framework. Remove at your own peril!
	// TODO: This needs thinking about...
	// &service.FileGeneratorServiceProvider{},
}
