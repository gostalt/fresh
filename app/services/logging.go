package services

import (
	"fmt"
	"gostalt/app/services/logging"
	"gostalt/config"

	"github.com/gostalt/framework/service"
	"github.com/gostalt/logger"
	"github.com/sarulabs/di/v2"
)

type LoggingServiceProvider struct {
	service.BaseProvider
}

func (p LoggingServiceProvider) Register(b *di.Builder) error {
	err := b.Add(di.Def{
		Name: "logger",
		Build: func(c di.Container) (interface{}, error) {
			var logger logger.Logger
			logger = p.getLogger(c)
			return logger, nil
		},
	})

	if err != nil {
		return fmt.Errorf("unable to register logging service: %w", err)
	}

	return nil
}

func (p LoggingServiceProvider) getLogger(c di.Container) (l logger.Logger) {
	logType := config.Get("logging", "driver")

	switch logType {
	case "stdout":
		l = logging.StdOut{}
	case "file":
		l = logging.MakeFile(
			config.Get("logging", "dir") + "gostalt.log",
		)
	default:
		l = logging.StdOut{}
	}

	return
}
