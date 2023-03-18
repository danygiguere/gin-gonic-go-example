package main

import (
	"example/web-service-gin/configs"
	"example/web-service-gin/initializers"
	"example/web-service-gin/routes"
)

func init() {
	initializers.LoadEnvVariables()
	configs.ConnectToDB()
}

func main() {

	router := routes.NewApiRoutes()
	router.Run()
}
