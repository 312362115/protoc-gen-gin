gen_rpc = protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/googleapis/googleapis
proto_file = api/demo.proto

rpc:
	$(gen_rpc) --gogo_out=plugins=grpc:. $(proto_file)
	$(gen_rpc) --gin_out=:. $(proto_file)