package routes

import (
	"gin-api/controllers/authcontroller"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authController authcontroller.AuthController
}

func NewAuthRoute(authcontroller authcontroller.AuthController) AuthRoute {
	return AuthRoute{
		authController: authcontroller,
	}
}

func (rc *AuthRoute) Route(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", rc.authController.Register)
}
