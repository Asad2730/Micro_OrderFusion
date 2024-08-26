package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Asad2730/Micro_OrderFusion/order/db"
	pb "github.com/Asad2730/Micro_OrderFusion/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	address string
}

func NewServer(address string) *server {
	return &server{address: address}
}

func (s *server) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	new_order := &pb.Order{
		Id:     int32(len(db.Order_db) + 1),
		UserId: request.UserId,
		Total:  request.Total,
		Status: request.Status,
	}

	db.Order_db = append(db.Order_db, new_order)
	response := &pb.CreateOrderResponse{Order: new_order}
	return response, nil
}

func (s *server) CreateOrderItem(ctx context.Context, request *pb.CreateOrderItemRequest) (*pb.CreateOrderItemResponse, error) {
	new_orderItem := &pb.OrderItem{
		Id:        int32(len(db.OrderItem_db) + 1),
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
		Price:     request.Price,
	}

	db.OrderItem_db = append(db.OrderItem_db, new_orderItem)
	response := &pb.CreateOrderItemResponse{OrderItem: new_orderItem}
	return response, nil
}

func (s *server) OrderByID(ctx context.Context, request *pb.OrderByIDRequest) (*pb.OrderByIDResponse, error) {

	for _, order := range db.Order_db {
		if order.Id == request.Id {
			var orderItems []*pb.OrderItem
			for _, item := range db.OrderItem_db {
				if item.OrderId == order.Id {
					orderItems = append(orderItems, item)
				}
			}
			response := &pb.OrderByIDResponse{
				Order:      order,
				OrderItems: orderItems,
			}

			return response, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Order Not found against id: %d\n", request.Id)
}

func (s *server) OrderList(context.Context, *pb.OrderListRequest) (*pb.OrderListResponse, error) {

	var result []*pb.OrderList

	for _, order := range db.Order_db {

		orderList := &pb.OrderList{
			Order:     order,
			OrderItem: []*pb.OrderItem{},
		}
		for _, item := range db.OrderItem_db {
			if item.OrderId == order.Id {
				orderList.OrderItem = append(orderList.OrderItem, item)
			}

			result = append(result, orderList)
		}
	}

	response := &pb.OrderListResponse{OrderList: result}
	return response, nil
}

func (s *server) Start() error {
	listeners, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	gRPC := grpc.NewServer()
	pb.RegisterOrderServiceServer(gRPC, &server{})
	fmt.Println("gRPC server is running at ", s.address)
	return gRPC.Serve(listeners)
}
