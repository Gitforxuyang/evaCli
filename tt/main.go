package main

import (
	"evaDemo/handler"
	"evaDemo/proto/hello"
	"github.com/Gitforxuyang/eva/server"
	"google.golang.org/grpc"
)

func main(){
	server.Init()
	server.RegisterGRpcService(func(server *grpc.Server) {
		hello.RegisterSayHelloServiceServer(server,&handler.HandlerService{})
	})
	server.Run()
}
