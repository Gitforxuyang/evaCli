package template

import (
	"html/template"
	"os"
	"path"
	"unicode"
)

const (
	makefile string = `
proto:
	-f mkdir ./proto/{{.Name}}
	protoc --eva_out=plugins=all:./proto/{{.Name}} -I=./proto {{.Name}}.proto
	protoc --go_out=plugins=grpc:./proto/{{.Name}} -I=./proto {{.Name}}.proto

.PHONY: proto
`
	gomod string = `
module {{.Name}}

go 1.12

require (
	github.com/Gitforxuyang/eva v0.0.0-20200812113758-270624ede9a9
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
)

`
	main string = `
package main

import (
	"{{.Name}}/handler"
	"{{.Name}}/proto/{{.Name}}"
	"github.com/Gitforxuyang/eva/server"
	"google.golang.org/grpc"
)

func main(){
	server.Init()
	server.RegisterGRpcService(func(server *grpc.Server) {
		{{.Name}}.RegisterSayHelloServiceServer(server,&handler.HandlerService{})
	})
	server.Run()
}
`
)

func Makefile(d Data) {
	f, err := os.Create(path.Join(d.Name, "Makefile"))
	CheckErr(err)
	tmp, err := template.New("test").Parse(makefile)
	CheckErr(err)
	err = tmp.Execute(f, d)
	CheckErr(err)
}

type Data struct {
	Name string
	Port int
}

func GoMod(d Data) {
	f, err := os.Create(path.Join(d.Name, "go.mod"))
	CheckErr(err)
	tmp, err := template.New("test").Parse(gomod)
	CheckErr(err)
	err = tmp.Execute(f, d)
	CheckErr(err)
}

func Main(d Data) {
	f, err := os.Create(path.Join(d.Name, "main.go"))
	CheckErr(err)
	tmp, err := template.New("test").Parse(main)
	CheckErr(err)
	err = tmp.Execute(f, d)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
