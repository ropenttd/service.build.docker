package models

import (
	r "github.com/ropenttd/tsubasa/generics/pkg/responses"
	"github.com/ropenttd/tsubasa/service.user.reddit/internal/user.reddit/app"
	"github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	"time"
)

type RedditorOauthToken struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();not null"`
	RedditorID uuid.UUID `json:"redditor_id" gorm:"index"`

	oauth2.Token

	CreatedAt time.Time  `json:"created" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (RedditOauthToken *RedditorOauthToken) Validate() (map[string]interface{}, bool) {
	if RedditOauthToken.RedditorID == uuid.Nil {
		return r.Message(false, "Redditor must be supplied"), false
	}

	if RedditOauthToken.AccessToken == "" {
		return r.Message(false, "Access token was not supplied"), false
	}

	//All the required parameters are present
	return r.Message(true, "success"), true
}

func (RedditorIn *Redditor) StoreTokenForRedditor(token *oauth2.Token) map[string]interface{} {

	result := RedditorOauthToken{}

	GetDB().Where(RedditorOauthToken{RedditorID: RedditorIn.ID}).FirstOrInit(&result, RedditorOauthToken{RedditorID: RedditorIn.ID})

	result.Token = *token
	GetDB().Save(&result)

	resp := r.Message(true, "success")
	resp["redditor"] = result
	return resp
}

func (Token *RedditorOauthToken) OAuthClient() (client *app.ExtendedOAuthSession, err error) {
	var oauthToken *oauth2.Token
	if !Token.Valid() {
		// The token is expired, override
		oauthToken = &oauth2.Token{RefreshToken: Token.RefreshToken}
	} else {
		oauthToken = &Token.Token
	}
	// New session.
	session, err := app.NewRedditOAuthSession()
	if err != nil {
		return nil, err
	}
	session.Client = session.OAuthConfig.Client(*session.GetCtx(), oauthToken)
	return session, nil
}
