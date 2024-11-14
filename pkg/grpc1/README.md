# Basic gRPC
This package is a basic gRPC implementation. For a project I am working on at the moment, I actually want it to work in the reverse way. I want a pluggable system, where the main process (the server), calls a process on the stub (the client). This is not possible in the architecture that is shown here.

An option is to instead configure my main process to be the client, and each client would then be a gRPC server. This would have some overheads and some additional complications that I need to investigate.

## codegen
To generate the code, update the protobuf file inside `./proto/sample`, then run the codegen with `protoc`:
```sh
cd pkg/grpc1/proto/sample
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative sample.proto
```
