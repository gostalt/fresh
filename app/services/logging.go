package services

import (
	"gostalt/app/services/logging"
	"gostalt/config"

	"github.com/gostalt/logger"
	"github.com/sarulabs/di"
)

type LoggingServiceProvider struct {
	BaseServiceProvider
}

func (p LoggingServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "logger",
		Build: func(c di.Container) (interface{}, error) {
			var logger logger.Logger
			logger = p.getLogger(c)
			return logger, nil
		},
	})
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
