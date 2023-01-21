package authcontroller

import (
	"gin-api/helpers"
	"gin-api/models/usermodel"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB: DB}
}

func (ac *AuthController) Register(c *gin.Context) {
	var payload usermodel.Register

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Password not match"})
		return
	}

	passwordHash, err := helpers.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	newUser := usermodel.User{
		Name:      payload.Name,
		Username:  payload.Username,
		Password:  passwordHash,
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := ac.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	c.JSON(http.StatusBadGateway, gin.H{"status": "success", "message": "Success register"})
}
