package api

import (
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"greeting": "Hello from Gostalt!"}`))
	return
}
