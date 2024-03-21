package server

import (
	"context"
	"fmt"
	"go-grpc-example/api"
)

// MyHelloServer 服务端  实现接口 HelloServer
type MyHelloServer struct {
}

func (h *MyHelloServer) NoArgs(ctx context.Context, request *api.NoArgsRequest) (*api.NoArgsResponse, error) {
	return &api.NoArgsResponse{}, nil
}

func (h *MyHelloServer) SayHello(ctx context.Context, request *api.HelloRequest) (*api.HelloResponse, error) {
	return &api.HelloResponse{Message: fmt.Sprintf("hello, %s", request.Username)}, nil
}
