package middleware

import (
	"context"
	"net/http"

	"github.com/sarulabs/di"
)

type ContainerResolver struct {
	di.Container
}

func (m ContainerResolver) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctn, err := m.Container.SubContainer()
			if err != nil {
				panic(err)
			}
			next.ServeHTTP(w, r.WithContext(
				context.WithValue(r.Context(), di.ContainerKey("di"), ctn),
			))
		},
	)
}
