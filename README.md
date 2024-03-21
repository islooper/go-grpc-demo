# go-grpc-demo
A demo of gRPC in Go. Turning on the reflection of gRPC.

```go
s := grpc.NewServer()
api.RegisterHelloServer(s, &server.MyHelloServer{})
api.RegisterUserServer(s, &server.MyUserServer{})
```