package main

import (
	"go-graphql-chat/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	routes.SetRouter(r)

	r.Run(":8000")
}