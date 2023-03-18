package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ProductCreateRequest struct {
	Code  string
	Price string
}

func init() {
	log.Info().Msg("validating ProductCreateRequest")
}

func ValidateProductCreateRequest(ctx *gin.Context) (ProductCreateRequest, ValidationErrors) {
	productCreateRequest := ProductCreateRequest{}
	//locale := ctx.GetHeader("Accept-Language")
	ctx.Bind(&productCreateRequest)
	fields := map[string]string{
		"Code":  productCreateRequest.Code,
		"Price": productCreateRequest.Price,
	}

	ve := make(ValidationErrors)

	code := fields["Code"]
	if len(code) < 6 || len(code) > 25 {
		ve["Code"] = append(ve["Code"], "The field Title must be between 6 and 25 characters long")
	}

	price := fields["Price"]
	if len(price) < 3 || len(price) > 25 {
		ve["Price"] = append(ve["Price"], "The field Description must be between 6 and 25 characters long")
	}

	if len(ve) > 0 {
		return productCreateRequest, ve
	}

	return productCreateRequest, nil
}
