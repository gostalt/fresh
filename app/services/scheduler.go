package services

import (
	"github.com/gostalt/framework/schedule"
	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"
)

type SchedulerServiceProvider struct {
	service.BaseProvider
}

func (p SchedulerServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "scheduler",
		Build: func(c di.Container) (interface{}, error) {
			s := schedule.Runner{}

			// Add jobs here.
			s.Add()

			return s, nil
		},
	})
}
