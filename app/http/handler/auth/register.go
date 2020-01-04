package auth

import (
	"gostalt/app/entity"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gostalt/validate"
	"github.com/sarulabs/di/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	views := di.Get(r, "views").(*template.Template)

	if r.Method == http.MethodGet {
		views.ExecuteTemplate(w, "auth.register", nil)
	}

	if r.Method == http.MethodPost {
		msgs, err := validate.Check(r, registerRules()...)
		if err != nil || len(msgs) > 0 {
			views.ExecuteTemplate(w, "auth.register", struct{ Failures validate.Message }{msgs})
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		client := di.Get(r, "entity-client").(*entity.Client)
		// crypt := di.Get(r, "crypto").(crypto.Crypto)
		// pw, _ := crypt.Encrypt([]byte(password))
		// pwBytes := base64.RawURLEncoding.EncodeToString(pw)
		encrypted, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		u, err := client.User.Create().SetUsername(username).SetPassword(encrypted).Save(r.Context())
		if u == nil || err != nil {
			log.Println("User", u)
			log.Println("Error", err)
		}

		store := di.Get(r, "session").(*sessions.CookieStore)
		auth := NewProvider(store)
		err = auth.ProcessLogin(w, r, u)
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
