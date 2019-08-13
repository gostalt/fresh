package command

import (
	"database/sql"
	"fmt"
	"gostalt/app"
	"gostalt/config"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gostalt/logger"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

// entityStub is the basic implementation of an entity.
// TODO: The go generate code here isn't great as it relies on shell.
// However, the generate code is called relatively, and invoking the
// gostalt binary with ../../gostalt doesn't pull in env files. Hmm...
const entityStub = `package entity

//go:gen sh -c "cd ../../ && go run main.go migrate magic $GOFILE"

type <Entity> struct {
	// Fields here
}

// Methods here
`

// repositoryStub is the basic implementation of a repository.
const repositoryStub = `package repository

import (
	"gostalt/app/entity"

	"github.com/jmoiron/sqlx"
)

type <Repository> struct {
	*sqlx.DB
}

func (r <Repository>) Fetch(id int) (entity.<Repository>, error) {
	<repository> := entity.<Repository>{}
	err := r.Get(&<repository>, "select * from <repository>s where id = $1 limit 1", id)
	if err != nil {
		r.Logger.Warning([]byte(err.Error()))
		return <repository>, err
	}

	return <repository>, nil
}

func (r <Repository>) FetchAll() []entity.<Repository> {
	<repository>s := []entity.<Repository>{}
	<repository>s, err := r.Select(&<repository>s, "select * from <repository>s")
	if err != nil {
		return []entity.<Repository>{}
	}

	return <repository>s
}
`

const handlerStub = `package <package>

import (
	"net/http"
)

func <handler>(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}`

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Make a new file",
}

var makeHandlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "Make a handler",
	Long: `Use dot notation to create a nested handler. For
example, "api.Welcome" would become api/Welcome.go`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := app.Make()
		handler := args[0]

		createHandler(handler, app)
	},
}

func createHandler(handler string, app *app.App) {
	handlerRoot := "app/http/handler/"
	pieces := strings.Split(handler, ".")

	// TODO: Tidy this shit up
	dir := filepath.Join(pieces[:len(pieces)-1]...)
	dir = handlerRoot + dir
	pkg := pieces[len(pieces)-2]
	file := pieces[len(pieces)-1]
	path := dir + "/" + file + ".go"
	os.MkdirAll(dir, os.ModePerm)

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}

	content := strings.Replace(handlerStub, "<package>", pkg, -1)
	content = strings.Replace(content, "<handler>", file, -1)

	f.Write([]byte(content))

	fmt.Println(path)
}

var makeRepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Make a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := app.Make()

		go createRepository(args[0], app, nil)
	},
}

var makeEntityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Make an entity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		name := args[0]

		app := app.Make()

		cmd.ParseFlags(args)
		migration := cmd.Flag("migration")
		repository := cmd.Flag("repository")

		if migration.Value.String() == "true" {
			wg.Add(1)
			path := fmt.Sprintf("create_%s_table", name)
			go createMigration(path, app, &wg)
		}

		if repository.Value.String() == "true" {
			wg.Add(1)
			go createRepository(name, app, &wg)
		}

		wg.Add(1)
		go createEntity(name, app, &wg)

		wg.Wait()
	},
}

// TODO: The createMigration, createRepository and createEntity
// calls shouldn't really live here, they should hand off to
// some sort of service to do the heavy lifting.

func createMigration(path string, app *app.App, wg *sync.WaitGroup) {
	db := app.Container.Get("database-basic").(*sql.DB)

	if err := goose.Create(
		db,
		config.Get("database", "migration_directory"),
		path,
		"sql",
	); err != nil {
		panic(err)
	}

	wg.Done()
}

func createRepository(name string, app *app.App, wg *sync.WaitGroup) {
	path := config.Get("maker", "repository_path")

	repository := strings.Title(strings.ToLower(name))

	path = path + repository + ".go"

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}

	content := strings.Replace(repositoryStub, "<repository>", strings.ToLower(repository), -1)
	content = strings.Replace(content, "<Repository>", repository, -1)

	f.Write([]byte(content))

	l := app.Container.Get("logger").(logger.Logger)
	l.Info([]byte("repository created"))
	wg.Done()
}

func createEntity(name string, app *app.App, wg *sync.WaitGroup) {
	path := config.Get("maker", "entity_path")

	entity := strings.Title(strings.ToLower(name))

	path = path + entity + ".go"

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}

	content := strings.Replace(entityStub, "<entity>", entity, -1)
	content = strings.Replace(content, "<Entity>", entity, -1)

	// Annoyingly, running `go generate ./...` will trigger the
	// command in the entityStub, so we must replace it here.
	content = strings.Replace(content, "go:gen", "go:generate", -1)

	f.Write([]byte(content))

	l := app.Container.Get("logger").(logger.Logger)
	l.Info([]byte("entity created"))
	wg.Done()
}

func init() {
	makeEntityCmd.Flags().BoolP("migration", "m", false, "Generate a migration file for this entity")
	makeEntityCmd.Flags().BoolP("repository", "r", false, "Generate a repository for this entity")
	makeCmd.AddCommand(makeEntityCmd)
	makeCmd.AddCommand(makeRepositoryCmd)
	makeCmd.AddCommand(makeHandlerCmd)
	rootCmd.AddCommand(makeCmd)
}
