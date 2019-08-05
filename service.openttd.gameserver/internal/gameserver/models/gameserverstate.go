package models

import (
	u "github.com/ropenttd/tsubasa/service.openttd.gameserver/pkg/utils"
	"github.com/satori/go.uuid"
	"time"
)

type GameserverState struct {
	ServerID       uuid.UUID  `gorm:"primary_key;type:uuid"`
	Server         Gameserver `gorm:"foreignkey:UserRefer"`
	LastSnapshot   uuid.UUID  `json:"snapshot"`
	LastSnapshotAt time.Time  `json:"snapshot_time"`
	GameID         uuid.UUID  `json:"game"`
}

func (GameserverState) TableName() string {
	return "gameservers_state"
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (gameserverState *GameserverState) Validate() (map[string]interface{}, bool) {

	if gameserverState.ServerID == uuid.Nil {
		return u.Message(false, "Server ID should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (gameserverState *GameserverState) Create() map[string]interface{} {

	if resp, ok := gameserverState.Validate(); !ok {
		return resp
	}

	GetDB().Create(gameserverState)

	resp := u.Message(true, "success")
	resp["gameserverState"] = gameserverState
	return resp
}
