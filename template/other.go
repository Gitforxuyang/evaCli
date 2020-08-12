package template

import (
	"html/template"
	"os"
	"path"
)

const (
	makefile string = `
proto:
	-f mkdir ./proto/hello
	protoc --eva_out=plugins=all:./proto/hello -I=./proto hello.proto
	protoc --go_out=plugins=grpc:./proto/hello -I=./proto hello.proto

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
)

func Makefile(app string) error {
	f, err := os.Create(path.Join(app, "Makefile"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(makefile)
	return nil
}

type Data struct {
	Name string
}

func GoMod(app string) {
	f, err := os.Create(path.Join(app, "go.mod"))
	CheckErr(err)
	tmp, err := template.New("test").Parse(gomod)
	CheckErr(err)
	err = tmp.Execute(f, Data{Name: app})
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
