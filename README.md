
# Book Management

This project is a Book Management service implemented using gRPC and Golang.

## Prerequisites

- Go 1.16 or higher
- `protoc` Protocol Buffers compiler
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins for Go

## Setup

### Install Protocol Buffers Compiler

You can download and install the Protocol Buffers compiler (`protoc`) from the official [Protocol Buffers releases](https://github.com/protocolbuffers/protobuf/releases) page.

### Install Go Protocol Buffers Plugins

To generate Go code from `.proto` files, you need to install the following plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Compiling Proto File
your path must inside services/"service name"
example:

```sh
protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative protos/book.proto
```

### Create Database From Docker
on root path ./mcs-book

```sh
docker-compose up --build
```

### Running Migration
on inside /migration-svc

#### Migrate
```sh
go run main.go db:migrate
```

#### Rollback
```sh
go run main.go db:rollback
```

### Running Server
with air inside on services/"service name"

```sh
cd services/book-svc
air
```

### Use Postman To Testing This GRPC
Use can read documentation in [Book Service](document/Book-Management%2089bcd24feea945a68cf9f0a1f3462d5a/Book%20Service%20c1afff5c0f874ef7916e1a690ece7328.md)

dont forget to import the proto file in /services/"service name"/protos/
postman-link:

[go-grpc][https://foryou99.postman.co/workspace/Private-Workspace~4b972bf2-8d51-404a-b0f3-0431f98f8f87/collection/66a59d691eb5d2b24269f612?action=share&creator=16466534&active-environment=16466534-560184da-4298-42c0-b923-d254f1cba339]