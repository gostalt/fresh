package services

import (
	"gostalt/app/jobs"

	"github.com/gostalt/framework/schedule"
	"github.com/gostalt/framework/service"
	"github.com/gostalt/logger"
	"github.com/sarulabs/di"
)

type SchedulerServiceProvider struct {
	service.BaseProvider
}

func (p SchedulerServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "scheduler",
		Build: func(c di.Container) (interface{}, error) {
			s := schedule.Runner{}

			sh := c.Get("hello-scheduled").(schedule.Job)
			s.Add(sh)

			return s, nil
		},
	})

	b.Add(di.Def{
		Name: "hello-scheduled",
		Build: func(c di.Container) (interface{}, error) {
			l := c.Get("logger").(logger.Logger)
			h := jobs.MakeSayHello(l)
			return h, nil
		},
	})
}
