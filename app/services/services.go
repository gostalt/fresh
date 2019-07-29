package services

import "github.com/gostalt/framework/service"

// Providers is a list of ServiceProviders that are registered
// and booted by the app when it is launched.
var Providers = []service.Provider{
	&AppServiceProvider{},
	&DatabaseServiceProvider{},
	&TLSServiceProvider{},
	&RouteServiceProvider{},
	&ViewServiceProvider{},
	&LoggingServiceProvider{},
	&SchedulerServiceProvider{},
}
