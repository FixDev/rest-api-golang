package controllers

import (
	"go-gin-mod/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func FindBooks(ctx *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBooksById(ctx *gin.Context) {
	var book []models.Book

	err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(ctx *gin.Context) {
	var input CreateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}

	models.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(ctx *gin.Context) {
	var book []models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Update(input)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(ctx *gin.Context) {
	var book []models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
