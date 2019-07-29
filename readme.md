# Gostalt

The Go Framework.

-----

Gostalt is a clean, minimal framework for the Go programming language.
It aims to offer “just enough”, whilst also eliminating a number
of pain points when creating new Go projects.

Feature highlights include:

- Easy route registration
- App-wide and route-group middleware
- DI Container
- Config resolution using `.env` files
- Database migrations
- Commands and scheduled jobs
- Automatic SSL certificate generation

## Features

### Container

The container makes it easy to handle dependencies and utilise
dependency injection on custom types, rather than hard-coding
the dependencies into the type.

#### Binding to the Container

Gostalt uses `sarulabs/di` as a service container. To bind a new
service, add a new `di.Def{}` to the `services` slice in
`app/services/app.go`:

```go
var services = []di.Def{
    Name: "UserRepository",
    Build: func(c di.Container) (interface{}, error) {
        db := c.Get("database").(*sql.DB)
        return &repository.User{
            DB: db,
        }, nil
    }
}
```

As you can see, a new services needs a `Name` and `Build` field.
The build field's value should be a function that accepts the
container as a parameter and returns an interface and an error.

> Note that, to retrieve an item from the container, you should
> use the `Get` method to pass in the name of an item. Because
> this returns an `interface{}`, it needs to be cast to the
> appropriate type—in the example above, we use `*sql.DB`.

For more complex services, it makes sense to create a dedicated
Service Provider, for example the `LoggingServiceProvider` or
the `TLSServiceProvider`. Service providers have two methods,
`Register` and `Boot`. The Register method is called on each
provider to create the container, and then the boot method is
called on each service provider.

You can create a service provider by declaring a new type:

```go
type ExampleServiceProvider struct {
    BaseServiceProvider
}
```

As you can see above, the new service provider has a promoted
`BaseServiceProvider` field. This implements the `Boot` method
for us, so we only need to define a `Register` method on the new
service provider. Of course, you can override the `Boot` method
of the `BaseServiceProvider` by defining one on your new provider.

> The naming is obvious, but you should only *Register* items
> into the container in the `Register` method. If you attempt
> anything else in the Register method, the service may not
> yet have been created.

### Granular Routing and Middleware

In the `./routes` directory, you can define routes for your app.
By default, there are two subrouters: `api` and `web`, but you
can easily create your own.

The subrouters can resolve items out of the container, and add
new routes to the router (behind the scenes, routing uses the
`gorilla/mux` package, so anything that works there will work
here). Basic routes can be defined using a `http.HandlerFunc`:

```go
s.Methods("GET").Path("/").HandlerFunc(
    func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello world!"))
    }
)
```

For more complex routes, you should create a new `http.Handler`
instead.

You can also add middleware to a set of routes. For example, in
the `api.go` file:

```go
s.Use(
    middleware.JSONHeader,
)
```

Because Gostalt uses `gorilla/mux`, you can use parameters in
your paths:

```go
s.Methods("GET").Path("/hello/{name}").Handler(...)
```

To retrieve these parameters, you should use `mux.Vars(r)` on
the request inside a handler. However, because this is a common
operation, you can optionally add the `AddURIParametersToRequest`
middleware to your router (this is added to the `web` router by
default). This middleware parses the URI params and adds them
to the Request's `Form` field.

This adds a header of `Content-Type: application/json` to every
route in the api subrouter.

### Migrations

Gostalt uses the `pressly/goose` library for creating, running
and managing migrations.

To create a new migration, run `gostalt migrate create <name>`.
This will generate a migration in the `/database/migrations/`
directory. You should fill this migration with the desired `up`
and `down` commands for the migration, i.e., a creation and a
reversal.

To run all pending migrations, run `gostalt migrate up`. This
will gather the migrations that have not been executed and run
them against the database defined in `/config/database.go`.

### Commands

When you're ready to run your app, run `gostalt serve`. You'll
be able to visit the address defined in `/config/app.go`.

`serve`, along with `migrate`, is a cobra command. You are free
to register additional commands by creating a command and adding
it to the `rootCmd`:

```go
// app/command/greet.go

package command

import (
	"gostalt/app"

	"github.com/spf13/cobra"
)

// serveCmd builds our app and runs it.
var greet = &cobra.Command{
	Use:   "greet",
	Short: "Greet the user",
	Run: func(cmd *cobra.Command, args []string) {
        // If you need to access config or resolve an item from
        // the DI Container, you can Make() the app here.
        a := app.Make()
        name := config.Get("app", "name")

        fmt.Printf("Hello from %s!\n", name)
        // Hello from Gostalt!
	},
}

func init() {
	rootCmd.AddCommand(greet)
}
```

### Scheduling

Gostalt aims to make scheduled jobs easier: instead of creating
loads of cron jobs on your environment, you can instead create one:

```bash
* * * * * ./gostalt schedule
```

Gostalt has a component called `schedule`, which you can use to
register commands. Commands should have two methods: `Handle()`
and `ShouldFire()`.

The `Handle()` method is responsible for running the actual job; you
can use the full power of the DI Container to inject the database
and any other requirements.

The `ShouldFire()` method determines if the job should be ran.
You can use Go's `time` package to create easy-to-understand
conditions:

```go
func ShouldFire() bool {
	if time.Now().Minute() == 0 {
		return true
    }

	return false
}
```

But, because this function returns a boolean, you can more powerfully
determine if a job should execute or not: is the environment
`staging`? If so, don't run the job. Is there only 5 records in
the database? Don't run the job until there are 10. As long as
the conditions can be boiled down to a true or false value, they
can be used to determine if a job should fire.

#### To write:

- Automatic Let's Encrypt certificate generation.