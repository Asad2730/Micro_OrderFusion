package main

import (
	"log"

	authcontroller "github.com/Asad2730/Micro_OrderFusion/api_gateway/authController"
	authroutes "github.com/Asad2730/Micro_OrderFusion/api_gateway/authRoutes"
	order "github.com/Asad2730/Micro_OrderFusion/proto/order"
	product "github.com/Asad2730/Micro_OrderFusion/proto/product"
	user "github.com/Asad2730/Micro_OrderFusion/proto/user"
	"github.com/rs/cors/wrapper/gin"
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
	orderClient := order.NewOrderServiceClient(conn)
	productClient := product.NewProductServiceClient(conn)

	authController := authcontroller.NewAuthClient(userClient)

	authroutes.RegisterAuthRoutes(r, authController)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
