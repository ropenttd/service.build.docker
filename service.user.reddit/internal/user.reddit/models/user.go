package models

import (
	r "github.com/ropenttd/tsubasa/generics/pkg/responses"
	"github.com/ropenttd/tsubasa/service.user.reddit/internal/user.reddit/app"
	"github.com/satori/go.uuid"
	"time"
)

type Redditor struct {
	ID     uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();not null"`
	UserID uuid.UUID `json:"user_id" gorm:"index"`

	RedditID       string `json:"reddit_id" gorm:"not null;index"`
	RedditUsername string `json:"reddit_username" gorm:"not null;index"`
	RedditToken    RedditorOauthToken

	CreatedAt time.Time  `json:"created" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted,omitempty"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (redditor *Redditor) Validate() (map[string]interface{}, bool) {
	if redditor.RedditID == "" {
		return r.Message(false, "Reddit ID is required"), false
	}
	if redditor.RedditUsername == "" {
		return r.Message(false, "Reddit Username is required"), false
	}

	//All the required parameters are present
	return r.Message(true, "success"), true
}

func (RedditorIn *Redditor) Create() map[string]interface{} {
	if resp, ok := RedditorIn.Validate(); !ok {
		return resp
	}

	GetDB().Preload("RedditToken").Where(Redditor{RedditID: RedditorIn.RedditID}).FirstOrCreate(&RedditorIn, RedditorIn)
	resp := r.Message(true, "success")
	resp["redditor"] = RedditorIn
	return resp
}

func (RedditorIn *Redditor) OAuthClient() (client *app.ExtendedOAuthSession, err error) {
	// Get the Redditor's token entry.
	var token *RedditorOauthToken
	if GetDB().Model(RedditorIn).Related(&token).Error != nil {
		return nil, err
	}

	return token.OAuthClient()
}
