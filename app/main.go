
package main

import (
	"app/handler"
	"app/proto/app"
	"github.com/Gitforxuyang/eva/server"
	"google.golang.org/grpc"
)

func main(){
	server.Init()
	server.RegisterGRpcService(func(server *grpc.Server) {
		app.RegisterSayHelloServiceServer(server,&handler.HandlerService{})
	})
	server.Run()
}
