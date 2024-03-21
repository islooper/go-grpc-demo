package server

import (
	context "context"
	"fmt"
	"go-grpc-example/api"
)

type MyUserServer struct {
}

func (m *MyUserServer) Register(ctx context.Context, in *api.RegisterIn) (*api.RegisterOut, error) {
	fmt.Printf("Register: %v", in)
	return &api.RegisterOut{Code: 200, Msg: "register ok"}, nil
}
