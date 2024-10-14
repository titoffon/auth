package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/titoffon/auth/pkg/auth_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserServiceServer
}

// Get ...
func (s *server) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &desc.GetUserResponse{
	    Id:       req.GetId(),
        Name:     gofakeit.Name(),
        Email:    gofakeit.Email(),
        Role:     desc.Role_USER,
        CreatedAt: timestamppb.New(gofakeit.Date()),
        UpdatedAt: timestamppb.New(gofakeit.Date()),
    }, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // создаём объект нового сервера
	reflection.Register(s) // включаем возможность сервера выдавать информацию о себе 
	desc.RegisterUserServiceServer(s, &server{}) //второй параметр это структура, которая имплементировала API

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil { // запускаем сервер
		log.Fatalf("failed to serve: %v", err)
	}
}
