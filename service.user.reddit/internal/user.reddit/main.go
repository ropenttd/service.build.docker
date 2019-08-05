package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/service.user.reddit/internal/user.reddit/controllers"
	"log"
	"net/http"
)

func RunServer(configPort *int) {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/reddit/auth", controllers.SendRedirect).Methods("GET")
	router.HandleFunc("/api/user/reddit/auth/callback", controllers.ReceiveCallback).Methods("GET")

	log.Printf("🚀️ service.user.reddit - ready to serve")
	err := http.ListenAndServe(fmt.Sprintf(":%v", *configPort), router)
	if err != nil {
		fmt.Print(err)
	}
}
