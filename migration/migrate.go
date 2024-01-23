package migration

import (
	"go-gin-tuts/initializers"
	"go-gin-tuts/models"
)

func MigrateTables() {
	initializers.DB.AutoMigrate(&models.Post{})
}
