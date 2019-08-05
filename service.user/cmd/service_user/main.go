package main

import (
	"github.com/ropenttd/tsubasa/generics/pkg/environment"
	handler "github.com/ropenttd/tsubasa/service.user/internal/user"
	"log"
	"strconv"
)

func main() {
	configPortS := environment.GetEnv("PORT", "8000")

	configPort, err := strconv.Atoi(configPortS)
	if err != nil {
		log.Fatalf("Port %s is not an integer", configPortS)
	}
	handler.RunServer(&configPort)
}
