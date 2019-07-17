# Gostalt

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

Easily add global middleware to the application, and create 'subrouters'
with their own prefixes and middleware stack using `gorilla/mux`.

- Automatic Let's Encrypt certificate generation.
- SQL Migrations built in.
- Console command extensibility using spf13/cobra.
- Scheduling.