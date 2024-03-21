package main

import (
	"fmt"
	"go-grpc-example/api"
	"go-grpc-example/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func StartServer() (stopServer func()) {
	lis, err := net.Listen("tcp", ":6002")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	api.RegisterHelloServer(s, &server.MyHelloServer{})
	api.RegisterUserServer(s, &server.MyUserServer{})
	// 开启反射服务
	reflection.Register(s)

	go func() {
		err = s.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	stopServer = s.Stop

	return stopServer
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	fmt.Println("startup...")
	stopServer := StartServer()
	s := <-c
	stopServer()
	fmt.Println("end...", s)
}
