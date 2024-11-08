package main

import (
	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/titoffon/auth/internal/config"
	"github.com/titoffon/auth/internal/server"
	"github.com/titoffon/auth/internal/storage"
	desc "github.com/titoffon/auth/pkg/auth_v1"
)

const grpcPort = 50052

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: %v", err)
	}

	ctx := context.Background()

	//Инициадизируем подключение к БД
	db, err := storage.NewDB(ctx, cfg)
	if err != nil{
		log.Fatal("failed to connect to database: %v", err)
	}

	// Создаем экземпляр сервера с подключением к БД
    srv := server.NewUserServiceServer(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()                        // создаём объект нового сервера
	reflection.Register(s)                       // включаем возможность сервера выдавать информацию о себе	
	desc.RegisterUserServiceServer(s, srv) //второй параметр это структура, которая имплементировала API

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil { // запускаем сервер
		log.Fatalf("failed to serve: %v", err)
	}
}
