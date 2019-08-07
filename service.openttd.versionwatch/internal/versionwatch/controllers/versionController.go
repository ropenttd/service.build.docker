package controllers

import (
	"github.com/gorilla/schema"
	"github.com/ropenttd/tsubasa/generics/pkg/responses"
	"github.com/ropenttd/tsubasa/service.openttd.versionwatch/internal/versionwatch/models"
	"net/http"
)

// GetVersions returns a list of versions, optionally filtered by the given query parameters.
// If you supply ?latest, the query will be limited to the newest version of each build train. (Filters will be disregarded)
var GetVersions = func(w http.ResponseWriter, r *http.Request) {
	// Decode the query into a new object, if present.
	queryParams := r.URL.Query()
	var decoder = schema.NewDecoder()
	var q = models.OpenttdGameVersion{}
	// decoder.Decode returns errors if things don't match the target, but we don't mind about those.
	decoder.Decode(&q, queryParams)

	var data []*models.OpenttdGameVersion
	if _, ok := queryParams["latest"]; ok {
		data = models.ListLatestVersions()
	} else {
		data = models.ListVersions(&q)
	}
	resp := responses.Message(true, "success")
	resp["data"] = data
	responses.Respond(w, resp)
}
