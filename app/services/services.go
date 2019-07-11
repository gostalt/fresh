package services

import (
	"github.com/sarulabs/di"
)

// ServiceProvider defines an interface for providers that need
// more complex setup.
type ServiceProvider interface {
	Register(*di.Builder)
}

// Providers is a list of ServiceProviders that are registered
// and booted by the app when it is launched.
var Providers = []ServiceProvider{
	&AppServiceProvider{},
	&RouteServiceProvider{},
	&ViewServiceProvider{},
}
