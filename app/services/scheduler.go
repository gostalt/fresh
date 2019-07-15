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
			l := c.Get("logger").(*jww.Notepad)
			// Scheduled jobs should just print to STDOUT for now.
			// TODO: Remove this - maybe have a separate log for
			// scheduled jobs?
			l.SetLogOutput(os.Stdout)

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
			l := c.Get("logger").(*jww.Notepad)
			h := cron.MakeSayHello(l)
			return h, nil
		},
	})
}
