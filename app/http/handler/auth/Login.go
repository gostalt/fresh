package auth

import (
	"gostalt/app/entity"
	"gostalt/app/entity/user"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gostalt/validate"
	"github.com/sarulabs/di/v2"
	"golang.org/x/crypto/bcrypt"
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

		client := di.Get(r, "entity-client").(*entity.Client)
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		u, err := client.User.Query().Where(user.UsernameEQ(username)).First(r.Context())
		if err != nil || bcrypt.CompareHashAndPassword(u.Password, []byte(password)) != nil {
			// TODO: Return errors here about invalid password or no user found.
			views.ExecuteTemplate(w, "auth.login", nil)
			return
		}

		err = auth.ProcessLogin(w, r, u)
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
