package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ropenttd/tsubasa/generics/pkg/environment"
	"log"
)

var db *gorm.DB

func init() {

	username := environment.GetEnv("db_user", "root")
	password := environment.GetEnv("db_pass", "")
	dbName := environment.GetEnv("db_name", "tsubasa")
	dbHost := environment.GetEnv("db_host", "localhost")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		// We can't talk to the database, that's fatal
		log.Fatal(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Gameserver{}, &GameserverState{})
}

func GetDB() *gorm.DB {
	return db
}
