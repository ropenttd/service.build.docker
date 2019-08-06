package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	resp "github.com/ropenttd/tsubasa/generics/pkg/responses"
	"github.com/ropenttd/tsubasa/service.user/internal/user/models"
	"github.com/satori/go.uuid"
	"net/http"
)

var CreateGameserver = func(w http.ResponseWriter, r *http.Request) {
	gameserver := &models.Gameserver{}

	err := json.NewDecoder(r.Body).Decode(gameserver)
	if err != nil {
		resp.Respond(w, resp.Message(false, "Error while decoding request body"))
		return
	}

	response := gameserver.Create()
	resp.Respond(w, response)
}

var GetGameserverByID = func(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(mux.Vars(r)["id"])

	if err != nil {
		resp.Respond(w, resp.Message(false, "Invalid UUID supplied"))
	}

	data := models.GetGameserverByID(id)
	response := resp.Message(true, "success")
	response["data"] = data
	resp.Respond(w, response)
}

var GetGameserverByShortname = func(w http.ResponseWriter, r *http.Request) {
	sn := mux.Vars(r)["shortname"]

	data := models.GetGameserverByShortname(&sn)
	response := resp.Message(true, "success")
	response["data"] = data
	resp.Respond(w, response)
}

/*var SearchGameserver = func(w http.ResponseWriter, r *http.Request) {
	gameserver := &models.Gameserver{}

	err := json.NewDecoder(r.Body).Decode(gameserver) //decode the request body into struct and failed if any error occur
	if err != nil {
		resp.Respond(w, resp.Message(false, "Invalid request"))
		return
	}

	data := models.SearchGameserver(gameserver)
	response := resp.Message(true, "success")
	response["data"] = data
	resp.Respond(w, response)
}*/
