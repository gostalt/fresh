package auth

import (
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
		// Create a new auth.Provider to handle the logic for
		// the login attempt.
		store := di.Get(r, "session").(*sessions.CookieStore)
		auth := NewProvider(store)

		// First, check that the login rules are satisfied by the
		// request. If not, redirect back to the login page with
		// appropriate error messages for the login attempt.
		msgs, err := validate.Check(r, loginRules()...)
		if err != nil || len(msgs) > 0 {
			views.ExecuteTemplate(w, "auth.login", getErrorsFromMessage(msgs))
			return
		}

		// After, try to retrieve the user from the database. If
		// the user does not exist or the password is incorrect,
		// then the user is redirected back to the login page
		// with an error indicating why the login failed.
		user, err := auth.RetrieveUser(r)
		if err != nil {
			views.ExecuteTemplate(w, "auth.login", []string{err.Error()})
			return
		}

		// If the login is successful, process the user's login
		// and redirect them to the appropriate location.
		err = auth.ProcessLogin(w, r, user)
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
