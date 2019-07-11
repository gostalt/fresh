package middleware

import (
	"database/sql"
	"log"
	"net/http"
)

type TokenAuthentication struct {
	DB *sql.DB
}

func (mw TokenAuthentication) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			stmt, err := mw.DB.Prepare("SELECT id FROM tokens WHERE token = $1")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				log.Println(err)
				return
			}
			usr := stmt.QueryRow(token)
			var id int
			if err := usr.Scan(&id); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				log.Println(err)
				return
			}

			if id == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				log.Println(err)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
