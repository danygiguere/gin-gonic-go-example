package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type DemoController struct {}

func NewDemoController() DemoController {
	return DemoController{}
}

func inc(x *int) {
	*x++
}

func (controller *DemoController) Index(ctx *gin.Context) {
	a := 8
	fmt.Println("a = ", a)
	// “&a” simply denotes that the system is to provide the memory address of the variable a.
	fmt.Println("&a = ", &a) // 0xc0004000d8
	b := &a
	fmt.Println("b = ", b)
	// “*” tells the system to use the value as a pointer and return whatever is at that address.
	fmt.Println("*b = ", *b) // 8

	i := 7
	inc(&i)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": i})
}