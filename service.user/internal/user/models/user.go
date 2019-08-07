package models

import (
	"github.com/jinzhu/gorm"
	resp "github.com/ropenttd/tsubasa/generics/pkg/responses"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();not null"`
	Username string    `json:"username" gorm:"not null;index"`
	Email    string    `json:"email" gorm:"not null;index"`

	CreatedAt time.Time  `json:"created" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted,omitempty"`
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (user *User) Validate() (map[string]interface{}, bool) {
	if user.Username == "" {
		return resp.Message(false, "Username is required"), false
	}

	if !strings.Contains(user.Email, "@") {
		return resp.Message(false, "Email address is required"), false
	}

	//Email must be unique
	t := &User{}

	//check for errors and duplicates
	err := GetDB().Table("users").Where("email = ? OR username = ?", user.Email, user.Username).First(t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return resp.Message(false, "Connection error. Please retry"), false
	}
	if t.Email == user.Email {
		return resp.Message(false, "Email address already in use by another user."), false
	} else if t.Username == user.Username {
		return resp.Message(false, "Username already in use."), false
	} else {
		return resp.Message(false, "This user is already taken."), false
	}

	//All the required parameters are present
	return resp.Message(true, "success"), true
}

func (user *User) Create() map[string]interface{} {
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	GetDB().Create(user)

	resp := resp.Message(true, "success")
	resp["user"] = user
	return resp
}
