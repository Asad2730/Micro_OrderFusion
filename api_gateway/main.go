package main

import (
	"log"

	controller "github.com/Asad2730/Micro_OrderFusion/api_gateway/Controller"
	routes "github.com/Asad2730/Micro_OrderFusion/api_gateway/Routes"
	order "github.com/Asad2730/Micro_OrderFusion/proto/order"
	product "github.com/Asad2730/Micro_OrderFusion/proto/product"
	user "github.com/Asad2730/Micro_OrderFusion/proto/user"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	r := gin.Default()

	userClient := user.NewUserServiceClient(conn)
	productClient := product.NewProductServiceClient(conn)
	orderClient := order.NewOrderServiceClient(conn)

	authController := controller.NewAuthClient(userClient)
	productController := controller.NewProductClient(productClient)
	orderController := controller.NewOrderClient(orderClient)

	routes.RegisterAuthRoutes(r, authController)
	routes.RegisterProductRoutes(r, productController)
	routes.RegisterOrderRoutes(r, orderController)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
