package cron

import (
	"time"

	"github.com/sarulabs/di"
	jww "github.com/spf13/jwalterweatherman"
)

type Jobber interface {
	Handle(di.Container) error
	ShouldFire(di.Container) bool
}

type SayHello struct{}

func (SayHello) Handle(c di.Container) error {
	l := c.Get("logger").(*jww.Notepad)

	l.INFO.Println("Hello!")

	return nil
}

func (SayHello) ShouldFire(c di.Container) bool {
	if time.Now().Minute() == 0 {
		return true
	}
	return false
}
