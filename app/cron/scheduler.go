package cron

import (
	"github.com/sarulabs/di"
	jww "github.com/spf13/jwalterweatherman"
)

type Scheduler struct {
	Logger    *jww.Notepad
	jobs      []Jobber
	Container di.Container
}

func (s *Scheduler) Run() {
	s.Logger.INFO.Println("running scheduler")

	for i, j := range s.jobs {
		if j.ShouldFire(s.Container) {
			s.Logger.INFO.Printf("Running job %d", i)
			j.Handle(s.Container)
		}
	}
}

func (s *Scheduler) Add(jobs ...Jobber) {
	s.jobs = append(s.jobs, jobs...)
}
