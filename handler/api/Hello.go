package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sarulabs/di"
)

// Hello is an example Handler to show the basic idea of how to
// inject the Container into a Handler. All a handler needs to
// be valid is to satisfy the http.Handler interface:
//
//     type http.Handler interface{ServeHTTP(http.ResponseWriter, *http.Request)}
//
type Hello struct {
	// The container is a promoted field on this struct, meaning
	// that is can be accessed directly from the struct instead
	// of making verbose calls to Hello.Container.Get.
	Container di.Container
}

// ServeHTTP is called on a Handler when the route is is registered
// against is called.
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Here, we add a header so that the content returned is
	// formatted as JSON.
	w.Header().Set("Content-Type", "application/json")

	// Here, the container is accessed to load an environment
	// variable: APP_NAME.
	greeting := make(map[string]string)
	env := h.Container.Get("env").(map[string]string)
	greeting["greeting"] = fmt.Sprintf("Hello, %s!", env["APP_NAME"])

	// The variable is marshalled into a slice of bytes, which
	// we can pass to the Write function on the ResponseWriter.
	val, _ := json.Marshal(greeting)
	w.Write(val)
}
