# julian protos

# Pre-requisites
```sh
# Installs protobuf (which includes protoc)
brew install protobuf

# Installs Go generator
brew install protoc-gen-go
brew install protoc-gen-go-grpc

# Installs TS generator
npm install -g protoc-gen-ts
npm install -g protoc-gen-js
```

# Building for Go (gRPC rest demo)

```sh
cd grpc-rest-demo
make build_go
```
