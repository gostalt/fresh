package services

import (
	"fmt"

	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"
)

// AppServiceProvider is a more generic ServiceProvider that you
// can use for any misc initialisation that doesn't warrant a
// dedicated ServiceProvider.
type AppServiceProvider struct {
	service.BaseProvider
}

var services = []di.Def{}

// Register registers the list of services in the Container's
// build definition.
func (p AppServiceProvider) Register(b *di.Builder) error {
	if err := b.Add(services...); err != nil {
		return fmt.Errorf("unable to register app service: %w", err)
	}

	return nil
}
