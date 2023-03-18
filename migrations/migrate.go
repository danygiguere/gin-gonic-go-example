package migrations

import (
	"example/web-service-gin/configs"
	"example/web-service-gin/models"
)

// to execute this file, run: go run migrations/migrations.go

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Product{})
}
