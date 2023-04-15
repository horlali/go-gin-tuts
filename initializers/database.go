package initializers

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error
	DB, err := gorm.Open(sqlite.Open("gorm.sqlite3"), &gorm.Config{})

	if DB != nil {
		log.Println("database connected")
	}

	if err != nil {
		log.Fatal("failed to connect database")
	}

	return DB, err
}
