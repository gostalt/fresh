package api

import (
	"fmt"
	"gostalt/app/entity"
	"net/http"

	"github.com/gostalt/validate"
	"github.com/sarulabs/di/v2"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	client := di.Get(r, "entity-client").(*entity.Client)
	name := r.Form.Get(":name")

	msgs, err := validate.Check(r, validate.Rule{Param: ":name", Check: validate.MinLength, Options: validate.Options{"length": 6}})
	if err != nil {
		validate.Respond(w, msgs)
		return
	}

	u, err := client.User.
		Create().
		SetName(name).
		Save(r.Context())
	if err != nil {
		w.Write([]byte(`{"error": "user could not be created"}`))
		return
	}

	w.Write([]byte(fmt.Sprintf(`{"greeting": "Hello, %s! From Gostalt."}`, u.Name)))
	return
}
