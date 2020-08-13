package main

import (
	"flag"
	"fmt"
	"github.com/Gitforxuyang/evaCli/template"
	"os"
	"path"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	tmp string = "evaCli -name=xxx -port=xxx"
)

func main() {
	name := flag.String("name", "", "tt name")
	port := flag.Int("port", 0, "tt port")
	flag.Parse()
	if *name == "" {
		panic(fmt.Sprintf("name不能为空 \n示例：%s", tmp))
	}
	if *port == 0 {
		panic(fmt.Sprintf("port不能为0 \n示例：%s", tmp))
	}
	//TODO:增加name的正则判断。只能创建驼峰命名
	//if isExist(*name) {
	//	panic("期望创建的服务已存在")
	//}
	//创建文件夹
	err := os.MkdirAll(path.Join(*name), 0777)
	if err != nil {
		panic(err)
	}
	os.MkdirAll(path.Join(*name, "app", "service"), 0777)
	os.MkdirAll(path.Join(*name, "app", "assembler"), 0777)
	os.MkdirAll(path.Join(*name, "conf"), 0777)
	os.MkdirAll(path.Join(*name, "handler"), 0777)
	os.MkdirAll(path.Join(*name, "domain", "entity"), 0777)
	os.MkdirAll(path.Join(*name, "domain", "event"), 0777)
	os.MkdirAll(path.Join(*name, "domain", "repo"), 0777)
	os.MkdirAll(path.Join(*name, "domain", "service"), 0777)
	os.MkdirAll(path.Join(*name, "domain", "util"), 0777)
	os.MkdirAll(path.Join(*name, "infra", "err"), 0777)
	os.MkdirAll(path.Join(*name, "infra", "helper"), 0777)
	os.MkdirAll(path.Join(*name, "infra", "repo"), 0777)
	os.MkdirAll(path.Join(*name, "infra", "util"), 0777)
	os.MkdirAll(path.Join(*name, "proto"), 0777)
	d := template.Data{Name: *name, Port: *port}
	template.Makefile(d)
	template.GoMod(d)
	template.Main(d)
	template.Conf(d)
	template.Proto(d)
	template.Handler(d)

}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
