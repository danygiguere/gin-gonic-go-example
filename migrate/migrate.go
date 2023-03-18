package migrate

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
)

// to execute this file, run: go run migrate/migrate.go

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Product{})
}
