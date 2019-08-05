package main

import (
	"github.com/ropenttd/tsubasa/generics/pkg/environment"
	handler "github.com/ropenttd/tsubasa/service.build.docker/internal/service-builder-docker/grpc"
	"log"
	"strconv"
)

func main() {
	// Pull configuration envvars.
	configHostname := environment.GetEnv("DOCKER_HOSTNAME", "unix://var/run/docker.sock")
	configAuthUser := environment.GetEnv("DOCKER_AUTH_USER", "")
	configAuthPass := environment.GetEnv("DOCKER_AUTH_PASS", "")
	configPortS := environment.GetEnv("SERVICE-BUILD-DOCKER_PORT", "80")
	configDebugS := environment.GetEnv("SERVICE-BUILD-DOCKER_DEBUG", "false")

	configPort, err := strconv.Atoi(configPortS)
	if err != nil {
		log.Fatalf("Port %s is not an integer", configPortS)
	}

	configDebug, err := strconv.ParseBool(configDebugS)
	if err != nil {
		log.Fatalf("Debug envvar %s is not parsable as a boolean", configDebugS)
	}

	handler.RunServer(configHostname, configAuthUser, configAuthPass, configPort, configDebug)
}
