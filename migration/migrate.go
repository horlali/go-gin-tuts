package migration

import (
	"go-gin-tuts/initializers"
	"go-gin-tuts/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
