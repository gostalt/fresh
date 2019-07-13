package services

import (
	"log"

	"github.com/sarulabs/di"
)

// ServiceProvider defines an interface for providers that need
// more complex setup.
type ServiceProvider interface {
	Register(*di.Builder)
	Boot()
}

// Providers is a list of ServiceProviders that are registered
// and booted by the app when it is launched.
var Providers = []ServiceProvider{
	&AppServiceProvider{},
	&DatabaseServiceProvider{},
	&TLSServiceProvider{},
	&RouteServiceProvider{},
	&ViewServiceProvider{},
}

type BaseServiceProvider struct{}

func (p BaseServiceProvider) Boot() {
	log.Println("booting provider")
}
