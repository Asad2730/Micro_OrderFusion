package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Asad2730/Micro_OrderFusion/product/db"
	pb "github.com/Asad2730/Micro_OrderFusion/proto/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedProductServiceServer
	address string
}

func NewServer(address string) *server {
	return &server{address: address}
}

func (s *server) CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	new_product := &pb.Product{
		Id:          int32(len(db.Product_db) + 1),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		StockQty:    request.StockQty,
	}

	db.Product_db = append(db.Product_db, new_product)
	response := &pb.CreateProductResponse{Message: "Product Added Succesfully!"}
	return response, nil
}

func (s *server) ProductList(ctx context.Context, request *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	response := &pb.ListProductResponse{
		Products: db.Product_db,
	}
	return response, nil
}

func (s *server) ProductByID(ctx context.Context, request *pb.RequestProductID) (*pb.SingleProductResponse, error) {

	for _, item := range db.Product_db {
		if item.Id == request.Id {
			response := &pb.SingleProductResponse{Product: item}
			return response, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Product Not found against id: %d\n", request.Id)

}

func (s *server) UpdateProduct(ctx context.Context, request *pb.RequestProductUpdate) (*pb.SingleProductResponse, error) {
	for index, item := range db.Product_db {
		if item.Id == request.Id {
			db.Product_db[index] = &pb.Product{
				Id:          item.Id,
				Name:        item.Name,
				Description: item.Description,
				Price:       item.Price,
				StockQty:    item.StockQty,
			}
			response := &pb.SingleProductResponse{Product: db.Product_db[index]}
			return response, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Product Not found id: %d\n", request.Id)

}

func (s *server) DeleteProduct(ctx context.Context, request *pb.RequestProductID) (*pb.DeleteProductResponse, error) {

	for index, item := range db.Product_db {
		if item.Id == request.Id {
			db.Product_db = append(db.Product_db[:index], db.Product_db[index+1:]...)
			response := &pb.DeleteProductResponse{Message: "Product Deleted Successfully!"}
			return response, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Product Not found id: %d\n", request.Id)

}

func (s *server) Start() error {
	listeners, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	gRPC := grpc.NewServer()
	pb.RegisterProductServiceServer(gRPC, &server{})
	fmt.Println("gRPC server is running at ", s.address)
	return gRPC.Serve(listeners)
}
