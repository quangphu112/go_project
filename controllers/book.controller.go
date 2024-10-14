package controllers

import (
	"go_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context, db *gorm.DB) {
	var book []models.Book
	db.Find(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBookByID(c *gin.Context, db *gorm.DB) {
	var book models.Book
	id := c.Param("id")
	err := db.First(&book, id).Error
	 if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context, db *gorm.DB) {
	var book models.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context, db *gorm.DB) {
	var book models.Book
	id := c.Param("id")
	err := db.First(&book, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}
	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context, db *gorm.DB) {
	var book models.Book
	id := c.Param("id")
	err := db.First(&book, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": "Book deleted successfully"})
}
