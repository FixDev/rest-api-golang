package main

import (
	"net/http"

	"github.com/fixdev/rest-api-golang/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"say": "Hello All!"})
	})

	models.ConnectDatabase()

	server.Run()
}
