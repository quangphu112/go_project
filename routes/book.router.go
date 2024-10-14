package routes

import (
	"go_project/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/books", func(c *gin.Context) {
		controllers.GetBooks(c, db)
	})
	r.GET("/book/:id", func(c *gin.Context) {
		controllers.GetBookByID(c, db)
	})
	r.POST("/book", func(c *gin.Context) {
		controllers.CreateBook(c, db)
	})
	r.PUT("/book/:id", func(c *gin.Context) {
		controllers.UpdateBook(c, db)
	})
	r.DELETE("/book/:id", func(c *gin.Context) {
		controllers.DeleteBook(c, db)
	})
	return r
}
