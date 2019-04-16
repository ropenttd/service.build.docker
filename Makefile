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

protobuf:
	$(MAKE) -C generics protobuf
	$(MAKE) -C service.build.docker protobuf