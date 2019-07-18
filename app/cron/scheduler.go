package cron

import (
	"fmt"

	"github.com/gostalt/logger"
)

type Scheduler struct {
	Logger logger.Logger
	jobs   []Jobber
}

func (s *Scheduler) Run() {
	s.Logger.Info([]byte("Running scheduler"))

	for i, j := range s.jobs {
		if j.ShouldFire() {
			p := fmt.Sprintf("Running job %d", i)
			s.Logger.Info([]byte(p))
			j.Handle()
		}
	}
}

func (s *Scheduler) Add(jobs ...Jobber) {
	s.jobs = append(s.jobs, jobs...)
}
