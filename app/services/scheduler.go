package services

import (
	"gostalt/app/cron"

	"github.com/sarulabs/di"
	"github.com/gostalt/logger"
)

type SchedulerServiceProvider struct {
	BaseServiceProvider
}

func (p SchedulerServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "scheduler",
		Build: func(c di.Container) (interface{}, error) {
			l := c.Get("logger").(logger.Logger)
			// Scheduled jobs should just print to STDOUT for now.
			// TODO: Remove this - maybe have a separate log for
			// scheduled jobs?

			s := &cron.Scheduler{
				Logger: l,
			}

			// TODO: Here is where jobs should be added to the scheduler.
			sh := c.Get("hello-scheduled").(cron.Jobber)
			s.Add(sh)

			return s, nil
		},
	})

	b.Add(di.Def{
		Name: "hello-scheduled",
		Build: func(c di.Container) (interface{}, error) {
			l := c.Get("logger").(logger.Logger)
			h := cron.MakeSayHello(l)
			return h, nil
		},
	})
}