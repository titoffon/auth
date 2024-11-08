package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	desc "github.com/titoffon/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceServer struct {
    desc.UnimplementedUserServiceServer
    db *pgxpool.Pool
}

func NewUserServiceServer(db *pgxpool.Pool) *UserServiceServer {
    return &UserServiceServer{
        db: db,
    }
}

// Реализация методов Create, Get, Update, Delete с использованием db
func (s *UserServiceServer) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
    // Логика сохранения пользователя в БД
    var userID int64
    query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
    err := s.db.QueryRow(ctx, query, req.Name, req.Email, req.Password, req.Role.String()).Scan(&userID)
    if err != nil { 
        return nil, err
    }

    return &desc.CreateUserResponse{Id: userID}, nil
}

func (s *UserServiceServer) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
    // Логика получения пользователя из БД
    var (
        id        int64
        name      string
        email     string
        role      string
		createdAt time.Time
		updatedAt sql.NullTime
    )

    query := `SELECT id, name, email, role, created_at, updated_at FROM users WHERE id = $1`
    err := s.db.QueryRow(ctx, query, req.Id).Scan(&id, &name, &email, &role, &createdAt, &updatedAt)
    if err != nil {
        return nil, err
    }

	response := &desc.GetUserResponse{
		Id:        id,
		Name:      name,
		Email:     email,
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(createdAt),
	}

	if updatedAt.Valid {
		response.UpdatedAt = timestamppb.New(updatedAt.Time)
	}

	return response, nil
}

func (s *UserServiceServer) Update(ctx context.Context, req *desc.UpdateUserRequest) (*emptypb.Empty, error) {
    // Логика обновления пользователя в БД
    query := `UPDATE users SET `
    args := []interface{}{}
    argID := 1

    if req.Name != nil {
        query += fmt.Sprintf("name = $%d, ", argID)
        args = append(args, req.Name.Value)
        argID++
    }
    if req.Email != nil {
        query += fmt.Sprintf("email = $%d, ", argID)
        args = append(args, req.Email.Value)
        argID++
    }

    query += fmt.Sprintf("updated_at = NOW() WHERE id = $%d", argID)
    args = append(args, req.Id)

    _, err := s.db.Exec(ctx, query, args...)
    if err != nil {
        return nil, err
    }

    return &emptypb.Empty{}, nil
}

func (s *UserServiceServer) Delete(ctx context.Context, req *desc.DeleteUserRequest) (*emptypb.Empty, error) {
    // Логика удаления пользователя из БД
    query := `DELETE FROM users WHERE id = $1`
    _, err := s.db.Exec(ctx, query, req.Id)
    if err != nil {
        return nil, err
    }

    return &emptypb.Empty{}, nil
}
