package services

import (
	"fmt"
	"gostalt/config"
	"os"

	"github.com/sarulabs/di"
	"github.com/tmus/logger"
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
	default:
		l = StdOutLogger{}
	}

	return
}

type StdOutLogger struct{}

func (l StdOutLogger) Alert(p []byte) {
	l.log("Alert", p)
}

func (l StdOutLogger) Critical(p []byte) {
	l.log("Critical", p)
}

func (l StdOutLogger) Debug(p []byte) {
	l.log("Debug", p)
}

func (l StdOutLogger) Emergency(p []byte) {
	l.log("Emergency", p)
}

func (l StdOutLogger) Error(p []byte) {
	l.log("Error", p)
}

func (l StdOutLogger) Info(p []byte) {
	l.log("Info", p)
}

func (l StdOutLogger) Notice(p []byte) {
	l.log("Notice", p)
}

func (l StdOutLogger) Warning(p []byte) {
	l.log("Warning", p)
}

func (l StdOutLogger) log(level string, p []byte) {
	fmt.Fprintf(
		os.Stdout,
		"%s: %s\n",
		level,
		p,
	)
}
