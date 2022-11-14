package user_repository

import (
	"context"
	"go-flutter-bootcamp/models"
)

type Repository interface {
	Create(ctx context.Context, req *models.CreateUserRequest) error
	Login(ctx context.Context, req *models.LoginRequest) (*models.UserDetailResponse, error)
	Detail(ctx context.Context, userId string) (*models.UserDetailResponse, error)
}
