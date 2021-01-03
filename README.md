# protoc-gen-gin

generate gin http server for grpc.

## Install
```shell
go install github.com/312362115/protoc-gen-gin
```

## Usage
```shell
protoc --gin_out=:. path_to_your_proto
```
check [Makefile](example/Makefile) for more usage.
