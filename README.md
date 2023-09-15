# grpc_tutorial

# Build protobufs example
> pwd

> /Users/home/go/src/github.com/grpc_tutorial

> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative account/account.proto


API flows i.e.,

* Server
* Middleware
* Routes
* Handlers
  * Validators
    * Types
* Services
  * Clients
  * Types

