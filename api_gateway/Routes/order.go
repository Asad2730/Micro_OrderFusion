package routes

import (
	controller "github.com/Asad2730/Micro_OrderFusion/api_gateway/Controller"
	"github.com/Asad2730/Micro_OrderFusion/api_gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.Engine, controller *controller.OrderClient) {
	orderGroup := router.Group("/product")
	{
		orderGroup.Use(middleware.AuthToken())
		orderGroup.POST("/", controller.CreateOrder)
		orderGroup.POST("/item", controller.CreateOrderItem)
		orderGroup.GET("/:id", controller.OrderByID)
		orderGroup.GET("/", controller.OrderList)
	}
}
