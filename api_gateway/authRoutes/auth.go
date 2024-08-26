package authroutes

import (
	authcontroller "github.com/Asad2730/Micro_OrderFusion/api_gateway/authController"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, controller *authcontroller.AuthClient) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controller.SignUp)
		auth.POST("/login", controller.Login)
	}
}
