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

	"github.com/gostalt/framework/service/maker"
	"github.com/gostalt/logger"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
)

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
	m := app.Container.Get("RepositoryMaker").(maker.RepositoryMaker)
	m.Make(name)

	l := app.Container.Get("logger").(logger.Logger)
	l.Info([]byte("repository created"))
	wg.Done()
}

func createEntity(name string, app *app.App, wg *sync.WaitGroup) {
	m := app.Container.Get("EntityMaker").(maker.EntityMaker)
	m.Make(name)

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
