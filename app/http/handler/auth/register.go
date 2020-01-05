package auth

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gostalt/validate"
	"github.com/sarulabs/di/v2"
)

func Register(w http.ResponseWriter, r *http.Request) {
	views := di.Get(r, "views").(*template.Template)

	if r.Method == http.MethodGet {
		views.ExecuteTemplate(w, "auth.register", nil)
	}

	if r.Method == http.MethodPost {
		// Create a new auth.Provider to handle the logic for
		// the registration attempt.
		store := di.Get(r, "session").(*sessions.CookieStore)
		auth := NewProvider(store)

		// First, check that the validation rules are satisfied by
		// the request. If not, redirect to the register page
		// with appropriate error messages for the attempt.
		msgs, err := validate.Check(r, registerRules()...)
		if err != nil || len(msgs) > 0 {
			views.ExecuteTemplate(w, "auth.register", getErrorsFromMessage(msgs))
			return
		}

		// Create the user ...
		user, err := auth.CreateUser(r)
		if err != nil {
			views.ExecuteTemplate(w, "auth.register", []string{err.Error()})
			return
		}

		// If the attempt is successful, redirect the user to
		// the appropriate location.
		err = auth.ProcessLogin(w, r, user)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, auth.DefaultRedirect(), http.StatusSeeOther)
	}
}

func registerRules() []validate.Rule {
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
			Param: "username",
			Check: validate.Alphanumeric,
		},
		validate.Rule{
			Param: "password",
			Check: validate.Required,
		},
		validate.Rule{
			Param: "password",
			Check: validate.Empty,
		},
		validate.Rule{
			Param:   "password",
			Check:   validate.MinLength,
			Options: validate.Options{"length": 8},
		},
	}
}
