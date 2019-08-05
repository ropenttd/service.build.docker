package app

import (
	u "github.com/ropenttd/tsubasa/service.openttd.gameserver/pkg/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "Resource not found"))
		next.ServeHTTP(w, r)
	})
}
