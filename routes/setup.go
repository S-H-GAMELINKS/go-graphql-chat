package routes

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})
}
