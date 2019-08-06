package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/generics/pkg/serviceserve"
	"github.com/ropenttd/tsubasa/service.user/internal/user/controllers"
)

func RunServer() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/login", controllers.GetUser).Methods("GET")
	//router.HandleFunc("/api/user/search", controllers.SearchUser).Methods("GET")
	router.HandleFunc("/api/user", controllers.GetUser).Methods("GET")

	err := serviceserve.Serve(router)
	if err != nil {
		fmt.Print(err)
	}
}
