# microservice

Test Bed

Add/Multiple Microservice CRUD 

TechStack

0. how to entity -> a. swagger / b. Go codes(protobuf/grpc) - golang template - ytd
1. Go language - done
2. GRPC - done
3. REST interface - mux / Gin - done
 
3.1 Mock unit test & Integration test - ytd 

4. Swagger - done
5. Docker - ytd
6. Kubernetes - ytd




Command Run:


https://github.com/grpc-ecosystem/grpc-gateway

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

protoc --proto_path=proto -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:proto service.proto

protoc --proto_path=proto -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:proto service.proto

protoc --proto_path=proto  -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:proto service.proto

protoc --proto_path=proto --go_out=plugins=grpc:proto service.proto
