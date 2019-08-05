package versionwatch

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/service.openttd.gameserver/internal/gameserver/controllers"
	"log"
	"net/http"
)

func RunServer(configPort *int) {

	router := mux.NewRouter()

	router.HandleFunc("/api/gameserver", controllers.CreateGameserver).Methods("POST")
	router.HandleFunc("/api/gameserver/find/{shortname}", controllers.GetGameserverByShortname).Methods("GET")
	//router.HandleFunc("/api/gameserver/search", controllers.SearchGameserver).Methods("GET")
	router.HandleFunc("/api/gameserver/{id}", controllers.GetGameserverByID).Methods("GET")
	//router.HandleFunc("/api/gameserver/{id}", controllers.PutGameserver).Methods("PUT")
	//router.HandleFunc("/api/gameserver/{id}", controllers.PatchGameserver).Methods("PATCH")
	//router.HandleFunc("/api/gameserver/{id}", controllers.DeleteGameserver).Methods("DELETE")
	//router.HandleFunc("/api/gameserver/{id}/state", controllers.PutGameserverState).Methods("PUT")
	//router.HandleFunc("/api/gameserver/{id}/state", controllers.PatchGameserverState).Methods("PATCH")

	log.Printf("🚀️ service.openttd.gameserver - ready to serve")
	err := http.ListenAndServe(fmt.Sprintf(":%v", *configPort), router)
	if err != nil {
		fmt.Print(err)
	}
}
