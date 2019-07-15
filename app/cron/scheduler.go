package cron

import (
	jww "github.com/spf13/jwalterweatherman"
)

type Scheduler struct {
	Logger *jww.Notepad
	jobs   []Jobber
}

func (s *Scheduler) Run() {
	s.Logger.INFO.Println("running scheduler")

	for i, j := range s.jobs {
		if j.ShouldFire() {
			s.Logger.INFO.Printf("Running job %d", i)
			j.Handle()
		}
	}
}

func (s *Scheduler) Add(jobs ...Jobber) {
	s.jobs = append(s.jobs, jobs...)
}
