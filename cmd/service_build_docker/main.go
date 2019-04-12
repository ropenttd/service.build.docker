package main

import (
	internal_service_http "github.com/ropenttd/service.build.docker/internal/service-builder-docker/http"
	"github.com/ropenttd/service.build.docker/pkg/builder"
	"github.com/ropenttd/service.build.docker/pkg/utility"
	"log"
	"net/http"
)

func main() {
	// Pull configuration envvars.
	configHostname := utility.GetEnv("DOCKER_HOSTNAME", "unix://")
	configAuth := utility.GetEnv("DOCKER_AUTH", "")
	// Initialize the Docker config.
	var err error
	builder.BuildClientConfig, err = builder.CreatedockerConfig(configHostname, configAuth)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/api/build", internal_service_http.ApiHandler(internal_service_http.BuildHandler))

	log.Printf("🚀️ service.build.docker - ready")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
