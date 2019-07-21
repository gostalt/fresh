package api

import (
	"fmt"
	"net/http"

	"github.com/gostalt/validate"
)

type PostTest struct{}

func (h PostTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if msgs, err := validate.Check(r, h.rules()...); err != nil {
		validate.Respond(w, msgs)
		return
	}

	w.Write([]byte("Hello"))
}

func (h PostTest) rules() []validate.Rule {
	return []validate.Rule{
		validate.Rule{
			Param: "forename",
			Check: func(r *http.Request, param string) error {
				if _, exists := r.Form[param]; !exists {
					return fmt.Errorf("The %s param must exist", param)
				}

				return nil
			},
		},
		validate.Rule{
			Param: "doohickey",
			Check: func(r *http.Request, param string) error {
				if _, exists := r.Form[param]; !exists {
					return fmt.Errorf("The %s param must exist", param)
				}

				return nil
			},
		},
	}
}
