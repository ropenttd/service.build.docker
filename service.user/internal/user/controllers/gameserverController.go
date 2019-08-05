package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ropenttd/tsubasa/service.openttd.gameserver/internal/gameserver/models"
	u "github.com/ropenttd/tsubasa/service.openttd.gameserver/pkg/utils"
	"github.com/satori/go.uuid"
	"net/http"
)

var CreateGameserver = func(w http.ResponseWriter, r *http.Request) {
	gameserver := &models.Gameserver{}

	err := json.NewDecoder(r.Body).Decode(gameserver)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := gameserver.Create()
	u.Respond(w, resp)
}

var GetGameserverByID = func(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(mux.Vars(r)["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid UUID supplied"))
	}

	data := models.GetGameserverByID(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetGameserverByShortname = func(w http.ResponseWriter, r *http.Request) {
	sn := mux.Vars(r)["shortname"]

	data := models.GetGameserverByShortname(&sn)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

/*var SearchGameserver = func(w http.ResponseWriter, r *http.Request) {
	gameserver := &models.Gameserver{}

	err := json.NewDecoder(r.Body).Decode(gameserver) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := models.SearchGameserver(gameserver)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}*/
