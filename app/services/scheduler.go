package services

import (
	"fmt"
	"gostalt/app/jobs"

	"github.com/gostalt/framework/schedule"
	"github.com/gostalt/framework/service"
	"github.com/gostalt/logger"
	"github.com/sarulabs/di/v2"
)

type SchedulerServiceProvider struct {
	service.BaseProvider
}

func (p SchedulerServiceProvider) jobs(c di.Container) []schedule.Job {
	// To add more scheduled jobs to your application, add them
	// to this array.
	return []schedule.Job{
		// The `Quote` job demonstrates how easy it is to resolve
		// a job from the container, allowing you to pass in
		// dependencies.
		c.Get("job-quote").(schedule.Job),
	}
}

func (p SchedulerServiceProvider) Register(b *di.Builder) error {
	err := b.Add(di.Def{
		Name: "scheduler",
		Build: func(c di.Container) (interface{}, error) {
			s := schedule.NewRunner(p.jobs(c)...)

			return s, nil
		},
	})

	if err != nil {
		return fmt.Errorf("unable to register schedule service: %w", err)
	}

	// This is just an example, you're free to delete this at
	// job at any time.
	err = b.Add(di.Def{
		Name: "job-quote",
		Build: func(c di.Container) (interface{}, error) {
			logger := c.Get("logger").(logger.Logger)

			return jobs.Quote{Logger: logger}, nil
		},
	})

	if err != nil {
		return fmt.Errorf("unable to register schedule service: %w", err)
	}

	return nil
}
