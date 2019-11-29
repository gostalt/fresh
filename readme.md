# Gostalt

The Go Framework.

[![Code Report](https://goreportcard.com/badge/github.com/gostalt/gostalt)](https://goreportcard.com/report/github.com/gostalt/gostalt)

> Gostalt is a clean, minimal framework for the Go programming language.
> It aims to offer just enough, whilst also eliminating a number
> of pain points when creating new Go projects.

> ⚠️ **Gostalt is nowhere near a v1.0.0 release, but you're welcome
> to take a look around, pinch parts that interest you or attempt
> to use it to create an app.**

Feature highlights include:

-   Easy route registration
-   App-wide and route-group middleware
-   DI Container
-   Config resolution using `.env` files
-   Database migrations
-   Commands and scheduled jobs
-   Automatic SSL certificate generation

## Installation

### Requirements

To compile the finished binary, the latest version of Go is
recommended.

To run the binary, you'll just need a web server with root access.

### Installation

Clone the `gostalt/gostalt` repo to your development environment.
In a terminal, navigate to the cloned directory and run `go build`,
then `./gostalt serve`. You should see a message from Gostalt
that it is running:

```bash
$ ./gostalt serve
  Info: Server running at https://localhost:8080
```

That's it — you've downloaded and build the Gostalt binary! Visit
the server's address and you should see the Gostalt splash screen.

## Routing

Gostalt uses an expressive syntax to make registering new routes
for your application easy. Two sets of route groups are set up
for you with a fresh install of Gostalt: `api` and `web`. Both
are stored in the `routes` folder.

Creating a new route is as easy as adding a new line to the
`route.Collection` function call:

```go
var Web = route.Collection(
    route.Get("/", web.HomeHandler{}),
    route.Get("about", http.Handler(http.HandlerFunc(web.AboutHandler))),
    // Your routes ...
)
```

Of course, you can use `route.Get`, `route.Post`, `route.Put`,
`route.Patch` and `route.Delete` to register routes.

The second parameter of these route registration methods should
be an `http.Handler`. To use a function, you should wrap it in as
an `http.Handler`, as seen in the `about` route above.

You can use `{param}` notation to add parameters to a route:

```go
route.Get("greet/{name}", web.GreetHandler{})
```

### Redirect Routes

Rather than create an entire Handler to manage redirecting from
one page to another, Gostalt includes a `Redirect` method to
make this process a cinch:

```go
route.Redirect("/old", "/new")
```

### Middleware

You can add Middleware to route groups by calling the `Middleware`
method on a route Collection:

```go
var Web = route.Collection(
	// ...
).Middleware(mw.AddURIParametersToRequest)
```

You'll see that the `web` and `api` routes already have some
sensible default Middleware assigned to them.

#### Writing Middleware

Middleware should be of type `route.Middleware`:

```go
type Middleware func(http.Handler) http.Handler
```

A middleware accepts an `http.Handler` and returns an `http.Handler`.
It uses this setup to run through each "layer" of middleware for
each request, passing the successful Handler to the "next" handler.

## Handlers

Gostalt uses Go's excellent `net/http` package Handlers. You should
add new Handlers to the `app/http/handler` folder, but you are
free to create them anywhere in your app.

### Accessing the Container

To access items from the container inside a Handler, you can use
the `di.Get(*http.Request, string)` method.

> Behind the scenes, Gostalt uses the `sarulabs/di` package. The
> `di.Get` package adds the container to the request context,
> allowing you to resolve items inside a Handler.

The `web.Welcome` handler (at `app/http/handler/web/Welcome.go`)
uses this to load the "welcome" view:

```go
func Welcome(w http.ResponseWriter, r *http.Request) {
	views := di.Get(r, "views").(*template.Template)

	views.ExecuteTemplate(w, "welcome", nil)
}
```

### Retrieving Items from the Request

You can access items from the Request as you would any other Go
project. Usually in Go, you'd need to call `r.ParseForm()`. This
is handled automatically, meaning you can access form values using
`r.Form.Get()`.

To access URI parameters, you should use `r.Form.Get()`. Note that,
to prevent collisions with form values, URI parameters are prepended
with a `:`. For example:

```go
// at /posts/{slug}

slug := r.Form.Get(":slug")
```

## Views and Writing Responses

Views use the built-in Go `html/template` package. View files are
`.html` files and should be stored in the `resources/views`
directory. You can nest folders in this directory as deep as you'd
like.

### Executing Views

You can execute a view inside a Handler or HandlerFunc by
resolving the views from the Container and calling the
`ExecuteTemplate` method on the view:

```go
func Welcome(w http.ResponseWriter, r *http.Request) {
	views := di.Get(r, "views").(*template.Template)

	views.ExecuteTemplate(w, "welcome", nil)
}
```

When accessing a view, you should omit the `.html` extension.
Nested views use dot-notation rather than directory separators:

```go
    // admin.users.index would load the view file from
    // resources/views/admin/users/index.html
    views.ExecuteTemplate(w, "admin.users.index", nil)
```

To pass data to a view, you should first assemble all the data
you intend to pass to the view into a single item. Often, this
will be a response from a Repository.

If you need to assemble a variety of data, you should use an
"anonymous struct" to define it:

```go
func ShowTodos(w http.ResponseWriter, r *http.Request) {
    views := di.Get(r, "views").(*template.Template)

    data := struct{
        Name string
        Age  int
        Todos []Todo
    }{
        Name: "Tomy Smith",
        Age:  27,
        Todos: todoRepository.FetchAll(),
    }

	views.ExecuteTemplate(w, "user.account", data)
}
```

### Responses

Because Gostalt uses Go's `net/http` package, writing responses
is exactly the same as you'd expect. You can write additional
headers to the response by using the `Header` method on the
`ResponseWriter`:

```go
func HeaderExample(w http.ResponseWriter, r *http.Request) {
    // Add a Header key and value
    w.Header().Set("Content-Type", "application/json")

    // Set an HTTP Status Code
    w.WriteHeader(http.StatusTeapot)

	views.ExecuteTemplate(w, "welcome", nil)
}
```

Because of the way the ResponseWriter works, you can write to it
at any point in the middleware stack. This is used in the `api`
routes to automatically add a JSON header to all API responses:

```go
// routes/api.go
var API = route.Collection(
	route.Get("welcome", http.Handler(http.HandlerFunc(api.Hello))),
).Prefix("api").Middleware(mw.AddJSONContentTypeHeader)

// AddJSONContentTypeHeader.go
func AddJSONContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		},
	)
}
```

## The Environment and Config

Modern web development dictates that config should be stored in
the environment: configuration will vary between deployments,
and indeed even between developer environments.

Gostalt enables each environment to maintain unique credentials,
whilst still allowing common code to be written.

During the app's startup, the contents of the `.env` file in the
root of the project is read, and is passed to config, which defines
many different domains and values. You can view these in the `config`
folder in the root of the project.

You can use `config.Get` to retrieve the value of any set key
at any point in your application:

```go
name := config.Get("app", "name") // Gostalt
```

You're free to create your own config files to use in your app.
You should pass in the `env` variable if you wish to load the
config from the `.env` file, as shown below:

```go
// config/config.go
cfg = map[string]map[string]string{
    "logging":  logging(env),
    // your config here...
}
```

## Container

The container makes it easy to handle dependencies and utilise
dependency injection on custom types, rather than hard-coding
the dependencies into the type.

### Binding to the Container

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

As you can see, a new service needs a `Name` and `Build` field.
The build field's value should be a function that accepts the
container as a parameter and returns an interface and an error.
In this example, we are resolving the "database" item out of the
container and using it as a value in the initialisation of a new
`respository.User` struct. This means that, should our database
details change, we don't need to change them in every area of our
codebase.

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
`BaseServiceProvider` field. This implements stub methods for
`Register` and `Boot`, meaning that we can omit them from our
Provider whilst still implementing the `Provider` interface.
Of course, you can override the `Register` and `Boot` methods
of the `BaseServiceProvider` by defining them on your new provider.

> The naming is obvious, but you should only _Register_ items
> into the container in the `Register` method. If you attempt
> anything else in the Register method, the service may not
> yet have been created.

## Logging

Gostalt uses a logging interface that allows you to swap logger
implementations easily, without needing to change any code.

The `logger.Logger` interface implements IETF's RFC 5254 to define
eight levels of logging: Emergency, Alert, Critical, Error,
Warning, Notice, Info and Debug.

Gostalt ships with two log drivers: `stdout` and `file`. The File
log driver writes all calls to a single file in the `storage/logs`
directory. By default, any calls to the logger are written to
stdout, but you're welcome to change this by updating the "driver"
key in `config/logging.go`.

To use the logger, you can add a dependency to any items that you
add to the DI Container. To log messages inside handlers, you can
resolve the logger from the container using the `di.Get` method:

```go
l := di.Get(r, "logging").(logger.Logger)
l.Alert([]byte("Something bad happened."))
```

## Database Migrations

Gostalt makes it easy to create and run migrations and manage the
state of your database.

To create a new migration, run `gostalt migrate create <name>`.
This will generate a migration in the `database/migrations/`
directory. You should fill this migration with the desired `up`
and `down` commands for the migration, i.e., a creation and a
reversal.

> Currently, Gostalt only supports `sql` migrations, that must be
> written in pure SQL.

To run all pending migrations, run `gostalt migrate up`. This
will gather the migrations that have not been executed and run
them against the database defined in `/config/database.go`.

To reverse a migration, you can use `gostalt migrate down`.

## Commands

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

With Commands, you have the full power of the `spf13/cobra`
library combined with the Gostalt DI Container.

## Scheduling

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

## Automatic SSL

When deploying your app to production, Gostalt will automatically
provision a Let's Encrypt certificate. Ensure that ports `443`
and `80` can be accessed.
