
proto:
	-f mkdir ./proto/app
	protoc --eva_out=plugins=all:./proto/app -I=./proto app.proto
	protoc --go_out=plugins=grpc:./proto/app -I=./proto app.proto

.PHONY: proto
