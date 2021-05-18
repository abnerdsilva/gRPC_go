# gRPC_go


# Generate pb
    $ protoc --go_out=. --go-grpc_out=. proto/*.proto

# Run gRPC Server
    $ go run cmd/server/main.go

# Run gRPC Client
    $ go run cmd/client/main.go