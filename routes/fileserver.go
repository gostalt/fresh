package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// fileServer creates an http.FileServer at `path` that loads
// files from the `path` directory from the project root.
func fileServer(r *mux.Router, path string, dir string) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			path,
			http.FileServer(http.Dir(dir)),
		),
	)
}
