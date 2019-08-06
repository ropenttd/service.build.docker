package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/generics/pkg/serviceserve"
	"github.com/ropenttd/tsubasa/service.user.reddit/internal/user.reddit/controllers"
)

func RunServer() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/reddit/auth", controllers.SendRedirect).Methods("GET")
	router.HandleFunc("/api/user/reddit/auth/callback", controllers.ReceiveCallback).Methods("GET")

	err := serviceserve.Serve(router)
	if err != nil {
		fmt.Print(err)
	}
}
