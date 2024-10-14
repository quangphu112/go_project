package controllers

import (
	"go_project/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context, db *gorm.DB) {
	var body models.User
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request !"})
		return
	}

	var exitingUser models.User
	if err := db.Where("email = ?", body.Email).First(&exitingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists !"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": " Failed to hashing password !"})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	db.Create(&user)
	
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": db.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context, db *gorm.DB) {
	var body models.User
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request !"})
		return
	}

	var user models.User
	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found !"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password !"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token !"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"data": tokenString})
}

func Validate (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Valid token !"})
}