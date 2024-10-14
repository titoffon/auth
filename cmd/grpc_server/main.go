package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/titoffon/auth/pkg/auth_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserServiceServer
}

// Create ...
func (s *server) Create(_ context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	fmt.Printf("Received Create request: Name=%s, Email=%s\n", req.Name, req.Email)
	log.Printf("Received Create request: Name=%s, Email=%s", req.Name, req.Email)
	return &desc.CreateUserResponse{Id: 1}, nil
}

// Get ...
func (s *server) Get(_ context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &desc.GetUserResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

// Update ...
func (s *server) Update(_ context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
	fmt.Printf("Received Update request: ID=%d, Name=%v, Email=%v, Role=%v\n", req.Id, req.Name, req.Email, req.Role)
	log.Printf("Received Update request: ID=%d, Name=%v, Email=%v, Role=%v", req.Id, req.Name, req.Email, req.Role)

	//логика обновления пользователя в базе данных

	return &emptypb.Empty{}, nil
}

// Delete ...
func (s *server) Delete(_ context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
	fmt.Printf("Received Delete request: ID=%d\n", req.Id)
	log.Printf("Received Update request: ID=%d", req.Id)

	// логика удаления пользователя из базы данных

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                        // создаём объект нового сервера
	reflection.Register(s)                       // включаем возможность сервера выдавать информацию о себе
	desc.RegisterUserServiceServer(s, &server{}) //второй параметр это структура, которая имплементировала API

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil { // запускаем сервер
		log.Fatalf("failed to serve: %v", err)
	}
}
