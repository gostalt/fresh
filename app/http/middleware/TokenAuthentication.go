package middleware

import "net/http"

type TokenAuthentication struct {
	Valid []string
}

func (mw TokenAuthentication) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			for _, v := range mw.Valid {
				if token == v {
					next.ServeHTTP(w, r)
				}
			}

			w.WriteHeader(http.StatusUnauthorized)
		},
	)
}
