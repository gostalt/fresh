package middleware

import (
	"log"
	"net/http"
)

func Varsity(next http.Handler) http.Handler {
	log.Println("Varsity happening")
	fn := func(w http.ResponseWriter, r *http.Request) {
		// vs := mux.Vars(r)

		r.ParseForm()
		formVals := r.Form
		formVals.Add("gorilla", "hello")
		formVals.Add("gorilla", "world")

		r.Form = formVals

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
