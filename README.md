[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/petersantoso94/golang-microservices)

### Protoc generate command:
```
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative *.proto
```

### Run user service:
```
GRPC_ADDR=":9000" ./user
```

To talk to the gRPC service and see how its responses mirror using the core service directly, build the `test` executable:
```
cd test
export GRPC_ADDR=":9000"
go build -o ./test .
# sample commands:
./test 1
./test 1 2 3
./test 5
./test 1 2 5