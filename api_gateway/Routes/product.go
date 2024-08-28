package routes

import (
	controller "github.com/Asad2730/Micro_OrderFusion/api_gateway/Controller"
	"github.com/Asad2730/Micro_OrderFusion/api_gateway/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, controller *controller.ProductClient) {
	productGroup := router.Group("/product")
	{
		productGroup.Use(middleware.AuthToken())
		productGroup.POST("/", controller.CreateProduct)
		productGroup.GET("/", controller.ProductList)
		productGroup.GET("/:id", controller.ProductByID)
		productGroup.PUT("/", controller.UpdateProduct)
		productGroup.DELETE("/:id", controller.DeleteProduct)
	}
}
