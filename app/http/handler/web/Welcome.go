package web

import (
	"html/template"
	"net/http"

	"github.com/sarulabs/di"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	views := di.Get(r, "views").(*template.Template)
	r.ParseForm()
	name := r.Form.Get(":name")

	views.ExecuteTemplate(w, "welcome", name)
}
