cmdGOPATH:=$(shell go env GOPATH)
BIN_PATH:=cmd/main.go
PROJECT:=go-grpc-demo
IMAGE:=helloworldyu/go-grpc-demo:latest

.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ${PROJECT} ${BIN_PATH}


.PHONY: darwin
darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -o ${PROJECT} ${BIN_PATH}

.PHONY: test
test:
	go test -v ./... -cover


.PHONY: proto
proto:
	protoc --go_out=plugins=grpc:. --proto_path=./proto/ ./proto/*.proto

.PHONY: clean
clean:
	rm -rf ${PROJECT}

.PHONY: docker-build
docker-build: linux
	docker build . -t ${IMAGE}

.PHONY: docker-push
docker-push: docker-build
	docker push ${IMAGE}
