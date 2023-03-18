package controllers

import (
	"net/http"

	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"example/web-service-gin/request"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func NewProductController() ProductController {
	return ProductController{}
}

func (pc *ProductController) Create(ctx *gin.Context) {
	validatedRequest, err := request.ValidateProductCreateRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"status": "errors", "data": validatedRequest})
		return
	}

	product := models.Product{Code: validatedRequest.Code, Price: validatedRequest.Price}
	config.DB.Create(&product)
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "product": product})
}

func (pc *ProductController) Index(ctx *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "products": products})
}

func (pc *ProductController) Show(ctx *gin.Context) {
	var product models.Product
	config.DB.First(&product, ctx.Param("id"))
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "product": product})
}

func (pc *ProductController) Update(ctx *gin.Context) {
	var body struct {
		Code  string
		Price string
	}

	ctx.Bind(&body)
	var product models.Product
	config.DB.First(&product, ctx.Param("id"))

	updatedProduct := models.Product{Code: body.Code, Price: body.Price}
	config.DB.Model(&product).Updates(updatedProduct)
	ctx.Status(http.StatusNoContent)
}

func (pc *ProductController) Delete(ctx *gin.Context) {
	config.DB.Delete(&models.Product{}, ctx.Param("id"))
	ctx.Status(http.StatusNoContent)
}
