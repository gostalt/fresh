package api

import (
	"encoding/json"
	"fmt"
	"gostalt/config"
	"net/http"
)

// Hello is a super basic example of using a http.HandlerFunc as
// a route handler. You can see it being added to APIRoutes in
// the `routes/api.go` file.
func Hello(w http.ResponseWriter, r *http.Request) {
	greeting := make(map[string]string)
	greeting["greeting"] = fmt.Sprintf("Hello from %s!", config.Get("app", "name"))

	val, _ := json.Marshal(greeting)
	w.Write(val)
}
