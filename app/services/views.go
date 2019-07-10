package services

import (
	"gostalt/resources/views"

	"github.com/sarulabs/di"
)

type ViewServiceProvider struct{}

// path is the directory, relative to the project root, that the
// view files will be loaded from. It is walked recursively.
var path = "resources/views"

func (p ViewServiceProvider) Register(b *di.Builder) {
	b.Add(di.Def{
		Name: "views",
		Build: func(c di.Container) (interface{}, error) {
			return views.Load(path), nil
		},
	})
}
