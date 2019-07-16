#### install protobuf
```shell
brew install protobuf
```

#### install dependencies
python
```shell
pip3 install grpcio
pip3 install grpcio-tools googleapis-common-protos
```
go
```shell
go get -u -v google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### generate code
```shell
cd hello
```
python
```shell
python3 -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. hello.proto
```
go
```shell
protoc hello.proto --go_out=plugins=grpc:.
```

#### run
```shell
go run server.go
python3 client.py
```