package services

import (
	"gostalt/app/cron"
	"os"

	"github.com/sarulabs/di"
	jww "github.com/spf13/jwalterweatherman"
)

type SchedulerServiceProvider struct {
	BaseServiceProvider
}

func (p SchedulerServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "scheduler",
		Build: func(c di.Container) (interface{}, error) {
			logger := c.Get("logger").(*jww.Notepad)

			// Scheduled jobs should just print to STDOUT for now.
			// TODO: Remove this - maybe have a separate log for
			// scheduled jobs?
			logger.SetLogOutput(os.Stdout)

			s := &cron.Scheduler{
				Logger:    logger,
				Container: c,
			}

			// Should loop through the jobs and add them here.
			s.Add(cron.SayHello{})

			return s, nil
		},
	})
}
