package controllers

import (
	"github.com/gin-gonic/gin"
)

type helloController struct {
}

type HelloController interface {
	Index(ctx *gin.Context)
}

func NewHelloController() HelloController {
	return &helloController{}
}

func (helloController *helloController) Index(ctx *gin.Context) {
	ctx.JSON(200, "Hello World!")
}
