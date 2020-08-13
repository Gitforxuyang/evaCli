
package handler

import (
	"context"
	"app/proto/app"
)

type HandlerService struct {

}

func (m *HandlerService) Ping(context.Context, *app.Nil) (*app.Nil, error) {
	return &app.Nil{}, nil
}

