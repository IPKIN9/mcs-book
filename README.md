
## Compiling Proto File

```go
protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative protos/book.proto
```