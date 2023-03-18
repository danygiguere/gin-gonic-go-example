package main

import (
	"example/web-service-gin/config"
	"example/web-service-gin/initializers"
	"example/web-service-gin/routes"
)

func init() {
	initializers.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {

	router := routes.NewApiRoutes()
	router.Run()
}
