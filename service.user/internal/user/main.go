package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/service.user/internal/user/controllers"
	"log"
	"net/http"
)

func RunServer(configPort *int) {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/login", controllers.GetUser).Methods("GET")
	//router.HandleFunc("/api/user/search", controllers.SearchUser).Methods("GET")
	router.HandleFunc("/api/user", controllers.GetUser).Methods("GET")

	log.Printf("🚀️ service.user - ready to serve")
	err := http.ListenAndServe(fmt.Sprintf(":%v", *configPort), router)
	if err != nil {
		fmt.Print(err)
	}
}
