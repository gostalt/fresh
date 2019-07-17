module gostalt

go 1.12

require (
	github.com/gorilla/mux v1.7.3

	github.com/joho/godotenv v1.3.0
	github.com/kabukky/httpscerts v0.0.0-20150320125433-617593d7dcb3
	github.com/lib/pq v1.1.1
	github.com/pkg/errors v0.8.1 // indirect
	github.com/pressly/goose v2.6.0+incompatible
	github.com/sarulabs/di v2.0.0+incompatible
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/tmus/logger v0.0.0
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9
)

replace github.com/tmus/logger => ../logger-interface
