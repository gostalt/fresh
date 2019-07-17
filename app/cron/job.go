package cron

import (
	"time"

	"github.com/tmus/logger"
)

type Jobber interface {
	Handle() error
	ShouldFire() bool
}

type SayHello struct {
	logger logger.Logger
}

func MakeSayHello(l logger.Logger) *SayHello {
	return &SayHello{
		logger: l,
	}
}

func (s SayHello) Handle() error {
	s.logger.Info([]byte("Hello!"))

	return nil
}

func (SayHello) ShouldFire() bool {
	if time.Now().Minute() == 0 {
		return true
	}
	return false
}
