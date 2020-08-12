
proto:
	-f mkdir ./proto/hello
	protoc --eva_out=plugins=all:./proto/hello -I=./proto hello.proto
	protoc --go_out=plugins=grpc:./proto/hello -I=./proto hello.proto

.PHONY: proto
