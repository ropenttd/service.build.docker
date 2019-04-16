package grpc

import (
	"fmt"
	"github.com/ropenttd/tsubasa/service.build.docker/api/protobuf"
	"github.com/ropenttd/tsubasa/service.build.docker/pkg/builder"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunServer(configHostname string, configAuthUser string, configAuthPass string, configPort int) {
	// Initialize the Docker config.
	var err error
	builder.BuildClientConfig, err = builder.CreatedockerConfig(configHostname, configAuthUser, configAuthPass)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	v1.RegisterDockerBuildServer(grpcServer, &BuildDockerHandler{})
	// ... // determine whether to use TLS

	// Serve
	log.Printf("🚀️ service.build.docker - ready to serve GRPC")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
