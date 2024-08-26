package main

import (
	"log"

	"github.com/Asad2730/Micro_OrderFusion/user/service"
)

func main() {
	gRpcService := service.NewServer(":8000")
	if err := gRpcService.Start(); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
