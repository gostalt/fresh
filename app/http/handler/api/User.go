package api

import (
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}