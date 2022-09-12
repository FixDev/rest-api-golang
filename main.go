package main

import (
	"go-gin-mod/controllers"
	"go-gin-mod/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	models.ConnectDatabase()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"say": "Hello All!"})
	})
	server.GET("/books", controllers.FindBooks)
	server.GET("/books/:id", controllers.FindBooksById)
	server.POST("/books", controllers.CreateBook)
	server.PATCH("/books/:id", controllers.UpdateBook)
	server.DELETE("/books/:id", controllers.DeleteBook)

	server.Run()
}
