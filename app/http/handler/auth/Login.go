package auth

import (
	"gostalt/app/entity"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gostalt/validate"
	"github.com/sarulabs/di/v2"
)

func Login(w http.ResponseWriter, r *http.Request) {
	views := di.Get(r, "views").(*template.Template)

	if r.Method == http.MethodGet {
		views.ExecuteTemplate(w, "auth.login", nil)
	}

	if r.Method == http.MethodPost {
		msgs, err := validate.Check(r, loginRules()...)
		if err != nil || len(msgs) > 0 {
			views.ExecuteTemplate(w, "auth.login", struct{ Failures validate.Message }{msgs})
			return
		}

		store := di.Get(r, "session").(*sessions.CookieStore)
		auth := NewProvider(store)

		err = auth.ProcessLogin(w, r, entity.User{Username: "Tomy"})
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, auth.DefaultRedirect(), http.StatusSeeOther)
	}
}

func loginRules() []validate.Rule {
	return []validate.Rule{
		validate.Rule{
			Param: "username",
			Check: validate.Required,
		},
		validate.Rule{
			Param: "username",
			Check: validate.Empty,
		},
		validate.Rule{
			Param: "password",
			Check: validate.Required,
		},
		validate.Rule{
			Param: "password",
			Check: validate.Empty,
		},
	}
}
