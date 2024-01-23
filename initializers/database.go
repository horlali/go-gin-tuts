package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open("./db.sqlite3"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	return DB, err
}
