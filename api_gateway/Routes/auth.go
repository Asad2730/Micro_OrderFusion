package routes

import (
	controller "github.com/Asad2730/Micro_OrderFusion/api_gateway/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, controller *controller.AuthClient) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controller.SignUp)
		auth.POST("/login", controller.Login)
	}
}
