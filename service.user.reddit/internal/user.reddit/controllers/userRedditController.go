package controllers

import (
	"fmt"
	u "github.com/ropenttd/tsubasa/service.openttd.gameserver/pkg/utils"
	"github.com/ropenttd/tsubasa/service.user.reddit/internal/user.reddit/app"
	"github.com/ropenttd/tsubasa/service.user.reddit/internal/user.reddit/models"
	"net/http"
)

var SendRedirect = func(w http.ResponseWriter, r *http.Request) {
	url, err := app.GetRedditOAuthURL()
	if err != nil {
		u.Respond(w, u.Message(false, fmt.Sprint(err)))
	}
	u.RedirectFound(w, r, *url)
}

var ReceiveCallback = func(w http.ResponseWriter, r *http.Request) {
	// Temporary, this should be randomly set / checked.
	if r.URL.Query().Get("state") != "state" {
		u.Respond(w, u.Message(false, "State mismatch (possible CSRF attack, or your session expired)"))
		return
	}

	session, err := app.NewRedditOAuthSession()
	if err != nil {
		u.Respond(w, u.Message(false, fmt.Sprint(err)))
		return
	}
	token, err := session.CodeAuthWithToken(r.URL.Query().Get("code"))
	if err != nil {
		u.Respond(w, u.Message(false, fmt.Sprint(err)))
		return
	}

	currentRedditUser, err := session.Me()
	if err != nil {
		u.Respond(w, u.Message(false, fmt.Sprint(err)))
		return
	}

	// We have a valid redditor, let's add / get them to/from the Redditor database
	dbRedditor := &models.Redditor{
		RedditID:       currentRedditUser.ID,
		RedditUsername: currentRedditUser.Name,
	}
	dbRedditor.Create()

	// and make a note of their AccessToken
	dbRedditor.StoreTokenForRedditor(token)

	// Finally, request the actual currentRedditUser
	// TODO from service.user GET "/api/user?id=dbRedditor.UserID"

	// if serviceUser != nil {
	// Create a new user.
	// TODO from service.user POST "/api/user"
	// dbRedditor.UserID = user.ID
	// dbRedditor.Save()
	// }

	// TODO Log the user in (this will mean getting them a token, setting it in their browser, etc etc standard session stuff)
	// TODO from service.user GET "/api/user/{serviceUser.ID}/token"
	//

	resp := u.Message(true, "success")
	u.Respond(w, resp)
}
