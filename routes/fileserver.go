package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// fileServer creates an http.FileServer at `/assets/*` that
// loads files from the `./assets` directory.
//
// TODO: There should be the ability to set this dynamically.
// Or maybe just leave this here and people can change it.
func fileServer(r *mux.Router) {
	// `/assets/` is the URI path that files will be served from.
	r.PathPrefix("/assets/").Handler(
		// The prefix needs stripping, because otherwise Go
		// would look for files in `/assets/assets/*`.
		http.StripPrefix(
			"/assets/",
			// `./assets` is the local directory to serve assets
			// out of. It doesn't need to be the same as above.
			http.FileServer(http.Dir("./assets")),
		),
	)
}
