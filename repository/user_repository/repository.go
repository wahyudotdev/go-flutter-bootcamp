package user_repository

import (
	"context"
	"go-flutter-bootcamp/models"
	"mime/multipart"
)

type Repository interface {
	Create(ctx context.Context, req *models.CreateUserRequest) error
	Login(ctx context.Context, req *models.LoginRequest) (*models.UserDetailResponse, error)
	Detail(ctx context.Context, userId string) (*models.UserDetailResponse, error)
	Update(ctx context.Context, userId string, file multipart.File, req *models.UpdateProfileRequest) (*models.UserDetailResponse, error)
}
