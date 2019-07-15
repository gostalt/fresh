package cron

import (
	"time"

	jww "github.com/spf13/jwalterweatherman"
)

type Jobber interface {
	Handle() error
	ShouldFire() bool
}

type SayHello struct {
	logger *jww.Notepad
}

func MakeSayHello(l *jww.Notepad) *SayHello {
	return &SayHello{
		logger: l,
	}
}

func (s SayHello) Handle() error {
	s.logger.INFO.Println("Hello!")

	return nil
}

func (SayHello) ShouldFire() bool {
	if time.Now().Minute() == 0 {
		return true
	}
	return false
}
