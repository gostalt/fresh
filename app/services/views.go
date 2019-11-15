package services

import (
	"gostalt/config"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gostalt/framework/service"
	"github.com/sarulabs/di/v2"
)

type ViewServiceProvider struct {
	service.BaseProvider
}

// path is the directory, relative to the project root, that the
// view files will be loaded from. It is walked recursively.
var path = "resources/views"

func (p ViewServiceProvider) Register(b *di.Builder) {
	shared := config.Get("app", "environment") != "production"

	b.Add(di.Def{
		Name: "views",
		Build: func(c di.Container) (interface{}, error) {
			return load(path), nil
		},

		// If the app's environment is not production, then set this
		// definition to `Unshare`. This means that a new instance
		// will be created per request, which is inefficient,
		// but useful for testing.
		Unshared: shared,
	})
}

// load walks through the directory provided and loads all the
// `.html` files.
func load(path string) *template.Template {
	path = filepath.Clean(path)

	tmpls, err := findAndParseTemplates(path, viewFunctions())
	if err != nil {
		log.Fatalln("unable to load templates:", err)
	}

	return tmpls
}

func viewFunctions() template.FuncMap {
	return template.FuncMap{
		"asset": func(path string) string {
			return "/assets/" + path
		},
	}
}

func findAndParseTemplates(
	path string,
	funcMap template.FuncMap,
) (*template.Template, error) {
	pfx := len(path) + 1
	root := template.New("")

	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, e1 error) error {
			if !info.IsDir() && strings.HasSuffix(path, ".html") {
				if e1 != nil {
					return e1
				}

				b, e2 := ioutil.ReadFile(path)
				if e2 != nil {
					return e2
				}

				// Strip the `.html` string from the end of the
				// template so we can execute it using `name`
				// rather than `name.html`.
				name := path[pfx : len(path)-5]

				name = strings.Join(
					strings.Split(name, "/"),
					".",
				)

				t := root.New(name).Funcs(funcMap)
				t, e2 = t.Parse(string(b))
				if e2 != nil {
					return e2
				}
			}

			return nil
		},
	)

	return root, err
}
