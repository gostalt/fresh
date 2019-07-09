package middleware

import (
	"log"
	"github.com/gostalt/container"
	"net/http"
)

type Token struct {
	Container container.Container
}

func (t *Token) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		at := r.Header.Get("Authorization")

		if at == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// TODO: temporary
		tkns, err := t.Container.Get("TokenRepository")
		if err != nil {
			log.Println("unable to load token repository")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, tkn := range tkns.([]string) {
			if tkn == at {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}
