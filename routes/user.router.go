package routes

import (
	"go_project/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/signup", func(c *gin.Context) {
		controllers.Signup(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	return r
}
