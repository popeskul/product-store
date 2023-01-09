# Product store 🚀
This project is a simple product store. It's a simple gRPC server and client.
The server is a gRPC server and the client is a gRPC client.

### Prerequisites
- [Go](https://golang.org/dl/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/downloads)
- [gRPC](https://grpc.io/docs/quickstart/go/)
- [Docker](https://docs.docker.com/get-docker/)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

## Usage

### Generate docker
```bash
docker run --rm -d --name product-store -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=123 -p 27017:27017 mongo:latest
```

### Generate protobuf files
```bash
protoc --go_out=. --go-grpc_out=. proto/products.proto
```

### Run the server
```bash
go run server/main.go
```

### Run the client
```bash
go run client/main.go
```

You can also find more examples in the cmd/client/main.go file.
