FROM alpine:latest
WORKDIR /
ADD go-grpc-demo /go-grpc-demo
ENTRYPOINT [ "./go-grpc-demo" ]
