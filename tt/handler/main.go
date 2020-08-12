package handler

import (
	"context"
	"evaDemo/proto/hello"
)

type HandlerService struct {

}

func (m *HandlerService) Hello(context.Context, *hello.String) (*hello.String, error) {
	return &hello.String{}, nil
}

