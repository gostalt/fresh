package config

import (
	"github.com/sarulabs/di"
)

var cfg map[string]map[string]string

func Load(c di.Container) map[string]map[string]string {
	// If the number of items in cfg is greater than zero, then
	// the config has already been loaded and we don't need to
	// reload it.
	if len(cfg) > 0 {
		return cfg
	}

	// Otherwise, load the env key from the container and loop
	// over all the configs. Currently these need registering
	// manually.
	//
	// TODO: Make this automatic?
	env := c.Get("env").(map[string]string)
	cfg = map[string]map[string]string{
		"main": app(env),
	}

	return cfg
}
