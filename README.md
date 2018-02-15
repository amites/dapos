# go-dapos

# GRPC Setup
- Install [protoc](https://github.com/google/protobuf/releases) compiler manually or by homebrew `$ brew install protobuf`
- Install `protoc-gen-go plugin`: `go get -u github.com/golang/protobuf/protoc-gen-go`
- Build Go bindings from `.proto` file. `protoc --go_out=plugins=grpc:. grpc/dapos_grpc.proto`