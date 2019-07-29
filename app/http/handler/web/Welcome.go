package web

import (
	"html/template"
	"net/http"

	"github.com/sarulabs/di"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	views := di.Get(r, "views").(*template.Template)

	views.ExecuteTemplate(w, "welcome", nil)
}
