package versionwatch

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/generics/pkg/serviceserve"
	"github.com/ropenttd/tsubasa/service.openttd.versionwatch/internal/versionwatch/controllers"
)

func RunServer() {
	router := mux.NewRouter()

	router.HandleFunc("/api/openttd/version", controllers.GetVersions).Methods("GET")

	err := serviceserve.Serve(router)
	if err != nil {
		fmt.Print(err)
	}
}
