package models

import (
	u "github.com/ropenttd/tsubasa/service.openttd.gameserver/pkg/utils"
	"github.com/satori/go.uuid"
	"time"
)

type Gameserver struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();not null"`
	NameShort   string    `json:"name_short" gorm:"not null;index"`
	NameLong    string    `json:"name_long"`
	Description string    `json:"description"`

	Hostname    string `json:"hostname" gorm:"index"`
	PortPublic  uint16 `json:"public_port"`
	PortPrivate uint16 `json:"private_port"`

	CreatedAt time.Time  `json:"created" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (gameserver *Gameserver) Validate() (map[string]interface{}, bool) {
	if gameserver.NameShort == "" {
		return u.Message(false, "Short name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (gameserver *Gameserver) Create() map[string]interface{} {
	if resp, ok := gameserver.Validate(); !ok {
		return resp
	}

	GetDB().Create(gameserver)

	resp := u.Message(true, "success")
	resp["gameserver"] = gameserver
	return resp
}

func GetGameserverByID(id uuid.UUID) *Gameserver {
	gameserver := &Gameserver{}
	err := GetDB().Where("id = ?", id).Preload("GameserverState").First(gameserver).Error
	if err != nil {
		return nil
	}
	return gameserver
}

func GetGameserverByShortname(sn *string) *Gameserver {
	gameserver := &Gameserver{}
	err := GetDB().Where("shortname = ?", sn).Preload("GameserverState").Find(gameserver).Error
	if err != nil {
		return nil
	}
	return gameserver
}

func SearchGameserver(query *Gameserver) *Gameserver {
	gameserver := &Gameserver{}
	err := GetDB().Where(query).Preload("GameserverState").Find(gameserver).Error
	if err != nil {
		return nil
	}
	return gameserver
}
