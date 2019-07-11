package api

import (
	"encoding/json"
	"fmt"
	"gostalt/config"
	"net/http"
)

// Hello is an example Handler to show the basic idea of how to
// inject the Container into a Handler. All a handler needs to
// be valid is to satisfy the http.Handler interface:
//
//     type http.Handler interface{ServeHTTP(http.ResponseWriter, *http.Request)}
//
type Hello struct{}

// ServeHTTP is called on a Handler when the route is is registered
// against is called.
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	greeting := make(map[string]string)
	greeting["greeting"] = fmt.Sprintf("Hello, %s!", config.Get("main", "app_name"))

	// The variable is marshalled into a slice of bytes, which
	// we can pass to the Write function on the ResponseWriter.
	val, _ := json.Marshal(greeting)
	w.Write(val)
}
