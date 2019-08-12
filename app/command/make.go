package command

import (
	"fmt"
	"gostalt/app"
	"gostalt/config"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// entityStub is the basic implementation of an entity.
const entityStub = `package entity

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

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Make a new file",
}

var makeRepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Make a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.Make()
		path := config.Get("maker", "repository_path")

		repository := strings.Title(strings.ToLower(args[0]))

		path = path + repository + ".go"

		f, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}

		content := strings.Replace(repositoryStub, "<repository>", strings.ToLower(repository), -1)
		content = strings.Replace(content, "<Repository>", repository, -1)

		f.Write([]byte(content))
	},
}

var makeEntityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Make an entity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.Make()
		path := config.Get("maker", "entity_path")

		entity := strings.Title(strings.ToLower(args[0]))

		path = path + entity + ".go"

		f, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}

		content := strings.Replace(entityStub, "<Entity>", entity, -1)

		f.Write([]byte(content))
	},
}

func init() {
	makeCmd.AddCommand(makeEntityCmd)
	makeCmd.AddCommand(makeRepositoryCmd)
	rootCmd.AddCommand(makeCmd)
}
