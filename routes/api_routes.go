package routes

import (
	"example/web-service-gin/controllers"
	"github.com/gin-gonic/gin"
)

func NewApiRoutes() *gin.Engine {
	router := gin.Default()

	productController := controllers.NewProductController()
	demoController := controllers.NewDemoController()

	apiRouter := router.Group("/api")
	// Route
	apiRouter.GET("/products", productController.Index)
	apiRouter.GET("/products/:id", productController.Show)
	apiRouter.POST("/products", productController.Create)
	apiRouter.PUT("/products/:id", productController.Update)
	apiRouter.DELETE("/products/:id", productController.Delete)

	apiRouter.GET("/demo", demoController.Index)

	return router
}
