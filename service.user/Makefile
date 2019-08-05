# Go parameters
GOCMD=go
DEPCMD=dep
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODEP=$(DEPCMD) ensure
BINARY_NAME=server
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/service_openttd_gameserver/main.go
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GODEP)
protobuf:
	cd api/protobuf && \
	PATH=${PATH}:${GOPATH}/bin protoc --proto_path=${GOPATH}/src --proto_path=${GOPATH}/src/github.com/google/protobuf/src --proto_path=. --go_out=plugins=grpc:. *.proto
