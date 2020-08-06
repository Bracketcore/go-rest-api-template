package app

import (
	"github.com/bracketcore/go-rest-api-template/app/models"
	"github.com/bracketcore/go-rest-api-template/app/utils"
	"github.com/jinzhu/gorm"
	"log"
	// using sqlite3 db
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDB for app
func InitDB(config *utils.Config) *gorm.DB {
	db, err := gorm.Open("sqlite3", config.DBName)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	log.Println("Database connection successful")
	return db
}
