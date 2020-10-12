#!/usr/bin/env bash

cd ../
#protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=$GOPATH/src internal/proto-files/domain/repository.proto
#protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. internal/proto-files/service/repository-service.proto
#protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. --go-grpc_out=. ./api/service/user.proto
#protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:. --go-grpc_out=. api/service/user.proto

protoc -I/usr/local/include -I. -I$GOPATH/src --go_out=plugins=grpc:.  api/service/user.proto

protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out ./gen/go --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative api/service/user.proto

protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out . --grpc-gateway_opt logtostderr=true api/service/user.proto
