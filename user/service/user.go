package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Asad2730/Micro_OrderFusion/common"
	pb "github.com/Asad2730/Micro_OrderFusion/proto"
	"github.com/Asad2730/Micro_OrderFusion/user/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	address string
}

func NewServer(address string) *server {
	return &server{address: address}
}

func (s *server) Login(ctx context.Context, request *pb.RequestLogin) (*pb.LoginResponse, error) {
	var logged_user *pb.User

	for _, user := range db.User_db {
		if user.Email == request.Email && user.Password == request.Password {
			logged_user = user
			break
		}
	}

	if logged_user == nil {
		return nil, status.Errorf(codes.NotFound, "User not found with Email : %s\n", request.Email)
	}

	token, err := common.GenerateJWT(logged_user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	response := &pb.LoginResponse{
		Id:    logged_user.Id,
		Name:  logged_user.Name,
		Email: logged_user.Email,
		Token: token,
	}

	return response, nil
}

func (s *server) SignUp(ctx context.Context, request *pb.RequestSignup) (*pb.SignupResponse, error) {
	new_user := &pb.User{
		Id:       int32(len(db.User_db) + 1),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	db.User_db = append(db.User_db, new_user)
	response := &pb.SignupResponse{Message: "User created Successfully!"}
	return response, nil
}

func (s *server) Start() error {
	listeners, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	gRPC := grpc.NewServer()
	pb.RegisterUserServiceServer(gRPC, &server{})
	fmt.Println("gRPC server is running at ", s.address)
	return gRPC.Serve(listeners)
}
